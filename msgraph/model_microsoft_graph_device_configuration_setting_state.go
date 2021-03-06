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

// MicrosoftGraphDeviceConfigurationSettingState Device Configuration Setting State for a given device.
type MicrosoftGraphDeviceConfigurationSettingState struct {
	// Current value of setting on device
	CurrentValue NullableString `json:"currentValue,omitempty"`
	// Error code for the setting
	ErrorCode *int64 `json:"errorCode,omitempty"`
	// Error description
	ErrorDescription NullableString `json:"errorDescription,omitempty"`
	// Name of setting instance that is being reported.
	InstanceDisplayName NullableString `json:"instanceDisplayName,omitempty"`
	// The setting that is being reported
	Setting NullableString `json:"setting,omitempty"`
	// Localized/user friendly setting name that is being reported
	SettingName NullableString `json:"settingName,omitempty"`
	// Contributing policies
	Sources *[]*AnyOfmicrosoftGraphSettingSource `json:"sources,omitempty"`
	// The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
	State AnyOfmicrosoftGraphComplianceStatus `json:"state,omitempty"`
	// UserEmail
	UserEmail NullableString `json:"userEmail,omitempty"`
	// UserId
	UserId NullableString `json:"userId,omitempty"`
	// UserName
	UserName NullableString `json:"userName,omitempty"`
	// UserPrincipalName.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewMicrosoftGraphDeviceConfigurationSettingState instantiates a new MicrosoftGraphDeviceConfigurationSettingState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceConfigurationSettingState() *MicrosoftGraphDeviceConfigurationSettingState {
	this := MicrosoftGraphDeviceConfigurationSettingState{}
	return &this
}

// NewMicrosoftGraphDeviceConfigurationSettingStateWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationSettingState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceConfigurationSettingStateWithDefaults() *MicrosoftGraphDeviceConfigurationSettingState {
	this := MicrosoftGraphDeviceConfigurationSettingState{}
	return &this
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetCurrentValue() string {
	if o == nil || o.CurrentValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.CurrentValue.Get()
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetCurrentValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentValue.Get(), o.CurrentValue.IsSet()
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasCurrentValue() bool {
	if o != nil && o.CurrentValue.IsSet() {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given NullableString and assigns it to the CurrentValue field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetCurrentValue(v string) {
	o.CurrentValue.Set(&v)
}
// SetCurrentValueNil sets the value for CurrentValue to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetCurrentValueNil() {
	o.CurrentValue.Set(nil)
}

// UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetCurrentValue() {
	o.CurrentValue.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorCodeOk() (*int64, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorDescription() string {
	if o == nil || o.ErrorDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription.Get()
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorDescription.Get(), o.ErrorDescription.IsSet()
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription.IsSet() {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given NullableString and assigns it to the ErrorDescription field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorDescription(v string) {
	o.ErrorDescription.Set(&v)
}
// SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorDescriptionNil() {
	o.ErrorDescription.Set(nil)
}

// UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetErrorDescription() {
	o.ErrorDescription.Unset()
}

// GetInstanceDisplayName returns the InstanceDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetInstanceDisplayName() string {
	if o == nil || o.InstanceDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstanceDisplayName.Get()
}

// GetInstanceDisplayNameOk returns a tuple with the InstanceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetInstanceDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstanceDisplayName.Get(), o.InstanceDisplayName.IsSet()
}

// HasInstanceDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasInstanceDisplayName() bool {
	if o != nil && o.InstanceDisplayName.IsSet() {
		return true
	}

	return false
}

// SetInstanceDisplayName gets a reference to the given NullableString and assigns it to the InstanceDisplayName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetInstanceDisplayName(v string) {
	o.InstanceDisplayName.Set(&v)
}
// SetInstanceDisplayNameNil sets the value for InstanceDisplayName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetInstanceDisplayNameNil() {
	o.InstanceDisplayName.Set(nil)
}

// UnsetInstanceDisplayName ensures that no value is present for InstanceDisplayName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetInstanceDisplayName() {
	o.InstanceDisplayName.Unset()
}

// GetSetting returns the Setting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSetting() string {
	if o == nil || o.Setting.Get() == nil {
		var ret string
		return ret
	}
	return *o.Setting.Get()
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Setting.Get(), o.Setting.IsSet()
}

// HasSetting returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSetting() bool {
	if o != nil && o.Setting.IsSet() {
		return true
	}

	return false
}

// SetSetting gets a reference to the given NullableString and assigns it to the Setting field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSetting(v string) {
	o.Setting.Set(&v)
}
// SetSettingNil sets the value for Setting to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingNil() {
	o.Setting.Set(nil)
}

// UnsetSetting ensures that no value is present for Setting, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetSetting() {
	o.Setting.Unset()
}

// GetSettingName returns the SettingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingName() string {
	if o == nil || o.SettingName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SettingName.Get()
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SettingName.Get(), o.SettingName.IsSet()
}

// HasSettingName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSettingName() bool {
	if o != nil && o.SettingName.IsSet() {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given NullableString and assigns it to the SettingName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingName(v string) {
	o.SettingName.Set(&v)
}
// SetSettingNameNil sets the value for SettingName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingNameNil() {
	o.SettingName.Set(nil)
}

// UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetSettingName() {
	o.SettingName.Unset()
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSources() []*AnyOfmicrosoftGraphSettingSource {
	if o == nil || o.Sources == nil {
		var ret []*AnyOfmicrosoftGraphSettingSource
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSourcesOk() (*[]*AnyOfmicrosoftGraphSettingSource, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []*AnyOfmicrosoftGraphSettingSource and assigns it to the Sources field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSources(v []*AnyOfmicrosoftGraphSettingSource) {
	o.Sources = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus) {
	o.State = v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserEmail() string {
	if o == nil || o.UserEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserEmail.Get()
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserEmail.Get(), o.UserEmail.IsSet()
}

// HasUserEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserEmail() bool {
	if o != nil && o.UserEmail.IsSet() {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given NullableString and assigns it to the UserEmail field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserEmail(v string) {
	o.UserEmail.Set(&v)
}
// SetUserEmailNil sets the value for UserEmail to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserEmailNil() {
	o.UserEmail.Set(nil)
}

// UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserEmail() {
	o.UserEmail.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserName() {
	o.UserName.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o MicrosoftGraphDeviceConfigurationSettingState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentValue.IsSet() {
		toSerialize["currentValue"] = o.CurrentValue.Get()
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription.IsSet() {
		toSerialize["errorDescription"] = o.ErrorDescription.Get()
	}
	if o.InstanceDisplayName.IsSet() {
		toSerialize["instanceDisplayName"] = o.InstanceDisplayName.Get()
	}
	if o.Setting.IsSet() {
		toSerialize["setting"] = o.Setting.Get()
	}
	if o.SettingName.IsSet() {
		toSerialize["settingName"] = o.SettingName.Get()
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UserEmail.IsSet() {
		toSerialize["userEmail"] = o.UserEmail.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceConfigurationSettingState struct {
	value *MicrosoftGraphDeviceConfigurationSettingState
	isSet bool
}

func (v NullableMicrosoftGraphDeviceConfigurationSettingState) Get() *MicrosoftGraphDeviceConfigurationSettingState {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceConfigurationSettingState) Set(val *MicrosoftGraphDeviceConfigurationSettingState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceConfigurationSettingState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceConfigurationSettingState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceConfigurationSettingState(val *MicrosoftGraphDeviceConfigurationSettingState) *NullableMicrosoftGraphDeviceConfigurationSettingState {
	return &NullableMicrosoftGraphDeviceConfigurationSettingState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceConfigurationSettingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceConfigurationSettingState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


