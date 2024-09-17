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

package secretmanagerregional

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const SecretManagerRegionalRegionalSecretIAMAssetType string = "secretmanager.{{location}}.rep.googleapis.com/RegionalSecret"

func ResourceConverterSecretManagerRegionalRegionalSecretIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecretManagerRegionalRegionalSecretIAMAssetType,
		Convert:           GetSecretManagerRegionalRegionalSecretIamPolicyCaiObject,
		MergeCreateUpdate: MergeSecretManagerRegionalRegionalSecretIamPolicy,
	}
}

func ResourceConverterSecretManagerRegionalRegionalSecretIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecretManagerRegionalRegionalSecretIAMAssetType,
		Convert:           GetSecretManagerRegionalRegionalSecretIamBindingCaiObject,
		FetchFullResource: FetchSecretManagerRegionalRegionalSecretIamPolicy,
		MergeCreateUpdate: MergeSecretManagerRegionalRegionalSecretIamBinding,
		MergeDelete:       MergeSecretManagerRegionalRegionalSecretIamBindingDelete,
	}
}

func ResourceConverterSecretManagerRegionalRegionalSecretIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecretManagerRegionalRegionalSecretIAMAssetType,
		Convert:           GetSecretManagerRegionalRegionalSecretIamMemberCaiObject,
		FetchFullResource: FetchSecretManagerRegionalRegionalSecretIamPolicy,
		MergeCreateUpdate: MergeSecretManagerRegionalRegionalSecretIamMember,
		MergeDelete:       MergeSecretManagerRegionalRegionalSecretIamMemberDelete,
	}
}

func GetSecretManagerRegionalRegionalSecretIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecretManagerRegionalRegionalSecretIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetSecretManagerRegionalRegionalSecretIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecretManagerRegionalRegionalSecretIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetSecretManagerRegionalRegionalSecretIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecretManagerRegionalRegionalSecretIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeSecretManagerRegionalRegionalSecretIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeSecretManagerRegionalRegionalSecretIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeSecretManagerRegionalRegionalSecretIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeSecretManagerRegionalRegionalSecretIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeSecretManagerRegionalRegionalSecretIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newSecretManagerRegionalRegionalSecretIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//secretmanager.{{location}}.rep.googleapis.com/projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: SecretManagerRegionalRegionalSecretIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchSecretManagerRegionalRegionalSecretIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("secret_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		SecretManagerRegionalRegionalSecretIamUpdaterProducer,
		d,
		config,
		"//secretmanager.{{location}}.rep.googleapis.com/projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}",
		SecretManagerRegionalRegionalSecretIAMAssetType,
	)
}