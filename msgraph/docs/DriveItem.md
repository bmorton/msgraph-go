# DriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDriveItem

`func NewDriveItem() *DriveItem`

NewDriveItem instantiates a new DriveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveItemWithDefaults

`func NewDriveItemWithDefaults() *DriveItem`

NewDriveItemWithDefaults instantiates a new DriveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudio

`func (o *DriveItem) GetAudio() AnyOfmicrosoftGraphAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *DriveItem) GetAudioOk() (*AnyOfmicrosoftGraphAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *DriveItem) SetAudio(v AnyOfmicrosoftGraphAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *DriveItem) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *DriveItem) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *DriveItem) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetBundle

`func (o *DriveItem) GetBundle() AnyOfmicrosoftGraphBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *DriveItem) GetBundleOk() (*AnyOfmicrosoftGraphBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *DriveItem) SetBundle(v AnyOfmicrosoftGraphBundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *DriveItem) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### SetBundleNil

`func (o *DriveItem) SetBundleNil(b bool)`

 SetBundleNil sets the value for Bundle to be an explicit nil

### UnsetBundle
`func (o *DriveItem) UnsetBundle()`

UnsetBundle ensures that no value is present for Bundle, not even an explicit nil
### GetContent

`func (o *DriveItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DriveItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DriveItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DriveItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *DriveItem) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *DriveItem) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCTag

`func (o *DriveItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *DriveItem) GetCTagOk() (*string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTag

`func (o *DriveItem) SetCTag(v string)`

SetCTag sets CTag field to given value.

### HasCTag

`func (o *DriveItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### SetCTagNil

`func (o *DriveItem) SetCTagNil(b bool)`

 SetCTagNil sets the value for CTag to be an explicit nil

### UnsetCTag
`func (o *DriveItem) UnsetCTag()`

UnsetCTag ensures that no value is present for CTag, not even an explicit nil
### GetDeleted

`func (o *DriveItem) GetDeleted() AnyOfmicrosoftGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DriveItem) GetDeletedOk() (*AnyOfmicrosoftGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DriveItem) SetDeleted(v AnyOfmicrosoftGraphDeleted)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *DriveItem) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *DriveItem) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetFile

`func (o *DriveItem) GetFile() AnyOfmicrosoftGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DriveItem) GetFileOk() (*AnyOfmicrosoftGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DriveItem) SetFile(v AnyOfmicrosoftGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *DriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *DriveItem) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *DriveItem) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFileSystemInfo

`func (o *DriveItem) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *DriveItem) GetFileSystemInfoOk() (*AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *DriveItem) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *DriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfoNil

`func (o *DriveItem) SetFileSystemInfoNil(b bool)`

 SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil

### UnsetFileSystemInfo
`func (o *DriveItem) UnsetFileSystemInfo()`

UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
### GetFolder

`func (o *DriveItem) GetFolder() AnyOfmicrosoftGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *DriveItem) GetFolderOk() (*AnyOfmicrosoftGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *DriveItem) SetFolder(v AnyOfmicrosoftGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *DriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *DriveItem) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *DriveItem) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetImage

`func (o *DriveItem) GetImage() AnyOfmicrosoftGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DriveItem) GetImageOk() (*AnyOfmicrosoftGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DriveItem) SetImage(v AnyOfmicrosoftGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *DriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *DriveItem) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *DriveItem) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLocation

`func (o *DriveItem) GetLocation() AnyOfmicrosoftGraphGeoCoordinates`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DriveItem) GetLocationOk() (*AnyOfmicrosoftGraphGeoCoordinates, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DriveItem) SetLocation(v AnyOfmicrosoftGraphGeoCoordinates)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DriveItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DriveItem) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DriveItem) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetMalware

`func (o *DriveItem) GetMalware() AnyOfmicrosoftGraphMalware`

GetMalware returns the Malware field if non-nil, zero value otherwise.

### GetMalwareOk

`func (o *DriveItem) GetMalwareOk() (*AnyOfmicrosoftGraphMalware, bool)`

GetMalwareOk returns a tuple with the Malware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalware

`func (o *DriveItem) SetMalware(v AnyOfmicrosoftGraphMalware)`

SetMalware sets Malware field to given value.

### HasMalware

`func (o *DriveItem) HasMalware() bool`

HasMalware returns a boolean if a field has been set.

### SetMalwareNil

`func (o *DriveItem) SetMalwareNil(b bool)`

 SetMalwareNil sets the value for Malware to be an explicit nil

### UnsetMalware
`func (o *DriveItem) UnsetMalware()`

UnsetMalware ensures that no value is present for Malware, not even an explicit nil
### GetPackage

`func (o *DriveItem) GetPackage() AnyOfmicrosoftGraphPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DriveItem) GetPackageOk() (*AnyOfmicrosoftGraphPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DriveItem) SetPackage(v AnyOfmicrosoftGraphPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *DriveItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### SetPackageNil

`func (o *DriveItem) SetPackageNil(b bool)`

 SetPackageNil sets the value for Package to be an explicit nil

### UnsetPackage
`func (o *DriveItem) UnsetPackage()`

UnsetPackage ensures that no value is present for Package, not even an explicit nil
### GetPendingOperations

`func (o *DriveItem) GetPendingOperations() AnyOfmicrosoftGraphPendingOperations`

GetPendingOperations returns the PendingOperations field if non-nil, zero value otherwise.

### GetPendingOperationsOk

`func (o *DriveItem) GetPendingOperationsOk() (*AnyOfmicrosoftGraphPendingOperations, bool)`

GetPendingOperationsOk returns a tuple with the PendingOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOperations

`func (o *DriveItem) SetPendingOperations(v AnyOfmicrosoftGraphPendingOperations)`

SetPendingOperations sets PendingOperations field to given value.

### HasPendingOperations

`func (o *DriveItem) HasPendingOperations() bool`

HasPendingOperations returns a boolean if a field has been set.

### SetPendingOperationsNil

`func (o *DriveItem) SetPendingOperationsNil(b bool)`

 SetPendingOperationsNil sets the value for PendingOperations to be an explicit nil

### UnsetPendingOperations
`func (o *DriveItem) UnsetPendingOperations()`

UnsetPendingOperations ensures that no value is present for PendingOperations, not even an explicit nil
### GetPhoto

`func (o *DriveItem) GetPhoto() AnyOfmicrosoftGraphPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *DriveItem) GetPhotoOk() (*AnyOfmicrosoftGraphPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *DriveItem) SetPhoto(v AnyOfmicrosoftGraphPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *DriveItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *DriveItem) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *DriveItem) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPublication

`func (o *DriveItem) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *DriveItem) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *DriveItem) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *DriveItem) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *DriveItem) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *DriveItem) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetRemoteItem

`func (o *DriveItem) GetRemoteItem() AnyOfmicrosoftGraphRemoteItem`

GetRemoteItem returns the RemoteItem field if non-nil, zero value otherwise.

### GetRemoteItemOk

`func (o *DriveItem) GetRemoteItemOk() (*AnyOfmicrosoftGraphRemoteItem, bool)`

GetRemoteItemOk returns a tuple with the RemoteItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteItem

`func (o *DriveItem) SetRemoteItem(v AnyOfmicrosoftGraphRemoteItem)`

SetRemoteItem sets RemoteItem field to given value.

### HasRemoteItem

`func (o *DriveItem) HasRemoteItem() bool`

HasRemoteItem returns a boolean if a field has been set.

### SetRemoteItemNil

`func (o *DriveItem) SetRemoteItemNil(b bool)`

 SetRemoteItemNil sets the value for RemoteItem to be an explicit nil

### UnsetRemoteItem
`func (o *DriveItem) UnsetRemoteItem()`

UnsetRemoteItem ensures that no value is present for RemoteItem, not even an explicit nil
### GetRoot

`func (o *DriveItem) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *DriveItem) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *DriveItem) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *DriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *DriveItem) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *DriveItem) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil
### GetSearchResult

`func (o *DriveItem) GetSearchResult() AnyOfmicrosoftGraphSearchResult`

GetSearchResult returns the SearchResult field if non-nil, zero value otherwise.

### GetSearchResultOk

`func (o *DriveItem) GetSearchResultOk() (*AnyOfmicrosoftGraphSearchResult, bool)`

GetSearchResultOk returns a tuple with the SearchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResult

`func (o *DriveItem) SetSearchResult(v AnyOfmicrosoftGraphSearchResult)`

SetSearchResult sets SearchResult field to given value.

### HasSearchResult

`func (o *DriveItem) HasSearchResult() bool`

HasSearchResult returns a boolean if a field has been set.

### SetSearchResultNil

`func (o *DriveItem) SetSearchResultNil(b bool)`

 SetSearchResultNil sets the value for SearchResult to be an explicit nil

### UnsetSearchResult
`func (o *DriveItem) UnsetSearchResult()`

UnsetSearchResult ensures that no value is present for SearchResult, not even an explicit nil
### GetShared

`func (o *DriveItem) GetShared() AnyOfmicrosoftGraphShared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DriveItem) GetSharedOk() (*AnyOfmicrosoftGraphShared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DriveItem) SetShared(v AnyOfmicrosoftGraphShared)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DriveItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *DriveItem) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *DriveItem) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetSharepointIds

`func (o *DriveItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *DriveItem) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *DriveItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *DriveItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *DriveItem) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *DriveItem) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSize

`func (o *DriveItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DriveItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DriveItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DriveItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *DriveItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *DriveItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSpecialFolder

`func (o *DriveItem) GetSpecialFolder() AnyOfmicrosoftGraphSpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *DriveItem) GetSpecialFolderOk() (*AnyOfmicrosoftGraphSpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFolder

`func (o *DriveItem) SetSpecialFolder(v AnyOfmicrosoftGraphSpecialFolder)`

SetSpecialFolder sets SpecialFolder field to given value.

### HasSpecialFolder

`func (o *DriveItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### SetSpecialFolderNil

`func (o *DriveItem) SetSpecialFolderNil(b bool)`

 SetSpecialFolderNil sets the value for SpecialFolder to be an explicit nil

### UnsetSpecialFolder
`func (o *DriveItem) UnsetSpecialFolder()`

UnsetSpecialFolder ensures that no value is present for SpecialFolder, not even an explicit nil
### GetVideo

`func (o *DriveItem) GetVideo() AnyOfmicrosoftGraphVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *DriveItem) GetVideoOk() (*AnyOfmicrosoftGraphVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *DriveItem) SetVideo(v AnyOfmicrosoftGraphVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *DriveItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *DriveItem) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *DriveItem) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil
### GetWebDavUrl

`func (o *DriveItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *DriveItem) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *DriveItem) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *DriveItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrlNil

`func (o *DriveItem) SetWebDavUrlNil(b bool)`

 SetWebDavUrlNil sets the value for WebDavUrl to be an explicit nil

### UnsetWebDavUrl
`func (o *DriveItem) UnsetWebDavUrl()`

UnsetWebDavUrl ensures that no value is present for WebDavUrl, not even an explicit nil
### GetWorkbook

`func (o *DriveItem) GetWorkbook() AnyOfmicrosoftGraphWorkbook`

GetWorkbook returns the Workbook field if non-nil, zero value otherwise.

### GetWorkbookOk

`func (o *DriveItem) GetWorkbookOk() (*AnyOfmicrosoftGraphWorkbook, bool)`

GetWorkbookOk returns a tuple with the Workbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkbook

`func (o *DriveItem) SetWorkbook(v AnyOfmicrosoftGraphWorkbook)`

SetWorkbook sets Workbook field to given value.

### HasWorkbook

`func (o *DriveItem) HasWorkbook() bool`

HasWorkbook returns a boolean if a field has been set.

### SetWorkbookNil

`func (o *DriveItem) SetWorkbookNil(b bool)`

 SetWorkbookNil sets the value for Workbook to be an explicit nil

### UnsetWorkbook
`func (o *DriveItem) UnsetWorkbook()`

UnsetWorkbook ensures that no value is present for Workbook, not even an explicit nil
### GetAnalytics

`func (o *DriveItem) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *DriveItem) GetAnalyticsOk() (*AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *DriveItem) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *DriveItem) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalyticsNil

`func (o *DriveItem) SetAnalyticsNil(b bool)`

 SetAnalyticsNil sets the value for Analytics to be an explicit nil

### UnsetAnalytics
`func (o *DriveItem) UnsetAnalytics()`

UnsetAnalytics ensures that no value is present for Analytics, not even an explicit nil
### GetChildren

`func (o *DriveItem) GetChildren() []MicrosoftGraphDriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DriveItem) GetChildrenOk() (*[]MicrosoftGraphDriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *DriveItem) SetChildren(v []MicrosoftGraphDriveItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *DriveItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetListItem

`func (o *DriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *DriveItem) GetListItemOk() (*AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItem

`func (o *DriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem sets ListItem field to given value.

### HasListItem

`func (o *DriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItemNil

`func (o *DriveItem) SetListItemNil(b bool)`

 SetListItemNil sets the value for ListItem to be an explicit nil

### UnsetListItem
`func (o *DriveItem) UnsetListItem()`

UnsetListItem ensures that no value is present for ListItem, not even an explicit nil
### GetPermissions

`func (o *DriveItem) GetPermissions() []MicrosoftGraphPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DriveItem) GetPermissionsOk() (*[]MicrosoftGraphPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DriveItem) SetPermissions(v []MicrosoftGraphPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DriveItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DriveItem) GetSubscriptions() []MicrosoftGraphSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DriveItem) GetSubscriptionsOk() (*[]MicrosoftGraphSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DriveItem) SetSubscriptions(v []MicrosoftGraphSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *DriveItem) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetThumbnails

`func (o *DriveItem) GetThumbnails() []MicrosoftGraphThumbnailSet`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *DriveItem) GetThumbnailsOk() (*[]MicrosoftGraphThumbnailSet, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *DriveItem) SetThumbnails(v []MicrosoftGraphThumbnailSet)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *DriveItem) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetVersions

`func (o *DriveItem) GetVersions() []MicrosoftGraphDriveItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DriveItem) GetVersionsOk() (*[]MicrosoftGraphDriveItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *DriveItem) SetVersions(v []MicrosoftGraphDriveItemVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *DriveItem) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


