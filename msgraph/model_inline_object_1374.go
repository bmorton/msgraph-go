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

// InlineObject1374 struct for InlineObject1374
type InlineObject1374 struct {
	Inumber AnyOfobject `json:"inumber,omitempty"`
}

// NewInlineObject1374 instantiates a new InlineObject1374 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1374() *InlineObject1374 {
	this := InlineObject1374{}
	return &this
}

// NewInlineObject1374WithDefaults instantiates a new InlineObject1374 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1374WithDefaults() *InlineObject1374 {
	this := InlineObject1374{}
	return &this
}

// GetInumber returns the Inumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1374) GetInumber() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Inumber
}

// GetInumberOk returns a tuple with the Inumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1374) GetInumberOk() (*AnyOfobject, bool) {
	if o == nil || o.Inumber == nil {
		return nil, false
	}
	return &o.Inumber, true
}

// HasInumber returns a boolean if a field has been set.
func (o *InlineObject1374) HasInumber() bool {
	if o != nil && o.Inumber != nil {
		return true
	}

	return false
}

// SetInumber gets a reference to the given AnyOfobject and assigns it to the Inumber field.
func (o *InlineObject1374) SetInumber(v AnyOfobject) {
	o.Inumber = v
}

func (o InlineObject1374) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inumber != nil {
		toSerialize["inumber"] = o.Inumber
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1374 struct {
	value *InlineObject1374
	isSet bool
}

func (v NullableInlineObject1374) Get() *InlineObject1374 {
	return v.value
}

func (v *NullableInlineObject1374) Set(val *InlineObject1374) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1374) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1374) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1374(val *InlineObject1374) *NullableInlineObject1374 {
	return &NullableInlineObject1374{value: val, isSet: true}
}

func (v NullableInlineObject1374) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1374) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


