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

// MicrosoftGraphPrintJobStatus struct for MicrosoftGraphPrintJobStatus
type MicrosoftGraphPrintJobStatus struct {
	// A human-readable description of the print job's current processing state. Read-only.
	Description *string `json:"description,omitempty"`
	// Additional details for print job state. Valid values are described in the following table. Read-only.
	Details *[]AnyOfmicrosoftGraphPrintJobStateDetail `json:"details,omitempty"`
	// True if the job was acknowledged by a printer; false otherwise. Read-only.
	IsAcquiredByPrinter *bool `json:"isAcquiredByPrinter,omitempty"`
	// The print job's current processing state. Valid values are described in the following table. Read-only.
	State AnyOfmicrosoftGraphPrintJobProcessingState `json:"state,omitempty"`
}

// NewMicrosoftGraphPrintJobStatus instantiates a new MicrosoftGraphPrintJobStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintJobStatus() *MicrosoftGraphPrintJobStatus {
	this := MicrosoftGraphPrintJobStatus{}
	return &this
}

// NewMicrosoftGraphPrintJobStatusWithDefaults instantiates a new MicrosoftGraphPrintJobStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintJobStatusWithDefaults() *MicrosoftGraphPrintJobStatus {
	this := MicrosoftGraphPrintJobStatus{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintJobStatus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintJobStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobStatus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphPrintJobStatus) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintJobStatus) GetDetails() []AnyOfmicrosoftGraphPrintJobStateDetail {
	if o == nil || o.Details == nil {
		var ret []AnyOfmicrosoftGraphPrintJobStateDetail
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintJobStatus) GetDetailsOk() (*[]AnyOfmicrosoftGraphPrintJobStateDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobStatus) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []AnyOfmicrosoftGraphPrintJobStateDetail and assigns it to the Details field.
func (o *MicrosoftGraphPrintJobStatus) SetDetails(v []AnyOfmicrosoftGraphPrintJobStateDetail) {
	o.Details = &v
}

// GetIsAcquiredByPrinter returns the IsAcquiredByPrinter field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintJobStatus) GetIsAcquiredByPrinter() bool {
	if o == nil || o.IsAcquiredByPrinter == nil {
		var ret bool
		return ret
	}
	return *o.IsAcquiredByPrinter
}

// GetIsAcquiredByPrinterOk returns a tuple with the IsAcquiredByPrinter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintJobStatus) GetIsAcquiredByPrinterOk() (*bool, bool) {
	if o == nil || o.IsAcquiredByPrinter == nil {
		return nil, false
	}
	return o.IsAcquiredByPrinter, true
}

// HasIsAcquiredByPrinter returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobStatus) HasIsAcquiredByPrinter() bool {
	if o != nil && o.IsAcquiredByPrinter != nil {
		return true
	}

	return false
}

// SetIsAcquiredByPrinter gets a reference to the given bool and assigns it to the IsAcquiredByPrinter field.
func (o *MicrosoftGraphPrintJobStatus) SetIsAcquiredByPrinter(v bool) {
	o.IsAcquiredByPrinter = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintJobStatus) GetState() AnyOfmicrosoftGraphPrintJobProcessingState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintJobProcessingState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintJobStatus) GetStateOk() (*AnyOfmicrosoftGraphPrintJobProcessingState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintJobStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphPrintJobProcessingState and assigns it to the State field.
func (o *MicrosoftGraphPrintJobStatus) SetState(v AnyOfmicrosoftGraphPrintJobProcessingState) {
	o.State = v
}

func (o MicrosoftGraphPrintJobStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.IsAcquiredByPrinter != nil {
		toSerialize["isAcquiredByPrinter"] = o.IsAcquiredByPrinter
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintJobStatus struct {
	value *MicrosoftGraphPrintJobStatus
	isSet bool
}

func (v NullableMicrosoftGraphPrintJobStatus) Get() *MicrosoftGraphPrintJobStatus {
	return v.value
}

func (v *NullableMicrosoftGraphPrintJobStatus) Set(val *MicrosoftGraphPrintJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintJobStatus(val *MicrosoftGraphPrintJobStatus) *NullableMicrosoftGraphPrintJobStatus {
	return &NullableMicrosoftGraphPrintJobStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


