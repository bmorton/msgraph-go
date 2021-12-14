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

// MicrosoftGraphOnenoteEntityBaseModel struct for MicrosoftGraphOnenoteEntityBaseModel
type MicrosoftGraphOnenoteEntityBaseModel struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The endpoint where you can get details about the page. Read-only.
	Self NullableString `json:"self,omitempty"`
}

// NewMicrosoftGraphOnenoteEntityBaseModel instantiates a new MicrosoftGraphOnenoteEntityBaseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOnenoteEntityBaseModel() *MicrosoftGraphOnenoteEntityBaseModel {
	this := MicrosoftGraphOnenoteEntityBaseModel{}
	return &this
}

// NewMicrosoftGraphOnenoteEntityBaseModelWithDefaults instantiates a new MicrosoftGraphOnenoteEntityBaseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOnenoteEntityBaseModelWithDefaults() *MicrosoftGraphOnenoteEntityBaseModel {
	this := MicrosoftGraphOnenoteEntityBaseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOnenoteEntityBaseModel) SetId(v string) {
	o.Id = &v
}

// GetSelf returns the Self field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetSelf() string {
	if o == nil || o.Self.Get() == nil {
		var ret string
		return ret
	}
	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetSelfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// HasSelf returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) HasSelf() bool {
	if o != nil && o.Self.IsSet() {
		return true
	}

	return false
}

// SetSelf gets a reference to the given NullableString and assigns it to the Self field.
func (o *MicrosoftGraphOnenoteEntityBaseModel) SetSelf(v string) {
	o.Self.Set(&v)
}
// SetSelfNil sets the value for Self to be an explicit nil
func (o *MicrosoftGraphOnenoteEntityBaseModel) SetSelfNil() {
	o.Self.Set(nil)
}

// UnsetSelf ensures that no value is present for Self, not even an explicit nil
func (o *MicrosoftGraphOnenoteEntityBaseModel) UnsetSelf() {
	o.Self.Unset()
}

func (o MicrosoftGraphOnenoteEntityBaseModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Self.IsSet() {
		toSerialize["self"] = o.Self.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOnenoteEntityBaseModel struct {
	value *MicrosoftGraphOnenoteEntityBaseModel
	isSet bool
}

func (v NullableMicrosoftGraphOnenoteEntityBaseModel) Get() *MicrosoftGraphOnenoteEntityBaseModel {
	return v.value
}

func (v *NullableMicrosoftGraphOnenoteEntityBaseModel) Set(val *MicrosoftGraphOnenoteEntityBaseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOnenoteEntityBaseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOnenoteEntityBaseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOnenoteEntityBaseModel(val *MicrosoftGraphOnenoteEntityBaseModel) *NullableMicrosoftGraphOnenoteEntityBaseModel {
	return &NullableMicrosoftGraphOnenoteEntityBaseModel{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOnenoteEntityBaseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOnenoteEntityBaseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


