# MicrosoftGraphSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HitsContainers** | Pointer to [**[]AnyOfmicrosoftGraphSearchHitsContainer**](AnyOfmicrosoftGraphSearchHitsContainer.md) | A collection of search results. | [optional] 
**SearchTerms** | Pointer to **[]string** | Contains the search terms sent in the initial search query. | [optional] 

## Methods

### NewMicrosoftGraphSearchResponse

`func NewMicrosoftGraphSearchResponse() *MicrosoftGraphSearchResponse`

NewMicrosoftGraphSearchResponse instantiates a new MicrosoftGraphSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSearchResponseWithDefaults

`func NewMicrosoftGraphSearchResponseWithDefaults() *MicrosoftGraphSearchResponse`

NewMicrosoftGraphSearchResponseWithDefaults instantiates a new MicrosoftGraphSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHitsContainers

`func (o *MicrosoftGraphSearchResponse) GetHitsContainers() []*AnyOfmicrosoftGraphSearchHitsContainer`

GetHitsContainers returns the HitsContainers field if non-nil, zero value otherwise.

### GetHitsContainersOk

`func (o *MicrosoftGraphSearchResponse) GetHitsContainersOk() (*[]*AnyOfmicrosoftGraphSearchHitsContainer, bool)`

GetHitsContainersOk returns a tuple with the HitsContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsContainers

`func (o *MicrosoftGraphSearchResponse) SetHitsContainers(v []*AnyOfmicrosoftGraphSearchHitsContainer)`

SetHitsContainers sets HitsContainers field to given value.

### HasHitsContainers

`func (o *MicrosoftGraphSearchResponse) HasHitsContainers() bool`

HasHitsContainers returns a boolean if a field has been set.

### GetSearchTerms

`func (o *MicrosoftGraphSearchResponse) GetSearchTerms() []*string`

GetSearchTerms returns the SearchTerms field if non-nil, zero value otherwise.

### GetSearchTermsOk

`func (o *MicrosoftGraphSearchResponse) GetSearchTermsOk() (*[]*string, bool)`

GetSearchTermsOk returns a tuple with the SearchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerms

`func (o *MicrosoftGraphSearchResponse) SetSearchTerms(v []*string)`

SetSearchTerms sets SearchTerms field to given value.

### HasSearchTerms

`func (o *MicrosoftGraphSearchResponse) HasSearchTerms() bool`

HasSearchTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


