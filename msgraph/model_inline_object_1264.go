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

// InlineObject1264 struct for InlineObject1264
type InlineObject1264 struct {
	Alpha AnyOfobject `json:"alpha,omitempty"`
	StandardDev AnyOfobject `json:"standardDev,omitempty"`
	Size AnyOfobject `json:"size,omitempty"`
}

// NewInlineObject1264 instantiates a new InlineObject1264 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1264() *InlineObject1264 {
	this := InlineObject1264{}
	return &this
}

// NewInlineObject1264WithDefaults instantiates a new InlineObject1264 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1264WithDefaults() *InlineObject1264 {
	this := InlineObject1264{}
	return &this
}

// GetAlpha returns the Alpha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1264) GetAlpha() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1264) GetAlphaOk() (*AnyOfobject, bool) {
	if o == nil || o.Alpha == nil {
		return nil, false
	}
	return &o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *InlineObject1264) HasAlpha() bool {
	if o != nil && o.Alpha != nil {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given AnyOfobject and assigns it to the Alpha field.
func (o *InlineObject1264) SetAlpha(v AnyOfobject) {
	o.Alpha = v
}

// GetStandardDev returns the StandardDev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1264) GetStandardDev() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StandardDev
}

// GetStandardDevOk returns a tuple with the StandardDev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1264) GetStandardDevOk() (*AnyOfobject, bool) {
	if o == nil || o.StandardDev == nil {
		return nil, false
	}
	return &o.StandardDev, true
}

// HasStandardDev returns a boolean if a field has been set.
func (o *InlineObject1264) HasStandardDev() bool {
	if o != nil && o.StandardDev != nil {
		return true
	}

	return false
}

// SetStandardDev gets a reference to the given AnyOfobject and assigns it to the StandardDev field.
func (o *InlineObject1264) SetStandardDev(v AnyOfobject) {
	o.StandardDev = v
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1264) GetSize() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1264) GetSizeOk() (*AnyOfobject, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return &o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InlineObject1264) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given AnyOfobject and assigns it to the Size field.
func (o *InlineObject1264) SetSize(v AnyOfobject) {
	o.Size = v
}

func (o InlineObject1264) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alpha != nil {
		toSerialize["alpha"] = o.Alpha
	}
	if o.StandardDev != nil {
		toSerialize["standardDev"] = o.StandardDev
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1264 struct {
	value *InlineObject1264
	isSet bool
}

func (v NullableInlineObject1264) Get() *InlineObject1264 {
	return v.value
}

func (v *NullableInlineObject1264) Set(val *InlineObject1264) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1264) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1264) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1264(val *InlineObject1264) *NullableInlineObject1264 {
	return &NullableInlineObject1264{value: val, isSet: true}
}

func (v NullableInlineObject1264) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1264) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


