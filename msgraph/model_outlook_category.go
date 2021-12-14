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

// OutlookCategory struct for OutlookCategory
type OutlookCategory struct {
	// A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
	Color AnyOfmicrosoftGraphCategoryColor `json:"color,omitempty"`
	// A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewOutlookCategory instantiates a new OutlookCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlookCategory() *OutlookCategory {
	this := OutlookCategory{}
	return &this
}

// NewOutlookCategoryWithDefaults instantiates a new OutlookCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlookCategoryWithDefaults() *OutlookCategory {
	this := OutlookCategory{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutlookCategory) GetColor() AnyOfmicrosoftGraphCategoryColor {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCategoryColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutlookCategory) GetColorOk() (*AnyOfmicrosoftGraphCategoryColor, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return &o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *OutlookCategory) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given AnyOfmicrosoftGraphCategoryColor and assigns it to the Color field.
func (o *OutlookCategory) SetColor(v AnyOfmicrosoftGraphCategoryColor) {
	o.Color = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutlookCategory) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutlookCategory) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OutlookCategory) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *OutlookCategory) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *OutlookCategory) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *OutlookCategory) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o OutlookCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOutlookCategory struct {
	value *OutlookCategory
	isSet bool
}

func (v NullableOutlookCategory) Get() *OutlookCategory {
	return v.value
}

func (v *NullableOutlookCategory) Set(val *OutlookCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlookCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlookCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlookCategory(val *OutlookCategory) *NullableOutlookCategory {
	return &NullableOutlookCategory{value: val, isSet: true}
}

func (v NullableOutlookCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlookCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


