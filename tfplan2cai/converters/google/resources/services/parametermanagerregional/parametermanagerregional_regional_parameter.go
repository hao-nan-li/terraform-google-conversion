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

package parametermanagerregional

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ParameterManagerRegionalRegionalParameterAssetType string = "parametermanager.googleapis.com/RegionalParameter"

func ResourceConverterParameterManagerRegionalRegionalParameter() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ParameterManagerRegionalRegionalParameterAssetType,
		Convert:   GetParameterManagerRegionalRegionalParameterCaiObject,
	}
}

func GetParameterManagerRegionalRegionalParameterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//parametermanager.googleapis.com/projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetParameterManagerRegionalRegionalParameterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ParameterManagerRegionalRegionalParameterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/parametermanager/v1/rest",
				DiscoveryName:        "RegionalParameter",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetParameterManagerRegionalRegionalParameterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	formatProp, err := expandParameterManagerRegionalRegionalParameterFormat(d.Get("format"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("format"); !tpgresource.IsEmptyValue(reflect.ValueOf(formatProp)) && (ok || !reflect.DeepEqual(v, formatProp)) {
		obj["format"] = formatProp
	}
	labelsProp, err := expandParameterManagerRegionalRegionalParameterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandParameterManagerRegionalRegionalParameterFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParameterManagerRegionalRegionalParameterEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}