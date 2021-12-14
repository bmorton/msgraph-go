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

// CollectionOfExternalConnection struct for CollectionOfExternalConnection
type CollectionOfExternalConnection struct {
	Value *[]MicrosoftGraphExternalConnectorsExternalConnection `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfExternalConnection instantiates a new CollectionOfExternalConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfExternalConnection() *CollectionOfExternalConnection {
	this := CollectionOfExternalConnection{}
	return &this
}

// NewCollectionOfExternalConnectionWithDefaults instantiates a new CollectionOfExternalConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfExternalConnectionWithDefaults() *CollectionOfExternalConnection {
	this := CollectionOfExternalConnection{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfExternalConnection) GetValue() []MicrosoftGraphExternalConnectorsExternalConnection {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphExternalConnectorsExternalConnection
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfExternalConnection) GetValueOk() (*[]MicrosoftGraphExternalConnectorsExternalConnection, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfExternalConnection) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphExternalConnectorsExternalConnection and assigns it to the Value field.
func (o *CollectionOfExternalConnection) SetValue(v []MicrosoftGraphExternalConnectorsExternalConnection) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfExternalConnection) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfExternalConnection) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfExternalConnection) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfExternalConnection) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfExternalConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfExternalConnection struct {
	value *CollectionOfExternalConnection
	isSet bool
}

func (v NullableCollectionOfExternalConnection) Get() *CollectionOfExternalConnection {
	return v.value
}

func (v *NullableCollectionOfExternalConnection) Set(val *CollectionOfExternalConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfExternalConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfExternalConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfExternalConnection(val *CollectionOfExternalConnection) *NullableCollectionOfExternalConnection {
	return &NullableCollectionOfExternalConnection{value: val, isSet: true}
}

func (v NullableCollectionOfExternalConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfExternalConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


