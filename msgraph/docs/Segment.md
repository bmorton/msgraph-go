# Segment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Callee** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that answered this segment. | [optional] 
**Caller** | Pointer to [**AnyOfmicrosoftGraphCallRecordsEndpoint**](anyOf&lt;microsoft.graph.callRecords.endpoint&gt;.md) | Endpoint that initiated this segment. | [optional] 
**EndDateTime** | Pointer to **time.Time** | UTC time when the segment ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**FailureInfo** | Pointer to [**AnyOfmicrosoftGraphCallRecordsFailureInfo**](anyOf&lt;microsoft.graph.callRecords.failureInfo&gt;.md) | Failure information associated with the segment if it failed. | [optional] 
**Media** | Pointer to [**[]AnyOfmicrosoftGraphCallRecordsMedia**](AnyOfmicrosoftGraphCallRecordsMedia.md) | Media associated with this segment. | [optional] 
**StartDateTime** | Pointer to **time.Time** | UTC time when the segment started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 

## Methods

### NewSegment

`func NewSegment() *Segment`

NewSegment instantiates a new Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentWithDefaults

`func NewSegmentWithDefaults() *Segment`

NewSegmentWithDefaults instantiates a new Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallee

`func (o *Segment) GetCallee() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCallee returns the Callee field if non-nil, zero value otherwise.

### GetCalleeOk

`func (o *Segment) GetCalleeOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCalleeOk returns a tuple with the Callee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallee

`func (o *Segment) SetCallee(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCallee sets Callee field to given value.

### HasCallee

`func (o *Segment) HasCallee() bool`

HasCallee returns a boolean if a field has been set.

### SetCalleeNil

`func (o *Segment) SetCalleeNil(b bool)`

 SetCalleeNil sets the value for Callee to be an explicit nil

### UnsetCallee
`func (o *Segment) UnsetCallee()`

UnsetCallee ensures that no value is present for Callee, not even an explicit nil
### GetCaller

`func (o *Segment) GetCaller() AnyOfmicrosoftGraphCallRecordsEndpoint`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Segment) GetCallerOk() (*AnyOfmicrosoftGraphCallRecordsEndpoint, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Segment) SetCaller(v AnyOfmicrosoftGraphCallRecordsEndpoint)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Segment) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Segment) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Segment) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetEndDateTime

`func (o *Segment) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Segment) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Segment) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Segment) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetFailureInfo

`func (o *Segment) GetFailureInfo() AnyOfmicrosoftGraphCallRecordsFailureInfo`

GetFailureInfo returns the FailureInfo field if non-nil, zero value otherwise.

### GetFailureInfoOk

`func (o *Segment) GetFailureInfoOk() (*AnyOfmicrosoftGraphCallRecordsFailureInfo, bool)`

GetFailureInfoOk returns a tuple with the FailureInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureInfo

`func (o *Segment) SetFailureInfo(v AnyOfmicrosoftGraphCallRecordsFailureInfo)`

SetFailureInfo sets FailureInfo field to given value.

### HasFailureInfo

`func (o *Segment) HasFailureInfo() bool`

HasFailureInfo returns a boolean if a field has been set.

### SetFailureInfoNil

`func (o *Segment) SetFailureInfoNil(b bool)`

 SetFailureInfoNil sets the value for FailureInfo to be an explicit nil

### UnsetFailureInfo
`func (o *Segment) UnsetFailureInfo()`

UnsetFailureInfo ensures that no value is present for FailureInfo, not even an explicit nil
### GetMedia

`func (o *Segment) GetMedia() []*AnyOfmicrosoftGraphCallRecordsMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Segment) GetMediaOk() (*[]*AnyOfmicrosoftGraphCallRecordsMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Segment) SetMedia(v []*AnyOfmicrosoftGraphCallRecordsMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Segment) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetStartDateTime

`func (o *Segment) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Segment) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Segment) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *Segment) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


