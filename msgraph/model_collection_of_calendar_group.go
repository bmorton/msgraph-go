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

// CollectionOfCalendarGroup struct for CollectionOfCalendarGroup
type CollectionOfCalendarGroup struct {
	Value *[]MicrosoftGraphCalendarGroup `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfCalendarGroup instantiates a new CollectionOfCalendarGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfCalendarGroup() *CollectionOfCalendarGroup {
	this := CollectionOfCalendarGroup{}
	return &this
}

// NewCollectionOfCalendarGroupWithDefaults instantiates a new CollectionOfCalendarGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfCalendarGroupWithDefaults() *CollectionOfCalendarGroup {
	this := CollectionOfCalendarGroup{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfCalendarGroup) GetValue() []MicrosoftGraphCalendarGroup {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphCalendarGroup
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfCalendarGroup) GetValueOk() (*[]MicrosoftGraphCalendarGroup, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfCalendarGroup) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphCalendarGroup and assigns it to the Value field.
func (o *CollectionOfCalendarGroup) SetValue(v []MicrosoftGraphCalendarGroup) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfCalendarGroup) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfCalendarGroup) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfCalendarGroup) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfCalendarGroup) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfCalendarGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfCalendarGroup struct {
	value *CollectionOfCalendarGroup
	isSet bool
}

func (v NullableCollectionOfCalendarGroup) Get() *CollectionOfCalendarGroup {
	return v.value
}

func (v *NullableCollectionOfCalendarGroup) Set(val *CollectionOfCalendarGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfCalendarGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfCalendarGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfCalendarGroup(val *CollectionOfCalendarGroup) *NullableCollectionOfCalendarGroup {
	return &NullableCollectionOfCalendarGroup{value: val, isSet: true}
}

func (v NullableCollectionOfCalendarGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfCalendarGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


