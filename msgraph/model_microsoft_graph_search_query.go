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

// MicrosoftGraphSearchQuery struct for MicrosoftGraphSearchQuery
type MicrosoftGraphSearchQuery struct {
	// The search query containing the search terms. Required.
	QueryString *string `json:"queryString,omitempty"`
}

// NewMicrosoftGraphSearchQuery instantiates a new MicrosoftGraphSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSearchQuery() *MicrosoftGraphSearchQuery {
	this := MicrosoftGraphSearchQuery{}
	return &this
}

// NewMicrosoftGraphSearchQueryWithDefaults instantiates a new MicrosoftGraphSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSearchQueryWithDefaults() *MicrosoftGraphSearchQuery {
	this := MicrosoftGraphSearchQuery{}
	return &this
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *MicrosoftGraphSearchQuery) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSearchQuery) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *MicrosoftGraphSearchQuery) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *MicrosoftGraphSearchQuery) SetQueryString(v string) {
	o.QueryString = &v
}

func (o MicrosoftGraphSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSearchQuery struct {
	value *MicrosoftGraphSearchQuery
	isSet bool
}

func (v NullableMicrosoftGraphSearchQuery) Get() *MicrosoftGraphSearchQuery {
	return v.value
}

func (v *NullableMicrosoftGraphSearchQuery) Set(val *MicrosoftGraphSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSearchQuery(val *MicrosoftGraphSearchQuery) *NullableMicrosoftGraphSearchQuery {
	return &NullableMicrosoftGraphSearchQuery{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

