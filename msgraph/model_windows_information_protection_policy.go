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

// WindowsInformationProtectionPolicy Policy for Windows information protection without MDM
type WindowsInformationProtectionPolicy struct {
	// Offline interval before app data is wiped (days)
	DaysWithoutContactBeforeUnenroll *int32 `json:"daysWithoutContactBeforeUnenroll,omitempty"`
	// Enrollment url for the MDM
	MdmEnrollmentUrl NullableString `json:"mdmEnrollmentUrl,omitempty"`
	// Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 <= X <= 999.
	MinutesOfInactivityBeforeDeviceLock *int32 `json:"minutesOfInactivityBeforeDeviceLock,omitempty"`
	// Integer value that specifies the number of past PINs that can be associated to a user account that can't be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0.
	NumberOfPastPinsRemembered *int32 `json:"numberOfPastPinsRemembered,omitempty"`
	// The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 <= X <= 16 for desktop and 0 <= X <= 999 for mobile devices.
	PasswordMaximumAttemptCount *int32 `json:"passwordMaximumAttemptCount,omitempty"`
	// Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user's PIN will never expire. This node was added in Windows 10, version 1511. Default is 0.
	PinExpirationDays *int32 `json:"pinExpirationDays,omitempty"`
	// Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
	PinLowercaseLetters AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements `json:"pinLowercaseLetters,omitempty"`
	// Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest.
	PinMinimumLength *int32 `json:"pinMinimumLength,omitempty"`
	// Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! ' # $ % & ' ( )  + , - . / : ; < = > ? @ [ / ] ^  ` {
	PinSpecialCharacters AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements `json:"pinSpecialCharacters,omitempty"`
	// Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. Possible values are: notAllow, requireAtLeastOne, allow.
	PinUppercaseLetters AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements `json:"pinUppercaseLetters,omitempty"`
	// New property in RS2, pending documentation
	RevokeOnMdmHandoffDisabled *bool `json:"revokeOnMdmHandoffDisabled,omitempty"`
	// Boolean value that sets Windows Hello for Business as a method for signing into Windows.
	WindowsHelloForBusinessBlocked *bool `json:"windowsHelloForBusinessBlocked,omitempty"`
}

// NewWindowsInformationProtectionPolicy instantiates a new WindowsInformationProtectionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsInformationProtectionPolicy() *WindowsInformationProtectionPolicy {
	this := WindowsInformationProtectionPolicy{}
	return &this
}

// NewWindowsInformationProtectionPolicyWithDefaults instantiates a new WindowsInformationProtectionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsInformationProtectionPolicyWithDefaults() *WindowsInformationProtectionPolicy {
	this := WindowsInformationProtectionPolicy{}
	return &this
}

// GetDaysWithoutContactBeforeUnenroll returns the DaysWithoutContactBeforeUnenroll field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenroll() int32 {
	if o == nil || o.DaysWithoutContactBeforeUnenroll == nil {
		var ret int32
		return ret
	}
	return *o.DaysWithoutContactBeforeUnenroll
}

// GetDaysWithoutContactBeforeUnenrollOk returns a tuple with the DaysWithoutContactBeforeUnenroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenrollOk() (*int32, bool) {
	if o == nil || o.DaysWithoutContactBeforeUnenroll == nil {
		return nil, false
	}
	return o.DaysWithoutContactBeforeUnenroll, true
}

// HasDaysWithoutContactBeforeUnenroll returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasDaysWithoutContactBeforeUnenroll() bool {
	if o != nil && o.DaysWithoutContactBeforeUnenroll != nil {
		return true
	}

	return false
}

// SetDaysWithoutContactBeforeUnenroll gets a reference to the given int32 and assigns it to the DaysWithoutContactBeforeUnenroll field.
func (o *WindowsInformationProtectionPolicy) SetDaysWithoutContactBeforeUnenroll(v int32) {
	o.DaysWithoutContactBeforeUnenroll = &v
}

// GetMdmEnrollmentUrl returns the MdmEnrollmentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrl() string {
	if o == nil || o.MdmEnrollmentUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.MdmEnrollmentUrl.Get()
}

// GetMdmEnrollmentUrlOk returns a tuple with the MdmEnrollmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MdmEnrollmentUrl.Get(), o.MdmEnrollmentUrl.IsSet()
}

// HasMdmEnrollmentUrl returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasMdmEnrollmentUrl() bool {
	if o != nil && o.MdmEnrollmentUrl.IsSet() {
		return true
	}

	return false
}

// SetMdmEnrollmentUrl gets a reference to the given NullableString and assigns it to the MdmEnrollmentUrl field.
func (o *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrl(v string) {
	o.MdmEnrollmentUrl.Set(&v)
}
// SetMdmEnrollmentUrlNil sets the value for MdmEnrollmentUrl to be an explicit nil
func (o *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrlNil() {
	o.MdmEnrollmentUrl.Set(nil)
}

// UnsetMdmEnrollmentUrl ensures that no value is present for MdmEnrollmentUrl, not even an explicit nil
func (o *WindowsInformationProtectionPolicy) UnsetMdmEnrollmentUrl() {
	o.MdmEnrollmentUrl.Unset()
}

// GetMinutesOfInactivityBeforeDeviceLock returns the MinutesOfInactivityBeforeDeviceLock field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLock() int32 {
	if o == nil || o.MinutesOfInactivityBeforeDeviceLock == nil {
		var ret int32
		return ret
	}
	return *o.MinutesOfInactivityBeforeDeviceLock
}

// GetMinutesOfInactivityBeforeDeviceLockOk returns a tuple with the MinutesOfInactivityBeforeDeviceLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLockOk() (*int32, bool) {
	if o == nil || o.MinutesOfInactivityBeforeDeviceLock == nil {
		return nil, false
	}
	return o.MinutesOfInactivityBeforeDeviceLock, true
}

// HasMinutesOfInactivityBeforeDeviceLock returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasMinutesOfInactivityBeforeDeviceLock() bool {
	if o != nil && o.MinutesOfInactivityBeforeDeviceLock != nil {
		return true
	}

	return false
}

// SetMinutesOfInactivityBeforeDeviceLock gets a reference to the given int32 and assigns it to the MinutesOfInactivityBeforeDeviceLock field.
func (o *WindowsInformationProtectionPolicy) SetMinutesOfInactivityBeforeDeviceLock(v int32) {
	o.MinutesOfInactivityBeforeDeviceLock = &v
}

// GetNumberOfPastPinsRemembered returns the NumberOfPastPinsRemembered field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRemembered() int32 {
	if o == nil || o.NumberOfPastPinsRemembered == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfPastPinsRemembered
}

// GetNumberOfPastPinsRememberedOk returns a tuple with the NumberOfPastPinsRemembered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRememberedOk() (*int32, bool) {
	if o == nil || o.NumberOfPastPinsRemembered == nil {
		return nil, false
	}
	return o.NumberOfPastPinsRemembered, true
}

// HasNumberOfPastPinsRemembered returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasNumberOfPastPinsRemembered() bool {
	if o != nil && o.NumberOfPastPinsRemembered != nil {
		return true
	}

	return false
}

// SetNumberOfPastPinsRemembered gets a reference to the given int32 and assigns it to the NumberOfPastPinsRemembered field.
func (o *WindowsInformationProtectionPolicy) SetNumberOfPastPinsRemembered(v int32) {
	o.NumberOfPastPinsRemembered = &v
}

// GetPasswordMaximumAttemptCount returns the PasswordMaximumAttemptCount field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCount() int32 {
	if o == nil || o.PasswordMaximumAttemptCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMaximumAttemptCount
}

// GetPasswordMaximumAttemptCountOk returns a tuple with the PasswordMaximumAttemptCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCountOk() (*int32, bool) {
	if o == nil || o.PasswordMaximumAttemptCount == nil {
		return nil, false
	}
	return o.PasswordMaximumAttemptCount, true
}

// HasPasswordMaximumAttemptCount returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasPasswordMaximumAttemptCount() bool {
	if o != nil && o.PasswordMaximumAttemptCount != nil {
		return true
	}

	return false
}

// SetPasswordMaximumAttemptCount gets a reference to the given int32 and assigns it to the PasswordMaximumAttemptCount field.
func (o *WindowsInformationProtectionPolicy) SetPasswordMaximumAttemptCount(v int32) {
	o.PasswordMaximumAttemptCount = &v
}

// GetPinExpirationDays returns the PinExpirationDays field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetPinExpirationDays() int32 {
	if o == nil || o.PinExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.PinExpirationDays
}

// GetPinExpirationDaysOk returns a tuple with the PinExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetPinExpirationDaysOk() (*int32, bool) {
	if o == nil || o.PinExpirationDays == nil {
		return nil, false
	}
	return o.PinExpirationDays, true
}

// HasPinExpirationDays returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasPinExpirationDays() bool {
	if o != nil && o.PinExpirationDays != nil {
		return true
	}

	return false
}

// SetPinExpirationDays gets a reference to the given int32 and assigns it to the PinExpirationDays field.
func (o *WindowsInformationProtectionPolicy) SetPinExpirationDays(v int32) {
	o.PinExpirationDays = &v
}

// GetPinLowercaseLetters returns the PinLowercaseLetters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WindowsInformationProtectionPolicy) GetPinLowercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements
		return ret
	}
	return o.PinLowercaseLetters
}

// GetPinLowercaseLettersOk returns a tuple with the PinLowercaseLetters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WindowsInformationProtectionPolicy) GetPinLowercaseLettersOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool) {
	if o == nil || o.PinLowercaseLetters == nil {
		return nil, false
	}
	return &o.PinLowercaseLetters, true
}

// HasPinLowercaseLetters returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasPinLowercaseLetters() bool {
	if o != nil && o.PinLowercaseLetters != nil {
		return true
	}

	return false
}

// SetPinLowercaseLetters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinLowercaseLetters field.
func (o *WindowsInformationProtectionPolicy) SetPinLowercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) {
	o.PinLowercaseLetters = v
}

// GetPinMinimumLength returns the PinMinimumLength field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetPinMinimumLength() int32 {
	if o == nil || o.PinMinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.PinMinimumLength
}

// GetPinMinimumLengthOk returns a tuple with the PinMinimumLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetPinMinimumLengthOk() (*int32, bool) {
	if o == nil || o.PinMinimumLength == nil {
		return nil, false
	}
	return o.PinMinimumLength, true
}

// HasPinMinimumLength returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasPinMinimumLength() bool {
	if o != nil && o.PinMinimumLength != nil {
		return true
	}

	return false
}

// SetPinMinimumLength gets a reference to the given int32 and assigns it to the PinMinimumLength field.
func (o *WindowsInformationProtectionPolicy) SetPinMinimumLength(v int32) {
	o.PinMinimumLength = &v
}

// GetPinSpecialCharacters returns the PinSpecialCharacters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WindowsInformationProtectionPolicy) GetPinSpecialCharacters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements
		return ret
	}
	return o.PinSpecialCharacters
}

// GetPinSpecialCharactersOk returns a tuple with the PinSpecialCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WindowsInformationProtectionPolicy) GetPinSpecialCharactersOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool) {
	if o == nil || o.PinSpecialCharacters == nil {
		return nil, false
	}
	return &o.PinSpecialCharacters, true
}

// HasPinSpecialCharacters returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasPinSpecialCharacters() bool {
	if o != nil && o.PinSpecialCharacters != nil {
		return true
	}

	return false
}

// SetPinSpecialCharacters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinSpecialCharacters field.
func (o *WindowsInformationProtectionPolicy) SetPinSpecialCharacters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) {
	o.PinSpecialCharacters = v
}

// GetPinUppercaseLetters returns the PinUppercaseLetters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WindowsInformationProtectionPolicy) GetPinUppercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements
		return ret
	}
	return o.PinUppercaseLetters
}

// GetPinUppercaseLettersOk returns a tuple with the PinUppercaseLetters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WindowsInformationProtectionPolicy) GetPinUppercaseLettersOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool) {
	if o == nil || o.PinUppercaseLetters == nil {
		return nil, false
	}
	return &o.PinUppercaseLetters, true
}

// HasPinUppercaseLetters returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasPinUppercaseLetters() bool {
	if o != nil && o.PinUppercaseLetters != nil {
		return true
	}

	return false
}

// SetPinUppercaseLetters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinUppercaseLetters field.
func (o *WindowsInformationProtectionPolicy) SetPinUppercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements) {
	o.PinUppercaseLetters = v
}

// GetRevokeOnMdmHandoffDisabled returns the RevokeOnMdmHandoffDisabled field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabled() bool {
	if o == nil || o.RevokeOnMdmHandoffDisabled == nil {
		var ret bool
		return ret
	}
	return *o.RevokeOnMdmHandoffDisabled
}

// GetRevokeOnMdmHandoffDisabledOk returns a tuple with the RevokeOnMdmHandoffDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabledOk() (*bool, bool) {
	if o == nil || o.RevokeOnMdmHandoffDisabled == nil {
		return nil, false
	}
	return o.RevokeOnMdmHandoffDisabled, true
}

// HasRevokeOnMdmHandoffDisabled returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasRevokeOnMdmHandoffDisabled() bool {
	if o != nil && o.RevokeOnMdmHandoffDisabled != nil {
		return true
	}

	return false
}

// SetRevokeOnMdmHandoffDisabled gets a reference to the given bool and assigns it to the RevokeOnMdmHandoffDisabled field.
func (o *WindowsInformationProtectionPolicy) SetRevokeOnMdmHandoffDisabled(v bool) {
	o.RevokeOnMdmHandoffDisabled = &v
}

// GetWindowsHelloForBusinessBlocked returns the WindowsHelloForBusinessBlocked field value if set, zero value otherwise.
func (o *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlocked() bool {
	if o == nil || o.WindowsHelloForBusinessBlocked == nil {
		var ret bool
		return ret
	}
	return *o.WindowsHelloForBusinessBlocked
}

// GetWindowsHelloForBusinessBlockedOk returns a tuple with the WindowsHelloForBusinessBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlockedOk() (*bool, bool) {
	if o == nil || o.WindowsHelloForBusinessBlocked == nil {
		return nil, false
	}
	return o.WindowsHelloForBusinessBlocked, true
}

// HasWindowsHelloForBusinessBlocked returns a boolean if a field has been set.
func (o *WindowsInformationProtectionPolicy) HasWindowsHelloForBusinessBlocked() bool {
	if o != nil && o.WindowsHelloForBusinessBlocked != nil {
		return true
	}

	return false
}

// SetWindowsHelloForBusinessBlocked gets a reference to the given bool and assigns it to the WindowsHelloForBusinessBlocked field.
func (o *WindowsInformationProtectionPolicy) SetWindowsHelloForBusinessBlocked(v bool) {
	o.WindowsHelloForBusinessBlocked = &v
}

func (o WindowsInformationProtectionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DaysWithoutContactBeforeUnenroll != nil {
		toSerialize["daysWithoutContactBeforeUnenroll"] = o.DaysWithoutContactBeforeUnenroll
	}
	if o.MdmEnrollmentUrl.IsSet() {
		toSerialize["mdmEnrollmentUrl"] = o.MdmEnrollmentUrl.Get()
	}
	if o.MinutesOfInactivityBeforeDeviceLock != nil {
		toSerialize["minutesOfInactivityBeforeDeviceLock"] = o.MinutesOfInactivityBeforeDeviceLock
	}
	if o.NumberOfPastPinsRemembered != nil {
		toSerialize["numberOfPastPinsRemembered"] = o.NumberOfPastPinsRemembered
	}
	if o.PasswordMaximumAttemptCount != nil {
		toSerialize["passwordMaximumAttemptCount"] = o.PasswordMaximumAttemptCount
	}
	if o.PinExpirationDays != nil {
		toSerialize["pinExpirationDays"] = o.PinExpirationDays
	}
	if o.PinLowercaseLetters != nil {
		toSerialize["pinLowercaseLetters"] = o.PinLowercaseLetters
	}
	if o.PinMinimumLength != nil {
		toSerialize["pinMinimumLength"] = o.PinMinimumLength
	}
	if o.PinSpecialCharacters != nil {
		toSerialize["pinSpecialCharacters"] = o.PinSpecialCharacters
	}
	if o.PinUppercaseLetters != nil {
		toSerialize["pinUppercaseLetters"] = o.PinUppercaseLetters
	}
	if o.RevokeOnMdmHandoffDisabled != nil {
		toSerialize["revokeOnMdmHandoffDisabled"] = o.RevokeOnMdmHandoffDisabled
	}
	if o.WindowsHelloForBusinessBlocked != nil {
		toSerialize["windowsHelloForBusinessBlocked"] = o.WindowsHelloForBusinessBlocked
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsInformationProtectionPolicy struct {
	value *WindowsInformationProtectionPolicy
	isSet bool
}

func (v NullableWindowsInformationProtectionPolicy) Get() *WindowsInformationProtectionPolicy {
	return v.value
}

func (v *NullableWindowsInformationProtectionPolicy) Set(val *WindowsInformationProtectionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsInformationProtectionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsInformationProtectionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsInformationProtectionPolicy(val *WindowsInformationProtectionPolicy) *NullableWindowsInformationProtectionPolicy {
	return &NullableWindowsInformationProtectionPolicy{value: val, isSet: true}
}

func (v NullableWindowsInformationProtectionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsInformationProtectionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

