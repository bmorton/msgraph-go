# CalendarGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**ClassId** | Pointer to **NullableString** | The class identifier. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | The group name. | [optional] 
**Calendars** | Pointer to [**[]MicrosoftGraphCalendar**](MicrosoftGraphCalendar.md) | The calendars in the calendar group. Navigation property. Read-only. Nullable. | [optional] 

## Methods

### NewCalendarGroup

`func NewCalendarGroup() *CalendarGroup`

NewCalendarGroup instantiates a new CalendarGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarGroupWithDefaults

`func NewCalendarGroupWithDefaults() *CalendarGroup`

NewCalendarGroupWithDefaults instantiates a new CalendarGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeKey

`func (o *CalendarGroup) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *CalendarGroup) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *CalendarGroup) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *CalendarGroup) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *CalendarGroup) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *CalendarGroup) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetClassId

`func (o *CalendarGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CalendarGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CalendarGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *CalendarGroup) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### SetClassIdNil

`func (o *CalendarGroup) SetClassIdNil(b bool)`

 SetClassIdNil sets the value for ClassId to be an explicit nil

### UnsetClassId
`func (o *CalendarGroup) UnsetClassId()`

UnsetClassId ensures that no value is present for ClassId, not even an explicit nil
### GetName

`func (o *CalendarGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalendarGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalendarGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CalendarGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CalendarGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CalendarGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCalendars

`func (o *CalendarGroup) GetCalendars() []MicrosoftGraphCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *CalendarGroup) GetCalendarsOk() (*[]MicrosoftGraphCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendars

`func (o *CalendarGroup) SetCalendars(v []MicrosoftGraphCalendar)`

SetCalendars sets Calendars field to given value.

### HasCalendars

`func (o *CalendarGroup) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


