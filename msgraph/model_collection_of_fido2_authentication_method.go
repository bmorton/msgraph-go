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

// CollectionOfFido2AuthenticationMethod struct for CollectionOfFido2AuthenticationMethod
type CollectionOfFido2AuthenticationMethod struct {
	Value *[]MicrosoftGraphFido2AuthenticationMethod `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfFido2AuthenticationMethod instantiates a new CollectionOfFido2AuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfFido2AuthenticationMethod() *CollectionOfFido2AuthenticationMethod {
	this := CollectionOfFido2AuthenticationMethod{}
	return &this
}

// NewCollectionOfFido2AuthenticationMethodWithDefaults instantiates a new CollectionOfFido2AuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfFido2AuthenticationMethodWithDefaults() *CollectionOfFido2AuthenticationMethod {
	this := CollectionOfFido2AuthenticationMethod{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfFido2AuthenticationMethod) GetValue() []MicrosoftGraphFido2AuthenticationMethod {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphFido2AuthenticationMethod
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfFido2AuthenticationMethod) GetValueOk() (*[]MicrosoftGraphFido2AuthenticationMethod, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfFido2AuthenticationMethod) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphFido2AuthenticationMethod and assigns it to the Value field.
func (o *CollectionOfFido2AuthenticationMethod) SetValue(v []MicrosoftGraphFido2AuthenticationMethod) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfFido2AuthenticationMethod) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfFido2AuthenticationMethod) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfFido2AuthenticationMethod) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfFido2AuthenticationMethod) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfFido2AuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfFido2AuthenticationMethod struct {
	value *CollectionOfFido2AuthenticationMethod
	isSet bool
}

func (v NullableCollectionOfFido2AuthenticationMethod) Get() *CollectionOfFido2AuthenticationMethod {
	return v.value
}

func (v *NullableCollectionOfFido2AuthenticationMethod) Set(val *CollectionOfFido2AuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfFido2AuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfFido2AuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfFido2AuthenticationMethod(val *CollectionOfFido2AuthenticationMethod) *NullableCollectionOfFido2AuthenticationMethod {
	return &NullableCollectionOfFido2AuthenticationMethod{value: val, isSet: true}
}

func (v NullableCollectionOfFido2AuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfFido2AuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


