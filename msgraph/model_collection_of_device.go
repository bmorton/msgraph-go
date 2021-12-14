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

// CollectionOfDevice struct for CollectionOfDevice
type CollectionOfDevice struct {
	Value *[]MicrosoftGraphDevice `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDevice instantiates a new CollectionOfDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDevice() *CollectionOfDevice {
	this := CollectionOfDevice{}
	return &this
}

// NewCollectionOfDeviceWithDefaults instantiates a new CollectionOfDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceWithDefaults() *CollectionOfDevice {
	this := CollectionOfDevice{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDevice) GetValue() []MicrosoftGraphDevice {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDevice
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDevice) GetValueOk() (*[]MicrosoftGraphDevice, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDevice) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDevice and assigns it to the Value field.
func (o *CollectionOfDevice) SetValue(v []MicrosoftGraphDevice) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDevice) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDevice) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDevice) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDevice) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDevice struct {
	value *CollectionOfDevice
	isSet bool
}

func (v NullableCollectionOfDevice) Get() *CollectionOfDevice {
	return v.value
}

func (v *NullableCollectionOfDevice) Set(val *CollectionOfDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDevice(val *CollectionOfDevice) *NullableCollectionOfDevice {
	return &NullableCollectionOfDevice{value: val, isSet: true}
}

func (v NullableCollectionOfDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


