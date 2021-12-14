# MicrosoftGraphMeetingParticipantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity information of the participant. | [optional] 
**Role** | Pointer to [**AnyOfmicrosoftGraphOnlineMeetingRole**](anyOf&lt;microsoft.graph.onlineMeetingRole&gt;.md) | Specifies the participant&#39;s role in the meeting.  Possible values are attendee, presenter, producer, and unknownFutureValue. | [optional] 
**Upn** | Pointer to **NullableString** | User principal name of the participant. | [optional] 

## Methods

### NewMicrosoftGraphMeetingParticipantInfo

`func NewMicrosoftGraphMeetingParticipantInfo() *MicrosoftGraphMeetingParticipantInfo`

NewMicrosoftGraphMeetingParticipantInfo instantiates a new MicrosoftGraphMeetingParticipantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMeetingParticipantInfoWithDefaults

`func NewMicrosoftGraphMeetingParticipantInfoWithDefaults() *MicrosoftGraphMeetingParticipantInfo`

NewMicrosoftGraphMeetingParticipantInfoWithDefaults instantiates a new MicrosoftGraphMeetingParticipantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *MicrosoftGraphMeetingParticipantInfo) GetIdentity() AnyOfmicrosoftGraphIdentitySet`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MicrosoftGraphMeetingParticipantInfo) GetIdentityOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MicrosoftGraphMeetingParticipantInfo) SetIdentity(v AnyOfmicrosoftGraphIdentitySet)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MicrosoftGraphMeetingParticipantInfo) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *MicrosoftGraphMeetingParticipantInfo) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *MicrosoftGraphMeetingParticipantInfo) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetRole

`func (o *MicrosoftGraphMeetingParticipantInfo) GetRole() AnyOfmicrosoftGraphOnlineMeetingRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MicrosoftGraphMeetingParticipantInfo) GetRoleOk() (*AnyOfmicrosoftGraphOnlineMeetingRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MicrosoftGraphMeetingParticipantInfo) SetRole(v AnyOfmicrosoftGraphOnlineMeetingRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *MicrosoftGraphMeetingParticipantInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *MicrosoftGraphMeetingParticipantInfo) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *MicrosoftGraphMeetingParticipantInfo) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetUpn

`func (o *MicrosoftGraphMeetingParticipantInfo) GetUpn() string`

GetUpn returns the Upn field if non-nil, zero value otherwise.

### GetUpnOk

`func (o *MicrosoftGraphMeetingParticipantInfo) GetUpnOk() (*string, bool)`

GetUpnOk returns a tuple with the Upn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpn

`func (o *MicrosoftGraphMeetingParticipantInfo) SetUpn(v string)`

SetUpn sets Upn field to given value.

### HasUpn

`func (o *MicrosoftGraphMeetingParticipantInfo) HasUpn() bool`

HasUpn returns a boolean if a field has been set.

### SetUpnNil

`func (o *MicrosoftGraphMeetingParticipantInfo) SetUpnNil(b bool)`

 SetUpnNil sets the value for Upn to be an explicit nil

### UnsetUpn
`func (o *MicrosoftGraphMeetingParticipantInfo) UnsetUpn()`

UnsetUpn ensures that no value is present for Upn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


