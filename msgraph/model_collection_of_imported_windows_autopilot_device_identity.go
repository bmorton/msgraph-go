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

// CollectionOfImportedWindowsAutopilotDeviceIdentity struct for CollectionOfImportedWindowsAutopilotDeviceIdentity
type CollectionOfImportedWindowsAutopilotDeviceIdentity struct {
	Value *[]MicrosoftGraphImportedWindowsAutopilotDeviceIdentity `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfImportedWindowsAutopilotDeviceIdentity instantiates a new CollectionOfImportedWindowsAutopilotDeviceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfImportedWindowsAutopilotDeviceIdentity() *CollectionOfImportedWindowsAutopilotDeviceIdentity {
	this := CollectionOfImportedWindowsAutopilotDeviceIdentity{}
	return &this
}

// NewCollectionOfImportedWindowsAutopilotDeviceIdentityWithDefaults instantiates a new CollectionOfImportedWindowsAutopilotDeviceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfImportedWindowsAutopilotDeviceIdentityWithDefaults() *CollectionOfImportedWindowsAutopilotDeviceIdentity {
	this := CollectionOfImportedWindowsAutopilotDeviceIdentity{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) GetValue() []MicrosoftGraphImportedWindowsAutopilotDeviceIdentity {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphImportedWindowsAutopilotDeviceIdentity
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) GetValueOk() (*[]MicrosoftGraphImportedWindowsAutopilotDeviceIdentity, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphImportedWindowsAutopilotDeviceIdentity and assigns it to the Value field.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) SetValue(v []MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfImportedWindowsAutopilotDeviceIdentity) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfImportedWindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfImportedWindowsAutopilotDeviceIdentity struct {
	value *CollectionOfImportedWindowsAutopilotDeviceIdentity
	isSet bool
}

func (v NullableCollectionOfImportedWindowsAutopilotDeviceIdentity) Get() *CollectionOfImportedWindowsAutopilotDeviceIdentity {
	return v.value
}

func (v *NullableCollectionOfImportedWindowsAutopilotDeviceIdentity) Set(val *CollectionOfImportedWindowsAutopilotDeviceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfImportedWindowsAutopilotDeviceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfImportedWindowsAutopilotDeviceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfImportedWindowsAutopilotDeviceIdentity(val *CollectionOfImportedWindowsAutopilotDeviceIdentity) *NullableCollectionOfImportedWindowsAutopilotDeviceIdentity {
	return &NullableCollectionOfImportedWindowsAutopilotDeviceIdentity{value: val, isSet: true}
}

func (v NullableCollectionOfImportedWindowsAutopilotDeviceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfImportedWindowsAutopilotDeviceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


