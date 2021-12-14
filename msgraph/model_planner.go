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

// Planner struct for Planner
type Planner struct {
	// Read-only. Nullable. Returns a collection of the specified buckets
	Buckets *[]MicrosoftGraphPlannerBucket `json:"buckets,omitempty"`
	// Read-only. Nullable. Returns a collection of the specified plans
	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`
	// Read-only. Nullable. Returns a collection of the specified tasks
	Tasks *[]MicrosoftGraphPlannerTask `json:"tasks,omitempty"`
}

// NewPlanner instantiates a new Planner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanner() *Planner {
	this := Planner{}
	return &this
}

// NewPlannerWithDefaults instantiates a new Planner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannerWithDefaults() *Planner {
	this := Planner{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *Planner) GetBuckets() []MicrosoftGraphPlannerBucket {
	if o == nil || o.Buckets == nil {
		var ret []MicrosoftGraphPlannerBucket
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Planner) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *Planner) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []MicrosoftGraphPlannerBucket and assigns it to the Buckets field.
func (o *Planner) SetBuckets(v []MicrosoftGraphPlannerBucket) {
	o.Buckets = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *Planner) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Planner) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *Planner) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *Planner) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *Planner) GetTasks() []MicrosoftGraphPlannerTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPlannerTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Planner) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *Planner) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.
func (o *Planner) SetTasks(v []MicrosoftGraphPlannerTask) {
	o.Tasks = &v
}

func (o Planner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullablePlanner struct {
	value *Planner
	isSet bool
}

func (v NullablePlanner) Get() *Planner {
	return v.value
}

func (v *NullablePlanner) Set(val *Planner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanner(val *Planner) *NullablePlanner {
	return &NullablePlanner{value: val, isSet: true}
}

func (v NullablePlanner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


