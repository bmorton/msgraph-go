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

// MicrosoftGraphWorkbookChartGridlines struct for MicrosoftGraphWorkbookChartGridlines
type MicrosoftGraphWorkbookChartGridlines struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Boolean value representing if the axis gridlines are visible or not.
	Visible *bool `json:"visible,omitempty"`
	// Represents the formatting of chart gridlines. Read-only.
	Format AnyOfmicrosoftGraphWorkbookChartGridlinesFormat `json:"format,omitempty"`
}

// NewMicrosoftGraphWorkbookChartGridlines instantiates a new MicrosoftGraphWorkbookChartGridlines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookChartGridlines() *MicrosoftGraphWorkbookChartGridlines {
	this := MicrosoftGraphWorkbookChartGridlines{}
	return &this
}

// NewMicrosoftGraphWorkbookChartGridlinesWithDefaults instantiates a new MicrosoftGraphWorkbookChartGridlines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookChartGridlinesWithDefaults() *MicrosoftGraphWorkbookChartGridlines {
	this := MicrosoftGraphWorkbookChartGridlines{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartGridlines) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartGridlines) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartGridlines) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartGridlines) SetId(v string) {
	o.Id = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartGridlines) GetVisible() bool {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartGridlines) GetVisibleOk() (*bool, bool) {
	if o == nil || o.Visible == nil {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartGridlines) HasVisible() bool {
	if o != nil && o.Visible != nil {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *MicrosoftGraphWorkbookChartGridlines) SetVisible(v bool) {
	o.Visible = &v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookChartGridlines) GetFormat() AnyOfmicrosoftGraphWorkbookChartGridlinesFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookChartGridlinesFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookChartGridlines) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartGridlinesFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartGridlines) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartGridlinesFormat and assigns it to the Format field.
func (o *MicrosoftGraphWorkbookChartGridlines) SetFormat(v AnyOfmicrosoftGraphWorkbookChartGridlinesFormat) {
	o.Format = v
}

func (o MicrosoftGraphWorkbookChartGridlines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Visible != nil {
		toSerialize["visible"] = o.Visible
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookChartGridlines struct {
	value *MicrosoftGraphWorkbookChartGridlines
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookChartGridlines) Get() *MicrosoftGraphWorkbookChartGridlines {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookChartGridlines) Set(val *MicrosoftGraphWorkbookChartGridlines) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookChartGridlines) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookChartGridlines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookChartGridlines(val *MicrosoftGraphWorkbookChartGridlines) *NullableMicrosoftGraphWorkbookChartGridlines {
	return &NullableMicrosoftGraphWorkbookChartGridlines{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookChartGridlines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookChartGridlines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


