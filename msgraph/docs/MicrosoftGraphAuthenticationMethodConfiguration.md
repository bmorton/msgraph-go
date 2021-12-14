# MicrosoftGraphAuthenticationMethodConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodState**](anyOf&lt;microsoft.graph.authenticationMethodState&gt;.md) | The state of the policy. Possible values are: enabled, disabled. | [optional] 

## Methods

### NewMicrosoftGraphAuthenticationMethodConfiguration

`func NewMicrosoftGraphAuthenticationMethodConfiguration() *MicrosoftGraphAuthenticationMethodConfiguration`

NewMicrosoftGraphAuthenticationMethodConfiguration instantiates a new MicrosoftGraphAuthenticationMethodConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthenticationMethodConfigurationWithDefaults

`func NewMicrosoftGraphAuthenticationMethodConfigurationWithDefaults() *MicrosoftGraphAuthenticationMethodConfiguration`

NewMicrosoftGraphAuthenticationMethodConfigurationWithDefaults instantiates a new MicrosoftGraphAuthenticationMethodConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) GetState() AnyOfmicrosoftGraphAuthenticationMethodState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) GetStateOk() (*AnyOfmicrosoftGraphAuthenticationMethodState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) SetState(v AnyOfmicrosoftGraphAuthenticationMethodState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphAuthenticationMethodConfiguration) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphAuthenticationMethodConfiguration) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


