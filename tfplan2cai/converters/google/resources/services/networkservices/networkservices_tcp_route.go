// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkservices/TcpRoute.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networkservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkServicesTcpRouteAssetType string = "networkservices.googleapis.com/TcpRoute"

func ResourceConverterNetworkServicesTcpRoute() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkServicesTcpRouteAssetType,
		Convert:   GetNetworkServicesTcpRouteCaiObject,
	}
}

func GetNetworkServicesTcpRouteCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/global/tcpRoutes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkServicesTcpRouteApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkServicesTcpRouteAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "TcpRoute",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkServicesTcpRouteApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesTcpRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	meshesProp, err := expandNetworkServicesTcpRouteMeshes(d.Get("meshes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("meshes"); ok || !reflect.DeepEqual(v, meshesProp) {
		obj["meshes"] = meshesProp
	}
	gatewaysProp, err := expandNetworkServicesTcpRouteGateways(d.Get("gateways"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gateways"); ok || !reflect.DeepEqual(v, gatewaysProp) {
		obj["gateways"] = gatewaysProp
	}
	rulesProp, err := expandNetworkServicesTcpRouteRules(d.Get("rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rules"); ok || !reflect.DeepEqual(v, rulesProp) {
		obj["rules"] = rulesProp
	}
	labelsProp, err := expandNetworkServicesTcpRouteEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkServicesTcpRouteDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteMeshes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteGateways(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMatches, err := expandNetworkServicesTcpRouteRulesMatches(original["matches"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["matches"] = transformedMatches
		}

		transformedAction, err := expandNetworkServicesTcpRouteRulesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTcpRouteRulesMatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAddress, err := expandNetworkServicesTcpRouteRulesMatchesAddress(original["address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["address"] = transformedAddress
		}

		transformedPort, err := expandNetworkServicesTcpRouteRulesMatchesPort(original["port"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["port"] = transformedPort
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTcpRouteRulesMatchesAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesMatchesPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDestinations, err := expandNetworkServicesTcpRouteRulesActionDestinations(original["destinations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinations"] = transformedDestinations
	}

	transformedOriginalDestination, err := expandNetworkServicesTcpRouteRulesActionOriginalDestination(original["original_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOriginalDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["originalDestination"] = transformedOriginalDestination
	}

	transformedIdleTimeout, err := expandNetworkServicesTcpRouteRulesActionIdleTimeout(original["idle_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdleTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idleTimeout"] = transformedIdleTimeout
	}

	return transformed, nil
}

func expandNetworkServicesTcpRouteRulesActionDestinations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedServiceName, err := expandNetworkServicesTcpRouteRulesActionDestinationsServiceName(original["service_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["serviceName"] = transformedServiceName
		}

		transformedWeight, err := expandNetworkServicesTcpRouteRulesActionDestinationsWeight(original["weight"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["weight"] = transformedWeight
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTcpRouteRulesActionDestinationsServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesActionDestinationsWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesActionOriginalDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesActionIdleTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
