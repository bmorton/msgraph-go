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

// InlineObject521 struct for InlineObject521
type InlineObject521 struct {
	Apps *[]*AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
	AppGroupType AnyOfmicrosoftGraphTargetedManagedAppGroupType `json:"appGroupType,omitempty"`
}

// NewInlineObject521 instantiates a new InlineObject521 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject521() *InlineObject521 {
	this := InlineObject521{}
	return &this
}

// NewInlineObject521WithDefaults instantiates a new InlineObject521 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject521WithDefaults() *InlineObject521 {
	this := InlineObject521{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *InlineObject521) GetApps() []*AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []*AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject521) GetAppsOk() (*[]*AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject521) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []*AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject521) SetApps(v []*AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

// GetAppGroupType returns the AppGroupType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject521) GetAppGroupType() AnyOfmicrosoftGraphTargetedManagedAppGroupType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTargetedManagedAppGroupType
		return ret
	}
	return o.AppGroupType
}

// GetAppGroupTypeOk returns a tuple with the AppGroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject521) GetAppGroupTypeOk() (*AnyOfmicrosoftGraphTargetedManagedAppGroupType, bool) {
	if o == nil || o.AppGroupType == nil {
		return nil, false
	}
	return &o.AppGroupType, true
}

// HasAppGroupType returns a boolean if a field has been set.
func (o *InlineObject521) HasAppGroupType() bool {
	if o != nil && o.AppGroupType != nil {
		return true
	}

	return false
}

// SetAppGroupType gets a reference to the given AnyOfmicrosoftGraphTargetedManagedAppGroupType and assigns it to the AppGroupType field.
func (o *InlineObject521) SetAppGroupType(v AnyOfmicrosoftGraphTargetedManagedAppGroupType) {
	o.AppGroupType = v
}

func (o InlineObject521) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.AppGroupType != nil {
		toSerialize["appGroupType"] = o.AppGroupType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject521 struct {
	value *InlineObject521
	isSet bool
}

func (v NullableInlineObject521) Get() *InlineObject521 {
	return v.value
}

func (v *NullableInlineObject521) Set(val *InlineObject521) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject521) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject521) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject521(val *InlineObject521) *NullableInlineObject521 {
	return &NullableInlineObject521{value: val, isSet: true}
}

func (v NullableInlineObject521) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject521) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


