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

// CollectionOfAppRoleAssignment struct for CollectionOfAppRoleAssignment
type CollectionOfAppRoleAssignment struct {
	Value *[]MicrosoftGraphAppRoleAssignment `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfAppRoleAssignment instantiates a new CollectionOfAppRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfAppRoleAssignment() *CollectionOfAppRoleAssignment {
	this := CollectionOfAppRoleAssignment{}
	return &this
}

// NewCollectionOfAppRoleAssignmentWithDefaults instantiates a new CollectionOfAppRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfAppRoleAssignmentWithDefaults() *CollectionOfAppRoleAssignment {
	this := CollectionOfAppRoleAssignment{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfAppRoleAssignment) GetValue() []MicrosoftGraphAppRoleAssignment {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphAppRoleAssignment
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAppRoleAssignment) GetValueOk() (*[]MicrosoftGraphAppRoleAssignment, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfAppRoleAssignment) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphAppRoleAssignment and assigns it to the Value field.
func (o *CollectionOfAppRoleAssignment) SetValue(v []MicrosoftGraphAppRoleAssignment) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfAppRoleAssignment) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAppRoleAssignment) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfAppRoleAssignment) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfAppRoleAssignment) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfAppRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfAppRoleAssignment struct {
	value *CollectionOfAppRoleAssignment
	isSet bool
}

func (v NullableCollectionOfAppRoleAssignment) Get() *CollectionOfAppRoleAssignment {
	return v.value
}

func (v *NullableCollectionOfAppRoleAssignment) Set(val *CollectionOfAppRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfAppRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfAppRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfAppRoleAssignment(val *CollectionOfAppRoleAssignment) *NullableCollectionOfAppRoleAssignment {
	return &NullableCollectionOfAppRoleAssignment{value: val, isSet: true}
}

func (v NullableCollectionOfAppRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfAppRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

