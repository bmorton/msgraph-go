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

// TargetedManagedAppPolicyAssignment The type for deployment of groups or apps.
type TargetedManagedAppPolicyAssignment struct {
	// Identifier for deployment to a group or app
	Target AnyOfobject `json:"target,omitempty"`
}

// NewTargetedManagedAppPolicyAssignment instantiates a new TargetedManagedAppPolicyAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetedManagedAppPolicyAssignment() *TargetedManagedAppPolicyAssignment {
	this := TargetedManagedAppPolicyAssignment{}
	return &this
}

// NewTargetedManagedAppPolicyAssignmentWithDefaults instantiates a new TargetedManagedAppPolicyAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetedManagedAppPolicyAssignmentWithDefaults() *TargetedManagedAppPolicyAssignment {
	this := TargetedManagedAppPolicyAssignment{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetedManagedAppPolicyAssignment) GetTarget() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetedManagedAppPolicyAssignment) GetTargetOk() (*AnyOfobject, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *TargetedManagedAppPolicyAssignment) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.
func (o *TargetedManagedAppPolicyAssignment) SetTarget(v AnyOfobject) {
	o.Target = v
}

func (o TargetedManagedAppPolicyAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableTargetedManagedAppPolicyAssignment struct {
	value *TargetedManagedAppPolicyAssignment
	isSet bool
}

func (v NullableTargetedManagedAppPolicyAssignment) Get() *TargetedManagedAppPolicyAssignment {
	return v.value
}

func (v *NullableTargetedManagedAppPolicyAssignment) Set(val *TargetedManagedAppPolicyAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetedManagedAppPolicyAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetedManagedAppPolicyAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetedManagedAppPolicyAssignment(val *TargetedManagedAppPolicyAssignment) *NullableTargetedManagedAppPolicyAssignment {
	return &NullableTargetedManagedAppPolicyAssignment{value: val, isSet: true}
}

func (v NullableTargetedManagedAppPolicyAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetedManagedAppPolicyAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


