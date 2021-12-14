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

// MicrosoftGraphSearchEntity struct for MicrosoftGraphSearchEntity
type MicrosoftGraphSearchEntity struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
}

// NewMicrosoftGraphSearchEntity instantiates a new MicrosoftGraphSearchEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSearchEntity() *MicrosoftGraphSearchEntity {
	this := MicrosoftGraphSearchEntity{}
	return &this
}

// NewMicrosoftGraphSearchEntityWithDefaults instantiates a new MicrosoftGraphSearchEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSearchEntityWithDefaults() *MicrosoftGraphSearchEntity {
	this := MicrosoftGraphSearchEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphSearchEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSearchEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSearchEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSearchEntity) SetId(v string) {
	o.Id = &v
}

func (o MicrosoftGraphSearchEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSearchEntity struct {
	value *MicrosoftGraphSearchEntity
	isSet bool
}

func (v NullableMicrosoftGraphSearchEntity) Get() *MicrosoftGraphSearchEntity {
	return v.value
}

func (v *NullableMicrosoftGraphSearchEntity) Set(val *MicrosoftGraphSearchEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSearchEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSearchEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSearchEntity(val *MicrosoftGraphSearchEntity) *NullableMicrosoftGraphSearchEntity {
	return &NullableMicrosoftGraphSearchEntity{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSearchEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSearchEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


