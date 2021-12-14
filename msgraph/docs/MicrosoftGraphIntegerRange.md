# MicrosoftGraphIntegerRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **NullableInt64** | The inclusive upper bound of the integer range. | [optional] 
**Start** | Pointer to **NullableInt64** | The inclusive lower bound of the integer range. | [optional] 

## Methods

### NewMicrosoftGraphIntegerRange

`func NewMicrosoftGraphIntegerRange() *MicrosoftGraphIntegerRange`

NewMicrosoftGraphIntegerRange instantiates a new MicrosoftGraphIntegerRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIntegerRangeWithDefaults

`func NewMicrosoftGraphIntegerRangeWithDefaults() *MicrosoftGraphIntegerRange`

NewMicrosoftGraphIntegerRangeWithDefaults instantiates a new MicrosoftGraphIntegerRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MicrosoftGraphIntegerRange) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MicrosoftGraphIntegerRange) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MicrosoftGraphIntegerRange) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MicrosoftGraphIntegerRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *MicrosoftGraphIntegerRange) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *MicrosoftGraphIntegerRange) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStart

`func (o *MicrosoftGraphIntegerRange) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MicrosoftGraphIntegerRange) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MicrosoftGraphIntegerRange) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *MicrosoftGraphIntegerRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *MicrosoftGraphIntegerRange) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *MicrosoftGraphIntegerRange) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


