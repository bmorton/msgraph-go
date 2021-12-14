# MicrosoftGraphDefaultManagedAppProtection

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

### NewMicrosoftGraphDefaultManagedAppProtection

`func NewMicrosoftGraphDefaultManagedAppProtection() *MicrosoftGraphDefaultManagedAppProtection`

NewMicrosoftGraphDefaultManagedAppProtection instantiates a new MicrosoftGraphDefaultManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDefaultManagedAppProtectionWithDefaults

`func NewMicrosoftGraphDefaultManagedAppProtectionWithDefaults() *MicrosoftGraphDefaultManagedAppProtection`

NewMicrosoftGraphDefaultManagedAppProtectionWithDefaults instantiates a new MicrosoftGraphDefaultManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedDataStorageLocationsOk() (*[]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations sets AllowedDataStorageLocations field to given value.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources sets AllowedInboundDataTransferSources field to given value.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSourcesNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedInboundDataTransferSourcesNil(b bool)`

 SetAllowedInboundDataTransferSourcesNil sets the value for AllowedInboundDataTransferSources to be an explicit nil

### UnsetAllowedInboundDataTransferSources
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetAllowedInboundDataTransferSources()`

UnsetAllowedInboundDataTransferSources ensures that no value is present for AllowedInboundDataTransferSources, not even an explicit nil
### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (*AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel sets AllowedOutboundClipboardSharingLevel field to given value.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevelNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedOutboundClipboardSharingLevelNil(b bool)`

 SetAllowedOutboundClipboardSharingLevelNil sets the value for AllowedOutboundClipboardSharingLevel to be an explicit nil

### UnsetAllowedOutboundClipboardSharingLevel
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetAllowedOutboundClipboardSharingLevel()`

UnsetAllowedOutboundClipboardSharingLevel ensures that no value is present for AllowedOutboundClipboardSharingLevel, not even an explicit nil
### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations sets AllowedOutboundDataTransferDestinations field to given value.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinationsNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedOutboundDataTransferDestinationsNil(b bool)`

 SetAllowedOutboundDataTransferDestinationsNil sets the value for AllowedOutboundDataTransferDestinations to be an explicit nil

### UnsetAllowedOutboundDataTransferDestinations
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetAllowedOutboundDataTransferDestinations()`

UnsetAllowedOutboundDataTransferDestinations ensures that no value is present for AllowedOutboundDataTransferDestinations, not even an explicit nil
### GetContactSyncBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetContactSyncBlockedOk() (*bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked sets ContactSyncBlocked field to given value.

### HasContactSyncBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### GetDataBackupBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDataBackupBlockedOk() (*bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked sets DataBackupBlocked field to given value.

### HasDataBackupBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeviceComplianceRequiredOk() (*bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired sets DeviceComplianceRequired field to given value.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (*bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet sets DisableAppPinIfDevicePinIsSet field to given value.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### GetFingerprintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFingerprintBlockedOk() (*bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked sets FingerprintBlocked field to given value.

### HasFingerprintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### GetManagedBrowser

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetManagedBrowser() AnyOfmicrosoftGraphManagedBrowserType`

GetManagedBrowser returns the ManagedBrowser field if non-nil, zero value otherwise.

### GetManagedBrowserOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetManagedBrowserOk() (*AnyOfmicrosoftGraphManagedBrowserType, bool)`

GetManagedBrowserOk returns a tuple with the ManagedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowser

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetManagedBrowser(v AnyOfmicrosoftGraphManagedBrowserType)`

SetManagedBrowser sets ManagedBrowser field to given value.

### HasManagedBrowser

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasManagedBrowser() bool`

HasManagedBrowser returns a boolean if a field has been set.

### SetManagedBrowserNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetManagedBrowserNil(b bool)`

 SetManagedBrowserNil sets the value for ManagedBrowser to be an explicit nil

### UnsetManagedBrowser
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetManagedBrowser()`

UnsetManagedBrowser ensures that no value is present for ManagedBrowser, not even an explicit nil
### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (*bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired sets ManagedBrowserToOpenLinksRequired field to given value.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### GetMaximumPinRetries

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMaximumPinRetriesOk() (*int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries sets MaximumPinRetries field to given value.

### HasMaximumPinRetries

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### GetMinimumPinLength

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumPinLengthOk() (*int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength sets MinimumPinLength field to given value.

### HasMinimumPinLength

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredAppVersionOk() (*string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion sets MinimumRequiredAppVersion field to given value.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredAppVersionNil(b bool)`

 SetMinimumRequiredAppVersionNil sets the value for MinimumRequiredAppVersion to be an explicit nil

### UnsetMinimumRequiredAppVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumRequiredAppVersion()`

UnsetMinimumRequiredAppVersion ensures that no value is present for MinimumRequiredAppVersion, not even an explicit nil
### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredOsVersionOk() (*string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion sets MinimumRequiredOsVersion field to given value.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredOsVersionNil(b bool)`

 SetMinimumRequiredOsVersionNil sets the value for MinimumRequiredOsVersion to be an explicit nil

### UnsetMinimumRequiredOsVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumRequiredOsVersion()`

UnsetMinimumRequiredOsVersion ensures that no value is present for MinimumRequiredOsVersion, not even an explicit nil
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningAppVersionOk() (*string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion sets MinimumWarningAppVersion field to given value.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningAppVersionNil(b bool)`

 SetMinimumWarningAppVersionNil sets the value for MinimumWarningAppVersion to be an explicit nil

### UnsetMinimumWarningAppVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumWarningAppVersion()`

UnsetMinimumWarningAppVersion ensures that no value is present for MinimumWarningAppVersion, not even an explicit nil
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningOsVersionOk() (*string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion sets MinimumWarningOsVersion field to given value.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningOsVersionNil(b bool)`

 SetMinimumWarningOsVersionNil sets the value for MinimumWarningOsVersion to be an explicit nil

### UnsetMinimumWarningOsVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumWarningOsVersion()`

UnsetMinimumWarningOsVersion ensures that no value is present for MinimumWarningOsVersion, not even an explicit nil
### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (*bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired sets OrganizationalCredentialsRequired field to given value.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodBeforePinResetOk() (*string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset sets PeriodBeforePinReset field to given value.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck sets PeriodOfflineBeforeAccessCheck field to given value.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (*string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced sets PeriodOfflineBeforeWipeIsEnforced field to given value.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck sets PeriodOnlineBeforeAccessCheck field to given value.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### GetPinCharacterSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinCharacterSetOk() (*AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet sets PinCharacterSet field to given value.

### HasPinCharacterSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSetNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPinCharacterSetNil(b bool)`

 SetPinCharacterSetNil sets the value for PinCharacterSet to be an explicit nil

### UnsetPinCharacterSet
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetPinCharacterSet()`

UnsetPinCharacterSet ensures that no value is present for PinCharacterSet, not even an explicit nil
### GetPinRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinRequiredOk() (*bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired sets PinRequired field to given value.

### HasPinRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### GetPrintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPrintBlockedOk() (*bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked sets PrintBlocked field to given value.

### HasPrintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### GetSaveAsBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSaveAsBlockedOk() (*bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked sets SaveAsBlocked field to given value.

### HasSaveAsBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### GetSimplePinBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSimplePinBlockedOk() (*bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked sets SimplePinBlocked field to given value.

### HasSimplePinBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### GetAppDataEncryptionType

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAppDataEncryptionTypeOk() (*AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataEncryptionType

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType sets AppDataEncryptionType field to given value.

### HasAppDataEncryptionType

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionTypeNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAppDataEncryptionTypeNil(b bool)`

 SetAppDataEncryptionTypeNil sets the value for AppDataEncryptionType to be an explicit nil

### UnsetAppDataEncryptionType
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetAppDataEncryptionType()`

UnsetAppDataEncryptionType ensures that no value is present for AppDataEncryptionType, not even an explicit nil
### GetCustomSettings

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCustomSettings() []MicrosoftGraphKeyValuePair`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCustomSettingsOk() (*[]MicrosoftGraphKeyValuePair, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSettings

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetCustomSettings(v []MicrosoftGraphKeyValuePair)`

SetCustomSettings sets CustomSettings field to given value.

### HasCustomSettings

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.

### GetDeployedAppCount

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (*bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets DisableAppEncryptionIfDeviceEncryptionIsEnabled field to given value.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### GetEncryptAppData

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetEncryptAppDataOk() (*bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAppData

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData sets EncryptAppData field to given value.

### HasEncryptAppData

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### GetFaceIdBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFaceIdBlockedOk() (*bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceIdBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked sets FaceIdBlocked field to given value.

### HasFaceIdBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### GetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredPatchVersionOk() (*string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion sets MinimumRequiredPatchVersion field to given value.

### HasMinimumRequiredPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredPatchVersionNil(b bool)`

 SetMinimumRequiredPatchVersionNil sets the value for MinimumRequiredPatchVersion to be an explicit nil

### UnsetMinimumRequiredPatchVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumRequiredPatchVersion()`

UnsetMinimumRequiredPatchVersion ensures that no value is present for MinimumRequiredPatchVersion, not even an explicit nil
### GetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredSdkVersionOk() (*string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion sets MinimumRequiredSdkVersion field to given value.

### HasMinimumRequiredSdkVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredSdkVersionNil(b bool)`

 SetMinimumRequiredSdkVersionNil sets the value for MinimumRequiredSdkVersion to be an explicit nil

### UnsetMinimumRequiredSdkVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumRequiredSdkVersion()`

UnsetMinimumRequiredSdkVersion ensures that no value is present for MinimumRequiredSdkVersion, not even an explicit nil
### GetMinimumWarningPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningPatchVersionOk() (*string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion sets MinimumWarningPatchVersion field to given value.

### HasMinimumWarningPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersionNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningPatchVersionNil(b bool)`

 SetMinimumWarningPatchVersionNil sets the value for MinimumWarningPatchVersion to be an explicit nil

### UnsetMinimumWarningPatchVersion
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetMinimumWarningPatchVersion()`

UnsetMinimumWarningPatchVersion ensures that no value is present for MinimumWarningPatchVersion, not even an explicit nil
### GetScreenCaptureBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetScreenCaptureBlockedOk() (*bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked sets ScreenCaptureBlocked field to given value.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### GetApps

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *MicrosoftGraphDefaultManagedAppProtection) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


