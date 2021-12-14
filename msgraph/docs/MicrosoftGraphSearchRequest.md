# MicrosoftGraphSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationFilters** | Pointer to **[]string** | Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format &#39;{field}:/&#39;{aggregationFilterToken}/&#39;&#39;. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format &#39;{field}:or(/&#39;{aggregationFilterToken1}/&#39;,/&#39;{aggregationFilterToken2}/&#39;)&#39;. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses. | [optional] 
**Aggregations** | Pointer to [**[]AnyOfmicrosoftGraphAggregationOption**](AnyOfmicrosoftGraphAggregationOption.md) | Specifies aggregations (also known as refiners) to be returned alongside search results. Optional. | [optional] 
**ContentSources** | Pointer to **[]string** | Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType&#x3D;externalItem. Optional. | [optional] 
**EnableTopResults** | Pointer to **NullableBool** | This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType&#x3D;message. Optional. | [optional] 
**EntityTypes** | Pointer to [**[]AnyOfmicrosoftGraphEntityType**](AnyOfmicrosoftGraphEntityType.md) | One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem. See known limitations for those combinations of two or more entity types that are supported in the same search request. Required. | [optional] 
**Fields** | Pointer to **[]string** | Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional. | [optional] 
**From** | Pointer to **int32** | Specifies the offset for the search results. Offset 0 returns the very first result. Optional. | [optional] 
**Query** | Pointer to [**MicrosoftGraphSearchQuery**](MicrosoftGraphSearchQuery.md) |  | [optional] 
**Size** | Pointer to **int32** | The size of the page to be retrieved. Optional. | [optional] 
**SortProperties** | Pointer to [**[]AnyOfmicrosoftGraphSortProperty**](AnyOfmicrosoftGraphSortProperty.md) | Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional. | [optional] 

## Methods

### NewMicrosoftGraphSearchRequest

`func NewMicrosoftGraphSearchRequest() *MicrosoftGraphSearchRequest`

NewMicrosoftGraphSearchRequest instantiates a new MicrosoftGraphSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSearchRequestWithDefaults

`func NewMicrosoftGraphSearchRequestWithDefaults() *MicrosoftGraphSearchRequest`

NewMicrosoftGraphSearchRequestWithDefaults instantiates a new MicrosoftGraphSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationFilters

`func (o *MicrosoftGraphSearchRequest) GetAggregationFilters() []*string`

GetAggregationFilters returns the AggregationFilters field if non-nil, zero value otherwise.

### GetAggregationFiltersOk

`func (o *MicrosoftGraphSearchRequest) GetAggregationFiltersOk() (*[]*string, bool)`

GetAggregationFiltersOk returns a tuple with the AggregationFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFilters

`func (o *MicrosoftGraphSearchRequest) SetAggregationFilters(v []*string)`

SetAggregationFilters sets AggregationFilters field to given value.

### HasAggregationFilters

`func (o *MicrosoftGraphSearchRequest) HasAggregationFilters() bool`

HasAggregationFilters returns a boolean if a field has been set.

### GetAggregations

`func (o *MicrosoftGraphSearchRequest) GetAggregations() []*AnyOfmicrosoftGraphAggregationOption`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MicrosoftGraphSearchRequest) GetAggregationsOk() (*[]*AnyOfmicrosoftGraphAggregationOption, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MicrosoftGraphSearchRequest) SetAggregations(v []*AnyOfmicrosoftGraphAggregationOption)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *MicrosoftGraphSearchRequest) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetContentSources

`func (o *MicrosoftGraphSearchRequest) GetContentSources() []*string`

GetContentSources returns the ContentSources field if non-nil, zero value otherwise.

### GetContentSourcesOk

`func (o *MicrosoftGraphSearchRequest) GetContentSourcesOk() (*[]*string, bool)`

GetContentSourcesOk returns a tuple with the ContentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSources

`func (o *MicrosoftGraphSearchRequest) SetContentSources(v []*string)`

SetContentSources sets ContentSources field to given value.

### HasContentSources

`func (o *MicrosoftGraphSearchRequest) HasContentSources() bool`

HasContentSources returns a boolean if a field has been set.

### GetEnableTopResults

`func (o *MicrosoftGraphSearchRequest) GetEnableTopResults() bool`

GetEnableTopResults returns the EnableTopResults field if non-nil, zero value otherwise.

### GetEnableTopResultsOk

`func (o *MicrosoftGraphSearchRequest) GetEnableTopResultsOk() (*bool, bool)`

GetEnableTopResultsOk returns a tuple with the EnableTopResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTopResults

`func (o *MicrosoftGraphSearchRequest) SetEnableTopResults(v bool)`

SetEnableTopResults sets EnableTopResults field to given value.

### HasEnableTopResults

`func (o *MicrosoftGraphSearchRequest) HasEnableTopResults() bool`

HasEnableTopResults returns a boolean if a field has been set.

### SetEnableTopResultsNil

`func (o *MicrosoftGraphSearchRequest) SetEnableTopResultsNil(b bool)`

 SetEnableTopResultsNil sets the value for EnableTopResults to be an explicit nil

### UnsetEnableTopResults
`func (o *MicrosoftGraphSearchRequest) UnsetEnableTopResults()`

UnsetEnableTopResults ensures that no value is present for EnableTopResults, not even an explicit nil
### GetEntityTypes

`func (o *MicrosoftGraphSearchRequest) GetEntityTypes() []*AnyOfmicrosoftGraphEntityType`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *MicrosoftGraphSearchRequest) GetEntityTypesOk() (*[]*AnyOfmicrosoftGraphEntityType, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *MicrosoftGraphSearchRequest) SetEntityTypes(v []*AnyOfmicrosoftGraphEntityType)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *MicrosoftGraphSearchRequest) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### GetFields

`func (o *MicrosoftGraphSearchRequest) GetFields() []*string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MicrosoftGraphSearchRequest) GetFieldsOk() (*[]*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MicrosoftGraphSearchRequest) SetFields(v []*string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MicrosoftGraphSearchRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFrom

`func (o *MicrosoftGraphSearchRequest) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphSearchRequest) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MicrosoftGraphSearchRequest) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MicrosoftGraphSearchRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQuery

`func (o *MicrosoftGraphSearchRequest) GetQuery() MicrosoftGraphSearchQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MicrosoftGraphSearchRequest) GetQueryOk() (*MicrosoftGraphSearchQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MicrosoftGraphSearchRequest) SetQuery(v MicrosoftGraphSearchQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MicrosoftGraphSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSize

`func (o *MicrosoftGraphSearchRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphSearchRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphSearchRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphSearchRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSortProperties

`func (o *MicrosoftGraphSearchRequest) GetSortProperties() []*AnyOfmicrosoftGraphSortProperty`

GetSortProperties returns the SortProperties field if non-nil, zero value otherwise.

### GetSortPropertiesOk

`func (o *MicrosoftGraphSearchRequest) GetSortPropertiesOk() (*[]*AnyOfmicrosoftGraphSortProperty, bool)`

GetSortPropertiesOk returns a tuple with the SortProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortProperties

`func (o *MicrosoftGraphSearchRequest) SetSortProperties(v []*AnyOfmicrosoftGraphSortProperty)`

SetSortProperties sets SortProperties field to given value.

### HasSortProperties

`func (o *MicrosoftGraphSearchRequest) HasSortProperties() bool`

HasSortProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


