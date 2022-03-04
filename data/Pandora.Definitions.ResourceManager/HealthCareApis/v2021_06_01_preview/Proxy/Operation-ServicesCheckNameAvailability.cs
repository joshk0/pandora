using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System;
using System.Collections.Generic;
using System.Net;

namespace Pandora.Definitions.ResourceManager.HealthCareApis.v2021_06_01_preview.Proxy;

internal class ServicesCheckNameAvailabilityOperation : Operations.PostOperation
{
    public override IEnumerable<HttpStatusCode> ExpectedStatusCodes() => new List<HttpStatusCode>
        {
                HttpStatusCode.OK,
        };

    public override Type? RequestObject() => typeof(CheckNameAvailabilityParametersModel);

    public override ResourceID? ResourceId() => new SubscriptionId();

    public override Type? ResponseObject() => typeof(ServicesNameAvailabilityInfoModel);

    public override string? UriSuffix() => "/providers/Microsoft.HealthcareApis/checkNameAvailability";


}
