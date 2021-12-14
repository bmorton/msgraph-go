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

// InlineObject512 struct for InlineObject512
type InlineObject512 struct {
	Fields *[]*AnyOfmicrosoftGraphWorkbookSortField `json:"fields,omitempty"`
	MatchCase *bool `json:"matchCase,omitempty"`
	HasHeaders *bool `json:"hasHeaders,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Method *string `json:"method,omitempty"`
}

// NewInlineObject512 instantiates a new InlineObject512 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject512() *InlineObject512 {
	this := InlineObject512{}
	var matchCase bool = false
	this.MatchCase = &matchCase
	var hasHeaders bool = false
	this.HasHeaders = &hasHeaders
	return &this
}

// NewInlineObject512WithDefaults instantiates a new InlineObject512 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject512WithDefaults() *InlineObject512 {
	this := InlineObject512{}
	var matchCase bool = false
	this.MatchCase = &matchCase
	var hasHeaders bool = false
	this.HasHeaders = &hasHeaders
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *InlineObject512) GetFields() []*AnyOfmicrosoftGraphWorkbookSortField {
	if o == nil || o.Fields == nil {
		var ret []*AnyOfmicrosoftGraphWorkbookSortField
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject512) GetFieldsOk() (*[]*AnyOfmicrosoftGraphWorkbookSortField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *InlineObject512) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []*AnyOfmicrosoftGraphWorkbookSortField and assigns it to the Fields field.
func (o *InlineObject512) SetFields(v []*AnyOfmicrosoftGraphWorkbookSortField) {
	o.Fields = &v
}

// GetMatchCase returns the MatchCase field value if set, zero value otherwise.
func (o *InlineObject512) GetMatchCase() bool {
	if o == nil || o.MatchCase == nil {
		var ret bool
		return ret
	}
	return *o.MatchCase
}

// GetMatchCaseOk returns a tuple with the MatchCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject512) GetMatchCaseOk() (*bool, bool) {
	if o == nil || o.MatchCase == nil {
		return nil, false
	}
	return o.MatchCase, true
}

// HasMatchCase returns a boolean if a field has been set.
func (o *InlineObject512) HasMatchCase() bool {
	if o != nil && o.MatchCase != nil {
		return true
	}

	return false
}

// SetMatchCase gets a reference to the given bool and assigns it to the MatchCase field.
func (o *InlineObject512) SetMatchCase(v bool) {
	o.MatchCase = &v
}

// GetHasHeaders returns the HasHeaders field value if set, zero value otherwise.
func (o *InlineObject512) GetHasHeaders() bool {
	if o == nil || o.HasHeaders == nil {
		var ret bool
		return ret
	}
	return *o.HasHeaders
}

// GetHasHeadersOk returns a tuple with the HasHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject512) GetHasHeadersOk() (*bool, bool) {
	if o == nil || o.HasHeaders == nil {
		return nil, false
	}
	return o.HasHeaders, true
}

// HasHasHeaders returns a boolean if a field has been set.
func (o *InlineObject512) HasHasHeaders() bool {
	if o != nil && o.HasHeaders != nil {
		return true
	}

	return false
}

// SetHasHeaders gets a reference to the given bool and assigns it to the HasHeaders field.
func (o *InlineObject512) SetHasHeaders(v bool) {
	o.HasHeaders = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *InlineObject512) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject512) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *InlineObject512) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *InlineObject512) SetOrientation(v string) {
	o.Orientation = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineObject512) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject512) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineObject512) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineObject512) SetMethod(v string) {
	o.Method = &v
}

func (o InlineObject512) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.MatchCase != nil {
		toSerialize["matchCase"] = o.MatchCase
	}
	if o.HasHeaders != nil {
		toSerialize["hasHeaders"] = o.HasHeaders
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject512 struct {
	value *InlineObject512
	isSet bool
}

func (v NullableInlineObject512) Get() *InlineObject512 {
	return v.value
}

func (v *NullableInlineObject512) Set(val *InlineObject512) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject512) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject512) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject512(val *InlineObject512) *NullableInlineObject512 {
	return &NullableInlineObject512{value: val, isSet: true}
}

func (v NullableInlineObject512) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject512) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

