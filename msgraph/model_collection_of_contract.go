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

// CollectionOfContract struct for CollectionOfContract
type CollectionOfContract struct {
	Value *[]MicrosoftGraphContract `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfContract instantiates a new CollectionOfContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfContract() *CollectionOfContract {
	this := CollectionOfContract{}
	return &this
}

// NewCollectionOfContractWithDefaults instantiates a new CollectionOfContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfContractWithDefaults() *CollectionOfContract {
	this := CollectionOfContract{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfContract) GetValue() []MicrosoftGraphContract {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphContract
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfContract) GetValueOk() (*[]MicrosoftGraphContract, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfContract) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphContract and assigns it to the Value field.
func (o *CollectionOfContract) SetValue(v []MicrosoftGraphContract) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfContract) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfContract) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfContract) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfContract) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfContract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfContract struct {
	value *CollectionOfContract
	isSet bool
}

func (v NullableCollectionOfContract) Get() *CollectionOfContract {
	return v.value
}

func (v *NullableCollectionOfContract) Set(val *CollectionOfContract) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfContract) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfContract(val *CollectionOfContract) *NullableCollectionOfContract {
	return &NullableCollectionOfContract{value: val, isSet: true}
}

func (v NullableCollectionOfContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

