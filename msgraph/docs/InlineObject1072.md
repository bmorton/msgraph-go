# InlineObject1072

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) |  | [optional] 
**Message** | Pointer to [**AnyOfmicrosoftGraphMessage**](anyOf&lt;microsoft.graph.message&gt;.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject1072

`func NewInlineObject1072() *InlineObject1072`

NewInlineObject1072 instantiates a new InlineObject1072 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1072WithDefaults

`func NewInlineObject1072WithDefaults() *InlineObject1072`

NewInlineObject1072WithDefaults instantiates a new InlineObject1072 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToRecipients

`func (o *InlineObject1072) GetToRecipients() []*AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *InlineObject1072) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRecipients

`func (o *InlineObject1072) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetToRecipients sets ToRecipients field to given value.

### HasToRecipients

`func (o *InlineObject1072) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### GetMessage

`func (o *InlineObject1072) GetMessage() AnyOfmicrosoftGraphMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject1072) GetMessageOk() (*AnyOfmicrosoftGraphMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject1072) SetMessage(v AnyOfmicrosoftGraphMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineObject1072) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InlineObject1072) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InlineObject1072) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetComment

`func (o *InlineObject1072) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject1072) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineObject1072) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineObject1072) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *InlineObject1072) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *InlineObject1072) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


