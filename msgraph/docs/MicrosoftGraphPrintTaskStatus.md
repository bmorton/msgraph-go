# MicrosoftGraphPrintTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A human-readable description of the current processing state of the printTask. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphPrintTaskProcessingState**](anyOf&lt;microsoft.graph.printTaskProcessingState&gt;.md) | The current processing state of the printTask. Valid values are described in the following table. | [optional] 

## Methods

### NewMicrosoftGraphPrintTaskStatus

`func NewMicrosoftGraphPrintTaskStatus() *MicrosoftGraphPrintTaskStatus`

NewMicrosoftGraphPrintTaskStatus instantiates a new MicrosoftGraphPrintTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintTaskStatusWithDefaults

`func NewMicrosoftGraphPrintTaskStatusWithDefaults() *MicrosoftGraphPrintTaskStatus`

NewMicrosoftGraphPrintTaskStatusWithDefaults instantiates a new MicrosoftGraphPrintTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphPrintTaskStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPrintTaskStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphPrintTaskStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphPrintTaskStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphPrintTaskStatus) GetState() AnyOfmicrosoftGraphPrintTaskProcessingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphPrintTaskStatus) GetStateOk() (*AnyOfmicrosoftGraphPrintTaskProcessingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphPrintTaskStatus) SetState(v AnyOfmicrosoftGraphPrintTaskProcessingState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphPrintTaskStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphPrintTaskStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphPrintTaskStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


