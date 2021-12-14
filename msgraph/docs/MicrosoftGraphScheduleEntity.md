# MicrosoftGraphScheduleEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**Theme** | Pointer to [**AnyOfmicrosoftGraphScheduleEntityTheme**](anyOf&lt;microsoft.graph.scheduleEntityTheme&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphScheduleEntity

`func NewMicrosoftGraphScheduleEntity() *MicrosoftGraphScheduleEntity`

NewMicrosoftGraphScheduleEntity instantiates a new MicrosoftGraphScheduleEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScheduleEntityWithDefaults

`func NewMicrosoftGraphScheduleEntityWithDefaults() *MicrosoftGraphScheduleEntity`

NewMicrosoftGraphScheduleEntityWithDefaults instantiates a new MicrosoftGraphScheduleEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *MicrosoftGraphScheduleEntity) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphScheduleEntity) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphScheduleEntity) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphScheduleEntity) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphScheduleEntity) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphScheduleEntity) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphScheduleEntity) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphScheduleEntity) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphScheduleEntity) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphScheduleEntity) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphScheduleEntity) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphScheduleEntity) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTheme

`func (o *MicrosoftGraphScheduleEntity) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MicrosoftGraphScheduleEntity) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MicrosoftGraphScheduleEntity) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MicrosoftGraphScheduleEntity) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *MicrosoftGraphScheduleEntity) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *MicrosoftGraphScheduleEntity) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


