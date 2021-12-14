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

// CollectionOfTeamsAppDefinition struct for CollectionOfTeamsAppDefinition
type CollectionOfTeamsAppDefinition struct {
	Value *[]MicrosoftGraphTeamsAppDefinition `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfTeamsAppDefinition instantiates a new CollectionOfTeamsAppDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfTeamsAppDefinition() *CollectionOfTeamsAppDefinition {
	this := CollectionOfTeamsAppDefinition{}
	return &this
}

// NewCollectionOfTeamsAppDefinitionWithDefaults instantiates a new CollectionOfTeamsAppDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfTeamsAppDefinitionWithDefaults() *CollectionOfTeamsAppDefinition {
	this := CollectionOfTeamsAppDefinition{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfTeamsAppDefinition) GetValue() []MicrosoftGraphTeamsAppDefinition {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTeamsAppDefinition
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTeamsAppDefinition) GetValueOk() (*[]MicrosoftGraphTeamsAppDefinition, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfTeamsAppDefinition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTeamsAppDefinition and assigns it to the Value field.
func (o *CollectionOfTeamsAppDefinition) SetValue(v []MicrosoftGraphTeamsAppDefinition) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfTeamsAppDefinition) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTeamsAppDefinition) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfTeamsAppDefinition) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfTeamsAppDefinition) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfTeamsAppDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfTeamsAppDefinition struct {
	value *CollectionOfTeamsAppDefinition
	isSet bool
}

func (v NullableCollectionOfTeamsAppDefinition) Get() *CollectionOfTeamsAppDefinition {
	return v.value
}

func (v *NullableCollectionOfTeamsAppDefinition) Set(val *CollectionOfTeamsAppDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfTeamsAppDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfTeamsAppDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfTeamsAppDefinition(val *CollectionOfTeamsAppDefinition) *NullableCollectionOfTeamsAppDefinition {
	return &NullableCollectionOfTeamsAppDefinition{value: val, isSet: true}
}

func (v NullableCollectionOfTeamsAppDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfTeamsAppDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


