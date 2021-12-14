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

// MicrosoftGraphCallRecordsEndpoint struct for MicrosoftGraphCallRecordsEndpoint
type MicrosoftGraphCallRecordsEndpoint struct {
	// User-agent reported by this endpoint.
	UserAgent AnyOfmicrosoftGraphCallRecordsUserAgent `json:"userAgent,omitempty"`
}

// NewMicrosoftGraphCallRecordsEndpoint instantiates a new MicrosoftGraphCallRecordsEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphCallRecordsEndpoint() *MicrosoftGraphCallRecordsEndpoint {
	this := MicrosoftGraphCallRecordsEndpoint{}
	return &this
}

// NewMicrosoftGraphCallRecordsEndpointWithDefaults instantiates a new MicrosoftGraphCallRecordsEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphCallRecordsEndpointWithDefaults() *MicrosoftGraphCallRecordsEndpoint {
	this := MicrosoftGraphCallRecordsEndpoint{}
	return &this
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsEndpoint) GetUserAgent() AnyOfmicrosoftGraphCallRecordsUserAgent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsUserAgent
		return ret
	}
	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsEndpoint) GetUserAgentOk() (*AnyOfmicrosoftGraphCallRecordsUserAgent, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsEndpoint) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given AnyOfmicrosoftGraphCallRecordsUserAgent and assigns it to the UserAgent field.
func (o *MicrosoftGraphCallRecordsEndpoint) SetUserAgent(v AnyOfmicrosoftGraphCallRecordsUserAgent) {
	o.UserAgent = v
}

func (o MicrosoftGraphCallRecordsEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserAgent != nil {
		toSerialize["userAgent"] = o.UserAgent
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphCallRecordsEndpoint struct {
	value *MicrosoftGraphCallRecordsEndpoint
	isSet bool
}

func (v NullableMicrosoftGraphCallRecordsEndpoint) Get() *MicrosoftGraphCallRecordsEndpoint {
	return v.value
}

func (v *NullableMicrosoftGraphCallRecordsEndpoint) Set(val *MicrosoftGraphCallRecordsEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCallRecordsEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCallRecordsEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCallRecordsEndpoint(val *MicrosoftGraphCallRecordsEndpoint) *NullableMicrosoftGraphCallRecordsEndpoint {
	return &NullableMicrosoftGraphCallRecordsEndpoint{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCallRecordsEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCallRecordsEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


