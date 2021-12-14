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

// CollectionOfCalendarPermission struct for CollectionOfCalendarPermission
type CollectionOfCalendarPermission struct {
	Value *[]MicrosoftGraphCalendarPermission `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfCalendarPermission instantiates a new CollectionOfCalendarPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfCalendarPermission() *CollectionOfCalendarPermission {
	this := CollectionOfCalendarPermission{}
	return &this
}

// NewCollectionOfCalendarPermissionWithDefaults instantiates a new CollectionOfCalendarPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfCalendarPermissionWithDefaults() *CollectionOfCalendarPermission {
	this := CollectionOfCalendarPermission{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfCalendarPermission) GetValue() []MicrosoftGraphCalendarPermission {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphCalendarPermission
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfCalendarPermission) GetValueOk() (*[]MicrosoftGraphCalendarPermission, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfCalendarPermission) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphCalendarPermission and assigns it to the Value field.
func (o *CollectionOfCalendarPermission) SetValue(v []MicrosoftGraphCalendarPermission) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfCalendarPermission) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfCalendarPermission) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfCalendarPermission) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfCalendarPermission) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfCalendarPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfCalendarPermission struct {
	value *CollectionOfCalendarPermission
	isSet bool
}

func (v NullableCollectionOfCalendarPermission) Get() *CollectionOfCalendarPermission {
	return v.value
}

func (v *NullableCollectionOfCalendarPermission) Set(val *CollectionOfCalendarPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfCalendarPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfCalendarPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfCalendarPermission(val *CollectionOfCalendarPermission) *NullableCollectionOfCalendarPermission {
	return &NullableCollectionOfCalendarPermission{value: val, isSet: true}
}

func (v NullableCollectionOfCalendarPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfCalendarPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

