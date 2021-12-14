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

// ResourceOperation Describes the resourceOperation resource (entity) of the Microsoft Graph API (REST), which supports Intune workflows related to role-based access control (RBAC).
type ResourceOperation struct {
	// Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
	ActionName NullableString `json:"actionName,omitempty"`
	// Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
	Description NullableString `json:"description,omitempty"`
	// Name of the Resource this operation is performed on.
	ResourceName NullableString `json:"resourceName,omitempty"`
}

// NewResourceOperation instantiates a new ResourceOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOperation() *ResourceOperation {
	this := ResourceOperation{}
	return &this
}

// NewResourceOperationWithDefaults instantiates a new ResourceOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOperationWithDefaults() *ResourceOperation {
	this := ResourceOperation{}
	return &this
}

// GetActionName returns the ActionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceOperation) GetActionName() string {
	if o == nil || o.ActionName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionName.Get()
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceOperation) GetActionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionName.Get(), o.ActionName.IsSet()
}

// HasActionName returns a boolean if a field has been set.
func (o *ResourceOperation) HasActionName() bool {
	if o != nil && o.ActionName.IsSet() {
		return true
	}

	return false
}

// SetActionName gets a reference to the given NullableString and assigns it to the ActionName field.
func (o *ResourceOperation) SetActionName(v string) {
	o.ActionName.Set(&v)
}
// SetActionNameNil sets the value for ActionName to be an explicit nil
func (o *ResourceOperation) SetActionNameNil() {
	o.ActionName.Set(nil)
}

// UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
func (o *ResourceOperation) UnsetActionName() {
	o.ActionName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceOperation) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceOperation) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceOperation) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ResourceOperation) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ResourceOperation) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ResourceOperation) UnsetDescription() {
	o.Description.Unset()
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceOperation) GetResourceName() string {
	if o == nil || o.ResourceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceName.Get()
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceOperation) GetResourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceName.Get(), o.ResourceName.IsSet()
}

// HasResourceName returns a boolean if a field has been set.
func (o *ResourceOperation) HasResourceName() bool {
	if o != nil && o.ResourceName.IsSet() {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given NullableString and assigns it to the ResourceName field.
func (o *ResourceOperation) SetResourceName(v string) {
	o.ResourceName.Set(&v)
}
// SetResourceNameNil sets the value for ResourceName to be an explicit nil
func (o *ResourceOperation) SetResourceNameNil() {
	o.ResourceName.Set(nil)
}

// UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
func (o *ResourceOperation) UnsetResourceName() {
	o.ResourceName.Unset()
}

func (o ResourceOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionName.IsSet() {
		toSerialize["actionName"] = o.ActionName.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ResourceName.IsSet() {
		toSerialize["resourceName"] = o.ResourceName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceOperation struct {
	value *ResourceOperation
	isSet bool
}

func (v NullableResourceOperation) Get() *ResourceOperation {
	return v.value
}

func (v *NullableResourceOperation) Set(val *ResourceOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOperation(val *ResourceOperation) *NullableResourceOperation {
	return &NullableResourceOperation{value: val, isSet: true}
}

func (v NullableResourceOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


