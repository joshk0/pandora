using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Synapse.v2021_06_01.IntegrationRuntime;


internal class ManagedIntegrationRuntimeStatusTypePropertiesModel
{
    [DateFormat(DateFormatAttribute.DateFormat.RFC3339)]
    [JsonPropertyName("createTime")]
    public DateTime? CreateTime { get; set; }

    [JsonPropertyName("lastOperation")]
    public ManagedIntegrationRuntimeOperationResultModel? LastOperation { get; set; }

    [JsonPropertyName("nodes")]
    public List<ManagedIntegrationRuntimeNodeModel>? Nodes { get; set; }

    [JsonPropertyName("otherErrors")]
    public List<ManagedIntegrationRuntimeErrorModel>? OtherErrors { get; set; }
}