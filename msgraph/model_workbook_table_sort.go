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

// WorkbookTableSort struct for WorkbookTableSort
type WorkbookTableSort struct {
	// Represents the current conditions used to last sort the table. Read-only.
	Fields *[]*AnyOfmicrosoftGraphWorkbookSortField `json:"fields,omitempty"`
	// Represents whether the casing impacted the last sort of the table. Read-only.
	MatchCase *bool `json:"matchCase,omitempty"`
	// Represents Chinese character ordering method last used to sort the table. The possible values are: PinYin, StrokeCount. Read-only.
	Method *string `json:"method,omitempty"`
}

// NewWorkbookTableSort instantiates a new WorkbookTableSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookTableSort() *WorkbookTableSort {
	this := WorkbookTableSort{}
	return &this
}

// NewWorkbookTableSortWithDefaults instantiates a new WorkbookTableSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookTableSortWithDefaults() *WorkbookTableSort {
	this := WorkbookTableSort{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *WorkbookTableSort) GetFields() []*AnyOfmicrosoftGraphWorkbookSortField {
	if o == nil || o.Fields == nil {
		var ret []*AnyOfmicrosoftGraphWorkbookSortField
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableSort) GetFieldsOk() (*[]*AnyOfmicrosoftGraphWorkbookSortField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *WorkbookTableSort) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []*AnyOfmicrosoftGraphWorkbookSortField and assigns it to the Fields field.
func (o *WorkbookTableSort) SetFields(v []*AnyOfmicrosoftGraphWorkbookSortField) {
	o.Fields = &v
}

// GetMatchCase returns the MatchCase field value if set, zero value otherwise.
func (o *WorkbookTableSort) GetMatchCase() bool {
	if o == nil || o.MatchCase == nil {
		var ret bool
		return ret
	}
	return *o.MatchCase
}

// GetMatchCaseOk returns a tuple with the MatchCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableSort) GetMatchCaseOk() (*bool, bool) {
	if o == nil || o.MatchCase == nil {
		return nil, false
	}
	return o.MatchCase, true
}

// HasMatchCase returns a boolean if a field has been set.
func (o *WorkbookTableSort) HasMatchCase() bool {
	if o != nil && o.MatchCase != nil {
		return true
	}

	return false
}

// SetMatchCase gets a reference to the given bool and assigns it to the MatchCase field.
func (o *WorkbookTableSort) SetMatchCase(v bool) {
	o.MatchCase = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WorkbookTableSort) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableSort) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkbookTableSort) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkbookTableSort) SetMethod(v string) {
	o.Method = &v
}

func (o WorkbookTableSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.MatchCase != nil {
		toSerialize["matchCase"] = o.MatchCase
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookTableSort struct {
	value *WorkbookTableSort
	isSet bool
}

func (v NullableWorkbookTableSort) Get() *WorkbookTableSort {
	return v.value
}

func (v *NullableWorkbookTableSort) Set(val *WorkbookTableSort) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookTableSort) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookTableSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookTableSort(val *WorkbookTableSort) *NullableWorkbookTableSort {
	return &NullableWorkbookTableSort{value: val, isSet: true}
}

func (v NullableWorkbookTableSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookTableSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

