# EducationAssignmentDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedStudentAction** | Pointer to [**AnyOfmicrosoftGraphEducationAddedStudentAction**](anyOf&lt;microsoft.graph.educationAddedStudentAction&gt;.md) | Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen. | [optional] 
**DueTime** | Pointer to **NullableString** | Class-level default value for due time field. Default value is 23:59:00. | [optional] 
**NotificationChannelUrl** | Pointer to **NullableString** | Default Teams channel to which notifications will be sent. Default value is null. | [optional] 

## Methods

### NewEducationAssignmentDefaults

`func NewEducationAssignmentDefaults() *EducationAssignmentDefaults`

NewEducationAssignmentDefaults instantiates a new EducationAssignmentDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationAssignmentDefaultsWithDefaults

`func NewEducationAssignmentDefaultsWithDefaults() *EducationAssignmentDefaults`

NewEducationAssignmentDefaultsWithDefaults instantiates a new EducationAssignmentDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedStudentAction

`func (o *EducationAssignmentDefaults) GetAddedStudentAction() AnyOfmicrosoftGraphEducationAddedStudentAction`

GetAddedStudentAction returns the AddedStudentAction field if non-nil, zero value otherwise.

### GetAddedStudentActionOk

`func (o *EducationAssignmentDefaults) GetAddedStudentActionOk() (*AnyOfmicrosoftGraphEducationAddedStudentAction, bool)`

GetAddedStudentActionOk returns a tuple with the AddedStudentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedStudentAction

`func (o *EducationAssignmentDefaults) SetAddedStudentAction(v AnyOfmicrosoftGraphEducationAddedStudentAction)`

SetAddedStudentAction sets AddedStudentAction field to given value.

### HasAddedStudentAction

`func (o *EducationAssignmentDefaults) HasAddedStudentAction() bool`

HasAddedStudentAction returns a boolean if a field has been set.

### SetAddedStudentActionNil

`func (o *EducationAssignmentDefaults) SetAddedStudentActionNil(b bool)`

 SetAddedStudentActionNil sets the value for AddedStudentAction to be an explicit nil

### UnsetAddedStudentAction
`func (o *EducationAssignmentDefaults) UnsetAddedStudentAction()`

UnsetAddedStudentAction ensures that no value is present for AddedStudentAction, not even an explicit nil
### GetDueTime

`func (o *EducationAssignmentDefaults) GetDueTime() string`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *EducationAssignmentDefaults) GetDueTimeOk() (*string, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *EducationAssignmentDefaults) SetDueTime(v string)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *EducationAssignmentDefaults) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### SetDueTimeNil

`func (o *EducationAssignmentDefaults) SetDueTimeNil(b bool)`

 SetDueTimeNil sets the value for DueTime to be an explicit nil

### UnsetDueTime
`func (o *EducationAssignmentDefaults) UnsetDueTime()`

UnsetDueTime ensures that no value is present for DueTime, not even an explicit nil
### GetNotificationChannelUrl

`func (o *EducationAssignmentDefaults) GetNotificationChannelUrl() string`

GetNotificationChannelUrl returns the NotificationChannelUrl field if non-nil, zero value otherwise.

### GetNotificationChannelUrlOk

`func (o *EducationAssignmentDefaults) GetNotificationChannelUrlOk() (*string, bool)`

GetNotificationChannelUrlOk returns a tuple with the NotificationChannelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannelUrl

`func (o *EducationAssignmentDefaults) SetNotificationChannelUrl(v string)`

SetNotificationChannelUrl sets NotificationChannelUrl field to given value.

### HasNotificationChannelUrl

`func (o *EducationAssignmentDefaults) HasNotificationChannelUrl() bool`

HasNotificationChannelUrl returns a boolean if a field has been set.

### SetNotificationChannelUrlNil

`func (o *EducationAssignmentDefaults) SetNotificationChannelUrlNil(b bool)`

 SetNotificationChannelUrlNil sets the value for NotificationChannelUrl to be an explicit nil

### UnsetNotificationChannelUrl
`func (o *EducationAssignmentDefaults) UnsetNotificationChannelUrl()`

UnsetNotificationChannelUrl ensures that no value is present for NotificationChannelUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


