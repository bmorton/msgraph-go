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

// MicrosoftGraphOnPremisesExtensionAttributes struct for MicrosoftGraphOnPremisesExtensionAttributes
type MicrosoftGraphOnPremisesExtensionAttributes struct {
	// First customizable extension attribute.
	ExtensionAttribute1 NullableString `json:"extensionAttribute1,omitempty"`
	// Tenth customizable extension attribute.
	ExtensionAttribute10 NullableString `json:"extensionAttribute10,omitempty"`
	// Eleventh customizable extension attribute.
	ExtensionAttribute11 NullableString `json:"extensionAttribute11,omitempty"`
	// Twelfth customizable extension attribute.
	ExtensionAttribute12 NullableString `json:"extensionAttribute12,omitempty"`
	// Thirteenth customizable extension attribute.
	ExtensionAttribute13 NullableString `json:"extensionAttribute13,omitempty"`
	// Fourteenth customizable extension attribute.
	ExtensionAttribute14 NullableString `json:"extensionAttribute14,omitempty"`
	// Fifteenth customizable extension attribute.
	ExtensionAttribute15 NullableString `json:"extensionAttribute15,omitempty"`
	// Second customizable extension attribute.
	ExtensionAttribute2 NullableString `json:"extensionAttribute2,omitempty"`
	// Third customizable extension attribute.
	ExtensionAttribute3 NullableString `json:"extensionAttribute3,omitempty"`
	// Fourth customizable extension attribute.
	ExtensionAttribute4 NullableString `json:"extensionAttribute4,omitempty"`
	// Fifth customizable extension attribute.
	ExtensionAttribute5 NullableString `json:"extensionAttribute5,omitempty"`
	// Sixth customizable extension attribute.
	ExtensionAttribute6 NullableString `json:"extensionAttribute6,omitempty"`
	// Seventh customizable extension attribute.
	ExtensionAttribute7 NullableString `json:"extensionAttribute7,omitempty"`
	// Eighth customizable extension attribute.
	ExtensionAttribute8 NullableString `json:"extensionAttribute8,omitempty"`
	// Ninth customizable extension attribute.
	ExtensionAttribute9 NullableString `json:"extensionAttribute9,omitempty"`
}

// NewMicrosoftGraphOnPremisesExtensionAttributes instantiates a new MicrosoftGraphOnPremisesExtensionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOnPremisesExtensionAttributes() *MicrosoftGraphOnPremisesExtensionAttributes {
	this := MicrosoftGraphOnPremisesExtensionAttributes{}
	return &this
}

// NewMicrosoftGraphOnPremisesExtensionAttributesWithDefaults instantiates a new MicrosoftGraphOnPremisesExtensionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOnPremisesExtensionAttributesWithDefaults() *MicrosoftGraphOnPremisesExtensionAttributes {
	this := MicrosoftGraphOnPremisesExtensionAttributes{}
	return &this
}

// GetExtensionAttribute1 returns the ExtensionAttribute1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute1() string {
	if o == nil || o.ExtensionAttribute1.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute1.Get()
}

// GetExtensionAttribute1Ok returns a tuple with the ExtensionAttribute1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute1.Get(), o.ExtensionAttribute1.IsSet()
}

// HasExtensionAttribute1 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute1() bool {
	if o != nil && o.ExtensionAttribute1.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute1 gets a reference to the given NullableString and assigns it to the ExtensionAttribute1 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute1(v string) {
	o.ExtensionAttribute1.Set(&v)
}
// SetExtensionAttribute1Nil sets the value for ExtensionAttribute1 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute1Nil() {
	o.ExtensionAttribute1.Set(nil)
}

// UnsetExtensionAttribute1 ensures that no value is present for ExtensionAttribute1, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute1() {
	o.ExtensionAttribute1.Unset()
}

// GetExtensionAttribute10 returns the ExtensionAttribute10 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute10() string {
	if o == nil || o.ExtensionAttribute10.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute10.Get()
}

// GetExtensionAttribute10Ok returns a tuple with the ExtensionAttribute10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute10Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute10.Get(), o.ExtensionAttribute10.IsSet()
}

// HasExtensionAttribute10 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute10() bool {
	if o != nil && o.ExtensionAttribute10.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute10 gets a reference to the given NullableString and assigns it to the ExtensionAttribute10 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute10(v string) {
	o.ExtensionAttribute10.Set(&v)
}
// SetExtensionAttribute10Nil sets the value for ExtensionAttribute10 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute10Nil() {
	o.ExtensionAttribute10.Set(nil)
}

// UnsetExtensionAttribute10 ensures that no value is present for ExtensionAttribute10, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute10() {
	o.ExtensionAttribute10.Unset()
}

// GetExtensionAttribute11 returns the ExtensionAttribute11 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute11() string {
	if o == nil || o.ExtensionAttribute11.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute11.Get()
}

// GetExtensionAttribute11Ok returns a tuple with the ExtensionAttribute11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute11Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute11.Get(), o.ExtensionAttribute11.IsSet()
}

// HasExtensionAttribute11 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute11() bool {
	if o != nil && o.ExtensionAttribute11.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute11 gets a reference to the given NullableString and assigns it to the ExtensionAttribute11 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute11(v string) {
	o.ExtensionAttribute11.Set(&v)
}
// SetExtensionAttribute11Nil sets the value for ExtensionAttribute11 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute11Nil() {
	o.ExtensionAttribute11.Set(nil)
}

// UnsetExtensionAttribute11 ensures that no value is present for ExtensionAttribute11, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute11() {
	o.ExtensionAttribute11.Unset()
}

// GetExtensionAttribute12 returns the ExtensionAttribute12 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute12() string {
	if o == nil || o.ExtensionAttribute12.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute12.Get()
}

// GetExtensionAttribute12Ok returns a tuple with the ExtensionAttribute12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute12Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute12.Get(), o.ExtensionAttribute12.IsSet()
}

// HasExtensionAttribute12 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute12() bool {
	if o != nil && o.ExtensionAttribute12.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute12 gets a reference to the given NullableString and assigns it to the ExtensionAttribute12 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute12(v string) {
	o.ExtensionAttribute12.Set(&v)
}
// SetExtensionAttribute12Nil sets the value for ExtensionAttribute12 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute12Nil() {
	o.ExtensionAttribute12.Set(nil)
}

// UnsetExtensionAttribute12 ensures that no value is present for ExtensionAttribute12, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute12() {
	o.ExtensionAttribute12.Unset()
}

// GetExtensionAttribute13 returns the ExtensionAttribute13 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute13() string {
	if o == nil || o.ExtensionAttribute13.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute13.Get()
}

// GetExtensionAttribute13Ok returns a tuple with the ExtensionAttribute13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute13Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute13.Get(), o.ExtensionAttribute13.IsSet()
}

// HasExtensionAttribute13 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute13() bool {
	if o != nil && o.ExtensionAttribute13.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute13 gets a reference to the given NullableString and assigns it to the ExtensionAttribute13 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute13(v string) {
	o.ExtensionAttribute13.Set(&v)
}
// SetExtensionAttribute13Nil sets the value for ExtensionAttribute13 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute13Nil() {
	o.ExtensionAttribute13.Set(nil)
}

// UnsetExtensionAttribute13 ensures that no value is present for ExtensionAttribute13, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute13() {
	o.ExtensionAttribute13.Unset()
}

// GetExtensionAttribute14 returns the ExtensionAttribute14 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute14() string {
	if o == nil || o.ExtensionAttribute14.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute14.Get()
}

// GetExtensionAttribute14Ok returns a tuple with the ExtensionAttribute14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute14Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute14.Get(), o.ExtensionAttribute14.IsSet()
}

// HasExtensionAttribute14 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute14() bool {
	if o != nil && o.ExtensionAttribute14.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute14 gets a reference to the given NullableString and assigns it to the ExtensionAttribute14 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute14(v string) {
	o.ExtensionAttribute14.Set(&v)
}
// SetExtensionAttribute14Nil sets the value for ExtensionAttribute14 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute14Nil() {
	o.ExtensionAttribute14.Set(nil)
}

// UnsetExtensionAttribute14 ensures that no value is present for ExtensionAttribute14, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute14() {
	o.ExtensionAttribute14.Unset()
}

// GetExtensionAttribute15 returns the ExtensionAttribute15 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute15() string {
	if o == nil || o.ExtensionAttribute15.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute15.Get()
}

// GetExtensionAttribute15Ok returns a tuple with the ExtensionAttribute15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute15Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute15.Get(), o.ExtensionAttribute15.IsSet()
}

// HasExtensionAttribute15 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute15() bool {
	if o != nil && o.ExtensionAttribute15.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute15 gets a reference to the given NullableString and assigns it to the ExtensionAttribute15 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute15(v string) {
	o.ExtensionAttribute15.Set(&v)
}
// SetExtensionAttribute15Nil sets the value for ExtensionAttribute15 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute15Nil() {
	o.ExtensionAttribute15.Set(nil)
}

// UnsetExtensionAttribute15 ensures that no value is present for ExtensionAttribute15, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute15() {
	o.ExtensionAttribute15.Unset()
}

// GetExtensionAttribute2 returns the ExtensionAttribute2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute2() string {
	if o == nil || o.ExtensionAttribute2.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute2.Get()
}

// GetExtensionAttribute2Ok returns a tuple with the ExtensionAttribute2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute2.Get(), o.ExtensionAttribute2.IsSet()
}

// HasExtensionAttribute2 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute2() bool {
	if o != nil && o.ExtensionAttribute2.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute2 gets a reference to the given NullableString and assigns it to the ExtensionAttribute2 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute2(v string) {
	o.ExtensionAttribute2.Set(&v)
}
// SetExtensionAttribute2Nil sets the value for ExtensionAttribute2 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute2Nil() {
	o.ExtensionAttribute2.Set(nil)
}

// UnsetExtensionAttribute2 ensures that no value is present for ExtensionAttribute2, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute2() {
	o.ExtensionAttribute2.Unset()
}

// GetExtensionAttribute3 returns the ExtensionAttribute3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute3() string {
	if o == nil || o.ExtensionAttribute3.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute3.Get()
}

// GetExtensionAttribute3Ok returns a tuple with the ExtensionAttribute3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute3Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute3.Get(), o.ExtensionAttribute3.IsSet()
}

// HasExtensionAttribute3 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute3() bool {
	if o != nil && o.ExtensionAttribute3.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute3 gets a reference to the given NullableString and assigns it to the ExtensionAttribute3 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute3(v string) {
	o.ExtensionAttribute3.Set(&v)
}
// SetExtensionAttribute3Nil sets the value for ExtensionAttribute3 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute3Nil() {
	o.ExtensionAttribute3.Set(nil)
}

// UnsetExtensionAttribute3 ensures that no value is present for ExtensionAttribute3, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute3() {
	o.ExtensionAttribute3.Unset()
}

// GetExtensionAttribute4 returns the ExtensionAttribute4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute4() string {
	if o == nil || o.ExtensionAttribute4.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute4.Get()
}

// GetExtensionAttribute4Ok returns a tuple with the ExtensionAttribute4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute4Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute4.Get(), o.ExtensionAttribute4.IsSet()
}

// HasExtensionAttribute4 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute4() bool {
	if o != nil && o.ExtensionAttribute4.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute4 gets a reference to the given NullableString and assigns it to the ExtensionAttribute4 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute4(v string) {
	o.ExtensionAttribute4.Set(&v)
}
// SetExtensionAttribute4Nil sets the value for ExtensionAttribute4 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute4Nil() {
	o.ExtensionAttribute4.Set(nil)
}

// UnsetExtensionAttribute4 ensures that no value is present for ExtensionAttribute4, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute4() {
	o.ExtensionAttribute4.Unset()
}

// GetExtensionAttribute5 returns the ExtensionAttribute5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute5() string {
	if o == nil || o.ExtensionAttribute5.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute5.Get()
}

// GetExtensionAttribute5Ok returns a tuple with the ExtensionAttribute5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute5Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute5.Get(), o.ExtensionAttribute5.IsSet()
}

// HasExtensionAttribute5 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute5() bool {
	if o != nil && o.ExtensionAttribute5.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute5 gets a reference to the given NullableString and assigns it to the ExtensionAttribute5 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute5(v string) {
	o.ExtensionAttribute5.Set(&v)
}
// SetExtensionAttribute5Nil sets the value for ExtensionAttribute5 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute5Nil() {
	o.ExtensionAttribute5.Set(nil)
}

// UnsetExtensionAttribute5 ensures that no value is present for ExtensionAttribute5, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute5() {
	o.ExtensionAttribute5.Unset()
}

// GetExtensionAttribute6 returns the ExtensionAttribute6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute6() string {
	if o == nil || o.ExtensionAttribute6.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute6.Get()
}

// GetExtensionAttribute6Ok returns a tuple with the ExtensionAttribute6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute6Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute6.Get(), o.ExtensionAttribute6.IsSet()
}

// HasExtensionAttribute6 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute6() bool {
	if o != nil && o.ExtensionAttribute6.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute6 gets a reference to the given NullableString and assigns it to the ExtensionAttribute6 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute6(v string) {
	o.ExtensionAttribute6.Set(&v)
}
// SetExtensionAttribute6Nil sets the value for ExtensionAttribute6 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute6Nil() {
	o.ExtensionAttribute6.Set(nil)
}

// UnsetExtensionAttribute6 ensures that no value is present for ExtensionAttribute6, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute6() {
	o.ExtensionAttribute6.Unset()
}

// GetExtensionAttribute7 returns the ExtensionAttribute7 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute7() string {
	if o == nil || o.ExtensionAttribute7.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute7.Get()
}

// GetExtensionAttribute7Ok returns a tuple with the ExtensionAttribute7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute7Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute7.Get(), o.ExtensionAttribute7.IsSet()
}

// HasExtensionAttribute7 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute7() bool {
	if o != nil && o.ExtensionAttribute7.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute7 gets a reference to the given NullableString and assigns it to the ExtensionAttribute7 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute7(v string) {
	o.ExtensionAttribute7.Set(&v)
}
// SetExtensionAttribute7Nil sets the value for ExtensionAttribute7 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute7Nil() {
	o.ExtensionAttribute7.Set(nil)
}

// UnsetExtensionAttribute7 ensures that no value is present for ExtensionAttribute7, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute7() {
	o.ExtensionAttribute7.Unset()
}

// GetExtensionAttribute8 returns the ExtensionAttribute8 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute8() string {
	if o == nil || o.ExtensionAttribute8.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute8.Get()
}

// GetExtensionAttribute8Ok returns a tuple with the ExtensionAttribute8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute8Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute8.Get(), o.ExtensionAttribute8.IsSet()
}

// HasExtensionAttribute8 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute8() bool {
	if o != nil && o.ExtensionAttribute8.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute8 gets a reference to the given NullableString and assigns it to the ExtensionAttribute8 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute8(v string) {
	o.ExtensionAttribute8.Set(&v)
}
// SetExtensionAttribute8Nil sets the value for ExtensionAttribute8 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute8Nil() {
	o.ExtensionAttribute8.Set(nil)
}

// UnsetExtensionAttribute8 ensures that no value is present for ExtensionAttribute8, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute8() {
	o.ExtensionAttribute8.Unset()
}

// GetExtensionAttribute9 returns the ExtensionAttribute9 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute9() string {
	if o == nil || o.ExtensionAttribute9.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtensionAttribute9.Get()
}

// GetExtensionAttribute9Ok returns a tuple with the ExtensionAttribute9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnPremisesExtensionAttributes) GetExtensionAttribute9Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtensionAttribute9.Get(), o.ExtensionAttribute9.IsSet()
}

// HasExtensionAttribute9 returns a boolean if a field has been set.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) HasExtensionAttribute9() bool {
	if o != nil && o.ExtensionAttribute9.IsSet() {
		return true
	}

	return false
}

// SetExtensionAttribute9 gets a reference to the given NullableString and assigns it to the ExtensionAttribute9 field.
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute9(v string) {
	o.ExtensionAttribute9.Set(&v)
}
// SetExtensionAttribute9Nil sets the value for ExtensionAttribute9 to be an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) SetExtensionAttribute9Nil() {
	o.ExtensionAttribute9.Set(nil)
}

// UnsetExtensionAttribute9 ensures that no value is present for ExtensionAttribute9, not even an explicit nil
func (o *MicrosoftGraphOnPremisesExtensionAttributes) UnsetExtensionAttribute9() {
	o.ExtensionAttribute9.Unset()
}

func (o MicrosoftGraphOnPremisesExtensionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtensionAttribute1.IsSet() {
		toSerialize["extensionAttribute1"] = o.ExtensionAttribute1.Get()
	}
	if o.ExtensionAttribute10.IsSet() {
		toSerialize["extensionAttribute10"] = o.ExtensionAttribute10.Get()
	}
	if o.ExtensionAttribute11.IsSet() {
		toSerialize["extensionAttribute11"] = o.ExtensionAttribute11.Get()
	}
	if o.ExtensionAttribute12.IsSet() {
		toSerialize["extensionAttribute12"] = o.ExtensionAttribute12.Get()
	}
	if o.ExtensionAttribute13.IsSet() {
		toSerialize["extensionAttribute13"] = o.ExtensionAttribute13.Get()
	}
	if o.ExtensionAttribute14.IsSet() {
		toSerialize["extensionAttribute14"] = o.ExtensionAttribute14.Get()
	}
	if o.ExtensionAttribute15.IsSet() {
		toSerialize["extensionAttribute15"] = o.ExtensionAttribute15.Get()
	}
	if o.ExtensionAttribute2.IsSet() {
		toSerialize["extensionAttribute2"] = o.ExtensionAttribute2.Get()
	}
	if o.ExtensionAttribute3.IsSet() {
		toSerialize["extensionAttribute3"] = o.ExtensionAttribute3.Get()
	}
	if o.ExtensionAttribute4.IsSet() {
		toSerialize["extensionAttribute4"] = o.ExtensionAttribute4.Get()
	}
	if o.ExtensionAttribute5.IsSet() {
		toSerialize["extensionAttribute5"] = o.ExtensionAttribute5.Get()
	}
	if o.ExtensionAttribute6.IsSet() {
		toSerialize["extensionAttribute6"] = o.ExtensionAttribute6.Get()
	}
	if o.ExtensionAttribute7.IsSet() {
		toSerialize["extensionAttribute7"] = o.ExtensionAttribute7.Get()
	}
	if o.ExtensionAttribute8.IsSet() {
		toSerialize["extensionAttribute8"] = o.ExtensionAttribute8.Get()
	}
	if o.ExtensionAttribute9.IsSet() {
		toSerialize["extensionAttribute9"] = o.ExtensionAttribute9.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOnPremisesExtensionAttributes struct {
	value *MicrosoftGraphOnPremisesExtensionAttributes
	isSet bool
}

func (v NullableMicrosoftGraphOnPremisesExtensionAttributes) Get() *MicrosoftGraphOnPremisesExtensionAttributes {
	return v.value
}

func (v *NullableMicrosoftGraphOnPremisesExtensionAttributes) Set(val *MicrosoftGraphOnPremisesExtensionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOnPremisesExtensionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOnPremisesExtensionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOnPremisesExtensionAttributes(val *MicrosoftGraphOnPremisesExtensionAttributes) *NullableMicrosoftGraphOnPremisesExtensionAttributes {
	return &NullableMicrosoftGraphOnPremisesExtensionAttributes{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOnPremisesExtensionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOnPremisesExtensionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

