// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/orgpolicy/Policy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package orgpolicy

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func resourceOrgpolicyPolicyRulesConditionExpressionDiffSuppress(_, old, new string, d *schema.ResourceData) bool {
	oldReplaced := strings.ReplaceAll(strings.ReplaceAll(old, "Labels", "TagId"), "label", "tag")
	newReplaced := strings.ReplaceAll(strings.ReplaceAll(new, "Labels", "TagId"), "label", "tag")
	return oldReplaced == newReplaced
}

const OrgPolicyPolicyAssetType string = "orgpolicy.googleapis.com/Policy"

func ResourceConverterOrgPolicyPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: OrgPolicyPolicyAssetType,
		Convert:   GetOrgPolicyPolicyCaiObject,
	}
}

func GetOrgPolicyPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//orgpolicy.googleapis.com/{{parent}}/policies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetOrgPolicyPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: OrgPolicyPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/orgpolicy/v2/rest",
				DiscoveryName:        "Policy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetOrgPolicyPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandOrgPolicyPolicyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	specProp, err := expandOrgPolicyPolicySpec(d.Get("spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}
	dryRunSpecProp, err := expandOrgPolicyPolicyDryRunSpec(d.Get("dry_run_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dry_run_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(dryRunSpecProp)) && (ok || !reflect.DeepEqual(v, dryRunSpecProp)) {
		obj["dryRunSpec"] = dryRunSpecProp
	}

	return resourceOrgPolicyPolicyEncoder(d, config, obj)
}

func resourceOrgPolicyPolicyEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	name := d.Get("name").(string)
	d.Set("name", tpgresource.GetResourceNameFromSelfLink(name))
	return obj, nil
}

func expandOrgPolicyPolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEtag, err := expandOrgPolicyPolicySpecEtag(original["etag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEtag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["etag"] = transformedEtag
	}

	transformedUpdateTime, err := expandOrgPolicyPolicySpecUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	transformedRules, err := expandOrgPolicyPolicySpecRules(original["rules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rules"] = transformedRules
	}

	transformedInheritFromParent, err := expandOrgPolicyPolicySpecInheritFromParent(original["inherit_from_parent"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["inheritFromParent"] = transformedInheritFromParent
	}

	transformedReset, err := expandOrgPolicyPolicySpecReset(original["reset"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["reset"] = transformedReset
	}

	return transformed, nil
}

func expandOrgPolicyPolicySpecEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValues, err := expandOrgPolicyPolicySpecRulesValues(original["values"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["values"] = transformedValues
		}

		transformedAllowAll, err := expandOrgPolicyPolicySpecRulesAllowAll(original["allow_all"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["allowAll"] = transformedAllowAll
		}

		transformedDenyAll, err := expandOrgPolicyPolicySpecRulesDenyAll(original["deny_all"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["denyAll"] = transformedDenyAll
		}

		transformedEnforce, err := expandOrgPolicyPolicySpecRulesEnforce(original["enforce"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["enforce"] = transformedEnforce
		}

		transformedParameters, err := expandOrgPolicyPolicySpecRulesParameters(original["parameters"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["parameters"] = transformedParameters
		}

		transformedCondition, err := expandOrgPolicyPolicySpecRulesCondition(original["condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["condition"] = transformedCondition
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandOrgPolicyPolicySpecRulesValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedValues, err := expandOrgPolicyPolicySpecRulesValuesAllowedValues(original["allowed_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedValues"] = transformedAllowedValues
	}

	transformedDeniedValues, err := expandOrgPolicyPolicySpecRulesValuesDeniedValues(original["denied_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeniedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deniedValues"] = transformedDeniedValues
	}

	return transformed, nil
}

func expandOrgPolicyPolicySpecRulesValuesAllowedValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecRulesValuesDeniedValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecRulesAllowAll(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	b, err := strconv.ParseBool(v.(string))
	if err != nil {
		return nil, nil
	}
	return b, nil
}

func expandOrgPolicyPolicySpecRulesDenyAll(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	b, err := strconv.ParseBool(v.(string))
	if err != nil {
		return nil, nil
	}
	return b, nil
}

func expandOrgPolicyPolicySpecRulesEnforce(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	b, err := strconv.ParseBool(v.(string))
	if err != nil {
		return nil, nil
	}
	return b, nil
}

func expandOrgPolicyPolicySpecRulesParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandOrgPolicyPolicySpecRulesCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandOrgPolicyPolicySpecRulesConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandOrgPolicyPolicySpecRulesConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandOrgPolicyPolicySpecRulesConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandOrgPolicyPolicySpecRulesConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandOrgPolicyPolicySpecRulesConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecRulesConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecRulesConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecRulesConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecInheritFromParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicySpecReset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEtag, err := expandOrgPolicyPolicyDryRunSpecEtag(original["etag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEtag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["etag"] = transformedEtag
	}

	transformedUpdateTime, err := expandOrgPolicyPolicyDryRunSpecUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	transformedRules, err := expandOrgPolicyPolicyDryRunSpecRules(original["rules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rules"] = transformedRules
	}

	transformedInheritFromParent, err := expandOrgPolicyPolicyDryRunSpecInheritFromParent(original["inherit_from_parent"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["inheritFromParent"] = transformedInheritFromParent
	}

	transformedReset, err := expandOrgPolicyPolicyDryRunSpecReset(original["reset"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["reset"] = transformedReset
	}

	return transformed, nil
}

func expandOrgPolicyPolicyDryRunSpecEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValues, err := expandOrgPolicyPolicyDryRunSpecRulesValues(original["values"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["values"] = transformedValues
		}

		transformedAllowAll, err := expandOrgPolicyPolicyDryRunSpecRulesAllowAll(original["allow_all"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["allowAll"] = transformedAllowAll
		}

		transformedDenyAll, err := expandOrgPolicyPolicyDryRunSpecRulesDenyAll(original["deny_all"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["denyAll"] = transformedDenyAll
		}

		transformedEnforce, err := expandOrgPolicyPolicyDryRunSpecRulesEnforce(original["enforce"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["enforce"] = transformedEnforce
		}

		transformedParameters, err := expandOrgPolicyPolicyDryRunSpecRulesParameters(original["parameters"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["parameters"] = transformedParameters
		}

		transformedCondition, err := expandOrgPolicyPolicyDryRunSpecRulesCondition(original["condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["condition"] = transformedCondition
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedValues, err := expandOrgPolicyPolicyDryRunSpecRulesValuesAllowedValues(original["allowed_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedValues"] = transformedAllowedValues
	}

	transformedDeniedValues, err := expandOrgPolicyPolicyDryRunSpecRulesValuesDeniedValues(original["denied_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeniedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deniedValues"] = transformedDeniedValues
	}

	return transformed, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesValuesAllowedValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesValuesDeniedValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesAllowAll(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	b, err := strconv.ParseBool(v.(string))
	if err != nil {
		return nil, nil
	}
	return b, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesDenyAll(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	b, err := strconv.ParseBool(v.(string))
	if err != nil {
		return nil, nil
	}
	return b, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesEnforce(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	b, err := strconv.ParseBool(v.(string))
	if err != nil {
		return nil, nil
	}
	return b, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandOrgPolicyPolicyDryRunSpecRulesConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandOrgPolicyPolicyDryRunSpecRulesConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandOrgPolicyPolicyDryRunSpecRulesConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandOrgPolicyPolicyDryRunSpecRulesConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecRulesConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecInheritFromParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyPolicyDryRunSpecReset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
