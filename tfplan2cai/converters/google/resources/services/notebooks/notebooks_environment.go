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

package notebooks

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const NotebooksEnvironmentAssetType string = "notebooks.googleapis.com/Environment"

func ResourceConverterNotebooksEnvironment() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: NotebooksEnvironmentAssetType,
		Convert:   GetNotebooksEnvironmentCaiObject,
	}
}

func GetNotebooksEnvironmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//notebooks.googleapis.com/projects/{{project}}/locations/{{location}}/environments/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetNotebooksEnvironmentApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: NotebooksEnvironmentAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/notebooks/v1/rest",
				DiscoveryName:        "Environment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetNotebooksEnvironmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandNotebooksEnvironmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandNotebooksEnvironmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	postStartupScriptProp, err := expandNotebooksEnvironmentPostStartupScript(d.Get("post_startup_script"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("post_startup_script"); !tpgresource.IsEmptyValue(reflect.ValueOf(postStartupScriptProp)) && (ok || !reflect.DeepEqual(v, postStartupScriptProp)) {
		obj["postStartupScript"] = postStartupScriptProp
	}
	vmImageProp, err := expandNotebooksEnvironmentVmImage(d.Get("vm_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vm_image"); !tpgresource.IsEmptyValue(reflect.ValueOf(vmImageProp)) && (ok || !reflect.DeepEqual(v, vmImageProp)) {
		obj["vmImage"] = vmImageProp
	}
	containerImageProp, err := expandNotebooksEnvironmentContainerImage(d.Get("container_image"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("container_image"); !tpgresource.IsEmptyValue(reflect.ValueOf(containerImageProp)) && (ok || !reflect.DeepEqual(v, containerImageProp)) {
		obj["containerImage"] = containerImageProp
	}

	return obj, nil
}

func expandNotebooksEnvironmentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentPostStartupScript(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentVmImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProject, err := expandNotebooksEnvironmentVmImageProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	transformedImageName, err := expandNotebooksEnvironmentVmImageImageName(original["image_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["imageName"] = transformedImageName
	}

	transformedImageFamily, err := expandNotebooksEnvironmentVmImageImageFamily(original["image_family"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageFamily); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["imageFamily"] = transformedImageFamily
	}

	return transformed, nil
}

func expandNotebooksEnvironmentVmImageProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentVmImageImageName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentVmImageImageFamily(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentContainerImage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepository, err := expandNotebooksEnvironmentContainerImageRepository(original["repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepository); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["repository"] = transformedRepository
	}

	transformedTag, err := expandNotebooksEnvironmentContainerImageTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandNotebooksEnvironmentContainerImageRepository(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksEnvironmentContainerImageTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}