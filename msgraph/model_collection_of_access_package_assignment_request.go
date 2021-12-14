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

// CollectionOfAccessPackageAssignmentRequest struct for CollectionOfAccessPackageAssignmentRequest
type CollectionOfAccessPackageAssignmentRequest struct {
	Value *[]MicrosoftGraphAccessPackageAssignmentRequest `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfAccessPackageAssignmentRequest instantiates a new CollectionOfAccessPackageAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfAccessPackageAssignmentRequest() *CollectionOfAccessPackageAssignmentRequest {
	this := CollectionOfAccessPackageAssignmentRequest{}
	return &this
}

// NewCollectionOfAccessPackageAssignmentRequestWithDefaults instantiates a new CollectionOfAccessPackageAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfAccessPackageAssignmentRequestWithDefaults() *CollectionOfAccessPackageAssignmentRequest {
	this := CollectionOfAccessPackageAssignmentRequest{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfAccessPackageAssignmentRequest) GetValue() []MicrosoftGraphAccessPackageAssignmentRequest {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphAccessPackageAssignmentRequest
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAccessPackageAssignmentRequest) GetValueOk() (*[]MicrosoftGraphAccessPackageAssignmentRequest, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfAccessPackageAssignmentRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphAccessPackageAssignmentRequest and assigns it to the Value field.
func (o *CollectionOfAccessPackageAssignmentRequest) SetValue(v []MicrosoftGraphAccessPackageAssignmentRequest) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfAccessPackageAssignmentRequest) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAccessPackageAssignmentRequest) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfAccessPackageAssignmentRequest) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfAccessPackageAssignmentRequest) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfAccessPackageAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfAccessPackageAssignmentRequest struct {
	value *CollectionOfAccessPackageAssignmentRequest
	isSet bool
}

func (v NullableCollectionOfAccessPackageAssignmentRequest) Get() *CollectionOfAccessPackageAssignmentRequest {
	return v.value
}

func (v *NullableCollectionOfAccessPackageAssignmentRequest) Set(val *CollectionOfAccessPackageAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfAccessPackageAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfAccessPackageAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfAccessPackageAssignmentRequest(val *CollectionOfAccessPackageAssignmentRequest) *NullableCollectionOfAccessPackageAssignmentRequest {
	return &NullableCollectionOfAccessPackageAssignmentRequest{value: val, isSet: true}
}

func (v NullableCollectionOfAccessPackageAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfAccessPackageAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


