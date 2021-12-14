# MicrosoftGraphResourceVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerDisplayName** | Pointer to **NullableString** | A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item. | [optional] 
**ContainerType** | Pointer to **NullableString** | Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness. | [optional] 
**ContainerWebUrl** | Pointer to **NullableString** | A path leading to the folder in which the item is stored. | [optional] 
**MediaType** | Pointer to **NullableString** | The item&#39;s media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported. | [optional] 
**PreviewImageUrl** | Pointer to **NullableString** | A URL leading to the preview image for the item. | [optional] 
**PreviewText** | Pointer to **NullableString** | A preview text for the item. | [optional] 
**Title** | Pointer to **NullableString** | The item&#39;s title text. | [optional] 
**Type** | Pointer to **NullableString** | The item&#39;s media type. Can be used for filtering for a specific file based on a specific type. See below for supported types. | [optional] 

## Methods

### NewMicrosoftGraphResourceVisualization

`func NewMicrosoftGraphResourceVisualization() *MicrosoftGraphResourceVisualization`

NewMicrosoftGraphResourceVisualization instantiates a new MicrosoftGraphResourceVisualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphResourceVisualizationWithDefaults

`func NewMicrosoftGraphResourceVisualizationWithDefaults() *MicrosoftGraphResourceVisualization`

NewMicrosoftGraphResourceVisualizationWithDefaults instantiates a new MicrosoftGraphResourceVisualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerDisplayName

`func (o *MicrosoftGraphResourceVisualization) GetContainerDisplayName() string`

GetContainerDisplayName returns the ContainerDisplayName field if non-nil, zero value otherwise.

### GetContainerDisplayNameOk

`func (o *MicrosoftGraphResourceVisualization) GetContainerDisplayNameOk() (*string, bool)`

GetContainerDisplayNameOk returns a tuple with the ContainerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDisplayName

`func (o *MicrosoftGraphResourceVisualization) SetContainerDisplayName(v string)`

SetContainerDisplayName sets ContainerDisplayName field to given value.

### HasContainerDisplayName

`func (o *MicrosoftGraphResourceVisualization) HasContainerDisplayName() bool`

HasContainerDisplayName returns a boolean if a field has been set.

### SetContainerDisplayNameNil

`func (o *MicrosoftGraphResourceVisualization) SetContainerDisplayNameNil(b bool)`

 SetContainerDisplayNameNil sets the value for ContainerDisplayName to be an explicit nil

### UnsetContainerDisplayName
`func (o *MicrosoftGraphResourceVisualization) UnsetContainerDisplayName()`

UnsetContainerDisplayName ensures that no value is present for ContainerDisplayName, not even an explicit nil
### GetContainerType

`func (o *MicrosoftGraphResourceVisualization) GetContainerType() string`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *MicrosoftGraphResourceVisualization) GetContainerTypeOk() (*string, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *MicrosoftGraphResourceVisualization) SetContainerType(v string)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *MicrosoftGraphResourceVisualization) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### SetContainerTypeNil

`func (o *MicrosoftGraphResourceVisualization) SetContainerTypeNil(b bool)`

 SetContainerTypeNil sets the value for ContainerType to be an explicit nil

### UnsetContainerType
`func (o *MicrosoftGraphResourceVisualization) UnsetContainerType()`

UnsetContainerType ensures that no value is present for ContainerType, not even an explicit nil
### GetContainerWebUrl

`func (o *MicrosoftGraphResourceVisualization) GetContainerWebUrl() string`

GetContainerWebUrl returns the ContainerWebUrl field if non-nil, zero value otherwise.

### GetContainerWebUrlOk

`func (o *MicrosoftGraphResourceVisualization) GetContainerWebUrlOk() (*string, bool)`

GetContainerWebUrlOk returns a tuple with the ContainerWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerWebUrl

`func (o *MicrosoftGraphResourceVisualization) SetContainerWebUrl(v string)`

SetContainerWebUrl sets ContainerWebUrl field to given value.

### HasContainerWebUrl

`func (o *MicrosoftGraphResourceVisualization) HasContainerWebUrl() bool`

HasContainerWebUrl returns a boolean if a field has been set.

### SetContainerWebUrlNil

`func (o *MicrosoftGraphResourceVisualization) SetContainerWebUrlNil(b bool)`

 SetContainerWebUrlNil sets the value for ContainerWebUrl to be an explicit nil

### UnsetContainerWebUrl
`func (o *MicrosoftGraphResourceVisualization) UnsetContainerWebUrl()`

UnsetContainerWebUrl ensures that no value is present for ContainerWebUrl, not even an explicit nil
### GetMediaType

`func (o *MicrosoftGraphResourceVisualization) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MicrosoftGraphResourceVisualization) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MicrosoftGraphResourceVisualization) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *MicrosoftGraphResourceVisualization) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *MicrosoftGraphResourceVisualization) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *MicrosoftGraphResourceVisualization) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetPreviewImageUrl

`func (o *MicrosoftGraphResourceVisualization) GetPreviewImageUrl() string`

GetPreviewImageUrl returns the PreviewImageUrl field if non-nil, zero value otherwise.

### GetPreviewImageUrlOk

`func (o *MicrosoftGraphResourceVisualization) GetPreviewImageUrlOk() (*string, bool)`

GetPreviewImageUrlOk returns a tuple with the PreviewImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImageUrl

`func (o *MicrosoftGraphResourceVisualization) SetPreviewImageUrl(v string)`

SetPreviewImageUrl sets PreviewImageUrl field to given value.

### HasPreviewImageUrl

`func (o *MicrosoftGraphResourceVisualization) HasPreviewImageUrl() bool`

HasPreviewImageUrl returns a boolean if a field has been set.

### SetPreviewImageUrlNil

`func (o *MicrosoftGraphResourceVisualization) SetPreviewImageUrlNil(b bool)`

 SetPreviewImageUrlNil sets the value for PreviewImageUrl to be an explicit nil

### UnsetPreviewImageUrl
`func (o *MicrosoftGraphResourceVisualization) UnsetPreviewImageUrl()`

UnsetPreviewImageUrl ensures that no value is present for PreviewImageUrl, not even an explicit nil
### GetPreviewText

`func (o *MicrosoftGraphResourceVisualization) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *MicrosoftGraphResourceVisualization) GetPreviewTextOk() (*string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *MicrosoftGraphResourceVisualization) SetPreviewText(v string)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *MicrosoftGraphResourceVisualization) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### SetPreviewTextNil

`func (o *MicrosoftGraphResourceVisualization) SetPreviewTextNil(b bool)`

 SetPreviewTextNil sets the value for PreviewText to be an explicit nil

### UnsetPreviewText
`func (o *MicrosoftGraphResourceVisualization) UnsetPreviewText()`

UnsetPreviewText ensures that no value is present for PreviewText, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphResourceVisualization) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphResourceVisualization) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphResourceVisualization) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphResourceVisualization) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphResourceVisualization) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphResourceVisualization) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *MicrosoftGraphResourceVisualization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphResourceVisualization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphResourceVisualization) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphResourceVisualization) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphResourceVisualization) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphResourceVisualization) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


