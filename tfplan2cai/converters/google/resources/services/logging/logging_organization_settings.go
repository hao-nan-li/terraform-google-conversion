// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/logging/OrganizationSettings.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package logging

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const LoggingOrganizationSettingsAssetType string = "logging.googleapis.com/OrganizationSettings"

func ResourceConverterLoggingOrganizationSettings() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LoggingOrganizationSettingsAssetType,
		Convert:   GetLoggingOrganizationSettingsCaiObject,
	}
}

func GetLoggingOrganizationSettingsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//logging.googleapis.com/organizations/{{organization}}/settings")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetLoggingOrganizationSettingsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: LoggingOrganizationSettingsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/logging/v2/rest",
				DiscoveryName:        "OrganizationSettings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetLoggingOrganizationSettingsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	kmsKeyNameProp, err := expandLoggingOrganizationSettingsKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	storageLocationProp, err := expandLoggingOrganizationSettingsStorageLocation(d.Get("storage_location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("storage_location"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageLocationProp)) && (ok || !reflect.DeepEqual(v, storageLocationProp)) {
		obj["storageLocation"] = storageLocationProp
	}
	disableDefaultSinkProp, err := expandLoggingOrganizationSettingsDisableDefaultSink(d.Get("disable_default_sink"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disable_default_sink"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableDefaultSinkProp)) && (ok || !reflect.DeepEqual(v, disableDefaultSinkProp)) {
		obj["disableDefaultSink"] = disableDefaultSinkProp
	}

	return obj, nil
}

func expandLoggingOrganizationSettingsKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingOrganizationSettingsStorageLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingOrganizationSettingsDisableDefaultSink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
