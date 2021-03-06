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

// Bitlocker struct for Bitlocker
type Bitlocker struct {
	// The recovery keys associated with the bitlocker entity.
	RecoveryKeys *[]MicrosoftGraphBitlockerRecoveryKey `json:"recoveryKeys,omitempty"`
}

// NewBitlocker instantiates a new Bitlocker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBitlocker() *Bitlocker {
	this := Bitlocker{}
	return &this
}

// NewBitlockerWithDefaults instantiates a new Bitlocker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBitlockerWithDefaults() *Bitlocker {
	this := Bitlocker{}
	return &this
}

// GetRecoveryKeys returns the RecoveryKeys field value if set, zero value otherwise.
func (o *Bitlocker) GetRecoveryKeys() []MicrosoftGraphBitlockerRecoveryKey {
	if o == nil || o.RecoveryKeys == nil {
		var ret []MicrosoftGraphBitlockerRecoveryKey
		return ret
	}
	return *o.RecoveryKeys
}

// GetRecoveryKeysOk returns a tuple with the RecoveryKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bitlocker) GetRecoveryKeysOk() (*[]MicrosoftGraphBitlockerRecoveryKey, bool) {
	if o == nil || o.RecoveryKeys == nil {
		return nil, false
	}
	return o.RecoveryKeys, true
}

// HasRecoveryKeys returns a boolean if a field has been set.
func (o *Bitlocker) HasRecoveryKeys() bool {
	if o != nil && o.RecoveryKeys != nil {
		return true
	}

	return false
}

// SetRecoveryKeys gets a reference to the given []MicrosoftGraphBitlockerRecoveryKey and assigns it to the RecoveryKeys field.
func (o *Bitlocker) SetRecoveryKeys(v []MicrosoftGraphBitlockerRecoveryKey) {
	o.RecoveryKeys = &v
}

func (o Bitlocker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryKeys != nil {
		toSerialize["recoveryKeys"] = o.RecoveryKeys
	}
	return json.Marshal(toSerialize)
}

type NullableBitlocker struct {
	value *Bitlocker
	isSet bool
}

func (v NullableBitlocker) Get() *Bitlocker {
	return v.value
}

func (v *NullableBitlocker) Set(val *Bitlocker) {
	v.value = val
	v.isSet = true
}

func (v NullableBitlocker) IsSet() bool {
	return v.isSet
}

func (v *NullableBitlocker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBitlocker(val *Bitlocker) *NullableBitlocker {
	return &NullableBitlocker{value: val, isSet: true}
}

func (v NullableBitlocker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBitlocker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


