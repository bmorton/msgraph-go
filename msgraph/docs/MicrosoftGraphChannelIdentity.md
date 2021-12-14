# MicrosoftGraphChannelIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **NullableString** | The identity of the channel in which the message was posted. | [optional] 
**TeamId** | Pointer to **NullableString** | The identity of the team in which the message was posted. | [optional] 

## Methods

### NewMicrosoftGraphChannelIdentity

`func NewMicrosoftGraphChannelIdentity() *MicrosoftGraphChannelIdentity`

NewMicrosoftGraphChannelIdentity instantiates a new MicrosoftGraphChannelIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChannelIdentityWithDefaults

`func NewMicrosoftGraphChannelIdentityWithDefaults() *MicrosoftGraphChannelIdentity`

NewMicrosoftGraphChannelIdentityWithDefaults instantiates a new MicrosoftGraphChannelIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *MicrosoftGraphChannelIdentity) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *MicrosoftGraphChannelIdentity) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *MicrosoftGraphChannelIdentity) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *MicrosoftGraphChannelIdentity) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *MicrosoftGraphChannelIdentity) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *MicrosoftGraphChannelIdentity) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetTeamId

`func (o *MicrosoftGraphChannelIdentity) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MicrosoftGraphChannelIdentity) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MicrosoftGraphChannelIdentity) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *MicrosoftGraphChannelIdentity) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *MicrosoftGraphChannelIdentity) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *MicrosoftGraphChannelIdentity) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


