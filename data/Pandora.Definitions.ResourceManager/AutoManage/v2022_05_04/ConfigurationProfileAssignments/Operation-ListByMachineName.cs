using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.AutoManage.v2022_05_04.ConfigurationProfileAssignments;

internal class ListByMachineNameOperation : Pandora.Definitions.Operations.GetOperation
{
    public override ResourceID? ResourceId() => new MachineId();

    public override Type? ResponseObject() => typeof(ConfigurationProfileAssignmentListModel);

    public override string? UriSuffix() => "/providers/Microsoft.AutoManage/configurationProfileAssignments";


}
