# MicrosoftGraphSearchAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]AnyOfmicrosoftGraphSearchBucket**](AnyOfmicrosoftGraphSearchBucket.md) | Defines the actual buckets of the computed aggregation. | [optional] 
**Field** | Pointer to **NullableString** | Defines on which field the aggregation was computed on. | [optional] 

## Methods

### NewMicrosoftGraphSearchAggregation

`func NewMicrosoftGraphSearchAggregation() *MicrosoftGraphSearchAggregation`

NewMicrosoftGraphSearchAggregation instantiates a new MicrosoftGraphSearchAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSearchAggregationWithDefaults

`func NewMicrosoftGraphSearchAggregationWithDefaults() *MicrosoftGraphSearchAggregation`

NewMicrosoftGraphSearchAggregationWithDefaults instantiates a new MicrosoftGraphSearchAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *MicrosoftGraphSearchAggregation) GetBuckets() []*AnyOfmicrosoftGraphSearchBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MicrosoftGraphSearchAggregation) GetBucketsOk() (*[]*AnyOfmicrosoftGraphSearchBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MicrosoftGraphSearchAggregation) SetBuckets(v []*AnyOfmicrosoftGraphSearchBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MicrosoftGraphSearchAggregation) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetField

`func (o *MicrosoftGraphSearchAggregation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MicrosoftGraphSearchAggregation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MicrosoftGraphSearchAggregation) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MicrosoftGraphSearchAggregation) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *MicrosoftGraphSearchAggregation) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *MicrosoftGraphSearchAggregation) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


