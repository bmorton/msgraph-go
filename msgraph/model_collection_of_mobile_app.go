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

// CollectionOfMobileApp struct for CollectionOfMobileApp
type CollectionOfMobileApp struct {
	Value *[]MicrosoftGraphMobileApp `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfMobileApp instantiates a new CollectionOfMobileApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfMobileApp() *CollectionOfMobileApp {
	this := CollectionOfMobileApp{}
	return &this
}

// NewCollectionOfMobileAppWithDefaults instantiates a new CollectionOfMobileApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfMobileAppWithDefaults() *CollectionOfMobileApp {
	this := CollectionOfMobileApp{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfMobileApp) GetValue() []MicrosoftGraphMobileApp {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphMobileApp
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMobileApp) GetValueOk() (*[]MicrosoftGraphMobileApp, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfMobileApp) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphMobileApp and assigns it to the Value field.
func (o *CollectionOfMobileApp) SetValue(v []MicrosoftGraphMobileApp) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfMobileApp) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMobileApp) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfMobileApp) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfMobileApp) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfMobileApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfMobileApp struct {
	value *CollectionOfMobileApp
	isSet bool
}

func (v NullableCollectionOfMobileApp) Get() *CollectionOfMobileApp {
	return v.value
}

func (v *NullableCollectionOfMobileApp) Set(val *CollectionOfMobileApp) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfMobileApp) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfMobileApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfMobileApp(val *CollectionOfMobileApp) *NullableCollectionOfMobileApp {
	return &NullableCollectionOfMobileApp{value: val, isSet: true}
}

func (v NullableCollectionOfMobileApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfMobileApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

