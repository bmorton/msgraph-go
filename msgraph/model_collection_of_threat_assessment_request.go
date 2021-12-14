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

// CollectionOfThreatAssessmentRequest struct for CollectionOfThreatAssessmentRequest
type CollectionOfThreatAssessmentRequest struct {
	Value *[]MicrosoftGraphThreatAssessmentRequest `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfThreatAssessmentRequest instantiates a new CollectionOfThreatAssessmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfThreatAssessmentRequest() *CollectionOfThreatAssessmentRequest {
	this := CollectionOfThreatAssessmentRequest{}
	return &this
}

// NewCollectionOfThreatAssessmentRequestWithDefaults instantiates a new CollectionOfThreatAssessmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfThreatAssessmentRequestWithDefaults() *CollectionOfThreatAssessmentRequest {
	this := CollectionOfThreatAssessmentRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfThreatAssessmentRequest) GetValue() []MicrosoftGraphThreatAssessmentRequest {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphThreatAssessmentRequest
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfThreatAssessmentRequest) GetValueOk() (*[]MicrosoftGraphThreatAssessmentRequest, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfThreatAssessmentRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphThreatAssessmentRequest and assigns it to the Value field.
func (o *CollectionOfThreatAssessmentRequest) SetValue(v []MicrosoftGraphThreatAssessmentRequest) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfThreatAssessmentRequest) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfThreatAssessmentRequest) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfThreatAssessmentRequest) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfThreatAssessmentRequest) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfThreatAssessmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfThreatAssessmentRequest struct {
	value *CollectionOfThreatAssessmentRequest
	isSet bool
}

func (v NullableCollectionOfThreatAssessmentRequest) Get() *CollectionOfThreatAssessmentRequest {
	return v.value
}

func (v *NullableCollectionOfThreatAssessmentRequest) Set(val *CollectionOfThreatAssessmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfThreatAssessmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfThreatAssessmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfThreatAssessmentRequest(val *CollectionOfThreatAssessmentRequest) *NullableCollectionOfThreatAssessmentRequest {
	return &NullableCollectionOfThreatAssessmentRequest{value: val, isSet: true}
}

func (v NullableCollectionOfThreatAssessmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfThreatAssessmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

