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

// CollectionOfColumnDefinition struct for CollectionOfColumnDefinition
type CollectionOfColumnDefinition struct {
	Value *[]MicrosoftGraphColumnDefinition `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfColumnDefinition instantiates a new CollectionOfColumnDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfColumnDefinition() *CollectionOfColumnDefinition {
	this := CollectionOfColumnDefinition{}
	return &this
}

// NewCollectionOfColumnDefinitionWithDefaults instantiates a new CollectionOfColumnDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfColumnDefinitionWithDefaults() *CollectionOfColumnDefinition {
	this := CollectionOfColumnDefinition{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfColumnDefinition) GetValue() []MicrosoftGraphColumnDefinition {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfColumnDefinition) GetValueOk() (*[]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfColumnDefinition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Value field.
func (o *CollectionOfColumnDefinition) SetValue(v []MicrosoftGraphColumnDefinition) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfColumnDefinition) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfColumnDefinition) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfColumnDefinition) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfColumnDefinition) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfColumnDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfColumnDefinition struct {
	value *CollectionOfColumnDefinition
	isSet bool
}

func (v NullableCollectionOfColumnDefinition) Get() *CollectionOfColumnDefinition {
	return v.value
}

func (v *NullableCollectionOfColumnDefinition) Set(val *CollectionOfColumnDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfColumnDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfColumnDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfColumnDefinition(val *CollectionOfColumnDefinition) *NullableCollectionOfColumnDefinition {
	return &NullableCollectionOfColumnDefinition{value: val, isSet: true}
}

func (v NullableCollectionOfColumnDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfColumnDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

