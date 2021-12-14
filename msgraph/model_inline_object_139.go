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

// InlineObject139 struct for InlineObject139
type InlineObject139 struct {
	SourceFile *MicrosoftGraphItemReference `json:"sourceFile,omitempty"`
	DestinationFileName NullableString `json:"destinationFileName,omitempty"`
}

// NewInlineObject139 instantiates a new InlineObject139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject139() *InlineObject139 {
	this := InlineObject139{}
	return &this
}

// NewInlineObject139WithDefaults instantiates a new InlineObject139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject139WithDefaults() *InlineObject139 {
	this := InlineObject139{}
	return &this
}

// GetSourceFile returns the SourceFile field value if set, zero value otherwise.
func (o *InlineObject139) GetSourceFile() MicrosoftGraphItemReference {
	if o == nil || o.SourceFile == nil {
		var ret MicrosoftGraphItemReference
		return ret
	}
	return *o.SourceFile
}

// GetSourceFileOk returns a tuple with the SourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetSourceFileOk() (*MicrosoftGraphItemReference, bool) {
	if o == nil || o.SourceFile == nil {
		return nil, false
	}
	return o.SourceFile, true
}

// HasSourceFile returns a boolean if a field has been set.
func (o *InlineObject139) HasSourceFile() bool {
	if o != nil && o.SourceFile != nil {
		return true
	}

	return false
}

// SetSourceFile gets a reference to the given MicrosoftGraphItemReference and assigns it to the SourceFile field.
func (o *InlineObject139) SetSourceFile(v MicrosoftGraphItemReference) {
	o.SourceFile = &v
}

// GetDestinationFileName returns the DestinationFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject139) GetDestinationFileName() string {
	if o == nil || o.DestinationFileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationFileName.Get()
}

// GetDestinationFileNameOk returns a tuple with the DestinationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject139) GetDestinationFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DestinationFileName.Get(), o.DestinationFileName.IsSet()
}

// HasDestinationFileName returns a boolean if a field has been set.
func (o *InlineObject139) HasDestinationFileName() bool {
	if o != nil && o.DestinationFileName.IsSet() {
		return true
	}

	return false
}

// SetDestinationFileName gets a reference to the given NullableString and assigns it to the DestinationFileName field.
func (o *InlineObject139) SetDestinationFileName(v string) {
	o.DestinationFileName.Set(&v)
}
// SetDestinationFileNameNil sets the value for DestinationFileName to be an explicit nil
func (o *InlineObject139) SetDestinationFileNameNil() {
	o.DestinationFileName.Set(nil)
}

// UnsetDestinationFileName ensures that no value is present for DestinationFileName, not even an explicit nil
func (o *InlineObject139) UnsetDestinationFileName() {
	o.DestinationFileName.Unset()
}

func (o InlineObject139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceFile != nil {
		toSerialize["sourceFile"] = o.SourceFile
	}
	if o.DestinationFileName.IsSet() {
		toSerialize["destinationFileName"] = o.DestinationFileName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject139 struct {
	value *InlineObject139
	isSet bool
}

func (v NullableInlineObject139) Get() *InlineObject139 {
	return v.value
}

func (v *NullableInlineObject139) Set(val *InlineObject139) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject139) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject139(val *InlineObject139) *NullableInlineObject139 {
	return &NullableInlineObject139{value: val, isSet: true}
}

func (v NullableInlineObject139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


