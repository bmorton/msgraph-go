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

// Approval struct for Approval
type Approval struct {
	// A collection of stages in the approval decision.
	Stages *[]MicrosoftGraphApprovalStage `json:"stages,omitempty"`
}

// NewApproval instantiates a new Approval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproval() *Approval {
	this := Approval{}
	return &this
}

// NewApprovalWithDefaults instantiates a new Approval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalWithDefaults() *Approval {
	this := Approval{}
	return &this
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *Approval) GetStages() []MicrosoftGraphApprovalStage {
	if o == nil || o.Stages == nil {
		var ret []MicrosoftGraphApprovalStage
		return ret
	}
	return *o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetStagesOk() (*[]MicrosoftGraphApprovalStage, bool) {
	if o == nil || o.Stages == nil {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *Approval) HasStages() bool {
	if o != nil && o.Stages != nil {
		return true
	}

	return false
}

// SetStages gets a reference to the given []MicrosoftGraphApprovalStage and assigns it to the Stages field.
func (o *Approval) SetStages(v []MicrosoftGraphApprovalStage) {
	o.Stages = &v
}

func (o Approval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Stages != nil {
		toSerialize["stages"] = o.Stages
	}
	return json.Marshal(toSerialize)
}

type NullableApproval struct {
	value *Approval
	isSet bool
}

func (v NullableApproval) Get() *Approval {
	return v.value
}

func (v *NullableApproval) Set(val *Approval) {
	v.value = val
	v.isSet = true
}

func (v NullableApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproval(val *Approval) *NullableApproval {
	return &NullableApproval{value: val, isSet: true}
}

func (v NullableApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


