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

// InlineObject385 struct for InlineObject385
type InlineObject385 struct {
	NewReminderTime *MicrosoftGraphDateTimeTimeZone `json:"NewReminderTime,omitempty"`
}

// NewInlineObject385 instantiates a new InlineObject385 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject385() *InlineObject385 {
	this := InlineObject385{}
	return &this
}

// NewInlineObject385WithDefaults instantiates a new InlineObject385 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject385WithDefaults() *InlineObject385 {
	this := InlineObject385{}
	return &this
}

// GetNewReminderTime returns the NewReminderTime field value if set, zero value otherwise.
func (o *InlineObject385) GetNewReminderTime() MicrosoftGraphDateTimeTimeZone {
	if o == nil || o.NewReminderTime == nil {
		var ret MicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.NewReminderTime
}

// GetNewReminderTimeOk returns a tuple with the NewReminderTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject385) GetNewReminderTimeOk() (*MicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.NewReminderTime == nil {
		return nil, false
	}
	return o.NewReminderTime, true
}

// HasNewReminderTime returns a boolean if a field has been set.
func (o *InlineObject385) HasNewReminderTime() bool {
	if o != nil && o.NewReminderTime != nil {
		return true
	}

	return false
}

// SetNewReminderTime gets a reference to the given MicrosoftGraphDateTimeTimeZone and assigns it to the NewReminderTime field.
func (o *InlineObject385) SetNewReminderTime(v MicrosoftGraphDateTimeTimeZone) {
	o.NewReminderTime = &v
}

func (o InlineObject385) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewReminderTime != nil {
		toSerialize["NewReminderTime"] = o.NewReminderTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject385 struct {
	value *InlineObject385
	isSet bool
}

func (v NullableInlineObject385) Get() *InlineObject385 {
	return v.value
}

func (v *NullableInlineObject385) Set(val *InlineObject385) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject385) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject385) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject385(val *InlineObject385) *NullableInlineObject385 {
	return &NullableInlineObject385{value: val, isSet: true}
}

func (v NullableInlineObject385) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject385) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

