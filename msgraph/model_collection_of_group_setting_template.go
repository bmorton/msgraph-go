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

// CollectionOfGroupSettingTemplate struct for CollectionOfGroupSettingTemplate
type CollectionOfGroupSettingTemplate struct {
	Value *[]MicrosoftGraphGroupSettingTemplate `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfGroupSettingTemplate instantiates a new CollectionOfGroupSettingTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfGroupSettingTemplate() *CollectionOfGroupSettingTemplate {
	this := CollectionOfGroupSettingTemplate{}
	return &this
}

// NewCollectionOfGroupSettingTemplateWithDefaults instantiates a new CollectionOfGroupSettingTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfGroupSettingTemplateWithDefaults() *CollectionOfGroupSettingTemplate {
	this := CollectionOfGroupSettingTemplate{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfGroupSettingTemplate) GetValue() []MicrosoftGraphGroupSettingTemplate {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphGroupSettingTemplate
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroupSettingTemplate) GetValueOk() (*[]MicrosoftGraphGroupSettingTemplate, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfGroupSettingTemplate) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphGroupSettingTemplate and assigns it to the Value field.
func (o *CollectionOfGroupSettingTemplate) SetValue(v []MicrosoftGraphGroupSettingTemplate) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfGroupSettingTemplate) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroupSettingTemplate) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfGroupSettingTemplate) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfGroupSettingTemplate) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfGroupSettingTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfGroupSettingTemplate struct {
	value *CollectionOfGroupSettingTemplate
	isSet bool
}

func (v NullableCollectionOfGroupSettingTemplate) Get() *CollectionOfGroupSettingTemplate {
	return v.value
}

func (v *NullableCollectionOfGroupSettingTemplate) Set(val *CollectionOfGroupSettingTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfGroupSettingTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfGroupSettingTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfGroupSettingTemplate(val *CollectionOfGroupSettingTemplate) *NullableCollectionOfGroupSettingTemplate {
	return &NullableCollectionOfGroupSettingTemplate{value: val, isSet: true}
}

func (v NullableCollectionOfGroupSettingTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfGroupSettingTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


