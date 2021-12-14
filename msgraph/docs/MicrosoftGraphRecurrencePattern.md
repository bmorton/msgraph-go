# MicrosoftGraphRecurrencePattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **int32** | The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly. | [optional] 
**DaysOfWeek** | Pointer to [**[]AnyOfmicrosoftGraphDayOfWeek**](AnyOfmicrosoftGraphDayOfWeek.md) | A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly. | [optional] 
**FirstDayOfWeek** | Pointer to [**AnyOfmicrosoftGraphDayOfWeek**](anyOf&lt;microsoft.graph.dayOfWeek&gt;.md) | The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly. | [optional] 
**Index** | Pointer to [**AnyOfmicrosoftGraphWeekIndex**](anyOf&lt;microsoft.graph.weekIndex&gt;.md) | Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly. | [optional] 
**Interval** | Pointer to **int32** | The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required. | [optional] 
**Month** | Pointer to **int32** | The month in which the event occurs.  This is a number from 1 to 12. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphRecurrencePatternType**](anyOf&lt;microsoft.graph.recurrencePatternType&gt;.md) | The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. | [optional] 

## Methods

### NewMicrosoftGraphRecurrencePattern

`func NewMicrosoftGraphRecurrencePattern() *MicrosoftGraphRecurrencePattern`

NewMicrosoftGraphRecurrencePattern instantiates a new MicrosoftGraphRecurrencePattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRecurrencePatternWithDefaults

`func NewMicrosoftGraphRecurrencePatternWithDefaults() *MicrosoftGraphRecurrencePattern`

NewMicrosoftGraphRecurrencePatternWithDefaults instantiates a new MicrosoftGraphRecurrencePattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *MicrosoftGraphRecurrencePattern) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MicrosoftGraphRecurrencePattern) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *MicrosoftGraphRecurrencePattern) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *MicrosoftGraphRecurrencePattern) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetDaysOfWeek

`func (o *MicrosoftGraphRecurrencePattern) GetDaysOfWeek() []*AnyOfmicrosoftGraphDayOfWeek`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *MicrosoftGraphRecurrencePattern) GetDaysOfWeekOk() (*[]*AnyOfmicrosoftGraphDayOfWeek, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *MicrosoftGraphRecurrencePattern) SetDaysOfWeek(v []*AnyOfmicrosoftGraphDayOfWeek)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *MicrosoftGraphRecurrencePattern) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### GetFirstDayOfWeek

`func (o *MicrosoftGraphRecurrencePattern) GetFirstDayOfWeek() AnyOfmicrosoftGraphDayOfWeek`

GetFirstDayOfWeek returns the FirstDayOfWeek field if non-nil, zero value otherwise.

### GetFirstDayOfWeekOk

`func (o *MicrosoftGraphRecurrencePattern) GetFirstDayOfWeekOk() (*AnyOfmicrosoftGraphDayOfWeek, bool)`

GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDayOfWeek

`func (o *MicrosoftGraphRecurrencePattern) SetFirstDayOfWeek(v AnyOfmicrosoftGraphDayOfWeek)`

SetFirstDayOfWeek sets FirstDayOfWeek field to given value.

### HasFirstDayOfWeek

`func (o *MicrosoftGraphRecurrencePattern) HasFirstDayOfWeek() bool`

HasFirstDayOfWeek returns a boolean if a field has been set.

### SetFirstDayOfWeekNil

`func (o *MicrosoftGraphRecurrencePattern) SetFirstDayOfWeekNil(b bool)`

 SetFirstDayOfWeekNil sets the value for FirstDayOfWeek to be an explicit nil

### UnsetFirstDayOfWeek
`func (o *MicrosoftGraphRecurrencePattern) UnsetFirstDayOfWeek()`

UnsetFirstDayOfWeek ensures that no value is present for FirstDayOfWeek, not even an explicit nil
### GetIndex

`func (o *MicrosoftGraphRecurrencePattern) GetIndex() AnyOfmicrosoftGraphWeekIndex`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphRecurrencePattern) GetIndexOk() (*AnyOfmicrosoftGraphWeekIndex, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MicrosoftGraphRecurrencePattern) SetIndex(v AnyOfmicrosoftGraphWeekIndex)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MicrosoftGraphRecurrencePattern) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *MicrosoftGraphRecurrencePattern) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *MicrosoftGraphRecurrencePattern) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetInterval

`func (o *MicrosoftGraphRecurrencePattern) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MicrosoftGraphRecurrencePattern) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MicrosoftGraphRecurrencePattern) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MicrosoftGraphRecurrencePattern) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMonth

`func (o *MicrosoftGraphRecurrencePattern) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MicrosoftGraphRecurrencePattern) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *MicrosoftGraphRecurrencePattern) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *MicrosoftGraphRecurrencePattern) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftGraphRecurrencePattern) GetType() AnyOfmicrosoftGraphRecurrencePatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphRecurrencePattern) GetTypeOk() (*AnyOfmicrosoftGraphRecurrencePatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphRecurrencePattern) SetType(v AnyOfmicrosoftGraphRecurrencePatternType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphRecurrencePattern) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphRecurrencePattern) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphRecurrencePattern) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


