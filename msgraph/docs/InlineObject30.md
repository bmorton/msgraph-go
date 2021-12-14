# InlineObject30

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferTarget** | Pointer to [**MicrosoftGraphInvitationParticipantInfo**](MicrosoftGraphInvitationParticipantInfo.md) |  | [optional] 
**Transferee** | Pointer to [**AnyOfmicrosoftGraphParticipantInfo**](anyOf&lt;microsoft.graph.participantInfo&gt;.md) |  | [optional] 

## Methods

### NewInlineObject30

`func NewInlineObject30() *InlineObject30`

NewInlineObject30 instantiates a new InlineObject30 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject30WithDefaults

`func NewInlineObject30WithDefaults() *InlineObject30`

NewInlineObject30WithDefaults instantiates a new InlineObject30 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferTarget

`func (o *InlineObject30) GetTransferTarget() MicrosoftGraphInvitationParticipantInfo`

GetTransferTarget returns the TransferTarget field if non-nil, zero value otherwise.

### GetTransferTargetOk

`func (o *InlineObject30) GetTransferTargetOk() (*MicrosoftGraphInvitationParticipantInfo, bool)`

GetTransferTargetOk returns a tuple with the TransferTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTarget

`func (o *InlineObject30) SetTransferTarget(v MicrosoftGraphInvitationParticipantInfo)`

SetTransferTarget sets TransferTarget field to given value.

### HasTransferTarget

`func (o *InlineObject30) HasTransferTarget() bool`

HasTransferTarget returns a boolean if a field has been set.

### GetTransferee

`func (o *InlineObject30) GetTransferee() AnyOfmicrosoftGraphParticipantInfo`

GetTransferee returns the Transferee field if non-nil, zero value otherwise.

### GetTransfereeOk

`func (o *InlineObject30) GetTransfereeOk() (*AnyOfmicrosoftGraphParticipantInfo, bool)`

GetTransfereeOk returns a tuple with the Transferee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferee

`func (o *InlineObject30) SetTransferee(v AnyOfmicrosoftGraphParticipantInfo)`

SetTransferee sets Transferee field to given value.

### HasTransferee

`func (o *InlineObject30) HasTransferee() bool`

HasTransferee returns a boolean if a field has been set.

### SetTransfereeNil

`func (o *InlineObject30) SetTransfereeNil(b bool)`

 SetTransfereeNil sets the value for Transferee to be an explicit nil

### UnsetTransferee
`func (o *InlineObject30) UnsetTransferee()`

UnsetTransferee ensures that no value is present for Transferee, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


