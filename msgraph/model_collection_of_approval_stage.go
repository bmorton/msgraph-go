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

// CollectionOfApprovalStage struct for CollectionOfApprovalStage
type CollectionOfApprovalStage struct {
	Value *[]MicrosoftGraphApprovalStage `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfApprovalStage instantiates a new CollectionOfApprovalStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfApprovalStage() *CollectionOfApprovalStage {
	this := CollectionOfApprovalStage{}
	return &this
}

// NewCollectionOfApprovalStageWithDefaults instantiates a new CollectionOfApprovalStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfApprovalStageWithDefaults() *CollectionOfApprovalStage {
	this := CollectionOfApprovalStage{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfApprovalStage) GetValue() []MicrosoftGraphApprovalStage {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphApprovalStage
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfApprovalStage) GetValueOk() (*[]MicrosoftGraphApprovalStage, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfApprovalStage) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphApprovalStage and assigns it to the Value field.
func (o *CollectionOfApprovalStage) SetValue(v []MicrosoftGraphApprovalStage) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfApprovalStage) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfApprovalStage) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfApprovalStage) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfApprovalStage) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfApprovalStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfApprovalStage struct {
	value *CollectionOfApprovalStage
	isSet bool
}

func (v NullableCollectionOfApprovalStage) Get() *CollectionOfApprovalStage {
	return v.value
}

func (v *NullableCollectionOfApprovalStage) Set(val *CollectionOfApprovalStage) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfApprovalStage) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfApprovalStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfApprovalStage(val *CollectionOfApprovalStage) *NullableCollectionOfApprovalStage {
	return &NullableCollectionOfApprovalStage{value: val, isSet: true}
}

func (v NullableCollectionOfApprovalStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfApprovalStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


