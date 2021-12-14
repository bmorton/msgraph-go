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

// CollectionOfTargetedManagedAppConfiguration struct for CollectionOfTargetedManagedAppConfiguration
type CollectionOfTargetedManagedAppConfiguration struct {
	Value *[]MicrosoftGraphTargetedManagedAppConfiguration `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfTargetedManagedAppConfiguration instantiates a new CollectionOfTargetedManagedAppConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfTargetedManagedAppConfiguration() *CollectionOfTargetedManagedAppConfiguration {
	this := CollectionOfTargetedManagedAppConfiguration{}
	return &this
}

// NewCollectionOfTargetedManagedAppConfigurationWithDefaults instantiates a new CollectionOfTargetedManagedAppConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfTargetedManagedAppConfigurationWithDefaults() *CollectionOfTargetedManagedAppConfiguration {
	this := CollectionOfTargetedManagedAppConfiguration{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfTargetedManagedAppConfiguration) GetValue() []MicrosoftGraphTargetedManagedAppConfiguration {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTargetedManagedAppConfiguration
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTargetedManagedAppConfiguration) GetValueOk() (*[]MicrosoftGraphTargetedManagedAppConfiguration, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfTargetedManagedAppConfiguration) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTargetedManagedAppConfiguration and assigns it to the Value field.
func (o *CollectionOfTargetedManagedAppConfiguration) SetValue(v []MicrosoftGraphTargetedManagedAppConfiguration) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfTargetedManagedAppConfiguration) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTargetedManagedAppConfiguration) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfTargetedManagedAppConfiguration) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfTargetedManagedAppConfiguration) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfTargetedManagedAppConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfTargetedManagedAppConfiguration struct {
	value *CollectionOfTargetedManagedAppConfiguration
	isSet bool
}

func (v NullableCollectionOfTargetedManagedAppConfiguration) Get() *CollectionOfTargetedManagedAppConfiguration {
	return v.value
}

func (v *NullableCollectionOfTargetedManagedAppConfiguration) Set(val *CollectionOfTargetedManagedAppConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfTargetedManagedAppConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfTargetedManagedAppConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfTargetedManagedAppConfiguration(val *CollectionOfTargetedManagedAppConfiguration) *NullableCollectionOfTargetedManagedAppConfiguration {
	return &NullableCollectionOfTargetedManagedAppConfiguration{value: val, isSet: true}
}

func (v NullableCollectionOfTargetedManagedAppConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfTargetedManagedAppConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

