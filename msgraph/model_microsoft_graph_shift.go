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

// MicrosoftGraphShift struct for MicrosoftGraphShift
type MicrosoftGraphShift struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Identity of the person who last modified the entity.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// The draft version of this shift that is viewable by managers. Required.
	DraftShift AnyOfmicrosoftGraphShiftItem `json:"draftShift,omitempty"`
	// ID of the scheduling group the shift is part of. Required.
	SchedulingGroupId NullableString `json:"schedulingGroupId,omitempty"`
	// The shared version of this shift that is viewable by both employees and managers. Required.
	SharedShift AnyOfmicrosoftGraphShiftItem `json:"sharedShift,omitempty"`
	// ID of the user assigned to the shift. Required.
	UserId NullableString `json:"userId,omitempty"`
}

// NewMicrosoftGraphShift instantiates a new MicrosoftGraphShift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphShift() *MicrosoftGraphShift {
	this := MicrosoftGraphShift{}
	return &this
}

// NewMicrosoftGraphShiftWithDefaults instantiates a new MicrosoftGraphShift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphShiftWithDefaults() *MicrosoftGraphShift {
	this := MicrosoftGraphShift{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphShift) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphShift) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphShift) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphShift) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphShift) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphShift) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphShift) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphShift) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphShift) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphShift) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetDraftShift returns the DraftShift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetDraftShift() AnyOfmicrosoftGraphShiftItem {
	if o == nil  {
		var ret AnyOfmicrosoftGraphShiftItem
		return ret
	}
	return o.DraftShift
}

// GetDraftShiftOk returns a tuple with the DraftShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetDraftShiftOk() (*AnyOfmicrosoftGraphShiftItem, bool) {
	if o == nil || o.DraftShift == nil {
		return nil, false
	}
	return &o.DraftShift, true
}

// HasDraftShift returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasDraftShift() bool {
	if o != nil && o.DraftShift != nil {
		return true
	}

	return false
}

// SetDraftShift gets a reference to the given AnyOfmicrosoftGraphShiftItem and assigns it to the DraftShift field.
func (o *MicrosoftGraphShift) SetDraftShift(v AnyOfmicrosoftGraphShiftItem) {
	o.DraftShift = v
}

// GetSchedulingGroupId returns the SchedulingGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetSchedulingGroupId() string {
	if o == nil || o.SchedulingGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SchedulingGroupId.Get()
}

// GetSchedulingGroupIdOk returns a tuple with the SchedulingGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetSchedulingGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchedulingGroupId.Get(), o.SchedulingGroupId.IsSet()
}

// HasSchedulingGroupId returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasSchedulingGroupId() bool {
	if o != nil && o.SchedulingGroupId.IsSet() {
		return true
	}

	return false
}

// SetSchedulingGroupId gets a reference to the given NullableString and assigns it to the SchedulingGroupId field.
func (o *MicrosoftGraphShift) SetSchedulingGroupId(v string) {
	o.SchedulingGroupId.Set(&v)
}
// SetSchedulingGroupIdNil sets the value for SchedulingGroupId to be an explicit nil
func (o *MicrosoftGraphShift) SetSchedulingGroupIdNil() {
	o.SchedulingGroupId.Set(nil)
}

// UnsetSchedulingGroupId ensures that no value is present for SchedulingGroupId, not even an explicit nil
func (o *MicrosoftGraphShift) UnsetSchedulingGroupId() {
	o.SchedulingGroupId.Unset()
}

// GetSharedShift returns the SharedShift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetSharedShift() AnyOfmicrosoftGraphShiftItem {
	if o == nil  {
		var ret AnyOfmicrosoftGraphShiftItem
		return ret
	}
	return o.SharedShift
}

// GetSharedShiftOk returns a tuple with the SharedShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetSharedShiftOk() (*AnyOfmicrosoftGraphShiftItem, bool) {
	if o == nil || o.SharedShift == nil {
		return nil, false
	}
	return &o.SharedShift, true
}

// HasSharedShift returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasSharedShift() bool {
	if o != nil && o.SharedShift != nil {
		return true
	}

	return false
}

// SetSharedShift gets a reference to the given AnyOfmicrosoftGraphShiftItem and assigns it to the SharedShift field.
func (o *MicrosoftGraphShift) SetSharedShift(v AnyOfmicrosoftGraphShiftItem) {
	o.SharedShift = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphShift) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphShift) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MicrosoftGraphShift) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *MicrosoftGraphShift) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MicrosoftGraphShift) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MicrosoftGraphShift) UnsetUserId() {
	o.UserId.Unset()
}

func (o MicrosoftGraphShift) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.DraftShift != nil {
		toSerialize["draftShift"] = o.DraftShift
	}
	if o.SchedulingGroupId.IsSet() {
		toSerialize["schedulingGroupId"] = o.SchedulingGroupId.Get()
	}
	if o.SharedShift != nil {
		toSerialize["sharedShift"] = o.SharedShift
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphShift struct {
	value *MicrosoftGraphShift
	isSet bool
}

func (v NullableMicrosoftGraphShift) Get() *MicrosoftGraphShift {
	return v.value
}

func (v *NullableMicrosoftGraphShift) Set(val *MicrosoftGraphShift) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphShift) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphShift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphShift(val *MicrosoftGraphShift) *NullableMicrosoftGraphShift {
	return &NullableMicrosoftGraphShift{value: val, isSet: true}
}

func (v NullableMicrosoftGraphShift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphShift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


