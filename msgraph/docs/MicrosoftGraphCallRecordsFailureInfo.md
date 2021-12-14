# MicrosoftGraphCallRecordsFailureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** | Classification of why a call or portion of a call failed. | [optional] 
**Stage** | Pointer to [**AnyOfmicrosoftGraphCallRecordsFailureStage**](anyOf&lt;microsoft.graph.callRecords.failureStage&gt;.md) | The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphCallRecordsFailureInfo

`func NewMicrosoftGraphCallRecordsFailureInfo() *MicrosoftGraphCallRecordsFailureInfo`

NewMicrosoftGraphCallRecordsFailureInfo instantiates a new MicrosoftGraphCallRecordsFailureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsFailureInfoWithDefaults

`func NewMicrosoftGraphCallRecordsFailureInfoWithDefaults() *MicrosoftGraphCallRecordsFailureInfo`

NewMicrosoftGraphCallRecordsFailureInfoWithDefaults instantiates a new MicrosoftGraphCallRecordsFailureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *MicrosoftGraphCallRecordsFailureInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MicrosoftGraphCallRecordsFailureInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MicrosoftGraphCallRecordsFailureInfo) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MicrosoftGraphCallRecordsFailureInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *MicrosoftGraphCallRecordsFailureInfo) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MicrosoftGraphCallRecordsFailureInfo) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetStage

`func (o *MicrosoftGraphCallRecordsFailureInfo) GetStage() AnyOfmicrosoftGraphCallRecordsFailureStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MicrosoftGraphCallRecordsFailureInfo) GetStageOk() (*AnyOfmicrosoftGraphCallRecordsFailureStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MicrosoftGraphCallRecordsFailureInfo) SetStage(v AnyOfmicrosoftGraphCallRecordsFailureStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *MicrosoftGraphCallRecordsFailureInfo) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *MicrosoftGraphCallRecordsFailureInfo) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *MicrosoftGraphCallRecordsFailureInfo) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


