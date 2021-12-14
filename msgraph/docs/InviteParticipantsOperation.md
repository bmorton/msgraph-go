# InviteParticipantsOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participants** | Pointer to [**[]MicrosoftGraphInvitationParticipantInfo**](MicrosoftGraphInvitationParticipantInfo.md) | The participants to invite. | [optional] 

## Methods

### NewInviteParticipantsOperation

`func NewInviteParticipantsOperation() *InviteParticipantsOperation`

NewInviteParticipantsOperation instantiates a new InviteParticipantsOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteParticipantsOperationWithDefaults

`func NewInviteParticipantsOperationWithDefaults() *InviteParticipantsOperation`

NewInviteParticipantsOperationWithDefaults instantiates a new InviteParticipantsOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipants

`func (o *InviteParticipantsOperation) GetParticipants() []MicrosoftGraphInvitationParticipantInfo`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *InviteParticipantsOperation) GetParticipantsOk() (*[]MicrosoftGraphInvitationParticipantInfo, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *InviteParticipantsOperation) SetParticipants(v []MicrosoftGraphInvitationParticipantInfo)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *InviteParticipantsOperation) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


