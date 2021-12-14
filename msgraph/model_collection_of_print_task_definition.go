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

// CollectionOfPrintTaskDefinition struct for CollectionOfPrintTaskDefinition
type CollectionOfPrintTaskDefinition struct {
	Value *[]MicrosoftGraphPrintTaskDefinition `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfPrintTaskDefinition instantiates a new CollectionOfPrintTaskDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfPrintTaskDefinition() *CollectionOfPrintTaskDefinition {
	this := CollectionOfPrintTaskDefinition{}
	return &this
}

// NewCollectionOfPrintTaskDefinitionWithDefaults instantiates a new CollectionOfPrintTaskDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfPrintTaskDefinitionWithDefaults() *CollectionOfPrintTaskDefinition {
	this := CollectionOfPrintTaskDefinition{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfPrintTaskDefinition) GetValue() []MicrosoftGraphPrintTaskDefinition {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphPrintTaskDefinition
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPrintTaskDefinition) GetValueOk() (*[]MicrosoftGraphPrintTaskDefinition, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfPrintTaskDefinition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphPrintTaskDefinition and assigns it to the Value field.
func (o *CollectionOfPrintTaskDefinition) SetValue(v []MicrosoftGraphPrintTaskDefinition) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfPrintTaskDefinition) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPrintTaskDefinition) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfPrintTaskDefinition) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfPrintTaskDefinition) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfPrintTaskDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfPrintTaskDefinition struct {
	value *CollectionOfPrintTaskDefinition
	isSet bool
}

func (v NullableCollectionOfPrintTaskDefinition) Get() *CollectionOfPrintTaskDefinition {
	return v.value
}

func (v *NullableCollectionOfPrintTaskDefinition) Set(val *CollectionOfPrintTaskDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfPrintTaskDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfPrintTaskDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfPrintTaskDefinition(val *CollectionOfPrintTaskDefinition) *NullableCollectionOfPrintTaskDefinition {
	return &NullableCollectionOfPrintTaskDefinition{value: val, isSet: true}
}

func (v NullableCollectionOfPrintTaskDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfPrintTaskDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

