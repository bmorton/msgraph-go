# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fido2Methods** | Pointer to [**[]MicrosoftGraphFido2AuthenticationMethod**](MicrosoftGraphFido2AuthenticationMethod.md) |  | [optional] 
**Methods** | Pointer to [**[]MicrosoftGraphAuthenticationMethod**](MicrosoftGraphAuthenticationMethod.md) |  | [optional] 
**MicrosoftAuthenticatorMethods** | Pointer to [**[]MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod**](MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod.md) |  | [optional] 
**WindowsHelloForBusinessMethods** | Pointer to [**[]MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod**](MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod.md) |  | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFido2Methods

`func (o *Authentication) GetFido2Methods() []MicrosoftGraphFido2AuthenticationMethod`

GetFido2Methods returns the Fido2Methods field if non-nil, zero value otherwise.

### GetFido2MethodsOk

`func (o *Authentication) GetFido2MethodsOk() (*[]MicrosoftGraphFido2AuthenticationMethod, bool)`

GetFido2MethodsOk returns a tuple with the Fido2Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2Methods

`func (o *Authentication) SetFido2Methods(v []MicrosoftGraphFido2AuthenticationMethod)`

SetFido2Methods sets Fido2Methods field to given value.

### HasFido2Methods

`func (o *Authentication) HasFido2Methods() bool`

HasFido2Methods returns a boolean if a field has been set.

### GetMethods

`func (o *Authentication) GetMethods() []MicrosoftGraphAuthenticationMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Authentication) GetMethodsOk() (*[]MicrosoftGraphAuthenticationMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Authentication) SetMethods(v []MicrosoftGraphAuthenticationMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *Authentication) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetMicrosoftAuthenticatorMethods

`func (o *Authentication) GetMicrosoftAuthenticatorMethods() []MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod`

GetMicrosoftAuthenticatorMethods returns the MicrosoftAuthenticatorMethods field if non-nil, zero value otherwise.

### GetMicrosoftAuthenticatorMethodsOk

`func (o *Authentication) GetMicrosoftAuthenticatorMethodsOk() (*[]MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod, bool)`

GetMicrosoftAuthenticatorMethodsOk returns a tuple with the MicrosoftAuthenticatorMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftAuthenticatorMethods

`func (o *Authentication) SetMicrosoftAuthenticatorMethods(v []MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod)`

SetMicrosoftAuthenticatorMethods sets MicrosoftAuthenticatorMethods field to given value.

### HasMicrosoftAuthenticatorMethods

`func (o *Authentication) HasMicrosoftAuthenticatorMethods() bool`

HasMicrosoftAuthenticatorMethods returns a boolean if a field has been set.

### GetWindowsHelloForBusinessMethods

`func (o *Authentication) GetWindowsHelloForBusinessMethods() []MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod`

GetWindowsHelloForBusinessMethods returns the WindowsHelloForBusinessMethods field if non-nil, zero value otherwise.

### GetWindowsHelloForBusinessMethodsOk

`func (o *Authentication) GetWindowsHelloForBusinessMethodsOk() (*[]MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod, bool)`

GetWindowsHelloForBusinessMethodsOk returns a tuple with the WindowsHelloForBusinessMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsHelloForBusinessMethods

`func (o *Authentication) SetWindowsHelloForBusinessMethods(v []MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod)`

SetWindowsHelloForBusinessMethods sets WindowsHelloForBusinessMethods field to given value.

### HasWindowsHelloForBusinessMethods

`func (o *Authentication) HasWindowsHelloForBusinessMethods() bool`

HasWindowsHelloForBusinessMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


