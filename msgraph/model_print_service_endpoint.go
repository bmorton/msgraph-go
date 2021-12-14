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

// PrintServiceEndpoint struct for PrintServiceEndpoint
type PrintServiceEndpoint struct {
	// A human-readable display name for the endpoint.
	DisplayName *string `json:"displayName,omitempty"`
	// The URI that can be used to access the service.
	Uri *string `json:"uri,omitempty"`
}

// NewPrintServiceEndpoint instantiates a new PrintServiceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintServiceEndpoint() *PrintServiceEndpoint {
	this := PrintServiceEndpoint{}
	return &this
}

// NewPrintServiceEndpointWithDefaults instantiates a new PrintServiceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintServiceEndpointWithDefaults() *PrintServiceEndpoint {
	this := PrintServiceEndpoint{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PrintServiceEndpoint) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintServiceEndpoint) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PrintServiceEndpoint) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PrintServiceEndpoint) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PrintServiceEndpoint) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintServiceEndpoint) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PrintServiceEndpoint) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *PrintServiceEndpoint) SetUri(v string) {
	o.Uri = &v
}

func (o PrintServiceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullablePrintServiceEndpoint struct {
	value *PrintServiceEndpoint
	isSet bool
}

func (v NullablePrintServiceEndpoint) Get() *PrintServiceEndpoint {
	return v.value
}

func (v *NullablePrintServiceEndpoint) Set(val *PrintServiceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintServiceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintServiceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintServiceEndpoint(val *PrintServiceEndpoint) *NullablePrintServiceEndpoint {
	return &NullablePrintServiceEndpoint{value: val, isSet: true}
}

func (v NullablePrintServiceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintServiceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


