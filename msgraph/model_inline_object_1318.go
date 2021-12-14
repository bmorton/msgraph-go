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

// InlineObject1318 struct for InlineObject1318
type InlineObject1318 struct {
	StartDate AnyOfobject `json:"startDate,omitempty"`
	Months AnyOfobject `json:"months,omitempty"`
}

// NewInlineObject1318 instantiates a new InlineObject1318 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1318() *InlineObject1318 {
	this := InlineObject1318{}
	return &this
}

// NewInlineObject1318WithDefaults instantiates a new InlineObject1318 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1318WithDefaults() *InlineObject1318 {
	this := InlineObject1318{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1318) GetStartDate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1318) GetStartDateOk() (*AnyOfobject, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineObject1318) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given AnyOfobject and assigns it to the StartDate field.
func (o *InlineObject1318) SetStartDate(v AnyOfobject) {
	o.StartDate = v
}

// GetMonths returns the Months field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1318) GetMonths() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1318) GetMonthsOk() (*AnyOfobject, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return &o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *InlineObject1318) HasMonths() bool {
	if o != nil && o.Months != nil {
		return true
	}

	return false
}

// SetMonths gets a reference to the given AnyOfobject and assigns it to the Months field.
func (o *InlineObject1318) SetMonths(v AnyOfobject) {
	o.Months = v
}

func (o InlineObject1318) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1318 struct {
	value *InlineObject1318
	isSet bool
}

func (v NullableInlineObject1318) Get() *InlineObject1318 {
	return v.value
}

func (v *NullableInlineObject1318) Set(val *InlineObject1318) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1318) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1318) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1318(val *InlineObject1318) *NullableInlineObject1318 {
	return &NullableInlineObject1318{value: val, isSet: true}
}

func (v NullableInlineObject1318) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1318) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


