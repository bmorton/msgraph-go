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

// MicrosoftGraphAccessReviewSet struct for MicrosoftGraphAccessReviewSet
type MicrosoftGraphAccessReviewSet struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	Definitions *[]MicrosoftGraphAccessReviewScheduleDefinition `json:"definitions,omitempty"`
}

// NewMicrosoftGraphAccessReviewSet instantiates a new MicrosoftGraphAccessReviewSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAccessReviewSet() *MicrosoftGraphAccessReviewSet {
	this := MicrosoftGraphAccessReviewSet{}
	return &this
}

// NewMicrosoftGraphAccessReviewSetWithDefaults instantiates a new MicrosoftGraphAccessReviewSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAccessReviewSetWithDefaults() *MicrosoftGraphAccessReviewSet {
	this := MicrosoftGraphAccessReviewSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessReviewSet) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessReviewSet) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewSet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAccessReviewSet) SetId(v string) {
	o.Id = &v
}

// GetDefinitions returns the Definitions field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessReviewSet) GetDefinitions() []MicrosoftGraphAccessReviewScheduleDefinition {
	if o == nil || o.Definitions == nil {
		var ret []MicrosoftGraphAccessReviewScheduleDefinition
		return ret
	}
	return *o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessReviewSet) GetDefinitionsOk() (*[]MicrosoftGraphAccessReviewScheduleDefinition, bool) {
	if o == nil || o.Definitions == nil {
		return nil, false
	}
	return o.Definitions, true
}

// HasDefinitions returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewSet) HasDefinitions() bool {
	if o != nil && o.Definitions != nil {
		return true
	}

	return false
}

// SetDefinitions gets a reference to the given []MicrosoftGraphAccessReviewScheduleDefinition and assigns it to the Definitions field.
func (o *MicrosoftGraphAccessReviewSet) SetDefinitions(v []MicrosoftGraphAccessReviewScheduleDefinition) {
	o.Definitions = &v
}

func (o MicrosoftGraphAccessReviewSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Definitions != nil {
		toSerialize["definitions"] = o.Definitions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAccessReviewSet struct {
	value *MicrosoftGraphAccessReviewSet
	isSet bool
}

func (v NullableMicrosoftGraphAccessReviewSet) Get() *MicrosoftGraphAccessReviewSet {
	return v.value
}

func (v *NullableMicrosoftGraphAccessReviewSet) Set(val *MicrosoftGraphAccessReviewSet) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessReviewSet) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessReviewSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessReviewSet(val *MicrosoftGraphAccessReviewSet) *NullableMicrosoftGraphAccessReviewSet {
	return &NullableMicrosoftGraphAccessReviewSet{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessReviewSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessReviewSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


