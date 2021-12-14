# MicrosoftGraphRecurrenceRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **NullableString** | The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate. | [optional] 
**NumberOfOccurrences** | Pointer to **int32** | The number of times to repeat the event. Required and must be positive if type is numbered. | [optional] 
**RecurrenceTimeZone** | Pointer to **NullableString** | Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used. | [optional] 
**StartDate** | Pointer to **NullableString** | The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphRecurrenceRangeType**](anyOf&lt;microsoft.graph.recurrenceRangeType&gt;.md) | The recurrence range. The possible values are: endDate, noEnd, numbered. Required. | [optional] 

## Methods

### NewMicrosoftGraphRecurrenceRange

`func NewMicrosoftGraphRecurrenceRange() *MicrosoftGraphRecurrenceRange`

NewMicrosoftGraphRecurrenceRange instantiates a new MicrosoftGraphRecurrenceRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRecurrenceRangeWithDefaults

`func NewMicrosoftGraphRecurrenceRangeWithDefaults() *MicrosoftGraphRecurrenceRange`

NewMicrosoftGraphRecurrenceRangeWithDefaults instantiates a new MicrosoftGraphRecurrenceRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *MicrosoftGraphRecurrenceRange) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MicrosoftGraphRecurrenceRange) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MicrosoftGraphRecurrenceRange) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MicrosoftGraphRecurrenceRange) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *MicrosoftGraphRecurrenceRange) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *MicrosoftGraphRecurrenceRange) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetNumberOfOccurrences

`func (o *MicrosoftGraphRecurrenceRange) GetNumberOfOccurrences() int32`

GetNumberOfOccurrences returns the NumberOfOccurrences field if non-nil, zero value otherwise.

### GetNumberOfOccurrencesOk

`func (o *MicrosoftGraphRecurrenceRange) GetNumberOfOccurrencesOk() (*int32, bool)`

GetNumberOfOccurrencesOk returns a tuple with the NumberOfOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOccurrences

`func (o *MicrosoftGraphRecurrenceRange) SetNumberOfOccurrences(v int32)`

SetNumberOfOccurrences sets NumberOfOccurrences field to given value.

### HasNumberOfOccurrences

`func (o *MicrosoftGraphRecurrenceRange) HasNumberOfOccurrences() bool`

HasNumberOfOccurrences returns a boolean if a field has been set.

### GetRecurrenceTimeZone

`func (o *MicrosoftGraphRecurrenceRange) GetRecurrenceTimeZone() string`

GetRecurrenceTimeZone returns the RecurrenceTimeZone field if non-nil, zero value otherwise.

### GetRecurrenceTimeZoneOk

`func (o *MicrosoftGraphRecurrenceRange) GetRecurrenceTimeZoneOk() (*string, bool)`

GetRecurrenceTimeZoneOk returns a tuple with the RecurrenceTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceTimeZone

`func (o *MicrosoftGraphRecurrenceRange) SetRecurrenceTimeZone(v string)`

SetRecurrenceTimeZone sets RecurrenceTimeZone field to given value.

### HasRecurrenceTimeZone

`func (o *MicrosoftGraphRecurrenceRange) HasRecurrenceTimeZone() bool`

HasRecurrenceTimeZone returns a boolean if a field has been set.

### SetRecurrenceTimeZoneNil

`func (o *MicrosoftGraphRecurrenceRange) SetRecurrenceTimeZoneNil(b bool)`

 SetRecurrenceTimeZoneNil sets the value for RecurrenceTimeZone to be an explicit nil

### UnsetRecurrenceTimeZone
`func (o *MicrosoftGraphRecurrenceRange) UnsetRecurrenceTimeZone()`

UnsetRecurrenceTimeZone ensures that no value is present for RecurrenceTimeZone, not even an explicit nil
### GetStartDate

`func (o *MicrosoftGraphRecurrenceRange) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MicrosoftGraphRecurrenceRange) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MicrosoftGraphRecurrenceRange) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MicrosoftGraphRecurrenceRange) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *MicrosoftGraphRecurrenceRange) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *MicrosoftGraphRecurrenceRange) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetType

`func (o *MicrosoftGraphRecurrenceRange) GetType() AnyOfmicrosoftGraphRecurrenceRangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphRecurrenceRange) GetTypeOk() (*AnyOfmicrosoftGraphRecurrenceRangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphRecurrenceRange) SetType(v AnyOfmicrosoftGraphRecurrenceRangeType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphRecurrenceRange) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphRecurrenceRange) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphRecurrenceRange) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


