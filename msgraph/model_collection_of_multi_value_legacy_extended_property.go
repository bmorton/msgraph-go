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

// CollectionOfMultiValueLegacyExtendedProperty struct for CollectionOfMultiValueLegacyExtendedProperty
type CollectionOfMultiValueLegacyExtendedProperty struct {
	Value *[]MicrosoftGraphMultiValueLegacyExtendedProperty `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfMultiValueLegacyExtendedProperty instantiates a new CollectionOfMultiValueLegacyExtendedProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfMultiValueLegacyExtendedProperty() *CollectionOfMultiValueLegacyExtendedProperty {
	this := CollectionOfMultiValueLegacyExtendedProperty{}
	return &this
}

// NewCollectionOfMultiValueLegacyExtendedPropertyWithDefaults instantiates a new CollectionOfMultiValueLegacyExtendedProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfMultiValueLegacyExtendedPropertyWithDefaults() *CollectionOfMultiValueLegacyExtendedProperty {
	this := CollectionOfMultiValueLegacyExtendedProperty{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfMultiValueLegacyExtendedProperty) GetValue() []MicrosoftGraphMultiValueLegacyExtendedProperty {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphMultiValueLegacyExtendedProperty
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMultiValueLegacyExtendedProperty) GetValueOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfMultiValueLegacyExtendedProperty) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the Value field.
func (o *CollectionOfMultiValueLegacyExtendedProperty) SetValue(v []MicrosoftGraphMultiValueLegacyExtendedProperty) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfMultiValueLegacyExtendedProperty) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMultiValueLegacyExtendedProperty) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfMultiValueLegacyExtendedProperty) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfMultiValueLegacyExtendedProperty) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfMultiValueLegacyExtendedProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfMultiValueLegacyExtendedProperty struct {
	value *CollectionOfMultiValueLegacyExtendedProperty
	isSet bool
}

func (v NullableCollectionOfMultiValueLegacyExtendedProperty) Get() *CollectionOfMultiValueLegacyExtendedProperty {
	return v.value
}

func (v *NullableCollectionOfMultiValueLegacyExtendedProperty) Set(val *CollectionOfMultiValueLegacyExtendedProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfMultiValueLegacyExtendedProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfMultiValueLegacyExtendedProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfMultiValueLegacyExtendedProperty(val *CollectionOfMultiValueLegacyExtendedProperty) *NullableCollectionOfMultiValueLegacyExtendedProperty {
	return &NullableCollectionOfMultiValueLegacyExtendedProperty{value: val, isSet: true}
}

func (v NullableCollectionOfMultiValueLegacyExtendedProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfMultiValueLegacyExtendedProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


