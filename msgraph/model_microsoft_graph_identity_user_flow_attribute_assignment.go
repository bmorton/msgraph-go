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

// MicrosoftGraphIdentityUserFlowAttributeAssignment struct for MicrosoftGraphIdentityUserFlowAttributeAssignment
type MicrosoftGraphIdentityUserFlowAttributeAssignment struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The display name of the identityUserFlowAttribute within a user flow.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
	IsOptional *bool `json:"isOptional,omitempty"`
	// Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
	RequiresVerification *bool `json:"requiresVerification,omitempty"`
	// The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
	UserAttributeValues *[]*AnyOfmicrosoftGraphUserAttributeValuesItem `json:"userAttributeValues,omitempty"`
	// The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect.
	UserInputType AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType `json:"userInputType,omitempty"`
	// The user attribute that you want to add to your user flow.
	UserAttribute AnyOfmicrosoftGraphIdentityUserFlowAttribute `json:"userAttribute,omitempty"`
}

// NewMicrosoftGraphIdentityUserFlowAttributeAssignment instantiates a new MicrosoftGraphIdentityUserFlowAttributeAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIdentityUserFlowAttributeAssignment() *MicrosoftGraphIdentityUserFlowAttributeAssignment {
	this := MicrosoftGraphIdentityUserFlowAttributeAssignment{}
	return &this
}

// NewMicrosoftGraphIdentityUserFlowAttributeAssignmentWithDefaults instantiates a new MicrosoftGraphIdentityUserFlowAttributeAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIdentityUserFlowAttributeAssignmentWithDefaults() *MicrosoftGraphIdentityUserFlowAttributeAssignment {
	this := MicrosoftGraphIdentityUserFlowAttributeAssignment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsOptional returns the IsOptional field value if set, zero value otherwise.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetIsOptional() bool {
	if o == nil || o.IsOptional == nil {
		var ret bool
		return ret
	}
	return *o.IsOptional
}

// GetIsOptionalOk returns a tuple with the IsOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetIsOptionalOk() (*bool, bool) {
	if o == nil || o.IsOptional == nil {
		return nil, false
	}
	return o.IsOptional, true
}

// HasIsOptional returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasIsOptional() bool {
	if o != nil && o.IsOptional != nil {
		return true
	}

	return false
}

// SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetIsOptional(v bool) {
	o.IsOptional = &v
}

// GetRequiresVerification returns the RequiresVerification field value if set, zero value otherwise.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetRequiresVerification() bool {
	if o == nil || o.RequiresVerification == nil {
		var ret bool
		return ret
	}
	return *o.RequiresVerification
}

// GetRequiresVerificationOk returns a tuple with the RequiresVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetRequiresVerificationOk() (*bool, bool) {
	if o == nil || o.RequiresVerification == nil {
		return nil, false
	}
	return o.RequiresVerification, true
}

// HasRequiresVerification returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasRequiresVerification() bool {
	if o != nil && o.RequiresVerification != nil {
		return true
	}

	return false
}

// SetRequiresVerification gets a reference to the given bool and assigns it to the RequiresVerification field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetRequiresVerification(v bool) {
	o.RequiresVerification = &v
}

// GetUserAttributeValues returns the UserAttributeValues field value if set, zero value otherwise.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttributeValues() []*AnyOfmicrosoftGraphUserAttributeValuesItem {
	if o == nil || o.UserAttributeValues == nil {
		var ret []*AnyOfmicrosoftGraphUserAttributeValuesItem
		return ret
	}
	return *o.UserAttributeValues
}

// GetUserAttributeValuesOk returns a tuple with the UserAttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttributeValuesOk() (*[]*AnyOfmicrosoftGraphUserAttributeValuesItem, bool) {
	if o == nil || o.UserAttributeValues == nil {
		return nil, false
	}
	return o.UserAttributeValues, true
}

// HasUserAttributeValues returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasUserAttributeValues() bool {
	if o != nil && o.UserAttributeValues != nil {
		return true
	}

	return false
}

// SetUserAttributeValues gets a reference to the given []*AnyOfmicrosoftGraphUserAttributeValuesItem and assigns it to the UserAttributeValues field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserAttributeValues(v []*AnyOfmicrosoftGraphUserAttributeValuesItem) {
	o.UserAttributeValues = &v
}

// GetUserInputType returns the UserInputType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserInputType() AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType
		return ret
	}
	return o.UserInputType
}

// GetUserInputTypeOk returns a tuple with the UserInputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserInputTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType, bool) {
	if o == nil || o.UserInputType == nil {
		return nil, false
	}
	return &o.UserInputType, true
}

// HasUserInputType returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasUserInputType() bool {
	if o != nil && o.UserInputType != nil {
		return true
	}

	return false
}

// SetUserInputType gets a reference to the given AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType and assigns it to the UserInputType field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserInputType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType) {
	o.UserInputType = v
}

// GetUserAttribute returns the UserAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttribute() AnyOfmicrosoftGraphIdentityUserFlowAttribute {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentityUserFlowAttribute
		return ret
	}
	return o.UserAttribute
}

// GetUserAttributeOk returns a tuple with the UserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttributeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttribute, bool) {
	if o == nil || o.UserAttribute == nil {
		return nil, false
	}
	return &o.UserAttribute, true
}

// HasUserAttribute returns a boolean if a field has been set.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasUserAttribute() bool {
	if o != nil && o.UserAttribute != nil {
		return true
	}

	return false
}

// SetUserAttribute gets a reference to the given AnyOfmicrosoftGraphIdentityUserFlowAttribute and assigns it to the UserAttribute field.
func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserAttribute(v AnyOfmicrosoftGraphIdentityUserFlowAttribute) {
	o.UserAttribute = v
}

func (o MicrosoftGraphIdentityUserFlowAttributeAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsOptional != nil {
		toSerialize["isOptional"] = o.IsOptional
	}
	if o.RequiresVerification != nil {
		toSerialize["requiresVerification"] = o.RequiresVerification
	}
	if o.UserAttributeValues != nil {
		toSerialize["userAttributeValues"] = o.UserAttributeValues
	}
	if o.UserInputType != nil {
		toSerialize["userInputType"] = o.UserInputType
	}
	if o.UserAttribute != nil {
		toSerialize["userAttribute"] = o.UserAttribute
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIdentityUserFlowAttributeAssignment struct {
	value *MicrosoftGraphIdentityUserFlowAttributeAssignment
	isSet bool
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeAssignment) Get() *MicrosoftGraphIdentityUserFlowAttributeAssignment {
	return v.value
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeAssignment) Set(val *MicrosoftGraphIdentityUserFlowAttributeAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIdentityUserFlowAttributeAssignment(val *MicrosoftGraphIdentityUserFlowAttributeAssignment) *NullableMicrosoftGraphIdentityUserFlowAttributeAssignment {
	return &NullableMicrosoftGraphIdentityUserFlowAttributeAssignment{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIdentityUserFlowAttributeAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIdentityUserFlowAttributeAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


