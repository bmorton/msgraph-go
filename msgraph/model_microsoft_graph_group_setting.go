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

// MicrosoftGraphGroupSetting struct for MicrosoftGraphGroupSetting
type MicrosoftGraphGroupSetting struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Display name of this group of settings, which comes from the associated template.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique identifier for the template used to create this group of settings. Read-only.
	TemplateId NullableString `json:"templateId,omitempty"`
	// Collection of name value pairs. Must contain and set all the settings defined in the template.
	Values *[]MicrosoftGraphSettingValue `json:"values,omitempty"`
}

// NewMicrosoftGraphGroupSetting instantiates a new MicrosoftGraphGroupSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphGroupSetting() *MicrosoftGraphGroupSetting {
	this := MicrosoftGraphGroupSetting{}
	return &this
}

// NewMicrosoftGraphGroupSettingWithDefaults instantiates a new MicrosoftGraphGroupSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphGroupSettingWithDefaults() *MicrosoftGraphGroupSetting {
	this := MicrosoftGraphGroupSetting{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphGroupSetting) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGroupSetting) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSetting) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphGroupSetting) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGroupSetting) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGroupSetting) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSetting) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphGroupSetting) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphGroupSetting) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphGroupSetting) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGroupSetting) GetTemplateId() string {
	if o == nil || o.TemplateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGroupSetting) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSetting) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *MicrosoftGraphGroupSetting) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *MicrosoftGraphGroupSetting) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *MicrosoftGraphGroupSetting) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MicrosoftGraphGroupSetting) GetValues() []MicrosoftGraphSettingValue {
	if o == nil || o.Values == nil {
		var ret []MicrosoftGraphSettingValue
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGroupSetting) GetValuesOk() (*[]MicrosoftGraphSettingValue, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSetting) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []MicrosoftGraphSettingValue and assigns it to the Values field.
func (o *MicrosoftGraphGroupSetting) SetValues(v []MicrosoftGraphSettingValue) {
	o.Values = &v
}

func (o MicrosoftGraphGroupSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphGroupSetting struct {
	value *MicrosoftGraphGroupSetting
	isSet bool
}

func (v NullableMicrosoftGraphGroupSetting) Get() *MicrosoftGraphGroupSetting {
	return v.value
}

func (v *NullableMicrosoftGraphGroupSetting) Set(val *MicrosoftGraphGroupSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphGroupSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphGroupSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphGroupSetting(val *MicrosoftGraphGroupSetting) *NullableMicrosoftGraphGroupSetting {
	return &NullableMicrosoftGraphGroupSetting{value: val, isSet: true}
}

func (v NullableMicrosoftGraphGroupSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphGroupSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


