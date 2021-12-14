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

// CollectionOfPermission struct for CollectionOfPermission
type CollectionOfPermission struct {
	Value *[]MicrosoftGraphPermission `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfPermission instantiates a new CollectionOfPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfPermission() *CollectionOfPermission {
	this := CollectionOfPermission{}
	return &this
}

// NewCollectionOfPermissionWithDefaults instantiates a new CollectionOfPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfPermissionWithDefaults() *CollectionOfPermission {
	this := CollectionOfPermission{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfPermission) GetValue() []MicrosoftGraphPermission {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphPermission
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPermission) GetValueOk() (*[]MicrosoftGraphPermission, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfPermission) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphPermission and assigns it to the Value field.
func (o *CollectionOfPermission) SetValue(v []MicrosoftGraphPermission) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfPermission) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPermission) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfPermission) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfPermission) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfPermission struct {
	value *CollectionOfPermission
	isSet bool
}

func (v NullableCollectionOfPermission) Get() *CollectionOfPermission {
	return v.value
}

func (v *NullableCollectionOfPermission) Set(val *CollectionOfPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfPermission(val *CollectionOfPermission) *NullableCollectionOfPermission {
	return &NullableCollectionOfPermission{value: val, isSet: true}
}

func (v NullableCollectionOfPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


