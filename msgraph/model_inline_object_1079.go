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

// InlineObject1079 struct for InlineObject1079
type InlineObject1079 struct {
	AddLicenses *[]MicrosoftGraphAssignedLicense `json:"addLicenses,omitempty"`
	RemoveLicenses *[]string `json:"removeLicenses,omitempty"`
}

// NewInlineObject1079 instantiates a new InlineObject1079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1079() *InlineObject1079 {
	this := InlineObject1079{}
	return &this
}

// NewInlineObject1079WithDefaults instantiates a new InlineObject1079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1079WithDefaults() *InlineObject1079 {
	this := InlineObject1079{}
	return &this
}

// GetAddLicenses returns the AddLicenses field value if set, zero value otherwise.
func (o *InlineObject1079) GetAddLicenses() []MicrosoftGraphAssignedLicense {
	if o == nil || o.AddLicenses == nil {
		var ret []MicrosoftGraphAssignedLicense
		return ret
	}
	return *o.AddLicenses
}

// GetAddLicensesOk returns a tuple with the AddLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1079) GetAddLicensesOk() (*[]MicrosoftGraphAssignedLicense, bool) {
	if o == nil || o.AddLicenses == nil {
		return nil, false
	}
	return o.AddLicenses, true
}

// HasAddLicenses returns a boolean if a field has been set.
func (o *InlineObject1079) HasAddLicenses() bool {
	if o != nil && o.AddLicenses != nil {
		return true
	}

	return false
}

// SetAddLicenses gets a reference to the given []MicrosoftGraphAssignedLicense and assigns it to the AddLicenses field.
func (o *InlineObject1079) SetAddLicenses(v []MicrosoftGraphAssignedLicense) {
	o.AddLicenses = &v
}

// GetRemoveLicenses returns the RemoveLicenses field value if set, zero value otherwise.
func (o *InlineObject1079) GetRemoveLicenses() []string {
	if o == nil || o.RemoveLicenses == nil {
		var ret []string
		return ret
	}
	return *o.RemoveLicenses
}

// GetRemoveLicensesOk returns a tuple with the RemoveLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1079) GetRemoveLicensesOk() (*[]string, bool) {
	if o == nil || o.RemoveLicenses == nil {
		return nil, false
	}
	return o.RemoveLicenses, true
}

// HasRemoveLicenses returns a boolean if a field has been set.
func (o *InlineObject1079) HasRemoveLicenses() bool {
	if o != nil && o.RemoveLicenses != nil {
		return true
	}

	return false
}

// SetRemoveLicenses gets a reference to the given []string and assigns it to the RemoveLicenses field.
func (o *InlineObject1079) SetRemoveLicenses(v []string) {
	o.RemoveLicenses = &v
}

func (o InlineObject1079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddLicenses != nil {
		toSerialize["addLicenses"] = o.AddLicenses
	}
	if o.RemoveLicenses != nil {
		toSerialize["removeLicenses"] = o.RemoveLicenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1079 struct {
	value *InlineObject1079
	isSet bool
}

func (v NullableInlineObject1079) Get() *InlineObject1079 {
	return v.value
}

func (v *NullableInlineObject1079) Set(val *InlineObject1079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1079(val *InlineObject1079) *NullableInlineObject1079 {
	return &NullableInlineObject1079{value: val, isSet: true}
}

func (v NullableInlineObject1079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


