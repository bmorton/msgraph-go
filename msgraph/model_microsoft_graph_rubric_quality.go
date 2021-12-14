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

// MicrosoftGraphRubricQuality struct for MicrosoftGraphRubricQuality
type MicrosoftGraphRubricQuality struct {
	// The collection of criteria for this rubric quality.
	Criteria *[]*AnyOfmicrosoftGraphRubricCriterion `json:"criteria,omitempty"`
	// The description of this rubric quality.
	Description AnyOfmicrosoftGraphEducationItemBody `json:"description,omitempty"`
	// The name of this rubric quality.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The ID of this resource.
	QualityId NullableString `json:"qualityId,omitempty"`
	// If present, a numerical weight for this quality.  Weights must add up to 100.
	Weight AnyOfnumberstringstring `json:"weight,omitempty"`
}

// NewMicrosoftGraphRubricQuality instantiates a new MicrosoftGraphRubricQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRubricQuality() *MicrosoftGraphRubricQuality {
	this := MicrosoftGraphRubricQuality{}
	return &this
}

// NewMicrosoftGraphRubricQualityWithDefaults instantiates a new MicrosoftGraphRubricQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRubricQualityWithDefaults() *MicrosoftGraphRubricQuality {
	this := MicrosoftGraphRubricQuality{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *MicrosoftGraphRubricQuality) GetCriteria() []*AnyOfmicrosoftGraphRubricCriterion {
	if o == nil || o.Criteria == nil {
		var ret []*AnyOfmicrosoftGraphRubricCriterion
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRubricQuality) GetCriteriaOk() (*[]*AnyOfmicrosoftGraphRubricCriterion, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *MicrosoftGraphRubricQuality) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []*AnyOfmicrosoftGraphRubricCriterion and assigns it to the Criteria field.
func (o *MicrosoftGraphRubricQuality) SetCriteria(v []*AnyOfmicrosoftGraphRubricCriterion) {
	o.Criteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRubricQuality) GetDescription() AnyOfmicrosoftGraphEducationItemBody {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationItemBody
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRubricQuality) GetDescriptionOk() (*AnyOfmicrosoftGraphEducationItemBody, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphRubricQuality) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AnyOfmicrosoftGraphEducationItemBody and assigns it to the Description field.
func (o *MicrosoftGraphRubricQuality) SetDescription(v AnyOfmicrosoftGraphEducationItemBody) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRubricQuality) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRubricQuality) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphRubricQuality) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphRubricQuality) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphRubricQuality) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphRubricQuality) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetQualityId returns the QualityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRubricQuality) GetQualityId() string {
	if o == nil || o.QualityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.QualityId.Get()
}

// GetQualityIdOk returns a tuple with the QualityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRubricQuality) GetQualityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QualityId.Get(), o.QualityId.IsSet()
}

// HasQualityId returns a boolean if a field has been set.
func (o *MicrosoftGraphRubricQuality) HasQualityId() bool {
	if o != nil && o.QualityId.IsSet() {
		return true
	}

	return false
}

// SetQualityId gets a reference to the given NullableString and assigns it to the QualityId field.
func (o *MicrosoftGraphRubricQuality) SetQualityId(v string) {
	o.QualityId.Set(&v)
}
// SetQualityIdNil sets the value for QualityId to be an explicit nil
func (o *MicrosoftGraphRubricQuality) SetQualityIdNil() {
	o.QualityId.Set(nil)
}

// UnsetQualityId ensures that no value is present for QualityId, not even an explicit nil
func (o *MicrosoftGraphRubricQuality) UnsetQualityId() {
	o.QualityId.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRubricQuality) GetWeight() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRubricQuality) GetWeightOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return &o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *MicrosoftGraphRubricQuality) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Weight field.
func (o *MicrosoftGraphRubricQuality) SetWeight(v AnyOfnumberstringstring) {
	o.Weight = v
}

func (o MicrosoftGraphRubricQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.QualityId.IsSet() {
		toSerialize["qualityId"] = o.QualityId.Get()
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRubricQuality struct {
	value *MicrosoftGraphRubricQuality
	isSet bool
}

func (v NullableMicrosoftGraphRubricQuality) Get() *MicrosoftGraphRubricQuality {
	return v.value
}

func (v *NullableMicrosoftGraphRubricQuality) Set(val *MicrosoftGraphRubricQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRubricQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRubricQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRubricQuality(val *MicrosoftGraphRubricQuality) *NullableMicrosoftGraphRubricQuality {
	return &NullableMicrosoftGraphRubricQuality{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRubricQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRubricQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


