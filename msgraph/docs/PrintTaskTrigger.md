# PrintTaskTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**AnyOfmicrosoftGraphPrintEvent**](anyOf&lt;microsoft.graph.printEvent&gt;.md) | The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table. | [optional] 
**Definition** | Pointer to [**MicrosoftGraphPrintTaskDefinition**](MicrosoftGraphPrintTaskDefinition.md) |  | [optional] 

## Methods

### NewPrintTaskTrigger

`func NewPrintTaskTrigger() *PrintTaskTrigger`

NewPrintTaskTrigger instantiates a new PrintTaskTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintTaskTriggerWithDefaults

`func NewPrintTaskTriggerWithDefaults() *PrintTaskTrigger`

NewPrintTaskTriggerWithDefaults instantiates a new PrintTaskTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PrintTaskTrigger) GetEvent() AnyOfmicrosoftGraphPrintEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PrintTaskTrigger) GetEventOk() (*AnyOfmicrosoftGraphPrintEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PrintTaskTrigger) SetEvent(v AnyOfmicrosoftGraphPrintEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PrintTaskTrigger) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *PrintTaskTrigger) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *PrintTaskTrigger) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetDefinition

`func (o *PrintTaskTrigger) GetDefinition() MicrosoftGraphPrintTaskDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PrintTaskTrigger) GetDefinitionOk() (*MicrosoftGraphPrintTaskDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PrintTaskTrigger) SetDefinition(v MicrosoftGraphPrintTaskDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *PrintTaskTrigger) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


