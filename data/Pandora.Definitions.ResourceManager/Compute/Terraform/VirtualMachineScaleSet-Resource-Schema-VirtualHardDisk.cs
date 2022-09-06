using System.Collections.Generic;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.CommonSchema;

namespace Pandora.Definitions.ResourceManager.Compute.Terraform;

public class VirtualMachineScaleSetResourceVirtualHardDiskSchema
{

    [HclName("uri")]
    [Optional]
    public string Uri { get; set; }

}