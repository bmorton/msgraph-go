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

// InlineObject1360 struct for InlineObject1360
type InlineObject1360 struct {
	LinkLocation AnyOfobject `json:"linkLocation,omitempty"`
	FriendlyName AnyOfobject `json:"friendlyName,omitempty"`
}

// NewInlineObject1360 instantiates a new InlineObject1360 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1360() *InlineObject1360 {
	this := InlineObject1360{}
	return &this
}

// NewInlineObject1360WithDefaults instantiates a new InlineObject1360 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1360WithDefaults() *InlineObject1360 {
	this := InlineObject1360{}
	return &this
}

// GetLinkLocation returns the LinkLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1360) GetLinkLocation() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.LinkLocation
}

// GetLinkLocationOk returns a tuple with the LinkLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1360) GetLinkLocationOk() (*AnyOfobject, bool) {
	if o == nil || o.LinkLocation == nil {
		return nil, false
	}
	return &o.LinkLocation, true
}

// HasLinkLocation returns a boolean if a field has been set.
func (o *InlineObject1360) HasLinkLocation() bool {
	if o != nil && o.LinkLocation != nil {
		return true
	}

	return false
}

// SetLinkLocation gets a reference to the given AnyOfobject and assigns it to the LinkLocation field.
func (o *InlineObject1360) SetLinkLocation(v AnyOfobject) {
	o.LinkLocation = v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1360) GetFriendlyName() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1360) GetFriendlyNameOk() (*AnyOfobject, bool) {
	if o == nil || o.FriendlyName == nil {
		return nil, false
	}
	return &o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *InlineObject1360) HasFriendlyName() bool {
	if o != nil && o.FriendlyName != nil {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given AnyOfobject and assigns it to the FriendlyName field.
func (o *InlineObject1360) SetFriendlyName(v AnyOfobject) {
	o.FriendlyName = v
}

func (o InlineObject1360) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkLocation != nil {
		toSerialize["linkLocation"] = o.LinkLocation
	}
	if o.FriendlyName != nil {
		toSerialize["friendlyName"] = o.FriendlyName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1360 struct {
	value *InlineObject1360
	isSet bool
}

func (v NullableInlineObject1360) Get() *InlineObject1360 {
	return v.value
}

func (v *NullableInlineObject1360) Set(val *InlineObject1360) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1360) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1360) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1360(val *InlineObject1360) *NullableInlineObject1360 {
	return &NullableInlineObject1360{value: val, isSet: true}
}

func (v NullableInlineObject1360) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1360) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


