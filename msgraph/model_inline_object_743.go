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

// InlineObject743 struct for InlineObject743
type InlineObject743 struct {
	SourceFile *MicrosoftGraphItemReference `json:"sourceFile,omitempty"`
	DestinationFileName NullableString `json:"destinationFileName,omitempty"`
}

// NewInlineObject743 instantiates a new InlineObject743 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject743() *InlineObject743 {
	this := InlineObject743{}
	return &this
}

// NewInlineObject743WithDefaults instantiates a new InlineObject743 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject743WithDefaults() *InlineObject743 {
	this := InlineObject743{}
	return &this
}

// GetSourceFile returns the SourceFile field value if set, zero value otherwise.
func (o *InlineObject743) GetSourceFile() MicrosoftGraphItemReference {
	if o == nil || o.SourceFile == nil {
		var ret MicrosoftGraphItemReference
		return ret
	}
	return *o.SourceFile
}

// GetSourceFileOk returns a tuple with the SourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject743) GetSourceFileOk() (*MicrosoftGraphItemReference, bool) {
	if o == nil || o.SourceFile == nil {
		return nil, false
	}
	return o.SourceFile, true
}

// HasSourceFile returns a boolean if a field has been set.
func (o *InlineObject743) HasSourceFile() bool {
	if o != nil && o.SourceFile != nil {
		return true
	}

	return false
}

// SetSourceFile gets a reference to the given MicrosoftGraphItemReference and assigns it to the SourceFile field.
func (o *InlineObject743) SetSourceFile(v MicrosoftGraphItemReference) {
	o.SourceFile = &v
}

// GetDestinationFileName returns the DestinationFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject743) GetDestinationFileName() string {
	if o == nil || o.DestinationFileName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationFileName.Get()
}

// GetDestinationFileNameOk returns a tuple with the DestinationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject743) GetDestinationFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DestinationFileName.Get(), o.DestinationFileName.IsSet()
}

// HasDestinationFileName returns a boolean if a field has been set.
func (o *InlineObject743) HasDestinationFileName() bool {
	if o != nil && o.DestinationFileName.IsSet() {
		return true
	}

	return false
}

// SetDestinationFileName gets a reference to the given NullableString and assigns it to the DestinationFileName field.
func (o *InlineObject743) SetDestinationFileName(v string) {
	o.DestinationFileName.Set(&v)
}
// SetDestinationFileNameNil sets the value for DestinationFileName to be an explicit nil
func (o *InlineObject743) SetDestinationFileNameNil() {
	o.DestinationFileName.Set(nil)
}

// UnsetDestinationFileName ensures that no value is present for DestinationFileName, not even an explicit nil
func (o *InlineObject743) UnsetDestinationFileName() {
	o.DestinationFileName.Unset()
}

func (o InlineObject743) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceFile != nil {
		toSerialize["sourceFile"] = o.SourceFile
	}
	if o.DestinationFileName.IsSet() {
		toSerialize["destinationFileName"] = o.DestinationFileName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject743 struct {
	value *InlineObject743
	isSet bool
}

func (v NullableInlineObject743) Get() *InlineObject743 {
	return v.value
}

func (v *NullableInlineObject743) Set(val *InlineObject743) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject743) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject743) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject743(val *InlineObject743) *NullableInlineObject743 {
	return &NullableInlineObject743{value: val, isSet: true}
}

func (v NullableInlineObject743) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject743) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


