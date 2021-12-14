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

// ConnectionOperation struct for ConnectionOperation
type ConnectionOperation struct {
	// If status is failed, provides more information about the error that caused the failure.
	Error AnyOfmicrosoftGraphPublicError `json:"error,omitempty"`
	// Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue.
	Status AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus `json:"status,omitempty"`
}

// NewConnectionOperation instantiates a new ConnectionOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOperation() *ConnectionOperation {
	this := ConnectionOperation{}
	return &this
}

// NewConnectionOperationWithDefaults instantiates a new ConnectionOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOperationWithDefaults() *ConnectionOperation {
	this := ConnectionOperation{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOperation) GetError() AnyOfmicrosoftGraphPublicError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPublicError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOperation) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConnectionOperation) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphPublicError and assigns it to the Error field.
func (o *ConnectionOperation) SetError(v AnyOfmicrosoftGraphPublicError) {
	o.Error = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOperation) GetStatus() AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOperation) GetStatusOk() (*AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectionOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus and assigns it to the Status field.
func (o *ConnectionOperation) SetStatus(v AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus) {
	o.Status = v
}

func (o ConnectionOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOperation struct {
	value *ConnectionOperation
	isSet bool
}

func (v NullableConnectionOperation) Get() *ConnectionOperation {
	return v.value
}

func (v *NullableConnectionOperation) Set(val *ConnectionOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOperation(val *ConnectionOperation) *NullableConnectionOperation {
	return &NullableConnectionOperation{value: val, isSet: true}
}

func (v NullableConnectionOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

