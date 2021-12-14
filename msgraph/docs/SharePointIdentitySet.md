# SharePointIdentitySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | The group associated with this action. Optional. | [optional] 
**SiteGroup** | Pointer to [**AnyOfmicrosoftGraphSharePointIdentity**](anyOf&lt;microsoft.graph.sharePointIdentity&gt;.md) | The SharePoint group associated with this action. Optional. | [optional] 
**SiteUser** | Pointer to [**AnyOfmicrosoftGraphSharePointIdentity**](anyOf&lt;microsoft.graph.sharePointIdentity&gt;.md) | The SharePoint user associated with this action. Optional. | [optional] 

## Methods

### NewSharePointIdentitySet

`func NewSharePointIdentitySet() *SharePointIdentitySet`

NewSharePointIdentitySet instantiates a new SharePointIdentitySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePointIdentitySetWithDefaults

`func NewSharePointIdentitySetWithDefaults() *SharePointIdentitySet`

NewSharePointIdentitySetWithDefaults instantiates a new SharePointIdentitySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SharePointIdentitySet) GetGroup() AnyOfmicrosoftGraphIdentity`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SharePointIdentitySet) GetGroupOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SharePointIdentitySet) SetGroup(v AnyOfmicrosoftGraphIdentity)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SharePointIdentitySet) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *SharePointIdentitySet) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *SharePointIdentitySet) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSiteGroup

`func (o *SharePointIdentitySet) GetSiteGroup() AnyOfmicrosoftGraphSharePointIdentity`

GetSiteGroup returns the SiteGroup field if non-nil, zero value otherwise.

### GetSiteGroupOk

`func (o *SharePointIdentitySet) GetSiteGroupOk() (*AnyOfmicrosoftGraphSharePointIdentity, bool)`

GetSiteGroupOk returns a tuple with the SiteGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroup

`func (o *SharePointIdentitySet) SetSiteGroup(v AnyOfmicrosoftGraphSharePointIdentity)`

SetSiteGroup sets SiteGroup field to given value.

### HasSiteGroup

`func (o *SharePointIdentitySet) HasSiteGroup() bool`

HasSiteGroup returns a boolean if a field has been set.

### SetSiteGroupNil

`func (o *SharePointIdentitySet) SetSiteGroupNil(b bool)`

 SetSiteGroupNil sets the value for SiteGroup to be an explicit nil

### UnsetSiteGroup
`func (o *SharePointIdentitySet) UnsetSiteGroup()`

UnsetSiteGroup ensures that no value is present for SiteGroup, not even an explicit nil
### GetSiteUser

`func (o *SharePointIdentitySet) GetSiteUser() AnyOfmicrosoftGraphSharePointIdentity`

GetSiteUser returns the SiteUser field if non-nil, zero value otherwise.

### GetSiteUserOk

`func (o *SharePointIdentitySet) GetSiteUserOk() (*AnyOfmicrosoftGraphSharePointIdentity, bool)`

GetSiteUserOk returns a tuple with the SiteUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUser

`func (o *SharePointIdentitySet) SetSiteUser(v AnyOfmicrosoftGraphSharePointIdentity)`

SetSiteUser sets SiteUser field to given value.

### HasSiteUser

`func (o *SharePointIdentitySet) HasSiteUser() bool`

HasSiteUser returns a boolean if a field has been set.

### SetSiteUserNil

`func (o *SharePointIdentitySet) SetSiteUserNil(b bool)`

 SetSiteUserNil sets the value for SiteUser to be an explicit nil

### UnsetSiteUser
`func (o *SharePointIdentitySet) UnsetSiteUser()`

UnsetSiteUser ensures that no value is present for SiteUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


