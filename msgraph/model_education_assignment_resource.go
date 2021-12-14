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

// EducationAssignmentResource struct for EducationAssignmentResource
type EducationAssignmentResource struct {
	// Indicates whether this resource should be copied to each student submission for modification and submission. Required
	DistributeForStudentWork NullableBool `json:"distributeForStudentWork,omitempty"`
	// Resource object that has been associated with this assignment.
	Resource AnyOfmicrosoftGraphEducationResource `json:"resource,omitempty"`
}

// NewEducationAssignmentResource instantiates a new EducationAssignmentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationAssignmentResource() *EducationAssignmentResource {
	this := EducationAssignmentResource{}
	return &this
}

// NewEducationAssignmentResourceWithDefaults instantiates a new EducationAssignmentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationAssignmentResourceWithDefaults() *EducationAssignmentResource {
	this := EducationAssignmentResource{}
	return &this
}

// GetDistributeForStudentWork returns the DistributeForStudentWork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationAssignmentResource) GetDistributeForStudentWork() bool {
	if o == nil || o.DistributeForStudentWork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DistributeForStudentWork.Get()
}

// GetDistributeForStudentWorkOk returns a tuple with the DistributeForStudentWork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationAssignmentResource) GetDistributeForStudentWorkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DistributeForStudentWork.Get(), o.DistributeForStudentWork.IsSet()
}

// HasDistributeForStudentWork returns a boolean if a field has been set.
func (o *EducationAssignmentResource) HasDistributeForStudentWork() bool {
	if o != nil && o.DistributeForStudentWork.IsSet() {
		return true
	}

	return false
}

// SetDistributeForStudentWork gets a reference to the given NullableBool and assigns it to the DistributeForStudentWork field.
func (o *EducationAssignmentResource) SetDistributeForStudentWork(v bool) {
	o.DistributeForStudentWork.Set(&v)
}
// SetDistributeForStudentWorkNil sets the value for DistributeForStudentWork to be an explicit nil
func (o *EducationAssignmentResource) SetDistributeForStudentWorkNil() {
	o.DistributeForStudentWork.Set(nil)
}

// UnsetDistributeForStudentWork ensures that no value is present for DistributeForStudentWork, not even an explicit nil
func (o *EducationAssignmentResource) UnsetDistributeForStudentWork() {
	o.DistributeForStudentWork.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationAssignmentResource) GetResource() AnyOfmicrosoftGraphEducationResource {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationResource
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationAssignmentResource) GetResourceOk() (*AnyOfmicrosoftGraphEducationResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *EducationAssignmentResource) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEducationResource and assigns it to the Resource field.
func (o *EducationAssignmentResource) SetResource(v AnyOfmicrosoftGraphEducationResource) {
	o.Resource = v
}

func (o EducationAssignmentResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DistributeForStudentWork.IsSet() {
		toSerialize["distributeForStudentWork"] = o.DistributeForStudentWork.Get()
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableEducationAssignmentResource struct {
	value *EducationAssignmentResource
	isSet bool
}

func (v NullableEducationAssignmentResource) Get() *EducationAssignmentResource {
	return v.value
}

func (v *NullableEducationAssignmentResource) Set(val *EducationAssignmentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationAssignmentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationAssignmentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationAssignmentResource(val *EducationAssignmentResource) *NullableEducationAssignmentResource {
	return &NullableEducationAssignmentResource{value: val, isSet: true}
}

func (v NullableEducationAssignmentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationAssignmentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


