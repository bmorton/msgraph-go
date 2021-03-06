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

// InlineObject1273 struct for InlineObject1273
type InlineObject1273 struct {
	Range AnyOfobject `json:"range,omitempty"`
	Criteria AnyOfobject `json:"criteria,omitempty"`
}

// NewInlineObject1273 instantiates a new InlineObject1273 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1273() *InlineObject1273 {
	this := InlineObject1273{}
	return &this
}

// NewInlineObject1273WithDefaults instantiates a new InlineObject1273 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1273WithDefaults() *InlineObject1273 {
	this := InlineObject1273{}
	return &this
}

// GetRange returns the Range field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1273) GetRange() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1273) GetRangeOk() (*AnyOfobject, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return &o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *InlineObject1273) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given AnyOfobject and assigns it to the Range field.
func (o *InlineObject1273) SetRange(v AnyOfobject) {
	o.Range = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1273) GetCriteria() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1273) GetCriteriaOk() (*AnyOfobject, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject1273) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfobject and assigns it to the Criteria field.
func (o *InlineObject1273) SetCriteria(v AnyOfobject) {
	o.Criteria = v
}

func (o InlineObject1273) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1273 struct {
	value *InlineObject1273
	isSet bool
}

func (v NullableInlineObject1273) Get() *InlineObject1273 {
	return v.value
}

func (v *NullableInlineObject1273) Set(val *InlineObject1273) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1273) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1273) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1273(val *InlineObject1273) *NullableInlineObject1273 {
	return &NullableInlineObject1273{value: val, isSet: true}
}

func (v NullableInlineObject1273) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1273) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


