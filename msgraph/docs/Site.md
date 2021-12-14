# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The full title for the site. Read-only. | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphPublicError**](anyOf&lt;microsoft.graph.publicError&gt;.md) |  | [optional] 
**Root** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | If present, indicates that this is the root site in the site collection. Read-only. | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) | Returns identifiers useful for SharePoint REST compatibility. Read-only. | [optional] 
**SiteCollection** | Pointer to [**AnyOfmicrosoftGraphSiteCollection**](anyOf&lt;microsoft.graph.siteCollection&gt;.md) | Provides details about the site&#39;s site collection. Available only on the root site. Read-only. | [optional] 
**Analytics** | Pointer to [**AnyOfmicrosoftGraphItemAnalytics**](anyOf&lt;microsoft.graph.itemAnalytics&gt;.md) | Analytics about the view activities that took place in this site. | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | The collection of column definitions reusable across lists under this site. | [optional] 
**ContentTypes** | Pointer to [**[]MicrosoftGraphContentType**](MicrosoftGraphContentType.md) | The collection of content types defined for this site. | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) | The default drive (document library) for this site. | [optional] 
**Drives** | Pointer to [**[]MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | The collection of drives (document libraries) under this site. | [optional] 
**ExternalColumns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | The collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site. | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md) | Used to address any item contained in this site. This collection can&#39;t be enumerated. | [optional] 
**Lists** | Pointer to [**[]MicrosoftGraphList**](MicrosoftGraphList.md) | The collection of lists under this site. | [optional] 
**Permissions** | Pointer to [**[]MicrosoftGraphPermission**](MicrosoftGraphPermission.md) | The permissions associated with the site. Nullable. | [optional] 
**Sites** | Pointer to [**[]MicrosoftGraphSite**](MicrosoftGraphSite.md) | The collection of the sub-sites under this site. | [optional] 
**TermStore** | Pointer to [**AnyOfmicrosoftGraphTermStoreStore**](anyOf&lt;microsoft.graph.termStore.store&gt;.md) | The default termStore under this site. | [optional] 
**TermStores** | Pointer to [**[]MicrosoftGraphTermStoreStore**](MicrosoftGraphTermStoreStore.md) | The collection of termStores under this site. | [optional] 
**Onenote** | Pointer to [**AnyOfmicrosoftGraphOnenote**](anyOf&lt;microsoft.graph.onenote&gt;.md) | Calls the OneNote service for notebook related operations. | [optional] 

## Methods

### NewSite

`func NewSite() *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Site) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Site) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Site) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Site) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Site) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Site) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetError

`func (o *Site) GetError() AnyOfmicrosoftGraphPublicError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Site) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Site) SetError(v AnyOfmicrosoftGraphPublicError)`

SetError sets Error field to given value.

### HasError

`func (o *Site) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Site) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Site) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetRoot

`func (o *Site) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Site) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *Site) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *Site) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *Site) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *Site) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSharepointIds

`func (o *Site) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *Site) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *Site) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *Site) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *Site) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *Site) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSiteCollection

`func (o *Site) GetSiteCollection() AnyOfmicrosoftGraphSiteCollection`

GetSiteCollection returns the SiteCollection field if non-nil, zero value otherwise.

### GetSiteCollectionOk

`func (o *Site) GetSiteCollectionOk() (*AnyOfmicrosoftGraphSiteCollection, bool)`

GetSiteCollectionOk returns a tuple with the SiteCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCollection

`func (o *Site) SetSiteCollection(v AnyOfmicrosoftGraphSiteCollection)`

SetSiteCollection sets SiteCollection field to given value.

### HasSiteCollection

`func (o *Site) HasSiteCollection() bool`

HasSiteCollection returns a boolean if a field has been set.

### SetSiteCollectionNil

`func (o *Site) SetSiteCollectionNil(b bool)`

 SetSiteCollectionNil sets the value for SiteCollection to be an explicit nil

### UnsetSiteCollection
`func (o *Site) UnsetSiteCollection()`

UnsetSiteCollection ensures that no value is present for SiteCollection, not even an explicit nil
### GetAnalytics

`func (o *Site) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *Site) GetAnalyticsOk() (*AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *Site) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *Site) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalyticsNil

`func (o *Site) SetAnalyticsNil(b bool)`

 SetAnalyticsNil sets the value for Analytics to be an explicit nil

### UnsetAnalytics
`func (o *Site) UnsetAnalytics()`

UnsetAnalytics ensures that no value is present for Analytics, not even an explicit nil
### GetColumns

`func (o *Site) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Site) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Site) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Site) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetContentTypes

`func (o *Site) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *Site) GetContentTypesOk() (*[]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *Site) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *Site) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetDrive

`func (o *Site) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *Site) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *Site) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *Site) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *Site) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *Site) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetDrives

`func (o *Site) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *Site) GetDrivesOk() (*[]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *Site) SetDrives(v []MicrosoftGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *Site) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetExternalColumns

`func (o *Site) GetExternalColumns() []MicrosoftGraphColumnDefinition`

GetExternalColumns returns the ExternalColumns field if non-nil, zero value otherwise.

### GetExternalColumnsOk

`func (o *Site) GetExternalColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetExternalColumnsOk returns a tuple with the ExternalColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalColumns

`func (o *Site) SetExternalColumns(v []MicrosoftGraphColumnDefinition)`

SetExternalColumns sets ExternalColumns field to given value.

### HasExternalColumns

`func (o *Site) HasExternalColumns() bool`

HasExternalColumns returns a boolean if a field has been set.

### GetItems

`func (o *Site) GetItems() []MicrosoftGraphBaseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Site) GetItemsOk() (*[]MicrosoftGraphBaseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Site) SetItems(v []MicrosoftGraphBaseItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Site) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLists

`func (o *Site) GetLists() []MicrosoftGraphList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Site) GetListsOk() (*[]MicrosoftGraphList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *Site) SetLists(v []MicrosoftGraphList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *Site) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetPermissions

`func (o *Site) GetPermissions() []MicrosoftGraphPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Site) GetPermissionsOk() (*[]MicrosoftGraphPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Site) SetPermissions(v []MicrosoftGraphPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Site) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSites

`func (o *Site) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Site) GetSitesOk() (*[]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *Site) SetSites(v []MicrosoftGraphSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *Site) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetTermStore

`func (o *Site) GetTermStore() AnyOfmicrosoftGraphTermStoreStore`

GetTermStore returns the TermStore field if non-nil, zero value otherwise.

### GetTermStoreOk

`func (o *Site) GetTermStoreOk() (*AnyOfmicrosoftGraphTermStoreStore, bool)`

GetTermStoreOk returns a tuple with the TermStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermStore

`func (o *Site) SetTermStore(v AnyOfmicrosoftGraphTermStoreStore)`

SetTermStore sets TermStore field to given value.

### HasTermStore

`func (o *Site) HasTermStore() bool`

HasTermStore returns a boolean if a field has been set.

### SetTermStoreNil

`func (o *Site) SetTermStoreNil(b bool)`

 SetTermStoreNil sets the value for TermStore to be an explicit nil

### UnsetTermStore
`func (o *Site) UnsetTermStore()`

UnsetTermStore ensures that no value is present for TermStore, not even an explicit nil
### GetTermStores

`func (o *Site) GetTermStores() []MicrosoftGraphTermStoreStore`

GetTermStores returns the TermStores field if non-nil, zero value otherwise.

### GetTermStoresOk

`func (o *Site) GetTermStoresOk() (*[]MicrosoftGraphTermStoreStore, bool)`

GetTermStoresOk returns a tuple with the TermStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermStores

`func (o *Site) SetTermStores(v []MicrosoftGraphTermStoreStore)`

SetTermStores sets TermStores field to given value.

### HasTermStores

`func (o *Site) HasTermStores() bool`

HasTermStores returns a boolean if a field has been set.

### GetOnenote

`func (o *Site) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *Site) GetOnenoteOk() (*AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnenote

`func (o *Site) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote sets Onenote field to given value.

### HasOnenote

`func (o *Site) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenoteNil

`func (o *Site) SetOnenoteNil(b bool)`

 SetOnenoteNil sets the value for Onenote to be an explicit nil

### UnsetOnenote
`func (o *Site) UnsetOnenote()`

UnsetOnenote ensures that no value is present for Onenote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


