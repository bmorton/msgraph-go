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

// MicrosoftGraphLicenseUnitsDetail struct for MicrosoftGraphLicenseUnitsDetail
type MicrosoftGraphLicenseUnitsDetail struct {
	// The number of units that are enabled for the active subscription of the service SKU.
	Enabled NullableInt32 `json:"enabled,omitempty"`
	// The number of units that are suspended because the subscription of the service SKU has been cancelled. The units cannot be assigned but can still be reactivated before they are deleted.
	Suspended NullableInt32 `json:"suspended,omitempty"`
	// The number of units that are in warning status. When the subscription of the service SKU has expired, the customer has a grace period to renew their subscription before it is cancelled (moved to a suspended state).
	Warning NullableInt32 `json:"warning,omitempty"`
}

// NewMicrosoftGraphLicenseUnitsDetail instantiates a new MicrosoftGraphLicenseUnitsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphLicenseUnitsDetail() *MicrosoftGraphLicenseUnitsDetail {
	this := MicrosoftGraphLicenseUnitsDetail{}
	return &this
}

// NewMicrosoftGraphLicenseUnitsDetailWithDefaults instantiates a new MicrosoftGraphLicenseUnitsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphLicenseUnitsDetailWithDefaults() *MicrosoftGraphLicenseUnitsDetail {
	this := MicrosoftGraphLicenseUnitsDetail{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLicenseUnitsDetail) GetEnabled() int32 {
	if o == nil || o.Enabled.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLicenseUnitsDetail) GetEnabledOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphLicenseUnitsDetail) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableInt32 and assigns it to the Enabled field.
func (o *MicrosoftGraphLicenseUnitsDetail) SetEnabled(v int32) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *MicrosoftGraphLicenseUnitsDetail) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *MicrosoftGraphLicenseUnitsDetail) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetSuspended returns the Suspended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLicenseUnitsDetail) GetSuspended() int32 {
	if o == nil || o.Suspended.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Suspended.Get()
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLicenseUnitsDetail) GetSuspendedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suspended.Get(), o.Suspended.IsSet()
}

// HasSuspended returns a boolean if a field has been set.
func (o *MicrosoftGraphLicenseUnitsDetail) HasSuspended() bool {
	if o != nil && o.Suspended.IsSet() {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given NullableInt32 and assigns it to the Suspended field.
func (o *MicrosoftGraphLicenseUnitsDetail) SetSuspended(v int32) {
	o.Suspended.Set(&v)
}
// SetSuspendedNil sets the value for Suspended to be an explicit nil
func (o *MicrosoftGraphLicenseUnitsDetail) SetSuspendedNil() {
	o.Suspended.Set(nil)
}

// UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
func (o *MicrosoftGraphLicenseUnitsDetail) UnsetSuspended() {
	o.Suspended.Unset()
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLicenseUnitsDetail) GetWarning() int32 {
	if o == nil || o.Warning.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLicenseUnitsDetail) GetWarningOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *MicrosoftGraphLicenseUnitsDetail) HasWarning() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableInt32 and assigns it to the Warning field.
func (o *MicrosoftGraphLicenseUnitsDetail) SetWarning(v int32) {
	o.Warning.Set(&v)
}
// SetWarningNil sets the value for Warning to be an explicit nil
func (o *MicrosoftGraphLicenseUnitsDetail) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil
func (o *MicrosoftGraphLicenseUnitsDetail) UnsetWarning() {
	o.Warning.Unset()
}

func (o MicrosoftGraphLicenseUnitsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.Suspended.IsSet() {
		toSerialize["suspended"] = o.Suspended.Get()
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphLicenseUnitsDetail struct {
	value *MicrosoftGraphLicenseUnitsDetail
	isSet bool
}

func (v NullableMicrosoftGraphLicenseUnitsDetail) Get() *MicrosoftGraphLicenseUnitsDetail {
	return v.value
}

func (v *NullableMicrosoftGraphLicenseUnitsDetail) Set(val *MicrosoftGraphLicenseUnitsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphLicenseUnitsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphLicenseUnitsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphLicenseUnitsDetail(val *MicrosoftGraphLicenseUnitsDetail) *NullableMicrosoftGraphLicenseUnitsDetail {
	return &NullableMicrosoftGraphLicenseUnitsDetail{value: val, isSet: true}
}

func (v NullableMicrosoftGraphLicenseUnitsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphLicenseUnitsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

