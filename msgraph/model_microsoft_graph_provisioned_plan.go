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

// MicrosoftGraphProvisionedPlan struct for MicrosoftGraphProvisionedPlan
type MicrosoftGraphProvisionedPlan struct {
	// For example, 'Enabled'.
	CapabilityStatus NullableString `json:"capabilityStatus,omitempty"`
	// For example, 'Success'.
	ProvisioningStatus NullableString `json:"provisioningStatus,omitempty"`
	// The name of the service; for example, 'AccessControlS2S'
	Service NullableString `json:"service,omitempty"`
}

// NewMicrosoftGraphProvisionedPlan instantiates a new MicrosoftGraphProvisionedPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphProvisionedPlan() *MicrosoftGraphProvisionedPlan {
	this := MicrosoftGraphProvisionedPlan{}
	return &this
}

// NewMicrosoftGraphProvisionedPlanWithDefaults instantiates a new MicrosoftGraphProvisionedPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphProvisionedPlanWithDefaults() *MicrosoftGraphProvisionedPlan {
	this := MicrosoftGraphProvisionedPlan{}
	return &this
}

// GetCapabilityStatus returns the CapabilityStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphProvisionedPlan) GetCapabilityStatus() string {
	if o == nil || o.CapabilityStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.CapabilityStatus.Get()
}

// GetCapabilityStatusOk returns a tuple with the CapabilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphProvisionedPlan) GetCapabilityStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CapabilityStatus.Get(), o.CapabilityStatus.IsSet()
}

// HasCapabilityStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphProvisionedPlan) HasCapabilityStatus() bool {
	if o != nil && o.CapabilityStatus.IsSet() {
		return true
	}

	return false
}

// SetCapabilityStatus gets a reference to the given NullableString and assigns it to the CapabilityStatus field.
func (o *MicrosoftGraphProvisionedPlan) SetCapabilityStatus(v string) {
	o.CapabilityStatus.Set(&v)
}
// SetCapabilityStatusNil sets the value for CapabilityStatus to be an explicit nil
func (o *MicrosoftGraphProvisionedPlan) SetCapabilityStatusNil() {
	o.CapabilityStatus.Set(nil)
}

// UnsetCapabilityStatus ensures that no value is present for CapabilityStatus, not even an explicit nil
func (o *MicrosoftGraphProvisionedPlan) UnsetCapabilityStatus() {
	o.CapabilityStatus.Unset()
}

// GetProvisioningStatus returns the ProvisioningStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphProvisionedPlan) GetProvisioningStatus() string {
	if o == nil || o.ProvisioningStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProvisioningStatus.Get()
}

// GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphProvisionedPlan) GetProvisioningStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProvisioningStatus.Get(), o.ProvisioningStatus.IsSet()
}

// HasProvisioningStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphProvisionedPlan) HasProvisioningStatus() bool {
	if o != nil && o.ProvisioningStatus.IsSet() {
		return true
	}

	return false
}

// SetProvisioningStatus gets a reference to the given NullableString and assigns it to the ProvisioningStatus field.
func (o *MicrosoftGraphProvisionedPlan) SetProvisioningStatus(v string) {
	o.ProvisioningStatus.Set(&v)
}
// SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil
func (o *MicrosoftGraphProvisionedPlan) SetProvisioningStatusNil() {
	o.ProvisioningStatus.Set(nil)
}

// UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
func (o *MicrosoftGraphProvisionedPlan) UnsetProvisioningStatus() {
	o.ProvisioningStatus.Unset()
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphProvisionedPlan) GetService() string {
	if o == nil || o.Service.Get() == nil {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphProvisionedPlan) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *MicrosoftGraphProvisionedPlan) HasService() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *MicrosoftGraphProvisionedPlan) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *MicrosoftGraphProvisionedPlan) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *MicrosoftGraphProvisionedPlan) UnsetService() {
	o.Service.Unset()
}

func (o MicrosoftGraphProvisionedPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapabilityStatus.IsSet() {
		toSerialize["capabilityStatus"] = o.CapabilityStatus.Get()
	}
	if o.ProvisioningStatus.IsSet() {
		toSerialize["provisioningStatus"] = o.ProvisioningStatus.Get()
	}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphProvisionedPlan struct {
	value *MicrosoftGraphProvisionedPlan
	isSet bool
}

func (v NullableMicrosoftGraphProvisionedPlan) Get() *MicrosoftGraphProvisionedPlan {
	return v.value
}

func (v *NullableMicrosoftGraphProvisionedPlan) Set(val *MicrosoftGraphProvisionedPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphProvisionedPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphProvisionedPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphProvisionedPlan(val *MicrosoftGraphProvisionedPlan) *NullableMicrosoftGraphProvisionedPlan {
	return &NullableMicrosoftGraphProvisionedPlan{value: val, isSet: true}
}

func (v NullableMicrosoftGraphProvisionedPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphProvisionedPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


