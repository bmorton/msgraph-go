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

// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	Quality *MicrosoftGraphTeleconferenceDeviceQuality `json:"quality,omitempty"`
}

// NewInlineObject37 instantiates a new InlineObject37 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject37() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// NewInlineObject37WithDefaults instantiates a new InlineObject37 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject37WithDefaults() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *InlineObject37) GetQuality() MicrosoftGraphTeleconferenceDeviceQuality {
	if o == nil || o.Quality == nil {
		var ret MicrosoftGraphTeleconferenceDeviceQuality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetQualityOk() (*MicrosoftGraphTeleconferenceDeviceQuality, bool) {
	if o == nil || o.Quality == nil {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *InlineObject37) HasQuality() bool {
	if o != nil && o.Quality != nil {
		return true
	}

	return false
}

// SetQuality gets a reference to the given MicrosoftGraphTeleconferenceDeviceQuality and assigns it to the Quality field.
func (o *InlineObject37) SetQuality(v MicrosoftGraphTeleconferenceDeviceQuality) {
	o.Quality = &v
}

func (o InlineObject37) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quality != nil {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject37 struct {
	value *InlineObject37
	isSet bool
}

func (v NullableInlineObject37) Get() *InlineObject37 {
	return v.value
}

func (v *NullableInlineObject37) Set(val *InlineObject37) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject37) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject37) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject37(val *InlineObject37) *NullableInlineObject37 {
	return &NullableInlineObject37{value: val, isSet: true}
}

func (v NullableInlineObject37) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject37) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

