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

// MicrosoftGraphOnenoteOperation struct for MicrosoftGraphOnenoteOperation
type MicrosoftGraphOnenoteOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The start time of the operation.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The time of the last action of the operation.
	LastActionDateTime NullableTime `json:"lastActionDateTime,omitempty"`
	// The current status of the operation: notStarted, running, completed, failed
	Status AnyOfmicrosoftGraphOperationStatus `json:"status,omitempty"`
	// The error returned by the operation.
	Error AnyOfmicrosoftGraphOnenoteOperationError `json:"error,omitempty"`
	// The operation percent complete if the operation is still in running status.
	PercentComplete NullableString `json:"percentComplete,omitempty"`
	// The resource id.
	ResourceId NullableString `json:"resourceId,omitempty"`
	// The resource URI for the object. For example, the resource URI for a copied page or section.
	ResourceLocation NullableString `json:"resourceLocation,omitempty"`
}

// NewMicrosoftGraphOnenoteOperation instantiates a new MicrosoftGraphOnenoteOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOnenoteOperation() *MicrosoftGraphOnenoteOperation {
	this := MicrosoftGraphOnenoteOperation{}
	return &this
}

// NewMicrosoftGraphOnenoteOperationWithDefaults instantiates a new MicrosoftGraphOnenoteOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOnenoteOperationWithDefaults() *MicrosoftGraphOnenoteOperation {
	this := MicrosoftGraphOnenoteOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOnenoteOperation) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphOnenoteOperation) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastActionDateTime returns the LastActionDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime.Get()
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastActionDateTime.Get(), o.LastActionDateTime.IsSet()
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given NullableTime and assigns it to the LastActionDateTime field.
func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime.Set(&v)
}
// SetLastActionDateTimeNil sets the value for LastActionDateTime to be an explicit nil
func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTimeNil() {
	o.LastActionDateTime.Set(nil)
}

// UnsetLastActionDateTime ensures that no value is present for LastActionDateTime, not even an explicit nil
func (o *MicrosoftGraphOnenoteOperation) UnsetLastActionDateTime() {
	o.LastActionDateTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphOnenoteOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetError() AnyOfmicrosoftGraphOnenoteOperationError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnenoteOperationError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetErrorOk() (*AnyOfmicrosoftGraphOnenoteOperationError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphOnenoteOperationError and assigns it to the Error field.
func (o *MicrosoftGraphOnenoteOperation) SetError(v AnyOfmicrosoftGraphOnenoteOperationError) {
	o.Error = v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetPercentComplete() string {
	if o == nil || o.PercentComplete.Get() == nil {
		var ret string
		return ret
	}
	return *o.PercentComplete.Get()
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetPercentCompleteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PercentComplete.Get(), o.PercentComplete.IsSet()
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasPercentComplete() bool {
	if o != nil && o.PercentComplete.IsSet() {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given NullableString and assigns it to the PercentComplete field.
func (o *MicrosoftGraphOnenoteOperation) SetPercentComplete(v string) {
	o.PercentComplete.Set(&v)
}
// SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil
func (o *MicrosoftGraphOnenoteOperation) SetPercentCompleteNil() {
	o.PercentComplete.Set(nil)
}

// UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
func (o *MicrosoftGraphOnenoteOperation) UnsetPercentComplete() {
	o.PercentComplete.Unset()
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetResourceId() string {
	if o == nil || o.ResourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceId.Get()
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceId.Get(), o.ResourceId.IsSet()
}

// HasResourceId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasResourceId() bool {
	if o != nil && o.ResourceId.IsSet() {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given NullableString and assigns it to the ResourceId field.
func (o *MicrosoftGraphOnenoteOperation) SetResourceId(v string) {
	o.ResourceId.Set(&v)
}
// SetResourceIdNil sets the value for ResourceId to be an explicit nil
func (o *MicrosoftGraphOnenoteOperation) SetResourceIdNil() {
	o.ResourceId.Set(nil)
}

// UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
func (o *MicrosoftGraphOnenoteOperation) UnsetResourceId() {
	o.ResourceId.Unset()
}

// GetResourceLocation returns the ResourceLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenoteOperation) GetResourceLocation() string {
	if o == nil || o.ResourceLocation.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceLocation.Get()
}

// GetResourceLocationOk returns a tuple with the ResourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenoteOperation) GetResourceLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceLocation.Get(), o.ResourceLocation.IsSet()
}

// HasResourceLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasResourceLocation() bool {
	if o != nil && o.ResourceLocation.IsSet() {
		return true
	}

	return false
}

// SetResourceLocation gets a reference to the given NullableString and assigns it to the ResourceLocation field.
func (o *MicrosoftGraphOnenoteOperation) SetResourceLocation(v string) {
	o.ResourceLocation.Set(&v)
}
// SetResourceLocationNil sets the value for ResourceLocation to be an explicit nil
func (o *MicrosoftGraphOnenoteOperation) SetResourceLocationNil() {
	o.ResourceLocation.Set(nil)
}

// UnsetResourceLocation ensures that no value is present for ResourceLocation, not even an explicit nil
func (o *MicrosoftGraphOnenoteOperation) UnsetResourceLocation() {
	o.ResourceLocation.Unset()
}

func (o MicrosoftGraphOnenoteOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.LastActionDateTime.IsSet() {
		toSerialize["lastActionDateTime"] = o.LastActionDateTime.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.PercentComplete.IsSet() {
		toSerialize["percentComplete"] = o.PercentComplete.Get()
	}
	if o.ResourceId.IsSet() {
		toSerialize["resourceId"] = o.ResourceId.Get()
	}
	if o.ResourceLocation.IsSet() {
		toSerialize["resourceLocation"] = o.ResourceLocation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOnenoteOperation struct {
	value *MicrosoftGraphOnenoteOperation
	isSet bool
}

func (v NullableMicrosoftGraphOnenoteOperation) Get() *MicrosoftGraphOnenoteOperation {
	return v.value
}

func (v *NullableMicrosoftGraphOnenoteOperation) Set(val *MicrosoftGraphOnenoteOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOnenoteOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOnenoteOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOnenoteOperation(val *MicrosoftGraphOnenoteOperation) *NullableMicrosoftGraphOnenoteOperation {
	return &NullableMicrosoftGraphOnenoteOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOnenoteOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOnenoteOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


