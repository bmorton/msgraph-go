# ServiceUpdateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionRequiredByDateTime** | Pointer to **NullableTime** | The expected deadline of the action for the message. | [optional] 
**Body** | Pointer to [**MicrosoftGraphItemBody**](MicrosoftGraphItemBody.md) |  | [optional] 
**Category** | Pointer to [**AnyOfmicrosoftGraphServiceUpdateCategory**](anyOf&lt;microsoft.graph.serviceUpdateCategory&gt;.md) | The service message category. Possible values are: preventOrFixIssue, planForChange, stayInformed, unknownFutureValue. | [optional] 
**IsMajorChange** | Pointer to **NullableBool** | Indicates whether the message describes a major update for the service. | [optional] 
**Services** | Pointer to **[]string** | The affected services by the service message. | [optional] 
**Severity** | Pointer to [**AnyOfmicrosoftGraphServiceUpdateSeverity**](anyOf&lt;microsoft.graph.serviceUpdateSeverity&gt;.md) | The severity of the service message. Possible values are: normal, high, critical, unknownFutureValue. | [optional] 
**Tags** | Pointer to **[]string** | A collection of tags for the service message. | [optional] 
**ViewPoint** | Pointer to [**AnyOfmicrosoftGraphServiceUpdateMessageViewpoint**](anyOf&lt;microsoft.graph.serviceUpdateMessageViewpoint&gt;.md) | Represents user view points data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions. | [optional] 

## Methods

### NewServiceUpdateMessage

`func NewServiceUpdateMessage() *ServiceUpdateMessage`

NewServiceUpdateMessage instantiates a new ServiceUpdateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUpdateMessageWithDefaults

`func NewServiceUpdateMessageWithDefaults() *ServiceUpdateMessage`

NewServiceUpdateMessageWithDefaults instantiates a new ServiceUpdateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionRequiredByDateTime

`func (o *ServiceUpdateMessage) GetActionRequiredByDateTime() time.Time`

GetActionRequiredByDateTime returns the ActionRequiredByDateTime field if non-nil, zero value otherwise.

### GetActionRequiredByDateTimeOk

`func (o *ServiceUpdateMessage) GetActionRequiredByDateTimeOk() (*time.Time, bool)`

GetActionRequiredByDateTimeOk returns a tuple with the ActionRequiredByDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRequiredByDateTime

`func (o *ServiceUpdateMessage) SetActionRequiredByDateTime(v time.Time)`

SetActionRequiredByDateTime sets ActionRequiredByDateTime field to given value.

### HasActionRequiredByDateTime

`func (o *ServiceUpdateMessage) HasActionRequiredByDateTime() bool`

HasActionRequiredByDateTime returns a boolean if a field has been set.

### SetActionRequiredByDateTimeNil

`func (o *ServiceUpdateMessage) SetActionRequiredByDateTimeNil(b bool)`

 SetActionRequiredByDateTimeNil sets the value for ActionRequiredByDateTime to be an explicit nil

### UnsetActionRequiredByDateTime
`func (o *ServiceUpdateMessage) UnsetActionRequiredByDateTime()`

UnsetActionRequiredByDateTime ensures that no value is present for ActionRequiredByDateTime, not even an explicit nil
### GetBody

`func (o *ServiceUpdateMessage) GetBody() MicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ServiceUpdateMessage) GetBodyOk() (*MicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ServiceUpdateMessage) SetBody(v MicrosoftGraphItemBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *ServiceUpdateMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *ServiceUpdateMessage) GetCategory() AnyOfmicrosoftGraphServiceUpdateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ServiceUpdateMessage) GetCategoryOk() (*AnyOfmicrosoftGraphServiceUpdateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ServiceUpdateMessage) SetCategory(v AnyOfmicrosoftGraphServiceUpdateCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ServiceUpdateMessage) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ServiceUpdateMessage) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ServiceUpdateMessage) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsMajorChange

`func (o *ServiceUpdateMessage) GetIsMajorChange() bool`

GetIsMajorChange returns the IsMajorChange field if non-nil, zero value otherwise.

### GetIsMajorChangeOk

`func (o *ServiceUpdateMessage) GetIsMajorChangeOk() (*bool, bool)`

GetIsMajorChangeOk returns a tuple with the IsMajorChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorChange

`func (o *ServiceUpdateMessage) SetIsMajorChange(v bool)`

SetIsMajorChange sets IsMajorChange field to given value.

### HasIsMajorChange

`func (o *ServiceUpdateMessage) HasIsMajorChange() bool`

HasIsMajorChange returns a boolean if a field has been set.

### SetIsMajorChangeNil

`func (o *ServiceUpdateMessage) SetIsMajorChangeNil(b bool)`

 SetIsMajorChangeNil sets the value for IsMajorChange to be an explicit nil

### UnsetIsMajorChange
`func (o *ServiceUpdateMessage) UnsetIsMajorChange()`

UnsetIsMajorChange ensures that no value is present for IsMajorChange, not even an explicit nil
### GetServices

`func (o *ServiceUpdateMessage) GetServices() []*string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceUpdateMessage) GetServicesOk() (*[]*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceUpdateMessage) SetServices(v []*string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServiceUpdateMessage) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSeverity

`func (o *ServiceUpdateMessage) GetSeverity() AnyOfmicrosoftGraphServiceUpdateSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ServiceUpdateMessage) GetSeverityOk() (*AnyOfmicrosoftGraphServiceUpdateSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ServiceUpdateMessage) SetSeverity(v AnyOfmicrosoftGraphServiceUpdateSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ServiceUpdateMessage) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *ServiceUpdateMessage) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *ServiceUpdateMessage) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetTags

`func (o *ServiceUpdateMessage) GetTags() []*string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceUpdateMessage) GetTagsOk() (*[]*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceUpdateMessage) SetTags(v []*string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceUpdateMessage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewPoint

`func (o *ServiceUpdateMessage) GetViewPoint() AnyOfmicrosoftGraphServiceUpdateMessageViewpoint`

GetViewPoint returns the ViewPoint field if non-nil, zero value otherwise.

### GetViewPointOk

`func (o *ServiceUpdateMessage) GetViewPointOk() (*AnyOfmicrosoftGraphServiceUpdateMessageViewpoint, bool)`

GetViewPointOk returns a tuple with the ViewPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPoint

`func (o *ServiceUpdateMessage) SetViewPoint(v AnyOfmicrosoftGraphServiceUpdateMessageViewpoint)`

SetViewPoint sets ViewPoint field to given value.

### HasViewPoint

`func (o *ServiceUpdateMessage) HasViewPoint() bool`

HasViewPoint returns a boolean if a field has been set.

### SetViewPointNil

`func (o *ServiceUpdateMessage) SetViewPointNil(b bool)`

 SetViewPointNil sets the value for ViewPoint to be an explicit nil

### UnsetViewPoint
`func (o *ServiceUpdateMessage) UnsetViewPoint()`

UnsetViewPoint ensures that no value is present for ViewPoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


