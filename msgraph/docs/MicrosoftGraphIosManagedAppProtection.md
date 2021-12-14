# MicrosoftGraphIosManagedAppProtection

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
**AppDataEncryptionType** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataEncryptionType**](anyOf&lt;microsoft.graph.managedAppDataEncryptionType&gt;.md) | Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked. | [optional] 
**CustomBrowserProtocol** | Pointer to **NullableString** | A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**FaceIdBlocked** | Pointer to **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. | [optional] 
**MinimumRequiredSdkVersion** | Pointer to **NullableString** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | List of apps to which the policy is deployed. | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) | Navigation property to deployment summary of the configuration. | [optional] 

## Methods

### NewMicrosoftGraphIosManagedAppProtection

`func NewMicrosoftGraphIosManagedAppProtection() *MicrosoftGraphIosManagedAppProtection`

NewMicrosoftGraphIosManagedAppProtection instantiates a new MicrosoftGraphIosManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIosManagedAppProtectionWithDefaults

`func NewMicrosoftGraphIosManagedAppProtectionWithDefaults() *MicrosoftGraphIosManagedAppProtection`

NewMicrosoftGraphIosManagedAppProtectionWithDefaults instantiates a new MicrosoftGraphIosManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIosManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIosManagedAppProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIosManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphIosManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphIosManagedAppProtection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphIosManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosManagedAppProtection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphIosManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedDataStorageLocationsOk() (*[]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations sets AllowedDataStorageLocations field to given value.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources sets AllowedInboundDataTransferSources field to given value.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSourcesNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedInboundDataTransferSourcesNil(b bool)`

 SetAllowedInboundDataTransferSourcesNil sets the value for AllowedInboundDataTransferSources to be an explicit nil

### UnsetAllowedInboundDataTransferSources
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetAllowedInboundDataTransferSources()`

UnsetAllowedInboundDataTransferSources ensures that no value is present for AllowedInboundDataTransferSources, not even an explicit nil
### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (*AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel sets AllowedOutboundClipboardSharingLevel field to given value.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevelNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedOutboundClipboardSharingLevelNil(b bool)`

 SetAllowedOutboundClipboardSharingLevelNil sets the value for AllowedOutboundClipboardSharingLevel to be an explicit nil

### UnsetAllowedOutboundClipboardSharingLevel
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetAllowedOutboundClipboardSharingLevel()`

UnsetAllowedOutboundClipboardSharingLevel ensures that no value is present for AllowedOutboundClipboardSharingLevel, not even an explicit nil
### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations sets AllowedOutboundDataTransferDestinations field to given value.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinationsNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedOutboundDataTransferDestinationsNil(b bool)`

 SetAllowedOutboundDataTransferDestinationsNil sets the value for AllowedOutboundDataTransferDestinations to be an explicit nil

### UnsetAllowedOutboundDataTransferDestinations
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetAllowedOutboundDataTransferDestinations()`

UnsetAllowedOutboundDataTransferDestinations ensures that no value is present for AllowedOutboundDataTransferDestinations, not even an explicit nil
### GetContactSyncBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetContactSyncBlockedOk() (*bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked sets ContactSyncBlocked field to given value.

### HasContactSyncBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### GetDataBackupBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDataBackupBlockedOk() (*bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked sets DataBackupBlocked field to given value.

### HasDataBackupBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeviceComplianceRequiredOk() (*bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired sets DeviceComplianceRequired field to given value.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (*bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphIosManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet sets DisableAppPinIfDevicePinIsSet field to given value.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphIosManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### GetFingerprintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetFingerprintBlockedOk() (*bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked sets FingerprintBlocked field to given value.

### HasFingerprintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### GetManagedBrowser

`func (o *MicrosoftGraphIosManagedAppProtection) GetManagedBrowser() AnyOfmicrosoftGraphManagedBrowserType`

GetManagedBrowser returns the ManagedBrowser field if non-nil, zero value otherwise.

### GetManagedBrowserOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetManagedBrowserOk() (*AnyOfmicrosoftGraphManagedBrowserType, bool)`

GetManagedBrowserOk returns a tuple with the ManagedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowser

`func (o *MicrosoftGraphIosManagedAppProtection) SetManagedBrowser(v AnyOfmicrosoftGraphManagedBrowserType)`

SetManagedBrowser sets ManagedBrowser field to given value.

### HasManagedBrowser

`func (o *MicrosoftGraphIosManagedAppProtection) HasManagedBrowser() bool`

HasManagedBrowser returns a boolean if a field has been set.

### SetManagedBrowserNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetManagedBrowserNil(b bool)`

 SetManagedBrowserNil sets the value for ManagedBrowser to be an explicit nil

### UnsetManagedBrowser
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetManagedBrowser()`

UnsetManagedBrowser ensures that no value is present for ManagedBrowser, not even an explicit nil
### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (*bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired sets ManagedBrowserToOpenLinksRequired field to given value.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### GetMaximumPinRetries

`func (o *MicrosoftGraphIosManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMaximumPinRetriesOk() (*int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphIosManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries sets MaximumPinRetries field to given value.

### HasMaximumPinRetries

`func (o *MicrosoftGraphIosManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### GetMinimumPinLength

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumPinLengthOk() (*int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength sets MinimumPinLength field to given value.

### HasMinimumPinLength

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredAppVersionOk() (*string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion sets MinimumRequiredAppVersion field to given value.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredAppVersionNil(b bool)`

 SetMinimumRequiredAppVersionNil sets the value for MinimumRequiredAppVersion to be an explicit nil

### UnsetMinimumRequiredAppVersion
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetMinimumRequiredAppVersion()`

UnsetMinimumRequiredAppVersion ensures that no value is present for MinimumRequiredAppVersion, not even an explicit nil
### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredOsVersionOk() (*string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion sets MinimumRequiredOsVersion field to given value.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredOsVersionNil(b bool)`

 SetMinimumRequiredOsVersionNil sets the value for MinimumRequiredOsVersion to be an explicit nil

### UnsetMinimumRequiredOsVersion
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetMinimumRequiredOsVersion()`

UnsetMinimumRequiredOsVersion ensures that no value is present for MinimumRequiredOsVersion, not even an explicit nil
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningAppVersionOk() (*string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion sets MinimumWarningAppVersion field to given value.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningAppVersionNil(b bool)`

 SetMinimumWarningAppVersionNil sets the value for MinimumWarningAppVersion to be an explicit nil

### UnsetMinimumWarningAppVersion
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetMinimumWarningAppVersion()`

UnsetMinimumWarningAppVersion ensures that no value is present for MinimumWarningAppVersion, not even an explicit nil
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningOsVersionOk() (*string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion sets MinimumWarningOsVersion field to given value.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningOsVersionNil(b bool)`

 SetMinimumWarningOsVersionNil sets the value for MinimumWarningOsVersion to be an explicit nil

### UnsetMinimumWarningOsVersion
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetMinimumWarningOsVersion()`

UnsetMinimumWarningOsVersion ensures that no value is present for MinimumWarningOsVersion, not even an explicit nil
### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (*bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired sets OrganizationalCredentialsRequired field to given value.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodBeforePinResetOk() (*string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset sets PeriodBeforePinReset field to given value.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck sets PeriodOfflineBeforeAccessCheck field to given value.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (*string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced sets PeriodOfflineBeforeWipeIsEnforced field to given value.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck sets PeriodOnlineBeforeAccessCheck field to given value.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### GetPinCharacterSet

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinCharacterSetOk() (*AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphIosManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet sets PinCharacterSet field to given value.

### HasPinCharacterSet

`func (o *MicrosoftGraphIosManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSetNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetPinCharacterSetNil(b bool)`

 SetPinCharacterSetNil sets the value for PinCharacterSet to be an explicit nil

### UnsetPinCharacterSet
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetPinCharacterSet()`

UnsetPinCharacterSet ensures that no value is present for PinCharacterSet, not even an explicit nil
### GetPinRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinRequiredOk() (*bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired sets PinRequired field to given value.

### HasPinRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### GetPrintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPrintBlockedOk() (*bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked sets PrintBlocked field to given value.

### HasPrintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### GetSaveAsBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetSaveAsBlockedOk() (*bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked sets SaveAsBlocked field to given value.

### HasSaveAsBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### GetSimplePinBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetSimplePinBlockedOk() (*bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked sets SimplePinBlocked field to given value.

### HasSimplePinBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### GetIsAssigned

`func (o *MicrosoftGraphIosManagedAppProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MicrosoftGraphIosManagedAppProtection) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MicrosoftGraphIosManagedAppProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphIosManagedAppProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphIosManagedAppProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphIosManagedAppProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAppDataEncryptionType

`func (o *MicrosoftGraphIosManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAppDataEncryptionTypeOk() (*AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataEncryptionType

`func (o *MicrosoftGraphIosManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType sets AppDataEncryptionType field to given value.

### HasAppDataEncryptionType

`func (o *MicrosoftGraphIosManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionTypeNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetAppDataEncryptionTypeNil(b bool)`

 SetAppDataEncryptionTypeNil sets the value for AppDataEncryptionType to be an explicit nil

### UnsetAppDataEncryptionType
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetAppDataEncryptionType()`

UnsetAppDataEncryptionType ensures that no value is present for AppDataEncryptionType, not even an explicit nil
### GetCustomBrowserProtocol

`func (o *MicrosoftGraphIosManagedAppProtection) GetCustomBrowserProtocol() string`

GetCustomBrowserProtocol returns the CustomBrowserProtocol field if non-nil, zero value otherwise.

### GetCustomBrowserProtocolOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetCustomBrowserProtocolOk() (*string, bool)`

GetCustomBrowserProtocolOk returns a tuple with the CustomBrowserProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBrowserProtocol

`func (o *MicrosoftGraphIosManagedAppProtection) SetCustomBrowserProtocol(v string)`

SetCustomBrowserProtocol sets CustomBrowserProtocol field to given value.

### HasCustomBrowserProtocol

`func (o *MicrosoftGraphIosManagedAppProtection) HasCustomBrowserProtocol() bool`

HasCustomBrowserProtocol returns a boolean if a field has been set.

### SetCustomBrowserProtocolNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetCustomBrowserProtocolNil(b bool)`

 SetCustomBrowserProtocolNil sets the value for CustomBrowserProtocol to be an explicit nil

### UnsetCustomBrowserProtocol
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetCustomBrowserProtocol()`

UnsetCustomBrowserProtocol ensures that no value is present for CustomBrowserProtocol, not even an explicit nil
### GetDeployedAppCount

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *MicrosoftGraphIosManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetFaceIdBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetFaceIdBlockedOk() (*bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceIdBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked sets FaceIdBlocked field to given value.

### HasFaceIdBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### GetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredSdkVersionOk() (*string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion sets MinimumRequiredSdkVersion field to given value.

### HasMinimumRequiredSdkVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersionNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredSdkVersionNil(b bool)`

 SetMinimumRequiredSdkVersionNil sets the value for MinimumRequiredSdkVersion to be an explicit nil

### UnsetMinimumRequiredSdkVersion
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetMinimumRequiredSdkVersion()`

UnsetMinimumRequiredSdkVersion ensures that no value is present for MinimumRequiredSdkVersion, not even an explicit nil
### GetApps

`func (o *MicrosoftGraphIosManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *MicrosoftGraphIosManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *MicrosoftGraphIosManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *MicrosoftGraphIosManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *MicrosoftGraphIosManagedAppProtection) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


