# InlineObject26

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompts** | Pointer to [**[]AnyOfobject**](AnyOfobject.md) |  | [optional] 
**BargeInAllowed** | Pointer to **NullableBool** |  | [optional] [default to false]
**InitialSilenceTimeoutInSeconds** | Pointer to **NullableInt32** |  | [optional] 
**MaxSilenceTimeoutInSeconds** | Pointer to **NullableInt32** |  | [optional] 
**MaxRecordDurationInSeconds** | Pointer to **NullableInt32** |  | [optional] 
**PlayBeep** | Pointer to **NullableBool** |  | [optional] [default to false]
**StopTones** | Pointer to **[]string** |  | [optional] 
**ClientContext** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject26

`func NewInlineObject26() *InlineObject26`

NewInlineObject26 instantiates a new InlineObject26 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject26WithDefaults

`func NewInlineObject26WithDefaults() *InlineObject26`

NewInlineObject26WithDefaults instantiates a new InlineObject26 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompts

`func (o *InlineObject26) GetPrompts() []*AnyOfobject`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *InlineObject26) GetPromptsOk() (*[]*AnyOfobject, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *InlineObject26) SetPrompts(v []*AnyOfobject)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *InlineObject26) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetBargeInAllowed

`func (o *InlineObject26) GetBargeInAllowed() bool`

GetBargeInAllowed returns the BargeInAllowed field if non-nil, zero value otherwise.

### GetBargeInAllowedOk

`func (o *InlineObject26) GetBargeInAllowedOk() (*bool, bool)`

GetBargeInAllowedOk returns a tuple with the BargeInAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBargeInAllowed

`func (o *InlineObject26) SetBargeInAllowed(v bool)`

SetBargeInAllowed sets BargeInAllowed field to given value.

### HasBargeInAllowed

`func (o *InlineObject26) HasBargeInAllowed() bool`

HasBargeInAllowed returns a boolean if a field has been set.

### SetBargeInAllowedNil

`func (o *InlineObject26) SetBargeInAllowedNil(b bool)`

 SetBargeInAllowedNil sets the value for BargeInAllowed to be an explicit nil

### UnsetBargeInAllowed
`func (o *InlineObject26) UnsetBargeInAllowed()`

UnsetBargeInAllowed ensures that no value is present for BargeInAllowed, not even an explicit nil
### GetInitialSilenceTimeoutInSeconds

`func (o *InlineObject26) GetInitialSilenceTimeoutInSeconds() int32`

GetInitialSilenceTimeoutInSeconds returns the InitialSilenceTimeoutInSeconds field if non-nil, zero value otherwise.

### GetInitialSilenceTimeoutInSecondsOk

`func (o *InlineObject26) GetInitialSilenceTimeoutInSecondsOk() (*int32, bool)`

GetInitialSilenceTimeoutInSecondsOk returns a tuple with the InitialSilenceTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSilenceTimeoutInSeconds

`func (o *InlineObject26) SetInitialSilenceTimeoutInSeconds(v int32)`

SetInitialSilenceTimeoutInSeconds sets InitialSilenceTimeoutInSeconds field to given value.

### HasInitialSilenceTimeoutInSeconds

`func (o *InlineObject26) HasInitialSilenceTimeoutInSeconds() bool`

HasInitialSilenceTimeoutInSeconds returns a boolean if a field has been set.

### SetInitialSilenceTimeoutInSecondsNil

`func (o *InlineObject26) SetInitialSilenceTimeoutInSecondsNil(b bool)`

 SetInitialSilenceTimeoutInSecondsNil sets the value for InitialSilenceTimeoutInSeconds to be an explicit nil

### UnsetInitialSilenceTimeoutInSeconds
`func (o *InlineObject26) UnsetInitialSilenceTimeoutInSeconds()`

UnsetInitialSilenceTimeoutInSeconds ensures that no value is present for InitialSilenceTimeoutInSeconds, not even an explicit nil
### GetMaxSilenceTimeoutInSeconds

`func (o *InlineObject26) GetMaxSilenceTimeoutInSeconds() int32`

GetMaxSilenceTimeoutInSeconds returns the MaxSilenceTimeoutInSeconds field if non-nil, zero value otherwise.

### GetMaxSilenceTimeoutInSecondsOk

`func (o *InlineObject26) GetMaxSilenceTimeoutInSecondsOk() (*int32, bool)`

GetMaxSilenceTimeoutInSecondsOk returns a tuple with the MaxSilenceTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSilenceTimeoutInSeconds

`func (o *InlineObject26) SetMaxSilenceTimeoutInSeconds(v int32)`

SetMaxSilenceTimeoutInSeconds sets MaxSilenceTimeoutInSeconds field to given value.

### HasMaxSilenceTimeoutInSeconds

`func (o *InlineObject26) HasMaxSilenceTimeoutInSeconds() bool`

HasMaxSilenceTimeoutInSeconds returns a boolean if a field has been set.

### SetMaxSilenceTimeoutInSecondsNil

`func (o *InlineObject26) SetMaxSilenceTimeoutInSecondsNil(b bool)`

 SetMaxSilenceTimeoutInSecondsNil sets the value for MaxSilenceTimeoutInSeconds to be an explicit nil

### UnsetMaxSilenceTimeoutInSeconds
`func (o *InlineObject26) UnsetMaxSilenceTimeoutInSeconds()`

UnsetMaxSilenceTimeoutInSeconds ensures that no value is present for MaxSilenceTimeoutInSeconds, not even an explicit nil
### GetMaxRecordDurationInSeconds

`func (o *InlineObject26) GetMaxRecordDurationInSeconds() int32`

GetMaxRecordDurationInSeconds returns the MaxRecordDurationInSeconds field if non-nil, zero value otherwise.

### GetMaxRecordDurationInSecondsOk

`func (o *InlineObject26) GetMaxRecordDurationInSecondsOk() (*int32, bool)`

GetMaxRecordDurationInSecondsOk returns a tuple with the MaxRecordDurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordDurationInSeconds

`func (o *InlineObject26) SetMaxRecordDurationInSeconds(v int32)`

SetMaxRecordDurationInSeconds sets MaxRecordDurationInSeconds field to given value.

### HasMaxRecordDurationInSeconds

`func (o *InlineObject26) HasMaxRecordDurationInSeconds() bool`

HasMaxRecordDurationInSeconds returns a boolean if a field has been set.

### SetMaxRecordDurationInSecondsNil

`func (o *InlineObject26) SetMaxRecordDurationInSecondsNil(b bool)`

 SetMaxRecordDurationInSecondsNil sets the value for MaxRecordDurationInSeconds to be an explicit nil

### UnsetMaxRecordDurationInSeconds
`func (o *InlineObject26) UnsetMaxRecordDurationInSeconds()`

UnsetMaxRecordDurationInSeconds ensures that no value is present for MaxRecordDurationInSeconds, not even an explicit nil
### GetPlayBeep

`func (o *InlineObject26) GetPlayBeep() bool`

GetPlayBeep returns the PlayBeep field if non-nil, zero value otherwise.

### GetPlayBeepOk

`func (o *InlineObject26) GetPlayBeepOk() (*bool, bool)`

GetPlayBeepOk returns a tuple with the PlayBeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayBeep

`func (o *InlineObject26) SetPlayBeep(v bool)`

SetPlayBeep sets PlayBeep field to given value.

### HasPlayBeep

`func (o *InlineObject26) HasPlayBeep() bool`

HasPlayBeep returns a boolean if a field has been set.

### SetPlayBeepNil

`func (o *InlineObject26) SetPlayBeepNil(b bool)`

 SetPlayBeepNil sets the value for PlayBeep to be an explicit nil

### UnsetPlayBeep
`func (o *InlineObject26) UnsetPlayBeep()`

UnsetPlayBeep ensures that no value is present for PlayBeep, not even an explicit nil
### GetStopTones

`func (o *InlineObject26) GetStopTones() []*string`

GetStopTones returns the StopTones field if non-nil, zero value otherwise.

### GetStopTonesOk

`func (o *InlineObject26) GetStopTonesOk() (*[]*string, bool)`

GetStopTonesOk returns a tuple with the StopTones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTones

`func (o *InlineObject26) SetStopTones(v []*string)`

SetStopTones sets StopTones field to given value.

### HasStopTones

`func (o *InlineObject26) HasStopTones() bool`

HasStopTones returns a boolean if a field has been set.

### GetClientContext

`func (o *InlineObject26) GetClientContext() string`

GetClientContext returns the ClientContext field if non-nil, zero value otherwise.

### GetClientContextOk

`func (o *InlineObject26) GetClientContextOk() (*string, bool)`

GetClientContextOk returns a tuple with the ClientContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContext

`func (o *InlineObject26) SetClientContext(v string)`

SetClientContext sets ClientContext field to given value.

### HasClientContext

`func (o *InlineObject26) HasClientContext() bool`

HasClientContext returns a boolean if a field has been set.

### SetClientContextNil

`func (o *InlineObject26) SetClientContextNil(b bool)`

 SetClientContextNil sets the value for ClientContext to be an explicit nil

### UnsetClientContext
`func (o *InlineObject26) UnsetClientContext()`

UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


