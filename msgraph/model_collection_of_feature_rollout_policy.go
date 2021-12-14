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

// CollectionOfFeatureRolloutPolicy struct for CollectionOfFeatureRolloutPolicy
type CollectionOfFeatureRolloutPolicy struct {
	Value *[]MicrosoftGraphFeatureRolloutPolicy `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfFeatureRolloutPolicy instantiates a new CollectionOfFeatureRolloutPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfFeatureRolloutPolicy() *CollectionOfFeatureRolloutPolicy {
	this := CollectionOfFeatureRolloutPolicy{}
	return &this
}

// NewCollectionOfFeatureRolloutPolicyWithDefaults instantiates a new CollectionOfFeatureRolloutPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfFeatureRolloutPolicyWithDefaults() *CollectionOfFeatureRolloutPolicy {
	this := CollectionOfFeatureRolloutPolicy{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfFeatureRolloutPolicy) GetValue() []MicrosoftGraphFeatureRolloutPolicy {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphFeatureRolloutPolicy
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfFeatureRolloutPolicy) GetValueOk() (*[]MicrosoftGraphFeatureRolloutPolicy, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfFeatureRolloutPolicy) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphFeatureRolloutPolicy and assigns it to the Value field.
func (o *CollectionOfFeatureRolloutPolicy) SetValue(v []MicrosoftGraphFeatureRolloutPolicy) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfFeatureRolloutPolicy) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfFeatureRolloutPolicy) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfFeatureRolloutPolicy) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfFeatureRolloutPolicy) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfFeatureRolloutPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfFeatureRolloutPolicy struct {
	value *CollectionOfFeatureRolloutPolicy
	isSet bool
}

func (v NullableCollectionOfFeatureRolloutPolicy) Get() *CollectionOfFeatureRolloutPolicy {
	return v.value
}

func (v *NullableCollectionOfFeatureRolloutPolicy) Set(val *CollectionOfFeatureRolloutPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfFeatureRolloutPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfFeatureRolloutPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfFeatureRolloutPolicy(val *CollectionOfFeatureRolloutPolicy) *NullableCollectionOfFeatureRolloutPolicy {
	return &NullableCollectionOfFeatureRolloutPolicy{value: val, isSet: true}
}

func (v NullableCollectionOfFeatureRolloutPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfFeatureRolloutPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


