# MicrosoftGraphPatternedRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to [**AnyOfmicrosoftGraphRecurrencePattern**](anyOf&lt;microsoft.graph.recurrencePattern&gt;.md) | The frequency of an event. Do not specify for a one-time access review. | [optional] 
**Range** | Pointer to [**AnyOfmicrosoftGraphRecurrenceRange**](anyOf&lt;microsoft.graph.recurrenceRange&gt;.md) | The duration of an event. | [optional] 

## Methods

### NewMicrosoftGraphPatternedRecurrence

`func NewMicrosoftGraphPatternedRecurrence() *MicrosoftGraphPatternedRecurrence`

NewMicrosoftGraphPatternedRecurrence instantiates a new MicrosoftGraphPatternedRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPatternedRecurrenceWithDefaults

`func NewMicrosoftGraphPatternedRecurrenceWithDefaults() *MicrosoftGraphPatternedRecurrence`

NewMicrosoftGraphPatternedRecurrenceWithDefaults instantiates a new MicrosoftGraphPatternedRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *MicrosoftGraphPatternedRecurrence) GetPattern() AnyOfmicrosoftGraphRecurrencePattern`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *MicrosoftGraphPatternedRecurrence) GetPatternOk() (*AnyOfmicrosoftGraphRecurrencePattern, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *MicrosoftGraphPatternedRecurrence) SetPattern(v AnyOfmicrosoftGraphRecurrencePattern)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *MicrosoftGraphPatternedRecurrence) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *MicrosoftGraphPatternedRecurrence) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *MicrosoftGraphPatternedRecurrence) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetRange

`func (o *MicrosoftGraphPatternedRecurrence) GetRange() AnyOfmicrosoftGraphRecurrenceRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MicrosoftGraphPatternedRecurrence) GetRangeOk() (*AnyOfmicrosoftGraphRecurrenceRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *MicrosoftGraphPatternedRecurrence) SetRange(v AnyOfmicrosoftGraphRecurrenceRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *MicrosoftGraphPatternedRecurrence) HasRange() bool`

HasRange returns a boolean if a field has been set.

### SetRangeNil

`func (o *MicrosoftGraphPatternedRecurrence) SetRangeNil(b bool)`

 SetRangeNil sets the value for Range to be an explicit nil

### UnsetRange
`func (o *MicrosoftGraphPatternedRecurrence) UnsetRange()`

UnsetRange ensures that no value is present for Range, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


