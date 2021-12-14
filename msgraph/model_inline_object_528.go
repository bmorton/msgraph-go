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

// InlineObject528 struct for InlineObject528
type InlineObject528 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
}

// NewInlineObject528 instantiates a new InlineObject528 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject528() *InlineObject528 {
	this := InlineObject528{}
	return &this
}

// NewInlineObject528WithDefaults instantiates a new InlineObject528 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject528WithDefaults() *InlineObject528 {
	this := InlineObject528{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject528) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject528) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject528) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject528) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

func (o InlineObject528) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject528 struct {
	value *InlineObject528
	isSet bool
}

func (v NullableInlineObject528) Get() *InlineObject528 {
	return v.value
}

func (v *NullableInlineObject528) Set(val *InlineObject528) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject528) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject528) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject528(val *InlineObject528) *NullableInlineObject528 {
	return &NullableInlineObject528{value: val, isSet: true}
}

func (v NullableInlineObject528) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject528) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


