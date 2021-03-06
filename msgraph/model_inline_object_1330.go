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

// InlineObject1330 struct for InlineObject1330
type InlineObject1330 struct {
	X AnyOfobject `json:"x,omitempty"`
	DegFreedom1 AnyOfobject `json:"degFreedom1,omitempty"`
	DegFreedom2 AnyOfobject `json:"degFreedom2,omitempty"`
	Cumulative AnyOfobject `json:"cumulative,omitempty"`
}

// NewInlineObject1330 instantiates a new InlineObject1330 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1330() *InlineObject1330 {
	this := InlineObject1330{}
	return &this
}

// NewInlineObject1330WithDefaults instantiates a new InlineObject1330 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1330WithDefaults() *InlineObject1330 {
	this := InlineObject1330{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1330) GetX() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1330) GetXOk() (*AnyOfobject, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return &o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1330) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1330) SetX(v AnyOfobject) {
	o.X = v
}

// GetDegFreedom1 returns the DegFreedom1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1330) GetDegFreedom1() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.DegFreedom1
}

// GetDegFreedom1Ok returns a tuple with the DegFreedom1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1330) GetDegFreedom1Ok() (*AnyOfobject, bool) {
	if o == nil || o.DegFreedom1 == nil {
		return nil, false
	}
	return &o.DegFreedom1, true
}

// HasDegFreedom1 returns a boolean if a field has been set.
func (o *InlineObject1330) HasDegFreedom1() bool {
	if o != nil && o.DegFreedom1 != nil {
		return true
	}

	return false
}

// SetDegFreedom1 gets a reference to the given AnyOfobject and assigns it to the DegFreedom1 field.
func (o *InlineObject1330) SetDegFreedom1(v AnyOfobject) {
	o.DegFreedom1 = v
}

// GetDegFreedom2 returns the DegFreedom2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1330) GetDegFreedom2() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.DegFreedom2
}

// GetDegFreedom2Ok returns a tuple with the DegFreedom2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1330) GetDegFreedom2Ok() (*AnyOfobject, bool) {
	if o == nil || o.DegFreedom2 == nil {
		return nil, false
	}
	return &o.DegFreedom2, true
}

// HasDegFreedom2 returns a boolean if a field has been set.
func (o *InlineObject1330) HasDegFreedom2() bool {
	if o != nil && o.DegFreedom2 != nil {
		return true
	}

	return false
}

// SetDegFreedom2 gets a reference to the given AnyOfobject and assigns it to the DegFreedom2 field.
func (o *InlineObject1330) SetDegFreedom2(v AnyOfobject) {
	o.DegFreedom2 = v
}

// GetCumulative returns the Cumulative field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1330) GetCumulative() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Cumulative
}

// GetCumulativeOk returns a tuple with the Cumulative field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1330) GetCumulativeOk() (*AnyOfobject, bool) {
	if o == nil || o.Cumulative == nil {
		return nil, false
	}
	return &o.Cumulative, true
}

// HasCumulative returns a boolean if a field has been set.
func (o *InlineObject1330) HasCumulative() bool {
	if o != nil && o.Cumulative != nil {
		return true
	}

	return false
}

// SetCumulative gets a reference to the given AnyOfobject and assigns it to the Cumulative field.
func (o *InlineObject1330) SetCumulative(v AnyOfobject) {
	o.Cumulative = v
}

func (o InlineObject1330) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.DegFreedom1 != nil {
		toSerialize["degFreedom1"] = o.DegFreedom1
	}
	if o.DegFreedom2 != nil {
		toSerialize["degFreedom2"] = o.DegFreedom2
	}
	if o.Cumulative != nil {
		toSerialize["cumulative"] = o.Cumulative
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1330 struct {
	value *InlineObject1330
	isSet bool
}

func (v NullableInlineObject1330) Get() *InlineObject1330 {
	return v.value
}

func (v *NullableInlineObject1330) Set(val *InlineObject1330) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1330) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1330) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1330(val *InlineObject1330) *NullableInlineObject1330 {
	return &NullableInlineObject1330{value: val, isSet: true}
}

func (v NullableInlineObject1330) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1330) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


