# MicrosoftGraphEntitlementManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AccessPackageAssignmentApprovals** | Pointer to [**[]MicrosoftGraphApproval**](MicrosoftGraphApproval.md) |  | [optional] 
**AccessPackages** | Pointer to [**[]MicrosoftGraphAccessPackage**](MicrosoftGraphAccessPackage.md) | Access packages. | [optional] 
**AssignmentRequests** | Pointer to [**[]MicrosoftGraphAccessPackageAssignmentRequest**](MicrosoftGraphAccessPackageAssignmentRequest.md) | Access package assignment requests. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphAccessPackageAssignment**](MicrosoftGraphAccessPackageAssignment.md) | Access package assignments. | [optional] 
**Catalogs** | Pointer to [**[]MicrosoftGraphAccessPackageCatalog**](MicrosoftGraphAccessPackageCatalog.md) | Access package catalogs. | [optional] 
**ConnectedOrganizations** | Pointer to [**[]MicrosoftGraphConnectedOrganization**](MicrosoftGraphConnectedOrganization.md) | Connected organizations. | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphEntitlementManagementSettings**](anyOf&lt;microsoft.graph.entitlementManagementSettings&gt;.md) | Entitlement management settings. | [optional] 

## Methods

### NewMicrosoftGraphEntitlementManagement

`func NewMicrosoftGraphEntitlementManagement() *MicrosoftGraphEntitlementManagement`

NewMicrosoftGraphEntitlementManagement instantiates a new MicrosoftGraphEntitlementManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEntitlementManagementWithDefaults

`func NewMicrosoftGraphEntitlementManagementWithDefaults() *MicrosoftGraphEntitlementManagement`

NewMicrosoftGraphEntitlementManagementWithDefaults instantiates a new MicrosoftGraphEntitlementManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEntitlementManagement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEntitlementManagement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEntitlementManagement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEntitlementManagement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessPackageAssignmentApprovals

`func (o *MicrosoftGraphEntitlementManagement) GetAccessPackageAssignmentApprovals() []MicrosoftGraphApproval`

GetAccessPackageAssignmentApprovals returns the AccessPackageAssignmentApprovals field if non-nil, zero value otherwise.

### GetAccessPackageAssignmentApprovalsOk

`func (o *MicrosoftGraphEntitlementManagement) GetAccessPackageAssignmentApprovalsOk() (*[]MicrosoftGraphApproval, bool)`

GetAccessPackageAssignmentApprovalsOk returns a tuple with the AccessPackageAssignmentApprovals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackageAssignmentApprovals

`func (o *MicrosoftGraphEntitlementManagement) SetAccessPackageAssignmentApprovals(v []MicrosoftGraphApproval)`

SetAccessPackageAssignmentApprovals sets AccessPackageAssignmentApprovals field to given value.

### HasAccessPackageAssignmentApprovals

`func (o *MicrosoftGraphEntitlementManagement) HasAccessPackageAssignmentApprovals() bool`

HasAccessPackageAssignmentApprovals returns a boolean if a field has been set.

### GetAccessPackages

`func (o *MicrosoftGraphEntitlementManagement) GetAccessPackages() []MicrosoftGraphAccessPackage`

GetAccessPackages returns the AccessPackages field if non-nil, zero value otherwise.

### GetAccessPackagesOk

`func (o *MicrosoftGraphEntitlementManagement) GetAccessPackagesOk() (*[]MicrosoftGraphAccessPackage, bool)`

GetAccessPackagesOk returns a tuple with the AccessPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackages

`func (o *MicrosoftGraphEntitlementManagement) SetAccessPackages(v []MicrosoftGraphAccessPackage)`

SetAccessPackages sets AccessPackages field to given value.

### HasAccessPackages

`func (o *MicrosoftGraphEntitlementManagement) HasAccessPackages() bool`

HasAccessPackages returns a boolean if a field has been set.

### GetAssignmentRequests

`func (o *MicrosoftGraphEntitlementManagement) GetAssignmentRequests() []MicrosoftGraphAccessPackageAssignmentRequest`

GetAssignmentRequests returns the AssignmentRequests field if non-nil, zero value otherwise.

### GetAssignmentRequestsOk

`func (o *MicrosoftGraphEntitlementManagement) GetAssignmentRequestsOk() (*[]MicrosoftGraphAccessPackageAssignmentRequest, bool)`

GetAssignmentRequestsOk returns a tuple with the AssignmentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentRequests

`func (o *MicrosoftGraphEntitlementManagement) SetAssignmentRequests(v []MicrosoftGraphAccessPackageAssignmentRequest)`

SetAssignmentRequests sets AssignmentRequests field to given value.

### HasAssignmentRequests

`func (o *MicrosoftGraphEntitlementManagement) HasAssignmentRequests() bool`

HasAssignmentRequests returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphEntitlementManagement) GetAssignments() []MicrosoftGraphAccessPackageAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphEntitlementManagement) GetAssignmentsOk() (*[]MicrosoftGraphAccessPackageAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphEntitlementManagement) SetAssignments(v []MicrosoftGraphAccessPackageAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphEntitlementManagement) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetCatalogs

`func (o *MicrosoftGraphEntitlementManagement) GetCatalogs() []MicrosoftGraphAccessPackageCatalog`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *MicrosoftGraphEntitlementManagement) GetCatalogsOk() (*[]MicrosoftGraphAccessPackageCatalog, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *MicrosoftGraphEntitlementManagement) SetCatalogs(v []MicrosoftGraphAccessPackageCatalog)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *MicrosoftGraphEntitlementManagement) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetConnectedOrganizations

`func (o *MicrosoftGraphEntitlementManagement) GetConnectedOrganizations() []MicrosoftGraphConnectedOrganization`

GetConnectedOrganizations returns the ConnectedOrganizations field if non-nil, zero value otherwise.

### GetConnectedOrganizationsOk

`func (o *MicrosoftGraphEntitlementManagement) GetConnectedOrganizationsOk() (*[]MicrosoftGraphConnectedOrganization, bool)`

GetConnectedOrganizationsOk returns a tuple with the ConnectedOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedOrganizations

`func (o *MicrosoftGraphEntitlementManagement) SetConnectedOrganizations(v []MicrosoftGraphConnectedOrganization)`

SetConnectedOrganizations sets ConnectedOrganizations field to given value.

### HasConnectedOrganizations

`func (o *MicrosoftGraphEntitlementManagement) HasConnectedOrganizations() bool`

HasConnectedOrganizations returns a boolean if a field has been set.

### GetSettings

`func (o *MicrosoftGraphEntitlementManagement) GetSettings() AnyOfmicrosoftGraphEntitlementManagementSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphEntitlementManagement) GetSettingsOk() (*AnyOfmicrosoftGraphEntitlementManagementSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MicrosoftGraphEntitlementManagement) SetSettings(v AnyOfmicrosoftGraphEntitlementManagementSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MicrosoftGraphEntitlementManagement) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *MicrosoftGraphEntitlementManagement) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *MicrosoftGraphEntitlementManagement) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


