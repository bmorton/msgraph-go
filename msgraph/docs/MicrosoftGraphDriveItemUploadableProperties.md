# MicrosoftGraphDriveItemUploadableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Provides a user-visible description of the item. Read-write. Only on OneDrive Personal. | [optional] 
**FileSize** | Pointer to **NullableInt64** | Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal. | [optional] 
**FileSystemInfo** | Pointer to [**AnyOfmicrosoftGraphFileSystemInfo**](anyOf&lt;microsoft.graph.fileSystemInfo&gt;.md) | File system information on client. Read-write. | [optional] 
**Name** | Pointer to **NullableString** | The name of the item (filename and extension). Read-write. | [optional] 

## Methods

### NewMicrosoftGraphDriveItemUploadableProperties

`func NewMicrosoftGraphDriveItemUploadableProperties() *MicrosoftGraphDriveItemUploadableProperties`

NewMicrosoftGraphDriveItemUploadableProperties instantiates a new MicrosoftGraphDriveItemUploadableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDriveItemUploadablePropertiesWithDefaults

`func NewMicrosoftGraphDriveItemUploadablePropertiesWithDefaults() *MicrosoftGraphDriveItemUploadableProperties`

NewMicrosoftGraphDriveItemUploadablePropertiesWithDefaults instantiates a new MicrosoftGraphDriveItemUploadableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphDriveItemUploadableProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphDriveItemUploadableProperties) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFileSize

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *MicrosoftGraphDriveItemUploadableProperties) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *MicrosoftGraphDriveItemUploadableProperties) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetFileSystemInfo

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetFileSystemInfoOk() (*AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemInfo

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo sets FileSystemInfo field to given value.

### HasFileSystemInfo

`func (o *MicrosoftGraphDriveItemUploadableProperties) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfoNil

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetFileSystemInfoNil(b bool)`

 SetFileSystemInfoNil sets the value for FileSystemInfo to be an explicit nil

### UnsetFileSystemInfo
`func (o *MicrosoftGraphDriveItemUploadableProperties) UnsetFileSystemInfo()`

UnsetFileSystemInfo ensures that no value is present for FileSystemInfo, not even an explicit nil
### GetName

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphDriveItemUploadableProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphDriveItemUploadableProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphDriveItemUploadableProperties) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphDriveItemUploadableProperties) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


