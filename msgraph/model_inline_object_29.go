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

// InlineObject29 struct for InlineObject29
type InlineObject29 struct {
	ClientContext NullableString `json:"clientContext,omitempty"`
}

// NewInlineObject29 instantiates a new InlineObject29 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject29() *InlineObject29 {
	this := InlineObject29{}
	return &this
}

// NewInlineObject29WithDefaults instantiates a new InlineObject29 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject29WithDefaults() *InlineObject29 {
	this := InlineObject29{}
	return &this
}

// GetClientContext returns the ClientContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject29) GetClientContext() string {
	if o == nil || o.ClientContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientContext.Get()
}

// GetClientContextOk returns a tuple with the ClientContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject29) GetClientContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientContext.Get(), o.ClientContext.IsSet()
}

// HasClientContext returns a boolean if a field has been set.
func (o *InlineObject29) HasClientContext() bool {
	if o != nil && o.ClientContext.IsSet() {
		return true
	}

	return false
}

// SetClientContext gets a reference to the given NullableString and assigns it to the ClientContext field.
func (o *InlineObject29) SetClientContext(v string) {
	o.ClientContext.Set(&v)
}
// SetClientContextNil sets the value for ClientContext to be an explicit nil
func (o *InlineObject29) SetClientContextNil() {
	o.ClientContext.Set(nil)
}

// UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
func (o *InlineObject29) UnsetClientContext() {
	o.ClientContext.Unset()
}

func (o InlineObject29) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientContext.IsSet() {
		toSerialize["clientContext"] = o.ClientContext.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject29 struct {
	value *InlineObject29
	isSet bool
}

func (v NullableInlineObject29) Get() *InlineObject29 {
	return v.value
}

func (v *NullableInlineObject29) Set(val *InlineObject29) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject29) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject29) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject29(val *InlineObject29) *NullableInlineObject29 {
	return &NullableInlineObject29{value: val, isSet: true}
}

func (v NullableInlineObject29) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject29) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


