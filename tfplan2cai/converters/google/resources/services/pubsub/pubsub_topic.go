// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/pubsub/Topic.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package pubsub

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const PubsubTopicAssetType string = "pubsub.googleapis.com/Topic"

func ResourceConverterPubsubTopic() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: PubsubTopicAssetType,
		Convert:   GetPubsubTopicCaiObject,
	}
}

func GetPubsubTopicCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/topics/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetPubsubTopicApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: PubsubTopicAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/pubsub/v1/rest",
				DiscoveryName:        "Topic",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetPubsubTopicApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandPubsubTopicName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	kmsKeyNameProp, err := expandPubsubTopicKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	messageStoragePolicyProp, err := expandPubsubTopicMessageStoragePolicy(d.Get("message_storage_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("message_storage_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(messageStoragePolicyProp)) && (ok || !reflect.DeepEqual(v, messageStoragePolicyProp)) {
		obj["messageStoragePolicy"] = messageStoragePolicyProp
	}
	schemaSettingsProp, err := expandPubsubTopicSchemaSettings(d.Get("schema_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schema_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(schemaSettingsProp)) && (ok || !reflect.DeepEqual(v, schemaSettingsProp)) {
		obj["schemaSettings"] = schemaSettingsProp
	}
	messageRetentionDurationProp, err := expandPubsubTopicMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !tpgresource.IsEmptyValue(reflect.ValueOf(messageRetentionDurationProp)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	ingestionDataSourceSettingsProp, err := expandPubsubTopicIngestionDataSourceSettings(d.Get("ingestion_data_source_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ingestion_data_source_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(ingestionDataSourceSettingsProp)) && (ok || !reflect.DeepEqual(v, ingestionDataSourceSettingsProp)) {
		obj["ingestionDataSourceSettings"] = ingestionDataSourceSettingsProp
	}
	labelsProp, err := expandPubsubTopicEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourcePubsubTopicEncoder(d, config, obj)
}

func resourcePubsubTopicEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func expandPubsubTopicName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func expandPubsubTopicKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicMessageStoragePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedPersistenceRegions, err := expandPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(original["allowed_persistence_regions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedPersistenceRegions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedPersistenceRegions"] = transformedAllowedPersistenceRegions
	}

	transformedEnforceInTransit, err := expandPubsubTopicMessageStoragePolicyEnforceInTransit(original["enforce_in_transit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnforceInTransit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enforceInTransit"] = transformedEnforceInTransit
	}

	return transformed, nil
}

func expandPubsubTopicMessageStoragePolicyAllowedPersistenceRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicMessageStoragePolicyEnforceInTransit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicSchemaSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchema, err := expandPubsubTopicSchemaSettingsSchema(original["schema"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchema); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schema"] = transformedSchema
	}

	transformedEncoding, err := expandPubsubTopicSchemaSettingsEncoding(original["encoding"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncoding); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["encoding"] = transformedEncoding
	}

	return transformed, nil
}

func expandPubsubTopicSchemaSettingsSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicSchemaSettingsEncoding(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicMessageRetentionDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAwsKinesis, err := expandPubsubTopicIngestionDataSourceSettingsAwsKinesis(original["aws_kinesis"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAwsKinesis); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["awsKinesis"] = transformedAwsKinesis
	}

	transformedCloudStorage, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorage(original["cloud_storage"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStorage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cloudStorage"] = transformedCloudStorage
	}

	transformedPlatformLogsSettings, err := expandPubsubTopicIngestionDataSourceSettingsPlatformLogsSettings(original["platform_logs_settings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPlatformLogsSettings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["platformLogsSettings"] = transformedPlatformLogsSettings
	}

	transformedAzureEventHubs, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubs(original["azure_event_hubs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAzureEventHubs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["azureEventHubs"] = transformedAzureEventHubs
	}

	transformedAwsMsk, err := expandPubsubTopicIngestionDataSourceSettingsAwsMsk(original["aws_msk"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAwsMsk); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["awsMsk"] = transformedAwsMsk
	}

	transformedConfluentCloud, err := expandPubsubTopicIngestionDataSourceSettingsConfluentCloud(original["confluent_cloud"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfluentCloud); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["confluentCloud"] = transformedConfluentCloud
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsKinesis(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStreamArn, err := expandPubsubTopicIngestionDataSourceSettingsAwsKinesisStreamArn(original["stream_arn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStreamArn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["streamArn"] = transformedStreamArn
	}

	transformedConsumerArn, err := expandPubsubTopicIngestionDataSourceSettingsAwsKinesisConsumerArn(original["consumer_arn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConsumerArn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["consumerArn"] = transformedConsumerArn
	}

	transformedAwsRoleArn, err := expandPubsubTopicIngestionDataSourceSettingsAwsKinesisAwsRoleArn(original["aws_role_arn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAwsRoleArn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["awsRoleArn"] = transformedAwsRoleArn
	}

	transformedGcpServiceAccount, err := expandPubsubTopicIngestionDataSourceSettingsAwsKinesisGcpServiceAccount(original["gcp_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpServiceAccount"] = transformedGcpServiceAccount
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsKinesisStreamArn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsKinesisConsumerArn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsKinesisAwsRoleArn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsKinesisGcpServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsCloudStorage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorageBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedTextFormat, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat(original["text_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTextFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["textFormat"] = transformedTextFormat
	}

	transformedAvroFormat, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat(original["avro_format"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["avroFormat"] = transformedAvroFormat
	}

	transformedPubsubAvroFormat, err := expandPubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat(original["pubsub_avro_format"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["pubsubAvroFormat"] = transformedPubsubAvroFormat
	}

	transformedMinimumObjectCreateTime, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorageMinimumObjectCreateTime(original["minimum_object_create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinimumObjectCreateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minimumObjectCreateTime"] = transformedMinimumObjectCreateTime
	}

	transformedMatchGlob, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorageMatchGlob(original["match_glob"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMatchGlob); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["matchGlob"] = transformedMatchGlob
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsCloudStorageBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDelimiter, err := expandPubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatDelimiter(original["delimiter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelimiter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["delimiter"] = transformedDelimiter
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatDelimiter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

func expandPubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

func expandPubsubTopicIngestionDataSourceSettingsCloudStorageMinimumObjectCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsCloudStorageMatchGlob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsPlatformLogsSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSeverity, err := expandPubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsSeverity(original["severity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severity"] = transformedSeverity
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceGroup, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsResourceGroup(original["resource_group"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceGroup"] = transformedResourceGroup
	}

	transformedNamespace, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	transformedEventHub, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsEventHub(original["event_hub"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventHub); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["eventHub"] = transformedEventHub
	}

	transformedClientId, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsClientId(original["client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientId"] = transformedClientId
	}

	transformedTenantId, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsTenantId(original["tenant_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTenantId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tenantId"] = transformedTenantId
	}

	transformedSubscriptionId, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsSubscriptionId(original["subscription_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubscriptionId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subscriptionId"] = transformedSubscriptionId
	}

	transformedGcpServiceAccount, err := expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsGcpServiceAccount(original["gcp_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpServiceAccount"] = transformedGcpServiceAccount
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsResourceGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsEventHub(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsTenantId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsSubscriptionId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAzureEventHubsGcpServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsMsk(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClusterArn, err := expandPubsubTopicIngestionDataSourceSettingsAwsMskClusterArn(original["cluster_arn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterArn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterArn"] = transformedClusterArn
	}

	transformedTopic, err := expandPubsubTopicIngestionDataSourceSettingsAwsMskTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	transformedAwsRoleArn, err := expandPubsubTopicIngestionDataSourceSettingsAwsMskAwsRoleArn(original["aws_role_arn"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAwsRoleArn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["awsRoleArn"] = transformedAwsRoleArn
	}

	transformedGcpServiceAccount, err := expandPubsubTopicIngestionDataSourceSettingsAwsMskGcpServiceAccount(original["gcp_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpServiceAccount"] = transformedGcpServiceAccount
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsMskClusterArn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsMskTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsMskAwsRoleArn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsAwsMskGcpServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsConfluentCloud(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBootstrapServer, err := expandPubsubTopicIngestionDataSourceSettingsConfluentCloudBootstrapServer(original["bootstrap_server"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBootstrapServer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bootstrapServer"] = transformedBootstrapServer
	}

	transformedClusterId, err := expandPubsubTopicIngestionDataSourceSettingsConfluentCloudClusterId(original["cluster_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterId"] = transformedClusterId
	}

	transformedTopic, err := expandPubsubTopicIngestionDataSourceSettingsConfluentCloudTopic(original["topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["topic"] = transformedTopic
	}

	transformedIdentityPoolId, err := expandPubsubTopicIngestionDataSourceSettingsConfluentCloudIdentityPoolId(original["identity_pool_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdentityPoolId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["identityPoolId"] = transformedIdentityPoolId
	}

	transformedGcpServiceAccount, err := expandPubsubTopicIngestionDataSourceSettingsConfluentCloudGcpServiceAccount(original["gcp_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpServiceAccount"] = transformedGcpServiceAccount
	}

	return transformed, nil
}

func expandPubsubTopicIngestionDataSourceSettingsConfluentCloudBootstrapServer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsConfluentCloudClusterId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsConfluentCloudTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsConfluentCloudIdentityPoolId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicIngestionDataSourceSettingsConfluentCloudGcpServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPubsubTopicEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
