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

// MicrosoftGraphWorkbookChartPointFormat struct for MicrosoftGraphWorkbookChartPointFormat
type MicrosoftGraphWorkbookChartPointFormat struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Represents the fill format of a chart, which includes background formating information. Read-only.
	Fill AnyOfmicrosoftGraphWorkbookChartFill `json:"fill,omitempty"`
}

// NewMicrosoftGraphWorkbookChartPointFormat instantiates a new MicrosoftGraphWorkbookChartPointFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookChartPointFormat() *MicrosoftGraphWorkbookChartPointFormat {
	this := MicrosoftGraphWorkbookChartPointFormat{}
	return &this
}

// NewMicrosoftGraphWorkbookChartPointFormatWithDefaults instantiates a new MicrosoftGraphWorkbookChartPointFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookChartPointFormatWithDefaults() *MicrosoftGraphWorkbookChartPointFormat {
	this := MicrosoftGraphWorkbookChartPointFormat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartPointFormat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartPointFormat) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartPointFormat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartPointFormat) SetId(v string) {
	o.Id = &v
}

// GetFill returns the Fill field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartPointFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartFill
		return ret
	}
	return o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartPointFormat) GetFillOk() (*AnyOfmicrosoftGraphWorkbookChartFill, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return &o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartPointFormat) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFill and assigns it to the Fill field.
func (o *MicrosoftGraphWorkbookChartPointFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill) {
	o.Fill = v
}

func (o MicrosoftGraphWorkbookChartPointFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookChartPointFormat struct {
	value *MicrosoftGraphWorkbookChartPointFormat
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookChartPointFormat) Get() *MicrosoftGraphWorkbookChartPointFormat {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookChartPointFormat) Set(val *MicrosoftGraphWorkbookChartPointFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookChartPointFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookChartPointFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookChartPointFormat(val *MicrosoftGraphWorkbookChartPointFormat) *NullableMicrosoftGraphWorkbookChartPointFormat {
	return &NullableMicrosoftGraphWorkbookChartPointFormat{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookChartPointFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookChartPointFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


