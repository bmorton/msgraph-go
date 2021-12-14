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

// CollectionOfThreatAssessmentResult struct for CollectionOfThreatAssessmentResult
type CollectionOfThreatAssessmentResult struct {
	Value *[]MicrosoftGraphThreatAssessmentResult `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfThreatAssessmentResult instantiates a new CollectionOfThreatAssessmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfThreatAssessmentResult() *CollectionOfThreatAssessmentResult {
	this := CollectionOfThreatAssessmentResult{}
	return &this
}

// NewCollectionOfThreatAssessmentResultWithDefaults instantiates a new CollectionOfThreatAssessmentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfThreatAssessmentResultWithDefaults() *CollectionOfThreatAssessmentResult {
	this := CollectionOfThreatAssessmentResult{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfThreatAssessmentResult) GetValue() []MicrosoftGraphThreatAssessmentResult {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphThreatAssessmentResult
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfThreatAssessmentResult) GetValueOk() (*[]MicrosoftGraphThreatAssessmentResult, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfThreatAssessmentResult) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphThreatAssessmentResult and assigns it to the Value field.
func (o *CollectionOfThreatAssessmentResult) SetValue(v []MicrosoftGraphThreatAssessmentResult) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfThreatAssessmentResult) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfThreatAssessmentResult) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfThreatAssessmentResult) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfThreatAssessmentResult) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfThreatAssessmentResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfThreatAssessmentResult struct {
	value *CollectionOfThreatAssessmentResult
	isSet bool
}

func (v NullableCollectionOfThreatAssessmentResult) Get() *CollectionOfThreatAssessmentResult {
	return v.value
}

func (v *NullableCollectionOfThreatAssessmentResult) Set(val *CollectionOfThreatAssessmentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfThreatAssessmentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfThreatAssessmentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfThreatAssessmentResult(val *CollectionOfThreatAssessmentResult) *NullableCollectionOfThreatAssessmentResult {
	return &NullableCollectionOfThreatAssessmentResult{value: val, isSet: true}
}

func (v NullableCollectionOfThreatAssessmentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfThreatAssessmentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

