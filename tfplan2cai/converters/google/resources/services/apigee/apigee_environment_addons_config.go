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

const ApigeeEnvironmentAddonsConfigAssetType string = "apigee.googleapis.com/EnvironmentAddonsConfig"

func ResourceConverterApigeeEnvironmentAddonsConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeEnvironmentAddonsConfigAssetType,
		Convert:   GetApigeeEnvironmentAddonsConfigCaiObject,
	}
}

func GetApigeeEnvironmentAddonsConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{env_id}}/addonsConfig")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeEnvironmentAddonsConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeEnvironmentAddonsConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "EnvironmentAddonsConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeEnvironmentAddonsConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	analyticsEnabledProp, err := expandApigeeEnvironmentAddonsConfigAnalyticsEnabled(d.Get("analytics_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("analytics_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(analyticsEnabledProp)) && (ok || !reflect.DeepEqual(v, analyticsEnabledProp)) {
		obj["analyticsEnabled"] = analyticsEnabledProp
	}

	return obj, nil
}

func expandApigeeEnvironmentAddonsConfigAnalyticsEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}