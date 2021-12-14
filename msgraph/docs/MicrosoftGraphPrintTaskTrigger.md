# MicrosoftGraphPrintTaskTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Event** | Pointer to [**AnyOfmicrosoftGraphPrintEvent**](anyOf&lt;microsoft.graph.printEvent&gt;.md) | The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table. | [optional] 
**Definition** | Pointer to [**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) |  | [optional] 

## Methods

### NewMicrosoftGraphPrintTaskTrigger

`func NewMicrosoftGraphPrintTaskTrigger() *MicrosoftGraphPrintTaskTrigger`

NewMicrosoftGraphPrintTaskTrigger instantiates a new MicrosoftGraphPrintTaskTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintTaskTriggerWithDefaults

`func NewMicrosoftGraphPrintTaskTriggerWithDefaults() *MicrosoftGraphPrintTaskTrigger`

NewMicrosoftGraphPrintTaskTriggerWithDefaults instantiates a new MicrosoftGraphPrintTaskTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintTaskTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintTaskTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintTaskTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintTaskTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *MicrosoftGraphPrintTaskTrigger) GetEvent() AnyOfmicrosoftGraphPrintEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MicrosoftGraphPrintTaskTrigger) GetEventOk() (*AnyOfmicrosoftGraphPrintEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MicrosoftGraphPrintTaskTrigger) SetEvent(v AnyOfmicrosoftGraphPrintEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *MicrosoftGraphPrintTaskTrigger) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *MicrosoftGraphPrintTaskTrigger) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *MicrosoftGraphPrintTaskTrigger) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetDefinition

`func (o *MicrosoftGraphPrintTaskTrigger) GetDefinition() MicrosoftGraphPrintTaskDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *MicrosoftGraphPrintTaskTrigger) GetDefinitionOk() (*MicrosoftGraphPrintTaskDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *MicrosoftGraphPrintTaskTrigger) SetDefinition(v MicrosoftGraphPrintTaskDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *MicrosoftGraphPrintTaskTrigger) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


