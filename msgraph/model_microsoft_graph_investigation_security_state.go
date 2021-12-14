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

// MicrosoftGraphInvestigationSecurityState struct for MicrosoftGraphInvestigationSecurityState
type MicrosoftGraphInvestigationSecurityState struct {
	Name NullableString `json:"name,omitempty"`
	Status NullableString `json:"status,omitempty"`
}

// NewMicrosoftGraphInvestigationSecurityState instantiates a new MicrosoftGraphInvestigationSecurityState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphInvestigationSecurityState() *MicrosoftGraphInvestigationSecurityState {
	this := MicrosoftGraphInvestigationSecurityState{}
	return &this
}

// NewMicrosoftGraphInvestigationSecurityStateWithDefaults instantiates a new MicrosoftGraphInvestigationSecurityState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphInvestigationSecurityStateWithDefaults() *MicrosoftGraphInvestigationSecurityState {
	this := MicrosoftGraphInvestigationSecurityState{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInvestigationSecurityState) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInvestigationSecurityState) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphInvestigationSecurityState) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphInvestigationSecurityState) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphInvestigationSecurityState) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphInvestigationSecurityState) UnsetName() {
	o.Name.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInvestigationSecurityState) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInvestigationSecurityState) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphInvestigationSecurityState) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *MicrosoftGraphInvestigationSecurityState) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *MicrosoftGraphInvestigationSecurityState) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *MicrosoftGraphInvestigationSecurityState) UnsetStatus() {
	o.Status.Unset()
}

func (o MicrosoftGraphInvestigationSecurityState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphInvestigationSecurityState struct {
	value *MicrosoftGraphInvestigationSecurityState
	isSet bool
}

func (v NullableMicrosoftGraphInvestigationSecurityState) Get() *MicrosoftGraphInvestigationSecurityState {
	return v.value
}

func (v *NullableMicrosoftGraphInvestigationSecurityState) Set(val *MicrosoftGraphInvestigationSecurityState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphInvestigationSecurityState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphInvestigationSecurityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphInvestigationSecurityState(val *MicrosoftGraphInvestigationSecurityState) *NullableMicrosoftGraphInvestigationSecurityState {
	return &NullableMicrosoftGraphInvestigationSecurityState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphInvestigationSecurityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphInvestigationSecurityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

