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

// InlineObject527 struct for InlineObject527
type InlineObject527 struct {
	Fields *[]*AnyOfmicrosoftGraphWorkbookSortField `json:"fields,omitempty"`
	MatchCase *bool `json:"matchCase,omitempty"`
	HasHeaders *bool `json:"hasHeaders,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Method *string `json:"method,omitempty"`
}

// NewInlineObject527 instantiates a new InlineObject527 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject527() *InlineObject527 {
	this := InlineObject527{}
	var matchCase bool = false
	this.MatchCase = &matchCase
	var hasHeaders bool = false
	this.HasHeaders = &hasHeaders
	return &this
}

// NewInlineObject527WithDefaults instantiates a new InlineObject527 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject527WithDefaults() *InlineObject527 {
	this := InlineObject527{}
	var matchCase bool = false
	this.MatchCase = &matchCase
	var hasHeaders bool = false
	this.HasHeaders = &hasHeaders
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *InlineObject527) GetFields() []*AnyOfmicrosoftGraphWorkbookSortField {
	if o == nil || o.Fields == nil {
		var ret []*AnyOfmicrosoftGraphWorkbookSortField
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject527) GetFieldsOk() (*[]*AnyOfmicrosoftGraphWorkbookSortField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *InlineObject527) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []*AnyOfmicrosoftGraphWorkbookSortField and assigns it to the Fields field.
func (o *InlineObject527) SetFields(v []*AnyOfmicrosoftGraphWorkbookSortField) {
	o.Fields = &v
}

// GetMatchCase returns the MatchCase field value if set, zero value otherwise.
func (o *InlineObject527) GetMatchCase() bool {
	if o == nil || o.MatchCase == nil {
		var ret bool
		return ret
	}
	return *o.MatchCase
}

// GetMatchCaseOk returns a tuple with the MatchCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject527) GetMatchCaseOk() (*bool, bool) {
	if o == nil || o.MatchCase == nil {
		return nil, false
	}
	return o.MatchCase, true
}

// HasMatchCase returns a boolean if a field has been set.
func (o *InlineObject527) HasMatchCase() bool {
	if o != nil && o.MatchCase != nil {
		return true
	}

	return false
}

// SetMatchCase gets a reference to the given bool and assigns it to the MatchCase field.
func (o *InlineObject527) SetMatchCase(v bool) {
	o.MatchCase = &v
}

// GetHasHeaders returns the HasHeaders field value if set, zero value otherwise.
func (o *InlineObject527) GetHasHeaders() bool {
	if o == nil || o.HasHeaders == nil {
		var ret bool
		return ret
	}
	return *o.HasHeaders
}

// GetHasHeadersOk returns a tuple with the HasHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject527) GetHasHeadersOk() (*bool, bool) {
	if o == nil || o.HasHeaders == nil {
		return nil, false
	}
	return o.HasHeaders, true
}

// HasHasHeaders returns a boolean if a field has been set.
func (o *InlineObject527) HasHasHeaders() bool {
	if o != nil && o.HasHeaders != nil {
		return true
	}

	return false
}

// SetHasHeaders gets a reference to the given bool and assigns it to the HasHeaders field.
func (o *InlineObject527) SetHasHeaders(v bool) {
	o.HasHeaders = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *InlineObject527) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject527) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *InlineObject527) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *InlineObject527) SetOrientation(v string) {
	o.Orientation = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineObject527) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject527) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineObject527) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineObject527) SetMethod(v string) {
	o.Method = &v
}

func (o InlineObject527) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject527 struct {
	value *InlineObject527
	isSet bool
}

func (v NullableInlineObject527) Get() *InlineObject527 {
	return v.value
}

func (v *NullableInlineObject527) Set(val *InlineObject527) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject527) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject527) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject527(val *InlineObject527) *NullableInlineObject527 {
	return &NullableInlineObject527{value: val, isSet: true}
}

func (v NullableInlineObject527) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject527) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


