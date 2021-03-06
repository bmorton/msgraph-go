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

// CollectionOfSectionGroup struct for CollectionOfSectionGroup
type CollectionOfSectionGroup struct {
	Value *[]MicrosoftGraphSectionGroup `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfSectionGroup instantiates a new CollectionOfSectionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfSectionGroup() *CollectionOfSectionGroup {
	this := CollectionOfSectionGroup{}
	return &this
}

// NewCollectionOfSectionGroupWithDefaults instantiates a new CollectionOfSectionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfSectionGroupWithDefaults() *CollectionOfSectionGroup {
	this := CollectionOfSectionGroup{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfSectionGroup) GetValue() []MicrosoftGraphSectionGroup {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphSectionGroup
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSectionGroup) GetValueOk() (*[]MicrosoftGraphSectionGroup, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfSectionGroup) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the Value field.
func (o *CollectionOfSectionGroup) SetValue(v []MicrosoftGraphSectionGroup) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfSectionGroup) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfSectionGroup) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfSectionGroup) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfSectionGroup) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfSectionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfSectionGroup struct {
	value *CollectionOfSectionGroup
	isSet bool
}

func (v NullableCollectionOfSectionGroup) Get() *CollectionOfSectionGroup {
	return v.value
}

func (v *NullableCollectionOfSectionGroup) Set(val *CollectionOfSectionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfSectionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfSectionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfSectionGroup(val *CollectionOfSectionGroup) *NullableCollectionOfSectionGroup {
	return &NullableCollectionOfSectionGroup{value: val, isSet: true}
}

func (v NullableCollectionOfSectionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfSectionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


