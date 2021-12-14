# MicrosoftGraphAlertHistoryState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableString** | The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be extracted from the auth token and not entered manually by the calling application. | [optional] 
**AssignedTo** | Pointer to **NullableString** | UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN). | [optional] 
**Comments** | Pointer to **[]string** | Comment entered by signed-in user. | [optional] 
**Feedback** | Pointer to [**AnyOfmicrosoftGraphAlertFeedback**](anyOf&lt;microsoft.graph.alertFeedback&gt;.md) | Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive, benignPositive. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphAlertStatus**](anyOf&lt;microsoft.graph.alertStatus&gt;.md) | Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed. | [optional] 
**UpdatedDateTime** | Pointer to **NullableTime** | Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**User** | Pointer to **NullableString** | UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode). | [optional] 

## Methods

### NewMicrosoftGraphAlertHistoryState

`func NewMicrosoftGraphAlertHistoryState() *MicrosoftGraphAlertHistoryState`

NewMicrosoftGraphAlertHistoryState instantiates a new MicrosoftGraphAlertHistoryState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAlertHistoryStateWithDefaults

`func NewMicrosoftGraphAlertHistoryStateWithDefaults() *MicrosoftGraphAlertHistoryState`

NewMicrosoftGraphAlertHistoryStateWithDefaults instantiates a new MicrosoftGraphAlertHistoryState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *MicrosoftGraphAlertHistoryState) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphAlertHistoryState) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphAlertHistoryState) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphAlertHistoryState) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *MicrosoftGraphAlertHistoryState) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *MicrosoftGraphAlertHistoryState) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAssignedTo

`func (o *MicrosoftGraphAlertHistoryState) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphAlertHistoryState) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *MicrosoftGraphAlertHistoryState) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *MicrosoftGraphAlertHistoryState) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *MicrosoftGraphAlertHistoryState) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *MicrosoftGraphAlertHistoryState) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetComments

`func (o *MicrosoftGraphAlertHistoryState) GetComments() []*string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MicrosoftGraphAlertHistoryState) GetCommentsOk() (*[]*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *MicrosoftGraphAlertHistoryState) SetComments(v []*string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *MicrosoftGraphAlertHistoryState) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetFeedback

`func (o *MicrosoftGraphAlertHistoryState) GetFeedback() AnyOfmicrosoftGraphAlertFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *MicrosoftGraphAlertHistoryState) GetFeedbackOk() (*AnyOfmicrosoftGraphAlertFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *MicrosoftGraphAlertHistoryState) SetFeedback(v AnyOfmicrosoftGraphAlertFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *MicrosoftGraphAlertHistoryState) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### SetFeedbackNil

`func (o *MicrosoftGraphAlertHistoryState) SetFeedbackNil(b bool)`

 SetFeedbackNil sets the value for Feedback to be an explicit nil

### UnsetFeedback
`func (o *MicrosoftGraphAlertHistoryState) UnsetFeedback()`

UnsetFeedback ensures that no value is present for Feedback, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphAlertHistoryState) GetStatus() AnyOfmicrosoftGraphAlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAlertHistoryState) GetStatusOk() (*AnyOfmicrosoftGraphAlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphAlertHistoryState) SetStatus(v AnyOfmicrosoftGraphAlertStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphAlertHistoryState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphAlertHistoryState) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphAlertHistoryState) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUpdatedDateTime

`func (o *MicrosoftGraphAlertHistoryState) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *MicrosoftGraphAlertHistoryState) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *MicrosoftGraphAlertHistoryState) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *MicrosoftGraphAlertHistoryState) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### SetUpdatedDateTimeNil

`func (o *MicrosoftGraphAlertHistoryState) SetUpdatedDateTimeNil(b bool)`

 SetUpdatedDateTimeNil sets the value for UpdatedDateTime to be an explicit nil

### UnsetUpdatedDateTime
`func (o *MicrosoftGraphAlertHistoryState) UnsetUpdatedDateTime()`

UnsetUpdatedDateTime ensures that no value is present for UpdatedDateTime, not even an explicit nil
### GetUser

`func (o *MicrosoftGraphAlertHistoryState) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphAlertHistoryState) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftGraphAlertHistoryState) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MicrosoftGraphAlertHistoryState) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MicrosoftGraphAlertHistoryState) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MicrosoftGraphAlertHistoryState) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


