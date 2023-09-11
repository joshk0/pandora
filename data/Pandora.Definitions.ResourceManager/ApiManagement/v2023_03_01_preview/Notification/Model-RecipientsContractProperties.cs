using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ApiManagement.v2023_03_01_preview.Notification;


internal class RecipientsContractPropertiesModel
{
    [JsonPropertyName("emails")]
    public List<string>? Emails { get; set; }

    [JsonPropertyName("users")]
    public List<string>? Users { get; set; }
}
