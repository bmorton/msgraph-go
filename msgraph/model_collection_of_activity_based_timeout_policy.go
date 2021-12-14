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

// CollectionOfActivityBasedTimeoutPolicy struct for CollectionOfActivityBasedTimeoutPolicy
type CollectionOfActivityBasedTimeoutPolicy struct {
	Value *[]MicrosoftGraphActivityBasedTimeoutPolicy `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfActivityBasedTimeoutPolicy instantiates a new CollectionOfActivityBasedTimeoutPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfActivityBasedTimeoutPolicy() *CollectionOfActivityBasedTimeoutPolicy {
	this := CollectionOfActivityBasedTimeoutPolicy{}
	return &this
}

// NewCollectionOfActivityBasedTimeoutPolicyWithDefaults instantiates a new CollectionOfActivityBasedTimeoutPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfActivityBasedTimeoutPolicyWithDefaults() *CollectionOfActivityBasedTimeoutPolicy {
	this := CollectionOfActivityBasedTimeoutPolicy{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfActivityBasedTimeoutPolicy) GetValue() []MicrosoftGraphActivityBasedTimeoutPolicy {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphActivityBasedTimeoutPolicy
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfActivityBasedTimeoutPolicy) GetValueOk() (*[]MicrosoftGraphActivityBasedTimeoutPolicy, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfActivityBasedTimeoutPolicy) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphActivityBasedTimeoutPolicy and assigns it to the Value field.
func (o *CollectionOfActivityBasedTimeoutPolicy) SetValue(v []MicrosoftGraphActivityBasedTimeoutPolicy) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfActivityBasedTimeoutPolicy) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfActivityBasedTimeoutPolicy) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfActivityBasedTimeoutPolicy) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfActivityBasedTimeoutPolicy) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfActivityBasedTimeoutPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfActivityBasedTimeoutPolicy struct {
	value *CollectionOfActivityBasedTimeoutPolicy
	isSet bool
}

func (v NullableCollectionOfActivityBasedTimeoutPolicy) Get() *CollectionOfActivityBasedTimeoutPolicy {
	return v.value
}

func (v *NullableCollectionOfActivityBasedTimeoutPolicy) Set(val *CollectionOfActivityBasedTimeoutPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfActivityBasedTimeoutPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfActivityBasedTimeoutPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfActivityBasedTimeoutPolicy(val *CollectionOfActivityBasedTimeoutPolicy) *NullableCollectionOfActivityBasedTimeoutPolicy {
	return &NullableCollectionOfActivityBasedTimeoutPolicy{value: val, isSet: true}
}

func (v NullableCollectionOfActivityBasedTimeoutPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfActivityBasedTimeoutPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


