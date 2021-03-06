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

// InlineObject618 struct for InlineObject618
type InlineObject618 struct {
	WebUrl NullableString `json:"webUrl,omitempty"`
}

// NewInlineObject618 instantiates a new InlineObject618 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject618() *InlineObject618 {
	this := InlineObject618{}
	return &this
}

// NewInlineObject618WithDefaults instantiates a new InlineObject618 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject618WithDefaults() *InlineObject618 {
	this := InlineObject618{}
	return &this
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject618) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject618) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *InlineObject618) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *InlineObject618) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *InlineObject618) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *InlineObject618) UnsetWebUrl() {
	o.WebUrl.Unset()
}

func (o InlineObject618) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject618 struct {
	value *InlineObject618
	isSet bool
}

func (v NullableInlineObject618) Get() *InlineObject618 {
	return v.value
}

func (v *NullableInlineObject618) Set(val *InlineObject618) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject618) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject618) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject618(val *InlineObject618) *NullableInlineObject618 {
	return &NullableInlineObject618{value: val, isSet: true}
}

func (v NullableInlineObject618) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject618) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


