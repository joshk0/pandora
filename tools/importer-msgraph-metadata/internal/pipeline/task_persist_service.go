package pipeline

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	generatorjson "github.com/hashicorp/pandora/tools/data-api-sdk/v1/generator-json"
	dataapisdk "github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
	sdkModels "github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
)

func (p pipelineForService) persistApiDefinitions(sdkServices map[string]sdkModels.Service, commonTypes map[string]sdkModels.CommonTypes) error {
	for serviceName, service := range sdkServices {
		for version := range service.APIVersions {
			p.logger.Debug(fmt.Sprintf("removing any existing API Definitions for the Service %q with Version %q", serviceName, version))

			removeServiceOpts := generatorjson.RemoveServiceOptions{
				ServiceName:      serviceName,
				SourceDataOrigin: dataapisdk.MicrosoftGraphMetaDataSourceDataOrigin,
				SourceDataType:   dataapisdk.MicrosoftGraphSourceDataType,
				Version:          version,
			}

			if err := p.repo.RemoveService(removeServiceOpts); err != nil {
				return fmt.Errorf("removing existing API Definitions for Service %q with Version %q: %+v", serviceName, version, err)
			}
		}

		p.logger.Info(fmt.Sprintf("persisting API Definitions for Service %q..", serviceName))

		opts := generatorjson.SaveServiceOptions{
			SourceDataGitSHA: pointer.To(p.metadataGitSha),
			CommonTypes:      commonTypes,
			Service:          service,
			ServiceName:      serviceName,
			SourceDataOrigin: dataapisdk.MicrosoftGraphMetaDataSourceDataOrigin,
			SourceDataType:   dataapisdk.MicrosoftGraphSourceDataType,
		}

		if err := p.repo.SaveService(opts); err != nil {
			return fmt.Errorf("persisting Data API Definitions for Service %q: %+v", serviceName, err)
		}
	}

	return nil
}
