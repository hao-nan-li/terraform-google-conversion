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

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeNetworkPeeringRoutesConfigAssetType string = "compute.googleapis.com/NetworkPeeringRoutesConfig"

func ResourceConverterComputeNetworkPeeringRoutesConfig() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeNetworkPeeringRoutesConfigAssetType,
		Convert:   GetComputeNetworkPeeringRoutesConfigCaiObject,
	}
}

func GetComputeNetworkPeeringRoutesConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/networks/{{network}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeNetworkPeeringRoutesConfigApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeNetworkPeeringRoutesConfigAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "NetworkPeeringRoutesConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeNetworkPeeringRoutesConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	exportCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_custom_routes"); ok || !reflect.DeepEqual(v, exportCustomRoutesProp) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_custom_routes"); ok || !reflect.DeepEqual(v, importCustomRoutesProp) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}

	return resourceComputeNetworkPeeringRoutesConfigEncoder(d, config, obj)
}

func resourceComputeNetworkPeeringRoutesConfigEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Stick request in a networkPeering block as in
	// https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering
	newObj := make(map[string]interface{})
	newObj["networkPeering"] = obj
	return newObj, nil
}

func expandComputeNetworkPeeringRoutesConfigPeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}