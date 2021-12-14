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

// InlineObject500 struct for InlineObject500
type InlineObject500 struct {
	Properties *MicrosoftGraphPrintDocumentUploadProperties `json:"properties,omitempty"`
}

// NewInlineObject500 instantiates a new InlineObject500 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject500() *InlineObject500 {
	this := InlineObject500{}
	return &this
}

// NewInlineObject500WithDefaults instantiates a new InlineObject500 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject500WithDefaults() *InlineObject500 {
	this := InlineObject500{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *InlineObject500) GetProperties() MicrosoftGraphPrintDocumentUploadProperties {
	if o == nil || o.Properties == nil {
		var ret MicrosoftGraphPrintDocumentUploadProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject500) GetPropertiesOk() (*MicrosoftGraphPrintDocumentUploadProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *InlineObject500) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given MicrosoftGraphPrintDocumentUploadProperties and assigns it to the Properties field.
func (o *InlineObject500) SetProperties(v MicrosoftGraphPrintDocumentUploadProperties) {
	o.Properties = &v
}

func (o InlineObject500) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject500 struct {
	value *InlineObject500
	isSet bool
}

func (v NullableInlineObject500) Get() *InlineObject500 {
	return v.value
}

func (v *NullableInlineObject500) Set(val *InlineObject500) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject500) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject500) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject500(val *InlineObject500) *NullableInlineObject500 {
	return &NullableInlineObject500{value: val, isSet: true}
}

func (v NullableInlineObject500) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject500) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


