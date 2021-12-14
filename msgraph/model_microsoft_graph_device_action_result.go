/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphDeviceActionResult Device action result
type MicrosoftGraphDeviceActionResult struct {
	// Action name
	ActionName NullableString `json:"actionName,omitempty"`
	// State of the action. Possible values are: none, pending, canceled, active, done, failed, notSupported.
	ActionState AnyOfmicrosoftGraphActionState `json:"actionState,omitempty"`
	// Time the action state was last updated
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
	// Time the action was initiated
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

// NewMicrosoftGraphDeviceActionResult instantiates a new MicrosoftGraphDeviceActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceActionResult() *MicrosoftGraphDeviceActionResult {
	this := MicrosoftGraphDeviceActionResult{}
	return &this
}

// NewMicrosoftGraphDeviceActionResultWithDefaults instantiates a new MicrosoftGraphDeviceActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceActionResultWithDefaults() *MicrosoftGraphDeviceActionResult {
	this := MicrosoftGraphDeviceActionResult{}
	return &this
}

// GetActionName returns the ActionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceActionResult) GetActionName() string {
	if o == nil || o.ActionName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionName.Get()
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceActionResult) GetActionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionName.Get(), o.ActionName.IsSet()
}

// HasActionName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceActionResult) HasActionName() bool {
	if o != nil && o.ActionName.IsSet() {
		return true
	}

	return false
}

// SetActionName gets a reference to the given NullableString and assigns it to the ActionName field.
func (o *MicrosoftGraphDeviceActionResult) SetActionName(v string) {
	o.ActionName.Set(&v)
}
// SetActionNameNil sets the value for ActionName to be an explicit nil
func (o *MicrosoftGraphDeviceActionResult) SetActionNameNil() {
	o.ActionName.Set(nil)
}

// UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
func (o *MicrosoftGraphDeviceActionResult) UnsetActionName() {
	o.ActionName.Unset()
}

// GetActionState returns the ActionState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceActionResult) GetActionState() AnyOfmicrosoftGraphActionState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphActionState
		return ret
	}
	return o.ActionState
}

// GetActionStateOk returns a tuple with the ActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceActionResult) GetActionStateOk() (*AnyOfmicrosoftGraphActionState, bool) {
	if o == nil || o.ActionState == nil {
		return nil, false
	}
	return &o.ActionState, true
}

// HasActionState returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceActionResult) HasActionState() bool {
	if o != nil && o.ActionState != nil {
		return true
	}

	return false
}

// SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.
func (o *MicrosoftGraphDeviceActionResult) SetActionState(v AnyOfmicrosoftGraphActionState) {
	o.ActionState = v
}

// GetLastUpdatedDateTime returns the LastUpdatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceActionResult) GetLastUpdatedDateTime() time.Time {
	if o == nil || o.LastUpdatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDateTime
}

// GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceActionResult) GetLastUpdatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedDateTime == nil {
		return nil, false
	}
	return o.LastUpdatedDateTime, true
}

// HasLastUpdatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceActionResult) HasLastUpdatedDateTime() bool {
	if o != nil && o.LastUpdatedDateTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.
func (o *MicrosoftGraphDeviceActionResult) SetLastUpdatedDateTime(v time.Time) {
	o.LastUpdatedDateTime = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceActionResult) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceActionResult) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceActionResult) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MicrosoftGraphDeviceActionResult) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

func (o MicrosoftGraphDeviceActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionName.IsSet() {
		toSerialize["actionName"] = o.ActionName.Get()
	}
	if o.ActionState != nil {
		toSerialize["actionState"] = o.ActionState
	}
	if o.LastUpdatedDateTime != nil {
		toSerialize["lastUpdatedDateTime"] = o.LastUpdatedDateTime
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceActionResult struct {
	value *MicrosoftGraphDeviceActionResult
	isSet bool
}

func (v NullableMicrosoftGraphDeviceActionResult) Get() *MicrosoftGraphDeviceActionResult {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceActionResult) Set(val *MicrosoftGraphDeviceActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceActionResult(val *MicrosoftGraphDeviceActionResult) *NullableMicrosoftGraphDeviceActionResult {
	return &NullableMicrosoftGraphDeviceActionResult{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


