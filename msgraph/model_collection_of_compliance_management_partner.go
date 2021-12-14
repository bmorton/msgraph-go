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

// CollectionOfComplianceManagementPartner struct for CollectionOfComplianceManagementPartner
type CollectionOfComplianceManagementPartner struct {
	Value *[]MicrosoftGraphComplianceManagementPartner `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfComplianceManagementPartner instantiates a new CollectionOfComplianceManagementPartner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfComplianceManagementPartner() *CollectionOfComplianceManagementPartner {
	this := CollectionOfComplianceManagementPartner{}
	return &this
}

// NewCollectionOfComplianceManagementPartnerWithDefaults instantiates a new CollectionOfComplianceManagementPartner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfComplianceManagementPartnerWithDefaults() *CollectionOfComplianceManagementPartner {
	this := CollectionOfComplianceManagementPartner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfComplianceManagementPartner) GetValue() []MicrosoftGraphComplianceManagementPartner {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphComplianceManagementPartner
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfComplianceManagementPartner) GetValueOk() (*[]MicrosoftGraphComplianceManagementPartner, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfComplianceManagementPartner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphComplianceManagementPartner and assigns it to the Value field.
func (o *CollectionOfComplianceManagementPartner) SetValue(v []MicrosoftGraphComplianceManagementPartner) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfComplianceManagementPartner) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfComplianceManagementPartner) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfComplianceManagementPartner) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfComplianceManagementPartner) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfComplianceManagementPartner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfComplianceManagementPartner struct {
	value *CollectionOfComplianceManagementPartner
	isSet bool
}

func (v NullableCollectionOfComplianceManagementPartner) Get() *CollectionOfComplianceManagementPartner {
	return v.value
}

func (v *NullableCollectionOfComplianceManagementPartner) Set(val *CollectionOfComplianceManagementPartner) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfComplianceManagementPartner) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfComplianceManagementPartner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfComplianceManagementPartner(val *CollectionOfComplianceManagementPartner) *NullableCollectionOfComplianceManagementPartner {
	return &NullableCollectionOfComplianceManagementPartner{value: val, isSet: true}
}

func (v NullableCollectionOfComplianceManagementPartner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfComplianceManagementPartner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


