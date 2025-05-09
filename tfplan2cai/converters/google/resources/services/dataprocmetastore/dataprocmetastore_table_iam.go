// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/metastore/Table.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataprocmetastore

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataprocMetastoreTableIAMAssetType string = "metastore.googleapis.com/Table"

func ResourceConverterDataprocMetastoreTableIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreTableIAMAssetType,
		Convert:           GetDataprocMetastoreTableIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataprocMetastoreTableIamPolicy,
	}
}

func ResourceConverterDataprocMetastoreTableIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreTableIAMAssetType,
		Convert:           GetDataprocMetastoreTableIamBindingCaiObject,
		FetchFullResource: FetchDataprocMetastoreTableIamPolicy,
		MergeCreateUpdate: MergeDataprocMetastoreTableIamBinding,
		MergeDelete:       MergeDataprocMetastoreTableIamBindingDelete,
	}
}

func ResourceConverterDataprocMetastoreTableIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreTableIAMAssetType,
		Convert:           GetDataprocMetastoreTableIamMemberCaiObject,
		FetchFullResource: FetchDataprocMetastoreTableIamPolicy,
		MergeCreateUpdate: MergeDataprocMetastoreTableIamMember,
		MergeDelete:       MergeDataprocMetastoreTableIamMemberDelete,
	}
}

func GetDataprocMetastoreTableIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreTableIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataprocMetastoreTableIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreTableIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataprocMetastoreTableIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreTableIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataprocMetastoreTableIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataprocMetastoreTableIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataprocMetastoreTableIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataprocMetastoreTableIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataprocMetastoreTableIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataprocMetastoreTableIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{serviceId}}/databases/{{databaseId}}/tables/{{table}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataprocMetastoreTableIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataprocMetastoreTableIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("serviceId"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("databaseId"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("table"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataprocMetastoreTableIamUpdaterProducer,
		d,
		config,
		"//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{serviceId}}/databases/{{databaseId}}/tables/{{table}}",
		DataprocMetastoreTableIAMAssetType,
	)
}
