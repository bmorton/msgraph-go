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

// InlineObject879 struct for InlineObject879
type InlineObject879 struct {
	NewReminderTime *MicrosoftGraphDateTimeTimeZone `json:"NewReminderTime,omitempty"`
}

// NewInlineObject879 instantiates a new InlineObject879 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject879() *InlineObject879 {
	this := InlineObject879{}
	return &this
}

// NewInlineObject879WithDefaults instantiates a new InlineObject879 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject879WithDefaults() *InlineObject879 {
	this := InlineObject879{}
	return &this
}

// GetNewReminderTime returns the NewReminderTime field value if set, zero value otherwise.
func (o *InlineObject879) GetNewReminderTime() MicrosoftGraphDateTimeTimeZone {
	if o == nil || o.NewReminderTime == nil {
		var ret MicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.NewReminderTime
}

// GetNewReminderTimeOk returns a tuple with the NewReminderTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject879) GetNewReminderTimeOk() (*MicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.NewReminderTime == nil {
		return nil, false
	}
	return o.NewReminderTime, true
}

// HasNewReminderTime returns a boolean if a field has been set.
func (o *InlineObject879) HasNewReminderTime() bool {
	if o != nil && o.NewReminderTime != nil {
		return true
	}

	return false
}

// SetNewReminderTime gets a reference to the given MicrosoftGraphDateTimeTimeZone and assigns it to the NewReminderTime field.
func (o *InlineObject879) SetNewReminderTime(v MicrosoftGraphDateTimeTimeZone) {
	o.NewReminderTime = &v
}

func (o InlineObject879) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewReminderTime != nil {
		toSerialize["NewReminderTime"] = o.NewReminderTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject879 struct {
	value *InlineObject879
	isSet bool
}

func (v NullableInlineObject879) Get() *InlineObject879 {
	return v.value
}

func (v *NullableInlineObject879) Set(val *InlineObject879) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject879) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject879) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject879(val *InlineObject879) *NullableInlineObject879 {
	return &NullableInlineObject879{value: val, isSet: true}
}

func (v NullableInlineObject879) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject879) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


