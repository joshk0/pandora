using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.KeyVault.v2021_10_01.Secrets;

internal class Definition : ResourceDefinition
{
    public string Name => "Secrets";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new CreateOrUpdateOperation(),
        new GetOperation(),
        new ListOperation(),
        new UpdateOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(AccessPolicyUpdateKindConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(AttributesModel),
        typeof(SecretModel),
        typeof(SecretCreateOrUpdateParametersModel),
        typeof(SecretPatchParametersModel),
        typeof(SecretPatchPropertiesModel),
        typeof(SecretPropertiesModel),
    };
}