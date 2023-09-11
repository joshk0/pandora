using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ApiManagement.v2023_03_01_preview.Product;

internal class Definition : ResourceDefinition
{
    public string Name => "Product";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new CreateOrUpdateOperation(),
        new DeleteOperation(),
        new GetOperation(),
        new GetEntityTagOperation(),
        new ListByServiceOperation(),
        new UpdateOperation(),
        new WorkspaceProductCreateOrUpdateOperation(),
        new WorkspaceProductDeleteOperation(),
        new WorkspaceProductGetOperation(),
        new WorkspaceProductGetEntityTagOperation(),
        new WorkspaceProductListByServiceOperation(),
        new WorkspaceProductUpdateOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(ProductStateConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(ProductContractModel),
        typeof(ProductContractPropertiesModel),
        typeof(ProductUpdateParametersModel),
        typeof(ProductUpdatePropertiesModel),
    };
}
