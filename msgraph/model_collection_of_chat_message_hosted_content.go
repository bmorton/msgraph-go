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

// CollectionOfChatMessageHostedContent struct for CollectionOfChatMessageHostedContent
type CollectionOfChatMessageHostedContent struct {
	Value *[]MicrosoftGraphChatMessageHostedContent `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfChatMessageHostedContent instantiates a new CollectionOfChatMessageHostedContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfChatMessageHostedContent() *CollectionOfChatMessageHostedContent {
	this := CollectionOfChatMessageHostedContent{}
	return &this
}

// NewCollectionOfChatMessageHostedContentWithDefaults instantiates a new CollectionOfChatMessageHostedContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfChatMessageHostedContentWithDefaults() *CollectionOfChatMessageHostedContent {
	this := CollectionOfChatMessageHostedContent{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfChatMessageHostedContent) GetValue() []MicrosoftGraphChatMessageHostedContent {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphChatMessageHostedContent
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChatMessageHostedContent) GetValueOk() (*[]MicrosoftGraphChatMessageHostedContent, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfChatMessageHostedContent) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphChatMessageHostedContent and assigns it to the Value field.
func (o *CollectionOfChatMessageHostedContent) SetValue(v []MicrosoftGraphChatMessageHostedContent) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfChatMessageHostedContent) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfChatMessageHostedContent) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfChatMessageHostedContent) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfChatMessageHostedContent) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfChatMessageHostedContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfChatMessageHostedContent struct {
	value *CollectionOfChatMessageHostedContent
	isSet bool
}

func (v NullableCollectionOfChatMessageHostedContent) Get() *CollectionOfChatMessageHostedContent {
	return v.value
}

func (v *NullableCollectionOfChatMessageHostedContent) Set(val *CollectionOfChatMessageHostedContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfChatMessageHostedContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfChatMessageHostedContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfChatMessageHostedContent(val *CollectionOfChatMessageHostedContent) *NullableCollectionOfChatMessageHostedContent {
	return &NullableCollectionOfChatMessageHostedContent{value: val, isSet: true}
}

func (v NullableCollectionOfChatMessageHostedContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfChatMessageHostedContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

