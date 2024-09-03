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

package netapp

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Custom function to wait for mirrorState target states
func NetappVolumeReplicationWaitForMirror(d *schema.ResourceData, meta interface{}, targetState string) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/replications/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for volume replication: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	for {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetappVolumeReplication %q", d.Id()))
		}

		log.Printf("[DEBUG] waiting for mirrorState. actual: %v, target: %v", res["mirrorState"], targetState)

		if res["mirrorState"] == targetState {
			break
		}

		time.Sleep(30 * time.Second)
		// This method can potentially run for days, e.g. when setting up a replication for a source volume
		// with dozens of TiB of data. Timeout handling yes/no?
	}

	return nil
}

const NetappVolumeReplicationAssetType string = "netapp.googleapis.com/VolumeReplication"

func ResourceConverterNetappVolumeReplication() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetappVolumeReplicationAssetType,
		Convert:   GetNetappVolumeReplicationCaiObject,
	}
}

func GetNetappVolumeReplicationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//netapp.googleapis.com/projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/replications/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetappVolumeReplicationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetappVolumeReplicationAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/netapp/v1beta1/rest",
				DiscoveryName:        "VolumeReplication",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetappVolumeReplicationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	replicationScheduleProp, err := expandNetappVolumeReplicationReplicationSchedule(d.Get("replication_schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replication_schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(replicationScheduleProp)) && (ok || !reflect.DeepEqual(v, replicationScheduleProp)) {
		obj["replicationSchedule"] = replicationScheduleProp
	}
	destinationVolumeParametersProp, err := expandNetappVolumeReplicationDestinationVolumeParameters(d.Get("destination_volume_parameters"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_volume_parameters"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationVolumeParametersProp)) && (ok || !reflect.DeepEqual(v, destinationVolumeParametersProp)) {
		obj["destinationVolumeParameters"] = destinationVolumeParametersProp
	}
	descriptionProp, err := expandNetappVolumeReplicationDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetappVolumeReplicationEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetappVolumeReplicationReplicationSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeReplicationDestinationVolumeParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStoragePool, err := expandNetappVolumeReplicationDestinationVolumeParametersStoragePool(original["storage_pool"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStoragePool); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storagePool"] = transformedStoragePool
	}

	transformedVolumeId, err := expandNetappVolumeReplicationDestinationVolumeParametersVolumeId(original["volume_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVolumeId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["volumeId"] = transformedVolumeId
	}

	transformedShareName, err := expandNetappVolumeReplicationDestinationVolumeParametersShareName(original["share_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShareName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["shareName"] = transformedShareName
	}

	transformedDescription, err := expandNetappVolumeReplicationDestinationVolumeParametersDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	return transformed, nil
}

func expandNetappVolumeReplicationDestinationVolumeParametersStoragePool(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeReplicationDestinationVolumeParametersVolumeId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeReplicationDestinationVolumeParametersShareName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeReplicationDestinationVolumeParametersDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeReplicationDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeReplicationEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
