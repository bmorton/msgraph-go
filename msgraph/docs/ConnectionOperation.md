# ConnectionOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AnyOfmicrosoftGraphPublicError**](anyOf&lt;microsoft.graph.publicError&gt;.md) | If status is failed, provides more information about the error that caused the failure. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus**](anyOf&lt;microsoft.graph.externalConnectors.connectionOperationStatus&gt;.md) | Indicates the status of the asynchronous operation. Possible values are: unspecified, inprogress, completed, failed, unknownFutureValue. | [optional] 

## Methods

### NewConnectionOperation

`func NewConnectionOperation() *ConnectionOperation`

NewConnectionOperation instantiates a new ConnectionOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOperationWithDefaults

`func NewConnectionOperationWithDefaults() *ConnectionOperation`

NewConnectionOperationWithDefaults instantiates a new ConnectionOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ConnectionOperation) GetError() AnyOfmicrosoftGraphPublicError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectionOperation) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectionOperation) SetError(v AnyOfmicrosoftGraphPublicError)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectionOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConnectionOperation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConnectionOperation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStatus

`func (o *ConnectionOperation) GetStatus() AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionOperation) GetStatusOk() (*AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionOperation) SetStatus(v AnyOfmicrosoftGraphExternalConnectorsConnectionOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ConnectionOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ConnectionOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


