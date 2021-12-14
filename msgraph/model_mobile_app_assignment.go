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

// MobileAppAssignment A class containing the properties used for Group Assignment of a Mobile App.
type MobileAppAssignment struct {
	// The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
	Intent AnyOfmicrosoftGraphInstallIntent `json:"intent,omitempty"`
	// The settings for target assignment defined by the admin.
	Settings AnyOfobject `json:"settings,omitempty"`
	// The target group assignment defined by the admin.
	Target AnyOfobject `json:"target,omitempty"`
}

// NewMobileAppAssignment instantiates a new MobileAppAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileAppAssignment() *MobileAppAssignment {
	this := MobileAppAssignment{}
	return &this
}

// NewMobileAppAssignmentWithDefaults instantiates a new MobileAppAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileAppAssignmentWithDefaults() *MobileAppAssignment {
	this := MobileAppAssignment{}
	return &this
}

// GetIntent returns the Intent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MobileAppAssignment) GetIntent() AnyOfmicrosoftGraphInstallIntent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphInstallIntent
		return ret
	}
	return o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MobileAppAssignment) GetIntentOk() (*AnyOfmicrosoftGraphInstallIntent, bool) {
	if o == nil || o.Intent == nil {
		return nil, false
	}
	return &o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *MobileAppAssignment) HasIntent() bool {
	if o != nil && o.Intent != nil {
		return true
	}

	return false
}

// SetIntent gets a reference to the given AnyOfmicrosoftGraphInstallIntent and assigns it to the Intent field.
func (o *MobileAppAssignment) SetIntent(v AnyOfmicrosoftGraphInstallIntent) {
	o.Intent = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MobileAppAssignment) GetSettings() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MobileAppAssignment) GetSettingsOk() (*AnyOfobject, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *MobileAppAssignment) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AnyOfobject and assigns it to the Settings field.
func (o *MobileAppAssignment) SetSettings(v AnyOfobject) {
	o.Settings = v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MobileAppAssignment) GetTarget() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MobileAppAssignment) GetTargetOk() (*AnyOfobject, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *MobileAppAssignment) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.
func (o *MobileAppAssignment) SetTarget(v AnyOfobject) {
	o.Target = v
}

func (o MobileAppAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Intent != nil {
		toSerialize["intent"] = o.Intent
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableMobileAppAssignment struct {
	value *MobileAppAssignment
	isSet bool
}

func (v NullableMobileAppAssignment) Get() *MobileAppAssignment {
	return v.value
}

func (v *NullableMobileAppAssignment) Set(val *MobileAppAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileAppAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileAppAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileAppAssignment(val *MobileAppAssignment) *NullableMobileAppAssignment {
	return &NullableMobileAppAssignment{value: val, isSet: true}
}

func (v NullableMobileAppAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileAppAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

