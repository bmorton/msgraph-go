# MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The object identifier of an Azure Active Directory user or group. | [optional] 
**TargetedAuthenticationMethod** | Pointer to **NullableString** | The authentication method that the user is prompted to register. The value must be microsoftAuthenticator. | [optional] 
**TargetType** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodTargetType**](anyOf&lt;microsoft.graph.authenticationMethodTargetType&gt;.md) | The type of the authentication method target. Possible values are: user, group, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget

`func NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget() *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget`

NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget instantiates a new MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTargetWithDefaults

`func NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTargetWithDefaults() *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget`

NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTargetWithDefaults instantiates a new MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetedAuthenticationMethod

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetedAuthenticationMethod() string`

GetTargetedAuthenticationMethod returns the TargetedAuthenticationMethod field if non-nil, zero value otherwise.

### GetTargetedAuthenticationMethodOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetedAuthenticationMethodOk() (*string, bool)`

GetTargetedAuthenticationMethodOk returns a tuple with the TargetedAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedAuthenticationMethod

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetedAuthenticationMethod(v string)`

SetTargetedAuthenticationMethod sets TargetedAuthenticationMethod field to given value.

### HasTargetedAuthenticationMethod

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) HasTargetedAuthenticationMethod() bool`

HasTargetedAuthenticationMethod returns a boolean if a field has been set.

### SetTargetedAuthenticationMethodNil

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetedAuthenticationMethodNil(b bool)`

 SetTargetedAuthenticationMethodNil sets the value for TargetedAuthenticationMethod to be an explicit nil

### UnsetTargetedAuthenticationMethod
`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) UnsetTargetedAuthenticationMethod()`

UnsetTargetedAuthenticationMethod ensures that no value is present for TargetedAuthenticationMethod, not even an explicit nil
### GetTargetType

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetType() AnyOfmicrosoftGraphAuthenticationMethodTargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) GetTargetTypeOk() (*AnyOfmicrosoftGraphAuthenticationMethodTargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetType(v AnyOfmicrosoftGraphAuthenticationMethodTargetType)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


