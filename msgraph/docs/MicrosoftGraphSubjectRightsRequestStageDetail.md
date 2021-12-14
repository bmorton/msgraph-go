# MicrosoftGraphSubjectRightsRequestStageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AnyOfmicrosoftGraphPublicError**](anyOf&lt;microsoft.graph.publicError&gt;.md) | Describes the error, if any, for the current stage. | [optional] 
**Stage** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestStage**](anyOf&lt;microsoft.graph.subjectRightsRequestStage&gt;.md) | The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestStageStatus**](anyOf&lt;microsoft.graph.subjectRightsRequestStageStatus&gt;.md) | Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphSubjectRightsRequestStageDetail

`func NewMicrosoftGraphSubjectRightsRequestStageDetail() *MicrosoftGraphSubjectRightsRequestStageDetail`

NewMicrosoftGraphSubjectRightsRequestStageDetail instantiates a new MicrosoftGraphSubjectRightsRequestStageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSubjectRightsRequestStageDetailWithDefaults

`func NewMicrosoftGraphSubjectRightsRequestStageDetailWithDefaults() *MicrosoftGraphSubjectRightsRequestStageDetail`

NewMicrosoftGraphSubjectRightsRequestStageDetailWithDefaults instantiates a new MicrosoftGraphSubjectRightsRequestStageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetError() AnyOfmicrosoftGraphPublicError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetError(v AnyOfmicrosoftGraphPublicError)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStage

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStage() AnyOfmicrosoftGraphSubjectRightsRequestStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStageOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetStage(v AnyOfmicrosoftGraphSubjectRightsRequestStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStatus() AnyOfmicrosoftGraphSubjectRightsRequestStageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStatusOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetStatus(v AnyOfmicrosoftGraphSubjectRightsRequestStageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphSubjectRightsRequestStageDetail) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


