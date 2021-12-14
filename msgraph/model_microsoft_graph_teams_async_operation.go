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

// MicrosoftGraphTeamsAsyncOperation struct for MicrosoftGraphTeamsAsyncOperation
type MicrosoftGraphTeamsAsyncOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Number of times the operation was attempted before being marked successful or failed.
	AttemptsCount *int32 `json:"attemptsCount,omitempty"`
	// Time when the operation was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Any error that causes the async operation to fail.
	Error AnyOfmicrosoftGraphOperationError `json:"error,omitempty"`
	// Time when the async operation was last updated.
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	// Denotes which type of operation is being described.
	OperationType AnyOfmicrosoftGraphTeamsAsyncOperationType `json:"operationType,omitempty"`
	// Operation status.
	Status AnyOfmicrosoftGraphTeamsAsyncOperationStatus `json:"status,omitempty"`
	// The ID of the object that's created or modified as result of this async operation, typically a team.
	TargetResourceId NullableString `json:"targetResourceId,omitempty"`
	// The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
	TargetResourceLocation NullableString `json:"targetResourceLocation,omitempty"`
}

// NewMicrosoftGraphTeamsAsyncOperation instantiates a new MicrosoftGraphTeamsAsyncOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTeamsAsyncOperation() *MicrosoftGraphTeamsAsyncOperation {
	this := MicrosoftGraphTeamsAsyncOperation{}
	return &this
}

// NewMicrosoftGraphTeamsAsyncOperationWithDefaults instantiates a new MicrosoftGraphTeamsAsyncOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTeamsAsyncOperationWithDefaults() *MicrosoftGraphTeamsAsyncOperation {
	this := MicrosoftGraphTeamsAsyncOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsAsyncOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetId(v string) {
	o.Id = &v
}

// GetAttemptsCount returns the AttemptsCount field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsAsyncOperation) GetAttemptsCount() int32 {
	if o == nil || o.AttemptsCount == nil {
		var ret int32
		return ret
	}
	return *o.AttemptsCount
}

// GetAttemptsCountOk returns a tuple with the AttemptsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) GetAttemptsCountOk() (*int32, bool) {
	if o == nil || o.AttemptsCount == nil {
		return nil, false
	}
	return o.AttemptsCount, true
}

// HasAttemptsCount returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasAttemptsCount() bool {
	if o != nil && o.AttemptsCount != nil {
		return true
	}

	return false
}

// SetAttemptsCount gets a reference to the given int32 and assigns it to the AttemptsCount field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetAttemptsCount(v int32) {
	o.AttemptsCount = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsAsyncOperation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAsyncOperation) GetError() AnyOfmicrosoftGraphOperationError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOperationError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAsyncOperation) GetErrorOk() (*AnyOfmicrosoftGraphOperationError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphOperationError and assigns it to the Error field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetError(v AnyOfmicrosoftGraphOperationError) {
	o.Error = v
}

// GetLastActionDateTime returns the LastActionDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsAsyncOperation) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) GetLastActionDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastActionDateTime == nil {
		return nil, false
	}
	return o.LastActionDateTime, true
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime != nil {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAsyncOperation) GetOperationType() AnyOfmicrosoftGraphTeamsAsyncOperationType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamsAsyncOperationType
		return ret
	}
	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAsyncOperation) GetOperationTypeOk() (*AnyOfmicrosoftGraphTeamsAsyncOperationType, bool) {
	if o == nil || o.OperationType == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasOperationType() bool {
	if o != nil && o.OperationType != nil {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationType and assigns it to the OperationType field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetOperationType(v AnyOfmicrosoftGraphTeamsAsyncOperationType) {
	o.OperationType = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAsyncOperation) GetStatus() AnyOfmicrosoftGraphTeamsAsyncOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamsAsyncOperationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAsyncOperation) GetStatusOk() (*AnyOfmicrosoftGraphTeamsAsyncOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetStatus(v AnyOfmicrosoftGraphTeamsAsyncOperationStatus) {
	o.Status = v
}

// GetTargetResourceId returns the TargetResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceId() string {
	if o == nil || o.TargetResourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetResourceId.Get()
}

// GetTargetResourceIdOk returns a tuple with the TargetResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetResourceId.Get(), o.TargetResourceId.IsSet()
}

// HasTargetResourceId returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasTargetResourceId() bool {
	if o != nil && o.TargetResourceId.IsSet() {
		return true
	}

	return false
}

// SetTargetResourceId gets a reference to the given NullableString and assigns it to the TargetResourceId field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceId(v string) {
	o.TargetResourceId.Set(&v)
}
// SetTargetResourceIdNil sets the value for TargetResourceId to be an explicit nil
func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceIdNil() {
	o.TargetResourceId.Set(nil)
}

// UnsetTargetResourceId ensures that no value is present for TargetResourceId, not even an explicit nil
func (o *MicrosoftGraphTeamsAsyncOperation) UnsetTargetResourceId() {
	o.TargetResourceId.Unset()
}

// GetTargetResourceLocation returns the TargetResourceLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceLocation() string {
	if o == nil || o.TargetResourceLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetResourceLocation.Get()
}

// GetTargetResourceLocationOk returns a tuple with the TargetResourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetResourceLocation.Get(), o.TargetResourceLocation.IsSet()
}

// HasTargetResourceLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAsyncOperation) HasTargetResourceLocation() bool {
	if o != nil && o.TargetResourceLocation.IsSet() {
		return true
	}

	return false
}

// SetTargetResourceLocation gets a reference to the given NullableString and assigns it to the TargetResourceLocation field.
func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceLocation(v string) {
	o.TargetResourceLocation.Set(&v)
}
// SetTargetResourceLocationNil sets the value for TargetResourceLocation to be an explicit nil
func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceLocationNil() {
	o.TargetResourceLocation.Set(nil)
}

// UnsetTargetResourceLocation ensures that no value is present for TargetResourceLocation, not even an explicit nil
func (o *MicrosoftGraphTeamsAsyncOperation) UnsetTargetResourceLocation() {
	o.TargetResourceLocation.Unset()
}

func (o MicrosoftGraphTeamsAsyncOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AttemptsCount != nil {
		toSerialize["attemptsCount"] = o.AttemptsCount
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.LastActionDateTime != nil {
		toSerialize["lastActionDateTime"] = o.LastActionDateTime
	}
	if o.OperationType != nil {
		toSerialize["operationType"] = o.OperationType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TargetResourceId.IsSet() {
		toSerialize["targetResourceId"] = o.TargetResourceId.Get()
	}
	if o.TargetResourceLocation.IsSet() {
		toSerialize["targetResourceLocation"] = o.TargetResourceLocation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTeamsAsyncOperation struct {
	value *MicrosoftGraphTeamsAsyncOperation
	isSet bool
}

func (v NullableMicrosoftGraphTeamsAsyncOperation) Get() *MicrosoftGraphTeamsAsyncOperation {
	return v.value
}

func (v *NullableMicrosoftGraphTeamsAsyncOperation) Set(val *MicrosoftGraphTeamsAsyncOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamsAsyncOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamsAsyncOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamsAsyncOperation(val *MicrosoftGraphTeamsAsyncOperation) *NullableMicrosoftGraphTeamsAsyncOperation {
	return &NullableMicrosoftGraphTeamsAsyncOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamsAsyncOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamsAsyncOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

