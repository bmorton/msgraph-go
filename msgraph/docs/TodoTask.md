# TodoTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The task body that typically contains information about the task. | [optional] 
**BodyLastModifiedDateTime** | Pointer to **time.Time** | The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: &#39;2020-01-01T00:00:00Z&#39;. | [optional] 
**CompletedDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date in the specified time zone that the task was finished. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: &#39;2020-01-01T00:00:00Z&#39;. | [optional] 
**DueDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date in the specified time zone that the task is to be finished. | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) | The importance of the task. Possible values are: low, normal, high. | [optional] 
**IsReminderOn** | Pointer to **bool** | Set to true if an alert is set to remind the user of the task. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: &#39;2020-01-01T00:00:00Z&#39;. | [optional] 
**Recurrence** | Pointer to [**AnyOfmicrosoftGraphPatternedRecurrence**](anyOf&lt;microsoft.graph.patternedRecurrence&gt;.md) | The recurrence pattern for the task. | [optional] 
**ReminderDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) | The date and time for a reminder alert of the task to occur. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphTaskStatus**](anyOf&lt;microsoft.graph.taskStatus&gt;.md) | Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed, waitingOnOthers, deferred. | [optional] 
**Title** | Pointer to **NullableString** | A brief description of the task. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the task. Nullable. | [optional] 
**LinkedResources** | Pointer to [**[]MicrosoftGraphLinkedResource**](MicrosoftGraphLinkedResource.md) | A collection of resources linked to the task. | [optional] 

## Methods

### NewTodoTask

`func NewTodoTask() *TodoTask`

NewTodoTask instantiates a new TodoTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoTaskWithDefaults

`func NewTodoTaskWithDefaults() *TodoTask`

NewTodoTaskWithDefaults instantiates a new TodoTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TodoTask) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TodoTask) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TodoTask) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *TodoTask) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *TodoTask) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *TodoTask) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyLastModifiedDateTime

`func (o *TodoTask) GetBodyLastModifiedDateTime() time.Time`

GetBodyLastModifiedDateTime returns the BodyLastModifiedDateTime field if non-nil, zero value otherwise.

### GetBodyLastModifiedDateTimeOk

`func (o *TodoTask) GetBodyLastModifiedDateTimeOk() (*time.Time, bool)`

GetBodyLastModifiedDateTimeOk returns a tuple with the BodyLastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyLastModifiedDateTime

`func (o *TodoTask) SetBodyLastModifiedDateTime(v time.Time)`

SetBodyLastModifiedDateTime sets BodyLastModifiedDateTime field to given value.

### HasBodyLastModifiedDateTime

`func (o *TodoTask) HasBodyLastModifiedDateTime() bool`

HasBodyLastModifiedDateTime returns a boolean if a field has been set.

### GetCompletedDateTime

`func (o *TodoTask) GetCompletedDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *TodoTask) GetCompletedDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *TodoTask) SetCompletedDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *TodoTask) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *TodoTask) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *TodoTask) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetCreatedDateTime

`func (o *TodoTask) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *TodoTask) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *TodoTask) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *TodoTask) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDueDateTime

`func (o *TodoTask) GetDueDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *TodoTask) GetDueDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *TodoTask) SetDueDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *TodoTask) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *TodoTask) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *TodoTask) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetImportance

`func (o *TodoTask) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *TodoTask) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *TodoTask) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *TodoTask) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *TodoTask) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *TodoTask) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetIsReminderOn

`func (o *TodoTask) GetIsReminderOn() bool`

GetIsReminderOn returns the IsReminderOn field if non-nil, zero value otherwise.

### GetIsReminderOnOk

`func (o *TodoTask) GetIsReminderOnOk() (*bool, bool)`

GetIsReminderOnOk returns a tuple with the IsReminderOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReminderOn

`func (o *TodoTask) SetIsReminderOn(v bool)`

SetIsReminderOn sets IsReminderOn field to given value.

### HasIsReminderOn

`func (o *TodoTask) HasIsReminderOn() bool`

HasIsReminderOn returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *TodoTask) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *TodoTask) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *TodoTask) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *TodoTask) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetRecurrence

`func (o *TodoTask) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *TodoTask) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *TodoTask) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *TodoTask) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *TodoTask) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *TodoTask) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetReminderDateTime

`func (o *TodoTask) GetReminderDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetReminderDateTime returns the ReminderDateTime field if non-nil, zero value otherwise.

### GetReminderDateTimeOk

`func (o *TodoTask) GetReminderDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetReminderDateTimeOk returns a tuple with the ReminderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDateTime

`func (o *TodoTask) SetReminderDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetReminderDateTime sets ReminderDateTime field to given value.

### HasReminderDateTime

`func (o *TodoTask) HasReminderDateTime() bool`

HasReminderDateTime returns a boolean if a field has been set.

### SetReminderDateTimeNil

`func (o *TodoTask) SetReminderDateTimeNil(b bool)`

 SetReminderDateTimeNil sets the value for ReminderDateTime to be an explicit nil

### UnsetReminderDateTime
`func (o *TodoTask) UnsetReminderDateTime()`

UnsetReminderDateTime ensures that no value is present for ReminderDateTime, not even an explicit nil
### GetStatus

`func (o *TodoTask) GetStatus() AnyOfmicrosoftGraphTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TodoTask) GetStatusOk() (*AnyOfmicrosoftGraphTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TodoTask) SetStatus(v AnyOfmicrosoftGraphTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TodoTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TodoTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TodoTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTitle

`func (o *TodoTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TodoTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TodoTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TodoTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TodoTask) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TodoTask) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExtensions

`func (o *TodoTask) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *TodoTask) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *TodoTask) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *TodoTask) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetLinkedResources

`func (o *TodoTask) GetLinkedResources() []MicrosoftGraphLinkedResource`

GetLinkedResources returns the LinkedResources field if non-nil, zero value otherwise.

### GetLinkedResourcesOk

`func (o *TodoTask) GetLinkedResourcesOk() (*[]MicrosoftGraphLinkedResource, bool)`

GetLinkedResourcesOk returns a tuple with the LinkedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResources

`func (o *TodoTask) SetLinkedResources(v []MicrosoftGraphLinkedResource)`

SetLinkedResources sets LinkedResources field to given value.

### HasLinkedResources

`func (o *TodoTask) HasLinkedResources() bool`

HasLinkedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


