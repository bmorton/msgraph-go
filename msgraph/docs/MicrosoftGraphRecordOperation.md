# MicrosoftGraphRecordOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClientContext** | Pointer to **NullableString** | Unique Client Context string. Max limit is 256 chars. | [optional] 
**ResultInfo** | Pointer to [**AnyOfmicrosoftGraphResultInfo**](anyOf&lt;microsoft.graph.resultInfo&gt;.md) | The result information. Read-only. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) | Possible values are: notStarted, running, completed, failed. Read-only. | [optional] 
**RecordingAccessToken** | Pointer to **NullableString** | The access token required to retrieve the recording. | [optional] 
**RecordingLocation** | Pointer to **NullableString** | The location where the recording is located. | [optional] 

## Methods

### NewMicrosoftGraphRecordOperation

`func NewMicrosoftGraphRecordOperation() *MicrosoftGraphRecordOperation`

NewMicrosoftGraphRecordOperation instantiates a new MicrosoftGraphRecordOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRecordOperationWithDefaults

`func NewMicrosoftGraphRecordOperationWithDefaults() *MicrosoftGraphRecordOperation`

NewMicrosoftGraphRecordOperationWithDefaults instantiates a new MicrosoftGraphRecordOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphRecordOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRecordOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphRecordOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphRecordOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientContext

`func (o *MicrosoftGraphRecordOperation) GetClientContext() string`

GetClientContext returns the ClientContext field if non-nil, zero value otherwise.

### GetClientContextOk

`func (o *MicrosoftGraphRecordOperation) GetClientContextOk() (*string, bool)`

GetClientContextOk returns a tuple with the ClientContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContext

`func (o *MicrosoftGraphRecordOperation) SetClientContext(v string)`

SetClientContext sets ClientContext field to given value.

### HasClientContext

`func (o *MicrosoftGraphRecordOperation) HasClientContext() bool`

HasClientContext returns a boolean if a field has been set.

### SetClientContextNil

`func (o *MicrosoftGraphRecordOperation) SetClientContextNil(b bool)`

 SetClientContextNil sets the value for ClientContext to be an explicit nil

### UnsetClientContext
`func (o *MicrosoftGraphRecordOperation) UnsetClientContext()`

UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
### GetResultInfo

`func (o *MicrosoftGraphRecordOperation) GetResultInfo() AnyOfmicrosoftGraphResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *MicrosoftGraphRecordOperation) GetResultInfoOk() (*AnyOfmicrosoftGraphResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *MicrosoftGraphRecordOperation) SetResultInfo(v AnyOfmicrosoftGraphResultInfo)`

SetResultInfo sets ResultInfo field to given value.

### HasResultInfo

`func (o *MicrosoftGraphRecordOperation) HasResultInfo() bool`

HasResultInfo returns a boolean if a field has been set.

### SetResultInfoNil

`func (o *MicrosoftGraphRecordOperation) SetResultInfoNil(b bool)`

 SetResultInfoNil sets the value for ResultInfo to be an explicit nil

### UnsetResultInfo
`func (o *MicrosoftGraphRecordOperation) UnsetResultInfo()`

UnsetResultInfo ensures that no value is present for ResultInfo, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphRecordOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphRecordOperation) GetStatusOk() (*AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphRecordOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphRecordOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphRecordOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphRecordOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRecordingAccessToken

`func (o *MicrosoftGraphRecordOperation) GetRecordingAccessToken() string`

GetRecordingAccessToken returns the RecordingAccessToken field if non-nil, zero value otherwise.

### GetRecordingAccessTokenOk

`func (o *MicrosoftGraphRecordOperation) GetRecordingAccessTokenOk() (*string, bool)`

GetRecordingAccessTokenOk returns a tuple with the RecordingAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingAccessToken

`func (o *MicrosoftGraphRecordOperation) SetRecordingAccessToken(v string)`

SetRecordingAccessToken sets RecordingAccessToken field to given value.

### HasRecordingAccessToken

`func (o *MicrosoftGraphRecordOperation) HasRecordingAccessToken() bool`

HasRecordingAccessToken returns a boolean if a field has been set.

### SetRecordingAccessTokenNil

`func (o *MicrosoftGraphRecordOperation) SetRecordingAccessTokenNil(b bool)`

 SetRecordingAccessTokenNil sets the value for RecordingAccessToken to be an explicit nil

### UnsetRecordingAccessToken
`func (o *MicrosoftGraphRecordOperation) UnsetRecordingAccessToken()`

UnsetRecordingAccessToken ensures that no value is present for RecordingAccessToken, not even an explicit nil
### GetRecordingLocation

`func (o *MicrosoftGraphRecordOperation) GetRecordingLocation() string`

GetRecordingLocation returns the RecordingLocation field if non-nil, zero value otherwise.

### GetRecordingLocationOk

`func (o *MicrosoftGraphRecordOperation) GetRecordingLocationOk() (*string, bool)`

GetRecordingLocationOk returns a tuple with the RecordingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingLocation

`func (o *MicrosoftGraphRecordOperation) SetRecordingLocation(v string)`

SetRecordingLocation sets RecordingLocation field to given value.

### HasRecordingLocation

`func (o *MicrosoftGraphRecordOperation) HasRecordingLocation() bool`

HasRecordingLocation returns a boolean if a field has been set.

### SetRecordingLocationNil

`func (o *MicrosoftGraphRecordOperation) SetRecordingLocationNil(b bool)`

 SetRecordingLocationNil sets the value for RecordingLocation to be an explicit nil

### UnsetRecordingLocation
`func (o *MicrosoftGraphRecordOperation) UnsetRecordingLocation()`

UnsetRecordingLocation ensures that no value is present for RecordingLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


