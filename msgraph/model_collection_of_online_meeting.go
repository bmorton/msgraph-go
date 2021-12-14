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

// CollectionOfOnlineMeeting struct for CollectionOfOnlineMeeting
type CollectionOfOnlineMeeting struct {
	Value *[]MicrosoftGraphOnlineMeeting `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOnlineMeeting instantiates a new CollectionOfOnlineMeeting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOnlineMeeting() *CollectionOfOnlineMeeting {
	this := CollectionOfOnlineMeeting{}
	return &this
}

// NewCollectionOfOnlineMeetingWithDefaults instantiates a new CollectionOfOnlineMeeting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOnlineMeetingWithDefaults() *CollectionOfOnlineMeeting {
	this := CollectionOfOnlineMeeting{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOnlineMeeting) GetValue() []MicrosoftGraphOnlineMeeting {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOnlineMeeting
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOnlineMeeting) GetValueOk() (*[]MicrosoftGraphOnlineMeeting, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOnlineMeeting) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOnlineMeeting and assigns it to the Value field.
func (o *CollectionOfOnlineMeeting) SetValue(v []MicrosoftGraphOnlineMeeting) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOnlineMeeting) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOnlineMeeting) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOnlineMeeting) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOnlineMeeting) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOnlineMeeting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOnlineMeeting struct {
	value *CollectionOfOnlineMeeting
	isSet bool
}

func (v NullableCollectionOfOnlineMeeting) Get() *CollectionOfOnlineMeeting {
	return v.value
}

func (v *NullableCollectionOfOnlineMeeting) Set(val *CollectionOfOnlineMeeting) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOnlineMeeting) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOnlineMeeting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOnlineMeeting(val *CollectionOfOnlineMeeting) *NullableCollectionOfOnlineMeeting {
	return &NullableCollectionOfOnlineMeeting{value: val, isSet: true}
}

func (v NullableCollectionOfOnlineMeeting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOnlineMeeting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


