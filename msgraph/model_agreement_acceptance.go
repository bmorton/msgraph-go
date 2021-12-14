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

// AgreementAcceptance struct for AgreementAcceptance
type AgreementAcceptance struct {
	// The identifier of the agreement file accepted by the user.
	AgreementFileId NullableString `json:"agreementFileId,omitempty"`
	// The identifier of the agreement.
	AgreementId NullableString `json:"agreementId,omitempty"`
	// The display name of the device used for accepting the agreement.
	DeviceDisplayName NullableString `json:"deviceDisplayName,omitempty"`
	// The unique identifier of the device used for accepting the agreement.
	DeviceId NullableString `json:"deviceId,omitempty"`
	// The operating system used to accept the agreement.
	DeviceOSType NullableString `json:"deviceOSType,omitempty"`
	// The operating system version of the device used to accept the agreement.
	DeviceOSVersion NullableString `json:"deviceOSVersion,omitempty"`
	// The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ExpirationDateTime NullableTime `json:"expirationDateTime,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	RecordedDateTime NullableTime `json:"recordedDateTime,omitempty"`
	// The state of the agreement acceptance. Possible values are: accepted, declined.
	State AnyOfmicrosoftGraphAgreementAcceptanceState `json:"state,omitempty"`
	// Display name of the user when the acceptance was recorded.
	UserDisplayName NullableString `json:"userDisplayName,omitempty"`
	// Email of the user when the acceptance was recorded.
	UserEmail NullableString `json:"userEmail,omitempty"`
	// The identifier of the user who accepted the agreement.
	UserId NullableString `json:"userId,omitempty"`
	// UPN of the user when the acceptance was recorded.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewAgreementAcceptance instantiates a new AgreementAcceptance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementAcceptance() *AgreementAcceptance {
	this := AgreementAcceptance{}
	return &this
}

// NewAgreementAcceptanceWithDefaults instantiates a new AgreementAcceptance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementAcceptanceWithDefaults() *AgreementAcceptance {
	this := AgreementAcceptance{}
	return &this
}

// GetAgreementFileId returns the AgreementFileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetAgreementFileId() string {
	if o == nil || o.AgreementFileId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgreementFileId.Get()
}

// GetAgreementFileIdOk returns a tuple with the AgreementFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetAgreementFileIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgreementFileId.Get(), o.AgreementFileId.IsSet()
}

// HasAgreementFileId returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasAgreementFileId() bool {
	if o != nil && o.AgreementFileId.IsSet() {
		return true
	}

	return false
}

// SetAgreementFileId gets a reference to the given NullableString and assigns it to the AgreementFileId field.
func (o *AgreementAcceptance) SetAgreementFileId(v string) {
	o.AgreementFileId.Set(&v)
}
// SetAgreementFileIdNil sets the value for AgreementFileId to be an explicit nil
func (o *AgreementAcceptance) SetAgreementFileIdNil() {
	o.AgreementFileId.Set(nil)
}

// UnsetAgreementFileId ensures that no value is present for AgreementFileId, not even an explicit nil
func (o *AgreementAcceptance) UnsetAgreementFileId() {
	o.AgreementFileId.Unset()
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetAgreementId() string {
	if o == nil || o.AgreementId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgreementId.Get()
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetAgreementIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgreementId.Get(), o.AgreementId.IsSet()
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasAgreementId() bool {
	if o != nil && o.AgreementId.IsSet() {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given NullableString and assigns it to the AgreementId field.
func (o *AgreementAcceptance) SetAgreementId(v string) {
	o.AgreementId.Set(&v)
}
// SetAgreementIdNil sets the value for AgreementId to be an explicit nil
func (o *AgreementAcceptance) SetAgreementIdNil() {
	o.AgreementId.Set(nil)
}

// UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
func (o *AgreementAcceptance) UnsetAgreementId() {
	o.AgreementId.Unset()
}

// GetDeviceDisplayName returns the DeviceDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetDeviceDisplayName() string {
	if o == nil || o.DeviceDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceDisplayName.Get()
}

// GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetDeviceDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceDisplayName.Get(), o.DeviceDisplayName.IsSet()
}

// HasDeviceDisplayName returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasDeviceDisplayName() bool {
	if o != nil && o.DeviceDisplayName.IsSet() {
		return true
	}

	return false
}

// SetDeviceDisplayName gets a reference to the given NullableString and assigns it to the DeviceDisplayName field.
func (o *AgreementAcceptance) SetDeviceDisplayName(v string) {
	o.DeviceDisplayName.Set(&v)
}
// SetDeviceDisplayNameNil sets the value for DeviceDisplayName to be an explicit nil
func (o *AgreementAcceptance) SetDeviceDisplayNameNil() {
	o.DeviceDisplayName.Set(nil)
}

// UnsetDeviceDisplayName ensures that no value is present for DeviceDisplayName, not even an explicit nil
func (o *AgreementAcceptance) UnsetDeviceDisplayName() {
	o.DeviceDisplayName.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetDeviceId() string {
	if o == nil || o.DeviceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetDeviceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *AgreementAcceptance) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *AgreementAcceptance) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *AgreementAcceptance) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetDeviceOSType returns the DeviceOSType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetDeviceOSType() string {
	if o == nil || o.DeviceOSType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceOSType.Get()
}

// GetDeviceOSTypeOk returns a tuple with the DeviceOSType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetDeviceOSTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceOSType.Get(), o.DeviceOSType.IsSet()
}

// HasDeviceOSType returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasDeviceOSType() bool {
	if o != nil && o.DeviceOSType.IsSet() {
		return true
	}

	return false
}

// SetDeviceOSType gets a reference to the given NullableString and assigns it to the DeviceOSType field.
func (o *AgreementAcceptance) SetDeviceOSType(v string) {
	o.DeviceOSType.Set(&v)
}
// SetDeviceOSTypeNil sets the value for DeviceOSType to be an explicit nil
func (o *AgreementAcceptance) SetDeviceOSTypeNil() {
	o.DeviceOSType.Set(nil)
}

// UnsetDeviceOSType ensures that no value is present for DeviceOSType, not even an explicit nil
func (o *AgreementAcceptance) UnsetDeviceOSType() {
	o.DeviceOSType.Unset()
}

// GetDeviceOSVersion returns the DeviceOSVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetDeviceOSVersion() string {
	if o == nil || o.DeviceOSVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceOSVersion.Get()
}

// GetDeviceOSVersionOk returns a tuple with the DeviceOSVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetDeviceOSVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceOSVersion.Get(), o.DeviceOSVersion.IsSet()
}

// HasDeviceOSVersion returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasDeviceOSVersion() bool {
	if o != nil && o.DeviceOSVersion.IsSet() {
		return true
	}

	return false
}

// SetDeviceOSVersion gets a reference to the given NullableString and assigns it to the DeviceOSVersion field.
func (o *AgreementAcceptance) SetDeviceOSVersion(v string) {
	o.DeviceOSVersion.Set(&v)
}
// SetDeviceOSVersionNil sets the value for DeviceOSVersion to be an explicit nil
func (o *AgreementAcceptance) SetDeviceOSVersionNil() {
	o.DeviceOSVersion.Set(nil)
}

// UnsetDeviceOSVersion ensures that no value is present for DeviceOSVersion, not even an explicit nil
func (o *AgreementAcceptance) UnsetDeviceOSVersion() {
	o.DeviceOSVersion.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime.Get()
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDateTime.Get(), o.ExpirationDateTime.IsSet()
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given NullableTime and assigns it to the ExpirationDateTime field.
func (o *AgreementAcceptance) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime.Set(&v)
}
// SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil
func (o *AgreementAcceptance) SetExpirationDateTimeNil() {
	o.ExpirationDateTime.Set(nil)
}

// UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
func (o *AgreementAcceptance) UnsetExpirationDateTime() {
	o.ExpirationDateTime.Unset()
}

// GetRecordedDateTime returns the RecordedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetRecordedDateTime() time.Time {
	if o == nil || o.RecordedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RecordedDateTime.Get()
}

// GetRecordedDateTimeOk returns a tuple with the RecordedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetRecordedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordedDateTime.Get(), o.RecordedDateTime.IsSet()
}

// HasRecordedDateTime returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasRecordedDateTime() bool {
	if o != nil && o.RecordedDateTime.IsSet() {
		return true
	}

	return false
}

// SetRecordedDateTime gets a reference to the given NullableTime and assigns it to the RecordedDateTime field.
func (o *AgreementAcceptance) SetRecordedDateTime(v time.Time) {
	o.RecordedDateTime.Set(&v)
}
// SetRecordedDateTimeNil sets the value for RecordedDateTime to be an explicit nil
func (o *AgreementAcceptance) SetRecordedDateTimeNil() {
	o.RecordedDateTime.Set(nil)
}

// UnsetRecordedDateTime ensures that no value is present for RecordedDateTime, not even an explicit nil
func (o *AgreementAcceptance) UnsetRecordedDateTime() {
	o.RecordedDateTime.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetState() AnyOfmicrosoftGraphAgreementAcceptanceState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAgreementAcceptanceState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetStateOk() (*AnyOfmicrosoftGraphAgreementAcceptanceState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphAgreementAcceptanceState and assigns it to the State field.
func (o *AgreementAcceptance) SetState(v AnyOfmicrosoftGraphAgreementAcceptanceState) {
	o.State = v
}

// GetUserDisplayName returns the UserDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetUserDisplayName() string {
	if o == nil || o.UserDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserDisplayName.Get()
}

// GetUserDisplayNameOk returns a tuple with the UserDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetUserDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserDisplayName.Get(), o.UserDisplayName.IsSet()
}

// HasUserDisplayName returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasUserDisplayName() bool {
	if o != nil && o.UserDisplayName.IsSet() {
		return true
	}

	return false
}

// SetUserDisplayName gets a reference to the given NullableString and assigns it to the UserDisplayName field.
func (o *AgreementAcceptance) SetUserDisplayName(v string) {
	o.UserDisplayName.Set(&v)
}
// SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil
func (o *AgreementAcceptance) SetUserDisplayNameNil() {
	o.UserDisplayName.Set(nil)
}

// UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
func (o *AgreementAcceptance) UnsetUserDisplayName() {
	o.UserDisplayName.Unset()
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetUserEmail() string {
	if o == nil || o.UserEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserEmail.Get()
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetUserEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserEmail.Get(), o.UserEmail.IsSet()
}

// HasUserEmail returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasUserEmail() bool {
	if o != nil && o.UserEmail.IsSet() {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given NullableString and assigns it to the UserEmail field.
func (o *AgreementAcceptance) SetUserEmail(v string) {
	o.UserEmail.Set(&v)
}
// SetUserEmailNil sets the value for UserEmail to be an explicit nil
func (o *AgreementAcceptance) SetUserEmailNil() {
	o.UserEmail.Set(nil)
}

// UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
func (o *AgreementAcceptance) UnsetUserEmail() {
	o.UserEmail.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *AgreementAcceptance) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *AgreementAcceptance) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *AgreementAcceptance) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreementAcceptance) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreementAcceptance) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *AgreementAcceptance) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *AgreementAcceptance) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *AgreementAcceptance) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *AgreementAcceptance) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o AgreementAcceptance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgreementFileId.IsSet() {
		toSerialize["agreementFileId"] = o.AgreementFileId.Get()
	}
	if o.AgreementId.IsSet() {
		toSerialize["agreementId"] = o.AgreementId.Get()
	}
	if o.DeviceDisplayName.IsSet() {
		toSerialize["deviceDisplayName"] = o.DeviceDisplayName.Get()
	}
	if o.DeviceId.IsSet() {
		toSerialize["deviceId"] = o.DeviceId.Get()
	}
	if o.DeviceOSType.IsSet() {
		toSerialize["deviceOSType"] = o.DeviceOSType.Get()
	}
	if o.DeviceOSVersion.IsSet() {
		toSerialize["deviceOSVersion"] = o.DeviceOSVersion.Get()
	}
	if o.ExpirationDateTime.IsSet() {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime.Get()
	}
	if o.RecordedDateTime.IsSet() {
		toSerialize["recordedDateTime"] = o.RecordedDateTime.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UserDisplayName.IsSet() {
		toSerialize["userDisplayName"] = o.UserDisplayName.Get()
	}
	if o.UserEmail.IsSet() {
		toSerialize["userEmail"] = o.UserEmail.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAgreementAcceptance struct {
	value *AgreementAcceptance
	isSet bool
}

func (v NullableAgreementAcceptance) Get() *AgreementAcceptance {
	return v.value
}

func (v *NullableAgreementAcceptance) Set(val *AgreementAcceptance) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementAcceptance) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementAcceptance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementAcceptance(val *AgreementAcceptance) *NullableAgreementAcceptance {
	return &NullableAgreementAcceptance{value: val, isSet: true}
}

func (v NullableAgreementAcceptance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementAcceptance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


