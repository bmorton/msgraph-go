# MicrosoftGraphDeviceAppManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**IsEnabledForMicrosoftStoreForBusiness** | Pointer to **bool** | Whether the account is enabled for syncing applications from the Microsoft Store for Business. | [optional] 
**MicrosoftStoreForBusinessLanguage** | Pointer to **NullableString** | The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -&lt;country/regioncode2&gt;, where  is a lowercase two-letter code derived from ISO 639-1 and &lt;country/regioncode2&gt; is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture. | [optional] 
**MicrosoftStoreForBusinessLastCompletedApplicationSyncTime** | Pointer to **time.Time** | The last time an application sync from the Microsoft Store for Business was completed. | [optional] 
**MicrosoftStoreForBusinessLastSuccessfulSyncDateTime** | Pointer to **time.Time** | The last time the apps from the Microsoft Store for Business were synced successfully for the account. | [optional] 
**ManagedEBooks** | Pointer to [**[]MicrosoftGraphManagedEBook**](MicrosoftGraphManagedEBook.md) | The Managed eBook. | [optional] 
**MobileAppCategories** | Pointer to [**[]MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md) | The mobile app categories. | [optional] 
**MobileAppConfigurations** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfiguration**](MicrosoftGraphManagedDeviceMobileAppConfiguration.md) | The Managed Device Mobile Application Configurations. | [optional] 
**MobileApps** | Pointer to [**[]MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md) | The mobile apps. | [optional] 
**VppTokens** | Pointer to [**[]MicrosoftGraphVppToken**](MicrosoftGraphVppToken.md) | List of Vpp tokens for this organization. | [optional] 
**AndroidManagedAppProtections** | Pointer to [**[]MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md) | Android managed app policies. | [optional] 
**DefaultManagedAppProtections** | Pointer to [**[]MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md) | Default managed app policies. | [optional] 
**IosManagedAppProtections** | Pointer to [**[]MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md) | iOS managed app policies. | [optional] 
**ManagedAppPolicies** | Pointer to [**[]MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md) | Managed app policies. | [optional] 
**ManagedAppRegistrations** | Pointer to [**[]MicrosoftGraphManagedAppRegistration**](MicrosoftGraphManagedAppRegistration.md) | The managed app registrations. | [optional] 
**ManagedAppStatuses** | Pointer to [**[]MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md) | The managed app statuses. | [optional] 
**MdmWindowsInformationProtectionPolicies** | Pointer to [**[]MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md) | Windows information protection for apps running on devices which are MDM enrolled. | [optional] 
**TargetedManagedAppConfigurations** | Pointer to [**[]MicrosoftGraphTargetedManagedAppConfiguration**](MicrosoftGraphTargetedManagedAppConfiguration.md) | Targeted managed app configurations. | [optional] 
**WindowsInformationProtectionPolicies** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md) | Windows information protection for apps running on devices which are not MDM enrolled. | [optional] 

## Methods

### NewMicrosoftGraphDeviceAppManagement

`func NewMicrosoftGraphDeviceAppManagement() *MicrosoftGraphDeviceAppManagement`

NewMicrosoftGraphDeviceAppManagement instantiates a new MicrosoftGraphDeviceAppManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceAppManagementWithDefaults

`func NewMicrosoftGraphDeviceAppManagementWithDefaults() *MicrosoftGraphDeviceAppManagement`

NewMicrosoftGraphDeviceAppManagementWithDefaults instantiates a new MicrosoftGraphDeviceAppManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceAppManagement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceAppManagement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceAppManagement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceAppManagement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabledForMicrosoftStoreForBusiness

`func (o *MicrosoftGraphDeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness() bool`

GetIsEnabledForMicrosoftStoreForBusiness returns the IsEnabledForMicrosoftStoreForBusiness field if non-nil, zero value otherwise.

### GetIsEnabledForMicrosoftStoreForBusinessOk

`func (o *MicrosoftGraphDeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusinessOk() (*bool, bool)`

GetIsEnabledForMicrosoftStoreForBusinessOk returns a tuple with the IsEnabledForMicrosoftStoreForBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledForMicrosoftStoreForBusiness

`func (o *MicrosoftGraphDeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(v bool)`

SetIsEnabledForMicrosoftStoreForBusiness sets IsEnabledForMicrosoftStoreForBusiness field to given value.

### HasIsEnabledForMicrosoftStoreForBusiness

`func (o *MicrosoftGraphDeviceAppManagement) HasIsEnabledForMicrosoftStoreForBusiness() bool`

HasIsEnabledForMicrosoftStoreForBusiness returns a boolean if a field has been set.

### GetMicrosoftStoreForBusinessLanguage

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLanguage() string`

GetMicrosoftStoreForBusinessLanguage returns the MicrosoftStoreForBusinessLanguage field if non-nil, zero value otherwise.

### GetMicrosoftStoreForBusinessLanguageOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLanguageOk() (*string, bool)`

GetMicrosoftStoreForBusinessLanguageOk returns a tuple with the MicrosoftStoreForBusinessLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftStoreForBusinessLanguage

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(v string)`

SetMicrosoftStoreForBusinessLanguage sets MicrosoftStoreForBusinessLanguage field to given value.

### HasMicrosoftStoreForBusinessLanguage

`func (o *MicrosoftGraphDeviceAppManagement) HasMicrosoftStoreForBusinessLanguage() bool`

HasMicrosoftStoreForBusinessLanguage returns a boolean if a field has been set.

### SetMicrosoftStoreForBusinessLanguageNil

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLanguageNil(b bool)`

 SetMicrosoftStoreForBusinessLanguageNil sets the value for MicrosoftStoreForBusinessLanguage to be an explicit nil

### UnsetMicrosoftStoreForBusinessLanguage
`func (o *MicrosoftGraphDeviceAppManagement) UnsetMicrosoftStoreForBusinessLanguage()`

UnsetMicrosoftStoreForBusinessLanguage ensures that no value is present for MicrosoftStoreForBusinessLanguage, not even an explicit nil
### GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime() time.Time`

GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime returns the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field if non-nil, zero value otherwise.

### GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk() (*time.Time, bool)`

GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk returns a tuple with the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(v time.Time)`

SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime sets MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field to given value.

### HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime

`func (o *MicrosoftGraphDeviceAppManagement) HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime() bool`

HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime returns a boolean if a field has been set.

### GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime() time.Time`

GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime returns the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field if non-nil, zero value otherwise.

### GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk() (*time.Time, bool)`

GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk returns a tuple with the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(v time.Time)`

SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime sets MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field to given value.

### HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphDeviceAppManagement) HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime() bool`

HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime returns a boolean if a field has been set.

### GetManagedEBooks

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedEBooks() []MicrosoftGraphManagedEBook`

GetManagedEBooks returns the ManagedEBooks field if non-nil, zero value otherwise.

### GetManagedEBooksOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedEBooksOk() (*[]MicrosoftGraphManagedEBook, bool)`

GetManagedEBooksOk returns a tuple with the ManagedEBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedEBooks

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedEBooks(v []MicrosoftGraphManagedEBook)`

SetManagedEBooks sets ManagedEBooks field to given value.

### HasManagedEBooks

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedEBooks() bool`

HasManagedEBooks returns a boolean if a field has been set.

### GetMobileAppCategories

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppCategories() []MicrosoftGraphMobileAppCategory`

GetMobileAppCategories returns the MobileAppCategories field if non-nil, zero value otherwise.

### GetMobileAppCategoriesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppCategoriesOk() (*[]MicrosoftGraphMobileAppCategory, bool)`

GetMobileAppCategoriesOk returns a tuple with the MobileAppCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppCategories

`func (o *MicrosoftGraphDeviceAppManagement) SetMobileAppCategories(v []MicrosoftGraphMobileAppCategory)`

SetMobileAppCategories sets MobileAppCategories field to given value.

### HasMobileAppCategories

`func (o *MicrosoftGraphDeviceAppManagement) HasMobileAppCategories() bool`

HasMobileAppCategories returns a boolean if a field has been set.

### GetMobileAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppConfigurations() []MicrosoftGraphManagedDeviceMobileAppConfiguration`

GetMobileAppConfigurations returns the MobileAppConfigurations field if non-nil, zero value otherwise.

### GetMobileAppConfigurationsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppConfigurationsOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfiguration, bool)`

GetMobileAppConfigurationsOk returns a tuple with the MobileAppConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) SetMobileAppConfigurations(v []MicrosoftGraphManagedDeviceMobileAppConfiguration)`

SetMobileAppConfigurations sets MobileAppConfigurations field to given value.

### HasMobileAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) HasMobileAppConfigurations() bool`

HasMobileAppConfigurations returns a boolean if a field has been set.

### GetMobileApps

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileApps() []MicrosoftGraphMobileApp`

GetMobileApps returns the MobileApps field if non-nil, zero value otherwise.

### GetMobileAppsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppsOk() (*[]MicrosoftGraphMobileApp, bool)`

GetMobileAppsOk returns a tuple with the MobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileApps

`func (o *MicrosoftGraphDeviceAppManagement) SetMobileApps(v []MicrosoftGraphMobileApp)`

SetMobileApps sets MobileApps field to given value.

### HasMobileApps

`func (o *MicrosoftGraphDeviceAppManagement) HasMobileApps() bool`

HasMobileApps returns a boolean if a field has been set.

### GetVppTokens

`func (o *MicrosoftGraphDeviceAppManagement) GetVppTokens() []MicrosoftGraphVppToken`

GetVppTokens returns the VppTokens field if non-nil, zero value otherwise.

### GetVppTokensOk

`func (o *MicrosoftGraphDeviceAppManagement) GetVppTokensOk() (*[]MicrosoftGraphVppToken, bool)`

GetVppTokensOk returns a tuple with the VppTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTokens

`func (o *MicrosoftGraphDeviceAppManagement) SetVppTokens(v []MicrosoftGraphVppToken)`

SetVppTokens sets VppTokens field to given value.

### HasVppTokens

`func (o *MicrosoftGraphDeviceAppManagement) HasVppTokens() bool`

HasVppTokens returns a boolean if a field has been set.

### GetAndroidManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) GetAndroidManagedAppProtections() []MicrosoftGraphAndroidManagedAppProtection`

GetAndroidManagedAppProtections returns the AndroidManagedAppProtections field if non-nil, zero value otherwise.

### GetAndroidManagedAppProtectionsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetAndroidManagedAppProtectionsOk() (*[]MicrosoftGraphAndroidManagedAppProtection, bool)`

GetAndroidManagedAppProtectionsOk returns a tuple with the AndroidManagedAppProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) SetAndroidManagedAppProtections(v []MicrosoftGraphAndroidManagedAppProtection)`

SetAndroidManagedAppProtections sets AndroidManagedAppProtections field to given value.

### HasAndroidManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) HasAndroidManagedAppProtections() bool`

HasAndroidManagedAppProtections returns a boolean if a field has been set.

### GetDefaultManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) GetDefaultManagedAppProtections() []MicrosoftGraphDefaultManagedAppProtection`

GetDefaultManagedAppProtections returns the DefaultManagedAppProtections field if non-nil, zero value otherwise.

### GetDefaultManagedAppProtectionsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetDefaultManagedAppProtectionsOk() (*[]MicrosoftGraphDefaultManagedAppProtection, bool)`

GetDefaultManagedAppProtectionsOk returns a tuple with the DefaultManagedAppProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) SetDefaultManagedAppProtections(v []MicrosoftGraphDefaultManagedAppProtection)`

SetDefaultManagedAppProtections sets DefaultManagedAppProtections field to given value.

### HasDefaultManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) HasDefaultManagedAppProtections() bool`

HasDefaultManagedAppProtections returns a boolean if a field has been set.

### GetIosManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) GetIosManagedAppProtections() []MicrosoftGraphIosManagedAppProtection`

GetIosManagedAppProtections returns the IosManagedAppProtections field if non-nil, zero value otherwise.

### GetIosManagedAppProtectionsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetIosManagedAppProtectionsOk() (*[]MicrosoftGraphIosManagedAppProtection, bool)`

GetIosManagedAppProtectionsOk returns a tuple with the IosManagedAppProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) SetIosManagedAppProtections(v []MicrosoftGraphIosManagedAppProtection)`

SetIosManagedAppProtections sets IosManagedAppProtections field to given value.

### HasIosManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) HasIosManagedAppProtections() bool`

HasIosManagedAppProtections returns a boolean if a field has been set.

### GetManagedAppPolicies

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppPolicies() []MicrosoftGraphManagedAppPolicy`

GetManagedAppPolicies returns the ManagedAppPolicies field if non-nil, zero value otherwise.

### GetManagedAppPoliciesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppPoliciesOk() (*[]MicrosoftGraphManagedAppPolicy, bool)`

GetManagedAppPoliciesOk returns a tuple with the ManagedAppPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppPolicies

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedAppPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetManagedAppPolicies sets ManagedAppPolicies field to given value.

### HasManagedAppPolicies

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedAppPolicies() bool`

HasManagedAppPolicies returns a boolean if a field has been set.

### GetManagedAppRegistrations

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppRegistrations() []MicrosoftGraphManagedAppRegistration`

GetManagedAppRegistrations returns the ManagedAppRegistrations field if non-nil, zero value otherwise.

### GetManagedAppRegistrationsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppRegistrationsOk() (*[]MicrosoftGraphManagedAppRegistration, bool)`

GetManagedAppRegistrationsOk returns a tuple with the ManagedAppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppRegistrations

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedAppRegistrations(v []MicrosoftGraphManagedAppRegistration)`

SetManagedAppRegistrations sets ManagedAppRegistrations field to given value.

### HasManagedAppRegistrations

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedAppRegistrations() bool`

HasManagedAppRegistrations returns a boolean if a field has been set.

### GetManagedAppStatuses

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppStatuses() []MicrosoftGraphManagedAppStatus`

GetManagedAppStatuses returns the ManagedAppStatuses field if non-nil, zero value otherwise.

### GetManagedAppStatusesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppStatusesOk() (*[]MicrosoftGraphManagedAppStatus, bool)`

GetManagedAppStatusesOk returns a tuple with the ManagedAppStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppStatuses

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedAppStatuses(v []MicrosoftGraphManagedAppStatus)`

SetManagedAppStatuses sets ManagedAppStatuses field to given value.

### HasManagedAppStatuses

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedAppStatuses() bool`

HasManagedAppStatuses returns a boolean if a field has been set.

### GetMdmWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) GetMdmWindowsInformationProtectionPolicies() []MicrosoftGraphMdmWindowsInformationProtectionPolicy`

GetMdmWindowsInformationProtectionPolicies returns the MdmWindowsInformationProtectionPolicies field if non-nil, zero value otherwise.

### GetMdmWindowsInformationProtectionPoliciesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMdmWindowsInformationProtectionPoliciesOk() (*[]MicrosoftGraphMdmWindowsInformationProtectionPolicy, bool)`

GetMdmWindowsInformationProtectionPoliciesOk returns a tuple with the MdmWindowsInformationProtectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(v []MicrosoftGraphMdmWindowsInformationProtectionPolicy)`

SetMdmWindowsInformationProtectionPolicies sets MdmWindowsInformationProtectionPolicies field to given value.

### HasMdmWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) HasMdmWindowsInformationProtectionPolicies() bool`

HasMdmWindowsInformationProtectionPolicies returns a boolean if a field has been set.

### GetTargetedManagedAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) GetTargetedManagedAppConfigurations() []MicrosoftGraphTargetedManagedAppConfiguration`

GetTargetedManagedAppConfigurations returns the TargetedManagedAppConfigurations field if non-nil, zero value otherwise.

### GetTargetedManagedAppConfigurationsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetTargetedManagedAppConfigurationsOk() (*[]MicrosoftGraphTargetedManagedAppConfiguration, bool)`

GetTargetedManagedAppConfigurationsOk returns a tuple with the TargetedManagedAppConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedManagedAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) SetTargetedManagedAppConfigurations(v []MicrosoftGraphTargetedManagedAppConfiguration)`

SetTargetedManagedAppConfigurations sets TargetedManagedAppConfigurations field to given value.

### HasTargetedManagedAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) HasTargetedManagedAppConfigurations() bool`

HasTargetedManagedAppConfigurations returns a boolean if a field has been set.

### GetWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) GetWindowsInformationProtectionPolicies() []MicrosoftGraphWindowsInformationProtectionPolicy`

GetWindowsInformationProtectionPolicies returns the WindowsInformationProtectionPolicies field if non-nil, zero value otherwise.

### GetWindowsInformationProtectionPoliciesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetWindowsInformationProtectionPoliciesOk() (*[]MicrosoftGraphWindowsInformationProtectionPolicy, bool)`

GetWindowsInformationProtectionPoliciesOk returns a tuple with the WindowsInformationProtectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) SetWindowsInformationProtectionPolicies(v []MicrosoftGraphWindowsInformationProtectionPolicy)`

SetWindowsInformationProtectionPolicies sets WindowsInformationProtectionPolicies field to given value.

### HasWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) HasWindowsInformationProtectionPolicies() bool`

HasWindowsInformationProtectionPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


