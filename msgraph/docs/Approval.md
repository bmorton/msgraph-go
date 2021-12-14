# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | Pointer to [**[]MicrosoftGraphApprovalStage**](MicrosoftGraphApprovalStage.md) | A collection of stages in the approval decision. | [optional] 

## Methods

### NewApproval

`func NewApproval() *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *Approval) GetStages() []MicrosoftGraphApprovalStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *Approval) GetStagesOk() (*[]MicrosoftGraphApprovalStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *Approval) SetStages(v []MicrosoftGraphApprovalStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *Approval) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


