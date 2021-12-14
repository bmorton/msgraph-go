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

// MicrosoftGraphPlannerUser struct for MicrosoftGraphPlannerUser
type MicrosoftGraphPlannerUser struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Read-only. Nullable. Returns the plannerTasks assigned to the user.
	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`
	// Read-only. Nullable. Returns the plannerPlans shared with the user.
	Tasks *[]MicrosoftGraphPlannerTask `json:"tasks,omitempty"`
}

// NewMicrosoftGraphPlannerUser instantiates a new MicrosoftGraphPlannerUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPlannerUser() *MicrosoftGraphPlannerUser {
	this := MicrosoftGraphPlannerUser{}
	return &this
}

// NewMicrosoftGraphPlannerUserWithDefaults instantiates a new MicrosoftGraphPlannerUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPlannerUserWithDefaults() *MicrosoftGraphPlannerUser {
	this := MicrosoftGraphPlannerUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPlannerUser) SetId(v string) {
	o.Id = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerUser) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerUser) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerUser) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *MicrosoftGraphPlannerUser) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerUser) GetTasks() []MicrosoftGraphPlannerTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPlannerTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerUser) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerUser) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.
func (o *MicrosoftGraphPlannerUser) SetTasks(v []MicrosoftGraphPlannerTask) {
	o.Tasks = &v
}

func (o MicrosoftGraphPlannerUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPlannerUser struct {
	value *MicrosoftGraphPlannerUser
	isSet bool
}

func (v NullableMicrosoftGraphPlannerUser) Get() *MicrosoftGraphPlannerUser {
	return v.value
}

func (v *NullableMicrosoftGraphPlannerUser) Set(val *MicrosoftGraphPlannerUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPlannerUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPlannerUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPlannerUser(val *MicrosoftGraphPlannerUser) *NullableMicrosoftGraphPlannerUser {
	return &NullableMicrosoftGraphPlannerUser{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPlannerUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPlannerUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


