# MicrosoftGraphCalendarGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**ClassId** | Pointer to **NullableString** | The class identifier. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | The group name. | [optional] 
**Calendars** | Pointer to [**[]MicrosoftGraphCalendar**](MicrosoftGraphCalendar.md) | The calendars in the calendar group. Navigation property. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphCalendarGroup

`func NewMicrosoftGraphCalendarGroup() *MicrosoftGraphCalendarGroup`

NewMicrosoftGraphCalendarGroup instantiates a new MicrosoftGraphCalendarGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCalendarGroupWithDefaults

`func NewMicrosoftGraphCalendarGroupWithDefaults() *MicrosoftGraphCalendarGroup`

NewMicrosoftGraphCalendarGroupWithDefaults instantiates a new MicrosoftGraphCalendarGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphCalendarGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCalendarGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCalendarGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCalendarGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChangeKey

`func (o *MicrosoftGraphCalendarGroup) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphCalendarGroup) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *MicrosoftGraphCalendarGroup) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *MicrosoftGraphCalendarGroup) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *MicrosoftGraphCalendarGroup) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *MicrosoftGraphCalendarGroup) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetClassId

`func (o *MicrosoftGraphCalendarGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MicrosoftGraphCalendarGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MicrosoftGraphCalendarGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *MicrosoftGraphCalendarGroup) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### SetClassIdNil

`func (o *MicrosoftGraphCalendarGroup) SetClassIdNil(b bool)`

 SetClassIdNil sets the value for ClassId to be an explicit nil

### UnsetClassId
`func (o *MicrosoftGraphCalendarGroup) UnsetClassId()`

UnsetClassId ensures that no value is present for ClassId, not even an explicit nil
### GetName

`func (o *MicrosoftGraphCalendarGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphCalendarGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphCalendarGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphCalendarGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphCalendarGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphCalendarGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCalendars

`func (o *MicrosoftGraphCalendarGroup) GetCalendars() []MicrosoftGraphCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *MicrosoftGraphCalendarGroup) GetCalendarsOk() (*[]MicrosoftGraphCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendars

`func (o *MicrosoftGraphCalendarGroup) SetCalendars(v []MicrosoftGraphCalendar)`

SetCalendars sets Calendars field to given value.

### HasCalendars

`func (o *MicrosoftGraphCalendarGroup) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


