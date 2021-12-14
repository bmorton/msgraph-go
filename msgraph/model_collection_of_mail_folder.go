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

// CollectionOfMailFolder struct for CollectionOfMailFolder
type CollectionOfMailFolder struct {
	Value *[]MicrosoftGraphMailFolder `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfMailFolder instantiates a new CollectionOfMailFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfMailFolder() *CollectionOfMailFolder {
	this := CollectionOfMailFolder{}
	return &this
}

// NewCollectionOfMailFolderWithDefaults instantiates a new CollectionOfMailFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfMailFolderWithDefaults() *CollectionOfMailFolder {
	this := CollectionOfMailFolder{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfMailFolder) GetValue() []MicrosoftGraphMailFolder {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphMailFolder
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMailFolder) GetValueOk() (*[]MicrosoftGraphMailFolder, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfMailFolder) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphMailFolder and assigns it to the Value field.
func (o *CollectionOfMailFolder) SetValue(v []MicrosoftGraphMailFolder) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfMailFolder) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMailFolder) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfMailFolder) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfMailFolder) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfMailFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfMailFolder struct {
	value *CollectionOfMailFolder
	isSet bool
}

func (v NullableCollectionOfMailFolder) Get() *CollectionOfMailFolder {
	return v.value
}

func (v *NullableCollectionOfMailFolder) Set(val *CollectionOfMailFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfMailFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfMailFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfMailFolder(val *CollectionOfMailFolder) *NullableCollectionOfMailFolder {
	return &NullableCollectionOfMailFolder{value: val, isSet: true}
}

func (v NullableCollectionOfMailFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfMailFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

