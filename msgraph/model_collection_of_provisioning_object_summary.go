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

// CollectionOfProvisioningObjectSummary struct for CollectionOfProvisioningObjectSummary
type CollectionOfProvisioningObjectSummary struct {
	Value *[]MicrosoftGraphProvisioningObjectSummary `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfProvisioningObjectSummary instantiates a new CollectionOfProvisioningObjectSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfProvisioningObjectSummary() *CollectionOfProvisioningObjectSummary {
	this := CollectionOfProvisioningObjectSummary{}
	return &this
}

// NewCollectionOfProvisioningObjectSummaryWithDefaults instantiates a new CollectionOfProvisioningObjectSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfProvisioningObjectSummaryWithDefaults() *CollectionOfProvisioningObjectSummary {
	this := CollectionOfProvisioningObjectSummary{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfProvisioningObjectSummary) GetValue() []MicrosoftGraphProvisioningObjectSummary {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphProvisioningObjectSummary
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfProvisioningObjectSummary) GetValueOk() (*[]MicrosoftGraphProvisioningObjectSummary, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfProvisioningObjectSummary) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphProvisioningObjectSummary and assigns it to the Value field.
func (o *CollectionOfProvisioningObjectSummary) SetValue(v []MicrosoftGraphProvisioningObjectSummary) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfProvisioningObjectSummary) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfProvisioningObjectSummary) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfProvisioningObjectSummary) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfProvisioningObjectSummary) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfProvisioningObjectSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfProvisioningObjectSummary struct {
	value *CollectionOfProvisioningObjectSummary
	isSet bool
}

func (v NullableCollectionOfProvisioningObjectSummary) Get() *CollectionOfProvisioningObjectSummary {
	return v.value
}

func (v *NullableCollectionOfProvisioningObjectSummary) Set(val *CollectionOfProvisioningObjectSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfProvisioningObjectSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfProvisioningObjectSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfProvisioningObjectSummary(val *CollectionOfProvisioningObjectSummary) *NullableCollectionOfProvisioningObjectSummary {
	return &NullableCollectionOfProvisioningObjectSummary{value: val, isSet: true}
}

func (v NullableCollectionOfProvisioningObjectSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfProvisioningObjectSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

