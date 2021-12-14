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

// MicrosoftGraphScheduleChangeRequest struct for MicrosoftGraphScheduleChangeRequest
type MicrosoftGraphScheduleChangeRequest struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Identity of the person who last modified the entity.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	AssignedTo AnyOfmicrosoftGraphScheduleChangeRequestActor `json:"assignedTo,omitempty"`
	ManagerActionDateTime NullableTime `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage NullableString `json:"managerActionMessage,omitempty"`
	ManagerUserId NullableString `json:"managerUserId,omitempty"`
	SenderDateTime NullableTime `json:"senderDateTime,omitempty"`
	SenderMessage NullableString `json:"senderMessage,omitempty"`
	SenderUserId NullableString `json:"senderUserId,omitempty"`
	State AnyOfmicrosoftGraphScheduleChangeState `json:"state,omitempty"`
}

// NewMicrosoftGraphScheduleChangeRequest instantiates a new MicrosoftGraphScheduleChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphScheduleChangeRequest() *MicrosoftGraphScheduleChangeRequest {
	this := MicrosoftGraphScheduleChangeRequest{}
	return &this
}

// NewMicrosoftGraphScheduleChangeRequestWithDefaults instantiates a new MicrosoftGraphScheduleChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphScheduleChangeRequestWithDefaults() *MicrosoftGraphScheduleChangeRequest {
	this := MicrosoftGraphScheduleChangeRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphScheduleChangeRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphScheduleChangeRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphScheduleChangeRequest) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphScheduleChangeRequest) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetAssignedTo() AnyOfmicrosoftGraphScheduleChangeRequestActor {
	if o == nil  {
		var ret AnyOfmicrosoftGraphScheduleChangeRequestActor
		return ret
	}
	return o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphScheduleChangeRequestActor, bool) {
	if o == nil || o.AssignedTo == nil {
		return nil, false
	}
	return &o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasAssignedTo() bool {
	if o != nil && o.AssignedTo != nil {
		return true
	}

	return false
}

// SetAssignedTo gets a reference to the given AnyOfmicrosoftGraphScheduleChangeRequestActor and assigns it to the AssignedTo field.
func (o *MicrosoftGraphScheduleChangeRequest) SetAssignedTo(v AnyOfmicrosoftGraphScheduleChangeRequestActor) {
	o.AssignedTo = v
}

// GetManagerActionDateTime returns the ManagerActionDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionDateTime() time.Time {
	if o == nil || o.ManagerActionDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ManagerActionDateTime.Get()
}

// GetManagerActionDateTimeOk returns a tuple with the ManagerActionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ManagerActionDateTime.Get(), o.ManagerActionDateTime.IsSet()
}

// HasManagerActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasManagerActionDateTime() bool {
	if o != nil && o.ManagerActionDateTime.IsSet() {
		return true
	}

	return false
}

// SetManagerActionDateTime gets a reference to the given NullableTime and assigns it to the ManagerActionDateTime field.
func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionDateTime(v time.Time) {
	o.ManagerActionDateTime.Set(&v)
}
// SetManagerActionDateTimeNil sets the value for ManagerActionDateTime to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionDateTimeNil() {
	o.ManagerActionDateTime.Set(nil)
}

// UnsetManagerActionDateTime ensures that no value is present for ManagerActionDateTime, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetManagerActionDateTime() {
	o.ManagerActionDateTime.Unset()
}

// GetManagerActionMessage returns the ManagerActionMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionMessage() string {
	if o == nil || o.ManagerActionMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ManagerActionMessage.Get()
}

// GetManagerActionMessageOk returns a tuple with the ManagerActionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetManagerActionMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ManagerActionMessage.Get(), o.ManagerActionMessage.IsSet()
}

// HasManagerActionMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasManagerActionMessage() bool {
	if o != nil && o.ManagerActionMessage.IsSet() {
		return true
	}

	return false
}

// SetManagerActionMessage gets a reference to the given NullableString and assigns it to the ManagerActionMessage field.
func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionMessage(v string) {
	o.ManagerActionMessage.Set(&v)
}
// SetManagerActionMessageNil sets the value for ManagerActionMessage to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetManagerActionMessageNil() {
	o.ManagerActionMessage.Set(nil)
}

// UnsetManagerActionMessage ensures that no value is present for ManagerActionMessage, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetManagerActionMessage() {
	o.ManagerActionMessage.Unset()
}

// GetManagerUserId returns the ManagerUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetManagerUserId() string {
	if o == nil || o.ManagerUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ManagerUserId.Get()
}

// GetManagerUserIdOk returns a tuple with the ManagerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetManagerUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ManagerUserId.Get(), o.ManagerUserId.IsSet()
}

// HasManagerUserId returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasManagerUserId() bool {
	if o != nil && o.ManagerUserId.IsSet() {
		return true
	}

	return false
}

// SetManagerUserId gets a reference to the given NullableString and assigns it to the ManagerUserId field.
func (o *MicrosoftGraphScheduleChangeRequest) SetManagerUserId(v string) {
	o.ManagerUserId.Set(&v)
}
// SetManagerUserIdNil sets the value for ManagerUserId to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetManagerUserIdNil() {
	o.ManagerUserId.Set(nil)
}

// UnsetManagerUserId ensures that no value is present for ManagerUserId, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetManagerUserId() {
	o.ManagerUserId.Unset()
}

// GetSenderDateTime returns the SenderDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetSenderDateTime() time.Time {
	if o == nil || o.SenderDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SenderDateTime.Get()
}

// GetSenderDateTimeOk returns a tuple with the SenderDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetSenderDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SenderDateTime.Get(), o.SenderDateTime.IsSet()
}

// HasSenderDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasSenderDateTime() bool {
	if o != nil && o.SenderDateTime.IsSet() {
		return true
	}

	return false
}

// SetSenderDateTime gets a reference to the given NullableTime and assigns it to the SenderDateTime field.
func (o *MicrosoftGraphScheduleChangeRequest) SetSenderDateTime(v time.Time) {
	o.SenderDateTime.Set(&v)
}
// SetSenderDateTimeNil sets the value for SenderDateTime to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetSenderDateTimeNil() {
	o.SenderDateTime.Set(nil)
}

// UnsetSenderDateTime ensures that no value is present for SenderDateTime, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetSenderDateTime() {
	o.SenderDateTime.Unset()
}

// GetSenderMessage returns the SenderMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetSenderMessage() string {
	if o == nil || o.SenderMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.SenderMessage.Get()
}

// GetSenderMessageOk returns a tuple with the SenderMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetSenderMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SenderMessage.Get(), o.SenderMessage.IsSet()
}

// HasSenderMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasSenderMessage() bool {
	if o != nil && o.SenderMessage.IsSet() {
		return true
	}

	return false
}

// SetSenderMessage gets a reference to the given NullableString and assigns it to the SenderMessage field.
func (o *MicrosoftGraphScheduleChangeRequest) SetSenderMessage(v string) {
	o.SenderMessage.Set(&v)
}
// SetSenderMessageNil sets the value for SenderMessage to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetSenderMessageNil() {
	o.SenderMessage.Set(nil)
}

// UnsetSenderMessage ensures that no value is present for SenderMessage, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetSenderMessage() {
	o.SenderMessage.Unset()
}

// GetSenderUserId returns the SenderUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetSenderUserId() string {
	if o == nil || o.SenderUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SenderUserId.Get()
}

// GetSenderUserIdOk returns a tuple with the SenderUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetSenderUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SenderUserId.Get(), o.SenderUserId.IsSet()
}

// HasSenderUserId returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasSenderUserId() bool {
	if o != nil && o.SenderUserId.IsSet() {
		return true
	}

	return false
}

// SetSenderUserId gets a reference to the given NullableString and assigns it to the SenderUserId field.
func (o *MicrosoftGraphScheduleChangeRequest) SetSenderUserId(v string) {
	o.SenderUserId.Set(&v)
}
// SetSenderUserIdNil sets the value for SenderUserId to be an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) SetSenderUserIdNil() {
	o.SenderUserId.Set(nil)
}

// UnsetSenderUserId ensures that no value is present for SenderUserId, not even an explicit nil
func (o *MicrosoftGraphScheduleChangeRequest) UnsetSenderUserId() {
	o.SenderUserId.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphScheduleChangeRequest) GetState() AnyOfmicrosoftGraphScheduleChangeState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphScheduleChangeState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphScheduleChangeRequest) GetStateOk() (*AnyOfmicrosoftGraphScheduleChangeState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphScheduleChangeRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphScheduleChangeState and assigns it to the State field.
func (o *MicrosoftGraphScheduleChangeRequest) SetState(v AnyOfmicrosoftGraphScheduleChangeState) {
	o.State = v
}

func (o MicrosoftGraphScheduleChangeRequest) MarshalJSON() ([]byte, error) {
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
	if o.AssignedTo != nil {
		toSerialize["assignedTo"] = o.AssignedTo
	}
	if o.ManagerActionDateTime.IsSet() {
		toSerialize["managerActionDateTime"] = o.ManagerActionDateTime.Get()
	}
	if o.ManagerActionMessage.IsSet() {
		toSerialize["managerActionMessage"] = o.ManagerActionMessage.Get()
	}
	if o.ManagerUserId.IsSet() {
		toSerialize["managerUserId"] = o.ManagerUserId.Get()
	}
	if o.SenderDateTime.IsSet() {
		toSerialize["senderDateTime"] = o.SenderDateTime.Get()
	}
	if o.SenderMessage.IsSet() {
		toSerialize["senderMessage"] = o.SenderMessage.Get()
	}
	if o.SenderUserId.IsSet() {
		toSerialize["senderUserId"] = o.SenderUserId.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphScheduleChangeRequest struct {
	value *MicrosoftGraphScheduleChangeRequest
	isSet bool
}

func (v NullableMicrosoftGraphScheduleChangeRequest) Get() *MicrosoftGraphScheduleChangeRequest {
	return v.value
}

func (v *NullableMicrosoftGraphScheduleChangeRequest) Set(val *MicrosoftGraphScheduleChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphScheduleChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphScheduleChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphScheduleChangeRequest(val *MicrosoftGraphScheduleChangeRequest) *NullableMicrosoftGraphScheduleChangeRequest {
	return &NullableMicrosoftGraphScheduleChangeRequest{value: val, isSet: true}
}

func (v NullableMicrosoftGraphScheduleChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphScheduleChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


