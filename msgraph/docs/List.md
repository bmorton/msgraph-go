# List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The displayable title of the list. | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphListInfo**](anyOf&lt;microsoft.graph.listInfo&gt;.md) | Provides additional details about the list. | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) | Returns identifiers useful for SharePoint REST compatibility. Read-only. | [optional] 
**System** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | If present, indicates that this is a system-managed list. Read-only. | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | The collection of field definitions for this list. | [optional] 
**ContentTypes** | Pointer to [**[]MicrosoftGraphContentType**](MicrosoftGraphContentType.md) | The collection of content types present in this list. | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) | Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem]. | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphListItem**](MicrosoftGraphListItem.md) | All items contained in the list. | [optional] 
**Subscriptions** | Pointer to [**[]MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | The set of subscriptions on the list. | [optional] 

## Methods

### NewList

`func NewList() *List`

NewList instantiates a new List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWithDefaults

`func NewListWithDefaults() *List`

NewListWithDefaults instantiates a new List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *List) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *List) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *List) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *List) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *List) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *List) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetList

`func (o *List) GetList() AnyOfmicrosoftGraphListInfo`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *List) GetListOk() (*AnyOfmicrosoftGraphListInfo, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *List) SetList(v AnyOfmicrosoftGraphListInfo)`

SetList sets List field to given value.

### HasList

`func (o *List) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *List) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *List) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetSharepointIds

`func (o *List) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *List) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *List) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *List) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *List) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *List) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSystem

`func (o *List) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *List) GetSystemOk() (*AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *List) SetSystem(v AnyOfobject)`

SetSystem sets System field to given value.

### HasSystem

`func (o *List) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *List) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *List) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetColumns

`func (o *List) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *List) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *List) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *List) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetContentTypes

`func (o *List) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *List) GetContentTypesOk() (*[]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *List) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *List) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetDrive

`func (o *List) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *List) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *List) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *List) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *List) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *List) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetItems

`func (o *List) GetItems() []MicrosoftGraphListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *List) GetItemsOk() (*[]MicrosoftGraphListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *List) SetItems(v []MicrosoftGraphListItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *List) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSubscriptions

`func (o *List) GetSubscriptions() []MicrosoftGraphSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *List) GetSubscriptionsOk() (*[]MicrosoftGraphSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *List) SetSubscriptions(v []MicrosoftGraphSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *List) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


