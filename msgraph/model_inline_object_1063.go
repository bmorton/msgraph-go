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

// InlineObject1063 struct for InlineObject1063
type InlineObject1063 struct {
	DestinationId *string `json:"DestinationId,omitempty"`
}

// NewInlineObject1063 instantiates a new InlineObject1063 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1063() *InlineObject1063 {
	this := InlineObject1063{}
	return &this
}

// NewInlineObject1063WithDefaults instantiates a new InlineObject1063 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1063WithDefaults() *InlineObject1063 {
	this := InlineObject1063{}
	return &this
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *InlineObject1063) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1063) GetDestinationIdOk() (*string, bool) {
	if o == nil || o.DestinationId == nil {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *InlineObject1063) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *InlineObject1063) SetDestinationId(v string) {
	o.DestinationId = &v
}

func (o InlineObject1063) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationId != nil {
		toSerialize["DestinationId"] = o.DestinationId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1063 struct {
	value *InlineObject1063
	isSet bool
}

func (v NullableInlineObject1063) Get() *InlineObject1063 {
	return v.value
}

func (v *NullableInlineObject1063) Set(val *InlineObject1063) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1063) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1063) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1063(val *InlineObject1063) *NullableInlineObject1063 {
	return &NullableInlineObject1063{value: val, isSet: true}
}

func (v NullableInlineObject1063) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1063) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


