// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dialogflowcx/TestCase.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dialogflowcx

import (
	"encoding/json"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXTestCaseAssetType string = "{{location}}-dialogflow.googleapis.com/TestCase"

func ResourceConverterDialogflowCXTestCase() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowCXTestCaseAssetType,
		Convert:   GetDialogflowCXTestCaseCaiObject,
	}
}

func GetDialogflowCXTestCaseCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/{{parent}}/testCases/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowCXTestCaseApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowCXTestCaseAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "TestCase",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowCXTestCaseApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	tagsProp, err := expandDialogflowCXTestCaseTags(d.Get("tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tags"); !tpgresource.IsEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	displayNameProp, err := expandDialogflowCXTestCaseDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	notesProp, err := expandDialogflowCXTestCaseNotes(d.Get("notes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notes"); !tpgresource.IsEmptyValue(reflect.ValueOf(notesProp)) && (ok || !reflect.DeepEqual(v, notesProp)) {
		obj["notes"] = notesProp
	}
	testConfigProp, err := expandDialogflowCXTestCaseTestConfig(d.Get("test_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("test_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(testConfigProp)) && (ok || !reflect.DeepEqual(v, testConfigProp)) {
		obj["testConfig"] = testConfigProp
	}
	testCaseConversationTurnsProp, err := expandDialogflowCXTestCaseTestCaseConversationTurns(d.Get("test_case_conversation_turns"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("test_case_conversation_turns"); !tpgresource.IsEmptyValue(reflect.ValueOf(testCaseConversationTurnsProp)) && (ok || !reflect.DeepEqual(v, testCaseConversationTurnsProp)) {
		obj["testCaseConversationTurns"] = testCaseConversationTurnsProp
	}

	return obj, nil
}

func expandDialogflowCXTestCaseTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseNotes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTrackingParameters, err := expandDialogflowCXTestCaseTestConfigTrackingParameters(original["tracking_parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrackingParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["trackingParameters"] = transformedTrackingParameters
	}

	transformedFlow, err := expandDialogflowCXTestCaseTestConfigFlow(original["flow"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFlow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["flow"] = transformedFlow
	}

	transformedPage, err := expandDialogflowCXTestCaseTestConfigPage(original["page"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["page"] = transformedPage
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestConfigTrackingParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestConfigFlow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestConfigPage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUserInput, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInput(original["user_input"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUserInput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["userInput"] = transformedUserInput
		}

		transformedVirtualAgentOutput, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutput(original["virtual_agent_output"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVirtualAgentOutput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["virtualAgentOutput"] = transformedVirtualAgentOutput
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInput, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInput(original["input"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["input"] = transformedInput
	}

	transformedInjectedParameters, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInjectedParameters(original["injected_parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInjectedParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["injectedParameters"] = transformedInjectedParameters
	}

	transformedIsWebhookEnabled, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputIsWebhookEnabled(original["is_webhook_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIsWebhookEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["isWebhookEnabled"] = transformedIsWebhookEnabled
	}

	transformedEnableSentimentAnalysis, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputEnableSentimentAnalysis(original["enable_sentiment_analysis"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSentimentAnalysis); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableSentimentAnalysis"] = transformedEnableSentimentAnalysis
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLanguageCode, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputLanguageCode(original["language_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLanguageCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["languageCode"] = transformedLanguageCode
	}

	transformedText, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["text"] = transformedText
	}

	transformedEvent, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputEvent(original["event"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEvent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["event"] = transformedEvent
	}

	transformedDtmf, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputDtmf(original["dtmf"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDtmf); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dtmf"] = transformedDtmf
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputText(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedText, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputTextText(original["text"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedText); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["text"] = transformedText
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputTextText(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputEvent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEvent, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputEventEvent(original["event"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEvent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["event"] = transformedEvent
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputEventEvent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputDtmf(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDigits, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputDtmfDigits(original["digits"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDigits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["digits"] = transformedDigits
	}

	transformedFinishDigit, err := expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputDtmfFinishDigit(original["finish_digit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFinishDigit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["finishDigit"] = transformedFinishDigit
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputDtmfDigits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInputDtmfFinishDigit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputInjectedParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputIsWebhookEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsUserInputEnableSentimentAnalysis(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSessionParameters, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputSessionParameters(original["session_parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSessionParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sessionParameters"] = transformedSessionParameters
	}

	transformedTriggeredIntent, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent(original["triggered_intent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTriggeredIntent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["triggeredIntent"] = transformedTriggeredIntent
	}

	transformedCurrentPage, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage(original["current_page"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCurrentPage); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["currentPage"] = transformedCurrentPage
	}

	transformedTextResponses, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponses(original["text_responses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTextResponses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["textResponses"] = transformedTextResponses
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputSessionParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedDisplayName, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentDisplayName(original["display_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["displayName"] = transformedDisplayName
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTriggeredIntentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedDisplayName, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageDisplayName(original["display_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["displayName"] = transformedDisplayName
	}

	return transformed, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputCurrentPageDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedText, err := expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponsesText(original["text"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedText); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["text"] = transformedText
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowCXTestCaseTestCaseConversationTurnsVirtualAgentOutputTextResponsesText(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
