// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataproc/Batch.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataproc

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

/*
 * Dataproc Batch api apends subminor version to the provided
 * version. We are suppressing this server generated subminor.
 */
func CloudDataprocBatchRuntimeConfigVersionDiffSuppressFunc(old, new string) bool {
	if old != "" && strings.HasPrefix(new, old) || (new != "" && strings.HasPrefix(old, new)) {
		return true
	}

	return old == new
}

func CloudDataprocBatchRuntimeConfigVersionDiffSuppress(_, old, new string, d *schema.ResourceData) bool {
	return CloudDataprocBatchRuntimeConfigVersionDiffSuppressFunc(old, new)
}

const DataprocBatchAssetType string = "dataproc.googleapis.com/Batch"

func ResourceConverterDataprocBatch() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataprocBatchAssetType,
		Convert:   GetDataprocBatchCaiObject,
	}
}

func GetDataprocBatchCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataproc.googleapis.com/projects/{{project}}/locations/{{location}}/batches/{{batch_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataprocBatchApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataprocBatchAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataproc/v1/rest",
				DiscoveryName:        "Batch",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataprocBatchApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	runtimeConfigProp, err := expandDataprocBatchRuntimeConfig(d.Get("runtime_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeConfigProp)) && (ok || !reflect.DeepEqual(v, runtimeConfigProp)) {
		obj["runtimeConfig"] = runtimeConfigProp
	}
	environmentConfigProp, err := expandDataprocBatchEnvironmentConfig(d.Get("environment_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("environment_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(environmentConfigProp)) && (ok || !reflect.DeepEqual(v, environmentConfigProp)) {
		obj["environmentConfig"] = environmentConfigProp
	}
	pysparkBatchProp, err := expandDataprocBatchPysparkBatch(d.Get("pyspark_batch"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pyspark_batch"); !tpgresource.IsEmptyValue(reflect.ValueOf(pysparkBatchProp)) && (ok || !reflect.DeepEqual(v, pysparkBatchProp)) {
		obj["pysparkBatch"] = pysparkBatchProp
	}
	sparkBatchProp, err := expandDataprocBatchSparkBatch(d.Get("spark_batch"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark_batch"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkBatchProp)) && (ok || !reflect.DeepEqual(v, sparkBatchProp)) {
		obj["sparkBatch"] = sparkBatchProp
	}
	sparkRBatchProp, err := expandDataprocBatchSparkRBatch(d.Get("spark_r_batch"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark_r_batch"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkRBatchProp)) && (ok || !reflect.DeepEqual(v, sparkRBatchProp)) {
		obj["sparkRBatch"] = sparkRBatchProp
	}
	sparkSqlBatchProp, err := expandDataprocBatchSparkSqlBatch(d.Get("spark_sql_batch"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spark_sql_batch"); !tpgresource.IsEmptyValue(reflect.ValueOf(sparkSqlBatchProp)) && (ok || !reflect.DeepEqual(v, sparkSqlBatchProp)) {
		obj["sparkSqlBatch"] = sparkSqlBatchProp
	}
	labelsProp, err := expandDataprocBatchEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDataprocBatchRuntimeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVersion, err := expandDataprocBatchRuntimeConfigVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedContainerImage, err := expandDataprocBatchRuntimeConfigContainerImage(original["container_image"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContainerImage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["containerImage"] = transformedContainerImage
	}

	transformedProperties, err := expandDataprocBatchRuntimeConfigProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	transformedEffectiveProperties, err := expandDataprocBatchRuntimeConfigEffectiveProperties(original["effective_properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEffectiveProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["effective_properties"] = transformedEffectiveProperties
	}

	transformedAutotuningConfig, err := expandDataprocBatchRuntimeConfigAutotuningConfig(original["autotuning_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutotuningConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autotuningConfig"] = transformedAutotuningConfig
	}

	transformedCohort, err := expandDataprocBatchRuntimeConfigCohort(original["cohort"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCohort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cohort"] = transformedCohort
	}

	return transformed, nil
}

func expandDataprocBatchRuntimeConfigVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchRuntimeConfigContainerImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchRuntimeConfigProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocBatchRuntimeConfigEffectiveProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocBatchRuntimeConfigAutotuningConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScenarios, err := expandDataprocBatchRuntimeConfigAutotuningConfigScenarios(original["scenarios"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScenarios); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scenarios"] = transformedScenarios
	}

	return transformed, nil
}

func expandDataprocBatchRuntimeConfigAutotuningConfigScenarios(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchRuntimeConfigCohort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExecutionConfig, err := expandDataprocBatchEnvironmentConfigExecutionConfig(original["execution_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExecutionConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["executionConfig"] = transformedExecutionConfig
	}

	transformedPeripheralsConfig, err := expandDataprocBatchEnvironmentConfigPeripheralsConfig(original["peripherals_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPeripheralsConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["peripheralsConfig"] = transformedPeripheralsConfig
	}

	return transformed, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccount, err := expandDataprocBatchEnvironmentConfigExecutionConfigServiceAccount(original["service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccount"] = transformedServiceAccount
	}

	transformedNetworkTags, err := expandDataprocBatchEnvironmentConfigExecutionConfigNetworkTags(original["network_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkTags"] = transformedNetworkTags
	}

	transformedKmsKey, err := expandDataprocBatchEnvironmentConfigExecutionConfigKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	transformedTtl, err := expandDataprocBatchEnvironmentConfigExecutionConfigTtl(original["ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTtl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ttl"] = transformedTtl
	}

	transformedStagingBucket, err := expandDataprocBatchEnvironmentConfigExecutionConfigStagingBucket(original["staging_bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStagingBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["stagingBucket"] = transformedStagingBucket
	}

	transformedNetworkUri, err := expandDataprocBatchEnvironmentConfigExecutionConfigNetworkUri(original["network_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkUri"] = transformedNetworkUri
	}

	transformedSubnetworkUri, err := expandDataprocBatchEnvironmentConfigExecutionConfigSubnetworkUri(original["subnetwork_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnetworkUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnetworkUri"] = transformedSubnetworkUri
	}

	return transformed, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigNetworkTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigStagingBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigNetworkUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigExecutionConfigSubnetworkUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigPeripheralsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMetastoreService, err := expandDataprocBatchEnvironmentConfigPeripheralsConfigMetastoreService(original["metastore_service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetastoreService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["metastoreService"] = transformedMetastoreService
	}

	transformedSparkHistoryServerConfig, err := expandDataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig(original["spark_history_server_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSparkHistoryServerConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sparkHistoryServerConfig"] = transformedSparkHistoryServerConfig
	}

	return transformed, nil
}

func expandDataprocBatchEnvironmentConfigPeripheralsConfigMetastoreService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDataprocCluster, err := expandDataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigDataprocCluster(original["dataproc_cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataprocCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataprocCluster"] = transformedDataprocCluster
	}

	return transformed, nil
}

func expandDataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigDataprocCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchPysparkBatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMainPythonFileUri, err := expandDataprocBatchPysparkBatchMainPythonFileUri(original["main_python_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainPythonFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainPythonFileUri"] = transformedMainPythonFileUri
	}

	transformedArgs, err := expandDataprocBatchPysparkBatchArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedPythonFileUris, err := expandDataprocBatchPysparkBatchPythonFileUris(original["python_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPythonFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pythonFileUris"] = transformedPythonFileUris
	}

	transformedJarFileUris, err := expandDataprocBatchPysparkBatchJarFileUris(original["jar_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJarFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jarFileUris"] = transformedJarFileUris
	}

	transformedFileUris, err := expandDataprocBatchPysparkBatchFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataprocBatchPysparkBatchArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	return transformed, nil
}

func expandDataprocBatchPysparkBatchMainPythonFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchPysparkBatchArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchPysparkBatchPythonFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchPysparkBatchJarFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchPysparkBatchFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchPysparkBatchArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkBatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedArgs, err := expandDataprocBatchSparkBatchArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedJarFileUris, err := expandDataprocBatchSparkBatchJarFileUris(original["jar_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJarFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jarFileUris"] = transformedJarFileUris
	}

	transformedFileUris, err := expandDataprocBatchSparkBatchFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataprocBatchSparkBatchArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	transformedMainJarFileUri, err := expandDataprocBatchSparkBatchMainJarFileUri(original["main_jar_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainJarFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainJarFileUri"] = transformedMainJarFileUri
	}

	transformedMainClass, err := expandDataprocBatchSparkBatchMainClass(original["main_class"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainClass); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainClass"] = transformedMainClass
	}

	return transformed, nil
}

func expandDataprocBatchSparkBatchArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkBatchJarFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkBatchFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkBatchArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkBatchMainJarFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkBatchMainClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkRBatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMainRFileUri, err := expandDataprocBatchSparkRBatchMainRFileUri(original["main_r_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMainRFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mainRFileUri"] = transformedMainRFileUri
	}

	transformedArgs, err := expandDataprocBatchSparkRBatchArgs(original["args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["args"] = transformedArgs
	}

	transformedFileUris, err := expandDataprocBatchSparkRBatchFileUris(original["file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fileUris"] = transformedFileUris
	}

	transformedArchiveUris, err := expandDataprocBatchSparkRBatchArchiveUris(original["archive_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedArchiveUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["archiveUris"] = transformedArchiveUris
	}

	return transformed, nil
}

func expandDataprocBatchSparkRBatchMainRFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkRBatchArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkRBatchFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkRBatchArchiveUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkSqlBatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedQueryFileUri, err := expandDataprocBatchSparkSqlBatchQueryFileUri(original["query_file_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryFileUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryFileUri"] = transformedQueryFileUri
	}

	transformedJarFileUris, err := expandDataprocBatchSparkSqlBatchJarFileUris(original["jar_file_uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedJarFileUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["jarFileUris"] = transformedJarFileUris
	}

	transformedQueryVariables, err := expandDataprocBatchSparkSqlBatchQueryVariables(original["query_variables"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQueryVariables); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["queryVariables"] = transformedQueryVariables
	}

	return transformed, nil
}

func expandDataprocBatchSparkSqlBatchQueryFileUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkSqlBatchJarFileUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocBatchSparkSqlBatchQueryVariables(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDataprocBatchEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
