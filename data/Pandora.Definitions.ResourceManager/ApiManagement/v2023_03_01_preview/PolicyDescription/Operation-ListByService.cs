using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ApiManagement.v2023_03_01_preview.PolicyDescription;

internal class ListByServiceOperation : Pandora.Definitions.Operations.GetOperation
{
    public override ResourceID? ResourceId() => new ServiceId();

    public override Type? ResponseObject() => typeof(PolicyDescriptionCollectionModel);

    public override Type? OptionsObject() => typeof(ListByServiceOperation.ListByServiceOptions);

    public override string? UriSuffix() => "/policyDescriptions";

    internal class ListByServiceOptions
    {
        [QueryStringName("scope")]
        [Optional]
        public PolicyScopeContractConstant Scope { get; set; }
    }
}
