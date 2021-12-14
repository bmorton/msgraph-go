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

// MicrosoftGraphCloudAppSecuritySessionControl struct for MicrosoftGraphCloudAppSecuritySessionControl
type MicrosoftGraphCloudAppSecuritySessionControl struct {
	// Specifies whether the session control is enabled.
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
	// Possible values are: mcasConfigured, monitorOnly, blockDownloads, unknownFutureValue. For more information, see Deploy Conditional Access App Control for featured apps.
	CloudAppSecurityType AnyOfmicrosoftGraphCloudAppSecuritySessionControlType `json:"cloudAppSecurityType,omitempty"`
}

// NewMicrosoftGraphCloudAppSecuritySessionControl instantiates a new MicrosoftGraphCloudAppSecuritySessionControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphCloudAppSecuritySessionControl() *MicrosoftGraphCloudAppSecuritySessionControl {
	this := MicrosoftGraphCloudAppSecuritySessionControl{}
	return &this
}

// NewMicrosoftGraphCloudAppSecuritySessionControlWithDefaults instantiates a new MicrosoftGraphCloudAppSecuritySessionControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphCloudAppSecuritySessionControlWithDefaults() *MicrosoftGraphCloudAppSecuritySessionControl {
	this := MicrosoftGraphCloudAppSecuritySessionControl{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetIsEnabled() bool {
	if o == nil || o.IsEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphCloudAppSecuritySessionControl) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *MicrosoftGraphCloudAppSecuritySessionControl) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetCloudAppSecurityType returns the CloudAppSecurityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetCloudAppSecurityType() AnyOfmicrosoftGraphCloudAppSecuritySessionControlType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCloudAppSecuritySessionControlType
		return ret
	}
	return o.CloudAppSecurityType
}

// GetCloudAppSecurityTypeOk returns a tuple with the CloudAppSecurityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCloudAppSecuritySessionControl) GetCloudAppSecurityTypeOk() (*AnyOfmicrosoftGraphCloudAppSecuritySessionControlType, bool) {
	if o == nil || o.CloudAppSecurityType == nil {
		return nil, false
	}
	return &o.CloudAppSecurityType, true
}

// HasCloudAppSecurityType returns a boolean if a field has been set.
func (o *MicrosoftGraphCloudAppSecuritySessionControl) HasCloudAppSecurityType() bool {
	if o != nil && o.CloudAppSecurityType != nil {
		return true
	}

	return false
}

// SetCloudAppSecurityType gets a reference to the given AnyOfmicrosoftGraphCloudAppSecuritySessionControlType and assigns it to the CloudAppSecurityType field.
func (o *MicrosoftGraphCloudAppSecuritySessionControl) SetCloudAppSecurityType(v AnyOfmicrosoftGraphCloudAppSecuritySessionControlType) {
	o.CloudAppSecurityType = v
}

func (o MicrosoftGraphCloudAppSecuritySessionControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	if o.CloudAppSecurityType != nil {
		toSerialize["cloudAppSecurityType"] = o.CloudAppSecurityType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphCloudAppSecuritySessionControl struct {
	value *MicrosoftGraphCloudAppSecuritySessionControl
	isSet bool
}

func (v NullableMicrosoftGraphCloudAppSecuritySessionControl) Get() *MicrosoftGraphCloudAppSecuritySessionControl {
	return v.value
}

func (v *NullableMicrosoftGraphCloudAppSecuritySessionControl) Set(val *MicrosoftGraphCloudAppSecuritySessionControl) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCloudAppSecuritySessionControl) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCloudAppSecuritySessionControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCloudAppSecuritySessionControl(val *MicrosoftGraphCloudAppSecuritySessionControl) *NullableMicrosoftGraphCloudAppSecuritySessionControl {
	return &NullableMicrosoftGraphCloudAppSecuritySessionControl{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCloudAppSecuritySessionControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCloudAppSecuritySessionControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


