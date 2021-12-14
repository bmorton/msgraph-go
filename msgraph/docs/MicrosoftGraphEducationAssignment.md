# MicrosoftGraphEducationAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AddedStudentAction** | Pointer to [**AnyOfmicrosoftGraphEducationAddedStudentAction**](anyOf&lt;microsoft.graph.educationAddedStudentAction&gt;.md) | Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none value. Currently supports only two values: none or assignIfOpen. | [optional] 
**AllowLateSubmissions** | Pointer to **NullableBool** | Identifies whether students can submit after the due date. If this property isn&#39;t specified during create, it defaults to true. | [optional] 
**AllowStudentsToAddResourcesToSubmission** | Pointer to **NullableBool** | Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher. | [optional] 
**AssignDateTime** | Pointer to **NullableTime** | The date when the assignment should become active.  If in the future, the assignment isn&#39;t shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**AssignedDateTime** | Pointer to **NullableTime** | The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**AssignTo** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Which users, or whole class should receive a submission object once the assignment is published. | [optional] 
**ClassId** | Pointer to **NullableString** | Class which this assignment belongs. | [optional] 
**CloseDateTime** | Pointer to **NullableTime** | Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Who created the assignment. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the assignment. | [optional] 
**DueDateTime** | Pointer to **NullableTime** | Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Grading** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | How the assignment will be graded. | [optional] 
**Instructions** | Pointer to [**AnyOfmicrosoftGraphEducationItemBody**](anyOf&lt;microsoft.graph.educationItemBody&gt;.md) | Instructions for the assignment.  This along with the display name tell the student what to do. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Who last modified the assignment. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**NotificationChannelUrl** | Pointer to **NullableString** | Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn&#39;t allowed after the assignment has been published. | [optional] 
**ResourcesFolderUrl** | Pointer to **NullableString** | Folder URL where all the file resources for this assignment are stored. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphEducationAssignmentStatus**](anyOf&lt;microsoft.graph.educationAssignmentStatus&gt;.md) | Status of the Assignment.  You can&#39;t PATCH this value.  Possible values are: draft, scheduled, published, assigned. | [optional] 
**WebUrl** | Pointer to **NullableString** | The deep link URL for the given assignment. | [optional] 
**Categories** | Pointer to [**[]MicrosoftGraphEducationCategory**](MicrosoftGraphEducationCategory.md) | When set, enables users to easily find assignments of a given type.  Read-only. Nullable. | [optional] 
**Resources** | Pointer to [**[]MicrosoftGraphEducationAssignmentResource**](MicrosoftGraphEducationAssignmentResource.md) | Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable. | [optional] 
**Rubric** | Pointer to [**AnyOfmicrosoftGraphEducationRubric**](anyOf&lt;microsoft.graph.educationRubric&gt;.md) | When set, the grading rubric attached to this assignment. | [optional] 
**Submissions** | Pointer to [**[]MicrosoftGraphEducationSubmission**](MicrosoftGraphEducationSubmission.md) | Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphEducationAssignment

`func NewMicrosoftGraphEducationAssignment() *MicrosoftGraphEducationAssignment`

NewMicrosoftGraphEducationAssignment instantiates a new MicrosoftGraphEducationAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationAssignmentWithDefaults

`func NewMicrosoftGraphEducationAssignmentWithDefaults() *MicrosoftGraphEducationAssignment`

NewMicrosoftGraphEducationAssignmentWithDefaults instantiates a new MicrosoftGraphEducationAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddedStudentAction

`func (o *MicrosoftGraphEducationAssignment) GetAddedStudentAction() AnyOfmicrosoftGraphEducationAddedStudentAction`

GetAddedStudentAction returns the AddedStudentAction field if non-nil, zero value otherwise.

### GetAddedStudentActionOk

`func (o *MicrosoftGraphEducationAssignment) GetAddedStudentActionOk() (*AnyOfmicrosoftGraphEducationAddedStudentAction, bool)`

GetAddedStudentActionOk returns a tuple with the AddedStudentAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedStudentAction

`func (o *MicrosoftGraphEducationAssignment) SetAddedStudentAction(v AnyOfmicrosoftGraphEducationAddedStudentAction)`

SetAddedStudentAction sets AddedStudentAction field to given value.

### HasAddedStudentAction

`func (o *MicrosoftGraphEducationAssignment) HasAddedStudentAction() bool`

HasAddedStudentAction returns a boolean if a field has been set.

### SetAddedStudentActionNil

`func (o *MicrosoftGraphEducationAssignment) SetAddedStudentActionNil(b bool)`

 SetAddedStudentActionNil sets the value for AddedStudentAction to be an explicit nil

### UnsetAddedStudentAction
`func (o *MicrosoftGraphEducationAssignment) UnsetAddedStudentAction()`

UnsetAddedStudentAction ensures that no value is present for AddedStudentAction, not even an explicit nil
### GetAllowLateSubmissions

`func (o *MicrosoftGraphEducationAssignment) GetAllowLateSubmissions() bool`

GetAllowLateSubmissions returns the AllowLateSubmissions field if non-nil, zero value otherwise.

### GetAllowLateSubmissionsOk

`func (o *MicrosoftGraphEducationAssignment) GetAllowLateSubmissionsOk() (*bool, bool)`

GetAllowLateSubmissionsOk returns a tuple with the AllowLateSubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLateSubmissions

`func (o *MicrosoftGraphEducationAssignment) SetAllowLateSubmissions(v bool)`

SetAllowLateSubmissions sets AllowLateSubmissions field to given value.

### HasAllowLateSubmissions

`func (o *MicrosoftGraphEducationAssignment) HasAllowLateSubmissions() bool`

HasAllowLateSubmissions returns a boolean if a field has been set.

### SetAllowLateSubmissionsNil

`func (o *MicrosoftGraphEducationAssignment) SetAllowLateSubmissionsNil(b bool)`

 SetAllowLateSubmissionsNil sets the value for AllowLateSubmissions to be an explicit nil

### UnsetAllowLateSubmissions
`func (o *MicrosoftGraphEducationAssignment) UnsetAllowLateSubmissions()`

UnsetAllowLateSubmissions ensures that no value is present for AllowLateSubmissions, not even an explicit nil
### GetAllowStudentsToAddResourcesToSubmission

`func (o *MicrosoftGraphEducationAssignment) GetAllowStudentsToAddResourcesToSubmission() bool`

GetAllowStudentsToAddResourcesToSubmission returns the AllowStudentsToAddResourcesToSubmission field if non-nil, zero value otherwise.

### GetAllowStudentsToAddResourcesToSubmissionOk

`func (o *MicrosoftGraphEducationAssignment) GetAllowStudentsToAddResourcesToSubmissionOk() (*bool, bool)`

GetAllowStudentsToAddResourcesToSubmissionOk returns a tuple with the AllowStudentsToAddResourcesToSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStudentsToAddResourcesToSubmission

`func (o *MicrosoftGraphEducationAssignment) SetAllowStudentsToAddResourcesToSubmission(v bool)`

SetAllowStudentsToAddResourcesToSubmission sets AllowStudentsToAddResourcesToSubmission field to given value.

### HasAllowStudentsToAddResourcesToSubmission

`func (o *MicrosoftGraphEducationAssignment) HasAllowStudentsToAddResourcesToSubmission() bool`

HasAllowStudentsToAddResourcesToSubmission returns a boolean if a field has been set.

### SetAllowStudentsToAddResourcesToSubmissionNil

`func (o *MicrosoftGraphEducationAssignment) SetAllowStudentsToAddResourcesToSubmissionNil(b bool)`

 SetAllowStudentsToAddResourcesToSubmissionNil sets the value for AllowStudentsToAddResourcesToSubmission to be an explicit nil

### UnsetAllowStudentsToAddResourcesToSubmission
`func (o *MicrosoftGraphEducationAssignment) UnsetAllowStudentsToAddResourcesToSubmission()`

UnsetAllowStudentsToAddResourcesToSubmission ensures that no value is present for AllowStudentsToAddResourcesToSubmission, not even an explicit nil
### GetAssignDateTime

`func (o *MicrosoftGraphEducationAssignment) GetAssignDateTime() time.Time`

GetAssignDateTime returns the AssignDateTime field if non-nil, zero value otherwise.

### GetAssignDateTimeOk

`func (o *MicrosoftGraphEducationAssignment) GetAssignDateTimeOk() (*time.Time, bool)`

GetAssignDateTimeOk returns a tuple with the AssignDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignDateTime

`func (o *MicrosoftGraphEducationAssignment) SetAssignDateTime(v time.Time)`

SetAssignDateTime sets AssignDateTime field to given value.

### HasAssignDateTime

`func (o *MicrosoftGraphEducationAssignment) HasAssignDateTime() bool`

HasAssignDateTime returns a boolean if a field has been set.

### SetAssignDateTimeNil

`func (o *MicrosoftGraphEducationAssignment) SetAssignDateTimeNil(b bool)`

 SetAssignDateTimeNil sets the value for AssignDateTime to be an explicit nil

### UnsetAssignDateTime
`func (o *MicrosoftGraphEducationAssignment) UnsetAssignDateTime()`

UnsetAssignDateTime ensures that no value is present for AssignDateTime, not even an explicit nil
### GetAssignedDateTime

`func (o *MicrosoftGraphEducationAssignment) GetAssignedDateTime() time.Time`

GetAssignedDateTime returns the AssignedDateTime field if non-nil, zero value otherwise.

### GetAssignedDateTimeOk

`func (o *MicrosoftGraphEducationAssignment) GetAssignedDateTimeOk() (*time.Time, bool)`

GetAssignedDateTimeOk returns a tuple with the AssignedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDateTime

`func (o *MicrosoftGraphEducationAssignment) SetAssignedDateTime(v time.Time)`

SetAssignedDateTime sets AssignedDateTime field to given value.

### HasAssignedDateTime

`func (o *MicrosoftGraphEducationAssignment) HasAssignedDateTime() bool`

HasAssignedDateTime returns a boolean if a field has been set.

### SetAssignedDateTimeNil

`func (o *MicrosoftGraphEducationAssignment) SetAssignedDateTimeNil(b bool)`

 SetAssignedDateTimeNil sets the value for AssignedDateTime to be an explicit nil

### UnsetAssignedDateTime
`func (o *MicrosoftGraphEducationAssignment) UnsetAssignedDateTime()`

UnsetAssignedDateTime ensures that no value is present for AssignedDateTime, not even an explicit nil
### GetAssignTo

`func (o *MicrosoftGraphEducationAssignment) GetAssignTo() AnyOfobject`

GetAssignTo returns the AssignTo field if non-nil, zero value otherwise.

### GetAssignToOk

`func (o *MicrosoftGraphEducationAssignment) GetAssignToOk() (*AnyOfobject, bool)`

GetAssignToOk returns a tuple with the AssignTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignTo

`func (o *MicrosoftGraphEducationAssignment) SetAssignTo(v AnyOfobject)`

SetAssignTo sets AssignTo field to given value.

### HasAssignTo

`func (o *MicrosoftGraphEducationAssignment) HasAssignTo() bool`

HasAssignTo returns a boolean if a field has been set.

### SetAssignToNil

`func (o *MicrosoftGraphEducationAssignment) SetAssignToNil(b bool)`

 SetAssignToNil sets the value for AssignTo to be an explicit nil

### UnsetAssignTo
`func (o *MicrosoftGraphEducationAssignment) UnsetAssignTo()`

UnsetAssignTo ensures that no value is present for AssignTo, not even an explicit nil
### GetClassId

`func (o *MicrosoftGraphEducationAssignment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MicrosoftGraphEducationAssignment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MicrosoftGraphEducationAssignment) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *MicrosoftGraphEducationAssignment) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### SetClassIdNil

`func (o *MicrosoftGraphEducationAssignment) SetClassIdNil(b bool)`

 SetClassIdNil sets the value for ClassId to be an explicit nil

### UnsetClassId
`func (o *MicrosoftGraphEducationAssignment) UnsetClassId()`

UnsetClassId ensures that no value is present for ClassId, not even an explicit nil
### GetCloseDateTime

`func (o *MicrosoftGraphEducationAssignment) GetCloseDateTime() time.Time`

GetCloseDateTime returns the CloseDateTime field if non-nil, zero value otherwise.

### GetCloseDateTimeOk

`func (o *MicrosoftGraphEducationAssignment) GetCloseDateTimeOk() (*time.Time, bool)`

GetCloseDateTimeOk returns a tuple with the CloseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDateTime

`func (o *MicrosoftGraphEducationAssignment) SetCloseDateTime(v time.Time)`

SetCloseDateTime sets CloseDateTime field to given value.

### HasCloseDateTime

`func (o *MicrosoftGraphEducationAssignment) HasCloseDateTime() bool`

HasCloseDateTime returns a boolean if a field has been set.

### SetCloseDateTimeNil

`func (o *MicrosoftGraphEducationAssignment) SetCloseDateTimeNil(b bool)`

 SetCloseDateTimeNil sets the value for CloseDateTime to be an explicit nil

### UnsetCloseDateTime
`func (o *MicrosoftGraphEducationAssignment) UnsetCloseDateTime()`

UnsetCloseDateTime ensures that no value is present for CloseDateTime, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphEducationAssignment) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphEducationAssignment) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphEducationAssignment) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphEducationAssignment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphEducationAssignment) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphEducationAssignment) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphEducationAssignment) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphEducationAssignment) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphEducationAssignment) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphEducationAssignment) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphEducationAssignment) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphEducationAssignment) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphEducationAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphEducationAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphEducationAssignment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphEducationAssignment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDueDateTime

`func (o *MicrosoftGraphEducationAssignment) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *MicrosoftGraphEducationAssignment) GetDueDateTimeOk() (*time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDateTime

`func (o *MicrosoftGraphEducationAssignment) SetDueDateTime(v time.Time)`

SetDueDateTime sets DueDateTime field to given value.

### HasDueDateTime

`func (o *MicrosoftGraphEducationAssignment) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTimeNil

`func (o *MicrosoftGraphEducationAssignment) SetDueDateTimeNil(b bool)`

 SetDueDateTimeNil sets the value for DueDateTime to be an explicit nil

### UnsetDueDateTime
`func (o *MicrosoftGraphEducationAssignment) UnsetDueDateTime()`

UnsetDueDateTime ensures that no value is present for DueDateTime, not even an explicit nil
### GetGrading

`func (o *MicrosoftGraphEducationAssignment) GetGrading() AnyOfobject`

GetGrading returns the Grading field if non-nil, zero value otherwise.

### GetGradingOk

`func (o *MicrosoftGraphEducationAssignment) GetGradingOk() (*AnyOfobject, bool)`

GetGradingOk returns a tuple with the Grading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrading

`func (o *MicrosoftGraphEducationAssignment) SetGrading(v AnyOfobject)`

SetGrading sets Grading field to given value.

### HasGrading

`func (o *MicrosoftGraphEducationAssignment) HasGrading() bool`

HasGrading returns a boolean if a field has been set.

### SetGradingNil

`func (o *MicrosoftGraphEducationAssignment) SetGradingNil(b bool)`

 SetGradingNil sets the value for Grading to be an explicit nil

### UnsetGrading
`func (o *MicrosoftGraphEducationAssignment) UnsetGrading()`

UnsetGrading ensures that no value is present for Grading, not even an explicit nil
### GetInstructions

`func (o *MicrosoftGraphEducationAssignment) GetInstructions() AnyOfmicrosoftGraphEducationItemBody`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *MicrosoftGraphEducationAssignment) GetInstructionsOk() (*AnyOfmicrosoftGraphEducationItemBody, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *MicrosoftGraphEducationAssignment) SetInstructions(v AnyOfmicrosoftGraphEducationItemBody)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *MicrosoftGraphEducationAssignment) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *MicrosoftGraphEducationAssignment) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *MicrosoftGraphEducationAssignment) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphEducationAssignment) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphEducationAssignment) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphEducationAssignment) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphEducationAssignment) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphEducationAssignment) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphEducationAssignment) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphEducationAssignment) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphEducationAssignment) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphEducationAssignment) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphEducationAssignment) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphEducationAssignment) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphEducationAssignment) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetNotificationChannelUrl

`func (o *MicrosoftGraphEducationAssignment) GetNotificationChannelUrl() string`

GetNotificationChannelUrl returns the NotificationChannelUrl field if non-nil, zero value otherwise.

### GetNotificationChannelUrlOk

`func (o *MicrosoftGraphEducationAssignment) GetNotificationChannelUrlOk() (*string, bool)`

GetNotificationChannelUrlOk returns a tuple with the NotificationChannelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannelUrl

`func (o *MicrosoftGraphEducationAssignment) SetNotificationChannelUrl(v string)`

SetNotificationChannelUrl sets NotificationChannelUrl field to given value.

### HasNotificationChannelUrl

`func (o *MicrosoftGraphEducationAssignment) HasNotificationChannelUrl() bool`

HasNotificationChannelUrl returns a boolean if a field has been set.

### SetNotificationChannelUrlNil

`func (o *MicrosoftGraphEducationAssignment) SetNotificationChannelUrlNil(b bool)`

 SetNotificationChannelUrlNil sets the value for NotificationChannelUrl to be an explicit nil

### UnsetNotificationChannelUrl
`func (o *MicrosoftGraphEducationAssignment) UnsetNotificationChannelUrl()`

UnsetNotificationChannelUrl ensures that no value is present for NotificationChannelUrl, not even an explicit nil
### GetResourcesFolderUrl

`func (o *MicrosoftGraphEducationAssignment) GetResourcesFolderUrl() string`

GetResourcesFolderUrl returns the ResourcesFolderUrl field if non-nil, zero value otherwise.

### GetResourcesFolderUrlOk

`func (o *MicrosoftGraphEducationAssignment) GetResourcesFolderUrlOk() (*string, bool)`

GetResourcesFolderUrlOk returns a tuple with the ResourcesFolderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesFolderUrl

`func (o *MicrosoftGraphEducationAssignment) SetResourcesFolderUrl(v string)`

SetResourcesFolderUrl sets ResourcesFolderUrl field to given value.

### HasResourcesFolderUrl

`func (o *MicrosoftGraphEducationAssignment) HasResourcesFolderUrl() bool`

HasResourcesFolderUrl returns a boolean if a field has been set.

### SetResourcesFolderUrlNil

`func (o *MicrosoftGraphEducationAssignment) SetResourcesFolderUrlNil(b bool)`

 SetResourcesFolderUrlNil sets the value for ResourcesFolderUrl to be an explicit nil

### UnsetResourcesFolderUrl
`func (o *MicrosoftGraphEducationAssignment) UnsetResourcesFolderUrl()`

UnsetResourcesFolderUrl ensures that no value is present for ResourcesFolderUrl, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphEducationAssignment) GetStatus() AnyOfmicrosoftGraphEducationAssignmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphEducationAssignment) GetStatusOk() (*AnyOfmicrosoftGraphEducationAssignmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphEducationAssignment) SetStatus(v AnyOfmicrosoftGraphEducationAssignmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphEducationAssignment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphEducationAssignment) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphEducationAssignment) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphEducationAssignment) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphEducationAssignment) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphEducationAssignment) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphEducationAssignment) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphEducationAssignment) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphEducationAssignment) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetCategories

`func (o *MicrosoftGraphEducationAssignment) GetCategories() []MicrosoftGraphEducationCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphEducationAssignment) GetCategoriesOk() (*[]MicrosoftGraphEducationCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphEducationAssignment) SetCategories(v []MicrosoftGraphEducationCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphEducationAssignment) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetResources

`func (o *MicrosoftGraphEducationAssignment) GetResources() []MicrosoftGraphEducationAssignmentResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MicrosoftGraphEducationAssignment) GetResourcesOk() (*[]MicrosoftGraphEducationAssignmentResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MicrosoftGraphEducationAssignment) SetResources(v []MicrosoftGraphEducationAssignmentResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MicrosoftGraphEducationAssignment) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRubric

`func (o *MicrosoftGraphEducationAssignment) GetRubric() AnyOfmicrosoftGraphEducationRubric`

GetRubric returns the Rubric field if non-nil, zero value otherwise.

### GetRubricOk

`func (o *MicrosoftGraphEducationAssignment) GetRubricOk() (*AnyOfmicrosoftGraphEducationRubric, bool)`

GetRubricOk returns a tuple with the Rubric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubric

`func (o *MicrosoftGraphEducationAssignment) SetRubric(v AnyOfmicrosoftGraphEducationRubric)`

SetRubric sets Rubric field to given value.

### HasRubric

`func (o *MicrosoftGraphEducationAssignment) HasRubric() bool`

HasRubric returns a boolean if a field has been set.

### SetRubricNil

`func (o *MicrosoftGraphEducationAssignment) SetRubricNil(b bool)`

 SetRubricNil sets the value for Rubric to be an explicit nil

### UnsetRubric
`func (o *MicrosoftGraphEducationAssignment) UnsetRubric()`

UnsetRubric ensures that no value is present for Rubric, not even an explicit nil
### GetSubmissions

`func (o *MicrosoftGraphEducationAssignment) GetSubmissions() []MicrosoftGraphEducationSubmission`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *MicrosoftGraphEducationAssignment) GetSubmissionsOk() (*[]MicrosoftGraphEducationSubmission, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *MicrosoftGraphEducationAssignment) SetSubmissions(v []MicrosoftGraphEducationSubmission)`

SetSubmissions sets Submissions field to given value.

### HasSubmissions

`func (o *MicrosoftGraphEducationAssignment) HasSubmissions() bool`

HasSubmissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


