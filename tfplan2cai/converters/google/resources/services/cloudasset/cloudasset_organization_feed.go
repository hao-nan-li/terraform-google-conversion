// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudasset/OrganizationFeed.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudasset

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CloudAssetOrganizationFeedAssetType string = "cloudasset.googleapis.com/OrganizationFeed"

func ResourceConverterCloudAssetOrganizationFeed() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudAssetOrganizationFeedAssetType,
		Convert:   GetCloudAssetOrganizationFeedCaiObject,
	}
}

func GetCloudAssetOrganizationFeedCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudasset.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudAssetOrganizationFeedApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CloudAssetOrganizationFeedAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudasset/v1/rest",
				DiscoveryName:        "OrganizationFeed",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudAssetOrganizationFeedApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	assetNamesProp, err := expandCloudAssetOrganizationFeedAssetNames(d.Get("asset_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetNamesProp)) && (ok || !reflect.DeepEqual(v, assetNamesProp)) {
		obj["assetNames"] = assetNamesProp
	}
	assetTypesProp, err := expandCloudAssetOrganizationFeedAssetTypes(d.Get("asset_types"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("asset_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(assetTypesProp)) && (ok || !reflect.DeepEqual(v, assetTypesProp)) {
		obj["assetTypes"] = assetTypesProp
	}
	contentTypeProp, err := expandCloudAssetOrganizationFeedContentType(d.Get("content_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("content_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(contentTypeProp)) && (ok || !reflect.DeepEqual(v, contentTypeProp)) {
		obj["contentType"] = contentTypeProp
	}
	feedOutputConfigProp, err := expandCloudAssetOrganizationFeedFeedOutputConfig(d.Get("feed_output_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("feed_output_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(feedOutputConfigProp)) && (ok || !reflect.DeepEqual(v, feedOutputConfigProp)) {
		obj["feedOutputConfig"] = feedOutputConfigProp
	}
	conditionProp, err := expandCloudAssetOrganizationFeedCondition(d.Get("condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}

	return resourceCloudAssetOrganizationFeedEncoder(d, config, obj)
}

func resourceCloudAssetOrganizationFeedEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Remove the "folders/" prefix from the folder ID
	if folder, ok := d.GetOkExists("folder"); ok {
		if err := d.Set("folder_id", strings.TrimPrefix(folder.(string), "folders/")); err != nil {
			return nil, fmt.Errorf("Error setting folder_id: %s", err)
		}
	}
	// The feed object must be under the "feed" attribute on the request.
	newObj := make(map[string]interface{})
	newObj["feed"] = obj
	return newObj, nil
}

func expandCloudAssetOrganizationFeedAssetNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedAssetTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedContentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubDestination, err := expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(original["pubsub_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubDestination"] = transformedPubsubDestination
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTopic, err := expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedFeedOutputConfigPubsubDestinationTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandCloudAssetOrganizationFeedConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandCloudAssetOrganizationFeedConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandCloudAssetOrganizationFeedConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandCloudAssetOrganizationFeedConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandCloudAssetOrganizationFeedConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudAssetOrganizationFeedConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
