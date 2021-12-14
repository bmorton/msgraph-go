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

// ShiftItem struct for ShiftItem
type ShiftItem struct {
	// An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
	Activities *[]*AnyOfmicrosoftGraphShiftActivity `json:"activities,omitempty"`
	// The shift label of the shiftItem.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The shift notes for the shiftItem.
	Notes NullableString `json:"notes,omitempty"`
}

// NewShiftItem instantiates a new ShiftItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShiftItem() *ShiftItem {
	this := ShiftItem{}
	return &this
}

// NewShiftItemWithDefaults instantiates a new ShiftItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShiftItemWithDefaults() *ShiftItem {
	this := ShiftItem{}
	return &this
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *ShiftItem) GetActivities() []*AnyOfmicrosoftGraphShiftActivity {
	if o == nil || o.Activities == nil {
		var ret []*AnyOfmicrosoftGraphShiftActivity
		return ret
	}
	return *o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShiftItem) GetActivitiesOk() (*[]*AnyOfmicrosoftGraphShiftActivity, bool) {
	if o == nil || o.Activities == nil {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *ShiftItem) HasActivities() bool {
	if o != nil && o.Activities != nil {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []*AnyOfmicrosoftGraphShiftActivity and assigns it to the Activities field.
func (o *ShiftItem) SetActivities(v []*AnyOfmicrosoftGraphShiftActivity) {
	o.Activities = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShiftItem) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShiftItem) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ShiftItem) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ShiftItem) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ShiftItem) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ShiftItem) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShiftItem) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShiftItem) GetNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *ShiftItem) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *ShiftItem) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *ShiftItem) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *ShiftItem) UnsetNotes() {
	o.Notes.Unset()
}

func (o ShiftItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activities != nil {
		toSerialize["activities"] = o.Activities
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableShiftItem struct {
	value *ShiftItem
	isSet bool
}

func (v NullableShiftItem) Get() *ShiftItem {
	return v.value
}

func (v *NullableShiftItem) Set(val *ShiftItem) {
	v.value = val
	v.isSet = true
}

func (v NullableShiftItem) IsSet() bool {
	return v.isSet
}

func (v *NullableShiftItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShiftItem(val *ShiftItem) *NullableShiftItem {
	return &NullableShiftItem{value: val, isSet: true}
}

func (v NullableShiftItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShiftItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


