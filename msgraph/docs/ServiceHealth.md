# ServiceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphServiceHealthStatus**](anyOf&lt;microsoft.graph.serviceHealthStatus&gt;.md) | Show the overral service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. | [optional] 
**Issues** | Pointer to [**[]MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | A collection of issues happened on the service, with detailed information for each issue. | [optional] 

## Methods

### NewServiceHealth

`func NewServiceHealth() *ServiceHealth`

NewServiceHealth instantiates a new ServiceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHealthWithDefaults

`func NewServiceHealthWithDefaults() *ServiceHealth`

NewServiceHealthWithDefaults instantiates a new ServiceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceHealth) GetStatus() AnyOfmicrosoftGraphServiceHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceHealth) GetStatusOk() (*AnyOfmicrosoftGraphServiceHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceHealth) SetStatus(v AnyOfmicrosoftGraphServiceHealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServiceHealth) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServiceHealth) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIssues

`func (o *ServiceHealth) GetIssues() []MicrosoftGraphServiceHealthIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ServiceHealth) GetIssuesOk() (*[]MicrosoftGraphServiceHealthIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ServiceHealth) SetIssues(v []MicrosoftGraphServiceHealthIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ServiceHealth) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


