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

// MicrosoftGraphPrintOperation struct for MicrosoftGraphPrintOperation
type MicrosoftGraphPrintOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The DateTimeOffset when the operation was created. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	Status *MicrosoftGraphPrintOperationStatus `json:"status,omitempty"`
}

// NewMicrosoftGraphPrintOperation instantiates a new MicrosoftGraphPrintOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintOperation() *MicrosoftGraphPrintOperation {
	this := MicrosoftGraphPrintOperation{}
	return &this
}

// NewMicrosoftGraphPrintOperationWithDefaults instantiates a new MicrosoftGraphPrintOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintOperationWithDefaults() *MicrosoftGraphPrintOperation {
	this := MicrosoftGraphPrintOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintOperation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPrintOperation) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintOperation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintOperation) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintOperation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphPrintOperation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintOperation) GetStatus() MicrosoftGraphPrintOperationStatus {
	if o == nil || o.Status == nil {
		var ret MicrosoftGraphPrintOperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintOperation) GetStatusOk() (*MicrosoftGraphPrintOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MicrosoftGraphPrintOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphPrintOperation) SetStatus(v MicrosoftGraphPrintOperationStatus) {
	o.Status = &v
}

func (o MicrosoftGraphPrintOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintOperation struct {
	value *MicrosoftGraphPrintOperation
	isSet bool
}

func (v NullableMicrosoftGraphPrintOperation) Get() *MicrosoftGraphPrintOperation {
	return v.value
}

func (v *NullableMicrosoftGraphPrintOperation) Set(val *MicrosoftGraphPrintOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintOperation(val *MicrosoftGraphPrintOperation) *NullableMicrosoftGraphPrintOperation {
	return &NullableMicrosoftGraphPrintOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


