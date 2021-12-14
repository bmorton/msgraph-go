# MicrosoftGraphShiftAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recurrence** | Pointer to [**AnyOfmicrosoftGraphPatternedRecurrence**](anyOf&lt;microsoft.graph.patternedRecurrence&gt;.md) | Specifies the pattern for recurrence | [optional] 
**TimeSlots** | Pointer to [**[]AnyOfmicrosoftGraphTimeRange**](AnyOfmicrosoftGraphTimeRange.md) | The time slot(s) preferred by the user. | [optional] 
**TimeZone** | Pointer to **NullableString** | Specifies the time zone for the indicated time. | [optional] 

## Methods

### NewMicrosoftGraphShiftAvailability

`func NewMicrosoftGraphShiftAvailability() *MicrosoftGraphShiftAvailability`

NewMicrosoftGraphShiftAvailability instantiates a new MicrosoftGraphShiftAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphShiftAvailabilityWithDefaults

`func NewMicrosoftGraphShiftAvailabilityWithDefaults() *MicrosoftGraphShiftAvailability`

NewMicrosoftGraphShiftAvailabilityWithDefaults instantiates a new MicrosoftGraphShiftAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurrence

`func (o *MicrosoftGraphShiftAvailability) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *MicrosoftGraphShiftAvailability) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *MicrosoftGraphShiftAvailability) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *MicrosoftGraphShiftAvailability) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *MicrosoftGraphShiftAvailability) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *MicrosoftGraphShiftAvailability) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetTimeSlots

`func (o *MicrosoftGraphShiftAvailability) GetTimeSlots() []*AnyOfmicrosoftGraphTimeRange`

GetTimeSlots returns the TimeSlots field if non-nil, zero value otherwise.

### GetTimeSlotsOk

`func (o *MicrosoftGraphShiftAvailability) GetTimeSlotsOk() (*[]*AnyOfmicrosoftGraphTimeRange, bool)`

GetTimeSlotsOk returns a tuple with the TimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlots

`func (o *MicrosoftGraphShiftAvailability) SetTimeSlots(v []*AnyOfmicrosoftGraphTimeRange)`

SetTimeSlots sets TimeSlots field to given value.

### HasTimeSlots

`func (o *MicrosoftGraphShiftAvailability) HasTimeSlots() bool`

HasTimeSlots returns a boolean if a field has been set.

### GetTimeZone

`func (o *MicrosoftGraphShiftAvailability) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphShiftAvailability) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MicrosoftGraphShiftAvailability) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MicrosoftGraphShiftAvailability) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *MicrosoftGraphShiftAvailability) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *MicrosoftGraphShiftAvailability) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


