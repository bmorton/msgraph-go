# AuthenticationFlowsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Inherited property. A description of the policy. Optional. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | Inherited property. The human-readable name of the policy. Optional. Read-only. | [optional] 
**SelfServiceSignUp** | Pointer to [**AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration**](anyOf&lt;microsoft.graph.selfServiceSignUpAuthenticationFlowConfiguration&gt;.md) | Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only. | [optional] 

## Methods

### NewAuthenticationFlowsPolicy

`func NewAuthenticationFlowsPolicy() *AuthenticationFlowsPolicy`

NewAuthenticationFlowsPolicy instantiates a new AuthenticationFlowsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationFlowsPolicyWithDefaults

`func NewAuthenticationFlowsPolicyWithDefaults() *AuthenticationFlowsPolicy`

NewAuthenticationFlowsPolicyWithDefaults instantiates a new AuthenticationFlowsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AuthenticationFlowsPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationFlowsPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationFlowsPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationFlowsPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AuthenticationFlowsPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AuthenticationFlowsPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *AuthenticationFlowsPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthenticationFlowsPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthenticationFlowsPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthenticationFlowsPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AuthenticationFlowsPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AuthenticationFlowsPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSelfServiceSignUp

`func (o *AuthenticationFlowsPolicy) GetSelfServiceSignUp() AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration`

GetSelfServiceSignUp returns the SelfServiceSignUp field if non-nil, zero value otherwise.

### GetSelfServiceSignUpOk

`func (o *AuthenticationFlowsPolicy) GetSelfServiceSignUpOk() (*AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration, bool)`

GetSelfServiceSignUpOk returns a tuple with the SelfServiceSignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceSignUp

`func (o *AuthenticationFlowsPolicy) SetSelfServiceSignUp(v AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration)`

SetSelfServiceSignUp sets SelfServiceSignUp field to given value.

### HasSelfServiceSignUp

`func (o *AuthenticationFlowsPolicy) HasSelfServiceSignUp() bool`

HasSelfServiceSignUp returns a boolean if a field has been set.

### SetSelfServiceSignUpNil

`func (o *AuthenticationFlowsPolicy) SetSelfServiceSignUpNil(b bool)`

 SetSelfServiceSignUpNil sets the value for SelfServiceSignUp to be an explicit nil

### UnsetSelfServiceSignUp
`func (o *AuthenticationFlowsPolicy) UnsetSelfServiceSignUp()`

UnsetSelfServiceSignUp ensures that no value is present for SelfServiceSignUp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


