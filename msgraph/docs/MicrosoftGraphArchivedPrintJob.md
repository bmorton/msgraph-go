# MicrosoftGraphArchivedPrintJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquiredByPrinter** | Pointer to **bool** | True if the job was acquired by a printer; false otherwise. Read-only. | [optional] 
**AcquiredDateTime** | Pointer to **NullableTime** | The dateTimeOffset when the job was acquired by the printer, if any. Read-only. | [optional] 
**CompletionDateTime** | Pointer to **NullableTime** | The dateTimeOffset when the job was completed, canceled or aborted. Read-only. | [optional] 
**CopiesPrinted** | Pointer to **int32** | The number of copies that were printed. Read-only. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) | The user who created the print job. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The dateTimeOffset when the job was created. Read-only. | [optional] 
**Id** | Pointer to **string** | The archived print job&#39;s GUID. Read-only. | [optional] 
**PrinterId** | Pointer to **NullableString** | The printer ID that the job was queued for. Read-only. | [optional] 
**ProcessingState** | Pointer to [**AnyOfmicrosoftGraphPrintJobProcessingState**](anyOf&lt;microsoft.graph.printJobProcessingState&gt;.md) | The print job&#39;s final processing state. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphArchivedPrintJob

`func NewMicrosoftGraphArchivedPrintJob() *MicrosoftGraphArchivedPrintJob`

NewMicrosoftGraphArchivedPrintJob instantiates a new MicrosoftGraphArchivedPrintJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphArchivedPrintJobWithDefaults

`func NewMicrosoftGraphArchivedPrintJobWithDefaults() *MicrosoftGraphArchivedPrintJob`

NewMicrosoftGraphArchivedPrintJobWithDefaults instantiates a new MicrosoftGraphArchivedPrintJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquiredByPrinter

`func (o *MicrosoftGraphArchivedPrintJob) GetAcquiredByPrinter() bool`

GetAcquiredByPrinter returns the AcquiredByPrinter field if non-nil, zero value otherwise.

### GetAcquiredByPrinterOk

`func (o *MicrosoftGraphArchivedPrintJob) GetAcquiredByPrinterOk() (*bool, bool)`

GetAcquiredByPrinterOk returns a tuple with the AcquiredByPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiredByPrinter

`func (o *MicrosoftGraphArchivedPrintJob) SetAcquiredByPrinter(v bool)`

SetAcquiredByPrinter sets AcquiredByPrinter field to given value.

### HasAcquiredByPrinter

`func (o *MicrosoftGraphArchivedPrintJob) HasAcquiredByPrinter() bool`

HasAcquiredByPrinter returns a boolean if a field has been set.

### GetAcquiredDateTime

`func (o *MicrosoftGraphArchivedPrintJob) GetAcquiredDateTime() time.Time`

GetAcquiredDateTime returns the AcquiredDateTime field if non-nil, zero value otherwise.

### GetAcquiredDateTimeOk

`func (o *MicrosoftGraphArchivedPrintJob) GetAcquiredDateTimeOk() (*time.Time, bool)`

GetAcquiredDateTimeOk returns a tuple with the AcquiredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiredDateTime

`func (o *MicrosoftGraphArchivedPrintJob) SetAcquiredDateTime(v time.Time)`

SetAcquiredDateTime sets AcquiredDateTime field to given value.

### HasAcquiredDateTime

`func (o *MicrosoftGraphArchivedPrintJob) HasAcquiredDateTime() bool`

HasAcquiredDateTime returns a boolean if a field has been set.

### SetAcquiredDateTimeNil

`func (o *MicrosoftGraphArchivedPrintJob) SetAcquiredDateTimeNil(b bool)`

 SetAcquiredDateTimeNil sets the value for AcquiredDateTime to be an explicit nil

### UnsetAcquiredDateTime
`func (o *MicrosoftGraphArchivedPrintJob) UnsetAcquiredDateTime()`

UnsetAcquiredDateTime ensures that no value is present for AcquiredDateTime, not even an explicit nil
### GetCompletionDateTime

`func (o *MicrosoftGraphArchivedPrintJob) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *MicrosoftGraphArchivedPrintJob) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *MicrosoftGraphArchivedPrintJob) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *MicrosoftGraphArchivedPrintJob) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### SetCompletionDateTimeNil

`func (o *MicrosoftGraphArchivedPrintJob) SetCompletionDateTimeNil(b bool)`

 SetCompletionDateTimeNil sets the value for CompletionDateTime to be an explicit nil

### UnsetCompletionDateTime
`func (o *MicrosoftGraphArchivedPrintJob) UnsetCompletionDateTime()`

UnsetCompletionDateTime ensures that no value is present for CompletionDateTime, not even an explicit nil
### GetCopiesPrinted

`func (o *MicrosoftGraphArchivedPrintJob) GetCopiesPrinted() int32`

GetCopiesPrinted returns the CopiesPrinted field if non-nil, zero value otherwise.

### GetCopiesPrintedOk

`func (o *MicrosoftGraphArchivedPrintJob) GetCopiesPrintedOk() (*int32, bool)`

GetCopiesPrintedOk returns a tuple with the CopiesPrinted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopiesPrinted

`func (o *MicrosoftGraphArchivedPrintJob) SetCopiesPrinted(v int32)`

SetCopiesPrinted sets CopiesPrinted field to given value.

### HasCopiesPrinted

`func (o *MicrosoftGraphArchivedPrintJob) HasCopiesPrinted() bool`

HasCopiesPrinted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphArchivedPrintJob) GetCreatedBy() AnyOfmicrosoftGraphUserIdentity`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphArchivedPrintJob) GetCreatedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphArchivedPrintJob) SetCreatedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphArchivedPrintJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphArchivedPrintJob) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphArchivedPrintJob) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphArchivedPrintJob) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphArchivedPrintJob) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphArchivedPrintJob) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphArchivedPrintJob) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetId

`func (o *MicrosoftGraphArchivedPrintJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphArchivedPrintJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphArchivedPrintJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphArchivedPrintJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrinterId

`func (o *MicrosoftGraphArchivedPrintJob) GetPrinterId() string`

GetPrinterId returns the PrinterId field if non-nil, zero value otherwise.

### GetPrinterIdOk

`func (o *MicrosoftGraphArchivedPrintJob) GetPrinterIdOk() (*string, bool)`

GetPrinterIdOk returns a tuple with the PrinterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterId

`func (o *MicrosoftGraphArchivedPrintJob) SetPrinterId(v string)`

SetPrinterId sets PrinterId field to given value.

### HasPrinterId

`func (o *MicrosoftGraphArchivedPrintJob) HasPrinterId() bool`

HasPrinterId returns a boolean if a field has been set.

### SetPrinterIdNil

`func (o *MicrosoftGraphArchivedPrintJob) SetPrinterIdNil(b bool)`

 SetPrinterIdNil sets the value for PrinterId to be an explicit nil

### UnsetPrinterId
`func (o *MicrosoftGraphArchivedPrintJob) UnsetPrinterId()`

UnsetPrinterId ensures that no value is present for PrinterId, not even an explicit nil
### GetProcessingState

`func (o *MicrosoftGraphArchivedPrintJob) GetProcessingState() AnyOfmicrosoftGraphPrintJobProcessingState`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *MicrosoftGraphArchivedPrintJob) GetProcessingStateOk() (*AnyOfmicrosoftGraphPrintJobProcessingState, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *MicrosoftGraphArchivedPrintJob) SetProcessingState(v AnyOfmicrosoftGraphPrintJobProcessingState)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *MicrosoftGraphArchivedPrintJob) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *MicrosoftGraphArchivedPrintJob) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *MicrosoftGraphArchivedPrintJob) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


