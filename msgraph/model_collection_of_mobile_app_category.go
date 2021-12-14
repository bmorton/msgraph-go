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

// CollectionOfMobileAppCategory struct for CollectionOfMobileAppCategory
type CollectionOfMobileAppCategory struct {
	Value *[]MicrosoftGraphMobileAppCategory `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfMobileAppCategory instantiates a new CollectionOfMobileAppCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfMobileAppCategory() *CollectionOfMobileAppCategory {
	this := CollectionOfMobileAppCategory{}
	return &this
}

// NewCollectionOfMobileAppCategoryWithDefaults instantiates a new CollectionOfMobileAppCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfMobileAppCategoryWithDefaults() *CollectionOfMobileAppCategory {
	this := CollectionOfMobileAppCategory{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfMobileAppCategory) GetValue() []MicrosoftGraphMobileAppCategory {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphMobileAppCategory
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMobileAppCategory) GetValueOk() (*[]MicrosoftGraphMobileAppCategory, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfMobileAppCategory) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Value field.
func (o *CollectionOfMobileAppCategory) SetValue(v []MicrosoftGraphMobileAppCategory) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfMobileAppCategory) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMobileAppCategory) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfMobileAppCategory) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfMobileAppCategory) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfMobileAppCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfMobileAppCategory struct {
	value *CollectionOfMobileAppCategory
	isSet bool
}

func (v NullableCollectionOfMobileAppCategory) Get() *CollectionOfMobileAppCategory {
	return v.value
}

func (v *NullableCollectionOfMobileAppCategory) Set(val *CollectionOfMobileAppCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfMobileAppCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfMobileAppCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfMobileAppCategory(val *CollectionOfMobileAppCategory) *NullableCollectionOfMobileAppCategory {
	return &NullableCollectionOfMobileAppCategory{value: val, isSet: true}
}

func (v NullableCollectionOfMobileAppCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfMobileAppCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

