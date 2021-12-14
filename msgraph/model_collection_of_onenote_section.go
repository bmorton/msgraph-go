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

// CollectionOfOnenoteSection struct for CollectionOfOnenoteSection
type CollectionOfOnenoteSection struct {
	Value *[]MicrosoftGraphOnenoteSection `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOnenoteSection instantiates a new CollectionOfOnenoteSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOnenoteSection() *CollectionOfOnenoteSection {
	this := CollectionOfOnenoteSection{}
	return &this
}

// NewCollectionOfOnenoteSectionWithDefaults instantiates a new CollectionOfOnenoteSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOnenoteSectionWithDefaults() *CollectionOfOnenoteSection {
	this := CollectionOfOnenoteSection{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOnenoteSection) GetValue() []MicrosoftGraphOnenoteSection {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOnenoteSection
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOnenoteSection) GetValueOk() (*[]MicrosoftGraphOnenoteSection, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOnenoteSection) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Value field.
func (o *CollectionOfOnenoteSection) SetValue(v []MicrosoftGraphOnenoteSection) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOnenoteSection) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOnenoteSection) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOnenoteSection) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOnenoteSection) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOnenoteSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOnenoteSection struct {
	value *CollectionOfOnenoteSection
	isSet bool
}

func (v NullableCollectionOfOnenoteSection) Get() *CollectionOfOnenoteSection {
	return v.value
}

func (v *NullableCollectionOfOnenoteSection) Set(val *CollectionOfOnenoteSection) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOnenoteSection) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOnenoteSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOnenoteSection(val *CollectionOfOnenoteSection) *NullableCollectionOfOnenoteSection {
	return &NullableCollectionOfOnenoteSection{value: val, isSet: true}
}

func (v NullableCollectionOfOnenoteSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOnenoteSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


