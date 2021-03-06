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

// CollectionOfOutlookCategory struct for CollectionOfOutlookCategory
type CollectionOfOutlookCategory struct {
	Value *[]MicrosoftGraphOutlookCategory `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOutlookCategory instantiates a new CollectionOfOutlookCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOutlookCategory() *CollectionOfOutlookCategory {
	this := CollectionOfOutlookCategory{}
	return &this
}

// NewCollectionOfOutlookCategoryWithDefaults instantiates a new CollectionOfOutlookCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOutlookCategoryWithDefaults() *CollectionOfOutlookCategory {
	this := CollectionOfOutlookCategory{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOutlookCategory) GetValue() []MicrosoftGraphOutlookCategory {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOutlookCategory
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOutlookCategory) GetValueOk() (*[]MicrosoftGraphOutlookCategory, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOutlookCategory) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOutlookCategory and assigns it to the Value field.
func (o *CollectionOfOutlookCategory) SetValue(v []MicrosoftGraphOutlookCategory) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOutlookCategory) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOutlookCategory) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOutlookCategory) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOutlookCategory) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOutlookCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOutlookCategory struct {
	value *CollectionOfOutlookCategory
	isSet bool
}

func (v NullableCollectionOfOutlookCategory) Get() *CollectionOfOutlookCategory {
	return v.value
}

func (v *NullableCollectionOfOutlookCategory) Set(val *CollectionOfOutlookCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOutlookCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOutlookCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOutlookCategory(val *CollectionOfOutlookCategory) *NullableCollectionOfOutlookCategory {
	return &NullableCollectionOfOutlookCategory{value: val, isSet: true}
}

func (v NullableCollectionOfOutlookCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOutlookCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


