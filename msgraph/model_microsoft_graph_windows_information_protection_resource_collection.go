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

// MicrosoftGraphWindowsInformationProtectionResourceCollection Windows Information Protection Resource Collection
type MicrosoftGraphWindowsInformationProtectionResourceCollection struct {
	// Display name
	DisplayName *string `json:"displayName,omitempty"`
	// Collection of resources
	Resources *[]*string `json:"resources,omitempty"`
}

// NewMicrosoftGraphWindowsInformationProtectionResourceCollection instantiates a new MicrosoftGraphWindowsInformationProtectionResourceCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWindowsInformationProtectionResourceCollection() *MicrosoftGraphWindowsInformationProtectionResourceCollection {
	this := MicrosoftGraphWindowsInformationProtectionResourceCollection{}
	return &this
}

// NewMicrosoftGraphWindowsInformationProtectionResourceCollectionWithDefaults instantiates a new MicrosoftGraphWindowsInformationProtectionResourceCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWindowsInformationProtectionResourceCollectionWithDefaults() *MicrosoftGraphWindowsInformationProtectionResourceCollection {
	this := MicrosoftGraphWindowsInformationProtectionResourceCollection{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) GetResources() []*string {
	if o == nil || o.Resources == nil {
		var ret []*string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) GetResourcesOk() (*[]*string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []*string and assigns it to the Resources field.
func (o *MicrosoftGraphWindowsInformationProtectionResourceCollection) SetResources(v []*string) {
	o.Resources = &v
}

func (o MicrosoftGraphWindowsInformationProtectionResourceCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWindowsInformationProtectionResourceCollection struct {
	value *MicrosoftGraphWindowsInformationProtectionResourceCollection
	isSet bool
}

func (v NullableMicrosoftGraphWindowsInformationProtectionResourceCollection) Get() *MicrosoftGraphWindowsInformationProtectionResourceCollection {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionResourceCollection) Set(val *MicrosoftGraphWindowsInformationProtectionResourceCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsInformationProtectionResourceCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionResourceCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsInformationProtectionResourceCollection(val *MicrosoftGraphWindowsInformationProtectionResourceCollection) *NullableMicrosoftGraphWindowsInformationProtectionResourceCollection {
	return &NullableMicrosoftGraphWindowsInformationProtectionResourceCollection{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsInformationProtectionResourceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionResourceCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


