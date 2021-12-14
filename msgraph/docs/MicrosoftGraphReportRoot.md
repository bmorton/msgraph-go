# MicrosoftGraphReportRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DailyPrintUsageByPrinter** | Pointer to [**[]MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md) |  | [optional] 
**DailyPrintUsageByUser** | Pointer to [**[]MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md) |  | [optional] 
**MonthlyPrintUsageByPrinter** | Pointer to [**[]MicrosoftGraphPrintUsageByPrinter**](MicrosoftGraphPrintUsageByPrinter.md) |  | [optional] 
**MonthlyPrintUsageByUser** | Pointer to [**[]MicrosoftGraphPrintUsageByUser**](MicrosoftGraphPrintUsageByUser.md) |  | [optional] 

## Methods

### NewMicrosoftGraphReportRoot

`func NewMicrosoftGraphReportRoot() *MicrosoftGraphReportRoot`

NewMicrosoftGraphReportRoot instantiates a new MicrosoftGraphReportRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphReportRootWithDefaults

`func NewMicrosoftGraphReportRootWithDefaults() *MicrosoftGraphReportRoot`

NewMicrosoftGraphReportRootWithDefaults instantiates a new MicrosoftGraphReportRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphReportRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphReportRoot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphReportRoot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphReportRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDailyPrintUsageByPrinter

`func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByPrinter() []MicrosoftGraphPrintUsageByPrinter`

GetDailyPrintUsageByPrinter returns the DailyPrintUsageByPrinter field if non-nil, zero value otherwise.

### GetDailyPrintUsageByPrinterOk

`func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByPrinterOk() (*[]MicrosoftGraphPrintUsageByPrinter, bool)`

GetDailyPrintUsageByPrinterOk returns a tuple with the DailyPrintUsageByPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyPrintUsageByPrinter

`func (o *MicrosoftGraphReportRoot) SetDailyPrintUsageByPrinter(v []MicrosoftGraphPrintUsageByPrinter)`

SetDailyPrintUsageByPrinter sets DailyPrintUsageByPrinter field to given value.

### HasDailyPrintUsageByPrinter

`func (o *MicrosoftGraphReportRoot) HasDailyPrintUsageByPrinter() bool`

HasDailyPrintUsageByPrinter returns a boolean if a field has been set.

### GetDailyPrintUsageByUser

`func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByUser() []MicrosoftGraphPrintUsageByUser`

GetDailyPrintUsageByUser returns the DailyPrintUsageByUser field if non-nil, zero value otherwise.

### GetDailyPrintUsageByUserOk

`func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByUserOk() (*[]MicrosoftGraphPrintUsageByUser, bool)`

GetDailyPrintUsageByUserOk returns a tuple with the DailyPrintUsageByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyPrintUsageByUser

`func (o *MicrosoftGraphReportRoot) SetDailyPrintUsageByUser(v []MicrosoftGraphPrintUsageByUser)`

SetDailyPrintUsageByUser sets DailyPrintUsageByUser field to given value.

### HasDailyPrintUsageByUser

`func (o *MicrosoftGraphReportRoot) HasDailyPrintUsageByUser() bool`

HasDailyPrintUsageByUser returns a boolean if a field has been set.

### GetMonthlyPrintUsageByPrinter

`func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByPrinter() []MicrosoftGraphPrintUsageByPrinter`

GetMonthlyPrintUsageByPrinter returns the MonthlyPrintUsageByPrinter field if non-nil, zero value otherwise.

### GetMonthlyPrintUsageByPrinterOk

`func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByPrinterOk() (*[]MicrosoftGraphPrintUsageByPrinter, bool)`

GetMonthlyPrintUsageByPrinterOk returns a tuple with the MonthlyPrintUsageByPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrintUsageByPrinter

`func (o *MicrosoftGraphReportRoot) SetMonthlyPrintUsageByPrinter(v []MicrosoftGraphPrintUsageByPrinter)`

SetMonthlyPrintUsageByPrinter sets MonthlyPrintUsageByPrinter field to given value.

### HasMonthlyPrintUsageByPrinter

`func (o *MicrosoftGraphReportRoot) HasMonthlyPrintUsageByPrinter() bool`

HasMonthlyPrintUsageByPrinter returns a boolean if a field has been set.

### GetMonthlyPrintUsageByUser

`func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByUser() []MicrosoftGraphPrintUsageByUser`

GetMonthlyPrintUsageByUser returns the MonthlyPrintUsageByUser field if non-nil, zero value otherwise.

### GetMonthlyPrintUsageByUserOk

`func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByUserOk() (*[]MicrosoftGraphPrintUsageByUser, bool)`

GetMonthlyPrintUsageByUserOk returns a tuple with the MonthlyPrintUsageByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrintUsageByUser

`func (o *MicrosoftGraphReportRoot) SetMonthlyPrintUsageByUser(v []MicrosoftGraphPrintUsageByUser)`

SetMonthlyPrintUsageByUser sets MonthlyPrintUsageByUser field to given value.

### HasMonthlyPrintUsageByUser

`func (o *MicrosoftGraphReportRoot) HasMonthlyPrintUsageByUser() bool`

HasMonthlyPrintUsageByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


