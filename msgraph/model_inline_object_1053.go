/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// InlineObject1053 struct for InlineObject1053
type InlineObject1053 struct {
	DestinationId *string `json:"DestinationId,omitempty"`
}

// NewInlineObject1053 instantiates a new InlineObject1053 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1053() *InlineObject1053 {
	this := InlineObject1053{}
	return &this
}

// NewInlineObject1053WithDefaults instantiates a new InlineObject1053 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1053WithDefaults() *InlineObject1053 {
	this := InlineObject1053{}
	return &this
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *InlineObject1053) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1053) GetDestinationIdOk() (*string, bool) {
	if o == nil || o.DestinationId == nil {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *InlineObject1053) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *InlineObject1053) SetDestinationId(v string) {
	o.DestinationId = &v
}

func (o InlineObject1053) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationId != nil {
		toSerialize["DestinationId"] = o.DestinationId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1053 struct {
	value *InlineObject1053
	isSet bool
}

func (v NullableInlineObject1053) Get() *InlineObject1053 {
	return v.value
}

func (v *NullableInlineObject1053) Set(val *InlineObject1053) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1053) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1053) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1053(val *InlineObject1053) *NullableInlineObject1053 {
	return &NullableInlineObject1053{value: val, isSet: true}
}

func (v NullableInlineObject1053) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1053) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


