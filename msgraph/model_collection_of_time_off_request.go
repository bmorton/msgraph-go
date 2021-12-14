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

// CollectionOfTimeOffRequest struct for CollectionOfTimeOffRequest
type CollectionOfTimeOffRequest struct {
	Value *[]MicrosoftGraphTimeOffRequest `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfTimeOffRequest instantiates a new CollectionOfTimeOffRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfTimeOffRequest() *CollectionOfTimeOffRequest {
	this := CollectionOfTimeOffRequest{}
	return &this
}

// NewCollectionOfTimeOffRequestWithDefaults instantiates a new CollectionOfTimeOffRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfTimeOffRequestWithDefaults() *CollectionOfTimeOffRequest {
	this := CollectionOfTimeOffRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfTimeOffRequest) GetValue() []MicrosoftGraphTimeOffRequest {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphTimeOffRequest
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTimeOffRequest) GetValueOk() (*[]MicrosoftGraphTimeOffRequest, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfTimeOffRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphTimeOffRequest and assigns it to the Value field.
func (o *CollectionOfTimeOffRequest) SetValue(v []MicrosoftGraphTimeOffRequest) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfTimeOffRequest) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfTimeOffRequest) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfTimeOffRequest) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfTimeOffRequest) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfTimeOffRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfTimeOffRequest struct {
	value *CollectionOfTimeOffRequest
	isSet bool
}

func (v NullableCollectionOfTimeOffRequest) Get() *CollectionOfTimeOffRequest {
	return v.value
}

func (v *NullableCollectionOfTimeOffRequest) Set(val *CollectionOfTimeOffRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfTimeOffRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfTimeOffRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfTimeOffRequest(val *CollectionOfTimeOffRequest) *NullableCollectionOfTimeOffRequest {
	return &NullableCollectionOfTimeOffRequest{value: val, isSet: true}
}

func (v NullableCollectionOfTimeOffRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfTimeOffRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


