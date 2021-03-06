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

// CollectionOfUserConsentRequest struct for CollectionOfUserConsentRequest
type CollectionOfUserConsentRequest struct {
	Value *[]MicrosoftGraphUserConsentRequest `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfUserConsentRequest instantiates a new CollectionOfUserConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfUserConsentRequest() *CollectionOfUserConsentRequest {
	this := CollectionOfUserConsentRequest{}
	return &this
}

// NewCollectionOfUserConsentRequestWithDefaults instantiates a new CollectionOfUserConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfUserConsentRequestWithDefaults() *CollectionOfUserConsentRequest {
	this := CollectionOfUserConsentRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfUserConsentRequest) GetValue() []MicrosoftGraphUserConsentRequest {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphUserConsentRequest
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserConsentRequest) GetValueOk() (*[]MicrosoftGraphUserConsentRequest, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfUserConsentRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphUserConsentRequest and assigns it to the Value field.
func (o *CollectionOfUserConsentRequest) SetValue(v []MicrosoftGraphUserConsentRequest) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfUserConsentRequest) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserConsentRequest) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfUserConsentRequest) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfUserConsentRequest) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfUserConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfUserConsentRequest struct {
	value *CollectionOfUserConsentRequest
	isSet bool
}

func (v NullableCollectionOfUserConsentRequest) Get() *CollectionOfUserConsentRequest {
	return v.value
}

func (v *NullableCollectionOfUserConsentRequest) Set(val *CollectionOfUserConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfUserConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfUserConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfUserConsentRequest(val *CollectionOfUserConsentRequest) *NullableCollectionOfUserConsentRequest {
	return &NullableCollectionOfUserConsentRequest{value: val, isSet: true}
}

func (v NullableCollectionOfUserConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfUserConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


