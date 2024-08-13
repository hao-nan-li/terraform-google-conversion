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

package compute

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Whether the IP CIDR change shrinks the block.
func IsShrinkageIpCidr(_ context.Context, old, new, _ interface{}) bool {
	_, oldCidr, oldErr := net.ParseCIDR(old.(string))
	_, newCidr, newErr := net.ParseCIDR(new.(string))

	if oldErr != nil || newErr != nil {
		// This should never happen. The ValidateFunc on the field ensures it.
		return false
	}

	oldStart, oldEnd := cidr.AddressRange(oldCidr)

	if newCidr.Contains(oldStart) && newCidr.Contains(oldEnd) {
		// This is a CIDR range expansion, no need to ForceNew, we have an update method for it.
		return false
	}

	return true
}

func sendSecondaryIpRangeIfEmptyDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	// on create, return immediately as we don't need to determine if the value is empty or not
	if diff.Id() == "" {
		return nil
	}

	sendZero := diff.Get("send_secondary_ip_range_if_empty").(bool)
	if !sendZero {
		return nil
	}

	configSecondaryIpRange := diff.GetRawConfig().GetAttr("secondary_ip_range")
	if !configSecondaryIpRange.IsKnown() {
		return nil
	}
	configValueIsEmpty := configSecondaryIpRange.IsNull() || configSecondaryIpRange.LengthInt() == 0

	stateSecondaryIpRange := diff.GetRawState().GetAttr("secondary_ip_range")
	if !stateSecondaryIpRange.IsKnown() {
		return nil
	}
	stateValueIsEmpty := stateSecondaryIpRange.IsNull() || stateSecondaryIpRange.LengthInt() == 0

	if configValueIsEmpty && !stateValueIsEmpty {
		log.Printf("[DEBUG] setting secondary_ip_range to newly empty")
		diff.SetNew("secondary_ip_range", make([]interface{}, 0))
	}

	return nil
}

const ComputeSubnetworkAssetType string = "compute.googleapis.com/Subnetwork"

func ResourceConverterComputeSubnetwork() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeSubnetworkAssetType,
		Convert:   GetComputeSubnetworkCaiObject,
	}
}

func GetComputeSubnetworkCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeSubnetworkApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeSubnetworkAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "Subnetwork",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeSubnetworkApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeSubnetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	ipCidrRangeProp, err := expandComputeSubnetworkIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	nameProp, err := expandComputeSubnetworkName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeSubnetworkNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	purposeProp, err := expandComputeSubnetworkPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purpose"); !tpgresource.IsEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	roleProp, err := expandComputeSubnetworkRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !tpgresource.IsEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}
	secondaryIpRangesProp, err := expandComputeSubnetworkSecondaryIpRange(d.Get("secondary_ip_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("secondary_ip_range"); ok || !reflect.DeepEqual(v, secondaryIpRangesProp) {
		obj["secondaryIpRanges"] = secondaryIpRangesProp
	}
	privateIpGoogleAccessProp, err := expandComputeSubnetworkPrivateIpGoogleAccess(d.Get("private_ip_google_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_ip_google_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateIpGoogleAccessProp)) && (ok || !reflect.DeepEqual(v, privateIpGoogleAccessProp)) {
		obj["privateIpGoogleAccess"] = privateIpGoogleAccessProp
	}
	privateIpv6GoogleAccessProp, err := expandComputeSubnetworkPrivateIpv6GoogleAccess(d.Get("private_ipv6_google_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_ipv6_google_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateIpv6GoogleAccessProp)) && (ok || !reflect.DeepEqual(v, privateIpv6GoogleAccessProp)) {
		obj["privateIpv6GoogleAccess"] = privateIpv6GoogleAccessProp
	}
	regionProp, err := expandComputeSubnetworkRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	logConfigProp, err := expandComputeSubnetworkLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_config"); ok || !reflect.DeepEqual(v, logConfigProp) {
		obj["logConfig"] = logConfigProp
	}
	stackTypeProp, err := expandComputeSubnetworkStackType(d.Get("stack_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("stack_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(stackTypeProp)) && (ok || !reflect.DeepEqual(v, stackTypeProp)) {
		obj["stackType"] = stackTypeProp
	}
	ipv6AccessTypeProp, err := expandComputeSubnetworkIpv6AccessType(d.Get("ipv6_access_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ipv6_access_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipv6AccessTypeProp)) && (ok || !reflect.DeepEqual(v, ipv6AccessTypeProp)) {
		obj["ipv6AccessType"] = ipv6AccessTypeProp
	}
	externalIpv6PrefixProp, err := expandComputeSubnetworkExternalIpv6Prefix(d.Get("external_ipv6_prefix"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_ipv6_prefix"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalIpv6PrefixProp)) && (ok || !reflect.DeepEqual(v, externalIpv6PrefixProp)) {
		obj["externalIpv6Prefix"] = externalIpv6PrefixProp
	}
	allowSubnetCidrRoutesOverlapProp, err := expandComputeSubnetworkAllowSubnetCidrRoutesOverlap(d.Get("allow_subnet_cidr_routes_overlap"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_subnet_cidr_routes_overlap"); ok || !reflect.DeepEqual(v, allowSubnetCidrRoutesOverlapProp) {
		obj["allowSubnetCidrRoutesOverlap"] = allowSubnetCidrRoutesOverlapProp
	}

	return obj, nil
}

func expandComputeSubnetworkDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSubnetworkPurpose(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkSecondaryIpRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRangeName, err := expandComputeSubnetworkSecondaryIpRangeRangeName(original["range_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRangeName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["rangeName"] = transformedRangeName
		}

		transformedIpCidrRange, err := expandComputeSubnetworkSecondaryIpRangeIpCidrRange(original["ip_cidr_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpCidrRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipCidrRange"] = transformedIpCidrRange
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeSubnetworkSecondaryIpRangeRangeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkSecondaryIpRangeIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkPrivateIpGoogleAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkPrivateIpv6GoogleAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeSubnetworkLogConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	transformed := make(map[string]interface{})
	if len(l) == 0 || l[0] == nil {
		purpose, ok := d.GetOkExists("purpose")

		if ok && (purpose.(string) == "REGIONAL_MANAGED_PROXY" || purpose.(string) == "GLOBAL_MANAGED_PROXY" || purpose.(string) == "INTERNAL_HTTPS_LOAD_BALANCER") {
			// Subnetworks for regional L7 ILB/XLB or cross-regional L7 ILB do not accept any values for logConfig
			return nil, nil
		}
		// send enable = false to ensure logging is disabled if there is no config
		transformed["enable"] = false
		return transformed, nil
	}

	raw := l[0]
	original := raw.(map[string]interface{})

	// The log_config block is specified, so logging should be enabled
	transformed["enable"] = true
	transformed["aggregationInterval"] = original["aggregation_interval"]
	transformed["flowSampling"] = original["flow_sampling"]
	transformed["metadata"] = original["metadata"]
	transformed["filterExpr"] = original["filter_expr"]

	// make it JSON marshallable
	transformed["metadataFields"] = original["metadata_fields"].(*schema.Set).List()

	return transformed, nil
}

func expandComputeSubnetworkStackType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkIpv6AccessType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkExternalIpv6Prefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeSubnetworkAllowSubnetCidrRoutesOverlap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
