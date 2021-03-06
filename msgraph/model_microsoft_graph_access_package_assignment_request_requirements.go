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

// MicrosoftGraphAccessPackageAssignmentRequestRequirements struct for MicrosoftGraphAccessPackageAssignmentRequestRequirements
type MicrosoftGraphAccessPackageAssignmentRequestRequirements struct {
	// Indicates whether the requestor is allowed to set a custom schedule.
	AllowCustomAssignmentSchedule NullableBool `json:"allowCustomAssignmentSchedule,omitempty"`
	// Indicates whether a request to add must be approved by an approver.
	IsApprovalRequiredForAdd NullableBool `json:"isApprovalRequiredForAdd,omitempty"`
	// Indicates whether a request to update must be approved by an approver.
	IsApprovalRequiredForUpdate NullableBool `json:"isApprovalRequiredForUpdate,omitempty"`
	// The description of the policy that the user is trying to request access using.
	PolicyDescription NullableString `json:"policyDescription,omitempty"`
	// The display name of the policy that the user is trying to request access using.
	PolicyDisplayName NullableString `json:"policyDisplayName,omitempty"`
	// The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
	PolicyId NullableString `json:"policyId,omitempty"`
	// Schedule restrictions enforced, if any.
	Schedule AnyOfmicrosoftGraphEntitlementManagementSchedule `json:"schedule,omitempty"`
}

// NewMicrosoftGraphAccessPackageAssignmentRequestRequirements instantiates a new MicrosoftGraphAccessPackageAssignmentRequestRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAccessPackageAssignmentRequestRequirements() *MicrosoftGraphAccessPackageAssignmentRequestRequirements {
	this := MicrosoftGraphAccessPackageAssignmentRequestRequirements{}
	return &this
}

// NewMicrosoftGraphAccessPackageAssignmentRequestRequirementsWithDefaults instantiates a new MicrosoftGraphAccessPackageAssignmentRequestRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAccessPackageAssignmentRequestRequirementsWithDefaults() *MicrosoftGraphAccessPackageAssignmentRequestRequirements {
	this := MicrosoftGraphAccessPackageAssignmentRequestRequirements{}
	return &this
}

// GetAllowCustomAssignmentSchedule returns the AllowCustomAssignmentSchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetAllowCustomAssignmentSchedule() bool {
	if o == nil || o.AllowCustomAssignmentSchedule.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomAssignmentSchedule.Get()
}

// GetAllowCustomAssignmentScheduleOk returns a tuple with the AllowCustomAssignmentSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetAllowCustomAssignmentScheduleOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowCustomAssignmentSchedule.Get(), o.AllowCustomAssignmentSchedule.IsSet()
}

// HasAllowCustomAssignmentSchedule returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasAllowCustomAssignmentSchedule() bool {
	if o != nil && o.AllowCustomAssignmentSchedule.IsSet() {
		return true
	}

	return false
}

// SetAllowCustomAssignmentSchedule gets a reference to the given NullableBool and assigns it to the AllowCustomAssignmentSchedule field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetAllowCustomAssignmentSchedule(v bool) {
	o.AllowCustomAssignmentSchedule.Set(&v)
}
// SetAllowCustomAssignmentScheduleNil sets the value for AllowCustomAssignmentSchedule to be an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetAllowCustomAssignmentScheduleNil() {
	o.AllowCustomAssignmentSchedule.Set(nil)
}

// UnsetAllowCustomAssignmentSchedule ensures that no value is present for AllowCustomAssignmentSchedule, not even an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetAllowCustomAssignmentSchedule() {
	o.AllowCustomAssignmentSchedule.Unset()
}

// GetIsApprovalRequiredForAdd returns the IsApprovalRequiredForAdd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForAdd() bool {
	if o == nil || o.IsApprovalRequiredForAdd.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsApprovalRequiredForAdd.Get()
}

// GetIsApprovalRequiredForAddOk returns a tuple with the IsApprovalRequiredForAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForAddOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsApprovalRequiredForAdd.Get(), o.IsApprovalRequiredForAdd.IsSet()
}

// HasIsApprovalRequiredForAdd returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasIsApprovalRequiredForAdd() bool {
	if o != nil && o.IsApprovalRequiredForAdd.IsSet() {
		return true
	}

	return false
}

// SetIsApprovalRequiredForAdd gets a reference to the given NullableBool and assigns it to the IsApprovalRequiredForAdd field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForAdd(v bool) {
	o.IsApprovalRequiredForAdd.Set(&v)
}
// SetIsApprovalRequiredForAddNil sets the value for IsApprovalRequiredForAdd to be an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForAddNil() {
	o.IsApprovalRequiredForAdd.Set(nil)
}

// UnsetIsApprovalRequiredForAdd ensures that no value is present for IsApprovalRequiredForAdd, not even an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetIsApprovalRequiredForAdd() {
	o.IsApprovalRequiredForAdd.Unset()
}

// GetIsApprovalRequiredForUpdate returns the IsApprovalRequiredForUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForUpdate() bool {
	if o == nil || o.IsApprovalRequiredForUpdate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsApprovalRequiredForUpdate.Get()
}

// GetIsApprovalRequiredForUpdateOk returns a tuple with the IsApprovalRequiredForUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForUpdateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsApprovalRequiredForUpdate.Get(), o.IsApprovalRequiredForUpdate.IsSet()
}

// HasIsApprovalRequiredForUpdate returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasIsApprovalRequiredForUpdate() bool {
	if o != nil && o.IsApprovalRequiredForUpdate.IsSet() {
		return true
	}

	return false
}

// SetIsApprovalRequiredForUpdate gets a reference to the given NullableBool and assigns it to the IsApprovalRequiredForUpdate field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForUpdate(v bool) {
	o.IsApprovalRequiredForUpdate.Set(&v)
}
// SetIsApprovalRequiredForUpdateNil sets the value for IsApprovalRequiredForUpdate to be an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForUpdateNil() {
	o.IsApprovalRequiredForUpdate.Set(nil)
}

// UnsetIsApprovalRequiredForUpdate ensures that no value is present for IsApprovalRequiredForUpdate, not even an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetIsApprovalRequiredForUpdate() {
	o.IsApprovalRequiredForUpdate.Unset()
}

// GetPolicyDescription returns the PolicyDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDescription() string {
	if o == nil || o.PolicyDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyDescription.Get()
}

// GetPolicyDescriptionOk returns a tuple with the PolicyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyDescription.Get(), o.PolicyDescription.IsSet()
}

// HasPolicyDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasPolicyDescription() bool {
	if o != nil && o.PolicyDescription.IsSet() {
		return true
	}

	return false
}

// SetPolicyDescription gets a reference to the given NullableString and assigns it to the PolicyDescription field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDescription(v string) {
	o.PolicyDescription.Set(&v)
}
// SetPolicyDescriptionNil sets the value for PolicyDescription to be an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDescriptionNil() {
	o.PolicyDescription.Set(nil)
}

// UnsetPolicyDescription ensures that no value is present for PolicyDescription, not even an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetPolicyDescription() {
	o.PolicyDescription.Unset()
}

// GetPolicyDisplayName returns the PolicyDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDisplayName() string {
	if o == nil || o.PolicyDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyDisplayName.Get()
}

// GetPolicyDisplayNameOk returns a tuple with the PolicyDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyDisplayName.Get(), o.PolicyDisplayName.IsSet()
}

// HasPolicyDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasPolicyDisplayName() bool {
	if o != nil && o.PolicyDisplayName.IsSet() {
		return true
	}

	return false
}

// SetPolicyDisplayName gets a reference to the given NullableString and assigns it to the PolicyDisplayName field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDisplayName(v string) {
	o.PolicyDisplayName.Set(&v)
}
// SetPolicyDisplayNameNil sets the value for PolicyDisplayName to be an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDisplayNameNil() {
	o.PolicyDisplayName.Set(nil)
}

// UnsetPolicyDisplayName ensures that no value is present for PolicyDisplayName, not even an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetPolicyDisplayName() {
	o.PolicyDisplayName.Unset()
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// HasPolicyId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasPolicyId() bool {
	if o != nil && o.PolicyId.IsSet() {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given NullableString and assigns it to the PolicyId field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}
// SetPolicyIdNil sets the value for PolicyId to be an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyIdNil() {
	o.PolicyId.Set(nil)
}

// UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetPolicyId() {
	o.PolicyId.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetSchedule() AnyOfmicrosoftGraphEntitlementManagementSchedule {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEntitlementManagementSchedule
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetScheduleOk() (*AnyOfmicrosoftGraphEntitlementManagementSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given AnyOfmicrosoftGraphEntitlementManagementSchedule and assigns it to the Schedule field.
func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetSchedule(v AnyOfmicrosoftGraphEntitlementManagementSchedule) {
	o.Schedule = v
}

func (o MicrosoftGraphAccessPackageAssignmentRequestRequirements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowCustomAssignmentSchedule.IsSet() {
		toSerialize["allowCustomAssignmentSchedule"] = o.AllowCustomAssignmentSchedule.Get()
	}
	if o.IsApprovalRequiredForAdd.IsSet() {
		toSerialize["isApprovalRequiredForAdd"] = o.IsApprovalRequiredForAdd.Get()
	}
	if o.IsApprovalRequiredForUpdate.IsSet() {
		toSerialize["isApprovalRequiredForUpdate"] = o.IsApprovalRequiredForUpdate.Get()
	}
	if o.PolicyDescription.IsSet() {
		toSerialize["policyDescription"] = o.PolicyDescription.Get()
	}
	if o.PolicyDisplayName.IsSet() {
		toSerialize["policyDisplayName"] = o.PolicyDisplayName.Get()
	}
	if o.PolicyId.IsSet() {
		toSerialize["policyId"] = o.PolicyId.Get()
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements struct {
	value *MicrosoftGraphAccessPackageAssignmentRequestRequirements
	isSet bool
}

func (v NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements) Get() *MicrosoftGraphAccessPackageAssignmentRequestRequirements {
	return v.value
}

func (v *NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements) Set(val *MicrosoftGraphAccessPackageAssignmentRequestRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessPackageAssignmentRequestRequirements(val *MicrosoftGraphAccessPackageAssignmentRequestRequirements) *NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements {
	return &NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessPackageAssignmentRequestRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


