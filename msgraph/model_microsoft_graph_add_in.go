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

// MicrosoftGraphAddIn struct for MicrosoftGraphAddIn
type MicrosoftGraphAddIn struct {
	Id NullableString `json:"id,omitempty"`
	Properties *[]MicrosoftGraphKeyValue `json:"properties,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMicrosoftGraphAddIn instantiates a new MicrosoftGraphAddIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAddIn() *MicrosoftGraphAddIn {
	this := MicrosoftGraphAddIn{}
	return &this
}

// NewMicrosoftGraphAddInWithDefaults instantiates a new MicrosoftGraphAddIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAddInWithDefaults() *MicrosoftGraphAddIn {
	this := MicrosoftGraphAddIn{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAddIn) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAddIn) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAddIn) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphAddIn) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphAddIn) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphAddIn) UnsetId() {
	o.Id.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MicrosoftGraphAddIn) GetProperties() []MicrosoftGraphKeyValue {
	if o == nil || o.Properties == nil {
		var ret []MicrosoftGraphKeyValue
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAddIn) GetPropertiesOk() (*[]MicrosoftGraphKeyValue, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MicrosoftGraphAddIn) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []MicrosoftGraphKeyValue and assigns it to the Properties field.
func (o *MicrosoftGraphAddIn) SetProperties(v []MicrosoftGraphKeyValue) {
	o.Properties = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MicrosoftGraphAddIn) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAddIn) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphAddIn) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MicrosoftGraphAddIn) SetType(v string) {
	o.Type = &v
}

func (o MicrosoftGraphAddIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAddIn struct {
	value *MicrosoftGraphAddIn
	isSet bool
}

func (v NullableMicrosoftGraphAddIn) Get() *MicrosoftGraphAddIn {
	return v.value
}

func (v *NullableMicrosoftGraphAddIn) Set(val *MicrosoftGraphAddIn) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAddIn) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAddIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAddIn(val *MicrosoftGraphAddIn) *NullableMicrosoftGraphAddIn {
	return &NullableMicrosoftGraphAddIn{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAddIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAddIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


