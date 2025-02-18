// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/edgenetwork/Network.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package edgenetwork

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const EdgenetworkNetworkAssetType string = "edgenetwork.googleapis.com/Network"

func ResourceConverterEdgenetworkNetwork() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: EdgenetworkNetworkAssetType,
		Convert:   GetEdgenetworkNetworkCaiObject,
	}
}

func GetEdgenetworkNetworkCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//edgenetwork.googleapis.com/projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetEdgenetworkNetworkApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: EdgenetworkNetworkAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/edgenetwork/v1/rest",
				DiscoveryName:        "Network",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetEdgenetworkNetworkApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandEdgenetworkNetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	mtuProp, err := expandEdgenetworkNetworkMtu(d.Get("mtu"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mtu"); !tpgresource.IsEmptyValue(reflect.ValueOf(mtuProp)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}
	labelsProp, err := expandEdgenetworkNetworkEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandEdgenetworkNetworkDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgenetworkNetworkMtu(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgenetworkNetworkEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
