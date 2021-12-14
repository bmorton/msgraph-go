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

// CollectionOfDeviceManagementExchangeConnector struct for CollectionOfDeviceManagementExchangeConnector
type CollectionOfDeviceManagementExchangeConnector struct {
	Value *[]MicrosoftGraphDeviceManagementExchangeConnector `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDeviceManagementExchangeConnector instantiates a new CollectionOfDeviceManagementExchangeConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDeviceManagementExchangeConnector() *CollectionOfDeviceManagementExchangeConnector {
	this := CollectionOfDeviceManagementExchangeConnector{}
	return &this
}

// NewCollectionOfDeviceManagementExchangeConnectorWithDefaults instantiates a new CollectionOfDeviceManagementExchangeConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceManagementExchangeConnectorWithDefaults() *CollectionOfDeviceManagementExchangeConnector {
	this := CollectionOfDeviceManagementExchangeConnector{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDeviceManagementExchangeConnector) GetValue() []MicrosoftGraphDeviceManagementExchangeConnector {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceManagementExchangeConnector
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceManagementExchangeConnector) GetValueOk() (*[]MicrosoftGraphDeviceManagementExchangeConnector, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceManagementExchangeConnector) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceManagementExchangeConnector and assigns it to the Value field.
func (o *CollectionOfDeviceManagementExchangeConnector) SetValue(v []MicrosoftGraphDeviceManagementExchangeConnector) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDeviceManagementExchangeConnector) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceManagementExchangeConnector) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDeviceManagementExchangeConnector) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDeviceManagementExchangeConnector) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDeviceManagementExchangeConnector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDeviceManagementExchangeConnector struct {
	value *CollectionOfDeviceManagementExchangeConnector
	isSet bool
}

func (v NullableCollectionOfDeviceManagementExchangeConnector) Get() *CollectionOfDeviceManagementExchangeConnector {
	return v.value
}

func (v *NullableCollectionOfDeviceManagementExchangeConnector) Set(val *CollectionOfDeviceManagementExchangeConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDeviceManagementExchangeConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDeviceManagementExchangeConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDeviceManagementExchangeConnector(val *CollectionOfDeviceManagementExchangeConnector) *NullableCollectionOfDeviceManagementExchangeConnector {
	return &NullableCollectionOfDeviceManagementExchangeConnector{value: val, isSet: true}
}

func (v NullableCollectionOfDeviceManagementExchangeConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDeviceManagementExchangeConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

