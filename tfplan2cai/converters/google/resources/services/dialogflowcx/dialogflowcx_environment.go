// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dialogflowcx/Environment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dialogflowcx

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXEnvironmentAssetType string = "{{location}}-dialogflow.googleapis.com/Environment"

func ResourceConverterDialogflowCXEnvironment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowCXEnvironmentAssetType,
		Convert:   GetDialogflowCXEnvironmentCaiObject,
	}
}

func GetDialogflowCXEnvironmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/{{parent}}/environments/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowCXEnvironmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowCXEnvironmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "Environment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowCXEnvironmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDialogflowCXEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	versionConfigsProp, err := expandDialogflowCXEnvironmentVersionConfigs(d.Get("version_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionConfigsProp)) && (ok || !reflect.DeepEqual(v, versionConfigsProp)) {
		obj["versionConfigs"] = versionConfigsProp
	}

	return obj, nil
}

func expandDialogflowCXEnvironmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEnvironmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXEnvironmentVersionConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVersion, err := expandDialogflowCXEnvironmentVersionConfigsVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXEnvironmentVersionConfigsVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
