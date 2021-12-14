# MicrosoftGraphServiceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Service** | Pointer to **string** | The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphServiceHealthStatus**](anyOf&lt;microsoft.graph.serviceHealthStatus&gt;.md) | Show the overral service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. | [optional] 
**Issues** | Pointer to [**[]MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | A collection of issues happened on the service, with detailed information for each issue. | [optional] 

## Methods

### NewMicrosoftGraphServiceHealth

`func NewMicrosoftGraphServiceHealth() *MicrosoftGraphServiceHealth`

NewMicrosoftGraphServiceHealth instantiates a new MicrosoftGraphServiceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServiceHealthWithDefaults

`func NewMicrosoftGraphServiceHealthWithDefaults() *MicrosoftGraphServiceHealth`

NewMicrosoftGraphServiceHealthWithDefaults instantiates a new MicrosoftGraphServiceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphServiceHealth) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphServiceHealth) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphServiceHealth) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphServiceHealth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetService

`func (o *MicrosoftGraphServiceHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MicrosoftGraphServiceHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MicrosoftGraphServiceHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MicrosoftGraphServiceHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftGraphServiceHealth) GetStatus() AnyOfmicrosoftGraphServiceHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphServiceHealth) GetStatusOk() (*AnyOfmicrosoftGraphServiceHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphServiceHealth) SetStatus(v AnyOfmicrosoftGraphServiceHealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphServiceHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphServiceHealth) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphServiceHealth) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIssues

`func (o *MicrosoftGraphServiceHealth) GetIssues() []MicrosoftGraphServiceHealthIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *MicrosoftGraphServiceHealth) GetIssuesOk() (*[]MicrosoftGraphServiceHealthIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *MicrosoftGraphServiceHealth) SetIssues(v []MicrosoftGraphServiceHealthIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *MicrosoftGraphServiceHealth) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


