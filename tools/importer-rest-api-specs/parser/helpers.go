package parser

import (
	"fmt"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
	"strings"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/cleanup"

	"github.com/go-openapi/spec"
)

func fragmentNameFromReference(input spec.Ref) *string {
	fragmentName := input.String()
	return fragmentNameFromString(fragmentName)
}

func fragmentNameFromString(fragmentName string) *string {
	if fragmentName != "" {
		fragments := strings.Split(fragmentName, "/") // format #/definitions/ConfigurationStoreListResult
		referenceName := fragments[len(fragments)-1]
		return &referenceName
	}

	return nil
}

func inlinedModelName(parentModelName, fieldName string) string {
	// intentionally split out for consistency
	val := fmt.Sprintf("%s%s", strings.Title(parentModelName), strings.Title(fieldName))
	return cleanup.NormalizeName(val)
}

func operationMatchesTag(operation *spec.Operation, tag *string) bool {
	// if there's no tags defined, we should capture it when the tag matched
	if tag == nil {
		return len(operation.Tags) == 0
	}

	for _, thisTag := range operation.Tags {
		if thisTag == *tag {
			return true
		}
	}

	return false
}

func operationShouldBeIgnored(input string) bool {
	// we're not concerned with exposing the "operations" API's at this time - e.g.
	// /providers/Microsoft.EnterpriseKnowledgeGraph/services
	if strings.HasPrefix(strings.ToLower(input), "/providers/") {
		return true
	}

	// LRO's shouldn't be directly exposed
	if strings.Contains(strings.ToLower(input), "/operationresults/{operationid}") {
		return true
	}

	return false
}

// topLevelObjectDefinition returns the top level object definition, that is a Constant or Model (or simple type) directly
func topLevelObjectDefinition(input models.ObjectDefinition) models.ObjectDefinition {
	if input.NestedItem != nil {
		return topLevelObjectDefinition(*input.NestedItem)
	}

	return input
}
