# CommsOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientContext** | Pointer to **NullableString** | Unique Client Context string. Max limit is 256 chars. | [optional] 
**ResultInfo** | Pointer to [**AnyOfmicrosoftGraphResultInfo**](anyOf&lt;microsoft.graph.resultInfo&gt;.md) | The result information. Read-only. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) | Possible values are: notStarted, running, completed, failed. Read-only. | [optional] 

## Methods

### NewCommsOperation

`func NewCommsOperation() *CommsOperation`

NewCommsOperation instantiates a new CommsOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommsOperationWithDefaults

`func NewCommsOperationWithDefaults() *CommsOperation`

NewCommsOperationWithDefaults instantiates a new CommsOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientContext

`func (o *CommsOperation) GetClientContext() string`

GetClientContext returns the ClientContext field if non-nil, zero value otherwise.

### GetClientContextOk

`func (o *CommsOperation) GetClientContextOk() (*string, bool)`

GetClientContextOk returns a tuple with the ClientContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContext

`func (o *CommsOperation) SetClientContext(v string)`

SetClientContext sets ClientContext field to given value.

### HasClientContext

`func (o *CommsOperation) HasClientContext() bool`

HasClientContext returns a boolean if a field has been set.

### SetClientContextNil

`func (o *CommsOperation) SetClientContextNil(b bool)`

 SetClientContextNil sets the value for ClientContext to be an explicit nil

### UnsetClientContext
`func (o *CommsOperation) UnsetClientContext()`

UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
### GetResultInfo

`func (o *CommsOperation) GetResultInfo() AnyOfmicrosoftGraphResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *CommsOperation) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *CommsOperation) SetResultInfo(v AnyOfmicrosoftGraphResultInfo)`

SetResultInfo sets ResultInfo field to given value.

### HasResultInfo

`func (o *CommsOperation) HasResultInfo() bool`

HasResultInfo returns a boolean if a field has been set.

### SetResultInfoNil

`func (o *CommsOperation) SetResultInfoNil(b bool)`

 SetResultInfoNil sets the value for ResultInfo to be an explicit nil

### UnsetResultInfo
`func (o *CommsOperation) UnsetResultInfo()`

UnsetResultInfo ensures that no value is present for ResultInfo, not even an explicit nil
### GetStatus

`func (o *CommsOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommsOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommsOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommsOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CommsOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CommsOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


