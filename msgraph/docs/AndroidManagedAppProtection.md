# AndroidManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomBrowserDisplayName** | Pointer to **NullableString** | Friendly name of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**CustomBrowserPackageId** | Pointer to **NullableString** | Unique identifier of the preferred custom browser to open weblink on Android. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | Pointer to **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled | [optional] 
**EncryptAppData** | Pointer to **bool** | Indicates whether application data for managed apps should be encrypted | [optional] 
**MinimumRequiredPatchVersion** | Pointer to **NullableString** | Define the oldest required Android security patch level a user can have to gain secure access to the app. | [optional] 
**MinimumWarningPatchVersion** | Pointer to **NullableString** | Define the oldest recommended Android security patch level a user can have for secure access to the app. | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether a managed user can take screen captures of managed apps | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | List of apps to which the policy is deployed. | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) | Navigation property to deployment summary of the configuration. | [optional] 

## Methods

### NewAndroidManagedAppProtection

`func NewAndroidManagedAppProtection() *AndroidManagedAppProtection`

NewAndroidManagedAppProtection instantiates a new AndroidManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidManagedAppProtectionWithDefaults

`func NewAndroidManagedAppProtectionWithDefaults() *AndroidManagedAppProtection`

NewAndroidManagedAppProtectionWithDefaults instantiates a new AndroidManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomBrowserDisplayName

`func (o *AndroidManagedAppProtection) GetCustomBrowserDisplayName() string`

GetCustomBrowserDisplayName returns the CustomBrowserDisplayName field if non-nil, zero value otherwise.

### GetCustomBrowserDisplayNameOk

`func (o *AndroidManagedAppProtection) GetCustomBrowserDisplayNameOk() (*string, bool)`

GetCustomBrowserDisplayNameOk returns a tuple with the CustomBrowserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBrowserDisplayName

`func (o *AndroidManagedAppProtection) SetCustomBrowserDisplayName(v string)`

SetCustomBrowserDisplayName sets CustomBrowserDisplayName field to given value.

### HasCustomBrowserDisplayName

`func (o *AndroidManagedAppProtection) HasCustomBrowserDisplayName() bool`

HasCustomBrowserDisplayName returns a boolean if a field has been set.

### SetCustomBrowserDisplayNameNil

`func (o *AndroidManagedAppProtection) SetCustomBrowserDisplayNameNil(b bool)`

 SetCustomBrowserDisplayNameNil sets the value for CustomBrowserDisplayName to be an explicit nil

### UnsetCustomBrowserDisplayName
`func (o *AndroidManagedAppProtection) UnsetCustomBrowserDisplayName()`

UnsetCustomBrowserDisplayName ensures that no value is present for CustomBrowserDisplayName, not even an explicit nil
### GetCustomBrowserPackageId

`func (o *AndroidManagedAppProtection) GetCustomBrowserPackageId() string`

GetCustomBrowserPackageId returns the CustomBrowserPackageId field if non-nil, zero value otherwise.

### GetCustomBrowserPackageIdOk

`func (o *AndroidManagedAppProtection) GetCustomBrowserPackageIdOk() (*string, bool)`

GetCustomBrowserPackageIdOk returns a tuple with the CustomBrowserPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBrowserPackageId

`func (o *AndroidManagedAppProtection) SetCustomBrowserPackageId(v string)`

SetCustomBrowserPackageId sets CustomBrowserPackageId field to given value.

### HasCustomBrowserPackageId

`func (o *AndroidManagedAppProtection) HasCustomBrowserPackageId() bool`

HasCustomBrowserPackageId returns a boolean if a field has been set.

### SetCustomBrowserPackageIdNil

`func (o *AndroidManagedAppProtection) SetCustomBrowserPackageIdNil(b bool)`

 SetCustomBrowserPackageIdNil sets the value for CustomBrowserPackageId to be an explicit nil

### UnsetCustomBrowserPackageId
`func (o *AndroidManagedAppProtection) UnsetCustomBrowserPackageId()`

UnsetCustomBrowserPackageId ensures that no value is present for CustomBrowserPackageId, not even an explicit nil
### GetDeployedAppCount

`func (o *AndroidManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *AndroidManagedAppProtection) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *AndroidManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *AndroidManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (*bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *AndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets DisableAppEncryptionIfDeviceEncryptionIsEnabled field to given value.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *AndroidManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### GetEncryptAppData

`func (o *AndroidManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *AndroidManagedAppProtection) GetEncryptAppDataOk() (*bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAppData

`func (o *AndroidManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData sets EncryptAppData field to given value.

### HasEncryptAppData

`func (o *AndroidManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### GetMinimumRequiredPatchVersion

`func (o *AndroidManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *AndroidManagedAppProtection) GetMinimumRequiredPatchVersionOk() (*string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredPatchVersion

`func (o *AndroidManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion sets MinimumRequiredPatchVersion field to given value.

### HasMinimumRequiredPatchVersion

`func (o *AndroidManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersionNil

`func (o *AndroidManagedAppProtection) SetMinimumRequiredPatchVersionNil(b bool)`

 SetMinimumRequiredPatchVersionNil sets the value for MinimumRequiredPatchVersion to be an explicit nil

### UnsetMinimumRequiredPatchVersion
`func (o *AndroidManagedAppProtection) UnsetMinimumRequiredPatchVersion()`

UnsetMinimumRequiredPatchVersion ensures that no value is present for MinimumRequiredPatchVersion, not even an explicit nil
### GetMinimumWarningPatchVersion

`func (o *AndroidManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *AndroidManagedAppProtection) GetMinimumWarningPatchVersionOk() (*string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningPatchVersion

`func (o *AndroidManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion sets MinimumWarningPatchVersion field to given value.

### HasMinimumWarningPatchVersion

`func (o *AndroidManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersionNil

`func (o *AndroidManagedAppProtection) SetMinimumWarningPatchVersionNil(b bool)`

 SetMinimumWarningPatchVersionNil sets the value for MinimumWarningPatchVersion to be an explicit nil

### UnsetMinimumWarningPatchVersion
`func (o *AndroidManagedAppProtection) UnsetMinimumWarningPatchVersion()`

UnsetMinimumWarningPatchVersion ensures that no value is present for MinimumWarningPatchVersion, not even an explicit nil
### GetScreenCaptureBlocked

`func (o *AndroidManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *AndroidManagedAppProtection) GetScreenCaptureBlockedOk() (*bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenCaptureBlocked

`func (o *AndroidManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked sets ScreenCaptureBlocked field to given value.

### HasScreenCaptureBlocked

`func (o *AndroidManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### GetApps

`func (o *AndroidManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AndroidManagedAppProtection) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AndroidManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AndroidManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *AndroidManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *AndroidManagedAppProtection) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *AndroidManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *AndroidManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *AndroidManagedAppProtection) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *AndroidManagedAppProtection) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


