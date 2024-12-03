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

package gemini

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const GeminiRepositoryGroupIAMAssetType string = "cloudaicompanion.googleapis.com/RepositoryGroup"

func ResourceConverterGeminiRepositoryGroupIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GeminiRepositoryGroupIAMAssetType,
		Convert:           GetGeminiRepositoryGroupIamPolicyCaiObject,
		MergeCreateUpdate: MergeGeminiRepositoryGroupIamPolicy,
	}
}

func ResourceConverterGeminiRepositoryGroupIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GeminiRepositoryGroupIAMAssetType,
		Convert:           GetGeminiRepositoryGroupIamBindingCaiObject,
		FetchFullResource: FetchGeminiRepositoryGroupIamPolicy,
		MergeCreateUpdate: MergeGeminiRepositoryGroupIamBinding,
		MergeDelete:       MergeGeminiRepositoryGroupIamBindingDelete,
	}
}

func ResourceConverterGeminiRepositoryGroupIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GeminiRepositoryGroupIAMAssetType,
		Convert:           GetGeminiRepositoryGroupIamMemberCaiObject,
		FetchFullResource: FetchGeminiRepositoryGroupIamPolicy,
		MergeCreateUpdate: MergeGeminiRepositoryGroupIamMember,
		MergeDelete:       MergeGeminiRepositoryGroupIamMemberDelete,
	}
}

func GetGeminiRepositoryGroupIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGeminiRepositoryGroupIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetGeminiRepositoryGroupIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGeminiRepositoryGroupIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetGeminiRepositoryGroupIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGeminiRepositoryGroupIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeGeminiRepositoryGroupIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeGeminiRepositoryGroupIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeGeminiRepositoryGroupIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeGeminiRepositoryGroupIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeGeminiRepositoryGroupIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newGeminiRepositoryGroupIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//cloudaicompanion.googleapis.com/projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index}}/repositoryGroups/{{repository_group_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: GeminiRepositoryGroupIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchGeminiRepositoryGroupIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("code_repository_index"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("repository_group_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		GeminiRepositoryGroupIamUpdaterProducer,
		d,
		config,
		"//cloudaicompanion.googleapis.com/projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index}}/repositoryGroups/{{repository_group_id}}",
		GeminiRepositoryGroupIAMAssetType,
	)
}