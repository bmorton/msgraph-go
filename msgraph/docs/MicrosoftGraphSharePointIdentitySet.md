# MicrosoftGraphSharePointIdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Optional. The application associated with this action. | [optional] 
**Device** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Optional. The device associated with this action. | [optional] 
**User** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Optional. The user associated with this action. | [optional] 
**Group** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | The group associated with this action. Optional. | [optional] 
**SiteGroup** | Pointer to [**AnyOfmicrosoftGraphSharePointIdentity**](anyOf&lt;microsoft.graph.sharePointIdentity&gt;.md) | The SharePoint group associated with this action. Optional. | [optional] 
**SiteUser** | Pointer to [**AnyOfmicrosoftGraphSharePointIdentity**](anyOf&lt;microsoft.graph.sharePointIdentity&gt;.md) | The SharePoint user associated with this action. Optional. | [optional] 

## Methods

### NewMicrosoftGraphSharePointIdentitySet

`func NewMicrosoftGraphSharePointIdentitySet() *MicrosoftGraphSharePointIdentitySet`

NewMicrosoftGraphSharePointIdentitySet instantiates a new MicrosoftGraphSharePointIdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharePointIdentitySetWithDefaults

`func NewMicrosoftGraphSharePointIdentitySetWithDefaults() *MicrosoftGraphSharePointIdentitySet`

NewMicrosoftGraphSharePointIdentitySetWithDefaults instantiates a new MicrosoftGraphSharePointIdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *MicrosoftGraphSharePointIdentitySet) GetApplication() AnyOfmicrosoftGraphIdentity`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MicrosoftGraphSharePointIdentitySet) GetApplicationOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MicrosoftGraphSharePointIdentitySet) SetApplication(v AnyOfmicrosoftGraphIdentity)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MicrosoftGraphSharePointIdentitySet) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *MicrosoftGraphSharePointIdentitySet) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *MicrosoftGraphSharePointIdentitySet) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetDevice

`func (o *MicrosoftGraphSharePointIdentitySet) GetDevice() AnyOfmicrosoftGraphIdentity`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MicrosoftGraphSharePointIdentitySet) GetDeviceOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MicrosoftGraphSharePointIdentitySet) SetDevice(v AnyOfmicrosoftGraphIdentity)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MicrosoftGraphSharePointIdentitySet) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *MicrosoftGraphSharePointIdentitySet) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *MicrosoftGraphSharePointIdentitySet) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetUser

`func (o *MicrosoftGraphSharePointIdentitySet) GetUser() AnyOfmicrosoftGraphIdentity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphSharePointIdentitySet) GetUserOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftGraphSharePointIdentitySet) SetUser(v AnyOfmicrosoftGraphIdentity)`

SetUser sets User field to given value.

### HasUser

`func (o *MicrosoftGraphSharePointIdentitySet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MicrosoftGraphSharePointIdentitySet) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MicrosoftGraphSharePointIdentitySet) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGroup

`func (o *MicrosoftGraphSharePointIdentitySet) GetGroup() AnyOfmicrosoftGraphIdentity`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MicrosoftGraphSharePointIdentitySet) GetGroupOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MicrosoftGraphSharePointIdentitySet) SetGroup(v AnyOfmicrosoftGraphIdentity)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MicrosoftGraphSharePointIdentitySet) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *MicrosoftGraphSharePointIdentitySet) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *MicrosoftGraphSharePointIdentitySet) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSiteGroup

`func (o *MicrosoftGraphSharePointIdentitySet) GetSiteGroup() AnyOfmicrosoftGraphSharePointIdentity`

GetSiteGroup returns the SiteGroup field if non-nil, zero value otherwise.

### GetSiteGroupOk

`func (o *MicrosoftGraphSharePointIdentitySet) GetSiteGroupOk() (*AnyOfmicrosoftGraphSharePointIdentity, bool)`

GetSiteGroupOk returns a tuple with the SiteGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroup

`func (o *MicrosoftGraphSharePointIdentitySet) SetSiteGroup(v AnyOfmicrosoftGraphSharePointIdentity)`

SetSiteGroup sets SiteGroup field to given value.

### HasSiteGroup

`func (o *MicrosoftGraphSharePointIdentitySet) HasSiteGroup() bool`

HasSiteGroup returns a boolean if a field has been set.

### SetSiteGroupNil

`func (o *MicrosoftGraphSharePointIdentitySet) SetSiteGroupNil(b bool)`

 SetSiteGroupNil sets the value for SiteGroup to be an explicit nil

### UnsetSiteGroup
`func (o *MicrosoftGraphSharePointIdentitySet) UnsetSiteGroup()`

UnsetSiteGroup ensures that no value is present for SiteGroup, not even an explicit nil
### GetSiteUser

`func (o *MicrosoftGraphSharePointIdentitySet) GetSiteUser() AnyOfmicrosoftGraphSharePointIdentity`

GetSiteUser returns the SiteUser field if non-nil, zero value otherwise.

### GetSiteUserOk

`func (o *MicrosoftGraphSharePointIdentitySet) GetSiteUserOk() (*AnyOfmicrosoftGraphSharePointIdentity, bool)`

GetSiteUserOk returns a tuple with the SiteUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUser

`func (o *MicrosoftGraphSharePointIdentitySet) SetSiteUser(v AnyOfmicrosoftGraphSharePointIdentity)`

SetSiteUser sets SiteUser field to given value.

### HasSiteUser

`func (o *MicrosoftGraphSharePointIdentitySet) HasSiteUser() bool`

HasSiteUser returns a boolean if a field has been set.

### SetSiteUserNil

`func (o *MicrosoftGraphSharePointIdentitySet) SetSiteUserNil(b bool)`

 SetSiteUserNil sets the value for SiteUser to be an explicit nil

### UnsetSiteUser
`func (o *MicrosoftGraphSharePointIdentitySet) UnsetSiteUser()`

UnsetSiteUser ensures that no value is present for SiteUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


