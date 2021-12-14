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

// MicrosoftGraphContentTypeInfo struct for MicrosoftGraphContentTypeInfo
type MicrosoftGraphContentTypeInfo struct {
	// The id of the content type.
	Id NullableString `json:"id,omitempty"`
	// The name of the content type.
	Name NullableString `json:"name,omitempty"`
}

// NewMicrosoftGraphContentTypeInfo instantiates a new MicrosoftGraphContentTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphContentTypeInfo() *MicrosoftGraphContentTypeInfo {
	this := MicrosoftGraphContentTypeInfo{}
	return &this
}

// NewMicrosoftGraphContentTypeInfoWithDefaults instantiates a new MicrosoftGraphContentTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphContentTypeInfoWithDefaults() *MicrosoftGraphContentTypeInfo {
	this := MicrosoftGraphContentTypeInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphContentTypeInfo) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphContentTypeInfo) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphContentTypeInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphContentTypeInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphContentTypeInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphContentTypeInfo) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphContentTypeInfo) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphContentTypeInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphContentTypeInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphContentTypeInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphContentTypeInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphContentTypeInfo) UnsetName() {
	o.Name.Unset()
}

func (o MicrosoftGraphContentTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphContentTypeInfo struct {
	value *MicrosoftGraphContentTypeInfo
	isSet bool
}

func (v NullableMicrosoftGraphContentTypeInfo) Get() *MicrosoftGraphContentTypeInfo {
	return v.value
}

func (v *NullableMicrosoftGraphContentTypeInfo) Set(val *MicrosoftGraphContentTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphContentTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphContentTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphContentTypeInfo(val *MicrosoftGraphContentTypeInfo) *NullableMicrosoftGraphContentTypeInfo {
	return &NullableMicrosoftGraphContentTypeInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphContentTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphContentTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

