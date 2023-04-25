using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using System;
using System.Collections.Generic;
using System.Net;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.MachineLearningServices.v2023_04_01.EnvironmentVersion;

internal class RegistryEnvironmentVersionsCreateOrUpdateOperation : Pandora.Definitions.Operations.PutOperation
{
    public override bool LongRunning() => true;

    public override Type? RequestObject() => typeof(EnvironmentVersionResourceModel);

    public override ResourceID? ResourceId() => new RegistryEnvironmentVersionId();

    public override Type? ResponseObject() => typeof(EnvironmentVersionResourceModel);


}