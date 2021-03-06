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

// CollectionOfAlert struct for CollectionOfAlert
type CollectionOfAlert struct {
	Value *[]MicrosoftGraphAlert `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfAlert instantiates a new CollectionOfAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfAlert() *CollectionOfAlert {
	this := CollectionOfAlert{}
	return &this
}

// NewCollectionOfAlertWithDefaults instantiates a new CollectionOfAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfAlertWithDefaults() *CollectionOfAlert {
	this := CollectionOfAlert{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfAlert) GetValue() []MicrosoftGraphAlert {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphAlert
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAlert) GetValueOk() (*[]MicrosoftGraphAlert, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfAlert) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphAlert and assigns it to the Value field.
func (o *CollectionOfAlert) SetValue(v []MicrosoftGraphAlert) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfAlert) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAlert) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfAlert) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfAlert) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfAlert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfAlert struct {
	value *CollectionOfAlert
	isSet bool
}

func (v NullableCollectionOfAlert) Get() *CollectionOfAlert {
	return v.value
}

func (v *NullableCollectionOfAlert) Set(val *CollectionOfAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfAlert(val *CollectionOfAlert) *NullableCollectionOfAlert {
	return &NullableCollectionOfAlert{value: val, isSet: true}
}

func (v NullableCollectionOfAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


