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

// CollectionOfOpenShiftChangeRequest struct for CollectionOfOpenShiftChangeRequest
type CollectionOfOpenShiftChangeRequest struct {
	Value *[]MicrosoftGraphOpenShiftChangeRequest `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOpenShiftChangeRequest instantiates a new CollectionOfOpenShiftChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOpenShiftChangeRequest() *CollectionOfOpenShiftChangeRequest {
	this := CollectionOfOpenShiftChangeRequest{}
	return &this
}

// NewCollectionOfOpenShiftChangeRequestWithDefaults instantiates a new CollectionOfOpenShiftChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOpenShiftChangeRequestWithDefaults() *CollectionOfOpenShiftChangeRequest {
	this := CollectionOfOpenShiftChangeRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOpenShiftChangeRequest) GetValue() []MicrosoftGraphOpenShiftChangeRequest {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOpenShiftChangeRequest
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOpenShiftChangeRequest) GetValueOk() (*[]MicrosoftGraphOpenShiftChangeRequest, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOpenShiftChangeRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOpenShiftChangeRequest and assigns it to the Value field.
func (o *CollectionOfOpenShiftChangeRequest) SetValue(v []MicrosoftGraphOpenShiftChangeRequest) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOpenShiftChangeRequest) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOpenShiftChangeRequest) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOpenShiftChangeRequest) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOpenShiftChangeRequest) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOpenShiftChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOpenShiftChangeRequest struct {
	value *CollectionOfOpenShiftChangeRequest
	isSet bool
}

func (v NullableCollectionOfOpenShiftChangeRequest) Get() *CollectionOfOpenShiftChangeRequest {
	return v.value
}

func (v *NullableCollectionOfOpenShiftChangeRequest) Set(val *CollectionOfOpenShiftChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOpenShiftChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOpenShiftChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOpenShiftChangeRequest(val *CollectionOfOpenShiftChangeRequest) *NullableCollectionOfOpenShiftChangeRequest {
	return &NullableCollectionOfOpenShiftChangeRequest{value: val, isSet: true}
}

func (v NullableCollectionOfOpenShiftChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOpenShiftChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


