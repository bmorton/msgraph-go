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

// InlineObject727 struct for InlineObject727
type InlineObject727 struct {
	SourceFile *MicrosoftGraphItemReference `json:"sourceFile,omitempty"`
	DestinationFileName NullableString `json:"destinationFileName,omitempty"`
}

// NewInlineObject727 instantiates a new InlineObject727 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject727() *InlineObject727 {
	this := InlineObject727{}
	return &this
}

// NewInlineObject727WithDefaults instantiates a new InlineObject727 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject727WithDefaults() *InlineObject727 {
	this := InlineObject727{}
	return &this
}

// GetSourceFile returns the SourceFile field value if set, zero value otherwise.
func (o *InlineObject727) GetSourceFile() MicrosoftGraphItemReference {
	if o == nil || o.SourceFile == nil {
		var ret MicrosoftGraphItemReference
		return ret
	}
	return *o.SourceFile
}

// GetSourceFileOk returns a tuple with the SourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject727) GetSourceFileOk() (*MicrosoftGraphItemReference, bool) {
	if o == nil || o.SourceFile == nil {
		return nil, false
	}
	return o.SourceFile, true
}

// HasSourceFile returns a boolean if a field has been set.
func (o *InlineObject727) HasSourceFile() bool {
	if o != nil && o.SourceFile != nil {
		return true
	}

	return false
}

// SetSourceFile gets a reference to the given MicrosoftGraphItemReference and assigns it to the SourceFile field.
func (o *InlineObject727) SetSourceFile(v MicrosoftGraphItemReference) {
	o.SourceFile = &v
}

// GetDestinationFileName returns the DestinationFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject727) GetDestinationFileName() string {
	if o == nil || o.DestinationFileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationFileName.Get()
}

// GetDestinationFileNameOk returns a tuple with the DestinationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject727) GetDestinationFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DestinationFileName.Get(), o.DestinationFileName.IsSet()
}

// HasDestinationFileName returns a boolean if a field has been set.
func (o *InlineObject727) HasDestinationFileName() bool {
	if o != nil && o.DestinationFileName.IsSet() {
		return true
	}

	return false
}

// SetDestinationFileName gets a reference to the given NullableString and assigns it to the DestinationFileName field.
func (o *InlineObject727) SetDestinationFileName(v string) {
	o.DestinationFileName.Set(&v)
}
// SetDestinationFileNameNil sets the value for DestinationFileName to be an explicit nil
func (o *InlineObject727) SetDestinationFileNameNil() {
	o.DestinationFileName.Set(nil)
}

// UnsetDestinationFileName ensures that no value is present for DestinationFileName, not even an explicit nil
func (o *InlineObject727) UnsetDestinationFileName() {
	o.DestinationFileName.Unset()
}

func (o InlineObject727) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceFile != nil {
		toSerialize["sourceFile"] = o.SourceFile
	}
	if o.DestinationFileName.IsSet() {
		toSerialize["destinationFileName"] = o.DestinationFileName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject727 struct {
	value *InlineObject727
	isSet bool
}

func (v NullableInlineObject727) Get() *InlineObject727 {
	return v.value
}

func (v *NullableInlineObject727) Set(val *InlineObject727) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject727) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject727) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject727(val *InlineObject727) *NullableInlineObject727 {
	return &NullableInlineObject727{value: val, isSet: true}
}

func (v NullableInlineObject727) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject727) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


