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

// CollectionOfUserFlowLanguageConfiguration struct for CollectionOfUserFlowLanguageConfiguration
type CollectionOfUserFlowLanguageConfiguration struct {
	Value *[]MicrosoftGraphUserFlowLanguageConfiguration `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfUserFlowLanguageConfiguration instantiates a new CollectionOfUserFlowLanguageConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfUserFlowLanguageConfiguration() *CollectionOfUserFlowLanguageConfiguration {
	this := CollectionOfUserFlowLanguageConfiguration{}
	return &this
}

// NewCollectionOfUserFlowLanguageConfigurationWithDefaults instantiates a new CollectionOfUserFlowLanguageConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfUserFlowLanguageConfigurationWithDefaults() *CollectionOfUserFlowLanguageConfiguration {
	this := CollectionOfUserFlowLanguageConfiguration{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfUserFlowLanguageConfiguration) GetValue() []MicrosoftGraphUserFlowLanguageConfiguration {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphUserFlowLanguageConfiguration
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserFlowLanguageConfiguration) GetValueOk() (*[]MicrosoftGraphUserFlowLanguageConfiguration, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfUserFlowLanguageConfiguration) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphUserFlowLanguageConfiguration and assigns it to the Value field.
func (o *CollectionOfUserFlowLanguageConfiguration) SetValue(v []MicrosoftGraphUserFlowLanguageConfiguration) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfUserFlowLanguageConfiguration) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserFlowLanguageConfiguration) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfUserFlowLanguageConfiguration) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfUserFlowLanguageConfiguration) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfUserFlowLanguageConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfUserFlowLanguageConfiguration struct {
	value *CollectionOfUserFlowLanguageConfiguration
	isSet bool
}

func (v NullableCollectionOfUserFlowLanguageConfiguration) Get() *CollectionOfUserFlowLanguageConfiguration {
	return v.value
}

func (v *NullableCollectionOfUserFlowLanguageConfiguration) Set(val *CollectionOfUserFlowLanguageConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfUserFlowLanguageConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfUserFlowLanguageConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfUserFlowLanguageConfiguration(val *CollectionOfUserFlowLanguageConfiguration) *NullableCollectionOfUserFlowLanguageConfiguration {
	return &NullableCollectionOfUserFlowLanguageConfiguration{value: val, isSet: true}
}

func (v NullableCollectionOfUserFlowLanguageConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfUserFlowLanguageConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


