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

// InlineObject1557 struct for InlineObject1557
type InlineObject1557 struct {
	Rate AnyOfobject `json:"rate,omitempty"`
	Values AnyOfobject `json:"values,omitempty"`
	Dates AnyOfobject `json:"dates,omitempty"`
}

// NewInlineObject1557 instantiates a new InlineObject1557 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1557() *InlineObject1557 {
	this := InlineObject1557{}
	return &this
}

// NewInlineObject1557WithDefaults instantiates a new InlineObject1557 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1557WithDefaults() *InlineObject1557 {
	this := InlineObject1557{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1557) GetRate() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1557) GetRateOk() (*AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return &o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1557) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1557) SetRate(v AnyOfobject) {
	o.Rate = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1557) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1557) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1557) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1557) SetValues(v AnyOfobject) {
	o.Values = v
}

// GetDates returns the Dates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1557) GetDates() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1557) GetDatesOk() (*AnyOfobject, bool) {
	if o == nil || o.Dates == nil {
		return nil, false
	}
	return &o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *InlineObject1557) HasDates() bool {
	if o != nil && o.Dates != nil {
		return true
	}

	return false
}

// SetDates gets a reference to the given AnyOfobject and assigns it to the Dates field.
func (o *InlineObject1557) SetDates(v AnyOfobject) {
	o.Dates = v
}

func (o InlineObject1557) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Dates != nil {
		toSerialize["dates"] = o.Dates
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1557 struct {
	value *InlineObject1557
	isSet bool
}

func (v NullableInlineObject1557) Get() *InlineObject1557 {
	return v.value
}

func (v *NullableInlineObject1557) Set(val *InlineObject1557) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1557) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1557) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1557(val *InlineObject1557) *NullableInlineObject1557 {
	return &NullableInlineObject1557{value: val, isSet: true}
}

func (v NullableInlineObject1557) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1557) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


