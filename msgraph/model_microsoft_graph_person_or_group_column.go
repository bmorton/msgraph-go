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

// MicrosoftGraphPersonOrGroupColumn struct for MicrosoftGraphPersonOrGroupColumn
type MicrosoftGraphPersonOrGroupColumn struct {
	// Indicates whether multiple values can be selected from the source.
	AllowMultipleSelection NullableBool `json:"allowMultipleSelection,omitempty"`
	// Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
	ChooseFromType NullableString `json:"chooseFromType,omitempty"`
	// How to display the information about the person or group chosen. See below.
	DisplayAs NullableString `json:"displayAs,omitempty"`
}

// NewMicrosoftGraphPersonOrGroupColumn instantiates a new MicrosoftGraphPersonOrGroupColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPersonOrGroupColumn() *MicrosoftGraphPersonOrGroupColumn {
	this := MicrosoftGraphPersonOrGroupColumn{}
	return &this
}

// NewMicrosoftGraphPersonOrGroupColumnWithDefaults instantiates a new MicrosoftGraphPersonOrGroupColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPersonOrGroupColumnWithDefaults() *MicrosoftGraphPersonOrGroupColumn {
	this := MicrosoftGraphPersonOrGroupColumn{}
	return &this
}

// GetAllowMultipleSelection returns the AllowMultipleSelection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPersonOrGroupColumn) GetAllowMultipleSelection() bool {
	if o == nil || o.AllowMultipleSelection.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowMultipleSelection.Get()
}

// GetAllowMultipleSelectionOk returns a tuple with the AllowMultipleSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPersonOrGroupColumn) GetAllowMultipleSelectionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowMultipleSelection.Get(), o.AllowMultipleSelection.IsSet()
}

// HasAllowMultipleSelection returns a boolean if a field has been set.
func (o *MicrosoftGraphPersonOrGroupColumn) HasAllowMultipleSelection() bool {
	if o != nil && o.AllowMultipleSelection.IsSet() {
		return true
	}

	return false
}

// SetAllowMultipleSelection gets a reference to the given NullableBool and assigns it to the AllowMultipleSelection field.
func (o *MicrosoftGraphPersonOrGroupColumn) SetAllowMultipleSelection(v bool) {
	o.AllowMultipleSelection.Set(&v)
}
// SetAllowMultipleSelectionNil sets the value for AllowMultipleSelection to be an explicit nil
func (o *MicrosoftGraphPersonOrGroupColumn) SetAllowMultipleSelectionNil() {
	o.AllowMultipleSelection.Set(nil)
}

// UnsetAllowMultipleSelection ensures that no value is present for AllowMultipleSelection, not even an explicit nil
func (o *MicrosoftGraphPersonOrGroupColumn) UnsetAllowMultipleSelection() {
	o.AllowMultipleSelection.Unset()
}

// GetChooseFromType returns the ChooseFromType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPersonOrGroupColumn) GetChooseFromType() string {
	if o == nil || o.ChooseFromType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChooseFromType.Get()
}

// GetChooseFromTypeOk returns a tuple with the ChooseFromType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPersonOrGroupColumn) GetChooseFromTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChooseFromType.Get(), o.ChooseFromType.IsSet()
}

// HasChooseFromType returns a boolean if a field has been set.
func (o *MicrosoftGraphPersonOrGroupColumn) HasChooseFromType() bool {
	if o != nil && o.ChooseFromType.IsSet() {
		return true
	}

	return false
}

// SetChooseFromType gets a reference to the given NullableString and assigns it to the ChooseFromType field.
func (o *MicrosoftGraphPersonOrGroupColumn) SetChooseFromType(v string) {
	o.ChooseFromType.Set(&v)
}
// SetChooseFromTypeNil sets the value for ChooseFromType to be an explicit nil
func (o *MicrosoftGraphPersonOrGroupColumn) SetChooseFromTypeNil() {
	o.ChooseFromType.Set(nil)
}

// UnsetChooseFromType ensures that no value is present for ChooseFromType, not even an explicit nil
func (o *MicrosoftGraphPersonOrGroupColumn) UnsetChooseFromType() {
	o.ChooseFromType.Unset()
}

// GetDisplayAs returns the DisplayAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPersonOrGroupColumn) GetDisplayAs() string {
	if o == nil || o.DisplayAs.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayAs.Get()
}

// GetDisplayAsOk returns a tuple with the DisplayAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPersonOrGroupColumn) GetDisplayAsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayAs.Get(), o.DisplayAs.IsSet()
}

// HasDisplayAs returns a boolean if a field has been set.
func (o *MicrosoftGraphPersonOrGroupColumn) HasDisplayAs() bool {
	if o != nil && o.DisplayAs.IsSet() {
		return true
	}

	return false
}

// SetDisplayAs gets a reference to the given NullableString and assigns it to the DisplayAs field.
func (o *MicrosoftGraphPersonOrGroupColumn) SetDisplayAs(v string) {
	o.DisplayAs.Set(&v)
}
// SetDisplayAsNil sets the value for DisplayAs to be an explicit nil
func (o *MicrosoftGraphPersonOrGroupColumn) SetDisplayAsNil() {
	o.DisplayAs.Set(nil)
}

// UnsetDisplayAs ensures that no value is present for DisplayAs, not even an explicit nil
func (o *MicrosoftGraphPersonOrGroupColumn) UnsetDisplayAs() {
	o.DisplayAs.Unset()
}

func (o MicrosoftGraphPersonOrGroupColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowMultipleSelection.IsSet() {
		toSerialize["allowMultipleSelection"] = o.AllowMultipleSelection.Get()
	}
	if o.ChooseFromType.IsSet() {
		toSerialize["chooseFromType"] = o.ChooseFromType.Get()
	}
	if o.DisplayAs.IsSet() {
		toSerialize["displayAs"] = o.DisplayAs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPersonOrGroupColumn struct {
	value *MicrosoftGraphPersonOrGroupColumn
	isSet bool
}

func (v NullableMicrosoftGraphPersonOrGroupColumn) Get() *MicrosoftGraphPersonOrGroupColumn {
	return v.value
}

func (v *NullableMicrosoftGraphPersonOrGroupColumn) Set(val *MicrosoftGraphPersonOrGroupColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPersonOrGroupColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPersonOrGroupColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPersonOrGroupColumn(val *MicrosoftGraphPersonOrGroupColumn) *NullableMicrosoftGraphPersonOrGroupColumn {
	return &NullableMicrosoftGraphPersonOrGroupColumn{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPersonOrGroupColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPersonOrGroupColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

