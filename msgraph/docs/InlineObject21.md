# InlineObject21

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUri** | Pointer to **string** |  | [optional] 
**MediaConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**AcceptedModalities** | Pointer to [**[]AnyOfmicrosoftGraphModality**](AnyOfmicrosoftGraphModality.md) |  | [optional] 
**ParticipantCapacity** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewInlineObject21

`func NewInlineObject21() *InlineObject21`

NewInlineObject21 instantiates a new InlineObject21 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject21WithDefaults

`func NewInlineObject21WithDefaults() *InlineObject21`

NewInlineObject21WithDefaults instantiates a new InlineObject21 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUri

`func (o *InlineObject21) GetCallbackUri() string`

GetCallbackUri returns the CallbackUri field if non-nil, zero value otherwise.

### GetCallbackUriOk

`func (o *InlineObject21) GetCallbackUriOk() (*string, bool)`

GetCallbackUriOk returns a tuple with the CallbackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUri

`func (o *InlineObject21) SetCallbackUri(v string)`

SetCallbackUri sets CallbackUri field to given value.

### HasCallbackUri

`func (o *InlineObject21) HasCallbackUri() bool`

HasCallbackUri returns a boolean if a field has been set.

### GetMediaConfig

`func (o *InlineObject21) GetMediaConfig() map[string]interface{}`

GetMediaConfig returns the MediaConfig field if non-nil, zero value otherwise.

### GetMediaConfigOk

`func (o *InlineObject21) GetMediaConfigOk() (*map[string]interface{}, bool)`

GetMediaConfigOk returns a tuple with the MediaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaConfig

`func (o *InlineObject21) SetMediaConfig(v map[string]interface{})`

SetMediaConfig sets MediaConfig field to given value.

### HasMediaConfig

`func (o *InlineObject21) HasMediaConfig() bool`

HasMediaConfig returns a boolean if a field has been set.

### GetAcceptedModalities

`func (o *InlineObject21) GetAcceptedModalities() []*AnyOfmicrosoftGraphModality`

GetAcceptedModalities returns the AcceptedModalities field if non-nil, zero value otherwise.

### GetAcceptedModalitiesOk

`func (o *InlineObject21) GetAcceptedModalitiesOk() (*[]*AnyOfmicrosoftGraphModality, bool)`

GetAcceptedModalitiesOk returns a tuple with the AcceptedModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedModalities

`func (o *InlineObject21) SetAcceptedModalities(v []*AnyOfmicrosoftGraphModality)`

SetAcceptedModalities sets AcceptedModalities field to given value.

### HasAcceptedModalities

`func (o *InlineObject21) HasAcceptedModalities() bool`

HasAcceptedModalities returns a boolean if a field has been set.

### GetParticipantCapacity

`func (o *InlineObject21) GetParticipantCapacity() int32`

GetParticipantCapacity returns the ParticipantCapacity field if non-nil, zero value otherwise.

### GetParticipantCapacityOk

`func (o *InlineObject21) GetParticipantCapacityOk() (*int32, bool)`

GetParticipantCapacityOk returns a tuple with the ParticipantCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCapacity

`func (o *InlineObject21) SetParticipantCapacity(v int32)`

SetParticipantCapacity sets ParticipantCapacity field to given value.

### HasParticipantCapacity

`func (o *InlineObject21) HasParticipantCapacity() bool`

HasParticipantCapacity returns a boolean if a field has been set.

### SetParticipantCapacityNil

`func (o *InlineObject21) SetParticipantCapacityNil(b bool)`

 SetParticipantCapacityNil sets the value for ParticipantCapacity to be an explicit nil

### UnsetParticipantCapacity
`func (o *InlineObject21) UnsetParticipantCapacity()`

UnsetParticipantCapacity ensures that no value is present for ParticipantCapacity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


