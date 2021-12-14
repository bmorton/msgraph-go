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

// MicrosoftGraphDomainState struct for MicrosoftGraphDomainState
type MicrosoftGraphDomainState struct {
	// Timestamp for when the last activity occurred. The value is updated when an operation is scheduled, the asynchronous task starts, and when the operation completes.
	LastActionDateTime NullableTime `json:"lastActionDateTime,omitempty"`
	// Type of asynchronous operation. The values can be ForceDelete or Verification
	Operation NullableString `json:"operation,omitempty"`
	// Current status of the operation.  Scheduled - Operation has been scheduled but has not started.  InProgress - Task has started and is in progress.  Failed - Operation has failed.
	Status NullableString `json:"status,omitempty"`
}

// NewMicrosoftGraphDomainState instantiates a new MicrosoftGraphDomainState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDomainState() *MicrosoftGraphDomainState {
	this := MicrosoftGraphDomainState{}
	return &this
}

// NewMicrosoftGraphDomainStateWithDefaults instantiates a new MicrosoftGraphDomainState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDomainStateWithDefaults() *MicrosoftGraphDomainState {
	this := MicrosoftGraphDomainState{}
	return &this
}

// GetLastActionDateTime returns the LastActionDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomainState) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime.Get()
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomainState) GetLastActionDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActionDateTime.Get(), o.LastActionDateTime.IsSet()
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainState) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given NullableTime and assigns it to the LastActionDateTime field.
func (o *MicrosoftGraphDomainState) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime.Set(&v)
}
// SetLastActionDateTimeNil sets the value for LastActionDateTime to be an explicit nil
func (o *MicrosoftGraphDomainState) SetLastActionDateTimeNil() {
	o.LastActionDateTime.Set(nil)
}

// UnsetLastActionDateTime ensures that no value is present for LastActionDateTime, not even an explicit nil
func (o *MicrosoftGraphDomainState) UnsetLastActionDateTime() {
	o.LastActionDateTime.Unset()
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomainState) GetOperation() string {
	if o == nil || o.Operation.Get() == nil {
		var ret string
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomainState) GetOperationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainState) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableString and assigns it to the Operation field.
func (o *MicrosoftGraphDomainState) SetOperation(v string) {
	o.Operation.Set(&v)
}
// SetOperationNil sets the value for Operation to be an explicit nil
func (o *MicrosoftGraphDomainState) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *MicrosoftGraphDomainState) UnsetOperation() {
	o.Operation.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDomainState) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDomainState) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainState) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *MicrosoftGraphDomainState) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *MicrosoftGraphDomainState) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *MicrosoftGraphDomainState) UnsetStatus() {
	o.Status.Unset()
}

func (o MicrosoftGraphDomainState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastActionDateTime.IsSet() {
		toSerialize["lastActionDateTime"] = o.LastActionDateTime.Get()
	}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDomainState struct {
	value *MicrosoftGraphDomainState
	isSet bool
}

func (v NullableMicrosoftGraphDomainState) Get() *MicrosoftGraphDomainState {
	return v.value
}

func (v *NullableMicrosoftGraphDomainState) Set(val *MicrosoftGraphDomainState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDomainState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDomainState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDomainState(val *MicrosoftGraphDomainState) *NullableMicrosoftGraphDomainState {
	return &NullableMicrosoftGraphDomainState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDomainState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDomainState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


