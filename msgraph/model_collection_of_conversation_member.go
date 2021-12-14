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

// CollectionOfConversationMember struct for CollectionOfConversationMember
type CollectionOfConversationMember struct {
	Value *[]MicrosoftGraphConversationMember `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfConversationMember instantiates a new CollectionOfConversationMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfConversationMember() *CollectionOfConversationMember {
	this := CollectionOfConversationMember{}
	return &this
}

// NewCollectionOfConversationMemberWithDefaults instantiates a new CollectionOfConversationMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfConversationMemberWithDefaults() *CollectionOfConversationMember {
	this := CollectionOfConversationMember{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfConversationMember) GetValue() []MicrosoftGraphConversationMember {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphConversationMember
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConversationMember) GetValueOk() (*[]MicrosoftGraphConversationMember, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfConversationMember) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphConversationMember and assigns it to the Value field.
func (o *CollectionOfConversationMember) SetValue(v []MicrosoftGraphConversationMember) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfConversationMember) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConversationMember) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfConversationMember) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfConversationMember) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfConversationMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfConversationMember struct {
	value *CollectionOfConversationMember
	isSet bool
}

func (v NullableCollectionOfConversationMember) Get() *CollectionOfConversationMember {
	return v.value
}

func (v *NullableCollectionOfConversationMember) Set(val *CollectionOfConversationMember) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfConversationMember) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfConversationMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfConversationMember(val *CollectionOfConversationMember) *NullableCollectionOfConversationMember {
	return &NullableCollectionOfConversationMember{value: val, isSet: true}
}

func (v NullableCollectionOfConversationMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfConversationMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


