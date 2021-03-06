/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MobileAppCategory Contains properties for a single Intune app category.
type MobileAppCategory struct {
	// The name of the app category.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The date and time the mobileAppCategory was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// NewMobileAppCategory instantiates a new MobileAppCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileAppCategory() *MobileAppCategory {
	this := MobileAppCategory{}
	return &this
}

// NewMobileAppCategoryWithDefaults instantiates a new MobileAppCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileAppCategoryWithDefaults() *MobileAppCategory {
	this := MobileAppCategory{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MobileAppCategory) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MobileAppCategory) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MobileAppCategory) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MobileAppCategory) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MobileAppCategory) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MobileAppCategory) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MobileAppCategory) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppCategory) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MobileAppCategory) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MobileAppCategory) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

func (o MobileAppCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMobileAppCategory struct {
	value *MobileAppCategory
	isSet bool
}

func (v NullableMobileAppCategory) Get() *MobileAppCategory {
	return v.value
}

func (v *NullableMobileAppCategory) Set(val *MobileAppCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileAppCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileAppCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileAppCategory(val *MobileAppCategory) *NullableMobileAppCategory {
	return &NullableMobileAppCategory{value: val, isSet: true}
}

func (v NullableMobileAppCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileAppCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


