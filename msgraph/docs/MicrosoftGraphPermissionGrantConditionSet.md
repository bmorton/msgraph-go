# MicrosoftGraphPermissionGrantConditionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClientApplicationIds** | Pointer to **[]string** | A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all. | [optional] 
**ClientApplicationPublisherIds** | Pointer to **[]string** | A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all. | [optional] 
**ClientApplicationsFromVerifiedPublisherOnly** | Pointer to **NullableBool** | Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false. | [optional] 
**ClientApplicationTenantIds** | Pointer to **[]string** | A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all. | [optional] 
**PermissionClassification** | Pointer to **NullableString** | The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all. | [optional] 
**Permissions** | Pointer to **[]string** | The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API&#39;s **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API&#39;s **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API&#39;s **servicePrincipal** object. Default is the single value all. | [optional] 
**PermissionType** | Pointer to [**AnyOfmicrosoftGraphPermissionType**](anyOf&lt;microsoft.graph.permissionType&gt;.md) | The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required. | [optional] 
**ResourceApplication** | Pointer to **NullableString** | The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any. | [optional] 

## Methods

### NewMicrosoftGraphPermissionGrantConditionSet

`func NewMicrosoftGraphPermissionGrantConditionSet() *MicrosoftGraphPermissionGrantConditionSet`

NewMicrosoftGraphPermissionGrantConditionSet instantiates a new MicrosoftGraphPermissionGrantConditionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPermissionGrantConditionSetWithDefaults

`func NewMicrosoftGraphPermissionGrantConditionSetWithDefaults() *MicrosoftGraphPermissionGrantConditionSet`

NewMicrosoftGraphPermissionGrantConditionSetWithDefaults instantiates a new MicrosoftGraphPermissionGrantConditionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientApplicationIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationIds() []*string`

GetClientApplicationIds returns the ClientApplicationIds field if non-nil, zero value otherwise.

### GetClientApplicationIdsOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationIdsOk() (*[]*string, bool)`

GetClientApplicationIdsOk returns a tuple with the ClientApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetClientApplicationIds(v []*string)`

SetClientApplicationIds sets ClientApplicationIds field to given value.

### HasClientApplicationIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasClientApplicationIds() bool`

HasClientApplicationIds returns a boolean if a field has been set.

### GetClientApplicationPublisherIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationPublisherIds() []*string`

GetClientApplicationPublisherIds returns the ClientApplicationPublisherIds field if non-nil, zero value otherwise.

### GetClientApplicationPublisherIdsOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationPublisherIdsOk() (*[]*string, bool)`

GetClientApplicationPublisherIdsOk returns a tuple with the ClientApplicationPublisherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationPublisherIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetClientApplicationPublisherIds(v []*string)`

SetClientApplicationPublisherIds sets ClientApplicationPublisherIds field to given value.

### HasClientApplicationPublisherIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasClientApplicationPublisherIds() bool`

HasClientApplicationPublisherIds returns a boolean if a field has been set.

### GetClientApplicationsFromVerifiedPublisherOnly

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnly() bool`

GetClientApplicationsFromVerifiedPublisherOnly returns the ClientApplicationsFromVerifiedPublisherOnly field if non-nil, zero value otherwise.

### GetClientApplicationsFromVerifiedPublisherOnlyOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnlyOk() (*bool, bool)`

GetClientApplicationsFromVerifiedPublisherOnlyOk returns a tuple with the ClientApplicationsFromVerifiedPublisherOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationsFromVerifiedPublisherOnly

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnly(v bool)`

SetClientApplicationsFromVerifiedPublisherOnly sets ClientApplicationsFromVerifiedPublisherOnly field to given value.

### HasClientApplicationsFromVerifiedPublisherOnly

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasClientApplicationsFromVerifiedPublisherOnly() bool`

HasClientApplicationsFromVerifiedPublisherOnly returns a boolean if a field has been set.

### SetClientApplicationsFromVerifiedPublisherOnlyNil

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnlyNil(b bool)`

 SetClientApplicationsFromVerifiedPublisherOnlyNil sets the value for ClientApplicationsFromVerifiedPublisherOnly to be an explicit nil

### UnsetClientApplicationsFromVerifiedPublisherOnly
`func (o *MicrosoftGraphPermissionGrantConditionSet) UnsetClientApplicationsFromVerifiedPublisherOnly()`

UnsetClientApplicationsFromVerifiedPublisherOnly ensures that no value is present for ClientApplicationsFromVerifiedPublisherOnly, not even an explicit nil
### GetClientApplicationTenantIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationTenantIds() []*string`

GetClientApplicationTenantIds returns the ClientApplicationTenantIds field if non-nil, zero value otherwise.

### GetClientApplicationTenantIdsOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetClientApplicationTenantIdsOk() (*[]*string, bool)`

GetClientApplicationTenantIdsOk returns a tuple with the ClientApplicationTenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApplicationTenantIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetClientApplicationTenantIds(v []*string)`

SetClientApplicationTenantIds sets ClientApplicationTenantIds field to given value.

### HasClientApplicationTenantIds

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasClientApplicationTenantIds() bool`

HasClientApplicationTenantIds returns a boolean if a field has been set.

### GetPermissionClassification

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetPermissionClassification() string`

GetPermissionClassification returns the PermissionClassification field if non-nil, zero value otherwise.

### GetPermissionClassificationOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetPermissionClassificationOk() (*string, bool)`

GetPermissionClassificationOk returns a tuple with the PermissionClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionClassification

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetPermissionClassification(v string)`

SetPermissionClassification sets PermissionClassification field to given value.

### HasPermissionClassification

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasPermissionClassification() bool`

HasPermissionClassification returns a boolean if a field has been set.

### SetPermissionClassificationNil

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetPermissionClassificationNil(b bool)`

 SetPermissionClassificationNil sets the value for PermissionClassification to be an explicit nil

### UnsetPermissionClassification
`func (o *MicrosoftGraphPermissionGrantConditionSet) UnsetPermissionClassification()`

UnsetPermissionClassification ensures that no value is present for PermissionClassification, not even an explicit nil
### GetPermissions

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetPermissions() []*string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetPermissionsOk() (*[]*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetPermissions(v []*string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPermissionType

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetPermissionType() AnyOfmicrosoftGraphPermissionType`

GetPermissionType returns the PermissionType field if non-nil, zero value otherwise.

### GetPermissionTypeOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetPermissionTypeOk() (*AnyOfmicrosoftGraphPermissionType, bool)`

GetPermissionTypeOk returns a tuple with the PermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionType

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetPermissionType(v AnyOfmicrosoftGraphPermissionType)`

SetPermissionType sets PermissionType field to given value.

### HasPermissionType

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasPermissionType() bool`

HasPermissionType returns a boolean if a field has been set.

### SetPermissionTypeNil

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetPermissionTypeNil(b bool)`

 SetPermissionTypeNil sets the value for PermissionType to be an explicit nil

### UnsetPermissionType
`func (o *MicrosoftGraphPermissionGrantConditionSet) UnsetPermissionType()`

UnsetPermissionType ensures that no value is present for PermissionType, not even an explicit nil
### GetResourceApplication

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetResourceApplication() string`

GetResourceApplication returns the ResourceApplication field if non-nil, zero value otherwise.

### GetResourceApplicationOk

`func (o *MicrosoftGraphPermissionGrantConditionSet) GetResourceApplicationOk() (*string, bool)`

GetResourceApplicationOk returns a tuple with the ResourceApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceApplication

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetResourceApplication(v string)`

SetResourceApplication sets ResourceApplication field to given value.

### HasResourceApplication

`func (o *MicrosoftGraphPermissionGrantConditionSet) HasResourceApplication() bool`

HasResourceApplication returns a boolean if a field has been set.

### SetResourceApplicationNil

`func (o *MicrosoftGraphPermissionGrantConditionSet) SetResourceApplicationNil(b bool)`

 SetResourceApplicationNil sets the value for ResourceApplication to be an explicit nil

### UnsetResourceApplication
`func (o *MicrosoftGraphPermissionGrantConditionSet) UnsetResourceApplication()`

UnsetResourceApplication ensures that no value is present for ResourceApplication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


