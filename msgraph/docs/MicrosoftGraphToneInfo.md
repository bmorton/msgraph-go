# MicrosoftGraphToneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceId** | Pointer to **int64** | An incremental identifier used for ordering DTMF events. | [optional] 
**Tone** | Pointer to [**AnyOfmicrosoftGraphTone**](anyOf&lt;microsoft.graph.tone&gt;.md) | Possible values are: tone0, tone1, tone2, tone3, tone4, tone5, tone6, tone7, tone8, tone9, star, pound, a, b, c, d, flash. | [optional] 

## Methods

### NewMicrosoftGraphToneInfo

`func NewMicrosoftGraphToneInfo() *MicrosoftGraphToneInfo`

NewMicrosoftGraphToneInfo instantiates a new MicrosoftGraphToneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphToneInfoWithDefaults

`func NewMicrosoftGraphToneInfoWithDefaults() *MicrosoftGraphToneInfo`

NewMicrosoftGraphToneInfoWithDefaults instantiates a new MicrosoftGraphToneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceId

`func (o *MicrosoftGraphToneInfo) GetSequenceId() int64`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *MicrosoftGraphToneInfo) GetSequenceIdOk() (*int64, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *MicrosoftGraphToneInfo) SetSequenceId(v int64)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *MicrosoftGraphToneInfo) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetTone

`func (o *MicrosoftGraphToneInfo) GetTone() AnyOfmicrosoftGraphTone`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *MicrosoftGraphToneInfo) GetToneOk() (*AnyOfmicrosoftGraphTone, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *MicrosoftGraphToneInfo) SetTone(v AnyOfmicrosoftGraphTone)`

SetTone sets Tone field to given value.

### HasTone

`func (o *MicrosoftGraphToneInfo) HasTone() bool`

HasTone returns a boolean if a field has been set.

### SetToneNil

`func (o *MicrosoftGraphToneInfo) SetToneNil(b bool)`

 SetToneNil sets the value for Tone to be an explicit nil

### UnsetTone
`func (o *MicrosoftGraphToneInfo) UnsetTone()`

UnsetTone ensures that no value is present for Tone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


