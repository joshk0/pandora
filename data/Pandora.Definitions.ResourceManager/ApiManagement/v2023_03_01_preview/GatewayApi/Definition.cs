using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ApiManagement.v2023_03_01_preview.GatewayApi;

internal class Definition : ResourceDefinition
{
    public string Name => "GatewayApi";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new CreateOrUpdateOperation(),
        new DeleteOperation(),
        new GetEntityTagOperation(),
        new ListByServiceOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(ApiTypeConstant),
        typeof(BearerTokenSendingMethodsConstant),
        typeof(ProtocolConstant),
        typeof(ProvisioningStateConstant),
        typeof(VersioningSchemeConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(ApiContactInformationModel),
        typeof(ApiContractModel),
        typeof(ApiContractPropertiesModel),
        typeof(ApiLicenseInformationModel),
        typeof(ApiVersionSetContractDetailsModel),
        typeof(AssociationContractModel),
        typeof(AssociationContractPropertiesModel),
        typeof(AuthenticationSettingsContractModel),
        typeof(OAuth2AuthenticationSettingsContractModel),
        typeof(OpenIdAuthenticationSettingsContractModel),
        typeof(SubscriptionKeyParameterNamesContractModel),
    };
}
