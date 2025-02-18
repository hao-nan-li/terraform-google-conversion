// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/SyncAuthorization.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeSyncAuthorizationAssetType string = "apigee.googleapis.com/SyncAuthorization"

func ResourceConverterApigeeSyncAuthorization() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeSyncAuthorizationAssetType,
		Convert:   GetApigeeSyncAuthorizationCaiObject,
	}
}

func GetApigeeSyncAuthorizationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/organizations/{{name}}:getSyncAuthorization")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeSyncAuthorizationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeSyncAuthorizationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "SyncAuthorization",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeSyncAuthorizationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	identitiesProp, err := expandApigeeSyncAuthorizationIdentities(d.Get("identities"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("identities"); ok || !reflect.DeepEqual(v, identitiesProp) {
		obj["identities"] = identitiesProp
	}
	etagProp, err := expandApigeeSyncAuthorizationEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	return obj, nil
}

func expandApigeeSyncAuthorizationIdentities(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeSyncAuthorizationEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
