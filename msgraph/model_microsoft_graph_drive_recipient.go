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

// MicrosoftGraphDriveRecipient struct for MicrosoftGraphDriveRecipient
type MicrosoftGraphDriveRecipient struct {
	// The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
	Alias NullableString `json:"alias,omitempty"`
	// The email address for the recipient, if the recipient has an associated email address.
	Email NullableString `json:"email,omitempty"`
	// The unique identifier for the recipient in the directory.
	ObjectId NullableString `json:"objectId,omitempty"`
}

// NewMicrosoftGraphDriveRecipient instantiates a new MicrosoftGraphDriveRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDriveRecipient() *MicrosoftGraphDriveRecipient {
	this := MicrosoftGraphDriveRecipient{}
	return &this
}

// NewMicrosoftGraphDriveRecipientWithDefaults instantiates a new MicrosoftGraphDriveRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDriveRecipientWithDefaults() *MicrosoftGraphDriveRecipient {
	this := MicrosoftGraphDriveRecipient{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDriveRecipient) GetAlias() string {
	if o == nil || o.Alias.Get() == nil {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDriveRecipient) GetAliasOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveRecipient) HasAlias() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *MicrosoftGraphDriveRecipient) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *MicrosoftGraphDriveRecipient) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *MicrosoftGraphDriveRecipient) UnsetAlias() {
	o.Alias.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDriveRecipient) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDriveRecipient) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveRecipient) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MicrosoftGraphDriveRecipient) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MicrosoftGraphDriveRecipient) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MicrosoftGraphDriveRecipient) UnsetEmail() {
	o.Email.Unset()
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDriveRecipient) GetObjectId() string {
	if o == nil || o.ObjectId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectId.Get()
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDriveRecipient) GetObjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectId.Get(), o.ObjectId.IsSet()
}

// HasObjectId returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveRecipient) HasObjectId() bool {
	if o != nil && o.ObjectId.IsSet() {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given NullableString and assigns it to the ObjectId field.
func (o *MicrosoftGraphDriveRecipient) SetObjectId(v string) {
	o.ObjectId.Set(&v)
}
// SetObjectIdNil sets the value for ObjectId to be an explicit nil
func (o *MicrosoftGraphDriveRecipient) SetObjectIdNil() {
	o.ObjectId.Set(nil)
}

// UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
func (o *MicrosoftGraphDriveRecipient) UnsetObjectId() {
	o.ObjectId.Unset()
}

func (o MicrosoftGraphDriveRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.ObjectId.IsSet() {
		toSerialize["objectId"] = o.ObjectId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDriveRecipient struct {
	value *MicrosoftGraphDriveRecipient
	isSet bool
}

func (v NullableMicrosoftGraphDriveRecipient) Get() *MicrosoftGraphDriveRecipient {
	return v.value
}

func (v *NullableMicrosoftGraphDriveRecipient) Set(val *MicrosoftGraphDriveRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDriveRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDriveRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDriveRecipient(val *MicrosoftGraphDriveRecipient) *NullableMicrosoftGraphDriveRecipient {
	return &NullableMicrosoftGraphDriveRecipient{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDriveRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDriveRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


