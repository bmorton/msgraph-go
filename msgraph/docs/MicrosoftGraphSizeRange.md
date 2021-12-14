# MicrosoftGraphSizeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumSize** | Pointer to **NullableInt32** | The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. | [optional] 
**MinimumSize** | Pointer to **NullableInt32** | The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. | [optional] 

## Methods

### NewMicrosoftGraphSizeRange

`func NewMicrosoftGraphSizeRange() *MicrosoftGraphSizeRange`

NewMicrosoftGraphSizeRange instantiates a new MicrosoftGraphSizeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSizeRangeWithDefaults

`func NewMicrosoftGraphSizeRangeWithDefaults() *MicrosoftGraphSizeRange`

NewMicrosoftGraphSizeRangeWithDefaults instantiates a new MicrosoftGraphSizeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumSize

`func (o *MicrosoftGraphSizeRange) GetMaximumSize() int32`

GetMaximumSize returns the MaximumSize field if non-nil, zero value otherwise.

### GetMaximumSizeOk

`func (o *MicrosoftGraphSizeRange) GetMaximumSizeOk() (*int32, bool)`

GetMaximumSizeOk returns a tuple with the MaximumSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSize

`func (o *MicrosoftGraphSizeRange) SetMaximumSize(v int32)`

SetMaximumSize sets MaximumSize field to given value.

### HasMaximumSize

`func (o *MicrosoftGraphSizeRange) HasMaximumSize() bool`

HasMaximumSize returns a boolean if a field has been set.

### SetMaximumSizeNil

`func (o *MicrosoftGraphSizeRange) SetMaximumSizeNil(b bool)`

 SetMaximumSizeNil sets the value for MaximumSize to be an explicit nil

### UnsetMaximumSize
`func (o *MicrosoftGraphSizeRange) UnsetMaximumSize()`

UnsetMaximumSize ensures that no value is present for MaximumSize, not even an explicit nil
### GetMinimumSize

`func (o *MicrosoftGraphSizeRange) GetMinimumSize() int32`

GetMinimumSize returns the MinimumSize field if non-nil, zero value otherwise.

### GetMinimumSizeOk

`func (o *MicrosoftGraphSizeRange) GetMinimumSizeOk() (*int32, bool)`

GetMinimumSizeOk returns a tuple with the MinimumSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSize

`func (o *MicrosoftGraphSizeRange) SetMinimumSize(v int32)`

SetMinimumSize sets MinimumSize field to given value.

### HasMinimumSize

`func (o *MicrosoftGraphSizeRange) HasMinimumSize() bool`

HasMinimumSize returns a boolean if a field has been set.

### SetMinimumSizeNil

`func (o *MicrosoftGraphSizeRange) SetMinimumSizeNil(b bool)`

 SetMinimumSizeNil sets the value for MinimumSize to be an explicit nil

### UnsetMinimumSize
`func (o *MicrosoftGraphSizeRange) UnsetMinimumSize()`

UnsetMinimumSize ensures that no value is present for MinimumSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


