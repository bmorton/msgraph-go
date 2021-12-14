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

// CollectionOfWorkbookOperation struct for CollectionOfWorkbookOperation
type CollectionOfWorkbookOperation struct {
	Value *[]MicrosoftGraphWorkbookOperation `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfWorkbookOperation instantiates a new CollectionOfWorkbookOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfWorkbookOperation() *CollectionOfWorkbookOperation {
	this := CollectionOfWorkbookOperation{}
	return &this
}

// NewCollectionOfWorkbookOperationWithDefaults instantiates a new CollectionOfWorkbookOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfWorkbookOperationWithDefaults() *CollectionOfWorkbookOperation {
	this := CollectionOfWorkbookOperation{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfWorkbookOperation) GetValue() []MicrosoftGraphWorkbookOperation {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphWorkbookOperation
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfWorkbookOperation) GetValueOk() (*[]MicrosoftGraphWorkbookOperation, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfWorkbookOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphWorkbookOperation and assigns it to the Value field.
func (o *CollectionOfWorkbookOperation) SetValue(v []MicrosoftGraphWorkbookOperation) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfWorkbookOperation) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfWorkbookOperation) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfWorkbookOperation) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfWorkbookOperation) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfWorkbookOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfWorkbookOperation struct {
	value *CollectionOfWorkbookOperation
	isSet bool
}

func (v NullableCollectionOfWorkbookOperation) Get() *CollectionOfWorkbookOperation {
	return v.value
}

func (v *NullableCollectionOfWorkbookOperation) Set(val *CollectionOfWorkbookOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfWorkbookOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfWorkbookOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfWorkbookOperation(val *CollectionOfWorkbookOperation) *NullableCollectionOfWorkbookOperation {
	return &NullableCollectionOfWorkbookOperation{value: val, isSet: true}
}

func (v NullableCollectionOfWorkbookOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfWorkbookOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


