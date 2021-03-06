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

// IdentityUserFlowAttribute struct for IdentityUserFlowAttribute
type IdentityUserFlowAttribute struct {
	// The data type of the user flow attribute. This cannot be modified after the custom user flow attribute is created. The supported values for dataType are: string , boolean , int64 , stringCollection , dateTime.
	DataType AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType `json:"dataType,omitempty"`
	// The description of the user flow attribute that's shown to the user at the time of sign-up.
	Description NullableString `json:"description,omitempty"`
	// The display name of the user flow attribute.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The type of the user flow attribute. This is a read-only attribute that is automatically set. Depending on the type of attribute, the values for this property will be builtIn, custom, or required.
	UserFlowAttributeType AnyOfmicrosoftGraphIdentityUserFlowAttributeType `json:"userFlowAttributeType,omitempty"`
}

// NewIdentityUserFlowAttribute instantiates a new IdentityUserFlowAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUserFlowAttribute() *IdentityUserFlowAttribute {
	this := IdentityUserFlowAttribute{}
	return &this
}

// NewIdentityUserFlowAttributeWithDefaults instantiates a new IdentityUserFlowAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUserFlowAttributeWithDefaults() *IdentityUserFlowAttribute {
	this := IdentityUserFlowAttribute{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserFlowAttribute) GetDataType() AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType
		return ret
	}
	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserFlowAttribute) GetDataTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return &o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *IdentityUserFlowAttribute) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType and assigns it to the DataType field.
func (o *IdentityUserFlowAttribute) SetDataType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeDataType) {
	o.DataType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserFlowAttribute) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserFlowAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityUserFlowAttribute) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IdentityUserFlowAttribute) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IdentityUserFlowAttribute) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IdentityUserFlowAttribute) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserFlowAttribute) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserFlowAttribute) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityUserFlowAttribute) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *IdentityUserFlowAttribute) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *IdentityUserFlowAttribute) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *IdentityUserFlowAttribute) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetUserFlowAttributeType returns the UserFlowAttributeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityUserFlowAttribute) GetUserFlowAttributeType() AnyOfmicrosoftGraphIdentityUserFlowAttributeType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentityUserFlowAttributeType
		return ret
	}
	return o.UserFlowAttributeType
}

// GetUserFlowAttributeTypeOk returns a tuple with the UserFlowAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityUserFlowAttribute) GetUserFlowAttributeTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeType, bool) {
	if o == nil || o.UserFlowAttributeType == nil {
		return nil, false
	}
	return &o.UserFlowAttributeType, true
}

// HasUserFlowAttributeType returns a boolean if a field has been set.
func (o *IdentityUserFlowAttribute) HasUserFlowAttributeType() bool {
	if o != nil && o.UserFlowAttributeType != nil {
		return true
	}

	return false
}

// SetUserFlowAttributeType gets a reference to the given AnyOfmicrosoftGraphIdentityUserFlowAttributeType and assigns it to the UserFlowAttributeType field.
func (o *IdentityUserFlowAttribute) SetUserFlowAttributeType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeType) {
	o.UserFlowAttributeType = v
}

func (o IdentityUserFlowAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.UserFlowAttributeType != nil {
		toSerialize["userFlowAttributeType"] = o.UserFlowAttributeType
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityUserFlowAttribute struct {
	value *IdentityUserFlowAttribute
	isSet bool
}

func (v NullableIdentityUserFlowAttribute) Get() *IdentityUserFlowAttribute {
	return v.value
}

func (v *NullableIdentityUserFlowAttribute) Set(val *IdentityUserFlowAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUserFlowAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUserFlowAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUserFlowAttribute(val *IdentityUserFlowAttribute) *NullableIdentityUserFlowAttribute {
	return &NullableIdentityUserFlowAttribute{value: val, isSet: true}
}

func (v NullableIdentityUserFlowAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUserFlowAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


