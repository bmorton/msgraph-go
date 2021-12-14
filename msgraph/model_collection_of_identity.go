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

// CollectionOfIdentity struct for CollectionOfIdentity
type CollectionOfIdentity struct {
	Value *[]MicrosoftGraphExternalConnectorsIdentity `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfIdentity instantiates a new CollectionOfIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfIdentity() *CollectionOfIdentity {
	this := CollectionOfIdentity{}
	return &this
}

// NewCollectionOfIdentityWithDefaults instantiates a new CollectionOfIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfIdentityWithDefaults() *CollectionOfIdentity {
	this := CollectionOfIdentity{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfIdentity) GetValue() []MicrosoftGraphExternalConnectorsIdentity {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphExternalConnectorsIdentity
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfIdentity) GetValueOk() (*[]MicrosoftGraphExternalConnectorsIdentity, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfIdentity) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphExternalConnectorsIdentity and assigns it to the Value field.
func (o *CollectionOfIdentity) SetValue(v []MicrosoftGraphExternalConnectorsIdentity) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfIdentity) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfIdentity) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfIdentity) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfIdentity) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfIdentity struct {
	value *CollectionOfIdentity
	isSet bool
}

func (v NullableCollectionOfIdentity) Get() *CollectionOfIdentity {
	return v.value
}

func (v *NullableCollectionOfIdentity) Set(val *CollectionOfIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfIdentity(val *CollectionOfIdentity) *NullableCollectionOfIdentity {
	return &NullableCollectionOfIdentity{value: val, isSet: true}
}

func (v NullableCollectionOfIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


