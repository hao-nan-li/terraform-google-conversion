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

package dialogflow

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DialogflowEntityTypeAssetType string = "dialogflow.googleapis.com/EntityType"

func ResourceConverterDialogflowEntityType() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: DialogflowEntityTypeAssetType,
		Convert:   GetDialogflowEntityTypeCaiObject,
	}
}

func GetDialogflowEntityTypeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//dialogflow.googleapis.com/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetDialogflowEntityTypeApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: DialogflowEntityTypeAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflow/v2/rest",
				DiscoveryName:        "EntityType",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetDialogflowEntityTypeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowEntityTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	kindProp, err := expandDialogflowEntityTypeKind(d.Get("kind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kind"); !tpgresource.IsEmptyValue(reflect.ValueOf(kindProp)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	enableFuzzyExtractionProp, err := expandDialogflowEntityTypeEnableFuzzyExtraction(d.Get("enable_fuzzy_extraction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_fuzzy_extraction"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableFuzzyExtractionProp)) && (ok || !reflect.DeepEqual(v, enableFuzzyExtractionProp)) {
		obj["enableFuzzyExtraction"] = enableFuzzyExtractionProp
	}
	entitiesProp, err := expandDialogflowEntityTypeEntities(d.Get("entities"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entities"); !tpgresource.IsEmptyValue(reflect.ValueOf(entitiesProp)) && (ok || !reflect.DeepEqual(v, entitiesProp)) {
		obj["entities"] = entitiesProp
	}

	return obj, nil
}

func expandDialogflowEntityTypeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeEnableFuzzyExtraction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeEntities(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValue, err := expandDialogflowEntityTypeEntitiesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedSynonyms, err := expandDialogflowEntityTypeEntitiesSynonyms(original["synonyms"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSynonyms); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["synonyms"] = transformedSynonyms
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowEntityTypeEntitiesValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowEntityTypeEntitiesSynonyms(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}