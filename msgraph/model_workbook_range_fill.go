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

// WorkbookRangeFill struct for WorkbookRangeFill
type WorkbookRangeFill struct {
	// HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
	Color NullableString `json:"color,omitempty"`
}

// NewWorkbookRangeFill instantiates a new WorkbookRangeFill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookRangeFill() *WorkbookRangeFill {
	this := WorkbookRangeFill{}
	return &this
}

// NewWorkbookRangeFillWithDefaults instantiates a new WorkbookRangeFill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookRangeFillWithDefaults() *WorkbookRangeFill {
	this := WorkbookRangeFill{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookRangeFill) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookRangeFill) GetColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *WorkbookRangeFill) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *WorkbookRangeFill) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *WorkbookRangeFill) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *WorkbookRangeFill) UnsetColor() {
	o.Color.Unset()
}

func (o WorkbookRangeFill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookRangeFill struct {
	value *WorkbookRangeFill
	isSet bool
}

func (v NullableWorkbookRangeFill) Get() *WorkbookRangeFill {
	return v.value
}

func (v *NullableWorkbookRangeFill) Set(val *WorkbookRangeFill) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookRangeFill) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookRangeFill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookRangeFill(val *WorkbookRangeFill) *NullableWorkbookRangeFill {
	return &NullableWorkbookRangeFill{value: val, isSet: true}
}

func (v NullableWorkbookRangeFill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookRangeFill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


