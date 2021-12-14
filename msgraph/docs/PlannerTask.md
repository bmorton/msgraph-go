# PlannerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveChecklistItemCount** | Pointer to **NullableInt32** | Number of checklist items with value set to false, representing incomplete items. | [optional] 
**AppliedCategories** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The categories to which the task has been applied. See applied Categories for possible values. | [optional] 
**AssigneePriority** | Pointer to **NullableString** | Hint used to order items of this type in a list view. The format is defined as outlined here. | [optional] 
**Assignments** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The set of assignees the task is assigned to. | [optional] 
**BucketId** | Pointer to **NullableString** | Bucket ID to which the task belongs. The bucket needs to be in the plan that the task is in. It is 28 characters long and case-sensitive. Format validation is done on the service. | [optional] 
**ChecklistItemCount** | Pointer to **NullableInt32** | Number of checklist items that are present on the task. | [optional] 
**CompletedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user that completed the task. | [optional] 
**CompletedDateTime** | Pointer to **NullableTime** | Read-only. Date and time at which the &#39;percentComplete&#39; of the task is set to &#39;100&#39;. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**ConversationThreadId** | Pointer to **NullableString** | Thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user that created the task. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Read-only. Date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**DueDateTime** | Pointer to **NullableTime** | Date and time at which the task is due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**HasDescription** | Pointer to **NullableBool** | Read-only. Value is true if the details object of the task has a non-empty description and false otherwise. | [optional] 
**OrderHint** | Pointer to **NullableString** | Hint used to order items of this type in a list view. The format is defined as outlined here. | [optional] 
**PercentComplete** | Pointer to **NullableInt32** | Percentage of task completion. When set to 100, the task is considered completed. | [optional] 
**PlanId** | Pointer to **NullableString** | Plan ID to which the task belongs. | [optional] 
**PreviewType** | Pointer to [**AnyOfmicrosoftGraphPlannerPreviewType**](anyOf&lt;microsoft.graph.plannerPreviewType&gt;.md) | This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. | [optional] 
**ReferenceCount** | Pointer to **NullableInt32** | Number of external references that exist on the task. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Title** | Pointer to **string** | Title of the task. | [optional] 
**AssignedToTaskBoardFormat** | Pointer to [**AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](anyOf&lt;microsoft.graph.plannerAssignedToTaskBoardTaskFormat&gt;.md) | Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo. | [optional] 
**BucketTaskBoardFormat** | Pointer to [**AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat**](anyOf&lt;microsoft.graph.plannerBucketTaskBoardTaskFormat&gt;.md) | Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket. | [optional] 
**Details** | Pointer to [**AnyOfmicrosoftGraphPlannerTaskDetails**](anyOf&lt;microsoft.graph.plannerTaskDetails&gt;.md) | Read-only. Nullable. Additional details about the task. | [optional] 
**ProgressTaskBoardFormat** | Pointer to [**AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat**](anyOf&lt;microsoft.graph.plannerProgressTaskBoardTaskFormat&gt;.md) | Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress. | [optional] 

## Methods

### NewPlannerTask

`func NewPlannerTask() *PlannerTask`

NewPlannerTask instantiates a new PlannerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerTaskWithDefaults

`func NewPlannerTaskWithDefaults() *PlannerTask`

NewPlannerTaskWithDefaults instantiates a new PlannerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveChecklistItemCount

`func (o *PlannerTask) GetActiveChecklistItemCount() int32`

GetActiveChecklistItemCount returns the ActiveChecklistItemCount field if non-nil, zero value otherwise.

### GetActiveChecklistItemCountOk

`func (o *PlannerTask) GetActiveChecklistItemCountOk() (*int32, bool)`

GetActiveChecklistItemCountOk returns a tuple with the ActiveChecklistItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveChecklistItemCount

`func (o *PlannerTask) SetActiveChecklistItemCount(v int32)`

SetActiveChecklistItemCount sets ActiveChecklistItemCount field to given value.

### HasActiveChecklistItemCount

`func (o *PlannerTask) HasActiveChecklistItemCount() bool`

HasActiveChecklistItemCount returns a boolean if a field has been set.

### SetActiveChecklistItemCountNil

`func (o *PlannerTask) SetActiveChecklistItemCountNil(b bool)`

 SetActiveChecklistItemCountNil sets the value for ActiveChecklistItemCount to be an explicit nil

### UnsetActiveChecklistItemCount
`func (o *PlannerTask) UnsetActiveChecklistItemCount()`

UnsetActiveChecklistItemCount ensures that no value is present for ActiveChecklistItemCount, not even an explicit nil
### GetAppliedCategories

`func (o *PlannerTask) GetAppliedCategories() AnyOfobject`

GetAppliedCategories returns the AppliedCategories field if non-nil, zero value otherwise.

### GetAppliedCategoriesOk

`func (o *PlannerTask) GetAppliedCategoriesOk() (*AnyOfobject, bool)`

GetAppliedCategoriesOk returns a tuple with the AppliedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCategories

`func (o *PlannerTask) SetAppliedCategories(v AnyOfobject)`

SetAppliedCategories sets AppliedCategories field to given value.

### HasAppliedCategories

`func (o *PlannerTask) HasAppliedCategories() bool`

HasAppliedCategories returns a boolean if a field has been set.

### SetAppliedCategoriesNil

`func (o *PlannerTask) SetAppliedCategoriesNil(b bool)`

 SetAppliedCategoriesNil sets the value for AppliedCategories to be an explicit nil

### UnsetAppliedCategories
`func (o *PlannerTask) UnsetAppliedCategories()`

UnsetAppliedCategories ensures that no value is present for AppliedCategories, not even an explicit nil
### GetAssigneePriority

`func (o *PlannerTask) GetAssigneePriority() string`

GetAssigneePriority returns the AssigneePriority field if non-nil, zero value otherwise.

### GetAssigneePriorityOk

`func (o *PlannerTask) GetAssigneePriorityOk() (*string, bool)`

GetAssigneePriorityOk returns a tuple with the AssigneePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneePriority

`func (o *PlannerTask) SetAssigneePriority(v string)`

SetAssigneePriority sets AssigneePriority field to given value.

### HasAssigneePriority

`func (o *PlannerTask) HasAssigneePriority() bool`

HasAssigneePriority returns a boolean if a field has been set.

### SetAssigneePriorityNil

`func (o *PlannerTask) SetAssigneePriorityNil(b bool)`

 SetAssigneePriorityNil sets the value for AssigneePriority to be an explicit nil

### UnsetAssigneePriority
`func (o *PlannerTask) UnsetAssigneePriority()`

UnsetAssigneePriority ensures that no value is present for AssigneePriority, not even an explicit nil
### GetAssignments

`func (o *PlannerTask) GetAssignments() AnyOfobject`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *PlannerTask) GetAssignmentsOk() (*AnyOfobject, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *PlannerTask) SetAssignments(v AnyOfobject)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *PlannerTask) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignmentsNil

`func (o *PlannerTask) SetAssignmentsNil(b bool)`

 SetAssignmentsNil sets the value for Assignments to be an explicit nil

### UnsetAssignments
`func (o *PlannerTask) UnsetAssignments()`

UnsetAssignments ensures that no value is present for Assignments, not even an explicit nil
### GetBucketId

`func (o *PlannerTask) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *PlannerTask) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *PlannerTask) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *PlannerTask) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### SetBucketIdNil

`func (o *PlannerTask) SetBucketIdNil(b bool)`

 SetBucketIdNil sets the value for BucketId to be an explicit nil

### UnsetBucketId
`func (o *PlannerTask) UnsetBucketId()`

UnsetBucketId ensures that no value is present for BucketId, not even an explicit nil
### GetChecklistItemCount

`func (o *PlannerTask) GetChecklistItemCount() int32`

GetChecklistItemCount returns the ChecklistItemCount field if non-nil, zero value otherwise.

### GetChecklistItemCountOk

`func (o *PlannerTask) GetChecklistItemCountOk() (*int32, bool)`

GetChecklistItemCountOk returns a tuple with the ChecklistItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistItemCount

`func (o *PlannerTask) SetChecklistItemCount(v int32)`

SetChecklistItemCount sets ChecklistItemCount field to given value.

### HasChecklistItemCount

`func (o *PlannerTask) HasChecklistItemCount() bool`

HasChecklistItemCount returns a boolean if a field has been set.

### SetChecklistItemCountNil

`func (o *PlannerTask) SetChecklistItemCountNil(b bool)`

 SetChecklistItemCountNil sets the value for ChecklistItemCount to be an explicit nil

### UnsetChecklistItemCount
`func (o *PlannerTask) UnsetChecklistItemCount()`

UnsetChecklistItemCount ensures that no value is present for ChecklistItemCount, not even an explicit nil
### GetCompletedBy

`func (o *PlannerTask) GetCompletedBy() AnyOfmicrosoftGraphIdentitySet`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *PlannerTask) GetCompletedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *PlannerTask) SetCompletedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *PlannerTask) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### SetCompletedByNil

`func (o *PlannerTask) SetCompletedByNil(b bool)`

 SetCompletedByNil sets the value for CompletedBy to be an explicit nil

### UnsetCompletedBy
`func (o *PlannerTask) UnsetCompletedBy()`

UnsetCompletedBy ensures that no value is present for CompletedBy, not even an explicit nil
### GetCompletedDateTime

`func (o *PlannerTask) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *PlannerTask) GetCompletedDateTimeOk() (*time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *PlannerTask) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *PlannerTask) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *PlannerTask) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *PlannerTask) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetConversationThreadId

`func (o *PlannerTask) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *PlannerTask) GetConversationThreadIdOk() (*string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationThreadId

`func (o *PlannerTask) SetConversationThreadId(v string)`

SetConversationThreadId sets ConversationThreadId field to given value.

### HasConversationThreadId

`func (o *PlannerTask) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadIdNil

`func (o *PlannerTask) SetConversationThreadIdNil(b bool)`

 SetConversationThreadIdNil sets the value for ConversationThreadId to be an explicit nil

### UnsetConversationThreadId
`func (o *PlannerTask) UnsetConversationThreadId()`

UnsetConversationThreadId ensures that no value is present for ConversationThreadId, not even an explicit nil
### GetCreatedBy

`func (o *PlannerTask) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PlannerTask) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PlannerTask) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PlannerTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PlannerTask) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PlannerTask) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *PlannerTask) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PlannerTask) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *PlannerTask) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *PlannerTask) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *PlannerTask) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *PlannerTask) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDueDateTime

`func (o *PlannerTask) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *PlannerTask) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *PlannerTask) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *PlannerTask) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *PlannerTask) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *PlannerTask) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetHasDescription

`func (o *PlannerTask) GetHasDescription() bool`

GetHasDescription returns the HasDescription field if non-nil, zero value otherwise.

### GetHasDescriptionOk

`func (o *PlannerTask) GetHasDescriptionOk() (*bool, bool)`

GetHasDescriptionOk returns a tuple with the HasDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescription

`func (o *PlannerTask) SetHasDescription(v bool)`

SetHasDescription sets HasDescription field to given value.

### HasHasDescription

`func (o *PlannerTask) HasHasDescription() bool`

HasHasDescription returns a boolean if a field has been set.

### SetHasDescriptionNil

`func (o *PlannerTask) SetHasDescriptionNil(b bool)`

 SetHasDescriptionNil sets the value for HasDescription to be an explicit nil

### UnsetHasDescription
`func (o *PlannerTask) UnsetHasDescription()`

UnsetHasDescription ensures that no value is present for HasDescription, not even an explicit nil
### GetOrderHint

`func (o *PlannerTask) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *PlannerTask) GetOrderHintOk() (*string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHint

`func (o *PlannerTask) SetOrderHint(v string)`

SetOrderHint sets OrderHint field to given value.

### HasOrderHint

`func (o *PlannerTask) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHintNil

`func (o *PlannerTask) SetOrderHintNil(b bool)`

 SetOrderHintNil sets the value for OrderHint to be an explicit nil

### UnsetOrderHint
`func (o *PlannerTask) UnsetOrderHint()`

UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil
### GetPercentComplete

`func (o *PlannerTask) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *PlannerTask) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *PlannerTask) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *PlannerTask) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *PlannerTask) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *PlannerTask) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetPlanId

`func (o *PlannerTask) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlannerTask) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlannerTask) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlannerTask) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *PlannerTask) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *PlannerTask) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPreviewType

`func (o *PlannerTask) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *PlannerTask) GetPreviewTypeOk() (*AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewType

`func (o *PlannerTask) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType sets PreviewType field to given value.

### HasPreviewType

`func (o *PlannerTask) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewTypeNil

`func (o *PlannerTask) SetPreviewTypeNil(b bool)`

 SetPreviewTypeNil sets the value for PreviewType to be an explicit nil

### UnsetPreviewType
`func (o *PlannerTask) UnsetPreviewType()`

UnsetPreviewType ensures that no value is present for PreviewType, not even an explicit nil
### GetReferenceCount

`func (o *PlannerTask) GetReferenceCount() int32`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *PlannerTask) GetReferenceCountOk() (*int32, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *PlannerTask) SetReferenceCount(v int32)`

SetReferenceCount sets ReferenceCount field to given value.

### HasReferenceCount

`func (o *PlannerTask) HasReferenceCount() bool`

HasReferenceCount returns a boolean if a field has been set.

### SetReferenceCountNil

`func (o *PlannerTask) SetReferenceCountNil(b bool)`

 SetReferenceCountNil sets the value for ReferenceCount to be an explicit nil

### UnsetReferenceCount
`func (o *PlannerTask) UnsetReferenceCount()`

UnsetReferenceCount ensures that no value is present for ReferenceCount, not even an explicit nil
### GetStartDateTime

`func (o *PlannerTask) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *PlannerTask) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *PlannerTask) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *PlannerTask) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *PlannerTask) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *PlannerTask) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTitle

`func (o *PlannerTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlannerTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlannerTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlannerTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssignedToTaskBoardFormat

`func (o *PlannerTask) GetAssignedToTaskBoardFormat() AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat`

GetAssignedToTaskBoardFormat returns the AssignedToTaskBoardFormat field if non-nil, zero value otherwise.

### GetAssignedToTaskBoardFormatOk

`func (o *PlannerTask) GetAssignedToTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat, bool)`

GetAssignedToTaskBoardFormatOk returns a tuple with the AssignedToTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToTaskBoardFormat

`func (o *PlannerTask) SetAssignedToTaskBoardFormat(v AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat)`

SetAssignedToTaskBoardFormat sets AssignedToTaskBoardFormat field to given value.

### HasAssignedToTaskBoardFormat

`func (o *PlannerTask) HasAssignedToTaskBoardFormat() bool`

HasAssignedToTaskBoardFormat returns a boolean if a field has been set.

### SetAssignedToTaskBoardFormatNil

`func (o *PlannerTask) SetAssignedToTaskBoardFormatNil(b bool)`

 SetAssignedToTaskBoardFormatNil sets the value for AssignedToTaskBoardFormat to be an explicit nil

### UnsetAssignedToTaskBoardFormat
`func (o *PlannerTask) UnsetAssignedToTaskBoardFormat()`

UnsetAssignedToTaskBoardFormat ensures that no value is present for AssignedToTaskBoardFormat, not even an explicit nil
### GetBucketTaskBoardFormat

`func (o *PlannerTask) GetBucketTaskBoardFormat() AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat`

GetBucketTaskBoardFormat returns the BucketTaskBoardFormat field if non-nil, zero value otherwise.

### GetBucketTaskBoardFormatOk

`func (o *PlannerTask) GetBucketTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat, bool)`

GetBucketTaskBoardFormatOk returns a tuple with the BucketTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketTaskBoardFormat

`func (o *PlannerTask) SetBucketTaskBoardFormat(v AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat)`

SetBucketTaskBoardFormat sets BucketTaskBoardFormat field to given value.

### HasBucketTaskBoardFormat

`func (o *PlannerTask) HasBucketTaskBoardFormat() bool`

HasBucketTaskBoardFormat returns a boolean if a field has been set.

### SetBucketTaskBoardFormatNil

`func (o *PlannerTask) SetBucketTaskBoardFormatNil(b bool)`

 SetBucketTaskBoardFormatNil sets the value for BucketTaskBoardFormat to be an explicit nil

### UnsetBucketTaskBoardFormat
`func (o *PlannerTask) UnsetBucketTaskBoardFormat()`

UnsetBucketTaskBoardFormat ensures that no value is present for BucketTaskBoardFormat, not even an explicit nil
### GetDetails

`func (o *PlannerTask) GetDetails() AnyOfmicrosoftGraphPlannerTaskDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PlannerTask) GetDetailsOk() (*AnyOfmicrosoftGraphPlannerTaskDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PlannerTask) SetDetails(v AnyOfmicrosoftGraphPlannerTaskDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PlannerTask) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PlannerTask) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PlannerTask) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetProgressTaskBoardFormat

`func (o *PlannerTask) GetProgressTaskBoardFormat() AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat`

GetProgressTaskBoardFormat returns the ProgressTaskBoardFormat field if non-nil, zero value otherwise.

### GetProgressTaskBoardFormatOk

`func (o *PlannerTask) GetProgressTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat, bool)`

GetProgressTaskBoardFormatOk returns a tuple with the ProgressTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskBoardFormat

`func (o *PlannerTask) SetProgressTaskBoardFormat(v AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat)`

SetProgressTaskBoardFormat sets ProgressTaskBoardFormat field to given value.

### HasProgressTaskBoardFormat

`func (o *PlannerTask) HasProgressTaskBoardFormat() bool`

HasProgressTaskBoardFormat returns a boolean if a field has been set.

### SetProgressTaskBoardFormatNil

`func (o *PlannerTask) SetProgressTaskBoardFormatNil(b bool)`

 SetProgressTaskBoardFormatNil sets the value for ProgressTaskBoardFormat to be an explicit nil

### UnsetProgressTaskBoardFormat
`func (o *PlannerTask) UnsetProgressTaskBoardFormat()`

UnsetProgressTaskBoardFormat ensures that no value is present for ProgressTaskBoardFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


