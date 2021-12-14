# MicrosoftGraphPrinterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | A human-readable description of the printer&#39;s current processing state. Read-only. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphPrinterProcessingStateDetail**](AnyOfmicrosoftGraphPrinterProcessingStateDetail.md) | The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphPrinterProcessingState**](anyOf&lt;microsoft.graph.printerProcessingState&gt;.md) | The current processing state. Valid values are described in the following table. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPrinterStatus

`func NewMicrosoftGraphPrinterStatus() *MicrosoftGraphPrinterStatus`

NewMicrosoftGraphPrinterStatus instantiates a new MicrosoftGraphPrinterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterStatusWithDefaults

`func NewMicrosoftGraphPrinterStatusWithDefaults() *MicrosoftGraphPrinterStatus`

NewMicrosoftGraphPrinterStatusWithDefaults instantiates a new MicrosoftGraphPrinterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphPrinterStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPrinterStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphPrinterStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphPrinterStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphPrinterStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphPrinterStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphPrinterStatus) GetDetails() []AnyOfmicrosoftGraphPrinterProcessingStateDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPrinterStatus) GetDetailsOk() (*[]AnyOfmicrosoftGraphPrinterProcessingStateDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphPrinterStatus) SetDetails(v []AnyOfmicrosoftGraphPrinterProcessingStateDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphPrinterStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphPrinterStatus) GetState() AnyOfmicrosoftGraphPrinterProcessingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphPrinterStatus) GetStateOk() (*AnyOfmicrosoftGraphPrinterProcessingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphPrinterStatus) SetState(v AnyOfmicrosoftGraphPrinterProcessingState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphPrinterStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphPrinterStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphPrinterStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


