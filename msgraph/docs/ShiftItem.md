# ShiftItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to [**[]AnyOfmicrosoftGraphShiftActivity**](AnyOfmicrosoftGraphShiftActivity.md) | An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | The shift label of the shiftItem. | [optional] 
**Notes** | Pointer to **NullableString** | The shift notes for the shiftItem. | [optional] 

## Methods

### NewShiftItem

`func NewShiftItem() *ShiftItem`

NewShiftItem instantiates a new ShiftItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShiftItemWithDefaults

`func NewShiftItemWithDefaults() *ShiftItem`

NewShiftItemWithDefaults instantiates a new ShiftItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *ShiftItem) GetActivities() []*AnyOfmicrosoftGraphShiftActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ShiftItem) GetActivitiesOk() (*[]*AnyOfmicrosoftGraphShiftActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ShiftItem) SetActivities(v []*AnyOfmicrosoftGraphShiftActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *ShiftItem) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetDisplayName

`func (o *ShiftItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ShiftItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ShiftItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ShiftItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ShiftItem) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ShiftItem) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetNotes

`func (o *ShiftItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ShiftItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ShiftItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ShiftItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ShiftItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ShiftItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


