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

// CollectionOfChat struct for CollectionOfChat
type CollectionOfChat struct {
	Value *[]MicrosoftGraphChat `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfChat instantiates a new CollectionOfChat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfChat() *CollectionOfChat {
	this := CollectionOfChat{}
	return &this
}

// NewCollectionOfChatWithDefaults instantiates a new CollectionOfChat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfChatWithDefaults() *CollectionOfChat {
	this := CollectionOfChat{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfChat) GetValue() []MicrosoftGraphChat {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphChat
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChat) GetValueOk() (*[]MicrosoftGraphChat, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfChat) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphChat and assigns it to the Value field.
func (o *CollectionOfChat) SetValue(v []MicrosoftGraphChat) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfChat) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChat) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfChat) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfChat) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfChat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfChat struct {
	value *CollectionOfChat
	isSet bool
}

func (v NullableCollectionOfChat) Get() *CollectionOfChat {
	return v.value
}

func (v *NullableCollectionOfChat) Set(val *CollectionOfChat) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfChat) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfChat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfChat(val *CollectionOfChat) *NullableCollectionOfChat {
	return &NullableCollectionOfChat{value: val, isSet: true}
}

func (v NullableCollectionOfChat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfChat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

