package parser

import (
	"fmt"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/cleanup"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

type resourceIdParseResult struct {
	// nameToResourceIDs is a map[name]ParsedResourceID containing information about the Resource ID's
	nameToResourceIDs map[string]models.ParsedResourceId

	// nestedResult contains any constants and models found when parsing these ID's
	nestedResult parseResult

	// resourceUriMetadata is a map[uri]resourceUriMetadata for the Resource ID reference & any suffix
	resourceUrisToMetadata map[string]resourceUriMetadata
}

type resourceUriMetadata struct {
	// resourceIdName is the name of the ResourceID object, available once the unique names have been
	// identified (if there's a Resource ID)
	resourceIdName *string

	// resourceId is the parsed Resource ID object, if any
	resourceId *models.ParsedResourceId

	// uriSuffix is any suffix which should be applied to the URI
	uriSuffix *string
}

func (d *SwaggerDefinition) findResourceIdsForTag(tag *string) (*resourceIdParseResult, error) {
	result := resourceIdParseResult{
		nestedResult: parseResult{
			constants: map[string]models.ConstantDetails{},
		},

		nameToResourceIDs:      map[string]models.ParsedResourceId{},
		resourceUrisToMetadata: map[string]resourceUriMetadata{},
	}

	// first get a list of all of the Resource ID's present in these operations
	// where a Suffix is present on a Resource ID, we'll have 2 entries for the Suffix and the Resource ID directly
	urisToMetadata, nestedResult, err := d.parseResourceIdsFromOperations(tag)
	if err != nil {
		return nil, fmt.Errorf("parsing Resource ID's from Operations: %+v", err)
	}
	result.nestedResult = *nestedResult

	// next determine names for these
	namesToResourceUris, urisToNames, err := determineNamesForResourceIds(*urisToMetadata)
	if err != nil {
		return nil, fmt.Errorf("determining names for Resource ID's: %+v", err)
	}
	result.nameToResourceIDs = *namesToResourceUris

	// finally go over the existing results and swap out the Resource ID objects for the Name which should be used
	urisToMetadata, err = mapNamesToResourceIds(*urisToNames, *urisToMetadata)
	if err != nil {
		return nil, fmt.Errorf("mapping names back to Resource ID's: %+v", err)
	}
	result.resourceUrisToMetadata = *urisToMetadata

	return &result, nil
}

func (d SwaggerDefinition) parseResourceIdsFromOperations(tag *string) (*map[string]resourceUriMetadata, *parseResult, error) {
	result := parseResult{
		constants: map[string]models.ConstantDetails{},
	}
	urisToMetaData := make(map[string]resourceUriMetadata, 0)

	for _, operation := range d.swaggerSpecExpanded.Operations() {
		for uri, operationDetails := range operation {
			if !operationMatchesTag(operationDetails, tag) {
				continue
			}

			if operationShouldBeIgnored(uri) {
				continue
			}

			metadata, err := d.parseResourceIdFromOperation(uri, operationDetails)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing %q: %+v", uri, err)
			}

			// next, if it's based on a Resource ID, let's ensure that's added too
			resourceUri := uri
			if metadata.resourceId != nil {
				result.appendConstants(metadata.resourceId.Constants)

				resourceManagerUri := metadata.resourceId.NormalizedResourceManagerResourceId()
				if resourceUri != resourceManagerUri {
					urisToMetaData[resourceManagerUri] = resourceUriMetadata{
						resourceIdName: metadata.resourceIdName,
						resourceId:     metadata.resourceId,
						uriSuffix:      nil,
					}
				}
			}
			urisToMetaData[resourceUri] = *metadata
		}
	}

	return &urisToMetaData, &result, nil
}

func (d *SwaggerDefinition) parseResourceIdFromOperation(uri string, operationDetails *spec.Operation) (*resourceUriMetadata, error) {
	// TODO: unit tests for this method too
	segments := make([]models.ResourceIdSegment, 0)
	result := parseResult{
		constants: map[string]models.ConstantDetails{},
	}

	uriSegments := strings.Split(strings.TrimPrefix(uri, "/"), "/")
	for _, uriSegment := range uriSegments {
		normalizedSegment := cleanup.NormalizeSegment(uriSegment, true)
		normalizedSegment = strings.TrimSuffix(strings.TrimPrefix(normalizedSegment, "{"), "}")

		// intentionally check the pre-cut version
		if strings.HasPrefix(uriSegment, "{") && strings.HasSuffix(uriSegment, "}") {
			if strings.EqualFold(normalizedSegment, "scope") {
				segments = append(segments, models.ResourceIdSegment{
					Type: models.ScopeSegment,
					Name: normalizedSegment,
				})
				continue
			}

			if strings.EqualFold(normalizedSegment, "subscriptionId") {
				segments = append(segments, models.ResourceIdSegment{
					Type: models.SubscriptionIdSegment,
					Name: normalizedSegment,
				})
				continue
			}

			if strings.EqualFold(normalizedSegment, "resourceGroupName") {
				segments = append(segments, models.ResourceIdSegment{
					Type: models.ResourceGroupSegment,
					Name: normalizedSegment,
				})
				continue
			}

			isConstant := false
			for _, param := range operationDetails.Parameters {
				if strings.EqualFold(param.Name, normalizedSegment) && strings.EqualFold(param.In, "path") {
					if param.Ref.String() != "" {
						return nil, fmt.Errorf("TODO: Enum's aren't supported by Reference right now, but apparently should be for %q", uriSegment)
					}

					if param.Enum != nil {
						// then find the constant itself
						constant, err := mapConstant([]string{param.Type}, param.Enum, param.Extensions)
						if err != nil {
							return nil, fmt.Errorf("parsing constant from %q: %+v", uriSegment, err)
						}
						result.constants[constant.name] = constant.details
						segments = append(segments, models.ResourceIdSegment{
							Type:              models.ConstantSegment,
							Name:              normalizedSegment,
							ConstantReference: &constant.name,
						})
						isConstant = true
						break
					}
				}
			}
			if isConstant {
				continue
			}

			segments = append(segments, models.ResourceIdSegment{
				Type: models.UserSpecifiedSegment,
				Name: normalizedSegment,
			})
			continue
		}

		segments = append(segments, models.ResourceIdSegment{
			Type:       models.StaticSegment,
			Name:       normalizedSegment,
			FixedValue: &normalizedSegment,
		})
	}

	output := resourceUriMetadata{
		resourceIdName: nil,
		uriSuffix:      nil,
	}

	// UriSuffixes are "operations" on a given Resource ID/URI - for example `/restart`
	// or in the case of List operations /providers/Microsoft.Blah/listAllTheThings
	// we treat these as "operations" on the Resource ID and as such the "segments" should
	// only be for the Resource ID and not for the UriSuffix (which is as an additional field)
	lastUserValueSegment := -1
	for i, segment := range segments {
		// everything else technically is a user configurable component
		if segment.Type != models.StaticSegment {
			lastUserValueSegment = i
		}
	}
	if lastUserValueSegment >= 0 && len(segments) > lastUserValueSegment+1 {
		suffix := ""
		for _, segment := range segments[lastUserValueSegment+1:] {
			suffix += fmt.Sprintf("/%s", *segment.FixedValue)
		}
		output.uriSuffix = &suffix

		// remove any URI Suffix since this isn't relevant for the ID's
		segments = segments[0 : lastUserValueSegment+1]
	}

	allSegmentsAreStatic := true
	for _, segment := range segments {
		if segment.Type != models.StaticSegment {
			allSegmentsAreStatic = false
			break
		}
	}
	if allSegmentsAreStatic {
		// if it's not an ARM ID there's nothing to output here, but new up a placeholder
		// to be able to give us a normalized id for the suffix
		pri := models.ParsedResourceId{
			Constants: result.constants,
			Segments:  segments,
		}
		suffix := pri.NormalizedResourceId()
		output.uriSuffix = &suffix
	} else {
		output.resourceId = &models.ParsedResourceId{
			Constants: result.constants,
			Segments:  segments,
		}
	}

	return &output, nil
}

// determineNamesForResourceIds returns a map[name]ParsedResourceID and map[Uri]Name based on the Resource Manager URI's available
func determineNamesForResourceIds(urisToObjects map[string]resourceUriMetadata) (*map[string]models.ParsedResourceId, *map[string]string, error) {
	// now that we have all of the Resource ID's, we then need to go through and determine Unique ID's for those
	// we need all of them here to avoid conflicts, e.g. AuthorizationRule which can be a NamespaceAuthorizationRule
	// or an EventHubAuthorizationRule, but is named AuthorizationRule in both

	// now we need to go through and determine candidate names for these Resource ID's
	// we can do this using the last user configurable segment
	// first let's go through and determine if we have any conflicting 'key' segments
	uniqueNamesForUris := make(map[string]models.ParsedResourceId) // map[name]uri
	conflictingKeys := make(map[string][]models.ParsedResourceId)  // map[name]uris
	for _, resourceId := range urisToObjects {
		// if it's just a suffix (e.g. root-level ListAll calls) iterate over it
		if resourceId.resourceId == nil {
			continue
		}

		// when there's a Uri Suffix we should pass in both the full uri and just the resource manager uri so we can
		// skip it if this is a full uri (with a suffix), since the name comes from the resource manager uri instead
		if resourceId.uriSuffix != nil {
			continue
		}

		userSpecifiableSegments := resourceId.resourceId.UserSpecifiableSegmentNames()
		if len(userSpecifiableSegments) == 0 {
			return nil, nil, fmt.Errorf("no user specifiable segments for %+v - this is not a resource id?", *resourceId.resourceId)
		}

		lastSegment := userSpecifiableSegments[len(userSpecifiableSegments)-1]

		// however if this is an ARM Resource ID we should have Key-Value Pairs - in which case the name likely
		// wants to come from the Key and not the Value
		if len(resourceId.resourceId.Segments)%2 == 0 {
			lastSegment = keyForUserSpecifiableSegment(lastSegment, resourceId.resourceId.Segments)
		}

		lastSegment = strings.TrimSuffix(lastSegment, "Name")
		if len(resourceId.resourceId.Segments) > 1 {
			// if the first segment is a Scope, prefix the name with 'Scoped'
			if firstSegment := resourceId.resourceId.Segments[0]; firstSegment.Type == models.ScopeSegment {
				lastSegment = fmt.Sprintf("Scoped%s", normalizeSegmentName(lastSegment))
			}
		}
		normalizedKey := normalizeSegmentName(lastSegment)
		if !strings.HasSuffix(normalizedKey, "Id") {
			normalizedKey = normalizedKey + "Id"
		}

		if uris, existing := conflictingKeys[normalizedKey]; existing {
			uris = append(uris, *resourceId.resourceId)
			conflictingKeys[normalizedKey] = uris
			continue
		}

		if existingUri, existing := uniqueNamesForUris[normalizedKey]; existing {
			conflictingKeys[normalizedKey] = []models.ParsedResourceId{existingUri, *resourceId.resourceId}
			delete(uniqueNamesForUris, normalizedKey)
			continue
		}

		uniqueNamesForUris[normalizedKey] = *resourceId.resourceId
	}

	// at this stage uniqueNamesForUris contains the unique names : uris
	// so we need to iterate over conflictingKeys and find unique names for those
	if len(conflictingKeys) > 0 {
		uniqueSegments, err := determineUniqueSegmentNames(conflictingKeys)
		if err != nil {
			return nil, nil, fmt.Errorf("determining unique segment names: %+v", err)
		}
		for k, v := range *uniqueSegments {
			uniqueNamesForUris[k] = v
		}
	}

	// finally compose a list of uris -> names so these are easier to map back
	urisToNames := make(map[string]string, 0)
	for k, v := range uniqueNamesForUris {
		urisToNames[v.NormalizedResourceManagerResourceId()] = k
	}

	return &uniqueNamesForUris, &urisToNames, nil
}

func keyForUserSpecifiableSegment(value string, segments []models.ResourceIdSegment) string {
	index := -1
	for i, segment := range segments {
		if segment.Name == value {
			index = i
			break
		}
	}
	v := segments[index-1]
	return v.Name
}

// determineUniqueSegmentNames returns a map[name]uri
func determineUniqueSegmentNames(input map[string][]models.ParsedResourceId) (*map[string]models.ParsedResourceId, error) {
	identifiers := make(map[string]models.ParsedResourceId, 0)
	for _, uris := range input {
		proposed := make(map[string]models.ParsedResourceId)
		for _, resourceId := range uris {
			names := resourceId.UserSpecifiableSegmentNames()
			if len(names) < 2 {
				return nil, fmt.Errorf("insufficient segments to create a unique identifier from %+v", resourceId)
			}

			isResourceManagerId := len(resourceId.Segments)%2 == 0

			// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}
			// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}
			childName := names[len(names)-1]
			if isResourceManagerId {
				childName = keyForUserSpecifiableSegment(childName, resourceId.Segments)
			}
			childName += "Id"

			// the names must already conflict or we wouldn't be here
			parentName := names[len(names)-2]
			if isResourceManagerId {
				parentName = keyForUserSpecifiableSegment(parentName, resourceId.Segments)
			}
			key := fmt.Sprintf("%s%s", normalizeSegmentName(parentName), normalizeSegmentName(childName))

			// check if we have an existing match for ParentChild
			if v, ok := keyHasConflicts(key, identifiers, proposed); ok && *v != resourceId.NormalizedResourceManagerResourceId() {
				if len(names) < 3 {
					return nil, fmt.Errorf("need a third unique identifier to make %q unique", resourceId)
				}

				// prefix the parent name
				if len(names) >= 3 {
					grandparentName := names[len(names)-3]
					if isResourceManagerId {
						grandparentName = keyForUserSpecifiableSegment(grandparentName, resourceId.Segments)
					}
					normalized := normalizeSegmentName(grandparentName)
					key = fmt.Sprintf("%s%s", normalized, key)

					if v, ok := keyHasConflicts(key, identifiers, proposed); ok {
						if *v != resourceId.NormalizedResourceManagerResourceId() {
							// e.g. Web App Slots
							// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}
							// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}
							if len(names) < 4 {
								return nil, fmt.Errorf("need a fourth unique identifier to make %q unique", resourceId)
							}

							// prefix the grandparent name
							if len(names) >= 4 {
								greatGrandParentName := (names)[len(names)-4]
								if isResourceManagerId {
									greatGrandParentName = keyForUserSpecifiableSegment(greatGrandParentName, resourceId.Segments)
								}
								normalized := normalizeSegmentName(greatGrandParentName)
								key = fmt.Sprintf("%s%s", normalized, key)

								if v, ok := keyHasConflicts(key, identifiers, proposed); ok {
									if *v != resourceId.NormalizedResourceManagerResourceId() {
										return nil, fmt.Errorf("conflicting unique keys for %q - %q and %q conflict", key, *v, resourceId)
									}
								}
							}
						}
					}
				}
			}

			proposed[key] = resourceId
		}

		// at this point all of these must be unique, so let's add them to identifiers
		for k, v := range proposed {
			identifiers[k] = v
		}
	}
	return &identifiers, nil
}

func keyHasConflicts(key string, identifiers, proposed map[string]models.ParsedResourceId) (*string, bool) {
	if v, ok := identifiers[key]; ok {
		id := v.NormalizedResourceId()
		return &id, true
	}
	if v, ok := proposed[key]; ok {
		id := v.NormalizedResourceId()
		return &id, true
	}

	return nil, false
}

func mapNamesToResourceIds(urisToNames map[string]string, urisToMetadata map[string]resourceUriMetadata) (*map[string]resourceUriMetadata, error) {
	output := make(map[string]resourceUriMetadata, 0)

	for uri, metadata := range urisToMetadata {
		// ID's with just Suffixes are valid and won't have an ID Type, so skip those
		if metadata.resourceId == nil {
			output[uri] = metadata
			continue
		}

		name, ok := urisToNames[metadata.resourceId.NormalizedResourceManagerResourceId()]
		if !ok {
			return nil, fmt.Errorf("Resource ID : Name mapping not found for %q", uri)
		}

		output[metadata.resourceId.NormalizedResourceManagerResourceId()] = resourceUriMetadata{
			resourceIdName: &name,
			// intentionally don't map over the UriSuffix since this is handled above
		}

		// when there's a suffix, we need to output the full uri in the map too
		if metadata.uriSuffix != nil {
			metadata.resourceIdName = &name
			output[uri] = metadata
		}
	}

	return &output, nil
}
