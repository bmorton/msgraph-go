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

// TargetedManagedAppProtection Policy used to configure detailed management settings targeted to specific security groups
type TargetedManagedAppProtection struct {
	// Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
	Assignments *[]MicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// NewTargetedManagedAppProtection instantiates a new TargetedManagedAppProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetedManagedAppProtection() *TargetedManagedAppProtection {
	this := TargetedManagedAppProtection{}
	return &this
}

// NewTargetedManagedAppProtectionWithDefaults instantiates a new TargetedManagedAppProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetedManagedAppProtectionWithDefaults() *TargetedManagedAppProtection {
	this := TargetedManagedAppProtection{}
	return &this
}

// GetIsAssigned returns the IsAssigned field value if set, zero value otherwise.
func (o *TargetedManagedAppProtection) GetIsAssigned() bool {
	if o == nil || o.IsAssigned == nil {
		var ret bool
		return ret
	}
	return *o.IsAssigned
}

// GetIsAssignedOk returns a tuple with the IsAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetedManagedAppProtection) GetIsAssignedOk() (*bool, bool) {
	if o == nil || o.IsAssigned == nil {
		return nil, false
	}
	return o.IsAssigned, true
}

// HasIsAssigned returns a boolean if a field has been set.
func (o *TargetedManagedAppProtection) HasIsAssigned() bool {
	if o != nil && o.IsAssigned != nil {
		return true
	}

	return false
}

// SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.
func (o *TargetedManagedAppProtection) SetIsAssigned(v bool) {
	o.IsAssigned = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *TargetedManagedAppProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetedManagedAppProtection) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *TargetedManagedAppProtection) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *TargetedManagedAppProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

func (o TargetedManagedAppProtection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAssigned != nil {
		toSerialize["isAssigned"] = o.IsAssigned
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableTargetedManagedAppProtection struct {
	value *TargetedManagedAppProtection
	isSet bool
}

func (v NullableTargetedManagedAppProtection) Get() *TargetedManagedAppProtection {
	return v.value
}

func (v *NullableTargetedManagedAppProtection) Set(val *TargetedManagedAppProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetedManagedAppProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetedManagedAppProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetedManagedAppProtection(val *TargetedManagedAppProtection) *NullableTargetedManagedAppProtection {
	return &NullableTargetedManagedAppProtection{value: val, isSet: true}
}

func (v NullableTargetedManagedAppProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetedManagedAppProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


