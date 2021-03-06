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

// CollectionOfGroup1 struct for CollectionOfGroup1
type CollectionOfGroup1 struct {
	Value *[]MicrosoftGraphTermStoreGroup `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfGroup1 instantiates a new CollectionOfGroup1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfGroup1() *CollectionOfGroup1 {
	this := CollectionOfGroup1{}
	return &this
}

// NewCollectionOfGroup1WithDefaults instantiates a new CollectionOfGroup1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfGroup1WithDefaults() *CollectionOfGroup1 {
	this := CollectionOfGroup1{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfGroup1) GetValue() []MicrosoftGraphTermStoreGroup {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTermStoreGroup
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroup1) GetValueOk() (*[]MicrosoftGraphTermStoreGroup, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfGroup1) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTermStoreGroup and assigns it to the Value field.
func (o *CollectionOfGroup1) SetValue(v []MicrosoftGraphTermStoreGroup) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfGroup1) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroup1) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfGroup1) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfGroup1) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfGroup1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfGroup1 struct {
	value *CollectionOfGroup1
	isSet bool
}

func (v NullableCollectionOfGroup1) Get() *CollectionOfGroup1 {
	return v.value
}

func (v *NullableCollectionOfGroup1) Set(val *CollectionOfGroup1) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfGroup1) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfGroup1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfGroup1(val *CollectionOfGroup1) *NullableCollectionOfGroup1 {
	return &NullableCollectionOfGroup1{value: val, isSet: true}
}

func (v NullableCollectionOfGroup1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfGroup1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


