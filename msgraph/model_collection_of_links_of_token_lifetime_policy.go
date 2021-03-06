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

// CollectionOfLinksOfTokenLifetimePolicy struct for CollectionOfLinksOfTokenLifetimePolicy
type CollectionOfLinksOfTokenLifetimePolicy struct {
	Value *[]string `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfLinksOfTokenLifetimePolicy instantiates a new CollectionOfLinksOfTokenLifetimePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfLinksOfTokenLifetimePolicy() *CollectionOfLinksOfTokenLifetimePolicy {
	this := CollectionOfLinksOfTokenLifetimePolicy{}
	return &this
}

// NewCollectionOfLinksOfTokenLifetimePolicyWithDefaults instantiates a new CollectionOfLinksOfTokenLifetimePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfLinksOfTokenLifetimePolicyWithDefaults() *CollectionOfLinksOfTokenLifetimePolicy {
	this := CollectionOfLinksOfTokenLifetimePolicy{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfLinksOfTokenLifetimePolicy) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfLinksOfTokenLifetimePolicy) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfLinksOfTokenLifetimePolicy) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *CollectionOfLinksOfTokenLifetimePolicy) SetValue(v []string) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfLinksOfTokenLifetimePolicy) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfLinksOfTokenLifetimePolicy) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfLinksOfTokenLifetimePolicy) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfLinksOfTokenLifetimePolicy) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfLinksOfTokenLifetimePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfLinksOfTokenLifetimePolicy struct {
	value *CollectionOfLinksOfTokenLifetimePolicy
	isSet bool
}

func (v NullableCollectionOfLinksOfTokenLifetimePolicy) Get() *CollectionOfLinksOfTokenLifetimePolicy {
	return v.value
}

func (v *NullableCollectionOfLinksOfTokenLifetimePolicy) Set(val *CollectionOfLinksOfTokenLifetimePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfLinksOfTokenLifetimePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfLinksOfTokenLifetimePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfLinksOfTokenLifetimePolicy(val *CollectionOfLinksOfTokenLifetimePolicy) *NullableCollectionOfLinksOfTokenLifetimePolicy {
	return &NullableCollectionOfLinksOfTokenLifetimePolicy{value: val, isSet: true}
}

func (v NullableCollectionOfLinksOfTokenLifetimePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfLinksOfTokenLifetimePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


