/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphGroupSettingTemplate struct for MicrosoftGraphGroupSettingTemplate
type MicrosoftGraphGroupSettingTemplate struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
	// Description of the template.
	Description NullableString `json:"description,omitempty"`
	// Display name of the template.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.
	Values *[]MicrosoftGraphSettingTemplateValue `json:"values,omitempty"`
}

// NewMicrosoftGraphGroupSettingTemplate instantiates a new MicrosoftGraphGroupSettingTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphGroupSettingTemplate() *MicrosoftGraphGroupSettingTemplate {
	this := MicrosoftGraphGroupSettingTemplate{}
	return &this
}

// NewMicrosoftGraphGroupSettingTemplateWithDefaults instantiates a new MicrosoftGraphGroupSettingTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphGroupSettingTemplateWithDefaults() *MicrosoftGraphGroupSettingTemplate {
	this := MicrosoftGraphGroupSettingTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphGroupSettingTemplate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGroupSettingTemplate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSettingTemplate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphGroupSettingTemplate) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGroupSettingTemplate) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGroupSettingTemplate) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSettingTemplate) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *MicrosoftGraphGroupSettingTemplate) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *MicrosoftGraphGroupSettingTemplate) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *MicrosoftGraphGroupSettingTemplate) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGroupSettingTemplate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGroupSettingTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSettingTemplate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphGroupSettingTemplate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphGroupSettingTemplate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphGroupSettingTemplate) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGroupSettingTemplate) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGroupSettingTemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSettingTemplate) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphGroupSettingTemplate) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphGroupSettingTemplate) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphGroupSettingTemplate) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MicrosoftGraphGroupSettingTemplate) GetValues() []MicrosoftGraphSettingTemplateValue {
	if o == nil || o.Values == nil {
		var ret []MicrosoftGraphSettingTemplateValue
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGroupSettingTemplate) GetValuesOk() (*[]MicrosoftGraphSettingTemplateValue, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MicrosoftGraphGroupSettingTemplate) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []MicrosoftGraphSettingTemplateValue and assigns it to the Values field.
func (o *MicrosoftGraphGroupSettingTemplate) SetValues(v []MicrosoftGraphSettingTemplateValue) {
	o.Values = &v
}

func (o MicrosoftGraphGroupSettingTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime.IsSet() {
		toSerialize["deletedDateTime"] = o.DeletedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphGroupSettingTemplate struct {
	value *MicrosoftGraphGroupSettingTemplate
	isSet bool
}

func (v NullableMicrosoftGraphGroupSettingTemplate) Get() *MicrosoftGraphGroupSettingTemplate {
	return v.value
}

func (v *NullableMicrosoftGraphGroupSettingTemplate) Set(val *MicrosoftGraphGroupSettingTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphGroupSettingTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphGroupSettingTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphGroupSettingTemplate(val *MicrosoftGraphGroupSettingTemplate) *NullableMicrosoftGraphGroupSettingTemplate {
	return &NullableMicrosoftGraphGroupSettingTemplate{value: val, isSet: true}
}

func (v NullableMicrosoftGraphGroupSettingTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphGroupSettingTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

