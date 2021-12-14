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

// InlineObject721 struct for InlineObject721
type InlineObject721 struct {
	KeyId *string `json:"keyId,omitempty"`
	Proof *string `json:"proof,omitempty"`
}

// NewInlineObject721 instantiates a new InlineObject721 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject721() *InlineObject721 {
	this := InlineObject721{}
	return &this
}

// NewInlineObject721WithDefaults instantiates a new InlineObject721 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject721WithDefaults() *InlineObject721 {
	this := InlineObject721{}
	return &this
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *InlineObject721) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject721) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *InlineObject721) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *InlineObject721) SetKeyId(v string) {
	o.KeyId = &v
}

// GetProof returns the Proof field value if set, zero value otherwise.
func (o *InlineObject721) GetProof() string {
	if o == nil || o.Proof == nil {
		var ret string
		return ret
	}
	return *o.Proof
}

// GetProofOk returns a tuple with the Proof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject721) GetProofOk() (*string, bool) {
	if o == nil || o.Proof == nil {
		return nil, false
	}
	return o.Proof, true
}

// HasProof returns a boolean if a field has been set.
func (o *InlineObject721) HasProof() bool {
	if o != nil && o.Proof != nil {
		return true
	}

	return false
}

// SetProof gets a reference to the given string and assigns it to the Proof field.
func (o *InlineObject721) SetProof(v string) {
	o.Proof = &v
}

func (o InlineObject721) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyId != nil {
		toSerialize["keyId"] = o.KeyId
	}
	if o.Proof != nil {
		toSerialize["proof"] = o.Proof
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject721 struct {
	value *InlineObject721
	isSet bool
}

func (v NullableInlineObject721) Get() *InlineObject721 {
	return v.value
}

func (v *NullableInlineObject721) Set(val *InlineObject721) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject721) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject721) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject721(val *InlineObject721) *NullableInlineObject721 {
	return &NullableInlineObject721{value: val, isSet: true}
}

func (v NullableInlineObject721) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject721) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


