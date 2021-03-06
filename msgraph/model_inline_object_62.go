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

// InlineObject62 struct for InlineObject62
type InlineObject62 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject62 instantiates a new InlineObject62 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject62() *InlineObject62 {
	this := InlineObject62{}
	return &this
}

// NewInlineObject62WithDefaults instantiates a new InlineObject62 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject62WithDefaults() *InlineObject62 {
	this := InlineObject62{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject62) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject62) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject62) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject62) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject62) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject62 struct {
	value *InlineObject62
	isSet bool
}

func (v NullableInlineObject62) Get() *InlineObject62 {
	return v.value
}

func (v *NullableInlineObject62) Set(val *InlineObject62) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject62) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject62) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject62(val *InlineObject62) *NullableInlineObject62 {
	return &NullableInlineObject62{value: val, isSet: true}
}

func (v NullableInlineObject62) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject62) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


