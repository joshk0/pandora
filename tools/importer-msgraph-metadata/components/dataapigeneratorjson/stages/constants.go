// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stages

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
	"github.com/hashicorp/pandora/tools/importer-msgraph-metadata/components/dataapigeneratorjson/helpers"
	"github.com/hashicorp/pandora/tools/importer-msgraph-metadata/components/dataapigeneratorjson/transforms"
	"github.com/hashicorp/pandora/tools/importer-msgraph-metadata/internal/logging"
)

var _ Stage = ConstantStage{}

type ConstantStage struct {
	// APIVersion specifies the APIVersion within the Service where the Constants exist.
	APIVersion string

	// APIResource specifies the APIResource within the APIVersion where the Constants exist.
	APIResource string

	// Constants specifies the map of Constant Name (key) to SDKConstant (value) which should be
	// persisted.
	Constants map[string]models.SDKConstant

	// ResourceIDs specifies a map of Resource ID Name (key) to ResourceID (value) that should
	// be persisted.
	ResourceIDs map[string]models.ResourceID

	// ServiceName specifies the name of the Service within which the Constants exist.
	ServiceName string
}

func (g ConstantStage) Name() string {
	return "Constants"
}

func (g ConstantStage) Generate(input *helpers.FileSystem) error {
	logging.Log.Debug("Generating Constants")

	for constantName, constantVal := range g.Constants {
		logging.Log.Trace(fmt.Sprintf("Processing Constant %q", constantName))

		mapped, err := transforms.MapSDKConstantToRepository(constantName, constantVal)
		if err != nil {
			return fmt.Errorf("mapping SDKConstant %q: %+v", constantName, err)
		}

		// {workingDirectory}/Service/APIVersion/APIResource/Constant-{Name}.json
		path := filepath.Join(g.ServiceName, g.APIVersion, g.APIResource, fmt.Sprintf("Constant-%s.json", constantName))
		logging.Log.Trace(fmt.Sprintf("Staging to %s", path))
		if err := input.Stage(path, *mapped); err != nil {
			return fmt.Errorf("staging Constant %q: %+v", constantName, err)
		}
	}

	// ResourceIDs also contain Constants - so we need to pull those out and persist them too
	for resourceIdName, resourceId := range g.ResourceIDs {
		logging.Log.Trace(fmt.Sprintf("Processing Constants within the Resource ID %q", resourceIdName))

		for constantName, constantVal := range resourceId.Constants {
			logging.Log.Trace(fmt.Sprintf("Processing Constant %q", constantName))

			mapped, err := transforms.MapSDKConstantToRepository(constantName, constantVal)
			if err != nil {
				return fmt.Errorf("mapping SDKConstant %q: %+v", constantName, err)
			}

			// {workingDirectory}/Service/APIVersion/APIResource/Constant-{Name}.json
			path := filepath.Join(g.ServiceName, g.APIVersion, g.APIResource, fmt.Sprintf("Constant-%s.json", constantName))
			logging.Log.Trace(fmt.Sprintf("Staging to %s", path))
			if err := input.Stage(path, *mapped); err != nil {
				return fmt.Errorf("staging Constant %q: %+v", constantName, err)
			}
		}
	}

	return nil
}
