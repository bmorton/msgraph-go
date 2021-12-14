# InlineObject581

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](AnyOfmicrosoftGraphRecipient.md) |  | [optional] 
**Message** | Pointer to [**AnyOfmicrosoftGraphMessage**](anyOf&lt;microsoft.graph.message&gt;.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject581

`func NewInlineObject581() *InlineObject581`

NewInlineObject581 instantiates a new InlineObject581 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject581WithDefaults

`func NewInlineObject581WithDefaults() *InlineObject581`

NewInlineObject581WithDefaults instantiates a new InlineObject581 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToRecipients

`func (o *InlineObject581) GetToRecipients() []*AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *InlineObject581) GetToRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRecipients

`func (o *InlineObject581) SetToRecipients(v []*AnyOfmicrosoftGraphRecipient)`

SetToRecipients sets ToRecipients field to given value.

### HasToRecipients

`func (o *InlineObject581) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### GetMessage

`func (o *InlineObject581) GetMessage() AnyOfmicrosoftGraphMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject581) GetMessageOk() (*AnyOfmicrosoftGraphMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject581) SetMessage(v AnyOfmicrosoftGraphMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineObject581) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InlineObject581) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InlineObject581) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetComment

`func (o *InlineObject581) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject581) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineObject581) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineObject581) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *InlineObject581) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *InlineObject581) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


