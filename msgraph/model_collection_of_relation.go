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

// CollectionOfRelation struct for CollectionOfRelation
type CollectionOfRelation struct {
	Value *[]MicrosoftGraphTermStoreRelation `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfRelation instantiates a new CollectionOfRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfRelation() *CollectionOfRelation {
	this := CollectionOfRelation{}
	return &this
}

// NewCollectionOfRelationWithDefaults instantiates a new CollectionOfRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfRelationWithDefaults() *CollectionOfRelation {
	this := CollectionOfRelation{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfRelation) GetValue() []MicrosoftGraphTermStoreRelation {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTermStoreRelation
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfRelation) GetValueOk() (*[]MicrosoftGraphTermStoreRelation, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfRelation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTermStoreRelation and assigns it to the Value field.
func (o *CollectionOfRelation) SetValue(v []MicrosoftGraphTermStoreRelation) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfRelation) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfRelation) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfRelation) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfRelation) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfRelation struct {
	value *CollectionOfRelation
	isSet bool
}

func (v NullableCollectionOfRelation) Get() *CollectionOfRelation {
	return v.value
}

func (v *NullableCollectionOfRelation) Set(val *CollectionOfRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfRelation(val *CollectionOfRelation) *NullableCollectionOfRelation {
	return &NullableCollectionOfRelation{value: val, isSet: true}
}

func (v NullableCollectionOfRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


