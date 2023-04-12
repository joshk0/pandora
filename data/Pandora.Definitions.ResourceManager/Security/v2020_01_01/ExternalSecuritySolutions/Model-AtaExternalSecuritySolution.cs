using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Security.v2020_01_01.ExternalSecuritySolutions;

[ValueForType("ATA")]
internal class AtaExternalSecuritySolutionModel : ExternalSecuritySolutionModel
{
    [JsonPropertyName("properties")]
    public AtaSolutionPropertiesModel? Properties { get; set; }
}