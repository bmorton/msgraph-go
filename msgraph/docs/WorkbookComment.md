# WorkbookComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | The content of comment. | [optional] 
**ContentType** | Pointer to **string** | Indicates the type for the comment. | [optional] 
**Replies** | Pointer to [**[]MicrosoftGraphWorkbookCommentReply**](MicrosoftGraphWorkbookCommentReply.md) | Read-only. Nullable. | [optional] 

## Methods

### NewWorkbookComment

`func NewWorkbookComment() *WorkbookComment`

NewWorkbookComment instantiates a new WorkbookComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookCommentWithDefaults

`func NewWorkbookCommentWithDefaults() *WorkbookComment`

NewWorkbookCommentWithDefaults instantiates a new WorkbookComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *WorkbookComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WorkbookComment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WorkbookComment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WorkbookComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WorkbookComment) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WorkbookComment) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentType

`func (o *WorkbookComment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkbookComment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WorkbookComment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WorkbookComment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetReplies

`func (o *WorkbookComment) GetReplies() []MicrosoftGraphWorkbookCommentReply`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *WorkbookComment) GetRepliesOk() (*[]MicrosoftGraphWorkbookCommentReply, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *WorkbookComment) SetReplies(v []MicrosoftGraphWorkbookCommentReply)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *WorkbookComment) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


