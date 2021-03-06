/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphUpdateRecordingStatusOperation struct for MicrosoftGraphUpdateRecordingStatusOperation
type MicrosoftGraphUpdateRecordingStatusOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Unique Client Context string. Max limit is 256 chars.
	ClientContext NullableString `json:"clientContext,omitempty"`
	// The result information. Read-only.
	ResultInfo AnyOfmicrosoftGraphResultInfo `json:"resultInfo,omitempty"`
	// Possible values are: notStarted, running, completed, failed. Read-only.
	Status AnyOfmicrosoftGraphOperationStatus `json:"status,omitempty"`
}

// NewMicrosoftGraphUpdateRecordingStatusOperation instantiates a new MicrosoftGraphUpdateRecordingStatusOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUpdateRecordingStatusOperation() *MicrosoftGraphUpdateRecordingStatusOperation {
	this := MicrosoftGraphUpdateRecordingStatusOperation{}
	return &this
}

// NewMicrosoftGraphUpdateRecordingStatusOperationWithDefaults instantiates a new MicrosoftGraphUpdateRecordingStatusOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUpdateRecordingStatusOperationWithDefaults() *MicrosoftGraphUpdateRecordingStatusOperation {
	this := MicrosoftGraphUpdateRecordingStatusOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetId(v string) {
	o.Id = &v
}

// GetClientContext returns the ClientContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetClientContext() string {
	if o == nil || o.ClientContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientContext.Get()
}

// GetClientContextOk returns a tuple with the ClientContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetClientContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientContext.Get(), o.ClientContext.IsSet()
}

// HasClientContext returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasClientContext() bool {
	if o != nil && o.ClientContext.IsSet() {
		return true
	}

	return false
}

// SetClientContext gets a reference to the given NullableString and assigns it to the ClientContext field.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetClientContext(v string) {
	o.ClientContext.Set(&v)
}
// SetClientContextNil sets the value for ClientContext to be an explicit nil
func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetClientContextNil() {
	o.ClientContext.Set(nil)
}

// UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
func (o *MicrosoftGraphUpdateRecordingStatusOperation) UnsetClientContext() {
	o.ClientContext.Unset()
}

// GetResultInfo returns the ResultInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetResultInfo() AnyOfmicrosoftGraphResultInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResultInfo
		return ret
	}
	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool) {
	if o == nil || o.ResultInfo == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// HasResultInfo returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasResultInfo() bool {
	if o != nil && o.ResultInfo != nil {
		return true
	}

	return false
}

// SetResultInfo gets a reference to the given AnyOfmicrosoftGraphResultInfo and assigns it to the ResultInfo field.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetResultInfo(v AnyOfmicrosoftGraphResultInfo) {
	o.ResultInfo = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus) {
	o.Status = v
}

func (o MicrosoftGraphUpdateRecordingStatusOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClientContext.IsSet() {
		toSerialize["clientContext"] = o.ClientContext.Get()
	}
	if o.ResultInfo != nil {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUpdateRecordingStatusOperation struct {
	value *MicrosoftGraphUpdateRecordingStatusOperation
	isSet bool
}

func (v NullableMicrosoftGraphUpdateRecordingStatusOperation) Get() *MicrosoftGraphUpdateRecordingStatusOperation {
	return v.value
}

func (v *NullableMicrosoftGraphUpdateRecordingStatusOperation) Set(val *MicrosoftGraphUpdateRecordingStatusOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUpdateRecordingStatusOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUpdateRecordingStatusOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUpdateRecordingStatusOperation(val *MicrosoftGraphUpdateRecordingStatusOperation) *NullableMicrosoftGraphUpdateRecordingStatusOperation {
	return &NullableMicrosoftGraphUpdateRecordingStatusOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUpdateRecordingStatusOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUpdateRecordingStatusOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


