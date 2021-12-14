# MicrosoftGraphSubjectRightsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AssignedTo** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Identity that the request is assigned to. | [optional] 
**ClosedDateTime** | Pointer to **NullableTime** | The date and time when the request was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity information for the entity that created the request. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the request was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**DataSubject** | Pointer to [**AnyOfmicrosoftGraphDataSubject**](anyOf&lt;microsoft.graph.dataSubject&gt;.md) | Information about the data subject. | [optional] 
**DataSubjectType** | Pointer to [**AnyOfmicrosoftGraphDataSubjectType**](anyOf&lt;microsoft.graph.dataSubjectType&gt;.md) | The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee, student, teacher, faculty, other, unknownFutureValue. | [optional] 
**Description** | Pointer to **NullableString** | Description for the request. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the request. | [optional] 
**History** | Pointer to [**[]AnyOfmicrosoftGraphSubjectRightsRequestHistory**](AnyOfmicrosoftGraphSubjectRightsRequestHistory.md) | Collection of history change events. | [optional] 
**Insight** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestDetail**](anyOf&lt;microsoft.graph.subjectRightsRequestDetail&gt;.md) | Insight about the request. | [optional] 
**InternalDueDateTime** | Pointer to **NullableTime** | The date and time when the request is internally due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity information for the entity that last modified the request. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The date and time when the request was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Regulations** | Pointer to **[]string** | List of regulations that this request will fulfill. | [optional] 
**Stages** | Pointer to [**[]AnyOfmicrosoftGraphSubjectRightsRequestStageDetail**](AnyOfmicrosoftGraphSubjectRightsRequestStageDetail.md) | Information about the different stages for the request. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestStatus**](anyOf&lt;microsoft.graph.subjectRightsRequestStatus&gt;.md) | The status of the request.. Possible values are: active, closed, unknownFutureValue. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphSubjectRightsRequestType**](anyOf&lt;microsoft.graph.subjectRightsRequestType&gt;.md) | The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue. | [optional] 
**Notes** | Pointer to [**[]MicrosoftGraphAuthoredNote**](MicrosoftGraphAuthoredNote.md) | List of notes associcated with the request. | [optional] 
**Team** | Pointer to [**AnyOfmicrosoftGraphTeam**](anyOf&lt;microsoft.graph.team&gt;.md) | Information about the Microsoft Teams team that was created for the request. | [optional] 

## Methods

### NewMicrosoftGraphSubjectRightsRequest

`func NewMicrosoftGraphSubjectRightsRequest() *MicrosoftGraphSubjectRightsRequest`

NewMicrosoftGraphSubjectRightsRequest instantiates a new MicrosoftGraphSubjectRightsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSubjectRightsRequestWithDefaults

`func NewMicrosoftGraphSubjectRightsRequestWithDefaults() *MicrosoftGraphSubjectRightsRequest`

NewMicrosoftGraphSubjectRightsRequestWithDefaults instantiates a new MicrosoftGraphSubjectRightsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSubjectRightsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSubjectRightsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSubjectRightsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignedTo

`func (o *MicrosoftGraphSubjectRightsRequest) GetAssignedTo() AnyOfmicrosoftGraphIdentity`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphSubjectRightsRequest) SetAssignedTo(v AnyOfmicrosoftGraphIdentity)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphSubjectRightsRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetClosedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) GetClosedDateTime() time.Time`

GetClosedDateTime returns the ClosedDateTime field if non-nil, zero value otherwise.

### GetClosedDateTimeOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetClosedDateTimeOk() (*time.Time, bool)`

GetClosedDateTimeOk returns a tuple with the ClosedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) SetClosedDateTime(v time.Time)`

SetClosedDateTime sets ClosedDateTime field to given value.

### HasClosedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) HasClosedDateTime() bool`

HasClosedDateTime returns a boolean if a field has been set.

### SetClosedDateTimeNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetClosedDateTimeNil(b bool)`

 SetClosedDateTimeNil sets the value for ClosedDateTime to be an explicit nil

### UnsetClosedDateTime
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetClosedDateTime()`

UnsetClosedDateTime ensures that no value is present for ClosedDateTime, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphSubjectRightsRequest) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphSubjectRightsRequest) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphSubjectRightsRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDataSubject

`func (o *MicrosoftGraphSubjectRightsRequest) GetDataSubject() AnyOfmicrosoftGraphDataSubject`

GetDataSubject returns the DataSubject field if non-nil, zero value otherwise.

### GetDataSubjectOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetDataSubjectOk() (*AnyOfmicrosoftGraphDataSubject, bool)`

GetDataSubjectOk returns a tuple with the DataSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSubject

`func (o *MicrosoftGraphSubjectRightsRequest) SetDataSubject(v AnyOfmicrosoftGraphDataSubject)`

SetDataSubject sets DataSubject field to given value.

### HasDataSubject

`func (o *MicrosoftGraphSubjectRightsRequest) HasDataSubject() bool`

HasDataSubject returns a boolean if a field has been set.

### SetDataSubjectNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetDataSubjectNil(b bool)`

 SetDataSubjectNil sets the value for DataSubject to be an explicit nil

### UnsetDataSubject
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetDataSubject()`

UnsetDataSubject ensures that no value is present for DataSubject, not even an explicit nil
### GetDataSubjectType

`func (o *MicrosoftGraphSubjectRightsRequest) GetDataSubjectType() AnyOfmicrosoftGraphDataSubjectType`

GetDataSubjectType returns the DataSubjectType field if non-nil, zero value otherwise.

### GetDataSubjectTypeOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetDataSubjectTypeOk() (*AnyOfmicrosoftGraphDataSubjectType, bool)`

GetDataSubjectTypeOk returns a tuple with the DataSubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSubjectType

`func (o *MicrosoftGraphSubjectRightsRequest) SetDataSubjectType(v AnyOfmicrosoftGraphDataSubjectType)`

SetDataSubjectType sets DataSubjectType field to given value.

### HasDataSubjectType

`func (o *MicrosoftGraphSubjectRightsRequest) HasDataSubjectType() bool`

HasDataSubjectType returns a boolean if a field has been set.

### SetDataSubjectTypeNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetDataSubjectTypeNil(b bool)`

 SetDataSubjectTypeNil sets the value for DataSubjectType to be an explicit nil

### UnsetDataSubjectType
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetDataSubjectType()`

UnsetDataSubjectType ensures that no value is present for DataSubjectType, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphSubjectRightsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphSubjectRightsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphSubjectRightsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphSubjectRightsRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphSubjectRightsRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphSubjectRightsRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetHistory

`func (o *MicrosoftGraphSubjectRightsRequest) GetHistory() []*AnyOfmicrosoftGraphSubjectRightsRequestHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetHistoryOk() (*[]*AnyOfmicrosoftGraphSubjectRightsRequestHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *MicrosoftGraphSubjectRightsRequest) SetHistory(v []*AnyOfmicrosoftGraphSubjectRightsRequestHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *MicrosoftGraphSubjectRightsRequest) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetInsight

`func (o *MicrosoftGraphSubjectRightsRequest) GetInsight() AnyOfmicrosoftGraphSubjectRightsRequestDetail`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetInsightOk() (*AnyOfmicrosoftGraphSubjectRightsRequestDetail, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *MicrosoftGraphSubjectRightsRequest) SetInsight(v AnyOfmicrosoftGraphSubjectRightsRequestDetail)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *MicrosoftGraphSubjectRightsRequest) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### SetInsightNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetInsightNil(b bool)`

 SetInsightNil sets the value for Insight to be an explicit nil

### UnsetInsight
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetInsight()`

UnsetInsight ensures that no value is present for Insight, not even an explicit nil
### GetInternalDueDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) GetInternalDueDateTime() time.Time`

GetInternalDueDateTime returns the InternalDueDateTime field if non-nil, zero value otherwise.

### GetInternalDueDateTimeOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetInternalDueDateTimeOk() (*time.Time, bool)`

GetInternalDueDateTimeOk returns a tuple with the InternalDueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDueDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) SetInternalDueDateTime(v time.Time)`

SetInternalDueDateTime sets InternalDueDateTime field to given value.

### HasInternalDueDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) HasInternalDueDateTime() bool`

HasInternalDueDateTime returns a boolean if a field has been set.

### SetInternalDueDateTimeNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetInternalDueDateTimeNil(b bool)`

 SetInternalDueDateTimeNil sets the value for InternalDueDateTime to be an explicit nil

### UnsetInternalDueDateTime
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetInternalDueDateTime()`

UnsetInternalDueDateTime ensures that no value is present for InternalDueDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphSubjectRightsRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSubjectRightsRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphSubjectRightsRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSubjectRightsRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetRegulations

`func (o *MicrosoftGraphSubjectRightsRequest) GetRegulations() []*string`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetRegulationsOk() (*[]*string, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *MicrosoftGraphSubjectRightsRequest) SetRegulations(v []*string)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *MicrosoftGraphSubjectRightsRequest) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.

### GetStages

`func (o *MicrosoftGraphSubjectRightsRequest) GetStages() []*AnyOfmicrosoftGraphSubjectRightsRequestStageDetail`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetStagesOk() (*[]*AnyOfmicrosoftGraphSubjectRightsRequestStageDetail, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *MicrosoftGraphSubjectRightsRequest) SetStages(v []*AnyOfmicrosoftGraphSubjectRightsRequestStageDetail)`

SetStages sets Stages field to given value.

### HasStages

`func (o *MicrosoftGraphSubjectRightsRequest) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphSubjectRightsRequest) GetStatus() AnyOfmicrosoftGraphSubjectRightsRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetStatusOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphSubjectRightsRequest) SetStatus(v AnyOfmicrosoftGraphSubjectRightsRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphSubjectRightsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *MicrosoftGraphSubjectRightsRequest) GetType() AnyOfmicrosoftGraphSubjectRightsRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetTypeOk() (*AnyOfmicrosoftGraphSubjectRightsRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphSubjectRightsRequest) SetType(v AnyOfmicrosoftGraphSubjectRightsRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphSubjectRightsRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetNotes

`func (o *MicrosoftGraphSubjectRightsRequest) GetNotes() []MicrosoftGraphAuthoredNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetNotesOk() (*[]MicrosoftGraphAuthoredNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MicrosoftGraphSubjectRightsRequest) SetNotes(v []MicrosoftGraphAuthoredNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MicrosoftGraphSubjectRightsRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTeam

`func (o *MicrosoftGraphSubjectRightsRequest) GetTeam() AnyOfmicrosoftGraphTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MicrosoftGraphSubjectRightsRequest) GetTeamOk() (*AnyOfmicrosoftGraphTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MicrosoftGraphSubjectRightsRequest) SetTeam(v AnyOfmicrosoftGraphTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MicrosoftGraphSubjectRightsRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *MicrosoftGraphSubjectRightsRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *MicrosoftGraphSubjectRightsRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


