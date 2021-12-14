# MicrosoftGraphIncomingContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservedParticipantId** | Pointer to **NullableString** | The ID of the participant that is under observation. Read-only. | [optional] 
**OnBehalfOf** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The identity that the call is happening on behalf of. | [optional] 
**SourceParticipantId** | Pointer to **NullableString** | The ID of the participant that triggered the incoming call. Read-only. | [optional] 
**Transferor** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The identity that transferred the call. | [optional] 

## Methods

### NewMicrosoftGraphIncomingContext

`func NewMicrosoftGraphIncomingContext() *MicrosoftGraphIncomingContext`

NewMicrosoftGraphIncomingContext instantiates a new MicrosoftGraphIncomingContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIncomingContextWithDefaults

`func NewMicrosoftGraphIncomingContextWithDefaults() *MicrosoftGraphIncomingContext`

NewMicrosoftGraphIncomingContextWithDefaults instantiates a new MicrosoftGraphIncomingContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservedParticipantId

`func (o *MicrosoftGraphIncomingContext) GetObservedParticipantId() string`

GetObservedParticipantId returns the ObservedParticipantId field if non-nil, zero value otherwise.

### GetObservedParticipantIdOk

`func (o *MicrosoftGraphIncomingContext) GetObservedParticipantIdOk() (*string, bool)`

GetObservedParticipantIdOk returns a tuple with the ObservedParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedParticipantId

`func (o *MicrosoftGraphIncomingContext) SetObservedParticipantId(v string)`

SetObservedParticipantId sets ObservedParticipantId field to given value.

### HasObservedParticipantId

`func (o *MicrosoftGraphIncomingContext) HasObservedParticipantId() bool`

HasObservedParticipantId returns a boolean if a field has been set.

### SetObservedParticipantIdNil

`func (o *MicrosoftGraphIncomingContext) SetObservedParticipantIdNil(b bool)`

 SetObservedParticipantIdNil sets the value for ObservedParticipantId to be an explicit nil

### UnsetObservedParticipantId
`func (o *MicrosoftGraphIncomingContext) UnsetObservedParticipantId()`

UnsetObservedParticipantId ensures that no value is present for ObservedParticipantId, not even an explicit nil
### GetOnBehalfOf

`func (o *MicrosoftGraphIncomingContext) GetOnBehalfOf() AnyOfmicrosoftGraphIdentitySet`

GetOnBehalfOf returns the OnBehalfOf field if non-nil, zero value otherwise.

### GetOnBehalfOfOk

`func (o *MicrosoftGraphIncomingContext) GetOnBehalfOfOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetOnBehalfOfOk returns a tuple with the OnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOf

`func (o *MicrosoftGraphIncomingContext) SetOnBehalfOf(v AnyOfmicrosoftGraphIdentitySet)`

SetOnBehalfOf sets OnBehalfOf field to given value.

### HasOnBehalfOf

`func (o *MicrosoftGraphIncomingContext) HasOnBehalfOf() bool`

HasOnBehalfOf returns a boolean if a field has been set.

### SetOnBehalfOfNil

`func (o *MicrosoftGraphIncomingContext) SetOnBehalfOfNil(b bool)`

 SetOnBehalfOfNil sets the value for OnBehalfOf to be an explicit nil

### UnsetOnBehalfOf
`func (o *MicrosoftGraphIncomingContext) UnsetOnBehalfOf()`

UnsetOnBehalfOf ensures that no value is present for OnBehalfOf, not even an explicit nil
### GetSourceParticipantId

`func (o *MicrosoftGraphIncomingContext) GetSourceParticipantId() string`

GetSourceParticipantId returns the SourceParticipantId field if non-nil, zero value otherwise.

### GetSourceParticipantIdOk

`func (o *MicrosoftGraphIncomingContext) GetSourceParticipantIdOk() (*string, bool)`

GetSourceParticipantIdOk returns a tuple with the SourceParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParticipantId

`func (o *MicrosoftGraphIncomingContext) SetSourceParticipantId(v string)`

SetSourceParticipantId sets SourceParticipantId field to given value.

### HasSourceParticipantId

`func (o *MicrosoftGraphIncomingContext) HasSourceParticipantId() bool`

HasSourceParticipantId returns a boolean if a field has been set.

### SetSourceParticipantIdNil

`func (o *MicrosoftGraphIncomingContext) SetSourceParticipantIdNil(b bool)`

 SetSourceParticipantIdNil sets the value for SourceParticipantId to be an explicit nil

### UnsetSourceParticipantId
`func (o *MicrosoftGraphIncomingContext) UnsetSourceParticipantId()`

UnsetSourceParticipantId ensures that no value is present for SourceParticipantId, not even an explicit nil
### GetTransferor

`func (o *MicrosoftGraphIncomingContext) GetTransferor() AnyOfmicrosoftGraphIdentitySet`

GetTransferor returns the Transferor field if non-nil, zero value otherwise.

### GetTransferorOk

`func (o *MicrosoftGraphIncomingContext) GetTransferorOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetTransferorOk returns a tuple with the Transferor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferor

`func (o *MicrosoftGraphIncomingContext) SetTransferor(v AnyOfmicrosoftGraphIdentitySet)`

SetTransferor sets Transferor field to given value.

### HasTransferor

`func (o *MicrosoftGraphIncomingContext) HasTransferor() bool`

HasTransferor returns a boolean if a field has been set.

### SetTransferorNil

`func (o *MicrosoftGraphIncomingContext) SetTransferorNil(b bool)`

 SetTransferorNil sets the value for Transferor to be an explicit nil

### UnsetTransferor
`func (o *MicrosoftGraphIncomingContext) UnsetTransferor()`

UnsetTransferor ensures that no value is present for Transferor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


