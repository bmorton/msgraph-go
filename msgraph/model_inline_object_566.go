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

// InlineObject566 struct for InlineObject566
type InlineObject566 struct {
	DestinationId *string `json:"DestinationId,omitempty"`
}

// NewInlineObject566 instantiates a new InlineObject566 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject566() *InlineObject566 {
	this := InlineObject566{}
	return &this
}

// NewInlineObject566WithDefaults instantiates a new InlineObject566 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject566WithDefaults() *InlineObject566 {
	this := InlineObject566{}
	return &this
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *InlineObject566) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject566) GetDestinationIdOk() (*string, bool) {
	if o == nil || o.DestinationId == nil {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *InlineObject566) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *InlineObject566) SetDestinationId(v string) {
	o.DestinationId = &v
}

func (o InlineObject566) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationId != nil {
		toSerialize["DestinationId"] = o.DestinationId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject566 struct {
	value *InlineObject566
	isSet bool
}

func (v NullableInlineObject566) Get() *InlineObject566 {
	return v.value
}

func (v *NullableInlineObject566) Set(val *InlineObject566) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject566) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject566) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject566(val *InlineObject566) *NullableInlineObject566 {
	return &NullableInlineObject566{value: val, isSet: true}
}

func (v NullableInlineObject566) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject566) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


