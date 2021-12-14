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

// MicrosoftGraphOutlookCategory struct for MicrosoftGraphOutlookCategory
type MicrosoftGraphOutlookCategory struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
	Color AnyOfmicrosoftGraphCategoryColor `json:"color,omitempty"`
	// A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewMicrosoftGraphOutlookCategory instantiates a new MicrosoftGraphOutlookCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOutlookCategory() *MicrosoftGraphOutlookCategory {
	this := MicrosoftGraphOutlookCategory{}
	return &this
}

// NewMicrosoftGraphOutlookCategoryWithDefaults instantiates a new MicrosoftGraphOutlookCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOutlookCategoryWithDefaults() *MicrosoftGraphOutlookCategory {
	this := MicrosoftGraphOutlookCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOutlookCategory) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookCategory) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookCategory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOutlookCategory) SetId(v string) {
	o.Id = &v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOutlookCategory) GetColor() AnyOfmicrosoftGraphCategoryColor {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCategoryColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOutlookCategory) GetColorOk() (*AnyOfmicrosoftGraphCategoryColor, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return &o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookCategory) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given AnyOfmicrosoftGraphCategoryColor and assigns it to the Color field.
func (o *MicrosoftGraphOutlookCategory) SetColor(v AnyOfmicrosoftGraphCategoryColor) {
	o.Color = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOutlookCategory) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOutlookCategory) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookCategory) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphOutlookCategory) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphOutlookCategory) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphOutlookCategory) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o MicrosoftGraphOutlookCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOutlookCategory struct {
	value *MicrosoftGraphOutlookCategory
	isSet bool
}

func (v NullableMicrosoftGraphOutlookCategory) Get() *MicrosoftGraphOutlookCategory {
	return v.value
}

func (v *NullableMicrosoftGraphOutlookCategory) Set(val *MicrosoftGraphOutlookCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOutlookCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOutlookCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOutlookCategory(val *MicrosoftGraphOutlookCategory) *NullableMicrosoftGraphOutlookCategory {
	return &NullableMicrosoftGraphOutlookCategory{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOutlookCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOutlookCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


