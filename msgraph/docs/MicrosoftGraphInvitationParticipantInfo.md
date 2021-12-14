# MicrosoftGraphInvitationParticipantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**MicrosoftGraphIdentitySet**](MicrosoftGraphIdentitySet.md) |  | [optional] 
**ReplacesCallId** | Pointer to **NullableString** | Optional. The call which the target identity is currently a part of. This call will be dropped once the participant is added. | [optional] 

## Methods

### NewMicrosoftGraphInvitationParticipantInfo

`func NewMicrosoftGraphInvitationParticipantInfo() *MicrosoftGraphInvitationParticipantInfo`

NewMicrosoftGraphInvitationParticipantInfo instantiates a new MicrosoftGraphInvitationParticipantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphInvitationParticipantInfoWithDefaults

`func NewMicrosoftGraphInvitationParticipantInfoWithDefaults() *MicrosoftGraphInvitationParticipantInfo`

NewMicrosoftGraphInvitationParticipantInfoWithDefaults instantiates a new MicrosoftGraphInvitationParticipantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *MicrosoftGraphInvitationParticipantInfo) GetIdentity() MicrosoftGraphIdentitySet`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MicrosoftGraphInvitationParticipantInfo) GetIdentityOk() (*MicrosoftGraphIdentitySet, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MicrosoftGraphInvitationParticipantInfo) SetIdentity(v MicrosoftGraphIdentitySet)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MicrosoftGraphInvitationParticipantInfo) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetReplacesCallId

`func (o *MicrosoftGraphInvitationParticipantInfo) GetReplacesCallId() string`

GetReplacesCallId returns the ReplacesCallId field if non-nil, zero value otherwise.

### GetReplacesCallIdOk

`func (o *MicrosoftGraphInvitationParticipantInfo) GetReplacesCallIdOk() (*string, bool)`

GetReplacesCallIdOk returns a tuple with the ReplacesCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacesCallId

`func (o *MicrosoftGraphInvitationParticipantInfo) SetReplacesCallId(v string)`

SetReplacesCallId sets ReplacesCallId field to given value.

### HasReplacesCallId

`func (o *MicrosoftGraphInvitationParticipantInfo) HasReplacesCallId() bool`

HasReplacesCallId returns a boolean if a field has been set.

### SetReplacesCallIdNil

`func (o *MicrosoftGraphInvitationParticipantInfo) SetReplacesCallIdNil(b bool)`

 SetReplacesCallIdNil sets the value for ReplacesCallId to be an explicit nil

### UnsetReplacesCallId
`func (o *MicrosoftGraphInvitationParticipantInfo) UnsetReplacesCallId()`

UnsetReplacesCallId ensures that no value is present for ReplacesCallId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


