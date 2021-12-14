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

// CollectionOfOrganizationalBrandingLocalization struct for CollectionOfOrganizationalBrandingLocalization
type CollectionOfOrganizationalBrandingLocalization struct {
	Value *[]MicrosoftGraphOrganizationalBrandingLocalization `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOrganizationalBrandingLocalization instantiates a new CollectionOfOrganizationalBrandingLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOrganizationalBrandingLocalization() *CollectionOfOrganizationalBrandingLocalization {
	this := CollectionOfOrganizationalBrandingLocalization{}
	return &this
}

// NewCollectionOfOrganizationalBrandingLocalizationWithDefaults instantiates a new CollectionOfOrganizationalBrandingLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOrganizationalBrandingLocalizationWithDefaults() *CollectionOfOrganizationalBrandingLocalization {
	this := CollectionOfOrganizationalBrandingLocalization{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOrganizationalBrandingLocalization) GetValue() []MicrosoftGraphOrganizationalBrandingLocalization {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOrganizationalBrandingLocalization
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOrganizationalBrandingLocalization) GetValueOk() (*[]MicrosoftGraphOrganizationalBrandingLocalization, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOrganizationalBrandingLocalization) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOrganizationalBrandingLocalization and assigns it to the Value field.
func (o *CollectionOfOrganizationalBrandingLocalization) SetValue(v []MicrosoftGraphOrganizationalBrandingLocalization) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOrganizationalBrandingLocalization) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOrganizationalBrandingLocalization) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOrganizationalBrandingLocalization) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOrganizationalBrandingLocalization) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOrganizationalBrandingLocalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOrganizationalBrandingLocalization struct {
	value *CollectionOfOrganizationalBrandingLocalization
	isSet bool
}

func (v NullableCollectionOfOrganizationalBrandingLocalization) Get() *CollectionOfOrganizationalBrandingLocalization {
	return v.value
}

func (v *NullableCollectionOfOrganizationalBrandingLocalization) Set(val *CollectionOfOrganizationalBrandingLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOrganizationalBrandingLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOrganizationalBrandingLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOrganizationalBrandingLocalization(val *CollectionOfOrganizationalBrandingLocalization) *NullableCollectionOfOrganizationalBrandingLocalization {
	return &NullableCollectionOfOrganizationalBrandingLocalization{value: val, isSet: true}
}

func (v NullableCollectionOfOrganizationalBrandingLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOrganizationalBrandingLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

