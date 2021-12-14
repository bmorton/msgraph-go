/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphStsPolicy struct for MicrosoftGraphStsPolicy
type MicrosoftGraphStsPolicy struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime NullableTime `json:"deletedDateTime,omitempty"`
	// Description for this policy. Required.
	Description NullableString `json:"description,omitempty"`
	// Display name for this policy. Required.
	DisplayName NullableString `json:"displayName,omitempty"`
	// A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
	Definition *[]string `json:"definition,omitempty"`
	// If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
	IsOrganizationDefault NullableBool `json:"isOrganizationDefault,omitempty"`
	AppliesTo *[]MicrosoftGraphDirectoryObject `json:"appliesTo,omitempty"`
}

// NewMicrosoftGraphStsPolicy instantiates a new MicrosoftGraphStsPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphStsPolicy() *MicrosoftGraphStsPolicy {
	this := MicrosoftGraphStsPolicy{}
	return &this
}

// NewMicrosoftGraphStsPolicyWithDefaults instantiates a new MicrosoftGraphStsPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphStsPolicyWithDefaults() *MicrosoftGraphStsPolicy {
	this := MicrosoftGraphStsPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphStsPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphStsPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphStsPolicy) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphStsPolicy) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime.Get()
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphStsPolicy) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedDateTime.Get(), o.DeletedDateTime.IsSet()
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given NullableTime and assigns it to the DeletedDateTime field.
func (o *MicrosoftGraphStsPolicy) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime.Set(&v)
}
// SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil
func (o *MicrosoftGraphStsPolicy) SetDeletedDateTimeNil() {
	o.DeletedDateTime.Set(nil)
}

// UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
func (o *MicrosoftGraphStsPolicy) UnsetDeletedDateTime() {
	o.DeletedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphStsPolicy) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphStsPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphStsPolicy) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphStsPolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphStsPolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphStsPolicy) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphStsPolicy) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphStsPolicy) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphStsPolicy) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphStsPolicy) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *MicrosoftGraphStsPolicy) GetDefinition() []string {
	if o == nil || o.Definition == nil {
		var ret []string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphStsPolicy) GetDefinitionOk() (*[]string, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given []string and assigns it to the Definition field.
func (o *MicrosoftGraphStsPolicy) SetDefinition(v []string) {
	o.Definition = &v
}

// GetIsOrganizationDefault returns the IsOrganizationDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphStsPolicy) GetIsOrganizationDefault() bool {
	if o == nil || o.IsOrganizationDefault.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsOrganizationDefault.Get()
}

// GetIsOrganizationDefaultOk returns a tuple with the IsOrganizationDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphStsPolicy) GetIsOrganizationDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsOrganizationDefault.Get(), o.IsOrganizationDefault.IsSet()
}

// HasIsOrganizationDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasIsOrganizationDefault() bool {
	if o != nil && o.IsOrganizationDefault.IsSet() {
		return true
	}

	return false
}

// SetIsOrganizationDefault gets a reference to the given NullableBool and assigns it to the IsOrganizationDefault field.
func (o *MicrosoftGraphStsPolicy) SetIsOrganizationDefault(v bool) {
	o.IsOrganizationDefault.Set(&v)
}
// SetIsOrganizationDefaultNil sets the value for IsOrganizationDefault to be an explicit nil
func (o *MicrosoftGraphStsPolicy) SetIsOrganizationDefaultNil() {
	o.IsOrganizationDefault.Set(nil)
}

// UnsetIsOrganizationDefault ensures that no value is present for IsOrganizationDefault, not even an explicit nil
func (o *MicrosoftGraphStsPolicy) UnsetIsOrganizationDefault() {
	o.IsOrganizationDefault.Unset()
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *MicrosoftGraphStsPolicy) GetAppliesTo() []MicrosoftGraphDirectoryObject {
	if o == nil || o.AppliesTo == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphStsPolicy) GetAppliesToOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.AppliesTo == nil {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *MicrosoftGraphStsPolicy) HasAppliesTo() bool {
	if o != nil && o.AppliesTo != nil {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the AppliesTo field.
func (o *MicrosoftGraphStsPolicy) SetAppliesTo(v []MicrosoftGraphDirectoryObject) {
	o.AppliesTo = &v
}

func (o MicrosoftGraphStsPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime.IsSet() {
		toSerialize["deletedDateTime"] = o.DeletedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	if o.IsOrganizationDefault.IsSet() {
		toSerialize["isOrganizationDefault"] = o.IsOrganizationDefault.Get()
	}
	if o.AppliesTo != nil {
		toSerialize["appliesTo"] = o.AppliesTo
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphStsPolicy struct {
	value *MicrosoftGraphStsPolicy
	isSet bool
}

func (v NullableMicrosoftGraphStsPolicy) Get() *MicrosoftGraphStsPolicy {
	return v.value
}

func (v *NullableMicrosoftGraphStsPolicy) Set(val *MicrosoftGraphStsPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphStsPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphStsPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphStsPolicy(val *MicrosoftGraphStsPolicy) *NullableMicrosoftGraphStsPolicy {
	return &NullableMicrosoftGraphStsPolicy{value: val, isSet: true}
}

func (v NullableMicrosoftGraphStsPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphStsPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


