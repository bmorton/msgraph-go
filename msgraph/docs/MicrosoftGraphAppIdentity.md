# MicrosoftGraphAppIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableString** | Refers to the Unique GUID representing Application Id in the Azure Active Directory. | [optional] 
**DisplayName** | Pointer to **NullableString** | Refers to the Application Name displayed in the Azure Portal. | [optional] 
**ServicePrincipalId** | Pointer to **NullableString** | Refers to the Unique GUID indicating Service Principal Id in Azure Active Directory for the corresponding App. | [optional] 
**ServicePrincipalName** | Pointer to **NullableString** | Refers to the Service Principal Name is the Application name in the tenant. | [optional] 

## Methods

### NewMicrosoftGraphAppIdentity

`func NewMicrosoftGraphAppIdentity() *MicrosoftGraphAppIdentity`

NewMicrosoftGraphAppIdentity instantiates a new MicrosoftGraphAppIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAppIdentityWithDefaults

`func NewMicrosoftGraphAppIdentityWithDefaults() *MicrosoftGraphAppIdentity`

NewMicrosoftGraphAppIdentityWithDefaults instantiates a new MicrosoftGraphAppIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *MicrosoftGraphAppIdentity) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphAppIdentity) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphAppIdentity) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphAppIdentity) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *MicrosoftGraphAppIdentity) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *MicrosoftGraphAppIdentity) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAppIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAppIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAppIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAppIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAppIdentity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAppIdentity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetServicePrincipalId

`func (o *MicrosoftGraphAppIdentity) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *MicrosoftGraphAppIdentity) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *MicrosoftGraphAppIdentity) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *MicrosoftGraphAppIdentity) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.

### SetServicePrincipalIdNil

`func (o *MicrosoftGraphAppIdentity) SetServicePrincipalIdNil(b bool)`

 SetServicePrincipalIdNil sets the value for ServicePrincipalId to be an explicit nil

### UnsetServicePrincipalId
`func (o *MicrosoftGraphAppIdentity) UnsetServicePrincipalId()`

UnsetServicePrincipalId ensures that no value is present for ServicePrincipalId, not even an explicit nil
### GetServicePrincipalName

`func (o *MicrosoftGraphAppIdentity) GetServicePrincipalName() string`

GetServicePrincipalName returns the ServicePrincipalName field if non-nil, zero value otherwise.

### GetServicePrincipalNameOk

`func (o *MicrosoftGraphAppIdentity) GetServicePrincipalNameOk() (*string, bool)`

GetServicePrincipalNameOk returns a tuple with the ServicePrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalName

`func (o *MicrosoftGraphAppIdentity) SetServicePrincipalName(v string)`

SetServicePrincipalName sets ServicePrincipalName field to given value.

### HasServicePrincipalName

`func (o *MicrosoftGraphAppIdentity) HasServicePrincipalName() bool`

HasServicePrincipalName returns a boolean if a field has been set.

### SetServicePrincipalNameNil

`func (o *MicrosoftGraphAppIdentity) SetServicePrincipalNameNil(b bool)`

 SetServicePrincipalNameNil sets the value for ServicePrincipalName to be an explicit nil

### UnsetServicePrincipalName
`func (o *MicrosoftGraphAppIdentity) UnsetServicePrincipalName()`

UnsetServicePrincipalName ensures that no value is present for ServicePrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


