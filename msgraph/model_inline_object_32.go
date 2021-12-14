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

// InlineObject32 struct for InlineObject32
type InlineObject32 struct {
	Status AnyOfmicrosoftGraphRecordingStatus `json:"status,omitempty"`
	ClientContext NullableString `json:"clientContext,omitempty"`
}

// NewInlineObject32 instantiates a new InlineObject32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject32() *InlineObject32 {
	this := InlineObject32{}
	return &this
}

// NewInlineObject32WithDefaults instantiates a new InlineObject32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject32WithDefaults() *InlineObject32 {
	this := InlineObject32{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject32) GetStatus() AnyOfmicrosoftGraphRecordingStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphRecordingStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject32) GetStatusOk() (*AnyOfmicrosoftGraphRecordingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineObject32) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphRecordingStatus and assigns it to the Status field.
func (o *InlineObject32) SetStatus(v AnyOfmicrosoftGraphRecordingStatus) {
	o.Status = v
}

// GetClientContext returns the ClientContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject32) GetClientContext() string {
	if o == nil || o.ClientContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientContext.Get()
}

// GetClientContextOk returns a tuple with the ClientContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject32) GetClientContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientContext.Get(), o.ClientContext.IsSet()
}

// HasClientContext returns a boolean if a field has been set.
func (o *InlineObject32) HasClientContext() bool {
	if o != nil && o.ClientContext.IsSet() {
		return true
	}

	return false
}

// SetClientContext gets a reference to the given NullableString and assigns it to the ClientContext field.
func (o *InlineObject32) SetClientContext(v string) {
	o.ClientContext.Set(&v)
}
// SetClientContextNil sets the value for ClientContext to be an explicit nil
func (o *InlineObject32) SetClientContextNil() {
	o.ClientContext.Set(nil)
}

// UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
func (o *InlineObject32) UnsetClientContext() {
	o.ClientContext.Unset()
}

func (o InlineObject32) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ClientContext.IsSet() {
		toSerialize["clientContext"] = o.ClientContext.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject32 struct {
	value *InlineObject32
	isSet bool
}

func (v NullableInlineObject32) Get() *InlineObject32 {
	return v.value
}

func (v *NullableInlineObject32) Set(val *InlineObject32) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject32) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject32(val *InlineObject32) *NullableInlineObject32 {
	return &NullableInlineObject32{value: val, isSet: true}
}

func (v NullableInlineObject32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


