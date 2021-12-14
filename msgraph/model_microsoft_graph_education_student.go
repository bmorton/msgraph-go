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

// MicrosoftGraphEducationStudent struct for MicrosoftGraphEducationStudent
type MicrosoftGraphEducationStudent struct {
	// Birth date of the student.
	BirthDate NullableString `json:"birthDate,omitempty"`
	// ID of the student in the source system.
	ExternalId NullableString `json:"externalId,omitempty"`
	// The possible values are: female, male, other, unknownFutureValue.
	Gender AnyOfmicrosoftGraphEducationGender `json:"gender,omitempty"`
	// Current grade level of the student.
	Grade NullableString `json:"grade,omitempty"`
	// Year the student is graduating from the school.
	GraduationYear NullableString `json:"graduationYear,omitempty"`
	// Student Number.
	StudentNumber NullableString `json:"studentNumber,omitempty"`
}

// NewMicrosoftGraphEducationStudent instantiates a new MicrosoftGraphEducationStudent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphEducationStudent() *MicrosoftGraphEducationStudent {
	this := MicrosoftGraphEducationStudent{}
	return &this
}

// NewMicrosoftGraphEducationStudentWithDefaults instantiates a new MicrosoftGraphEducationStudent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphEducationStudentWithDefaults() *MicrosoftGraphEducationStudent {
	this := MicrosoftGraphEducationStudent{}
	return &this
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationStudent) GetBirthDate() string {
	if o == nil || o.BirthDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationStudent) GetBirthDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationStudent) HasBirthDate() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableString and assigns it to the BirthDate field.
func (o *MicrosoftGraphEducationStudent) SetBirthDate(v string) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *MicrosoftGraphEducationStudent) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *MicrosoftGraphEducationStudent) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationStudent) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationStudent) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationStudent) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *MicrosoftGraphEducationStudent) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *MicrosoftGraphEducationStudent) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *MicrosoftGraphEducationStudent) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationStudent) GetGender() AnyOfmicrosoftGraphEducationGender {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationGender
		return ret
	}
	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationStudent) GetGenderOk() (*AnyOfmicrosoftGraphEducationGender, bool) {
	if o == nil || o.Gender == nil {
		return nil, false
	}
	return &o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationStudent) HasGender() bool {
	if o != nil && o.Gender != nil {
		return true
	}

	return false
}

// SetGender gets a reference to the given AnyOfmicrosoftGraphEducationGender and assigns it to the Gender field.
func (o *MicrosoftGraphEducationStudent) SetGender(v AnyOfmicrosoftGraphEducationGender) {
	o.Gender = v
}

// GetGrade returns the Grade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationStudent) GetGrade() string {
	if o == nil || o.Grade.Get() == nil {
		var ret string
		return ret
	}
	return *o.Grade.Get()
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationStudent) GetGradeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Grade.Get(), o.Grade.IsSet()
}

// HasGrade returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationStudent) HasGrade() bool {
	if o != nil && o.Grade.IsSet() {
		return true
	}

	return false
}

// SetGrade gets a reference to the given NullableString and assigns it to the Grade field.
func (o *MicrosoftGraphEducationStudent) SetGrade(v string) {
	o.Grade.Set(&v)
}
// SetGradeNil sets the value for Grade to be an explicit nil
func (o *MicrosoftGraphEducationStudent) SetGradeNil() {
	o.Grade.Set(nil)
}

// UnsetGrade ensures that no value is present for Grade, not even an explicit nil
func (o *MicrosoftGraphEducationStudent) UnsetGrade() {
	o.Grade.Unset()
}

// GetGraduationYear returns the GraduationYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationStudent) GetGraduationYear() string {
	if o == nil || o.GraduationYear.Get() == nil {
		var ret string
		return ret
	}
	return *o.GraduationYear.Get()
}

// GetGraduationYearOk returns a tuple with the GraduationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationStudent) GetGraduationYearOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GraduationYear.Get(), o.GraduationYear.IsSet()
}

// HasGraduationYear returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationStudent) HasGraduationYear() bool {
	if o != nil && o.GraduationYear.IsSet() {
		return true
	}

	return false
}

// SetGraduationYear gets a reference to the given NullableString and assigns it to the GraduationYear field.
func (o *MicrosoftGraphEducationStudent) SetGraduationYear(v string) {
	o.GraduationYear.Set(&v)
}
// SetGraduationYearNil sets the value for GraduationYear to be an explicit nil
func (o *MicrosoftGraphEducationStudent) SetGraduationYearNil() {
	o.GraduationYear.Set(nil)
}

// UnsetGraduationYear ensures that no value is present for GraduationYear, not even an explicit nil
func (o *MicrosoftGraphEducationStudent) UnsetGraduationYear() {
	o.GraduationYear.Unset()
}

// GetStudentNumber returns the StudentNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationStudent) GetStudentNumber() string {
	if o == nil || o.StudentNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.StudentNumber.Get()
}

// GetStudentNumberOk returns a tuple with the StudentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationStudent) GetStudentNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StudentNumber.Get(), o.StudentNumber.IsSet()
}

// HasStudentNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationStudent) HasStudentNumber() bool {
	if o != nil && o.StudentNumber.IsSet() {
		return true
	}

	return false
}

// SetStudentNumber gets a reference to the given NullableString and assigns it to the StudentNumber field.
func (o *MicrosoftGraphEducationStudent) SetStudentNumber(v string) {
	o.StudentNumber.Set(&v)
}
// SetStudentNumberNil sets the value for StudentNumber to be an explicit nil
func (o *MicrosoftGraphEducationStudent) SetStudentNumberNil() {
	o.StudentNumber.Set(nil)
}

// UnsetStudentNumber ensures that no value is present for StudentNumber, not even an explicit nil
func (o *MicrosoftGraphEducationStudent) UnsetStudentNumber() {
	o.StudentNumber.Unset()
}

func (o MicrosoftGraphEducationStudent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BirthDate.IsSet() {
		toSerialize["birthDate"] = o.BirthDate.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.Gender != nil {
		toSerialize["gender"] = o.Gender
	}
	if o.Grade.IsSet() {
		toSerialize["grade"] = o.Grade.Get()
	}
	if o.GraduationYear.IsSet() {
		toSerialize["graduationYear"] = o.GraduationYear.Get()
	}
	if o.StudentNumber.IsSet() {
		toSerialize["studentNumber"] = o.StudentNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphEducationStudent struct {
	value *MicrosoftGraphEducationStudent
	isSet bool
}

func (v NullableMicrosoftGraphEducationStudent) Get() *MicrosoftGraphEducationStudent {
	return v.value
}

func (v *NullableMicrosoftGraphEducationStudent) Set(val *MicrosoftGraphEducationStudent) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEducationStudent) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEducationStudent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEducationStudent(val *MicrosoftGraphEducationStudent) *NullableMicrosoftGraphEducationStudent {
	return &NullableMicrosoftGraphEducationStudent{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEducationStudent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEducationStudent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


