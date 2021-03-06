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

// MicrosoftGraphRubricCriterion struct for MicrosoftGraphRubricCriterion
type MicrosoftGraphRubricCriterion struct {
	// The description of this criterion.
	Description AnyOfmicrosoftGraphEducationItemBody `json:"description,omitempty"`
}

// NewMicrosoftGraphRubricCriterion instantiates a new MicrosoftGraphRubricCriterion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRubricCriterion() *MicrosoftGraphRubricCriterion {
	this := MicrosoftGraphRubricCriterion{}
	return &this
}

// NewMicrosoftGraphRubricCriterionWithDefaults instantiates a new MicrosoftGraphRubricCriterion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRubricCriterionWithDefaults() *MicrosoftGraphRubricCriterion {
	this := MicrosoftGraphRubricCriterion{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRubricCriterion) GetDescription() AnyOfmicrosoftGraphEducationItemBody {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationItemBody
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRubricCriterion) GetDescriptionOk() (*AnyOfmicrosoftGraphEducationItemBody, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphRubricCriterion) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AnyOfmicrosoftGraphEducationItemBody and assigns it to the Description field.
func (o *MicrosoftGraphRubricCriterion) SetDescription(v AnyOfmicrosoftGraphEducationItemBody) {
	o.Description = v
}

func (o MicrosoftGraphRubricCriterion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRubricCriterion struct {
	value *MicrosoftGraphRubricCriterion
	isSet bool
}

func (v NullableMicrosoftGraphRubricCriterion) Get() *MicrosoftGraphRubricCriterion {
	return v.value
}

func (v *NullableMicrosoftGraphRubricCriterion) Set(val *MicrosoftGraphRubricCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRubricCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRubricCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRubricCriterion(val *MicrosoftGraphRubricCriterion) *NullableMicrosoftGraphRubricCriterion {
	return &NullableMicrosoftGraphRubricCriterion{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRubricCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRubricCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


