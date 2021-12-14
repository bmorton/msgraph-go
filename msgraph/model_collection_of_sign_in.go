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

// CollectionOfSignIn struct for CollectionOfSignIn
type CollectionOfSignIn struct {
	Value *[]MicrosoftGraphSignIn `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfSignIn instantiates a new CollectionOfSignIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfSignIn() *CollectionOfSignIn {
	this := CollectionOfSignIn{}
	return &this
}

// NewCollectionOfSignInWithDefaults instantiates a new CollectionOfSignIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfSignInWithDefaults() *CollectionOfSignIn {
	this := CollectionOfSignIn{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfSignIn) GetValue() []MicrosoftGraphSignIn {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphSignIn
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSignIn) GetValueOk() (*[]MicrosoftGraphSignIn, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfSignIn) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphSignIn and assigns it to the Value field.
func (o *CollectionOfSignIn) SetValue(v []MicrosoftGraphSignIn) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfSignIn) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSignIn) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfSignIn) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfSignIn) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfSignIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfSignIn struct {
	value *CollectionOfSignIn
	isSet bool
}

func (v NullableCollectionOfSignIn) Get() *CollectionOfSignIn {
	return v.value
}

func (v *NullableCollectionOfSignIn) Set(val *CollectionOfSignIn) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfSignIn) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfSignIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfSignIn(val *CollectionOfSignIn) *NullableCollectionOfSignIn {
	return &NullableCollectionOfSignIn{value: val, isSet: true}
}

func (v NullableCollectionOfSignIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfSignIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

