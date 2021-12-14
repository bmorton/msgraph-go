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

// CollectionOfSharedDriveItem struct for CollectionOfSharedDriveItem
type CollectionOfSharedDriveItem struct {
	Value *[]MicrosoftGraphSharedDriveItem `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfSharedDriveItem instantiates a new CollectionOfSharedDriveItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfSharedDriveItem() *CollectionOfSharedDriveItem {
	this := CollectionOfSharedDriveItem{}
	return &this
}

// NewCollectionOfSharedDriveItemWithDefaults instantiates a new CollectionOfSharedDriveItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfSharedDriveItemWithDefaults() *CollectionOfSharedDriveItem {
	this := CollectionOfSharedDriveItem{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfSharedDriveItem) GetValue() []MicrosoftGraphSharedDriveItem {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphSharedDriveItem
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSharedDriveItem) GetValueOk() (*[]MicrosoftGraphSharedDriveItem, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfSharedDriveItem) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphSharedDriveItem and assigns it to the Value field.
func (o *CollectionOfSharedDriveItem) SetValue(v []MicrosoftGraphSharedDriveItem) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfSharedDriveItem) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSharedDriveItem) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfSharedDriveItem) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfSharedDriveItem) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfSharedDriveItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfSharedDriveItem struct {
	value *CollectionOfSharedDriveItem
	isSet bool
}

func (v NullableCollectionOfSharedDriveItem) Get() *CollectionOfSharedDriveItem {
	return v.value
}

func (v *NullableCollectionOfSharedDriveItem) Set(val *CollectionOfSharedDriveItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfSharedDriveItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfSharedDriveItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfSharedDriveItem(val *CollectionOfSharedDriveItem) *NullableCollectionOfSharedDriveItem {
	return &NullableCollectionOfSharedDriveItem{value: val, isSet: true}
}

func (v NullableCollectionOfSharedDriveItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfSharedDriveItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


