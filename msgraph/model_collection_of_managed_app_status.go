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

// CollectionOfManagedAppStatus struct for CollectionOfManagedAppStatus
type CollectionOfManagedAppStatus struct {
	Value *[]MicrosoftGraphManagedAppStatus `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfManagedAppStatus instantiates a new CollectionOfManagedAppStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfManagedAppStatus() *CollectionOfManagedAppStatus {
	this := CollectionOfManagedAppStatus{}
	return &this
}

// NewCollectionOfManagedAppStatusWithDefaults instantiates a new CollectionOfManagedAppStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfManagedAppStatusWithDefaults() *CollectionOfManagedAppStatus {
	this := CollectionOfManagedAppStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfManagedAppStatus) GetValue() []MicrosoftGraphManagedAppStatus {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphManagedAppStatus
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfManagedAppStatus) GetValueOk() (*[]MicrosoftGraphManagedAppStatus, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfManagedAppStatus) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphManagedAppStatus and assigns it to the Value field.
func (o *CollectionOfManagedAppStatus) SetValue(v []MicrosoftGraphManagedAppStatus) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfManagedAppStatus) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfManagedAppStatus) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfManagedAppStatus) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfManagedAppStatus) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfManagedAppStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfManagedAppStatus struct {
	value *CollectionOfManagedAppStatus
	isSet bool
}

func (v NullableCollectionOfManagedAppStatus) Get() *CollectionOfManagedAppStatus {
	return v.value
}

func (v *NullableCollectionOfManagedAppStatus) Set(val *CollectionOfManagedAppStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfManagedAppStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfManagedAppStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfManagedAppStatus(val *CollectionOfManagedAppStatus) *NullableCollectionOfManagedAppStatus {
	return &NullableCollectionOfManagedAppStatus{value: val, isSet: true}
}

func (v NullableCollectionOfManagedAppStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfManagedAppStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


