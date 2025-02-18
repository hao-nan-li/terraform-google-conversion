// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudrunv2/Service.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudrunv2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudRunV2ServiceIAMAssetType string = "run.googleapis.com/Service"

func ResourceConverterCloudRunV2ServiceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunV2ServiceIAMAssetType,
		Convert:           GetCloudRunV2ServiceIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudRunV2ServiceIamPolicy,
	}
}

func ResourceConverterCloudRunV2ServiceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunV2ServiceIAMAssetType,
		Convert:           GetCloudRunV2ServiceIamBindingCaiObject,
		FetchFullResource: FetchCloudRunV2ServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2ServiceIamBinding,
		MergeDelete:       MergeCloudRunV2ServiceIamBindingDelete,
	}
}

func ResourceConverterCloudRunV2ServiceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunV2ServiceIAMAssetType,
		Convert:           GetCloudRunV2ServiceIamMemberCaiObject,
		FetchFullResource: FetchCloudRunV2ServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2ServiceIamMember,
		MergeDelete:       MergeCloudRunV2ServiceIamMemberDelete,
	}
}

func GetCloudRunV2ServiceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunV2ServiceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetCloudRunV2ServiceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunV2ServiceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetCloudRunV2ServiceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunV2ServiceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeCloudRunV2ServiceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudRunV2ServiceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeCloudRunV2ServiceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeCloudRunV2ServiceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeCloudRunV2ServiceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newCloudRunV2ServiceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: CloudRunV2ServiceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudRunV2ServiceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		CloudRunV2ServiceIamUpdaterProducer,
		d,
		config,
		"//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{name}}",
		CloudRunV2ServiceIAMAssetType,
	)
}
