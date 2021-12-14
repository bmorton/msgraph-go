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

// InlineObject33 struct for InlineObject33
type InlineObject33 struct {
	ClientContext NullableString `json:"clientContext,omitempty"`
}

// NewInlineObject33 instantiates a new InlineObject33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject33() *InlineObject33 {
	this := InlineObject33{}
	return &this
}

// NewInlineObject33WithDefaults instantiates a new InlineObject33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject33WithDefaults() *InlineObject33 {
	this := InlineObject33{}
	return &this
}

// GetClientContext returns the ClientContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject33) GetClientContext() string {
	if o == nil || o.ClientContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientContext.Get()
}

// GetClientContextOk returns a tuple with the ClientContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject33) GetClientContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientContext.Get(), o.ClientContext.IsSet()
}

// HasClientContext returns a boolean if a field has been set.
func (o *InlineObject33) HasClientContext() bool {
	if o != nil && o.ClientContext.IsSet() {
		return true
	}

	return false
}

// SetClientContext gets a reference to the given NullableString and assigns it to the ClientContext field.
func (o *InlineObject33) SetClientContext(v string) {
	o.ClientContext.Set(&v)
}
// SetClientContextNil sets the value for ClientContext to be an explicit nil
func (o *InlineObject33) SetClientContextNil() {
	o.ClientContext.Set(nil)
}

// UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
func (o *InlineObject33) UnsetClientContext() {
	o.ClientContext.Unset()
}

func (o InlineObject33) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientContext.IsSet() {
		toSerialize["clientContext"] = o.ClientContext.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject33 struct {
	value *InlineObject33
	isSet bool
}

func (v NullableInlineObject33) Get() *InlineObject33 {
	return v.value
}

func (v *NullableInlineObject33) Set(val *InlineObject33) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject33) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject33(val *InlineObject33) *NullableInlineObject33 {
	return &NullableInlineObject33{value: val, isSet: true}
}

func (v NullableInlineObject33) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


