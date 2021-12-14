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

// CollectionOfRemoteAssistancePartner struct for CollectionOfRemoteAssistancePartner
type CollectionOfRemoteAssistancePartner struct {
	Value *[]MicrosoftGraphRemoteAssistancePartner `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfRemoteAssistancePartner instantiates a new CollectionOfRemoteAssistancePartner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfRemoteAssistancePartner() *CollectionOfRemoteAssistancePartner {
	this := CollectionOfRemoteAssistancePartner{}
	return &this
}

// NewCollectionOfRemoteAssistancePartnerWithDefaults instantiates a new CollectionOfRemoteAssistancePartner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfRemoteAssistancePartnerWithDefaults() *CollectionOfRemoteAssistancePartner {
	this := CollectionOfRemoteAssistancePartner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfRemoteAssistancePartner) GetValue() []MicrosoftGraphRemoteAssistancePartner {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphRemoteAssistancePartner
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfRemoteAssistancePartner) GetValueOk() (*[]MicrosoftGraphRemoteAssistancePartner, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfRemoteAssistancePartner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphRemoteAssistancePartner and assigns it to the Value field.
func (o *CollectionOfRemoteAssistancePartner) SetValue(v []MicrosoftGraphRemoteAssistancePartner) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfRemoteAssistancePartner) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfRemoteAssistancePartner) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfRemoteAssistancePartner) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfRemoteAssistancePartner) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfRemoteAssistancePartner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfRemoteAssistancePartner struct {
	value *CollectionOfRemoteAssistancePartner
	isSet bool
}

func (v NullableCollectionOfRemoteAssistancePartner) Get() *CollectionOfRemoteAssistancePartner {
	return v.value
}

func (v *NullableCollectionOfRemoteAssistancePartner) Set(val *CollectionOfRemoteAssistancePartner) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfRemoteAssistancePartner) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfRemoteAssistancePartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfRemoteAssistancePartner(val *CollectionOfRemoteAssistancePartner) *NullableCollectionOfRemoteAssistancePartner {
	return &NullableCollectionOfRemoteAssistancePartner{value: val, isSet: true}
}

func (v NullableCollectionOfRemoteAssistancePartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfRemoteAssistancePartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


