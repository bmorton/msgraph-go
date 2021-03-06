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

// CollectionOfDeviceComplianceUserStatus struct for CollectionOfDeviceComplianceUserStatus
type CollectionOfDeviceComplianceUserStatus struct {
	Value *[]MicrosoftGraphDeviceComplianceUserStatus `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDeviceComplianceUserStatus instantiates a new CollectionOfDeviceComplianceUserStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDeviceComplianceUserStatus() *CollectionOfDeviceComplianceUserStatus {
	this := CollectionOfDeviceComplianceUserStatus{}
	return &this
}

// NewCollectionOfDeviceComplianceUserStatusWithDefaults instantiates a new CollectionOfDeviceComplianceUserStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceComplianceUserStatusWithDefaults() *CollectionOfDeviceComplianceUserStatus {
	this := CollectionOfDeviceComplianceUserStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDeviceComplianceUserStatus) GetValue() []MicrosoftGraphDeviceComplianceUserStatus {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceComplianceUserStatus
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceComplianceUserStatus) GetValueOk() (*[]MicrosoftGraphDeviceComplianceUserStatus, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceComplianceUserStatus) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceComplianceUserStatus and assigns it to the Value field.
func (o *CollectionOfDeviceComplianceUserStatus) SetValue(v []MicrosoftGraphDeviceComplianceUserStatus) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDeviceComplianceUserStatus) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceComplianceUserStatus) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDeviceComplianceUserStatus) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDeviceComplianceUserStatus) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDeviceComplianceUserStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDeviceComplianceUserStatus struct {
	value *CollectionOfDeviceComplianceUserStatus
	isSet bool
}

func (v NullableCollectionOfDeviceComplianceUserStatus) Get() *CollectionOfDeviceComplianceUserStatus {
	return v.value
}

func (v *NullableCollectionOfDeviceComplianceUserStatus) Set(val *CollectionOfDeviceComplianceUserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDeviceComplianceUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDeviceComplianceUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDeviceComplianceUserStatus(val *CollectionOfDeviceComplianceUserStatus) *NullableCollectionOfDeviceComplianceUserStatus {
	return &NullableCollectionOfDeviceComplianceUserStatus{value: val, isSet: true}
}

func (v NullableCollectionOfDeviceComplianceUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDeviceComplianceUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


