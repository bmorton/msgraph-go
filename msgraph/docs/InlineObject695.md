# InlineObject695

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatInfo** | Pointer to [**AnyOfmicrosoftGraphChatInfo**](anyOf&lt;microsoft.graph.chatInfo&gt;.md) |  | [optional] 
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to [**AnyOfmicrosoftGraphMeetingParticipants**](anyOf&lt;microsoft.graph.meetingParticipants&gt;.md) |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject695

`func NewInlineObject695() *InlineObject695`

NewInlineObject695 instantiates a new InlineObject695 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject695WithDefaults

`func NewInlineObject695WithDefaults() *InlineObject695`

NewInlineObject695WithDefaults instantiates a new InlineObject695 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatInfo

`func (o *InlineObject695) GetChatInfo() AnyOfmicrosoftGraphChatInfo`

GetChatInfo returns the ChatInfo field if non-nil, zero value otherwise.

### GetChatInfoOk

`func (o *InlineObject695) GetChatInfoOk() (*AnyOfmicrosoftGraphChatInfo, bool)`

GetChatInfoOk returns a tuple with the ChatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatInfo

`func (o *InlineObject695) SetChatInfo(v AnyOfmicrosoftGraphChatInfo)`

SetChatInfo sets ChatInfo field to given value.

### HasChatInfo

`func (o *InlineObject695) HasChatInfo() bool`

HasChatInfo returns a boolean if a field has been set.

### SetChatInfoNil

`func (o *InlineObject695) SetChatInfoNil(b bool)`

 SetChatInfoNil sets the value for ChatInfo to be an explicit nil

### UnsetChatInfo
`func (o *InlineObject695) UnsetChatInfo()`

UnsetChatInfo ensures that no value is present for ChatInfo, not even an explicit nil
### GetEndDateTime

`func (o *InlineObject695) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *InlineObject695) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *InlineObject695) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *InlineObject695) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *InlineObject695) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *InlineObject695) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetExternalId

`func (o *InlineObject695) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineObject695) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineObject695) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineObject695) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetParticipants

`func (o *InlineObject695) GetParticipants() AnyOfmicrosoftGraphMeetingParticipants`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *InlineObject695) GetParticipantsOk() (*AnyOfmicrosoftGraphMeetingParticipants, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *InlineObject695) SetParticipants(v AnyOfmicrosoftGraphMeetingParticipants)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *InlineObject695) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### SetParticipantsNil

`func (o *InlineObject695) SetParticipantsNil(b bool)`

 SetParticipantsNil sets the value for Participants to be an explicit nil

### UnsetParticipants
`func (o *InlineObject695) UnsetParticipants()`

UnsetParticipants ensures that no value is present for Participants, not even an explicit nil
### GetStartDateTime

`func (o *InlineObject695) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *InlineObject695) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *InlineObject695) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *InlineObject695) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *InlineObject695) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *InlineObject695) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetSubject

`func (o *InlineObject695) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InlineObject695) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InlineObject695) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *InlineObject695) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *InlineObject695) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *InlineObject695) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


