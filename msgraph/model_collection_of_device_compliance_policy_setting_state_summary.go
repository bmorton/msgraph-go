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

// CollectionOfDeviceCompliancePolicySettingStateSummary struct for CollectionOfDeviceCompliancePolicySettingStateSummary
type CollectionOfDeviceCompliancePolicySettingStateSummary struct {
	Value *[]MicrosoftGraphDeviceCompliancePolicySettingStateSummary `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDeviceCompliancePolicySettingStateSummary instantiates a new CollectionOfDeviceCompliancePolicySettingStateSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDeviceCompliancePolicySettingStateSummary() *CollectionOfDeviceCompliancePolicySettingStateSummary {
	this := CollectionOfDeviceCompliancePolicySettingStateSummary{}
	return &this
}

// NewCollectionOfDeviceCompliancePolicySettingStateSummaryWithDefaults instantiates a new CollectionOfDeviceCompliancePolicySettingStateSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceCompliancePolicySettingStateSummaryWithDefaults() *CollectionOfDeviceCompliancePolicySettingStateSummary {
	this := CollectionOfDeviceCompliancePolicySettingStateSummary{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) GetValue() []MicrosoftGraphDeviceCompliancePolicySettingStateSummary {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceCompliancePolicySettingStateSummary
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) GetValueOk() (*[]MicrosoftGraphDeviceCompliancePolicySettingStateSummary, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceCompliancePolicySettingStateSummary and assigns it to the Value field.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) SetValue(v []MicrosoftGraphDeviceCompliancePolicySettingStateSummary) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDeviceCompliancePolicySettingStateSummary) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDeviceCompliancePolicySettingStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDeviceCompliancePolicySettingStateSummary struct {
	value *CollectionOfDeviceCompliancePolicySettingStateSummary
	isSet bool
}

func (v NullableCollectionOfDeviceCompliancePolicySettingStateSummary) Get() *CollectionOfDeviceCompliancePolicySettingStateSummary {
	return v.value
}

func (v *NullableCollectionOfDeviceCompliancePolicySettingStateSummary) Set(val *CollectionOfDeviceCompliancePolicySettingStateSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDeviceCompliancePolicySettingStateSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDeviceCompliancePolicySettingStateSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDeviceCompliancePolicySettingStateSummary(val *CollectionOfDeviceCompliancePolicySettingStateSummary) *NullableCollectionOfDeviceCompliancePolicySettingStateSummary {
	return &NullableCollectionOfDeviceCompliancePolicySettingStateSummary{value: val, isSet: true}
}

func (v NullableCollectionOfDeviceCompliancePolicySettingStateSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDeviceCompliancePolicySettingStateSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


