/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InformationProtection struct for InformationProtection
type InformationProtection struct {
	Bitlocker AnyOfmicrosoftGraphBitlocker `json:"bitlocker,omitempty"`
	ThreatAssessmentRequests *[]MicrosoftGraphThreatAssessmentRequest `json:"threatAssessmentRequests,omitempty"`
}

// NewInformationProtection instantiates a new InformationProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInformationProtection() *InformationProtection {
	this := InformationProtection{}
	return &this
}

// NewInformationProtectionWithDefaults instantiates a new InformationProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInformationProtectionWithDefaults() *InformationProtection {
	this := InformationProtection{}
	return &this
}

// GetBitlocker returns the Bitlocker field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InformationProtection) GetBitlocker() AnyOfmicrosoftGraphBitlocker {
	if o == nil  {
		var ret AnyOfmicrosoftGraphBitlocker
		return ret
	}
	return o.Bitlocker
}

// GetBitlockerOk returns a tuple with the Bitlocker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InformationProtection) GetBitlockerOk() (*AnyOfmicrosoftGraphBitlocker, bool) {
	if o == nil || o.Bitlocker == nil {
		return nil, false
	}
	return &o.Bitlocker, true
}

// HasBitlocker returns a boolean if a field has been set.
func (o *InformationProtection) HasBitlocker() bool {
	if o != nil && o.Bitlocker != nil {
		return true
	}

	return false
}

// SetBitlocker gets a reference to the given AnyOfmicrosoftGraphBitlocker and assigns it to the Bitlocker field.
func (o *InformationProtection) SetBitlocker(v AnyOfmicrosoftGraphBitlocker) {
	o.Bitlocker = v
}

// GetThreatAssessmentRequests returns the ThreatAssessmentRequests field value if set, zero value otherwise.
func (o *InformationProtection) GetThreatAssessmentRequests() []MicrosoftGraphThreatAssessmentRequest {
	if o == nil || o.ThreatAssessmentRequests == nil {
		var ret []MicrosoftGraphThreatAssessmentRequest
		return ret
	}
	return *o.ThreatAssessmentRequests
}

// GetThreatAssessmentRequestsOk returns a tuple with the ThreatAssessmentRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InformationProtection) GetThreatAssessmentRequestsOk() (*[]MicrosoftGraphThreatAssessmentRequest, bool) {
	if o == nil || o.ThreatAssessmentRequests == nil {
		return nil, false
	}
	return o.ThreatAssessmentRequests, true
}

// HasThreatAssessmentRequests returns a boolean if a field has been set.
func (o *InformationProtection) HasThreatAssessmentRequests() bool {
	if o != nil && o.ThreatAssessmentRequests != nil {
		return true
	}

	return false
}

// SetThreatAssessmentRequests gets a reference to the given []MicrosoftGraphThreatAssessmentRequest and assigns it to the ThreatAssessmentRequests field.
func (o *InformationProtection) SetThreatAssessmentRequests(v []MicrosoftGraphThreatAssessmentRequest) {
	o.ThreatAssessmentRequests = &v
}

func (o InformationProtection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bitlocker != nil {
		toSerialize["bitlocker"] = o.Bitlocker
	}
	if o.ThreatAssessmentRequests != nil {
		toSerialize["threatAssessmentRequests"] = o.ThreatAssessmentRequests
	}
	return json.Marshal(toSerialize)
}

type NullableInformationProtection struct {
	value *InformationProtection
	isSet bool
}

func (v NullableInformationProtection) Get() *InformationProtection {
	return v.value
}

func (v *NullableInformationProtection) Set(val *InformationProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableInformationProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableInformationProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInformationProtection(val *InformationProtection) *NullableInformationProtection {
	return &NullableInformationProtection{value: val, isSet: true}
}

func (v NullableInformationProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInformationProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

