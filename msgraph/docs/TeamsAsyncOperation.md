# TeamsAsyncOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptsCount** | Pointer to **int32** | Number of times the operation was attempted before being marked successful or failed. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Time when the operation was created. | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphOperationError**](anyOf&lt;microsoft.graph.operationError&gt;.md) | Any error that causes the async operation to fail. | [optional] 
**LastActionDateTime** | Pointer to **time.Time** | Time when the async operation was last updated. | [optional] 
**OperationType** | Pointer to [**AnyOfmicrosoftGraphTeamsAsyncOperationType**](anyOf&lt;microsoft.graph.teamsAsyncOperationType&gt;.md) | Denotes which type of operation is being described. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphTeamsAsyncOperationStatus**](anyOf&lt;microsoft.graph.teamsAsyncOperationStatus&gt;.md) | Operation status. | [optional] 
**TargetResourceId** | Pointer to **NullableString** | The ID of the object that&#39;s created or modified as result of this async operation, typically a team. | [optional] 
**TargetResourceLocation** | Pointer to **NullableString** | The location of the object that&#39;s created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths. | [optional] 

## Methods

### NewTeamsAsyncOperation

`func NewTeamsAsyncOperation() *TeamsAsyncOperation`

NewTeamsAsyncOperation instantiates a new TeamsAsyncOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsAsyncOperationWithDefaults

`func NewTeamsAsyncOperationWithDefaults() *TeamsAsyncOperation`

NewTeamsAsyncOperationWithDefaults instantiates a new TeamsAsyncOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptsCount

`func (o *TeamsAsyncOperation) GetAttemptsCount() int32`

GetAttemptsCount returns the AttemptsCount field if non-nil, zero value otherwise.

### GetAttemptsCountOk

`func (o *TeamsAsyncOperation) GetAttemptsCountOk() (*int32, bool)`

GetAttemptsCountOk returns a tuple with the AttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsCount

`func (o *TeamsAsyncOperation) SetAttemptsCount(v int32)`

SetAttemptsCount sets AttemptsCount field to given value.

### HasAttemptsCount

`func (o *TeamsAsyncOperation) HasAttemptsCount() bool`

HasAttemptsCount returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *TeamsAsyncOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *TeamsAsyncOperation) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *TeamsAsyncOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *TeamsAsyncOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetError

`func (o *TeamsAsyncOperation) GetError() AnyOfmicrosoftGraphOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TeamsAsyncOperation) GetErrorOk() (*AnyOfmicrosoftGraphOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TeamsAsyncOperation) SetError(v AnyOfmicrosoftGraphOperationError)`

SetError sets Error field to given value.

### HasError

`func (o *TeamsAsyncOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *TeamsAsyncOperation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *TeamsAsyncOperation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetLastActionDateTime

`func (o *TeamsAsyncOperation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *TeamsAsyncOperation) GetLastActionDateTimeOk() (*time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionDateTime

`func (o *TeamsAsyncOperation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime sets LastActionDateTime field to given value.

### HasLastActionDateTime

`func (o *TeamsAsyncOperation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### GetOperationType

`func (o *TeamsAsyncOperation) GetOperationType() AnyOfmicrosoftGraphTeamsAsyncOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TeamsAsyncOperation) GetOperationTypeOk() (*AnyOfmicrosoftGraphTeamsAsyncOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TeamsAsyncOperation) SetOperationType(v AnyOfmicrosoftGraphTeamsAsyncOperationType)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *TeamsAsyncOperation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *TeamsAsyncOperation) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *TeamsAsyncOperation) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetStatus

`func (o *TeamsAsyncOperation) GetStatus() AnyOfmicrosoftGraphTeamsAsyncOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TeamsAsyncOperation) GetStatusOk() (*AnyOfmicrosoftGraphTeamsAsyncOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TeamsAsyncOperation) SetStatus(v AnyOfmicrosoftGraphTeamsAsyncOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TeamsAsyncOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TeamsAsyncOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TeamsAsyncOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTargetResourceId

`func (o *TeamsAsyncOperation) GetTargetResourceId() string`

GetTargetResourceId returns the TargetResourceId field if non-nil, zero value otherwise.

### GetTargetResourceIdOk

`func (o *TeamsAsyncOperation) GetTargetResourceIdOk() (*string, bool)`

GetTargetResourceIdOk returns a tuple with the TargetResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceId

`func (o *TeamsAsyncOperation) SetTargetResourceId(v string)`

SetTargetResourceId sets TargetResourceId field to given value.

### HasTargetResourceId

`func (o *TeamsAsyncOperation) HasTargetResourceId() bool`

HasTargetResourceId returns a boolean if a field has been set.

### SetTargetResourceIdNil

`func (o *TeamsAsyncOperation) SetTargetResourceIdNil(b bool)`

 SetTargetResourceIdNil sets the value for TargetResourceId to be an explicit nil

### UnsetTargetResourceId
`func (o *TeamsAsyncOperation) UnsetTargetResourceId()`

UnsetTargetResourceId ensures that no value is present for TargetResourceId, not even an explicit nil
### GetTargetResourceLocation

`func (o *TeamsAsyncOperation) GetTargetResourceLocation() string`

GetTargetResourceLocation returns the TargetResourceLocation field if non-nil, zero value otherwise.

### GetTargetResourceLocationOk

`func (o *TeamsAsyncOperation) GetTargetResourceLocationOk() (*string, bool)`

GetTargetResourceLocationOk returns a tuple with the TargetResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceLocation

`func (o *TeamsAsyncOperation) SetTargetResourceLocation(v string)`

SetTargetResourceLocation sets TargetResourceLocation field to given value.

### HasTargetResourceLocation

`func (o *TeamsAsyncOperation) HasTargetResourceLocation() bool`

HasTargetResourceLocation returns a boolean if a field has been set.

### SetTargetResourceLocationNil

`func (o *TeamsAsyncOperation) SetTargetResourceLocationNil(b bool)`

 SetTargetResourceLocationNil sets the value for TargetResourceLocation to be an explicit nil

### UnsetTargetResourceLocation
`func (o *TeamsAsyncOperation) UnsetTargetResourceLocation()`

UnsetTargetResourceLocation ensures that no value is present for TargetResourceLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


