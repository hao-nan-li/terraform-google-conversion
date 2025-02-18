// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/clouddeploy/CustomTargetType.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package clouddeploy

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ClouddeployCustomTargetTypeIAMAssetType string = "clouddeploy.googleapis.com/CustomTargetType"

func ResourceConverterClouddeployCustomTargetTypeIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ClouddeployCustomTargetTypeIAMAssetType,
		Convert:           GetClouddeployCustomTargetTypeIamPolicyCaiObject,
		MergeCreateUpdate: MergeClouddeployCustomTargetTypeIamPolicy,
	}
}

func ResourceConverterClouddeployCustomTargetTypeIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ClouddeployCustomTargetTypeIAMAssetType,
		Convert:           GetClouddeployCustomTargetTypeIamBindingCaiObject,
		FetchFullResource: FetchClouddeployCustomTargetTypeIamPolicy,
		MergeCreateUpdate: MergeClouddeployCustomTargetTypeIamBinding,
		MergeDelete:       MergeClouddeployCustomTargetTypeIamBindingDelete,
	}
}

func ResourceConverterClouddeployCustomTargetTypeIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ClouddeployCustomTargetTypeIAMAssetType,
		Convert:           GetClouddeployCustomTargetTypeIamMemberCaiObject,
		FetchFullResource: FetchClouddeployCustomTargetTypeIamPolicy,
		MergeCreateUpdate: MergeClouddeployCustomTargetTypeIamMember,
		MergeDelete:       MergeClouddeployCustomTargetTypeIamMemberDelete,
	}
}

func GetClouddeployCustomTargetTypeIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newClouddeployCustomTargetTypeIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetClouddeployCustomTargetTypeIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newClouddeployCustomTargetTypeIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetClouddeployCustomTargetTypeIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newClouddeployCustomTargetTypeIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeClouddeployCustomTargetTypeIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeClouddeployCustomTargetTypeIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeClouddeployCustomTargetTypeIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeClouddeployCustomTargetTypeIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeClouddeployCustomTargetTypeIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newClouddeployCustomTargetTypeIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//clouddeploy.googleapis.com/projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ClouddeployCustomTargetTypeIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchClouddeployCustomTargetTypeIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ClouddeployCustomTargetTypeIamUpdaterProducer,
		d,
		config,
		"//clouddeploy.googleapis.com/projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}",
		ClouddeployCustomTargetTypeIAMAssetType,
	)
}
