# MicrosoftGraphUpdateRecordingStatusOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClientContext** | Pointer to **NullableString** | Unique Client Context string. Max limit is 256 chars. | [optional] 
**ResultInfo** | Pointer to [**AnyOfmicrosoftGraphResultInfo**](anyOf&lt;microsoft.graph.resultInfo&gt;.md) | The result information. Read-only. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) | Possible values are: notStarted, running, completed, failed. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphUpdateRecordingStatusOperation

`func NewMicrosoftGraphUpdateRecordingStatusOperation() *MicrosoftGraphUpdateRecordingStatusOperation`

NewMicrosoftGraphUpdateRecordingStatusOperation instantiates a new MicrosoftGraphUpdateRecordingStatusOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUpdateRecordingStatusOperationWithDefaults

`func NewMicrosoftGraphUpdateRecordingStatusOperationWithDefaults() *MicrosoftGraphUpdateRecordingStatusOperation`

NewMicrosoftGraphUpdateRecordingStatusOperationWithDefaults instantiates a new MicrosoftGraphUpdateRecordingStatusOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientContext

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetClientContext() string`

GetClientContext returns the ClientContext field if non-nil, zero value otherwise.

### GetClientContextOk

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetClientContextOk() (*string, bool)`

GetClientContextOk returns a tuple with the ClientContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContext

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetClientContext(v string)`

SetClientContext sets ClientContext field to given value.

### HasClientContext

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasClientContext() bool`

HasClientContext returns a boolean if a field has been set.

### SetClientContextNil

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetClientContextNil(b bool)`

 SetClientContextNil sets the value for ClientContext to be an explicit nil

### UnsetClientContext
`func (o *MicrosoftGraphUpdateRecordingStatusOperation) UnsetClientContext()`

UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
### GetResultInfo

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetResultInfo() AnyOfmicrosoftGraphResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetResultInfo(v AnyOfmicrosoftGraphResultInfo)`

SetResultInfo sets ResultInfo field to given value.

### HasResultInfo

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasResultInfo() bool`

HasResultInfo returns a boolean if a field has been set.

### SetResultInfoNil

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetResultInfoNil(b bool)`

 SetResultInfoNil sets the value for ResultInfo to be an explicit nil

### UnsetResultInfo
`func (o *MicrosoftGraphUpdateRecordingStatusOperation) UnsetResultInfo()`

UnsetResultInfo ensures that no value is present for ResultInfo, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphUpdateRecordingStatusOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphUpdateRecordingStatusOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


