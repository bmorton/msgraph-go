# MicrosoftGraphSubjectRightsRequestHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user who changed the  subject rights request. | [optional] 
**EventDateTime** | Pointer to **NullableTime** | Data and time when the entity was changed. | [optional] 
**Stage** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestStage**](anyOf&lt;microsoft.graph.subjectRightsRequestStage&gt;.md) | The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue. | [optional] 
**StageStatus** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestStageStatus**](anyOf&lt;microsoft.graph.subjectRightsRequestStageStatus&gt;.md) | The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue. | [optional] 
**Type** | Pointer to **NullableString** | Type of history. | [optional] 

## Methods

### NewMicrosoftGraphSubjectRightsRequestHistory

`func NewMicrosoftGraphSubjectRightsRequestHistory() *MicrosoftGraphSubjectRightsRequestHistory`

NewMicrosoftGraphSubjectRightsRequestHistory instantiates a new MicrosoftGraphSubjectRightsRequestHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSubjectRightsRequestHistoryWithDefaults

`func NewMicrosoftGraphSubjectRightsRequestHistoryWithDefaults() *MicrosoftGraphSubjectRightsRequestHistory`

NewMicrosoftGraphSubjectRightsRequestHistoryWithDefaults instantiates a new MicrosoftGraphSubjectRightsRequestHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedBy

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetChangedBy() AnyOfmicrosoftGraphIdentitySet`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetChangedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetChangedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetChangedBy sets ChangedBy field to given value.

### HasChangedBy

`func (o *MicrosoftGraphSubjectRightsRequestHistory) HasChangedBy() bool`

HasChangedBy returns a boolean if a field has been set.

### SetChangedByNil

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetChangedByNil(b bool)`

 SetChangedByNil sets the value for ChangedBy to be an explicit nil

### UnsetChangedBy
`func (o *MicrosoftGraphSubjectRightsRequestHistory) UnsetChangedBy()`

UnsetChangedBy ensures that no value is present for ChangedBy, not even an explicit nil
### GetEventDateTime

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *MicrosoftGraphSubjectRightsRequestHistory) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTimeNil

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetEventDateTimeNil(b bool)`

 SetEventDateTimeNil sets the value for EventDateTime to be an explicit nil

### UnsetEventDateTime
`func (o *MicrosoftGraphSubjectRightsRequestHistory) UnsetEventDateTime()`

UnsetEventDateTime ensures that no value is present for EventDateTime, not even an explicit nil
### GetStage

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetStage() AnyOfmicrosoftGraphSubjectRightsRequestStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetStageOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetStage(v AnyOfmicrosoftGraphSubjectRightsRequestStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *MicrosoftGraphSubjectRightsRequestHistory) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *MicrosoftGraphSubjectRightsRequestHistory) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetStageStatus

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetStageStatus() AnyOfmicrosoftGraphSubjectRightsRequestStageStatus`

GetStageStatus returns the StageStatus field if non-nil, zero value otherwise.

### GetStageStatusOk

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetStageStatusOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStageStatus, bool)`

GetStageStatusOk returns a tuple with the StageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStatus

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetStageStatus(v AnyOfmicrosoftGraphSubjectRightsRequestStageStatus)`

SetStageStatus sets StageStatus field to given value.

### HasStageStatus

`func (o *MicrosoftGraphSubjectRightsRequestHistory) HasStageStatus() bool`

HasStageStatus returns a boolean if a field has been set.

### SetStageStatusNil

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetStageStatusNil(b bool)`

 SetStageStatusNil sets the value for StageStatus to be an explicit nil

### UnsetStageStatus
`func (o *MicrosoftGraphSubjectRightsRequestHistory) UnsetStageStatus()`

UnsetStageStatus ensures that no value is present for StageStatus, not even an explicit nil
### GetType

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphSubjectRightsRequestHistory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphSubjectRightsRequestHistory) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphSubjectRightsRequestHistory) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphSubjectRightsRequestHistory) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


