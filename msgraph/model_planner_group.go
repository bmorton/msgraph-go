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

// PlannerGroup struct for PlannerGroup
type PlannerGroup struct {
	// Read-only. Nullable. Returns the plannerPlans owned by the group.
	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`
}

// NewPlannerGroup instantiates a new PlannerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannerGroup() *PlannerGroup {
	this := PlannerGroup{}
	return &this
}

// NewPlannerGroupWithDefaults instantiates a new PlannerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannerGroupWithDefaults() *PlannerGroup {
	this := PlannerGroup{}
	return &this
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *PlannerGroup) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerGroup) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *PlannerGroup) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *PlannerGroup) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}

func (o PlannerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullablePlannerGroup struct {
	value *PlannerGroup
	isSet bool
}

func (v NullablePlannerGroup) Get() *PlannerGroup {
	return v.value
}

func (v *NullablePlannerGroup) Set(val *PlannerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannerGroup(val *PlannerGroup) *NullablePlannerGroup {
	return &NullablePlannerGroup{value: val, isSet: true}
}

func (v NullablePlannerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


