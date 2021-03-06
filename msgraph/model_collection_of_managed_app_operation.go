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

// CollectionOfManagedAppOperation struct for CollectionOfManagedAppOperation
type CollectionOfManagedAppOperation struct {
	Value *[]MicrosoftGraphManagedAppOperation `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfManagedAppOperation instantiates a new CollectionOfManagedAppOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfManagedAppOperation() *CollectionOfManagedAppOperation {
	this := CollectionOfManagedAppOperation{}
	return &this
}

// NewCollectionOfManagedAppOperationWithDefaults instantiates a new CollectionOfManagedAppOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfManagedAppOperationWithDefaults() *CollectionOfManagedAppOperation {
	this := CollectionOfManagedAppOperation{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfManagedAppOperation) GetValue() []MicrosoftGraphManagedAppOperation {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphManagedAppOperation
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfManagedAppOperation) GetValueOk() (*[]MicrosoftGraphManagedAppOperation, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfManagedAppOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphManagedAppOperation and assigns it to the Value field.
func (o *CollectionOfManagedAppOperation) SetValue(v []MicrosoftGraphManagedAppOperation) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfManagedAppOperation) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfManagedAppOperation) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfManagedAppOperation) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfManagedAppOperation) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfManagedAppOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfManagedAppOperation struct {
	value *CollectionOfManagedAppOperation
	isSet bool
}

func (v NullableCollectionOfManagedAppOperation) Get() *CollectionOfManagedAppOperation {
	return v.value
}

func (v *NullableCollectionOfManagedAppOperation) Set(val *CollectionOfManagedAppOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfManagedAppOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfManagedAppOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfManagedAppOperation(val *CollectionOfManagedAppOperation) *NullableCollectionOfManagedAppOperation {
	return &NullableCollectionOfManagedAppOperation{value: val, isSet: true}
}

func (v NullableCollectionOfManagedAppOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfManagedAppOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


