# MicrosoftGraphProvisioningStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Summary of what occurred during the step. | [optional] 
**Details** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Details of what occurred during the step. | [optional] 
**Name** | Pointer to **NullableString** | Name of the step. | [optional] 
**ProvisioningStepType** | Pointer to [**AnyOfmicrosoftGraphProvisioningStepType**](anyOf&lt;microsoft.graph.provisioningStepType&gt;.md) | Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export, unknownFutureValue. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphProvisioningResult**](anyOf&lt;microsoft.graph.provisioningResult&gt;.md) | Status of the step. Possible values are: success, warning,  failure, skipped, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphProvisioningStep

`func NewMicrosoftGraphProvisioningStep() *MicrosoftGraphProvisioningStep`

NewMicrosoftGraphProvisioningStep instantiates a new MicrosoftGraphProvisioningStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisioningStepWithDefaults

`func NewMicrosoftGraphProvisioningStepWithDefaults() *MicrosoftGraphProvisioningStep`

NewMicrosoftGraphProvisioningStepWithDefaults instantiates a new MicrosoftGraphProvisioningStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphProvisioningStep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphProvisioningStep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphProvisioningStep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphProvisioningStep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphProvisioningStep) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphProvisioningStep) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphProvisioningStep) GetDetails() AnyOfobject`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphProvisioningStep) GetDetailsOk() (*AnyOfobject, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphProvisioningStep) SetDetails(v AnyOfobject)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphProvisioningStep) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *MicrosoftGraphProvisioningStep) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *MicrosoftGraphProvisioningStep) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetName

`func (o *MicrosoftGraphProvisioningStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphProvisioningStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphProvisioningStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphProvisioningStep) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphProvisioningStep) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphProvisioningStep) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvisioningStepType

`func (o *MicrosoftGraphProvisioningStep) GetProvisioningStepType() AnyOfmicrosoftGraphProvisioningStepType`

GetProvisioningStepType returns the ProvisioningStepType field if non-nil, zero value otherwise.

### GetProvisioningStepTypeOk

`func (o *MicrosoftGraphProvisioningStep) GetProvisioningStepTypeOk() (*AnyOfmicrosoftGraphProvisioningStepType, bool)`

GetProvisioningStepTypeOk returns a tuple with the ProvisioningStepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStepType

`func (o *MicrosoftGraphProvisioningStep) SetProvisioningStepType(v AnyOfmicrosoftGraphProvisioningStepType)`

SetProvisioningStepType sets ProvisioningStepType field to given value.

### HasProvisioningStepType

`func (o *MicrosoftGraphProvisioningStep) HasProvisioningStepType() bool`

HasProvisioningStepType returns a boolean if a field has been set.

### SetProvisioningStepTypeNil

`func (o *MicrosoftGraphProvisioningStep) SetProvisioningStepTypeNil(b bool)`

 SetProvisioningStepTypeNil sets the value for ProvisioningStepType to be an explicit nil

### UnsetProvisioningStepType
`func (o *MicrosoftGraphProvisioningStep) UnsetProvisioningStepType()`

UnsetProvisioningStepType ensures that no value is present for ProvisioningStepType, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphProvisioningStep) GetStatus() AnyOfmicrosoftGraphProvisioningResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphProvisioningStep) GetStatusOk() (*AnyOfmicrosoftGraphProvisioningResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphProvisioningStep) SetStatus(v AnyOfmicrosoftGraphProvisioningResult)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphProvisioningStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphProvisioningStep) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphProvisioningStep) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


