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

// MicrosoftGraphDataSubject struct for MicrosoftGraphDataSubject
type MicrosoftGraphDataSubject struct {
	// Email of the data subject.
	Email NullableString `json:"email,omitempty"`
	// First name of the data subject.
	FirstName NullableString `json:"firstName,omitempty"`
	// Last Name of the data subject.
	LastName NullableString `json:"lastName,omitempty"`
	// The country/region of residency. The residency information is uesed only for internal reporting but not for the content search.
	Residency NullableString `json:"residency,omitempty"`
}

// NewMicrosoftGraphDataSubject instantiates a new MicrosoftGraphDataSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDataSubject() *MicrosoftGraphDataSubject {
	this := MicrosoftGraphDataSubject{}
	return &this
}

// NewMicrosoftGraphDataSubjectWithDefaults instantiates a new MicrosoftGraphDataSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDataSubjectWithDefaults() *MicrosoftGraphDataSubject {
	this := MicrosoftGraphDataSubject{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDataSubject) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDataSubject) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphDataSubject) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MicrosoftGraphDataSubject) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MicrosoftGraphDataSubject) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MicrosoftGraphDataSubject) UnsetEmail() {
	o.Email.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDataSubject) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDataSubject) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *MicrosoftGraphDataSubject) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *MicrosoftGraphDataSubject) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *MicrosoftGraphDataSubject) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *MicrosoftGraphDataSubject) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDataSubject) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDataSubject) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *MicrosoftGraphDataSubject) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *MicrosoftGraphDataSubject) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *MicrosoftGraphDataSubject) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *MicrosoftGraphDataSubject) UnsetLastName() {
	o.LastName.Unset()
}

// GetResidency returns the Residency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDataSubject) GetResidency() string {
	if o == nil || o.Residency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Residency.Get()
}

// GetResidencyOk returns a tuple with the Residency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDataSubject) GetResidencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Residency.Get(), o.Residency.IsSet()
}

// HasResidency returns a boolean if a field has been set.
func (o *MicrosoftGraphDataSubject) HasResidency() bool {
	if o != nil && o.Residency.IsSet() {
		return true
	}

	return false
}

// SetResidency gets a reference to the given NullableString and assigns it to the Residency field.
func (o *MicrosoftGraphDataSubject) SetResidency(v string) {
	o.Residency.Set(&v)
}
// SetResidencyNil sets the value for Residency to be an explicit nil
func (o *MicrosoftGraphDataSubject) SetResidencyNil() {
	o.Residency.Set(nil)
}

// UnsetResidency ensures that no value is present for Residency, not even an explicit nil
func (o *MicrosoftGraphDataSubject) UnsetResidency() {
	o.Residency.Unset()
}

func (o MicrosoftGraphDataSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Residency.IsSet() {
		toSerialize["residency"] = o.Residency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDataSubject struct {
	value *MicrosoftGraphDataSubject
	isSet bool
}

func (v NullableMicrosoftGraphDataSubject) Get() *MicrosoftGraphDataSubject {
	return v.value
}

func (v *NullableMicrosoftGraphDataSubject) Set(val *MicrosoftGraphDataSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDataSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDataSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDataSubject(val *MicrosoftGraphDataSubject) *NullableMicrosoftGraphDataSubject {
	return &NullableMicrosoftGraphDataSubject{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDataSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDataSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


