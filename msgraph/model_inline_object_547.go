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

// InlineObject547 struct for InlineObject547
type InlineObject547 struct {
	DestinationPrinterId *string `json:"destinationPrinterId,omitempty"`
	Configuration AnyOfmicrosoftGraphPrintJobConfiguration `json:"configuration,omitempty"`
}

// NewInlineObject547 instantiates a new InlineObject547 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject547() *InlineObject547 {
	this := InlineObject547{}
	return &this
}

// NewInlineObject547WithDefaults instantiates a new InlineObject547 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject547WithDefaults() *InlineObject547 {
	this := InlineObject547{}
	return &this
}

// GetDestinationPrinterId returns the DestinationPrinterId field value if set, zero value otherwise.
func (o *InlineObject547) GetDestinationPrinterId() string {
	if o == nil || o.DestinationPrinterId == nil {
		var ret string
		return ret
	}
	return *o.DestinationPrinterId
}

// GetDestinationPrinterIdOk returns a tuple with the DestinationPrinterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject547) GetDestinationPrinterIdOk() (*string, bool) {
	if o == nil || o.DestinationPrinterId == nil {
		return nil, false
	}
	return o.DestinationPrinterId, true
}

// HasDestinationPrinterId returns a boolean if a field has been set.
func (o *InlineObject547) HasDestinationPrinterId() bool {
	if o != nil && o.DestinationPrinterId != nil {
		return true
	}

	return false
}

// SetDestinationPrinterId gets a reference to the given string and assigns it to the DestinationPrinterId field.
func (o *InlineObject547) SetDestinationPrinterId(v string) {
	o.DestinationPrinterId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject547) GetConfiguration() AnyOfmicrosoftGraphPrintJobConfiguration {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintJobConfiguration
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject547) GetConfigurationOk() (*AnyOfmicrosoftGraphPrintJobConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *InlineObject547) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given AnyOfmicrosoftGraphPrintJobConfiguration and assigns it to the Configuration field.
func (o *InlineObject547) SetConfiguration(v AnyOfmicrosoftGraphPrintJobConfiguration) {
	o.Configuration = v
}

func (o InlineObject547) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationPrinterId != nil {
		toSerialize["destinationPrinterId"] = o.DestinationPrinterId
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject547 struct {
	value *InlineObject547
	isSet bool
}

func (v NullableInlineObject547) Get() *InlineObject547 {
	return v.value
}

func (v *NullableInlineObject547) Set(val *InlineObject547) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject547) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject547) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject547(val *InlineObject547) *NullableInlineObject547 {
	return &NullableInlineObject547{value: val, isSet: true}
}

func (v NullableInlineObject547) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject547) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


