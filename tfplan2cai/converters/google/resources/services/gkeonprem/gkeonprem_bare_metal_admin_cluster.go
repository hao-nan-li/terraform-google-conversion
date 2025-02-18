// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkeonprem/BareMetalAdminCluster.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gkeonprem

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GkeonpremBareMetalAdminClusterAssetType string = "gkeonprem.googleapis.com/BareMetalAdminCluster"

func ResourceConverterGkeonpremBareMetalAdminCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GkeonpremBareMetalAdminClusterAssetType,
		Convert:   GetGkeonpremBareMetalAdminClusterCaiObject,
	}
}

func GetGkeonpremBareMetalAdminClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkeonprem.googleapis.com/projects/{{project}}/locations/{{location}}/bareMetalAdminClusters/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGkeonpremBareMetalAdminClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GkeonpremBareMetalAdminClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkeonprem/v1/rest",
				DiscoveryName:        "BareMetalAdminCluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGkeonpremBareMetalAdminClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandGkeonpremBareMetalAdminClusterDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	bareMetalVersionProp, err := expandGkeonpremBareMetalAdminClusterBareMetalVersion(d.Get("bare_metal_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bare_metal_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(bareMetalVersionProp)) && (ok || !reflect.DeepEqual(v, bareMetalVersionProp)) {
		obj["bareMetalVersion"] = bareMetalVersionProp
	}
	networkConfigProp, err := expandGkeonpremBareMetalAdminClusterNetworkConfig(d.Get("network_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkConfigProp)) && (ok || !reflect.DeepEqual(v, networkConfigProp)) {
		obj["networkConfig"] = networkConfigProp
	}
	controlPlaneProp, err := expandGkeonpremBareMetalAdminClusterControlPlane(d.Get("control_plane"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("control_plane"); !tpgresource.IsEmptyValue(reflect.ValueOf(controlPlaneProp)) && (ok || !reflect.DeepEqual(v, controlPlaneProp)) {
		obj["controlPlane"] = controlPlaneProp
	}
	loadBalancerProp, err := expandGkeonpremBareMetalAdminClusterLoadBalancer(d.Get("load_balancer"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancer"); !tpgresource.IsEmptyValue(reflect.ValueOf(loadBalancerProp)) && (ok || !reflect.DeepEqual(v, loadBalancerProp)) {
		obj["loadBalancer"] = loadBalancerProp
	}
	storageProp, err := expandGkeonpremBareMetalAdminClusterStorage(d.Get("storage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("storage"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageProp)) && (ok || !reflect.DeepEqual(v, storageProp)) {
		obj["storage"] = storageProp
	}
	proxyProp, err := expandGkeonpremBareMetalAdminClusterProxy(d.Get("proxy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("proxy"); !tpgresource.IsEmptyValue(reflect.ValueOf(proxyProp)) && (ok || !reflect.DeepEqual(v, proxyProp)) {
		obj["proxy"] = proxyProp
	}
	clusterOperationsProp, err := expandGkeonpremBareMetalAdminClusterClusterOperations(d.Get("cluster_operations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cluster_operations"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterOperationsProp)) && (ok || !reflect.DeepEqual(v, clusterOperationsProp)) {
		obj["clusterOperations"] = clusterOperationsProp
	}
	maintenanceConfigProp, err := expandGkeonpremBareMetalAdminClusterMaintenanceConfig(d.Get("maintenance_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenanceConfigProp)) && (ok || !reflect.DeepEqual(v, maintenanceConfigProp)) {
		obj["maintenanceConfig"] = maintenanceConfigProp
	}
	nodeConfigProp, err := expandGkeonpremBareMetalAdminClusterNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	nodeAccessConfigProp, err := expandGkeonpremBareMetalAdminClusterNodeAccessConfig(d.Get("node_access_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_access_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeAccessConfigProp)) && (ok || !reflect.DeepEqual(v, nodeAccessConfigProp)) {
		obj["nodeAccessConfig"] = nodeAccessConfigProp
	}
	securityConfigProp, err := expandGkeonpremBareMetalAdminClusterSecurityConfig(d.Get("security_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(securityConfigProp)) && (ok || !reflect.DeepEqual(v, securityConfigProp)) {
		obj["securityConfig"] = securityConfigProp
	}
	annotationsProp, err := expandGkeonpremBareMetalAdminClusterEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandGkeonpremBareMetalAdminClusterDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterBareMetalVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterNetworkConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIslandModeCidr, err := expandGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidr(original["island_mode_cidr"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIslandModeCidr); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["islandModeCidr"] = transformedIslandModeCidr
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAddressCidrBlocks, err := expandGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidrServiceAddressCidrBlocks(original["service_address_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAddressCidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAddressCidrBlocks"] = transformedServiceAddressCidrBlocks
	}

	transformedPodAddressCidrBlocks, err := expandGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidrPodAddressCidrBlocks(original["pod_address_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPodAddressCidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["podAddressCidrBlocks"] = transformedPodAddressCidrBlocks
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidrServiceAddressCidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterNetworkConfigIslandModeCidrPodAddressCidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlane(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedControlPlaneNodePoolConfig, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig(original["control_plane_node_pool_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedControlPlaneNodePoolConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["controlPlaneNodePoolConfig"] = transformedControlPlaneNodePoolConfig
	}

	transformedApiServerArgs, err := expandGkeonpremBareMetalAdminClusterControlPlaneApiServerArgs(original["api_server_args"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedApiServerArgs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["apiServerArgs"] = transformedApiServerArgs
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodePoolConfig, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig(original["node_pool_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodePoolConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodePoolConfig"] = transformedNodePoolConfig
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeConfigs, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigs(original["node_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeConfigs"] = transformedNodeConfigs
	}

	transformedOperatingSystem, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigOperatingSystem(original["operating_system"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOperatingSystem); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["operatingSystem"] = transformedOperatingSystem
	}

	transformedTaints, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaints(original["taints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTaints); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["taints"] = transformedTaints
	}

	transformedLabels, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNodeIp, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigsNodeIp(original["node_ip"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNodeIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nodeIp"] = transformedNodeIp
		}

		transformedLabels, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigsLabels(original["labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["labels"] = transformedLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigsNodeIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigsLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigOperatingSystem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedEffect, err := expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintsEffect(original["effect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEffect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["effect"] = transformedEffect
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintsEffect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneApiServerArgs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedArgument, err := expandGkeonpremBareMetalAdminClusterControlPlaneApiServerArgsArgument(original["argument"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArgument); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["argument"] = transformedArgument
		}

		transformedValue, err := expandGkeonpremBareMetalAdminClusterControlPlaneApiServerArgsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneApiServerArgsArgument(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterControlPlaneApiServerArgsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVipConfig, err := expandGkeonpremBareMetalAdminClusterLoadBalancerVipConfig(original["vip_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVipConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vipConfig"] = transformedVipConfig
	}

	transformedPortConfig, err := expandGkeonpremBareMetalAdminClusterLoadBalancerPortConfig(original["port_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["portConfig"] = transformedPortConfig
	}

	transformedManualLbConfig, err := expandGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig(original["manual_lb_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedManualLbConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["manualLbConfig"] = transformedManualLbConfig
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancerVipConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedControlPlaneVip, err := expandGkeonpremBareMetalAdminClusterLoadBalancerVipConfigControlPlaneVip(original["control_plane_vip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedControlPlaneVip); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["controlPlaneVip"] = transformedControlPlaneVip
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancerVipConfigControlPlaneVip(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancerPortConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedControlPlaneLoadBalancerPort, err := expandGkeonpremBareMetalAdminClusterLoadBalancerPortConfigControlPlaneLoadBalancerPort(original["control_plane_load_balancer_port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedControlPlaneLoadBalancerPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["controlPlaneLoadBalancerPort"] = transformedControlPlaneLoadBalancerPort
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancerPortConfigControlPlaneLoadBalancerPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfigEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterLoadBalancerManualLbConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterStorage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLvpShareConfig, err := expandGkeonpremBareMetalAdminClusterStorageLvpShareConfig(original["lvp_share_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLvpShareConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["lvpShareConfig"] = transformedLvpShareConfig
	}

	transformedLvpNodeMountsConfig, err := expandGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfig(original["lvp_node_mounts_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLvpNodeMountsConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["lvpNodeMountsConfig"] = transformedLvpNodeMountsConfig
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpShareConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLvpConfig, err := expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfig(original["lvp_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLvpConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["lvpConfig"] = transformedLvpConfig
	}

	transformedSharedPathPvCount, err := expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigSharedPathPvCount(original["shared_path_pv_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSharedPathPvCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sharedPathPvCount"] = transformedSharedPathPvCount
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfigPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedStorageClass, err := expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfigStorageClass(original["storage_class"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageClass); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageClass"] = transformedStorageClass
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfigPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigLvpConfigStorageClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpShareConfigSharedPathPvCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfigPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedStorageClass, err := expandGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfigStorageClass(original["storage_class"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageClass); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageClass"] = transformedStorageClass
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfigPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterStorageLvpNodeMountsConfigStorageClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterProxy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandGkeonpremBareMetalAdminClusterProxyUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedNoProxy, err := expandGkeonpremBareMetalAdminClusterProxyNoProxy(original["no_proxy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoProxy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["noProxy"] = transformedNoProxy
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterProxyUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterProxyNoProxy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterClusterOperations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableApplicationLogs, err := expandGkeonpremBareMetalAdminClusterClusterOperationsEnableApplicationLogs(original["enable_application_logs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableApplicationLogs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableApplicationLogs"] = transformedEnableApplicationLogs
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterClusterOperationsEnableApplicationLogs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterMaintenanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaintenanceAddressCidrBlocks, err := expandGkeonpremBareMetalAdminClusterMaintenanceConfigMaintenanceAddressCidrBlocks(original["maintenance_address_cidr_blocks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaintenanceAddressCidrBlocks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maintenanceAddressCidrBlocks"] = transformedMaintenanceAddressCidrBlocks
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterMaintenanceConfigMaintenanceAddressCidrBlocks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterNodeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxPodsPerNode, err := expandGkeonpremBareMetalAdminClusterNodeConfigMaxPodsPerNode(original["max_pods_per_node"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxPodsPerNode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxPodsPerNode"] = transformedMaxPodsPerNode
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterNodeConfigMaxPodsPerNode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterNodeAccessConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLoginUser, err := expandGkeonpremBareMetalAdminClusterNodeAccessConfigLoginUser(original["login_user"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoginUser); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loginUser"] = transformedLoginUser
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterNodeAccessConfigLoginUser(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterSecurityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAuthorization, err := expandGkeonpremBareMetalAdminClusterSecurityConfigAuthorization(original["authorization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthorization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["authorization"] = transformedAuthorization
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterSecurityConfigAuthorization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAdminUsers, err := expandGkeonpremBareMetalAdminClusterSecurityConfigAuthorizationAdminUsers(original["admin_users"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdminUsers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["adminUsers"] = transformedAdminUsers
	}

	return transformed, nil
}

func expandGkeonpremBareMetalAdminClusterSecurityConfigAuthorizationAdminUsers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUsername, err := expandGkeonpremBareMetalAdminClusterSecurityConfigAuthorizationAdminUsersUsername(original["username"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["username"] = transformedUsername
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalAdminClusterSecurityConfigAuthorizationAdminUsersUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalAdminClusterEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
