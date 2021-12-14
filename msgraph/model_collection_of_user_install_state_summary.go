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

// CollectionOfUserInstallStateSummary struct for CollectionOfUserInstallStateSummary
type CollectionOfUserInstallStateSummary struct {
	Value *[]MicrosoftGraphUserInstallStateSummary `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfUserInstallStateSummary instantiates a new CollectionOfUserInstallStateSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfUserInstallStateSummary() *CollectionOfUserInstallStateSummary {
	this := CollectionOfUserInstallStateSummary{}
	return &this
}

// NewCollectionOfUserInstallStateSummaryWithDefaults instantiates a new CollectionOfUserInstallStateSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfUserInstallStateSummaryWithDefaults() *CollectionOfUserInstallStateSummary {
	this := CollectionOfUserInstallStateSummary{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfUserInstallStateSummary) GetValue() []MicrosoftGraphUserInstallStateSummary {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphUserInstallStateSummary
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserInstallStateSummary) GetValueOk() (*[]MicrosoftGraphUserInstallStateSummary, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfUserInstallStateSummary) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphUserInstallStateSummary and assigns it to the Value field.
func (o *CollectionOfUserInstallStateSummary) SetValue(v []MicrosoftGraphUserInstallStateSummary) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfUserInstallStateSummary) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserInstallStateSummary) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfUserInstallStateSummary) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfUserInstallStateSummary) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfUserInstallStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfUserInstallStateSummary struct {
	value *CollectionOfUserInstallStateSummary
	isSet bool
}

func (v NullableCollectionOfUserInstallStateSummary) Get() *CollectionOfUserInstallStateSummary {
	return v.value
}

func (v *NullableCollectionOfUserInstallStateSummary) Set(val *CollectionOfUserInstallStateSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfUserInstallStateSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfUserInstallStateSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfUserInstallStateSummary(val *CollectionOfUserInstallStateSummary) *NullableCollectionOfUserInstallStateSummary {
	return &NullableCollectionOfUserInstallStateSummary{value: val, isSet: true}
}

func (v NullableCollectionOfUserInstallStateSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfUserInstallStateSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


