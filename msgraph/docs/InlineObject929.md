# InlineObject929

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** |  | [optional] 
**SendResponse** | Pointer to **NullableBool** |  | [optional] [default to false]
**ProposedNewTime** | Pointer to [**AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) |  | [optional] 

## Methods

### NewInlineObject929

`func NewInlineObject929() *InlineObject929`

NewInlineObject929 instantiates a new InlineObject929 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject929WithDefaults

`func NewInlineObject929WithDefaults() *InlineObject929`

NewInlineObject929WithDefaults instantiates a new InlineObject929 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *InlineObject929) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject929) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InlineObject929) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InlineObject929) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *InlineObject929) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *InlineObject929) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetSendResponse

`func (o *InlineObject929) GetSendResponse() bool`

GetSendResponse returns the SendResponse field if non-nil, zero value otherwise.

### GetSendResponseOk

`func (o *InlineObject929) GetSendResponseOk() (*bool, bool)`

GetSendResponseOk returns a tuple with the SendResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResponse

`func (o *InlineObject929) SetSendResponse(v bool)`

SetSendResponse sets SendResponse field to given value.

### HasSendResponse

`func (o *InlineObject929) HasSendResponse() bool`

HasSendResponse returns a boolean if a field has been set.

### SetSendResponseNil

`func (o *InlineObject929) SetSendResponseNil(b bool)`

 SetSendResponseNil sets the value for SendResponse to be an explicit nil

### UnsetSendResponse
`func (o *InlineObject929) UnsetSendResponse()`

UnsetSendResponse ensures that no value is present for SendResponse, not even an explicit nil
### GetProposedNewTime

`func (o *InlineObject929) GetProposedNewTime() AnyOfmicrosoftGraphTimeSlot`

GetProposedNewTime returns the ProposedNewTime field if non-nil, zero value otherwise.

### GetProposedNewTimeOk

`func (o *InlineObject929) GetProposedNewTimeOk() (*AnyOfmicrosoftGraphTimeSlot, bool)`

GetProposedNewTimeOk returns a tuple with the ProposedNewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedNewTime

`func (o *InlineObject929) SetProposedNewTime(v AnyOfmicrosoftGraphTimeSlot)`

SetProposedNewTime sets ProposedNewTime field to given value.

### HasProposedNewTime

`func (o *InlineObject929) HasProposedNewTime() bool`

HasProposedNewTime returns a boolean if a field has been set.

### SetProposedNewTimeNil

`func (o *InlineObject929) SetProposedNewTimeNil(b bool)`

 SetProposedNewTimeNil sets the value for ProposedNewTime to be an explicit nil

### UnsetProposedNewTime
`func (o *InlineObject929) UnsetProposedNewTime()`

UnsetProposedNewTime ensures that no value is present for ProposedNewTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


