// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataplex/Task.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataplex

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataplexTaskIAMAssetType string = "dataplex.googleapis.com/Task"

func ResourceConverterDataplexTaskIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexTaskIAMAssetType,
		Convert:           GetDataplexTaskIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataplexTaskIamPolicy,
	}
}

func ResourceConverterDataplexTaskIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexTaskIAMAssetType,
		Convert:           GetDataplexTaskIamBindingCaiObject,
		FetchFullResource: FetchDataplexTaskIamPolicy,
		MergeCreateUpdate: MergeDataplexTaskIamBinding,
		MergeDelete:       MergeDataplexTaskIamBindingDelete,
	}
}

func ResourceConverterDataplexTaskIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexTaskIAMAssetType,
		Convert:           GetDataplexTaskIamMemberCaiObject,
		FetchFullResource: FetchDataplexTaskIamPolicy,
		MergeCreateUpdate: MergeDataplexTaskIamMember,
		MergeDelete:       MergeDataplexTaskIamMemberDelete,
	}
}

func GetDataplexTaskIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexTaskIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataplexTaskIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexTaskIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataplexTaskIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexTaskIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataplexTaskIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataplexTaskIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataplexTaskIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataplexTaskIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataplexTaskIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataplexTaskIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataplexTaskIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataplexTaskIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("lake"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("task_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataplexTaskIamUpdaterProducer,
		d,
		config,
		"//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}",
		DataplexTaskIAMAssetType,
	)
}
