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

// CollectionOfEducationAssignment struct for CollectionOfEducationAssignment
type CollectionOfEducationAssignment struct {
	Value *[]MicrosoftGraphEducationAssignment `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfEducationAssignment instantiates a new CollectionOfEducationAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfEducationAssignment() *CollectionOfEducationAssignment {
	this := CollectionOfEducationAssignment{}
	return &this
}

// NewCollectionOfEducationAssignmentWithDefaults instantiates a new CollectionOfEducationAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfEducationAssignmentWithDefaults() *CollectionOfEducationAssignment {
	this := CollectionOfEducationAssignment{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfEducationAssignment) GetValue() []MicrosoftGraphEducationAssignment {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphEducationAssignment
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationAssignment) GetValueOk() (*[]MicrosoftGraphEducationAssignment, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfEducationAssignment) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphEducationAssignment and assigns it to the Value field.
func (o *CollectionOfEducationAssignment) SetValue(v []MicrosoftGraphEducationAssignment) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfEducationAssignment) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationAssignment) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfEducationAssignment) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfEducationAssignment) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfEducationAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfEducationAssignment struct {
	value *CollectionOfEducationAssignment
	isSet bool
}

func (v NullableCollectionOfEducationAssignment) Get() *CollectionOfEducationAssignment {
	return v.value
}

func (v *NullableCollectionOfEducationAssignment) Set(val *CollectionOfEducationAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfEducationAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfEducationAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfEducationAssignment(val *CollectionOfEducationAssignment) *NullableCollectionOfEducationAssignment {
	return &NullableCollectionOfEducationAssignment{value: val, isSet: true}
}

func (v NullableCollectionOfEducationAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfEducationAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


