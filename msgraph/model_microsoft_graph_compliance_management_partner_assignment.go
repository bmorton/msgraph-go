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

// MicrosoftGraphComplianceManagementPartnerAssignment User group targeting for Compliance Management Partner
type MicrosoftGraphComplianceManagementPartnerAssignment struct {
	// Group assignment target.
	Target AnyOfobject `json:"target,omitempty"`
}

// NewMicrosoftGraphComplianceManagementPartnerAssignment instantiates a new MicrosoftGraphComplianceManagementPartnerAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphComplianceManagementPartnerAssignment() *MicrosoftGraphComplianceManagementPartnerAssignment {
	this := MicrosoftGraphComplianceManagementPartnerAssignment{}
	return &this
}

// NewMicrosoftGraphComplianceManagementPartnerAssignmentWithDefaults instantiates a new MicrosoftGraphComplianceManagementPartnerAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphComplianceManagementPartnerAssignmentWithDefaults() *MicrosoftGraphComplianceManagementPartnerAssignment {
	this := MicrosoftGraphComplianceManagementPartnerAssignment{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphComplianceManagementPartnerAssignment) GetTarget() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphComplianceManagementPartnerAssignment) GetTargetOk() (*AnyOfobject, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return &o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *MicrosoftGraphComplianceManagementPartnerAssignment) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.
func (o *MicrosoftGraphComplianceManagementPartnerAssignment) SetTarget(v AnyOfobject) {
	o.Target = v
}

func (o MicrosoftGraphComplianceManagementPartnerAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphComplianceManagementPartnerAssignment struct {
	value *MicrosoftGraphComplianceManagementPartnerAssignment
	isSet bool
}

func (v NullableMicrosoftGraphComplianceManagementPartnerAssignment) Get() *MicrosoftGraphComplianceManagementPartnerAssignment {
	return v.value
}

func (v *NullableMicrosoftGraphComplianceManagementPartnerAssignment) Set(val *MicrosoftGraphComplianceManagementPartnerAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphComplianceManagementPartnerAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphComplianceManagementPartnerAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphComplianceManagementPartnerAssignment(val *MicrosoftGraphComplianceManagementPartnerAssignment) *NullableMicrosoftGraphComplianceManagementPartnerAssignment {
	return &NullableMicrosoftGraphComplianceManagementPartnerAssignment{value: val, isSet: true}
}

func (v NullableMicrosoftGraphComplianceManagementPartnerAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphComplianceManagementPartnerAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


