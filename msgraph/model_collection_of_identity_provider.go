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

// CollectionOfIdentityProvider struct for CollectionOfIdentityProvider
type CollectionOfIdentityProvider struct {
	Value *[]MicrosoftGraphIdentityProvider `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfIdentityProvider instantiates a new CollectionOfIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfIdentityProvider() *CollectionOfIdentityProvider {
	this := CollectionOfIdentityProvider{}
	return &this
}

// NewCollectionOfIdentityProviderWithDefaults instantiates a new CollectionOfIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfIdentityProviderWithDefaults() *CollectionOfIdentityProvider {
	this := CollectionOfIdentityProvider{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfIdentityProvider) GetValue() []MicrosoftGraphIdentityProvider {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphIdentityProvider
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfIdentityProvider) GetValueOk() (*[]MicrosoftGraphIdentityProvider, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfIdentityProvider) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphIdentityProvider and assigns it to the Value field.
func (o *CollectionOfIdentityProvider) SetValue(v []MicrosoftGraphIdentityProvider) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfIdentityProvider) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfIdentityProvider) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfIdentityProvider) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfIdentityProvider) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfIdentityProvider struct {
	value *CollectionOfIdentityProvider
	isSet bool
}

func (v NullableCollectionOfIdentityProvider) Get() *CollectionOfIdentityProvider {
	return v.value
}

func (v *NullableCollectionOfIdentityProvider) Set(val *CollectionOfIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfIdentityProvider(val *CollectionOfIdentityProvider) *NullableCollectionOfIdentityProvider {
	return &NullableCollectionOfIdentityProvider{value: val, isSet: true}
}

func (v NullableCollectionOfIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

