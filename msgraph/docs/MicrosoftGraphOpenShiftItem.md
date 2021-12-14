# MicrosoftGraphOpenShiftItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **NullableTime** |  | [optional] 
**StartDateTime** | Pointer to **NullableTime** |  | [optional] 
**Theme** | Pointer to [**AnyOfmicrosoftGraphScheduleEntityTheme**](anyOf&lt;microsoft.graph.scheduleEntityTheme&gt;.md) |  | [optional] 
**Activities** | Pointer to [**[]AnyOfmicrosoftGraphShiftActivity**](AnyOfmicrosoftGraphShiftActivity.md) | An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | The shift label of the shiftItem. | [optional] 
**Notes** | Pointer to **NullableString** | The shift notes for the shiftItem. | [optional] 
**OpenSlotCount** | Pointer to **int32** | Count of the number of slots for the given open shift. | [optional] 

## Methods

### NewMicrosoftGraphOpenShiftItem

`func NewMicrosoftGraphOpenShiftItem() *MicrosoftGraphOpenShiftItem`

NewMicrosoftGraphOpenShiftItem instantiates a new MicrosoftGraphOpenShiftItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOpenShiftItemWithDefaults

`func NewMicrosoftGraphOpenShiftItemWithDefaults() *MicrosoftGraphOpenShiftItem`

NewMicrosoftGraphOpenShiftItemWithDefaults instantiates a new MicrosoftGraphOpenShiftItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *MicrosoftGraphOpenShiftItem) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphOpenShiftItem) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphOpenShiftItem) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphOpenShiftItem) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphOpenShiftItem) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphOpenShiftItem) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphOpenShiftItem) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphOpenShiftItem) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphOpenShiftItem) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphOpenShiftItem) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphOpenShiftItem) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphOpenShiftItem) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTheme

`func (o *MicrosoftGraphOpenShiftItem) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MicrosoftGraphOpenShiftItem) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MicrosoftGraphOpenShiftItem) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MicrosoftGraphOpenShiftItem) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *MicrosoftGraphOpenShiftItem) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *MicrosoftGraphOpenShiftItem) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil
### GetActivities

`func (o *MicrosoftGraphOpenShiftItem) GetActivities() []*AnyOfmicrosoftGraphShiftActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *MicrosoftGraphOpenShiftItem) GetActivitiesOk() (*[]*AnyOfmicrosoftGraphShiftActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *MicrosoftGraphOpenShiftItem) SetActivities(v []*AnyOfmicrosoftGraphShiftActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *MicrosoftGraphOpenShiftItem) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphOpenShiftItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOpenShiftItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphOpenShiftItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphOpenShiftItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphOpenShiftItem) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphOpenShiftItem) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetNotes

`func (o *MicrosoftGraphOpenShiftItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphOpenShiftItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MicrosoftGraphOpenShiftItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MicrosoftGraphOpenShiftItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *MicrosoftGraphOpenShiftItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *MicrosoftGraphOpenShiftItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOpenSlotCount

`func (o *MicrosoftGraphOpenShiftItem) GetOpenSlotCount() int32`

GetOpenSlotCount returns the OpenSlotCount field if non-nil, zero value otherwise.

### GetOpenSlotCountOk

`func (o *MicrosoftGraphOpenShiftItem) GetOpenSlotCountOk() (*int32, bool)`

GetOpenSlotCountOk returns a tuple with the OpenSlotCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSlotCount

`func (o *MicrosoftGraphOpenShiftItem) SetOpenSlotCount(v int32)`

SetOpenSlotCount sets OpenSlotCount field to given value.

### HasOpenSlotCount

`func (o *MicrosoftGraphOpenShiftItem) HasOpenSlotCount() bool`

HasOpenSlotCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


