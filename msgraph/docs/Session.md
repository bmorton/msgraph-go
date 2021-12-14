# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callee** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that answered the session. | [optional] 
**Caller** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that initiated the session. | [optional] 
**EndDateTime** | Pointer to **time.Time** | UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**FailureInfo** | Pointer to [**AnyOfmicrosoftGraphCallRecordsFailureInfo**](anyOf&lt;microsoft.graph.callRecords.failureInfo&gt;.md) | Failure information associated with the session if the session failed. | [optional] 
**Modalities** | Pointer to [**[]AnyOfmicrosoftGraphCallRecordsModality**](AnyOfmicrosoftGraphCallRecordsModality.md) | List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue. | [optional] 
**StartDateTime** | Pointer to **time.Time** | UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Segments** | Pointer to [**[]MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md) | The list of segments involved in the session. Read-only. Nullable. | [optional] 

## Methods

### NewSession

`func NewSession() *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallee

`func (o *Session) GetCallee() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCallee returns the Callee field if non-nil, zero value otherwise.

### GetCalleeOk

`func (o *Session) GetCalleeOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCalleeOk returns a tuple with the Callee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallee

`func (o *Session) SetCallee(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCallee sets Callee field to given value.

### HasCallee

`func (o *Session) HasCallee() bool`

HasCallee returns a boolean if a field has been set.

### SetCalleeNil

`func (o *Session) SetCalleeNil(b bool)`

 SetCalleeNil sets the value for Callee to be an explicit nil

### UnsetCallee
`func (o *Session) UnsetCallee()`

UnsetCallee ensures that no value is present for Callee, not even an explicit nil
### GetCaller

`func (o *Session) GetCaller() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Session) GetCallerOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Session) SetCaller(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Session) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Session) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Session) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetEndDateTime

`func (o *Session) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Session) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Session) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Session) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetFailureInfo

`func (o *Session) GetFailureInfo() AnyOfmicrosoftGraphCallRecordsFailureInfo`

GetFailureInfo returns the FailureInfo field if non-nil, zero value otherwise.

### GetFailureInfoOk

`func (o *Session) GetFailureInfoOk() (*AnyOfmicrosoftGraphCallRecordsFailureInfo, bool)`

GetFailureInfoOk returns a tuple with the FailureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureInfo

`func (o *Session) SetFailureInfo(v AnyOfmicrosoftGraphCallRecordsFailureInfo)`

SetFailureInfo sets FailureInfo field to given value.

### HasFailureInfo

`func (o *Session) HasFailureInfo() bool`

HasFailureInfo returns a boolean if a field has been set.

### SetFailureInfoNil

`func (o *Session) SetFailureInfoNil(b bool)`

 SetFailureInfoNil sets the value for FailureInfo to be an explicit nil

### UnsetFailureInfo
`func (o *Session) UnsetFailureInfo()`

UnsetFailureInfo ensures that no value is present for FailureInfo, not even an explicit nil
### GetModalities

`func (o *Session) GetModalities() []AnyOfmicrosoftGraphCallRecordsModality`

GetModalities returns the Modalities field if non-nil, zero value otherwise.

### GetModalitiesOk

`func (o *Session) GetModalitiesOk() (*[]AnyOfmicrosoftGraphCallRecordsModality, bool)`

GetModalitiesOk returns a tuple with the Modalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalities

`func (o *Session) SetModalities(v []AnyOfmicrosoftGraphCallRecordsModality)`

SetModalities sets Modalities field to given value.

### HasModalities

`func (o *Session) HasModalities() bool`

HasModalities returns a boolean if a field has been set.

### GetStartDateTime

`func (o *Session) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Session) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Session) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *Session) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetSegments

`func (o *Session) GetSegments() []MicrosoftGraphCallRecordsSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Session) GetSegmentsOk() (*[]MicrosoftGraphCallRecordsSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Session) SetSegments(v []MicrosoftGraphCallRecordsSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Session) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


