# MicrosoftGraphPrint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**AnyOfmicrosoftGraphPrintSettings**](anyOf&lt;microsoft.graph.printSettings&gt;.md) | Tenant-wide settings for the Universal Print service. | [optional] 
**Connectors** | Pointer to [**[]MicrosoftGraphPrintConnector**](MicrosoftGraphPrintConnector.md) | The list of available print connectors. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphPrintOperation**](MicrosoftGraphPrintOperation.md) | The list of print long running operations. | [optional] 
**Printers** | Pointer to [**[]MicrosoftGraphPrinter**](MicrosoftGraphPrinter.md) | The list of printers registered in the tenant. | [optional] 
**Services** | Pointer to [**[]MicrosoftGraphPrintService**](MicrosoftGraphPrintService.md) | The list of available Universal Print service endpoints. | [optional] 
**Shares** | Pointer to [**[]MicrosoftGraphPrinterShare**](MicrosoftGraphPrinterShare.md) | The list of printer shares registered in the tenant. | [optional] 
**TaskDefinitions** | Pointer to [**[]MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) | List of abstract definition for a task that can be triggered when various events occur within Universal Print. | [optional] 

## Methods

### NewMicrosoftGraphPrint

`func NewMicrosoftGraphPrint() *MicrosoftGraphPrint`

NewMicrosoftGraphPrint instantiates a new MicrosoftGraphPrint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintWithDefaults

`func NewMicrosoftGraphPrintWithDefaults() *MicrosoftGraphPrint`

NewMicrosoftGraphPrintWithDefaults instantiates a new MicrosoftGraphPrint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *MicrosoftGraphPrint) GetSettings() AnyOfmicrosoftGraphPrintSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphPrint) GetSettingsOk() (*AnyOfmicrosoftGraphPrintSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MicrosoftGraphPrint) SetSettings(v AnyOfmicrosoftGraphPrintSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MicrosoftGraphPrint) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *MicrosoftGraphPrint) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *MicrosoftGraphPrint) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetConnectors

`func (o *MicrosoftGraphPrint) GetConnectors() []MicrosoftGraphPrintConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *MicrosoftGraphPrint) GetConnectorsOk() (*[]MicrosoftGraphPrintConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *MicrosoftGraphPrint) SetConnectors(v []MicrosoftGraphPrintConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *MicrosoftGraphPrint) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetOperations

`func (o *MicrosoftGraphPrint) GetOperations() []MicrosoftGraphPrintOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphPrint) GetOperationsOk() (*[]MicrosoftGraphPrintOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *MicrosoftGraphPrint) SetOperations(v []MicrosoftGraphPrintOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *MicrosoftGraphPrint) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPrinters

`func (o *MicrosoftGraphPrint) GetPrinters() []MicrosoftGraphPrinter`

GetPrinters returns the Printers field if non-nil, zero value otherwise.

### GetPrintersOk

`func (o *MicrosoftGraphPrint) GetPrintersOk() (*[]MicrosoftGraphPrinter, bool)`

GetPrintersOk returns a tuple with the Printers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinters

`func (o *MicrosoftGraphPrint) SetPrinters(v []MicrosoftGraphPrinter)`

SetPrinters sets Printers field to given value.

### HasPrinters

`func (o *MicrosoftGraphPrint) HasPrinters() bool`

HasPrinters returns a boolean if a field has been set.

### GetServices

`func (o *MicrosoftGraphPrint) GetServices() []MicrosoftGraphPrintService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MicrosoftGraphPrint) GetServicesOk() (*[]MicrosoftGraphPrintService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MicrosoftGraphPrint) SetServices(v []MicrosoftGraphPrintService)`

SetServices sets Services field to given value.

### HasServices

`func (o *MicrosoftGraphPrint) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetShares

`func (o *MicrosoftGraphPrint) GetShares() []MicrosoftGraphPrinterShare`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *MicrosoftGraphPrint) GetSharesOk() (*[]MicrosoftGraphPrinterShare, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *MicrosoftGraphPrint) SetShares(v []MicrosoftGraphPrinterShare)`

SetShares sets Shares field to given value.

### HasShares

`func (o *MicrosoftGraphPrint) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTaskDefinitions

`func (o *MicrosoftGraphPrint) GetTaskDefinitions() []MicrosoftGraphPrintTaskDefinition`

GetTaskDefinitions returns the TaskDefinitions field if non-nil, zero value otherwise.

### GetTaskDefinitionsOk

`func (o *MicrosoftGraphPrint) GetTaskDefinitionsOk() (*[]MicrosoftGraphPrintTaskDefinition, bool)`

GetTaskDefinitionsOk returns a tuple with the TaskDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitions

`func (o *MicrosoftGraphPrint) SetTaskDefinitions(v []MicrosoftGraphPrintTaskDefinition)`

SetTaskDefinitions sets TaskDefinitions field to given value.

### HasTaskDefinitions

`func (o *MicrosoftGraphPrint) HasTaskDefinitions() bool`

HasTaskDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


