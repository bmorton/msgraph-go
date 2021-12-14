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

// MicrosoftGraphExtensionProperty struct for MicrosoftGraphExtensionProperty
type MicrosoftGraphExtensionProperty struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
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

// NewMicrosoftGraphExtensionProperty instantiates a new MicrosoftGraphExtensionProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphExtensionProperty() *MicrosoftGraphExtensionProperty {
	this := MicrosoftGraphExtensionProperty{}
	return &this
}

// NewMicrosoftGraphExtensionPropertyWithDefaults instantiates a new MicrosoftGraphExtensionProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphExtensionPropertyWithDefaults() *MicrosoftGraphExtensionProperty {
	this := MicrosoftGraphExtensionProperty{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphExtensionProperty) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphExtensionProperty) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphExtensionProperty) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphExtensionProperty) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphExtensionProperty) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *MicrosoftGraphExtensionProperty) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *MicrosoftGraphExtensionProperty) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *MicrosoftGraphExtensionProperty) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

// GetAppDisplayName returns the AppDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphExtensionProperty) GetAppDisplayName() string {
	if o == nil || o.AppDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppDisplayName.Get()
}

// GetAppDisplayNameOk returns a tuple with the AppDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphExtensionProperty) GetAppDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppDisplayName.Get(), o.AppDisplayName.IsSet()
}

// HasAppDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasAppDisplayName() bool {
	if o != nil && o.AppDisplayName.IsSet() {
		return true
	}

	return false
}

// SetAppDisplayName gets a reference to the given NullableString and assigns it to the AppDisplayName field.
func (o *MicrosoftGraphExtensionProperty) SetAppDisplayName(v string) {
	o.AppDisplayName.Set(&v)
}
// SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil
func (o *MicrosoftGraphExtensionProperty) SetAppDisplayNameNil() {
	o.AppDisplayName.Set(nil)
}

// UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
func (o *MicrosoftGraphExtensionProperty) UnsetAppDisplayName() {
	o.AppDisplayName.Unset()
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *MicrosoftGraphExtensionProperty) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphExtensionProperty) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *MicrosoftGraphExtensionProperty) SetDataType(v string) {
	o.DataType = &v
}

// GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphExtensionProperty) GetIsSyncedFromOnPremises() bool {
	if o == nil || o.IsSyncedFromOnPremises.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsSyncedFromOnPremises.Get()
}

// GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphExtensionProperty) GetIsSyncedFromOnPremisesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsSyncedFromOnPremises.Get(), o.IsSyncedFromOnPremises.IsSet()
}

// HasIsSyncedFromOnPremises returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasIsSyncedFromOnPremises() bool {
	if o != nil && o.IsSyncedFromOnPremises.IsSet() {
		return true
	}

	return false
}

// SetIsSyncedFromOnPremises gets a reference to the given NullableBool and assigns it to the IsSyncedFromOnPremises field.
func (o *MicrosoftGraphExtensionProperty) SetIsSyncedFromOnPremises(v bool) {
	o.IsSyncedFromOnPremises.Set(&v)
}
// SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil
func (o *MicrosoftGraphExtensionProperty) SetIsSyncedFromOnPremisesNil() {
	o.IsSyncedFromOnPremises.Set(nil)
}

// UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
func (o *MicrosoftGraphExtensionProperty) UnsetIsSyncedFromOnPremises() {
	o.IsSyncedFromOnPremises.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MicrosoftGraphExtensionProperty) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphExtensionProperty) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphExtensionProperty) SetName(v string) {
	o.Name = &v
}

// GetTargetObjects returns the TargetObjects field value if set, zero value otherwise.
func (o *MicrosoftGraphExtensionProperty) GetTargetObjects() []string {
	if o == nil || o.TargetObjects == nil {
		var ret []string
		return ret
	}
	return *o.TargetObjects
}

// GetTargetObjectsOk returns a tuple with the TargetObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphExtensionProperty) GetTargetObjectsOk() (*[]string, bool) {
	if o == nil || o.TargetObjects == nil {
		return nil, false
	}
	return o.TargetObjects, true
}

// HasTargetObjects returns a boolean if a field has been set.
func (o *MicrosoftGraphExtensionProperty) HasTargetObjects() bool {
	if o != nil && o.TargetObjects != nil {
		return true
	}

	return false
}

// SetTargetObjects gets a reference to the given []string and assigns it to the TargetObjects field.
func (o *MicrosoftGraphExtensionProperty) SetTargetObjects(v []string) {
	o.TargetObjects = &v
}

func (o MicrosoftGraphExtensionProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime.IsSet() {
		toSerialize["deletedDateTime"] = o.DeletedDateTime.Get()
	}
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

type NullableMicrosoftGraphExtensionProperty struct {
	value *MicrosoftGraphExtensionProperty
	isSet bool
}

func (v NullableMicrosoftGraphExtensionProperty) Get() *MicrosoftGraphExtensionProperty {
	return v.value
}

func (v *NullableMicrosoftGraphExtensionProperty) Set(val *MicrosoftGraphExtensionProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphExtensionProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphExtensionProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphExtensionProperty(val *MicrosoftGraphExtensionProperty) *NullableMicrosoftGraphExtensionProperty {
	return &NullableMicrosoftGraphExtensionProperty{value: val, isSet: true}
}

func (v NullableMicrosoftGraphExtensionProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphExtensionProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


