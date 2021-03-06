/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphTermsAndConditionsAcceptanceStatus struct for MicrosoftGraphTermsAndConditionsAcceptanceStatus
type MicrosoftGraphTermsAndConditionsAcceptanceStatus struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// DateTime when the terms were last accepted by the user.
	AcceptedDateTime *time.Time `json:"acceptedDateTime,omitempty"`
	// Most recent version number of the T&C accepted by the user.
	AcceptedVersion *int32 `json:"acceptedVersion,omitempty"`
	// Display name of the user whose acceptance the entity represents.
	UserDisplayName NullableString `json:"userDisplayName,omitempty"`
	// The userPrincipalName of the User that accepted the term.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
	// Navigation link to the terms and conditions that are assigned.
	TermsAndConditions AnyOfmicrosoftGraphTermsAndConditions `json:"termsAndConditions,omitempty"`
}

// NewMicrosoftGraphTermsAndConditionsAcceptanceStatus instantiates a new MicrosoftGraphTermsAndConditionsAcceptanceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTermsAndConditionsAcceptanceStatus() *MicrosoftGraphTermsAndConditionsAcceptanceStatus {
	this := MicrosoftGraphTermsAndConditionsAcceptanceStatus{}
	return &this
}

// NewMicrosoftGraphTermsAndConditionsAcceptanceStatusWithDefaults instantiates a new MicrosoftGraphTermsAndConditionsAcceptanceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTermsAndConditionsAcceptanceStatusWithDefaults() *MicrosoftGraphTermsAndConditionsAcceptanceStatus {
	this := MicrosoftGraphTermsAndConditionsAcceptanceStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetId(v string) {
	o.Id = &v
}

// GetAcceptedDateTime returns the AcceptedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedDateTime() time.Time {
	if o == nil || o.AcceptedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AcceptedDateTime
}

// GetAcceptedDateTimeOk returns a tuple with the AcceptedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.AcceptedDateTime == nil {
		return nil, false
	}
	return o.AcceptedDateTime, true
}

// HasAcceptedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasAcceptedDateTime() bool {
	if o != nil && o.AcceptedDateTime != nil {
		return true
	}

	return false
}

// SetAcceptedDateTime gets a reference to the given time.Time and assigns it to the AcceptedDateTime field.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(v time.Time) {
	o.AcceptedDateTime = &v
}

// GetAcceptedVersion returns the AcceptedVersion field value if set, zero value otherwise.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedVersion() int32 {
	if o == nil || o.AcceptedVersion == nil {
		var ret int32
		return ret
	}
	return *o.AcceptedVersion
}

// GetAcceptedVersionOk returns a tuple with the AcceptedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedVersionOk() (*int32, bool) {
	if o == nil || o.AcceptedVersion == nil {
		return nil, false
	}
	return o.AcceptedVersion, true
}

// HasAcceptedVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasAcceptedVersion() bool {
	if o != nil && o.AcceptedVersion != nil {
		return true
	}

	return false
}

// SetAcceptedVersion gets a reference to the given int32 and assigns it to the AcceptedVersion field.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetAcceptedVersion(v int32) {
	o.AcceptedVersion = &v
}

// GetUserDisplayName returns the UserDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetUserDisplayName() string {
	if o == nil || o.UserDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayName.Get()
}

// GetUserDisplayNameOk returns a tuple with the UserDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetUserDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDisplayName.Get(), o.UserDisplayName.IsSet()
}

// HasUserDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasUserDisplayName() bool {
	if o != nil && o.UserDisplayName.IsSet() {
		return true
	}

	return false
}

// SetUserDisplayName gets a reference to the given NullableString and assigns it to the UserDisplayName field.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetUserDisplayName(v string) {
	o.UserDisplayName.Set(&v)
}
// SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetUserDisplayNameNil() {
	o.UserDisplayName.Set(nil)
}

// UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) UnsetUserDisplayName() {
	o.UserDisplayName.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

// GetTermsAndConditions returns the TermsAndConditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetTermsAndConditions() AnyOfmicrosoftGraphTermsAndConditions {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermsAndConditions
		return ret
	}
	return o.TermsAndConditions
}

// GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetTermsAndConditionsOk() (*AnyOfmicrosoftGraphTermsAndConditions, bool) {
	if o == nil || o.TermsAndConditions == nil {
		return nil, false
	}
	return &o.TermsAndConditions, true
}

// HasTermsAndConditions returns a boolean if a field has been set.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasTermsAndConditions() bool {
	if o != nil && o.TermsAndConditions != nil {
		return true
	}

	return false
}

// SetTermsAndConditions gets a reference to the given AnyOfmicrosoftGraphTermsAndConditions and assigns it to the TermsAndConditions field.
func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetTermsAndConditions(v AnyOfmicrosoftGraphTermsAndConditions) {
	o.TermsAndConditions = v
}

func (o MicrosoftGraphTermsAndConditionsAcceptanceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AcceptedDateTime != nil {
		toSerialize["acceptedDateTime"] = o.AcceptedDateTime
	}
	if o.AcceptedVersion != nil {
		toSerialize["acceptedVersion"] = o.AcceptedVersion
	}
	if o.UserDisplayName.IsSet() {
		toSerialize["userDisplayName"] = o.UserDisplayName.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	if o.TermsAndConditions != nil {
		toSerialize["termsAndConditions"] = o.TermsAndConditions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus struct {
	value *MicrosoftGraphTermsAndConditionsAcceptanceStatus
	isSet bool
}

func (v NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus) Get() *MicrosoftGraphTermsAndConditionsAcceptanceStatus {
	return v.value
}

func (v *NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus) Set(val *MicrosoftGraphTermsAndConditionsAcceptanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTermsAndConditionsAcceptanceStatus(val *MicrosoftGraphTermsAndConditionsAcceptanceStatus) *NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus {
	return &NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTermsAndConditionsAcceptanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


