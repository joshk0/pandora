package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hashicorp/pandora/tools/generator-go-sdk/generator"
	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
	"github.com/hashicorp/pandora/tools/sdk/services"
)

type GeneratorInput struct {
	apiServerEndpoint string
	outputDirectory   string
	sdkName           string
	services          []string
	settings          generator.Settings
}

type SDK struct {
	baseClientMethod   string
	baseClientPackage  string
	clients            []func(string) resourcemanager.Client
	outputSubDirectory string
	versionMapper      func(string) (*string, error)
}

var availableSDKs = map[string]SDK{
	"resource-manager": {
		baseClientMethod:  "NewResourceManagerClient",
		baseClientPackage: "resourcemanager",
		clients: []func(string) resourcemanager.Client{
			resourcemanager.NewResourceManagerClient,
		},
		outputSubDirectory: "resource-manager",
		versionMapper:      func(in string) (*string, error) { return &in, nil },
	},

	"microsoft-graph": {
		baseClientMethod:  "NewMsGraphClient",
		baseClientPackage: "msgraph",
		clients: []func(string) resourcemanager.Client{
			resourcemanager.NewMicrosoftGraphStableV1Client,
			resourcemanager.NewMicrosoftGraphBetaClient,
		},
		outputSubDirectory: "microsoft-graph",
		versionMapper: func(in string) (*string, error) {
			version := ""
			switch in {
			case "StableV1":
				version = "v1.0"
			case "Beta":
				version = "beta"
			}
			if version == "" {
				return nil, fmt.Errorf("unsupported API version: %q", in)
			}
			return &version, nil
		},
	},
}

func main() {
	input := GeneratorInput{
		settings: generator.Settings{},
	}

	var serviceNames string

	f := flag.NewFlagSet("generator-go-sdk", flag.ExitOnError)
	f.StringVar(&input.apiServerEndpoint, "data-api", "http://localhost:5000", "-data-api=http://localhost:5000")
	f.StringVar(&input.outputDirectory, "output-dir", "", "-output-dir=../generated-sdk-dev")
	f.StringVar(&input.sdkName, "sdk", "resource-manager", "-sdk=resource-manager")
	f.StringVar(&serviceNames, "services", "", "A list of comma separated Service named from the Data API to import")
	if err := f.Parse(os.Args[1:]); err != nil {
		log.Fatalf("parsing arguments: %+v", err)
	}

	if serviceNames != "" {
		input.services = strings.Split(serviceNames, ",")
	}

	if input.outputDirectory == "" {
		homeDir, _ := os.UserHomeDir()
		input.outputDirectory = filepath.Join(homeDir, "/Desktop/generated-sdk-dev")
	}

	if input.sdkName == "resource-manager" {
		input.settings.UseOldBaseLayerFor(
			// @tombuildsstuff: New Services should now use the `hashicorp/go-azure-sdk` base layer by default
			// instead of the base layer from `Azure/go-autorest` - as such this list is for compatibility purposes
			// with services already used in `terraform-provider-azurerm`. These services will be gradually removed
			// from this list to ensure they're migrated across to using `hashicorp/go-azure-sdk`s base layer.

			// NOTE: also see the list in ./tools/generator-terraform/generator/definitions/template_service_client.go
			// for services/resources which are auto-generated
			"ContainerApps",
			"ContainerInstance",
			"CosmosDB",
			"DataShare",
			"FrontDoor",
			"Insights",
			"Kusto",
			"Maintenance",
			"ManagedServices",
			"RecoveryServicesBackup", // error: generating Service "RecoveryServicesBackup" / Version "2023-04-01" / Resource "Operation": generating methods: templating methods (using hashicorp/go-azure-sdk): templating: building methods: building response struct template: existing model "ValidateOperationResponse" conflicts with the operation response model for "Validate"
			"Security",
			"SecurityInsights",
			"ServiceFabric",
			"ServiceFabricManagedCluster",
			"ServiceLinker",
			"SqlVirtualMachine",
			"Storage",
			"StreamAnalytics",
			"Subscription",
			"TimeSeriesInsights",

			// Automation @ 2022-08-08 uses the new base layer, so let's invert the older versions for now
			"Automation@2015-10-31",
			"Automation@2019-06-01",
			"Automation@2020-01-13-preview",
			"Automation@2021-06-22",

			// @tombuildsstuff: there's generated resources associated with these three - please check before removing these
			"ContainerService",
			"LoadTestService",
			"ManagedIdentity",

			// @tombuildsstuff: KeyVault requires that the exact casing retrieved from the API is re-sent back to the API
			// as such will require custom work in the Provider (potentially a custom unmarshaller from the HTTP Body) to support this
			"KeyVault",
		)
	}

	if err := run(input); err != nil {
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run(input GeneratorInput) error {
	sdk, ok := availableSDKs[input.sdkName]
	if !ok {
		return fmt.Errorf("unsupported SDK %q", input.sdkName)
	}

	input.outputDirectory = path.Join(input.outputDirectory, sdk.outputSubDirectory)

	for _, apiClient := range sdk.clients {
		client := apiClient(input.apiServerEndpoint)

		var loadedServices services.ResourceManagerServices
		if len(input.services) > 0 {
			log.Printf("[DEBUG] Loading the Services from the Data API %q..", strings.Join(input.services, " / "))
			servicesResult, err := services.GetResourceManagerServicesByName(client, input.services)
			if err != nil {
				return fmt.Errorf("retrieving resource manager services: %+v", err)
			}
			loadedServices = *servicesResult
		} else {
			log.Printf("[DEBUG] Loading All Services..")
			servicesResult, err := services.GetResourceManagerServices(client)
			if err != nil {
				return fmt.Errorf("retrieving resource manager services: %+v", err)
			}
			loadedServices = *servicesResult
		}

		errCh := make(chan error, 1)
		waitDone := make(chan struct{}, 1)
		var wg sync.WaitGroup
		addErr := func(err error) {
			// only put one err to channel
			select {
			case errCh <- err:
			default:
			}
		}

		generatorService := generator.NewServiceGenerator(input.settings)
		for serviceName, service := range loadedServices.Services {
			log.Printf("[DEBUG] Service %q..", serviceName)
			if !service.Details.Generate {
				log.Printf("[DEBUG] .. is opted out of generation, skipping..")
				continue
			}

			wg.Add(1)
			go func(serviceName string, service services.ResourceManagerService, input GeneratorInput) {
				defer wg.Done()
				log.Printf("[DEBUG] Service %q", serviceName)
				for versionNumber, versionDetails := range service.Versions {
					log.Printf("[DEBUG]   Version %q", versionNumber)
					versionName, err := sdk.versionMapper(versionNumber)
					if err != nil {
						addErr(fmt.Errorf("generating Service %q / Version %q: %+v", serviceName, versionNumber, err))
						return
					}
					for resourceName, resourceDetails := range versionDetails.Resources {
						log.Printf("[DEBUG]      Resource %q..", resourceName)
						generatorData := generator.ServiceGeneratorInput{
							ServiceName:       serviceName,
							ServiceDetails:    service,
							VersionName:       *versionName,
							VersionDetails:    versionDetails,
							ResourceName:      resourceName,
							ResourceDetails:   resourceDetails,
							BaseClientMethod:  sdk.baseClientMethod,
							BaseClientPackage: sdk.baseClientPackage,
							OutputDirectory:   input.outputDirectory,
							Source:            versionDetails.Details.Source,
						}
						log.Printf("[DEBUG] Generating Service %q / Version %q / Resource %q..", serviceName, versionNumber, resourceName)
						if err := generatorService.Generate(generatorData); err != nil {
							addErr(fmt.Errorf("generating Service %q / Version %q / Resource %q: %+v", serviceName, versionNumber, resourceName, err))
							return
						}
						log.Printf("[DEBUG] Generated Service %q / Version %q / Resource %q..", serviceName, versionNumber, resourceName)
					}

					// then output the Meta Client
					generatorData := generator.VersionInput{
						OutputDirectory:   input.outputDirectory,
						ServiceName:       serviceName,
						VersionName:       *versionName,
						BaseClientPackage: sdk.baseClientPackage,
						SdkPackage:        sdk.outputSubDirectory,
						Resources:         versionDetails.Resources,
						Source:            versionDetails.Details.Source,
					}
					generatorData.UseNewBaseLayer = false
					if input.settings.ShouldUseNewBaseLayer(serviceName, versionNumber) {
						generatorData.UseNewBaseLayer = true
					}
					log.Printf("[DEBUG] Generating Service %q / Version %q..", serviceName, versionNumber)
					if err := generatorService.GenerateForVersion(generatorData); err != nil {
						addErr(fmt.Errorf("generating Service %q / Version %q: %+v", serviceName, versionNumber, err))
						return
					}
					log.Printf("[DEBUG] Generated Service %q / Version %q..", serviceName, versionNumber)
				}
			}(serviceName, service, input)
		}

		go func() {
			wg.Wait()
			waitDone <- struct{}{}
		}()

		select {
		case <-waitDone:
			break
		case err := <-errCh:
			return err
		}
	}

	return nil
}
