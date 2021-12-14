# MicrosoftGraphSearchBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationFilterToken** | Pointer to **NullableString** | A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format &#39;{field}:/&#39;{aggregationFilterToken}/&#39;&#39;. See an example. | [optional] 
**Count** | Pointer to **NullableInt32** | The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches. | [optional] 
**Key** | Pointer to **NullableString** | The discrete value of the field that an aggregation was computed on. | [optional] 

## Methods

### NewMicrosoftGraphSearchBucket

`func NewMicrosoftGraphSearchBucket() *MicrosoftGraphSearchBucket`

NewMicrosoftGraphSearchBucket instantiates a new MicrosoftGraphSearchBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSearchBucketWithDefaults

`func NewMicrosoftGraphSearchBucketWithDefaults() *MicrosoftGraphSearchBucket`

NewMicrosoftGraphSearchBucketWithDefaults instantiates a new MicrosoftGraphSearchBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationFilterToken

`func (o *MicrosoftGraphSearchBucket) GetAggregationFilterToken() string`

GetAggregationFilterToken returns the AggregationFilterToken field if non-nil, zero value otherwise.

### GetAggregationFilterTokenOk

`func (o *MicrosoftGraphSearchBucket) GetAggregationFilterTokenOk() (*string, bool)`

GetAggregationFilterTokenOk returns a tuple with the AggregationFilterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFilterToken

`func (o *MicrosoftGraphSearchBucket) SetAggregationFilterToken(v string)`

SetAggregationFilterToken sets AggregationFilterToken field to given value.

### HasAggregationFilterToken

`func (o *MicrosoftGraphSearchBucket) HasAggregationFilterToken() bool`

HasAggregationFilterToken returns a boolean if a field has been set.

### SetAggregationFilterTokenNil

`func (o *MicrosoftGraphSearchBucket) SetAggregationFilterTokenNil(b bool)`

 SetAggregationFilterTokenNil sets the value for AggregationFilterToken to be an explicit nil

### UnsetAggregationFilterToken
`func (o *MicrosoftGraphSearchBucket) UnsetAggregationFilterToken()`

UnsetAggregationFilterToken ensures that no value is present for AggregationFilterToken, not even an explicit nil
### GetCount

`func (o *MicrosoftGraphSearchBucket) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MicrosoftGraphSearchBucket) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MicrosoftGraphSearchBucket) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MicrosoftGraphSearchBucket) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *MicrosoftGraphSearchBucket) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *MicrosoftGraphSearchBucket) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetKey

`func (o *MicrosoftGraphSearchBucket) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MicrosoftGraphSearchBucket) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MicrosoftGraphSearchBucket) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MicrosoftGraphSearchBucket) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MicrosoftGraphSearchBucket) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MicrosoftGraphSearchBucket) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


