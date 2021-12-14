# AuthenticationMethodConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodState**](anyOf&lt;microsoft.graph.authenticationMethodState&gt;.md) | The state of the policy. Possible values are: enabled, disabled. | [optional] 

## Methods

### NewAuthenticationMethodConfiguration

`func NewAuthenticationMethodConfiguration() *AuthenticationMethodConfiguration`

NewAuthenticationMethodConfiguration instantiates a new AuthenticationMethodConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodConfigurationWithDefaults

`func NewAuthenticationMethodConfigurationWithDefaults() *AuthenticationMethodConfiguration`

NewAuthenticationMethodConfigurationWithDefaults instantiates a new AuthenticationMethodConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AuthenticationMethodConfiguration) GetState() AnyOfmicrosoftGraphAuthenticationMethodState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AuthenticationMethodConfiguration) GetStateOk() (*AnyOfmicrosoftGraphAuthenticationMethodState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AuthenticationMethodConfiguration) SetState(v AnyOfmicrosoftGraphAuthenticationMethodState)`

SetState sets State field to given value.

### HasState

`func (o *AuthenticationMethodConfiguration) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AuthenticationMethodConfiguration) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AuthenticationMethodConfiguration) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


