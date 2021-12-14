# MicrosoftGraphPrintTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ParentUrl** | Pointer to **string** | The URL for the print entity that triggered this task. For example, https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only. | [optional] 
**Status** | Pointer to [**MicrosoftGraphPrintTaskStatus**](MicrosoftGraphPrintTaskStatus.md) |  | [optional] 
**Definition** | Pointer to [**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) |  | [optional] 
**Trigger** | Pointer to [**MicrosoftGraphPrintTaskTrigger**](MicrosoftGraphPrintTaskTrigger.md) |  | [optional] 

## Methods

### NewMicrosoftGraphPrintTask

`func NewMicrosoftGraphPrintTask() *MicrosoftGraphPrintTask`

NewMicrosoftGraphPrintTask instantiates a new MicrosoftGraphPrintTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintTaskWithDefaults

`func NewMicrosoftGraphPrintTaskWithDefaults() *MicrosoftGraphPrintTask`

NewMicrosoftGraphPrintTaskWithDefaults instantiates a new MicrosoftGraphPrintTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentUrl

`func (o *MicrosoftGraphPrintTask) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *MicrosoftGraphPrintTask) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *MicrosoftGraphPrintTask) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *MicrosoftGraphPrintTask) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphPrintTask) GetStatus() MicrosoftGraphPrintTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphPrintTask) GetStatusOk() (*MicrosoftGraphPrintTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphPrintTask) SetStatus(v MicrosoftGraphPrintTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphPrintTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefinition

`func (o *MicrosoftGraphPrintTask) GetDefinition() MicrosoftGraphPrintTaskDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *MicrosoftGraphPrintTask) GetDefinitionOk() (*MicrosoftGraphPrintTaskDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *MicrosoftGraphPrintTask) SetDefinition(v MicrosoftGraphPrintTaskDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *MicrosoftGraphPrintTask) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetTrigger

`func (o *MicrosoftGraphPrintTask) GetTrigger() MicrosoftGraphPrintTaskTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *MicrosoftGraphPrintTask) GetTriggerOk() (*MicrosoftGraphPrintTaskTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *MicrosoftGraphPrintTask) SetTrigger(v MicrosoftGraphPrintTaskTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *MicrosoftGraphPrintTask) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


