/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphEducationRubric struct for MicrosoftGraphEducationRubric
type MicrosoftGraphEducationRubric struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The user who created this resource.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The description of this rubric.
	Description AnyOfmicrosoftGraphEducationItemBody `json:"description,omitempty"`
	// The name of this rubric.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The grading type of this rubric -- null for a no-points rubric, or educationAssignmentPointsGradeType for a points rubric.
	Grading AnyOfobject `json:"grading,omitempty"`
	// The last user to modify the resource.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// The collection of levels making up this rubric.
	Levels *[]*AnyOfmicrosoftGraphRubricLevel `json:"levels,omitempty"`
	// The collection of qualities making up this rubric.
	Qualities *[]*AnyOfmicrosoftGraphRubricQuality `json:"qualities,omitempty"`
}

// NewMicrosoftGraphEducationRubric instantiates a new MicrosoftGraphEducationRubric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphEducationRubric() *MicrosoftGraphEducationRubric {
	this := MicrosoftGraphEducationRubric{}
	return &this
}

// NewMicrosoftGraphEducationRubricWithDefaults instantiates a new MicrosoftGraphEducationRubric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphEducationRubricWithDefaults() *MicrosoftGraphEducationRubric {
	this := MicrosoftGraphEducationRubric{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationRubric) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationRubric) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphEducationRubric) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphEducationRubric) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphEducationRubric) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphEducationRubric) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphEducationRubric) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetDescription() AnyOfmicrosoftGraphEducationItemBody {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationItemBody
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetDescriptionOk() (*AnyOfmicrosoftGraphEducationItemBody, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AnyOfmicrosoftGraphEducationItemBody and assigns it to the Description field.
func (o *MicrosoftGraphEducationRubric) SetDescription(v AnyOfmicrosoftGraphEducationItemBody) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphEducationRubric) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphEducationRubric) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphEducationRubric) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetGrading returns the Grading field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetGrading() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Grading
}

// GetGradingOk returns a tuple with the Grading field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetGradingOk() (*AnyOfobject, bool) {
	if o == nil || o.Grading == nil {
		return nil, false
	}
	return &o.Grading, true
}

// HasGrading returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasGrading() bool {
	if o != nil && o.Grading != nil {
		return true
	}

	return false
}

// SetGrading gets a reference to the given AnyOfobject and assigns it to the Grading field.
func (o *MicrosoftGraphEducationRubric) SetGrading(v AnyOfobject) {
	o.Grading = v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphEducationRubric) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationRubric) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationRubric) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphEducationRubric) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphEducationRubric) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphEducationRubric) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetLevels returns the Levels field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationRubric) GetLevels() []*AnyOfmicrosoftGraphRubricLevel {
	if o == nil || o.Levels == nil {
		var ret []*AnyOfmicrosoftGraphRubricLevel
		return ret
	}
	return *o.Levels
}

// GetLevelsOk returns a tuple with the Levels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationRubric) GetLevelsOk() (*[]*AnyOfmicrosoftGraphRubricLevel, bool) {
	if o == nil || o.Levels == nil {
		return nil, false
	}
	return o.Levels, true
}

// HasLevels returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasLevels() bool {
	if o != nil && o.Levels != nil {
		return true
	}

	return false
}

// SetLevels gets a reference to the given []*AnyOfmicrosoftGraphRubricLevel and assigns it to the Levels field.
func (o *MicrosoftGraphEducationRubric) SetLevels(v []*AnyOfmicrosoftGraphRubricLevel) {
	o.Levels = &v
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationRubric) GetQualities() []*AnyOfmicrosoftGraphRubricQuality {
	if o == nil || o.Qualities == nil {
		var ret []*AnyOfmicrosoftGraphRubricQuality
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationRubric) GetQualitiesOk() (*[]*AnyOfmicrosoftGraphRubricQuality, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationRubric) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given []*AnyOfmicrosoftGraphRubricQuality and assigns it to the Qualities field.
func (o *MicrosoftGraphEducationRubric) SetQualities(v []*AnyOfmicrosoftGraphRubricQuality) {
	o.Qualities = &v
}

func (o MicrosoftGraphEducationRubric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Grading != nil {
		toSerialize["grading"] = o.Grading
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.Levels != nil {
		toSerialize["levels"] = o.Levels
	}
	if o.Qualities != nil {
		toSerialize["qualities"] = o.Qualities
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphEducationRubric struct {
	value *MicrosoftGraphEducationRubric
	isSet bool
}

func (v NullableMicrosoftGraphEducationRubric) Get() *MicrosoftGraphEducationRubric {
	return v.value
}

func (v *NullableMicrosoftGraphEducationRubric) Set(val *MicrosoftGraphEducationRubric) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEducationRubric) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEducationRubric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEducationRubric(val *MicrosoftGraphEducationRubric) *NullableMicrosoftGraphEducationRubric {
	return &NullableMicrosoftGraphEducationRubric{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEducationRubric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEducationRubric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


