# PrintTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentUrl** | Pointer to **string** | The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only. | [optional] 
**Status** | Pointer to [**MicrosoftGraphPrintTaskStatus**](MicrosoftGraphPrintTaskStatus.md) |  | [optional] 
**Definition** | Pointer to [**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) |  | [optional] 
**Trigger** | Pointer to [**MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md) |  | [optional] 

## Methods

### NewPrintTask

`func NewPrintTask() *PrintTask`

NewPrintTask instantiates a new PrintTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintTaskWithDefaults

`func NewPrintTaskWithDefaults() *PrintTask`

NewPrintTaskWithDefaults instantiates a new PrintTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentUrl

`func (o *PrintTask) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *PrintTask) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *PrintTask) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *PrintTask) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetStatus

`func (o *PrintTask) GetStatus() MicrosoftGraphPrintTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrintTask) GetStatusOk() (*MicrosoftGraphPrintTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrintTask) SetStatus(v MicrosoftGraphPrintTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrintTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefinition

`func (o *PrintTask) GetDefinition() MicrosoftGraphPrintTaskDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PrintTask) GetDefinitionOk() (*MicrosoftGraphPrintTaskDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PrintTask) SetDefinition(v MicrosoftGraphPrintTaskDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *PrintTask) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetTrigger

`func (o *PrintTask) GetTrigger() MicrosoftGraphPrintTaskTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PrintTask) GetTriggerOk() (*MicrosoftGraphPrintTaskTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PrintTask) SetTrigger(v MicrosoftGraphPrintTaskTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *PrintTask) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


