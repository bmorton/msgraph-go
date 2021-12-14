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

// MicrosoftGraphInitiator struct for MicrosoftGraphInitiator
type MicrosoftGraphInitiator struct {
	// The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique identifier for the identity.
	Id NullableString `json:"id,omitempty"`
	// Type of initiator. Possible values are: user, application, system, unknownFutureValue.
	InitiatorType AnyOfmicrosoftGraphInitiatorType `json:"initiatorType,omitempty"`
}

// NewMicrosoftGraphInitiator instantiates a new MicrosoftGraphInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphInitiator() *MicrosoftGraphInitiator {
	this := MicrosoftGraphInitiator{}
	return &this
}

// NewMicrosoftGraphInitiatorWithDefaults instantiates a new MicrosoftGraphInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphInitiatorWithDefaults() *MicrosoftGraphInitiator {
	this := MicrosoftGraphInitiator{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInitiator) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInitiator) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphInitiator) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphInitiator) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphInitiator) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphInitiator) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInitiator) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInitiator) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphInitiator) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphInitiator) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphInitiator) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphInitiator) UnsetId() {
	o.Id.Unset()
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInitiator) GetInitiatorType() AnyOfmicrosoftGraphInitiatorType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphInitiatorType
		return ret
	}
	return o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInitiator) GetInitiatorTypeOk() (*AnyOfmicrosoftGraphInitiatorType, bool) {
	if o == nil || o.InitiatorType == nil {
		return nil, false
	}
	return &o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *MicrosoftGraphInitiator) HasInitiatorType() bool {
	if o != nil && o.InitiatorType != nil {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given AnyOfmicrosoftGraphInitiatorType and assigns it to the InitiatorType field.
func (o *MicrosoftGraphInitiator) SetInitiatorType(v AnyOfmicrosoftGraphInitiatorType) {
	o.InitiatorType = v
}

func (o MicrosoftGraphInitiator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InitiatorType != nil {
		toSerialize["initiatorType"] = o.InitiatorType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphInitiator struct {
	value *MicrosoftGraphInitiator
	isSet bool
}

func (v NullableMicrosoftGraphInitiator) Get() *MicrosoftGraphInitiator {
	return v.value
}

func (v *NullableMicrosoftGraphInitiator) Set(val *MicrosoftGraphInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphInitiator(val *MicrosoftGraphInitiator) *NullableMicrosoftGraphInitiator {
	return &NullableMicrosoftGraphInitiator{value: val, isSet: true}
}

func (v NullableMicrosoftGraphInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

