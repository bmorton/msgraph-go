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

// MicrosoftGraphWorkbookOperation struct for MicrosoftGraphWorkbookOperation
type MicrosoftGraphWorkbookOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The error returned by the operation.
	Error AnyOfmicrosoftGraphWorkbookOperationError `json:"error,omitempty"`
	// The resource URI for the result.
	ResourceLocation NullableString `json:"resourceLocation,omitempty"`
	// The current status of the operation. Possible values are: NotStarted, Running, Completed, Failed.
	Status AnyOfmicrosoftGraphWorkbookOperationStatus `json:"status,omitempty"`
}

// NewMicrosoftGraphWorkbookOperation instantiates a new MicrosoftGraphWorkbookOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookOperation() *MicrosoftGraphWorkbookOperation {
	this := MicrosoftGraphWorkbookOperation{}
	return &this
}

// NewMicrosoftGraphWorkbookOperationWithDefaults instantiates a new MicrosoftGraphWorkbookOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookOperationWithDefaults() *MicrosoftGraphWorkbookOperation {
	this := MicrosoftGraphWorkbookOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookOperation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookOperation) SetId(v string) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookOperation) GetError() AnyOfmicrosoftGraphWorkbookOperationError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookOperationError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookOperation) GetErrorOk() (*AnyOfmicrosoftGraphWorkbookOperationError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookOperation) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphWorkbookOperationError and assigns it to the Error field.
func (o *MicrosoftGraphWorkbookOperation) SetError(v AnyOfmicrosoftGraphWorkbookOperationError) {
	o.Error = v
}

// GetResourceLocation returns the ResourceLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookOperation) GetResourceLocation() string {
	if o == nil || o.ResourceLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceLocation.Get()
}

// GetResourceLocationOk returns a tuple with the ResourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookOperation) GetResourceLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceLocation.Get(), o.ResourceLocation.IsSet()
}

// HasResourceLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookOperation) HasResourceLocation() bool {
	if o != nil && o.ResourceLocation.IsSet() {
		return true
	}

	return false
}

// SetResourceLocation gets a reference to the given NullableString and assigns it to the ResourceLocation field.
func (o *MicrosoftGraphWorkbookOperation) SetResourceLocation(v string) {
	o.ResourceLocation.Set(&v)
}
// SetResourceLocationNil sets the value for ResourceLocation to be an explicit nil
func (o *MicrosoftGraphWorkbookOperation) SetResourceLocationNil() {
	o.ResourceLocation.Set(nil)
}

// UnsetResourceLocation ensures that no value is present for ResourceLocation, not even an explicit nil
func (o *MicrosoftGraphWorkbookOperation) UnsetResourceLocation() {
	o.ResourceLocation.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookOperation) GetStatus() AnyOfmicrosoftGraphWorkbookOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookOperationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookOperation) GetStatusOk() (*AnyOfmicrosoftGraphWorkbookOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphWorkbookOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphWorkbookOperation) SetStatus(v AnyOfmicrosoftGraphWorkbookOperationStatus) {
	o.Status = v
}

func (o MicrosoftGraphWorkbookOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ResourceLocation.IsSet() {
		toSerialize["resourceLocation"] = o.ResourceLocation.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookOperation struct {
	value *MicrosoftGraphWorkbookOperation
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookOperation) Get() *MicrosoftGraphWorkbookOperation {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookOperation) Set(val *MicrosoftGraphWorkbookOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookOperation(val *MicrosoftGraphWorkbookOperation) *NullableMicrosoftGraphWorkbookOperation {
	return &NullableMicrosoftGraphWorkbookOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

