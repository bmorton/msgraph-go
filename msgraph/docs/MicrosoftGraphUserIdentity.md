# MicrosoftGraphUserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The identity&#39;s display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won&#39;t show up as having changed when using delta. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier for the identity. | [optional] 
**IpAddress** | Pointer to **NullableString** | Indicates the client IP address used by user performing the activity (audit log only). | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The userPrincipalName attribute of the user. | [optional] 

## Methods

### NewMicrosoftGraphUserIdentity

`func NewMicrosoftGraphUserIdentity() *MicrosoftGraphUserIdentity`

NewMicrosoftGraphUserIdentity instantiates a new MicrosoftGraphUserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserIdentityWithDefaults

`func NewMicrosoftGraphUserIdentityWithDefaults() *MicrosoftGraphUserIdentity`

NewMicrosoftGraphUserIdentityWithDefaults instantiates a new MicrosoftGraphUserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphUserIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphUserIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphUserIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphUserIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphUserIdentity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphUserIdentity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphUserIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphUserIdentity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphUserIdentity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIpAddress

`func (o *MicrosoftGraphUserIdentity) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MicrosoftGraphUserIdentity) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MicrosoftGraphUserIdentity) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MicrosoftGraphUserIdentity) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *MicrosoftGraphUserIdentity) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *MicrosoftGraphUserIdentity) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphUserIdentity) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphUserIdentity) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphUserIdentity) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphUserIdentity) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphUserIdentity) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphUserIdentity) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


