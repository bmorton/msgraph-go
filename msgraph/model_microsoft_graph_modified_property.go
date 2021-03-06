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

// MicrosoftGraphModifiedProperty struct for MicrosoftGraphModifiedProperty
type MicrosoftGraphModifiedProperty struct {
	// Indicates the property name of the target attribute that was changed.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Indicates the updated value for the propery.
	NewValue NullableString `json:"newValue,omitempty"`
	// Indicates the previous value (before the update) for the property.
	OldValue NullableString `json:"oldValue,omitempty"`
}

// NewMicrosoftGraphModifiedProperty instantiates a new MicrosoftGraphModifiedProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphModifiedProperty() *MicrosoftGraphModifiedProperty {
	this := MicrosoftGraphModifiedProperty{}
	return &this
}

// NewMicrosoftGraphModifiedPropertyWithDefaults instantiates a new MicrosoftGraphModifiedProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphModifiedPropertyWithDefaults() *MicrosoftGraphModifiedProperty {
	this := MicrosoftGraphModifiedProperty{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphModifiedProperty) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphModifiedProperty) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphModifiedProperty) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphModifiedProperty) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphModifiedProperty) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphModifiedProperty) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphModifiedProperty) GetNewValue() string {
	if o == nil || o.NewValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphModifiedProperty) GetNewValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *MicrosoftGraphModifiedProperty) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *MicrosoftGraphModifiedProperty) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *MicrosoftGraphModifiedProperty) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *MicrosoftGraphModifiedProperty) UnsetNewValue() {
	o.NewValue.Unset()
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphModifiedProperty) GetOldValue() string {
	if o == nil || o.OldValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphModifiedProperty) GetOldValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *MicrosoftGraphModifiedProperty) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *MicrosoftGraphModifiedProperty) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *MicrosoftGraphModifiedProperty) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *MicrosoftGraphModifiedProperty) UnsetOldValue() {
	o.OldValue.Unset()
}

func (o MicrosoftGraphModifiedProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["newValue"] = o.NewValue.Get()
	}
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphModifiedProperty struct {
	value *MicrosoftGraphModifiedProperty
	isSet bool
}

func (v NullableMicrosoftGraphModifiedProperty) Get() *MicrosoftGraphModifiedProperty {
	return v.value
}

func (v *NullableMicrosoftGraphModifiedProperty) Set(val *MicrosoftGraphModifiedProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphModifiedProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphModifiedProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphModifiedProperty(val *MicrosoftGraphModifiedProperty) *NullableMicrosoftGraphModifiedProperty {
	return &NullableMicrosoftGraphModifiedProperty{value: val, isSet: true}
}

func (v NullableMicrosoftGraphModifiedProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphModifiedProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


