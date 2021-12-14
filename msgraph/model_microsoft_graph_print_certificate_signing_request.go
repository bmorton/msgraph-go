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

// MicrosoftGraphPrintCertificateSigningRequest struct for MicrosoftGraphPrintCertificateSigningRequest
type MicrosoftGraphPrintCertificateSigningRequest struct {
	// A base64-encoded pkcs10 certificate request. Read-only.
	Content *string `json:"content,omitempty"`
	// The base64-encoded public portion of an asymmetric key that is generated by the client. Read-only.
	TransportKey *string `json:"transportKey,omitempty"`
}

// NewMicrosoftGraphPrintCertificateSigningRequest instantiates a new MicrosoftGraphPrintCertificateSigningRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintCertificateSigningRequest() *MicrosoftGraphPrintCertificateSigningRequest {
	this := MicrosoftGraphPrintCertificateSigningRequest{}
	return &this
}

// NewMicrosoftGraphPrintCertificateSigningRequestWithDefaults instantiates a new MicrosoftGraphPrintCertificateSigningRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintCertificateSigningRequestWithDefaults() *MicrosoftGraphPrintCertificateSigningRequest {
	this := MicrosoftGraphPrintCertificateSigningRequest{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintCertificateSigningRequest) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintCertificateSigningRequest) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintCertificateSigningRequest) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *MicrosoftGraphPrintCertificateSigningRequest) SetContent(v string) {
	o.Content = &v
}

// GetTransportKey returns the TransportKey field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintCertificateSigningRequest) GetTransportKey() string {
	if o == nil || o.TransportKey == nil {
		var ret string
		return ret
	}
	return *o.TransportKey
}

// GetTransportKeyOk returns a tuple with the TransportKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintCertificateSigningRequest) GetTransportKeyOk() (*string, bool) {
	if o == nil || o.TransportKey == nil {
		return nil, false
	}
	return o.TransportKey, true
}

// HasTransportKey returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintCertificateSigningRequest) HasTransportKey() bool {
	if o != nil && o.TransportKey != nil {
		return true
	}

	return false
}

// SetTransportKey gets a reference to the given string and assigns it to the TransportKey field.
func (o *MicrosoftGraphPrintCertificateSigningRequest) SetTransportKey(v string) {
	o.TransportKey = &v
}

func (o MicrosoftGraphPrintCertificateSigningRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.TransportKey != nil {
		toSerialize["transportKey"] = o.TransportKey
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintCertificateSigningRequest struct {
	value *MicrosoftGraphPrintCertificateSigningRequest
	isSet bool
}

func (v NullableMicrosoftGraphPrintCertificateSigningRequest) Get() *MicrosoftGraphPrintCertificateSigningRequest {
	return v.value
}

func (v *NullableMicrosoftGraphPrintCertificateSigningRequest) Set(val *MicrosoftGraphPrintCertificateSigningRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintCertificateSigningRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintCertificateSigningRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintCertificateSigningRequest(val *MicrosoftGraphPrintCertificateSigningRequest) *NullableMicrosoftGraphPrintCertificateSigningRequest {
	return &NullableMicrosoftGraphPrintCertificateSigningRequest{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintCertificateSigningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintCertificateSigningRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


