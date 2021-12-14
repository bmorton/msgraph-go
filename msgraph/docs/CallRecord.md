# CallRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **time.Time** | UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**JoinWebUrl** | Pointer to **NullableString** | Meeting URL associated to the call. May not be available for a peerToPeer call record type. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Modalities** | Pointer to [**[]AnyOfmicrosoftGraphCallRecordsModality**](AnyOfmicrosoftGraphCallRecordsModality.md) | List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue. | [optional] 
**Organizer** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The organizing party&#39;s identity. | [optional] 
**Participants** | Pointer to [**[]AnyOfmicrosoftGraphIdentitySet**](AnyOfmicrosoftGraphIdentitySet.md) | List of distinct identities involved in the call. | [optional] 
**StartDateTime** | Pointer to **time.Time** | UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphCallRecordsCallType**](anyOf&lt;microsoft.graph.callRecords.callType&gt;.md) | Indicates the type of the call. Possible values are: unknown, groupCall, peerToPeer, unknownFutureValue. | [optional] 
**Version** | Pointer to **int64** | Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version. | [optional] 
**Sessions** | Pointer to [**[]MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md) | List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable. | [optional] 

## Methods

### NewCallRecord

`func NewCallRecord() *CallRecord`

NewCallRecord instantiates a new CallRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordWithDefaults

`func NewCallRecordWithDefaults() *CallRecord`

NewCallRecordWithDefaults instantiates a new CallRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *CallRecord) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CallRecord) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CallRecord) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CallRecord) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetJoinWebUrl

`func (o *CallRecord) GetJoinWebUrl() string`

GetJoinWebUrl returns the JoinWebUrl field if non-nil, zero value otherwise.

### GetJoinWebUrlOk

`func (o *CallRecord) GetJoinWebUrlOk() (*string, bool)`

GetJoinWebUrlOk returns a tuple with the JoinWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinWebUrl

`func (o *CallRecord) SetJoinWebUrl(v string)`

SetJoinWebUrl sets JoinWebUrl field to given value.

### HasJoinWebUrl

`func (o *CallRecord) HasJoinWebUrl() bool`

HasJoinWebUrl returns a boolean if a field has been set.

### SetJoinWebUrlNil

`func (o *CallRecord) SetJoinWebUrlNil(b bool)`

 SetJoinWebUrlNil sets the value for JoinWebUrl to be an explicit nil

### UnsetJoinWebUrl
`func (o *CallRecord) UnsetJoinWebUrl()`

UnsetJoinWebUrl ensures that no value is present for JoinWebUrl, not even an explicit nil
### GetLastModifiedDateTime

`func (o *CallRecord) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *CallRecord) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *CallRecord) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *CallRecord) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetModalities

`func (o *CallRecord) GetModalities() []AnyOfmicrosoftGraphCallRecordsModality`

GetModalities returns the Modalities field if non-nil, zero value otherwise.

### GetModalitiesOk

`func (o *CallRecord) GetModalitiesOk() (*[]AnyOfmicrosoftGraphCallRecordsModality, bool)`

GetModalitiesOk returns a tuple with the Modalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalities

`func (o *CallRecord) SetModalities(v []AnyOfmicrosoftGraphCallRecordsModality)`

SetModalities sets Modalities field to given value.

### HasModalities

`func (o *CallRecord) HasModalities() bool`

HasModalities returns a boolean if a field has been set.

### GetOrganizer

`func (o *CallRecord) GetOrganizer() AnyOfmicrosoftGraphIdentitySet`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *CallRecord) GetOrganizerOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizer

`func (o *CallRecord) SetOrganizer(v AnyOfmicrosoftGraphIdentitySet)`

SetOrganizer sets Organizer field to given value.

### HasOrganizer

`func (o *CallRecord) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### SetOrganizerNil

`func (o *CallRecord) SetOrganizerNil(b bool)`

 SetOrganizerNil sets the value for Organizer to be an explicit nil

### UnsetOrganizer
`func (o *CallRecord) UnsetOrganizer()`

UnsetOrganizer ensures that no value is present for Organizer, not even an explicit nil
### GetParticipants

`func (o *CallRecord) GetParticipants() []*AnyOfmicrosoftGraphIdentitySet`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *CallRecord) GetParticipantsOk() (*[]*AnyOfmicrosoftGraphIdentitySet, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *CallRecord) SetParticipants(v []*AnyOfmicrosoftGraphIdentitySet)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *CallRecord) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetStartDateTime

`func (o *CallRecord) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CallRecord) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CallRecord) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CallRecord) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetType

`func (o *CallRecord) GetType() AnyOfmicrosoftGraphCallRecordsCallType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CallRecord) GetTypeOk() (*AnyOfmicrosoftGraphCallRecordsCallType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CallRecord) SetType(v AnyOfmicrosoftGraphCallRecordsCallType)`

SetType sets Type field to given value.

### HasType

`func (o *CallRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CallRecord) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CallRecord) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVersion

`func (o *CallRecord) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CallRecord) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CallRecord) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CallRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSessions

`func (o *CallRecord) GetSessions() []MicrosoftGraphCallRecordsSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *CallRecord) GetSessionsOk() (*[]MicrosoftGraphCallRecordsSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *CallRecord) SetSessions(v []MicrosoftGraphCallRecordsSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *CallRecord) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


