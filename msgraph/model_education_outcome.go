/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// EducationOutcome struct for EducationOutcome
type EducationOutcome struct {
	// The individual who updated the resource.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Moment in time when the resource was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
}

// NewEducationOutcome instantiates a new EducationOutcome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationOutcome() *EducationOutcome {
	this := EducationOutcome{}
	return &this
}

// NewEducationOutcomeWithDefaults instantiates a new EducationOutcome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationOutcomeWithDefaults() *EducationOutcome {
	this := EducationOutcome{}
	return &this
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationOutcome) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationOutcome) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *EducationOutcome) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *EducationOutcome) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EducationOutcome) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EducationOutcome) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *EducationOutcome) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *EducationOutcome) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *EducationOutcome) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *EducationOutcome) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

func (o EducationOutcome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEducationOutcome struct {
	value *EducationOutcome
	isSet bool
}

func (v NullableEducationOutcome) Get() *EducationOutcome {
	return v.value
}

func (v *NullableEducationOutcome) Set(val *EducationOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationOutcome(val *EducationOutcome) *NullableEducationOutcome {
	return &NullableEducationOutcome{value: val, isSet: true}
}

func (v NullableEducationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


