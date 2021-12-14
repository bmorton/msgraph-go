# MicrosoftGraphDocumentSetContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to [**AnyOfmicrosoftGraphContentTypeInfo**](anyOf&lt;microsoft.graph.contentTypeInfo&gt;.md) | Content type information of the file. | [optional] 
**FileName** | Pointer to **NullableString** | Name of the file in resource folder that should be added as a default content or a template in the document set. | [optional] 
**FolderName** | Pointer to **NullableString** | Folder name in which the file will be placed when a new document set is created in the library. | [optional] 

## Methods

### NewMicrosoftGraphDocumentSetContent

`func NewMicrosoftGraphDocumentSetContent() *MicrosoftGraphDocumentSetContent`

NewMicrosoftGraphDocumentSetContent instantiates a new MicrosoftGraphDocumentSetContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDocumentSetContentWithDefaults

`func NewMicrosoftGraphDocumentSetContentWithDefaults() *MicrosoftGraphDocumentSetContent`

NewMicrosoftGraphDocumentSetContentWithDefaults instantiates a new MicrosoftGraphDocumentSetContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *MicrosoftGraphDocumentSetContent) GetContentType() AnyOfmicrosoftGraphContentTypeInfo`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphDocumentSetContent) GetContentTypeOk() (*AnyOfmicrosoftGraphContentTypeInfo, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphDocumentSetContent) SetContentType(v AnyOfmicrosoftGraphContentTypeInfo)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphDocumentSetContent) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MicrosoftGraphDocumentSetContent) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MicrosoftGraphDocumentSetContent) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetFileName

`func (o *MicrosoftGraphDocumentSetContent) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphDocumentSetContent) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MicrosoftGraphDocumentSetContent) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MicrosoftGraphDocumentSetContent) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MicrosoftGraphDocumentSetContent) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MicrosoftGraphDocumentSetContent) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFolderName

`func (o *MicrosoftGraphDocumentSetContent) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *MicrosoftGraphDocumentSetContent) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *MicrosoftGraphDocumentSetContent) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *MicrosoftGraphDocumentSetContent) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### SetFolderNameNil

`func (o *MicrosoftGraphDocumentSetContent) SetFolderNameNil(b bool)`

 SetFolderNameNil sets the value for FolderName to be an explicit nil

### UnsetFolderName
`func (o *MicrosoftGraphDocumentSetContent) UnsetFolderName()`

UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


