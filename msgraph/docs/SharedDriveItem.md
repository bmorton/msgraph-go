# SharedDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Information about the owner of the shared item being referenced. | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) | Used to access the underlying driveItem | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | All driveItems contained in the sharing root. This collection cannot be enumerated. | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphList**](anyOf&lt;microsoft.graph.list&gt;.md) | Used to access the underlying list | [optional] 
**ListItem** | Pointer to [**AnyOfmicrosoftGraphListItem**](anyOf&lt;microsoft.graph.listItem&gt;.md) | Used to access the underlying listItem | [optional] 
**Permission** | Pointer to [**AnyOfmicrosoftGraphPermission**](anyOf&lt;microsoft.graph.permission&gt;.md) | Used to access the permission representing the underlying sharing link | [optional] 
**Root** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) | Used to access the underlying driveItem. Deprecated -- use driveItem instead. | [optional] 
**Site** | Pointer to [**AnyOfmicrosoftGraphSite**](anyOf&lt;microsoft.graph.site&gt;.md) | Used to access the underlying site | [optional] 

## Methods

### NewSharedDriveItem

`func NewSharedDriveItem() *SharedDriveItem`

NewSharedDriveItem instantiates a new SharedDriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveItemWithDefaults

`func NewSharedDriveItemWithDefaults() *SharedDriveItem`

NewSharedDriveItemWithDefaults instantiates a new SharedDriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *SharedDriveItem) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SharedDriveItem) GetOwnerOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SharedDriveItem) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SharedDriveItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *SharedDriveItem) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *SharedDriveItem) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetDriveItem

`func (o *SharedDriveItem) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *SharedDriveItem) GetDriveItemOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveItem

`func (o *SharedDriveItem) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem sets DriveItem field to given value.

### HasDriveItem

`func (o *SharedDriveItem) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItemNil

`func (o *SharedDriveItem) SetDriveItemNil(b bool)`

 SetDriveItemNil sets the value for DriveItem to be an explicit nil

### UnsetDriveItem
`func (o *SharedDriveItem) UnsetDriveItem()`

UnsetDriveItem ensures that no value is present for DriveItem, not even an explicit nil
### GetItems

`func (o *SharedDriveItem) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SharedDriveItem) GetItemsOk() (*[]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SharedDriveItem) SetItems(v []MicrosoftGraphDriveItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *SharedDriveItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetList

`func (o *SharedDriveItem) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *SharedDriveItem) GetListOk() (*AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *SharedDriveItem) SetList(v AnyOfmicrosoftGraphList)`

SetList sets List field to given value.

### HasList

`func (o *SharedDriveItem) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *SharedDriveItem) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *SharedDriveItem) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetListItem

`func (o *SharedDriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *SharedDriveItem) GetListItemOk() (*AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItem

`func (o *SharedDriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem sets ListItem field to given value.

### HasListItem

`func (o *SharedDriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItemNil

`func (o *SharedDriveItem) SetListItemNil(b bool)`

 SetListItemNil sets the value for ListItem to be an explicit nil

### UnsetListItem
`func (o *SharedDriveItem) UnsetListItem()`

UnsetListItem ensures that no value is present for ListItem, not even an explicit nil
### GetPermission

`func (o *SharedDriveItem) GetPermission() AnyOfmicrosoftGraphPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *SharedDriveItem) GetPermissionOk() (*AnyOfmicrosoftGraphPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *SharedDriveItem) SetPermission(v AnyOfmicrosoftGraphPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *SharedDriveItem) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *SharedDriveItem) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *SharedDriveItem) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetRoot

`func (o *SharedDriveItem) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *SharedDriveItem) GetRootOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *SharedDriveItem) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *SharedDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *SharedDriveItem) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *SharedDriveItem) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSite

`func (o *SharedDriveItem) GetSite() AnyOfmicrosoftGraphSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SharedDriveItem) GetSiteOk() (*AnyOfmicrosoftGraphSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SharedDriveItem) SetSite(v AnyOfmicrosoftGraphSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *SharedDriveItem) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *SharedDriveItem) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *SharedDriveItem) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


