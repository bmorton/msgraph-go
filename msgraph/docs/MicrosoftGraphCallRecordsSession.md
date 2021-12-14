# MicrosoftGraphCallRecordsSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Callee** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that answered the session. | [optional] 
**Caller** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that initiated the session. | [optional] 
**EndDateTime** | Pointer to **time.Time** | UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**FailureInfo** | Pointer to [**AnyOfmicrosoftGraphCallRecordsFailureInfo**](anyOf&lt;microsoft.graph.callRecords.failureInfo&gt;.md) | Failure information associated with the session if the session failed. | [optional] 
**Modalities** | Pointer to [**[]AnyOfmicrosoftGraphCallRecordsModality**](AnyOfmicrosoftGraphCallRecordsModality.md) | List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue. | [optional] 
**StartDateTime** | Pointer to **time.Time** | UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Segments** | Pointer to [**[]MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md) | The list of segments involved in the session. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphCallRecordsSession

`func NewMicrosoftGraphCallRecordsSession() *MicrosoftGraphCallRecordsSession`

NewMicrosoftGraphCallRecordsSession instantiates a new MicrosoftGraphCallRecordsSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsSessionWithDefaults

`func NewMicrosoftGraphCallRecordsSessionWithDefaults() *MicrosoftGraphCallRecordsSession`

NewMicrosoftGraphCallRecordsSessionWithDefaults instantiates a new MicrosoftGraphCallRecordsSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphCallRecordsSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCallRecordsSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCallRecordsSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCallRecordsSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCallee

`func (o *MicrosoftGraphCallRecordsSession) GetCallee() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCallee returns the Callee field if non-nil, zero value otherwise.

### GetCalleeOk

`func (o *MicrosoftGraphCallRecordsSession) GetCalleeOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCalleeOk returns a tuple with the Callee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallee

`func (o *MicrosoftGraphCallRecordsSession) SetCallee(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCallee sets Callee field to given value.

### HasCallee

`func (o *MicrosoftGraphCallRecordsSession) HasCallee() bool`

HasCallee returns a boolean if a field has been set.

### SetCalleeNil

`func (o *MicrosoftGraphCallRecordsSession) SetCalleeNil(b bool)`

 SetCalleeNil sets the value for Callee to be an explicit nil

### UnsetCallee
`func (o *MicrosoftGraphCallRecordsSession) UnsetCallee()`

UnsetCallee ensures that no value is present for Callee, not even an explicit nil
### GetCaller

`func (o *MicrosoftGraphCallRecordsSession) GetCaller() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MicrosoftGraphCallRecordsSession) GetCallerOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MicrosoftGraphCallRecordsSession) SetCaller(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MicrosoftGraphCallRecordsSession) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MicrosoftGraphCallRecordsSession) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MicrosoftGraphCallRecordsSession) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphCallRecordsSession) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphCallRecordsSession) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphCallRecordsSession) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphCallRecordsSession) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetFailureInfo

`func (o *MicrosoftGraphCallRecordsSession) GetFailureInfo() AnyOfmicrosoftGraphCallRecordsFailureInfo`

GetFailureInfo returns the FailureInfo field if non-nil, zero value otherwise.

### GetFailureInfoOk

`func (o *MicrosoftGraphCallRecordsSession) GetFailureInfoOk() (*AnyOfmicrosoftGraphCallRecordsFailureInfo, bool)`

GetFailureInfoOk returns a tuple with the FailureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureInfo

`func (o *MicrosoftGraphCallRecordsSession) SetFailureInfo(v AnyOfmicrosoftGraphCallRecordsFailureInfo)`

SetFailureInfo sets FailureInfo field to given value.

### HasFailureInfo

`func (o *MicrosoftGraphCallRecordsSession) HasFailureInfo() bool`

HasFailureInfo returns a boolean if a field has been set.

### SetFailureInfoNil

`func (o *MicrosoftGraphCallRecordsSession) SetFailureInfoNil(b bool)`

 SetFailureInfoNil sets the value for FailureInfo to be an explicit nil

### UnsetFailureInfo
`func (o *MicrosoftGraphCallRecordsSession) UnsetFailureInfo()`

UnsetFailureInfo ensures that no value is present for FailureInfo, not even an explicit nil
### GetModalities

`func (o *MicrosoftGraphCallRecordsSession) GetModalities() []AnyOfmicrosoftGraphCallRecordsModality`

GetModalities returns the Modalities field if non-nil, zero value otherwise.

### GetModalitiesOk

`func (o *MicrosoftGraphCallRecordsSession) GetModalitiesOk() (*[]AnyOfmicrosoftGraphCallRecordsModality, bool)`

GetModalitiesOk returns a tuple with the Modalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalities

`func (o *MicrosoftGraphCallRecordsSession) SetModalities(v []AnyOfmicrosoftGraphCallRecordsModality)`

SetModalities sets Modalities field to given value.

### HasModalities

`func (o *MicrosoftGraphCallRecordsSession) HasModalities() bool`

HasModalities returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphCallRecordsSession) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphCallRecordsSession) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphCallRecordsSession) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphCallRecordsSession) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetSegments

`func (o *MicrosoftGraphCallRecordsSession) GetSegments() []MicrosoftGraphCallRecordsSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *MicrosoftGraphCallRecordsSession) GetSegmentsOk() (*[]MicrosoftGraphCallRecordsSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *MicrosoftGraphCallRecordsSession) SetSegments(v []MicrosoftGraphCallRecordsSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *MicrosoftGraphCallRecordsSession) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


