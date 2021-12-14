# MicrosoftGraphPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphPermission

`func NewMicrosoftGraphPermission() *MicrosoftGraphPermission`

NewMicrosoftGraphPermission instantiates a new MicrosoftGraphPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPermissionWithDefaults

`func NewMicrosoftGraphPermissionWithDefaults() *MicrosoftGraphPermission`

NewMicrosoftGraphPermissionWithDefaults instantiates a new MicrosoftGraphPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *MicrosoftGraphPermission) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphPermission) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphPermission) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphPermission) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *MicrosoftGraphPermission) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *MicrosoftGraphPermission) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetGrantedTo

`func (o *MicrosoftGraphPermission) GetGrantedTo() AnyOfmicrosoftGraphIdentitySet`

GetGrantedTo returns the GrantedTo field if non-nil, zero value otherwise.

### GetGrantedToOk

`func (o *MicrosoftGraphPermission) GetGrantedToOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetGrantedToOk returns a tuple with the GrantedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedTo

`func (o *MicrosoftGraphPermission) SetGrantedTo(v AnyOfmicrosoftGraphIdentitySet)`

SetGrantedTo sets GrantedTo field to given value.

### HasGrantedTo

`func (o *MicrosoftGraphPermission) HasGrantedTo() bool`

HasGrantedTo returns a boolean if a field has been set.

### SetGrantedToNil

`func (o *MicrosoftGraphPermission) SetGrantedToNil(b bool)`

 SetGrantedToNil sets the value for GrantedTo to be an explicit nil

### UnsetGrantedTo
`func (o *MicrosoftGraphPermission) UnsetGrantedTo()`

UnsetGrantedTo ensures that no value is present for GrantedTo, not even an explicit nil
### GetGrantedToIdentities

`func (o *MicrosoftGraphPermission) GetGrantedToIdentities() []*AnyOfmicrosoftGraphIdentitySet`

GetGrantedToIdentities returns the GrantedToIdentities field if non-nil, zero value otherwise.

### GetGrantedToIdentitiesOk

`func (o *MicrosoftGraphPermission) GetGrantedToIdentitiesOk() (*[]*AnyOfmicrosoftGraphIdentitySet, bool)`

GetGrantedToIdentitiesOk returns a tuple with the GrantedToIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToIdentities

`func (o *MicrosoftGraphPermission) SetGrantedToIdentities(v []*AnyOfmicrosoftGraphIdentitySet)`

SetGrantedToIdentities sets GrantedToIdentities field to given value.

### HasGrantedToIdentities

`func (o *MicrosoftGraphPermission) HasGrantedToIdentities() bool`

HasGrantedToIdentities returns a boolean if a field has been set.

### GetGrantedToIdentitiesV2

`func (o *MicrosoftGraphPermission) GetGrantedToIdentitiesV2() []*AnyOfmicrosoftGraphSharePointIdentitySet`

GetGrantedToIdentitiesV2 returns the GrantedToIdentitiesV2 field if non-nil, zero value otherwise.

### GetGrantedToIdentitiesV2Ok

`func (o *MicrosoftGraphPermission) GetGrantedToIdentitiesV2Ok() (*[]*AnyOfmicrosoftGraphSharePointIdentitySet, bool)`

GetGrantedToIdentitiesV2Ok returns a tuple with the GrantedToIdentitiesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToIdentitiesV2

`func (o *MicrosoftGraphPermission) SetGrantedToIdentitiesV2(v []*AnyOfmicrosoftGraphSharePointIdentitySet)`

SetGrantedToIdentitiesV2 sets GrantedToIdentitiesV2 field to given value.

### HasGrantedToIdentitiesV2

`func (o *MicrosoftGraphPermission) HasGrantedToIdentitiesV2() bool`

HasGrantedToIdentitiesV2 returns a boolean if a field has been set.

### GetGrantedToV2

`func (o *MicrosoftGraphPermission) GetGrantedToV2() AnyOfmicrosoftGraphSharePointIdentitySet`

GetGrantedToV2 returns the GrantedToV2 field if non-nil, zero value otherwise.

### GetGrantedToV2Ok

`func (o *MicrosoftGraphPermission) GetGrantedToV2Ok() (*AnyOfmicrosoftGraphSharePointIdentitySet, bool)`

GetGrantedToV2Ok returns a tuple with the GrantedToV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedToV2

`func (o *MicrosoftGraphPermission) SetGrantedToV2(v AnyOfmicrosoftGraphSharePointIdentitySet)`

SetGrantedToV2 sets GrantedToV2 field to given value.

### HasGrantedToV2

`func (o *MicrosoftGraphPermission) HasGrantedToV2() bool`

HasGrantedToV2 returns a boolean if a field has been set.

### SetGrantedToV2Nil

`func (o *MicrosoftGraphPermission) SetGrantedToV2Nil(b bool)`

 SetGrantedToV2Nil sets the value for GrantedToV2 to be an explicit nil

### UnsetGrantedToV2
`func (o *MicrosoftGraphPermission) UnsetGrantedToV2()`

UnsetGrantedToV2 ensures that no value is present for GrantedToV2, not even an explicit nil
### GetHasPassword

`func (o *MicrosoftGraphPermission) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *MicrosoftGraphPermission) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *MicrosoftGraphPermission) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *MicrosoftGraphPermission) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### SetHasPasswordNil

`func (o *MicrosoftGraphPermission) SetHasPasswordNil(b bool)`

 SetHasPasswordNil sets the value for HasPassword to be an explicit nil

### UnsetHasPassword
`func (o *MicrosoftGraphPermission) UnsetHasPassword()`

UnsetHasPassword ensures that no value is present for HasPassword, not even an explicit nil
### GetInheritedFrom

`func (o *MicrosoftGraphPermission) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *MicrosoftGraphPermission) GetInheritedFromOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFrom

`func (o *MicrosoftGraphPermission) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom sets InheritedFrom field to given value.

### HasInheritedFrom

`func (o *MicrosoftGraphPermission) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFromNil

`func (o *MicrosoftGraphPermission) SetInheritedFromNil(b bool)`

 SetInheritedFromNil sets the value for InheritedFrom to be an explicit nil

### UnsetInheritedFrom
`func (o *MicrosoftGraphPermission) UnsetInheritedFrom()`

UnsetInheritedFrom ensures that no value is present for InheritedFrom, not even an explicit nil
### GetInvitation

`func (o *MicrosoftGraphPermission) GetInvitation() AnyOfmicrosoftGraphSharingInvitation`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *MicrosoftGraphPermission) GetInvitationOk() (*AnyOfmicrosoftGraphSharingInvitation, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *MicrosoftGraphPermission) SetInvitation(v AnyOfmicrosoftGraphSharingInvitation)`

SetInvitation sets Invitation field to given value.

### HasInvitation

`func (o *MicrosoftGraphPermission) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### SetInvitationNil

`func (o *MicrosoftGraphPermission) SetInvitationNil(b bool)`

 SetInvitationNil sets the value for Invitation to be an explicit nil

### UnsetInvitation
`func (o *MicrosoftGraphPermission) UnsetInvitation()`

UnsetInvitation ensures that no value is present for Invitation, not even an explicit nil
### GetLink

`func (o *MicrosoftGraphPermission) GetLink() AnyOfmicrosoftGraphSharingLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MicrosoftGraphPermission) GetLinkOk() (*AnyOfmicrosoftGraphSharingLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *MicrosoftGraphPermission) SetLink(v AnyOfmicrosoftGraphSharingLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *MicrosoftGraphPermission) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *MicrosoftGraphPermission) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *MicrosoftGraphPermission) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetRoles

`func (o *MicrosoftGraphPermission) GetRoles() []*string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MicrosoftGraphPermission) GetRolesOk() (*[]*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MicrosoftGraphPermission) SetRoles(v []*string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *MicrosoftGraphPermission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetShareId

`func (o *MicrosoftGraphPermission) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *MicrosoftGraphPermission) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *MicrosoftGraphPermission) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *MicrosoftGraphPermission) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareIdNil

`func (o *MicrosoftGraphPermission) SetShareIdNil(b bool)`

 SetShareIdNil sets the value for ShareId to be an explicit nil

### UnsetShareId
`func (o *MicrosoftGraphPermission) UnsetShareId()`

UnsetShareId ensures that no value is present for ShareId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


