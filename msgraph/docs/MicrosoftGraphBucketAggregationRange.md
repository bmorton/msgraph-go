# MicrosoftGraphBucketAggregationRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Defines the lower bound from which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required. | [optional] 
**To** | Pointer to **string** | Defines the upper bound up to which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required. | [optional] 

## Methods

### NewMicrosoftGraphBucketAggregationRange

`func NewMicrosoftGraphBucketAggregationRange() *MicrosoftGraphBucketAggregationRange`

NewMicrosoftGraphBucketAggregationRange instantiates a new MicrosoftGraphBucketAggregationRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphBucketAggregationRangeWithDefaults

`func NewMicrosoftGraphBucketAggregationRangeWithDefaults() *MicrosoftGraphBucketAggregationRange`

NewMicrosoftGraphBucketAggregationRangeWithDefaults instantiates a new MicrosoftGraphBucketAggregationRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MicrosoftGraphBucketAggregationRange) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphBucketAggregationRange) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MicrosoftGraphBucketAggregationRange) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MicrosoftGraphBucketAggregationRange) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MicrosoftGraphBucketAggregationRange) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MicrosoftGraphBucketAggregationRange) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MicrosoftGraphBucketAggregationRange) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MicrosoftGraphBucketAggregationRange) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


