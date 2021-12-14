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

// InlineObject1289 struct for InlineObject1289
type InlineObject1289 struct {
	EndDate AnyOfobject `json:"endDate,omitempty"`
	StartDate AnyOfobject `json:"startDate,omitempty"`
}

// NewInlineObject1289 instantiates a new InlineObject1289 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1289() *InlineObject1289 {
	this := InlineObject1289{}
	return &this
}

// NewInlineObject1289WithDefaults instantiates a new InlineObject1289 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1289WithDefaults() *InlineObject1289 {
	this := InlineObject1289{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1289) GetEndDate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1289) GetEndDateOk() (*AnyOfobject, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InlineObject1289) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given AnyOfobject and assigns it to the EndDate field.
func (o *InlineObject1289) SetEndDate(v AnyOfobject) {
	o.EndDate = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1289) GetStartDate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1289) GetStartDateOk() (*AnyOfobject, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineObject1289) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given AnyOfobject and assigns it to the StartDate field.
func (o *InlineObject1289) SetStartDate(v AnyOfobject) {
	o.StartDate = v
}

func (o InlineObject1289) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1289 struct {
	value *InlineObject1289
	isSet bool
}

func (v NullableInlineObject1289) Get() *InlineObject1289 {
	return v.value
}

func (v *NullableInlineObject1289) Set(val *InlineObject1289) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1289) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1289) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1289(val *InlineObject1289) *NullableInlineObject1289 {
	return &NullableInlineObject1289{value: val, isSet: true}
}

func (v NullableInlineObject1289) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1289) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

