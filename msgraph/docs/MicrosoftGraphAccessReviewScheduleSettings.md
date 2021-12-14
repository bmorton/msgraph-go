# MicrosoftGraphAccessReviewScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyActions** | Pointer to [**[]AnyOfobject**](AnyOfobject.md) | Optional field. Describes the  actions to take once a review is complete. There are two types that are currently supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in the case of disableAndDeleteUserApplyAction. | [optional] 
**AutoApplyDecisionsEnabled** | Pointer to **bool** | Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually once the reviewer completes the access review. When set to true, decisions are applied automatically after the access review instance duration ends, whether or not the reviewers have responded. Default value is false. | [optional] 
**DefaultDecision** | Pointer to **NullableString** | Decision chosen if defaultDecisionEnabled is true. Can be one of Approve, Deny, or Recommendation. | [optional] 
**DefaultDecisionEnabled** | Pointer to **bool** | Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false. | [optional] 
**InstanceDurationInDays** | Pointer to **int32** | Duration of each recurrence of review (accessReviewInstance) in number of days. | [optional] 
**JustificationRequiredOnApproval** | Pointer to **bool** | Indicates whether reviewers are required to provide justification with their decision. Default value is false. | [optional] 
**MailNotificationsEnabled** | Pointer to **bool** | Indicates whether emails are enabled or disabled. Default value is false. | [optional] 
**RecommendationsEnabled** | Pointer to **bool** | Indicates whether decision recommendations are enabled or disabled. | [optional] 
**Recurrence** | Pointer to [**AnyOfmicrosoftGraphPatternedRecurrence**](anyOf&lt;microsoft.graph.patternedRecurrence&gt;.md) | Detailed settings for recurrence using the standard Outlook recurrence object. Only weekly and absoluteMonthly on recurrencePattern are supported. Use the property startDate on recurrenceRange to determine the day the review starts. | [optional] 
**ReminderNotificationsEnabled** | Pointer to **bool** | Indicates whether reminders are enabled or disabled. Default value is false. | [optional] 

## Methods

### NewMicrosoftGraphAccessReviewScheduleSettings

`func NewMicrosoftGraphAccessReviewScheduleSettings() *MicrosoftGraphAccessReviewScheduleSettings`

NewMicrosoftGraphAccessReviewScheduleSettings instantiates a new MicrosoftGraphAccessReviewScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessReviewScheduleSettingsWithDefaults

`func NewMicrosoftGraphAccessReviewScheduleSettingsWithDefaults() *MicrosoftGraphAccessReviewScheduleSettings`

NewMicrosoftGraphAccessReviewScheduleSettingsWithDefaults instantiates a new MicrosoftGraphAccessReviewScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyActions

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetApplyActions() []*AnyOfobject`

GetApplyActions returns the ApplyActions field if non-nil, zero value otherwise.

### GetApplyActionsOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetApplyActionsOk() (*[]*AnyOfobject, bool)`

GetApplyActionsOk returns a tuple with the ApplyActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyActions

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetApplyActions(v []*AnyOfobject)`

SetApplyActions sets ApplyActions field to given value.

### HasApplyActions

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasApplyActions() bool`

HasApplyActions returns a boolean if a field has been set.

### GetAutoApplyDecisionsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetAutoApplyDecisionsEnabled() bool`

GetAutoApplyDecisionsEnabled returns the AutoApplyDecisionsEnabled field if non-nil, zero value otherwise.

### GetAutoApplyDecisionsEnabledOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetAutoApplyDecisionsEnabledOk() (*bool, bool)`

GetAutoApplyDecisionsEnabledOk returns a tuple with the AutoApplyDecisionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApplyDecisionsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetAutoApplyDecisionsEnabled(v bool)`

SetAutoApplyDecisionsEnabled sets AutoApplyDecisionsEnabled field to given value.

### HasAutoApplyDecisionsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasAutoApplyDecisionsEnabled() bool`

HasAutoApplyDecisionsEnabled returns a boolean if a field has been set.

### GetDefaultDecision

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetDefaultDecision() string`

GetDefaultDecision returns the DefaultDecision field if non-nil, zero value otherwise.

### GetDefaultDecisionOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetDefaultDecisionOk() (*string, bool)`

GetDefaultDecisionOk returns a tuple with the DefaultDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDecision

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetDefaultDecision(v string)`

SetDefaultDecision sets DefaultDecision field to given value.

### HasDefaultDecision

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasDefaultDecision() bool`

HasDefaultDecision returns a boolean if a field has been set.

### SetDefaultDecisionNil

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetDefaultDecisionNil(b bool)`

 SetDefaultDecisionNil sets the value for DefaultDecision to be an explicit nil

### UnsetDefaultDecision
`func (o *MicrosoftGraphAccessReviewScheduleSettings) UnsetDefaultDecision()`

UnsetDefaultDecision ensures that no value is present for DefaultDecision, not even an explicit nil
### GetDefaultDecisionEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetDefaultDecisionEnabled() bool`

GetDefaultDecisionEnabled returns the DefaultDecisionEnabled field if non-nil, zero value otherwise.

### GetDefaultDecisionEnabledOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetDefaultDecisionEnabledOk() (*bool, bool)`

GetDefaultDecisionEnabledOk returns a tuple with the DefaultDecisionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDecisionEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetDefaultDecisionEnabled(v bool)`

SetDefaultDecisionEnabled sets DefaultDecisionEnabled field to given value.

### HasDefaultDecisionEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasDefaultDecisionEnabled() bool`

HasDefaultDecisionEnabled returns a boolean if a field has been set.

### GetInstanceDurationInDays

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetInstanceDurationInDays() int32`

GetInstanceDurationInDays returns the InstanceDurationInDays field if non-nil, zero value otherwise.

### GetInstanceDurationInDaysOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetInstanceDurationInDaysOk() (*int32, bool)`

GetInstanceDurationInDaysOk returns a tuple with the InstanceDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDurationInDays

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetInstanceDurationInDays(v int32)`

SetInstanceDurationInDays sets InstanceDurationInDays field to given value.

### HasInstanceDurationInDays

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasInstanceDurationInDays() bool`

HasInstanceDurationInDays returns a boolean if a field has been set.

### GetJustificationRequiredOnApproval

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetJustificationRequiredOnApproval() bool`

GetJustificationRequiredOnApproval returns the JustificationRequiredOnApproval field if non-nil, zero value otherwise.

### GetJustificationRequiredOnApprovalOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetJustificationRequiredOnApprovalOk() (*bool, bool)`

GetJustificationRequiredOnApprovalOk returns a tuple with the JustificationRequiredOnApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustificationRequiredOnApproval

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetJustificationRequiredOnApproval(v bool)`

SetJustificationRequiredOnApproval sets JustificationRequiredOnApproval field to given value.

### HasJustificationRequiredOnApproval

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasJustificationRequiredOnApproval() bool`

HasJustificationRequiredOnApproval returns a boolean if a field has been set.

### GetMailNotificationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetMailNotificationsEnabled() bool`

GetMailNotificationsEnabled returns the MailNotificationsEnabled field if non-nil, zero value otherwise.

### GetMailNotificationsEnabledOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetMailNotificationsEnabledOk() (*bool, bool)`

GetMailNotificationsEnabledOk returns a tuple with the MailNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNotificationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetMailNotificationsEnabled(v bool)`

SetMailNotificationsEnabled sets MailNotificationsEnabled field to given value.

### HasMailNotificationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasMailNotificationsEnabled() bool`

HasMailNotificationsEnabled returns a boolean if a field has been set.

### GetRecommendationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetRecommendationsEnabled() bool`

GetRecommendationsEnabled returns the RecommendationsEnabled field if non-nil, zero value otherwise.

### GetRecommendationsEnabledOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetRecommendationsEnabledOk() (*bool, bool)`

GetRecommendationsEnabledOk returns a tuple with the RecommendationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetRecommendationsEnabled(v bool)`

SetRecommendationsEnabled sets RecommendationsEnabled field to given value.

### HasRecommendationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasRecommendationsEnabled() bool`

HasRecommendationsEnabled returns a boolean if a field has been set.

### GetRecurrence

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *MicrosoftGraphAccessReviewScheduleSettings) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetReminderNotificationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetReminderNotificationsEnabled() bool`

GetReminderNotificationsEnabled returns the ReminderNotificationsEnabled field if non-nil, zero value otherwise.

### GetReminderNotificationsEnabledOk

`func (o *MicrosoftGraphAccessReviewScheduleSettings) GetReminderNotificationsEnabledOk() (*bool, bool)`

GetReminderNotificationsEnabledOk returns a tuple with the ReminderNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderNotificationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) SetReminderNotificationsEnabled(v bool)`

SetReminderNotificationsEnabled sets ReminderNotificationsEnabled field to given value.

### HasReminderNotificationsEnabled

`func (o *MicrosoftGraphAccessReviewScheduleSettings) HasReminderNotificationsEnabled() bool`

HasReminderNotificationsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


