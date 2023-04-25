using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ContainerInstance.v2023_05_01.ContainerInstance;


internal class AzureFileVolumeModel
{
    [JsonPropertyName("readOnly")]
    public bool? ReadOnly { get; set; }

    [JsonPropertyName("shareName")]
    [Required]
    public string ShareName { get; set; }

    [JsonPropertyName("storageAccountKey")]
    public string? StorageAccountKey { get; set; }

    [JsonPropertyName("storageAccountName")]
    [Required]
    public string StorageAccountName { get; set; }
}