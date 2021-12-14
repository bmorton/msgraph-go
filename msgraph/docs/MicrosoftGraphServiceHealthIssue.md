# MicrosoftGraphServiceHealthIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) | Additional details about service event. This property doesn&#39;t support filters. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The end time of the service event. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The last modified time of the service event. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start time of the service event. | [optional] 
**Title** | Pointer to **string** | The title of the service event. | [optional] 
**Classification** | Pointer to [**AnyOfmicrosoftGraphServiceHealthClassificationType**](anyOf&lt;microsoft.graph.serviceHealthClassificationType&gt;.md) | The type of service health issue. Possible values are: advisory, incident, unknownFutureValue. | [optional] 
**Feature** | Pointer to **NullableString** | The feature name of the service issue. | [optional] 
**FeatureGroup** | Pointer to **NullableString** | The feature group name of the service issue. | [optional] 
**ImpactDescription** | Pointer to **string** | The description of the service issue impact. | [optional] 
**IsResolved** | Pointer to **bool** | Indicates whether the issue is resolved. | [optional] 
**Origin** | Pointer to [**AnyOfmicrosoftGraphServiceHealthOrigin**](anyOf&lt;microsoft.graph.serviceHealthOrigin&gt;.md) | Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue. | [optional] 
**Posts** | Pointer to [**[]MicrosoftGraphServiceHealthIssuePost**](MicrosoftGraphServiceHealthIssuePost.md) | Collection of historical posts for the service issue. | [optional] 
**Service** | Pointer to **string** | Indicates the service affected by the issue. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphServiceHealthStatus**](anyOf&lt;microsoft.graph.serviceHealthStatus&gt;.md) | The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphServiceHealthIssue

`func NewMicrosoftGraphServiceHealthIssue() *MicrosoftGraphServiceHealthIssue`

NewMicrosoftGraphServiceHealthIssue instantiates a new MicrosoftGraphServiceHealthIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServiceHealthIssueWithDefaults

`func NewMicrosoftGraphServiceHealthIssueWithDefaults() *MicrosoftGraphServiceHealthIssue`

NewMicrosoftGraphServiceHealthIssueWithDefaults instantiates a new MicrosoftGraphServiceHealthIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphServiceHealthIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphServiceHealthIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphServiceHealthIssue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphServiceHealthIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDetails

`func (o *MicrosoftGraphServiceHealthIssue) GetDetails() []*AnyOfmicrosoftGraphKeyValuePair`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphServiceHealthIssue) GetDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphServiceHealthIssue) SetDetails(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphServiceHealthIssue) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MicrosoftGraphServiceHealthIssue) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphServiceHealthIssue) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphServiceHealthIssue) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphServiceHealthIssue) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphServiceHealthIssue) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphServiceHealthIssue) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphServiceHealthIssue) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphServiceHealthIssue) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphServiceHealthIssue) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphServiceHealthIssue) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphServiceHealthIssue) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphServiceHealthIssue) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphServiceHealthIssue) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphServiceHealthIssue) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetTitle

`func (o *MicrosoftGraphServiceHealthIssue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphServiceHealthIssue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphServiceHealthIssue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphServiceHealthIssue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetClassification

`func (o *MicrosoftGraphServiceHealthIssue) GetClassification() AnyOfmicrosoftGraphServiceHealthClassificationType`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *MicrosoftGraphServiceHealthIssue) GetClassificationOk() (*AnyOfmicrosoftGraphServiceHealthClassificationType, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *MicrosoftGraphServiceHealthIssue) SetClassification(v AnyOfmicrosoftGraphServiceHealthClassificationType)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *MicrosoftGraphServiceHealthIssue) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *MicrosoftGraphServiceHealthIssue) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *MicrosoftGraphServiceHealthIssue) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetFeature

`func (o *MicrosoftGraphServiceHealthIssue) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *MicrosoftGraphServiceHealthIssue) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *MicrosoftGraphServiceHealthIssue) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *MicrosoftGraphServiceHealthIssue) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### SetFeatureNil

`func (o *MicrosoftGraphServiceHealthIssue) SetFeatureNil(b bool)`

 SetFeatureNil sets the value for Feature to be an explicit nil

### UnsetFeature
`func (o *MicrosoftGraphServiceHealthIssue) UnsetFeature()`

UnsetFeature ensures that no value is present for Feature, not even an explicit nil
### GetFeatureGroup

`func (o *MicrosoftGraphServiceHealthIssue) GetFeatureGroup() string`

GetFeatureGroup returns the FeatureGroup field if non-nil, zero value otherwise.

### GetFeatureGroupOk

`func (o *MicrosoftGraphServiceHealthIssue) GetFeatureGroupOk() (*string, bool)`

GetFeatureGroupOk returns a tuple with the FeatureGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureGroup

`func (o *MicrosoftGraphServiceHealthIssue) SetFeatureGroup(v string)`

SetFeatureGroup sets FeatureGroup field to given value.

### HasFeatureGroup

`func (o *MicrosoftGraphServiceHealthIssue) HasFeatureGroup() bool`

HasFeatureGroup returns a boolean if a field has been set.

### SetFeatureGroupNil

`func (o *MicrosoftGraphServiceHealthIssue) SetFeatureGroupNil(b bool)`

 SetFeatureGroupNil sets the value for FeatureGroup to be an explicit nil

### UnsetFeatureGroup
`func (o *MicrosoftGraphServiceHealthIssue) UnsetFeatureGroup()`

UnsetFeatureGroup ensures that no value is present for FeatureGroup, not even an explicit nil
### GetImpactDescription

`func (o *MicrosoftGraphServiceHealthIssue) GetImpactDescription() string`

GetImpactDescription returns the ImpactDescription field if non-nil, zero value otherwise.

### GetImpactDescriptionOk

`func (o *MicrosoftGraphServiceHealthIssue) GetImpactDescriptionOk() (*string, bool)`

GetImpactDescriptionOk returns a tuple with the ImpactDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactDescription

`func (o *MicrosoftGraphServiceHealthIssue) SetImpactDescription(v string)`

SetImpactDescription sets ImpactDescription field to given value.

### HasImpactDescription

`func (o *MicrosoftGraphServiceHealthIssue) HasImpactDescription() bool`

HasImpactDescription returns a boolean if a field has been set.

### GetIsResolved

`func (o *MicrosoftGraphServiceHealthIssue) GetIsResolved() bool`

GetIsResolved returns the IsResolved field if non-nil, zero value otherwise.

### GetIsResolvedOk

`func (o *MicrosoftGraphServiceHealthIssue) GetIsResolvedOk() (*bool, bool)`

GetIsResolvedOk returns a tuple with the IsResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolved

`func (o *MicrosoftGraphServiceHealthIssue) SetIsResolved(v bool)`

SetIsResolved sets IsResolved field to given value.

### HasIsResolved

`func (o *MicrosoftGraphServiceHealthIssue) HasIsResolved() bool`

HasIsResolved returns a boolean if a field has been set.

### GetOrigin

`func (o *MicrosoftGraphServiceHealthIssue) GetOrigin() AnyOfmicrosoftGraphServiceHealthOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MicrosoftGraphServiceHealthIssue) GetOriginOk() (*AnyOfmicrosoftGraphServiceHealthOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MicrosoftGraphServiceHealthIssue) SetOrigin(v AnyOfmicrosoftGraphServiceHealthOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MicrosoftGraphServiceHealthIssue) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *MicrosoftGraphServiceHealthIssue) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *MicrosoftGraphServiceHealthIssue) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPosts

`func (o *MicrosoftGraphServiceHealthIssue) GetPosts() []MicrosoftGraphServiceHealthIssuePost`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *MicrosoftGraphServiceHealthIssue) GetPostsOk() (*[]MicrosoftGraphServiceHealthIssuePost, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *MicrosoftGraphServiceHealthIssue) SetPosts(v []MicrosoftGraphServiceHealthIssuePost)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *MicrosoftGraphServiceHealthIssue) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetService

`func (o *MicrosoftGraphServiceHealthIssue) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MicrosoftGraphServiceHealthIssue) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MicrosoftGraphServiceHealthIssue) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MicrosoftGraphServiceHealthIssue) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphServiceHealthIssue) GetStatus() AnyOfmicrosoftGraphServiceHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphServiceHealthIssue) GetStatusOk() (*AnyOfmicrosoftGraphServiceHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphServiceHealthIssue) SetStatus(v AnyOfmicrosoftGraphServiceHealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphServiceHealthIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphServiceHealthIssue) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphServiceHealthIssue) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


