# MicrosoftGraphCallRecordsCallRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphCallRecordsCallRecord

`func NewMicrosoftGraphCallRecordsCallRecord() *MicrosoftGraphCallRecordsCallRecord`

NewMicrosoftGraphCallRecordsCallRecord instantiates a new MicrosoftGraphCallRecordsCallRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsCallRecordWithDefaults

`func NewMicrosoftGraphCallRecordsCallRecordWithDefaults() *MicrosoftGraphCallRecordsCallRecord`

NewMicrosoftGraphCallRecordsCallRecordWithDefaults instantiates a new MicrosoftGraphCallRecordsCallRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphCallRecordsCallRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCallRecordsCallRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCallRecordsCallRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetJoinWebUrl

`func (o *MicrosoftGraphCallRecordsCallRecord) GetJoinWebUrl() string`

GetJoinWebUrl returns the JoinWebUrl field if non-nil, zero value otherwise.

### GetJoinWebUrlOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetJoinWebUrlOk() (*string, bool)`

GetJoinWebUrlOk returns a tuple with the JoinWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinWebUrl

`func (o *MicrosoftGraphCallRecordsCallRecord) SetJoinWebUrl(v string)`

SetJoinWebUrl sets JoinWebUrl field to given value.

### HasJoinWebUrl

`func (o *MicrosoftGraphCallRecordsCallRecord) HasJoinWebUrl() bool`

HasJoinWebUrl returns a boolean if a field has been set.

### SetJoinWebUrlNil

`func (o *MicrosoftGraphCallRecordsCallRecord) SetJoinWebUrlNil(b bool)`

 SetJoinWebUrlNil sets the value for JoinWebUrl to be an explicit nil

### UnsetJoinWebUrl
`func (o *MicrosoftGraphCallRecordsCallRecord) UnsetJoinWebUrl()`

UnsetJoinWebUrl ensures that no value is present for JoinWebUrl, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetModalities

`func (o *MicrosoftGraphCallRecordsCallRecord) GetModalities() []AnyOfmicrosoftGraphCallRecordsModality`

GetModalities returns the Modalities field if non-nil, zero value otherwise.

### GetModalitiesOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetModalitiesOk() (*[]AnyOfmicrosoftGraphCallRecordsModality, bool)`

GetModalitiesOk returns a tuple with the Modalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalities

`func (o *MicrosoftGraphCallRecordsCallRecord) SetModalities(v []AnyOfmicrosoftGraphCallRecordsModality)`

SetModalities sets Modalities field to given value.

### HasModalities

`func (o *MicrosoftGraphCallRecordsCallRecord) HasModalities() bool`

HasModalities returns a boolean if a field has been set.

### GetOrganizer

`func (o *MicrosoftGraphCallRecordsCallRecord) GetOrganizer() AnyOfmicrosoftGraphIdentitySet`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetOrganizerOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizer

`func (o *MicrosoftGraphCallRecordsCallRecord) SetOrganizer(v AnyOfmicrosoftGraphIdentitySet)`

SetOrganizer sets Organizer field to given value.

### HasOrganizer

`func (o *MicrosoftGraphCallRecordsCallRecord) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### SetOrganizerNil

`func (o *MicrosoftGraphCallRecordsCallRecord) SetOrganizerNil(b bool)`

 SetOrganizerNil sets the value for Organizer to be an explicit nil

### UnsetOrganizer
`func (o *MicrosoftGraphCallRecordsCallRecord) UnsetOrganizer()`

UnsetOrganizer ensures that no value is present for Organizer, not even an explicit nil
### GetParticipants

`func (o *MicrosoftGraphCallRecordsCallRecord) GetParticipants() []*AnyOfmicrosoftGraphIdentitySet`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetParticipantsOk() (*[]*AnyOfmicrosoftGraphIdentitySet, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *MicrosoftGraphCallRecordsCallRecord) SetParticipants(v []*AnyOfmicrosoftGraphIdentitySet)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *MicrosoftGraphCallRecordsCallRecord) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphCallRecordsCallRecord) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftGraphCallRecordsCallRecord) GetType() AnyOfmicrosoftGraphCallRecordsCallType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetTypeOk() (*AnyOfmicrosoftGraphCallRecordsCallType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphCallRecordsCallRecord) SetType(v AnyOfmicrosoftGraphCallRecordsCallType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphCallRecordsCallRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphCallRecordsCallRecord) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphCallRecordsCallRecord) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphCallRecordsCallRecord) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphCallRecordsCallRecord) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphCallRecordsCallRecord) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSessions

`func (o *MicrosoftGraphCallRecordsCallRecord) GetSessions() []MicrosoftGraphCallRecordsSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *MicrosoftGraphCallRecordsCallRecord) GetSessionsOk() (*[]MicrosoftGraphCallRecordsSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *MicrosoftGraphCallRecordsCallRecord) SetSessions(v []MicrosoftGraphCallRecordsSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *MicrosoftGraphCallRecordsCallRecord) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


