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

// CollectionOfTimeOff struct for CollectionOfTimeOff
type CollectionOfTimeOff struct {
	Value *[]MicrosoftGraphTimeOff `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfTimeOff instantiates a new CollectionOfTimeOff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfTimeOff() *CollectionOfTimeOff {
	this := CollectionOfTimeOff{}
	return &this
}

// NewCollectionOfTimeOffWithDefaults instantiates a new CollectionOfTimeOff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfTimeOffWithDefaults() *CollectionOfTimeOff {
	this := CollectionOfTimeOff{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfTimeOff) GetValue() []MicrosoftGraphTimeOff {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTimeOff
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTimeOff) GetValueOk() (*[]MicrosoftGraphTimeOff, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfTimeOff) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTimeOff and assigns it to the Value field.
func (o *CollectionOfTimeOff) SetValue(v []MicrosoftGraphTimeOff) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfTimeOff) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTimeOff) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfTimeOff) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfTimeOff) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfTimeOff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfTimeOff struct {
	value *CollectionOfTimeOff
	isSet bool
}

func (v NullableCollectionOfTimeOff) Get() *CollectionOfTimeOff {
	return v.value
}

func (v *NullableCollectionOfTimeOff) Set(val *CollectionOfTimeOff) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfTimeOff) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfTimeOff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfTimeOff(val *CollectionOfTimeOff) *NullableCollectionOfTimeOff {
	return &NullableCollectionOfTimeOff{value: val, isSet: true}
}

func (v NullableCollectionOfTimeOff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfTimeOff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


