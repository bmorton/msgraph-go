# MicrosoftGraphInviteParticipantsOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClientContext** | Pointer to **NullableString** | Unique Client Context string. Max limit is 256 chars. | [optional] 
**ResultInfo** | Pointer to [**AnyOfmicrosoftGraphResultInfo**](anyOf&lt;microsoft.graph.resultInfo&gt;.md) | The result information. Read-only. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) | Possible values are: notStarted, running, completed, failed. Read-only. | [optional] 
**Participants** | Pointer to [**[]MicrosoftGraphInvitationParticipantInfo**](MicrosoftGraphInvitationParticipantInfo.md) | The participants to invite. | [optional] 

## Methods

### NewMicrosoftGraphInviteParticipantsOperation

`func NewMicrosoftGraphInviteParticipantsOperation() *MicrosoftGraphInviteParticipantsOperation`

NewMicrosoftGraphInviteParticipantsOperation instantiates a new MicrosoftGraphInviteParticipantsOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphInviteParticipantsOperationWithDefaults

`func NewMicrosoftGraphInviteParticipantsOperationWithDefaults() *MicrosoftGraphInviteParticipantsOperation`

NewMicrosoftGraphInviteParticipantsOperationWithDefaults instantiates a new MicrosoftGraphInviteParticipantsOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphInviteParticipantsOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphInviteParticipantsOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphInviteParticipantsOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphInviteParticipantsOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientContext

`func (o *MicrosoftGraphInviteParticipantsOperation) GetClientContext() string`

GetClientContext returns the ClientContext field if non-nil, zero value otherwise.

### GetClientContextOk

`func (o *MicrosoftGraphInviteParticipantsOperation) GetClientContextOk() (*string, bool)`

GetClientContextOk returns a tuple with the ClientContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContext

`func (o *MicrosoftGraphInviteParticipantsOperation) SetClientContext(v string)`

SetClientContext sets ClientContext field to given value.

### HasClientContext

`func (o *MicrosoftGraphInviteParticipantsOperation) HasClientContext() bool`

HasClientContext returns a boolean if a field has been set.

### SetClientContextNil

`func (o *MicrosoftGraphInviteParticipantsOperation) SetClientContextNil(b bool)`

 SetClientContextNil sets the value for ClientContext to be an explicit nil

### UnsetClientContext
`func (o *MicrosoftGraphInviteParticipantsOperation) UnsetClientContext()`

UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
### GetResultInfo

`func (o *MicrosoftGraphInviteParticipantsOperation) GetResultInfo() AnyOfmicrosoftGraphResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *MicrosoftGraphInviteParticipantsOperation) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *MicrosoftGraphInviteParticipantsOperation) SetResultInfo(v AnyOfmicrosoftGraphResultInfo)`

SetResultInfo sets ResultInfo field to given value.

### HasResultInfo

`func (o *MicrosoftGraphInviteParticipantsOperation) HasResultInfo() bool`

HasResultInfo returns a boolean if a field has been set.

### SetResultInfoNil

`func (o *MicrosoftGraphInviteParticipantsOperation) SetResultInfoNil(b bool)`

 SetResultInfoNil sets the value for ResultInfo to be an explicit nil

### UnsetResultInfo
`func (o *MicrosoftGraphInviteParticipantsOperation) UnsetResultInfo()`

UnsetResultInfo ensures that no value is present for ResultInfo, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphInviteParticipantsOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphInviteParticipantsOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphInviteParticipantsOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphInviteParticipantsOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphInviteParticipantsOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphInviteParticipantsOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetParticipants

`func (o *MicrosoftGraphInviteParticipantsOperation) GetParticipants() []MicrosoftGraphInvitationParticipantInfo`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *MicrosoftGraphInviteParticipantsOperation) GetParticipantsOk() (*[]MicrosoftGraphInvitationParticipantInfo, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *MicrosoftGraphInviteParticipantsOperation) SetParticipants(v []MicrosoftGraphInvitationParticipantInfo)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *MicrosoftGraphInviteParticipantsOperation) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


