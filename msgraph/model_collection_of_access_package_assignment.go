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

// CollectionOfAccessPackageAssignment struct for CollectionOfAccessPackageAssignment
type CollectionOfAccessPackageAssignment struct {
	Value *[]MicrosoftGraphAccessPackageAssignment `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfAccessPackageAssignment instantiates a new CollectionOfAccessPackageAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfAccessPackageAssignment() *CollectionOfAccessPackageAssignment {
	this := CollectionOfAccessPackageAssignment{}
	return &this
}

// NewCollectionOfAccessPackageAssignmentWithDefaults instantiates a new CollectionOfAccessPackageAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfAccessPackageAssignmentWithDefaults() *CollectionOfAccessPackageAssignment {
	this := CollectionOfAccessPackageAssignment{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfAccessPackageAssignment) GetValue() []MicrosoftGraphAccessPackageAssignment {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphAccessPackageAssignment
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAccessPackageAssignment) GetValueOk() (*[]MicrosoftGraphAccessPackageAssignment, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfAccessPackageAssignment) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphAccessPackageAssignment and assigns it to the Value field.
func (o *CollectionOfAccessPackageAssignment) SetValue(v []MicrosoftGraphAccessPackageAssignment) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfAccessPackageAssignment) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAccessPackageAssignment) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfAccessPackageAssignment) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfAccessPackageAssignment) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfAccessPackageAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfAccessPackageAssignment struct {
	value *CollectionOfAccessPackageAssignment
	isSet bool
}

func (v NullableCollectionOfAccessPackageAssignment) Get() *CollectionOfAccessPackageAssignment {
	return v.value
}

func (v *NullableCollectionOfAccessPackageAssignment) Set(val *CollectionOfAccessPackageAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfAccessPackageAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfAccessPackageAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfAccessPackageAssignment(val *CollectionOfAccessPackageAssignment) *NullableCollectionOfAccessPackageAssignment {
	return &NullableCollectionOfAccessPackageAssignment{value: val, isSet: true}
}

func (v NullableCollectionOfAccessPackageAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfAccessPackageAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


