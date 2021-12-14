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

// RecordOperation struct for RecordOperation
type RecordOperation struct {
	// The access token required to retrieve the recording.
	RecordingAccessToken NullableString `json:"recordingAccessToken,omitempty"`
	// The location where the recording is located.
	RecordingLocation NullableString `json:"recordingLocation,omitempty"`
}

// NewRecordOperation instantiates a new RecordOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordOperation() *RecordOperation {
	this := RecordOperation{}
	return &this
}

// NewRecordOperationWithDefaults instantiates a new RecordOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordOperationWithDefaults() *RecordOperation {
	this := RecordOperation{}
	return &this
}

// GetRecordingAccessToken returns the RecordingAccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecordOperation) GetRecordingAccessToken() string {
	if o == nil || o.RecordingAccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecordingAccessToken.Get()
}

// GetRecordingAccessTokenOk returns a tuple with the RecordingAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordOperation) GetRecordingAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordingAccessToken.Get(), o.RecordingAccessToken.IsSet()
}

// HasRecordingAccessToken returns a boolean if a field has been set.
func (o *RecordOperation) HasRecordingAccessToken() bool {
	if o != nil && o.RecordingAccessToken.IsSet() {
		return true
	}

	return false
}

// SetRecordingAccessToken gets a reference to the given NullableString and assigns it to the RecordingAccessToken field.
func (o *RecordOperation) SetRecordingAccessToken(v string) {
	o.RecordingAccessToken.Set(&v)
}
// SetRecordingAccessTokenNil sets the value for RecordingAccessToken to be an explicit nil
func (o *RecordOperation) SetRecordingAccessTokenNil() {
	o.RecordingAccessToken.Set(nil)
}

// UnsetRecordingAccessToken ensures that no value is present for RecordingAccessToken, not even an explicit nil
func (o *RecordOperation) UnsetRecordingAccessToken() {
	o.RecordingAccessToken.Unset()
}

// GetRecordingLocation returns the RecordingLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecordOperation) GetRecordingLocation() string {
	if o == nil || o.RecordingLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecordingLocation.Get()
}

// GetRecordingLocationOk returns a tuple with the RecordingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordOperation) GetRecordingLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordingLocation.Get(), o.RecordingLocation.IsSet()
}

// HasRecordingLocation returns a boolean if a field has been set.
func (o *RecordOperation) HasRecordingLocation() bool {
	if o != nil && o.RecordingLocation.IsSet() {
		return true
	}

	return false
}

// SetRecordingLocation gets a reference to the given NullableString and assigns it to the RecordingLocation field.
func (o *RecordOperation) SetRecordingLocation(v string) {
	o.RecordingLocation.Set(&v)
}
// SetRecordingLocationNil sets the value for RecordingLocation to be an explicit nil
func (o *RecordOperation) SetRecordingLocationNil() {
	o.RecordingLocation.Set(nil)
}

// UnsetRecordingLocation ensures that no value is present for RecordingLocation, not even an explicit nil
func (o *RecordOperation) UnsetRecordingLocation() {
	o.RecordingLocation.Unset()
}

func (o RecordOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecordingAccessToken.IsSet() {
		toSerialize["recordingAccessToken"] = o.RecordingAccessToken.Get()
	}
	if o.RecordingLocation.IsSet() {
		toSerialize["recordingLocation"] = o.RecordingLocation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecordOperation struct {
	value *RecordOperation
	isSet bool
}

func (v NullableRecordOperation) Get() *RecordOperation {
	return v.value
}

func (v *NullableRecordOperation) Set(val *RecordOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordOperation(val *RecordOperation) *NullableRecordOperation {
	return &NullableRecordOperation{value: val, isSet: true}
}

func (v NullableRecordOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

