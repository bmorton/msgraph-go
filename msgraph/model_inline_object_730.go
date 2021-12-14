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

// InlineObject730 struct for InlineObject730
type InlineObject730 struct {
	SourceFile *MicrosoftGraphItemReference `json:"sourceFile,omitempty"`
	DestinationFileName NullableString `json:"destinationFileName,omitempty"`
}

// NewInlineObject730 instantiates a new InlineObject730 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject730() *InlineObject730 {
	this := InlineObject730{}
	return &this
}

// NewInlineObject730WithDefaults instantiates a new InlineObject730 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject730WithDefaults() *InlineObject730 {
	this := InlineObject730{}
	return &this
}

// GetSourceFile returns the SourceFile field value if set, zero value otherwise.
func (o *InlineObject730) GetSourceFile() MicrosoftGraphItemReference {
	if o == nil || o.SourceFile == nil {
		var ret MicrosoftGraphItemReference
		return ret
	}
	return *o.SourceFile
}

// GetSourceFileOk returns a tuple with the SourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject730) GetSourceFileOk() (*MicrosoftGraphItemReference, bool) {
	if o == nil || o.SourceFile == nil {
		return nil, false
	}
	return o.SourceFile, true
}

// HasSourceFile returns a boolean if a field has been set.
func (o *InlineObject730) HasSourceFile() bool {
	if o != nil && o.SourceFile != nil {
		return true
	}

	return false
}

// SetSourceFile gets a reference to the given MicrosoftGraphItemReference and assigns it to the SourceFile field.
func (o *InlineObject730) SetSourceFile(v MicrosoftGraphItemReference) {
	o.SourceFile = &v
}

// GetDestinationFileName returns the DestinationFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject730) GetDestinationFileName() string {
	if o == nil || o.DestinationFileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationFileName.Get()
}

// GetDestinationFileNameOk returns a tuple with the DestinationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject730) GetDestinationFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DestinationFileName.Get(), o.DestinationFileName.IsSet()
}

// HasDestinationFileName returns a boolean if a field has been set.
func (o *InlineObject730) HasDestinationFileName() bool {
	if o != nil && o.DestinationFileName.IsSet() {
		return true
	}

	return false
}

// SetDestinationFileName gets a reference to the given NullableString and assigns it to the DestinationFileName field.
func (o *InlineObject730) SetDestinationFileName(v string) {
	o.DestinationFileName.Set(&v)
}
// SetDestinationFileNameNil sets the value for DestinationFileName to be an explicit nil
func (o *InlineObject730) SetDestinationFileNameNil() {
	o.DestinationFileName.Set(nil)
}

// UnsetDestinationFileName ensures that no value is present for DestinationFileName, not even an explicit nil
func (o *InlineObject730) UnsetDestinationFileName() {
	o.DestinationFileName.Unset()
}

func (o InlineObject730) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceFile != nil {
		toSerialize["sourceFile"] = o.SourceFile
	}
	if o.DestinationFileName.IsSet() {
		toSerialize["destinationFileName"] = o.DestinationFileName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject730 struct {
	value *InlineObject730
	isSet bool
}

func (v NullableInlineObject730) Get() *InlineObject730 {
	return v.value
}

func (v *NullableInlineObject730) Set(val *InlineObject730) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject730) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject730) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject730(val *InlineObject730) *NullableInlineObject730 {
	return &NullableInlineObject730{value: val, isSet: true}
}

func (v NullableInlineObject730) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject730) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

