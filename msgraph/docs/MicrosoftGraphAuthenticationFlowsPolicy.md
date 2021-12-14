# MicrosoftGraphAuthenticationFlowsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Inherited property. A description of the policy. Optional. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | Inherited property. The human-readable name of the policy. Optional. Read-only. | [optional] 
**SelfServiceSignUp** | Pointer to [**AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration**](anyOf&lt;microsoft.graph.selfServiceSignUpAuthenticationFlowConfiguration&gt;.md) | Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphAuthenticationFlowsPolicy

`func NewMicrosoftGraphAuthenticationFlowsPolicy() *MicrosoftGraphAuthenticationFlowsPolicy`

NewMicrosoftGraphAuthenticationFlowsPolicy instantiates a new MicrosoftGraphAuthenticationFlowsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthenticationFlowsPolicyWithDefaults

`func NewMicrosoftGraphAuthenticationFlowsPolicyWithDefaults() *MicrosoftGraphAuthenticationFlowsPolicy`

NewMicrosoftGraphAuthenticationFlowsPolicyWithDefaults instantiates a new MicrosoftGraphAuthenticationFlowsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAuthenticationFlowsPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAuthenticationFlowsPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSelfServiceSignUp

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetSelfServiceSignUp() AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration`

GetSelfServiceSignUp returns the SelfServiceSignUp field if non-nil, zero value otherwise.

### GetSelfServiceSignUpOk

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) GetSelfServiceSignUpOk() (*AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration, bool)`

GetSelfServiceSignUpOk returns a tuple with the SelfServiceSignUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceSignUp

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetSelfServiceSignUp(v AnyOfmicrosoftGraphSelfServiceSignUpAuthenticationFlowConfiguration)`

SetSelfServiceSignUp sets SelfServiceSignUp field to given value.

### HasSelfServiceSignUp

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) HasSelfServiceSignUp() bool`

HasSelfServiceSignUp returns a boolean if a field has been set.

### SetSelfServiceSignUpNil

`func (o *MicrosoftGraphAuthenticationFlowsPolicy) SetSelfServiceSignUpNil(b bool)`

 SetSelfServiceSignUpNil sets the value for SelfServiceSignUp to be an explicit nil

### UnsetSelfServiceSignUp
`func (o *MicrosoftGraphAuthenticationFlowsPolicy) UnsetSelfServiceSignUp()`

UnsetSelfServiceSignUp ensures that no value is present for SelfServiceSignUp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


