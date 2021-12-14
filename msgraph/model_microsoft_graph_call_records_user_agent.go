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

// MicrosoftGraphCallRecordsUserAgent struct for MicrosoftGraphCallRecordsUserAgent
type MicrosoftGraphCallRecordsUserAgent struct {
	// Identifies the version of application software used by this endpoint.
	ApplicationVersion NullableString `json:"applicationVersion,omitempty"`
	// User-agent header value reported by this endpoint.
	HeaderValue NullableString `json:"headerValue,omitempty"`
}

// NewMicrosoftGraphCallRecordsUserAgent instantiates a new MicrosoftGraphCallRecordsUserAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphCallRecordsUserAgent() *MicrosoftGraphCallRecordsUserAgent {
	this := MicrosoftGraphCallRecordsUserAgent{}
	return &this
}

// NewMicrosoftGraphCallRecordsUserAgentWithDefaults instantiates a new MicrosoftGraphCallRecordsUserAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphCallRecordsUserAgentWithDefaults() *MicrosoftGraphCallRecordsUserAgent {
	this := MicrosoftGraphCallRecordsUserAgent{}
	return &this
}

// GetApplicationVersion returns the ApplicationVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsUserAgent) GetApplicationVersion() string {
	if o == nil || o.ApplicationVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationVersion.Get()
}

// GetApplicationVersionOk returns a tuple with the ApplicationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsUserAgent) GetApplicationVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationVersion.Get(), o.ApplicationVersion.IsSet()
}

// HasApplicationVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsUserAgent) HasApplicationVersion() bool {
	if o != nil && o.ApplicationVersion.IsSet() {
		return true
	}

	return false
}

// SetApplicationVersion gets a reference to the given NullableString and assigns it to the ApplicationVersion field.
func (o *MicrosoftGraphCallRecordsUserAgent) SetApplicationVersion(v string) {
	o.ApplicationVersion.Set(&v)
}
// SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil
func (o *MicrosoftGraphCallRecordsUserAgent) SetApplicationVersionNil() {
	o.ApplicationVersion.Set(nil)
}

// UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
func (o *MicrosoftGraphCallRecordsUserAgent) UnsetApplicationVersion() {
	o.ApplicationVersion.Unset()
}

// GetHeaderValue returns the HeaderValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsUserAgent) GetHeaderValue() string {
	if o == nil || o.HeaderValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderValue.Get()
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsUserAgent) GetHeaderValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HeaderValue.Get(), o.HeaderValue.IsSet()
}

// HasHeaderValue returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsUserAgent) HasHeaderValue() bool {
	if o != nil && o.HeaderValue.IsSet() {
		return true
	}

	return false
}

// SetHeaderValue gets a reference to the given NullableString and assigns it to the HeaderValue field.
func (o *MicrosoftGraphCallRecordsUserAgent) SetHeaderValue(v string) {
	o.HeaderValue.Set(&v)
}
// SetHeaderValueNil sets the value for HeaderValue to be an explicit nil
func (o *MicrosoftGraphCallRecordsUserAgent) SetHeaderValueNil() {
	o.HeaderValue.Set(nil)
}

// UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
func (o *MicrosoftGraphCallRecordsUserAgent) UnsetHeaderValue() {
	o.HeaderValue.Unset()
}

func (o MicrosoftGraphCallRecordsUserAgent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationVersion.IsSet() {
		toSerialize["applicationVersion"] = o.ApplicationVersion.Get()
	}
	if o.HeaderValue.IsSet() {
		toSerialize["headerValue"] = o.HeaderValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphCallRecordsUserAgent struct {
	value *MicrosoftGraphCallRecordsUserAgent
	isSet bool
}

func (v NullableMicrosoftGraphCallRecordsUserAgent) Get() *MicrosoftGraphCallRecordsUserAgent {
	return v.value
}

func (v *NullableMicrosoftGraphCallRecordsUserAgent) Set(val *MicrosoftGraphCallRecordsUserAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCallRecordsUserAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCallRecordsUserAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCallRecordsUserAgent(val *MicrosoftGraphCallRecordsUserAgent) *NullableMicrosoftGraphCallRecordsUserAgent {
	return &NullableMicrosoftGraphCallRecordsUserAgent{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCallRecordsUserAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCallRecordsUserAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


