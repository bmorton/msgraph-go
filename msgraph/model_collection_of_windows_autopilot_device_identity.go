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

// CollectionOfWindowsAutopilotDeviceIdentity struct for CollectionOfWindowsAutopilotDeviceIdentity
type CollectionOfWindowsAutopilotDeviceIdentity struct {
	Value *[]MicrosoftGraphWindowsAutopilotDeviceIdentity `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfWindowsAutopilotDeviceIdentity instantiates a new CollectionOfWindowsAutopilotDeviceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfWindowsAutopilotDeviceIdentity() *CollectionOfWindowsAutopilotDeviceIdentity {
	this := CollectionOfWindowsAutopilotDeviceIdentity{}
	return &this
}

// NewCollectionOfWindowsAutopilotDeviceIdentityWithDefaults instantiates a new CollectionOfWindowsAutopilotDeviceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfWindowsAutopilotDeviceIdentityWithDefaults() *CollectionOfWindowsAutopilotDeviceIdentity {
	this := CollectionOfWindowsAutopilotDeviceIdentity{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) GetValue() []MicrosoftGraphWindowsAutopilotDeviceIdentity {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphWindowsAutopilotDeviceIdentity
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) GetValueOk() (*[]MicrosoftGraphWindowsAutopilotDeviceIdentity, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphWindowsAutopilotDeviceIdentity and assigns it to the Value field.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) SetValue(v []MicrosoftGraphWindowsAutopilotDeviceIdentity) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfWindowsAutopilotDeviceIdentity) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfWindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfWindowsAutopilotDeviceIdentity struct {
	value *CollectionOfWindowsAutopilotDeviceIdentity
	isSet bool
}

func (v NullableCollectionOfWindowsAutopilotDeviceIdentity) Get() *CollectionOfWindowsAutopilotDeviceIdentity {
	return v.value
}

func (v *NullableCollectionOfWindowsAutopilotDeviceIdentity) Set(val *CollectionOfWindowsAutopilotDeviceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfWindowsAutopilotDeviceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfWindowsAutopilotDeviceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfWindowsAutopilotDeviceIdentity(val *CollectionOfWindowsAutopilotDeviceIdentity) *NullableCollectionOfWindowsAutopilotDeviceIdentity {
	return &NullableCollectionOfWindowsAutopilotDeviceIdentity{value: val, isSet: true}
}

func (v NullableCollectionOfWindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfWindowsAutopilotDeviceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


