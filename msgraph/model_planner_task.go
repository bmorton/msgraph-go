/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// PlannerTask struct for PlannerTask
type PlannerTask struct {
	// Number of checklist items with value set to false, representing incomplete items.
	ActiveChecklistItemCount NullableInt32 `json:"activeChecklistItemCount,omitempty"`
	// The categories to which the task has been applied. See applied Categories for possible values.
	AppliedCategories AnyOfobject `json:"appliedCategories,omitempty"`
	// Hint used to order items of this type in a list view. The format is defined as outlined here.
	AssigneePriority NullableString `json:"assigneePriority,omitempty"`
	// The set of assignees the task is assigned to.
	Assignments AnyOfobject `json:"assignments,omitempty"`
	// Bucket ID to which the task belongs. The bucket needs to be in the plan that the task is in. It is 28 characters long and case-sensitive. Format validation is done on the service.
	BucketId NullableString `json:"bucketId,omitempty"`
	// Number of checklist items that are present on the task.
	ChecklistItemCount NullableInt32 `json:"checklistItemCount,omitempty"`
	// Identity of the user that completed the task.
	CompletedBy AnyOfmicrosoftGraphIdentitySet `json:"completedBy,omitempty"`
	// Read-only. Date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CompletedDateTime NullableTime `json:"completedDateTime,omitempty"`
	// Thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
	ConversationThreadId NullableString `json:"conversationThreadId,omitempty"`
	// Identity of the user that created the task.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// Read-only. Date and time at which the task is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Date and time at which the task is due. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DueDateTime NullableTime `json:"dueDateTime,omitempty"`
	// Read-only. Value is true if the details object of the task has a non-empty description and false otherwise.
	HasDescription NullableBool `json:"hasDescription,omitempty"`
	// Hint used to order items of this type in a list view. The format is defined as outlined here.
	OrderHint NullableString `json:"orderHint,omitempty"`
	// Percentage of task completion. When set to 100, the task is considered completed.
	PercentComplete NullableInt32 `json:"percentComplete,omitempty"`
	// Plan ID to which the task belongs.
	PlanId NullableString `json:"planId,omitempty"`
	// This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference.
	PreviewType AnyOfmicrosoftGraphPlannerPreviewType `json:"previewType,omitempty"`
	// Number of external references that exist on the task.
	ReferenceCount NullableInt32 `json:"referenceCount,omitempty"`
	// Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	StartDateTime NullableTime `json:"startDateTime,omitempty"`
	// Title of the task.
	Title *string `json:"title,omitempty"`
	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
	AssignedToTaskBoardFormat AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`
	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
	BucketTaskBoardFormat AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat `json:"bucketTaskBoardFormat,omitempty"`
	// Read-only. Nullable. Additional details about the task.
	Details AnyOfmicrosoftGraphPlannerTaskDetails `json:"details,omitempty"`
	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
	ProgressTaskBoardFormat AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat `json:"progressTaskBoardFormat,omitempty"`
}

// NewPlannerTask instantiates a new PlannerTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannerTask() *PlannerTask {
	this := PlannerTask{}
	return &this
}

// NewPlannerTaskWithDefaults instantiates a new PlannerTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannerTaskWithDefaults() *PlannerTask {
	this := PlannerTask{}
	return &this
}

// GetActiveChecklistItemCount returns the ActiveChecklistItemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetActiveChecklistItemCount() int32 {
	if o == nil || o.ActiveChecklistItemCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ActiveChecklistItemCount.Get()
}

// GetActiveChecklistItemCountOk returns a tuple with the ActiveChecklistItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetActiveChecklistItemCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActiveChecklistItemCount.Get(), o.ActiveChecklistItemCount.IsSet()
}

// HasActiveChecklistItemCount returns a boolean if a field has been set.
func (o *PlannerTask) HasActiveChecklistItemCount() bool {
	if o != nil && o.ActiveChecklistItemCount.IsSet() {
		return true
	}

	return false
}

// SetActiveChecklistItemCount gets a reference to the given NullableInt32 and assigns it to the ActiveChecklistItemCount field.
func (o *PlannerTask) SetActiveChecklistItemCount(v int32) {
	o.ActiveChecklistItemCount.Set(&v)
}
// SetActiveChecklistItemCountNil sets the value for ActiveChecklistItemCount to be an explicit nil
func (o *PlannerTask) SetActiveChecklistItemCountNil() {
	o.ActiveChecklistItemCount.Set(nil)
}

// UnsetActiveChecklistItemCount ensures that no value is present for ActiveChecklistItemCount, not even an explicit nil
func (o *PlannerTask) UnsetActiveChecklistItemCount() {
	o.ActiveChecklistItemCount.Unset()
}

// GetAppliedCategories returns the AppliedCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetAppliedCategories() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.AppliedCategories
}

// GetAppliedCategoriesOk returns a tuple with the AppliedCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetAppliedCategoriesOk() (*AnyOfobject, bool) {
	if o == nil || o.AppliedCategories == nil {
		return nil, false
	}
	return &o.AppliedCategories, true
}

// HasAppliedCategories returns a boolean if a field has been set.
func (o *PlannerTask) HasAppliedCategories() bool {
	if o != nil && o.AppliedCategories != nil {
		return true
	}

	return false
}

// SetAppliedCategories gets a reference to the given AnyOfobject and assigns it to the AppliedCategories field.
func (o *PlannerTask) SetAppliedCategories(v AnyOfobject) {
	o.AppliedCategories = v
}

// GetAssigneePriority returns the AssigneePriority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetAssigneePriority() string {
	if o == nil || o.AssigneePriority.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssigneePriority.Get()
}

// GetAssigneePriorityOk returns a tuple with the AssigneePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetAssigneePriorityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssigneePriority.Get(), o.AssigneePriority.IsSet()
}

// HasAssigneePriority returns a boolean if a field has been set.
func (o *PlannerTask) HasAssigneePriority() bool {
	if o != nil && o.AssigneePriority.IsSet() {
		return true
	}

	return false
}

// SetAssigneePriority gets a reference to the given NullableString and assigns it to the AssigneePriority field.
func (o *PlannerTask) SetAssigneePriority(v string) {
	o.AssigneePriority.Set(&v)
}
// SetAssigneePriorityNil sets the value for AssigneePriority to be an explicit nil
func (o *PlannerTask) SetAssigneePriorityNil() {
	o.AssigneePriority.Set(nil)
}

// UnsetAssigneePriority ensures that no value is present for AssigneePriority, not even an explicit nil
func (o *PlannerTask) UnsetAssigneePriority() {
	o.AssigneePriority.Unset()
}

// GetAssignments returns the Assignments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetAssignments() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetAssignmentsOk() (*AnyOfobject, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return &o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *PlannerTask) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given AnyOfobject and assigns it to the Assignments field.
func (o *PlannerTask) SetAssignments(v AnyOfobject) {
	o.Assignments = v
}

// GetBucketId returns the BucketId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetBucketId() string {
	if o == nil || o.BucketId.Get() == nil {
		var ret string
		return ret
	}
	return *o.BucketId.Get()
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetBucketIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BucketId.Get(), o.BucketId.IsSet()
}

// HasBucketId returns a boolean if a field has been set.
func (o *PlannerTask) HasBucketId() bool {
	if o != nil && o.BucketId.IsSet() {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given NullableString and assigns it to the BucketId field.
func (o *PlannerTask) SetBucketId(v string) {
	o.BucketId.Set(&v)
}
// SetBucketIdNil sets the value for BucketId to be an explicit nil
func (o *PlannerTask) SetBucketIdNil() {
	o.BucketId.Set(nil)
}

// UnsetBucketId ensures that no value is present for BucketId, not even an explicit nil
func (o *PlannerTask) UnsetBucketId() {
	o.BucketId.Unset()
}

// GetChecklistItemCount returns the ChecklistItemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetChecklistItemCount() int32 {
	if o == nil || o.ChecklistItemCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ChecklistItemCount.Get()
}

// GetChecklistItemCountOk returns a tuple with the ChecklistItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetChecklistItemCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChecklistItemCount.Get(), o.ChecklistItemCount.IsSet()
}

// HasChecklistItemCount returns a boolean if a field has been set.
func (o *PlannerTask) HasChecklistItemCount() bool {
	if o != nil && o.ChecklistItemCount.IsSet() {
		return true
	}

	return false
}

// SetChecklistItemCount gets a reference to the given NullableInt32 and assigns it to the ChecklistItemCount field.
func (o *PlannerTask) SetChecklistItemCount(v int32) {
	o.ChecklistItemCount.Set(&v)
}
// SetChecklistItemCountNil sets the value for ChecklistItemCount to be an explicit nil
func (o *PlannerTask) SetChecklistItemCountNil() {
	o.ChecklistItemCount.Set(nil)
}

// UnsetChecklistItemCount ensures that no value is present for ChecklistItemCount, not even an explicit nil
func (o *PlannerTask) UnsetChecklistItemCount() {
	o.ChecklistItemCount.Unset()
}

// GetCompletedBy returns the CompletedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetCompletedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CompletedBy
}

// GetCompletedByOk returns a tuple with the CompletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetCompletedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CompletedBy == nil {
		return nil, false
	}
	return &o.CompletedBy, true
}

// HasCompletedBy returns a boolean if a field has been set.
func (o *PlannerTask) HasCompletedBy() bool {
	if o != nil && o.CompletedBy != nil {
		return true
	}

	return false
}

// SetCompletedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CompletedBy field.
func (o *PlannerTask) SetCompletedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CompletedBy = v
}

// GetCompletedDateTime returns the CompletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetCompletedDateTime() time.Time {
	if o == nil || o.CompletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedDateTime.Get()
}

// GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetCompletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedDateTime.Get(), o.CompletedDateTime.IsSet()
}

// HasCompletedDateTime returns a boolean if a field has been set.
func (o *PlannerTask) HasCompletedDateTime() bool {
	if o != nil && o.CompletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCompletedDateTime gets a reference to the given NullableTime and assigns it to the CompletedDateTime field.
func (o *PlannerTask) SetCompletedDateTime(v time.Time) {
	o.CompletedDateTime.Set(&v)
}
// SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil
func (o *PlannerTask) SetCompletedDateTimeNil() {
	o.CompletedDateTime.Set(nil)
}

// UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
func (o *PlannerTask) UnsetCompletedDateTime() {
	o.CompletedDateTime.Unset()
}

// GetConversationThreadId returns the ConversationThreadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetConversationThreadId() string {
	if o == nil || o.ConversationThreadId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConversationThreadId.Get()
}

// GetConversationThreadIdOk returns a tuple with the ConversationThreadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetConversationThreadIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConversationThreadId.Get(), o.ConversationThreadId.IsSet()
}

// HasConversationThreadId returns a boolean if a field has been set.
func (o *PlannerTask) HasConversationThreadId() bool {
	if o != nil && o.ConversationThreadId.IsSet() {
		return true
	}

	return false
}

// SetConversationThreadId gets a reference to the given NullableString and assigns it to the ConversationThreadId field.
func (o *PlannerTask) SetConversationThreadId(v string) {
	o.ConversationThreadId.Set(&v)
}
// SetConversationThreadIdNil sets the value for ConversationThreadId to be an explicit nil
func (o *PlannerTask) SetConversationThreadIdNil() {
	o.ConversationThreadId.Set(nil)
}

// UnsetConversationThreadId ensures that no value is present for ConversationThreadId, not even an explicit nil
func (o *PlannerTask) UnsetConversationThreadId() {
	o.ConversationThreadId.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PlannerTask) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *PlannerTask) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *PlannerTask) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *PlannerTask) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *PlannerTask) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *PlannerTask) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDueDateTime returns the DueDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetDueDateTime() time.Time {
	if o == nil || o.DueDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DueDateTime.Get()
}

// GetDueDateTimeOk returns a tuple with the DueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetDueDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DueDateTime.Get(), o.DueDateTime.IsSet()
}

// HasDueDateTime returns a boolean if a field has been set.
func (o *PlannerTask) HasDueDateTime() bool {
	if o != nil && o.DueDateTime.IsSet() {
		return true
	}

	return false
}

// SetDueDateTime gets a reference to the given NullableTime and assigns it to the DueDateTime field.
func (o *PlannerTask) SetDueDateTime(v time.Time) {
	o.DueDateTime.Set(&v)
}
// SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil
func (o *PlannerTask) SetDueDateTimeNil() {
	o.DueDateTime.Set(nil)
}

// UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
func (o *PlannerTask) UnsetDueDateTime() {
	o.DueDateTime.Unset()
}

// GetHasDescription returns the HasDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetHasDescription() bool {
	if o == nil || o.HasDescription.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasDescription.Get()
}

// GetHasDescriptionOk returns a tuple with the HasDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetHasDescriptionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasDescription.Get(), o.HasDescription.IsSet()
}

// HasHasDescription returns a boolean if a field has been set.
func (o *PlannerTask) HasHasDescription() bool {
	if o != nil && o.HasDescription.IsSet() {
		return true
	}

	return false
}

// SetHasDescription gets a reference to the given NullableBool and assigns it to the HasDescription field.
func (o *PlannerTask) SetHasDescription(v bool) {
	o.HasDescription.Set(&v)
}
// SetHasDescriptionNil sets the value for HasDescription to be an explicit nil
func (o *PlannerTask) SetHasDescriptionNil() {
	o.HasDescription.Set(nil)
}

// UnsetHasDescription ensures that no value is present for HasDescription, not even an explicit nil
func (o *PlannerTask) UnsetHasDescription() {
	o.HasDescription.Unset()
}

// GetOrderHint returns the OrderHint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetOrderHint() string {
	if o == nil || o.OrderHint.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrderHint.Get()
}

// GetOrderHintOk returns a tuple with the OrderHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetOrderHintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OrderHint.Get(), o.OrderHint.IsSet()
}

// HasOrderHint returns a boolean if a field has been set.
func (o *PlannerTask) HasOrderHint() bool {
	if o != nil && o.OrderHint.IsSet() {
		return true
	}

	return false
}

// SetOrderHint gets a reference to the given NullableString and assigns it to the OrderHint field.
func (o *PlannerTask) SetOrderHint(v string) {
	o.OrderHint.Set(&v)
}
// SetOrderHintNil sets the value for OrderHint to be an explicit nil
func (o *PlannerTask) SetOrderHintNil() {
	o.OrderHint.Set(nil)
}

// UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil
func (o *PlannerTask) UnsetOrderHint() {
	o.OrderHint.Unset()
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetPercentComplete() int32 {
	if o == nil || o.PercentComplete.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PercentComplete.Get()
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetPercentCompleteOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PercentComplete.Get(), o.PercentComplete.IsSet()
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *PlannerTask) HasPercentComplete() bool {
	if o != nil && o.PercentComplete.IsSet() {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given NullableInt32 and assigns it to the PercentComplete field.
func (o *PlannerTask) SetPercentComplete(v int32) {
	o.PercentComplete.Set(&v)
}
// SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil
func (o *PlannerTask) SetPercentCompleteNil() {
	o.PercentComplete.Set(nil)
}

// UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
func (o *PlannerTask) UnsetPercentComplete() {
	o.PercentComplete.Unset()
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetPlanId() string {
	if o == nil || o.PlanId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetPlanIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *PlannerTask) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *PlannerTask) SetPlanId(v string) {
	o.PlanId.Set(&v)
}
// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *PlannerTask) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *PlannerTask) UnsetPlanId() {
	o.PlanId.Unset()
}

// GetPreviewType returns the PreviewType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerPreviewType
		return ret
	}
	return o.PreviewType
}

// GetPreviewTypeOk returns a tuple with the PreviewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetPreviewTypeOk() (*AnyOfmicrosoftGraphPlannerPreviewType, bool) {
	if o == nil || o.PreviewType == nil {
		return nil, false
	}
	return &o.PreviewType, true
}

// HasPreviewType returns a boolean if a field has been set.
func (o *PlannerTask) HasPreviewType() bool {
	if o != nil && o.PreviewType != nil {
		return true
	}

	return false
}

// SetPreviewType gets a reference to the given AnyOfmicrosoftGraphPlannerPreviewType and assigns it to the PreviewType field.
func (o *PlannerTask) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType) {
	o.PreviewType = v
}

// GetReferenceCount returns the ReferenceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetReferenceCount() int32 {
	if o == nil || o.ReferenceCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceCount.Get()
}

// GetReferenceCountOk returns a tuple with the ReferenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetReferenceCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReferenceCount.Get(), o.ReferenceCount.IsSet()
}

// HasReferenceCount returns a boolean if a field has been set.
func (o *PlannerTask) HasReferenceCount() bool {
	if o != nil && o.ReferenceCount.IsSet() {
		return true
	}

	return false
}

// SetReferenceCount gets a reference to the given NullableInt32 and assigns it to the ReferenceCount field.
func (o *PlannerTask) SetReferenceCount(v int32) {
	o.ReferenceCount.Set(&v)
}
// SetReferenceCountNil sets the value for ReferenceCount to be an explicit nil
func (o *PlannerTask) SetReferenceCountNil() {
	o.ReferenceCount.Set(nil)
}

// UnsetReferenceCount ensures that no value is present for ReferenceCount, not even an explicit nil
func (o *PlannerTask) UnsetReferenceCount() {
	o.ReferenceCount.Unset()
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime.Get()
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDateTime.Get(), o.StartDateTime.IsSet()
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *PlannerTask) HasStartDateTime() bool {
	if o != nil && o.StartDateTime.IsSet() {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given NullableTime and assigns it to the StartDateTime field.
func (o *PlannerTask) SetStartDateTime(v time.Time) {
	o.StartDateTime.Set(&v)
}
// SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil
func (o *PlannerTask) SetStartDateTimeNil() {
	o.StartDateTime.Set(nil)
}

// UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
func (o *PlannerTask) UnsetStartDateTime() {
	o.StartDateTime.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlannerTask) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerTask) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlannerTask) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlannerTask) SetTitle(v string) {
	o.Title = &v
}

// GetAssignedToTaskBoardFormat returns the AssignedToTaskBoardFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetAssignedToTaskBoardFormat() AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
		return ret
	}
	return o.AssignedToTaskBoardFormat
}

// GetAssignedToTaskBoardFormatOk returns a tuple with the AssignedToTaskBoardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetAssignedToTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat, bool) {
	if o == nil || o.AssignedToTaskBoardFormat == nil {
		return nil, false
	}
	return &o.AssignedToTaskBoardFormat, true
}

// HasAssignedToTaskBoardFormat returns a boolean if a field has been set.
func (o *PlannerTask) HasAssignedToTaskBoardFormat() bool {
	if o != nil && o.AssignedToTaskBoardFormat != nil {
		return true
	}

	return false
}

// SetAssignedToTaskBoardFormat gets a reference to the given AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat and assigns it to the AssignedToTaskBoardFormat field.
func (o *PlannerTask) SetAssignedToTaskBoardFormat(v AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) {
	o.AssignedToTaskBoardFormat = v
}

// GetBucketTaskBoardFormat returns the BucketTaskBoardFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetBucketTaskBoardFormat() AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat
		return ret
	}
	return o.BucketTaskBoardFormat
}

// GetBucketTaskBoardFormatOk returns a tuple with the BucketTaskBoardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetBucketTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat, bool) {
	if o == nil || o.BucketTaskBoardFormat == nil {
		return nil, false
	}
	return &o.BucketTaskBoardFormat, true
}

// HasBucketTaskBoardFormat returns a boolean if a field has been set.
func (o *PlannerTask) HasBucketTaskBoardFormat() bool {
	if o != nil && o.BucketTaskBoardFormat != nil {
		return true
	}

	return false
}

// SetBucketTaskBoardFormat gets a reference to the given AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat and assigns it to the BucketTaskBoardFormat field.
func (o *PlannerTask) SetBucketTaskBoardFormat(v AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat) {
	o.BucketTaskBoardFormat = v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetDetails() AnyOfmicrosoftGraphPlannerTaskDetails {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerTaskDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetDetailsOk() (*AnyOfmicrosoftGraphPlannerTaskDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PlannerTask) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given AnyOfmicrosoftGraphPlannerTaskDetails and assigns it to the Details field.
func (o *PlannerTask) SetDetails(v AnyOfmicrosoftGraphPlannerTaskDetails) {
	o.Details = v
}

// GetProgressTaskBoardFormat returns the ProgressTaskBoardFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerTask) GetProgressTaskBoardFormat() AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat
		return ret
	}
	return o.ProgressTaskBoardFormat
}

// GetProgressTaskBoardFormatOk returns a tuple with the ProgressTaskBoardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerTask) GetProgressTaskBoardFormatOk() (*AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat, bool) {
	if o == nil || o.ProgressTaskBoardFormat == nil {
		return nil, false
	}
	return &o.ProgressTaskBoardFormat, true
}

// HasProgressTaskBoardFormat returns a boolean if a field has been set.
func (o *PlannerTask) HasProgressTaskBoardFormat() bool {
	if o != nil && o.ProgressTaskBoardFormat != nil {
		return true
	}

	return false
}

// SetProgressTaskBoardFormat gets a reference to the given AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat and assigns it to the ProgressTaskBoardFormat field.
func (o *PlannerTask) SetProgressTaskBoardFormat(v AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat) {
	o.ProgressTaskBoardFormat = v
}

func (o PlannerTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveChecklistItemCount.IsSet() {
		toSerialize["activeChecklistItemCount"] = o.ActiveChecklistItemCount.Get()
	}
	if o.AppliedCategories != nil {
		toSerialize["appliedCategories"] = o.AppliedCategories
	}
	if o.AssigneePriority.IsSet() {
		toSerialize["assigneePriority"] = o.AssigneePriority.Get()
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.BucketId.IsSet() {
		toSerialize["bucketId"] = o.BucketId.Get()
	}
	if o.ChecklistItemCount.IsSet() {
		toSerialize["checklistItemCount"] = o.ChecklistItemCount.Get()
	}
	if o.CompletedBy != nil {
		toSerialize["completedBy"] = o.CompletedBy
	}
	if o.CompletedDateTime.IsSet() {
		toSerialize["completedDateTime"] = o.CompletedDateTime.Get()
	}
	if o.ConversationThreadId.IsSet() {
		toSerialize["conversationThreadId"] = o.ConversationThreadId.Get()
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.DueDateTime.IsSet() {
		toSerialize["dueDateTime"] = o.DueDateTime.Get()
	}
	if o.HasDescription.IsSet() {
		toSerialize["hasDescription"] = o.HasDescription.Get()
	}
	if o.OrderHint.IsSet() {
		toSerialize["orderHint"] = o.OrderHint.Get()
	}
	if o.PercentComplete.IsSet() {
		toSerialize["percentComplete"] = o.PercentComplete.Get()
	}
	if o.PlanId.IsSet() {
		toSerialize["planId"] = o.PlanId.Get()
	}
	if o.PreviewType != nil {
		toSerialize["previewType"] = o.PreviewType
	}
	if o.ReferenceCount.IsSet() {
		toSerialize["referenceCount"] = o.ReferenceCount.Get()
	}
	if o.StartDateTime.IsSet() {
		toSerialize["startDateTime"] = o.StartDateTime.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.AssignedToTaskBoardFormat != nil {
		toSerialize["assignedToTaskBoardFormat"] = o.AssignedToTaskBoardFormat
	}
	if o.BucketTaskBoardFormat != nil {
		toSerialize["bucketTaskBoardFormat"] = o.BucketTaskBoardFormat
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.ProgressTaskBoardFormat != nil {
		toSerialize["progressTaskBoardFormat"] = o.ProgressTaskBoardFormat
	}
	return json.Marshal(toSerialize)
}

type NullablePlannerTask struct {
	value *PlannerTask
	isSet bool
}

func (v NullablePlannerTask) Get() *PlannerTask {
	return v.value
}

func (v *NullablePlannerTask) Set(val *PlannerTask) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannerTask) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannerTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannerTask(val *PlannerTask) *NullablePlannerTask {
	return &NullablePlannerTask{value: val, isSet: true}
}

func (v NullablePlannerTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannerTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


