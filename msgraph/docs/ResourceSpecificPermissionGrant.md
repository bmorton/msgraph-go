# ResourceSpecificPermissionGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAppId** | Pointer to **NullableString** | ID of the service principal of the Azure AD app that has been granted access. Read-only. | [optional] 
**ClientId** | Pointer to **NullableString** | ID of the Azure AD app that has been granted access. Read-only. | [optional] 
**Permission** | Pointer to **NullableString** | The name of the resource-specific permission. Read-only. | [optional] 
**PermissionType** | Pointer to **NullableString** | The type of permission. Possible values are: Application, Delegated. Read-only. | [optional] 
**ResourceAppId** | Pointer to **NullableString** | ID of the Azure AD app that is hosting the resource. Read-only. | [optional] 

## Methods

### NewResourceSpecificPermissionGrant

`func NewResourceSpecificPermissionGrant() *ResourceSpecificPermissionGrant`

NewResourceSpecificPermissionGrant instantiates a new ResourceSpecificPermissionGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSpecificPermissionGrantWithDefaults

`func NewResourceSpecificPermissionGrantWithDefaults() *ResourceSpecificPermissionGrant`

NewResourceSpecificPermissionGrantWithDefaults instantiates a new ResourceSpecificPermissionGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAppId

`func (o *ResourceSpecificPermissionGrant) GetClientAppId() string`

GetClientAppId returns the ClientAppId field if non-nil, zero value otherwise.

### GetClientAppIdOk

`func (o *ResourceSpecificPermissionGrant) GetClientAppIdOk() (*string, bool)`

GetClientAppIdOk returns a tuple with the ClientAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppId

`func (o *ResourceSpecificPermissionGrant) SetClientAppId(v string)`

SetClientAppId sets ClientAppId field to given value.

### HasClientAppId

`func (o *ResourceSpecificPermissionGrant) HasClientAppId() bool`

HasClientAppId returns a boolean if a field has been set.

### SetClientAppIdNil

`func (o *ResourceSpecificPermissionGrant) SetClientAppIdNil(b bool)`

 SetClientAppIdNil sets the value for ClientAppId to be an explicit nil

### UnsetClientAppId
`func (o *ResourceSpecificPermissionGrant) UnsetClientAppId()`

UnsetClientAppId ensures that no value is present for ClientAppId, not even an explicit nil
### GetClientId

`func (o *ResourceSpecificPermissionGrant) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ResourceSpecificPermissionGrant) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ResourceSpecificPermissionGrant) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ResourceSpecificPermissionGrant) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *ResourceSpecificPermissionGrant) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *ResourceSpecificPermissionGrant) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetPermission

`func (o *ResourceSpecificPermissionGrant) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ResourceSpecificPermissionGrant) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ResourceSpecificPermissionGrant) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ResourceSpecificPermissionGrant) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *ResourceSpecificPermissionGrant) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *ResourceSpecificPermissionGrant) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetPermissionType

`func (o *ResourceSpecificPermissionGrant) GetPermissionType() string`

GetPermissionType returns the PermissionType field if non-nil, zero value otherwise.

### GetPermissionTypeOk

`func (o *ResourceSpecificPermissionGrant) GetPermissionTypeOk() (*string, bool)`

GetPermissionTypeOk returns a tuple with the PermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionType

`func (o *ResourceSpecificPermissionGrant) SetPermissionType(v string)`

SetPermissionType sets PermissionType field to given value.

### HasPermissionType

`func (o *ResourceSpecificPermissionGrant) HasPermissionType() bool`

HasPermissionType returns a boolean if a field has been set.

### SetPermissionTypeNil

`func (o *ResourceSpecificPermissionGrant) SetPermissionTypeNil(b bool)`

 SetPermissionTypeNil sets the value for PermissionType to be an explicit nil

### UnsetPermissionType
`func (o *ResourceSpecificPermissionGrant) UnsetPermissionType()`

UnsetPermissionType ensures that no value is present for PermissionType, not even an explicit nil
### GetResourceAppId

`func (o *ResourceSpecificPermissionGrant) GetResourceAppId() string`

GetResourceAppId returns the ResourceAppId field if non-nil, zero value otherwise.

### GetResourceAppIdOk

`func (o *ResourceSpecificPermissionGrant) GetResourceAppIdOk() (*string, bool)`

GetResourceAppIdOk returns a tuple with the ResourceAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAppId

`func (o *ResourceSpecificPermissionGrant) SetResourceAppId(v string)`

SetResourceAppId sets ResourceAppId field to given value.

### HasResourceAppId

`func (o *ResourceSpecificPermissionGrant) HasResourceAppId() bool`

HasResourceAppId returns a boolean if a field has been set.

### SetResourceAppIdNil

`func (o *ResourceSpecificPermissionGrant) SetResourceAppIdNil(b bool)`

 SetResourceAppIdNil sets the value for ResourceAppId to be an explicit nil

### UnsetResourceAppId
`func (o *ResourceSpecificPermissionGrant) UnsetResourceAppId()`

UnsetResourceAppId ensures that no value is present for ResourceAppId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


