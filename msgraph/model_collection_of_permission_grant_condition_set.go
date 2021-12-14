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

// CollectionOfPermissionGrantConditionSet struct for CollectionOfPermissionGrantConditionSet
type CollectionOfPermissionGrantConditionSet struct {
	Value *[]MicrosoftGraphPermissionGrantConditionSet `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfPermissionGrantConditionSet instantiates a new CollectionOfPermissionGrantConditionSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfPermissionGrantConditionSet() *CollectionOfPermissionGrantConditionSet {
	this := CollectionOfPermissionGrantConditionSet{}
	return &this
}

// NewCollectionOfPermissionGrantConditionSetWithDefaults instantiates a new CollectionOfPermissionGrantConditionSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfPermissionGrantConditionSetWithDefaults() *CollectionOfPermissionGrantConditionSet {
	this := CollectionOfPermissionGrantConditionSet{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfPermissionGrantConditionSet) GetValue() []MicrosoftGraphPermissionGrantConditionSet {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphPermissionGrantConditionSet
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPermissionGrantConditionSet) GetValueOk() (*[]MicrosoftGraphPermissionGrantConditionSet, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfPermissionGrantConditionSet) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphPermissionGrantConditionSet and assigns it to the Value field.
func (o *CollectionOfPermissionGrantConditionSet) SetValue(v []MicrosoftGraphPermissionGrantConditionSet) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfPermissionGrantConditionSet) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPermissionGrantConditionSet) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfPermissionGrantConditionSet) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfPermissionGrantConditionSet) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfPermissionGrantConditionSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfPermissionGrantConditionSet struct {
	value *CollectionOfPermissionGrantConditionSet
	isSet bool
}

func (v NullableCollectionOfPermissionGrantConditionSet) Get() *CollectionOfPermissionGrantConditionSet {
	return v.value
}

func (v *NullableCollectionOfPermissionGrantConditionSet) Set(val *CollectionOfPermissionGrantConditionSet) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfPermissionGrantConditionSet) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfPermissionGrantConditionSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfPermissionGrantConditionSet(val *CollectionOfPermissionGrantConditionSet) *NullableCollectionOfPermissionGrantConditionSet {
	return &NullableCollectionOfPermissionGrantConditionSet{value: val, isSet: true}
}

func (v NullableCollectionOfPermissionGrantConditionSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfPermissionGrantConditionSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


