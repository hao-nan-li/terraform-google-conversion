// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firestore/Index.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firestore

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

/*
 * FirestoreIndex api apends __name__ as an item to the
 * end of the fields list if not present. We are suppressing
 * this server generated field.
 */
func FirestoreIFieldsDiffSuppressFunc(k, old, new string, d tpgresource.TerraformResourceDataChange) bool {
	kLength := "fields.#"
	oldLength, newLength := d.GetChange(kLength)
	oldInt, ok := oldLength.(int)
	if !ok {
		return false
	}
	newInt, ok := newLength.(int)
	if !ok {
		return false
	}

	if oldInt == newInt+1 {
		kold := fmt.Sprintf("fields.%v.field_path", oldInt-1)
		knew := fmt.Sprintf("fields.%v.field_path", newInt-1)

		oldLastIndexName, _ := d.GetChange(kold)
		_, newLastIndexName := d.GetChange(knew)
		if oldLastIndexName == "__name__" && newLastIndexName != "__name__" {
			oldBase := fmt.Sprintf("fields.%v", oldInt-1)
			if strings.HasPrefix(k, oldBase) || k == kLength {
				return true
			}
		}
	}
	return false
}

func firestoreIFieldsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	return FirestoreIFieldsDiffSuppressFunc(k, old, new, d)
}

const FirestoreIndexAssetType string = "firestore.googleapis.com/Index"

func ResourceConverterFirestoreIndex() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirestoreIndexAssetType,
		Convert:   GetFirestoreIndexCaiObject,
	}
}

func GetFirestoreIndexCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firestore.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirestoreIndexApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirestoreIndexAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firestore/v1/rest",
				DiscoveryName:        "Index",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirestoreIndexApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	databaseProp, err := expandFirestoreIndexDatabase(d.Get("database"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseProp)) && (ok || !reflect.DeepEqual(v, databaseProp)) {
		obj["database"] = databaseProp
	}
	collectionProp, err := expandFirestoreIndexCollection(d.Get("collection"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("collection"); !tpgresource.IsEmptyValue(reflect.ValueOf(collectionProp)) && (ok || !reflect.DeepEqual(v, collectionProp)) {
		obj["collection"] = collectionProp
	}
	queryScopeProp, err := expandFirestoreIndexQueryScope(d.Get("query_scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("query_scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(queryScopeProp)) && (ok || !reflect.DeepEqual(v, queryScopeProp)) {
		obj["queryScope"] = queryScopeProp
	}
	apiScopeProp, err := expandFirestoreIndexApiScope(d.Get("api_scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("api_scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiScopeProp)) && (ok || !reflect.DeepEqual(v, apiScopeProp)) {
		obj["apiScope"] = apiScopeProp
	}
	fieldsProp, err := expandFirestoreIndexFields(d.Get("fields"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	return resourceFirestoreIndexEncoder(d, config, obj)
}

func resourceFirestoreIndexEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// We've added project / database / collection as split fields of the name, but
	// the API doesn't expect them.  Make sure we remove them from any requests.

	delete(obj, "project")
	delete(obj, "database")
	delete(obj, "collection")
	return obj, nil
}

func expandFirestoreIndexDatabase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexCollection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexQueryScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexApiScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFieldPath, err := expandFirestoreIndexFieldsFieldPath(original["field_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFieldPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["fieldPath"] = transformedFieldPath
		}

		transformedOrder, err := expandFirestoreIndexFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedArrayConfig, err := expandFirestoreIndexFieldsArrayConfig(original["array_config"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArrayConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["arrayConfig"] = transformedArrayConfig
		}

		transformedVectorConfig, err := expandFirestoreIndexFieldsVectorConfig(original["vector_config"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVectorConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["vectorConfig"] = transformedVectorConfig
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirestoreIndexFieldsFieldPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsOrder(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsArrayConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsVectorConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDimension, err := expandFirestoreIndexFieldsVectorConfigDimension(original["dimension"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDimension); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dimension"] = transformedDimension
	}

	transformedFlat, err := expandFirestoreIndexFieldsVectorConfigFlat(original["flat"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["flat"] = transformedFlat
	}

	return transformed, nil
}

func expandFirestoreIndexFieldsVectorConfigDimension(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsVectorConfigFlat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}
