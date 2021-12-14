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

// CollectionOfRoleAssignment struct for CollectionOfRoleAssignment
type CollectionOfRoleAssignment struct {
	Value *[]MicrosoftGraphRoleAssignment `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfRoleAssignment instantiates a new CollectionOfRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfRoleAssignment() *CollectionOfRoleAssignment {
	this := CollectionOfRoleAssignment{}
	return &this
}

// NewCollectionOfRoleAssignmentWithDefaults instantiates a new CollectionOfRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfRoleAssignmentWithDefaults() *CollectionOfRoleAssignment {
	this := CollectionOfRoleAssignment{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfRoleAssignment) GetValue() []MicrosoftGraphRoleAssignment {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphRoleAssignment
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfRoleAssignment) GetValueOk() (*[]MicrosoftGraphRoleAssignment, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfRoleAssignment) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphRoleAssignment and assigns it to the Value field.
func (o *CollectionOfRoleAssignment) SetValue(v []MicrosoftGraphRoleAssignment) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfRoleAssignment) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfRoleAssignment) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfRoleAssignment) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfRoleAssignment) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfRoleAssignment struct {
	value *CollectionOfRoleAssignment
	isSet bool
}

func (v NullableCollectionOfRoleAssignment) Get() *CollectionOfRoleAssignment {
	return v.value
}

func (v *NullableCollectionOfRoleAssignment) Set(val *CollectionOfRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfRoleAssignment(val *CollectionOfRoleAssignment) *NullableCollectionOfRoleAssignment {
	return &NullableCollectionOfRoleAssignment{value: val, isSet: true}
}

func (v NullableCollectionOfRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


