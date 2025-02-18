// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/PublicAdvertisedPrefix.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputePublicAdvertisedPrefixAssetType string = "compute.googleapis.com/PublicAdvertisedPrefix"

func ResourceConverterComputePublicAdvertisedPrefix() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputePublicAdvertisedPrefixAssetType,
		Convert:   GetComputePublicAdvertisedPrefixCaiObject,
	}
}

func GetComputePublicAdvertisedPrefixCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/publicAdvertisedPrefixes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputePublicAdvertisedPrefixApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputePublicAdvertisedPrefixAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "PublicAdvertisedPrefix",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputePublicAdvertisedPrefixApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputePublicAdvertisedPrefixDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputePublicAdvertisedPrefixName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	dnsVerificationIpProp, err := expandComputePublicAdvertisedPrefixDnsVerificationIp(d.Get("dns_verification_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dns_verification_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnsVerificationIpProp)) && (ok || !reflect.DeepEqual(v, dnsVerificationIpProp)) {
		obj["dnsVerificationIp"] = dnsVerificationIpProp
	}
	ipCidrRangeProp, err := expandComputePublicAdvertisedPrefixIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	pdpScopeProp, err := expandComputePublicAdvertisedPrefixPdpScope(d.Get("pdp_scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pdp_scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(pdpScopeProp)) && (ok || !reflect.DeepEqual(v, pdpScopeProp)) {
		obj["pdpScope"] = pdpScopeProp
	}

	return obj, nil
}

func expandComputePublicAdvertisedPrefixDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixDnsVerificationIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicAdvertisedPrefixPdpScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
