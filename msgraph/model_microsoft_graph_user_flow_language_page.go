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

// MicrosoftGraphUserFlowLanguagePage struct for MicrosoftGraphUserFlowLanguagePage
type MicrosoftGraphUserFlowLanguagePage struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
}

// NewMicrosoftGraphUserFlowLanguagePage instantiates a new MicrosoftGraphUserFlowLanguagePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserFlowLanguagePage() *MicrosoftGraphUserFlowLanguagePage {
	this := MicrosoftGraphUserFlowLanguagePage{}
	return &this
}

// NewMicrosoftGraphUserFlowLanguagePageWithDefaults instantiates a new MicrosoftGraphUserFlowLanguagePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserFlowLanguagePageWithDefaults() *MicrosoftGraphUserFlowLanguagePage {
	this := MicrosoftGraphUserFlowLanguagePage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUserFlowLanguagePage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserFlowLanguagePage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserFlowLanguagePage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUserFlowLanguagePage) SetId(v string) {
	o.Id = &v
}

func (o MicrosoftGraphUserFlowLanguagePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserFlowLanguagePage struct {
	value *MicrosoftGraphUserFlowLanguagePage
	isSet bool
}

func (v NullableMicrosoftGraphUserFlowLanguagePage) Get() *MicrosoftGraphUserFlowLanguagePage {
	return v.value
}

func (v *NullableMicrosoftGraphUserFlowLanguagePage) Set(val *MicrosoftGraphUserFlowLanguagePage) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserFlowLanguagePage) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserFlowLanguagePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserFlowLanguagePage(val *MicrosoftGraphUserFlowLanguagePage) *NullableMicrosoftGraphUserFlowLanguagePage {
	return &NullableMicrosoftGraphUserFlowLanguagePage{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserFlowLanguagePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserFlowLanguagePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


