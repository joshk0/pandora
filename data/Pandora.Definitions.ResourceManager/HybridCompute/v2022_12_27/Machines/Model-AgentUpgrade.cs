using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.HybridCompute.v2022_12_27.Machines;


internal class AgentUpgradeModel
{
    [JsonPropertyName("correlationId")]
    public string? CorrelationId { get; set; }

    [JsonPropertyName("desiredVersion")]
    public string? DesiredVersion { get; set; }

    [JsonPropertyName("enableAutomaticUpgrade")]
    public bool? EnableAutomaticUpgrade { get; set; }

    [JsonPropertyName("lastAttemptMessage")]
    public string? LastAttemptMessage { get; set; }

    [JsonPropertyName("lastAttemptStatus")]
    public LastAttemptStatusEnumConstant? LastAttemptStatus { get; set; }

    [JsonPropertyName("lastAttemptTimestamp")]
    public string? LastAttemptTimestamp { get; set; }
}