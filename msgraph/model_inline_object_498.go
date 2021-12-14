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

// InlineObject498 struct for InlineObject498
type InlineObject498 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject498 instantiates a new InlineObject498 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject498() *InlineObject498 {
	this := InlineObject498{}
	return &this
}

// NewInlineObject498WithDefaults instantiates a new InlineObject498 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject498WithDefaults() *InlineObject498 {
	this := InlineObject498{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject498) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject498) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject498) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject498) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject498) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject498 struct {
	value *InlineObject498
	isSet bool
}

func (v NullableInlineObject498) Get() *InlineObject498 {
	return v.value
}

func (v *NullableInlineObject498) Set(val *InlineObject498) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject498) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject498) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject498(val *InlineObject498) *NullableInlineObject498 {
	return &NullableInlineObject498{value: val, isSet: true}
}

func (v NullableInlineObject498) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject498) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


