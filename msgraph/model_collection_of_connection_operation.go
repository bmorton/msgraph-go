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

// CollectionOfConnectionOperation struct for CollectionOfConnectionOperation
type CollectionOfConnectionOperation struct {
	Value *[]MicrosoftGraphExternalConnectorsConnectionOperation `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfConnectionOperation instantiates a new CollectionOfConnectionOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfConnectionOperation() *CollectionOfConnectionOperation {
	this := CollectionOfConnectionOperation{}
	return &this
}

// NewCollectionOfConnectionOperationWithDefaults instantiates a new CollectionOfConnectionOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfConnectionOperationWithDefaults() *CollectionOfConnectionOperation {
	this := CollectionOfConnectionOperation{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfConnectionOperation) GetValue() []MicrosoftGraphExternalConnectorsConnectionOperation {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphExternalConnectorsConnectionOperation
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConnectionOperation) GetValueOk() (*[]MicrosoftGraphExternalConnectorsConnectionOperation, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfConnectionOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphExternalConnectorsConnectionOperation and assigns it to the Value field.
func (o *CollectionOfConnectionOperation) SetValue(v []MicrosoftGraphExternalConnectorsConnectionOperation) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfConnectionOperation) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfConnectionOperation) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfConnectionOperation) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfConnectionOperation) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfConnectionOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfConnectionOperation struct {
	value *CollectionOfConnectionOperation
	isSet bool
}

func (v NullableCollectionOfConnectionOperation) Get() *CollectionOfConnectionOperation {
	return v.value
}

func (v *NullableCollectionOfConnectionOperation) Set(val *CollectionOfConnectionOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfConnectionOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfConnectionOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfConnectionOperation(val *CollectionOfConnectionOperation) *NullableCollectionOfConnectionOperation {
	return &NullableCollectionOfConnectionOperation{value: val, isSet: true}
}

func (v NullableCollectionOfConnectionOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfConnectionOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


