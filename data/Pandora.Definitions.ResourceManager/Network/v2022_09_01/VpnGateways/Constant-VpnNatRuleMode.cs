using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Network.v2022_09_01.VpnGateways;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum VpnNatRuleModeConstant
{
    [Description("EgressSnat")]
    EgressSnat,

    [Description("IngressSnat")]
    IngressSnat,
}