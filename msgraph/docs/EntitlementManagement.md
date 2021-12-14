# EntitlementManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPackageAssignmentApprovals** | Pointer to [**[]MicrosoftGraphApproval**](MicrosoftGraphApproval.md) |  | [optional] 
**AccessPackages** | Pointer to [**[]MicrosoftGraphAccessPackage**](MicrosoftGraphAccessPackage.md) | Access packages. | [optional] 
**AssignmentRequests** | Pointer to [**[]MicrosoftGraphAccessPackageAssignmentRequest**](MicrosoftGraphAccessPackageAssignmentRequest.md) | Access package assignment requests. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphAccessPackageAssignment**](MicrosoftGraphAccessPackageAssignment.md) | Access package assignments. | [optional] 
**Catalogs** | Pointer to [**[]MicrosoftGraphAccessPackageCatalog**](MicrosoftGraphAccessPackageCatalog.md) | Access package catalogs. | [optional] 
**ConnectedOrganizations** | Pointer to [**[]MicrosoftGraphConnectedOrganization**](MicrosoftGraphConnectedOrganization.md) | Connected organizations. | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphEntitlementManagementSettings**](anyOf&lt;microsoft.graph.entitlementManagementSettings&gt;.md) | Entitlement management settings. | [optional] 

## Methods

### NewEntitlementManagement

`func NewEntitlementManagement() *EntitlementManagement`

NewEntitlementManagement instantiates a new EntitlementManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementManagementWithDefaults

`func NewEntitlementManagementWithDefaults() *EntitlementManagement`

NewEntitlementManagementWithDefaults instantiates a new EntitlementManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPackageAssignmentApprovals

`func (o *EntitlementManagement) GetAccessPackageAssignmentApprovals() []MicrosoftGraphApproval`

GetAccessPackageAssignmentApprovals returns the AccessPackageAssignmentApprovals field if non-nil, zero value otherwise.

### GetAccessPackageAssignmentApprovalsOk

`func (o *EntitlementManagement) GetAccessPackageAssignmentApprovalsOk() (*[]MicrosoftGraphApproval, bool)`

GetAccessPackageAssignmentApprovalsOk returns a tuple with the AccessPackageAssignmentApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackageAssignmentApprovals

`func (o *EntitlementManagement) SetAccessPackageAssignmentApprovals(v []MicrosoftGraphApproval)`

SetAccessPackageAssignmentApprovals sets AccessPackageAssignmentApprovals field to given value.

### HasAccessPackageAssignmentApprovals

`func (o *EntitlementManagement) HasAccessPackageAssignmentApprovals() bool`

HasAccessPackageAssignmentApprovals returns a boolean if a field has been set.

### GetAccessPackages

`func (o *EntitlementManagement) GetAccessPackages() []MicrosoftGraphAccessPackage`

GetAccessPackages returns the AccessPackages field if non-nil, zero value otherwise.

### GetAccessPackagesOk

`func (o *EntitlementManagement) GetAccessPackagesOk() (*[]MicrosoftGraphAccessPackage, bool)`

GetAccessPackagesOk returns a tuple with the AccessPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackages

`func (o *EntitlementManagement) SetAccessPackages(v []MicrosoftGraphAccessPackage)`

SetAccessPackages sets AccessPackages field to given value.

### HasAccessPackages

`func (o *EntitlementManagement) HasAccessPackages() bool`

HasAccessPackages returns a boolean if a field has been set.

### GetAssignmentRequests

`func (o *EntitlementManagement) GetAssignmentRequests() []MicrosoftGraphAccessPackageAssignmentRequest`

GetAssignmentRequests returns the AssignmentRequests field if non-nil, zero value otherwise.

### GetAssignmentRequestsOk

`func (o *EntitlementManagement) GetAssignmentRequestsOk() (*[]MicrosoftGraphAccessPackageAssignmentRequest, bool)`

GetAssignmentRequestsOk returns a tuple with the AssignmentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRequests

`func (o *EntitlementManagement) SetAssignmentRequests(v []MicrosoftGraphAccessPackageAssignmentRequest)`

SetAssignmentRequests sets AssignmentRequests field to given value.

### HasAssignmentRequests

`func (o *EntitlementManagement) HasAssignmentRequests() bool`

HasAssignmentRequests returns a boolean if a field has been set.

### GetAssignments

`func (o *EntitlementManagement) GetAssignments() []MicrosoftGraphAccessPackageAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *EntitlementManagement) GetAssignmentsOk() (*[]MicrosoftGraphAccessPackageAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *EntitlementManagement) SetAssignments(v []MicrosoftGraphAccessPackageAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *EntitlementManagement) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetCatalogs

`func (o *EntitlementManagement) GetCatalogs() []MicrosoftGraphAccessPackageCatalog`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *EntitlementManagement) GetCatalogsOk() (*[]MicrosoftGraphAccessPackageCatalog, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *EntitlementManagement) SetCatalogs(v []MicrosoftGraphAccessPackageCatalog)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *EntitlementManagement) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetConnectedOrganizations

`func (o *EntitlementManagement) GetConnectedOrganizations() []MicrosoftGraphConnectedOrganization`

GetConnectedOrganizations returns the ConnectedOrganizations field if non-nil, zero value otherwise.

### GetConnectedOrganizationsOk

`func (o *EntitlementManagement) GetConnectedOrganizationsOk() (*[]MicrosoftGraphConnectedOrganization, bool)`

GetConnectedOrganizationsOk returns a tuple with the ConnectedOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedOrganizations

`func (o *EntitlementManagement) SetConnectedOrganizations(v []MicrosoftGraphConnectedOrganization)`

SetConnectedOrganizations sets ConnectedOrganizations field to given value.

### HasConnectedOrganizations

`func (o *EntitlementManagement) HasConnectedOrganizations() bool`

HasConnectedOrganizations returns a boolean if a field has been set.

### GetSettings

`func (o *EntitlementManagement) GetSettings() AnyOfmicrosoftGraphEntitlementManagementSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *EntitlementManagement) GetSettingsOk() (*AnyOfmicrosoftGraphEntitlementManagementSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *EntitlementManagement) SetSettings(v AnyOfmicrosoftGraphEntitlementManagementSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *EntitlementManagement) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *EntitlementManagement) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *EntitlementManagement) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


