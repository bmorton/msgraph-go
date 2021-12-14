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

// MicrosoftGraphSpecialFolder struct for MicrosoftGraphSpecialFolder
type MicrosoftGraphSpecialFolder struct {
	// The unique identifier for this item in the /drive/special collection
	Name NullableString `json:"name,omitempty"`
}

// NewMicrosoftGraphSpecialFolder instantiates a new MicrosoftGraphSpecialFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSpecialFolder() *MicrosoftGraphSpecialFolder {
	this := MicrosoftGraphSpecialFolder{}
	return &this
}

// NewMicrosoftGraphSpecialFolderWithDefaults instantiates a new MicrosoftGraphSpecialFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSpecialFolderWithDefaults() *MicrosoftGraphSpecialFolder {
	this := MicrosoftGraphSpecialFolder{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSpecialFolder) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSpecialFolder) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphSpecialFolder) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphSpecialFolder) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphSpecialFolder) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphSpecialFolder) UnsetName() {
	o.Name.Unset()
}

func (o MicrosoftGraphSpecialFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSpecialFolder struct {
	value *MicrosoftGraphSpecialFolder
	isSet bool
}

func (v NullableMicrosoftGraphSpecialFolder) Get() *MicrosoftGraphSpecialFolder {
	return v.value
}

func (v *NullableMicrosoftGraphSpecialFolder) Set(val *MicrosoftGraphSpecialFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSpecialFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSpecialFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSpecialFolder(val *MicrosoftGraphSpecialFolder) *NullableMicrosoftGraphSpecialFolder {
	return &NullableMicrosoftGraphSpecialFolder{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSpecialFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSpecialFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

