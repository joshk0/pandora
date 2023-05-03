using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Network.v2022_09_01.ApplicationGateways;

internal class Definition : ResourceDefinition
{
    public string Name => "ApplicationGateways";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new BackendHealthOperation(),
        new BackendHealthOnDemandOperation(),
        new CreateOrUpdateOperation(),
        new DeleteOperation(),
        new GetOperation(),
        new GetSslPredefinedPolicyOperation(),
        new ListOperation(),
        new ListAllOperation(),
        new ListAvailableRequestHeadersOperation(),
        new ListAvailableResponseHeadersOperation(),
        new ListAvailableServerVariablesOperation(),
        new ListAvailableSslOptionsOperation(),
        new ListAvailableSslPredefinedPoliciesOperation(),
        new ListAvailableWafRuleSetsOperation(),
        new StartOperation(),
        new StopOperation(),
        new UpdateTagsOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(ApplicationGatewayBackendHealthServerHealthConstant),
        typeof(ApplicationGatewayClientRevocationOptionsConstant),
        typeof(ApplicationGatewayCookieBasedAffinityConstant),
        typeof(ApplicationGatewayCustomErrorStatusCodeConstant),
        typeof(ApplicationGatewayFirewallModeConstant),
        typeof(ApplicationGatewayLoadDistributionAlgorithmConstant),
        typeof(ApplicationGatewayOperationalStateConstant),
        typeof(ApplicationGatewayProtocolConstant),
        typeof(ApplicationGatewayRedirectTypeConstant),
        typeof(ApplicationGatewayRequestRoutingRuleTypeConstant),
        typeof(ApplicationGatewaySkuNameConstant),
        typeof(ApplicationGatewaySslCipherSuiteConstant),
        typeof(ApplicationGatewaySslPolicyNameConstant),
        typeof(ApplicationGatewaySslPolicyTypeConstant),
        typeof(ApplicationGatewaySslProtocolConstant),
        typeof(ApplicationGatewayTierConstant),
        typeof(ApplicationGatewayTierTypesConstant),
        typeof(ApplicationGatewayWafRuleActionTypesConstant),
        typeof(ApplicationGatewayWafRuleStateTypesConstant),
        typeof(DdosSettingsProtectionModeConstant),
        typeof(DeleteOptionsConstant),
        typeof(FlowLogFormatTypeConstant),
        typeof(GatewayLoadBalancerTunnelInterfaceTypeConstant),
        typeof(GatewayLoadBalancerTunnelProtocolConstant),
        typeof(IPAllocationMethodConstant),
        typeof(IPVersionConstant),
        typeof(LoadBalancerBackendAddressAdminStateConstant),
        typeof(NatGatewaySkuNameConstant),
        typeof(NetworkInterfaceAuxiliaryModeConstant),
        typeof(NetworkInterfaceMigrationPhaseConstant),
        typeof(NetworkInterfaceNicTypeConstant),
        typeof(ProvisioningStateConstant),
        typeof(PublicIPAddressMigrationPhaseConstant),
        typeof(PublicIPAddressSkuNameConstant),
        typeof(PublicIPAddressSkuTierConstant),
        typeof(RouteNextHopTypeConstant),
        typeof(SecurityRuleAccessConstant),
        typeof(SecurityRuleDirectionConstant),
        typeof(SecurityRuleProtocolConstant),
        typeof(TransportProtocolConstant),
        typeof(VirtualNetworkPrivateEndpointNetworkPoliciesConstant),
        typeof(VirtualNetworkPrivateLinkServiceNetworkPoliciesConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(ApplicationGatewayModel),
        typeof(ApplicationGatewayAuthenticationCertificateModel),
        typeof(ApplicationGatewayAuthenticationCertificatePropertiesFormatModel),
        typeof(ApplicationGatewayAutoscaleConfigurationModel),
        typeof(ApplicationGatewayAvailableSslOptionsModel),
        typeof(ApplicationGatewayAvailableSslOptionsPropertiesFormatModel),
        typeof(ApplicationGatewayAvailableWafRuleSetsResultModel),
        typeof(ApplicationGatewayBackendAddressModel),
        typeof(ApplicationGatewayBackendAddressPoolModel),
        typeof(ApplicationGatewayBackendAddressPoolPropertiesFormatModel),
        typeof(ApplicationGatewayBackendHTTPSettingsModel),
        typeof(ApplicationGatewayBackendHTTPSettingsPropertiesFormatModel),
        typeof(ApplicationGatewayBackendHealthModel),
        typeof(ApplicationGatewayBackendHealthHTTPSettingsModel),
        typeof(ApplicationGatewayBackendHealthOnDemandModel),
        typeof(ApplicationGatewayBackendHealthPoolModel),
        typeof(ApplicationGatewayBackendHealthServerModel),
        typeof(ApplicationGatewayBackendSettingsModel),
        typeof(ApplicationGatewayBackendSettingsPropertiesFormatModel),
        typeof(ApplicationGatewayClientAuthConfigurationModel),
        typeof(ApplicationGatewayConnectionDrainingModel),
        typeof(ApplicationGatewayCustomErrorModel),
        typeof(ApplicationGatewayFirewallDisabledRuleGroupModel),
        typeof(ApplicationGatewayFirewallExclusionModel),
        typeof(ApplicationGatewayFirewallRuleModel),
        typeof(ApplicationGatewayFirewallRuleGroupModel),
        typeof(ApplicationGatewayFirewallRuleSetModel),
        typeof(ApplicationGatewayFirewallRuleSetPropertiesFormatModel),
        typeof(ApplicationGatewayFrontendIPConfigurationModel),
        typeof(ApplicationGatewayFrontendIPConfigurationPropertiesFormatModel),
        typeof(ApplicationGatewayFrontendPortModel),
        typeof(ApplicationGatewayFrontendPortPropertiesFormatModel),
        typeof(ApplicationGatewayGlobalConfigurationModel),
        typeof(ApplicationGatewayHTTPListenerModel),
        typeof(ApplicationGatewayHTTPListenerPropertiesFormatModel),
        typeof(ApplicationGatewayHeaderConfigurationModel),
        typeof(ApplicationGatewayIPConfigurationModel),
        typeof(ApplicationGatewayIPConfigurationPropertiesFormatModel),
        typeof(ApplicationGatewayListenerModel),
        typeof(ApplicationGatewayListenerPropertiesFormatModel),
        typeof(ApplicationGatewayLoadDistributionPolicyModel),
        typeof(ApplicationGatewayLoadDistributionPolicyPropertiesFormatModel),
        typeof(ApplicationGatewayLoadDistributionTargetModel),
        typeof(ApplicationGatewayLoadDistributionTargetPropertiesFormatModel),
        typeof(ApplicationGatewayOnDemandProbeModel),
        typeof(ApplicationGatewayPathRuleModel),
        typeof(ApplicationGatewayPathRulePropertiesFormatModel),
        typeof(ApplicationGatewayPrivateEndpointConnectionModel),
        typeof(ApplicationGatewayPrivateEndpointConnectionPropertiesModel),
        typeof(ApplicationGatewayPrivateLinkConfigurationModel),
        typeof(ApplicationGatewayPrivateLinkConfigurationPropertiesModel),
        typeof(ApplicationGatewayPrivateLinkIPConfigurationModel),
        typeof(ApplicationGatewayPrivateLinkIPConfigurationPropertiesModel),
        typeof(ApplicationGatewayProbeModel),
        typeof(ApplicationGatewayProbeHealthResponseMatchModel),
        typeof(ApplicationGatewayProbePropertiesFormatModel),
        typeof(ApplicationGatewayPropertiesFormatModel),
        typeof(ApplicationGatewayRedirectConfigurationModel),
        typeof(ApplicationGatewayRedirectConfigurationPropertiesFormatModel),
        typeof(ApplicationGatewayRequestRoutingRuleModel),
        typeof(ApplicationGatewayRequestRoutingRulePropertiesFormatModel),
        typeof(ApplicationGatewayRewriteRuleModel),
        typeof(ApplicationGatewayRewriteRuleActionSetModel),
        typeof(ApplicationGatewayRewriteRuleConditionModel),
        typeof(ApplicationGatewayRewriteRuleSetModel),
        typeof(ApplicationGatewayRewriteRuleSetPropertiesFormatModel),
        typeof(ApplicationGatewayRoutingRuleModel),
        typeof(ApplicationGatewayRoutingRulePropertiesFormatModel),
        typeof(ApplicationGatewaySkuModel),
        typeof(ApplicationGatewaySslCertificateModel),
        typeof(ApplicationGatewaySslCertificatePropertiesFormatModel),
        typeof(ApplicationGatewaySslPolicyModel),
        typeof(ApplicationGatewaySslPredefinedPolicyModel),
        typeof(ApplicationGatewaySslPredefinedPolicyPropertiesFormatModel),
        typeof(ApplicationGatewaySslProfileModel),
        typeof(ApplicationGatewaySslProfilePropertiesFormatModel),
        typeof(ApplicationGatewayTrustedClientCertificateModel),
        typeof(ApplicationGatewayTrustedClientCertificatePropertiesFormatModel),
        typeof(ApplicationGatewayTrustedRootCertificateModel),
        typeof(ApplicationGatewayTrustedRootCertificatePropertiesFormatModel),
        typeof(ApplicationGatewayUrlConfigurationModel),
        typeof(ApplicationGatewayUrlPathMapModel),
        typeof(ApplicationGatewayUrlPathMapPropertiesFormatModel),
        typeof(ApplicationGatewayWebApplicationFirewallConfigurationModel),
        typeof(ApplicationSecurityGroupModel),
        typeof(ApplicationSecurityGroupPropertiesFormatModel),
        typeof(BackendAddressPoolModel),
        typeof(BackendAddressPoolPropertiesFormatModel),
        typeof(CustomDnsConfigPropertiesFormatModel),
        typeof(DdosSettingsModel),
        typeof(DelegationModel),
        typeof(FlowLogModel),
        typeof(FlowLogFormatParametersModel),
        typeof(FlowLogPropertiesFormatModel),
        typeof(FrontendIPConfigurationModel),
        typeof(FrontendIPConfigurationPropertiesFormatModel),
        typeof(GatewayLoadBalancerTunnelInterfaceModel),
        typeof(IPConfigurationModel),
        typeof(IPConfigurationProfileModel),
        typeof(IPConfigurationProfilePropertiesFormatModel),
        typeof(IPConfigurationPropertiesFormatModel),
        typeof(IPTagModel),
        typeof(InboundNatRuleModel),
        typeof(InboundNatRulePropertiesFormatModel),
        typeof(LoadBalancerBackendAddressModel),
        typeof(LoadBalancerBackendAddressPropertiesFormatModel),
        typeof(NatGatewayModel),
        typeof(NatGatewayPropertiesFormatModel),
        typeof(NatGatewaySkuModel),
        typeof(NatRulePortMappingModel),
        typeof(NetworkInterfaceModel),
        typeof(NetworkInterfaceDnsSettingsModel),
        typeof(NetworkInterfaceIPConfigurationModel),
        typeof(NetworkInterfaceIPConfigurationPrivateLinkConnectionPropertiesModel),
        typeof(NetworkInterfaceIPConfigurationPropertiesFormatModel),
        typeof(NetworkInterfacePropertiesFormatModel),
        typeof(NetworkInterfaceTapConfigurationModel),
        typeof(NetworkInterfaceTapConfigurationPropertiesFormatModel),
        typeof(NetworkSecurityGroupModel),
        typeof(NetworkSecurityGroupPropertiesFormatModel),
        typeof(PrivateEndpointModel),
        typeof(PrivateEndpointConnectionModel),
        typeof(PrivateEndpointConnectionPropertiesModel),
        typeof(PrivateEndpointIPConfigurationModel),
        typeof(PrivateEndpointIPConfigurationPropertiesModel),
        typeof(PrivateEndpointPropertiesModel),
        typeof(PrivateLinkServiceModel),
        typeof(PrivateLinkServiceConnectionModel),
        typeof(PrivateLinkServiceConnectionPropertiesModel),
        typeof(PrivateLinkServiceConnectionStateModel),
        typeof(PrivateLinkServiceIPConfigurationModel),
        typeof(PrivateLinkServiceIPConfigurationPropertiesModel),
        typeof(PrivateLinkServicePropertiesModel),
        typeof(PublicIPAddressModel),
        typeof(PublicIPAddressDnsSettingsModel),
        typeof(PublicIPAddressPropertiesFormatModel),
        typeof(PublicIPAddressSkuModel),
        typeof(ResourceNavigationLinkModel),
        typeof(ResourceNavigationLinkFormatModel),
        typeof(ResourceSetModel),
        typeof(RetentionPolicyParametersModel),
        typeof(RouteModel),
        typeof(RoutePropertiesFormatModel),
        typeof(RouteTableModel),
        typeof(RouteTablePropertiesFormatModel),
        typeof(SecurityRuleModel),
        typeof(SecurityRulePropertiesFormatModel),
        typeof(ServiceAssociationLinkModel),
        typeof(ServiceAssociationLinkPropertiesFormatModel),
        typeof(ServiceDelegationPropertiesFormatModel),
        typeof(ServiceEndpointPolicyModel),
        typeof(ServiceEndpointPolicyDefinitionModel),
        typeof(ServiceEndpointPolicyDefinitionPropertiesFormatModel),
        typeof(ServiceEndpointPolicyPropertiesFormatModel),
        typeof(ServiceEndpointPropertiesFormatModel),
        typeof(SubResourceModel),
        typeof(SubnetModel),
        typeof(SubnetPropertiesFormatModel),
        typeof(TagsObjectModel),
        typeof(TrafficAnalyticsConfigurationPropertiesModel),
        typeof(TrafficAnalyticsPropertiesModel),
        typeof(VirtualNetworkTapModel),
        typeof(VirtualNetworkTapPropertiesFormatModel),
    };
}