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

// MicrosoftGraphCallRoute struct for MicrosoftGraphCallRoute
type MicrosoftGraphCallRoute struct {
	Final *MicrosoftGraphIdentitySet `json:"final,omitempty"`
	Original *MicrosoftGraphIdentitySet `json:"original,omitempty"`
	// Possible values are: forwarded, lookup, selfFork.
	RoutingType AnyOfmicrosoftGraphRoutingType `json:"routingType,omitempty"`
}

// NewMicrosoftGraphCallRoute instantiates a new MicrosoftGraphCallRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphCallRoute() *MicrosoftGraphCallRoute {
	this := MicrosoftGraphCallRoute{}
	return &this
}

// NewMicrosoftGraphCallRouteWithDefaults instantiates a new MicrosoftGraphCallRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphCallRouteWithDefaults() *MicrosoftGraphCallRoute {
	this := MicrosoftGraphCallRoute{}
	return &this
}

// GetFinal returns the Final field value if set, zero value otherwise.
func (o *MicrosoftGraphCallRoute) GetFinal() MicrosoftGraphIdentitySet {
	if o == nil || o.Final == nil {
		var ret MicrosoftGraphIdentitySet
		return ret
	}
	return *o.Final
}

// GetFinalOk returns a tuple with the Final field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCallRoute) GetFinalOk() (*MicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Final == nil {
		return nil, false
	}
	return o.Final, true
}

// HasFinal returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRoute) HasFinal() bool {
	if o != nil && o.Final != nil {
		return true
	}

	return false
}

// SetFinal gets a reference to the given MicrosoftGraphIdentitySet and assigns it to the Final field.
func (o *MicrosoftGraphCallRoute) SetFinal(v MicrosoftGraphIdentitySet) {
	o.Final = &v
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *MicrosoftGraphCallRoute) GetOriginal() MicrosoftGraphIdentitySet {
	if o == nil || o.Original == nil {
		var ret MicrosoftGraphIdentitySet
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCallRoute) GetOriginalOk() (*MicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Original == nil {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRoute) HasOriginal() bool {
	if o != nil && o.Original != nil {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given MicrosoftGraphIdentitySet and assigns it to the Original field.
func (o *MicrosoftGraphCallRoute) SetOriginal(v MicrosoftGraphIdentitySet) {
	o.Original = &v
}

// GetRoutingType returns the RoutingType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRoute) GetRoutingType() AnyOfmicrosoftGraphRoutingType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphRoutingType
		return ret
	}
	return o.RoutingType
}

// GetRoutingTypeOk returns a tuple with the RoutingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRoute) GetRoutingTypeOk() (*AnyOfmicrosoftGraphRoutingType, bool) {
	if o == nil || o.RoutingType == nil {
		return nil, false
	}
	return &o.RoutingType, true
}

// HasRoutingType returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRoute) HasRoutingType() bool {
	if o != nil && o.RoutingType != nil {
		return true
	}

	return false
}

// SetRoutingType gets a reference to the given AnyOfmicrosoftGraphRoutingType and assigns it to the RoutingType field.
func (o *MicrosoftGraphCallRoute) SetRoutingType(v AnyOfmicrosoftGraphRoutingType) {
	o.RoutingType = v
}

func (o MicrosoftGraphCallRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Final != nil {
		toSerialize["final"] = o.Final
	}
	if o.Original != nil {
		toSerialize["original"] = o.Original
	}
	if o.RoutingType != nil {
		toSerialize["routingType"] = o.RoutingType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphCallRoute struct {
	value *MicrosoftGraphCallRoute
	isSet bool
}

func (v NullableMicrosoftGraphCallRoute) Get() *MicrosoftGraphCallRoute {
	return v.value
}

func (v *NullableMicrosoftGraphCallRoute) Set(val *MicrosoftGraphCallRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCallRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCallRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCallRoute(val *MicrosoftGraphCallRoute) *NullableMicrosoftGraphCallRoute {
	return &NullableMicrosoftGraphCallRoute{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCallRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCallRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

