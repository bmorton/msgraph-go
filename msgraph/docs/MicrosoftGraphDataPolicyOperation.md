# MicrosoftGraphDataPolicyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CompletedDateTime** | Pointer to **NullableTime** | Represents when the request for this data policy operation was completed, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Null until the operation completes. | [optional] 
**Progress** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Specifies the progress of an operation. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphDataPolicyOperationStatus**](anyOf&lt;microsoft.graph.dataPolicyOperationStatus&gt;.md) | Possible values are: notStarted, running, complete, failed, unknownFutureValue. | [optional] 
**StorageLocation** | Pointer to **NullableString** | The URL location to where data is being exported for export requests. | [optional] 
**SubmittedDateTime** | Pointer to **time.Time** | Represents when the request for this data operation was submitted, in UTC time, using the ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**UserId** | Pointer to **string** | The id for the user on whom the operation is performed. | [optional] 

## Methods

### NewMicrosoftGraphDataPolicyOperation

`func NewMicrosoftGraphDataPolicyOperation() *MicrosoftGraphDataPolicyOperation`

NewMicrosoftGraphDataPolicyOperation instantiates a new MicrosoftGraphDataPolicyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDataPolicyOperationWithDefaults

`func NewMicrosoftGraphDataPolicyOperationWithDefaults() *MicrosoftGraphDataPolicyOperation`

NewMicrosoftGraphDataPolicyOperationWithDefaults instantiates a new MicrosoftGraphDataPolicyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDataPolicyOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDataPolicyOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDataPolicyOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDataPolicyOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompletedDateTime

`func (o *MicrosoftGraphDataPolicyOperation) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *MicrosoftGraphDataPolicyOperation) GetCompletedDateTimeOk() (*time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *MicrosoftGraphDataPolicyOperation) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *MicrosoftGraphDataPolicyOperation) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *MicrosoftGraphDataPolicyOperation) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *MicrosoftGraphDataPolicyOperation) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetProgress

`func (o *MicrosoftGraphDataPolicyOperation) GetProgress() AnyOfnumberstringstring`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *MicrosoftGraphDataPolicyOperation) GetProgressOk() (*AnyOfnumberstringstring, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *MicrosoftGraphDataPolicyOperation) SetProgress(v AnyOfnumberstringstring)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *MicrosoftGraphDataPolicyOperation) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *MicrosoftGraphDataPolicyOperation) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *MicrosoftGraphDataPolicyOperation) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphDataPolicyOperation) GetStatus() AnyOfmicrosoftGraphDataPolicyOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDataPolicyOperation) GetStatusOk() (*AnyOfmicrosoftGraphDataPolicyOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphDataPolicyOperation) SetStatus(v AnyOfmicrosoftGraphDataPolicyOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphDataPolicyOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphDataPolicyOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphDataPolicyOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStorageLocation

`func (o *MicrosoftGraphDataPolicyOperation) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *MicrosoftGraphDataPolicyOperation) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *MicrosoftGraphDataPolicyOperation) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *MicrosoftGraphDataPolicyOperation) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### SetStorageLocationNil

`func (o *MicrosoftGraphDataPolicyOperation) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *MicrosoftGraphDataPolicyOperation) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
### GetSubmittedDateTime

`func (o *MicrosoftGraphDataPolicyOperation) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *MicrosoftGraphDataPolicyOperation) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *MicrosoftGraphDataPolicyOperation) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.

### HasSubmittedDateTime

`func (o *MicrosoftGraphDataPolicyOperation) HasSubmittedDateTime() bool`

HasSubmittedDateTime returns a boolean if a field has been set.

### GetUserId

`func (o *MicrosoftGraphDataPolicyOperation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphDataPolicyOperation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphDataPolicyOperation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphDataPolicyOperation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


