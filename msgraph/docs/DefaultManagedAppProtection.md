# DefaultManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDataEncryptionType** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataEncryptionType**](anyOf&lt;microsoft.graph.managedAppDataEncryptionType&gt;.md) | Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked. | [optional] 
**CustomSettings** | Pointer to [**[]MicrosoftGraphKeyValuePair**](MicrosoftGraphKeyValuePair.md) | A set of string key and string value pairs to be sent to the affected users, unalterned by this service | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | Pointer to **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only) | [optional] 
**EncryptAppData** | Pointer to **bool** | Indicates whether managed-app data should be encrypted. (Android only) | [optional] 
**FaceIdBlocked** | Pointer to **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only) | [optional] 
**MinimumRequiredPatchVersion** | Pointer to **NullableString** | Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only) | [optional] 
**MinimumRequiredSdkVersion** | Pointer to **NullableString** | Versions less than the specified version will block the managed app from accessing company data. (iOS Only) | [optional] 
**MinimumWarningPatchVersion** | Pointer to **NullableString** | Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only) | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether screen capture is blocked. (Android only) | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | List of apps to which the policy is deployed. | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) | Navigation property to deployment summary of the configuration. | [optional] 

## Methods

### NewDefaultManagedAppProtection

`func NewDefaultManagedAppProtection() *DefaultManagedAppProtection`

NewDefaultManagedAppProtection instantiates a new DefaultManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultManagedAppProtectionWithDefaults

`func NewDefaultManagedAppProtectionWithDefaults() *DefaultManagedAppProtection`

NewDefaultManagedAppProtectionWithDefaults instantiates a new DefaultManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDataEncryptionType

`func (o *DefaultManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *DefaultManagedAppProtection) GetAppDataEncryptionTypeOk() (*AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataEncryptionType

`func (o *DefaultManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType sets AppDataEncryptionType field to given value.

### HasAppDataEncryptionType

`func (o *DefaultManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionTypeNil

`func (o *DefaultManagedAppProtection) SetAppDataEncryptionTypeNil(b bool)`

 SetAppDataEncryptionTypeNil sets the value for AppDataEncryptionType to be an explicit nil

### UnsetAppDataEncryptionType
`func (o *DefaultManagedAppProtection) UnsetAppDataEncryptionType()`

UnsetAppDataEncryptionType ensures that no value is present for AppDataEncryptionType, not even an explicit nil
### GetCustomSettings

`func (o *DefaultManagedAppProtection) GetCustomSettings() []MicrosoftGraphKeyValuePair`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *DefaultManagedAppProtection) GetCustomSettingsOk() (*[]MicrosoftGraphKeyValuePair, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSettings

`func (o *DefaultManagedAppProtection) SetCustomSettings(v []MicrosoftGraphKeyValuePair)`

SetCustomSettings sets CustomSettings field to given value.

### HasCustomSettings

`func (o *DefaultManagedAppProtection) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.

### GetDeployedAppCount

`func (o *DefaultManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *DefaultManagedAppProtection) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *DefaultManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *DefaultManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (*bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *DefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets DisableAppEncryptionIfDeviceEncryptionIsEnabled field to given value.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *DefaultManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### GetEncryptAppData

`func (o *DefaultManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *DefaultManagedAppProtection) GetEncryptAppDataOk() (*bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAppData

`func (o *DefaultManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData sets EncryptAppData field to given value.

### HasEncryptAppData

`func (o *DefaultManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### GetFaceIdBlocked

`func (o *DefaultManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *DefaultManagedAppProtection) GetFaceIdBlockedOk() (*bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceIdBlocked

`func (o *DefaultManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked sets FaceIdBlocked field to given value.

### HasFaceIdBlocked

`func (o *DefaultManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### GetMinimumRequiredPatchVersion

`func (o *DefaultManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *DefaultManagedAppProtection) GetMinimumRequiredPatchVersionOk() (*string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredPatchVersion

`func (o *DefaultManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion sets MinimumRequiredPatchVersion field to given value.

### HasMinimumRequiredPatchVersion

`func (o *DefaultManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersionNil

`func (o *DefaultManagedAppProtection) SetMinimumRequiredPatchVersionNil(b bool)`

 SetMinimumRequiredPatchVersionNil sets the value for MinimumRequiredPatchVersion to be an explicit nil

### UnsetMinimumRequiredPatchVersion
`func (o *DefaultManagedAppProtection) UnsetMinimumRequiredPatchVersion()`

UnsetMinimumRequiredPatchVersion ensures that no value is present for MinimumRequiredPatchVersion, not even an explicit nil
### GetMinimumRequiredSdkVersion

`func (o *DefaultManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *DefaultManagedAppProtection) GetMinimumRequiredSdkVersionOk() (*string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredSdkVersion

`func (o *DefaultManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion sets MinimumRequiredSdkVersion field to given value.

### HasMinimumRequiredSdkVersion

`func (o *DefaultManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersionNil

`func (o *DefaultManagedAppProtection) SetMinimumRequiredSdkVersionNil(b bool)`

 SetMinimumRequiredSdkVersionNil sets the value for MinimumRequiredSdkVersion to be an explicit nil

### UnsetMinimumRequiredSdkVersion
`func (o *DefaultManagedAppProtection) UnsetMinimumRequiredSdkVersion()`

UnsetMinimumRequiredSdkVersion ensures that no value is present for MinimumRequiredSdkVersion, not even an explicit nil
### GetMinimumWarningPatchVersion

`func (o *DefaultManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *DefaultManagedAppProtection) GetMinimumWarningPatchVersionOk() (*string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningPatchVersion

`func (o *DefaultManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion sets MinimumWarningPatchVersion field to given value.

### HasMinimumWarningPatchVersion

`func (o *DefaultManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersionNil

`func (o *DefaultManagedAppProtection) SetMinimumWarningPatchVersionNil(b bool)`

 SetMinimumWarningPatchVersionNil sets the value for MinimumWarningPatchVersion to be an explicit nil

### UnsetMinimumWarningPatchVersion
`func (o *DefaultManagedAppProtection) UnsetMinimumWarningPatchVersion()`

UnsetMinimumWarningPatchVersion ensures that no value is present for MinimumWarningPatchVersion, not even an explicit nil
### GetScreenCaptureBlocked

`func (o *DefaultManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *DefaultManagedAppProtection) GetScreenCaptureBlockedOk() (*bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenCaptureBlocked

`func (o *DefaultManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked sets ScreenCaptureBlocked field to given value.

### HasScreenCaptureBlocked

`func (o *DefaultManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### GetApps

`func (o *DefaultManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *DefaultManagedAppProtection) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *DefaultManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *DefaultManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *DefaultManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *DefaultManagedAppProtection) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *DefaultManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *DefaultManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *DefaultManagedAppProtection) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *DefaultManagedAppProtection) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


