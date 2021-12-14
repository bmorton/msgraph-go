# ManagedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationLockBypassCode** | Pointer to **NullableString** | Code that allows the Activation Lock on a device to be bypassed. This property is read-only. | [optional] 
**AndroidSecurityPatchLevel** | Pointer to **NullableString** | Android security patch level. This property is read-only. | [optional] 
**AzureADDeviceId** | Pointer to **NullableString** | The unique identifier for the Azure Active Directory device. Read only. This property is read-only. | [optional] 
**AzureADRegistered** | Pointer to **NullableBool** | Whether the device is Azure Active Directory registered. This property is read-only. | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to **time.Time** | The DateTime when device compliance grace period expires. This property is read-only. | [optional] 
**ComplianceState** | Pointer to [**AnyOfmicrosoftGraphComplianceState**](anyOf&lt;microsoft.graph.complianceState&gt;.md) | Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. | [optional] 
**ConfigurationManagerClientEnabledFeatures** | Pointer to [**AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures**](anyOf&lt;microsoft.graph.configurationManagerClientEnabledFeatures&gt;.md) | ConfigrMgr client enabled features. This property is read-only. | [optional] 
**DeviceActionResults** | Pointer to [**[]AnyOfmicrosoftGraphDeviceActionResult**](AnyOfmicrosoftGraphDeviceActionResult.md) | List of ComplexType deviceActionResult objects. This property is read-only. | [optional] 
**DeviceCategoryDisplayName** | Pointer to **NullableString** | Device category display name. This property is read-only. | [optional] 
**DeviceEnrollmentType** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentType**](anyOf&lt;microsoft.graph.deviceEnrollmentType&gt;.md) | Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount. | [optional] 
**DeviceHealthAttestationState** | Pointer to [**AnyOfmicrosoftGraphDeviceHealthAttestationState**](anyOf&lt;microsoft.graph.deviceHealthAttestationState&gt;.md) | The device health attestation state. This property is read-only. | [optional] 
**DeviceName** | Pointer to **NullableString** | Name of the device. This property is read-only. | [optional] 
**DeviceRegistrationState** | Pointer to [**AnyOfmicrosoftGraphDeviceRegistrationState**](anyOf&lt;microsoft.graph.deviceRegistrationState&gt;.md) | Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown. | [optional] 
**EasActivated** | Pointer to **bool** | Whether the device is Exchange ActiveSync activated. This property is read-only. | [optional] 
**EasActivationDateTime** | Pointer to **time.Time** | Exchange ActivationSync activation time of the device. This property is read-only. | [optional] 
**EasDeviceId** | Pointer to **NullableString** | Exchange ActiveSync Id of the device. This property is read-only. | [optional] 
**EmailAddress** | Pointer to **NullableString** | Email(s) for the user associated with the device. This property is read-only. | [optional] 
**EnrolledDateTime** | Pointer to **time.Time** | Enrollment time of the device. This property is read-only. | [optional] 
**EthernetMacAddress** | Pointer to **NullableString** | Ethernet MAC. This property is read-only. | [optional] 
**ExchangeAccessState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeAccessState**](anyOf&lt;microsoft.graph.deviceManagementExchangeAccessState&gt;.md) | The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined. | [optional] 
**ExchangeAccessStateReason** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason**](anyOf&lt;microsoft.graph.deviceManagementExchangeAccessStateReason&gt;.md) | The reason for the device&#39;s access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp. | [optional] 
**ExchangeLastSuccessfulSyncDateTime** | Pointer to **time.Time** | Last time the device contacted Exchange. This property is read-only. | [optional] 
**FreeStorageSpaceInBytes** | Pointer to **int64** | Free Storage in Bytes. This property is read-only. | [optional] 
**Iccid** | Pointer to **NullableString** | Integrated Circuit Card Identifier, it is A SIM card&#39;s unique identification number. This property is read-only. | [optional] 
**Imei** | Pointer to **NullableString** | IMEI. This property is read-only. | [optional] 
**IsEncrypted** | Pointer to **bool** | Device encryption status. This property is read-only. | [optional] 
**IsSupervised** | Pointer to **bool** | Device supervised status. This property is read-only. | [optional] 
**JailBroken** | Pointer to **NullableString** | whether the device is jail broken or rooted. This property is read-only. | [optional] 
**LastSyncDateTime** | Pointer to **time.Time** | The date and time that the device last completed a successful sync with Intune. This property is read-only. | [optional] 
**ManagedDeviceName** | Pointer to **NullableString** | Automatically generated name to identify a device. Can be overwritten to a user friendly name. | [optional] 
**ManagedDeviceOwnerType** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceOwnerType**](anyOf&lt;microsoft.graph.managedDeviceOwnerType&gt;.md) | Ownership of the device. Can be &#39;company&#39; or &#39;personal&#39;. Possible values are: unknown, company, personal. | [optional] 
**ManagementAgent** | Pointer to [**AnyOfmicrosoftGraphManagementAgentType**](anyOf&lt;microsoft.graph.managementAgentType&gt;.md) | Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController. | [optional] 
**Manufacturer** | Pointer to **NullableString** | Manufacturer of the device. This property is read-only. | [optional] 
**Meid** | Pointer to **NullableString** | MEID. This property is read-only. | [optional] 
**Model** | Pointer to **NullableString** | Model of the device. This property is read-only. | [optional] 
**Notes** | Pointer to **NullableString** | Notes on the device created by IT Admin | [optional] 
**OperatingSystem** | Pointer to **NullableString** | Operating system of the device. Windows, iOS, etc. This property is read-only. | [optional] 
**OsVersion** | Pointer to **NullableString** | Operating system version of the device. This property is read-only. | [optional] 
**PartnerReportedThreatState** | Pointer to [**AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState**](anyOf&lt;microsoft.graph.managedDevicePartnerReportedHealthState&gt;.md) | Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | Phone number of the device. This property is read-only. | [optional] 
**PhysicalMemoryInBytes** | Pointer to **int64** | Total Memory in Bytes. This property is read-only. | [optional] 
**RemoteAssistanceSessionErrorDetails** | Pointer to **NullableString** | An error string that identifies issues when creating Remote Assistance session objects. This property is read-only. | [optional] 
**RemoteAssistanceSessionUrl** | Pointer to **NullableString** | Url that allows a Remote Assistance session to be established with the device. This property is read-only. | [optional] 
**SerialNumber** | Pointer to **NullableString** | SerialNumber. This property is read-only. | [optional] 
**SubscriberCarrier** | Pointer to **NullableString** | Subscriber Carrier. This property is read-only. | [optional] 
**TotalStorageSpaceInBytes** | Pointer to **int64** | Total Storage in Bytes. This property is read-only. | [optional] 
**Udid** | Pointer to **NullableString** | Unique Device Identifier for iOS and macOS devices. This property is read-only. | [optional] 
**UserDisplayName** | Pointer to **NullableString** | User display name. This property is read-only. | [optional] 
**UserId** | Pointer to **NullableString** | Unique Identifier for the user associated with the device. This property is read-only. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | Device user principal name. This property is read-only. | [optional] 
**WiFiMacAddress** | Pointer to **NullableString** | Wi-Fi MAC. This property is read-only. | [optional] 
**DeviceCompliancePolicyStates** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md) | Device compliance policy states for this device. | [optional] 
**DeviceConfigurationStates** | Pointer to [**[]MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md) | Device configuration states for this device. | [optional] 
**DeviceCategory** | Pointer to [**AnyOfmicrosoftGraphDeviceCategory**](anyOf&lt;microsoft.graph.deviceCategory&gt;.md) | Device category | [optional] 

## Methods

### NewManagedDevice

`func NewManagedDevice() *ManagedDevice`

NewManagedDevice instantiates a new ManagedDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDeviceWithDefaults

`func NewManagedDeviceWithDefaults() *ManagedDevice`

NewManagedDeviceWithDefaults instantiates a new ManagedDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationLockBypassCode

`func (o *ManagedDevice) GetActivationLockBypassCode() string`

GetActivationLockBypassCode returns the ActivationLockBypassCode field if non-nil, zero value otherwise.

### GetActivationLockBypassCodeOk

`func (o *ManagedDevice) GetActivationLockBypassCodeOk() (*string, bool)`

GetActivationLockBypassCodeOk returns a tuple with the ActivationLockBypassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockBypassCode

`func (o *ManagedDevice) SetActivationLockBypassCode(v string)`

SetActivationLockBypassCode sets ActivationLockBypassCode field to given value.

### HasActivationLockBypassCode

`func (o *ManagedDevice) HasActivationLockBypassCode() bool`

HasActivationLockBypassCode returns a boolean if a field has been set.

### SetActivationLockBypassCodeNil

`func (o *ManagedDevice) SetActivationLockBypassCodeNil(b bool)`

 SetActivationLockBypassCodeNil sets the value for ActivationLockBypassCode to be an explicit nil

### UnsetActivationLockBypassCode
`func (o *ManagedDevice) UnsetActivationLockBypassCode()`

UnsetActivationLockBypassCode ensures that no value is present for ActivationLockBypassCode, not even an explicit nil
### GetAndroidSecurityPatchLevel

`func (o *ManagedDevice) GetAndroidSecurityPatchLevel() string`

GetAndroidSecurityPatchLevel returns the AndroidSecurityPatchLevel field if non-nil, zero value otherwise.

### GetAndroidSecurityPatchLevelOk

`func (o *ManagedDevice) GetAndroidSecurityPatchLevelOk() (*string, bool)`

GetAndroidSecurityPatchLevelOk returns a tuple with the AndroidSecurityPatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidSecurityPatchLevel

`func (o *ManagedDevice) SetAndroidSecurityPatchLevel(v string)`

SetAndroidSecurityPatchLevel sets AndroidSecurityPatchLevel field to given value.

### HasAndroidSecurityPatchLevel

`func (o *ManagedDevice) HasAndroidSecurityPatchLevel() bool`

HasAndroidSecurityPatchLevel returns a boolean if a field has been set.

### SetAndroidSecurityPatchLevelNil

`func (o *ManagedDevice) SetAndroidSecurityPatchLevelNil(b bool)`

 SetAndroidSecurityPatchLevelNil sets the value for AndroidSecurityPatchLevel to be an explicit nil

### UnsetAndroidSecurityPatchLevel
`func (o *ManagedDevice) UnsetAndroidSecurityPatchLevel()`

UnsetAndroidSecurityPatchLevel ensures that no value is present for AndroidSecurityPatchLevel, not even an explicit nil
### GetAzureADDeviceId

`func (o *ManagedDevice) GetAzureADDeviceId() string`

GetAzureADDeviceId returns the AzureADDeviceId field if non-nil, zero value otherwise.

### GetAzureADDeviceIdOk

`func (o *ManagedDevice) GetAzureADDeviceIdOk() (*string, bool)`

GetAzureADDeviceIdOk returns a tuple with the AzureADDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADDeviceId

`func (o *ManagedDevice) SetAzureADDeviceId(v string)`

SetAzureADDeviceId sets AzureADDeviceId field to given value.

### HasAzureADDeviceId

`func (o *ManagedDevice) HasAzureADDeviceId() bool`

HasAzureADDeviceId returns a boolean if a field has been set.

### SetAzureADDeviceIdNil

`func (o *ManagedDevice) SetAzureADDeviceIdNil(b bool)`

 SetAzureADDeviceIdNil sets the value for AzureADDeviceId to be an explicit nil

### UnsetAzureADDeviceId
`func (o *ManagedDevice) UnsetAzureADDeviceId()`

UnsetAzureADDeviceId ensures that no value is present for AzureADDeviceId, not even an explicit nil
### GetAzureADRegistered

`func (o *ManagedDevice) GetAzureADRegistered() bool`

GetAzureADRegistered returns the AzureADRegistered field if non-nil, zero value otherwise.

### GetAzureADRegisteredOk

`func (o *ManagedDevice) GetAzureADRegisteredOk() (*bool, bool)`

GetAzureADRegisteredOk returns a tuple with the AzureADRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADRegistered

`func (o *ManagedDevice) SetAzureADRegistered(v bool)`

SetAzureADRegistered sets AzureADRegistered field to given value.

### HasAzureADRegistered

`func (o *ManagedDevice) HasAzureADRegistered() bool`

HasAzureADRegistered returns a boolean if a field has been set.

### SetAzureADRegisteredNil

`func (o *ManagedDevice) SetAzureADRegisteredNil(b bool)`

 SetAzureADRegisteredNil sets the value for AzureADRegistered to be an explicit nil

### UnsetAzureADRegistered
`func (o *ManagedDevice) UnsetAzureADRegistered()`

UnsetAzureADRegistered ensures that no value is present for AzureADRegistered, not even an explicit nil
### GetComplianceGracePeriodExpirationDateTime

`func (o *ManagedDevice) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *ManagedDevice) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime sets ComplianceGracePeriodExpirationDateTime field to given value.

### HasComplianceGracePeriodExpirationDateTime

`func (o *ManagedDevice) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### GetComplianceState

`func (o *ManagedDevice) GetComplianceState() AnyOfmicrosoftGraphComplianceState`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *ManagedDevice) GetComplianceStateOk() (*AnyOfmicrosoftGraphComplianceState, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *ManagedDevice) SetComplianceState(v AnyOfmicrosoftGraphComplianceState)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *ManagedDevice) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### SetComplianceStateNil

`func (o *ManagedDevice) SetComplianceStateNil(b bool)`

 SetComplianceStateNil sets the value for ComplianceState to be an explicit nil

### UnsetComplianceState
`func (o *ManagedDevice) UnsetComplianceState()`

UnsetComplianceState ensures that no value is present for ComplianceState, not even an explicit nil
### GetConfigurationManagerClientEnabledFeatures

`func (o *ManagedDevice) GetConfigurationManagerClientEnabledFeatures() AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures`

GetConfigurationManagerClientEnabledFeatures returns the ConfigurationManagerClientEnabledFeatures field if non-nil, zero value otherwise.

### GetConfigurationManagerClientEnabledFeaturesOk

`func (o *ManagedDevice) GetConfigurationManagerClientEnabledFeaturesOk() (*AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures, bool)`

GetConfigurationManagerClientEnabledFeaturesOk returns a tuple with the ConfigurationManagerClientEnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationManagerClientEnabledFeatures

`func (o *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(v AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures)`

SetConfigurationManagerClientEnabledFeatures sets ConfigurationManagerClientEnabledFeatures field to given value.

### HasConfigurationManagerClientEnabledFeatures

`func (o *ManagedDevice) HasConfigurationManagerClientEnabledFeatures() bool`

HasConfigurationManagerClientEnabledFeatures returns a boolean if a field has been set.

### SetConfigurationManagerClientEnabledFeaturesNil

`func (o *ManagedDevice) SetConfigurationManagerClientEnabledFeaturesNil(b bool)`

 SetConfigurationManagerClientEnabledFeaturesNil sets the value for ConfigurationManagerClientEnabledFeatures to be an explicit nil

### UnsetConfigurationManagerClientEnabledFeatures
`func (o *ManagedDevice) UnsetConfigurationManagerClientEnabledFeatures()`

UnsetConfigurationManagerClientEnabledFeatures ensures that no value is present for ConfigurationManagerClientEnabledFeatures, not even an explicit nil
### GetDeviceActionResults

`func (o *ManagedDevice) GetDeviceActionResults() []*AnyOfmicrosoftGraphDeviceActionResult`

GetDeviceActionResults returns the DeviceActionResults field if non-nil, zero value otherwise.

### GetDeviceActionResultsOk

`func (o *ManagedDevice) GetDeviceActionResultsOk() (*[]*AnyOfmicrosoftGraphDeviceActionResult, bool)`

GetDeviceActionResultsOk returns a tuple with the DeviceActionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceActionResults

`func (o *ManagedDevice) SetDeviceActionResults(v []*AnyOfmicrosoftGraphDeviceActionResult)`

SetDeviceActionResults sets DeviceActionResults field to given value.

### HasDeviceActionResults

`func (o *ManagedDevice) HasDeviceActionResults() bool`

HasDeviceActionResults returns a boolean if a field has been set.

### GetDeviceCategoryDisplayName

`func (o *ManagedDevice) GetDeviceCategoryDisplayName() string`

GetDeviceCategoryDisplayName returns the DeviceCategoryDisplayName field if non-nil, zero value otherwise.

### GetDeviceCategoryDisplayNameOk

`func (o *ManagedDevice) GetDeviceCategoryDisplayNameOk() (*string, bool)`

GetDeviceCategoryDisplayNameOk returns a tuple with the DeviceCategoryDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategoryDisplayName

`func (o *ManagedDevice) SetDeviceCategoryDisplayName(v string)`

SetDeviceCategoryDisplayName sets DeviceCategoryDisplayName field to given value.

### HasDeviceCategoryDisplayName

`func (o *ManagedDevice) HasDeviceCategoryDisplayName() bool`

HasDeviceCategoryDisplayName returns a boolean if a field has been set.

### SetDeviceCategoryDisplayNameNil

`func (o *ManagedDevice) SetDeviceCategoryDisplayNameNil(b bool)`

 SetDeviceCategoryDisplayNameNil sets the value for DeviceCategoryDisplayName to be an explicit nil

### UnsetDeviceCategoryDisplayName
`func (o *ManagedDevice) UnsetDeviceCategoryDisplayName()`

UnsetDeviceCategoryDisplayName ensures that no value is present for DeviceCategoryDisplayName, not even an explicit nil
### GetDeviceEnrollmentType

`func (o *ManagedDevice) GetDeviceEnrollmentType() AnyOfmicrosoftGraphDeviceEnrollmentType`

GetDeviceEnrollmentType returns the DeviceEnrollmentType field if non-nil, zero value otherwise.

### GetDeviceEnrollmentTypeOk

`func (o *ManagedDevice) GetDeviceEnrollmentTypeOk() (*AnyOfmicrosoftGraphDeviceEnrollmentType, bool)`

GetDeviceEnrollmentTypeOk returns a tuple with the DeviceEnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentType

`func (o *ManagedDevice) SetDeviceEnrollmentType(v AnyOfmicrosoftGraphDeviceEnrollmentType)`

SetDeviceEnrollmentType sets DeviceEnrollmentType field to given value.

### HasDeviceEnrollmentType

`func (o *ManagedDevice) HasDeviceEnrollmentType() bool`

HasDeviceEnrollmentType returns a boolean if a field has been set.

### SetDeviceEnrollmentTypeNil

`func (o *ManagedDevice) SetDeviceEnrollmentTypeNil(b bool)`

 SetDeviceEnrollmentTypeNil sets the value for DeviceEnrollmentType to be an explicit nil

### UnsetDeviceEnrollmentType
`func (o *ManagedDevice) UnsetDeviceEnrollmentType()`

UnsetDeviceEnrollmentType ensures that no value is present for DeviceEnrollmentType, not even an explicit nil
### GetDeviceHealthAttestationState

`func (o *ManagedDevice) GetDeviceHealthAttestationState() AnyOfmicrosoftGraphDeviceHealthAttestationState`

GetDeviceHealthAttestationState returns the DeviceHealthAttestationState field if non-nil, zero value otherwise.

### GetDeviceHealthAttestationStateOk

`func (o *ManagedDevice) GetDeviceHealthAttestationStateOk() (*AnyOfmicrosoftGraphDeviceHealthAttestationState, bool)`

GetDeviceHealthAttestationStateOk returns a tuple with the DeviceHealthAttestationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealthAttestationState

`func (o *ManagedDevice) SetDeviceHealthAttestationState(v AnyOfmicrosoftGraphDeviceHealthAttestationState)`

SetDeviceHealthAttestationState sets DeviceHealthAttestationState field to given value.

### HasDeviceHealthAttestationState

`func (o *ManagedDevice) HasDeviceHealthAttestationState() bool`

HasDeviceHealthAttestationState returns a boolean if a field has been set.

### SetDeviceHealthAttestationStateNil

`func (o *ManagedDevice) SetDeviceHealthAttestationStateNil(b bool)`

 SetDeviceHealthAttestationStateNil sets the value for DeviceHealthAttestationState to be an explicit nil

### UnsetDeviceHealthAttestationState
`func (o *ManagedDevice) UnsetDeviceHealthAttestationState()`

UnsetDeviceHealthAttestationState ensures that no value is present for DeviceHealthAttestationState, not even an explicit nil
### GetDeviceName

`func (o *ManagedDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ManagedDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ManagedDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ManagedDevice) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *ManagedDevice) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *ManagedDevice) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceRegistrationState

`func (o *ManagedDevice) GetDeviceRegistrationState() AnyOfmicrosoftGraphDeviceRegistrationState`

GetDeviceRegistrationState returns the DeviceRegistrationState field if non-nil, zero value otherwise.

### GetDeviceRegistrationStateOk

`func (o *ManagedDevice) GetDeviceRegistrationStateOk() (*AnyOfmicrosoftGraphDeviceRegistrationState, bool)`

GetDeviceRegistrationStateOk returns a tuple with the DeviceRegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationState

`func (o *ManagedDevice) SetDeviceRegistrationState(v AnyOfmicrosoftGraphDeviceRegistrationState)`

SetDeviceRegistrationState sets DeviceRegistrationState field to given value.

### HasDeviceRegistrationState

`func (o *ManagedDevice) HasDeviceRegistrationState() bool`

HasDeviceRegistrationState returns a boolean if a field has been set.

### SetDeviceRegistrationStateNil

`func (o *ManagedDevice) SetDeviceRegistrationStateNil(b bool)`

 SetDeviceRegistrationStateNil sets the value for DeviceRegistrationState to be an explicit nil

### UnsetDeviceRegistrationState
`func (o *ManagedDevice) UnsetDeviceRegistrationState()`

UnsetDeviceRegistrationState ensures that no value is present for DeviceRegistrationState, not even an explicit nil
### GetEasActivated

`func (o *ManagedDevice) GetEasActivated() bool`

GetEasActivated returns the EasActivated field if non-nil, zero value otherwise.

### GetEasActivatedOk

`func (o *ManagedDevice) GetEasActivatedOk() (*bool, bool)`

GetEasActivatedOk returns a tuple with the EasActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasActivated

`func (o *ManagedDevice) SetEasActivated(v bool)`

SetEasActivated sets EasActivated field to given value.

### HasEasActivated

`func (o *ManagedDevice) HasEasActivated() bool`

HasEasActivated returns a boolean if a field has been set.

### GetEasActivationDateTime

`func (o *ManagedDevice) GetEasActivationDateTime() time.Time`

GetEasActivationDateTime returns the EasActivationDateTime field if non-nil, zero value otherwise.

### GetEasActivationDateTimeOk

`func (o *ManagedDevice) GetEasActivationDateTimeOk() (*time.Time, bool)`

GetEasActivationDateTimeOk returns a tuple with the EasActivationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasActivationDateTime

`func (o *ManagedDevice) SetEasActivationDateTime(v time.Time)`

SetEasActivationDateTime sets EasActivationDateTime field to given value.

### HasEasActivationDateTime

`func (o *ManagedDevice) HasEasActivationDateTime() bool`

HasEasActivationDateTime returns a boolean if a field has been set.

### GetEasDeviceId

`func (o *ManagedDevice) GetEasDeviceId() string`

GetEasDeviceId returns the EasDeviceId field if non-nil, zero value otherwise.

### GetEasDeviceIdOk

`func (o *ManagedDevice) GetEasDeviceIdOk() (*string, bool)`

GetEasDeviceIdOk returns a tuple with the EasDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDeviceId

`func (o *ManagedDevice) SetEasDeviceId(v string)`

SetEasDeviceId sets EasDeviceId field to given value.

### HasEasDeviceId

`func (o *ManagedDevice) HasEasDeviceId() bool`

HasEasDeviceId returns a boolean if a field has been set.

### SetEasDeviceIdNil

`func (o *ManagedDevice) SetEasDeviceIdNil(b bool)`

 SetEasDeviceIdNil sets the value for EasDeviceId to be an explicit nil

### UnsetEasDeviceId
`func (o *ManagedDevice) UnsetEasDeviceId()`

UnsetEasDeviceId ensures that no value is present for EasDeviceId, not even an explicit nil
### GetEmailAddress

`func (o *ManagedDevice) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ManagedDevice) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ManagedDevice) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ManagedDevice) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *ManagedDevice) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *ManagedDevice) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEnrolledDateTime

`func (o *ManagedDevice) GetEnrolledDateTime() time.Time`

GetEnrolledDateTime returns the EnrolledDateTime field if non-nil, zero value otherwise.

### GetEnrolledDateTimeOk

`func (o *ManagedDevice) GetEnrolledDateTimeOk() (*time.Time, bool)`

GetEnrolledDateTimeOk returns a tuple with the EnrolledDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledDateTime

`func (o *ManagedDevice) SetEnrolledDateTime(v time.Time)`

SetEnrolledDateTime sets EnrolledDateTime field to given value.

### HasEnrolledDateTime

`func (o *ManagedDevice) HasEnrolledDateTime() bool`

HasEnrolledDateTime returns a boolean if a field has been set.

### GetEthernetMacAddress

`func (o *ManagedDevice) GetEthernetMacAddress() string`

GetEthernetMacAddress returns the EthernetMacAddress field if non-nil, zero value otherwise.

### GetEthernetMacAddressOk

`func (o *ManagedDevice) GetEthernetMacAddressOk() (*string, bool)`

GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMacAddress

`func (o *ManagedDevice) SetEthernetMacAddress(v string)`

SetEthernetMacAddress sets EthernetMacAddress field to given value.

### HasEthernetMacAddress

`func (o *ManagedDevice) HasEthernetMacAddress() bool`

HasEthernetMacAddress returns a boolean if a field has been set.

### SetEthernetMacAddressNil

`func (o *ManagedDevice) SetEthernetMacAddressNil(b bool)`

 SetEthernetMacAddressNil sets the value for EthernetMacAddress to be an explicit nil

### UnsetEthernetMacAddress
`func (o *ManagedDevice) UnsetEthernetMacAddress()`

UnsetEthernetMacAddress ensures that no value is present for EthernetMacAddress, not even an explicit nil
### GetExchangeAccessState

`func (o *ManagedDevice) GetExchangeAccessState() AnyOfmicrosoftGraphDeviceManagementExchangeAccessState`

GetExchangeAccessState returns the ExchangeAccessState field if non-nil, zero value otherwise.

### GetExchangeAccessStateOk

`func (o *ManagedDevice) GetExchangeAccessStateOk() (*AnyOfmicrosoftGraphDeviceManagementExchangeAccessState, bool)`

GetExchangeAccessStateOk returns a tuple with the ExchangeAccessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAccessState

`func (o *ManagedDevice) SetExchangeAccessState(v AnyOfmicrosoftGraphDeviceManagementExchangeAccessState)`

SetExchangeAccessState sets ExchangeAccessState field to given value.

### HasExchangeAccessState

`func (o *ManagedDevice) HasExchangeAccessState() bool`

HasExchangeAccessState returns a boolean if a field has been set.

### SetExchangeAccessStateNil

`func (o *ManagedDevice) SetExchangeAccessStateNil(b bool)`

 SetExchangeAccessStateNil sets the value for ExchangeAccessState to be an explicit nil

### UnsetExchangeAccessState
`func (o *ManagedDevice) UnsetExchangeAccessState()`

UnsetExchangeAccessState ensures that no value is present for ExchangeAccessState, not even an explicit nil
### GetExchangeAccessStateReason

`func (o *ManagedDevice) GetExchangeAccessStateReason() AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason`

GetExchangeAccessStateReason returns the ExchangeAccessStateReason field if non-nil, zero value otherwise.

### GetExchangeAccessStateReasonOk

`func (o *ManagedDevice) GetExchangeAccessStateReasonOk() (*AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason, bool)`

GetExchangeAccessStateReasonOk returns a tuple with the ExchangeAccessStateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAccessStateReason

`func (o *ManagedDevice) SetExchangeAccessStateReason(v AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason)`

SetExchangeAccessStateReason sets ExchangeAccessStateReason field to given value.

### HasExchangeAccessStateReason

`func (o *ManagedDevice) HasExchangeAccessStateReason() bool`

HasExchangeAccessStateReason returns a boolean if a field has been set.

### SetExchangeAccessStateReasonNil

`func (o *ManagedDevice) SetExchangeAccessStateReasonNil(b bool)`

 SetExchangeAccessStateReasonNil sets the value for ExchangeAccessStateReason to be an explicit nil

### UnsetExchangeAccessStateReason
`func (o *ManagedDevice) UnsetExchangeAccessStateReason()`

UnsetExchangeAccessStateReason ensures that no value is present for ExchangeAccessStateReason, not even an explicit nil
### GetExchangeLastSuccessfulSyncDateTime

`func (o *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime() time.Time`

GetExchangeLastSuccessfulSyncDateTime returns the ExchangeLastSuccessfulSyncDateTime field if non-nil, zero value otherwise.

### GetExchangeLastSuccessfulSyncDateTimeOk

`func (o *ManagedDevice) GetExchangeLastSuccessfulSyncDateTimeOk() (*time.Time, bool)`

GetExchangeLastSuccessfulSyncDateTimeOk returns a tuple with the ExchangeLastSuccessfulSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeLastSuccessfulSyncDateTime

`func (o *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(v time.Time)`

SetExchangeLastSuccessfulSyncDateTime sets ExchangeLastSuccessfulSyncDateTime field to given value.

### HasExchangeLastSuccessfulSyncDateTime

`func (o *ManagedDevice) HasExchangeLastSuccessfulSyncDateTime() bool`

HasExchangeLastSuccessfulSyncDateTime returns a boolean if a field has been set.

### GetFreeStorageSpaceInBytes

`func (o *ManagedDevice) GetFreeStorageSpaceInBytes() int64`

GetFreeStorageSpaceInBytes returns the FreeStorageSpaceInBytes field if non-nil, zero value otherwise.

### GetFreeStorageSpaceInBytesOk

`func (o *ManagedDevice) GetFreeStorageSpaceInBytesOk() (*int64, bool)`

GetFreeStorageSpaceInBytesOk returns a tuple with the FreeStorageSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStorageSpaceInBytes

`func (o *ManagedDevice) SetFreeStorageSpaceInBytes(v int64)`

SetFreeStorageSpaceInBytes sets FreeStorageSpaceInBytes field to given value.

### HasFreeStorageSpaceInBytes

`func (o *ManagedDevice) HasFreeStorageSpaceInBytes() bool`

HasFreeStorageSpaceInBytes returns a boolean if a field has been set.

### GetIccid

`func (o *ManagedDevice) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *ManagedDevice) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *ManagedDevice) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *ManagedDevice) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### SetIccidNil

`func (o *ManagedDevice) SetIccidNil(b bool)`

 SetIccidNil sets the value for Iccid to be an explicit nil

### UnsetIccid
`func (o *ManagedDevice) UnsetIccid()`

UnsetIccid ensures that no value is present for Iccid, not even an explicit nil
### GetImei

`func (o *ManagedDevice) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *ManagedDevice) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *ManagedDevice) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *ManagedDevice) HasImei() bool`

HasImei returns a boolean if a field has been set.

### SetImeiNil

`func (o *ManagedDevice) SetImeiNil(b bool)`

 SetImeiNil sets the value for Imei to be an explicit nil

### UnsetImei
`func (o *ManagedDevice) UnsetImei()`

UnsetImei ensures that no value is present for Imei, not even an explicit nil
### GetIsEncrypted

`func (o *ManagedDevice) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *ManagedDevice) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *ManagedDevice) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *ManagedDevice) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### GetIsSupervised

`func (o *ManagedDevice) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *ManagedDevice) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *ManagedDevice) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.

### HasIsSupervised

`func (o *ManagedDevice) HasIsSupervised() bool`

HasIsSupervised returns a boolean if a field has been set.

### GetJailBroken

`func (o *ManagedDevice) GetJailBroken() string`

GetJailBroken returns the JailBroken field if non-nil, zero value otherwise.

### GetJailBrokenOk

`func (o *ManagedDevice) GetJailBrokenOk() (*string, bool)`

GetJailBrokenOk returns a tuple with the JailBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailBroken

`func (o *ManagedDevice) SetJailBroken(v string)`

SetJailBroken sets JailBroken field to given value.

### HasJailBroken

`func (o *ManagedDevice) HasJailBroken() bool`

HasJailBroken returns a boolean if a field has been set.

### SetJailBrokenNil

`func (o *ManagedDevice) SetJailBrokenNil(b bool)`

 SetJailBrokenNil sets the value for JailBroken to be an explicit nil

### UnsetJailBroken
`func (o *ManagedDevice) UnsetJailBroken()`

UnsetJailBroken ensures that no value is present for JailBroken, not even an explicit nil
### GetLastSyncDateTime

`func (o *ManagedDevice) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *ManagedDevice) GetLastSyncDateTimeOk() (*time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDateTime

`func (o *ManagedDevice) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime sets LastSyncDateTime field to given value.

### HasLastSyncDateTime

`func (o *ManagedDevice) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### GetManagedDeviceName

`func (o *ManagedDevice) GetManagedDeviceName() string`

GetManagedDeviceName returns the ManagedDeviceName field if non-nil, zero value otherwise.

### GetManagedDeviceNameOk

`func (o *ManagedDevice) GetManagedDeviceNameOk() (*string, bool)`

GetManagedDeviceNameOk returns a tuple with the ManagedDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDeviceName

`func (o *ManagedDevice) SetManagedDeviceName(v string)`

SetManagedDeviceName sets ManagedDeviceName field to given value.

### HasManagedDeviceName

`func (o *ManagedDevice) HasManagedDeviceName() bool`

HasManagedDeviceName returns a boolean if a field has been set.

### SetManagedDeviceNameNil

`func (o *ManagedDevice) SetManagedDeviceNameNil(b bool)`

 SetManagedDeviceNameNil sets the value for ManagedDeviceName to be an explicit nil

### UnsetManagedDeviceName
`func (o *ManagedDevice) UnsetManagedDeviceName()`

UnsetManagedDeviceName ensures that no value is present for ManagedDeviceName, not even an explicit nil
### GetManagedDeviceOwnerType

`func (o *ManagedDevice) GetManagedDeviceOwnerType() AnyOfmicrosoftGraphManagedDeviceOwnerType`

GetManagedDeviceOwnerType returns the ManagedDeviceOwnerType field if non-nil, zero value otherwise.

### GetManagedDeviceOwnerTypeOk

`func (o *ManagedDevice) GetManagedDeviceOwnerTypeOk() (*AnyOfmicrosoftGraphManagedDeviceOwnerType, bool)`

GetManagedDeviceOwnerTypeOk returns a tuple with the ManagedDeviceOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDeviceOwnerType

`func (o *ManagedDevice) SetManagedDeviceOwnerType(v AnyOfmicrosoftGraphManagedDeviceOwnerType)`

SetManagedDeviceOwnerType sets ManagedDeviceOwnerType field to given value.

### HasManagedDeviceOwnerType

`func (o *ManagedDevice) HasManagedDeviceOwnerType() bool`

HasManagedDeviceOwnerType returns a boolean if a field has been set.

### SetManagedDeviceOwnerTypeNil

`func (o *ManagedDevice) SetManagedDeviceOwnerTypeNil(b bool)`

 SetManagedDeviceOwnerTypeNil sets the value for ManagedDeviceOwnerType to be an explicit nil

### UnsetManagedDeviceOwnerType
`func (o *ManagedDevice) UnsetManagedDeviceOwnerType()`

UnsetManagedDeviceOwnerType ensures that no value is present for ManagedDeviceOwnerType, not even an explicit nil
### GetManagementAgent

`func (o *ManagedDevice) GetManagementAgent() AnyOfmicrosoftGraphManagementAgentType`

GetManagementAgent returns the ManagementAgent field if non-nil, zero value otherwise.

### GetManagementAgentOk

`func (o *ManagedDevice) GetManagementAgentOk() (*AnyOfmicrosoftGraphManagementAgentType, bool)`

GetManagementAgentOk returns a tuple with the ManagementAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAgent

`func (o *ManagedDevice) SetManagementAgent(v AnyOfmicrosoftGraphManagementAgentType)`

SetManagementAgent sets ManagementAgent field to given value.

### HasManagementAgent

`func (o *ManagedDevice) HasManagementAgent() bool`

HasManagementAgent returns a boolean if a field has been set.

### SetManagementAgentNil

`func (o *ManagedDevice) SetManagementAgentNil(b bool)`

 SetManagementAgentNil sets the value for ManagementAgent to be an explicit nil

### UnsetManagementAgent
`func (o *ManagedDevice) UnsetManagementAgent()`

UnsetManagementAgent ensures that no value is present for ManagementAgent, not even an explicit nil
### GetManufacturer

`func (o *ManagedDevice) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ManagedDevice) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ManagedDevice) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ManagedDevice) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *ManagedDevice) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *ManagedDevice) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetMeid

`func (o *ManagedDevice) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *ManagedDevice) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *ManagedDevice) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *ManagedDevice) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### SetMeidNil

`func (o *ManagedDevice) SetMeidNil(b bool)`

 SetMeidNil sets the value for Meid to be an explicit nil

### UnsetMeid
`func (o *ManagedDevice) UnsetMeid()`

UnsetMeid ensures that no value is present for Meid, not even an explicit nil
### GetModel

`func (o *ManagedDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ManagedDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ManagedDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ManagedDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ManagedDevice) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ManagedDevice) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetNotes

`func (o *ManagedDevice) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ManagedDevice) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ManagedDevice) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ManagedDevice) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ManagedDevice) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ManagedDevice) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOperatingSystem

`func (o *ManagedDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ManagedDevice) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ManagedDevice) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ManagedDevice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *ManagedDevice) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *ManagedDevice) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetOsVersion

`func (o *ManagedDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ManagedDevice) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ManagedDevice) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ManagedDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *ManagedDevice) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *ManagedDevice) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetPartnerReportedThreatState

`func (o *ManagedDevice) GetPartnerReportedThreatState() AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState`

GetPartnerReportedThreatState returns the PartnerReportedThreatState field if non-nil, zero value otherwise.

### GetPartnerReportedThreatStateOk

`func (o *ManagedDevice) GetPartnerReportedThreatStateOk() (*AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState, bool)`

GetPartnerReportedThreatStateOk returns a tuple with the PartnerReportedThreatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReportedThreatState

`func (o *ManagedDevice) SetPartnerReportedThreatState(v AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState)`

SetPartnerReportedThreatState sets PartnerReportedThreatState field to given value.

### HasPartnerReportedThreatState

`func (o *ManagedDevice) HasPartnerReportedThreatState() bool`

HasPartnerReportedThreatState returns a boolean if a field has been set.

### SetPartnerReportedThreatStateNil

`func (o *ManagedDevice) SetPartnerReportedThreatStateNil(b bool)`

 SetPartnerReportedThreatStateNil sets the value for PartnerReportedThreatState to be an explicit nil

### UnsetPartnerReportedThreatState
`func (o *ManagedDevice) UnsetPartnerReportedThreatState()`

UnsetPartnerReportedThreatState ensures that no value is present for PartnerReportedThreatState, not even an explicit nil
### GetPhoneNumber

`func (o *ManagedDevice) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ManagedDevice) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ManagedDevice) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ManagedDevice) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ManagedDevice) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ManagedDevice) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPhysicalMemoryInBytes

`func (o *ManagedDevice) GetPhysicalMemoryInBytes() int64`

GetPhysicalMemoryInBytes returns the PhysicalMemoryInBytes field if non-nil, zero value otherwise.

### GetPhysicalMemoryInBytesOk

`func (o *ManagedDevice) GetPhysicalMemoryInBytesOk() (*int64, bool)`

GetPhysicalMemoryInBytesOk returns a tuple with the PhysicalMemoryInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalMemoryInBytes

`func (o *ManagedDevice) SetPhysicalMemoryInBytes(v int64)`

SetPhysicalMemoryInBytes sets PhysicalMemoryInBytes field to given value.

### HasPhysicalMemoryInBytes

`func (o *ManagedDevice) HasPhysicalMemoryInBytes() bool`

HasPhysicalMemoryInBytes returns a boolean if a field has been set.

### GetRemoteAssistanceSessionErrorDetails

`func (o *ManagedDevice) GetRemoteAssistanceSessionErrorDetails() string`

GetRemoteAssistanceSessionErrorDetails returns the RemoteAssistanceSessionErrorDetails field if non-nil, zero value otherwise.

### GetRemoteAssistanceSessionErrorDetailsOk

`func (o *ManagedDevice) GetRemoteAssistanceSessionErrorDetailsOk() (*string, bool)`

GetRemoteAssistanceSessionErrorDetailsOk returns a tuple with the RemoteAssistanceSessionErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssistanceSessionErrorDetails

`func (o *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(v string)`

SetRemoteAssistanceSessionErrorDetails sets RemoteAssistanceSessionErrorDetails field to given value.

### HasRemoteAssistanceSessionErrorDetails

`func (o *ManagedDevice) HasRemoteAssistanceSessionErrorDetails() bool`

HasRemoteAssistanceSessionErrorDetails returns a boolean if a field has been set.

### SetRemoteAssistanceSessionErrorDetailsNil

`func (o *ManagedDevice) SetRemoteAssistanceSessionErrorDetailsNil(b bool)`

 SetRemoteAssistanceSessionErrorDetailsNil sets the value for RemoteAssistanceSessionErrorDetails to be an explicit nil

### UnsetRemoteAssistanceSessionErrorDetails
`func (o *ManagedDevice) UnsetRemoteAssistanceSessionErrorDetails()`

UnsetRemoteAssistanceSessionErrorDetails ensures that no value is present for RemoteAssistanceSessionErrorDetails, not even an explicit nil
### GetRemoteAssistanceSessionUrl

`func (o *ManagedDevice) GetRemoteAssistanceSessionUrl() string`

GetRemoteAssistanceSessionUrl returns the RemoteAssistanceSessionUrl field if non-nil, zero value otherwise.

### GetRemoteAssistanceSessionUrlOk

`func (o *ManagedDevice) GetRemoteAssistanceSessionUrlOk() (*string, bool)`

GetRemoteAssistanceSessionUrlOk returns a tuple with the RemoteAssistanceSessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssistanceSessionUrl

`func (o *ManagedDevice) SetRemoteAssistanceSessionUrl(v string)`

SetRemoteAssistanceSessionUrl sets RemoteAssistanceSessionUrl field to given value.

### HasRemoteAssistanceSessionUrl

`func (o *ManagedDevice) HasRemoteAssistanceSessionUrl() bool`

HasRemoteAssistanceSessionUrl returns a boolean if a field has been set.

### SetRemoteAssistanceSessionUrlNil

`func (o *ManagedDevice) SetRemoteAssistanceSessionUrlNil(b bool)`

 SetRemoteAssistanceSessionUrlNil sets the value for RemoteAssistanceSessionUrl to be an explicit nil

### UnsetRemoteAssistanceSessionUrl
`func (o *ManagedDevice) UnsetRemoteAssistanceSessionUrl()`

UnsetRemoteAssistanceSessionUrl ensures that no value is present for RemoteAssistanceSessionUrl, not even an explicit nil
### GetSerialNumber

`func (o *ManagedDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ManagedDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ManagedDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ManagedDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *ManagedDevice) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *ManagedDevice) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetSubscriberCarrier

`func (o *ManagedDevice) GetSubscriberCarrier() string`

GetSubscriberCarrier returns the SubscriberCarrier field if non-nil, zero value otherwise.

### GetSubscriberCarrierOk

`func (o *ManagedDevice) GetSubscriberCarrierOk() (*string, bool)`

GetSubscriberCarrierOk returns a tuple with the SubscriberCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCarrier

`func (o *ManagedDevice) SetSubscriberCarrier(v string)`

SetSubscriberCarrier sets SubscriberCarrier field to given value.

### HasSubscriberCarrier

`func (o *ManagedDevice) HasSubscriberCarrier() bool`

HasSubscriberCarrier returns a boolean if a field has been set.

### SetSubscriberCarrierNil

`func (o *ManagedDevice) SetSubscriberCarrierNil(b bool)`

 SetSubscriberCarrierNil sets the value for SubscriberCarrier to be an explicit nil

### UnsetSubscriberCarrier
`func (o *ManagedDevice) UnsetSubscriberCarrier()`

UnsetSubscriberCarrier ensures that no value is present for SubscriberCarrier, not even an explicit nil
### GetTotalStorageSpaceInBytes

`func (o *ManagedDevice) GetTotalStorageSpaceInBytes() int64`

GetTotalStorageSpaceInBytes returns the TotalStorageSpaceInBytes field if non-nil, zero value otherwise.

### GetTotalStorageSpaceInBytesOk

`func (o *ManagedDevice) GetTotalStorageSpaceInBytesOk() (*int64, bool)`

GetTotalStorageSpaceInBytesOk returns a tuple with the TotalStorageSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorageSpaceInBytes

`func (o *ManagedDevice) SetTotalStorageSpaceInBytes(v int64)`

SetTotalStorageSpaceInBytes sets TotalStorageSpaceInBytes field to given value.

### HasTotalStorageSpaceInBytes

`func (o *ManagedDevice) HasTotalStorageSpaceInBytes() bool`

HasTotalStorageSpaceInBytes returns a boolean if a field has been set.

### GetUdid

`func (o *ManagedDevice) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ManagedDevice) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ManagedDevice) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ManagedDevice) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### SetUdidNil

`func (o *ManagedDevice) SetUdidNil(b bool)`

 SetUdidNil sets the value for Udid to be an explicit nil

### UnsetUdid
`func (o *ManagedDevice) UnsetUdid()`

UnsetUdid ensures that no value is present for Udid, not even an explicit nil
### GetUserDisplayName

`func (o *ManagedDevice) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *ManagedDevice) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *ManagedDevice) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *ManagedDevice) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *ManagedDevice) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *ManagedDevice) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserId

`func (o *ManagedDevice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ManagedDevice) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ManagedDevice) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ManagedDevice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ManagedDevice) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ManagedDevice) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserPrincipalName

`func (o *ManagedDevice) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ManagedDevice) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *ManagedDevice) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *ManagedDevice) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *ManagedDevice) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *ManagedDevice) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetWiFiMacAddress

`func (o *ManagedDevice) GetWiFiMacAddress() string`

GetWiFiMacAddress returns the WiFiMacAddress field if non-nil, zero value otherwise.

### GetWiFiMacAddressOk

`func (o *ManagedDevice) GetWiFiMacAddressOk() (*string, bool)`

GetWiFiMacAddressOk returns a tuple with the WiFiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiFiMacAddress

`func (o *ManagedDevice) SetWiFiMacAddress(v string)`

SetWiFiMacAddress sets WiFiMacAddress field to given value.

### HasWiFiMacAddress

`func (o *ManagedDevice) HasWiFiMacAddress() bool`

HasWiFiMacAddress returns a boolean if a field has been set.

### SetWiFiMacAddressNil

`func (o *ManagedDevice) SetWiFiMacAddressNil(b bool)`

 SetWiFiMacAddressNil sets the value for WiFiMacAddress to be an explicit nil

### UnsetWiFiMacAddress
`func (o *ManagedDevice) UnsetWiFiMacAddress()`

UnsetWiFiMacAddress ensures that no value is present for WiFiMacAddress, not even an explicit nil
### GetDeviceCompliancePolicyStates

`func (o *ManagedDevice) GetDeviceCompliancePolicyStates() []MicrosoftGraphDeviceCompliancePolicyState`

GetDeviceCompliancePolicyStates returns the DeviceCompliancePolicyStates field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicyStatesOk

`func (o *ManagedDevice) GetDeviceCompliancePolicyStatesOk() (*[]MicrosoftGraphDeviceCompliancePolicyState, bool)`

GetDeviceCompliancePolicyStatesOk returns a tuple with the DeviceCompliancePolicyStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCompliancePolicyStates

`func (o *ManagedDevice) SetDeviceCompliancePolicyStates(v []MicrosoftGraphDeviceCompliancePolicyState)`

SetDeviceCompliancePolicyStates sets DeviceCompliancePolicyStates field to given value.

### HasDeviceCompliancePolicyStates

`func (o *ManagedDevice) HasDeviceCompliancePolicyStates() bool`

HasDeviceCompliancePolicyStates returns a boolean if a field has been set.

### GetDeviceConfigurationStates

`func (o *ManagedDevice) GetDeviceConfigurationStates() []MicrosoftGraphDeviceConfigurationState`

GetDeviceConfigurationStates returns the DeviceConfigurationStates field if non-nil, zero value otherwise.

### GetDeviceConfigurationStatesOk

`func (o *ManagedDevice) GetDeviceConfigurationStatesOk() (*[]MicrosoftGraphDeviceConfigurationState, bool)`

GetDeviceConfigurationStatesOk returns a tuple with the DeviceConfigurationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurationStates

`func (o *ManagedDevice) SetDeviceConfigurationStates(v []MicrosoftGraphDeviceConfigurationState)`

SetDeviceConfigurationStates sets DeviceConfigurationStates field to given value.

### HasDeviceConfigurationStates

`func (o *ManagedDevice) HasDeviceConfigurationStates() bool`

HasDeviceConfigurationStates returns a boolean if a field has been set.

### GetDeviceCategory

`func (o *ManagedDevice) GetDeviceCategory() AnyOfmicrosoftGraphDeviceCategory`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *ManagedDevice) GetDeviceCategoryOk() (*AnyOfmicrosoftGraphDeviceCategory, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategory

`func (o *ManagedDevice) SetDeviceCategory(v AnyOfmicrosoftGraphDeviceCategory)`

SetDeviceCategory sets DeviceCategory field to given value.

### HasDeviceCategory

`func (o *ManagedDevice) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### SetDeviceCategoryNil

`func (o *ManagedDevice) SetDeviceCategoryNil(b bool)`

 SetDeviceCategoryNil sets the value for DeviceCategory to be an explicit nil

### UnsetDeviceCategory
`func (o *ManagedDevice) UnsetDeviceCategory()`

UnsetDeviceCategory ensures that no value is present for DeviceCategory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


