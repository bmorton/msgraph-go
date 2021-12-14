# MicrosoftGraphOnenoteOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The start time of the operation. | [optional] 
**LastActionDateTime** | Pointer to **NullableTime** | The time of the last action of the operation. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) | The current status of the operation: notStarted, running, completed, failed | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphOnenoteOperationError**](anyOf&lt;microsoft.graph.onenoteOperationError&gt;.md) | The error returned by the operation. | [optional] 
**PercentComplete** | Pointer to **NullableString** | The operation percent complete if the operation is still in running status. | [optional] 
**ResourceId** | Pointer to **NullableString** | The resource id. | [optional] 
**ResourceLocation** | Pointer to **NullableString** | The resource URI for the object. For example, the resource URI for a copied page or section. | [optional] 

## Methods

### NewMicrosoftGraphOnenoteOperation

`func NewMicrosoftGraphOnenoteOperation() *MicrosoftGraphOnenoteOperation`

NewMicrosoftGraphOnenoteOperation instantiates a new MicrosoftGraphOnenoteOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnenoteOperationWithDefaults

`func NewMicrosoftGraphOnenoteOperationWithDefaults() *MicrosoftGraphOnenoteOperation`

NewMicrosoftGraphOnenoteOperationWithDefaults instantiates a new MicrosoftGraphOnenoteOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOnenoteOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenoteOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOnenoteOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOnenoteOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenoteOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOnenoteOperation) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastActionDateTime

`func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTimeOk() (*time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionDateTime

`func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime sets LastActionDateTime field to given value.

### HasLastActionDateTime

`func (o *MicrosoftGraphOnenoteOperation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTimeNil

`func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTimeNil(b bool)`

 SetLastActionDateTimeNil sets the value for LastActionDateTime to be an explicit nil

### UnsetLastActionDateTime
`func (o *MicrosoftGraphOnenoteOperation) UnsetLastActionDateTime()`

UnsetLastActionDateTime ensures that no value is present for LastActionDateTime, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphOnenoteOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphOnenoteOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphOnenoteOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphOnenoteOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphOnenoteOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphOnenoteOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetError

`func (o *MicrosoftGraphOnenoteOperation) GetError() AnyOfmicrosoftGraphOnenoteOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphOnenoteOperation) GetErrorOk() (*AnyOfmicrosoftGraphOnenoteOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphOnenoteOperation) SetError(v AnyOfmicrosoftGraphOnenoteOperationError)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphOnenoteOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphOnenoteOperation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphOnenoteOperation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPercentComplete

`func (o *MicrosoftGraphOnenoteOperation) GetPercentComplete() string`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MicrosoftGraphOnenoteOperation) GetPercentCompleteOk() (*string, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *MicrosoftGraphOnenoteOperation) SetPercentComplete(v string)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *MicrosoftGraphOnenoteOperation) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *MicrosoftGraphOnenoteOperation) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *MicrosoftGraphOnenoteOperation) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetResourceId

`func (o *MicrosoftGraphOnenoteOperation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphOnenoteOperation) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MicrosoftGraphOnenoteOperation) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MicrosoftGraphOnenoteOperation) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *MicrosoftGraphOnenoteOperation) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *MicrosoftGraphOnenoteOperation) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceLocation

`func (o *MicrosoftGraphOnenoteOperation) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *MicrosoftGraphOnenoteOperation) GetResourceLocationOk() (*string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *MicrosoftGraphOnenoteOperation) SetResourceLocation(v string)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *MicrosoftGraphOnenoteOperation) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### SetResourceLocationNil

`func (o *MicrosoftGraphOnenoteOperation) SetResourceLocationNil(b bool)`

 SetResourceLocationNil sets the value for ResourceLocation to be an explicit nil

### UnsetResourceLocation
`func (o *MicrosoftGraphOnenoteOperation) UnsetResourceLocation()`

UnsetResourceLocation ensures that no value is present for ResourceLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


