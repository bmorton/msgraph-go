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

// CollectionOfTeamsTab struct for CollectionOfTeamsTab
type CollectionOfTeamsTab struct {
	Value *[]MicrosoftGraphTeamsTab `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfTeamsTab instantiates a new CollectionOfTeamsTab object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfTeamsTab() *CollectionOfTeamsTab {
	this := CollectionOfTeamsTab{}
	return &this
}

// NewCollectionOfTeamsTabWithDefaults instantiates a new CollectionOfTeamsTab object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfTeamsTabWithDefaults() *CollectionOfTeamsTab {
	this := CollectionOfTeamsTab{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfTeamsTab) GetValue() []MicrosoftGraphTeamsTab {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTeamsTab
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTeamsTab) GetValueOk() (*[]MicrosoftGraphTeamsTab, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfTeamsTab) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTeamsTab and assigns it to the Value field.
func (o *CollectionOfTeamsTab) SetValue(v []MicrosoftGraphTeamsTab) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfTeamsTab) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTeamsTab) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfTeamsTab) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfTeamsTab) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfTeamsTab) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfTeamsTab struct {
	value *CollectionOfTeamsTab
	isSet bool
}

func (v NullableCollectionOfTeamsTab) Get() *CollectionOfTeamsTab {
	return v.value
}

func (v *NullableCollectionOfTeamsTab) Set(val *CollectionOfTeamsTab) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfTeamsTab) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfTeamsTab) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfTeamsTab(val *CollectionOfTeamsTab) *NullableCollectionOfTeamsTab {
	return &NullableCollectionOfTeamsTab{value: val, isSet: true}
}

func (v NullableCollectionOfTeamsTab) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfTeamsTab) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


