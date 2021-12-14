# MicrosoftGraphCallTranscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModifiedDateTime** | Pointer to **NullableTime** | The state modified time in UTC. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphCallTranscriptionState**](anyOf&lt;microsoft.graph.callTranscriptionState&gt;.md) | Possible values are: notStarted, active, inactive. | [optional] 

## Methods

### NewMicrosoftGraphCallTranscriptionInfo

`func NewMicrosoftGraphCallTranscriptionInfo() *MicrosoftGraphCallTranscriptionInfo`

NewMicrosoftGraphCallTranscriptionInfo instantiates a new MicrosoftGraphCallTranscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallTranscriptionInfoWithDefaults

`func NewMicrosoftGraphCallTranscriptionInfoWithDefaults() *MicrosoftGraphCallTranscriptionInfo`

NewMicrosoftGraphCallTranscriptionInfoWithDefaults instantiates a new MicrosoftGraphCallTranscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedDateTime

`func (o *MicrosoftGraphCallTranscriptionInfo) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphCallTranscriptionInfo) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphCallTranscriptionInfo) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphCallTranscriptionInfo) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphCallTranscriptionInfo) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphCallTranscriptionInfo) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetState

`func (o *MicrosoftGraphCallTranscriptionInfo) GetState() AnyOfmicrosoftGraphCallTranscriptionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphCallTranscriptionInfo) GetStateOk() (*AnyOfmicrosoftGraphCallTranscriptionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphCallTranscriptionInfo) SetState(v AnyOfmicrosoftGraphCallTranscriptionState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphCallTranscriptionInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphCallTranscriptionInfo) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphCallTranscriptionInfo) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


