# ManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewManagedAppProtection

`func NewManagedAppProtection() *ManagedAppProtection`

NewManagedAppProtection instantiates a new ManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAppProtectionWithDefaults

`func NewManagedAppProtectionWithDefaults() *ManagedAppProtection`

NewManagedAppProtectionWithDefaults instantiates a new ManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDataStorageLocations

`func (o *ManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *ManagedAppProtection) GetAllowedDataStorageLocationsOk() (*[]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataStorageLocations

`func (o *ManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations sets AllowedDataStorageLocations field to given value.

### HasAllowedDataStorageLocations

`func (o *ManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### GetAllowedInboundDataTransferSources

`func (o *ManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *ManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInboundDataTransferSources

`func (o *ManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources sets AllowedInboundDataTransferSources field to given value.

### HasAllowedInboundDataTransferSources

`func (o *ManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSourcesNil

`func (o *ManagedAppProtection) SetAllowedInboundDataTransferSourcesNil(b bool)`

 SetAllowedInboundDataTransferSourcesNil sets the value for AllowedInboundDataTransferSources to be an explicit nil

### UnsetAllowedInboundDataTransferSources
`func (o *ManagedAppProtection) UnsetAllowedInboundDataTransferSources()`

UnsetAllowedInboundDataTransferSources ensures that no value is present for AllowedInboundDataTransferSources, not even an explicit nil
### GetAllowedOutboundClipboardSharingLevel

`func (o *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (*AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel sets AllowedOutboundClipboardSharingLevel field to given value.

### HasAllowedOutboundClipboardSharingLevel

`func (o *ManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevelNil

`func (o *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevelNil(b bool)`

 SetAllowedOutboundClipboardSharingLevelNil sets the value for AllowedOutboundClipboardSharingLevel to be an explicit nil

### UnsetAllowedOutboundClipboardSharingLevel
`func (o *ManagedAppProtection) UnsetAllowedOutboundClipboardSharingLevel()`

UnsetAllowedOutboundClipboardSharingLevel ensures that no value is present for AllowedOutboundClipboardSharingLevel, not even an explicit nil
### GetAllowedOutboundDataTransferDestinations

`func (o *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *ManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations sets AllowedOutboundDataTransferDestinations field to given value.

### HasAllowedOutboundDataTransferDestinations

`func (o *ManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinationsNil

`func (o *ManagedAppProtection) SetAllowedOutboundDataTransferDestinationsNil(b bool)`

 SetAllowedOutboundDataTransferDestinationsNil sets the value for AllowedOutboundDataTransferDestinations to be an explicit nil

### UnsetAllowedOutboundDataTransferDestinations
`func (o *ManagedAppProtection) UnsetAllowedOutboundDataTransferDestinations()`

UnsetAllowedOutboundDataTransferDestinations ensures that no value is present for AllowedOutboundDataTransferDestinations, not even an explicit nil
### GetContactSyncBlocked

`func (o *ManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *ManagedAppProtection) GetContactSyncBlockedOk() (*bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSyncBlocked

`func (o *ManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked sets ContactSyncBlocked field to given value.

### HasContactSyncBlocked

`func (o *ManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### GetDataBackupBlocked

`func (o *ManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *ManagedAppProtection) GetDataBackupBlockedOk() (*bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackupBlocked

`func (o *ManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked sets DataBackupBlocked field to given value.

### HasDataBackupBlocked

`func (o *ManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### GetDeviceComplianceRequired

`func (o *ManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *ManagedAppProtection) GetDeviceComplianceRequiredOk() (*bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceRequired

`func (o *ManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired sets DeviceComplianceRequired field to given value.

### HasDeviceComplianceRequired

`func (o *ManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### GetDisableAppPinIfDevicePinIsSet

`func (o *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (*bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet sets DisableAppPinIfDevicePinIsSet field to given value.

### HasDisableAppPinIfDevicePinIsSet

`func (o *ManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### GetFingerprintBlocked

`func (o *ManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *ManagedAppProtection) GetFingerprintBlockedOk() (*bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBlocked

`func (o *ManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked sets FingerprintBlocked field to given value.

### HasFingerprintBlocked

`func (o *ManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### GetManagedBrowser

`func (o *ManagedAppProtection) GetManagedBrowser() AnyOfmicrosoftGraphManagedBrowserType`

GetManagedBrowser returns the ManagedBrowser field if non-nil, zero value otherwise.

### GetManagedBrowserOk

`func (o *ManagedAppProtection) GetManagedBrowserOk() (*AnyOfmicrosoftGraphManagedBrowserType, bool)`

GetManagedBrowserOk returns a tuple with the ManagedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowser

`func (o *ManagedAppProtection) SetManagedBrowser(v AnyOfmicrosoftGraphManagedBrowserType)`

SetManagedBrowser sets ManagedBrowser field to given value.

### HasManagedBrowser

`func (o *ManagedAppProtection) HasManagedBrowser() bool`

HasManagedBrowser returns a boolean if a field has been set.

### SetManagedBrowserNil

`func (o *ManagedAppProtection) SetManagedBrowserNil(b bool)`

 SetManagedBrowserNil sets the value for ManagedBrowser to be an explicit nil

### UnsetManagedBrowser
`func (o *ManagedAppProtection) UnsetManagedBrowser()`

UnsetManagedBrowser ensures that no value is present for ManagedBrowser, not even an explicit nil
### GetManagedBrowserToOpenLinksRequired

`func (o *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *ManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (*bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired sets ManagedBrowserToOpenLinksRequired field to given value.

### HasManagedBrowserToOpenLinksRequired

`func (o *ManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### GetMaximumPinRetries

`func (o *ManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *ManagedAppProtection) GetMaximumPinRetriesOk() (*int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinRetries

`func (o *ManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries sets MaximumPinRetries field to given value.

### HasMaximumPinRetries

`func (o *ManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### GetMinimumPinLength

`func (o *ManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *ManagedAppProtection) GetMinimumPinLengthOk() (*int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPinLength

`func (o *ManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength sets MinimumPinLength field to given value.

### HasMinimumPinLength

`func (o *ManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### GetMinimumRequiredAppVersion

`func (o *ManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *ManagedAppProtection) GetMinimumRequiredAppVersionOk() (*string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredAppVersion

`func (o *ManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion sets MinimumRequiredAppVersion field to given value.

### HasMinimumRequiredAppVersion

`func (o *ManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersionNil

`func (o *ManagedAppProtection) SetMinimumRequiredAppVersionNil(b bool)`

 SetMinimumRequiredAppVersionNil sets the value for MinimumRequiredAppVersion to be an explicit nil

### UnsetMinimumRequiredAppVersion
`func (o *ManagedAppProtection) UnsetMinimumRequiredAppVersion()`

UnsetMinimumRequiredAppVersion ensures that no value is present for MinimumRequiredAppVersion, not even an explicit nil
### GetMinimumRequiredOsVersion

`func (o *ManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *ManagedAppProtection) GetMinimumRequiredOsVersionOk() (*string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredOsVersion

`func (o *ManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion sets MinimumRequiredOsVersion field to given value.

### HasMinimumRequiredOsVersion

`func (o *ManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersionNil

`func (o *ManagedAppProtection) SetMinimumRequiredOsVersionNil(b bool)`

 SetMinimumRequiredOsVersionNil sets the value for MinimumRequiredOsVersion to be an explicit nil

### UnsetMinimumRequiredOsVersion
`func (o *ManagedAppProtection) UnsetMinimumRequiredOsVersion()`

UnsetMinimumRequiredOsVersion ensures that no value is present for MinimumRequiredOsVersion, not even an explicit nil
### GetMinimumWarningAppVersion

`func (o *ManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *ManagedAppProtection) GetMinimumWarningAppVersionOk() (*string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningAppVersion

`func (o *ManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion sets MinimumWarningAppVersion field to given value.

### HasMinimumWarningAppVersion

`func (o *ManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersionNil

`func (o *ManagedAppProtection) SetMinimumWarningAppVersionNil(b bool)`

 SetMinimumWarningAppVersionNil sets the value for MinimumWarningAppVersion to be an explicit nil

### UnsetMinimumWarningAppVersion
`func (o *ManagedAppProtection) UnsetMinimumWarningAppVersion()`

UnsetMinimumWarningAppVersion ensures that no value is present for MinimumWarningAppVersion, not even an explicit nil
### GetMinimumWarningOsVersion

`func (o *ManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *ManagedAppProtection) GetMinimumWarningOsVersionOk() (*string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningOsVersion

`func (o *ManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion sets MinimumWarningOsVersion field to given value.

### HasMinimumWarningOsVersion

`func (o *ManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersionNil

`func (o *ManagedAppProtection) SetMinimumWarningOsVersionNil(b bool)`

 SetMinimumWarningOsVersionNil sets the value for MinimumWarningOsVersion to be an explicit nil

### UnsetMinimumWarningOsVersion
`func (o *ManagedAppProtection) UnsetMinimumWarningOsVersion()`

UnsetMinimumWarningOsVersion ensures that no value is present for MinimumWarningOsVersion, not even an explicit nil
### GetOrganizationalCredentialsRequired

`func (o *ManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *ManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (*bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalCredentialsRequired

`func (o *ManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired sets OrganizationalCredentialsRequired field to given value.

### HasOrganizationalCredentialsRequired

`func (o *ManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### GetPeriodBeforePinReset

`func (o *ManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *ManagedAppProtection) GetPeriodBeforePinResetOk() (*string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBeforePinReset

`func (o *ManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset sets PeriodBeforePinReset field to given value.

### HasPeriodBeforePinReset

`func (o *ManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### GetPeriodOfflineBeforeAccessCheck

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck sets PeriodOfflineBeforeAccessCheck field to given value.

### HasPeriodOfflineBeforeAccessCheck

`func (o *ManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (*string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced sets PeriodOfflineBeforeWipeIsEnforced field to given value.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *ManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### GetPeriodOnlineBeforeAccessCheck

`func (o *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck sets PeriodOnlineBeforeAccessCheck field to given value.

### HasPeriodOnlineBeforeAccessCheck

`func (o *ManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### GetPinCharacterSet

`func (o *ManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *ManagedAppProtection) GetPinCharacterSetOk() (*AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCharacterSet

`func (o *ManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet sets PinCharacterSet field to given value.

### HasPinCharacterSet

`func (o *ManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSetNil

`func (o *ManagedAppProtection) SetPinCharacterSetNil(b bool)`

 SetPinCharacterSetNil sets the value for PinCharacterSet to be an explicit nil

### UnsetPinCharacterSet
`func (o *ManagedAppProtection) UnsetPinCharacterSet()`

UnsetPinCharacterSet ensures that no value is present for PinCharacterSet, not even an explicit nil
### GetPinRequired

`func (o *ManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *ManagedAppProtection) GetPinRequiredOk() (*bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinRequired

`func (o *ManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired sets PinRequired field to given value.

### HasPinRequired

`func (o *ManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### GetPrintBlocked

`func (o *ManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *ManagedAppProtection) GetPrintBlockedOk() (*bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBlocked

`func (o *ManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked sets PrintBlocked field to given value.

### HasPrintBlocked

`func (o *ManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### GetSaveAsBlocked

`func (o *ManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *ManagedAppProtection) GetSaveAsBlockedOk() (*bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAsBlocked

`func (o *ManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked sets SaveAsBlocked field to given value.

### HasSaveAsBlocked

`func (o *ManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### GetSimplePinBlocked

`func (o *ManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *ManagedAppProtection) GetSimplePinBlockedOk() (*bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePinBlocked

`func (o *ManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked sets SimplePinBlocked field to given value.

### HasSimplePinBlocked

`func (o *ManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


