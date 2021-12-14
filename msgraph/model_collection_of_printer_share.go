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

// CollectionOfPrinterShare struct for CollectionOfPrinterShare
type CollectionOfPrinterShare struct {
	Value *[]MicrosoftGraphPrinterShare `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfPrinterShare instantiates a new CollectionOfPrinterShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfPrinterShare() *CollectionOfPrinterShare {
	this := CollectionOfPrinterShare{}
	return &this
}

// NewCollectionOfPrinterShareWithDefaults instantiates a new CollectionOfPrinterShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfPrinterShareWithDefaults() *CollectionOfPrinterShare {
	this := CollectionOfPrinterShare{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfPrinterShare) GetValue() []MicrosoftGraphPrinterShare {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphPrinterShare
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPrinterShare) GetValueOk() (*[]MicrosoftGraphPrinterShare, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfPrinterShare) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphPrinterShare and assigns it to the Value field.
func (o *CollectionOfPrinterShare) SetValue(v []MicrosoftGraphPrinterShare) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfPrinterShare) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPrinterShare) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfPrinterShare) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfPrinterShare) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfPrinterShare) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfPrinterShare struct {
	value *CollectionOfPrinterShare
	isSet bool
}

func (v NullableCollectionOfPrinterShare) Get() *CollectionOfPrinterShare {
	return v.value
}

func (v *NullableCollectionOfPrinterShare) Set(val *CollectionOfPrinterShare) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfPrinterShare) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfPrinterShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfPrinterShare(val *CollectionOfPrinterShare) *NullableCollectionOfPrinterShare {
	return &NullableCollectionOfPrinterShare{value: val, isSet: true}
}

func (v NullableCollectionOfPrinterShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfPrinterShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


