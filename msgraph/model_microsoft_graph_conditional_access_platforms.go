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

// MicrosoftGraphConditionalAccessPlatforms struct for MicrosoftGraphConditionalAccessPlatforms
type MicrosoftGraphConditionalAccessPlatforms struct {
	// Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
	ExcludePlatforms *[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform `json:"excludePlatforms,omitempty"`
	// Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
	IncludePlatforms *[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform `json:"includePlatforms,omitempty"`
}

// NewMicrosoftGraphConditionalAccessPlatforms instantiates a new MicrosoftGraphConditionalAccessPlatforms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphConditionalAccessPlatforms() *MicrosoftGraphConditionalAccessPlatforms {
	this := MicrosoftGraphConditionalAccessPlatforms{}
	return &this
}

// NewMicrosoftGraphConditionalAccessPlatformsWithDefaults instantiates a new MicrosoftGraphConditionalAccessPlatforms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphConditionalAccessPlatformsWithDefaults() *MicrosoftGraphConditionalAccessPlatforms {
	this := MicrosoftGraphConditionalAccessPlatforms{}
	return &this
}

// GetExcludePlatforms returns the ExcludePlatforms field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessPlatforms) GetExcludePlatforms() []AnyOfmicrosoftGraphConditionalAccessDevicePlatform {
	if o == nil || o.ExcludePlatforms == nil {
		var ret []AnyOfmicrosoftGraphConditionalAccessDevicePlatform
		return ret
	}
	return *o.ExcludePlatforms
}

// GetExcludePlatformsOk returns a tuple with the ExcludePlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessPlatforms) GetExcludePlatformsOk() (*[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform, bool) {
	if o == nil || o.ExcludePlatforms == nil {
		return nil, false
	}
	return o.ExcludePlatforms, true
}

// HasExcludePlatforms returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPlatforms) HasExcludePlatforms() bool {
	if o != nil && o.ExcludePlatforms != nil {
		return true
	}

	return false
}

// SetExcludePlatforms gets a reference to the given []AnyOfmicrosoftGraphConditionalAccessDevicePlatform and assigns it to the ExcludePlatforms field.
func (o *MicrosoftGraphConditionalAccessPlatforms) SetExcludePlatforms(v []AnyOfmicrosoftGraphConditionalAccessDevicePlatform) {
	o.ExcludePlatforms = &v
}

// GetIncludePlatforms returns the IncludePlatforms field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessPlatforms) GetIncludePlatforms() []AnyOfmicrosoftGraphConditionalAccessDevicePlatform {
	if o == nil || o.IncludePlatforms == nil {
		var ret []AnyOfmicrosoftGraphConditionalAccessDevicePlatform
		return ret
	}
	return *o.IncludePlatforms
}

// GetIncludePlatformsOk returns a tuple with the IncludePlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessPlatforms) GetIncludePlatformsOk() (*[]AnyOfmicrosoftGraphConditionalAccessDevicePlatform, bool) {
	if o == nil || o.IncludePlatforms == nil {
		return nil, false
	}
	return o.IncludePlatforms, true
}

// HasIncludePlatforms returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPlatforms) HasIncludePlatforms() bool {
	if o != nil && o.IncludePlatforms != nil {
		return true
	}

	return false
}

// SetIncludePlatforms gets a reference to the given []AnyOfmicrosoftGraphConditionalAccessDevicePlatform and assigns it to the IncludePlatforms field.
func (o *MicrosoftGraphConditionalAccessPlatforms) SetIncludePlatforms(v []AnyOfmicrosoftGraphConditionalAccessDevicePlatform) {
	o.IncludePlatforms = &v
}

func (o MicrosoftGraphConditionalAccessPlatforms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludePlatforms != nil {
		toSerialize["excludePlatforms"] = o.ExcludePlatforms
	}
	if o.IncludePlatforms != nil {
		toSerialize["includePlatforms"] = o.IncludePlatforms
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphConditionalAccessPlatforms struct {
	value *MicrosoftGraphConditionalAccessPlatforms
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessPlatforms) Get() *MicrosoftGraphConditionalAccessPlatforms {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessPlatforms) Set(val *MicrosoftGraphConditionalAccessPlatforms) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessPlatforms) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessPlatforms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessPlatforms(val *MicrosoftGraphConditionalAccessPlatforms) *NullableMicrosoftGraphConditionalAccessPlatforms {
	return &NullableMicrosoftGraphConditionalAccessPlatforms{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessPlatforms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessPlatforms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


