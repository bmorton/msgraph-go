# Drive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDrive

`func NewDrive() *Drive`

NewDrive instantiates a new Drive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveWithDefaults

`func NewDriveWithDefaults() *Drive`

NewDriveWithDefaults instantiates a new Drive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveType

`func (o *Drive) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *Drive) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *Drive) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *Drive) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### SetDriveTypeNil

`func (o *Drive) SetDriveTypeNil(b bool)`

 SetDriveTypeNil sets the value for DriveType to be an explicit nil

### UnsetDriveType
`func (o *Drive) UnsetDriveType()`

UnsetDriveType ensures that no value is present for DriveType, not even an explicit nil
### GetOwner

`func (o *Drive) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Drive) GetOwnerOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Drive) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Drive) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Drive) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Drive) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetQuota

`func (o *Drive) GetQuota() AnyOfmicrosoftGraphQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Drive) GetQuotaOk() (*AnyOfmicrosoftGraphQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Drive) SetQuota(v AnyOfmicrosoftGraphQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Drive) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *Drive) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *Drive) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetSharePointIds

`func (o *Drive) GetSharePointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharePointIds returns the SharePointIds field if non-nil, zero value otherwise.

### GetSharePointIdsOk

`func (o *Drive) GetSharePointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharePointIdsOk returns a tuple with the SharePointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePointIds

`func (o *Drive) SetSharePointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharePointIds sets SharePointIds field to given value.

### HasSharePointIds

`func (o *Drive) HasSharePointIds() bool`

HasSharePointIds returns a boolean if a field has been set.

### SetSharePointIdsNil

`func (o *Drive) SetSharePointIdsNil(b bool)`

 SetSharePointIdsNil sets the value for SharePointIds to be an explicit nil

### UnsetSharePointIds
`func (o *Drive) UnsetSharePointIds()`

UnsetSharePointIds ensures that no value is present for SharePointIds, not even an explicit nil
### GetSystem

`func (o *Drive) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Drive) GetSystemOk() (*AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Drive) SetSystem(v AnyOfobject)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Drive) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *Drive) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *Drive) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetBundles

`func (o *Drive) GetBundles() []MicrosoftGraphDriveItem`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *Drive) GetBundlesOk() (*[]MicrosoftGraphDriveItem, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *Drive) SetBundles(v []MicrosoftGraphDriveItem)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *Drive) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetFollowing

`func (o *Drive) GetFollowing() []MicrosoftGraphDriveItem`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *Drive) GetFollowingOk() (*[]MicrosoftGraphDriveItem, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *Drive) SetFollowing(v []MicrosoftGraphDriveItem)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *Drive) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetItems

`func (o *Drive) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Drive) GetItemsOk() (*[]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Drive) SetItems(v []MicrosoftGraphDriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Drive) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetList

`func (o *Drive) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *Drive) GetListOk() (*AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *Drive) SetList(v AnyOfmicrosoftGraphList)`

SetList sets List field to given value.

### HasList

`func (o *Drive) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *Drive) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *Drive) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetRoot

`func (o *Drive) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Drive) GetRootOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *Drive) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *Drive) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *Drive) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *Drive) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSpecial

`func (o *Drive) GetSpecial() []MicrosoftGraphDriveItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *Drive) GetSpecialOk() (*[]MicrosoftGraphDriveItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *Drive) SetSpecial(v []MicrosoftGraphDriveItem)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *Drive) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


