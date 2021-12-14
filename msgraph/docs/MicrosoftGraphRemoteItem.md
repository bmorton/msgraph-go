# MicrosoftGraphRemoteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user, device, and application which created the item. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of item creation. Read-only. | [optional] 
**File** | Pointer to [**AnyOfmicrosoftGraphFile**](anyOf&lt;microsoft.graph.file&gt;.md) | Indicates that the remote item is a file. Read-only. | [optional] 
**FileSystemInfo** | Pointer to [**AnyOfmicrosoftGraphFileSystemInfo**](anyOf&lt;microsoft.graph.fileSystemInfo&gt;.md) | Information about the remote item from the local file system. Read-only. | [optional] 
**Folder** | Pointer to [**AnyOfmicrosoftGraphFolder**](anyOf&lt;microsoft.graph.folder&gt;.md) | Indicates that the remote item is a folder. Read-only. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier for the remote item in its drive. Read-only. | [optional] 
**Image** | Pointer to [**AnyOfmicrosoftGraphImage**](anyOf&lt;microsoft.graph.image&gt;.md) | Image metadata, if the item is an image. Read-only. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user, device, and application which last modified the item. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Date and time the item was last modified. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | Optional. Filename of the remote item. Read-only. | [optional] 
**Package** | Pointer to [**AnyOfmicrosoftGraphPackage**](anyOf&lt;microsoft.graph.package&gt;.md) | If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only. | [optional] 
**ParentReference** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) | Properties of the parent of the remote item. Read-only. | [optional] 
**Shared** | Pointer to [**AnyOfmicrosoftGraphShared**](anyOf&lt;microsoft.graph.shared&gt;.md) | Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only. | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) | Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only. | [optional] 
**Size** | Pointer to **NullableInt64** | Size of the remote item. Read-only. | [optional] 
**SpecialFolder** | Pointer to [**AnyOfmicrosoftGraphSpecialFolder**](anyOf&lt;microsoft.graph.specialFolder&gt;.md) | If the current item is also available as a special folder, this facet is returned. Read-only. | [optional] 
**Video** | Pointer to [**AnyOfmicrosoftGraphVideo**](anyOf&lt;microsoft.graph.video&gt;.md) | Video metadata, if the item is a video. Read-only. | [optional] 
**WebDavUrl** | Pointer to **NullableString** | DAV compatible URL for the item. | [optional] 
**WebUrl** | Pointer to **NullableString** | URL that displays the resource in the browser. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphRemoteItem

`func NewMicrosoftGraphRemoteItem() *MicrosoftGraphRemoteItem`

NewMicrosoftGraphRemoteItem instantiates a new MicrosoftGraphRemoteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRemoteItemWithDefaults

`func NewMicrosoftGraphRemoteItemWithDefaults() *MicrosoftGraphRemoteItem`

NewMicrosoftGraphRemoteItemWithDefaults instantiates a new MicrosoftGraphRemoteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *MicrosoftGraphRemoteItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphRemoteItem) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphRemoteItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphRemoteItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphRemoteItem) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphRemoteItem) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphRemoteItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphRemoteItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphRemoteItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphRemoteItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphRemoteItem) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphRemoteItem) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetFile

`func (o *MicrosoftGraphRemoteItem) GetFile() AnyOfmicrosoftGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MicrosoftGraphRemoteItem) GetFileOk() (*AnyOfmicrosoftGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MicrosoftGraphRemoteItem) SetFile(v AnyOfmicrosoftGraphFile)`

SetFile sets File field to given value.

### HasFile

`func (o *MicrosoftGraphRemoteItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *MicrosoftGraphRemoteItem) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *MicrosoftGraphRemoteItem) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFileSystemInfo

`func (o *MicrosoftGraphRemoteItem) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *MicrosoftGraphRemoteItem) GetFileSystemInfoOk() (*AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *MicrosoftGraphRemoteItem) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *MicrosoftGraphRemoteItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfoNil

`func (o *MicrosoftGraphRemoteItem) SetFileSystemInfoNil(b bool)`

 SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil

### UnsetFileSystemInfo
`func (o *MicrosoftGraphRemoteItem) UnsetFileSystemInfo()`

UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
### GetFolder

`func (o *MicrosoftGraphRemoteItem) GetFolder() AnyOfmicrosoftGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MicrosoftGraphRemoteItem) GetFolderOk() (*AnyOfmicrosoftGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *MicrosoftGraphRemoteItem) SetFolder(v AnyOfmicrosoftGraphFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *MicrosoftGraphRemoteItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *MicrosoftGraphRemoteItem) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *MicrosoftGraphRemoteItem) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetId

`func (o *MicrosoftGraphRemoteItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRemoteItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphRemoteItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphRemoteItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphRemoteItem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphRemoteItem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetImage

`func (o *MicrosoftGraphRemoteItem) GetImage() AnyOfmicrosoftGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MicrosoftGraphRemoteItem) GetImageOk() (*AnyOfmicrosoftGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MicrosoftGraphRemoteItem) SetImage(v AnyOfmicrosoftGraphImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *MicrosoftGraphRemoteItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *MicrosoftGraphRemoteItem) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *MicrosoftGraphRemoteItem) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphRemoteItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphRemoteItem) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphRemoteItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphRemoteItem) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetName

`func (o *MicrosoftGraphRemoteItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphRemoteItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphRemoteItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphRemoteItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphRemoteItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphRemoteItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackage

`func (o *MicrosoftGraphRemoteItem) GetPackage() AnyOfmicrosoftGraphPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *MicrosoftGraphRemoteItem) GetPackageOk() (*AnyOfmicrosoftGraphPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *MicrosoftGraphRemoteItem) SetPackage(v AnyOfmicrosoftGraphPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *MicrosoftGraphRemoteItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### SetPackageNil

`func (o *MicrosoftGraphRemoteItem) SetPackageNil(b bool)`

 SetPackageNil sets the value for Package to be an explicit nil

### UnsetPackage
`func (o *MicrosoftGraphRemoteItem) UnsetPackage()`

UnsetPackage ensures that no value is present for Package, not even an explicit nil
### GetParentReference

`func (o *MicrosoftGraphRemoteItem) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphRemoteItem) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MicrosoftGraphRemoteItem) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MicrosoftGraphRemoteItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReferenceNil

`func (o *MicrosoftGraphRemoteItem) SetParentReferenceNil(b bool)`

 SetParentReferenceNil sets the value for ParentReference to be an explicit nil

### UnsetParentReference
`func (o *MicrosoftGraphRemoteItem) UnsetParentReference()`

UnsetParentReference ensures that no value is present for ParentReference, not even an explicit nil
### GetShared

`func (o *MicrosoftGraphRemoteItem) GetShared() AnyOfmicrosoftGraphShared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *MicrosoftGraphRemoteItem) GetSharedOk() (*AnyOfmicrosoftGraphShared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *MicrosoftGraphRemoteItem) SetShared(v AnyOfmicrosoftGraphShared)`

SetShared sets Shared field to given value.

### HasShared

`func (o *MicrosoftGraphRemoteItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *MicrosoftGraphRemoteItem) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *MicrosoftGraphRemoteItem) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetSharepointIds

`func (o *MicrosoftGraphRemoteItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphRemoteItem) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *MicrosoftGraphRemoteItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *MicrosoftGraphRemoteItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *MicrosoftGraphRemoteItem) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *MicrosoftGraphRemoteItem) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphRemoteItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphRemoteItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphRemoteItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphRemoteItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphRemoteItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphRemoteItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSpecialFolder

`func (o *MicrosoftGraphRemoteItem) GetSpecialFolder() AnyOfmicrosoftGraphSpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *MicrosoftGraphRemoteItem) GetSpecialFolderOk() (*AnyOfmicrosoftGraphSpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFolder

`func (o *MicrosoftGraphRemoteItem) SetSpecialFolder(v AnyOfmicrosoftGraphSpecialFolder)`

SetSpecialFolder sets SpecialFolder field to given value.

### HasSpecialFolder

`func (o *MicrosoftGraphRemoteItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### SetSpecialFolderNil

`func (o *MicrosoftGraphRemoteItem) SetSpecialFolderNil(b bool)`

 SetSpecialFolderNil sets the value for SpecialFolder to be an explicit nil

### UnsetSpecialFolder
`func (o *MicrosoftGraphRemoteItem) UnsetSpecialFolder()`

UnsetSpecialFolder ensures that no value is present for SpecialFolder, not even an explicit nil
### GetVideo

`func (o *MicrosoftGraphRemoteItem) GetVideo() AnyOfmicrosoftGraphVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *MicrosoftGraphRemoteItem) GetVideoOk() (*AnyOfmicrosoftGraphVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *MicrosoftGraphRemoteItem) SetVideo(v AnyOfmicrosoftGraphVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *MicrosoftGraphRemoteItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *MicrosoftGraphRemoteItem) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *MicrosoftGraphRemoteItem) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil
### GetWebDavUrl

`func (o *MicrosoftGraphRemoteItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *MicrosoftGraphRemoteItem) GetWebDavUrlOk() (*string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDavUrl

`func (o *MicrosoftGraphRemoteItem) SetWebDavUrl(v string)`

SetWebDavUrl sets WebDavUrl field to given value.

### HasWebDavUrl

`func (o *MicrosoftGraphRemoteItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrlNil

`func (o *MicrosoftGraphRemoteItem) SetWebDavUrlNil(b bool)`

 SetWebDavUrlNil sets the value for WebDavUrl to be an explicit nil

### UnsetWebDavUrl
`func (o *MicrosoftGraphRemoteItem) UnsetWebDavUrl()`

UnsetWebDavUrl ensures that no value is present for WebDavUrl, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphRemoteItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphRemoteItem) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphRemoteItem) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphRemoteItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphRemoteItem) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphRemoteItem) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


