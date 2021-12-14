# MicrosoftGraphAudioRoutingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Receivers** | Pointer to **[]string** | List of receiving participant ids. | [optional] 
**RoutingMode** | Pointer to [**AnyOfmicrosoftGraphRoutingMode**](anyOf&lt;microsoft.graph.routingMode&gt;.md) | Routing group mode.  Possible values are: oneToOne, multicast. | [optional] 
**Sources** | Pointer to **[]string** | List of source participant ids. | [optional] 

## Methods

### NewMicrosoftGraphAudioRoutingGroup

`func NewMicrosoftGraphAudioRoutingGroup() *MicrosoftGraphAudioRoutingGroup`

NewMicrosoftGraphAudioRoutingGroup instantiates a new MicrosoftGraphAudioRoutingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAudioRoutingGroupWithDefaults

`func NewMicrosoftGraphAudioRoutingGroupWithDefaults() *MicrosoftGraphAudioRoutingGroup`

NewMicrosoftGraphAudioRoutingGroupWithDefaults instantiates a new MicrosoftGraphAudioRoutingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAudioRoutingGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAudioRoutingGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAudioRoutingGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAudioRoutingGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReceivers

`func (o *MicrosoftGraphAudioRoutingGroup) GetReceivers() []*string`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *MicrosoftGraphAudioRoutingGroup) GetReceiversOk() (*[]*string, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *MicrosoftGraphAudioRoutingGroup) SetReceivers(v []*string)`

SetReceivers sets Receivers field to given value.

### HasReceivers

`func (o *MicrosoftGraphAudioRoutingGroup) HasReceivers() bool`

HasReceivers returns a boolean if a field has been set.

### GetRoutingMode

`func (o *MicrosoftGraphAudioRoutingGroup) GetRoutingMode() AnyOfmicrosoftGraphRoutingMode`

GetRoutingMode returns the RoutingMode field if non-nil, zero value otherwise.

### GetRoutingModeOk

`func (o *MicrosoftGraphAudioRoutingGroup) GetRoutingModeOk() (*AnyOfmicrosoftGraphRoutingMode, bool)`

GetRoutingModeOk returns a tuple with the RoutingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingMode

`func (o *MicrosoftGraphAudioRoutingGroup) SetRoutingMode(v AnyOfmicrosoftGraphRoutingMode)`

SetRoutingMode sets RoutingMode field to given value.

### HasRoutingMode

`func (o *MicrosoftGraphAudioRoutingGroup) HasRoutingMode() bool`

HasRoutingMode returns a boolean if a field has been set.

### SetRoutingModeNil

`func (o *MicrosoftGraphAudioRoutingGroup) SetRoutingModeNil(b bool)`

 SetRoutingModeNil sets the value for RoutingMode to be an explicit nil

### UnsetRoutingMode
`func (o *MicrosoftGraphAudioRoutingGroup) UnsetRoutingMode()`

UnsetRoutingMode ensures that no value is present for RoutingMode, not even an explicit nil
### GetSources

`func (o *MicrosoftGraphAudioRoutingGroup) GetSources() []*string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MicrosoftGraphAudioRoutingGroup) GetSourcesOk() (*[]*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MicrosoftGraphAudioRoutingGroup) SetSources(v []*string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MicrosoftGraphAudioRoutingGroup) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


