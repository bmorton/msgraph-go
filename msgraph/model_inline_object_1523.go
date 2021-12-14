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

// InlineObject1523 struct for InlineObject1523
type InlineObject1523 struct {
	X AnyOfobject `json:"x,omitempty"`
	DegFreedom AnyOfobject `json:"degFreedom,omitempty"`
	Cumulative AnyOfobject `json:"cumulative,omitempty"`
}

// NewInlineObject1523 instantiates a new InlineObject1523 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1523() *InlineObject1523 {
	this := InlineObject1523{}
	return &this
}

// NewInlineObject1523WithDefaults instantiates a new InlineObject1523 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1523WithDefaults() *InlineObject1523 {
	this := InlineObject1523{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1523) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1523) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1523) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1523) SetX(v AnyOfobject) {
	o.X = v
}

// GetDegFreedom returns the DegFreedom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1523) GetDegFreedom() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.DegFreedom
}

// GetDegFreedomOk returns a tuple with the DegFreedom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1523) GetDegFreedomOk() (*AnyOfobject, bool) {
	if o == nil || o.DegFreedom == nil {
		return nil, false
	}
	return &o.DegFreedom, true
}

// HasDegFreedom returns a boolean if a field has been set.
func (o *InlineObject1523) HasDegFreedom() bool {
	if o != nil && o.DegFreedom != nil {
		return true
	}

	return false
}

// SetDegFreedom gets a reference to the given AnyOfobject and assigns it to the DegFreedom field.
func (o *InlineObject1523) SetDegFreedom(v AnyOfobject) {
	o.DegFreedom = v
}

// GetCumulative returns the Cumulative field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1523) GetCumulative() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cumulative
}

// GetCumulativeOk returns a tuple with the Cumulative field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1523) GetCumulativeOk() (*AnyOfobject, bool) {
	if o == nil || o.Cumulative == nil {
		return nil, false
	}
	return &o.Cumulative, true
}

// HasCumulative returns a boolean if a field has been set.
func (o *InlineObject1523) HasCumulative() bool {
	if o != nil && o.Cumulative != nil {
		return true
	}

	return false
}

// SetCumulative gets a reference to the given AnyOfobject and assigns it to the Cumulative field.
func (o *InlineObject1523) SetCumulative(v AnyOfobject) {
	o.Cumulative = v
}

func (o InlineObject1523) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.DegFreedom != nil {
		toSerialize["degFreedom"] = o.DegFreedom
	}
	if o.Cumulative != nil {
		toSerialize["cumulative"] = o.Cumulative
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1523 struct {
	value *InlineObject1523
	isSet bool
}

func (v NullableInlineObject1523) Get() *InlineObject1523 {
	return v.value
}

func (v *NullableInlineObject1523) Set(val *InlineObject1523) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1523) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1523) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1523(val *InlineObject1523) *NullableInlineObject1523 {
	return &NullableInlineObject1523{value: val, isSet: true}
}

func (v NullableInlineObject1523) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1523) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


