# MicrosoftGraphTodoTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphTodoTask

`func NewMicrosoftGraphTodoTask() *MicrosoftGraphTodoTask`

NewMicrosoftGraphTodoTask instantiates a new MicrosoftGraphTodoTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTodoTaskWithDefaults

`func NewMicrosoftGraphTodoTaskWithDefaults() *MicrosoftGraphTodoTask`

NewMicrosoftGraphTodoTaskWithDefaults instantiates a new MicrosoftGraphTodoTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTodoTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTodoTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTodoTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTodoTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *MicrosoftGraphTodoTask) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphTodoTask) GetBodyOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MicrosoftGraphTodoTask) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *MicrosoftGraphTodoTask) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *MicrosoftGraphTodoTask) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *MicrosoftGraphTodoTask) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBodyLastModifiedDateTime

`func (o *MicrosoftGraphTodoTask) GetBodyLastModifiedDateTime() time.Time`

GetBodyLastModifiedDateTime returns the BodyLastModifiedDateTime field if non-nil, zero value otherwise.

### GetBodyLastModifiedDateTimeOk

`func (o *MicrosoftGraphTodoTask) GetBodyLastModifiedDateTimeOk() (*time.Time, bool)`

GetBodyLastModifiedDateTimeOk returns a tuple with the BodyLastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyLastModifiedDateTime

`func (o *MicrosoftGraphTodoTask) SetBodyLastModifiedDateTime(v time.Time)`

SetBodyLastModifiedDateTime sets BodyLastModifiedDateTime field to given value.

### HasBodyLastModifiedDateTime

`func (o *MicrosoftGraphTodoTask) HasBodyLastModifiedDateTime() bool`

HasBodyLastModifiedDateTime returns a boolean if a field has been set.

### GetCompletedDateTime

`func (o *MicrosoftGraphTodoTask) GetCompletedDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *MicrosoftGraphTodoTask) GetCompletedDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *MicrosoftGraphTodoTask) SetCompletedDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *MicrosoftGraphTodoTask) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *MicrosoftGraphTodoTask) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *MicrosoftGraphTodoTask) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphTodoTask) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTodoTask) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTodoTask) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTodoTask) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDueDateTime

`func (o *MicrosoftGraphTodoTask) GetDueDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *MicrosoftGraphTodoTask) GetDueDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *MicrosoftGraphTodoTask) SetDueDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *MicrosoftGraphTodoTask) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *MicrosoftGraphTodoTask) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *MicrosoftGraphTodoTask) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetImportance

`func (o *MicrosoftGraphTodoTask) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphTodoTask) GetImportanceOk() (*AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *MicrosoftGraphTodoTask) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *MicrosoftGraphTodoTask) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportanceNil

`func (o *MicrosoftGraphTodoTask) SetImportanceNil(b bool)`

 SetImportanceNil sets the value for Importance to be an explicit nil

### UnsetImportance
`func (o *MicrosoftGraphTodoTask) UnsetImportance()`

UnsetImportance ensures that no value is present for Importance, not even an explicit nil
### GetIsReminderOn

`func (o *MicrosoftGraphTodoTask) GetIsReminderOn() bool`

GetIsReminderOn returns the IsReminderOn field if non-nil, zero value otherwise.

### GetIsReminderOnOk

`func (o *MicrosoftGraphTodoTask) GetIsReminderOnOk() (*bool, bool)`

GetIsReminderOnOk returns a tuple with the IsReminderOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReminderOn

`func (o *MicrosoftGraphTodoTask) SetIsReminderOn(v bool)`

SetIsReminderOn sets IsReminderOn field to given value.

### HasIsReminderOn

`func (o *MicrosoftGraphTodoTask) HasIsReminderOn() bool`

HasIsReminderOn returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTodoTask) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTodoTask) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTodoTask) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTodoTask) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetRecurrence

`func (o *MicrosoftGraphTodoTask) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *MicrosoftGraphTodoTask) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *MicrosoftGraphTodoTask) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *MicrosoftGraphTodoTask) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *MicrosoftGraphTodoTask) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *MicrosoftGraphTodoTask) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetReminderDateTime

`func (o *MicrosoftGraphTodoTask) GetReminderDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetReminderDateTime returns the ReminderDateTime field if non-nil, zero value otherwise.

### GetReminderDateTimeOk

`func (o *MicrosoftGraphTodoTask) GetReminderDateTimeOk() (*AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetReminderDateTimeOk returns a tuple with the ReminderDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderDateTime

`func (o *MicrosoftGraphTodoTask) SetReminderDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetReminderDateTime sets ReminderDateTime field to given value.

### HasReminderDateTime

`func (o *MicrosoftGraphTodoTask) HasReminderDateTime() bool`

HasReminderDateTime returns a boolean if a field has been set.

### SetReminderDateTimeNil

`func (o *MicrosoftGraphTodoTask) SetReminderDateTimeNil(b bool)`

 SetReminderDateTimeNil sets the value for ReminderDateTime to be an explicit nil

### UnsetReminderDateTime
`func (o *MicrosoftGraphTodoTask) UnsetReminderDateTime()`

UnsetReminderDateTime ensures that no value is present for ReminderDateTime, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphTodoTask) GetStatus() AnyOfmicrosoftGraphTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphTodoTask) GetStatusOk() (*AnyOfmicrosoftGraphTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphTodoTask) SetStatus(v AnyOfmicrosoftGraphTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphTodoTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphTodoTask) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphTodoTask) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphTodoTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphTodoTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphTodoTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphTodoTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphTodoTask) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphTodoTask) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExtensions

`func (o *MicrosoftGraphTodoTask) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphTodoTask) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphTodoTask) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphTodoTask) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetLinkedResources

`func (o *MicrosoftGraphTodoTask) GetLinkedResources() []MicrosoftGraphLinkedResource`

GetLinkedResources returns the LinkedResources field if non-nil, zero value otherwise.

### GetLinkedResourcesOk

`func (o *MicrosoftGraphTodoTask) GetLinkedResourcesOk() (*[]MicrosoftGraphLinkedResource, bool)`

GetLinkedResourcesOk returns a tuple with the LinkedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResources

`func (o *MicrosoftGraphTodoTask) SetLinkedResources(v []MicrosoftGraphLinkedResource)`

SetLinkedResources sets LinkedResources field to given value.

### HasLinkedResources

`func (o *MicrosoftGraphTodoTask) HasLinkedResources() bool`

HasLinkedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


