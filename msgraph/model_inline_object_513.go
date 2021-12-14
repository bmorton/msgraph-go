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

// InlineObject513 struct for InlineObject513
type InlineObject513 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject513 instantiates a new InlineObject513 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject513() *InlineObject513 {
	this := InlineObject513{}
	return &this
}

// NewInlineObject513WithDefaults instantiates a new InlineObject513 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject513WithDefaults() *InlineObject513 {
	this := InlineObject513{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject513) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject513) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject513) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject513) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject513) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject513 struct {
	value *InlineObject513
	isSet bool
}

func (v NullableInlineObject513) Get() *InlineObject513 {
	return v.value
}

func (v *NullableInlineObject513) Set(val *InlineObject513) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject513) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject513) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject513(val *InlineObject513) *NullableInlineObject513 {
	return &NullableInlineObject513{value: val, isSet: true}
}

func (v NullableInlineObject513) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject513) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


