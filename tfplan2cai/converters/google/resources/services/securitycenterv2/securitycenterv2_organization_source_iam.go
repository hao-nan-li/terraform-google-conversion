// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenterv2/OrganizationSource.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenterv2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const SecurityCenterV2OrganizationSourceIAMAssetType string = "securitycenter.googleapis.com/OrganizationSource"

func ResourceConverterSecurityCenterV2OrganizationSourceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecurityCenterV2OrganizationSourceIAMAssetType,
		Convert:           GetSecurityCenterV2OrganizationSourceIamPolicyCaiObject,
		MergeCreateUpdate: MergeSecurityCenterV2OrganizationSourceIamPolicy,
	}
}

func ResourceConverterSecurityCenterV2OrganizationSourceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecurityCenterV2OrganizationSourceIAMAssetType,
		Convert:           GetSecurityCenterV2OrganizationSourceIamBindingCaiObject,
		FetchFullResource: FetchSecurityCenterV2OrganizationSourceIamPolicy,
		MergeCreateUpdate: MergeSecurityCenterV2OrganizationSourceIamBinding,
		MergeDelete:       MergeSecurityCenterV2OrganizationSourceIamBindingDelete,
	}
}

func ResourceConverterSecurityCenterV2OrganizationSourceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecurityCenterV2OrganizationSourceIAMAssetType,
		Convert:           GetSecurityCenterV2OrganizationSourceIamMemberCaiObject,
		FetchFullResource: FetchSecurityCenterV2OrganizationSourceIamPolicy,
		MergeCreateUpdate: MergeSecurityCenterV2OrganizationSourceIamMember,
		MergeDelete:       MergeSecurityCenterV2OrganizationSourceIamMemberDelete,
	}
}

func GetSecurityCenterV2OrganizationSourceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecurityCenterV2OrganizationSourceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetSecurityCenterV2OrganizationSourceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecurityCenterV2OrganizationSourceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetSecurityCenterV2OrganizationSourceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecurityCenterV2OrganizationSourceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeSecurityCenterV2OrganizationSourceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeSecurityCenterV2OrganizationSourceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeSecurityCenterV2OrganizationSourceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeSecurityCenterV2OrganizationSourceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeSecurityCenterV2OrganizationSourceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newSecurityCenterV2OrganizationSourceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/organizations/{{organization}}/sources/{{source}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: SecurityCenterV2OrganizationSourceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchSecurityCenterV2OrganizationSourceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("organization"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("source"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		SecurityCenterV2OrganizationSourceIamUpdaterProducer,
		d,
		config,
		"//securitycenter.googleapis.com/organizations/{{organization}}/sources/{{source}}",
		SecurityCenterV2OrganizationSourceIAMAssetType,
	)
}
