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

// MicrosoftGraphPlannerGroup struct for MicrosoftGraphPlannerGroup
type MicrosoftGraphPlannerGroup struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Read-only. Nullable. Returns the plannerPlans owned by the group.
	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`
}

// NewMicrosoftGraphPlannerGroup instantiates a new MicrosoftGraphPlannerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPlannerGroup() *MicrosoftGraphPlannerGroup {
	this := MicrosoftGraphPlannerGroup{}
	return &this
}

// NewMicrosoftGraphPlannerGroupWithDefaults instantiates a new MicrosoftGraphPlannerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPlannerGroupWithDefaults() *MicrosoftGraphPlannerGroup {
	this := MicrosoftGraphPlannerGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerGroup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPlannerGroup) SetId(v string) {
	o.Id = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerGroup) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerGroup) GetPlansOk() (*[]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerGroup) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *MicrosoftGraphPlannerGroup) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}

func (o MicrosoftGraphPlannerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPlannerGroup struct {
	value *MicrosoftGraphPlannerGroup
	isSet bool
}

func (v NullableMicrosoftGraphPlannerGroup) Get() *MicrosoftGraphPlannerGroup {
	return v.value
}

func (v *NullableMicrosoftGraphPlannerGroup) Set(val *MicrosoftGraphPlannerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPlannerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPlannerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPlannerGroup(val *MicrosoftGraphPlannerGroup) *NullableMicrosoftGraphPlannerGroup {
	return &NullableMicrosoftGraphPlannerGroup{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPlannerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPlannerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


