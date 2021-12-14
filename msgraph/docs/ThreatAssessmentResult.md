# ThreatAssessmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Message** | Pointer to **NullableString** | The result message for each threat assessment. | [optional] 
**ResultType** | Pointer to [**AnyOfmicrosoftGraphThreatAssessmentResultType**](anyOf&lt;microsoft.graph.threatAssessmentResultType&gt;.md) | The threat assessment result type. Possible values are: checkPolicy, rescan. | [optional] 

## Methods

### NewThreatAssessmentResult

`func NewThreatAssessmentResult() *ThreatAssessmentResult`

NewThreatAssessmentResult instantiates a new ThreatAssessmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatAssessmentResultWithDefaults

`func NewThreatAssessmentResultWithDefaults() *ThreatAssessmentResult`

NewThreatAssessmentResultWithDefaults instantiates a new ThreatAssessmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *ThreatAssessmentResult) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ThreatAssessmentResult) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ThreatAssessmentResult) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ThreatAssessmentResult) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *ThreatAssessmentResult) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *ThreatAssessmentResult) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetMessage

`func (o *ThreatAssessmentResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ThreatAssessmentResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ThreatAssessmentResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ThreatAssessmentResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ThreatAssessmentResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ThreatAssessmentResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetResultType

`func (o *ThreatAssessmentResult) GetResultType() AnyOfmicrosoftGraphThreatAssessmentResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ThreatAssessmentResult) GetResultTypeOk() (*AnyOfmicrosoftGraphThreatAssessmentResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ThreatAssessmentResult) SetResultType(v AnyOfmicrosoftGraphThreatAssessmentResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ThreatAssessmentResult) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *ThreatAssessmentResult) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *ThreatAssessmentResult) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


