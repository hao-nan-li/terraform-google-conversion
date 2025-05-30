// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryanalyticshub/ListingSubscription.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigqueryanalyticshub

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigqueryAnalyticsHubListingSubscriptionAssetType string = "analyticshub.googleapis.com/ListingSubscription"

func ResourceConverterBigqueryAnalyticsHubListingSubscription() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryAnalyticsHubListingSubscriptionAssetType,
		Convert:   GetBigqueryAnalyticsHubListingSubscriptionCaiObject,
	}
}

func GetBigqueryAnalyticsHubListingSubscriptionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//analyticshub.googleapis.com/projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryAnalyticsHubListingSubscriptionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryAnalyticsHubListingSubscriptionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/analyticshub/v1/rest",
				DiscoveryName:        "ListingSubscription",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryAnalyticsHubListingSubscriptionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	destinationDatasetProp, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDataset(d.Get("destination_dataset"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationDatasetProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetProp)) {
		obj["destinationDataset"] = destinationDatasetProp
	}

	return obj, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocation, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	transformedDatasetReference, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReference(original["dataset_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetReference); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetReference"] = transformedDatasetReference
	}

	transformedFriendlyName, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetFriendlyName(original["friendly_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFriendlyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["friendlyName"] = transformedFriendlyName
	}

	transformedDescription, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLabels, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReferenceDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReferenceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReferenceDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDatasetReferenceProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetFriendlyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubListingSubscriptionDestinationDatasetLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
