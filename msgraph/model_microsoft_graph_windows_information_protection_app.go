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

// MicrosoftGraphWindowsInformationProtectionApp App for Windows information protection
type MicrosoftGraphWindowsInformationProtectionApp struct {
	// If true, app is denied protection or exemption.
	Denied *bool `json:"denied,omitempty"`
	// The app's description.
	Description NullableString `json:"description,omitempty"`
	// App display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The product name.
	ProductName NullableString `json:"productName,omitempty"`
	// The publisher name
	PublisherName NullableString `json:"publisherName,omitempty"`
}

// NewMicrosoftGraphWindowsInformationProtectionApp instantiates a new MicrosoftGraphWindowsInformationProtectionApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWindowsInformationProtectionApp() *MicrosoftGraphWindowsInformationProtectionApp {
	this := MicrosoftGraphWindowsInformationProtectionApp{}
	return &this
}

// NewMicrosoftGraphWindowsInformationProtectionAppWithDefaults instantiates a new MicrosoftGraphWindowsInformationProtectionApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWindowsInformationProtectionAppWithDefaults() *MicrosoftGraphWindowsInformationProtectionApp {
	this := MicrosoftGraphWindowsInformationProtectionApp{}
	return &this
}

// GetDenied returns the Denied field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetDenied() bool {
	if o == nil || o.Denied == nil {
		var ret bool
		return ret
	}
	return *o.Denied
}

// GetDeniedOk returns a tuple with the Denied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetDeniedOk() (*bool, bool) {
	if o == nil || o.Denied == nil {
		return nil, false
	}
	return o.Denied, true
}

// HasDenied returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) HasDenied() bool {
	if o != nil && o.Denied != nil {
		return true
	}

	return false
}

// SetDenied gets a reference to the given bool and assigns it to the Denied field.
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetDenied(v bool) {
	o.Denied = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionApp) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetProductName() string {
	if o == nil || o.ProductName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetProductNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionApp) UnsetProductName() {
	o.ProductName.Unset()
}

// GetPublisherName returns the PublisherName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetPublisherName() string {
	if o == nil || o.PublisherName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublisherName.Get()
}

// GetPublisherNameOk returns a tuple with the PublisherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWindowsInformationProtectionApp) GetPublisherNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublisherName.Get(), o.PublisherName.IsSet()
}

// HasPublisherName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionApp) HasPublisherName() bool {
	if o != nil && o.PublisherName.IsSet() {
		return true
	}

	return false
}

// SetPublisherName gets a reference to the given NullableString and assigns it to the PublisherName field.
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetPublisherName(v string) {
	o.PublisherName.Set(&v)
}
// SetPublisherNameNil sets the value for PublisherName to be an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionApp) SetPublisherNameNil() {
	o.PublisherName.Set(nil)
}

// UnsetPublisherName ensures that no value is present for PublisherName, not even an explicit nil
func (o *MicrosoftGraphWindowsInformationProtectionApp) UnsetPublisherName() {
	o.PublisherName.Unset()
}

func (o MicrosoftGraphWindowsInformationProtectionApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Denied != nil {
		toSerialize["denied"] = o.Denied
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if o.PublisherName.IsSet() {
		toSerialize["publisherName"] = o.PublisherName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWindowsInformationProtectionApp struct {
	value *MicrosoftGraphWindowsInformationProtectionApp
	isSet bool
}

func (v NullableMicrosoftGraphWindowsInformationProtectionApp) Get() *MicrosoftGraphWindowsInformationProtectionApp {
	return v.value
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionApp) Set(val *MicrosoftGraphWindowsInformationProtectionApp) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWindowsInformationProtectionApp) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWindowsInformationProtectionApp(val *MicrosoftGraphWindowsInformationProtectionApp) *NullableMicrosoftGraphWindowsInformationProtectionApp {
	return &NullableMicrosoftGraphWindowsInformationProtectionApp{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWindowsInformationProtectionApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWindowsInformationProtectionApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


