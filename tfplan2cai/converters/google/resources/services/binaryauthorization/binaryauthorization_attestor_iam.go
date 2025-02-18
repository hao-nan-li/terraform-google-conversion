// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/binaryauthorization/Attestor.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package binaryauthorization

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const BinaryAuthorizationAttestorIAMAssetType string = "binaryauthorization.googleapis.com/Attestor"

func ResourceConverterBinaryAuthorizationAttestorIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BinaryAuthorizationAttestorIAMAssetType,
		Convert:           GetBinaryAuthorizationAttestorIamPolicyCaiObject,
		MergeCreateUpdate: MergeBinaryAuthorizationAttestorIamPolicy,
	}
}

func ResourceConverterBinaryAuthorizationAttestorIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BinaryAuthorizationAttestorIAMAssetType,
		Convert:           GetBinaryAuthorizationAttestorIamBindingCaiObject,
		FetchFullResource: FetchBinaryAuthorizationAttestorIamPolicy,
		MergeCreateUpdate: MergeBinaryAuthorizationAttestorIamBinding,
		MergeDelete:       MergeBinaryAuthorizationAttestorIamBindingDelete,
	}
}

func ResourceConverterBinaryAuthorizationAttestorIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BinaryAuthorizationAttestorIAMAssetType,
		Convert:           GetBinaryAuthorizationAttestorIamMemberCaiObject,
		FetchFullResource: FetchBinaryAuthorizationAttestorIamPolicy,
		MergeCreateUpdate: MergeBinaryAuthorizationAttestorIamMember,
		MergeDelete:       MergeBinaryAuthorizationAttestorIamMemberDelete,
	}
}

func GetBinaryAuthorizationAttestorIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBinaryAuthorizationAttestorIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetBinaryAuthorizationAttestorIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBinaryAuthorizationAttestorIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetBinaryAuthorizationAttestorIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBinaryAuthorizationAttestorIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeBinaryAuthorizationAttestorIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeBinaryAuthorizationAttestorIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeBinaryAuthorizationAttestorIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeBinaryAuthorizationAttestorIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeBinaryAuthorizationAttestorIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newBinaryAuthorizationAttestorIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//binaryauthorization.googleapis.com/projects/{{project}}/attestors/{{attestor}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: BinaryAuthorizationAttestorIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchBinaryAuthorizationAttestorIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("attestor"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		BinaryAuthorizationAttestorIamUpdaterProducer,
		d,
		config,
		"//binaryauthorization.googleapis.com/projects/{{project}}/attestors/{{attestor}}",
		BinaryAuthorizationAttestorIAMAssetType,
	)
}
