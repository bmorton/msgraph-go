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

// MicrosoftGraphActivityHistoryItem struct for MicrosoftGraphActivityHistoryItem
type MicrosoftGraphActivityHistoryItem struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Optional. The duration of active user engagement. if not supplied, this is calculated from the startedDateTime and lastActiveDateTime.
	ActiveDurationSeconds NullableInt32 `json:"activeDurationSeconds,omitempty"`
	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Optional. UTC DateTime when the historyItem will undergo hard-delete. Can be set by the client.
	ExpirationDateTime NullableTime `json:"expirationDateTime,omitempty"`
	// Optional. UTC DateTime when the historyItem (activity session) was last understood as active or finished - if null, historyItem status should be Ongoing.
	LastActiveDateTime NullableTime `json:"lastActiveDateTime,omitempty"`
	// Set by the server. DateTime in UTC when the object was modified on the server.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// Required. UTC DateTime when the historyItem (activity session) was started. Required for timeline history.
	StartedDateTime *time.Time `json:"startedDateTime,omitempty"`
	// Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
	Status AnyOfmicrosoftGraphStatus `json:"status,omitempty"`
	// Optional. The timezone in which the user's device used to generate the activity was located at activity creation time. Values supplied as Olson IDs in order to support cross-platform representation.
	UserTimezone NullableString `json:"userTimezone,omitempty"`
	Activity *MicrosoftGraphUserActivity `json:"activity,omitempty"`
}

// NewMicrosoftGraphActivityHistoryItem instantiates a new MicrosoftGraphActivityHistoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphActivityHistoryItem() *MicrosoftGraphActivityHistoryItem {
	this := MicrosoftGraphActivityHistoryItem{}
	return &this
}

// NewMicrosoftGraphActivityHistoryItemWithDefaults instantiates a new MicrosoftGraphActivityHistoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphActivityHistoryItemWithDefaults() *MicrosoftGraphActivityHistoryItem {
	this := MicrosoftGraphActivityHistoryItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphActivityHistoryItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphActivityHistoryItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphActivityHistoryItem) SetId(v string) {
	o.Id = &v
}

// GetActiveDurationSeconds returns the ActiveDurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetActiveDurationSeconds() int32 {
	if o == nil || o.ActiveDurationSeconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ActiveDurationSeconds.Get()
}

// GetActiveDurationSecondsOk returns a tuple with the ActiveDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetActiveDurationSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActiveDurationSeconds.Get(), o.ActiveDurationSeconds.IsSet()
}

// HasActiveDurationSeconds returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasActiveDurationSeconds() bool {
	if o != nil && o.ActiveDurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetActiveDurationSeconds gets a reference to the given NullableInt32 and assigns it to the ActiveDurationSeconds field.
func (o *MicrosoftGraphActivityHistoryItem) SetActiveDurationSeconds(v int32) {
	o.ActiveDurationSeconds.Set(&v)
}
// SetActiveDurationSecondsNil sets the value for ActiveDurationSeconds to be an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) SetActiveDurationSecondsNil() {
	o.ActiveDurationSeconds.Set(nil)
}

// UnsetActiveDurationSeconds ensures that no value is present for ActiveDurationSeconds, not even an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) UnsetActiveDurationSeconds() {
	o.ActiveDurationSeconds.Unset()
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphActivityHistoryItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime.Get()
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDateTime.Get(), o.ExpirationDateTime.IsSet()
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given NullableTime and assigns it to the ExpirationDateTime field.
func (o *MicrosoftGraphActivityHistoryItem) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime.Set(&v)
}
// SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) SetExpirationDateTimeNil() {
	o.ExpirationDateTime.Set(nil)
}

// UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) UnsetExpirationDateTime() {
	o.ExpirationDateTime.Unset()
}

// GetLastActiveDateTime returns the LastActiveDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetLastActiveDateTime() time.Time {
	if o == nil || o.LastActiveDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActiveDateTime.Get()
}

// GetLastActiveDateTimeOk returns a tuple with the LastActiveDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetLastActiveDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActiveDateTime.Get(), o.LastActiveDateTime.IsSet()
}

// HasLastActiveDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasLastActiveDateTime() bool {
	if o != nil && o.LastActiveDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastActiveDateTime gets a reference to the given NullableTime and assigns it to the LastActiveDateTime field.
func (o *MicrosoftGraphActivityHistoryItem) SetLastActiveDateTime(v time.Time) {
	o.LastActiveDateTime.Set(&v)
}
// SetLastActiveDateTimeNil sets the value for LastActiveDateTime to be an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) SetLastActiveDateTimeNil() {
	o.LastActiveDateTime.Set(nil)
}

// UnsetLastActiveDateTime ensures that no value is present for LastActiveDateTime, not even an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) UnsetLastActiveDateTime() {
	o.LastActiveDateTime.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphActivityHistoryItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetStartedDateTime returns the StartedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphActivityHistoryItem) GetStartedDateTime() time.Time {
	if o == nil || o.StartedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedDateTime
}

// GetStartedDateTimeOk returns a tuple with the StartedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphActivityHistoryItem) GetStartedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartedDateTime == nil {
		return nil, false
	}
	return o.StartedDateTime, true
}

// HasStartedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasStartedDateTime() bool {
	if o != nil && o.StartedDateTime != nil {
		return true
	}

	return false
}

// SetStartedDateTime gets a reference to the given time.Time and assigns it to the StartedDateTime field.
func (o *MicrosoftGraphActivityHistoryItem) SetStartedDateTime(v time.Time) {
	o.StartedDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetStatus() AnyOfmicrosoftGraphStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetStatusOk() (*AnyOfmicrosoftGraphStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphStatus and assigns it to the Status field.
func (o *MicrosoftGraphActivityHistoryItem) SetStatus(v AnyOfmicrosoftGraphStatus) {
	o.Status = v
}

// GetUserTimezone returns the UserTimezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActivityHistoryItem) GetUserTimezone() string {
	if o == nil || o.UserTimezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserTimezone.Get()
}

// GetUserTimezoneOk returns a tuple with the UserTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActivityHistoryItem) GetUserTimezoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserTimezone.Get(), o.UserTimezone.IsSet()
}

// HasUserTimezone returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasUserTimezone() bool {
	if o != nil && o.UserTimezone.IsSet() {
		return true
	}

	return false
}

// SetUserTimezone gets a reference to the given NullableString and assigns it to the UserTimezone field.
func (o *MicrosoftGraphActivityHistoryItem) SetUserTimezone(v string) {
	o.UserTimezone.Set(&v)
}
// SetUserTimezoneNil sets the value for UserTimezone to be an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) SetUserTimezoneNil() {
	o.UserTimezone.Set(nil)
}

// UnsetUserTimezone ensures that no value is present for UserTimezone, not even an explicit nil
func (o *MicrosoftGraphActivityHistoryItem) UnsetUserTimezone() {
	o.UserTimezone.Unset()
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *MicrosoftGraphActivityHistoryItem) GetActivity() MicrosoftGraphUserActivity {
	if o == nil || o.Activity == nil {
		var ret MicrosoftGraphUserActivity
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphActivityHistoryItem) GetActivityOk() (*MicrosoftGraphUserActivity, bool) {
	if o == nil || o.Activity == nil {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *MicrosoftGraphActivityHistoryItem) HasActivity() bool {
	if o != nil && o.Activity != nil {
		return true
	}

	return false
}

// SetActivity gets a reference to the given MicrosoftGraphUserActivity and assigns it to the Activity field.
func (o *MicrosoftGraphActivityHistoryItem) SetActivity(v MicrosoftGraphUserActivity) {
	o.Activity = &v
}

func (o MicrosoftGraphActivityHistoryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ActiveDurationSeconds.IsSet() {
		toSerialize["activeDurationSeconds"] = o.ActiveDurationSeconds.Get()
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.ExpirationDateTime.IsSet() {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime.Get()
	}
	if o.LastActiveDateTime.IsSet() {
		toSerialize["lastActiveDateTime"] = o.LastActiveDateTime.Get()
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.StartedDateTime != nil {
		toSerialize["startedDateTime"] = o.StartedDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserTimezone.IsSet() {
		toSerialize["userTimezone"] = o.UserTimezone.Get()
	}
	if o.Activity != nil {
		toSerialize["activity"] = o.Activity
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphActivityHistoryItem struct {
	value *MicrosoftGraphActivityHistoryItem
	isSet bool
}

func (v NullableMicrosoftGraphActivityHistoryItem) Get() *MicrosoftGraphActivityHistoryItem {
	return v.value
}

func (v *NullableMicrosoftGraphActivityHistoryItem) Set(val *MicrosoftGraphActivityHistoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphActivityHistoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphActivityHistoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphActivityHistoryItem(val *MicrosoftGraphActivityHistoryItem) *NullableMicrosoftGraphActivityHistoryItem {
	return &NullableMicrosoftGraphActivityHistoryItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphActivityHistoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphActivityHistoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


