# MicrosoftGraphBucketAggregationDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDescending** | Pointer to **NullableBool** | True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional. | [optional] 
**MinimumCount** | Pointer to **NullableInt32** | The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional. | [optional] 
**PrefixFilter** | Pointer to **NullableString** | A filter to define a matching criteria. The key should start with the specified prefix to be returned in the response. Optional. | [optional] 
**Ranges** | Pointer to [**[]AnyOfmicrosoftGraphBucketAggregationRange**](AnyOfmicrosoftGraphBucketAggregationRange.md) | Specifies the manual ranges to compute the aggregations. This is only valid for non-string refiners of date or numeric type. Optional. | [optional] 
**SortBy** | Pointer to [**AnyOfmicrosoftGraphBucketAggregationSortProperty**](anyOf&lt;microsoft.graph.bucketAggregationSortProperty&gt;.md) | The possible values are count to sort by the number of matches in the aggregation, keyAsStringto sort alphabeticaly based on the key in the aggregation, keyAsNumber for numerical sorting based on the key in the aggregation. Required. | [optional] 

## Methods

### NewMicrosoftGraphBucketAggregationDefinition

`func NewMicrosoftGraphBucketAggregationDefinition() *MicrosoftGraphBucketAggregationDefinition`

NewMicrosoftGraphBucketAggregationDefinition instantiates a new MicrosoftGraphBucketAggregationDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphBucketAggregationDefinitionWithDefaults

`func NewMicrosoftGraphBucketAggregationDefinitionWithDefaults() *MicrosoftGraphBucketAggregationDefinition`

NewMicrosoftGraphBucketAggregationDefinitionWithDefaults instantiates a new MicrosoftGraphBucketAggregationDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDescending

`func (o *MicrosoftGraphBucketAggregationDefinition) GetIsDescending() bool`

GetIsDescending returns the IsDescending field if non-nil, zero value otherwise.

### GetIsDescendingOk

`func (o *MicrosoftGraphBucketAggregationDefinition) GetIsDescendingOk() (*bool, bool)`

GetIsDescendingOk returns a tuple with the IsDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDescending

`func (o *MicrosoftGraphBucketAggregationDefinition) SetIsDescending(v bool)`

SetIsDescending sets IsDescending field to given value.

### HasIsDescending

`func (o *MicrosoftGraphBucketAggregationDefinition) HasIsDescending() bool`

HasIsDescending returns a boolean if a field has been set.

### SetIsDescendingNil

`func (o *MicrosoftGraphBucketAggregationDefinition) SetIsDescendingNil(b bool)`

 SetIsDescendingNil sets the value for IsDescending to be an explicit nil

### UnsetIsDescending
`func (o *MicrosoftGraphBucketAggregationDefinition) UnsetIsDescending()`

UnsetIsDescending ensures that no value is present for IsDescending, not even an explicit nil
### GetMinimumCount

`func (o *MicrosoftGraphBucketAggregationDefinition) GetMinimumCount() int32`

GetMinimumCount returns the MinimumCount field if non-nil, zero value otherwise.

### GetMinimumCountOk

`func (o *MicrosoftGraphBucketAggregationDefinition) GetMinimumCountOk() (*int32, bool)`

GetMinimumCountOk returns a tuple with the MinimumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCount

`func (o *MicrosoftGraphBucketAggregationDefinition) SetMinimumCount(v int32)`

SetMinimumCount sets MinimumCount field to given value.

### HasMinimumCount

`func (o *MicrosoftGraphBucketAggregationDefinition) HasMinimumCount() bool`

HasMinimumCount returns a boolean if a field has been set.

### SetMinimumCountNil

`func (o *MicrosoftGraphBucketAggregationDefinition) SetMinimumCountNil(b bool)`

 SetMinimumCountNil sets the value for MinimumCount to be an explicit nil

### UnsetMinimumCount
`func (o *MicrosoftGraphBucketAggregationDefinition) UnsetMinimumCount()`

UnsetMinimumCount ensures that no value is present for MinimumCount, not even an explicit nil
### GetPrefixFilter

`func (o *MicrosoftGraphBucketAggregationDefinition) GetPrefixFilter() string`

GetPrefixFilter returns the PrefixFilter field if non-nil, zero value otherwise.

### GetPrefixFilterOk

`func (o *MicrosoftGraphBucketAggregationDefinition) GetPrefixFilterOk() (*string, bool)`

GetPrefixFilterOk returns a tuple with the PrefixFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixFilter

`func (o *MicrosoftGraphBucketAggregationDefinition) SetPrefixFilter(v string)`

SetPrefixFilter sets PrefixFilter field to given value.

### HasPrefixFilter

`func (o *MicrosoftGraphBucketAggregationDefinition) HasPrefixFilter() bool`

HasPrefixFilter returns a boolean if a field has been set.

### SetPrefixFilterNil

`func (o *MicrosoftGraphBucketAggregationDefinition) SetPrefixFilterNil(b bool)`

 SetPrefixFilterNil sets the value for PrefixFilter to be an explicit nil

### UnsetPrefixFilter
`func (o *MicrosoftGraphBucketAggregationDefinition) UnsetPrefixFilter()`

UnsetPrefixFilter ensures that no value is present for PrefixFilter, not even an explicit nil
### GetRanges

`func (o *MicrosoftGraphBucketAggregationDefinition) GetRanges() []*AnyOfmicrosoftGraphBucketAggregationRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *MicrosoftGraphBucketAggregationDefinition) GetRangesOk() (*[]*AnyOfmicrosoftGraphBucketAggregationRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *MicrosoftGraphBucketAggregationDefinition) SetRanges(v []*AnyOfmicrosoftGraphBucketAggregationRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *MicrosoftGraphBucketAggregationDefinition) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetSortBy

`func (o *MicrosoftGraphBucketAggregationDefinition) GetSortBy() AnyOfmicrosoftGraphBucketAggregationSortProperty`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *MicrosoftGraphBucketAggregationDefinition) GetSortByOk() (*AnyOfmicrosoftGraphBucketAggregationSortProperty, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *MicrosoftGraphBucketAggregationDefinition) SetSortBy(v AnyOfmicrosoftGraphBucketAggregationSortProperty)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *MicrosoftGraphBucketAggregationDefinition) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *MicrosoftGraphBucketAggregationDefinition) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *MicrosoftGraphBucketAggregationDefinition) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


