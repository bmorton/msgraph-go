# AuthenticationMethodsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | A description of the policy. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the policy. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The date and time of the last update to the policy. Read-only. | [optional] 
**PolicyVersion** | Pointer to **NullableString** | The version of the policy in use. Read-only. | [optional] 
**ReconfirmationInDays** | Pointer to **NullableInt32** |  | [optional] 
**RegistrationEnforcement** | Pointer to [**AnyOfmicrosoftGraphRegistrationEnforcement**](anyOf&lt;microsoft.graph.registrationEnforcement&gt;.md) | Enforce registration at sign-in time. This property can be used to remind users to set up targeted authentication methods. | [optional] 
**AuthenticationMethodConfigurations** | Pointer to [**[]MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md) | Represents the settings for each authentication method. | [optional] 

## Methods

### NewAuthenticationMethodsPolicy

`func NewAuthenticationMethodsPolicy() *AuthenticationMethodsPolicy`

NewAuthenticationMethodsPolicy instantiates a new AuthenticationMethodsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodsPolicyWithDefaults

`func NewAuthenticationMethodsPolicyWithDefaults() *AuthenticationMethodsPolicy`

NewAuthenticationMethodsPolicyWithDefaults instantiates a new AuthenticationMethodsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AuthenticationMethodsPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationMethodsPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationMethodsPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationMethodsPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AuthenticationMethodsPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AuthenticationMethodsPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *AuthenticationMethodsPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthenticationMethodsPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthenticationMethodsPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthenticationMethodsPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AuthenticationMethodsPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AuthenticationMethodsPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedDateTime

`func (o *AuthenticationMethodsPolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *AuthenticationMethodsPolicy) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *AuthenticationMethodsPolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *AuthenticationMethodsPolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *AuthenticationMethodsPolicy) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *AuthenticationMethodsPolicy) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPolicyVersion

`func (o *AuthenticationMethodsPolicy) GetPolicyVersion() string`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *AuthenticationMethodsPolicy) GetPolicyVersionOk() (*string, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *AuthenticationMethodsPolicy) SetPolicyVersion(v string)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *AuthenticationMethodsPolicy) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### SetPolicyVersionNil

`func (o *AuthenticationMethodsPolicy) SetPolicyVersionNil(b bool)`

 SetPolicyVersionNil sets the value for PolicyVersion to be an explicit nil

### UnsetPolicyVersion
`func (o *AuthenticationMethodsPolicy) UnsetPolicyVersion()`

UnsetPolicyVersion ensures that no value is present for PolicyVersion, not even an explicit nil
### GetReconfirmationInDays

`func (o *AuthenticationMethodsPolicy) GetReconfirmationInDays() int32`

GetReconfirmationInDays returns the ReconfirmationInDays field if non-nil, zero value otherwise.

### GetReconfirmationInDaysOk

`func (o *AuthenticationMethodsPolicy) GetReconfirmationInDaysOk() (*int32, bool)`

GetReconfirmationInDaysOk returns a tuple with the ReconfirmationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconfirmationInDays

`func (o *AuthenticationMethodsPolicy) SetReconfirmationInDays(v int32)`

SetReconfirmationInDays sets ReconfirmationInDays field to given value.

### HasReconfirmationInDays

`func (o *AuthenticationMethodsPolicy) HasReconfirmationInDays() bool`

HasReconfirmationInDays returns a boolean if a field has been set.

### SetReconfirmationInDaysNil

`func (o *AuthenticationMethodsPolicy) SetReconfirmationInDaysNil(b bool)`

 SetReconfirmationInDaysNil sets the value for ReconfirmationInDays to be an explicit nil

### UnsetReconfirmationInDays
`func (o *AuthenticationMethodsPolicy) UnsetReconfirmationInDays()`

UnsetReconfirmationInDays ensures that no value is present for ReconfirmationInDays, not even an explicit nil
### GetRegistrationEnforcement

`func (o *AuthenticationMethodsPolicy) GetRegistrationEnforcement() AnyOfmicrosoftGraphRegistrationEnforcement`

GetRegistrationEnforcement returns the RegistrationEnforcement field if non-nil, zero value otherwise.

### GetRegistrationEnforcementOk

`func (o *AuthenticationMethodsPolicy) GetRegistrationEnforcementOk() (*AnyOfmicrosoftGraphRegistrationEnforcement, bool)`

GetRegistrationEnforcementOk returns a tuple with the RegistrationEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnforcement

`func (o *AuthenticationMethodsPolicy) SetRegistrationEnforcement(v AnyOfmicrosoftGraphRegistrationEnforcement)`

SetRegistrationEnforcement sets RegistrationEnforcement field to given value.

### HasRegistrationEnforcement

`func (o *AuthenticationMethodsPolicy) HasRegistrationEnforcement() bool`

HasRegistrationEnforcement returns a boolean if a field has been set.

### SetRegistrationEnforcementNil

`func (o *AuthenticationMethodsPolicy) SetRegistrationEnforcementNil(b bool)`

 SetRegistrationEnforcementNil sets the value for RegistrationEnforcement to be an explicit nil

### UnsetRegistrationEnforcement
`func (o *AuthenticationMethodsPolicy) UnsetRegistrationEnforcement()`

UnsetRegistrationEnforcement ensures that no value is present for RegistrationEnforcement, not even an explicit nil
### GetAuthenticationMethodConfigurations

`func (o *AuthenticationMethodsPolicy) GetAuthenticationMethodConfigurations() []MicrosoftGraphAuthenticationMethodConfiguration`

GetAuthenticationMethodConfigurations returns the AuthenticationMethodConfigurations field if non-nil, zero value otherwise.

### GetAuthenticationMethodConfigurationsOk

`func (o *AuthenticationMethodsPolicy) GetAuthenticationMethodConfigurationsOk() (*[]MicrosoftGraphAuthenticationMethodConfiguration, bool)`

GetAuthenticationMethodConfigurationsOk returns a tuple with the AuthenticationMethodConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethodConfigurations

`func (o *AuthenticationMethodsPolicy) SetAuthenticationMethodConfigurations(v []MicrosoftGraphAuthenticationMethodConfiguration)`

SetAuthenticationMethodConfigurations sets AuthenticationMethodConfigurations field to given value.

### HasAuthenticationMethodConfigurations

`func (o *AuthenticationMethodsPolicy) HasAuthenticationMethodConfigurations() bool`

HasAuthenticationMethodConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


