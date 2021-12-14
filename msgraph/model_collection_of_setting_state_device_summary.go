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

// CollectionOfSettingStateDeviceSummary struct for CollectionOfSettingStateDeviceSummary
type CollectionOfSettingStateDeviceSummary struct {
	Value *[]MicrosoftGraphSettingStateDeviceSummary `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfSettingStateDeviceSummary instantiates a new CollectionOfSettingStateDeviceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfSettingStateDeviceSummary() *CollectionOfSettingStateDeviceSummary {
	this := CollectionOfSettingStateDeviceSummary{}
	return &this
}

// NewCollectionOfSettingStateDeviceSummaryWithDefaults instantiates a new CollectionOfSettingStateDeviceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfSettingStateDeviceSummaryWithDefaults() *CollectionOfSettingStateDeviceSummary {
	this := CollectionOfSettingStateDeviceSummary{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfSettingStateDeviceSummary) GetValue() []MicrosoftGraphSettingStateDeviceSummary {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphSettingStateDeviceSummary
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSettingStateDeviceSummary) GetValueOk() (*[]MicrosoftGraphSettingStateDeviceSummary, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfSettingStateDeviceSummary) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the Value field.
func (o *CollectionOfSettingStateDeviceSummary) SetValue(v []MicrosoftGraphSettingStateDeviceSummary) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfSettingStateDeviceSummary) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSettingStateDeviceSummary) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfSettingStateDeviceSummary) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfSettingStateDeviceSummary) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfSettingStateDeviceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfSettingStateDeviceSummary struct {
	value *CollectionOfSettingStateDeviceSummary
	isSet bool
}

func (v NullableCollectionOfSettingStateDeviceSummary) Get() *CollectionOfSettingStateDeviceSummary {
	return v.value
}

func (v *NullableCollectionOfSettingStateDeviceSummary) Set(val *CollectionOfSettingStateDeviceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfSettingStateDeviceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfSettingStateDeviceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfSettingStateDeviceSummary(val *CollectionOfSettingStateDeviceSummary) *NullableCollectionOfSettingStateDeviceSummary {
	return &NullableCollectionOfSettingStateDeviceSummary{value: val, isSet: true}
}

func (v NullableCollectionOfSettingStateDeviceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfSettingStateDeviceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


