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

// InlineObject1639 struct for InlineObject1639
type InlineObject1639 struct {
	Criteria1 NullableString `json:"criteria1,omitempty"`
	Criteria2 NullableString `json:"criteria2,omitempty"`
	Oper *string `json:"oper,omitempty"`
}

// NewInlineObject1639 instantiates a new InlineObject1639 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1639() *InlineObject1639 {
	this := InlineObject1639{}
	return &this
}

// NewInlineObject1639WithDefaults instantiates a new InlineObject1639 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1639WithDefaults() *InlineObject1639 {
	this := InlineObject1639{}
	return &this
}

// GetCriteria1 returns the Criteria1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1639) GetCriteria1() string {
	if o == nil || o.Criteria1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Criteria1.Get()
}

// GetCriteria1Ok returns a tuple with the Criteria1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1639) GetCriteria1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Criteria1.Get(), o.Criteria1.IsSet()
}

// HasCriteria1 returns a boolean if a field has been set.
func (o *InlineObject1639) HasCriteria1() bool {
	if o != nil && o.Criteria1.IsSet() {
		return true
	}

	return false
}

// SetCriteria1 gets a reference to the given NullableString and assigns it to the Criteria1 field.
func (o *InlineObject1639) SetCriteria1(v string) {
	o.Criteria1.Set(&v)
}
// SetCriteria1Nil sets the value for Criteria1 to be an explicit nil
func (o *InlineObject1639) SetCriteria1Nil() {
	o.Criteria1.Set(nil)
}

// UnsetCriteria1 ensures that no value is present for Criteria1, not even an explicit nil
func (o *InlineObject1639) UnsetCriteria1() {
	o.Criteria1.Unset()
}

// GetCriteria2 returns the Criteria2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1639) GetCriteria2() string {
	if o == nil || o.Criteria2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Criteria2.Get()
}

// GetCriteria2Ok returns a tuple with the Criteria2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1639) GetCriteria2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Criteria2.Get(), o.Criteria2.IsSet()
}

// HasCriteria2 returns a boolean if a field has been set.
func (o *InlineObject1639) HasCriteria2() bool {
	if o != nil && o.Criteria2.IsSet() {
		return true
	}

	return false
}

// SetCriteria2 gets a reference to the given NullableString and assigns it to the Criteria2 field.
func (o *InlineObject1639) SetCriteria2(v string) {
	o.Criteria2.Set(&v)
}
// SetCriteria2Nil sets the value for Criteria2 to be an explicit nil
func (o *InlineObject1639) SetCriteria2Nil() {
	o.Criteria2.Set(nil)
}

// UnsetCriteria2 ensures that no value is present for Criteria2, not even an explicit nil
func (o *InlineObject1639) UnsetCriteria2() {
	o.Criteria2.Unset()
}

// GetOper returns the Oper field value if set, zero value otherwise.
func (o *InlineObject1639) GetOper() string {
	if o == nil || o.Oper == nil {
		var ret string
		return ret
	}
	return *o.Oper
}

// GetOperOk returns a tuple with the Oper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1639) GetOperOk() (*string, bool) {
	if o == nil || o.Oper == nil {
		return nil, false
	}
	return o.Oper, true
}

// HasOper returns a boolean if a field has been set.
func (o *InlineObject1639) HasOper() bool {
	if o != nil && o.Oper != nil {
		return true
	}

	return false
}

// SetOper gets a reference to the given string and assigns it to the Oper field.
func (o *InlineObject1639) SetOper(v string) {
	o.Oper = &v
}

func (o InlineObject1639) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Criteria1.IsSet() {
		toSerialize["criteria1"] = o.Criteria1.Get()
	}
	if o.Criteria2.IsSet() {
		toSerialize["criteria2"] = o.Criteria2.Get()
	}
	if o.Oper != nil {
		toSerialize["oper"] = o.Oper
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1639 struct {
	value *InlineObject1639
	isSet bool
}

func (v NullableInlineObject1639) Get() *InlineObject1639 {
	return v.value
}

func (v *NullableInlineObject1639) Set(val *InlineObject1639) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1639) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1639) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1639(val *InlineObject1639) *NullableInlineObject1639 {
	return &NullableInlineObject1639{value: val, isSet: true}
}

func (v NullableInlineObject1639) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1639) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


