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

// MicrosoftGraphRelatedContact struct for MicrosoftGraphRelatedContact
type MicrosoftGraphRelatedContact struct {
	// Indicates whether the user has been consented to access student data.
	AccessConsent NullableBool `json:"accessConsent,omitempty"`
	// Name of the contact. Required.
	DisplayName *string `json:"displayName,omitempty"`
	// Primary email address of the contact.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Mobile phone number of the contact.
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
	// Relationship to the user. Possible values are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
	Relationship AnyOfmicrosoftGraphContactRelationship `json:"relationship,omitempty"`
}

// NewMicrosoftGraphRelatedContact instantiates a new MicrosoftGraphRelatedContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRelatedContact() *MicrosoftGraphRelatedContact {
	this := MicrosoftGraphRelatedContact{}
	return &this
}

// NewMicrosoftGraphRelatedContactWithDefaults instantiates a new MicrosoftGraphRelatedContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRelatedContactWithDefaults() *MicrosoftGraphRelatedContact {
	this := MicrosoftGraphRelatedContact{}
	return &this
}

// GetAccessConsent returns the AccessConsent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRelatedContact) GetAccessConsent() bool {
	if o == nil || o.AccessConsent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AccessConsent.Get()
}

// GetAccessConsentOk returns a tuple with the AccessConsent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRelatedContact) GetAccessConsentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessConsent.Get(), o.AccessConsent.IsSet()
}

// HasAccessConsent returns a boolean if a field has been set.
func (o *MicrosoftGraphRelatedContact) HasAccessConsent() bool {
	if o != nil && o.AccessConsent.IsSet() {
		return true
	}

	return false
}

// SetAccessConsent gets a reference to the given NullableBool and assigns it to the AccessConsent field.
func (o *MicrosoftGraphRelatedContact) SetAccessConsent(v bool) {
	o.AccessConsent.Set(&v)
}
// SetAccessConsentNil sets the value for AccessConsent to be an explicit nil
func (o *MicrosoftGraphRelatedContact) SetAccessConsentNil() {
	o.AccessConsent.Set(nil)
}

// UnsetAccessConsent ensures that no value is present for AccessConsent, not even an explicit nil
func (o *MicrosoftGraphRelatedContact) UnsetAccessConsent() {
	o.AccessConsent.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphRelatedContact) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRelatedContact) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphRelatedContact) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphRelatedContact) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *MicrosoftGraphRelatedContact) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRelatedContact) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphRelatedContact) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *MicrosoftGraphRelatedContact) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRelatedContact) GetMobilePhone() string {
	if o == nil || o.MobilePhone.Get() == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRelatedContact) GetMobilePhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *MicrosoftGraphRelatedContact) HasMobilePhone() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *MicrosoftGraphRelatedContact) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}
// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *MicrosoftGraphRelatedContact) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *MicrosoftGraphRelatedContact) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

// GetRelationship returns the Relationship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRelatedContact) GetRelationship() AnyOfmicrosoftGraphContactRelationship {
	if o == nil  {
		var ret AnyOfmicrosoftGraphContactRelationship
		return ret
	}
	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRelatedContact) GetRelationshipOk() (*AnyOfmicrosoftGraphContactRelationship, bool) {
	if o == nil || o.Relationship == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *MicrosoftGraphRelatedContact) HasRelationship() bool {
	if o != nil && o.Relationship != nil {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given AnyOfmicrosoftGraphContactRelationship and assigns it to the Relationship field.
func (o *MicrosoftGraphRelatedContact) SetRelationship(v AnyOfmicrosoftGraphContactRelationship) {
	o.Relationship = v
}

func (o MicrosoftGraphRelatedContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessConsent.IsSet() {
		toSerialize["accessConsent"] = o.AccessConsent.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.MobilePhone.IsSet() {
		toSerialize["mobilePhone"] = o.MobilePhone.Get()
	}
	if o.Relationship != nil {
		toSerialize["relationship"] = o.Relationship
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRelatedContact struct {
	value *MicrosoftGraphRelatedContact
	isSet bool
}

func (v NullableMicrosoftGraphRelatedContact) Get() *MicrosoftGraphRelatedContact {
	return v.value
}

func (v *NullableMicrosoftGraphRelatedContact) Set(val *MicrosoftGraphRelatedContact) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRelatedContact) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRelatedContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRelatedContact(val *MicrosoftGraphRelatedContact) *NullableMicrosoftGraphRelatedContact {
	return &NullableMicrosoftGraphRelatedContact{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRelatedContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRelatedContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


