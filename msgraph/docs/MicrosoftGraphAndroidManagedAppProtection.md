# MicrosoftGraphAndroidManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time the policy was created. | [optional] 
**Description** | Pointer to **NullableString** | The policy&#39;s description. | [optional] 
**DisplayName** | Pointer to **string** | Policy display name. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Last time the policy was modified. | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 
**AllowedDataStorageLocations** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppDataStorageLocation**](AnyOfmicrosoftGraphManagedAppDataStorageLocation.md) | Data storage locations where a user may store managed data. | [optional] 
**AllowedInboundDataTransferSources** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataTransferLevel**](anyOf&lt;microsoft.graph.managedAppDataTransferLevel&gt;.md) | Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none. | [optional] 
**AllowedOutboundClipboardSharingLevel** | Pointer to [**AnyOfmicrosoftGraphManagedAppClipboardSharingLevel**](anyOf&lt;microsoft.graph.managedAppClipboardSharingLevel&gt;.md) | The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked. | [optional] 
**AllowedOutboundDataTransferDestinations** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataTransferLevel**](anyOf&lt;microsoft.graph.managedAppDataTransferLevel&gt;.md) | Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none. | [optional] 
**ContactSyncBlocked** | Pointer to **bool** | Indicates whether contacts can be synced to the user&#39;s device. | [optional] 
**DataBackupBlocked** | Pointer to **bool** | Indicates whether the backup of a managed app&#39;s data is blocked. | [optional] 
**DeviceComplianceRequired** | Pointer to **bool** | Indicates whether device compliance is required. | [optional] 
**DisableAppPinIfDevicePinIsSet** | Pointer to **bool** | Indicates whether use of the app pin is required if the device pin is set. | [optional] 
**FingerprintBlocked** | Pointer to **bool** | Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True. | [optional] 
**ManagedBrowser** | Pointer to [**AnyOfmicrosoftGraphManagedBrowserType**](anyOf&lt;microsoft.graph.managedBrowserType&gt;.md) | Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge. | [optional] 
**ManagedBrowserToOpenLinksRequired** | Pointer to **bool** | Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android) | [optional] 
**MaximumPinRetries** | Pointer to **int32** | Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped. | [optional] 
**MinimumPinLength** | Pointer to **int32** | Minimum pin length required for an app-level pin if PinRequired is set to True | [optional] 
**MinimumRequiredAppVersion** | Pointer to **NullableString** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**MinimumRequiredOsVersion** | Pointer to **NullableString** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**MinimumWarningAppVersion** | Pointer to **NullableString** | Versions less than the specified version will result in warning message on the managed app. | [optional] 
**MinimumWarningOsVersion** | Pointer to **NullableString** | Versions less than the specified version will result in warning message on the managed app from accessing company data. | [optional] 
**OrganizationalCredentialsRequired** | Pointer to **bool** | Indicates whether organizational credentials are required for app use. | [optional] 
**PeriodBeforePinReset** | Pointer to **string** | TimePeriod before the all-level pin must be reset if PinRequired is set to True. | [optional] 
**PeriodOfflineBeforeAccessCheck** | Pointer to **string** | The period after which access is checked when the device is not connected to the internet. | [optional] 
**PeriodOfflineBeforeWipeIsEnforced** | Pointer to **string** | The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. | [optional] 
**PeriodOnlineBeforeAccessCheck** | Pointer to **string** | The period after which access is checked when the device is connected to the internet. | [optional] 
**PinCharacterSet** | Pointer to [**AnyOfmicrosoftGraphManagedAppPinCharacterSet**](anyOf&lt;microsoft.graph.managedAppPinCharacterSet&gt;.md) | Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol. | [optional] 
**PinRequired** | Pointer to **bool** | Indicates whether an app-level pin is required. | [optional] 
**PrintBlocked** | Pointer to **bool** | Indicates whether printing is allowed from managed apps. | [optional] 
**SaveAsBlocked** | Pointer to **bool** | Indicates whether users may use the &#39;Save As&#39; menu item to save a copy of protected files. | [optional] 
**SimplePinBlocked** | Pointer to **bool** | Indicates whether simplePin is blocked. | [optional] 
**IsAssigned** | Pointer to **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md) | Navigation property to list of inclusion and exclusion groups to which the policy is deployed. | [optional] 
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

### NewMicrosoftGraphAndroidManagedAppProtection

`func NewMicrosoftGraphAndroidManagedAppProtection() *MicrosoftGraphAndroidManagedAppProtection`

NewMicrosoftGraphAndroidManagedAppProtection instantiates a new MicrosoftGraphAndroidManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAndroidManagedAppProtectionWithDefaults

`func NewMicrosoftGraphAndroidManagedAppProtectionWithDefaults() *MicrosoftGraphAndroidManagedAppProtection`

NewMicrosoftGraphAndroidManagedAppProtectionWithDefaults instantiates a new MicrosoftGraphAndroidManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedDataStorageLocationsOk() (*[]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations sets AllowedDataStorageLocations field to given value.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources sets AllowedInboundDataTransferSources field to given value.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSourcesNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedInboundDataTransferSourcesNil(b bool)`

 SetAllowedInboundDataTransferSourcesNil sets the value for AllowedInboundDataTransferSources to be an explicit nil

### UnsetAllowedInboundDataTransferSources
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetAllowedInboundDataTransferSources()`

UnsetAllowedInboundDataTransferSources ensures that no value is present for AllowedInboundDataTransferSources, not even an explicit nil
### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (*AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel sets AllowedOutboundClipboardSharingLevel field to given value.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevelNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedOutboundClipboardSharingLevelNil(b bool)`

 SetAllowedOutboundClipboardSharingLevelNil sets the value for AllowedOutboundClipboardSharingLevel to be an explicit nil

### UnsetAllowedOutboundClipboardSharingLevel
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetAllowedOutboundClipboardSharingLevel()`

UnsetAllowedOutboundClipboardSharingLevel ensures that no value is present for AllowedOutboundClipboardSharingLevel, not even an explicit nil
### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations sets AllowedOutboundDataTransferDestinations field to given value.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinationsNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedOutboundDataTransferDestinationsNil(b bool)`

 SetAllowedOutboundDataTransferDestinationsNil sets the value for AllowedOutboundDataTransferDestinations to be an explicit nil

### UnsetAllowedOutboundDataTransferDestinations
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetAllowedOutboundDataTransferDestinations()`

UnsetAllowedOutboundDataTransferDestinations ensures that no value is present for AllowedOutboundDataTransferDestinations, not even an explicit nil
### GetContactSyncBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetContactSyncBlockedOk() (*bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked sets ContactSyncBlocked field to given value.

### HasContactSyncBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### GetDataBackupBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDataBackupBlockedOk() (*bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked sets DataBackupBlocked field to given value.

### HasDataBackupBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeviceComplianceRequiredOk() (*bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired sets DeviceComplianceRequired field to given value.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (*bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet sets DisableAppPinIfDevicePinIsSet field to given value.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### GetFingerprintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetFingerprintBlockedOk() (*bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked sets FingerprintBlocked field to given value.

### HasFingerprintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### GetManagedBrowser

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetManagedBrowser() AnyOfmicrosoftGraphManagedBrowserType`

GetManagedBrowser returns the ManagedBrowser field if non-nil, zero value otherwise.

### GetManagedBrowserOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetManagedBrowserOk() (*AnyOfmicrosoftGraphManagedBrowserType, bool)`

GetManagedBrowserOk returns a tuple with the ManagedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowser

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetManagedBrowser(v AnyOfmicrosoftGraphManagedBrowserType)`

SetManagedBrowser sets ManagedBrowser field to given value.

### HasManagedBrowser

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasManagedBrowser() bool`

HasManagedBrowser returns a boolean if a field has been set.

### SetManagedBrowserNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetManagedBrowserNil(b bool)`

 SetManagedBrowserNil sets the value for ManagedBrowser to be an explicit nil

### UnsetManagedBrowser
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetManagedBrowser()`

UnsetManagedBrowser ensures that no value is present for ManagedBrowser, not even an explicit nil
### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (*bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired sets ManagedBrowserToOpenLinksRequired field to given value.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### GetMaximumPinRetries

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMaximumPinRetriesOk() (*int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries sets MaximumPinRetries field to given value.

### HasMaximumPinRetries

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### GetMinimumPinLength

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumPinLengthOk() (*int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength sets MinimumPinLength field to given value.

### HasMinimumPinLength

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredAppVersionOk() (*string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion sets MinimumRequiredAppVersion field to given value.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredAppVersionNil(b bool)`

 SetMinimumRequiredAppVersionNil sets the value for MinimumRequiredAppVersion to be an explicit nil

### UnsetMinimumRequiredAppVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetMinimumRequiredAppVersion()`

UnsetMinimumRequiredAppVersion ensures that no value is present for MinimumRequiredAppVersion, not even an explicit nil
### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredOsVersionOk() (*string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion sets MinimumRequiredOsVersion field to given value.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredOsVersionNil(b bool)`

 SetMinimumRequiredOsVersionNil sets the value for MinimumRequiredOsVersion to be an explicit nil

### UnsetMinimumRequiredOsVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetMinimumRequiredOsVersion()`

UnsetMinimumRequiredOsVersion ensures that no value is present for MinimumRequiredOsVersion, not even an explicit nil
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningAppVersionOk() (*string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion sets MinimumWarningAppVersion field to given value.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningAppVersionNil(b bool)`

 SetMinimumWarningAppVersionNil sets the value for MinimumWarningAppVersion to be an explicit nil

### UnsetMinimumWarningAppVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetMinimumWarningAppVersion()`

UnsetMinimumWarningAppVersion ensures that no value is present for MinimumWarningAppVersion, not even an explicit nil
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningOsVersionOk() (*string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion sets MinimumWarningOsVersion field to given value.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningOsVersionNil(b bool)`

 SetMinimumWarningOsVersionNil sets the value for MinimumWarningOsVersion to be an explicit nil

### UnsetMinimumWarningOsVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetMinimumWarningOsVersion()`

UnsetMinimumWarningOsVersion ensures that no value is present for MinimumWarningOsVersion, not even an explicit nil
### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (*bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired sets OrganizationalCredentialsRequired field to given value.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodBeforePinResetOk() (*string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset sets PeriodBeforePinReset field to given value.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck sets PeriodOfflineBeforeAccessCheck field to given value.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (*string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced sets PeriodOfflineBeforeWipeIsEnforced field to given value.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck sets PeriodOnlineBeforeAccessCheck field to given value.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### GetPinCharacterSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinCharacterSetOk() (*AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet sets PinCharacterSet field to given value.

### HasPinCharacterSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSetNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPinCharacterSetNil(b bool)`

 SetPinCharacterSetNil sets the value for PinCharacterSet to be an explicit nil

### UnsetPinCharacterSet
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetPinCharacterSet()`

UnsetPinCharacterSet ensures that no value is present for PinCharacterSet, not even an explicit nil
### GetPinRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinRequiredOk() (*bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired sets PinRequired field to given value.

### HasPinRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### GetPrintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPrintBlockedOk() (*bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked sets PrintBlocked field to given value.

### HasPrintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### GetSaveAsBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSaveAsBlockedOk() (*bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked sets SaveAsBlocked field to given value.

### HasSaveAsBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### GetSimplePinBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSimplePinBlockedOk() (*bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked sets SimplePinBlocked field to given value.

### HasSimplePinBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### GetIsAssigned

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetCustomBrowserDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCustomBrowserDisplayName() string`

GetCustomBrowserDisplayName returns the CustomBrowserDisplayName field if non-nil, zero value otherwise.

### GetCustomBrowserDisplayNameOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCustomBrowserDisplayNameOk() (*string, bool)`

GetCustomBrowserDisplayNameOk returns a tuple with the CustomBrowserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBrowserDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetCustomBrowserDisplayName(v string)`

SetCustomBrowserDisplayName sets CustomBrowserDisplayName field to given value.

### HasCustomBrowserDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasCustomBrowserDisplayName() bool`

HasCustomBrowserDisplayName returns a boolean if a field has been set.

### SetCustomBrowserDisplayNameNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetCustomBrowserDisplayNameNil(b bool)`

 SetCustomBrowserDisplayNameNil sets the value for CustomBrowserDisplayName to be an explicit nil

### UnsetCustomBrowserDisplayName
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetCustomBrowserDisplayName()`

UnsetCustomBrowserDisplayName ensures that no value is present for CustomBrowserDisplayName, not even an explicit nil
### GetCustomBrowserPackageId

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCustomBrowserPackageId() string`

GetCustomBrowserPackageId returns the CustomBrowserPackageId field if non-nil, zero value otherwise.

### GetCustomBrowserPackageIdOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCustomBrowserPackageIdOk() (*string, bool)`

GetCustomBrowserPackageIdOk returns a tuple with the CustomBrowserPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBrowserPackageId

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetCustomBrowserPackageId(v string)`

SetCustomBrowserPackageId sets CustomBrowserPackageId field to given value.

### HasCustomBrowserPackageId

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasCustomBrowserPackageId() bool`

HasCustomBrowserPackageId returns a boolean if a field has been set.

### SetCustomBrowserPackageIdNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetCustomBrowserPackageIdNil(b bool)`

 SetCustomBrowserPackageIdNil sets the value for CustomBrowserPackageId to be an explicit nil

### UnsetCustomBrowserPackageId
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetCustomBrowserPackageId()`

UnsetCustomBrowserPackageId ensures that no value is present for CustomBrowserPackageId, not even an explicit nil
### GetDeployedAppCount

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (*bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets DisableAppEncryptionIfDeviceEncryptionIsEnabled field to given value.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### GetEncryptAppData

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetEncryptAppDataOk() (*bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAppData

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData sets EncryptAppData field to given value.

### HasEncryptAppData

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### GetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredPatchVersionOk() (*string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion sets MinimumRequiredPatchVersion field to given value.

### HasMinimumRequiredPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredPatchVersionNil(b bool)`

 SetMinimumRequiredPatchVersionNil sets the value for MinimumRequiredPatchVersion to be an explicit nil

### UnsetMinimumRequiredPatchVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetMinimumRequiredPatchVersion()`

UnsetMinimumRequiredPatchVersion ensures that no value is present for MinimumRequiredPatchVersion, not even an explicit nil
### GetMinimumWarningPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningPatchVersionOk() (*string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion sets MinimumWarningPatchVersion field to given value.

### HasMinimumWarningPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersionNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningPatchVersionNil(b bool)`

 SetMinimumWarningPatchVersionNil sets the value for MinimumWarningPatchVersion to be an explicit nil

### UnsetMinimumWarningPatchVersion
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetMinimumWarningPatchVersion()`

UnsetMinimumWarningPatchVersion ensures that no value is present for MinimumWarningPatchVersion, not even an explicit nil
### GetScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetScreenCaptureBlockedOk() (*bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked sets ScreenCaptureBlocked field to given value.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### GetApps

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *MicrosoftGraphAndroidManagedAppProtection) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


