# InlineObject926

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to **[]string** |  | [optional] 
**EndTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**StartTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**AvailabilityViewInterval** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewInlineObject926

`func NewInlineObject926() *InlineObject926`

NewInlineObject926 instantiates a new InlineObject926 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject926WithDefaults

`func NewInlineObject926WithDefaults() *InlineObject926`

NewInlineObject926WithDefaults instantiates a new InlineObject926 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *InlineObject926) GetSchedules() []*string`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *InlineObject926) GetSchedulesOk() (*[]*string, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *InlineObject926) SetSchedules(v []*string)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *InlineObject926) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetEndTime

`func (o *InlineObject926) GetEndTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InlineObject926) GetEndTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InlineObject926) SetEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InlineObject926) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *InlineObject926) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *InlineObject926) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *InlineObject926) GetStartTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InlineObject926) GetStartTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InlineObject926) SetStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InlineObject926) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *InlineObject926) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *InlineObject926) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetAvailabilityViewInterval

`func (o *InlineObject926) GetAvailabilityViewInterval() int32`

GetAvailabilityViewInterval returns the AvailabilityViewInterval field if non-nil, zero value otherwise.

### GetAvailabilityViewIntervalOk

`func (o *InlineObject926) GetAvailabilityViewIntervalOk() (*int32, bool)`

GetAvailabilityViewIntervalOk returns a tuple with the AvailabilityViewInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityViewInterval

`func (o *InlineObject926) SetAvailabilityViewInterval(v int32)`

SetAvailabilityViewInterval sets AvailabilityViewInterval field to given value.

### HasAvailabilityViewInterval

`func (o *InlineObject926) HasAvailabilityViewInterval() bool`

HasAvailabilityViewInterval returns a boolean if a field has been set.

### SetAvailabilityViewIntervalNil

`func (o *InlineObject926) SetAvailabilityViewIntervalNil(b bool)`

 SetAvailabilityViewIntervalNil sets the value for AvailabilityViewInterval to be an explicit nil

### UnsetAvailabilityViewInterval
`func (o *InlineObject926) UnsetAvailabilityViewInterval()`

UnsetAvailabilityViewInterval ensures that no value is present for AvailabilityViewInterval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


