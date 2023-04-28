using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.HybridCompute.v2022_12_27.Extensions;

internal class Definition : ResourceDefinition
{
    public string Name => "Extensions";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new ExtensionMetadataGetOperation(),
        new ExtensionMetadataListOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {

    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(ExtensionValueModel),
        typeof(ExtensionValueListResultModel),
        typeof(ExtensionValuePropertiesModel),
    };
}