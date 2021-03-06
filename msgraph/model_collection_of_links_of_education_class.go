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

// CollectionOfLinksOfEducationClass struct for CollectionOfLinksOfEducationClass
type CollectionOfLinksOfEducationClass struct {
	Value *[]string `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfLinksOfEducationClass instantiates a new CollectionOfLinksOfEducationClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfLinksOfEducationClass() *CollectionOfLinksOfEducationClass {
	this := CollectionOfLinksOfEducationClass{}
	return &this
}

// NewCollectionOfLinksOfEducationClassWithDefaults instantiates a new CollectionOfLinksOfEducationClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfLinksOfEducationClassWithDefaults() *CollectionOfLinksOfEducationClass {
	this := CollectionOfLinksOfEducationClass{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfLinksOfEducationClass) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfLinksOfEducationClass) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfLinksOfEducationClass) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *CollectionOfLinksOfEducationClass) SetValue(v []string) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfLinksOfEducationClass) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfLinksOfEducationClass) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfLinksOfEducationClass) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfLinksOfEducationClass) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfLinksOfEducationClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfLinksOfEducationClass struct {
	value *CollectionOfLinksOfEducationClass
	isSet bool
}

func (v NullableCollectionOfLinksOfEducationClass) Get() *CollectionOfLinksOfEducationClass {
	return v.value
}

func (v *NullableCollectionOfLinksOfEducationClass) Set(val *CollectionOfLinksOfEducationClass) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfLinksOfEducationClass) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfLinksOfEducationClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfLinksOfEducationClass(val *CollectionOfLinksOfEducationClass) *NullableCollectionOfLinksOfEducationClass {
	return &NullableCollectionOfLinksOfEducationClass{value: val, isSet: true}
}

func (v NullableCollectionOfLinksOfEducationClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfLinksOfEducationClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


