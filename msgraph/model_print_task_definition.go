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

// PrintTaskDefinition struct for PrintTaskDefinition
type PrintTaskDefinition struct {
	CreatedBy *MicrosoftGraphAppIdentity `json:"createdBy,omitempty"`
	// The name of the printTaskDefinition.
	DisplayName *string `json:"displayName,omitempty"`
	// A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
	Tasks *[]MicrosoftGraphPrintTask `json:"tasks,omitempty"`
}

// NewPrintTaskDefinition instantiates a new PrintTaskDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintTaskDefinition() *PrintTaskDefinition {
	this := PrintTaskDefinition{}
	return &this
}

// NewPrintTaskDefinitionWithDefaults instantiates a new PrintTaskDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintTaskDefinitionWithDefaults() *PrintTaskDefinition {
	this := PrintTaskDefinition{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PrintTaskDefinition) GetCreatedBy() MicrosoftGraphAppIdentity {
	if o == nil || o.CreatedBy == nil {
		var ret MicrosoftGraphAppIdentity
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTaskDefinition) GetCreatedByOk() (*MicrosoftGraphAppIdentity, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PrintTaskDefinition) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given MicrosoftGraphAppIdentity and assigns it to the CreatedBy field.
func (o *PrintTaskDefinition) SetCreatedBy(v MicrosoftGraphAppIdentity) {
	o.CreatedBy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PrintTaskDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTaskDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PrintTaskDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PrintTaskDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *PrintTaskDefinition) GetTasks() []MicrosoftGraphPrintTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPrintTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintTaskDefinition) GetTasksOk() (*[]MicrosoftGraphPrintTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *PrintTaskDefinition) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphPrintTask and assigns it to the Tasks field.
func (o *PrintTaskDefinition) SetTasks(v []MicrosoftGraphPrintTask) {
	o.Tasks = &v
}

func (o PrintTaskDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullablePrintTaskDefinition struct {
	value *PrintTaskDefinition
	isSet bool
}

func (v NullablePrintTaskDefinition) Get() *PrintTaskDefinition {
	return v.value
}

func (v *NullablePrintTaskDefinition) Set(val *PrintTaskDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintTaskDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintTaskDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintTaskDefinition(val *PrintTaskDefinition) *NullablePrintTaskDefinition {
	return &NullablePrintTaskDefinition{value: val, isSet: true}
}

func (v NullablePrintTaskDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintTaskDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

