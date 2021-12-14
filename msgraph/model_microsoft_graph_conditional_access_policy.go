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

// MicrosoftGraphConditionalAccessPolicy struct for MicrosoftGraphConditionalAccessPolicy
type MicrosoftGraphConditionalAccessPolicy struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	Conditions *MicrosoftGraphConditionalAccessConditionSet `json:"conditions,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Not used.
	Description NullableString `json:"description,omitempty"`
	// Specifies a display name for the conditionalAccessPolicy object.
	DisplayName *string `json:"displayName,omitempty"`
	// Specifies the grant controls that must be fulfilled to pass the policy.
	GrantControls AnyOfmicrosoftGraphConditionalAccessGrantControls `json:"grantControls,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
	ModifiedDateTime NullableTime `json:"modifiedDateTime,omitempty"`
	// Specifies the session controls that are enforced after sign-in.
	SessionControls AnyOfmicrosoftGraphConditionalAccessSessionControls `json:"sessionControls,omitempty"`
	// Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced. Required.
	State AnyOfmicrosoftGraphConditionalAccessPolicyState `json:"state,omitempty"`
}

// NewMicrosoftGraphConditionalAccessPolicy instantiates a new MicrosoftGraphConditionalAccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphConditionalAccessPolicy() *MicrosoftGraphConditionalAccessPolicy {
	this := MicrosoftGraphConditionalAccessPolicy{}
	return &this
}

// NewMicrosoftGraphConditionalAccessPolicyWithDefaults instantiates a new MicrosoftGraphConditionalAccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphConditionalAccessPolicyWithDefaults() *MicrosoftGraphConditionalAccessPolicy {
	this := MicrosoftGraphConditionalAccessPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetId(v string) {
	o.Id = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessPolicy) GetConditions() MicrosoftGraphConditionalAccessConditionSet {
	if o == nil || o.Conditions == nil {
		var ret MicrosoftGraphConditionalAccessConditionSet
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) GetConditionsOk() (*MicrosoftGraphConditionalAccessConditionSet, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given MicrosoftGraphConditionalAccessConditionSet and assigns it to the Conditions field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetConditions(v MicrosoftGraphConditionalAccessConditionSet) {
	o.Conditions = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessPolicy) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessPolicy) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphConditionalAccessPolicy) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphConditionalAccessPolicy) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessPolicy) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphConditionalAccessPolicy) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphConditionalAccessPolicy) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphConditionalAccessPolicy) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGrantControls returns the GrantControls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessPolicy) GetGrantControls() AnyOfmicrosoftGraphConditionalAccessGrantControls {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessGrantControls
		return ret
	}
	return o.GrantControls
}

// GetGrantControlsOk returns a tuple with the GrantControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessPolicy) GetGrantControlsOk() (*AnyOfmicrosoftGraphConditionalAccessGrantControls, bool) {
	if o == nil || o.GrantControls == nil {
		return nil, false
	}
	return &o.GrantControls, true
}

// HasGrantControls returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasGrantControls() bool {
	if o != nil && o.GrantControls != nil {
		return true
	}

	return false
}

// SetGrantControls gets a reference to the given AnyOfmicrosoftGraphConditionalAccessGrantControls and assigns it to the GrantControls field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetGrantControls(v AnyOfmicrosoftGraphConditionalAccessGrantControls) {
	o.GrantControls = v
}

// GetModifiedDateTime returns the ModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessPolicy) GetModifiedDateTime() time.Time {
	if o == nil || o.ModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDateTime.Get()
}

// GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessPolicy) GetModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedDateTime.Get(), o.ModifiedDateTime.IsSet()
}

// HasModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasModifiedDateTime() bool {
	if o != nil && o.ModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetModifiedDateTime gets a reference to the given NullableTime and assigns it to the ModifiedDateTime field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetModifiedDateTime(v time.Time) {
	o.ModifiedDateTime.Set(&v)
}
// SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphConditionalAccessPolicy) SetModifiedDateTimeNil() {
	o.ModifiedDateTime.Set(nil)
}

// UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphConditionalAccessPolicy) UnsetModifiedDateTime() {
	o.ModifiedDateTime.Unset()
}

// GetSessionControls returns the SessionControls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessPolicy) GetSessionControls() AnyOfmicrosoftGraphConditionalAccessSessionControls {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessSessionControls
		return ret
	}
	return o.SessionControls
}

// GetSessionControlsOk returns a tuple with the SessionControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessPolicy) GetSessionControlsOk() (*AnyOfmicrosoftGraphConditionalAccessSessionControls, bool) {
	if o == nil || o.SessionControls == nil {
		return nil, false
	}
	return &o.SessionControls, true
}

// HasSessionControls returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasSessionControls() bool {
	if o != nil && o.SessionControls != nil {
		return true
	}

	return false
}

// SetSessionControls gets a reference to the given AnyOfmicrosoftGraphConditionalAccessSessionControls and assigns it to the SessionControls field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetSessionControls(v AnyOfmicrosoftGraphConditionalAccessSessionControls) {
	o.SessionControls = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConditionalAccessPolicy) GetState() AnyOfmicrosoftGraphConditionalAccessPolicyState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConditionalAccessPolicyState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConditionalAccessPolicy) GetStateOk() (*AnyOfmicrosoftGraphConditionalAccessPolicyState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphConditionalAccessPolicy) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphConditionalAccessPolicyState and assigns it to the State field.
func (o *MicrosoftGraphConditionalAccessPolicy) SetState(v AnyOfmicrosoftGraphConditionalAccessPolicyState) {
	o.State = v
}

func (o MicrosoftGraphConditionalAccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.GrantControls != nil {
		toSerialize["grantControls"] = o.GrantControls
	}
	if o.ModifiedDateTime.IsSet() {
		toSerialize["modifiedDateTime"] = o.ModifiedDateTime.Get()
	}
	if o.SessionControls != nil {
		toSerialize["sessionControls"] = o.SessionControls
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphConditionalAccessPolicy struct {
	value *MicrosoftGraphConditionalAccessPolicy
	isSet bool
}

func (v NullableMicrosoftGraphConditionalAccessPolicy) Get() *MicrosoftGraphConditionalAccessPolicy {
	return v.value
}

func (v *NullableMicrosoftGraphConditionalAccessPolicy) Set(val *MicrosoftGraphConditionalAccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConditionalAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConditionalAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConditionalAccessPolicy(val *MicrosoftGraphConditionalAccessPolicy) *NullableMicrosoftGraphConditionalAccessPolicy {
	return &NullableMicrosoftGraphConditionalAccessPolicy{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConditionalAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConditionalAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


