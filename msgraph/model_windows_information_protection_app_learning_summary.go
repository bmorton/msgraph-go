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

// WindowsInformationProtectionAppLearningSummary Windows Information Protection AppLearning Summary entity.
type WindowsInformationProtectionAppLearningSummary struct {
	// Application Name
	ApplicationName NullableString `json:"applicationName,omitempty"`
	// Application Type. Possible values are: universal, desktop.
	ApplicationType AnyOfmicrosoftGraphApplicationType `json:"applicationType,omitempty"`
	// Device Count
	DeviceCount *int32 `json:"deviceCount,omitempty"`
}

// NewWindowsInformationProtectionAppLearningSummary instantiates a new WindowsInformationProtectionAppLearningSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsInformationProtectionAppLearningSummary() *WindowsInformationProtectionAppLearningSummary {
	this := WindowsInformationProtectionAppLearningSummary{}
	return &this
}

// NewWindowsInformationProtectionAppLearningSummaryWithDefaults instantiates a new WindowsInformationProtectionAppLearningSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsInformationProtectionAppLearningSummaryWithDefaults() *WindowsInformationProtectionAppLearningSummary {
	this := WindowsInformationProtectionAppLearningSummary{}
	return &this
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationName() string {
	if o == nil || o.ApplicationName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName.Get()
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationName.Get(), o.ApplicationName.IsSet()
}

// HasApplicationName returns a boolean if a field has been set.
func (o *WindowsInformationProtectionAppLearningSummary) HasApplicationName() bool {
	if o != nil && o.ApplicationName.IsSet() {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given NullableString and assigns it to the ApplicationName field.
func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationName(v string) {
	o.ApplicationName.Set(&v)
}
// SetApplicationNameNil sets the value for ApplicationName to be an explicit nil
func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationNameNil() {
	o.ApplicationName.Set(nil)
}

// UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
func (o *WindowsInformationProtectionAppLearningSummary) UnsetApplicationName() {
	o.ApplicationName.Unset()
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationType() AnyOfmicrosoftGraphApplicationType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphApplicationType
		return ret
	}
	return o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationTypeOk() (*AnyOfmicrosoftGraphApplicationType, bool) {
	if o == nil || o.ApplicationType == nil {
		return nil, false
	}
	return &o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *WindowsInformationProtectionAppLearningSummary) HasApplicationType() bool {
	if o != nil && o.ApplicationType != nil {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given AnyOfmicrosoftGraphApplicationType and assigns it to the ApplicationType field.
func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationType(v AnyOfmicrosoftGraphApplicationType) {
	o.ApplicationType = v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *WindowsInformationProtectionAppLearningSummary) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsInformationProtectionAppLearningSummary) GetDeviceCountOk() (*int32, bool) {
	if o == nil || o.DeviceCount == nil {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *WindowsInformationProtectionAppLearningSummary) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *WindowsInformationProtectionAppLearningSummary) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

func (o WindowsInformationProtectionAppLearningSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationName.IsSet() {
		toSerialize["applicationName"] = o.ApplicationName.Get()
	}
	if o.ApplicationType != nil {
		toSerialize["applicationType"] = o.ApplicationType
	}
	if o.DeviceCount != nil {
		toSerialize["deviceCount"] = o.DeviceCount
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsInformationProtectionAppLearningSummary struct {
	value *WindowsInformationProtectionAppLearningSummary
	isSet bool
}

func (v NullableWindowsInformationProtectionAppLearningSummary) Get() *WindowsInformationProtectionAppLearningSummary {
	return v.value
}

func (v *NullableWindowsInformationProtectionAppLearningSummary) Set(val *WindowsInformationProtectionAppLearningSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsInformationProtectionAppLearningSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsInformationProtectionAppLearningSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsInformationProtectionAppLearningSummary(val *WindowsInformationProtectionAppLearningSummary) *NullableWindowsInformationProtectionAppLearningSummary {
	return &NullableWindowsInformationProtectionAppLearningSummary{value: val, isSet: true}
}

func (v NullableWindowsInformationProtectionAppLearningSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsInformationProtectionAppLearningSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


