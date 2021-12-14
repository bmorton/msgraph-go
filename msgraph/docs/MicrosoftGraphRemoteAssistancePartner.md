# MicrosoftGraphRemoteAssistancePartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name of the partner. | [optional] 
**LastConnectionDateTime** | Pointer to **time.Time** | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 
**OnboardingStatus** | Pointer to [**AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus**](anyOf&lt;microsoft.graph.remoteAssistanceOnboardingStatus&gt;.md) | A friendly description of the current TeamViewer connector status. Possible values are: notOnboarded, onboarding, onboarded. | [optional] 
**OnboardingUrl** | Pointer to **NullableString** | URL of the partner&#39;s onboarding portal, where an administrator can configure their Remote Assistance service. | [optional] 

## Methods

### NewMicrosoftGraphRemoteAssistancePartner

`func NewMicrosoftGraphRemoteAssistancePartner() *MicrosoftGraphRemoteAssistancePartner`

NewMicrosoftGraphRemoteAssistancePartner instantiates a new MicrosoftGraphRemoteAssistancePartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRemoteAssistancePartnerWithDefaults

`func NewMicrosoftGraphRemoteAssistancePartnerWithDefaults() *MicrosoftGraphRemoteAssistancePartner`

NewMicrosoftGraphRemoteAssistancePartnerWithDefaults instantiates a new MicrosoftGraphRemoteAssistancePartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphRemoteAssistancePartner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRemoteAssistancePartner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphRemoteAssistancePartner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphRemoteAssistancePartner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphRemoteAssistancePartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRemoteAssistancePartner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphRemoteAssistancePartner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphRemoteAssistancePartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphRemoteAssistancePartner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphRemoteAssistancePartner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastConnectionDateTime

`func (o *MicrosoftGraphRemoteAssistancePartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *MicrosoftGraphRemoteAssistancePartner) GetLastConnectionDateTimeOk() (*time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionDateTime

`func (o *MicrosoftGraphRemoteAssistancePartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime sets LastConnectionDateTime field to given value.

### HasLastConnectionDateTime

`func (o *MicrosoftGraphRemoteAssistancePartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### GetOnboardingStatus

`func (o *MicrosoftGraphRemoteAssistancePartner) GetOnboardingStatus() AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *MicrosoftGraphRemoteAssistancePartner) GetOnboardingStatusOk() (*AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingStatus

`func (o *MicrosoftGraphRemoteAssistancePartner) SetOnboardingStatus(v AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus)`

SetOnboardingStatus sets OnboardingStatus field to given value.

### HasOnboardingStatus

`func (o *MicrosoftGraphRemoteAssistancePartner) HasOnboardingStatus() bool`

HasOnboardingStatus returns a boolean if a field has been set.

### SetOnboardingStatusNil

`func (o *MicrosoftGraphRemoteAssistancePartner) SetOnboardingStatusNil(b bool)`

 SetOnboardingStatusNil sets the value for OnboardingStatus to be an explicit nil

### UnsetOnboardingStatus
`func (o *MicrosoftGraphRemoteAssistancePartner) UnsetOnboardingStatus()`

UnsetOnboardingStatus ensures that no value is present for OnboardingStatus, not even an explicit nil
### GetOnboardingUrl

`func (o *MicrosoftGraphRemoteAssistancePartner) GetOnboardingUrl() string`

GetOnboardingUrl returns the OnboardingUrl field if non-nil, zero value otherwise.

### GetOnboardingUrlOk

`func (o *MicrosoftGraphRemoteAssistancePartner) GetOnboardingUrlOk() (*string, bool)`

GetOnboardingUrlOk returns a tuple with the OnboardingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingUrl

`func (o *MicrosoftGraphRemoteAssistancePartner) SetOnboardingUrl(v string)`

SetOnboardingUrl sets OnboardingUrl field to given value.

### HasOnboardingUrl

`func (o *MicrosoftGraphRemoteAssistancePartner) HasOnboardingUrl() bool`

HasOnboardingUrl returns a boolean if a field has been set.

### SetOnboardingUrlNil

`func (o *MicrosoftGraphRemoteAssistancePartner) SetOnboardingUrlNil(b bool)`

 SetOnboardingUrlNil sets the value for OnboardingUrl to be an explicit nil

### UnsetOnboardingUrl
`func (o *MicrosoftGraphRemoteAssistancePartner) UnsetOnboardingUrl()`

UnsetOnboardingUrl ensures that no value is present for OnboardingUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


