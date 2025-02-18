// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accessapproval/ProjectSettings.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package accessapproval

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AccessApprovalProjectSettingsAssetType string = "accessapproval.googleapis.com/ProjectSettings"

func ResourceConverterAccessApprovalProjectSettings() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AccessApprovalProjectSettingsAssetType,
		Convert:   GetAccessApprovalProjectSettingsCaiObject,
	}
}

func GetAccessApprovalProjectSettingsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//accessapproval.googleapis.com/projects/{{project_id}}/accessApprovalSettings")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAccessApprovalProjectSettingsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AccessApprovalProjectSettingsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accessapproval/v1/rest",
				DiscoveryName:        "ProjectSettings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAccessApprovalProjectSettingsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalProjectSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_emails"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalProjectSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enrolled_services"); !tpgresource.IsEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalProjectSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("active_key_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(activeKeyVersionProp)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}
	projectProp, err := expandAccessApprovalProjectSettingsProject(d.Get("project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("project"); !tpgresource.IsEmptyValue(reflect.ValueOf(projectProp)) && (ok || !reflect.DeepEqual(v, projectProp)) {
		obj["project"] = projectProp
	}

	return obj, nil
}

func expandAccessApprovalProjectSettingsNotificationEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalProjectSettingsEnrolledServices(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalProjectSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalProjectSettingsEnrolledServicesCloudProduct(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsEnrolledServicesEnrollmentLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsActiveKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalProjectSettingsProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
