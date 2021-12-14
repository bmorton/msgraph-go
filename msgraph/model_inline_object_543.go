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

// InlineObject543 struct for InlineObject543
type InlineObject543 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject543 instantiates a new InlineObject543 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject543() *InlineObject543 {
	this := InlineObject543{}
	return &this
}

// NewInlineObject543WithDefaults instantiates a new InlineObject543 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject543WithDefaults() *InlineObject543 {
	this := InlineObject543{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject543) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject543) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject543) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject543) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject543) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject543 struct {
	value *InlineObject543
	isSet bool
}

func (v NullableInlineObject543) Get() *InlineObject543 {
	return v.value
}

func (v *NullableInlineObject543) Set(val *InlineObject543) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject543) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject543) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject543(val *InlineObject543) *NullableInlineObject543 {
	return &NullableInlineObject543{value: val, isSet: true}
}

func (v NullableInlineObject543) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject543) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

