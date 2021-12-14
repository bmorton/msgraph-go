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

// MicrosoftGraphSingleValueLegacyExtendedProperty struct for MicrosoftGraphSingleValueLegacyExtendedProperty
type MicrosoftGraphSingleValueLegacyExtendedProperty struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A property value.
	Value NullableString `json:"value,omitempty"`
}

// NewMicrosoftGraphSingleValueLegacyExtendedProperty instantiates a new MicrosoftGraphSingleValueLegacyExtendedProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSingleValueLegacyExtendedProperty() *MicrosoftGraphSingleValueLegacyExtendedProperty {
	this := MicrosoftGraphSingleValueLegacyExtendedProperty{}
	return &this
}

// NewMicrosoftGraphSingleValueLegacyExtendedPropertyWithDefaults instantiates a new MicrosoftGraphSingleValueLegacyExtendedProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSingleValueLegacyExtendedPropertyWithDefaults() *MicrosoftGraphSingleValueLegacyExtendedProperty {
	this := MicrosoftGraphSingleValueLegacyExtendedProperty{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) UnsetValue() {
	o.Value.Unset()
}

func (o MicrosoftGraphSingleValueLegacyExtendedProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSingleValueLegacyExtendedProperty struct {
	value *MicrosoftGraphSingleValueLegacyExtendedProperty
	isSet bool
}

func (v NullableMicrosoftGraphSingleValueLegacyExtendedProperty) Get() *MicrosoftGraphSingleValueLegacyExtendedProperty {
	return v.value
}

func (v *NullableMicrosoftGraphSingleValueLegacyExtendedProperty) Set(val *MicrosoftGraphSingleValueLegacyExtendedProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSingleValueLegacyExtendedProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSingleValueLegacyExtendedProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSingleValueLegacyExtendedProperty(val *MicrosoftGraphSingleValueLegacyExtendedProperty) *NullableMicrosoftGraphSingleValueLegacyExtendedProperty {
	return &NullableMicrosoftGraphSingleValueLegacyExtendedProperty{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSingleValueLegacyExtendedProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSingleValueLegacyExtendedProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


