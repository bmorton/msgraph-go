# MicrosoftGraphSite

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

### NewMicrosoftGraphSite

`func NewMicrosoftGraphSite() *MicrosoftGraphSite`

NewMicrosoftGraphSite instantiates a new MicrosoftGraphSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSiteWithDefaults

`func NewMicrosoftGraphSiteWithDefaults() *MicrosoftGraphSite`

NewMicrosoftGraphSiteWithDefaults instantiates a new MicrosoftGraphSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphSite) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphSite) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphSite) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphSite) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphSite) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphSite) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphSite) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSite) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSite) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphSite) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphSite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphSite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphSite) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphSite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphSite) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphSite) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetETag

`func (o *MicrosoftGraphSite) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphSite) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MicrosoftGraphSite) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MicrosoftGraphSite) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *MicrosoftGraphSite) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *MicrosoftGraphSite) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphSite) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSite) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSite) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphSite) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphSite) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphSite) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSite) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSite) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSite) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSite) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphSite) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphSite) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphSite) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentReference

`func (o *MicrosoftGraphSite) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphSite) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MicrosoftGraphSite) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MicrosoftGraphSite) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReferenceNil

`func (o *MicrosoftGraphSite) SetParentReferenceNil(b bool)`

 SetParentReferenceNil sets the value for ParentReference to be an explicit nil

### UnsetParentReference
`func (o *MicrosoftGraphSite) UnsetParentReference()`

UnsetParentReference ensures that no value is present for ParentReference, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphSite) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphSite) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphSite) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphSite) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphSite) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphSite) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetCreatedByUser

`func (o *MicrosoftGraphSite) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphSite) GetCreatedByUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphSite) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *MicrosoftGraphSite) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *MicrosoftGraphSite) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *MicrosoftGraphSite) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetLastModifiedByUser

`func (o *MicrosoftGraphSite) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphSite) GetLastModifiedByUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphSite) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *MicrosoftGraphSite) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUserNil

`func (o *MicrosoftGraphSite) SetLastModifiedByUserNil(b bool)`

 SetLastModifiedByUserNil sets the value for LastModifiedByUser to be an explicit nil

### UnsetLastModifiedByUser
`func (o *MicrosoftGraphSite) UnsetLastModifiedByUser()`

UnsetLastModifiedByUser ensures that no value is present for LastModifiedByUser, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphSite) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSite) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphSite) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphSite) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphSite) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphSite) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetError

`func (o *MicrosoftGraphSite) GetError() AnyOfmicrosoftGraphPublicError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphSite) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphSite) SetError(v AnyOfmicrosoftGraphPublicError)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphSite) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphSite) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphSite) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetRoot

`func (o *MicrosoftGraphSite) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphSite) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *MicrosoftGraphSite) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *MicrosoftGraphSite) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *MicrosoftGraphSite) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *MicrosoftGraphSite) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSharepointIds

`func (o *MicrosoftGraphSite) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphSite) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *MicrosoftGraphSite) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *MicrosoftGraphSite) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *MicrosoftGraphSite) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *MicrosoftGraphSite) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSiteCollection

`func (o *MicrosoftGraphSite) GetSiteCollection() AnyOfmicrosoftGraphSiteCollection`

GetSiteCollection returns the SiteCollection field if non-nil, zero value otherwise.

### GetSiteCollectionOk

`func (o *MicrosoftGraphSite) GetSiteCollectionOk() (*AnyOfmicrosoftGraphSiteCollection, bool)`

GetSiteCollectionOk returns a tuple with the SiteCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCollection

`func (o *MicrosoftGraphSite) SetSiteCollection(v AnyOfmicrosoftGraphSiteCollection)`

SetSiteCollection sets SiteCollection field to given value.

### HasSiteCollection

`func (o *MicrosoftGraphSite) HasSiteCollection() bool`

HasSiteCollection returns a boolean if a field has been set.

### SetSiteCollectionNil

`func (o *MicrosoftGraphSite) SetSiteCollectionNil(b bool)`

 SetSiteCollectionNil sets the value for SiteCollection to be an explicit nil

### UnsetSiteCollection
`func (o *MicrosoftGraphSite) UnsetSiteCollection()`

UnsetSiteCollection ensures that no value is present for SiteCollection, not even an explicit nil
### GetAnalytics

`func (o *MicrosoftGraphSite) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *MicrosoftGraphSite) GetAnalyticsOk() (*AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *MicrosoftGraphSite) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *MicrosoftGraphSite) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalyticsNil

`func (o *MicrosoftGraphSite) SetAnalyticsNil(b bool)`

 SetAnalyticsNil sets the value for Analytics to be an explicit nil

### UnsetAnalytics
`func (o *MicrosoftGraphSite) UnsetAnalytics()`

UnsetAnalytics ensures that no value is present for Analytics, not even an explicit nil
### GetColumns

`func (o *MicrosoftGraphSite) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MicrosoftGraphSite) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *MicrosoftGraphSite) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *MicrosoftGraphSite) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetContentTypes

`func (o *MicrosoftGraphSite) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *MicrosoftGraphSite) GetContentTypesOk() (*[]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *MicrosoftGraphSite) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *MicrosoftGraphSite) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetDrive

`func (o *MicrosoftGraphSite) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *MicrosoftGraphSite) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *MicrosoftGraphSite) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *MicrosoftGraphSite) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *MicrosoftGraphSite) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *MicrosoftGraphSite) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetDrives

`func (o *MicrosoftGraphSite) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *MicrosoftGraphSite) GetDrivesOk() (*[]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *MicrosoftGraphSite) SetDrives(v []MicrosoftGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *MicrosoftGraphSite) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetExternalColumns

`func (o *MicrosoftGraphSite) GetExternalColumns() []MicrosoftGraphColumnDefinition`

GetExternalColumns returns the ExternalColumns field if non-nil, zero value otherwise.

### GetExternalColumnsOk

`func (o *MicrosoftGraphSite) GetExternalColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetExternalColumnsOk returns a tuple with the ExternalColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalColumns

`func (o *MicrosoftGraphSite) SetExternalColumns(v []MicrosoftGraphColumnDefinition)`

SetExternalColumns sets ExternalColumns field to given value.

### HasExternalColumns

`func (o *MicrosoftGraphSite) HasExternalColumns() bool`

HasExternalColumns returns a boolean if a field has been set.

### GetItems

`func (o *MicrosoftGraphSite) GetItems() []MicrosoftGraphBaseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MicrosoftGraphSite) GetItemsOk() (*[]MicrosoftGraphBaseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MicrosoftGraphSite) SetItems(v []MicrosoftGraphBaseItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *MicrosoftGraphSite) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLists

`func (o *MicrosoftGraphSite) GetLists() []MicrosoftGraphList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *MicrosoftGraphSite) GetListsOk() (*[]MicrosoftGraphList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *MicrosoftGraphSite) SetLists(v []MicrosoftGraphList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *MicrosoftGraphSite) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetPermissions

`func (o *MicrosoftGraphSite) GetPermissions() []MicrosoftGraphPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MicrosoftGraphSite) GetPermissionsOk() (*[]MicrosoftGraphPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MicrosoftGraphSite) SetPermissions(v []MicrosoftGraphPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MicrosoftGraphSite) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSites

`func (o *MicrosoftGraphSite) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *MicrosoftGraphSite) GetSitesOk() (*[]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *MicrosoftGraphSite) SetSites(v []MicrosoftGraphSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *MicrosoftGraphSite) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetTermStore

`func (o *MicrosoftGraphSite) GetTermStore() AnyOfmicrosoftGraphTermStoreStore`

GetTermStore returns the TermStore field if non-nil, zero value otherwise.

### GetTermStoreOk

`func (o *MicrosoftGraphSite) GetTermStoreOk() (*AnyOfmicrosoftGraphTermStoreStore, bool)`

GetTermStoreOk returns a tuple with the TermStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermStore

`func (o *MicrosoftGraphSite) SetTermStore(v AnyOfmicrosoftGraphTermStoreStore)`

SetTermStore sets TermStore field to given value.

### HasTermStore

`func (o *MicrosoftGraphSite) HasTermStore() bool`

HasTermStore returns a boolean if a field has been set.

### SetTermStoreNil

`func (o *MicrosoftGraphSite) SetTermStoreNil(b bool)`

 SetTermStoreNil sets the value for TermStore to be an explicit nil

### UnsetTermStore
`func (o *MicrosoftGraphSite) UnsetTermStore()`

UnsetTermStore ensures that no value is present for TermStore, not even an explicit nil
### GetTermStores

`func (o *MicrosoftGraphSite) GetTermStores() []MicrosoftGraphTermStoreStore`

GetTermStores returns the TermStores field if non-nil, zero value otherwise.

### GetTermStoresOk

`func (o *MicrosoftGraphSite) GetTermStoresOk() (*[]MicrosoftGraphTermStoreStore, bool)`

GetTermStoresOk returns a tuple with the TermStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermStores

`func (o *MicrosoftGraphSite) SetTermStores(v []MicrosoftGraphTermStoreStore)`

SetTermStores sets TermStores field to given value.

### HasTermStores

`func (o *MicrosoftGraphSite) HasTermStores() bool`

HasTermStores returns a boolean if a field has been set.

### GetOnenote

`func (o *MicrosoftGraphSite) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *MicrosoftGraphSite) GetOnenoteOk() (*AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnenote

`func (o *MicrosoftGraphSite) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote sets Onenote field to given value.

### HasOnenote

`func (o *MicrosoftGraphSite) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenoteNil

`func (o *MicrosoftGraphSite) SetOnenoteNil(b bool)`

 SetOnenoteNil sets the value for Onenote to be an explicit nil

### UnsetOnenote
`func (o *MicrosoftGraphSite) UnsetOnenote()`

UnsetOnenote ensures that no value is present for Onenote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


