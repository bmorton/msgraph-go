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

// TargetedManagedAppConfiguration Configuration used to deliver a set of custom settings as-is to all users in the targeted security group
type TargetedManagedAppConfiguration struct {
	// Count of apps to which the current policy is deployed.
	DeployedAppCount *int32 `json:"deployedAppCount,omitempty"`
	// Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// List of apps to which the policy is deployed.
	Apps *[]MicrosoftGraphManagedMobileApp `json:"apps,omitempty"`
	// Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
	Assignments *[]MicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
	// Navigation property to deployment summary of the configuration.
	DeploymentSummary AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`
}

// NewTargetedManagedAppConfiguration instantiates a new TargetedManagedAppConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetedManagedAppConfiguration() *TargetedManagedAppConfiguration {
	this := TargetedManagedAppConfiguration{}
	return &this
}

// NewTargetedManagedAppConfigurationWithDefaults instantiates a new TargetedManagedAppConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetedManagedAppConfigurationWithDefaults() *TargetedManagedAppConfiguration {
	this := TargetedManagedAppConfiguration{}
	return &this
}

// GetDeployedAppCount returns the DeployedAppCount field value if set, zero value otherwise.
func (o *TargetedManagedAppConfiguration) GetDeployedAppCount() int32 {
	if o == nil || o.DeployedAppCount == nil {
		var ret int32
		return ret
	}
	return *o.DeployedAppCount
}

// GetDeployedAppCountOk returns a tuple with the DeployedAppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetedManagedAppConfiguration) GetDeployedAppCountOk() (*int32, bool) {
	if o == nil || o.DeployedAppCount == nil {
		return nil, false
	}
	return o.DeployedAppCount, true
}

// HasDeployedAppCount returns a boolean if a field has been set.
func (o *TargetedManagedAppConfiguration) HasDeployedAppCount() bool {
	if o != nil && o.DeployedAppCount != nil {
		return true
	}

	return false
}

// SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.
func (o *TargetedManagedAppConfiguration) SetDeployedAppCount(v int32) {
	o.DeployedAppCount = &v
}

// GetIsAssigned returns the IsAssigned field value if set, zero value otherwise.
func (o *TargetedManagedAppConfiguration) GetIsAssigned() bool {
	if o == nil || o.IsAssigned == nil {
		var ret bool
		return ret
	}
	return *o.IsAssigned
}

// GetIsAssignedOk returns a tuple with the IsAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetedManagedAppConfiguration) GetIsAssignedOk() (*bool, bool) {
	if o == nil || o.IsAssigned == nil {
		return nil, false
	}
	return o.IsAssigned, true
}

// HasIsAssigned returns a boolean if a field has been set.
func (o *TargetedManagedAppConfiguration) HasIsAssigned() bool {
	if o != nil && o.IsAssigned != nil {
		return true
	}

	return false
}

// SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.
func (o *TargetedManagedAppConfiguration) SetIsAssigned(v bool) {
	o.IsAssigned = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *TargetedManagedAppConfiguration) GetApps() []MicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []MicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetedManagedAppConfiguration) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *TargetedManagedAppConfiguration) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *TargetedManagedAppConfiguration) SetApps(v []MicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *TargetedManagedAppConfiguration) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetedManagedAppConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *TargetedManagedAppConfiguration) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *TargetedManagedAppConfiguration) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}

// GetDeploymentSummary returns the DeploymentSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetedManagedAppConfiguration) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary {
	if o == nil  {
		var ret AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary
		return ret
	}
	return o.DeploymentSummary
}

// GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetedManagedAppConfiguration) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool) {
	if o == nil || o.DeploymentSummary == nil {
		return nil, false
	}
	return &o.DeploymentSummary, true
}

// HasDeploymentSummary returns a boolean if a field has been set.
func (o *TargetedManagedAppConfiguration) HasDeploymentSummary() bool {
	if o != nil && o.DeploymentSummary != nil {
		return true
	}

	return false
}

// SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.
func (o *TargetedManagedAppConfiguration) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary) {
	o.DeploymentSummary = v
}

func (o TargetedManagedAppConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeployedAppCount != nil {
		toSerialize["deployedAppCount"] = o.DeployedAppCount
	}
	if o.IsAssigned != nil {
		toSerialize["isAssigned"] = o.IsAssigned
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.DeploymentSummary != nil {
		toSerialize["deploymentSummary"] = o.DeploymentSummary
	}
	return json.Marshal(toSerialize)
}

type NullableTargetedManagedAppConfiguration struct {
	value *TargetedManagedAppConfiguration
	isSet bool
}

func (v NullableTargetedManagedAppConfiguration) Get() *TargetedManagedAppConfiguration {
	return v.value
}

func (v *NullableTargetedManagedAppConfiguration) Set(val *TargetedManagedAppConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetedManagedAppConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetedManagedAppConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetedManagedAppConfiguration(val *TargetedManagedAppConfiguration) *NullableTargetedManagedAppConfiguration {
	return &NullableTargetedManagedAppConfiguration{value: val, isSet: true}
}

func (v NullableTargetedManagedAppConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetedManagedAppConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


