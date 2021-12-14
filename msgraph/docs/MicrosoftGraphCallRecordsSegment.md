# MicrosoftGraphCallRecordsSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Callee** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that answered this segment. | [optional] 
**Caller** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that initiated this segment. | [optional] 
**EndDateTime** | Pointer to **time.Time** | UTC time when the segment ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**FailureInfo** | Pointer to [**AnyOfmicrosoftGraphCallRecordsFailureInfo**](anyOf&lt;microsoft.graph.callRecords.failureInfo&gt;.md) | Failure information associated with the segment if it failed. | [optional] 
**Media** | Pointer to [**[]AnyOfmicrosoftGraphCallRecordsMedia**](AnyOfmicrosoftGraphCallRecordsMedia.md) | Media associated with this segment. | [optional] 
**StartDateTime** | Pointer to **time.Time** | UTC time when the segment started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 

## Methods

### NewMicrosoftGraphCallRecordsSegment

`func NewMicrosoftGraphCallRecordsSegment() *MicrosoftGraphCallRecordsSegment`

NewMicrosoftGraphCallRecordsSegment instantiates a new MicrosoftGraphCallRecordsSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsSegmentWithDefaults

`func NewMicrosoftGraphCallRecordsSegmentWithDefaults() *MicrosoftGraphCallRecordsSegment`

NewMicrosoftGraphCallRecordsSegmentWithDefaults instantiates a new MicrosoftGraphCallRecordsSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphCallRecordsSegment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCallRecordsSegment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCallRecordsSegment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCallRecordsSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCallee

`func (o *MicrosoftGraphCallRecordsSegment) GetCallee() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCallee returns the Callee field if non-nil, zero value otherwise.

### GetCalleeOk

`func (o *MicrosoftGraphCallRecordsSegment) GetCalleeOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCalleeOk returns a tuple with the Callee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallee

`func (o *MicrosoftGraphCallRecordsSegment) SetCallee(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCallee sets Callee field to given value.

### HasCallee

`func (o *MicrosoftGraphCallRecordsSegment) HasCallee() bool`

HasCallee returns a boolean if a field has been set.

### SetCalleeNil

`func (o *MicrosoftGraphCallRecordsSegment) SetCalleeNil(b bool)`

 SetCalleeNil sets the value for Callee to be an explicit nil

### UnsetCallee
`func (o *MicrosoftGraphCallRecordsSegment) UnsetCallee()`

UnsetCallee ensures that no value is present for Callee, not even an explicit nil
### GetCaller

`func (o *MicrosoftGraphCallRecordsSegment) GetCaller() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MicrosoftGraphCallRecordsSegment) GetCallerOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MicrosoftGraphCallRecordsSegment) SetCaller(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MicrosoftGraphCallRecordsSegment) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MicrosoftGraphCallRecordsSegment) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MicrosoftGraphCallRecordsSegment) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphCallRecordsSegment) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphCallRecordsSegment) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphCallRecordsSegment) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphCallRecordsSegment) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetFailureInfo

`func (o *MicrosoftGraphCallRecordsSegment) GetFailureInfo() AnyOfmicrosoftGraphCallRecordsFailureInfo`

GetFailureInfo returns the FailureInfo field if non-nil, zero value otherwise.

### GetFailureInfoOk

`func (o *MicrosoftGraphCallRecordsSegment) GetFailureInfoOk() (*AnyOfmicrosoftGraphCallRecordsFailureInfo, bool)`

GetFailureInfoOk returns a tuple with the FailureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureInfo

`func (o *MicrosoftGraphCallRecordsSegment) SetFailureInfo(v AnyOfmicrosoftGraphCallRecordsFailureInfo)`

SetFailureInfo sets FailureInfo field to given value.

### HasFailureInfo

`func (o *MicrosoftGraphCallRecordsSegment) HasFailureInfo() bool`

HasFailureInfo returns a boolean if a field has been set.

### SetFailureInfoNil

`func (o *MicrosoftGraphCallRecordsSegment) SetFailureInfoNil(b bool)`

 SetFailureInfoNil sets the value for FailureInfo to be an explicit nil

### UnsetFailureInfo
`func (o *MicrosoftGraphCallRecordsSegment) UnsetFailureInfo()`

UnsetFailureInfo ensures that no value is present for FailureInfo, not even an explicit nil
### GetMedia

`func (o *MicrosoftGraphCallRecordsSegment) GetMedia() []*AnyOfmicrosoftGraphCallRecordsMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MicrosoftGraphCallRecordsSegment) GetMediaOk() (*[]*AnyOfmicrosoftGraphCallRecordsMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MicrosoftGraphCallRecordsSegment) SetMedia(v []*AnyOfmicrosoftGraphCallRecordsMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MicrosoftGraphCallRecordsSegment) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphCallRecordsSegment) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphCallRecordsSegment) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphCallRecordsSegment) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphCallRecordsSegment) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


