# PermissionGrantConditionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientApplicationIds** | Pointer to **[]string** | A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all. | [optional] 
**ClientApplicationPublisherIds** | Pointer to **[]string** | A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all. | [optional] 
**ClientApplicationsFromVerifiedPublisherOnly** | Pointer to **NullableBool** | Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false. | [optional] 
**ClientApplicationTenantIds** | Pointer to **[]string** | A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all. | [optional] 
**PermissionClassification** | Pointer to **NullableString** | The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all. | [optional] 
**Permissions** | Pointer to **[]string** | The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API&#39;s **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API&#39;s **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API&#39;s **servicePrincipal** object. Default is the single value all. | [optional] 
**PermissionType** | Pointer to [**AnyOfmicrosoftGraphPermissionType**](anyOf&lt;microsoft.graph.permissionType&gt;.md) | The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required. | [optional] 
**ResourceApplication** | Pointer to **NullableString** | The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any. | [optional] 

## Methods

### NewPermissionGrantConditionSet

`func NewPermissionGrantConditionSet() *PermissionGrantConditionSet`

NewPermissionGrantConditionSet instantiates a new PermissionGrantConditionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantConditionSetWithDefaults

`func NewPermissionGrantConditionSetWithDefaults() *PermissionGrantConditionSet`

NewPermissionGrantConditionSetWithDefaults instantiates a new PermissionGrantConditionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientApplicationIds

`func (o *PermissionGrantConditionSet) GetClientApplicationIds() []*string`

GetClientApplicationIds returns the ClientApplicationIds field if non-nil, zero value otherwise.

### GetClientApplicationIdsOk

`func (o *PermissionGrantConditionSet) GetClientApplicationIdsOk() (*[]*string, bool)`

GetClientApplicationIdsOk returns a tuple with the ClientApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationIds

`func (o *PermissionGrantConditionSet) SetClientApplicationIds(v []*string)`

SetClientApplicationIds sets ClientApplicationIds field to given value.

### HasClientApplicationIds

`func (o *PermissionGrantConditionSet) HasClientApplicationIds() bool`

HasClientApplicationIds returns a boolean if a field has been set.

### GetClientApplicationPublisherIds

`func (o *PermissionGrantConditionSet) GetClientApplicationPublisherIds() []*string`

GetClientApplicationPublisherIds returns the ClientApplicationPublisherIds field if non-nil, zero value otherwise.

### GetClientApplicationPublisherIdsOk

`func (o *PermissionGrantConditionSet) GetClientApplicationPublisherIdsOk() (*[]*string, bool)`

GetClientApplicationPublisherIdsOk returns a tuple with the ClientApplicationPublisherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationPublisherIds

`func (o *PermissionGrantConditionSet) SetClientApplicationPublisherIds(v []*string)`

SetClientApplicationPublisherIds sets ClientApplicationPublisherIds field to given value.

### HasClientApplicationPublisherIds

`func (o *PermissionGrantConditionSet) HasClientApplicationPublisherIds() bool`

HasClientApplicationPublisherIds returns a boolean if a field has been set.

### GetClientApplicationsFromVerifiedPublisherOnly

`func (o *PermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnly() bool`

GetClientApplicationsFromVerifiedPublisherOnly returns the ClientApplicationsFromVerifiedPublisherOnly field if non-nil, zero value otherwise.

### GetClientApplicationsFromVerifiedPublisherOnlyOk

`func (o *PermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnlyOk() (*bool, bool)`

GetClientApplicationsFromVerifiedPublisherOnlyOk returns a tuple with the ClientApplicationsFromVerifiedPublisherOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationsFromVerifiedPublisherOnly

`func (o *PermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnly(v bool)`

SetClientApplicationsFromVerifiedPublisherOnly sets ClientApplicationsFromVerifiedPublisherOnly field to given value.

### HasClientApplicationsFromVerifiedPublisherOnly

`func (o *PermissionGrantConditionSet) HasClientApplicationsFromVerifiedPublisherOnly() bool`

HasClientApplicationsFromVerifiedPublisherOnly returns a boolean if a field has been set.

### SetClientApplicationsFromVerifiedPublisherOnlyNil

`func (o *PermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnlyNil(b bool)`

 SetClientApplicationsFromVerifiedPublisherOnlyNil sets the value for ClientApplicationsFromVerifiedPublisherOnly to be an explicit nil

### UnsetClientApplicationsFromVerifiedPublisherOnly
`func (o *PermissionGrantConditionSet) UnsetClientApplicationsFromVerifiedPublisherOnly()`

UnsetClientApplicationsFromVerifiedPublisherOnly ensures that no value is present for ClientApplicationsFromVerifiedPublisherOnly, not even an explicit nil
### GetClientApplicationTenantIds

`func (o *PermissionGrantConditionSet) GetClientApplicationTenantIds() []*string`

GetClientApplicationTenantIds returns the ClientApplicationTenantIds field if non-nil, zero value otherwise.

### GetClientApplicationTenantIdsOk

`func (o *PermissionGrantConditionSet) GetClientApplicationTenantIdsOk() (*[]*string, bool)`

GetClientApplicationTenantIdsOk returns a tuple with the ClientApplicationTenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationTenantIds

`func (o *PermissionGrantConditionSet) SetClientApplicationTenantIds(v []*string)`

SetClientApplicationTenantIds sets ClientApplicationTenantIds field to given value.

### HasClientApplicationTenantIds

`func (o *PermissionGrantConditionSet) HasClientApplicationTenantIds() bool`

HasClientApplicationTenantIds returns a boolean if a field has been set.

### GetPermissionClassification

`func (o *PermissionGrantConditionSet) GetPermissionClassification() string`

GetPermissionClassification returns the PermissionClassification field if non-nil, zero value otherwise.

### GetPermissionClassificationOk

`func (o *PermissionGrantConditionSet) GetPermissionClassificationOk() (*string, bool)`

GetPermissionClassificationOk returns a tuple with the PermissionClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionClassification

`func (o *PermissionGrantConditionSet) SetPermissionClassification(v string)`

SetPermissionClassification sets PermissionClassification field to given value.

### HasPermissionClassification

`func (o *PermissionGrantConditionSet) HasPermissionClassification() bool`

HasPermissionClassification returns a boolean if a field has been set.

### SetPermissionClassificationNil

`func (o *PermissionGrantConditionSet) SetPermissionClassificationNil(b bool)`

 SetPermissionClassificationNil sets the value for PermissionClassification to be an explicit nil

### UnsetPermissionClassification
`func (o *PermissionGrantConditionSet) UnsetPermissionClassification()`

UnsetPermissionClassification ensures that no value is present for PermissionClassification, not even an explicit nil
### GetPermissions

`func (o *PermissionGrantConditionSet) GetPermissions() []*string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionGrantConditionSet) GetPermissionsOk() (*[]*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionGrantConditionSet) SetPermissions(v []*string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PermissionGrantConditionSet) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPermissionType

`func (o *PermissionGrantConditionSet) GetPermissionType() AnyOfmicrosoftGraphPermissionType`

GetPermissionType returns the PermissionType field if non-nil, zero value otherwise.

### GetPermissionTypeOk

`func (o *PermissionGrantConditionSet) GetPermissionTypeOk() (*AnyOfmicrosoftGraphPermissionType, bool)`

GetPermissionTypeOk returns a tuple with the PermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionType

`func (o *PermissionGrantConditionSet) SetPermissionType(v AnyOfmicrosoftGraphPermissionType)`

SetPermissionType sets PermissionType field to given value.

### HasPermissionType

`func (o *PermissionGrantConditionSet) HasPermissionType() bool`

HasPermissionType returns a boolean if a field has been set.

### SetPermissionTypeNil

`func (o *PermissionGrantConditionSet) SetPermissionTypeNil(b bool)`

 SetPermissionTypeNil sets the value for PermissionType to be an explicit nil

### UnsetPermissionType
`func (o *PermissionGrantConditionSet) UnsetPermissionType()`

UnsetPermissionType ensures that no value is present for PermissionType, not even an explicit nil
### GetResourceApplication

`func (o *PermissionGrantConditionSet) GetResourceApplication() string`

GetResourceApplication returns the ResourceApplication field if non-nil, zero value otherwise.

### GetResourceApplicationOk

`func (o *PermissionGrantConditionSet) GetResourceApplicationOk() (*string, bool)`

GetResourceApplicationOk returns a tuple with the ResourceApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceApplication

`func (o *PermissionGrantConditionSet) SetResourceApplication(v string)`

SetResourceApplication sets ResourceApplication field to given value.

### HasResourceApplication

`func (o *PermissionGrantConditionSet) HasResourceApplication() bool`

HasResourceApplication returns a boolean if a field has been set.

### SetResourceApplicationNil

`func (o *PermissionGrantConditionSet) SetResourceApplicationNil(b bool)`

 SetResourceApplicationNil sets the value for ResourceApplication to be an explicit nil

### UnsetResourceApplication
`func (o *PermissionGrantConditionSet) UnsetResourceApplication()`

UnsetResourceApplication ensures that no value is present for ResourceApplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


