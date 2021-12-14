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

// PrintTask struct for PrintTask
type PrintTask struct {
	// The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only.
	ParentUrl *string `json:"parentUrl,omitempty"`
	Status *MicrosoftGraphPrintTaskStatus `json:"status,omitempty"`
	Definition *MicrosoftGraphPrintTaskDefinition `json:"definition,omitempty"`
	Trigger *MicrosoftGraphPrintTaskTrigger `json:"trigger,omitempty"`
}

// NewPrintTask instantiates a new PrintTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintTask() *PrintTask {
	this := PrintTask{}
	return &this
}

// NewPrintTaskWithDefaults instantiates a new PrintTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintTaskWithDefaults() *PrintTask {
	this := PrintTask{}
	return &this
}

// GetParentUrl returns the ParentUrl field value if set, zero value otherwise.
func (o *PrintTask) GetParentUrl() string {
	if o == nil || o.ParentUrl == nil {
		var ret string
		return ret
	}
	return *o.ParentUrl
}

// GetParentUrlOk returns a tuple with the ParentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTask) GetParentUrlOk() (*string, bool) {
	if o == nil || o.ParentUrl == nil {
		return nil, false
	}
	return o.ParentUrl, true
}

// HasParentUrl returns a boolean if a field has been set.
func (o *PrintTask) HasParentUrl() bool {
	if o != nil && o.ParentUrl != nil {
		return true
	}

	return false
}

// SetParentUrl gets a reference to the given string and assigns it to the ParentUrl field.
func (o *PrintTask) SetParentUrl(v string) {
	o.ParentUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrintTask) GetStatus() MicrosoftGraphPrintTaskStatus {
	if o == nil || o.Status == nil {
		var ret MicrosoftGraphPrintTaskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTask) GetStatusOk() (*MicrosoftGraphPrintTaskStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrintTask) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MicrosoftGraphPrintTaskStatus and assigns it to the Status field.
func (o *PrintTask) SetStatus(v MicrosoftGraphPrintTaskStatus) {
	o.Status = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *PrintTask) GetDefinition() MicrosoftGraphPrintTaskDefinition {
	if o == nil || o.Definition == nil {
		var ret MicrosoftGraphPrintTaskDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTask) GetDefinitionOk() (*MicrosoftGraphPrintTaskDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *PrintTask) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given MicrosoftGraphPrintTaskDefinition and assigns it to the Definition field.
func (o *PrintTask) SetDefinition(v MicrosoftGraphPrintTaskDefinition) {
	o.Definition = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *PrintTask) GetTrigger() MicrosoftGraphPrintTaskTrigger {
	if o == nil || o.Trigger == nil {
		var ret MicrosoftGraphPrintTaskTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTask) GetTriggerOk() (*MicrosoftGraphPrintTaskTrigger, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *PrintTask) HasTrigger() bool {
	if o != nil && o.Trigger != nil {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given MicrosoftGraphPrintTaskTrigger and assigns it to the Trigger field.
func (o *PrintTask) SetTrigger(v MicrosoftGraphPrintTaskTrigger) {
	o.Trigger = &v
}

func (o PrintTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentUrl != nil {
		toSerialize["parentUrl"] = o.ParentUrl
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}
	return json.Marshal(toSerialize)
}

type NullablePrintTask struct {
	value *PrintTask
	isSet bool
}

func (v NullablePrintTask) Get() *PrintTask {
	return v.value
}

func (v *NullablePrintTask) Set(val *PrintTask) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintTask) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintTask(val *PrintTask) *NullablePrintTask {
	return &NullablePrintTask{value: val, isSet: true}
}

func (v NullablePrintTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

