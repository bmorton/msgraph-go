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

// InlineObject1022 struct for InlineObject1022
type InlineObject1022 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject1022 instantiates a new InlineObject1022 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1022() *InlineObject1022 {
	this := InlineObject1022{}
	return &this
}

// NewInlineObject1022WithDefaults instantiates a new InlineObject1022 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1022WithDefaults() *InlineObject1022 {
	this := InlineObject1022{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject1022) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1022) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject1022) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject1022) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject1022) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1022 struct {
	value *InlineObject1022
	isSet bool
}

func (v NullableInlineObject1022) Get() *InlineObject1022 {
	return v.value
}

func (v *NullableInlineObject1022) Set(val *InlineObject1022) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1022) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1022) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1022(val *InlineObject1022) *NullableInlineObject1022 {
	return &NullableInlineObject1022{value: val, isSet: true}
}

func (v NullableInlineObject1022) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1022) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


