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

// MicrosoftGraphEmailAddress struct for MicrosoftGraphEmailAddress
type MicrosoftGraphEmailAddress struct {
	// The email address of the person or entity.
	Address NullableString `json:"address,omitempty"`
	// The display name of the person or entity.
	Name NullableString `json:"name,omitempty"`
}

// NewMicrosoftGraphEmailAddress instantiates a new MicrosoftGraphEmailAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphEmailAddress() *MicrosoftGraphEmailAddress {
	this := MicrosoftGraphEmailAddress{}
	return &this
}

// NewMicrosoftGraphEmailAddressWithDefaults instantiates a new MicrosoftGraphEmailAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphEmailAddressWithDefaults() *MicrosoftGraphEmailAddress {
	this := MicrosoftGraphEmailAddress{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEmailAddress) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEmailAddress) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphEmailAddress) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *MicrosoftGraphEmailAddress) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *MicrosoftGraphEmailAddress) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *MicrosoftGraphEmailAddress) UnsetAddress() {
	o.Address.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEmailAddress) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEmailAddress) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphEmailAddress) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphEmailAddress) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphEmailAddress) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphEmailAddress) UnsetName() {
	o.Name.Unset()
}

func (o MicrosoftGraphEmailAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphEmailAddress struct {
	value *MicrosoftGraphEmailAddress
	isSet bool
}

func (v NullableMicrosoftGraphEmailAddress) Get() *MicrosoftGraphEmailAddress {
	return v.value
}

func (v *NullableMicrosoftGraphEmailAddress) Set(val *MicrosoftGraphEmailAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEmailAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEmailAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEmailAddress(val *MicrosoftGraphEmailAddress) *NullableMicrosoftGraphEmailAddress {
	return &NullableMicrosoftGraphEmailAddress{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEmailAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEmailAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


