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

// MicrosoftGraphPrintMargin struct for MicrosoftGraphPrintMargin
type MicrosoftGraphPrintMargin struct {
	// The margin in microns from the bottom edge.
	Bottom NullableInt32 `json:"bottom,omitempty"`
	// The margin in microns from the left edge.
	Left NullableInt32 `json:"left,omitempty"`
	// The margin in microns from the right edge.
	Right NullableInt32 `json:"right,omitempty"`
	// The margin in microns from the top edge.
	Top NullableInt32 `json:"top,omitempty"`
}

// NewMicrosoftGraphPrintMargin instantiates a new MicrosoftGraphPrintMargin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintMargin() *MicrosoftGraphPrintMargin {
	this := MicrosoftGraphPrintMargin{}
	return &this
}

// NewMicrosoftGraphPrintMarginWithDefaults instantiates a new MicrosoftGraphPrintMargin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintMarginWithDefaults() *MicrosoftGraphPrintMargin {
	this := MicrosoftGraphPrintMargin{}
	return &this
}

// GetBottom returns the Bottom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintMargin) GetBottom() int32 {
	if o == nil || o.Bottom.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Bottom.Get()
}

// GetBottomOk returns a tuple with the Bottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintMargin) GetBottomOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bottom.Get(), o.Bottom.IsSet()
}

// HasBottom returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintMargin) HasBottom() bool {
	if o != nil && o.Bottom.IsSet() {
		return true
	}

	return false
}

// SetBottom gets a reference to the given NullableInt32 and assigns it to the Bottom field.
func (o *MicrosoftGraphPrintMargin) SetBottom(v int32) {
	o.Bottom.Set(&v)
}
// SetBottomNil sets the value for Bottom to be an explicit nil
func (o *MicrosoftGraphPrintMargin) SetBottomNil() {
	o.Bottom.Set(nil)
}

// UnsetBottom ensures that no value is present for Bottom, not even an explicit nil
func (o *MicrosoftGraphPrintMargin) UnsetBottom() {
	o.Bottom.Unset()
}

// GetLeft returns the Left field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintMargin) GetLeft() int32 {
	if o == nil || o.Left.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Left.Get()
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintMargin) GetLeftOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Left.Get(), o.Left.IsSet()
}

// HasLeft returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintMargin) HasLeft() bool {
	if o != nil && o.Left.IsSet() {
		return true
	}

	return false
}

// SetLeft gets a reference to the given NullableInt32 and assigns it to the Left field.
func (o *MicrosoftGraphPrintMargin) SetLeft(v int32) {
	o.Left.Set(&v)
}
// SetLeftNil sets the value for Left to be an explicit nil
func (o *MicrosoftGraphPrintMargin) SetLeftNil() {
	o.Left.Set(nil)
}

// UnsetLeft ensures that no value is present for Left, not even an explicit nil
func (o *MicrosoftGraphPrintMargin) UnsetLeft() {
	o.Left.Unset()
}

// GetRight returns the Right field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintMargin) GetRight() int32 {
	if o == nil || o.Right.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Right.Get()
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintMargin) GetRightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Right.Get(), o.Right.IsSet()
}

// HasRight returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintMargin) HasRight() bool {
	if o != nil && o.Right.IsSet() {
		return true
	}

	return false
}

// SetRight gets a reference to the given NullableInt32 and assigns it to the Right field.
func (o *MicrosoftGraphPrintMargin) SetRight(v int32) {
	o.Right.Set(&v)
}
// SetRightNil sets the value for Right to be an explicit nil
func (o *MicrosoftGraphPrintMargin) SetRightNil() {
	o.Right.Set(nil)
}

// UnsetRight ensures that no value is present for Right, not even an explicit nil
func (o *MicrosoftGraphPrintMargin) UnsetRight() {
	o.Right.Unset()
}

// GetTop returns the Top field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintMargin) GetTop() int32 {
	if o == nil || o.Top.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Top.Get()
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintMargin) GetTopOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Top.Get(), o.Top.IsSet()
}

// HasTop returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintMargin) HasTop() bool {
	if o != nil && o.Top.IsSet() {
		return true
	}

	return false
}

// SetTop gets a reference to the given NullableInt32 and assigns it to the Top field.
func (o *MicrosoftGraphPrintMargin) SetTop(v int32) {
	o.Top.Set(&v)
}
// SetTopNil sets the value for Top to be an explicit nil
func (o *MicrosoftGraphPrintMargin) SetTopNil() {
	o.Top.Set(nil)
}

// UnsetTop ensures that no value is present for Top, not even an explicit nil
func (o *MicrosoftGraphPrintMargin) UnsetTop() {
	o.Top.Unset()
}

func (o MicrosoftGraphPrintMargin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bottom.IsSet() {
		toSerialize["bottom"] = o.Bottom.Get()
	}
	if o.Left.IsSet() {
		toSerialize["left"] = o.Left.Get()
	}
	if o.Right.IsSet() {
		toSerialize["right"] = o.Right.Get()
	}
	if o.Top.IsSet() {
		toSerialize["top"] = o.Top.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintMargin struct {
	value *MicrosoftGraphPrintMargin
	isSet bool
}

func (v NullableMicrosoftGraphPrintMargin) Get() *MicrosoftGraphPrintMargin {
	return v.value
}

func (v *NullableMicrosoftGraphPrintMargin) Set(val *MicrosoftGraphPrintMargin) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintMargin) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintMargin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintMargin(val *MicrosoftGraphPrintMargin) *NullableMicrosoftGraphPrintMargin {
	return &NullableMicrosoftGraphPrintMargin{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintMargin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintMargin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

