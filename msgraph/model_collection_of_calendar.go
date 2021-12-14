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

// CollectionOfCalendar struct for CollectionOfCalendar
type CollectionOfCalendar struct {
	Value *[]MicrosoftGraphCalendar `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfCalendar instantiates a new CollectionOfCalendar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfCalendar() *CollectionOfCalendar {
	this := CollectionOfCalendar{}
	return &this
}

// NewCollectionOfCalendarWithDefaults instantiates a new CollectionOfCalendar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfCalendarWithDefaults() *CollectionOfCalendar {
	this := CollectionOfCalendar{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfCalendar) GetValue() []MicrosoftGraphCalendar {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphCalendar
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfCalendar) GetValueOk() (*[]MicrosoftGraphCalendar, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfCalendar) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphCalendar and assigns it to the Value field.
func (o *CollectionOfCalendar) SetValue(v []MicrosoftGraphCalendar) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfCalendar) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfCalendar) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfCalendar) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfCalendar) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfCalendar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfCalendar struct {
	value *CollectionOfCalendar
	isSet bool
}

func (v NullableCollectionOfCalendar) Get() *CollectionOfCalendar {
	return v.value
}

func (v *NullableCollectionOfCalendar) Set(val *CollectionOfCalendar) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfCalendar) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfCalendar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfCalendar(val *CollectionOfCalendar) *NullableCollectionOfCalendar {
	return &NullableCollectionOfCalendar{value: val, isSet: true}
}

func (v NullableCollectionOfCalendar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfCalendar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

