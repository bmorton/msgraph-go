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

// SchedulingGroup struct for SchedulingGroup
type SchedulingGroup struct {
	// The display name for the schedulingGroup. Required.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
	IsActive NullableBool `json:"isActive,omitempty"`
	// The list of user IDs that are a member of the schedulingGroup. Required.
	UserIds *[]*string `json:"userIds,omitempty"`
}

// NewSchedulingGroup instantiates a new SchedulingGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulingGroup() *SchedulingGroup {
	this := SchedulingGroup{}
	return &this
}

// NewSchedulingGroupWithDefaults instantiates a new SchedulingGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulingGroupWithDefaults() *SchedulingGroup {
	this := SchedulingGroup{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulingGroup) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulingGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchedulingGroup) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *SchedulingGroup) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *SchedulingGroup) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *SchedulingGroup) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulingGroup) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulingGroup) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *SchedulingGroup) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *SchedulingGroup) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *SchedulingGroup) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *SchedulingGroup) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *SchedulingGroup) GetUserIds() []*string {
	if o == nil || o.UserIds == nil {
		var ret []*string
		return ret
	}
	return *o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingGroup) GetUserIdsOk() (*[]*string, bool) {
	if o == nil || o.UserIds == nil {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *SchedulingGroup) HasUserIds() bool {
	if o != nil && o.UserIds != nil {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []*string and assigns it to the UserIds field.
func (o *SchedulingGroup) SetUserIds(v []*string) {
	o.UserIds = &v
}

func (o SchedulingGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	if o.UserIds != nil {
		toSerialize["userIds"] = o.UserIds
	}
	return json.Marshal(toSerialize)
}

type NullableSchedulingGroup struct {
	value *SchedulingGroup
	isSet bool
}

func (v NullableSchedulingGroup) Get() *SchedulingGroup {
	return v.value
}

func (v *NullableSchedulingGroup) Set(val *SchedulingGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulingGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulingGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulingGroup(val *SchedulingGroup) *NullableSchedulingGroup {
	return &NullableSchedulingGroup{value: val, isSet: true}
}

func (v NullableSchedulingGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulingGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


