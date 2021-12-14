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

// Operation struct for Operation
type Operation struct {
	// The start time of the operation.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The time of the last action of the operation.
	LastActionDateTime NullableTime `json:"lastActionDateTime,omitempty"`
	// The current status of the operation: notStarted, running, completed, failed
	Status AnyOfmicrosoftGraphOperationStatus `json:"status,omitempty"`
}

// NewOperation instantiates a new Operation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperation() *Operation {
	this := Operation{}
	return &this
}

// NewOperationWithDefaults instantiates a new Operation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationWithDefaults() *Operation {
	this := Operation{}
	return &this
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Operation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Operation) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Operation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *Operation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *Operation) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *Operation) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastActionDateTime returns the LastActionDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Operation) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime.Get()
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Operation) GetLastActionDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActionDateTime.Get(), o.LastActionDateTime.IsSet()
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *Operation) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given NullableTime and assigns it to the LastActionDateTime field.
func (o *Operation) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime.Set(&v)
}
// SetLastActionDateTimeNil sets the value for LastActionDateTime to be an explicit nil
func (o *Operation) SetLastActionDateTimeNil() {
	o.LastActionDateTime.Set(nil)
}

// UnsetLastActionDateTime ensures that no value is present for LastActionDateTime, not even an explicit nil
func (o *Operation) UnsetLastActionDateTime() {
	o.LastActionDateTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Operation) GetStatus() AnyOfmicrosoftGraphOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Operation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Operation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.
func (o *Operation) SetStatus(v AnyOfmicrosoftGraphOperationStatus) {
	o.Status = v
}

func (o Operation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.LastActionDateTime.IsSet() {
		toSerialize["lastActionDateTime"] = o.LastActionDateTime.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableOperation struct {
	value *Operation
	isSet bool
}

func (v NullableOperation) Get() *Operation {
	return v.value
}

func (v *NullableOperation) Set(val *Operation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation(val *Operation) *NullableOperation {
	return &NullableOperation{value: val, isSet: true}
}

func (v NullableOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


