# MicrosoftGraphWorkbookComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Content** | Pointer to **NullableString** | The content of comment. | [optional] 
**ContentType** | Pointer to **string** | Indicates the type for the comment. | [optional] 
**Replies** | Pointer to [**[]MicrosoftGraphWorkbookCommentReply**](MicrosoftGraphWorkbookCommentReply.md) | Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookComment

`func NewMicrosoftGraphWorkbookComment() *MicrosoftGraphWorkbookComment`

NewMicrosoftGraphWorkbookComment instantiates a new MicrosoftGraphWorkbookComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookCommentWithDefaults

`func NewMicrosoftGraphWorkbookCommentWithDefaults() *MicrosoftGraphWorkbookComment`

NewMicrosoftGraphWorkbookCommentWithDefaults instantiates a new MicrosoftGraphWorkbookComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookComment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *MicrosoftGraphWorkbookComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphWorkbookComment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphWorkbookComment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphWorkbookComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphWorkbookComment) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphWorkbookComment) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentType

`func (o *MicrosoftGraphWorkbookComment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphWorkbookComment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphWorkbookComment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphWorkbookComment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetReplies

`func (o *MicrosoftGraphWorkbookComment) GetReplies() []MicrosoftGraphWorkbookCommentReply`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *MicrosoftGraphWorkbookComment) GetRepliesOk() (*[]MicrosoftGraphWorkbookCommentReply, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *MicrosoftGraphWorkbookComment) SetReplies(v []MicrosoftGraphWorkbookCommentReply)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *MicrosoftGraphWorkbookComment) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


