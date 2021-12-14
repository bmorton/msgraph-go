# MicrosoftGraphServiceUpdateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) | Additional details about service event. This property doesn&#39;t support filters. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The end time of the service event. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The last modified time of the service event. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start time of the service event. | [optional] 
**Title** | Pointer to **string** | The title of the service event. | [optional] 
**ActionRequiredByDateTime** | Pointer to **NullableTime** | The expected deadline of the action for the message. | [optional] 
**Body** | Pointer to [**MicrosoftGraphItemBody**](MicrosoftGraphItemBody.md) |  | [optional] 
**Category** | Pointer to [**AnyOfmicrosoftGraphServiceUpdateCategory**](anyOf&lt;microsoft.graph.serviceUpdateCategory&gt;.md) | The service message category. Possible values are: preventOrFixIssue, planForChange, stayInformed, unknownFutureValue. | [optional] 
**IsMajorChange** | Pointer to **NullableBool** | Indicates whether the message describes a major update for the service. | [optional] 
**Services** | Pointer to **[]string** | The affected services by the service message. | [optional] 
**Severity** | Pointer to [**AnyOfmicrosoftGraphServiceUpdateSeverity**](anyOf&lt;microsoft.graph.serviceUpdateSeverity&gt;.md) | The severity of the service message. Possible values are: normal, high, critical, unknownFutureValue. | [optional] 
**Tags** | Pointer to **[]string** | A collection of tags for the service message. | [optional] 
**ViewPoint** | Pointer to [**AnyOfmicrosoftGraphServiceUpdateMessageViewpoint**](anyOf&lt;microsoft.graph.serviceUpdateMessageViewpoint&gt;.md) | Represents user view points data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions. | [optional] 

## Methods

### NewMicrosoftGraphServiceUpdateMessage

`func NewMicrosoftGraphServiceUpdateMessage() *MicrosoftGraphServiceUpdateMessage`

NewMicrosoftGraphServiceUpdateMessage instantiates a new MicrosoftGraphServiceUpdateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServiceUpdateMessageWithDefaults

`func NewMicrosoftGraphServiceUpdateMessageWithDefaults() *MicrosoftGraphServiceUpdateMessage`

NewMicrosoftGraphServiceUpdateMessageWithDefaults instantiates a new MicrosoftGraphServiceUpdateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphServiceUpdateMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphServiceUpdateMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphServiceUpdateMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDetails

`func (o *MicrosoftGraphServiceUpdateMessage) GetDetails() []*AnyOfmicrosoftGraphKeyValuePair`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphServiceUpdateMessage) SetDetails(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphServiceUpdateMessage) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphServiceUpdateMessage) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphServiceUpdateMessage) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetTitle

`func (o *MicrosoftGraphServiceUpdateMessage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphServiceUpdateMessage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphServiceUpdateMessage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetActionRequiredByDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) GetActionRequiredByDateTime() time.Time`

GetActionRequiredByDateTime returns the ActionRequiredByDateTime field if non-nil, zero value otherwise.

### GetActionRequiredByDateTimeOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetActionRequiredByDateTimeOk() (*time.Time, bool)`

GetActionRequiredByDateTimeOk returns a tuple with the ActionRequiredByDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRequiredByDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) SetActionRequiredByDateTime(v time.Time)`

SetActionRequiredByDateTime sets ActionRequiredByDateTime field to given value.

### HasActionRequiredByDateTime

`func (o *MicrosoftGraphServiceUpdateMessage) HasActionRequiredByDateTime() bool`

HasActionRequiredByDateTime returns a boolean if a field has been set.

### SetActionRequiredByDateTimeNil

`func (o *MicrosoftGraphServiceUpdateMessage) SetActionRequiredByDateTimeNil(b bool)`

 SetActionRequiredByDateTimeNil sets the value for ActionRequiredByDateTime to be an explicit nil

### UnsetActionRequiredByDateTime
`func (o *MicrosoftGraphServiceUpdateMessage) UnsetActionRequiredByDateTime()`

UnsetActionRequiredByDateTime ensures that no value is present for ActionRequiredByDateTime, not even an explicit nil
### GetBody

`func (o *MicrosoftGraphServiceUpdateMessage) GetBody() MicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetBodyOk() (*MicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MicrosoftGraphServiceUpdateMessage) SetBody(v MicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *MicrosoftGraphServiceUpdateMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *MicrosoftGraphServiceUpdateMessage) GetCategory() AnyOfmicrosoftGraphServiceUpdateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetCategoryOk() (*AnyOfmicrosoftGraphServiceUpdateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MicrosoftGraphServiceUpdateMessage) SetCategory(v AnyOfmicrosoftGraphServiceUpdateCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MicrosoftGraphServiceUpdateMessage) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MicrosoftGraphServiceUpdateMessage) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MicrosoftGraphServiceUpdateMessage) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsMajorChange

`func (o *MicrosoftGraphServiceUpdateMessage) GetIsMajorChange() bool`

GetIsMajorChange returns the IsMajorChange field if non-nil, zero value otherwise.

### GetIsMajorChangeOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetIsMajorChangeOk() (*bool, bool)`

GetIsMajorChangeOk returns a tuple with the IsMajorChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorChange

`func (o *MicrosoftGraphServiceUpdateMessage) SetIsMajorChange(v bool)`

SetIsMajorChange sets IsMajorChange field to given value.

### HasIsMajorChange

`func (o *MicrosoftGraphServiceUpdateMessage) HasIsMajorChange() bool`

HasIsMajorChange returns a boolean if a field has been set.

### SetIsMajorChangeNil

`func (o *MicrosoftGraphServiceUpdateMessage) SetIsMajorChangeNil(b bool)`

 SetIsMajorChangeNil sets the value for IsMajorChange to be an explicit nil

### UnsetIsMajorChange
`func (o *MicrosoftGraphServiceUpdateMessage) UnsetIsMajorChange()`

UnsetIsMajorChange ensures that no value is present for IsMajorChange, not even an explicit nil
### GetServices

`func (o *MicrosoftGraphServiceUpdateMessage) GetServices() []*string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetServicesOk() (*[]*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MicrosoftGraphServiceUpdateMessage) SetServices(v []*string)`

SetServices sets Services field to given value.

### HasServices

`func (o *MicrosoftGraphServiceUpdateMessage) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSeverity

`func (o *MicrosoftGraphServiceUpdateMessage) GetSeverity() AnyOfmicrosoftGraphServiceUpdateSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetSeverityOk() (*AnyOfmicrosoftGraphServiceUpdateSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MicrosoftGraphServiceUpdateMessage) SetSeverity(v AnyOfmicrosoftGraphServiceUpdateSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MicrosoftGraphServiceUpdateMessage) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *MicrosoftGraphServiceUpdateMessage) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *MicrosoftGraphServiceUpdateMessage) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetTags

`func (o *MicrosoftGraphServiceUpdateMessage) GetTags() []*string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetTagsOk() (*[]*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MicrosoftGraphServiceUpdateMessage) SetTags(v []*string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MicrosoftGraphServiceUpdateMessage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewPoint

`func (o *MicrosoftGraphServiceUpdateMessage) GetViewPoint() AnyOfmicrosoftGraphServiceUpdateMessageViewpoint`

GetViewPoint returns the ViewPoint field if non-nil, zero value otherwise.

### GetViewPointOk

`func (o *MicrosoftGraphServiceUpdateMessage) GetViewPointOk() (*AnyOfmicrosoftGraphServiceUpdateMessageViewpoint, bool)`

GetViewPointOk returns a tuple with the ViewPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPoint

`func (o *MicrosoftGraphServiceUpdateMessage) SetViewPoint(v AnyOfmicrosoftGraphServiceUpdateMessageViewpoint)`

SetViewPoint sets ViewPoint field to given value.

### HasViewPoint

`func (o *MicrosoftGraphServiceUpdateMessage) HasViewPoint() bool`

HasViewPoint returns a boolean if a field has been set.

### SetViewPointNil

`func (o *MicrosoftGraphServiceUpdateMessage) SetViewPointNil(b bool)`

 SetViewPointNil sets the value for ViewPoint to be an explicit nil

### UnsetViewPoint
`func (o *MicrosoftGraphServiceUpdateMessage) UnsetViewPoint()`

UnsetViewPoint ensures that no value is present for ViewPoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


