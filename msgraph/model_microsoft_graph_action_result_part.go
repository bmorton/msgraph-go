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

// MicrosoftGraphActionResultPart struct for MicrosoftGraphActionResultPart
type MicrosoftGraphActionResultPart struct {
	// The error that occurred, if any, during the course of the bulk operation.
	Error AnyOfmicrosoftGraphPublicError `json:"error,omitempty"`
}

// NewMicrosoftGraphActionResultPart instantiates a new MicrosoftGraphActionResultPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphActionResultPart() *MicrosoftGraphActionResultPart {
	this := MicrosoftGraphActionResultPart{}
	return &this
}

// NewMicrosoftGraphActionResultPartWithDefaults instantiates a new MicrosoftGraphActionResultPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphActionResultPartWithDefaults() *MicrosoftGraphActionResultPart {
	this := MicrosoftGraphActionResultPart{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphActionResultPart) GetError() AnyOfmicrosoftGraphPublicError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPublicError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphActionResultPart) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MicrosoftGraphActionResultPart) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphPublicError and assigns it to the Error field.
func (o *MicrosoftGraphActionResultPart) SetError(v AnyOfmicrosoftGraphPublicError) {
	o.Error = v
}

func (o MicrosoftGraphActionResultPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphActionResultPart struct {
	value *MicrosoftGraphActionResultPart
	isSet bool
}

func (v NullableMicrosoftGraphActionResultPart) Get() *MicrosoftGraphActionResultPart {
	return v.value
}

func (v *NullableMicrosoftGraphActionResultPart) Set(val *MicrosoftGraphActionResultPart) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphActionResultPart) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphActionResultPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphActionResultPart(val *MicrosoftGraphActionResultPart) *NullableMicrosoftGraphActionResultPart {
	return &NullableMicrosoftGraphActionResultPart{value: val, isSet: true}
}

func (v NullableMicrosoftGraphActionResultPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphActionResultPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


