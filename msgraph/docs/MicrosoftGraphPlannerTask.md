# MicrosoftGraphPlannerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphPlannerTask

`func NewMicrosoftGraphPlannerTask() *MicrosoftGraphPlannerTask`

NewMicrosoftGraphPlannerTask instantiates a new MicrosoftGraphPlannerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerTaskWithDefaults

`func NewMicrosoftGraphPlannerTaskWithDefaults() *MicrosoftGraphPlannerTask`

NewMicrosoftGraphPlannerTaskWithDefaults instantiates a new MicrosoftGraphPlannerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlannerTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlannerTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlannerTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActiveChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) GetActiveChecklistItemCount() int32`

GetActiveChecklistItemCount returns the ActiveChecklistItemCount field if non-nil, zero value otherwise.

### GetActiveChecklistItemCountOk

`func (o *MicrosoftGraphPlannerTask) GetActiveChecklistItemCountOk() (*int32, bool)`

GetActiveChecklistItemCountOk returns a tuple with the ActiveChecklistItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) SetActiveChecklistItemCount(v int32)`

SetActiveChecklistItemCount sets ActiveChecklistItemCount field to given value.

### HasActiveChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) HasActiveChecklistItemCount() bool`

HasActiveChecklistItemCount returns a boolean if a field has been set.

### SetActiveChecklistItemCountNil

`func (o *MicrosoftGraphPlannerTask) SetActiveChecklistItemCountNil(b bool)`

 SetActiveChecklistItemCountNil sets the value for ActiveChecklistItemCount to be an explicit nil

### UnsetActiveChecklistItemCount
`func (o *MicrosoftGraphPlannerTask) UnsetActiveChecklistItemCount()`

UnsetActiveChecklistItemCount ensures that no value is present for ActiveChecklistItemCount, not even an explicit nil
### GetAppliedCategories

`func (o *MicrosoftGraphPlannerTask) GetAppliedCategories() AnyOfobject`

GetAppliedCategories returns the AppliedCategories field if non-nil, zero value otherwise.

### GetAppliedCategoriesOk

`func (o *MicrosoftGraphPlannerTask) GetAppliedCategoriesOk() (*AnyOfobject, bool)`

GetAppliedCategoriesOk returns a tuple with the AppliedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCategories

`func (o *MicrosoftGraphPlannerTask) SetAppliedCategories(v AnyOfobject)`

SetAppliedCategories sets AppliedCategories field to given value.

### HasAppliedCategories

`func (o *MicrosoftGraphPlannerTask) HasAppliedCategories() bool`

HasAppliedCategories returns a boolean if a field has been set.

### SetAppliedCategoriesNil

`func (o *MicrosoftGraphPlannerTask) SetAppliedCategoriesNil(b bool)`

 SetAppliedCategoriesNil sets the value for AppliedCategories to be an explicit nil

### UnsetAppliedCategories
`func (o *MicrosoftGraphPlannerTask) UnsetAppliedCategories()`

UnsetAppliedCategories ensures that no value is present for AppliedCategories, not even an explicit nil
### GetAssigneePriority

`func (o *MicrosoftGraphPlannerTask) GetAssigneePriority() string`

GetAssigneePriority returns the AssigneePriority field if non-nil, zero value otherwise.

### GetAssigneePriorityOk

`func (o *MicrosoftGraphPlannerTask) GetAssigneePriorityOk() (*string, bool)`

GetAssigneePriorityOk returns a tuple with the AssigneePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneePriority

`func (o *MicrosoftGraphPlannerTask) SetAssigneePriority(v string)`

SetAssigneePriority sets AssigneePriority field to given value.

### HasAssigneePriority

`func (o *MicrosoftGraphPlannerTask) HasAssigneePriority() bool`

HasAssigneePriority returns a boolean if a field has been set.

### SetAssigneePriorityNil

`func (o *MicrosoftGraphPlannerTask) SetAssigneePriorityNil(b bool)`

 SetAssigneePriorityNil sets the value for AssigneePriority to be an explicit nil

### UnsetAssigneePriority
`func (o *MicrosoftGraphPlannerTask) UnsetAssigneePriority()`

UnsetAssigneePriority ensures that no value is present for AssigneePriority, not even an explicit nil
### GetAssignments

`func (o *MicrosoftGraphPlannerTask) GetAssignments() AnyOfobject`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphPlannerTask) GetAssignmentsOk() (*AnyOfobject, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphPlannerTask) SetAssignments(v AnyOfobject)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphPlannerTask) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignmentsNil

`func (o *MicrosoftGraphPlannerTask) SetAssignmentsNil(b bool)`

 SetAssignmentsNil sets the value for Assignments to be an explicit nil

### UnsetAssignments
`func (o *MicrosoftGraphPlannerTask) UnsetAssignments()`

UnsetAssignments ensures that no value is present for Assignments, not even an explicit nil
### GetBucketId

`func (o *MicrosoftGraphPlannerTask) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *MicrosoftGraphPlannerTask) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *MicrosoftGraphPlannerTask) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *MicrosoftGraphPlannerTask) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### SetBucketIdNil

`func (o *MicrosoftGraphPlannerTask) SetBucketIdNil(b bool)`

 SetBucketIdNil sets the value for BucketId to be an explicit nil

### UnsetBucketId
`func (o *MicrosoftGraphPlannerTask) UnsetBucketId()`

UnsetBucketId ensures that no value is present for BucketId, not even an explicit nil
### GetChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) GetChecklistItemCount() int32`

GetChecklistItemCount returns the ChecklistItemCount field if non-nil, zero value otherwise.

### GetChecklistItemCountOk

`func (o *MicrosoftGraphPlannerTask) GetChecklistItemCountOk() (*int32, bool)`

GetChecklistItemCountOk returns a tuple with the ChecklistItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) SetChecklistItemCount(v int32)`

SetChecklistItemCount sets ChecklistItemCount field to given value.

### HasChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) HasChecklistItemCount() bool`

HasChecklistItemCount returns a boolean if a field has been set.

### SetChecklistItemCountNil

`func (o *MicrosoftGraphPlannerTask) SetChecklistItemCountNil(b bool)`

 SetChecklistItemCountNil sets the value for ChecklistItemCount to be an explicit nil

### UnsetChecklistItemCount
`func (o *MicrosoftGraphPlannerTask) UnsetChecklistItemCount()`

UnsetChecklistItemCount ensures that no value is present for ChecklistItemCount, not even an explicit nil
### GetCompletedBy

`func (o *MicrosoftGraphPlannerTask) GetCompletedBy() AnyOfmicrosoftGraphIdentitySet`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *MicrosoftGraphPlannerTask) GetCompletedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedBy

`func (o *MicrosoftGraphPlannerTask) SetCompletedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCompletedBy sets CompletedBy field to given value.

### HasCompletedBy

`func (o *MicrosoftGraphPlannerTask) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### SetCompletedByNil

`func (o *MicrosoftGraphPlannerTask) SetCompletedByNil(b bool)`

 SetCompletedByNil sets the value for CompletedBy to be an explicit nil

### UnsetCompletedBy
`func (o *MicrosoftGraphPlannerTask) UnsetCompletedBy()`

UnsetCompletedBy ensures that no value is present for CompletedBy, not even an explicit nil
### GetCompletedDateTime

`func (o *MicrosoftGraphPlannerTask) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetCompletedDateTimeOk() (*time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *MicrosoftGraphPlannerTask) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *MicrosoftGraphPlannerTask) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *MicrosoftGraphPlannerTask) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *MicrosoftGraphPlannerTask) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetConversationThreadId

`func (o *MicrosoftGraphPlannerTask) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *MicrosoftGraphPlannerTask) GetConversationThreadIdOk() (*string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationThreadId

`func (o *MicrosoftGraphPlannerTask) SetConversationThreadId(v string)`

SetConversationThreadId sets ConversationThreadId field to given value.

### HasConversationThreadId

`func (o *MicrosoftGraphPlannerTask) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadIdNil

`func (o *MicrosoftGraphPlannerTask) SetConversationThreadIdNil(b bool)`

 SetConversationThreadIdNil sets the value for ConversationThreadId to be an explicit nil

### UnsetConversationThreadId
`func (o *MicrosoftGraphPlannerTask) UnsetConversationThreadId()`

UnsetConversationThreadId ensures that no value is present for ConversationThreadId, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphPlannerTask) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphPlannerTask) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphPlannerTask) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphPlannerTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphPlannerTask) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphPlannerTask) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphPlannerTask) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPlannerTask) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphPlannerTask) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphPlannerTask) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphPlannerTask) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDueDateTime

`func (o *MicrosoftGraphPlannerTask) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *MicrosoftGraphPlannerTask) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *MicrosoftGraphPlannerTask) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *MicrosoftGraphPlannerTask) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *MicrosoftGraphPlannerTask) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetHasDescription

`func (o *MicrosoftGraphPlannerTask) GetHasDescription() bool`

GetHasDescription returns the HasDescription field if non-nil, zero value otherwise.

### GetHasDescriptionOk

`func (o *MicrosoftGraphPlannerTask) GetHasDescriptionOk() (*bool, bool)`

GetHasDescriptionOk returns a tuple with the HasDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescription

`func (o *MicrosoftGraphPlannerTask) SetHasDescription(v bool)`

SetHasDescription sets HasDescription field to given value.

### HasHasDescription

`func (o *MicrosoftGraphPlannerTask) HasHasDescription() bool`

HasHasDescription returns a boolean if a field has been set.

### SetHasDescriptionNil

`func (o *MicrosoftGraphPlannerTask) SetHasDescriptionNil(b bool)`

 SetHasDescriptionNil sets the value for HasDescription to be an explicit nil

### UnsetHasDescription
`func (o *MicrosoftGraphPlannerTask) UnsetHasDescription()`

UnsetHasDescription ensures that no value is present for HasDescription, not even an explicit nil
### GetOrderHint

`func (o *MicrosoftGraphPlannerTask) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *MicrosoftGraphPlannerTask) GetOrderHintOk() (*string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHint

`func (o *MicrosoftGraphPlannerTask) SetOrderHint(v string)`

SetOrderHint sets OrderHint field to given value.

### HasOrderHint

`func (o *MicrosoftGraphPlannerTask) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHintNil

`func (o *MicrosoftGraphPlannerTask) SetOrderHintNil(b bool)`

 SetOrderHintNil sets the value for OrderHint to be an explicit nil

### UnsetOrderHint
`func (o *MicrosoftGraphPlannerTask) UnsetOrderHint()`

UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil
### GetPercentComplete

`func (o *MicrosoftGraphPlannerTask) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MicrosoftGraphPlannerTask) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *MicrosoftGraphPlannerTask) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *MicrosoftGraphPlannerTask) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *MicrosoftGraphPlannerTask) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *MicrosoftGraphPlannerTask) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetPlanId

`func (o *MicrosoftGraphPlannerTask) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *MicrosoftGraphPlannerTask) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *MicrosoftGraphPlannerTask) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *MicrosoftGraphPlannerTask) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *MicrosoftGraphPlannerTask) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *MicrosoftGraphPlannerTask) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPreviewType

`func (o *MicrosoftGraphPlannerTask) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *MicrosoftGraphPlannerTask) GetPreviewTypeOk() (*AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewType

`func (o *MicrosoftGraphPlannerTask) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType sets PreviewType field to given value.

### HasPreviewType

`func (o *MicrosoftGraphPlannerTask) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewTypeNil

`func (o *MicrosoftGraphPlannerTask) SetPreviewTypeNil(b bool)`

 SetPreviewTypeNil sets the value for PreviewType to be an explicit nil

### UnsetPreviewType
`func (o *MicrosoftGraphPlannerTask) UnsetPreviewType()`

UnsetPreviewType ensures that no value is present for PreviewType, not even an explicit nil
### GetReferenceCount

`func (o *MicrosoftGraphPlannerTask) GetReferenceCount() int32`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *MicrosoftGraphPlannerTask) GetReferenceCountOk() (*int32, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *MicrosoftGraphPlannerTask) SetReferenceCount(v int32)`

SetReferenceCount sets ReferenceCount field to given value.

### HasReferenceCount

`func (o *MicrosoftGraphPlannerTask) HasReferenceCount() bool`

HasReferenceCount returns a boolean if a field has been set.

### SetReferenceCountNil

`func (o *MicrosoftGraphPlannerTask) SetReferenceCountNil(b bool)`

 SetReferenceCountNil sets the value for ReferenceCount to be an explicit nil

### UnsetReferenceCount
`func (o *MicrosoftGraphPlannerTask) UnsetReferenceCount()`

UnsetReferenceCount ensures that no value is present for ReferenceCount, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphPlannerTask) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphPlannerTask) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphPlannerTask) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphPlannerTask) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphPlannerTask) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphPlannerTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphPlannerTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphPlannerTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphPlannerTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssignedToTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) GetAssignedToTaskBoardFormat() AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat`

GetAssignedToTaskBoardFormat returns the AssignedToTaskBoardFormat field if non-nil, zero value otherwise.

### GetAssignedToTaskBoardFormatOk

`func (o *MicrosoftGraphPlannerTask) GetAssignedToTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat, bool)`

GetAssignedToTaskBoardFormatOk returns a tuple with the AssignedToTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) SetAssignedToTaskBoardFormat(v AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat)`

SetAssignedToTaskBoardFormat sets AssignedToTaskBoardFormat field to given value.

### HasAssignedToTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) HasAssignedToTaskBoardFormat() bool`

HasAssignedToTaskBoardFormat returns a boolean if a field has been set.

### SetAssignedToTaskBoardFormatNil

`func (o *MicrosoftGraphPlannerTask) SetAssignedToTaskBoardFormatNil(b bool)`

 SetAssignedToTaskBoardFormatNil sets the value for AssignedToTaskBoardFormat to be an explicit nil

### UnsetAssignedToTaskBoardFormat
`func (o *MicrosoftGraphPlannerTask) UnsetAssignedToTaskBoardFormat()`

UnsetAssignedToTaskBoardFormat ensures that no value is present for AssignedToTaskBoardFormat, not even an explicit nil
### GetBucketTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) GetBucketTaskBoardFormat() AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat`

GetBucketTaskBoardFormat returns the BucketTaskBoardFormat field if non-nil, zero value otherwise.

### GetBucketTaskBoardFormatOk

`func (o *MicrosoftGraphPlannerTask) GetBucketTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat, bool)`

GetBucketTaskBoardFormatOk returns a tuple with the BucketTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) SetBucketTaskBoardFormat(v AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat)`

SetBucketTaskBoardFormat sets BucketTaskBoardFormat field to given value.

### HasBucketTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) HasBucketTaskBoardFormat() bool`

HasBucketTaskBoardFormat returns a boolean if a field has been set.

### SetBucketTaskBoardFormatNil

`func (o *MicrosoftGraphPlannerTask) SetBucketTaskBoardFormatNil(b bool)`

 SetBucketTaskBoardFormatNil sets the value for BucketTaskBoardFormat to be an explicit nil

### UnsetBucketTaskBoardFormat
`func (o *MicrosoftGraphPlannerTask) UnsetBucketTaskBoardFormat()`

UnsetBucketTaskBoardFormat ensures that no value is present for BucketTaskBoardFormat, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphPlannerTask) GetDetails() AnyOfmicrosoftGraphPlannerTaskDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPlannerTask) GetDetailsOk() (*AnyOfmicrosoftGraphPlannerTaskDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphPlannerTask) SetDetails(v AnyOfmicrosoftGraphPlannerTaskDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphPlannerTask) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *MicrosoftGraphPlannerTask) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *MicrosoftGraphPlannerTask) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetProgressTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) GetProgressTaskBoardFormat() AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat`

GetProgressTaskBoardFormat returns the ProgressTaskBoardFormat field if non-nil, zero value otherwise.

### GetProgressTaskBoardFormatOk

`func (o *MicrosoftGraphPlannerTask) GetProgressTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat, bool)`

GetProgressTaskBoardFormatOk returns a tuple with the ProgressTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) SetProgressTaskBoardFormat(v AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat)`

SetProgressTaskBoardFormat sets ProgressTaskBoardFormat field to given value.

### HasProgressTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) HasProgressTaskBoardFormat() bool`

HasProgressTaskBoardFormat returns a boolean if a field has been set.

### SetProgressTaskBoardFormatNil

`func (o *MicrosoftGraphPlannerTask) SetProgressTaskBoardFormatNil(b bool)`

 SetProgressTaskBoardFormatNil sets the value for ProgressTaskBoardFormat to be an explicit nil

### UnsetProgressTaskBoardFormat
`func (o *MicrosoftGraphPlannerTask) UnsetProgressTaskBoardFormat()`

UnsetProgressTaskBoardFormat ensures that no value is present for ProgressTaskBoardFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


