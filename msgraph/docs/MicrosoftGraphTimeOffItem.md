# MicrosoftGraphTimeOffItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**Theme** | Pointer to [**AnyOfmicrosoftGraphScheduleEntityTheme**](anyOf&lt;microsoft.graph.scheduleEntityTheme&gt;.md) |  | [optional] 
**TimeOffReasonId** | Pointer to **NullableString** | ID of the timeOffReason for this timeOffItem. Required. | [optional] 

## Methods

### NewMicrosoftGraphTimeOffItem

`func NewMicrosoftGraphTimeOffItem() *MicrosoftGraphTimeOffItem`

NewMicrosoftGraphTimeOffItem instantiates a new MicrosoftGraphTimeOffItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTimeOffItemWithDefaults

`func NewMicrosoftGraphTimeOffItemWithDefaults() *MicrosoftGraphTimeOffItem`

NewMicrosoftGraphTimeOffItemWithDefaults instantiates a new MicrosoftGraphTimeOffItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *MicrosoftGraphTimeOffItem) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphTimeOffItem) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphTimeOffItem) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphTimeOffItem) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphTimeOffItem) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphTimeOffItem) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphTimeOffItem) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphTimeOffItem) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphTimeOffItem) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphTimeOffItem) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphTimeOffItem) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphTimeOffItem) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTheme

`func (o *MicrosoftGraphTimeOffItem) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MicrosoftGraphTimeOffItem) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MicrosoftGraphTimeOffItem) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MicrosoftGraphTimeOffItem) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *MicrosoftGraphTimeOffItem) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *MicrosoftGraphTimeOffItem) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil
### GetTimeOffReasonId

`func (o *MicrosoftGraphTimeOffItem) GetTimeOffReasonId() string`

GetTimeOffReasonId returns the TimeOffReasonId field if non-nil, zero value otherwise.

### GetTimeOffReasonIdOk

`func (o *MicrosoftGraphTimeOffItem) GetTimeOffReasonIdOk() (*string, bool)`

GetTimeOffReasonIdOk returns a tuple with the TimeOffReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffReasonId

`func (o *MicrosoftGraphTimeOffItem) SetTimeOffReasonId(v string)`

SetTimeOffReasonId sets TimeOffReasonId field to given value.

### HasTimeOffReasonId

`func (o *MicrosoftGraphTimeOffItem) HasTimeOffReasonId() bool`

HasTimeOffReasonId returns a boolean if a field has been set.

### SetTimeOffReasonIdNil

`func (o *MicrosoftGraphTimeOffItem) SetTimeOffReasonIdNil(b bool)`

 SetTimeOffReasonIdNil sets the value for TimeOffReasonId to be an explicit nil

### UnsetTimeOffReasonId
`func (o *MicrosoftGraphTimeOffItem) UnsetTimeOffReasonId()`

UnsetTimeOffReasonId ensures that no value is present for TimeOffReasonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


