// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudtasks/Queue.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudtasks

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudTasksQueueIAMAssetType string = "cloudtasks.googleapis.com/Queue"

func ResourceConverterCloudTasksQueueIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudTasksQueueIAMAssetType,
		Convert:           GetCloudTasksQueueIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudTasksQueueIamPolicy,
	}
}

func ResourceConverterCloudTasksQueueIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudTasksQueueIAMAssetType,
		Convert:           GetCloudTasksQueueIamBindingCaiObject,
		FetchFullResource: FetchCloudTasksQueueIamPolicy,
		MergeCreateUpdate: MergeCloudTasksQueueIamBinding,
		MergeDelete:       MergeCloudTasksQueueIamBindingDelete,
	}
}

func ResourceConverterCloudTasksQueueIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudTasksQueueIAMAssetType,
		Convert:           GetCloudTasksQueueIamMemberCaiObject,
		FetchFullResource: FetchCloudTasksQueueIamPolicy,
		MergeCreateUpdate: MergeCloudTasksQueueIamMember,
		MergeDelete:       MergeCloudTasksQueueIamMemberDelete,
	}
}

func GetCloudTasksQueueIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudTasksQueueIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetCloudTasksQueueIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudTasksQueueIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetCloudTasksQueueIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudTasksQueueIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeCloudTasksQueueIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudTasksQueueIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeCloudTasksQueueIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeCloudTasksQueueIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeCloudTasksQueueIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newCloudTasksQueueIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//cloudtasks.googleapis.com/projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: CloudTasksQueueIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudTasksQueueIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		CloudTasksQueueIamUpdaterProducer,
		d,
		config,
		"//cloudtasks.googleapis.com/projects/{{project}}/locations/{{location}}/queues/{{name}}",
		CloudTasksQueueIAMAssetType,
	)
}
