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

// MicrosoftGraphEducationAssignmentResource struct for MicrosoftGraphEducationAssignmentResource
type MicrosoftGraphEducationAssignmentResource struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Indicates whether this resource should be copied to each student submission for modification and submission. Required
	DistributeForStudentWork NullableBool `json:"distributeForStudentWork,omitempty"`
	// Resource object that has been associated with this assignment.
	Resource AnyOfmicrosoftGraphEducationResource `json:"resource,omitempty"`
}

// NewMicrosoftGraphEducationAssignmentResource instantiates a new MicrosoftGraphEducationAssignmentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphEducationAssignmentResource() *MicrosoftGraphEducationAssignmentResource {
	this := MicrosoftGraphEducationAssignmentResource{}
	return &this
}

// NewMicrosoftGraphEducationAssignmentResourceWithDefaults instantiates a new MicrosoftGraphEducationAssignmentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphEducationAssignmentResourceWithDefaults() *MicrosoftGraphEducationAssignmentResource {
	this := MicrosoftGraphEducationAssignmentResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationAssignmentResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationAssignmentResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationAssignmentResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphEducationAssignmentResource) SetId(v string) {
	o.Id = &v
}

// GetDistributeForStudentWork returns the DistributeForStudentWork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationAssignmentResource) GetDistributeForStudentWork() bool {
	if o == nil || o.DistributeForStudentWork.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DistributeForStudentWork.Get()
}

// GetDistributeForStudentWorkOk returns a tuple with the DistributeForStudentWork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationAssignmentResource) GetDistributeForStudentWorkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DistributeForStudentWork.Get(), o.DistributeForStudentWork.IsSet()
}

// HasDistributeForStudentWork returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationAssignmentResource) HasDistributeForStudentWork() bool {
	if o != nil && o.DistributeForStudentWork.IsSet() {
		return true
	}

	return false
}

// SetDistributeForStudentWork gets a reference to the given NullableBool and assigns it to the DistributeForStudentWork field.
func (o *MicrosoftGraphEducationAssignmentResource) SetDistributeForStudentWork(v bool) {
	o.DistributeForStudentWork.Set(&v)
}
// SetDistributeForStudentWorkNil sets the value for DistributeForStudentWork to be an explicit nil
func (o *MicrosoftGraphEducationAssignmentResource) SetDistributeForStudentWorkNil() {
	o.DistributeForStudentWork.Set(nil)
}

// UnsetDistributeForStudentWork ensures that no value is present for DistributeForStudentWork, not even an explicit nil
func (o *MicrosoftGraphEducationAssignmentResource) UnsetDistributeForStudentWork() {
	o.DistributeForStudentWork.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationAssignmentResource) GetResource() AnyOfmicrosoftGraphEducationResource {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationResource
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationAssignmentResource) GetResourceOk() (*AnyOfmicrosoftGraphEducationResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationAssignmentResource) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEducationResource and assigns it to the Resource field.
func (o *MicrosoftGraphEducationAssignmentResource) SetResource(v AnyOfmicrosoftGraphEducationResource) {
	o.Resource = v
}

func (o MicrosoftGraphEducationAssignmentResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DistributeForStudentWork.IsSet() {
		toSerialize["distributeForStudentWork"] = o.DistributeForStudentWork.Get()
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphEducationAssignmentResource struct {
	value *MicrosoftGraphEducationAssignmentResource
	isSet bool
}

func (v NullableMicrosoftGraphEducationAssignmentResource) Get() *MicrosoftGraphEducationAssignmentResource {
	return v.value
}

func (v *NullableMicrosoftGraphEducationAssignmentResource) Set(val *MicrosoftGraphEducationAssignmentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEducationAssignmentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEducationAssignmentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEducationAssignmentResource(val *MicrosoftGraphEducationAssignmentResource) *NullableMicrosoftGraphEducationAssignmentResource {
	return &NullableMicrosoftGraphEducationAssignmentResource{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEducationAssignmentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEducationAssignmentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

