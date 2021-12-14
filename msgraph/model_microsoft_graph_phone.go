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

// MicrosoftGraphPhone struct for MicrosoftGraphPhone
type MicrosoftGraphPhone struct {
	Language NullableString `json:"language,omitempty"`
	// The phone number.
	Number NullableString `json:"number,omitempty"`
	Region NullableString `json:"region,omitempty"`
	// The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
	Type AnyOfmicrosoftGraphPhoneType `json:"type,omitempty"`
}

// NewMicrosoftGraphPhone instantiates a new MicrosoftGraphPhone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPhone() *MicrosoftGraphPhone {
	this := MicrosoftGraphPhone{}
	return &this
}

// NewMicrosoftGraphPhoneWithDefaults instantiates a new MicrosoftGraphPhone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPhoneWithDefaults() *MicrosoftGraphPhone {
	this := MicrosoftGraphPhone{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhone) GetLanguage() string {
	if o == nil || o.Language.Get() == nil {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhone) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *MicrosoftGraphPhone) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *MicrosoftGraphPhone) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *MicrosoftGraphPhone) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *MicrosoftGraphPhone) UnsetLanguage() {
	o.Language.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhone) GetNumber() string {
	if o == nil || o.Number.Get() == nil {
		var ret string
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhone) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphPhone) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableString and assigns it to the Number field.
func (o *MicrosoftGraphPhone) SetNumber(v string) {
	o.Number.Set(&v)
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *MicrosoftGraphPhone) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *MicrosoftGraphPhone) UnsetNumber() {
	o.Number.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhone) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhone) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *MicrosoftGraphPhone) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *MicrosoftGraphPhone) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *MicrosoftGraphPhone) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *MicrosoftGraphPhone) UnsetRegion() {
	o.Region.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhone) GetType() AnyOfmicrosoftGraphPhoneType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPhoneType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhone) GetTypeOk() (*AnyOfmicrosoftGraphPhoneType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphPhone) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfmicrosoftGraphPhoneType and assigns it to the Type field.
func (o *MicrosoftGraphPhone) SetType(v AnyOfmicrosoftGraphPhoneType) {
	o.Type = v
}

func (o MicrosoftGraphPhone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPhone struct {
	value *MicrosoftGraphPhone
	isSet bool
}

func (v NullableMicrosoftGraphPhone) Get() *MicrosoftGraphPhone {
	return v.value
}

func (v *NullableMicrosoftGraphPhone) Set(val *MicrosoftGraphPhone) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPhone) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPhone(val *MicrosoftGraphPhone) *NullableMicrosoftGraphPhone {
	return &NullableMicrosoftGraphPhone{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

