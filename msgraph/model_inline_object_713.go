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

// InlineObject713 struct for InlineObject713
type InlineObject713 struct {
	DisplayName *string `json:"displayName,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Model *string `json:"model,omitempty"`
	PhysicalDeviceId NullableString `json:"physicalDeviceId,omitempty"`
	HasPhysicalDevice NullableBool `json:"hasPhysicalDevice,omitempty"`
	CertificateSigningRequest *MicrosoftGraphPrintCertificateSigningRequest `json:"certificateSigningRequest,omitempty"`
	ConnectorId NullableString `json:"connectorId,omitempty"`
}

// NewInlineObject713 instantiates a new InlineObject713 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject713() *InlineObject713 {
	this := InlineObject713{}
	var hasPhysicalDevice bool = false
	this.HasPhysicalDevice = *NewNullableBool(&hasPhysicalDevice)
	return &this
}

// NewInlineObject713WithDefaults instantiates a new InlineObject713 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject713WithDefaults() *InlineObject713 {
	this := InlineObject713{}
	var hasPhysicalDevice bool = false
	this.HasPhysicalDevice = *NewNullableBool(&hasPhysicalDevice)
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InlineObject713) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject713) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InlineObject713) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InlineObject713) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *InlineObject713) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject713) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InlineObject713) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *InlineObject713) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineObject713) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject713) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineObject713) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineObject713) SetModel(v string) {
	o.Model = &v
}

// GetPhysicalDeviceId returns the PhysicalDeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject713) GetPhysicalDeviceId() string {
	if o == nil || o.PhysicalDeviceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhysicalDeviceId.Get()
}

// GetPhysicalDeviceIdOk returns a tuple with the PhysicalDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject713) GetPhysicalDeviceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhysicalDeviceId.Get(), o.PhysicalDeviceId.IsSet()
}

// HasPhysicalDeviceId returns a boolean if a field has been set.
func (o *InlineObject713) HasPhysicalDeviceId() bool {
	if o != nil && o.PhysicalDeviceId.IsSet() {
		return true
	}

	return false
}

// SetPhysicalDeviceId gets a reference to the given NullableString and assigns it to the PhysicalDeviceId field.
func (o *InlineObject713) SetPhysicalDeviceId(v string) {
	o.PhysicalDeviceId.Set(&v)
}
// SetPhysicalDeviceIdNil sets the value for PhysicalDeviceId to be an explicit nil
func (o *InlineObject713) SetPhysicalDeviceIdNil() {
	o.PhysicalDeviceId.Set(nil)
}

// UnsetPhysicalDeviceId ensures that no value is present for PhysicalDeviceId, not even an explicit nil
func (o *InlineObject713) UnsetPhysicalDeviceId() {
	o.PhysicalDeviceId.Unset()
}

// GetHasPhysicalDevice returns the HasPhysicalDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject713) GetHasPhysicalDevice() bool {
	if o == nil || o.HasPhysicalDevice.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasPhysicalDevice.Get()
}

// GetHasPhysicalDeviceOk returns a tuple with the HasPhysicalDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject713) GetHasPhysicalDeviceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasPhysicalDevice.Get(), o.HasPhysicalDevice.IsSet()
}

// HasHasPhysicalDevice returns a boolean if a field has been set.
func (o *InlineObject713) HasHasPhysicalDevice() bool {
	if o != nil && o.HasPhysicalDevice.IsSet() {
		return true
	}

	return false
}

// SetHasPhysicalDevice gets a reference to the given NullableBool and assigns it to the HasPhysicalDevice field.
func (o *InlineObject713) SetHasPhysicalDevice(v bool) {
	o.HasPhysicalDevice.Set(&v)
}
// SetHasPhysicalDeviceNil sets the value for HasPhysicalDevice to be an explicit nil
func (o *InlineObject713) SetHasPhysicalDeviceNil() {
	o.HasPhysicalDevice.Set(nil)
}

// UnsetHasPhysicalDevice ensures that no value is present for HasPhysicalDevice, not even an explicit nil
func (o *InlineObject713) UnsetHasPhysicalDevice() {
	o.HasPhysicalDevice.Unset()
}

// GetCertificateSigningRequest returns the CertificateSigningRequest field value if set, zero value otherwise.
func (o *InlineObject713) GetCertificateSigningRequest() MicrosoftGraphPrintCertificateSigningRequest {
	if o == nil || o.CertificateSigningRequest == nil {
		var ret MicrosoftGraphPrintCertificateSigningRequest
		return ret
	}
	return *o.CertificateSigningRequest
}

// GetCertificateSigningRequestOk returns a tuple with the CertificateSigningRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject713) GetCertificateSigningRequestOk() (*MicrosoftGraphPrintCertificateSigningRequest, bool) {
	if o == nil || o.CertificateSigningRequest == nil {
		return nil, false
	}
	return o.CertificateSigningRequest, true
}

// HasCertificateSigningRequest returns a boolean if a field has been set.
func (o *InlineObject713) HasCertificateSigningRequest() bool {
	if o != nil && o.CertificateSigningRequest != nil {
		return true
	}

	return false
}

// SetCertificateSigningRequest gets a reference to the given MicrosoftGraphPrintCertificateSigningRequest and assigns it to the CertificateSigningRequest field.
func (o *InlineObject713) SetCertificateSigningRequest(v MicrosoftGraphPrintCertificateSigningRequest) {
	o.CertificateSigningRequest = &v
}

// GetConnectorId returns the ConnectorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject713) GetConnectorId() string {
	if o == nil || o.ConnectorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectorId.Get()
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject713) GetConnectorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectorId.Get(), o.ConnectorId.IsSet()
}

// HasConnectorId returns a boolean if a field has been set.
func (o *InlineObject713) HasConnectorId() bool {
	if o != nil && o.ConnectorId.IsSet() {
		return true
	}

	return false
}

// SetConnectorId gets a reference to the given NullableString and assigns it to the ConnectorId field.
func (o *InlineObject713) SetConnectorId(v string) {
	o.ConnectorId.Set(&v)
}
// SetConnectorIdNil sets the value for ConnectorId to be an explicit nil
func (o *InlineObject713) SetConnectorIdNil() {
	o.ConnectorId.Set(nil)
}

// UnsetConnectorId ensures that no value is present for ConnectorId, not even an explicit nil
func (o *InlineObject713) UnsetConnectorId() {
	o.ConnectorId.Unset()
}

func (o InlineObject713) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.PhysicalDeviceId.IsSet() {
		toSerialize["physicalDeviceId"] = o.PhysicalDeviceId.Get()
	}
	if o.HasPhysicalDevice.IsSet() {
		toSerialize["hasPhysicalDevice"] = o.HasPhysicalDevice.Get()
	}
	if o.CertificateSigningRequest != nil {
		toSerialize["certificateSigningRequest"] = o.CertificateSigningRequest
	}
	if o.ConnectorId.IsSet() {
		toSerialize["connectorId"] = o.ConnectorId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject713 struct {
	value *InlineObject713
	isSet bool
}

func (v NullableInlineObject713) Get() *InlineObject713 {
	return v.value
}

func (v *NullableInlineObject713) Set(val *InlineObject713) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject713) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject713) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject713(val *InlineObject713) *NullableInlineObject713 {
	return &NullableInlineObject713{value: val, isSet: true}
}

func (v NullableInlineObject713) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject713) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

