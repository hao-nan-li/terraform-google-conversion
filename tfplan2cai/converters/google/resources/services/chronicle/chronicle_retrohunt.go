// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/chronicle/Retrohunt.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package chronicle

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ChronicleRetrohuntAssetType string = "{{location}}-chronicle.googleapis.com/Retrohunt"

func ResourceConverterChronicleRetrohunt() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ChronicleRetrohuntAssetType,
		Convert:   GetChronicleRetrohuntCaiObject,
	}
}

func GetChronicleRetrohuntCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-chronicle.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule}}/retrohunts/{{retrohunt}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetChronicleRetrohuntApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ChronicleRetrohuntAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-chronicle/v1beta/rest",
				DiscoveryName:        "Retrohunt",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetChronicleRetrohuntApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	processIntervalProp, err := expandChronicleRetrohuntProcessInterval(d.Get("process_interval"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("process_interval"); !tpgresource.IsEmptyValue(reflect.ValueOf(processIntervalProp)) && (ok || !reflect.DeepEqual(v, processIntervalProp)) {
		obj["processInterval"] = processIntervalProp
	}
	retrohuntProp, err := expandChronicleRetrohuntRetrohunt(d.Get("retrohunt"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retrohunt"); !tpgresource.IsEmptyValue(reflect.ValueOf(retrohuntProp)) && (ok || !reflect.DeepEqual(v, retrohuntProp)) {
		obj["retrohunt"] = retrohuntProp
	}

	return obj, nil
}

func expandChronicleRetrohuntProcessInterval(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandChronicleRetrohuntProcessIntervalStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandChronicleRetrohuntProcessIntervalEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandChronicleRetrohuntProcessIntervalStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleRetrohuntProcessIntervalEndTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleRetrohuntRetrohunt(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
