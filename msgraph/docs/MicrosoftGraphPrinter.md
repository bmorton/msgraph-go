# MicrosoftGraphPrinter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Capabilities** | Pointer to [**AnyOfmicrosoftGraphPrinterCapabilities**](anyOf&lt;microsoft.graph.printerCapabilities&gt;.md) | The capabilities of the printer/printerShare. | [optional] 
**Defaults** | Pointer to [**AnyOfmicrosoftGraphPrinterDefaults**](anyOf&lt;microsoft.graph.printerDefaults&gt;.md) | The default print settings of printer/printerShare. | [optional] 
**DisplayName** | Pointer to **string** | The name of the printer/printerShare. | [optional] 
**IsAcceptingJobs** | Pointer to **NullableBool** | Whether the printer/printerShare is currently accepting new print jobs. | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphPrinterLocation**](anyOf&lt;microsoft.graph.printerLocation&gt;.md) | The physical and/or organizational location of the printer/printerShare. | [optional] 
**Manufacturer** | Pointer to **NullableString** | The manufacturer of the printer/printerShare. | [optional] 
**Model** | Pointer to **NullableString** | The model name of the printer/printerShare. | [optional] 
**Status** | Pointer to [**MicrosoftGraphPrinterStatus**](MicrosoftGraphPrinterStatus.md) |  | [optional] 
**Jobs** | Pointer to [**[]MicrosoftGraphPrintJob**](MicrosoftGraphPrintJob.md) | The list of jobs that are queued for printing by the printer/printerShare. | [optional] 
**HasPhysicalDevice** | Pointer to **bool** | True if the printer has a physical device for printing. Read-only. | [optional] 
**IsShared** | Pointer to **bool** | True if the printer is shared; false otherwise. Read-only. | [optional] 
**LastSeenDateTime** | Pointer to **NullableTime** | The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only. | [optional] 
**RegisteredDateTime** | Pointer to **time.Time** | The DateTimeOffset when the printer was registered. Read-only. | [optional] 
**Connectors** | Pointer to [**[]MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md) | The connectors that are associated with the printer. | [optional] 
**Shares** | Pointer to [**[]MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md) | The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated with the printer. Read-only. Nullable. | [optional] 
**TaskTriggers** | Pointer to [**[]MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md) | A list of task triggers that are associated with the printer. | [optional] 

## Methods

### NewMicrosoftGraphPrinter

`func NewMicrosoftGraphPrinter() *MicrosoftGraphPrinter`

NewMicrosoftGraphPrinter instantiates a new MicrosoftGraphPrinter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterWithDefaults

`func NewMicrosoftGraphPrinterWithDefaults() *MicrosoftGraphPrinter`

NewMicrosoftGraphPrinterWithDefaults instantiates a new MicrosoftGraphPrinter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrinter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrinter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrinter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrinter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilities

`func (o *MicrosoftGraphPrinter) GetCapabilities() AnyOfmicrosoftGraphPrinterCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MicrosoftGraphPrinter) GetCapabilitiesOk() (*AnyOfmicrosoftGraphPrinterCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MicrosoftGraphPrinter) SetCapabilities(v AnyOfmicrosoftGraphPrinterCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MicrosoftGraphPrinter) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *MicrosoftGraphPrinter) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *MicrosoftGraphPrinter) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDefaults

`func (o *MicrosoftGraphPrinter) GetDefaults() AnyOfmicrosoftGraphPrinterDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *MicrosoftGraphPrinter) GetDefaultsOk() (*AnyOfmicrosoftGraphPrinterDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *MicrosoftGraphPrinter) SetDefaults(v AnyOfmicrosoftGraphPrinterDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *MicrosoftGraphPrinter) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *MicrosoftGraphPrinter) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *MicrosoftGraphPrinter) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphPrinter) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPrinter) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPrinter) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPrinter) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsAcceptingJobs

`func (o *MicrosoftGraphPrinter) GetIsAcceptingJobs() bool`

GetIsAcceptingJobs returns the IsAcceptingJobs field if non-nil, zero value otherwise.

### GetIsAcceptingJobsOk

`func (o *MicrosoftGraphPrinter) GetIsAcceptingJobsOk() (*bool, bool)`

GetIsAcceptingJobsOk returns a tuple with the IsAcceptingJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcceptingJobs

`func (o *MicrosoftGraphPrinter) SetIsAcceptingJobs(v bool)`

SetIsAcceptingJobs sets IsAcceptingJobs field to given value.

### HasIsAcceptingJobs

`func (o *MicrosoftGraphPrinter) HasIsAcceptingJobs() bool`

HasIsAcceptingJobs returns a boolean if a field has been set.

### SetIsAcceptingJobsNil

`func (o *MicrosoftGraphPrinter) SetIsAcceptingJobsNil(b bool)`

 SetIsAcceptingJobsNil sets the value for IsAcceptingJobs to be an explicit nil

### UnsetIsAcceptingJobs
`func (o *MicrosoftGraphPrinter) UnsetIsAcceptingJobs()`

UnsetIsAcceptingJobs ensures that no value is present for IsAcceptingJobs, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphPrinter) GetLocation() AnyOfmicrosoftGraphPrinterLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphPrinter) GetLocationOk() (*AnyOfmicrosoftGraphPrinterLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphPrinter) SetLocation(v AnyOfmicrosoftGraphPrinterLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphPrinter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphPrinter) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphPrinter) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetManufacturer

`func (o *MicrosoftGraphPrinter) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MicrosoftGraphPrinter) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MicrosoftGraphPrinter) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MicrosoftGraphPrinter) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *MicrosoftGraphPrinter) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *MicrosoftGraphPrinter) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetModel

`func (o *MicrosoftGraphPrinter) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MicrosoftGraphPrinter) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MicrosoftGraphPrinter) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MicrosoftGraphPrinter) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MicrosoftGraphPrinter) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MicrosoftGraphPrinter) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphPrinter) GetStatus() MicrosoftGraphPrinterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphPrinter) GetStatusOk() (*MicrosoftGraphPrinterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphPrinter) SetStatus(v MicrosoftGraphPrinterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphPrinter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobs

`func (o *MicrosoftGraphPrinter) GetJobs() []MicrosoftGraphPrintJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *MicrosoftGraphPrinter) GetJobsOk() (*[]MicrosoftGraphPrintJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *MicrosoftGraphPrinter) SetJobs(v []MicrosoftGraphPrintJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *MicrosoftGraphPrinter) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetHasPhysicalDevice

`func (o *MicrosoftGraphPrinter) GetHasPhysicalDevice() bool`

GetHasPhysicalDevice returns the HasPhysicalDevice field if non-nil, zero value otherwise.

### GetHasPhysicalDeviceOk

`func (o *MicrosoftGraphPrinter) GetHasPhysicalDeviceOk() (*bool, bool)`

GetHasPhysicalDeviceOk returns a tuple with the HasPhysicalDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPhysicalDevice

`func (o *MicrosoftGraphPrinter) SetHasPhysicalDevice(v bool)`

SetHasPhysicalDevice sets HasPhysicalDevice field to given value.

### HasHasPhysicalDevice

`func (o *MicrosoftGraphPrinter) HasHasPhysicalDevice() bool`

HasHasPhysicalDevice returns a boolean if a field has been set.

### GetIsShared

`func (o *MicrosoftGraphPrinter) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *MicrosoftGraphPrinter) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *MicrosoftGraphPrinter) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *MicrosoftGraphPrinter) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetLastSeenDateTime

`func (o *MicrosoftGraphPrinter) GetLastSeenDateTime() time.Time`

GetLastSeenDateTime returns the LastSeenDateTime field if non-nil, zero value otherwise.

### GetLastSeenDateTimeOk

`func (o *MicrosoftGraphPrinter) GetLastSeenDateTimeOk() (*time.Time, bool)`

GetLastSeenDateTimeOk returns a tuple with the LastSeenDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenDateTime

`func (o *MicrosoftGraphPrinter) SetLastSeenDateTime(v time.Time)`

SetLastSeenDateTime sets LastSeenDateTime field to given value.

### HasLastSeenDateTime

`func (o *MicrosoftGraphPrinter) HasLastSeenDateTime() bool`

HasLastSeenDateTime returns a boolean if a field has been set.

### SetLastSeenDateTimeNil

`func (o *MicrosoftGraphPrinter) SetLastSeenDateTimeNil(b bool)`

 SetLastSeenDateTimeNil sets the value for LastSeenDateTime to be an explicit nil

### UnsetLastSeenDateTime
`func (o *MicrosoftGraphPrinter) UnsetLastSeenDateTime()`

UnsetLastSeenDateTime ensures that no value is present for LastSeenDateTime, not even an explicit nil
### GetRegisteredDateTime

`func (o *MicrosoftGraphPrinter) GetRegisteredDateTime() time.Time`

GetRegisteredDateTime returns the RegisteredDateTime field if non-nil, zero value otherwise.

### GetRegisteredDateTimeOk

`func (o *MicrosoftGraphPrinter) GetRegisteredDateTimeOk() (*time.Time, bool)`

GetRegisteredDateTimeOk returns a tuple with the RegisteredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDateTime

`func (o *MicrosoftGraphPrinter) SetRegisteredDateTime(v time.Time)`

SetRegisteredDateTime sets RegisteredDateTime field to given value.

### HasRegisteredDateTime

`func (o *MicrosoftGraphPrinter) HasRegisteredDateTime() bool`

HasRegisteredDateTime returns a boolean if a field has been set.

### GetConnectors

`func (o *MicrosoftGraphPrinter) GetConnectors() []MicrosoftGraphPrintConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *MicrosoftGraphPrinter) GetConnectorsOk() (*[]MicrosoftGraphPrintConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *MicrosoftGraphPrinter) SetConnectors(v []MicrosoftGraphPrintConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *MicrosoftGraphPrinter) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetShares

`func (o *MicrosoftGraphPrinter) GetShares() []MicrosoftGraphPrinterShare`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *MicrosoftGraphPrinter) GetSharesOk() (*[]MicrosoftGraphPrinterShare, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *MicrosoftGraphPrinter) SetShares(v []MicrosoftGraphPrinterShare)`

SetShares sets Shares field to given value.

### HasShares

`func (o *MicrosoftGraphPrinter) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTaskTriggers

`func (o *MicrosoftGraphPrinter) GetTaskTriggers() []MicrosoftGraphPrintTaskTrigger`

GetTaskTriggers returns the TaskTriggers field if non-nil, zero value otherwise.

### GetTaskTriggersOk

`func (o *MicrosoftGraphPrinter) GetTaskTriggersOk() (*[]MicrosoftGraphPrintTaskTrigger, bool)`

GetTaskTriggersOk returns a tuple with the TaskTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTriggers

`func (o *MicrosoftGraphPrinter) SetTaskTriggers(v []MicrosoftGraphPrintTaskTrigger)`

SetTaskTriggers sets TaskTriggers field to given value.

### HasTaskTriggers

`func (o *MicrosoftGraphPrinter) HasTaskTriggers() bool`

HasTaskTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


