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

// MicrosoftGraphWorkbookTableSort struct for MicrosoftGraphWorkbookTableSort
type MicrosoftGraphWorkbookTableSort struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Represents the current conditions used to last sort the table. Read-only.
	Fields *[]*AnyOfmicrosoftGraphWorkbookSortField `json:"fields,omitempty"`
	// Represents whether the casing impacted the last sort of the table. Read-only.
	MatchCase *bool `json:"matchCase,omitempty"`
	// Represents Chinese character ordering method last used to sort the table. The possible values are: PinYin, StrokeCount. Read-only.
	Method *string `json:"method,omitempty"`
}

// NewMicrosoftGraphWorkbookTableSort instantiates a new MicrosoftGraphWorkbookTableSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookTableSort() *MicrosoftGraphWorkbookTableSort {
	this := MicrosoftGraphWorkbookTableSort{}
	return &this
}

// NewMicrosoftGraphWorkbookTableSortWithDefaults instantiates a new MicrosoftGraphWorkbookTableSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookTableSortWithDefaults() *MicrosoftGraphWorkbookTableSort {
	this := MicrosoftGraphWorkbookTableSort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTableSort) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTableSort) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTableSort) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookTableSort) SetId(v string) {
	o.Id = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTableSort) GetFields() []*AnyOfmicrosoftGraphWorkbookSortField {
	if o == nil || o.Fields == nil {
		var ret []*AnyOfmicrosoftGraphWorkbookSortField
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTableSort) GetFieldsOk() (*[]*AnyOfmicrosoftGraphWorkbookSortField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTableSort) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []*AnyOfmicrosoftGraphWorkbookSortField and assigns it to the Fields field.
func (o *MicrosoftGraphWorkbookTableSort) SetFields(v []*AnyOfmicrosoftGraphWorkbookSortField) {
	o.Fields = &v
}

// GetMatchCase returns the MatchCase field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTableSort) GetMatchCase() bool {
	if o == nil || o.MatchCase == nil {
		var ret bool
		return ret
	}
	return *o.MatchCase
}

// GetMatchCaseOk returns a tuple with the MatchCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTableSort) GetMatchCaseOk() (*bool, bool) {
	if o == nil || o.MatchCase == nil {
		return nil, false
	}
	return o.MatchCase, true
}

// HasMatchCase returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTableSort) HasMatchCase() bool {
	if o != nil && o.MatchCase != nil {
		return true
	}

	return false
}

// SetMatchCase gets a reference to the given bool and assigns it to the MatchCase field.
func (o *MicrosoftGraphWorkbookTableSort) SetMatchCase(v bool) {
	o.MatchCase = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookTableSort) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookTableSort) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookTableSort) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *MicrosoftGraphWorkbookTableSort) SetMethod(v string) {
	o.Method = &v
}

func (o MicrosoftGraphWorkbookTableSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphWorkbookTableSort struct {
	value *MicrosoftGraphWorkbookTableSort
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookTableSort) Get() *MicrosoftGraphWorkbookTableSort {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookTableSort) Set(val *MicrosoftGraphWorkbookTableSort) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookTableSort) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookTableSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookTableSort(val *MicrosoftGraphWorkbookTableSort) *NullableMicrosoftGraphWorkbookTableSort {
	return &NullableMicrosoftGraphWorkbookTableSort{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookTableSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookTableSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


