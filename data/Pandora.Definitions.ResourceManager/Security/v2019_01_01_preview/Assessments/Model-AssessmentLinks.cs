using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Security.v2019_01_01_preview.Assessments;


internal class AssessmentLinksModel
{
    [JsonPropertyName("azurePortalUri")]
    public string? AzurePortalUri { get; set; }
}