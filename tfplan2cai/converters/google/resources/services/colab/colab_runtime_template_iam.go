// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/colab/RuntimeTemplate.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package colab

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ColabRuntimeTemplateIAMAssetType string = "aiplatform.googleapis.com/RuntimeTemplate"

func ResourceConverterColabRuntimeTemplateIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ColabRuntimeTemplateIAMAssetType,
		Convert:           GetColabRuntimeTemplateIamPolicyCaiObject,
		MergeCreateUpdate: MergeColabRuntimeTemplateIamPolicy,
	}
}

func ResourceConverterColabRuntimeTemplateIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ColabRuntimeTemplateIAMAssetType,
		Convert:           GetColabRuntimeTemplateIamBindingCaiObject,
		FetchFullResource: FetchColabRuntimeTemplateIamPolicy,
		MergeCreateUpdate: MergeColabRuntimeTemplateIamBinding,
		MergeDelete:       MergeColabRuntimeTemplateIamBindingDelete,
	}
}

func ResourceConverterColabRuntimeTemplateIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ColabRuntimeTemplateIAMAssetType,
		Convert:           GetColabRuntimeTemplateIamMemberCaiObject,
		FetchFullResource: FetchColabRuntimeTemplateIamPolicy,
		MergeCreateUpdate: MergeColabRuntimeTemplateIamMember,
		MergeDelete:       MergeColabRuntimeTemplateIamMemberDelete,
	}
}

func GetColabRuntimeTemplateIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newColabRuntimeTemplateIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetColabRuntimeTemplateIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newColabRuntimeTemplateIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetColabRuntimeTemplateIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newColabRuntimeTemplateIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeColabRuntimeTemplateIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeColabRuntimeTemplateIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeColabRuntimeTemplateIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeColabRuntimeTemplateIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeColabRuntimeTemplateIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newColabRuntimeTemplateIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/projects/{{project}}/locations/{{location}}/notebookRuntimeTemplates/{{runtime_template}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ColabRuntimeTemplateIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchColabRuntimeTemplateIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("runtime_template"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ColabRuntimeTemplateIamUpdaterProducer,
		d,
		config,
		"//aiplatform.googleapis.com/projects/{{project}}/locations/{{location}}/notebookRuntimeTemplates/{{runtime_template}}",
		ColabRuntimeTemplateIAMAssetType,
	)
}
