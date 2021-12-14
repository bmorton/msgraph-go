# MicrosoftGraphSearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentSource** | Pointer to **NullableString** | The name of the content source which the externalItem is part of . | [optional] 
**HitId** | Pointer to **NullableString** | The internal identifier for the item. | [optional] 
**Rank** | Pointer to **NullableInt32** | The rank or the order of the result. | [optional] 
**Summary** | Pointer to **NullableString** | A summary of the result, if a summary is available. | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphSearchHit

`func NewMicrosoftGraphSearchHit() *MicrosoftGraphSearchHit`

NewMicrosoftGraphSearchHit instantiates a new MicrosoftGraphSearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSearchHitWithDefaults

`func NewMicrosoftGraphSearchHitWithDefaults() *MicrosoftGraphSearchHit`

NewMicrosoftGraphSearchHitWithDefaults instantiates a new MicrosoftGraphSearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentSource

`func (o *MicrosoftGraphSearchHit) GetContentSource() string`

GetContentSource returns the ContentSource field if non-nil, zero value otherwise.

### GetContentSourceOk

`func (o *MicrosoftGraphSearchHit) GetContentSourceOk() (*string, bool)`

GetContentSourceOk returns a tuple with the ContentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSource

`func (o *MicrosoftGraphSearchHit) SetContentSource(v string)`

SetContentSource sets ContentSource field to given value.

### HasContentSource

`func (o *MicrosoftGraphSearchHit) HasContentSource() bool`

HasContentSource returns a boolean if a field has been set.

### SetContentSourceNil

`func (o *MicrosoftGraphSearchHit) SetContentSourceNil(b bool)`

 SetContentSourceNil sets the value for ContentSource to be an explicit nil

### UnsetContentSource
`func (o *MicrosoftGraphSearchHit) UnsetContentSource()`

UnsetContentSource ensures that no value is present for ContentSource, not even an explicit nil
### GetHitId

`func (o *MicrosoftGraphSearchHit) GetHitId() string`

GetHitId returns the HitId field if non-nil, zero value otherwise.

### GetHitIdOk

`func (o *MicrosoftGraphSearchHit) GetHitIdOk() (*string, bool)`

GetHitIdOk returns a tuple with the HitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitId

`func (o *MicrosoftGraphSearchHit) SetHitId(v string)`

SetHitId sets HitId field to given value.

### HasHitId

`func (o *MicrosoftGraphSearchHit) HasHitId() bool`

HasHitId returns a boolean if a field has been set.

### SetHitIdNil

`func (o *MicrosoftGraphSearchHit) SetHitIdNil(b bool)`

 SetHitIdNil sets the value for HitId to be an explicit nil

### UnsetHitId
`func (o *MicrosoftGraphSearchHit) UnsetHitId()`

UnsetHitId ensures that no value is present for HitId, not even an explicit nil
### GetRank

`func (o *MicrosoftGraphSearchHit) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *MicrosoftGraphSearchHit) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *MicrosoftGraphSearchHit) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *MicrosoftGraphSearchHit) HasRank() bool`

HasRank returns a boolean if a field has been set.

### SetRankNil

`func (o *MicrosoftGraphSearchHit) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *MicrosoftGraphSearchHit) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetSummary

`func (o *MicrosoftGraphSearchHit) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *MicrosoftGraphSearchHit) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *MicrosoftGraphSearchHit) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *MicrosoftGraphSearchHit) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *MicrosoftGraphSearchHit) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *MicrosoftGraphSearchHit) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphSearchHit) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphSearchHit) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphSearchHit) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphSearchHit) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphSearchHit) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphSearchHit) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


