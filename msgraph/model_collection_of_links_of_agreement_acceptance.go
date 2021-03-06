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

// CollectionOfLinksOfAgreementAcceptance struct for CollectionOfLinksOfAgreementAcceptance
type CollectionOfLinksOfAgreementAcceptance struct {
	Value *[]string `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfLinksOfAgreementAcceptance instantiates a new CollectionOfLinksOfAgreementAcceptance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfLinksOfAgreementAcceptance() *CollectionOfLinksOfAgreementAcceptance {
	this := CollectionOfLinksOfAgreementAcceptance{}
	return &this
}

// NewCollectionOfLinksOfAgreementAcceptanceWithDefaults instantiates a new CollectionOfLinksOfAgreementAcceptance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfLinksOfAgreementAcceptanceWithDefaults() *CollectionOfLinksOfAgreementAcceptance {
	this := CollectionOfLinksOfAgreementAcceptance{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfLinksOfAgreementAcceptance) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfLinksOfAgreementAcceptance) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfLinksOfAgreementAcceptance) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *CollectionOfLinksOfAgreementAcceptance) SetValue(v []string) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfLinksOfAgreementAcceptance) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfLinksOfAgreementAcceptance) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfLinksOfAgreementAcceptance) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfLinksOfAgreementAcceptance) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfLinksOfAgreementAcceptance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfLinksOfAgreementAcceptance struct {
	value *CollectionOfLinksOfAgreementAcceptance
	isSet bool
}

func (v NullableCollectionOfLinksOfAgreementAcceptance) Get() *CollectionOfLinksOfAgreementAcceptance {
	return v.value
}

func (v *NullableCollectionOfLinksOfAgreementAcceptance) Set(val *CollectionOfLinksOfAgreementAcceptance) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfLinksOfAgreementAcceptance) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfLinksOfAgreementAcceptance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfLinksOfAgreementAcceptance(val *CollectionOfLinksOfAgreementAcceptance) *NullableCollectionOfLinksOfAgreementAcceptance {
	return &NullableCollectionOfLinksOfAgreementAcceptance{value: val, isSet: true}
}

func (v NullableCollectionOfLinksOfAgreementAcceptance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfLinksOfAgreementAcceptance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


