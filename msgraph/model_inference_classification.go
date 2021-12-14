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

// InferenceClassification struct for InferenceClassification
type InferenceClassification struct {
	// A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
	Overrides *[]MicrosoftGraphInferenceClassificationOverride `json:"overrides,omitempty"`
}

// NewInferenceClassification instantiates a new InferenceClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInferenceClassification() *InferenceClassification {
	this := InferenceClassification{}
	return &this
}

// NewInferenceClassificationWithDefaults instantiates a new InferenceClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInferenceClassificationWithDefaults() *InferenceClassification {
	this := InferenceClassification{}
	return &this
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InferenceClassification) GetOverrides() []MicrosoftGraphInferenceClassificationOverride {
	if o == nil || o.Overrides == nil {
		var ret []MicrosoftGraphInferenceClassificationOverride
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InferenceClassification) GetOverridesOk() (*[]MicrosoftGraphInferenceClassificationOverride, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InferenceClassification) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []MicrosoftGraphInferenceClassificationOverride and assigns it to the Overrides field.
func (o *InferenceClassification) SetOverrides(v []MicrosoftGraphInferenceClassificationOverride) {
	o.Overrides = &v
}

func (o InferenceClassification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInferenceClassification struct {
	value *InferenceClassification
	isSet bool
}

func (v NullableInferenceClassification) Get() *InferenceClassification {
	return v.value
}

func (v *NullableInferenceClassification) Set(val *InferenceClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableInferenceClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableInferenceClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInferenceClassification(val *InferenceClassification) *NullableInferenceClassification {
	return &NullableInferenceClassification{value: val, isSet: true}
}

func (v NullableInferenceClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInferenceClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


