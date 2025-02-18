// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenter/ProjectCustomModule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterProjectCustomModuleAssetType string = "securitycenter.googleapis.com/ProjectCustomModule"

func ResourceConverterSecurityCenterProjectCustomModule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterProjectCustomModuleAssetType,
		Convert:   GetSecurityCenterProjectCustomModuleCaiObject,
	}
}

func GetSecurityCenterProjectCustomModuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/projects/{{project}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterProjectCustomModuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterProjectCustomModuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v1/rest",
				DiscoveryName:        "ProjectCustomModule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterProjectCustomModuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityCenterProjectCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enablementStateProp, err := expandSecurityCenterProjectCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablementStateProp)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	customConfigProp, err := expandSecurityCenterProjectCustomModuleCustomConfig(d.Get("custom_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(customConfigProp)) && (ok || !reflect.DeepEqual(v, customConfigProp)) {
		obj["customConfig"] = customConfigProp
	}

	return obj, nil
}

func expandSecurityCenterProjectCustomModuleDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleEnablementState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredicate, err := expandSecurityCenterProjectCustomModuleCustomConfigPredicate(original["predicate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredicate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predicate"] = transformedPredicate
	}

	transformedCustomOutput, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutput(original["custom_output"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomOutput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customOutput"] = transformedCustomOutput
	}

	transformedResourceSelector, err := expandSecurityCenterProjectCustomModuleCustomConfigResourceSelector(original["resource_selector"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceSelector); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceSelector"] = transformedResourceSelector
	}

	transformedSeverity, err := expandSecurityCenterProjectCustomModuleCustomConfigSeverity(original["severity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severity"] = transformedSeverity
	}

	transformedDescription, err := expandSecurityCenterProjectCustomModuleCustomConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedRecommendation, err := expandSecurityCenterProjectCustomModuleCustomConfigRecommendation(original["recommendation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecommendation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recommendation"] = transformedRecommendation
	}

	return transformed, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigPredicate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterProjectCustomModuleCustomConfigPredicateExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterProjectCustomModuleCustomConfigPredicateTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterProjectCustomModuleCustomConfigPredicateDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterProjectCustomModuleCustomConfigPredicateLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigPredicateExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigPredicateTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigPredicateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigPredicateLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperties, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	return transformed, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValueExpression, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpression(original["value_expression"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["valueExpression"] = transformedValueExpression
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigResourceSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceTypes, err := expandSecurityCenterProjectCustomModuleCustomConfigResourceSelectorResourceTypes(original["resource_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceTypes"] = transformedResourceTypes
	}

	return transformed, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigResourceSelectorResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterProjectCustomModuleCustomConfigRecommendation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
