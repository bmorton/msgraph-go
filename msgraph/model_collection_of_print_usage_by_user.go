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

// CollectionOfPrintUsageByUser struct for CollectionOfPrintUsageByUser
type CollectionOfPrintUsageByUser struct {
	Value *[]MicrosoftGraphPrintUsageByUser `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfPrintUsageByUser instantiates a new CollectionOfPrintUsageByUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfPrintUsageByUser() *CollectionOfPrintUsageByUser {
	this := CollectionOfPrintUsageByUser{}
	return &this
}

// NewCollectionOfPrintUsageByUserWithDefaults instantiates a new CollectionOfPrintUsageByUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfPrintUsageByUserWithDefaults() *CollectionOfPrintUsageByUser {
	this := CollectionOfPrintUsageByUser{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfPrintUsageByUser) GetValue() []MicrosoftGraphPrintUsageByUser {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphPrintUsageByUser
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPrintUsageByUser) GetValueOk() (*[]MicrosoftGraphPrintUsageByUser, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfPrintUsageByUser) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphPrintUsageByUser and assigns it to the Value field.
func (o *CollectionOfPrintUsageByUser) SetValue(v []MicrosoftGraphPrintUsageByUser) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfPrintUsageByUser) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPrintUsageByUser) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfPrintUsageByUser) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfPrintUsageByUser) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfPrintUsageByUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfPrintUsageByUser struct {
	value *CollectionOfPrintUsageByUser
	isSet bool
}

func (v NullableCollectionOfPrintUsageByUser) Get() *CollectionOfPrintUsageByUser {
	return v.value
}

func (v *NullableCollectionOfPrintUsageByUser) Set(val *CollectionOfPrintUsageByUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfPrintUsageByUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfPrintUsageByUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfPrintUsageByUser(val *CollectionOfPrintUsageByUser) *NullableCollectionOfPrintUsageByUser {
	return &NullableCollectionOfPrintUsageByUser{value: val, isSet: true}
}

func (v NullableCollectionOfPrintUsageByUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfPrintUsageByUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


