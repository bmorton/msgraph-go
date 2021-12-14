# MicrosoftGraphDomainState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastActionDateTime** | Pointer to **NullableTime** | Timestamp for when the last activity occurred. The value is updated when an operation is scheduled, the asynchronous task starts, and when the operation completes. | [optional] 
**Operation** | Pointer to **NullableString** | Type of asynchronous operation. The values can be ForceDelete or Verification | [optional] 
**Status** | Pointer to **NullableString** | Current status of the operation.  Scheduled - Operation has been scheduled but has not started.  InProgress - Task has started and is in progress.  Failed - Operation has failed. | [optional] 

## Methods

### NewMicrosoftGraphDomainState

`func NewMicrosoftGraphDomainState() *MicrosoftGraphDomainState`

NewMicrosoftGraphDomainState instantiates a new MicrosoftGraphDomainState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDomainStateWithDefaults

`func NewMicrosoftGraphDomainStateWithDefaults() *MicrosoftGraphDomainState`

NewMicrosoftGraphDomainStateWithDefaults instantiates a new MicrosoftGraphDomainState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastActionDateTime

`func (o *MicrosoftGraphDomainState) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *MicrosoftGraphDomainState) GetLastActionDateTimeOk() (*time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionDateTime

`func (o *MicrosoftGraphDomainState) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime sets LastActionDateTime field to given value.

### HasLastActionDateTime

`func (o *MicrosoftGraphDomainState) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTimeNil

`func (o *MicrosoftGraphDomainState) SetLastActionDateTimeNil(b bool)`

 SetLastActionDateTimeNil sets the value for LastActionDateTime to be an explicit nil

### UnsetLastActionDateTime
`func (o *MicrosoftGraphDomainState) UnsetLastActionDateTime()`

UnsetLastActionDateTime ensures that no value is present for LastActionDateTime, not even an explicit nil
### GetOperation

`func (o *MicrosoftGraphDomainState) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MicrosoftGraphDomainState) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MicrosoftGraphDomainState) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MicrosoftGraphDomainState) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *MicrosoftGraphDomainState) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *MicrosoftGraphDomainState) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphDomainState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDomainState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphDomainState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphDomainState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphDomainState) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphDomainState) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


