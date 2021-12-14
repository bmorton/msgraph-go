# MicrosoftGraphPrintOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A human-readable description of the printOperation&#39;s current processing state. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphPrintOperationProcessingState**](anyOf&lt;microsoft.graph.printOperationProcessingState&gt;.md) | The printOperation&#39;s current processing state. Valid values are described in the following table. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPrintOperationStatus

`func NewMicrosoftGraphPrintOperationStatus() *MicrosoftGraphPrintOperationStatus`

NewMicrosoftGraphPrintOperationStatus instantiates a new MicrosoftGraphPrintOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintOperationStatusWithDefaults

`func NewMicrosoftGraphPrintOperationStatusWithDefaults() *MicrosoftGraphPrintOperationStatus`

NewMicrosoftGraphPrintOperationStatusWithDefaults instantiates a new MicrosoftGraphPrintOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphPrintOperationStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPrintOperationStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphPrintOperationStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphPrintOperationStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphPrintOperationStatus) GetState() AnyOfmicrosoftGraphPrintOperationProcessingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphPrintOperationStatus) GetStateOk() (*AnyOfmicrosoftGraphPrintOperationProcessingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphPrintOperationStatus) SetState(v AnyOfmicrosoftGraphPrintOperationProcessingState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphPrintOperationStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphPrintOperationStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphPrintOperationStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


