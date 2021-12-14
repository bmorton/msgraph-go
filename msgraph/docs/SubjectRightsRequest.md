# SubjectRightsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSubjectRightsRequest

`func NewSubjectRightsRequest() *SubjectRightsRequest`

NewSubjectRightsRequest instantiates a new SubjectRightsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectRightsRequestWithDefaults

`func NewSubjectRightsRequestWithDefaults() *SubjectRightsRequest`

NewSubjectRightsRequestWithDefaults instantiates a new SubjectRightsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *SubjectRightsRequest) GetAssignedTo() AnyOfmicrosoftGraphIdentity`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *SubjectRightsRequest) GetAssignedToOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *SubjectRightsRequest) SetAssignedTo(v AnyOfmicrosoftGraphIdentity)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *SubjectRightsRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *SubjectRightsRequest) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *SubjectRightsRequest) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetClosedDateTime

`func (o *SubjectRightsRequest) GetClosedDateTime() time.Time`

GetClosedDateTime returns the ClosedDateTime field if non-nil, zero value otherwise.

### GetClosedDateTimeOk

`func (o *SubjectRightsRequest) GetClosedDateTimeOk() (*time.Time, bool)`

GetClosedDateTimeOk returns a tuple with the ClosedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDateTime

`func (o *SubjectRightsRequest) SetClosedDateTime(v time.Time)`

SetClosedDateTime sets ClosedDateTime field to given value.

### HasClosedDateTime

`func (o *SubjectRightsRequest) HasClosedDateTime() bool`

HasClosedDateTime returns a boolean if a field has been set.

### SetClosedDateTimeNil

`func (o *SubjectRightsRequest) SetClosedDateTimeNil(b bool)`

 SetClosedDateTimeNil sets the value for ClosedDateTime to be an explicit nil

### UnsetClosedDateTime
`func (o *SubjectRightsRequest) UnsetClosedDateTime()`

UnsetClosedDateTime ensures that no value is present for ClosedDateTime, not even an explicit nil
### GetCreatedBy

`func (o *SubjectRightsRequest) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SubjectRightsRequest) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SubjectRightsRequest) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SubjectRightsRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *SubjectRightsRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *SubjectRightsRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *SubjectRightsRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *SubjectRightsRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *SubjectRightsRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *SubjectRightsRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *SubjectRightsRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *SubjectRightsRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDataSubject

`func (o *SubjectRightsRequest) GetDataSubject() AnyOfmicrosoftGraphDataSubject`

GetDataSubject returns the DataSubject field if non-nil, zero value otherwise.

### GetDataSubjectOk

`func (o *SubjectRightsRequest) GetDataSubjectOk() (*AnyOfmicrosoftGraphDataSubject, bool)`

GetDataSubjectOk returns a tuple with the DataSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSubject

`func (o *SubjectRightsRequest) SetDataSubject(v AnyOfmicrosoftGraphDataSubject)`

SetDataSubject sets DataSubject field to given value.

### HasDataSubject

`func (o *SubjectRightsRequest) HasDataSubject() bool`

HasDataSubject returns a boolean if a field has been set.

### SetDataSubjectNil

`func (o *SubjectRightsRequest) SetDataSubjectNil(b bool)`

 SetDataSubjectNil sets the value for DataSubject to be an explicit nil

### UnsetDataSubject
`func (o *SubjectRightsRequest) UnsetDataSubject()`

UnsetDataSubject ensures that no value is present for DataSubject, not even an explicit nil
### GetDataSubjectType

`func (o *SubjectRightsRequest) GetDataSubjectType() AnyOfmicrosoftGraphDataSubjectType`

GetDataSubjectType returns the DataSubjectType field if non-nil, zero value otherwise.

### GetDataSubjectTypeOk

`func (o *SubjectRightsRequest) GetDataSubjectTypeOk() (*AnyOfmicrosoftGraphDataSubjectType, bool)`

GetDataSubjectTypeOk returns a tuple with the DataSubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSubjectType

`func (o *SubjectRightsRequest) SetDataSubjectType(v AnyOfmicrosoftGraphDataSubjectType)`

SetDataSubjectType sets DataSubjectType field to given value.

### HasDataSubjectType

`func (o *SubjectRightsRequest) HasDataSubjectType() bool`

HasDataSubjectType returns a boolean if a field has been set.

### SetDataSubjectTypeNil

`func (o *SubjectRightsRequest) SetDataSubjectTypeNil(b bool)`

 SetDataSubjectTypeNil sets the value for DataSubjectType to be an explicit nil

### UnsetDataSubjectType
`func (o *SubjectRightsRequest) UnsetDataSubjectType()`

UnsetDataSubjectType ensures that no value is present for DataSubjectType, not even an explicit nil
### GetDescription

`func (o *SubjectRightsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubjectRightsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubjectRightsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubjectRightsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubjectRightsRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubjectRightsRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *SubjectRightsRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubjectRightsRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubjectRightsRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubjectRightsRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *SubjectRightsRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *SubjectRightsRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetHistory

`func (o *SubjectRightsRequest) GetHistory() []*AnyOfmicrosoftGraphSubjectRightsRequestHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *SubjectRightsRequest) GetHistoryOk() (*[]*AnyOfmicrosoftGraphSubjectRightsRequestHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *SubjectRightsRequest) SetHistory(v []*AnyOfmicrosoftGraphSubjectRightsRequestHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *SubjectRightsRequest) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetInsight

`func (o *SubjectRightsRequest) GetInsight() AnyOfmicrosoftGraphSubjectRightsRequestDetail`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *SubjectRightsRequest) GetInsightOk() (*AnyOfmicrosoftGraphSubjectRightsRequestDetail, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *SubjectRightsRequest) SetInsight(v AnyOfmicrosoftGraphSubjectRightsRequestDetail)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *SubjectRightsRequest) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### SetInsightNil

`func (o *SubjectRightsRequest) SetInsightNil(b bool)`

 SetInsightNil sets the value for Insight to be an explicit nil

### UnsetInsight
`func (o *SubjectRightsRequest) UnsetInsight()`

UnsetInsight ensures that no value is present for Insight, not even an explicit nil
### GetInternalDueDateTime

`func (o *SubjectRightsRequest) GetInternalDueDateTime() time.Time`

GetInternalDueDateTime returns the InternalDueDateTime field if non-nil, zero value otherwise.

### GetInternalDueDateTimeOk

`func (o *SubjectRightsRequest) GetInternalDueDateTimeOk() (*time.Time, bool)`

GetInternalDueDateTimeOk returns a tuple with the InternalDueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDueDateTime

`func (o *SubjectRightsRequest) SetInternalDueDateTime(v time.Time)`

SetInternalDueDateTime sets InternalDueDateTime field to given value.

### HasInternalDueDateTime

`func (o *SubjectRightsRequest) HasInternalDueDateTime() bool`

HasInternalDueDateTime returns a boolean if a field has been set.

### SetInternalDueDateTimeNil

`func (o *SubjectRightsRequest) SetInternalDueDateTimeNil(b bool)`

 SetInternalDueDateTimeNil sets the value for InternalDueDateTime to be an explicit nil

### UnsetInternalDueDateTime
`func (o *SubjectRightsRequest) UnsetInternalDueDateTime()`

UnsetInternalDueDateTime ensures that no value is present for InternalDueDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *SubjectRightsRequest) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *SubjectRightsRequest) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *SubjectRightsRequest) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *SubjectRightsRequest) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *SubjectRightsRequest) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *SubjectRightsRequest) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *SubjectRightsRequest) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *SubjectRightsRequest) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *SubjectRightsRequest) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *SubjectRightsRequest) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *SubjectRightsRequest) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *SubjectRightsRequest) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetRegulations

`func (o *SubjectRightsRequest) GetRegulations() []*string`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *SubjectRightsRequest) GetRegulationsOk() (*[]*string, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *SubjectRightsRequest) SetRegulations(v []*string)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *SubjectRightsRequest) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.

### GetStages

`func (o *SubjectRightsRequest) GetStages() []*AnyOfmicrosoftGraphSubjectRightsRequestStageDetail`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *SubjectRightsRequest) GetStagesOk() (*[]*AnyOfmicrosoftGraphSubjectRightsRequestStageDetail, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *SubjectRightsRequest) SetStages(v []*AnyOfmicrosoftGraphSubjectRightsRequestStageDetail)`

SetStages sets Stages field to given value.

### HasStages

`func (o *SubjectRightsRequest) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetStatus

`func (o *SubjectRightsRequest) GetStatus() AnyOfmicrosoftGraphSubjectRightsRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubjectRightsRequest) GetStatusOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubjectRightsRequest) SetStatus(v AnyOfmicrosoftGraphSubjectRightsRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubjectRightsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SubjectRightsRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SubjectRightsRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *SubjectRightsRequest) GetType() AnyOfmicrosoftGraphSubjectRightsRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubjectRightsRequest) GetTypeOk() (*AnyOfmicrosoftGraphSubjectRightsRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubjectRightsRequest) SetType(v AnyOfmicrosoftGraphSubjectRightsRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *SubjectRightsRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SubjectRightsRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SubjectRightsRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetNotes

`func (o *SubjectRightsRequest) GetNotes() []MicrosoftGraphAuthoredNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SubjectRightsRequest) GetNotesOk() (*[]MicrosoftGraphAuthoredNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SubjectRightsRequest) SetNotes(v []MicrosoftGraphAuthoredNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SubjectRightsRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTeam

`func (o *SubjectRightsRequest) GetTeam() AnyOfmicrosoftGraphTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *SubjectRightsRequest) GetTeamOk() (*AnyOfmicrosoftGraphTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *SubjectRightsRequest) SetTeam(v AnyOfmicrosoftGraphTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *SubjectRightsRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *SubjectRightsRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *SubjectRightsRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


