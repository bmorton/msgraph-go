# InlineObject401

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** |  | [optional] 
**SendResponse** | Pointer to **NullableBool** |  | [optional] [default to false]
**ProposedNewTime** | Pointer to [**AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) |  | [optional] 

## Methods

### NewInlineObject401

`func NewInlineObject401() *InlineObject401`

NewInlineObject401 instantiates a new InlineObject401 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject401WithDefaults

`func NewInlineObject401WithDefaults() *InlineObject401`

NewInlineObject401WithDefaults instantiates a new InlineObject401 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineObject401) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject401) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineObject401) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineObject401) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *InlineObject401) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *InlineObject401) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetSendResponse

`func (o *InlineObject401) GetSendResponse() bool`

GetSendResponse returns the SendResponse field if non-nil, zero value otherwise.

### GetSendResponseOk

`func (o *InlineObject401) GetSendResponseOk() (*bool, bool)`

GetSendResponseOk returns a tuple with the SendResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResponse

`func (o *InlineObject401) SetSendResponse(v bool)`

SetSendResponse sets SendResponse field to given value.

### HasSendResponse

`func (o *InlineObject401) HasSendResponse() bool`

HasSendResponse returns a boolean if a field has been set.

### SetSendResponseNil

`func (o *InlineObject401) SetSendResponseNil(b bool)`

 SetSendResponseNil sets the value for SendResponse to be an explicit nil

### UnsetSendResponse
`func (o *InlineObject401) UnsetSendResponse()`

UnsetSendResponse ensures that no value is present for SendResponse, not even an explicit nil
### GetProposedNewTime

`func (o *InlineObject401) GetProposedNewTime() AnyOfmicrosoftGraphTimeSlot`

GetProposedNewTime returns the ProposedNewTime field if non-nil, zero value otherwise.

### GetProposedNewTimeOk

`func (o *InlineObject401) GetProposedNewTimeOk() (*AnyOfmicrosoftGraphTimeSlot, bool)`

GetProposedNewTimeOk returns a tuple with the ProposedNewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedNewTime

`func (o *InlineObject401) SetProposedNewTime(v AnyOfmicrosoftGraphTimeSlot)`

SetProposedNewTime sets ProposedNewTime field to given value.

### HasProposedNewTime

`func (o *InlineObject401) HasProposedNewTime() bool`

HasProposedNewTime returns a boolean if a field has been set.

### SetProposedNewTimeNil

`func (o *InlineObject401) SetProposedNewTimeNil(b bool)`

 SetProposedNewTimeNil sets the value for ProposedNewTime to be an explicit nil

### UnsetProposedNewTime
`func (o *InlineObject401) UnsetProposedNewTime()`

UnsetProposedNewTime ensures that no value is present for ProposedNewTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


