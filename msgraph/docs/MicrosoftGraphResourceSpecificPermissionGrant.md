# MicrosoftGraphResourceSpecificPermissionGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**ClientAppId** | Pointer to **NullableString** | ID of the service principal of the Azure AD app that has been granted access. Read-only. | [optional] 
**ClientId** | Pointer to **NullableString** | ID of the Azure AD app that has been granted access. Read-only. | [optional] 
**Permission** | Pointer to **NullableString** | The name of the resource-specific permission. Read-only. | [optional] 
**PermissionType** | Pointer to **NullableString** | The type of permission. Possible values are: Application, Delegated. Read-only. | [optional] 
**ResourceAppId** | Pointer to **NullableString** | ID of the Azure AD app that is hosting the resource. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphResourceSpecificPermissionGrant

`func NewMicrosoftGraphResourceSpecificPermissionGrant() *MicrosoftGraphResourceSpecificPermissionGrant`

NewMicrosoftGraphResourceSpecificPermissionGrant instantiates a new MicrosoftGraphResourceSpecificPermissionGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphResourceSpecificPermissionGrantWithDefaults

`func NewMicrosoftGraphResourceSpecificPermissionGrantWithDefaults() *MicrosoftGraphResourceSpecificPermissionGrant`

NewMicrosoftGraphResourceSpecificPermissionGrantWithDefaults instantiates a new MicrosoftGraphResourceSpecificPermissionGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphResourceSpecificPermissionGrant) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetClientAppId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetClientAppId() string`

GetClientAppId returns the ClientAppId field if non-nil, zero value otherwise.

### GetClientAppIdOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetClientAppIdOk() (*string, bool)`

GetClientAppIdOk returns a tuple with the ClientAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetClientAppId(v string)`

SetClientAppId sets ClientAppId field to given value.

### HasClientAppId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasClientAppId() bool`

HasClientAppId returns a boolean if a field has been set.

### SetClientAppIdNil

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetClientAppIdNil(b bool)`

 SetClientAppIdNil sets the value for ClientAppId to be an explicit nil

### UnsetClientAppId
`func (o *MicrosoftGraphResourceSpecificPermissionGrant) UnsetClientAppId()`

UnsetClientAppId ensures that no value is present for ClientAppId, not even an explicit nil
### GetClientId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *MicrosoftGraphResourceSpecificPermissionGrant) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetPermission

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *MicrosoftGraphResourceSpecificPermissionGrant) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetPermissionType

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetPermissionType() string`

GetPermissionType returns the PermissionType field if non-nil, zero value otherwise.

### GetPermissionTypeOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetPermissionTypeOk() (*string, bool)`

GetPermissionTypeOk returns a tuple with the PermissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionType

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetPermissionType(v string)`

SetPermissionType sets PermissionType field to given value.

### HasPermissionType

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasPermissionType() bool`

HasPermissionType returns a boolean if a field has been set.

### SetPermissionTypeNil

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetPermissionTypeNil(b bool)`

 SetPermissionTypeNil sets the value for PermissionType to be an explicit nil

### UnsetPermissionType
`func (o *MicrosoftGraphResourceSpecificPermissionGrant) UnsetPermissionType()`

UnsetPermissionType ensures that no value is present for PermissionType, not even an explicit nil
### GetResourceAppId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetResourceAppId() string`

GetResourceAppId returns the ResourceAppId field if non-nil, zero value otherwise.

### GetResourceAppIdOk

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) GetResourceAppIdOk() (*string, bool)`

GetResourceAppIdOk returns a tuple with the ResourceAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAppId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetResourceAppId(v string)`

SetResourceAppId sets ResourceAppId field to given value.

### HasResourceAppId

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) HasResourceAppId() bool`

HasResourceAppId returns a boolean if a field has been set.

### SetResourceAppIdNil

`func (o *MicrosoftGraphResourceSpecificPermissionGrant) SetResourceAppIdNil(b bool)`

 SetResourceAppIdNil sets the value for ResourceAppId to be an explicit nil

### UnsetResourceAppId
`func (o *MicrosoftGraphResourceSpecificPermissionGrant) UnsetResourceAppId()`

UnsetResourceAppId ensures that no value is present for ResourceAppId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


