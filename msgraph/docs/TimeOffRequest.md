# TimeOffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**StartDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**TimeOffReasonId** | Pointer to **NullableString** | The reason for the time off. | [optional] 

## Methods

### NewTimeOffRequest

`func NewTimeOffRequest() *TimeOffRequest`

NewTimeOffRequest instantiates a new TimeOffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffRequestWithDefaults

`func NewTimeOffRequestWithDefaults() *TimeOffRequest`

NewTimeOffRequestWithDefaults instantiates a new TimeOffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *TimeOffRequest) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *TimeOffRequest) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *TimeOffRequest) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *TimeOffRequest) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *TimeOffRequest) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *TimeOffRequest) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetStartDateTime

`func (o *TimeOffRequest) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *TimeOffRequest) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *TimeOffRequest) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *TimeOffRequest) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *TimeOffRequest) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *TimeOffRequest) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTimeOffReasonId

`func (o *TimeOffRequest) GetTimeOffReasonId() string`

GetTimeOffReasonId returns the TimeOffReasonId field if non-nil, zero value otherwise.

### GetTimeOffReasonIdOk

`func (o *TimeOffRequest) GetTimeOffReasonIdOk() (*string, bool)`

GetTimeOffReasonIdOk returns a tuple with the TimeOffReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffReasonId

`func (o *TimeOffRequest) SetTimeOffReasonId(v string)`

SetTimeOffReasonId sets TimeOffReasonId field to given value.

### HasTimeOffReasonId

`func (o *TimeOffRequest) HasTimeOffReasonId() bool`

HasTimeOffReasonId returns a boolean if a field has been set.

### SetTimeOffReasonIdNil

`func (o *TimeOffRequest) SetTimeOffReasonIdNil(b bool)`

 SetTimeOffReasonIdNil sets the value for TimeOffReasonId to be an explicit nil

### UnsetTimeOffReasonId
`func (o *TimeOffRequest) UnsetTimeOffReasonId()`

UnsetTimeOffReasonId ensures that no value is present for TimeOffReasonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


