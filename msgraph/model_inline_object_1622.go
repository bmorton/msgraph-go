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

// InlineObject1622 struct for InlineObject1622
type InlineObject1622 struct {
	Address NullableString `json:"address,omitempty"`
	HasHeaders *bool `json:"hasHeaders,omitempty"`
}

// NewInlineObject1622 instantiates a new InlineObject1622 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1622() *InlineObject1622 {
	this := InlineObject1622{}
	var hasHeaders bool = false
	this.HasHeaders = &hasHeaders
	return &this
}

// NewInlineObject1622WithDefaults instantiates a new InlineObject1622 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1622WithDefaults() *InlineObject1622 {
	this := InlineObject1622{}
	var hasHeaders bool = false
	this.HasHeaders = &hasHeaders
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1622) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1622) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineObject1622) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *InlineObject1622) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *InlineObject1622) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *InlineObject1622) UnsetAddress() {
	o.Address.Unset()
}

// GetHasHeaders returns the HasHeaders field value if set, zero value otherwise.
func (o *InlineObject1622) GetHasHeaders() bool {
	if o == nil || o.HasHeaders == nil {
		var ret bool
		return ret
	}
	return *o.HasHeaders
}

// GetHasHeadersOk returns a tuple with the HasHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1622) GetHasHeadersOk() (*bool, bool) {
	if o == nil || o.HasHeaders == nil {
		return nil, false
	}
	return o.HasHeaders, true
}

// HasHasHeaders returns a boolean if a field has been set.
func (o *InlineObject1622) HasHasHeaders() bool {
	if o != nil && o.HasHeaders != nil {
		return true
	}

	return false
}

// SetHasHeaders gets a reference to the given bool and assigns it to the HasHeaders field.
func (o *InlineObject1622) SetHasHeaders(v bool) {
	o.HasHeaders = &v
}

func (o InlineObject1622) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.HasHeaders != nil {
		toSerialize["hasHeaders"] = o.HasHeaders
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1622 struct {
	value *InlineObject1622
	isSet bool
}

func (v NullableInlineObject1622) Get() *InlineObject1622 {
	return v.value
}

func (v *NullableInlineObject1622) Set(val *InlineObject1622) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1622) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1622) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1622(val *InlineObject1622) *NullableInlineObject1622 {
	return &NullableInlineObject1622{value: val, isSet: true}
}

func (v NullableInlineObject1622) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1622) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


