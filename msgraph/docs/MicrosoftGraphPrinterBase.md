# MicrosoftGraphPrinterBase

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

## Methods

### NewMicrosoftGraphPrinterBase

`func NewMicrosoftGraphPrinterBase() *MicrosoftGraphPrinterBase`

NewMicrosoftGraphPrinterBase instantiates a new MicrosoftGraphPrinterBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterBaseWithDefaults

`func NewMicrosoftGraphPrinterBaseWithDefaults() *MicrosoftGraphPrinterBase`

NewMicrosoftGraphPrinterBaseWithDefaults instantiates a new MicrosoftGraphPrinterBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrinterBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrinterBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrinterBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrinterBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilities

`func (o *MicrosoftGraphPrinterBase) GetCapabilities() AnyOfmicrosoftGraphPrinterCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MicrosoftGraphPrinterBase) GetCapabilitiesOk() (*AnyOfmicrosoftGraphPrinterCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MicrosoftGraphPrinterBase) SetCapabilities(v AnyOfmicrosoftGraphPrinterCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MicrosoftGraphPrinterBase) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *MicrosoftGraphPrinterBase) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *MicrosoftGraphPrinterBase) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDefaults

`func (o *MicrosoftGraphPrinterBase) GetDefaults() AnyOfmicrosoftGraphPrinterDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *MicrosoftGraphPrinterBase) GetDefaultsOk() (*AnyOfmicrosoftGraphPrinterDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *MicrosoftGraphPrinterBase) SetDefaults(v AnyOfmicrosoftGraphPrinterDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *MicrosoftGraphPrinterBase) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *MicrosoftGraphPrinterBase) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *MicrosoftGraphPrinterBase) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphPrinterBase) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPrinterBase) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPrinterBase) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPrinterBase) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsAcceptingJobs

`func (o *MicrosoftGraphPrinterBase) GetIsAcceptingJobs() bool`

GetIsAcceptingJobs returns the IsAcceptingJobs field if non-nil, zero value otherwise.

### GetIsAcceptingJobsOk

`func (o *MicrosoftGraphPrinterBase) GetIsAcceptingJobsOk() (*bool, bool)`

GetIsAcceptingJobsOk returns a tuple with the IsAcceptingJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcceptingJobs

`func (o *MicrosoftGraphPrinterBase) SetIsAcceptingJobs(v bool)`

SetIsAcceptingJobs sets IsAcceptingJobs field to given value.

### HasIsAcceptingJobs

`func (o *MicrosoftGraphPrinterBase) HasIsAcceptingJobs() bool`

HasIsAcceptingJobs returns a boolean if a field has been set.

### SetIsAcceptingJobsNil

`func (o *MicrosoftGraphPrinterBase) SetIsAcceptingJobsNil(b bool)`

 SetIsAcceptingJobsNil sets the value for IsAcceptingJobs to be an explicit nil

### UnsetIsAcceptingJobs
`func (o *MicrosoftGraphPrinterBase) UnsetIsAcceptingJobs()`

UnsetIsAcceptingJobs ensures that no value is present for IsAcceptingJobs, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphPrinterBase) GetLocation() AnyOfmicrosoftGraphPrinterLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphPrinterBase) GetLocationOk() (*AnyOfmicrosoftGraphPrinterLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphPrinterBase) SetLocation(v AnyOfmicrosoftGraphPrinterLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphPrinterBase) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphPrinterBase) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphPrinterBase) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetManufacturer

`func (o *MicrosoftGraphPrinterBase) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MicrosoftGraphPrinterBase) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MicrosoftGraphPrinterBase) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MicrosoftGraphPrinterBase) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *MicrosoftGraphPrinterBase) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *MicrosoftGraphPrinterBase) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetModel

`func (o *MicrosoftGraphPrinterBase) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MicrosoftGraphPrinterBase) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MicrosoftGraphPrinterBase) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MicrosoftGraphPrinterBase) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MicrosoftGraphPrinterBase) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MicrosoftGraphPrinterBase) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphPrinterBase) GetStatus() MicrosoftGraphPrinterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphPrinterBase) GetStatusOk() (*MicrosoftGraphPrinterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphPrinterBase) SetStatus(v MicrosoftGraphPrinterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphPrinterBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobs

`func (o *MicrosoftGraphPrinterBase) GetJobs() []MicrosoftGraphPrintJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *MicrosoftGraphPrinterBase) GetJobsOk() (*[]MicrosoftGraphPrintJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *MicrosoftGraphPrinterBase) SetJobs(v []MicrosoftGraphPrintJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *MicrosoftGraphPrinterBase) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


