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

// MicrosoftGraphPrintService struct for MicrosoftGraphPrintService
type MicrosoftGraphPrintService struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Endpoints that can be used to access the service. Read-only. Nullable.
	Endpoints *[]MicrosoftGraphPrintServiceEndpoint `json:"endpoints,omitempty"`
}

// NewMicrosoftGraphPrintService instantiates a new MicrosoftGraphPrintService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintService() *MicrosoftGraphPrintService {
	this := MicrosoftGraphPrintService{}
	return &this
}

// NewMicrosoftGraphPrintServiceWithDefaults instantiates a new MicrosoftGraphPrintService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintServiceWithDefaults() *MicrosoftGraphPrintService {
	this := MicrosoftGraphPrintService{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintService) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintService) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintService) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPrintService) SetId(v string) {
	o.Id = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintService) GetEndpoints() []MicrosoftGraphPrintServiceEndpoint {
	if o == nil || o.Endpoints == nil {
		var ret []MicrosoftGraphPrintServiceEndpoint
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintService) GetEndpointsOk() (*[]MicrosoftGraphPrintServiceEndpoint, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintService) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []MicrosoftGraphPrintServiceEndpoint and assigns it to the Endpoints field.
func (o *MicrosoftGraphPrintService) SetEndpoints(v []MicrosoftGraphPrintServiceEndpoint) {
	o.Endpoints = &v
}

func (o MicrosoftGraphPrintService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintService struct {
	value *MicrosoftGraphPrintService
	isSet bool
}

func (v NullableMicrosoftGraphPrintService) Get() *MicrosoftGraphPrintService {
	return v.value
}

func (v *NullableMicrosoftGraphPrintService) Set(val *MicrosoftGraphPrintService) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintService) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintService(val *MicrosoftGraphPrintService) *NullableMicrosoftGraphPrintService {
	return &NullableMicrosoftGraphPrintService{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


