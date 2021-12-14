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

// MicrosoftGraphDeviceCompliancePolicyAssignment struct for MicrosoftGraphDeviceCompliancePolicyAssignment
type MicrosoftGraphDeviceCompliancePolicyAssignment struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Target for the compliance policy assignment.
	Target AnyOfobject `json:"target,omitempty"`
}

// NewMicrosoftGraphDeviceCompliancePolicyAssignment instantiates a new MicrosoftGraphDeviceCompliancePolicyAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDeviceCompliancePolicyAssignment() *MicrosoftGraphDeviceCompliancePolicyAssignment {
	this := MicrosoftGraphDeviceCompliancePolicyAssignment{}
	return &this
}

// NewMicrosoftGraphDeviceCompliancePolicyAssignmentWithDefaults instantiates a new MicrosoftGraphDeviceCompliancePolicyAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDeviceCompliancePolicyAssignmentWithDefaults() *MicrosoftGraphDeviceCompliancePolicyAssignment {
	this := MicrosoftGraphDeviceCompliancePolicyAssignment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) SetId(v string) {
	o.Id = &v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) GetTarget() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) GetTargetOk() (*AnyOfobject, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.
func (o *MicrosoftGraphDeviceCompliancePolicyAssignment) SetTarget(v AnyOfobject) {
	o.Target = v
}

func (o MicrosoftGraphDeviceCompliancePolicyAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDeviceCompliancePolicyAssignment struct {
	value *MicrosoftGraphDeviceCompliancePolicyAssignment
	isSet bool
}

func (v NullableMicrosoftGraphDeviceCompliancePolicyAssignment) Get() *MicrosoftGraphDeviceCompliancePolicyAssignment {
	return v.value
}

func (v *NullableMicrosoftGraphDeviceCompliancePolicyAssignment) Set(val *MicrosoftGraphDeviceCompliancePolicyAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDeviceCompliancePolicyAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDeviceCompliancePolicyAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDeviceCompliancePolicyAssignment(val *MicrosoftGraphDeviceCompliancePolicyAssignment) *NullableMicrosoftGraphDeviceCompliancePolicyAssignment {
	return &NullableMicrosoftGraphDeviceCompliancePolicyAssignment{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDeviceCompliancePolicyAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDeviceCompliancePolicyAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


