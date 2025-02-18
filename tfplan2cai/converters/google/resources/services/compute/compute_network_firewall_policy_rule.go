// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/NetworkFirewallPolicyRule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeNetworkFirewallPolicyRuleAssetType string = "compute.googleapis.com/NetworkFirewallPolicyRule"

func ResourceConverterComputeNetworkFirewallPolicyRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNetworkFirewallPolicyRuleAssetType,
		Convert:   GetComputeNetworkFirewallPolicyRuleCaiObject,
	}
}

func GetComputeNetworkFirewallPolicyRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/getRule?priority={{priority}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNetworkFirewallPolicyRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNetworkFirewallPolicyRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NetworkFirewallPolicyRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNetworkFirewallPolicyRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	ruleNameProp, err := expandComputeNetworkFirewallPolicyRuleRuleName(d.Get("rule_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rule_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(ruleNameProp)) && (ok || !reflect.DeepEqual(v, ruleNameProp)) {
		obj["ruleName"] = ruleNameProp
	}
	descriptionProp, err := expandComputeNetworkFirewallPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeNetworkFirewallPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeNetworkFirewallPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("match"); !tpgresource.IsEmptyValue(reflect.ValueOf(matchProp)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeNetworkFirewallPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	securityProfileGroupProp, err := expandComputeNetworkFirewallPolicyRuleSecurityProfileGroup(d.Get("security_profile_group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_profile_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(securityProfileGroupProp)) && (ok || !reflect.DeepEqual(v, securityProfileGroupProp)) {
		obj["securityProfileGroup"] = securityProfileGroupProp
	}
	tlsInspectProp, err := expandComputeNetworkFirewallPolicyRuleTlsInspect(d.Get("tls_inspect"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tls_inspect"); !tpgresource.IsEmptyValue(reflect.ValueOf(tlsInspectProp)) && (ok || !reflect.DeepEqual(v, tlsInspectProp)) {
		obj["tlsInspect"] = tlsInspectProp
	}
	directionProp, err := expandComputeNetworkFirewallPolicyRuleDirection(d.Get("direction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("direction"); !tpgresource.IsEmptyValue(reflect.ValueOf(directionProp)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	enableLoggingProp, err := expandComputeNetworkFirewallPolicyRuleEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	targetServiceAccountsProp, err := expandComputeNetworkFirewallPolicyRuleTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_service_accounts"); ok || !reflect.DeepEqual(v, targetServiceAccountsProp) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}
	targetSecureTagsProp, err := expandComputeNetworkFirewallPolicyRuleTargetSecureTags(d.Get("target_secure_tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_secure_tags"); ok || !reflect.DeepEqual(v, targetSecureTagsProp) {
		obj["targetSecureTags"] = targetSecureTagsProp
	}
	disabledProp, err := expandComputeNetworkFirewallPolicyRuleDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}

	return obj, nil
}

func expandComputeNetworkFirewallPolicyRuleRuleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeNetworkFirewallPolicyRuleMatchSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeNetworkFirewallPolicyRuleMatchDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedSrcNetworkScope, err := expandComputeNetworkFirewallPolicyRuleMatchSrcNetworkScope(original["src_network_scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcNetworkScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcNetworkScope"] = transformedSrcNetworkScope
	}

	transformedSrcNetworks, err := expandComputeNetworkFirewallPolicyRuleMatchSrcNetworks(original["src_networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcNetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcNetworks"] = transformedSrcNetworks
	}

	transformedDestNetworkScope, err := expandComputeNetworkFirewallPolicyRuleMatchDestNetworkScope(original["dest_network_scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestNetworkScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destNetworkScope"] = transformedDestNetworkScope
	}

	transformedLayer4Configs, err := expandComputeNetworkFirewallPolicyRuleMatchLayer4Configs(original["layer4_configs"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["layer4Configs"] = transformedLayer4Configs
	}

	transformedSrcSecureTags, err := expandComputeNetworkFirewallPolicyRuleMatchSrcSecureTags(original["src_secure_tags"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["srcSecureTags"] = transformedSrcSecureTags
	}

	transformedDestAddressGroups, err := expandComputeNetworkFirewallPolicyRuleMatchDestAddressGroups(original["dest_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["destAddressGroups"] = transformedDestAddressGroups
	}

	transformedSrcAddressGroups, err := expandComputeNetworkFirewallPolicyRuleMatchSrcAddressGroups(original["src_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["srcAddressGroups"] = transformedSrcAddressGroups
	}

	transformedSrcFqdns, err := expandComputeNetworkFirewallPolicyRuleMatchSrcFqdns(original["src_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["srcFqdns"] = transformedSrcFqdns
	}

	transformedDestFqdns, err := expandComputeNetworkFirewallPolicyRuleMatchDestFqdns(original["dest_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["destFqdns"] = transformedDestFqdns
	}

	transformedSrcRegionCodes, err := expandComputeNetworkFirewallPolicyRuleMatchSrcRegionCodes(original["src_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["srcRegionCodes"] = transformedSrcRegionCodes
	}

	transformedDestRegionCodes, err := expandComputeNetworkFirewallPolicyRuleMatchDestRegionCodes(original["dest_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["destRegionCodes"] = transformedDestRegionCodes
	}

	transformedDestThreatIntelligences, err := expandComputeNetworkFirewallPolicyRuleMatchDestThreatIntelligences(original["dest_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["destThreatIntelligences"] = transformedDestThreatIntelligences
	}

	transformedSrcThreatIntelligences, err := expandComputeNetworkFirewallPolicyRuleMatchSrcThreatIntelligences(original["src_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["srcThreatIntelligences"] = transformedSrcThreatIntelligences
	}

	return transformed, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchDestIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcNetworkScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchDestNetworkScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchLayer4Configs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpProtocol, err := expandComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsIpProtocol(original["ip_protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipProtocol"] = transformedIpProtocol
		}

		transformedPorts, err := expandComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcSecureTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedState, err := expandComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchDestAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchDestFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchDestRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchDestThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleMatchSrcThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleSecurityProfileGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleTlsInspect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleTargetServiceAccounts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleTargetSecureTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeNetworkFirewallPolicyRuleTargetSecureTagsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedState, err := expandComputeNetworkFirewallPolicyRuleTargetSecureTagsState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyRuleTargetSecureTagsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleTargetSecureTagsState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyRuleDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
