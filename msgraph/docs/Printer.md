# Printer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPhysicalDevice** | Pointer to **bool** | True if the printer has a physical device for printing. Read-only. | [optional] 
**IsShared** | Pointer to **bool** | True if the printer is shared; false otherwise. Read-only. | [optional] 
**LastSeenDateTime** | Pointer to **NullableTime** | The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only. | [optional] 
**RegisteredDateTime** | Pointer to **time.Time** | The DateTimeOffset when the printer was registered. Read-only. | [optional] 
**Connectors** | Pointer to [**[]MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md) | The connectors that are associated with the printer. | [optional] 
**Shares** | Pointer to [**[]MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md) | The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable. | [optional] 
**TaskTriggers** | Pointer to [**[]MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md) | A list of task triggers that are associated with the printer. | [optional] 

## Methods

### NewPrinter

`func NewPrinter() *Printer`

NewPrinter instantiates a new Printer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterWithDefaults

`func NewPrinterWithDefaults() *Printer`

NewPrinterWithDefaults instantiates a new Printer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPhysicalDevice

`func (o *Printer) GetHasPhysicalDevice() bool`

GetHasPhysicalDevice returns the HasPhysicalDevice field if non-nil, zero value otherwise.

### GetHasPhysicalDeviceOk

`func (o *Printer) GetHasPhysicalDeviceOk() (*bool, bool)`

GetHasPhysicalDeviceOk returns a tuple with the HasPhysicalDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPhysicalDevice

`func (o *Printer) SetHasPhysicalDevice(v bool)`

SetHasPhysicalDevice sets HasPhysicalDevice field to given value.

### HasHasPhysicalDevice

`func (o *Printer) HasHasPhysicalDevice() bool`

HasHasPhysicalDevice returns a boolean if a field has been set.

### GetIsShared

`func (o *Printer) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *Printer) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *Printer) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *Printer) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetLastSeenDateTime

`func (o *Printer) GetLastSeenDateTime() time.Time`

GetLastSeenDateTime returns the LastSeenDateTime field if non-nil, zero value otherwise.

### GetLastSeenDateTimeOk

`func (o *Printer) GetLastSeenDateTimeOk() (*time.Time, bool)`

GetLastSeenDateTimeOk returns a tuple with the LastSeenDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenDateTime

`func (o *Printer) SetLastSeenDateTime(v time.Time)`

SetLastSeenDateTime sets LastSeenDateTime field to given value.

### HasLastSeenDateTime

`func (o *Printer) HasLastSeenDateTime() bool`

HasLastSeenDateTime returns a boolean if a field has been set.

### SetLastSeenDateTimeNil

`func (o *Printer) SetLastSeenDateTimeNil(b bool)`

 SetLastSeenDateTimeNil sets the value for LastSeenDateTime to be an explicit nil

### UnsetLastSeenDateTime
`func (o *Printer) UnsetLastSeenDateTime()`

UnsetLastSeenDateTime ensures that no value is present for LastSeenDateTime, not even an explicit nil
### GetRegisteredDateTime

`func (o *Printer) GetRegisteredDateTime() time.Time`

GetRegisteredDateTime returns the RegisteredDateTime field if non-nil, zero value otherwise.

### GetRegisteredDateTimeOk

`func (o *Printer) GetRegisteredDateTimeOk() (*time.Time, bool)`

GetRegisteredDateTimeOk returns a tuple with the RegisteredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDateTime

`func (o *Printer) SetRegisteredDateTime(v time.Time)`

SetRegisteredDateTime sets RegisteredDateTime field to given value.

### HasRegisteredDateTime

`func (o *Printer) HasRegisteredDateTime() bool`

HasRegisteredDateTime returns a boolean if a field has been set.

### GetConnectors

`func (o *Printer) GetConnectors() []MicrosoftGraphPrintConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *Printer) GetConnectorsOk() (*[]MicrosoftGraphPrintConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *Printer) SetConnectors(v []MicrosoftGraphPrintConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *Printer) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetShares

`func (o *Printer) GetShares() []MicrosoftGraphPrinterShare`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *Printer) GetSharesOk() (*[]MicrosoftGraphPrinterShare, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *Printer) SetShares(v []MicrosoftGraphPrinterShare)`

SetShares sets Shares field to given value.

### HasShares

`func (o *Printer) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTaskTriggers

`func (o *Printer) GetTaskTriggers() []MicrosoftGraphPrintTaskTrigger`

GetTaskTriggers returns the TaskTriggers field if non-nil, zero value otherwise.

### GetTaskTriggersOk

`func (o *Printer) GetTaskTriggersOk() (*[]MicrosoftGraphPrintTaskTrigger, bool)`

GetTaskTriggersOk returns a tuple with the TaskTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTriggers

`func (o *Printer) SetTaskTriggers(v []MicrosoftGraphPrintTaskTrigger)`

SetTaskTriggers sets TaskTriggers field to given value.

### HasTaskTriggers

`func (o *Printer) HasTaskTriggers() bool`

HasTaskTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


