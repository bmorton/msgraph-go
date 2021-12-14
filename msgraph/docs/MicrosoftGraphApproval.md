# MicrosoftGraphApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Stages** | Pointer to [**[]MicrosoftGraphApprovalStage**](MicrosoftGraphApprovalStage.md) | A collection of stages in the approval decision. | [optional] 

## Methods

### NewMicrosoftGraphApproval

`func NewMicrosoftGraphApproval() *MicrosoftGraphApproval`

NewMicrosoftGraphApproval instantiates a new MicrosoftGraphApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphApprovalWithDefaults

`func NewMicrosoftGraphApprovalWithDefaults() *MicrosoftGraphApproval`

NewMicrosoftGraphApprovalWithDefaults instantiates a new MicrosoftGraphApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphApproval) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphApproval) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphApproval) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphApproval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStages

`func (o *MicrosoftGraphApproval) GetStages() []MicrosoftGraphApprovalStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *MicrosoftGraphApproval) GetStagesOk() (*[]MicrosoftGraphApprovalStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *MicrosoftGraphApproval) SetStages(v []MicrosoftGraphApprovalStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *MicrosoftGraphApproval) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


