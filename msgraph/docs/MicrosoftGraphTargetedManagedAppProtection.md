# MicrosoftGraphTargetedManagedAppProtection

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

## Methods

### NewMicrosoftGraphTargetedManagedAppProtection

`func NewMicrosoftGraphTargetedManagedAppProtection() *MicrosoftGraphTargetedManagedAppProtection`

NewMicrosoftGraphTargetedManagedAppProtection instantiates a new MicrosoftGraphTargetedManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTargetedManagedAppProtectionWithDefaults

`func NewMicrosoftGraphTargetedManagedAppProtectionWithDefaults() *MicrosoftGraphTargetedManagedAppProtection`

NewMicrosoftGraphTargetedManagedAppProtectionWithDefaults instantiates a new MicrosoftGraphTargetedManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedDataStorageLocationsOk() (*[]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations sets AllowedDataStorageLocations field to given value.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources sets AllowedInboundDataTransferSources field to given value.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSourcesNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedInboundDataTransferSourcesNil(b bool)`

 SetAllowedInboundDataTransferSourcesNil sets the value for AllowedInboundDataTransferSources to be an explicit nil

### UnsetAllowedInboundDataTransferSources
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetAllowedInboundDataTransferSources()`

UnsetAllowedInboundDataTransferSources ensures that no value is present for AllowedInboundDataTransferSources, not even an explicit nil
### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (*AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel sets AllowedOutboundClipboardSharingLevel field to given value.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevelNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedOutboundClipboardSharingLevelNil(b bool)`

 SetAllowedOutboundClipboardSharingLevelNil sets the value for AllowedOutboundClipboardSharingLevel to be an explicit nil

### UnsetAllowedOutboundClipboardSharingLevel
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetAllowedOutboundClipboardSharingLevel()`

UnsetAllowedOutboundClipboardSharingLevel ensures that no value is present for AllowedOutboundClipboardSharingLevel, not even an explicit nil
### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (*AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations sets AllowedOutboundDataTransferDestinations field to given value.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinationsNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAllowedOutboundDataTransferDestinationsNil(b bool)`

 SetAllowedOutboundDataTransferDestinationsNil sets the value for AllowedOutboundDataTransferDestinations to be an explicit nil

### UnsetAllowedOutboundDataTransferDestinations
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetAllowedOutboundDataTransferDestinations()`

UnsetAllowedOutboundDataTransferDestinations ensures that no value is present for AllowedOutboundDataTransferDestinations, not even an explicit nil
### GetContactSyncBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetContactSyncBlockedOk() (*bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked sets ContactSyncBlocked field to given value.

### HasContactSyncBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### GetDataBackupBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDataBackupBlockedOk() (*bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked sets DataBackupBlocked field to given value.

### HasDataBackupBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDeviceComplianceRequiredOk() (*bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired sets DeviceComplianceRequired field to given value.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (*bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet sets DisableAppPinIfDevicePinIsSet field to given value.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### GetFingerprintBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetFingerprintBlockedOk() (*bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked sets FingerprintBlocked field to given value.

### HasFingerprintBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### GetManagedBrowser

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetManagedBrowser() AnyOfmicrosoftGraphManagedBrowserType`

GetManagedBrowser returns the ManagedBrowser field if non-nil, zero value otherwise.

### GetManagedBrowserOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetManagedBrowserOk() (*AnyOfmicrosoftGraphManagedBrowserType, bool)`

GetManagedBrowserOk returns a tuple with the ManagedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowser

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetManagedBrowser(v AnyOfmicrosoftGraphManagedBrowserType)`

SetManagedBrowser sets ManagedBrowser field to given value.

### HasManagedBrowser

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasManagedBrowser() bool`

HasManagedBrowser returns a boolean if a field has been set.

### SetManagedBrowserNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetManagedBrowserNil(b bool)`

 SetManagedBrowserNil sets the value for ManagedBrowser to be an explicit nil

### UnsetManagedBrowser
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetManagedBrowser()`

UnsetManagedBrowser ensures that no value is present for ManagedBrowser, not even an explicit nil
### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (*bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired sets ManagedBrowserToOpenLinksRequired field to given value.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### GetMaximumPinRetries

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMaximumPinRetriesOk() (*int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries sets MaximumPinRetries field to given value.

### HasMaximumPinRetries

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### GetMinimumPinLength

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumPinLengthOk() (*int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength sets MinimumPinLength field to given value.

### HasMinimumPinLength

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumRequiredAppVersionOk() (*string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion sets MinimumRequiredAppVersion field to given value.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersionNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumRequiredAppVersionNil(b bool)`

 SetMinimumRequiredAppVersionNil sets the value for MinimumRequiredAppVersion to be an explicit nil

### UnsetMinimumRequiredAppVersion
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetMinimumRequiredAppVersion()`

UnsetMinimumRequiredAppVersion ensures that no value is present for MinimumRequiredAppVersion, not even an explicit nil
### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumRequiredOsVersionOk() (*string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion sets MinimumRequiredOsVersion field to given value.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersionNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumRequiredOsVersionNil(b bool)`

 SetMinimumRequiredOsVersionNil sets the value for MinimumRequiredOsVersion to be an explicit nil

### UnsetMinimumRequiredOsVersion
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetMinimumRequiredOsVersion()`

UnsetMinimumRequiredOsVersion ensures that no value is present for MinimumRequiredOsVersion, not even an explicit nil
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumWarningAppVersionOk() (*string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion sets MinimumWarningAppVersion field to given value.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersionNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumWarningAppVersionNil(b bool)`

 SetMinimumWarningAppVersionNil sets the value for MinimumWarningAppVersion to be an explicit nil

### UnsetMinimumWarningAppVersion
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetMinimumWarningAppVersion()`

UnsetMinimumWarningAppVersion ensures that no value is present for MinimumWarningAppVersion, not even an explicit nil
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetMinimumWarningOsVersionOk() (*string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion sets MinimumWarningOsVersion field to given value.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersionNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetMinimumWarningOsVersionNil(b bool)`

 SetMinimumWarningOsVersionNil sets the value for MinimumWarningOsVersion to be an explicit nil

### UnsetMinimumWarningOsVersion
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetMinimumWarningOsVersion()`

UnsetMinimumWarningOsVersion ensures that no value is present for MinimumWarningOsVersion, not even an explicit nil
### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (*bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired sets OrganizationalCredentialsRequired field to given value.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodBeforePinResetOk() (*string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset sets PeriodBeforePinReset field to given value.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck sets PeriodOfflineBeforeAccessCheck field to given value.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (*string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced sets PeriodOfflineBeforeWipeIsEnforced field to given value.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (*string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck sets PeriodOnlineBeforeAccessCheck field to given value.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### GetPinCharacterSet

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPinCharacterSetOk() (*AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet sets PinCharacterSet field to given value.

### HasPinCharacterSet

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSetNil

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPinCharacterSetNil(b bool)`

 SetPinCharacterSetNil sets the value for PinCharacterSet to be an explicit nil

### UnsetPinCharacterSet
`func (o *MicrosoftGraphTargetedManagedAppProtection) UnsetPinCharacterSet()`

UnsetPinCharacterSet ensures that no value is present for PinCharacterSet, not even an explicit nil
### GetPinRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPinRequiredOk() (*bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired sets PinRequired field to given value.

### HasPinRequired

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### GetPrintBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetPrintBlockedOk() (*bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked sets PrintBlocked field to given value.

### HasPrintBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### GetSaveAsBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetSaveAsBlockedOk() (*bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked sets SaveAsBlocked field to given value.

### HasSaveAsBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### GetSimplePinBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetSimplePinBlockedOk() (*bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked sets SimplePinBlocked field to given value.

### HasSimplePinBlocked

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### GetIsAssigned

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphTargetedManagedAppProtection) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphTargetedManagedAppProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphTargetedManagedAppProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


