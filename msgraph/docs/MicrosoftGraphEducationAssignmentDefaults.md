# MicrosoftGraphEducationAssignmentDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AddedStudentAction** | Pointer to [**AnyOfmicrosoftGraphEducationAddedStudentAction**](anyOf&lt;microsoft.graph.educationAddedStudentAction&gt;.md) | Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen. | [optional] 
**DueTime** | Pointer to **NullableString** | Class-level default value for due time field. Default value is 23:59:00. | [optional] 
**NotificationChannelUrl** | Pointer to **NullableString** | Default Teams channel to which notifications will be sent. Default value is null. | [optional] 

## Methods

### NewMicrosoftGraphEducationAssignmentDefaults

`func NewMicrosoftGraphEducationAssignmentDefaults() *MicrosoftGraphEducationAssignmentDefaults`

NewMicrosoftGraphEducationAssignmentDefaults instantiates a new MicrosoftGraphEducationAssignmentDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationAssignmentDefaultsWithDefaults

`func NewMicrosoftGraphEducationAssignmentDefaultsWithDefaults() *MicrosoftGraphEducationAssignmentDefaults`

NewMicrosoftGraphEducationAssignmentDefaultsWithDefaults instantiates a new MicrosoftGraphEducationAssignmentDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationAssignmentDefaults) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddedStudentAction

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetAddedStudentAction() AnyOfmicrosoftGraphEducationAddedStudentAction`

GetAddedStudentAction returns the AddedStudentAction field if non-nil, zero value otherwise.

### GetAddedStudentActionOk

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetAddedStudentActionOk() (*AnyOfmicrosoftGraphEducationAddedStudentAction, bool)`

GetAddedStudentActionOk returns a tuple with the AddedStudentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedStudentAction

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetAddedStudentAction(v AnyOfmicrosoftGraphEducationAddedStudentAction)`

SetAddedStudentAction sets AddedStudentAction field to given value.

### HasAddedStudentAction

`func (o *MicrosoftGraphEducationAssignmentDefaults) HasAddedStudentAction() bool`

HasAddedStudentAction returns a boolean if a field has been set.

### SetAddedStudentActionNil

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetAddedStudentActionNil(b bool)`

 SetAddedStudentActionNil sets the value for AddedStudentAction to be an explicit nil

### UnsetAddedStudentAction
`func (o *MicrosoftGraphEducationAssignmentDefaults) UnsetAddedStudentAction()`

UnsetAddedStudentAction ensures that no value is present for AddedStudentAction, not even an explicit nil
### GetDueTime

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetDueTime() string`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetDueTimeOk() (*string, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetDueTime(v string)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *MicrosoftGraphEducationAssignmentDefaults) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### SetDueTimeNil

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetDueTimeNil(b bool)`

 SetDueTimeNil sets the value for DueTime to be an explicit nil

### UnsetDueTime
`func (o *MicrosoftGraphEducationAssignmentDefaults) UnsetDueTime()`

UnsetDueTime ensures that no value is present for DueTime, not even an explicit nil
### GetNotificationChannelUrl

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetNotificationChannelUrl() string`

GetNotificationChannelUrl returns the NotificationChannelUrl field if non-nil, zero value otherwise.

### GetNotificationChannelUrlOk

`func (o *MicrosoftGraphEducationAssignmentDefaults) GetNotificationChannelUrlOk() (*string, bool)`

GetNotificationChannelUrlOk returns a tuple with the NotificationChannelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannelUrl

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetNotificationChannelUrl(v string)`

SetNotificationChannelUrl sets NotificationChannelUrl field to given value.

### HasNotificationChannelUrl

`func (o *MicrosoftGraphEducationAssignmentDefaults) HasNotificationChannelUrl() bool`

HasNotificationChannelUrl returns a boolean if a field has been set.

### SetNotificationChannelUrlNil

`func (o *MicrosoftGraphEducationAssignmentDefaults) SetNotificationChannelUrlNil(b bool)`

 SetNotificationChannelUrlNil sets the value for NotificationChannelUrl to be an explicit nil

### UnsetNotificationChannelUrl
`func (o *MicrosoftGraphEducationAssignmentDefaults) UnsetNotificationChannelUrl()`

UnsetNotificationChannelUrl ensures that no value is present for NotificationChannelUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


