# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDateTime** | Pointer to **NullableTime** | A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional. | [optional] 
**GrantedTo** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**GrantedToIdentities** | Pointer to [**[]AnyOfmicrosoftGraphIdentitySet**](AnyOfmicrosoftGraphIdentitySet.md) |  | [optional] 
**GrantedToIdentitiesV2** | Pointer to [**[]AnyOfmicrosoftGraphSharePointIdentitySet**](AnyOfmicrosoftGraphSharePointIdentitySet.md) | For link type permissions, the details of the users to whom permission was granted. Read-only. | [optional] 
**GrantedToV2** | Pointer to [**AnyOfmicrosoftGraphSharePointIdentitySet**](anyOf&lt;microsoft.graph.sharePointIdentitySet&gt;.md) | For user type permissions, the details of the users and applications for this permission. Read-only. | [optional] 
**HasPassword** | Pointer to **NullableBool** | Indicates whether the password is set for this permission. This property only appears in the response. Optional. Read-only. For OneDrive Personal only.. | [optional] 
**InheritedFrom** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) | Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only. | [optional] 
**Invitation** | Pointer to [**AnyOfmicrosoftGraphSharingInvitation**](anyOf&lt;microsoft.graph.sharingInvitation&gt;.md) | Details of any associated sharing invitation for this permission. Read-only. | [optional] 
**Link** | Pointer to [**AnyOfmicrosoftGraphSharingLink**](anyOf&lt;microsoft.graph.sharingLink&gt;.md) | Provides the link details of the current permission, if it is a link type permissions. Read-only. | [optional] 
**Roles** | Pointer to **[]string** | The type of permission, for example, read. See below for the full list of roles. Read-only. | [optional] 
**ShareId** | Pointer to **NullableString** | A unique token that can be used to access this shared item via the **shares** API. Read-only. | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDateTime

`func (o *Permission) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *Permission) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *Permission) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *Permission) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *Permission) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *Permission) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetGrantedTo

`func (o *Permission) GetGrantedTo() AnyOfmicrosoftGraphIdentitySet`

GetGrantedTo returns the GrantedTo field if non-nil, zero value otherwise.

### GetGrantedToOk

`func (o *Permission) GetGrantedToOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetGrantedToOk returns a tuple with the GrantedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedTo

`func (o *Permission) SetGrantedTo(v AnyOfmicrosoftGraphIdentitySet)`

SetGrantedTo sets GrantedTo field to given value.

### HasGrantedTo

`func (o *Permission) HasGrantedTo() bool`

HasGrantedTo returns a boolean if a field has been set.

### SetGrantedToNil

`func (o *Permission) SetGrantedToNil(b bool)`

 SetGrantedToNil sets the value for GrantedTo to be an explicit nil

### UnsetGrantedTo
`func (o *Permission) UnsetGrantedTo()`

UnsetGrantedTo ensures that no value is present for GrantedTo, not even an explicit nil
### GetGrantedToIdentities

`func (o *Permission) GetGrantedToIdentities() []*AnyOfmicrosoftGraphIdentitySet`

GetGrantedToIdentities returns the GrantedToIdentities field if non-nil, zero value otherwise.

### GetGrantedToIdentitiesOk

`func (o *Permission) GetGrantedToIdentitiesOk() (*[]*AnyOfmicrosoftGraphIdentitySet, bool)`

GetGrantedToIdentitiesOk returns a tuple with the GrantedToIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToIdentities

`func (o *Permission) SetGrantedToIdentities(v []*AnyOfmicrosoftGraphIdentitySet)`

SetGrantedToIdentities sets GrantedToIdentities field to given value.

### HasGrantedToIdentities

`func (o *Permission) HasGrantedToIdentities() bool`

HasGrantedToIdentities returns a boolean if a field has been set.

### GetGrantedToIdentitiesV2

`func (o *Permission) GetGrantedToIdentitiesV2() []*AnyOfmicrosoftGraphSharePointIdentitySet`

GetGrantedToIdentitiesV2 returns the GrantedToIdentitiesV2 field if non-nil, zero value otherwise.

### GetGrantedToIdentitiesV2Ok

`func (o *Permission) GetGrantedToIdentitiesV2Ok() (*[]*AnyOfmicrosoftGraphSharePointIdentitySet, bool)`

GetGrantedToIdentitiesV2Ok returns a tuple with the GrantedToIdentitiesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToIdentitiesV2

`func (o *Permission) SetGrantedToIdentitiesV2(v []*AnyOfmicrosoftGraphSharePointIdentitySet)`

SetGrantedToIdentitiesV2 sets GrantedToIdentitiesV2 field to given value.

### HasGrantedToIdentitiesV2

`func (o *Permission) HasGrantedToIdentitiesV2() bool`

HasGrantedToIdentitiesV2 returns a boolean if a field has been set.

### GetGrantedToV2

`func (o *Permission) GetGrantedToV2() AnyOfmicrosoftGraphSharePointIdentitySet`

GetGrantedToV2 returns the GrantedToV2 field if non-nil, zero value otherwise.

### GetGrantedToV2Ok

`func (o *Permission) GetGrantedToV2Ok() (*AnyOfmicrosoftGraphSharePointIdentitySet, bool)`

GetGrantedToV2Ok returns a tuple with the GrantedToV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToV2

`func (o *Permission) SetGrantedToV2(v AnyOfmicrosoftGraphSharePointIdentitySet)`

SetGrantedToV2 sets GrantedToV2 field to given value.

### HasGrantedToV2

`func (o *Permission) HasGrantedToV2() bool`

HasGrantedToV2 returns a boolean if a field has been set.

### SetGrantedToV2Nil

`func (o *Permission) SetGrantedToV2Nil(b bool)`

 SetGrantedToV2Nil sets the value for GrantedToV2 to be an explicit nil

### UnsetGrantedToV2
`func (o *Permission) UnsetGrantedToV2()`

UnsetGrantedToV2 ensures that no value is present for GrantedToV2, not even an explicit nil
### GetHasPassword

`func (o *Permission) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *Permission) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *Permission) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *Permission) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### SetHasPasswordNil

`func (o *Permission) SetHasPasswordNil(b bool)`

 SetHasPasswordNil sets the value for HasPassword to be an explicit nil

### UnsetHasPassword
`func (o *Permission) UnsetHasPassword()`

UnsetHasPassword ensures that no value is present for HasPassword, not even an explicit nil
### GetInheritedFrom

`func (o *Permission) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *Permission) GetInheritedFromOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFrom

`func (o *Permission) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom sets InheritedFrom field to given value.

### HasInheritedFrom

`func (o *Permission) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFromNil

`func (o *Permission) SetInheritedFromNil(b bool)`

 SetInheritedFromNil sets the value for InheritedFrom to be an explicit nil

### UnsetInheritedFrom
`func (o *Permission) UnsetInheritedFrom()`

UnsetInheritedFrom ensures that no value is present for InheritedFrom, not even an explicit nil
### GetInvitation

`func (o *Permission) GetInvitation() AnyOfmicrosoftGraphSharingInvitation`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *Permission) GetInvitationOk() (*AnyOfmicrosoftGraphSharingInvitation, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *Permission) SetInvitation(v AnyOfmicrosoftGraphSharingInvitation)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *Permission) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### SetInvitationNil

`func (o *Permission) SetInvitationNil(b bool)`

 SetInvitationNil sets the value for Invitation to be an explicit nil

### UnsetInvitation
`func (o *Permission) UnsetInvitation()`

UnsetInvitation ensures that no value is present for Invitation, not even an explicit nil
### GetLink

`func (o *Permission) GetLink() AnyOfmicrosoftGraphSharingLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Permission) GetLinkOk() (*AnyOfmicrosoftGraphSharingLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Permission) SetLink(v AnyOfmicrosoftGraphSharingLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *Permission) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *Permission) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *Permission) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetRoles

`func (o *Permission) GetRoles() []*string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Permission) GetRolesOk() (*[]*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Permission) SetRoles(v []*string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Permission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetShareId

`func (o *Permission) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *Permission) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *Permission) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *Permission) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareIdNil

`func (o *Permission) SetShareIdNil(b bool)`

 SetShareIdNil sets the value for ShareId to be an explicit nil

### UnsetShareId
`func (o *Permission) UnsetShareId()`

UnsetShareId ensures that no value is present for ShareId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


