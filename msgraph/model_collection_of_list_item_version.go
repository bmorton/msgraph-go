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

// CollectionOfListItemVersion struct for CollectionOfListItemVersion
type CollectionOfListItemVersion struct {
	Value *[]MicrosoftGraphListItemVersion `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfListItemVersion instantiates a new CollectionOfListItemVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfListItemVersion() *CollectionOfListItemVersion {
	this := CollectionOfListItemVersion{}
	return &this
}

// NewCollectionOfListItemVersionWithDefaults instantiates a new CollectionOfListItemVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfListItemVersionWithDefaults() *CollectionOfListItemVersion {
	this := CollectionOfListItemVersion{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfListItemVersion) GetValue() []MicrosoftGraphListItemVersion {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphListItemVersion
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfListItemVersion) GetValueOk() (*[]MicrosoftGraphListItemVersion, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfListItemVersion) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphListItemVersion and assigns it to the Value field.
func (o *CollectionOfListItemVersion) SetValue(v []MicrosoftGraphListItemVersion) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfListItemVersion) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfListItemVersion) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfListItemVersion) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfListItemVersion) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfListItemVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfListItemVersion struct {
	value *CollectionOfListItemVersion
	isSet bool
}

func (v NullableCollectionOfListItemVersion) Get() *CollectionOfListItemVersion {
	return v.value
}

func (v *NullableCollectionOfListItemVersion) Set(val *CollectionOfListItemVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfListItemVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfListItemVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfListItemVersion(val *CollectionOfListItemVersion) *NullableCollectionOfListItemVersion {
	return &NullableCollectionOfListItemVersion{value: val, isSet: true}
}

func (v NullableCollectionOfListItemVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfListItemVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


