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

// CollectionOfEducationRubric struct for CollectionOfEducationRubric
type CollectionOfEducationRubric struct {
	Value *[]MicrosoftGraphEducationRubric `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfEducationRubric instantiates a new CollectionOfEducationRubric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfEducationRubric() *CollectionOfEducationRubric {
	this := CollectionOfEducationRubric{}
	return &this
}

// NewCollectionOfEducationRubricWithDefaults instantiates a new CollectionOfEducationRubric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfEducationRubricWithDefaults() *CollectionOfEducationRubric {
	this := CollectionOfEducationRubric{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfEducationRubric) GetValue() []MicrosoftGraphEducationRubric {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphEducationRubric
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationRubric) GetValueOk() (*[]MicrosoftGraphEducationRubric, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfEducationRubric) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphEducationRubric and assigns it to the Value field.
func (o *CollectionOfEducationRubric) SetValue(v []MicrosoftGraphEducationRubric) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfEducationRubric) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationRubric) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfEducationRubric) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfEducationRubric) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfEducationRubric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfEducationRubric struct {
	value *CollectionOfEducationRubric
	isSet bool
}

func (v NullableCollectionOfEducationRubric) Get() *CollectionOfEducationRubric {
	return v.value
}

func (v *NullableCollectionOfEducationRubric) Set(val *CollectionOfEducationRubric) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfEducationRubric) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfEducationRubric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfEducationRubric(val *CollectionOfEducationRubric) *NullableCollectionOfEducationRubric {
	return &NullableCollectionOfEducationRubric{value: val, isSet: true}
}

func (v NullableCollectionOfEducationRubric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfEducationRubric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


