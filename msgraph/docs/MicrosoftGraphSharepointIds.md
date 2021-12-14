# MicrosoftGraphSharepointIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | Pointer to **NullableString** | The unique identifier (guid) for the item&#39;s list in SharePoint. | [optional] 
**ListItemId** | Pointer to **NullableString** | An integer identifier for the item within the containing list. | [optional] 
**ListItemUniqueId** | Pointer to **NullableString** | The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site. | [optional] 
**SiteId** | Pointer to **NullableString** | The unique identifier (guid) for the item&#39;s site collection (SPSite). | [optional] 
**SiteUrl** | Pointer to **NullableString** | The SharePoint URL for the site that contains the item. | [optional] 
**TenantId** | Pointer to **NullableString** | The unique identifier (guid) for the tenancy. | [optional] 
**WebId** | Pointer to **NullableString** | The unique identifier (guid) for the item&#39;s site (SPWeb). | [optional] 

## Methods

### NewMicrosoftGraphSharepointIds

`func NewMicrosoftGraphSharepointIds() *MicrosoftGraphSharepointIds`

NewMicrosoftGraphSharepointIds instantiates a new MicrosoftGraphSharepointIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharepointIdsWithDefaults

`func NewMicrosoftGraphSharepointIdsWithDefaults() *MicrosoftGraphSharepointIds`

NewMicrosoftGraphSharepointIdsWithDefaults instantiates a new MicrosoftGraphSharepointIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListId

`func (o *MicrosoftGraphSharepointIds) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *MicrosoftGraphSharepointIds) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *MicrosoftGraphSharepointIds) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *MicrosoftGraphSharepointIds) HasListId() bool`

HasListId returns a boolean if a field has been set.

### SetListIdNil

`func (o *MicrosoftGraphSharepointIds) SetListIdNil(b bool)`

 SetListIdNil sets the value for ListId to be an explicit nil

### UnsetListId
`func (o *MicrosoftGraphSharepointIds) UnsetListId()`

UnsetListId ensures that no value is present for ListId, not even an explicit nil
### GetListItemId

`func (o *MicrosoftGraphSharepointIds) GetListItemId() string`

GetListItemId returns the ListItemId field if non-nil, zero value otherwise.

### GetListItemIdOk

`func (o *MicrosoftGraphSharepointIds) GetListItemIdOk() (*string, bool)`

GetListItemIdOk returns a tuple with the ListItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItemId

`func (o *MicrosoftGraphSharepointIds) SetListItemId(v string)`

SetListItemId sets ListItemId field to given value.

### HasListItemId

`func (o *MicrosoftGraphSharepointIds) HasListItemId() bool`

HasListItemId returns a boolean if a field has been set.

### SetListItemIdNil

`func (o *MicrosoftGraphSharepointIds) SetListItemIdNil(b bool)`

 SetListItemIdNil sets the value for ListItemId to be an explicit nil

### UnsetListItemId
`func (o *MicrosoftGraphSharepointIds) UnsetListItemId()`

UnsetListItemId ensures that no value is present for ListItemId, not even an explicit nil
### GetListItemUniqueId

`func (o *MicrosoftGraphSharepointIds) GetListItemUniqueId() string`

GetListItemUniqueId returns the ListItemUniqueId field if non-nil, zero value otherwise.

### GetListItemUniqueIdOk

`func (o *MicrosoftGraphSharepointIds) GetListItemUniqueIdOk() (*string, bool)`

GetListItemUniqueIdOk returns a tuple with the ListItemUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItemUniqueId

`func (o *MicrosoftGraphSharepointIds) SetListItemUniqueId(v string)`

SetListItemUniqueId sets ListItemUniqueId field to given value.

### HasListItemUniqueId

`func (o *MicrosoftGraphSharepointIds) HasListItemUniqueId() bool`

HasListItemUniqueId returns a boolean if a field has been set.

### SetListItemUniqueIdNil

`func (o *MicrosoftGraphSharepointIds) SetListItemUniqueIdNil(b bool)`

 SetListItemUniqueIdNil sets the value for ListItemUniqueId to be an explicit nil

### UnsetListItemUniqueId
`func (o *MicrosoftGraphSharepointIds) UnsetListItemUniqueId()`

UnsetListItemUniqueId ensures that no value is present for ListItemUniqueId, not even an explicit nil
### GetSiteId

`func (o *MicrosoftGraphSharepointIds) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MicrosoftGraphSharepointIds) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MicrosoftGraphSharepointIds) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MicrosoftGraphSharepointIds) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *MicrosoftGraphSharepointIds) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *MicrosoftGraphSharepointIds) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetSiteUrl

`func (o *MicrosoftGraphSharepointIds) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *MicrosoftGraphSharepointIds) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *MicrosoftGraphSharepointIds) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.

### HasSiteUrl

`func (o *MicrosoftGraphSharepointIds) HasSiteUrl() bool`

HasSiteUrl returns a boolean if a field has been set.

### SetSiteUrlNil

`func (o *MicrosoftGraphSharepointIds) SetSiteUrlNil(b bool)`

 SetSiteUrlNil sets the value for SiteUrl to be an explicit nil

### UnsetSiteUrl
`func (o *MicrosoftGraphSharepointIds) UnsetSiteUrl()`

UnsetSiteUrl ensures that no value is present for SiteUrl, not even an explicit nil
### GetTenantId

`func (o *MicrosoftGraphSharepointIds) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftGraphSharepointIds) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftGraphSharepointIds) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MicrosoftGraphSharepointIds) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MicrosoftGraphSharepointIds) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MicrosoftGraphSharepointIds) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWebId

`func (o *MicrosoftGraphSharepointIds) GetWebId() string`

GetWebId returns the WebId field if non-nil, zero value otherwise.

### GetWebIdOk

`func (o *MicrosoftGraphSharepointIds) GetWebIdOk() (*string, bool)`

GetWebIdOk returns a tuple with the WebId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebId

`func (o *MicrosoftGraphSharepointIds) SetWebId(v string)`

SetWebId sets WebId field to given value.

### HasWebId

`func (o *MicrosoftGraphSharepointIds) HasWebId() bool`

HasWebId returns a boolean if a field has been set.

### SetWebIdNil

`func (o *MicrosoftGraphSharepointIds) SetWebIdNil(b bool)`

 SetWebIdNil sets the value for WebId to be an explicit nil

### UnsetWebId
`func (o *MicrosoftGraphSharepointIds) UnsetWebId()`

UnsetWebId ensures that no value is present for WebId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


