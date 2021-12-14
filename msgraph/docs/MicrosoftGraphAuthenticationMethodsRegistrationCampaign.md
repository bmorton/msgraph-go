# MicrosoftGraphAuthenticationMethodsRegistrationCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeTargets** | Pointer to [**[]MicrosoftGraphExcludeTarget**](MicrosoftGraphExcludeTarget.md) | Users and groups of users that are excluded from being prompted to set up the authentication method. | [optional] 
**IncludeTargets** | Pointer to [**[]MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget**](MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget.md) | Users and groups of users that are prompted to set up the authentication method. | [optional] 
**SnoozeDurationInDays** | Pointer to **int32** | Specifies the number of days that the user sees a prompt again if they select &#39;Not now&#39; and snoozes the prompt. Minimum: 0 days. Maximum: 14 days. If the value is &#39;0&#39;, the user is prompted during every MFA attempt. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphAdvancedConfigState**](anyOf&lt;microsoft.graph.advancedConfigState&gt;.md) | Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn&#39;t been explicitly set and uses the default behavior of Azure Active Directory for the setting. The default value is disabled. | [optional] 

## Methods

### NewMicrosoftGraphAuthenticationMethodsRegistrationCampaign

`func NewMicrosoftGraphAuthenticationMethodsRegistrationCampaign() *MicrosoftGraphAuthenticationMethodsRegistrationCampaign`

NewMicrosoftGraphAuthenticationMethodsRegistrationCampaign instantiates a new MicrosoftGraphAuthenticationMethodsRegistrationCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignWithDefaults

`func NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignWithDefaults() *MicrosoftGraphAuthenticationMethodsRegistrationCampaign`

NewMicrosoftGraphAuthenticationMethodsRegistrationCampaignWithDefaults instantiates a new MicrosoftGraphAuthenticationMethodsRegistrationCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeTargets

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetExcludeTargets() []MicrosoftGraphExcludeTarget`

GetExcludeTargets returns the ExcludeTargets field if non-nil, zero value otherwise.

### GetExcludeTargetsOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetExcludeTargetsOk() (*[]MicrosoftGraphExcludeTarget, bool)`

GetExcludeTargetsOk returns a tuple with the ExcludeTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTargets

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) SetExcludeTargets(v []MicrosoftGraphExcludeTarget)`

SetExcludeTargets sets ExcludeTargets field to given value.

### HasExcludeTargets

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) HasExcludeTargets() bool`

HasExcludeTargets returns a boolean if a field has been set.

### GetIncludeTargets

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetIncludeTargets() []MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget`

GetIncludeTargets returns the IncludeTargets field if non-nil, zero value otherwise.

### GetIncludeTargetsOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetIncludeTargetsOk() (*[]MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget, bool)`

GetIncludeTargetsOk returns a tuple with the IncludeTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTargets

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) SetIncludeTargets(v []MicrosoftGraphAuthenticationMethodsRegistrationCampaignIncludeTarget)`

SetIncludeTargets sets IncludeTargets field to given value.

### HasIncludeTargets

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) HasIncludeTargets() bool`

HasIncludeTargets returns a boolean if a field has been set.

### GetSnoozeDurationInDays

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetSnoozeDurationInDays() int32`

GetSnoozeDurationInDays returns the SnoozeDurationInDays field if non-nil, zero value otherwise.

### GetSnoozeDurationInDaysOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetSnoozeDurationInDaysOk() (*int32, bool)`

GetSnoozeDurationInDaysOk returns a tuple with the SnoozeDurationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoozeDurationInDays

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) SetSnoozeDurationInDays(v int32)`

SetSnoozeDurationInDays sets SnoozeDurationInDays field to given value.

### HasSnoozeDurationInDays

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) HasSnoozeDurationInDays() bool`

HasSnoozeDurationInDays returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetState() AnyOfmicrosoftGraphAdvancedConfigState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) GetStateOk() (*AnyOfmicrosoftGraphAdvancedConfigState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) SetState(v AnyOfmicrosoftGraphAdvancedConfigState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphAuthenticationMethodsRegistrationCampaign) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


