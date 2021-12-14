# MicrosoftGraphShiftItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**Theme** | Pointer to [**AnyOfmicrosoftGraphScheduleEntityTheme**](anyOf&lt;microsoft.graph.scheduleEntityTheme&gt;.md) |  | [optional] 
**Activities** | Pointer to [**[]AnyOfmicrosoftGraphShiftActivity**](AnyOfmicrosoftGraphShiftActivity.md) | An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | The shift label of the shiftItem. | [optional] 
**Notes** | Pointer to **NullableString** | The shift notes for the shiftItem. | [optional] 

## Methods

### NewMicrosoftGraphShiftItem

`func NewMicrosoftGraphShiftItem() *MicrosoftGraphShiftItem`

NewMicrosoftGraphShiftItem instantiates a new MicrosoftGraphShiftItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphShiftItemWithDefaults

`func NewMicrosoftGraphShiftItemWithDefaults() *MicrosoftGraphShiftItem`

NewMicrosoftGraphShiftItemWithDefaults instantiates a new MicrosoftGraphShiftItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *MicrosoftGraphShiftItem) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphShiftItem) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphShiftItem) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphShiftItem) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphShiftItem) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphShiftItem) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphShiftItem) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphShiftItem) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphShiftItem) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphShiftItem) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphShiftItem) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphShiftItem) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTheme

`func (o *MicrosoftGraphShiftItem) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MicrosoftGraphShiftItem) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MicrosoftGraphShiftItem) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MicrosoftGraphShiftItem) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *MicrosoftGraphShiftItem) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *MicrosoftGraphShiftItem) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil
### GetActivities

`func (o *MicrosoftGraphShiftItem) GetActivities() []*AnyOfmicrosoftGraphShiftActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *MicrosoftGraphShiftItem) GetActivitiesOk() (*[]*AnyOfmicrosoftGraphShiftActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *MicrosoftGraphShiftItem) SetActivities(v []*AnyOfmicrosoftGraphShiftActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *MicrosoftGraphShiftItem) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphShiftItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphShiftItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphShiftItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphShiftItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphShiftItem) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphShiftItem) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetNotes

`func (o *MicrosoftGraphShiftItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphShiftItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MicrosoftGraphShiftItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MicrosoftGraphShiftItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *MicrosoftGraphShiftItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *MicrosoftGraphShiftItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


