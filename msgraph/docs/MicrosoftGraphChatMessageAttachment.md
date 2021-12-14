# MicrosoftGraphChatMessageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | The content of the attachment. If the attachment is a rich card, set the property to the rich card object. This property and contentUrl are mutually exclusive. | [optional] 
**ContentType** | Pointer to **NullableString** | The media type of the content attachment. It can have the following values: reference: Attachment is a link to another file. Populate the contentURL with the link to the object.Any contentTypes supported by the Bot Framework&#39;s Attachment objectapplication/vnd.microsoft.card.codesnippet: A code snippet. application/vnd.microsoft.card.announcement: An announcement header. | [optional] 
**ContentUrl** | Pointer to **NullableString** | URL for the content of the attachment. Supported protocols: http, https, file and data. | [optional] 
**Id** | Pointer to **NullableString** | Read-only. Unique id of the attachment. | [optional] 
**Name** | Pointer to **NullableString** | Name of the attachment. | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** | URL to a thumbnail image that the channel can use if it supports using an alternative, smaller form of content or contentUrl. For example, if you set contentType to application/word and set contentUrl to the location of the Word document, you might include a thumbnail image that represents the document. The channel could display the thumbnail image instead of the document. When the user clicks the image, the channel would open the document. | [optional] 

## Methods

### NewMicrosoftGraphChatMessageAttachment

`func NewMicrosoftGraphChatMessageAttachment() *MicrosoftGraphChatMessageAttachment`

NewMicrosoftGraphChatMessageAttachment instantiates a new MicrosoftGraphChatMessageAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessageAttachmentWithDefaults

`func NewMicrosoftGraphChatMessageAttachmentWithDefaults() *MicrosoftGraphChatMessageAttachment`

NewMicrosoftGraphChatMessageAttachmentWithDefaults instantiates a new MicrosoftGraphChatMessageAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MicrosoftGraphChatMessageAttachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphChatMessageAttachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphChatMessageAttachment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphChatMessageAttachment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphChatMessageAttachment) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphChatMessageAttachment) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentType

`func (o *MicrosoftGraphChatMessageAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphChatMessageAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphChatMessageAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphChatMessageAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MicrosoftGraphChatMessageAttachment) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MicrosoftGraphChatMessageAttachment) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentUrl

`func (o *MicrosoftGraphChatMessageAttachment) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *MicrosoftGraphChatMessageAttachment) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *MicrosoftGraphChatMessageAttachment) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *MicrosoftGraphChatMessageAttachment) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrlNil

`func (o *MicrosoftGraphChatMessageAttachment) SetContentUrlNil(b bool)`

 SetContentUrlNil sets the value for ContentUrl to be an explicit nil

### UnsetContentUrl
`func (o *MicrosoftGraphChatMessageAttachment) UnsetContentUrl()`

UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
### GetId

`func (o *MicrosoftGraphChatMessageAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphChatMessageAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphChatMessageAttachment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphChatMessageAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphChatMessageAttachment) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphChatMessageAttachment) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *MicrosoftGraphChatMessageAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphChatMessageAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphChatMessageAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphChatMessageAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphChatMessageAttachment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphChatMessageAttachment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetThumbnailUrl

`func (o *MicrosoftGraphChatMessageAttachment) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *MicrosoftGraphChatMessageAttachment) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *MicrosoftGraphChatMessageAttachment) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *MicrosoftGraphChatMessageAttachment) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *MicrosoftGraphChatMessageAttachment) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *MicrosoftGraphChatMessageAttachment) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


