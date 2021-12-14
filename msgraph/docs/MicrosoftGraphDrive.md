# MicrosoftGraphDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user, device, or application which created the item. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of item creation. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Provides a user-visible description of the item. Optional. | [optional] 
**ETag** | Pointer to **NullableString** | ETag for the item. Read-only. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user, device, and application which last modified the item. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Date and time the item was last modified. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | The name of the item. Read-write. | [optional] 
**ParentReference** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) | Parent information, if the item has a parent. Read-write. | [optional] 
**WebUrl** | Pointer to **NullableString** | URL that displays the resource in the browser. Read-only. | [optional] 
**CreatedByUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) | Identity of the user who created the item. Read-only. | [optional] 
**LastModifiedByUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) | Identity of the user who last modified the item. Read-only. | [optional] 
**DriveType** | Pointer to **NullableString** | Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only. | [optional] 
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Optional. The user account that owns the drive. Read-only. | [optional] 
**Quota** | Pointer to [**AnyOfmicrosoftGraphQuota**](anyOf&lt;microsoft.graph.quota&gt;.md) | Optional. Information about the drive&#39;s storage space quota. Read-only. | [optional] 
**SharePointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**System** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | If present, indicates that this is a system-managed drive. Read-only. | [optional] 
**Bundles** | Pointer to [**[]MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive. | [optional] 
**Following** | Pointer to [**[]MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | The list of items the user is following. Only in OneDrive for Business. | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | All items contained in the drive. Read-only. Nullable. | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphList**](anyOf&lt;microsoft.graph.list&gt;.md) | For drives in SharePoint, the underlying document library list. Read-only. Nullable. | [optional] 
**Root** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) | The root folder of the drive. Read-only. | [optional] 
**Special** | Pointer to [**[]MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | Collection of common folders available in OneDrive. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphDrive

`func NewMicrosoftGraphDrive() *MicrosoftGraphDrive`

NewMicrosoftGraphDrive instantiates a new MicrosoftGraphDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDriveWithDefaults

`func NewMicrosoftGraphDriveWithDefaults() *MicrosoftGraphDrive`

NewMicrosoftGraphDriveWithDefaults instantiates a new MicrosoftGraphDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDrive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDrive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDrive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDrive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphDrive) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphDrive) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphDrive) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphDrive) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphDrive) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphDrive) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphDrive) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDrive) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDrive) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphDrive) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphDrive) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDrive) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphDrive) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphDrive) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphDrive) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphDrive) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetETag

`func (o *MicrosoftGraphDrive) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphDrive) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MicrosoftGraphDrive) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MicrosoftGraphDrive) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *MicrosoftGraphDrive) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *MicrosoftGraphDrive) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphDrive) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphDrive) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphDrive) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphDrive) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphDrive) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphDrive) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphDrive) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDrive) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDrive) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDrive) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphDrive) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphDrive) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphDrive) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphDrive) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentReference

`func (o *MicrosoftGraphDrive) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphDrive) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MicrosoftGraphDrive) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MicrosoftGraphDrive) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReferenceNil

`func (o *MicrosoftGraphDrive) SetParentReferenceNil(b bool)`

 SetParentReferenceNil sets the value for ParentReference to be an explicit nil

### UnsetParentReference
`func (o *MicrosoftGraphDrive) UnsetParentReference()`

UnsetParentReference ensures that no value is present for ParentReference, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphDrive) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphDrive) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphDrive) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphDrive) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphDrive) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphDrive) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetCreatedByUser

`func (o *MicrosoftGraphDrive) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphDrive) GetCreatedByUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphDrive) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *MicrosoftGraphDrive) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *MicrosoftGraphDrive) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *MicrosoftGraphDrive) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetLastModifiedByUser

`func (o *MicrosoftGraphDrive) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphDrive) GetLastModifiedByUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphDrive) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *MicrosoftGraphDrive) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUserNil

`func (o *MicrosoftGraphDrive) SetLastModifiedByUserNil(b bool)`

 SetLastModifiedByUserNil sets the value for LastModifiedByUser to be an explicit nil

### UnsetLastModifiedByUser
`func (o *MicrosoftGraphDrive) UnsetLastModifiedByUser()`

UnsetLastModifiedByUser ensures that no value is present for LastModifiedByUser, not even an explicit nil
### GetDriveType

`func (o *MicrosoftGraphDrive) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *MicrosoftGraphDrive) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *MicrosoftGraphDrive) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *MicrosoftGraphDrive) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### SetDriveTypeNil

`func (o *MicrosoftGraphDrive) SetDriveTypeNil(b bool)`

 SetDriveTypeNil sets the value for DriveType to be an explicit nil

### UnsetDriveType
`func (o *MicrosoftGraphDrive) UnsetDriveType()`

UnsetDriveType ensures that no value is present for DriveType, not even an explicit nil
### GetOwner

`func (o *MicrosoftGraphDrive) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphDrive) GetOwnerOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MicrosoftGraphDrive) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MicrosoftGraphDrive) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *MicrosoftGraphDrive) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *MicrosoftGraphDrive) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetQuota

`func (o *MicrosoftGraphDrive) GetQuota() AnyOfmicrosoftGraphQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *MicrosoftGraphDrive) GetQuotaOk() (*AnyOfmicrosoftGraphQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *MicrosoftGraphDrive) SetQuota(v AnyOfmicrosoftGraphQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *MicrosoftGraphDrive) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *MicrosoftGraphDrive) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *MicrosoftGraphDrive) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetSharePointIds

`func (o *MicrosoftGraphDrive) GetSharePointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharePointIds returns the SharePointIds field if non-nil, zero value otherwise.

### GetSharePointIdsOk

`func (o *MicrosoftGraphDrive) GetSharePointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharePointIdsOk returns a tuple with the SharePointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePointIds

`func (o *MicrosoftGraphDrive) SetSharePointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharePointIds sets SharePointIds field to given value.

### HasSharePointIds

`func (o *MicrosoftGraphDrive) HasSharePointIds() bool`

HasSharePointIds returns a boolean if a field has been set.

### SetSharePointIdsNil

`func (o *MicrosoftGraphDrive) SetSharePointIdsNil(b bool)`

 SetSharePointIdsNil sets the value for SharePointIds to be an explicit nil

### UnsetSharePointIds
`func (o *MicrosoftGraphDrive) UnsetSharePointIds()`

UnsetSharePointIds ensures that no value is present for SharePointIds, not even an explicit nil
### GetSystem

`func (o *MicrosoftGraphDrive) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *MicrosoftGraphDrive) GetSystemOk() (*AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *MicrosoftGraphDrive) SetSystem(v AnyOfobject)`

SetSystem sets System field to given value.

### HasSystem

`func (o *MicrosoftGraphDrive) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *MicrosoftGraphDrive) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *MicrosoftGraphDrive) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetBundles

`func (o *MicrosoftGraphDrive) GetBundles() []MicrosoftGraphDriveItem`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *MicrosoftGraphDrive) GetBundlesOk() (*[]MicrosoftGraphDriveItem, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *MicrosoftGraphDrive) SetBundles(v []MicrosoftGraphDriveItem)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *MicrosoftGraphDrive) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetFollowing

`func (o *MicrosoftGraphDrive) GetFollowing() []MicrosoftGraphDriveItem`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *MicrosoftGraphDrive) GetFollowingOk() (*[]MicrosoftGraphDriveItem, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *MicrosoftGraphDrive) SetFollowing(v []MicrosoftGraphDriveItem)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *MicrosoftGraphDrive) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetItems

`func (o *MicrosoftGraphDrive) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MicrosoftGraphDrive) GetItemsOk() (*[]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MicrosoftGraphDrive) SetItems(v []MicrosoftGraphDriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *MicrosoftGraphDrive) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetList

`func (o *MicrosoftGraphDrive) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MicrosoftGraphDrive) GetListOk() (*AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MicrosoftGraphDrive) SetList(v AnyOfmicrosoftGraphList)`

SetList sets List field to given value.

### HasList

`func (o *MicrosoftGraphDrive) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *MicrosoftGraphDrive) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *MicrosoftGraphDrive) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetRoot

`func (o *MicrosoftGraphDrive) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphDrive) GetRootOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *MicrosoftGraphDrive) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *MicrosoftGraphDrive) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *MicrosoftGraphDrive) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *MicrosoftGraphDrive) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSpecial

`func (o *MicrosoftGraphDrive) GetSpecial() []MicrosoftGraphDriveItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *MicrosoftGraphDrive) GetSpecialOk() (*[]MicrosoftGraphDriveItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *MicrosoftGraphDrive) SetSpecial(v []MicrosoftGraphDriveItem)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *MicrosoftGraphDrive) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


