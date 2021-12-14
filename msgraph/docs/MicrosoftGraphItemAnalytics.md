# MicrosoftGraphItemAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AllTime** | Pointer to [**AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md) |  | [optional] 
**ItemActivityStats** | Pointer to [**[]MicrosoftGraphItemActivityStat**](MicrosoftGraphItemActivityStat.md) |  | [optional] 
**LastSevenDays** | Pointer to [**AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphItemAnalytics

`func NewMicrosoftGraphItemAnalytics() *MicrosoftGraphItemAnalytics`

NewMicrosoftGraphItemAnalytics instantiates a new MicrosoftGraphItemAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphItemAnalyticsWithDefaults

`func NewMicrosoftGraphItemAnalyticsWithDefaults() *MicrosoftGraphItemAnalytics`

NewMicrosoftGraphItemAnalyticsWithDefaults instantiates a new MicrosoftGraphItemAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphItemAnalytics) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphItemAnalytics) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphItemAnalytics) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphItemAnalytics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllTime

`func (o *MicrosoftGraphItemAnalytics) GetAllTime() AnyOfmicrosoftGraphItemActivityStat`

GetAllTime returns the AllTime field if non-nil, zero value otherwise.

### GetAllTimeOk

`func (o *MicrosoftGraphItemAnalytics) GetAllTimeOk() (*AnyOfmicrosoftGraphItemActivityStat, bool)`

GetAllTimeOk returns a tuple with the AllTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTime

`func (o *MicrosoftGraphItemAnalytics) SetAllTime(v AnyOfmicrosoftGraphItemActivityStat)`

SetAllTime sets AllTime field to given value.

### HasAllTime

`func (o *MicrosoftGraphItemAnalytics) HasAllTime() bool`

HasAllTime returns a boolean if a field has been set.

### SetAllTimeNil

`func (o *MicrosoftGraphItemAnalytics) SetAllTimeNil(b bool)`

 SetAllTimeNil sets the value for AllTime to be an explicit nil

### UnsetAllTime
`func (o *MicrosoftGraphItemAnalytics) UnsetAllTime()`

UnsetAllTime ensures that no value is present for AllTime, not even an explicit nil
### GetItemActivityStats

`func (o *MicrosoftGraphItemAnalytics) GetItemActivityStats() []MicrosoftGraphItemActivityStat`

GetItemActivityStats returns the ItemActivityStats field if non-nil, zero value otherwise.

### GetItemActivityStatsOk

`func (o *MicrosoftGraphItemAnalytics) GetItemActivityStatsOk() (*[]MicrosoftGraphItemActivityStat, bool)`

GetItemActivityStatsOk returns a tuple with the ItemActivityStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemActivityStats

`func (o *MicrosoftGraphItemAnalytics) SetItemActivityStats(v []MicrosoftGraphItemActivityStat)`

SetItemActivityStats sets ItemActivityStats field to given value.

### HasItemActivityStats

`func (o *MicrosoftGraphItemAnalytics) HasItemActivityStats() bool`

HasItemActivityStats returns a boolean if a field has been set.

### GetLastSevenDays

`func (o *MicrosoftGraphItemAnalytics) GetLastSevenDays() AnyOfmicrosoftGraphItemActivityStat`

GetLastSevenDays returns the LastSevenDays field if non-nil, zero value otherwise.

### GetLastSevenDaysOk

`func (o *MicrosoftGraphItemAnalytics) GetLastSevenDaysOk() (*AnyOfmicrosoftGraphItemActivityStat, bool)`

GetLastSevenDaysOk returns a tuple with the LastSevenDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSevenDays

`func (o *MicrosoftGraphItemAnalytics) SetLastSevenDays(v AnyOfmicrosoftGraphItemActivityStat)`

SetLastSevenDays sets LastSevenDays field to given value.

### HasLastSevenDays

`func (o *MicrosoftGraphItemAnalytics) HasLastSevenDays() bool`

HasLastSevenDays returns a boolean if a field has been set.

### SetLastSevenDaysNil

`func (o *MicrosoftGraphItemAnalytics) SetLastSevenDaysNil(b bool)`

 SetLastSevenDaysNil sets the value for LastSevenDays to be an explicit nil

### UnsetLastSevenDays
`func (o *MicrosoftGraphItemAnalytics) UnsetLastSevenDays()`

UnsetLastSevenDays ensures that no value is present for LastSevenDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


