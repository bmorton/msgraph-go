# MicrosoftGraphAuthenticationMethodsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | A description of the policy. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the policy. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The date and time of the last update to the policy. Read-only. | [optional] 
**PolicyVersion** | Pointer to **NullableString** | The version of the policy in use. Read-only. | [optional] 
**ReconfirmationInDays** | Pointer to **NullableInt32** |  | [optional] 
**RegistrationEnforcement** | Pointer to [**AnyOfmicrosoftGraphRegistrationEnforcement**](anyOf&lt;microsoft.graph.registrationEnforcement&gt;.md) | Enforce registration at sign-in time. This property can be used to remind users to set up targeted authentication methods. | [optional] 
**AuthenticationMethodConfigurations** | Pointer to [**[]MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md) | Represents the settings for each authentication method. | [optional] 

## Methods

### NewMicrosoftGraphAuthenticationMethodsPolicy

`func NewMicrosoftGraphAuthenticationMethodsPolicy() *MicrosoftGraphAuthenticationMethodsPolicy`

NewMicrosoftGraphAuthenticationMethodsPolicy instantiates a new MicrosoftGraphAuthenticationMethodsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthenticationMethodsPolicyWithDefaults

`func NewMicrosoftGraphAuthenticationMethodsPolicyWithDefaults() *MicrosoftGraphAuthenticationMethodsPolicy`

NewMicrosoftGraphAuthenticationMethodsPolicyWithDefaults instantiates a new MicrosoftGraphAuthenticationMethodsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAuthenticationMethodsPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAuthenticationMethodsPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphAuthenticationMethodsPolicy) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPolicyVersion

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetPolicyVersion() string`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetPolicyVersionOk() (*string, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetPolicyVersion(v string)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### SetPolicyVersionNil

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetPolicyVersionNil(b bool)`

 SetPolicyVersionNil sets the value for PolicyVersion to be an explicit nil

### UnsetPolicyVersion
`func (o *MicrosoftGraphAuthenticationMethodsPolicy) UnsetPolicyVersion()`

UnsetPolicyVersion ensures that no value is present for PolicyVersion, not even an explicit nil
### GetReconfirmationInDays

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetReconfirmationInDays() int32`

GetReconfirmationInDays returns the ReconfirmationInDays field if non-nil, zero value otherwise.

### GetReconfirmationInDaysOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetReconfirmationInDaysOk() (*int32, bool)`

GetReconfirmationInDaysOk returns a tuple with the ReconfirmationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconfirmationInDays

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetReconfirmationInDays(v int32)`

SetReconfirmationInDays sets ReconfirmationInDays field to given value.

### HasReconfirmationInDays

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasReconfirmationInDays() bool`

HasReconfirmationInDays returns a boolean if a field has been set.

### SetReconfirmationInDaysNil

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetReconfirmationInDaysNil(b bool)`

 SetReconfirmationInDaysNil sets the value for ReconfirmationInDays to be an explicit nil

### UnsetReconfirmationInDays
`func (o *MicrosoftGraphAuthenticationMethodsPolicy) UnsetReconfirmationInDays()`

UnsetReconfirmationInDays ensures that no value is present for ReconfirmationInDays, not even an explicit nil
### GetRegistrationEnforcement

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetRegistrationEnforcement() AnyOfmicrosoftGraphRegistrationEnforcement`

GetRegistrationEnforcement returns the RegistrationEnforcement field if non-nil, zero value otherwise.

### GetRegistrationEnforcementOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetRegistrationEnforcementOk() (*AnyOfmicrosoftGraphRegistrationEnforcement, bool)`

GetRegistrationEnforcementOk returns a tuple with the RegistrationEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnforcement

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetRegistrationEnforcement(v AnyOfmicrosoftGraphRegistrationEnforcement)`

SetRegistrationEnforcement sets RegistrationEnforcement field to given value.

### HasRegistrationEnforcement

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasRegistrationEnforcement() bool`

HasRegistrationEnforcement returns a boolean if a field has been set.

### SetRegistrationEnforcementNil

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetRegistrationEnforcementNil(b bool)`

 SetRegistrationEnforcementNil sets the value for RegistrationEnforcement to be an explicit nil

### UnsetRegistrationEnforcement
`func (o *MicrosoftGraphAuthenticationMethodsPolicy) UnsetRegistrationEnforcement()`

UnsetRegistrationEnforcement ensures that no value is present for RegistrationEnforcement, not even an explicit nil
### GetAuthenticationMethodConfigurations

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetAuthenticationMethodConfigurations() []MicrosoftGraphAuthenticationMethodConfiguration`

GetAuthenticationMethodConfigurations returns the AuthenticationMethodConfigurations field if non-nil, zero value otherwise.

### GetAuthenticationMethodConfigurationsOk

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) GetAuthenticationMethodConfigurationsOk() (*[]MicrosoftGraphAuthenticationMethodConfiguration, bool)`

GetAuthenticationMethodConfigurationsOk returns a tuple with the AuthenticationMethodConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethodConfigurations

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) SetAuthenticationMethodConfigurations(v []MicrosoftGraphAuthenticationMethodConfiguration)`

SetAuthenticationMethodConfigurations sets AuthenticationMethodConfigurations field to given value.

### HasAuthenticationMethodConfigurations

`func (o *MicrosoftGraphAuthenticationMethodsPolicy) HasAuthenticationMethodConfigurations() bool`

HasAuthenticationMethodConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


