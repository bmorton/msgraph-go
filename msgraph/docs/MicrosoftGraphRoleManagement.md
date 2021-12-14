# MicrosoftGraphRoleManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to [**AnyOfmicrosoftGraphRbacApplication**](anyOf&lt;microsoft.graph.rbacApplication&gt;.md) | Read-only. Nullable. | [optional] 
**EntitlementManagement** | Pointer to [**AnyOfmicrosoftGraphRbacApplication**](anyOf&lt;microsoft.graph.rbacApplication&gt;.md) | Container for all entitlement management resources in Azure AD identity governance. | [optional] 

## Methods

### NewMicrosoftGraphRoleManagement

`func NewMicrosoftGraphRoleManagement() *MicrosoftGraphRoleManagement`

NewMicrosoftGraphRoleManagement instantiates a new MicrosoftGraphRoleManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRoleManagementWithDefaults

`func NewMicrosoftGraphRoleManagementWithDefaults() *MicrosoftGraphRoleManagement`

NewMicrosoftGraphRoleManagementWithDefaults instantiates a new MicrosoftGraphRoleManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *MicrosoftGraphRoleManagement) GetDirectory() AnyOfmicrosoftGraphRbacApplication`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *MicrosoftGraphRoleManagement) GetDirectoryOk() (*AnyOfmicrosoftGraphRbacApplication, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *MicrosoftGraphRoleManagement) SetDirectory(v AnyOfmicrosoftGraphRbacApplication)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *MicrosoftGraphRoleManagement) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### SetDirectoryNil

`func (o *MicrosoftGraphRoleManagement) SetDirectoryNil(b bool)`

 SetDirectoryNil sets the value for Directory to be an explicit nil

### UnsetDirectory
`func (o *MicrosoftGraphRoleManagement) UnsetDirectory()`

UnsetDirectory ensures that no value is present for Directory, not even an explicit nil
### GetEntitlementManagement

`func (o *MicrosoftGraphRoleManagement) GetEntitlementManagement() AnyOfmicrosoftGraphRbacApplication`

GetEntitlementManagement returns the EntitlementManagement field if non-nil, zero value otherwise.

### GetEntitlementManagementOk

`func (o *MicrosoftGraphRoleManagement) GetEntitlementManagementOk() (*AnyOfmicrosoftGraphRbacApplication, bool)`

GetEntitlementManagementOk returns a tuple with the EntitlementManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementManagement

`func (o *MicrosoftGraphRoleManagement) SetEntitlementManagement(v AnyOfmicrosoftGraphRbacApplication)`

SetEntitlementManagement sets EntitlementManagement field to given value.

### HasEntitlementManagement

`func (o *MicrosoftGraphRoleManagement) HasEntitlementManagement() bool`

HasEntitlementManagement returns a boolean if a field has been set.

### SetEntitlementManagementNil

`func (o *MicrosoftGraphRoleManagement) SetEntitlementManagementNil(b bool)`

 SetEntitlementManagementNil sets the value for EntitlementManagement to be an explicit nil

### UnsetEntitlementManagement
`func (o *MicrosoftGraphRoleManagement) UnsetEntitlementManagement()`

UnsetEntitlementManagement ensures that no value is present for EntitlementManagement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


