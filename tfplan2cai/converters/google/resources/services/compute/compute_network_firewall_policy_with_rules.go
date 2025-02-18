// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/NetworkFirewallPolicyWithRules.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func networkFirewallPolicyWithRulesConvertPriorityToInt(v interface{}) (int64, error) {
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal, nil
		}
	}

	if intVal, ok := v.(int64); ok {
		return intVal, nil
	}

	if floatVal, ok := v.(float64); ok {
		intVal := int64(floatVal)
		return intVal, nil
	}

	return 0, fmt.Errorf("Incorrect rule priority: %s. Priority must be a number", v)
}

func networkFirewallPolicyWithRulesIsPredefinedRule(rule map[string]interface{}) (bool, error) {
	// Priorities from 2147483548 to 2147483647 are reserved and cannot be modified by the user.
	const ReservedPriorityStart = 2147483548

	priority := rule["priority"]
	priorityInt, err := networkFirewallPolicyWithRulesConvertPriorityToInt(priority)

	if err != nil {
		return false, err
	}

	return priorityInt >= ReservedPriorityStart, nil

}

func networkFirewallPolicyWithRulesSplitPredefinedRules(allRules []interface{}) ([]interface{}, []interface{}, error) {
	predefinedRules := make([]interface{}, 0)
	rules := make([]interface{}, 0)

	for _, rule := range allRules {
		isPredefined, err := networkFirewallPolicyWithRulesIsPredefinedRule(rule.(map[string]interface{}))
		if err != nil {
			return nil, nil, err
		}

		if isPredefined {
			predefinedRules = append(predefinedRules, rule)
		} else {
			rules = append(rules, rule)
		}
	}

	return rules, predefinedRules, nil
}

const ComputeNetworkFirewallPolicyWithRulesAssetType string = "compute.googleapis.com/NetworkFirewallPolicyWithRules"

func ResourceConverterComputeNetworkFirewallPolicyWithRules() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNetworkFirewallPolicyWithRulesAssetType,
		Convert:   GetComputeNetworkFirewallPolicyWithRulesCaiObject,
	}
}

func GetComputeNetworkFirewallPolicyWithRulesCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/firewallPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNetworkFirewallPolicyWithRulesApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNetworkFirewallPolicyWithRulesAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NetworkFirewallPolicyWithRules",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNetworkFirewallPolicyWithRulesApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkFirewallPolicyWithRulesName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeNetworkFirewallPolicyWithRulesDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	rulesProp, err := expandComputeNetworkFirewallPolicyWithRulesRule(d.Get("rule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}
	fingerprintProp, err := expandComputeNetworkFirewallPolicyWithRulesFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	return resourceComputeNetworkFirewallPolicyWithRulesEncoder(d, config, obj)
}

func resourceComputeNetworkFirewallPolicyWithRulesEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "rules") // Rules are not supported in the create API
	return obj, nil
}

func expandComputeNetworkFirewallPolicyWithRulesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeNetworkFirewallPolicyWithRulesRuleDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedRuleName, err := expandComputeNetworkFirewallPolicyWithRulesRuleRuleName(original["rule_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRuleName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ruleName"] = transformedRuleName
		}

		transformedPriority, err := expandComputeNetworkFirewallPolicyWithRulesRulePriority(original["priority"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPriority); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["priority"] = transformedPriority
		}

		transformedMatch, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatch(original["match"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMatch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["match"] = transformedMatch
		}

		transformedTargetSecureTag, err := expandComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTag(original["target_secure_tag"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetSecureTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["targetSecureTags"] = transformedTargetSecureTag
		}

		transformedAction, err := expandComputeNetworkFirewallPolicyWithRulesRuleAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedDirection, err := expandComputeNetworkFirewallPolicyWithRulesRuleDirection(original["direction"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDirection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["direction"] = transformedDirection
		}

		transformedEnableLogging, err := expandComputeNetworkFirewallPolicyWithRulesRuleEnableLogging(original["enable_logging"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["enableLogging"] = transformedEnableLogging
		}

		transformedTargetServiceAccounts, err := expandComputeNetworkFirewallPolicyWithRulesRuleTargetServiceAccounts(original["target_service_accounts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetServiceAccounts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["targetServiceAccounts"] = transformedTargetServiceAccounts
		}

		transformedSecurityProfileGroup, err := expandComputeNetworkFirewallPolicyWithRulesRuleSecurityProfileGroup(original["security_profile_group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecurityProfileGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["securityProfileGroup"] = transformedSecurityProfileGroup
		}

		transformedTlsInspect, err := expandComputeNetworkFirewallPolicyWithRulesRuleTlsInspect(original["tls_inspect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTlsInspect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["tlsInspect"] = transformedTlsInspect
		}

		transformedDisabled, err := expandComputeNetworkFirewallPolicyWithRulesRuleDisabled(original["disabled"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["disabled"] = transformedDisabled
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleRuleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedSrcAddressGroups, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcAddressGroups(original["src_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcAddressGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcAddressGroups"] = transformedSrcAddressGroups
	}

	transformedDestAddressGroups, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestAddressGroups(original["dest_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestAddressGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destAddressGroups"] = transformedDestAddressGroups
	}

	transformedSrcFqdns, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcFqdns(original["src_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcFqdns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcFqdns"] = transformedSrcFqdns
	}

	transformedDestFqdns, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestFqdns(original["dest_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestFqdns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destFqdns"] = transformedDestFqdns
	}

	transformedSrcRegionCodes, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcRegionCodes(original["src_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcRegionCodes"] = transformedSrcRegionCodes
	}

	transformedDestRegionCodes, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestRegionCodes(original["dest_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destRegionCodes"] = transformedDestRegionCodes
	}

	transformedSrcNetworkScope, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcNetworkScope(original["src_network_scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcNetworkScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcNetworkScope"] = transformedSrcNetworkScope
	}

	transformedSrcNetworks, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcNetworks(original["src_networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcNetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcNetworks"] = transformedSrcNetworks
	}

	transformedDestNetworkScope, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestNetworkScope(original["dest_network_scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestNetworkScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destNetworkScope"] = transformedDestNetworkScope
	}

	transformedSrcThreatIntelligences, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcThreatIntelligences(original["src_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcThreatIntelligences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcThreatIntelligences"] = transformedSrcThreatIntelligences
	}

	transformedDestThreatIntelligences, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestThreatIntelligences(original["dest_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestThreatIntelligences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destThreatIntelligences"] = transformedDestThreatIntelligences
	}

	transformedLayer4Config, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchLayer4Config(original["layer4_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLayer4Config); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["layer4Configs"] = transformedLayer4Config
	}

	transformedSrcSecureTag, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTag(original["src_secure_tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcSecureTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcSecureTags"] = transformedSrcSecureTag
	}

	return transformed, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcNetworkScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestNetworkScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchDestThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchLayer4Config(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpProtocol, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigIpProtocol(original["ip_protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipProtocol"] = transformedIpProtocol
		}

		transformedPorts, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedState, err := expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedState, err := expandComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["state"] = transformedState
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleTargetServiceAccounts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleSecurityProfileGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleTlsInspect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesRuleDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyWithRulesFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
