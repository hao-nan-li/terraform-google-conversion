// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/privateca/CertificateAuthority.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package privateca

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func resourcePrivateCaCACustomDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	if diff.HasChange("desired_state") {
		_, new := diff.GetChange("desired_state")

		if tpgresource.IsNewResource(diff) {
			if new.(string) != "STAGED" && new.(string) != "ENABLED" {
				return fmt.Errorf("`desired_state` can only be set to `STAGED` or `ENABLED` when creating a new CA")
			}
		} else {
			if new == "STAGED" && diff.Get("state") != new {
				return fmt.Errorf("Field `desired_state` can only be set to `STAGED` when creating a new CA")
			}
		}
	}
	return nil
}

const PrivatecaCertificateAuthorityAssetType string = "privateca.googleapis.com/CertificateAuthority"

func ResourceConverterPrivatecaCertificateAuthority() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: PrivatecaCertificateAuthorityAssetType,
		Convert:   GetPrivatecaCertificateAuthorityCaiObject,
	}
}

func GetPrivatecaCertificateAuthorityCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//privateca.googleapis.com/projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificateAuthorities/{{certificate_authority_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetPrivatecaCertificateAuthorityApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: PrivatecaCertificateAuthorityAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/privateca/v1/rest",
				DiscoveryName:        "CertificateAuthority",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetPrivatecaCertificateAuthorityApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	typeProp, err := expandPrivatecaCertificateAuthorityType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	configProp, err := expandPrivatecaCertificateAuthorityConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	lifetimeProp, err := expandPrivatecaCertificateAuthorityLifetime(d.Get("lifetime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("lifetime"); !tpgresource.IsEmptyValue(reflect.ValueOf(lifetimeProp)) && (ok || !reflect.DeepEqual(v, lifetimeProp)) {
		obj["lifetime"] = lifetimeProp
	}
	keySpecProp, err := expandPrivatecaCertificateAuthorityKeySpec(d.Get("key_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(keySpecProp)) && (ok || !reflect.DeepEqual(v, keySpecProp)) {
		obj["keySpec"] = keySpecProp
	}
	subordinateConfigProp, err := expandPrivatecaCertificateAuthoritySubordinateConfig(d.Get("subordinate_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subordinate_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(subordinateConfigProp)) && (ok || !reflect.DeepEqual(v, subordinateConfigProp)) {
		obj["subordinateConfig"] = subordinateConfigProp
	}
	gcsBucketProp, err := expandPrivatecaCertificateAuthorityGcsBucket(d.Get("gcs_bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gcs_bucket"); !tpgresource.IsEmptyValue(reflect.ValueOf(gcsBucketProp)) && (ok || !reflect.DeepEqual(v, gcsBucketProp)) {
		obj["gcsBucket"] = gcsBucketProp
	}
	userDefinedAccessUrlsProp, err := expandPrivatecaCertificateAuthorityUserDefinedAccessUrls(d.Get("user_defined_access_urls"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user_defined_access_urls"); !tpgresource.IsEmptyValue(reflect.ValueOf(userDefinedAccessUrlsProp)) && (ok || !reflect.DeepEqual(v, userDefinedAccessUrlsProp)) {
		obj["userDefinedAccessUrls"] = userDefinedAccessUrlsProp
	}
	labelsProp, err := expandPrivatecaCertificateAuthorityEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandPrivatecaCertificateAuthorityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubjectKeyId, err := expandPrivatecaCertificateAuthorityConfigSubjectKeyId(original["subject_key_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectKeyId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectKeyId"] = transformedSubjectKeyId
	}

	transformedX509Config, err := expandPrivatecaCertificateAuthorityConfigX509Config(original["x509_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedX509Config); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["x509Config"] = transformedX509Config
	}

	transformedSubjectConfig, err := expandPrivatecaCertificateAuthorityConfigSubjectConfig(original["subject_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectConfig"] = transformedSubjectConfig
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKeyId, err := expandPrivatecaCertificateAuthorityConfigSubjectKeyIdKeyId(original["key_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keyId"] = transformedKeyId
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectKeyIdKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigX509Config(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return v, nil
	}
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	if len(original) == 0 {
		return nil, nil
	}
	transformed := make(map[string]interface{})

	caOptions, err := expandPrivatecaCertificateConfigX509ConfigCaOptions(original["ca_options"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["caOptions"] = caOptions

	keyUsage, err := expandPrivatecaCertificateConfigX509ConfigKeyUsage(original["key_usage"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["keyUsage"] = keyUsage

	policyIds, err := expandPrivatecaCertificateConfigX509ConfigPolicyIds(original["policy_ids"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["policyIds"] = policyIds

	aiaOcspServers, err := expandPrivatecaCertificateConfigX509ConfigAiaOcspServers(original["aia_ocsp_servers"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["aiaOcspServers"] = aiaOcspServers

	addExts, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensions(original["additional_extensions"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["additionalExtensions"] = addExts

	nameConstraints, err := expandPrivatecaCertificateConfigX509ConfigNameConstraints(original["name_constraints"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["nameConstraints"] = nameConstraints
	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubject, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubject(original["subject"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subject"] = transformedSubject
	}

	transformedSubjectAltName, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName(original["subject_alt_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectAltName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectAltName"] = transformedSubjectAltName
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCountryCode, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCountryCode(original["country_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCountryCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["countryCode"] = transformedCountryCode
	}

	transformedOrganization, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedOrganizationalUnit, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganizationalUnit(original["organizational_unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganizationalUnit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organizationalUnit"] = transformedOrganizationalUnit
	}

	transformedLocality, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedProvince, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectProvince(original["province"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProvince); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["province"] = transformedProvince
	}

	transformedStreetAddress, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectStreetAddress(original["street_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStreetAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["streetAddress"] = transformedStreetAddress
	}

	transformedPostalCode, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedCommonName, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCommonName(original["common_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommonName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["commonName"] = transformedCommonName
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCountryCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectOrganizationalUnit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectProvince(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectStreetAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectPostalCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectCommonName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDnsNames, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameDnsNames(original["dns_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsNames); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dnsNames"] = transformedDnsNames
	}

	transformedUris, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedEmailAddresses, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameEmailAddresses(original["email_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmailAddresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["emailAddresses"] = transformedEmailAddresses
	}

	transformedIpAddresses, err := expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameIpAddresses(original["ip_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddresses"] = transformedIpAddresses
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameDnsNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameEmailAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltNameIpAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityLifetime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityKeySpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCloudKmsKeyVersion, err := expandPrivatecaCertificateAuthorityKeySpecCloudKmsKeyVersion(original["cloud_kms_key_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudKmsKeyVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cloudKmsKeyVersion"] = transformedCloudKmsKeyVersion
	}

	transformedAlgorithm, err := expandPrivatecaCertificateAuthorityKeySpecAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityKeySpecCloudKmsKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityKeySpecAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthoritySubordinateConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificateAuthority, err := expandPrivatecaCertificateAuthoritySubordinateConfigCertificateAuthority(original["certificate_authority"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateAuthority); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateAuthority"] = transformedCertificateAuthority
	}

	transformedPemIssuerChain, err := expandPrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain(original["pem_issuer_chain"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPemIssuerChain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pemIssuerChain"] = transformedPemIssuerChain
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthoritySubordinateConfigCertificateAuthority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPemCertificates, err := expandPrivatecaCertificateAuthoritySubordinateConfigPemIssuerChainPemCertificates(original["pem_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPemCertificates); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pemCertificates"] = transformedPemCertificates
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthoritySubordinateConfigPemIssuerChainPemCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityGcsBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityUserDefinedAccessUrls(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAiaIssuingCertificateUrls, err := expandPrivatecaCertificateAuthorityUserDefinedAccessUrlsAiaIssuingCertificateUrls(original["aia_issuing_certificate_urls"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAiaIssuingCertificateUrls); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["aiaIssuingCertificateUrls"] = transformedAiaIssuingCertificateUrls
	}

	transformedCrlAccessUrls, err := expandPrivatecaCertificateAuthorityUserDefinedAccessUrlsCrlAccessUrls(original["crl_access_urls"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCrlAccessUrls); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["crlAccessUrls"] = transformedCrlAccessUrls
	}

	return transformed, nil
}

func expandPrivatecaCertificateAuthorityUserDefinedAccessUrlsAiaIssuingCertificateUrls(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityUserDefinedAccessUrlsCrlAccessUrls(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateAuthorityEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
