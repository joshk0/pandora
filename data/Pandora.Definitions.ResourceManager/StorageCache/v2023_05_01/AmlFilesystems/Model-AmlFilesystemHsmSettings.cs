using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.StorageCache.v2023_05_01.AmlFilesystems;


internal class AmlFilesystemHsmSettingsModel
{
    [JsonPropertyName("container")]
    [Required]
    public string Container { get; set; }

    [JsonPropertyName("importPrefix")]
    public string? ImportPrefix { get; set; }

    [JsonPropertyName("loggingContainer")]
    [Required]
    public string LoggingContainer { get; set; }
}