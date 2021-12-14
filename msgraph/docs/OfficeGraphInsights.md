# OfficeGraphInsights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shared** | Pointer to [**[]MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md) | Calculated relationship identifying documents shared with or by the user. This includes URLs, file attachments, and reference attachments to OneDrive for Business and SharePoint files found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share. | [optional] 
**Trending** | Pointer to [**[]MicrosoftGraphTrending**](MicrosoftGraphTrending.md) | Calculated relationship identifying documents trending around a user. Trending documents are calculated based on activity of the user&#39;s closest network of people and include files stored in OneDrive for Business and SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but has never viewed before. | [optional] 
**Used** | Pointer to [**[]MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md) | Calculated relationship identifying the latest documents viewed or modified by a user, including OneDrive for Business and SharePoint documents, ranked by recency of use. | [optional] 

## Methods

### NewOfficeGraphInsights

`func NewOfficeGraphInsights() *OfficeGraphInsights`

NewOfficeGraphInsights instantiates a new OfficeGraphInsights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfficeGraphInsightsWithDefaults

`func NewOfficeGraphInsightsWithDefaults() *OfficeGraphInsights`

NewOfficeGraphInsightsWithDefaults instantiates a new OfficeGraphInsights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShared

`func (o *OfficeGraphInsights) GetShared() []MicrosoftGraphSharedInsight`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *OfficeGraphInsights) GetSharedOk() (*[]MicrosoftGraphSharedInsight, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *OfficeGraphInsights) SetShared(v []MicrosoftGraphSharedInsight)`

SetShared sets Shared field to given value.

### HasShared

`func (o *OfficeGraphInsights) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTrending

`func (o *OfficeGraphInsights) GetTrending() []MicrosoftGraphTrending`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *OfficeGraphInsights) GetTrendingOk() (*[]MicrosoftGraphTrending, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *OfficeGraphInsights) SetTrending(v []MicrosoftGraphTrending)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *OfficeGraphInsights) HasTrending() bool`

HasTrending returns a boolean if a field has been set.

### GetUsed

`func (o *OfficeGraphInsights) GetUsed() []MicrosoftGraphUsedInsight`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *OfficeGraphInsights) GetUsedOk() (*[]MicrosoftGraphUsedInsight, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *OfficeGraphInsights) SetUsed(v []MicrosoftGraphUsedInsight)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *OfficeGraphInsights) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


