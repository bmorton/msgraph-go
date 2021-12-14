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

// MicrosoftGraphResourceOperation struct for MicrosoftGraphResourceOperation
type MicrosoftGraphResourceOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
	ActionName NullableString `json:"actionName,omitempty"`
	// Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
	Description NullableString `json:"description,omitempty"`
	// Name of the Resource this operation is performed on.
	ResourceName NullableString `json:"resourceName,omitempty"`
}

// NewMicrosoftGraphResourceOperation instantiates a new MicrosoftGraphResourceOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphResourceOperation() *MicrosoftGraphResourceOperation {
	this := MicrosoftGraphResourceOperation{}
	return &this
}

// NewMicrosoftGraphResourceOperationWithDefaults instantiates a new MicrosoftGraphResourceOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphResourceOperationWithDefaults() *MicrosoftGraphResourceOperation {
	this := MicrosoftGraphResourceOperation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphResourceOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResourceOperation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphResourceOperation) SetId(v string) {
	o.Id = &v
}

// GetActionName returns the ActionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceOperation) GetActionName() string {
	if o == nil || o.ActionName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActionName.Get()
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceOperation) GetActionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionName.Get(), o.ActionName.IsSet()
}

// HasActionName returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasActionName() bool {
	if o != nil && o.ActionName.IsSet() {
		return true
	}

	return false
}

// SetActionName gets a reference to the given NullableString and assigns it to the ActionName field.
func (o *MicrosoftGraphResourceOperation) SetActionName(v string) {
	o.ActionName.Set(&v)
}
// SetActionNameNil sets the value for ActionName to be an explicit nil
func (o *MicrosoftGraphResourceOperation) SetActionNameNil() {
	o.ActionName.Set(nil)
}

// UnsetActionName ensures that no value is present for ActionName, not even an explicit nil
func (o *MicrosoftGraphResourceOperation) UnsetActionName() {
	o.ActionName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceOperation) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceOperation) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphResourceOperation) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphResourceOperation) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphResourceOperation) UnsetDescription() {
	o.Description.Unset()
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceOperation) GetResourceName() string {
	if o == nil || o.ResourceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceName.Get()
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceOperation) GetResourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceName.Get(), o.ResourceName.IsSet()
}

// HasResourceName returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasResourceName() bool {
	if o != nil && o.ResourceName.IsSet() {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given NullableString and assigns it to the ResourceName field.
func (o *MicrosoftGraphResourceOperation) SetResourceName(v string) {
	o.ResourceName.Set(&v)
}
// SetResourceNameNil sets the value for ResourceName to be an explicit nil
func (o *MicrosoftGraphResourceOperation) SetResourceNameNil() {
	o.ResourceName.Set(nil)
}

// UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
func (o *MicrosoftGraphResourceOperation) UnsetResourceName() {
	o.ResourceName.Unset()
}

func (o MicrosoftGraphResourceOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphResourceOperation struct {
	value *MicrosoftGraphResourceOperation
	isSet bool
}

func (v NullableMicrosoftGraphResourceOperation) Get() *MicrosoftGraphResourceOperation {
	return v.value
}

func (v *NullableMicrosoftGraphResourceOperation) Set(val *MicrosoftGraphResourceOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphResourceOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphResourceOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphResourceOperation(val *MicrosoftGraphResourceOperation) *NullableMicrosoftGraphResourceOperation {
	return &NullableMicrosoftGraphResourceOperation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphResourceOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphResourceOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


