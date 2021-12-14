# MicrosoftGraphManagedAppDiagnosticStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MitigationInstruction** | Pointer to **NullableString** | Instruction on how to mitigate a failed validation | [optional] 
**State** | Pointer to **NullableString** | The state of the operation | [optional] 
**ValidationName** | Pointer to **NullableString** | The validation friendly name | [optional] 

## Methods

### NewMicrosoftGraphManagedAppDiagnosticStatus

`func NewMicrosoftGraphManagedAppDiagnosticStatus() *MicrosoftGraphManagedAppDiagnosticStatus`

NewMicrosoftGraphManagedAppDiagnosticStatus instantiates a new MicrosoftGraphManagedAppDiagnosticStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedAppDiagnosticStatusWithDefaults

`func NewMicrosoftGraphManagedAppDiagnosticStatusWithDefaults() *MicrosoftGraphManagedAppDiagnosticStatus`

NewMicrosoftGraphManagedAppDiagnosticStatusWithDefaults instantiates a new MicrosoftGraphManagedAppDiagnosticStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMitigationInstruction

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetMitigationInstruction() string`

GetMitigationInstruction returns the MitigationInstruction field if non-nil, zero value otherwise.

### GetMitigationInstructionOk

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetMitigationInstructionOk() (*string, bool)`

GetMitigationInstructionOk returns a tuple with the MitigationInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationInstruction

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetMitigationInstruction(v string)`

SetMitigationInstruction sets MitigationInstruction field to given value.

### HasMitigationInstruction

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) HasMitigationInstruction() bool`

HasMitigationInstruction returns a boolean if a field has been set.

### SetMitigationInstructionNil

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetMitigationInstructionNil(b bool)`

 SetMitigationInstructionNil sets the value for MitigationInstruction to be an explicit nil

### UnsetMitigationInstruction
`func (o *MicrosoftGraphManagedAppDiagnosticStatus) UnsetMitigationInstruction()`

UnsetMitigationInstruction ensures that no value is present for MitigationInstruction, not even an explicit nil
### GetState

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphManagedAppDiagnosticStatus) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetValidationName

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetValidationName() string`

GetValidationName returns the ValidationName field if non-nil, zero value otherwise.

### GetValidationNameOk

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetValidationNameOk() (*string, bool)`

GetValidationNameOk returns a tuple with the ValidationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationName

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetValidationName(v string)`

SetValidationName sets ValidationName field to given value.

### HasValidationName

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) HasValidationName() bool`

HasValidationName returns a boolean if a field has been set.

### SetValidationNameNil

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetValidationNameNil(b bool)`

 SetValidationNameNil sets the value for ValidationName to be an explicit nil

### UnsetValidationName
`func (o *MicrosoftGraphManagedAppDiagnosticStatus) UnsetValidationName()`

UnsetValidationName ensures that no value is present for ValidationName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


