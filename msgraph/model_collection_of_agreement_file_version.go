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

// CollectionOfAgreementFileVersion struct for CollectionOfAgreementFileVersion
type CollectionOfAgreementFileVersion struct {
	Value *[]MicrosoftGraphAgreementFileVersion `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfAgreementFileVersion instantiates a new CollectionOfAgreementFileVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfAgreementFileVersion() *CollectionOfAgreementFileVersion {
	this := CollectionOfAgreementFileVersion{}
	return &this
}

// NewCollectionOfAgreementFileVersionWithDefaults instantiates a new CollectionOfAgreementFileVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfAgreementFileVersionWithDefaults() *CollectionOfAgreementFileVersion {
	this := CollectionOfAgreementFileVersion{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfAgreementFileVersion) GetValue() []MicrosoftGraphAgreementFileVersion {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphAgreementFileVersion
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAgreementFileVersion) GetValueOk() (*[]MicrosoftGraphAgreementFileVersion, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfAgreementFileVersion) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphAgreementFileVersion and assigns it to the Value field.
func (o *CollectionOfAgreementFileVersion) SetValue(v []MicrosoftGraphAgreementFileVersion) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfAgreementFileVersion) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAgreementFileVersion) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfAgreementFileVersion) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfAgreementFileVersion) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfAgreementFileVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfAgreementFileVersion struct {
	value *CollectionOfAgreementFileVersion
	isSet bool
}

func (v NullableCollectionOfAgreementFileVersion) Get() *CollectionOfAgreementFileVersion {
	return v.value
}

func (v *NullableCollectionOfAgreementFileVersion) Set(val *CollectionOfAgreementFileVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfAgreementFileVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfAgreementFileVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfAgreementFileVersion(val *CollectionOfAgreementFileVersion) *NullableCollectionOfAgreementFileVersion {
	return &NullableCollectionOfAgreementFileVersion{value: val, isSet: true}
}

func (v NullableCollectionOfAgreementFileVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfAgreementFileVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


