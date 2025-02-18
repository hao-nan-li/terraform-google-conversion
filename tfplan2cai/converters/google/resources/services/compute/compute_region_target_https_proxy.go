// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/RegionTargetHttpsProxy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeRegionTargetHttpsProxyAssetType string = "compute.googleapis.com/RegionTargetHttpsProxy"

func ResourceConverterComputeRegionTargetHttpsProxy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeRegionTargetHttpsProxyAssetType,
		Convert:   GetComputeRegionTargetHttpsProxyCaiObject,
	}
}

func GetComputeRegionTargetHttpsProxyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeRegionTargetHttpsProxyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeRegionTargetHttpsProxyAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionTargetHttpsProxy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeRegionTargetHttpsProxyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionTargetHttpsProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionTargetHttpsProxyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	certificateManagerCertificatesProp, err := expandComputeRegionTargetHttpsProxyCertificateManagerCertificates(d.Get("certificate_manager_certificates"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate_manager_certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateManagerCertificatesProp)) && (ok || !reflect.DeepEqual(v, certificateManagerCertificatesProp)) {
		obj["certificateManagerCertificates"] = certificateManagerCertificatesProp
	}
	sslCertificatesProp, err := expandComputeRegionTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ssl_certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(sslCertificatesProp)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
		obj["sslCertificates"] = sslCertificatesProp
	}
	sslPolicyProp, err := expandComputeRegionTargetHttpsProxySslPolicy(d.Get("ssl_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ssl_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(sslPolicyProp)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
		obj["sslPolicy"] = sslPolicyProp
	}
	urlMapProp, err := expandComputeRegionTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("url_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	httpKeepAliveTimeoutSecProp, err := expandComputeRegionTargetHttpsProxyHttpKeepAliveTimeoutSec(d.Get("http_keep_alive_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_keep_alive_timeout_sec"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpKeepAliveTimeoutSecProp)) && (ok || !reflect.DeepEqual(v, httpKeepAliveTimeoutSecProp)) {
		obj["httpKeepAliveTimeoutSec"] = httpKeepAliveTimeoutSecProp
	}
	serverTlsPolicyProp, err := expandComputeRegionTargetHttpsProxyServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverTlsPolicyProp)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}
	regionProp, err := expandComputeRegionTargetHttpsProxyRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return resourceComputeRegionTargetHttpsProxyEncoder(d, config, obj)
}

func resourceComputeRegionTargetHttpsProxyEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	if _, ok := obj["certificateManagerCertificates"]; ok {
		// The field certificateManagerCertificates should not be included in the API request, and it should be renamed to `sslCertificates`
		// The API does not allow using both certificate manager certificates and sslCertificates. If that changes
		// in the future, the encoder logic should change accordingly because this will mean that both fields are no longer mutual exclusive.
		log.Printf("[DEBUG] converting the field CertificateManagerCertificates to sslCertificates before sending the request")
		obj["sslCertificates"] = obj["certificateManagerCertificates"]
		delete(obj, "certificateManagerCertificates")
	}

	// Send null if serverTlsPolicy is not set. Without this, Terraform would not send any value for `serverTlsPolicy`
	// in the "PATCH" payload so if you were to remove a server TLS policy from a target HTTPS proxy, it would NOT remove
	// the association.
	if _, ok := obj["serverTlsPolicy"]; !ok {
		obj["serverTlsPolicy"] = nil
	}

	return obj, nil
}

func expandComputeRegionTargetHttpsProxyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpsProxyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpsProxyCertificateManagerCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for certificate_manager_certificates: nil")
		}
		if strings.HasPrefix(raw.(string), "//") || strings.HasPrefix(raw.(string), "https://") {
			// Any full URL will be passed to the API request (regardless of the resource type). This is to allow self_links of CertificateManagerCeritificate resources.
			// If the full URL is an invalid reference, that should be handled by the API.
			req = append(req, raw.(string))
		} else if reg, _ := regexp.Compile("projects/(.*)/locations/(.*)/certificates/(.*)"); reg.MatchString(raw.(string)) {
			// If the input is the id pattern of CertificateManagerCertificate resource, a prefix will be added to construct the full URL before constructing the API request.
			self_link := "https://certificatemanager.googleapis.com/v1/" + raw.(string)
			req = append(req, self_link)
		} else {
			return nil, fmt.Errorf("Invalid value for certificate_manager_certificates: %v is an invalid format for a certificateManagerCertificate resource", raw.(string))
		}
	}
	return req, nil
}

func expandComputeRegionTargetHttpsProxySslCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("sslCertificates", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeRegionTargetHttpsProxySslPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("sslPolicies", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for ssl_policy: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionTargetHttpsProxyUrlMap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("urlMaps", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionTargetHttpsProxyHttpKeepAliveTimeoutSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpsProxyServerTlsPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionTargetHttpsProxyRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
