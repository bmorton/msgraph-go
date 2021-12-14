# ItemAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTime** | Pointer to [**AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md) |  | [optional] 
**ItemActivityStats** | Pointer to [**[]MicrosoftGraphItemActivityStat**](MicrosoftGraphItemActivityStat.md) |  | [optional] 
**LastSevenDays** | Pointer to [**AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md) |  | [optional] 

## Methods

### NewItemAnalytics

`func NewItemAnalytics() *ItemAnalytics`

NewItemAnalytics instantiates a new ItemAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAnalyticsWithDefaults

`func NewItemAnalyticsWithDefaults() *ItemAnalytics`

NewItemAnalyticsWithDefaults instantiates a new ItemAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTime

`func (o *ItemAnalytics) GetAllTime() AnyOfmicrosoftGraphItemActivityStat`

GetAllTime returns the AllTime field if non-nil, zero value otherwise.

### GetAllTimeOk

`func (o *ItemAnalytics) GetAllTimeOk() (*AnyOfmicrosoftGraphItemActivityStat, bool)`

GetAllTimeOk returns a tuple with the AllTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTime

`func (o *ItemAnalytics) SetAllTime(v AnyOfmicrosoftGraphItemActivityStat)`

SetAllTime sets AllTime field to given value.

### HasAllTime

`func (o *ItemAnalytics) HasAllTime() bool`

HasAllTime returns a boolean if a field has been set.

### SetAllTimeNil

`func (o *ItemAnalytics) SetAllTimeNil(b bool)`

 SetAllTimeNil sets the value for AllTime to be an explicit nil

### UnsetAllTime
`func (o *ItemAnalytics) UnsetAllTime()`

UnsetAllTime ensures that no value is present for AllTime, not even an explicit nil
### GetItemActivityStats

`func (o *ItemAnalytics) GetItemActivityStats() []MicrosoftGraphItemActivityStat`

GetItemActivityStats returns the ItemActivityStats field if non-nil, zero value otherwise.

### GetItemActivityStatsOk

`func (o *ItemAnalytics) GetItemActivityStatsOk() (*[]MicrosoftGraphItemActivityStat, bool)`

GetItemActivityStatsOk returns a tuple with the ItemActivityStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemActivityStats

`func (o *ItemAnalytics) SetItemActivityStats(v []MicrosoftGraphItemActivityStat)`

SetItemActivityStats sets ItemActivityStats field to given value.

### HasItemActivityStats

`func (o *ItemAnalytics) HasItemActivityStats() bool`

HasItemActivityStats returns a boolean if a field has been set.

### GetLastSevenDays

`func (o *ItemAnalytics) GetLastSevenDays() AnyOfmicrosoftGraphItemActivityStat`

GetLastSevenDays returns the LastSevenDays field if non-nil, zero value otherwise.

### GetLastSevenDaysOk

`func (o *ItemAnalytics) GetLastSevenDaysOk() (*AnyOfmicrosoftGraphItemActivityStat, bool)`

GetLastSevenDaysOk returns a tuple with the LastSevenDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSevenDays

`func (o *ItemAnalytics) SetLastSevenDays(v AnyOfmicrosoftGraphItemActivityStat)`

SetLastSevenDays sets LastSevenDays field to given value.

### HasLastSevenDays

`func (o *ItemAnalytics) HasLastSevenDays() bool`

HasLastSevenDays returns a boolean if a field has been set.

### SetLastSevenDaysNil

`func (o *ItemAnalytics) SetLastSevenDaysNil(b bool)`

 SetLastSevenDaysNil sets the value for LastSevenDays to be an explicit nil

### UnsetLastSevenDays
`func (o *ItemAnalytics) UnsetLastSevenDays()`

UnsetLastSevenDays ensures that no value is present for LastSevenDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


