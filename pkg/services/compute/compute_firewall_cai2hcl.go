// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Firewall.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc_next/cai2hcl/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tpgresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/transport"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/transport"
)

type ComputeFirewallCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeFirewallCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	schema := provider.ResourcesMap[ComputeFirewallSchemaName].Schema

	return &ComputeFirewallCai2hclConverter{
		name:   ComputeFirewallSchemaName,
		schema: schema,
	}
}

// Convert converts asset to HCL resource blocks.
func (c *ComputeFirewallCai2hclConverter) Convert(asset caiasset.Asset) ([]*models.TerraformResourceBlock, error) {
	var blocks []*models.TerraformResourceBlock
	block, err := c.convertResourceData(asset)
	if err != nil {
		return nil, err
	}
	blocks = append(blocks, block)
	return blocks, nil
}

func (c *ComputeFirewallCai2hclConverter) convertResourceData(asset caiasset.Asset) (*models.TerraformResourceBlock, error) {
	if asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	var err error
	res := asset.Resource.Data
	config := transport.NewConfig()
	d := &schema.ResourceData{}

	assetNameParts := strings.Split(asset.Name, "/")
	hclBlockName := assetNameParts[len(assetNameParts)-1]

	hclData := make(map[string]interface{})

	outputFields := map[string]struct{}{"creation_timestamp": struct{}{}}
	utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/global/firewalls/{{name}}", outputFields, hclData)

	hclData["allow"] = flattenComputeFirewallAllow(res["allowed"], d, config)
	hclData["deny"] = flattenComputeFirewallDeny(res["denied"], d, config)
	hclData["description"] = flattenComputeFirewallDescription(res["description"], d, config)
	hclData["destination_ranges"] = flattenComputeFirewallDestinationRanges(res["destinationRanges"], d, config)
	hclData["direction"] = flattenComputeFirewallDirection(res["direction"], d, config)
	hclData["disabled"] = flattenComputeFirewallDisabled(res["disabled"], d, config)
	hclData["log_config"] = flattenComputeFirewallLogConfig(res["logConfig"], d, config)
	hclData["name"] = flattenComputeFirewallName(res["name"], d, config)
	hclData["network"] = flattenComputeFirewallNetwork(res["network"], d, config)
	hclData["priority"] = flattenComputeFirewallPriority(res["priority"], d, config)
	hclData["source_ranges"] = flattenComputeFirewallSourceRanges(res["sourceRanges"], d, config)
	hclData["source_service_accounts"] = flattenComputeFirewallSourceServiceAccounts(res["sourceServiceAccounts"], d, config)
	hclData["source_tags"] = flattenComputeFirewallSourceTags(res["sourceTags"], d, config)
	hclData["target_service_accounts"] = flattenComputeFirewallTargetServiceAccounts(res["targetServiceAccounts"], d, config)
	hclData["target_tags"] = flattenComputeFirewallTargetTags(res["targetTags"], d, config)
	hclData["params"] = flattenComputeFirewallParams(res["params"], d, config)

	ctyVal, err := utils.MapToCtyValWithSchema(hclData, c.schema)
	if err != nil {
		return nil, err
	}
	return &models.TerraformResourceBlock{
		Labels: []string{c.name, hclBlockName},
		Value:  ctyVal,
	}, nil
}

func flattenComputeFirewallAllow(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceComputeFirewallRuleHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"protocol": flattenComputeFirewallAllowProtocol(original["IPProtocol"], d, config),
			"ports":    flattenComputeFirewallAllowPorts(original["ports"], d, config),
		})
	}
	return transformed
}

func flattenComputeFirewallAllowProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallAllowPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDeny(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceComputeFirewallRuleHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"protocol": flattenComputeFirewallDenyProtocol(original["IPProtocol"], d, config),
			"ports":    flattenComputeFirewallDenyPorts(original["ports"], d, config),
		})
	}
	return transformed
}

func flattenComputeFirewallDenyProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDenyPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDestinationRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallLogConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}

	v, ok := original["enable"]
	if ok && !v.(bool) {
		return nil
	}

	transformed := make(map[string]interface{})
	transformed["metadata"] = original["metadata"]
	return []interface{}{transformed}
}

func flattenComputeFirewallName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	relative, err := tpgresource.GetRelativePath(v.(string))
	if err != nil {
		return v
	}
	return relative
}

func flattenComputeFirewallPriority(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeFirewallSourceRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallSourceServiceAccounts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallSourceTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallTargetServiceAccounts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallTargetTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeFirewallParams(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("params")
}

func flattenComputeFirewallParamsResourceManagerTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("params.0.resource_manager_tags")
}
