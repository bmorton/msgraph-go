/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MicrosoftGraphRiskEventType the model 'MicrosoftGraphRiskEventType'
type MicrosoftGraphRiskEventType string

// List of microsoft.graph.riskEventType
const (
	UNLIKELY_TRAVEL MicrosoftGraphRiskEventType = "unlikelyTravel"
	ANONYMIZED_IP_ADDRESS MicrosoftGraphRiskEventType = "anonymizedIPAddress"
	MALICIOUS_IP_ADDRESS MicrosoftGraphRiskEventType = "maliciousIPAddress"
	UNFAMILIAR_FEATURES MicrosoftGraphRiskEventType = "unfamiliarFeatures"
	MALWARE_INFECTED_IP_ADDRESS MicrosoftGraphRiskEventType = "malwareInfectedIPAddress"
	SUSPICIOUS_IP_ADDRESS MicrosoftGraphRiskEventType = "suspiciousIPAddress"
	LEAKED_CREDENTIALS MicrosoftGraphRiskEventType = "leakedCredentials"
	INVESTIGATIONS_THREAT_INTELLIGENCE MicrosoftGraphRiskEventType = "investigationsThreatIntelligence"
	GENERIC MicrosoftGraphRiskEventType = "generic"
	ADMIN_CONFIRMED_USER_COMPROMISED MicrosoftGraphRiskEventType = "adminConfirmedUserCompromised"
	MCAS_IMPOSSIBLE_TRAVEL MicrosoftGraphRiskEventType = "mcasImpossibleTravel"
	MCAS_SUSPICIOUS_INBOX_MANIPULATION_RULES MicrosoftGraphRiskEventType = "mcasSuspiciousInboxManipulationRules"
	INVESTIGATIONS_THREAT_INTELLIGENCE_SIGNIN_LINKED MicrosoftGraphRiskEventType = "investigationsThreatIntelligenceSigninLinked"
	MALICIOUS_IP_ADDRESS_VALID_CREDENTIALS_BLOCKED_IP MicrosoftGraphRiskEventType = "maliciousIPAddressValidCredentialsBlockedIP"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphRiskEventType = "unknownFutureValue"
)

// All allowed values of MicrosoftGraphRiskEventType enum
var AllowedMicrosoftGraphRiskEventTypeEnumValues = []MicrosoftGraphRiskEventType{
	"unlikelyTravel",
	"anonymizedIPAddress",
	"maliciousIPAddress",
	"unfamiliarFeatures",
	"malwareInfectedIPAddress",
	"suspiciousIPAddress",
	"leakedCredentials",
	"investigationsThreatIntelligence",
	"generic",
	"adminConfirmedUserCompromised",
	"mcasImpossibleTravel",
	"mcasSuspiciousInboxManipulationRules",
	"investigationsThreatIntelligenceSigninLinked",
	"maliciousIPAddressValidCredentialsBlockedIP",
	"unknownFutureValue",
}

func (v *MicrosoftGraphRiskEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftGraphRiskEventType(value)
	for _, existing := range AllowedMicrosoftGraphRiskEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftGraphRiskEventType", value)
}

// NewMicrosoftGraphRiskEventTypeFromValue returns a pointer to a valid MicrosoftGraphRiskEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftGraphRiskEventTypeFromValue(v string) (*MicrosoftGraphRiskEventType, error) {
	ev := MicrosoftGraphRiskEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftGraphRiskEventType: valid values are %v", v, AllowedMicrosoftGraphRiskEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftGraphRiskEventType) IsValid() bool {
	for _, existing := range AllowedMicrosoftGraphRiskEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to microsoft.graph.riskEventType value
func (v MicrosoftGraphRiskEventType) Ptr() *MicrosoftGraphRiskEventType {
	return &v
}

type NullableMicrosoftGraphRiskEventType struct {
	value *MicrosoftGraphRiskEventType
	isSet bool
}

func (v NullableMicrosoftGraphRiskEventType) Get() *MicrosoftGraphRiskEventType {
	return v.value
}

func (v *NullableMicrosoftGraphRiskEventType) Set(val *MicrosoftGraphRiskEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRiskEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRiskEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRiskEventType(val *MicrosoftGraphRiskEventType) *NullableMicrosoftGraphRiskEventType {
	return &NullableMicrosoftGraphRiskEventType{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRiskEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRiskEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
