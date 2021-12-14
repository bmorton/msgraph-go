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

// MicrosoftGraphRgbColor Color in RGB.
type MicrosoftGraphRgbColor struct {
	// Blue value
	B *int32 `json:"b,omitempty"`
	// Green value
	G *int32 `json:"g,omitempty"`
	// Red value
	R *int32 `json:"r,omitempty"`
}

// NewMicrosoftGraphRgbColor instantiates a new MicrosoftGraphRgbColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRgbColor() *MicrosoftGraphRgbColor {
	this := MicrosoftGraphRgbColor{}
	return &this
}

// NewMicrosoftGraphRgbColorWithDefaults instantiates a new MicrosoftGraphRgbColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRgbColorWithDefaults() *MicrosoftGraphRgbColor {
	this := MicrosoftGraphRgbColor{}
	return &this
}

// GetB returns the B field value if set, zero value otherwise.
func (o *MicrosoftGraphRgbColor) GetB() int32 {
	if o == nil || o.B == nil {
		var ret int32
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRgbColor) GetBOk() (*int32, bool) {
	if o == nil || o.B == nil {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *MicrosoftGraphRgbColor) HasB() bool {
	if o != nil && o.B != nil {
		return true
	}

	return false
}

// SetB gets a reference to the given int32 and assigns it to the B field.
func (o *MicrosoftGraphRgbColor) SetB(v int32) {
	o.B = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *MicrosoftGraphRgbColor) GetG() int32 {
	if o == nil || o.G == nil {
		var ret int32
		return ret
	}
	return *o.G
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRgbColor) GetGOk() (*int32, bool) {
	if o == nil || o.G == nil {
		return nil, false
	}
	return o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *MicrosoftGraphRgbColor) HasG() bool {
	if o != nil && o.G != nil {
		return true
	}

	return false
}

// SetG gets a reference to the given int32 and assigns it to the G field.
func (o *MicrosoftGraphRgbColor) SetG(v int32) {
	o.G = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *MicrosoftGraphRgbColor) GetR() int32 {
	if o == nil || o.R == nil {
		var ret int32
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRgbColor) GetROk() (*int32, bool) {
	if o == nil || o.R == nil {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *MicrosoftGraphRgbColor) HasR() bool {
	if o != nil && o.R != nil {
		return true
	}

	return false
}

// SetR gets a reference to the given int32 and assigns it to the R field.
func (o *MicrosoftGraphRgbColor) SetR(v int32) {
	o.R = &v
}

func (o MicrosoftGraphRgbColor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.B != nil {
		toSerialize["b"] = o.B
	}
	if o.G != nil {
		toSerialize["g"] = o.G
	}
	if o.R != nil {
		toSerialize["r"] = o.R
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRgbColor struct {
	value *MicrosoftGraphRgbColor
	isSet bool
}

func (v NullableMicrosoftGraphRgbColor) Get() *MicrosoftGraphRgbColor {
	return v.value
}

func (v *NullableMicrosoftGraphRgbColor) Set(val *MicrosoftGraphRgbColor) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRgbColor) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRgbColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRgbColor(val *MicrosoftGraphRgbColor) *NullableMicrosoftGraphRgbColor {
	return &NullableMicrosoftGraphRgbColor{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRgbColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRgbColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


