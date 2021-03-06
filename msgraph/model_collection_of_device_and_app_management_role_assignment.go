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

// CollectionOfDeviceAndAppManagementRoleAssignment struct for CollectionOfDeviceAndAppManagementRoleAssignment
type CollectionOfDeviceAndAppManagementRoleAssignment struct {
	Value *[]MicrosoftGraphDeviceAndAppManagementRoleAssignment `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDeviceAndAppManagementRoleAssignment instantiates a new CollectionOfDeviceAndAppManagementRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDeviceAndAppManagementRoleAssignment() *CollectionOfDeviceAndAppManagementRoleAssignment {
	this := CollectionOfDeviceAndAppManagementRoleAssignment{}
	return &this
}

// NewCollectionOfDeviceAndAppManagementRoleAssignmentWithDefaults instantiates a new CollectionOfDeviceAndAppManagementRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceAndAppManagementRoleAssignmentWithDefaults() *CollectionOfDeviceAndAppManagementRoleAssignment {
	this := CollectionOfDeviceAndAppManagementRoleAssignment{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) GetValue() []MicrosoftGraphDeviceAndAppManagementRoleAssignment {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceAndAppManagementRoleAssignment
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) GetValueOk() (*[]MicrosoftGraphDeviceAndAppManagementRoleAssignment, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceAndAppManagementRoleAssignment and assigns it to the Value field.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) SetValue(v []MicrosoftGraphDeviceAndAppManagementRoleAssignment) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDeviceAndAppManagementRoleAssignment struct {
	value *CollectionOfDeviceAndAppManagementRoleAssignment
	isSet bool
}

func (v NullableCollectionOfDeviceAndAppManagementRoleAssignment) Get() *CollectionOfDeviceAndAppManagementRoleAssignment {
	return v.value
}

func (v *NullableCollectionOfDeviceAndAppManagementRoleAssignment) Set(val *CollectionOfDeviceAndAppManagementRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDeviceAndAppManagementRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDeviceAndAppManagementRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDeviceAndAppManagementRoleAssignment(val *CollectionOfDeviceAndAppManagementRoleAssignment) *NullableCollectionOfDeviceAndAppManagementRoleAssignment {
	return &NullableCollectionOfDeviceAndAppManagementRoleAssignment{value: val, isSet: true}
}

func (v NullableCollectionOfDeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDeviceAndAppManagementRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


