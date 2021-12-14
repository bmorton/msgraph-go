# MicrosoftGraphWorkingHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfWeek** | Pointer to [**[]AnyOfmicrosoftGraphDayOfWeek**](AnyOfmicrosoftGraphDayOfWeek.md) | The days of the week on which the user works. | [optional] 
**EndTime** | Pointer to **NullableString** | The time of the day that the user stops working. | [optional] 
**StartTime** | Pointer to **NullableString** | The time of the day that the user starts working. | [optional] 
**TimeZone** | Pointer to [**AnyOfmicrosoftGraphTimeZoneBase**](anyOf&lt;microsoft.graph.timeZoneBase&gt;.md) | The time zone to which the working hours apply. | [optional] 

## Methods

### NewMicrosoftGraphWorkingHours

`func NewMicrosoftGraphWorkingHours() *MicrosoftGraphWorkingHours`

NewMicrosoftGraphWorkingHours instantiates a new MicrosoftGraphWorkingHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkingHoursWithDefaults

`func NewMicrosoftGraphWorkingHoursWithDefaults() *MicrosoftGraphWorkingHours`

NewMicrosoftGraphWorkingHoursWithDefaults instantiates a new MicrosoftGraphWorkingHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfWeek

`func (o *MicrosoftGraphWorkingHours) GetDaysOfWeek() []*AnyOfmicrosoftGraphDayOfWeek`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *MicrosoftGraphWorkingHours) GetDaysOfWeekOk() (*[]*AnyOfmicrosoftGraphDayOfWeek, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *MicrosoftGraphWorkingHours) SetDaysOfWeek(v []*AnyOfmicrosoftGraphDayOfWeek)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *MicrosoftGraphWorkingHours) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *MicrosoftGraphWorkingHours) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MicrosoftGraphWorkingHours) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MicrosoftGraphWorkingHours) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MicrosoftGraphWorkingHours) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *MicrosoftGraphWorkingHours) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *MicrosoftGraphWorkingHours) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetStartTime

`func (o *MicrosoftGraphWorkingHours) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MicrosoftGraphWorkingHours) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MicrosoftGraphWorkingHours) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MicrosoftGraphWorkingHours) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *MicrosoftGraphWorkingHours) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *MicrosoftGraphWorkingHours) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTimeZone

`func (o *MicrosoftGraphWorkingHours) GetTimeZone() AnyOfmicrosoftGraphTimeZoneBase`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphWorkingHours) GetTimeZoneOk() (*AnyOfmicrosoftGraphTimeZoneBase, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MicrosoftGraphWorkingHours) SetTimeZone(v AnyOfmicrosoftGraphTimeZoneBase)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MicrosoftGraphWorkingHours) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *MicrosoftGraphWorkingHours) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *MicrosoftGraphWorkingHours) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


