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

// CollectionOfOrganization struct for CollectionOfOrganization
type CollectionOfOrganization struct {
	Value *[]MicrosoftGraphOrganization `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOrganization instantiates a new CollectionOfOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOrganization() *CollectionOfOrganization {
	this := CollectionOfOrganization{}
	return &this
}

// NewCollectionOfOrganizationWithDefaults instantiates a new CollectionOfOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOrganizationWithDefaults() *CollectionOfOrganization {
	this := CollectionOfOrganization{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOrganization) GetValue() []MicrosoftGraphOrganization {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOrganization
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOrganization) GetValueOk() (*[]MicrosoftGraphOrganization, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOrganization) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOrganization and assigns it to the Value field.
func (o *CollectionOfOrganization) SetValue(v []MicrosoftGraphOrganization) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOrganization) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOrganization) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOrganization) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOrganization) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOrganization struct {
	value *CollectionOfOrganization
	isSet bool
}

func (v NullableCollectionOfOrganization) Get() *CollectionOfOrganization {
	return v.value
}

func (v *NullableCollectionOfOrganization) Set(val *CollectionOfOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOrganization(val *CollectionOfOrganization) *NullableCollectionOfOrganization {
	return &NullableCollectionOfOrganization{value: val, isSet: true}
}

func (v NullableCollectionOfOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

