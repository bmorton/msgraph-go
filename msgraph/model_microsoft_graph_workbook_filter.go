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

// MicrosoftGraphWorkbookFilter struct for MicrosoftGraphWorkbookFilter
type MicrosoftGraphWorkbookFilter struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The currently applied filter on the given column. Read-only.
	Criteria AnyOfmicrosoftGraphWorkbookFilterCriteria `json:"criteria,omitempty"`
}

// NewMicrosoftGraphWorkbookFilter instantiates a new MicrosoftGraphWorkbookFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookFilter() *MicrosoftGraphWorkbookFilter {
	this := MicrosoftGraphWorkbookFilter{}
	return &this
}

// NewMicrosoftGraphWorkbookFilterWithDefaults instantiates a new MicrosoftGraphWorkbookFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookFilterWithDefaults() *MicrosoftGraphWorkbookFilter {
	this := MicrosoftGraphWorkbookFilter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookFilter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookFilter) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookFilter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookFilter) SetId(v string) {
	o.Id = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookFilter) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookFilter) GetCriteriaOk() (*AnyOfmicrosoftGraphWorkbookFilterCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookFilter) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfmicrosoftGraphWorkbookFilterCriteria and assigns it to the Criteria field.
func (o *MicrosoftGraphWorkbookFilter) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria) {
	o.Criteria = v
}

func (o MicrosoftGraphWorkbookFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookFilter struct {
	value *MicrosoftGraphWorkbookFilter
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookFilter) Get() *MicrosoftGraphWorkbookFilter {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookFilter) Set(val *MicrosoftGraphWorkbookFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookFilter(val *MicrosoftGraphWorkbookFilter) *NullableMicrosoftGraphWorkbookFilter {
	return &NullableMicrosoftGraphWorkbookFilter{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


