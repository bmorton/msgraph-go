/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// ItemActivity struct for ItemActivity
type ItemActivity struct {
	// An item was accessed.
	Access AnyOfobject `json:"access,omitempty"`
	// Details about when the activity took place. Read-only.
	ActivityDateTime NullableTime `json:"activityDateTime,omitempty"`
	// Identity of who performed the action. Read-only.
	Actor AnyOfmicrosoftGraphIdentitySet `json:"actor,omitempty"`
	// Exposes the driveItem that was the target of this activity.
	DriveItem AnyOfmicrosoftGraphDriveItem `json:"driveItem,omitempty"`
}

// NewItemActivity instantiates a new ItemActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemActivity() *ItemActivity {
	this := ItemActivity{}
	return &this
}

// NewItemActivityWithDefaults instantiates a new ItemActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemActivityWithDefaults() *ItemActivity {
	this := ItemActivity{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemActivity) GetAccess() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemActivity) GetAccessOk() (*AnyOfobject, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return &o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ItemActivity) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AnyOfobject and assigns it to the Access field.
func (o *ItemActivity) SetAccess(v AnyOfobject) {
	o.Access = v
}

// GetActivityDateTime returns the ActivityDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemActivity) GetActivityDateTime() time.Time {
	if o == nil || o.ActivityDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivityDateTime.Get()
}

// GetActivityDateTimeOk returns a tuple with the ActivityDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemActivity) GetActivityDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActivityDateTime.Get(), o.ActivityDateTime.IsSet()
}

// HasActivityDateTime returns a boolean if a field has been set.
func (o *ItemActivity) HasActivityDateTime() bool {
	if o != nil && o.ActivityDateTime.IsSet() {
		return true
	}

	return false
}

// SetActivityDateTime gets a reference to the given NullableTime and assigns it to the ActivityDateTime field.
func (o *ItemActivity) SetActivityDateTime(v time.Time) {
	o.ActivityDateTime.Set(&v)
}
// SetActivityDateTimeNil sets the value for ActivityDateTime to be an explicit nil
func (o *ItemActivity) SetActivityDateTimeNil() {
	o.ActivityDateTime.Set(nil)
}

// UnsetActivityDateTime ensures that no value is present for ActivityDateTime, not even an explicit nil
func (o *ItemActivity) UnsetActivityDateTime() {
	o.ActivityDateTime.Unset()
}

// GetActor returns the Actor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemActivity) GetActor() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemActivity) GetActorOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return &o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *ItemActivity) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Actor field.
func (o *ItemActivity) SetActor(v AnyOfmicrosoftGraphIdentitySet) {
	o.Actor = v
}

// GetDriveItem returns the DriveItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemActivity) GetDriveItem() AnyOfmicrosoftGraphDriveItem {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDriveItem
		return ret
	}
	return o.DriveItem
}

// GetDriveItemOk returns a tuple with the DriveItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemActivity) GetDriveItemOk() (*AnyOfmicrosoftGraphDriveItem, bool) {
	if o == nil || o.DriveItem == nil {
		return nil, false
	}
	return &o.DriveItem, true
}

// HasDriveItem returns a boolean if a field has been set.
func (o *ItemActivity) HasDriveItem() bool {
	if o != nil && o.DriveItem != nil {
		return true
	}

	return false
}

// SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.
func (o *ItemActivity) SetDriveItem(v AnyOfmicrosoftGraphDriveItem) {
	o.DriveItem = v
}

func (o ItemActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.ActivityDateTime.IsSet() {
		toSerialize["activityDateTime"] = o.ActivityDateTime.Get()
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.DriveItem != nil {
		toSerialize["driveItem"] = o.DriveItem
	}
	return json.Marshal(toSerialize)
}

type NullableItemActivity struct {
	value *ItemActivity
	isSet bool
}

func (v NullableItemActivity) Get() *ItemActivity {
	return v.value
}

func (v *NullableItemActivity) Set(val *ItemActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableItemActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableItemActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemActivity(val *ItemActivity) *NullableItemActivity {
	return &NullableItemActivity{value: val, isSet: true}
}

func (v NullableItemActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


