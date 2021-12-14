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

// CollectionOfOfferShiftRequest struct for CollectionOfOfferShiftRequest
type CollectionOfOfferShiftRequest struct {
	Value *[]MicrosoftGraphOfferShiftRequest `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfOfferShiftRequest instantiates a new CollectionOfOfferShiftRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfOfferShiftRequest() *CollectionOfOfferShiftRequest {
	this := CollectionOfOfferShiftRequest{}
	return &this
}

// NewCollectionOfOfferShiftRequestWithDefaults instantiates a new CollectionOfOfferShiftRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfOfferShiftRequestWithDefaults() *CollectionOfOfferShiftRequest {
	this := CollectionOfOfferShiftRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfOfferShiftRequest) GetValue() []MicrosoftGraphOfferShiftRequest {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphOfferShiftRequest
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOfferShiftRequest) GetValueOk() (*[]MicrosoftGraphOfferShiftRequest, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfOfferShiftRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphOfferShiftRequest and assigns it to the Value field.
func (o *CollectionOfOfferShiftRequest) SetValue(v []MicrosoftGraphOfferShiftRequest) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfOfferShiftRequest) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfOfferShiftRequest) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfOfferShiftRequest) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfOfferShiftRequest) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfOfferShiftRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfOfferShiftRequest struct {
	value *CollectionOfOfferShiftRequest
	isSet bool
}

func (v NullableCollectionOfOfferShiftRequest) Get() *CollectionOfOfferShiftRequest {
	return v.value
}

func (v *NullableCollectionOfOfferShiftRequest) Set(val *CollectionOfOfferShiftRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfOfferShiftRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfOfferShiftRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfOfferShiftRequest(val *CollectionOfOfferShiftRequest) *NullableCollectionOfOfferShiftRequest {
	return &NullableCollectionOfOfferShiftRequest{value: val, isSet: true}
}

func (v NullableCollectionOfOfferShiftRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfOfferShiftRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


