# RemoteAssistancePartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Display name of the partner. | [optional] 
**LastConnectionDateTime** | Pointer to **time.Time** | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 
**OnboardingStatus** | Pointer to [**AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus**](anyOf&lt;microsoft.graph.remoteAssistanceOnboardingStatus&gt;.md) | A friendly description of the current TeamViewer connector status. Possible values are: notOnboarded, onboarding, onboarded. | [optional] 
**OnboardingUrl** | Pointer to **NullableString** | URL of the partner&#39;s onboarding portal, where an administrator can configure their Remote Assistance service. | [optional] 

## Methods

### NewRemoteAssistancePartner

`func NewRemoteAssistancePartner() *RemoteAssistancePartner`

NewRemoteAssistancePartner instantiates a new RemoteAssistancePartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAssistancePartnerWithDefaults

`func NewRemoteAssistancePartnerWithDefaults() *RemoteAssistancePartner`

NewRemoteAssistancePartnerWithDefaults instantiates a new RemoteAssistancePartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *RemoteAssistancePartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoteAssistancePartner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RemoteAssistancePartner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RemoteAssistancePartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *RemoteAssistancePartner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *RemoteAssistancePartner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastConnectionDateTime

`func (o *RemoteAssistancePartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *RemoteAssistancePartner) GetLastConnectionDateTimeOk() (*time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionDateTime

`func (o *RemoteAssistancePartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime sets LastConnectionDateTime field to given value.

### HasLastConnectionDateTime

`func (o *RemoteAssistancePartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### GetOnboardingStatus

`func (o *RemoteAssistancePartner) GetOnboardingStatus() AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *RemoteAssistancePartner) GetOnboardingStatusOk() (*AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingStatus

`func (o *RemoteAssistancePartner) SetOnboardingStatus(v AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus)`

SetOnboardingStatus sets OnboardingStatus field to given value.

### HasOnboardingStatus

`func (o *RemoteAssistancePartner) HasOnboardingStatus() bool`

HasOnboardingStatus returns a boolean if a field has been set.

### SetOnboardingStatusNil

`func (o *RemoteAssistancePartner) SetOnboardingStatusNil(b bool)`

 SetOnboardingStatusNil sets the value for OnboardingStatus to be an explicit nil

### UnsetOnboardingStatus
`func (o *RemoteAssistancePartner) UnsetOnboardingStatus()`

UnsetOnboardingStatus ensures that no value is present for OnboardingStatus, not even an explicit nil
### GetOnboardingUrl

`func (o *RemoteAssistancePartner) GetOnboardingUrl() string`

GetOnboardingUrl returns the OnboardingUrl field if non-nil, zero value otherwise.

### GetOnboardingUrlOk

`func (o *RemoteAssistancePartner) GetOnboardingUrlOk() (*string, bool)`

GetOnboardingUrlOk returns a tuple with the OnboardingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingUrl

`func (o *RemoteAssistancePartner) SetOnboardingUrl(v string)`

SetOnboardingUrl sets OnboardingUrl field to given value.

### HasOnboardingUrl

`func (o *RemoteAssistancePartner) HasOnboardingUrl() bool`

HasOnboardingUrl returns a boolean if a field has been set.

### SetOnboardingUrlNil

`func (o *RemoteAssistancePartner) SetOnboardingUrlNil(b bool)`

 SetOnboardingUrlNil sets the value for OnboardingUrl to be an explicit nil

### UnsetOnboardingUrl
`func (o *RemoteAssistancePartner) UnsetOnboardingUrl()`

UnsetOnboardingUrl ensures that no value is present for OnboardingUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


