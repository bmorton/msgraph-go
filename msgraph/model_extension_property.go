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

// ExtensionProperty struct for ExtensionProperty
type ExtensionProperty struct {
	// Display name of the application object on which this extension property is defined. Read-only.
	AppDisplayName NullableString `json:"appDisplayName,omitempty"`
	// Specifies the data type of the value the extension property can hold. Following values are supported. Not nullable. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximum
	DataType *string `json:"dataType,omitempty"`
	// Indicates if this extension property was sycned from onpremises directory using Azure AD Connect. Read-only.
	IsSyncedFromOnPremises NullableBool `json:"isSyncedFromOnPremises,omitempty"`
	// Name of the extension property. Not nullable.
	Name *string `json:"name,omitempty"`
	// Following values are supported. Not nullable. UserGroupOrganizationDeviceApplication
	TargetObjects *[]string `json:"targetObjects,omitempty"`
}

// NewExtensionProperty instantiates a new ExtensionProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionProperty() *ExtensionProperty {
	this := ExtensionProperty{}
	return &this
}

// NewExtensionPropertyWithDefaults instantiates a new ExtensionProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionPropertyWithDefaults() *ExtensionProperty {
	this := ExtensionProperty{}
	return &this
}

// GetAppDisplayName returns the AppDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionProperty) GetAppDisplayName() string {
	if o == nil || o.AppDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppDisplayName.Get()
}

// GetAppDisplayNameOk returns a tuple with the AppDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionProperty) GetAppDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppDisplayName.Get(), o.AppDisplayName.IsSet()
}

// HasAppDisplayName returns a boolean if a field has been set.
func (o *ExtensionProperty) HasAppDisplayName() bool {
	if o != nil && o.AppDisplayName.IsSet() {
		return true
	}

	return false
}

// SetAppDisplayName gets a reference to the given NullableString and assigns it to the AppDisplayName field.
func (o *ExtensionProperty) SetAppDisplayName(v string) {
	o.AppDisplayName.Set(&v)
}
// SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil
func (o *ExtensionProperty) SetAppDisplayNameNil() {
	o.AppDisplayName.Set(nil)
}

// UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
func (o *ExtensionProperty) UnsetAppDisplayName() {
	o.AppDisplayName.Unset()
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *ExtensionProperty) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionProperty) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *ExtensionProperty) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *ExtensionProperty) SetDataType(v string) {
	o.DataType = &v
}

// GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionProperty) GetIsSyncedFromOnPremises() bool {
	if o == nil || o.IsSyncedFromOnPremises.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsSyncedFromOnPremises.Get()
}

// GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionProperty) GetIsSyncedFromOnPremisesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsSyncedFromOnPremises.Get(), o.IsSyncedFromOnPremises.IsSet()
}

// HasIsSyncedFromOnPremises returns a boolean if a field has been set.
func (o *ExtensionProperty) HasIsSyncedFromOnPremises() bool {
	if o != nil && o.IsSyncedFromOnPremises.IsSet() {
		return true
	}

	return false
}

// SetIsSyncedFromOnPremises gets a reference to the given NullableBool and assigns it to the IsSyncedFromOnPremises field.
func (o *ExtensionProperty) SetIsSyncedFromOnPremises(v bool) {
	o.IsSyncedFromOnPremises.Set(&v)
}
// SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil
func (o *ExtensionProperty) SetIsSyncedFromOnPremisesNil() {
	o.IsSyncedFromOnPremises.Set(nil)
}

// UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
func (o *ExtensionProperty) UnsetIsSyncedFromOnPremises() {
	o.IsSyncedFromOnPremises.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtensionProperty) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionProperty) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtensionProperty) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtensionProperty) SetName(v string) {
	o.Name = &v
}

// GetTargetObjects returns the TargetObjects field value if set, zero value otherwise.
func (o *ExtensionProperty) GetTargetObjects() []string {
	if o == nil || o.TargetObjects == nil {
		var ret []string
		return ret
	}
	return *o.TargetObjects
}

// GetTargetObjectsOk returns a tuple with the TargetObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionProperty) GetTargetObjectsOk() (*[]string, bool) {
	if o == nil || o.TargetObjects == nil {
		return nil, false
	}
	return o.TargetObjects, true
}

// HasTargetObjects returns a boolean if a field has been set.
func (o *ExtensionProperty) HasTargetObjects() bool {
	if o != nil && o.TargetObjects != nil {
		return true
	}

	return false
}

// SetTargetObjects gets a reference to the given []string and assigns it to the TargetObjects field.
func (o *ExtensionProperty) SetTargetObjects(v []string) {
	o.TargetObjects = &v
}

func (o ExtensionProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppDisplayName.IsSet() {
		toSerialize["appDisplayName"] = o.AppDisplayName.Get()
	}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.IsSyncedFromOnPremises.IsSet() {
		toSerialize["isSyncedFromOnPremises"] = o.IsSyncedFromOnPremises.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TargetObjects != nil {
		toSerialize["targetObjects"] = o.TargetObjects
	}
	return json.Marshal(toSerialize)
}

type NullableExtensionProperty struct {
	value *ExtensionProperty
	isSet bool
}

func (v NullableExtensionProperty) Get() *ExtensionProperty {
	return v.value
}

func (v *NullableExtensionProperty) Set(val *ExtensionProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionProperty(val *ExtensionProperty) *NullableExtensionProperty {
	return &NullableExtensionProperty{value: val, isSet: true}
}

func (v NullableExtensionProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

