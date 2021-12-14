# MicrosoftGraphSearchHitsContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**[]AnyOfmicrosoftGraphSearchAggregation**](AnyOfmicrosoftGraphSearchAggregation.md) | Contains the collection of aggregations computed based on the provided aggregationOption specified in the request. | [optional] 
**Hits** | Pointer to [**[]AnyOfmicrosoftGraphSearchHit**](AnyOfmicrosoftGraphSearchHit.md) | A collection of the search results. | [optional] 
**MoreResultsAvailable** | Pointer to **NullableBool** | Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly. | [optional] 
**Total** | Pointer to **NullableInt32** | The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query. | [optional] 

## Methods

### NewMicrosoftGraphSearchHitsContainer

`func NewMicrosoftGraphSearchHitsContainer() *MicrosoftGraphSearchHitsContainer`

NewMicrosoftGraphSearchHitsContainer instantiates a new MicrosoftGraphSearchHitsContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSearchHitsContainerWithDefaults

`func NewMicrosoftGraphSearchHitsContainerWithDefaults() *MicrosoftGraphSearchHitsContainer`

NewMicrosoftGraphSearchHitsContainerWithDefaults instantiates a new MicrosoftGraphSearchHitsContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *MicrosoftGraphSearchHitsContainer) GetAggregations() []*AnyOfmicrosoftGraphSearchAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MicrosoftGraphSearchHitsContainer) GetAggregationsOk() (*[]*AnyOfmicrosoftGraphSearchAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MicrosoftGraphSearchHitsContainer) SetAggregations(v []*AnyOfmicrosoftGraphSearchAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *MicrosoftGraphSearchHitsContainer) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetHits

`func (o *MicrosoftGraphSearchHitsContainer) GetHits() []*AnyOfmicrosoftGraphSearchHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *MicrosoftGraphSearchHitsContainer) GetHitsOk() (*[]*AnyOfmicrosoftGraphSearchHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *MicrosoftGraphSearchHitsContainer) SetHits(v []*AnyOfmicrosoftGraphSearchHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *MicrosoftGraphSearchHitsContainer) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMoreResultsAvailable

`func (o *MicrosoftGraphSearchHitsContainer) GetMoreResultsAvailable() bool`

GetMoreResultsAvailable returns the MoreResultsAvailable field if non-nil, zero value otherwise.

### GetMoreResultsAvailableOk

`func (o *MicrosoftGraphSearchHitsContainer) GetMoreResultsAvailableOk() (*bool, bool)`

GetMoreResultsAvailableOk returns a tuple with the MoreResultsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreResultsAvailable

`func (o *MicrosoftGraphSearchHitsContainer) SetMoreResultsAvailable(v bool)`

SetMoreResultsAvailable sets MoreResultsAvailable field to given value.

### HasMoreResultsAvailable

`func (o *MicrosoftGraphSearchHitsContainer) HasMoreResultsAvailable() bool`

HasMoreResultsAvailable returns a boolean if a field has been set.

### SetMoreResultsAvailableNil

`func (o *MicrosoftGraphSearchHitsContainer) SetMoreResultsAvailableNil(b bool)`

 SetMoreResultsAvailableNil sets the value for MoreResultsAvailable to be an explicit nil

### UnsetMoreResultsAvailable
`func (o *MicrosoftGraphSearchHitsContainer) UnsetMoreResultsAvailable()`

UnsetMoreResultsAvailable ensures that no value is present for MoreResultsAvailable, not even an explicit nil
### GetTotal

`func (o *MicrosoftGraphSearchHitsContainer) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MicrosoftGraphSearchHitsContainer) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MicrosoftGraphSearchHitsContainer) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MicrosoftGraphSearchHitsContainer) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *MicrosoftGraphSearchHitsContainer) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *MicrosoftGraphSearchHitsContainer) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


