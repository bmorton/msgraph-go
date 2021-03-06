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

// InlineObject1600 struct for InlineObject1600
type InlineObject1600 struct {
	Criteria *string `json:"criteria,omitempty"`
}

// NewInlineObject1600 instantiates a new InlineObject1600 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1600() *InlineObject1600 {
	this := InlineObject1600{}
	return &this
}

// NewInlineObject1600WithDefaults instantiates a new InlineObject1600 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1600WithDefaults() *InlineObject1600 {
	this := InlineObject1600{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *InlineObject1600) GetCriteria() string {
	if o == nil || o.Criteria == nil {
		var ret string
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1600) GetCriteriaOk() (*string, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject1600) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given string and assigns it to the Criteria field.
func (o *InlineObject1600) SetCriteria(v string) {
	o.Criteria = &v
}

func (o InlineObject1600) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1600 struct {
	value *InlineObject1600
	isSet bool
}

func (v NullableInlineObject1600) Get() *InlineObject1600 {
	return v.value
}

func (v *NullableInlineObject1600) Set(val *InlineObject1600) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1600) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1600) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1600(val *InlineObject1600) *NullableInlineObject1600 {
	return &NullableInlineObject1600{value: val, isSet: true}
}

func (v NullableInlineObject1600) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1600) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


