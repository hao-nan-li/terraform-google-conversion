// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeEnvironmentAssetType string = "apigee.googleapis.com/Environment"

func ResourceConverterApigeeEnvironment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeEnvironmentAssetType,
		Convert:   GetApigeeEnvironmentCaiObject,
	}
}

func GetApigeeEnvironmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{org_id}}/environments/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeEnvironmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeEnvironmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "Environment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeEnvironmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvironmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandApigeeEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandApigeeEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	deploymentTypeProp, err := expandApigeeEnvironmentDeploymentType(d.Get("deployment_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployment_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(deploymentTypeProp)) && (ok || !reflect.DeepEqual(v, deploymentTypeProp)) {
		obj["deploymentType"] = deploymentTypeProp
	}
	apiProxyTypeProp, err := expandApigeeEnvironmentApiProxyType(d.Get("api_proxy_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("api_proxy_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiProxyTypeProp)) && (ok || !reflect.DeepEqual(v, apiProxyTypeProp)) {
		obj["apiProxyType"] = apiProxyTypeProp
	}
	nodeConfigProp, err := expandApigeeEnvironmentNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	typeProp, err := expandApigeeEnvironmentType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	forwardProxyUriProp, err := expandApigeeEnvironmentForwardProxyUri(d.Get("forward_proxy_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forward_proxy_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardProxyUriProp)) && (ok || !reflect.DeepEqual(v, forwardProxyUriProp)) {
		obj["forwardProxyUri"] = forwardProxyUriProp
	}
	propertiesProp, err := expandApigeeEnvironmentProperties(d.Get("properties"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}

	return obj, nil
}

func expandApigeeEnvironmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentDeploymentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentApiProxyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodeCount, err := expandApigeeEnvironmentNodeConfigMinNodeCount(original["min_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodeCount"] = transformedMinNodeCount
	}

	transformedMaxNodeCount, err := expandApigeeEnvironmentNodeConfigMaxNodeCount(original["max_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodeCount"] = transformedMaxNodeCount
	}

	transformedCurrentAggregateNodeCount, err := expandApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(original["current_aggregate_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentAggregateNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentAggregateNodeCount"] = transformedCurrentAggregateNodeCount
	}

	return transformed, nil
}

func expandApigeeEnvironmentNodeConfigMinNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfigMaxNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentNodeConfigCurrentAggregateNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentForwardProxyUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperty, err := expandApigeeEnvironmentPropertiesProperty(original["property"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperty); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["property"] = transformedProperty
	}

	return transformed, nil
}

func expandApigeeEnvironmentPropertiesProperty(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandApigeeEnvironmentPropertiesPropertyName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandApigeeEnvironmentPropertiesPropertyValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApigeeEnvironmentPropertiesPropertyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentPropertiesPropertyValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
