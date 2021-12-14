# MicrosoftGraphDriveItem

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
**Audio** | Pointer to [**AnyOfmicrosoftGraphAudio**](anyOf&lt;microsoft.graph.audio&gt;.md) | Audio metadata, if the item is an audio file. Read-only. | [optional] 
**Bundle** | Pointer to [**AnyOfmicrosoftGraphBundle**](anyOf&lt;microsoft.graph.bundle&gt;.md) |  | [optional] 
**Content** | Pointer to **NullableString** | The content stream, if the item represents a file. | [optional] 
**CTag** | Pointer to **NullableString** | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. | [optional] 
**Deleted** | Pointer to [**AnyOfmicrosoftGraphDeleted**](anyOf&lt;microsoft.graph.deleted&gt;.md) | Information about the deleted state of the item. Read-only. | [optional] 
**File** | Pointer to [**AnyOfmicrosoftGraphFile**](anyOf&lt;microsoft.graph.file&gt;.md) | File metadata, if the item is a file. Read-only. | [optional] 
**FileSystemInfo** | Pointer to [**AnyOfmicrosoftGraphFileSystemInfo**](anyOf&lt;microsoft.graph.fileSystemInfo&gt;.md) | File system information on client. Read-write. | [optional] 
**Folder** | Pointer to [**AnyOfmicrosoftGraphFolder**](anyOf&lt;microsoft.graph.folder&gt;.md) | Folder metadata, if the item is a folder. Read-only. | [optional] 
**Image** | Pointer to [**AnyOfmicrosoftGraphImage**](anyOf&lt;microsoft.graph.image&gt;.md) | Image metadata, if the item is an image. Read-only. | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphGeoCoordinates**](anyOf&lt;microsoft.graph.geoCoordinates&gt;.md) | Location metadata, if the item has location data. Read-only. | [optional] 
**Malware** | Pointer to [**AnyOfmicrosoftGraphMalware**](anyOf&lt;microsoft.graph.malware&gt;.md) | Malware metadata, if the item was detected to contain malware. Read-only. | [optional] 
**Package** | Pointer to [**AnyOfmicrosoftGraphPackage**](anyOf&lt;microsoft.graph.package&gt;.md) | If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only. | [optional] 
**PendingOperations** | Pointer to [**AnyOfmicrosoftGraphPendingOperations**](anyOf&lt;microsoft.graph.pendingOperations&gt;.md) | If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only. | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphPhoto**](anyOf&lt;microsoft.graph.photo&gt;.md) | Photo metadata, if the item is a photo. Read-only. | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) | Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only. | [optional] 
**RemoteItem** | Pointer to [**AnyOfmicrosoftGraphRemoteItem**](anyOf&lt;microsoft.graph.remoteItem&gt;.md) | Remote item data, if the item is shared from a drive other than the one being accessed. Read-only. | [optional] 
**Root** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive. | [optional] 
**SearchResult** | Pointer to [**AnyOfmicrosoftGraphSearchResult**](anyOf&lt;microsoft.graph.searchResult&gt;.md) | Search metadata, if the item is from a search result. Read-only. | [optional] 
**Shared** | Pointer to [**AnyOfmicrosoftGraphShared**](anyOf&lt;microsoft.graph.shared&gt;.md) | Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only. | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) | Returns identifiers useful for SharePoint REST compatibility. Read-only. | [optional] 
**Size** | Pointer to **NullableInt64** | Size of the item in bytes. Read-only. | [optional] 
**SpecialFolder** | Pointer to [**AnyOfmicrosoftGraphSpecialFolder**](anyOf&lt;microsoft.graph.specialFolder&gt;.md) | If the current item is also available as a special folder, this facet is returned. Read-only. | [optional] 
**Video** | Pointer to [**AnyOfmicrosoftGraphVideo**](anyOf&lt;microsoft.graph.video&gt;.md) | Video metadata, if the item is a video. Read-only. | [optional] 
**WebDavUrl** | Pointer to **NullableString** | WebDAV compatible URL for the item. | [optional] 
**Workbook** | Pointer to [**AnyOfmicrosoftGraphWorkbook**](anyOf&lt;microsoft.graph.workbook&gt;.md) | For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet&#39;s contents. Nullable. | [optional] 
**Analytics** | Pointer to [**AnyOfmicrosoftGraphItemAnalytics**](anyOf&lt;microsoft.graph.itemAnalytics&gt;.md) | Analytics about the view activities that took place on this item. | [optional] 
**Children** | Pointer to [**[]MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable. | [optional] 
**ListItem** | Pointer to [**AnyOfmicrosoftGraphListItem**](anyOf&lt;microsoft.graph.listItem&gt;.md) | For drives in SharePoint, the associated document library list item. Read-only. Nullable. | [optional] 
**Permissions** | Pointer to [**[]MicrosoftGraphPermission**](MicrosoftGraphPermission.md) | The set of permissions for the item. Read-only. Nullable. | [optional] 
**Subscriptions** | Pointer to [**[]MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | The set of subscriptions on the item. Only supported on the root of a drive. | [optional] 
**Thumbnails** | Pointer to [**[]MicrosoftGraphThumbnailSet**](MicrosoftGraphThumbnailSet.md) | Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable. | [optional] 
**Versions** | Pointer to [**[]MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md) | The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphDriveItem

`func NewMicrosoftGraphDriveItem() *MicrosoftGraphDriveItem`

NewMicrosoftGraphDriveItem instantiates a new MicrosoftGraphDriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDriveItemWithDefaults

`func NewMicrosoftGraphDriveItemWithDefaults() *MicrosoftGraphDriveItem`

NewMicrosoftGraphDriveItemWithDefaults instantiates a new MicrosoftGraphDriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDriveItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDriveItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDriveItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDriveItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphDriveItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphDriveItem) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphDriveItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphDriveItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphDriveItem) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphDriveItem) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphDriveItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDriveItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDriveItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphDriveItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphDriveItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDriveItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphDriveItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphDriveItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphDriveItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphDriveItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetETag

`func (o *MicrosoftGraphDriveItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphDriveItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MicrosoftGraphDriveItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MicrosoftGraphDriveItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *MicrosoftGraphDriveItem) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *MicrosoftGraphDriveItem) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphDriveItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphDriveItem) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphDriveItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphDriveItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphDriveItem) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphDriveItem) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDriveItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDriveItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphDriveItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphDriveItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphDriveItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphDriveItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphDriveItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphDriveItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentReference

`func (o *MicrosoftGraphDriveItem) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphDriveItem) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MicrosoftGraphDriveItem) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MicrosoftGraphDriveItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReferenceNil

`func (o *MicrosoftGraphDriveItem) SetParentReferenceNil(b bool)`

 SetParentReferenceNil sets the value for ParentReference to be an explicit nil

### UnsetParentReference
`func (o *MicrosoftGraphDriveItem) UnsetParentReference()`

UnsetParentReference ensures that no value is present for ParentReference, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphDriveItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphDriveItem) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphDriveItem) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphDriveItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphDriveItem) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphDriveItem) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetCreatedByUser

`func (o *MicrosoftGraphDriveItem) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphDriveItem) GetCreatedByUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphDriveItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *MicrosoftGraphDriveItem) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUserNil

`func (o *MicrosoftGraphDriveItem) SetCreatedByUserNil(b bool)`

 SetCreatedByUserNil sets the value for CreatedByUser to be an explicit nil

### UnsetCreatedByUser
`func (o *MicrosoftGraphDriveItem) UnsetCreatedByUser()`

UnsetCreatedByUser ensures that no value is present for CreatedByUser, not even an explicit nil
### GetLastModifiedByUser

`func (o *MicrosoftGraphDriveItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphDriveItem) GetLastModifiedByUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphDriveItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *MicrosoftGraphDriveItem) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUserNil

`func (o *MicrosoftGraphDriveItem) SetLastModifiedByUserNil(b bool)`

 SetLastModifiedByUserNil sets the value for LastModifiedByUser to be an explicit nil

### UnsetLastModifiedByUser
`func (o *MicrosoftGraphDriveItem) UnsetLastModifiedByUser()`

UnsetLastModifiedByUser ensures that no value is present for LastModifiedByUser, not even an explicit nil
### GetAudio

`func (o *MicrosoftGraphDriveItem) GetAudio() AnyOfmicrosoftGraphAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *MicrosoftGraphDriveItem) GetAudioOk() (*AnyOfmicrosoftGraphAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *MicrosoftGraphDriveItem) SetAudio(v AnyOfmicrosoftGraphAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *MicrosoftGraphDriveItem) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *MicrosoftGraphDriveItem) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *MicrosoftGraphDriveItem) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetBundle

`func (o *MicrosoftGraphDriveItem) GetBundle() AnyOfmicrosoftGraphBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *MicrosoftGraphDriveItem) GetBundleOk() (*AnyOfmicrosoftGraphBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *MicrosoftGraphDriveItem) SetBundle(v AnyOfmicrosoftGraphBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *MicrosoftGraphDriveItem) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### SetBundleNil

`func (o *MicrosoftGraphDriveItem) SetBundleNil(b bool)`

 SetBundleNil sets the value for Bundle to be an explicit nil

### UnsetBundle
`func (o *MicrosoftGraphDriveItem) UnsetBundle()`

UnsetBundle ensures that no value is present for Bundle, not even an explicit nil
### GetContent

`func (o *MicrosoftGraphDriveItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphDriveItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphDriveItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphDriveItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphDriveItem) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphDriveItem) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCTag

`func (o *MicrosoftGraphDriveItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *MicrosoftGraphDriveItem) GetCTagOk() (*string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTag

`func (o *MicrosoftGraphDriveItem) SetCTag(v string)`

SetCTag sets CTag field to given value.

### HasCTag

`func (o *MicrosoftGraphDriveItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### SetCTagNil

`func (o *MicrosoftGraphDriveItem) SetCTagNil(b bool)`

 SetCTagNil sets the value for CTag to be an explicit nil

### UnsetCTag
`func (o *MicrosoftGraphDriveItem) UnsetCTag()`

UnsetCTag ensures that no value is present for CTag, not even an explicit nil
### GetDeleted

`func (o *MicrosoftGraphDriveItem) GetDeleted() AnyOfmicrosoftGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MicrosoftGraphDriveItem) GetDeletedOk() (*AnyOfmicrosoftGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MicrosoftGraphDriveItem) SetDeleted(v AnyOfmicrosoftGraphDeleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MicrosoftGraphDriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *MicrosoftGraphDriveItem) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *MicrosoftGraphDriveItem) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetFile

`func (o *MicrosoftGraphDriveItem) GetFile() AnyOfmicrosoftGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MicrosoftGraphDriveItem) GetFileOk() (*AnyOfmicrosoftGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MicrosoftGraphDriveItem) SetFile(v AnyOfmicrosoftGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *MicrosoftGraphDriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *MicrosoftGraphDriveItem) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *MicrosoftGraphDriveItem) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFileSystemInfo

`func (o *MicrosoftGraphDriveItem) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *MicrosoftGraphDriveItem) GetFileSystemInfoOk() (*AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *MicrosoftGraphDriveItem) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *MicrosoftGraphDriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfoNil

`func (o *MicrosoftGraphDriveItem) SetFileSystemInfoNil(b bool)`

 SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil

### UnsetFileSystemInfo
`func (o *MicrosoftGraphDriveItem) UnsetFileSystemInfo()`

UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
### GetFolder

`func (o *MicrosoftGraphDriveItem) GetFolder() AnyOfmicrosoftGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MicrosoftGraphDriveItem) GetFolderOk() (*AnyOfmicrosoftGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *MicrosoftGraphDriveItem) SetFolder(v AnyOfmicrosoftGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *MicrosoftGraphDriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *MicrosoftGraphDriveItem) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *MicrosoftGraphDriveItem) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetImage

`func (o *MicrosoftGraphDriveItem) GetImage() AnyOfmicrosoftGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MicrosoftGraphDriveItem) GetImageOk() (*AnyOfmicrosoftGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MicrosoftGraphDriveItem) SetImage(v AnyOfmicrosoftGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *MicrosoftGraphDriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *MicrosoftGraphDriveItem) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *MicrosoftGraphDriveItem) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphDriveItem) GetLocation() AnyOfmicrosoftGraphGeoCoordinates`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphDriveItem) GetLocationOk() (*AnyOfmicrosoftGraphGeoCoordinates, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphDriveItem) SetLocation(v AnyOfmicrosoftGraphGeoCoordinates)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphDriveItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphDriveItem) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphDriveItem) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetMalware

`func (o *MicrosoftGraphDriveItem) GetMalware() AnyOfmicrosoftGraphMalware`

GetMalware returns the Malware field if non-nil, zero value otherwise.

### GetMalwareOk

`func (o *MicrosoftGraphDriveItem) GetMalwareOk() (*AnyOfmicrosoftGraphMalware, bool)`

GetMalwareOk returns a tuple with the Malware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalware

`func (o *MicrosoftGraphDriveItem) SetMalware(v AnyOfmicrosoftGraphMalware)`

SetMalware sets Malware field to given value.

### HasMalware

`func (o *MicrosoftGraphDriveItem) HasMalware() bool`

HasMalware returns a boolean if a field has been set.

### SetMalwareNil

`func (o *MicrosoftGraphDriveItem) SetMalwareNil(b bool)`

 SetMalwareNil sets the value for Malware to be an explicit nil

### UnsetMalware
`func (o *MicrosoftGraphDriveItem) UnsetMalware()`

UnsetMalware ensures that no value is present for Malware, not even an explicit nil
### GetPackage

`func (o *MicrosoftGraphDriveItem) GetPackage() AnyOfmicrosoftGraphPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *MicrosoftGraphDriveItem) GetPackageOk() (*AnyOfmicrosoftGraphPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *MicrosoftGraphDriveItem) SetPackage(v AnyOfmicrosoftGraphPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *MicrosoftGraphDriveItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### SetPackageNil

`func (o *MicrosoftGraphDriveItem) SetPackageNil(b bool)`

 SetPackageNil sets the value for Package to be an explicit nil

### UnsetPackage
`func (o *MicrosoftGraphDriveItem) UnsetPackage()`

UnsetPackage ensures that no value is present for Package, not even an explicit nil
### GetPendingOperations

`func (o *MicrosoftGraphDriveItem) GetPendingOperations() AnyOfmicrosoftGraphPendingOperations`

GetPendingOperations returns the PendingOperations field if non-nil, zero value otherwise.

### GetPendingOperationsOk

`func (o *MicrosoftGraphDriveItem) GetPendingOperationsOk() (*AnyOfmicrosoftGraphPendingOperations, bool)`

GetPendingOperationsOk returns a tuple with the PendingOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOperations

`func (o *MicrosoftGraphDriveItem) SetPendingOperations(v AnyOfmicrosoftGraphPendingOperations)`

SetPendingOperations sets PendingOperations field to given value.

### HasPendingOperations

`func (o *MicrosoftGraphDriveItem) HasPendingOperations() bool`

HasPendingOperations returns a boolean if a field has been set.

### SetPendingOperationsNil

`func (o *MicrosoftGraphDriveItem) SetPendingOperationsNil(b bool)`

 SetPendingOperationsNil sets the value for PendingOperations to be an explicit nil

### UnsetPendingOperations
`func (o *MicrosoftGraphDriveItem) UnsetPendingOperations()`

UnsetPendingOperations ensures that no value is present for PendingOperations, not even an explicit nil
### GetPhoto

`func (o *MicrosoftGraphDriveItem) GetPhoto() AnyOfmicrosoftGraphPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphDriveItem) GetPhotoOk() (*AnyOfmicrosoftGraphPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MicrosoftGraphDriveItem) SetPhoto(v AnyOfmicrosoftGraphPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MicrosoftGraphDriveItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *MicrosoftGraphDriveItem) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *MicrosoftGraphDriveItem) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPublication

`func (o *MicrosoftGraphDriveItem) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphDriveItem) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *MicrosoftGraphDriveItem) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *MicrosoftGraphDriveItem) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *MicrosoftGraphDriveItem) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *MicrosoftGraphDriveItem) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetRemoteItem

`func (o *MicrosoftGraphDriveItem) GetRemoteItem() AnyOfmicrosoftGraphRemoteItem`

GetRemoteItem returns the RemoteItem field if non-nil, zero value otherwise.

### GetRemoteItemOk

`func (o *MicrosoftGraphDriveItem) GetRemoteItemOk() (*AnyOfmicrosoftGraphRemoteItem, bool)`

GetRemoteItemOk returns a tuple with the RemoteItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteItem

`func (o *MicrosoftGraphDriveItem) SetRemoteItem(v AnyOfmicrosoftGraphRemoteItem)`

SetRemoteItem sets RemoteItem field to given value.

### HasRemoteItem

`func (o *MicrosoftGraphDriveItem) HasRemoteItem() bool`

HasRemoteItem returns a boolean if a field has been set.

### SetRemoteItemNil

`func (o *MicrosoftGraphDriveItem) SetRemoteItemNil(b bool)`

 SetRemoteItemNil sets the value for RemoteItem to be an explicit nil

### UnsetRemoteItem
`func (o *MicrosoftGraphDriveItem) UnsetRemoteItem()`

UnsetRemoteItem ensures that no value is present for RemoteItem, not even an explicit nil
### GetRoot

`func (o *MicrosoftGraphDriveItem) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphDriveItem) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *MicrosoftGraphDriveItem) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *MicrosoftGraphDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *MicrosoftGraphDriveItem) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *MicrosoftGraphDriveItem) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSearchResult

`func (o *MicrosoftGraphDriveItem) GetSearchResult() AnyOfmicrosoftGraphSearchResult`

GetSearchResult returns the SearchResult field if non-nil, zero value otherwise.

### GetSearchResultOk

`func (o *MicrosoftGraphDriveItem) GetSearchResultOk() (*AnyOfmicrosoftGraphSearchResult, bool)`

GetSearchResultOk returns a tuple with the SearchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResult

`func (o *MicrosoftGraphDriveItem) SetSearchResult(v AnyOfmicrosoftGraphSearchResult)`

SetSearchResult sets SearchResult field to given value.

### HasSearchResult

`func (o *MicrosoftGraphDriveItem) HasSearchResult() bool`

HasSearchResult returns a boolean if a field has been set.

### SetSearchResultNil

`func (o *MicrosoftGraphDriveItem) SetSearchResultNil(b bool)`

 SetSearchResultNil sets the value for SearchResult to be an explicit nil

### UnsetSearchResult
`func (o *MicrosoftGraphDriveItem) UnsetSearchResult()`

UnsetSearchResult ensures that no value is present for SearchResult, not even an explicit nil
### GetShared

`func (o *MicrosoftGraphDriveItem) GetShared() AnyOfmicrosoftGraphShared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *MicrosoftGraphDriveItem) GetSharedOk() (*AnyOfmicrosoftGraphShared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *MicrosoftGraphDriveItem) SetShared(v AnyOfmicrosoftGraphShared)`

SetShared sets Shared field to given value.

### HasShared

`func (o *MicrosoftGraphDriveItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *MicrosoftGraphDriveItem) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *MicrosoftGraphDriveItem) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetSharepointIds

`func (o *MicrosoftGraphDriveItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphDriveItem) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *MicrosoftGraphDriveItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *MicrosoftGraphDriveItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *MicrosoftGraphDriveItem) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *MicrosoftGraphDriveItem) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphDriveItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphDriveItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphDriveItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphDriveItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphDriveItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphDriveItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSpecialFolder

`func (o *MicrosoftGraphDriveItem) GetSpecialFolder() AnyOfmicrosoftGraphSpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *MicrosoftGraphDriveItem) GetSpecialFolderOk() (*AnyOfmicrosoftGraphSpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFolder

`func (o *MicrosoftGraphDriveItem) SetSpecialFolder(v AnyOfmicrosoftGraphSpecialFolder)`

SetSpecialFolder sets SpecialFolder field to given value.

### HasSpecialFolder

`func (o *MicrosoftGraphDriveItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### SetSpecialFolderNil

`func (o *MicrosoftGraphDriveItem) SetSpecialFolderNil(b bool)`

 SetSpecialFolderNil sets the value for SpecialFolder to be an explicit nil

### UnsetSpecialFolder
`func (o *MicrosoftGraphDriveItem) UnsetSpecialFolder()`

UnsetSpecialFolder ensures that no value is present for SpecialFolder, not even an explicit nil
### GetVideo

`func (o *MicrosoftGraphDriveItem) GetVideo() AnyOfmicrosoftGraphVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *MicrosoftGraphDriveItem) GetVideoOk() (*AnyOfmicrosoftGraphVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *MicrosoftGraphDriveItem) SetVideo(v AnyOfmicrosoftGraphVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *MicrosoftGraphDriveItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *MicrosoftGraphDriveItem) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *MicrosoftGraphDriveItem) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil
### GetWebDavUrl

`func (o *MicrosoftGraphDriveItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *MicrosoftGraphDriveItem) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *MicrosoftGraphDriveItem) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *MicrosoftGraphDriveItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrlNil

`func (o *MicrosoftGraphDriveItem) SetWebDavUrlNil(b bool)`

 SetWebDavUrlNil sets the value for WebDavUrl to be an explicit nil

### UnsetWebDavUrl
`func (o *MicrosoftGraphDriveItem) UnsetWebDavUrl()`

UnsetWebDavUrl ensures that no value is present for WebDavUrl, not even an explicit nil
### GetWorkbook

`func (o *MicrosoftGraphDriveItem) GetWorkbook() AnyOfmicrosoftGraphWorkbook`

GetWorkbook returns the Workbook field if non-nil, zero value otherwise.

### GetWorkbookOk

`func (o *MicrosoftGraphDriveItem) GetWorkbookOk() (*AnyOfmicrosoftGraphWorkbook, bool)`

GetWorkbookOk returns a tuple with the Workbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkbook

`func (o *MicrosoftGraphDriveItem) SetWorkbook(v AnyOfmicrosoftGraphWorkbook)`

SetWorkbook sets Workbook field to given value.

### HasWorkbook

`func (o *MicrosoftGraphDriveItem) HasWorkbook() bool`

HasWorkbook returns a boolean if a field has been set.

### SetWorkbookNil

`func (o *MicrosoftGraphDriveItem) SetWorkbookNil(b bool)`

 SetWorkbookNil sets the value for Workbook to be an explicit nil

### UnsetWorkbook
`func (o *MicrosoftGraphDriveItem) UnsetWorkbook()`

UnsetWorkbook ensures that no value is present for Workbook, not even an explicit nil
### GetAnalytics

`func (o *MicrosoftGraphDriveItem) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *MicrosoftGraphDriveItem) GetAnalyticsOk() (*AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *MicrosoftGraphDriveItem) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *MicrosoftGraphDriveItem) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalyticsNil

`func (o *MicrosoftGraphDriveItem) SetAnalyticsNil(b bool)`

 SetAnalyticsNil sets the value for Analytics to be an explicit nil

### UnsetAnalytics
`func (o *MicrosoftGraphDriveItem) UnsetAnalytics()`

UnsetAnalytics ensures that no value is present for Analytics, not even an explicit nil
### GetChildren

`func (o *MicrosoftGraphDriveItem) GetChildren() []MicrosoftGraphDriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MicrosoftGraphDriveItem) GetChildrenOk() (*[]MicrosoftGraphDriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MicrosoftGraphDriveItem) SetChildren(v []MicrosoftGraphDriveItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MicrosoftGraphDriveItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetListItem

`func (o *MicrosoftGraphDriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *MicrosoftGraphDriveItem) GetListItemOk() (*AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItem

`func (o *MicrosoftGraphDriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem sets ListItem field to given value.

### HasListItem

`func (o *MicrosoftGraphDriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItemNil

`func (o *MicrosoftGraphDriveItem) SetListItemNil(b bool)`

 SetListItemNil sets the value for ListItem to be an explicit nil

### UnsetListItem
`func (o *MicrosoftGraphDriveItem) UnsetListItem()`

UnsetListItem ensures that no value is present for ListItem, not even an explicit nil
### GetPermissions

`func (o *MicrosoftGraphDriveItem) GetPermissions() []MicrosoftGraphPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MicrosoftGraphDriveItem) GetPermissionsOk() (*[]MicrosoftGraphPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MicrosoftGraphDriveItem) SetPermissions(v []MicrosoftGraphPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MicrosoftGraphDriveItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSubscriptions

`func (o *MicrosoftGraphDriveItem) GetSubscriptions() []MicrosoftGraphSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MicrosoftGraphDriveItem) GetSubscriptionsOk() (*[]MicrosoftGraphSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MicrosoftGraphDriveItem) SetSubscriptions(v []MicrosoftGraphSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *MicrosoftGraphDriveItem) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetThumbnails

`func (o *MicrosoftGraphDriveItem) GetThumbnails() []MicrosoftGraphThumbnailSet`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *MicrosoftGraphDriveItem) GetThumbnailsOk() (*[]MicrosoftGraphThumbnailSet, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *MicrosoftGraphDriveItem) SetThumbnails(v []MicrosoftGraphThumbnailSet)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *MicrosoftGraphDriveItem) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetVersions

`func (o *MicrosoftGraphDriveItem) GetVersions() []MicrosoftGraphDriveItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MicrosoftGraphDriveItem) GetVersionsOk() (*[]MicrosoftGraphDriveItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MicrosoftGraphDriveItem) SetVersions(v []MicrosoftGraphDriveItemVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MicrosoftGraphDriveItem) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


