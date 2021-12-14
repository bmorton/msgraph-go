# MicrosoftGraphAggregationOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketDefinition** | Pointer to [**MicrosoftGraphBucketAggregationDefinition**](MicrosoftGraphBucketAggregationDefinition.md) |  | [optional] 
**Field** | Pointer to **string** | Computes aggregation on the field while the field exists in current entity type. Required. | [optional] 
**Size** | Pointer to **NullableInt32** | The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional. | [optional] 

## Methods

### NewMicrosoftGraphAggregationOption

`func NewMicrosoftGraphAggregationOption() *MicrosoftGraphAggregationOption`

NewMicrosoftGraphAggregationOption instantiates a new MicrosoftGraphAggregationOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAggregationOptionWithDefaults

`func NewMicrosoftGraphAggregationOptionWithDefaults() *MicrosoftGraphAggregationOption`

NewMicrosoftGraphAggregationOptionWithDefaults instantiates a new MicrosoftGraphAggregationOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketDefinition

`func (o *MicrosoftGraphAggregationOption) GetBucketDefinition() MicrosoftGraphBucketAggregationDefinition`

GetBucketDefinition returns the BucketDefinition field if non-nil, zero value otherwise.

### GetBucketDefinitionOk

`func (o *MicrosoftGraphAggregationOption) GetBucketDefinitionOk() (*MicrosoftGraphBucketAggregationDefinition, bool)`

GetBucketDefinitionOk returns a tuple with the BucketDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketDefinition

`func (o *MicrosoftGraphAggregationOption) SetBucketDefinition(v MicrosoftGraphBucketAggregationDefinition)`

SetBucketDefinition sets BucketDefinition field to given value.

### HasBucketDefinition

`func (o *MicrosoftGraphAggregationOption) HasBucketDefinition() bool`

HasBucketDefinition returns a boolean if a field has been set.

### GetField

`func (o *MicrosoftGraphAggregationOption) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MicrosoftGraphAggregationOption) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MicrosoftGraphAggregationOption) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MicrosoftGraphAggregationOption) HasField() bool`

HasField returns a boolean if a field has been set.

### GetSize

`func (o *MicrosoftGraphAggregationOption) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphAggregationOption) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphAggregationOption) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphAggregationOption) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphAggregationOption) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphAggregationOption) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


