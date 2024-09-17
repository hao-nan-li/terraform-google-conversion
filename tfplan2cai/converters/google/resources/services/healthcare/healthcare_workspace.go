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

package healthcare

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const HealthcareWorkspaceAssetType string = "healthcare.googleapis.com/Workspace"

func ResourceConverterHealthcareWorkspace() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: HealthcareWorkspaceAssetType,
		Convert:   GetHealthcareWorkspaceCaiObject,
	}
}

func GetHealthcareWorkspaceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//healthcare.googleapis.com/{{dataset}}/dataMapperWorkspaces/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetHealthcareWorkspaceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: HealthcareWorkspaceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/healthcare/v1beta1/rest",
				DiscoveryName:        "Workspace",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetHealthcareWorkspaceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareWorkspaceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	settingsProp, err := expandHealthcareWorkspaceSettings(d.Get("settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(settingsProp)) && (ok || !reflect.DeepEqual(v, settingsProp)) {
		obj["settings"] = settingsProp
	}
	labelsProp, err := expandHealthcareWorkspaceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandHealthcareWorkspaceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareWorkspaceSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDataProjectIds, err := expandHealthcareWorkspaceSettingsDataProjectIds(original["data_project_ids"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataProjectIds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataProjectIds"] = transformedDataProjectIds
	}

	return transformed, nil
}

func expandHealthcareWorkspaceSettingsDataProjectIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareWorkspaceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}