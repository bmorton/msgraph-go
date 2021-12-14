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

// MicrosoftGraphPlanner struct for MicrosoftGraphPlanner
type MicrosoftGraphPlanner struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Read-only. Nullable. Returns a collection of the specified buckets
	Buckets *[]MicrosoftGraphPlannerBucket `json:"buckets,omitempty"`
	// Read-only. Nullable. Returns a collection of the specified plans
	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`
	// Read-only. Nullable. Returns a collection of the specified tasks
	Tasks *[]MicrosoftGraphPlannerTask `json:"tasks,omitempty"`
}

// NewMicrosoftGraphPlanner instantiates a new MicrosoftGraphPlanner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPlanner() *MicrosoftGraphPlanner {
	this := MicrosoftGraphPlanner{}
	return &this
}

// NewMicrosoftGraphPlannerWithDefaults instantiates a new MicrosoftGraphPlanner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPlannerWithDefaults() *MicrosoftGraphPlanner {
	this := MicrosoftGraphPlanner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPlanner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlanner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPlanner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPlanner) SetId(v string) {
	o.Id = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *MicrosoftGraphPlanner) GetBuckets() []MicrosoftGraphPlannerBucket {
	if o == nil || o.Buckets == nil {
		var ret []MicrosoftGraphPlannerBucket
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlanner) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *MicrosoftGraphPlanner) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []MicrosoftGraphPlannerBucket and assigns it to the Buckets field.
func (o *MicrosoftGraphPlanner) SetBuckets(v []MicrosoftGraphPlannerBucket) {
	o.Buckets = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *MicrosoftGraphPlanner) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlanner) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *MicrosoftGraphPlanner) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *MicrosoftGraphPlanner) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *MicrosoftGraphPlanner) GetTasks() []MicrosoftGraphPlannerTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPlannerTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlanner) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *MicrosoftGraphPlanner) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.
func (o *MicrosoftGraphPlanner) SetTasks(v []MicrosoftGraphPlannerTask) {
	o.Tasks = &v
}

func (o MicrosoftGraphPlanner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPlanner struct {
	value *MicrosoftGraphPlanner
	isSet bool
}

func (v NullableMicrosoftGraphPlanner) Get() *MicrosoftGraphPlanner {
	return v.value
}

func (v *NullableMicrosoftGraphPlanner) Set(val *MicrosoftGraphPlanner) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPlanner) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPlanner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPlanner(val *MicrosoftGraphPlanner) *NullableMicrosoftGraphPlanner {
	return &NullableMicrosoftGraphPlanner{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPlanner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPlanner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


