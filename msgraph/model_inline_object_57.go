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

// InlineObject57 struct for InlineObject57
type InlineObject57 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject57 instantiates a new InlineObject57 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject57() *InlineObject57 {
	this := InlineObject57{}
	return &this
}

// NewInlineObject57WithDefaults instantiates a new InlineObject57 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject57WithDefaults() *InlineObject57 {
	this := InlineObject57{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject57) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject57) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject57) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject57) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject57) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject57 struct {
	value *InlineObject57
	isSet bool
}

func (v NullableInlineObject57) Get() *InlineObject57 {
	return v.value
}

func (v *NullableInlineObject57) Set(val *InlineObject57) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject57) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject57) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject57(val *InlineObject57) *NullableInlineObject57 {
	return &NullableInlineObject57{value: val, isSet: true}
}

func (v NullableInlineObject57) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject57) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


