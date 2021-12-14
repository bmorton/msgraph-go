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

// CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus struct for CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus
type CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus struct {
	Value *[]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus instantiates a new CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus() *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus {
	this := CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus{}
	return &this
}

// NewCollectionOfManagedDeviceMobileAppConfigurationDeviceStatusWithDefaults instantiates a new CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfManagedDeviceMobileAppConfigurationDeviceStatusWithDefaults() *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus {
	this := CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) GetValue() []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) GetValueOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus and assigns it to the Value field.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) SetValue(v []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus struct {
	value *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus
	isSet bool
}

func (v NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) Get() *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus {
	return v.value
}

func (v *NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) Set(val *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus(val *CollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) *NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus {
	return &NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus{value: val, isSet: true}
}

func (v NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfManagedDeviceMobileAppConfigurationDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

