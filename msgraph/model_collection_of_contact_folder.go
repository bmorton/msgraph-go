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

// CollectionOfContactFolder struct for CollectionOfContactFolder
type CollectionOfContactFolder struct {
	Value *[]MicrosoftGraphContactFolder `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfContactFolder instantiates a new CollectionOfContactFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfContactFolder() *CollectionOfContactFolder {
	this := CollectionOfContactFolder{}
	return &this
}

// NewCollectionOfContactFolderWithDefaults instantiates a new CollectionOfContactFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfContactFolderWithDefaults() *CollectionOfContactFolder {
	this := CollectionOfContactFolder{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfContactFolder) GetValue() []MicrosoftGraphContactFolder {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphContactFolder
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfContactFolder) GetValueOk() (*[]MicrosoftGraphContactFolder, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfContactFolder) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphContactFolder and assigns it to the Value field.
func (o *CollectionOfContactFolder) SetValue(v []MicrosoftGraphContactFolder) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfContactFolder) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfContactFolder) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfContactFolder) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfContactFolder) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfContactFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfContactFolder struct {
	value *CollectionOfContactFolder
	isSet bool
}

func (v NullableCollectionOfContactFolder) Get() *CollectionOfContactFolder {
	return v.value
}

func (v *NullableCollectionOfContactFolder) Set(val *CollectionOfContactFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfContactFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfContactFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfContactFolder(val *CollectionOfContactFolder) *NullableCollectionOfContactFolder {
	return &NullableCollectionOfContactFolder{value: val, isSet: true}
}

func (v NullableCollectionOfContactFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfContactFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


