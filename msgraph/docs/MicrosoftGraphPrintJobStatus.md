# MicrosoftGraphPrintJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A human-readable description of the print job&#39;s current processing state. Read-only. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphPrintJobStateDetail**](AnyOfmicrosoftGraphPrintJobStateDetail.md) | Additional details for print job state. Valid values are described in the following table. Read-only. | [optional] 
**IsAcquiredByPrinter** | Pointer to **bool** | True if the job was acknowledged by a printer; false otherwise. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphPrintJobProcessingState**](anyOf&lt;microsoft.graph.printJobProcessingState&gt;.md) | The print job&#39;s current processing state. Valid values are described in the following table. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPrintJobStatus

`func NewMicrosoftGraphPrintJobStatus() *MicrosoftGraphPrintJobStatus`

NewMicrosoftGraphPrintJobStatus instantiates a new MicrosoftGraphPrintJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintJobStatusWithDefaults

`func NewMicrosoftGraphPrintJobStatusWithDefaults() *MicrosoftGraphPrintJobStatus`

NewMicrosoftGraphPrintJobStatusWithDefaults instantiates a new MicrosoftGraphPrintJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphPrintJobStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPrintJobStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphPrintJobStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphPrintJobStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *MicrosoftGraphPrintJobStatus) GetDetails() []AnyOfmicrosoftGraphPrintJobStateDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPrintJobStatus) GetDetailsOk() (*[]AnyOfmicrosoftGraphPrintJobStateDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphPrintJobStatus) SetDetails(v []AnyOfmicrosoftGraphPrintJobStateDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphPrintJobStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetIsAcquiredByPrinter

`func (o *MicrosoftGraphPrintJobStatus) GetIsAcquiredByPrinter() bool`

GetIsAcquiredByPrinter returns the IsAcquiredByPrinter field if non-nil, zero value otherwise.

### GetIsAcquiredByPrinterOk

`func (o *MicrosoftGraphPrintJobStatus) GetIsAcquiredByPrinterOk() (*bool, bool)`

GetIsAcquiredByPrinterOk returns a tuple with the IsAcquiredByPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcquiredByPrinter

`func (o *MicrosoftGraphPrintJobStatus) SetIsAcquiredByPrinter(v bool)`

SetIsAcquiredByPrinter sets IsAcquiredByPrinter field to given value.

### HasIsAcquiredByPrinter

`func (o *MicrosoftGraphPrintJobStatus) HasIsAcquiredByPrinter() bool`

HasIsAcquiredByPrinter returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphPrintJobStatus) GetState() AnyOfmicrosoftGraphPrintJobProcessingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphPrintJobStatus) GetStateOk() (*AnyOfmicrosoftGraphPrintJobProcessingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphPrintJobStatus) SetState(v AnyOfmicrosoftGraphPrintJobProcessingState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphPrintJobStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphPrintJobStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphPrintJobStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


