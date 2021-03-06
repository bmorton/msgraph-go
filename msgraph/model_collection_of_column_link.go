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

// CollectionOfColumnLink struct for CollectionOfColumnLink
type CollectionOfColumnLink struct {
	Value *[]MicrosoftGraphColumnLink `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfColumnLink instantiates a new CollectionOfColumnLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfColumnLink() *CollectionOfColumnLink {
	this := CollectionOfColumnLink{}
	return &this
}

// NewCollectionOfColumnLinkWithDefaults instantiates a new CollectionOfColumnLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfColumnLinkWithDefaults() *CollectionOfColumnLink {
	this := CollectionOfColumnLink{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfColumnLink) GetValue() []MicrosoftGraphColumnLink {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphColumnLink
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfColumnLink) GetValueOk() (*[]MicrosoftGraphColumnLink, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfColumnLink) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphColumnLink and assigns it to the Value field.
func (o *CollectionOfColumnLink) SetValue(v []MicrosoftGraphColumnLink) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfColumnLink) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfColumnLink) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfColumnLink) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfColumnLink) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfColumnLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfColumnLink struct {
	value *CollectionOfColumnLink
	isSet bool
}

func (v NullableCollectionOfColumnLink) Get() *CollectionOfColumnLink {
	return v.value
}

func (v *NullableCollectionOfColumnLink) Set(val *CollectionOfColumnLink) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfColumnLink) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfColumnLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfColumnLink(val *CollectionOfColumnLink) *NullableCollectionOfColumnLink {
	return &NullableCollectionOfColumnLink{value: val, isSet: true}
}

func (v NullableCollectionOfColumnLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfColumnLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


