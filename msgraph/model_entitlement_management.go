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

// EntitlementManagement struct for EntitlementManagement
type EntitlementManagement struct {
	AccessPackageAssignmentApprovals *[]MicrosoftGraphApproval `json:"accessPackageAssignmentApprovals,omitempty"`
	// Access packages.
	AccessPackages *[]MicrosoftGraphAccessPackage `json:"accessPackages,omitempty"`
	// Access package assignment requests.
	AssignmentRequests *[]MicrosoftGraphAccessPackageAssignmentRequest `json:"assignmentRequests,omitempty"`
	// Access package assignments.
	Assignments *[]MicrosoftGraphAccessPackageAssignment `json:"assignments,omitempty"`
	// Access package catalogs.
	Catalogs *[]MicrosoftGraphAccessPackageCatalog `json:"catalogs,omitempty"`
	// Connected organizations.
	ConnectedOrganizations *[]MicrosoftGraphConnectedOrganization `json:"connectedOrganizations,omitempty"`
	// Entitlement management settings.
	Settings AnyOfmicrosoftGraphEntitlementManagementSettings `json:"settings,omitempty"`
}

// NewEntitlementManagement instantiates a new EntitlementManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementManagement() *EntitlementManagement {
	this := EntitlementManagement{}
	return &this
}

// NewEntitlementManagementWithDefaults instantiates a new EntitlementManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementManagementWithDefaults() *EntitlementManagement {
	this := EntitlementManagement{}
	return &this
}

// GetAccessPackageAssignmentApprovals returns the AccessPackageAssignmentApprovals field value if set, zero value otherwise.
func (o *EntitlementManagement) GetAccessPackageAssignmentApprovals() []MicrosoftGraphApproval {
	if o == nil || o.AccessPackageAssignmentApprovals == nil {
		var ret []MicrosoftGraphApproval
		return ret
	}
	return *o.AccessPackageAssignmentApprovals
}

// GetAccessPackageAssignmentApprovalsOk returns a tuple with the AccessPackageAssignmentApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementManagement) GetAccessPackageAssignmentApprovalsOk() (*[]MicrosoftGraphApproval, bool) {
	if o == nil || o.AccessPackageAssignmentApprovals == nil {
		return nil, false
	}
	return o.AccessPackageAssignmentApprovals, true
}

// HasAccessPackageAssignmentApprovals returns a boolean if a field has been set.
func (o *EntitlementManagement) HasAccessPackageAssignmentApprovals() bool {
	if o != nil && o.AccessPackageAssignmentApprovals != nil {
		return true
	}

	return false
}

// SetAccessPackageAssignmentApprovals gets a reference to the given []MicrosoftGraphApproval and assigns it to the AccessPackageAssignmentApprovals field.
func (o *EntitlementManagement) SetAccessPackageAssignmentApprovals(v []MicrosoftGraphApproval) {
	o.AccessPackageAssignmentApprovals = &v
}

// GetAccessPackages returns the AccessPackages field value if set, zero value otherwise.
func (o *EntitlementManagement) GetAccessPackages() []MicrosoftGraphAccessPackage {
	if o == nil || o.AccessPackages == nil {
		var ret []MicrosoftGraphAccessPackage
		return ret
	}
	return *o.AccessPackages
}

// GetAccessPackagesOk returns a tuple with the AccessPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementManagement) GetAccessPackagesOk() (*[]MicrosoftGraphAccessPackage, bool) {
	if o == nil || o.AccessPackages == nil {
		return nil, false
	}
	return o.AccessPackages, true
}

// HasAccessPackages returns a boolean if a field has been set.
func (o *EntitlementManagement) HasAccessPackages() bool {
	if o != nil && o.AccessPackages != nil {
		return true
	}

	return false
}

// SetAccessPackages gets a reference to the given []MicrosoftGraphAccessPackage and assigns it to the AccessPackages field.
func (o *EntitlementManagement) SetAccessPackages(v []MicrosoftGraphAccessPackage) {
	o.AccessPackages = &v
}

// GetAssignmentRequests returns the AssignmentRequests field value if set, zero value otherwise.
func (o *EntitlementManagement) GetAssignmentRequests() []MicrosoftGraphAccessPackageAssignmentRequest {
	if o == nil || o.AssignmentRequests == nil {
		var ret []MicrosoftGraphAccessPackageAssignmentRequest
		return ret
	}
	return *o.AssignmentRequests
}

// GetAssignmentRequestsOk returns a tuple with the AssignmentRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementManagement) GetAssignmentRequestsOk() (*[]MicrosoftGraphAccessPackageAssignmentRequest, bool) {
	if o == nil || o.AssignmentRequests == nil {
		return nil, false
	}
	return o.AssignmentRequests, true
}

// HasAssignmentRequests returns a boolean if a field has been set.
func (o *EntitlementManagement) HasAssignmentRequests() bool {
	if o != nil && o.AssignmentRequests != nil {
		return true
	}

	return false
}

// SetAssignmentRequests gets a reference to the given []MicrosoftGraphAccessPackageAssignmentRequest and assigns it to the AssignmentRequests field.
func (o *EntitlementManagement) SetAssignmentRequests(v []MicrosoftGraphAccessPackageAssignmentRequest) {
	o.AssignmentRequests = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *EntitlementManagement) GetAssignments() []MicrosoftGraphAccessPackageAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphAccessPackageAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementManagement) GetAssignmentsOk() (*[]MicrosoftGraphAccessPackageAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *EntitlementManagement) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphAccessPackageAssignment and assigns it to the Assignments field.
func (o *EntitlementManagement) SetAssignments(v []MicrosoftGraphAccessPackageAssignment) {
	o.Assignments = &v
}

// GetCatalogs returns the Catalogs field value if set, zero value otherwise.
func (o *EntitlementManagement) GetCatalogs() []MicrosoftGraphAccessPackageCatalog {
	if o == nil || o.Catalogs == nil {
		var ret []MicrosoftGraphAccessPackageCatalog
		return ret
	}
	return *o.Catalogs
}

// GetCatalogsOk returns a tuple with the Catalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementManagement) GetCatalogsOk() (*[]MicrosoftGraphAccessPackageCatalog, bool) {
	if o == nil || o.Catalogs == nil {
		return nil, false
	}
	return o.Catalogs, true
}

// HasCatalogs returns a boolean if a field has been set.
func (o *EntitlementManagement) HasCatalogs() bool {
	if o != nil && o.Catalogs != nil {
		return true
	}

	return false
}

// SetCatalogs gets a reference to the given []MicrosoftGraphAccessPackageCatalog and assigns it to the Catalogs field.
func (o *EntitlementManagement) SetCatalogs(v []MicrosoftGraphAccessPackageCatalog) {
	o.Catalogs = &v
}

// GetConnectedOrganizations returns the ConnectedOrganizations field value if set, zero value otherwise.
func (o *EntitlementManagement) GetConnectedOrganizations() []MicrosoftGraphConnectedOrganization {
	if o == nil || o.ConnectedOrganizations == nil {
		var ret []MicrosoftGraphConnectedOrganization
		return ret
	}
	return *o.ConnectedOrganizations
}

// GetConnectedOrganizationsOk returns a tuple with the ConnectedOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementManagement) GetConnectedOrganizationsOk() (*[]MicrosoftGraphConnectedOrganization, bool) {
	if o == nil || o.ConnectedOrganizations == nil {
		return nil, false
	}
	return o.ConnectedOrganizations, true
}

// HasConnectedOrganizations returns a boolean if a field has been set.
func (o *EntitlementManagement) HasConnectedOrganizations() bool {
	if o != nil && o.ConnectedOrganizations != nil {
		return true
	}

	return false
}

// SetConnectedOrganizations gets a reference to the given []MicrosoftGraphConnectedOrganization and assigns it to the ConnectedOrganizations field.
func (o *EntitlementManagement) SetConnectedOrganizations(v []MicrosoftGraphConnectedOrganization) {
	o.ConnectedOrganizations = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementManagement) GetSettings() AnyOfmicrosoftGraphEntitlementManagementSettings {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEntitlementManagementSettings
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementManagement) GetSettingsOk() (*AnyOfmicrosoftGraphEntitlementManagementSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *EntitlementManagement) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AnyOfmicrosoftGraphEntitlementManagementSettings and assigns it to the Settings field.
func (o *EntitlementManagement) SetSettings(v AnyOfmicrosoftGraphEntitlementManagementSettings) {
	o.Settings = v
}

func (o EntitlementManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessPackageAssignmentApprovals != nil {
		toSerialize["accessPackageAssignmentApprovals"] = o.AccessPackageAssignmentApprovals
	}
	if o.AccessPackages != nil {
		toSerialize["accessPackages"] = o.AccessPackages
	}
	if o.AssignmentRequests != nil {
		toSerialize["assignmentRequests"] = o.AssignmentRequests
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.Catalogs != nil {
		toSerialize["catalogs"] = o.Catalogs
	}
	if o.ConnectedOrganizations != nil {
		toSerialize["connectedOrganizations"] = o.ConnectedOrganizations
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableEntitlementManagement struct {
	value *EntitlementManagement
	isSet bool
}

func (v NullableEntitlementManagement) Get() *EntitlementManagement {
	return v.value
}

func (v *NullableEntitlementManagement) Set(val *EntitlementManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementManagement(val *EntitlementManagement) *NullableEntitlementManagement {
	return &NullableEntitlementManagement{value: val, isSet: true}
}

func (v NullableEntitlementManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


