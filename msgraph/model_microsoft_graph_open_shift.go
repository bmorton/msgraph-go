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

// MicrosoftGraphOpenShift struct for MicrosoftGraphOpenShift
type MicrosoftGraphOpenShift struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Identity of the person who last modified the entity.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// An unpublished open shift.
	DraftOpenShift AnyOfmicrosoftGraphOpenShiftItem `json:"draftOpenShift,omitempty"`
	// ID for the scheduling group that the open shift belongs to.
	SchedulingGroupId NullableString `json:"schedulingGroupId,omitempty"`
	// A published open shift.
	SharedOpenShift AnyOfmicrosoftGraphOpenShiftItem `json:"sharedOpenShift,omitempty"`
}

// NewMicrosoftGraphOpenShift instantiates a new MicrosoftGraphOpenShift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOpenShift() *MicrosoftGraphOpenShift {
	this := MicrosoftGraphOpenShift{}
	return &this
}

// NewMicrosoftGraphOpenShiftWithDefaults instantiates a new MicrosoftGraphOpenShift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOpenShiftWithDefaults() *MicrosoftGraphOpenShift {
	this := MicrosoftGraphOpenShift{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOpenShift) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOpenShift) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOpenShift) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShift) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShift) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOpenShift) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphOpenShift) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphOpenShift) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShift) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShift) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphOpenShift) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShift) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShift) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphOpenShift) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphOpenShift) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphOpenShift) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetDraftOpenShift returns the DraftOpenShift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShift) GetDraftOpenShift() AnyOfmicrosoftGraphOpenShiftItem {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOpenShiftItem
		return ret
	}
	return o.DraftOpenShift
}

// GetDraftOpenShiftOk returns a tuple with the DraftOpenShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShift) GetDraftOpenShiftOk() (*AnyOfmicrosoftGraphOpenShiftItem, bool) {
	if o == nil || o.DraftOpenShift == nil {
		return nil, false
	}
	return &o.DraftOpenShift, true
}

// HasDraftOpenShift returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasDraftOpenShift() bool {
	if o != nil && o.DraftOpenShift != nil {
		return true
	}

	return false
}

// SetDraftOpenShift gets a reference to the given AnyOfmicrosoftGraphOpenShiftItem and assigns it to the DraftOpenShift field.
func (o *MicrosoftGraphOpenShift) SetDraftOpenShift(v AnyOfmicrosoftGraphOpenShiftItem) {
	o.DraftOpenShift = v
}

// GetSchedulingGroupId returns the SchedulingGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShift) GetSchedulingGroupId() string {
	if o == nil || o.SchedulingGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SchedulingGroupId.Get()
}

// GetSchedulingGroupIdOk returns a tuple with the SchedulingGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShift) GetSchedulingGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SchedulingGroupId.Get(), o.SchedulingGroupId.IsSet()
}

// HasSchedulingGroupId returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasSchedulingGroupId() bool {
	if o != nil && o.SchedulingGroupId.IsSet() {
		return true
	}

	return false
}

// SetSchedulingGroupId gets a reference to the given NullableString and assigns it to the SchedulingGroupId field.
func (o *MicrosoftGraphOpenShift) SetSchedulingGroupId(v string) {
	o.SchedulingGroupId.Set(&v)
}
// SetSchedulingGroupIdNil sets the value for SchedulingGroupId to be an explicit nil
func (o *MicrosoftGraphOpenShift) SetSchedulingGroupIdNil() {
	o.SchedulingGroupId.Set(nil)
}

// UnsetSchedulingGroupId ensures that no value is present for SchedulingGroupId, not even an explicit nil
func (o *MicrosoftGraphOpenShift) UnsetSchedulingGroupId() {
	o.SchedulingGroupId.Unset()
}

// GetSharedOpenShift returns the SharedOpenShift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOpenShift) GetSharedOpenShift() AnyOfmicrosoftGraphOpenShiftItem {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOpenShiftItem
		return ret
	}
	return o.SharedOpenShift
}

// GetSharedOpenShiftOk returns a tuple with the SharedOpenShift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOpenShift) GetSharedOpenShiftOk() (*AnyOfmicrosoftGraphOpenShiftItem, bool) {
	if o == nil || o.SharedOpenShift == nil {
		return nil, false
	}
	return &o.SharedOpenShift, true
}

// HasSharedOpenShift returns a boolean if a field has been set.
func (o *MicrosoftGraphOpenShift) HasSharedOpenShift() bool {
	if o != nil && o.SharedOpenShift != nil {
		return true
	}

	return false
}

// SetSharedOpenShift gets a reference to the given AnyOfmicrosoftGraphOpenShiftItem and assigns it to the SharedOpenShift field.
func (o *MicrosoftGraphOpenShift) SetSharedOpenShift(v AnyOfmicrosoftGraphOpenShiftItem) {
	o.SharedOpenShift = v
}

func (o MicrosoftGraphOpenShift) MarshalJSON() ([]byte, error) {
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
	if o.DraftOpenShift != nil {
		toSerialize["draftOpenShift"] = o.DraftOpenShift
	}
	if o.SchedulingGroupId.IsSet() {
		toSerialize["schedulingGroupId"] = o.SchedulingGroupId.Get()
	}
	if o.SharedOpenShift != nil {
		toSerialize["sharedOpenShift"] = o.SharedOpenShift
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOpenShift struct {
	value *MicrosoftGraphOpenShift
	isSet bool
}

func (v NullableMicrosoftGraphOpenShift) Get() *MicrosoftGraphOpenShift {
	return v.value
}

func (v *NullableMicrosoftGraphOpenShift) Set(val *MicrosoftGraphOpenShift) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOpenShift) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOpenShift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOpenShift(val *MicrosoftGraphOpenShift) *NullableMicrosoftGraphOpenShift {
	return &NullableMicrosoftGraphOpenShift{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOpenShift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOpenShift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


