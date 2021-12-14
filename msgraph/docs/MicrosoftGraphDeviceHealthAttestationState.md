# MicrosoftGraphDeviceHealthAttestationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttestationIdentityKey** | Pointer to **NullableString** | TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate. | [optional] 
**BitLockerStatus** | Pointer to **NullableString** | On or Off of BitLocker Drive Encryption | [optional] 
**BootAppSecurityVersion** | Pointer to **NullableString** | The security version number of the Boot Application | [optional] 
**BootDebugging** | Pointer to **NullableString** | When bootDebugging is enabled, the device is used in development and testing | [optional] 
**BootManagerSecurityVersion** | Pointer to **NullableString** | The security version number of the Boot Application | [optional] 
**BootManagerVersion** | Pointer to **NullableString** | The version of the Boot Manager | [optional] 
**BootRevisionListInfo** | Pointer to **NullableString** | The Boot Revision List that was loaded during initial boot on the attested device | [optional] 
**CodeIntegrity** | Pointer to **NullableString** | When code integrity is enabled, code execution is restricted to integrity verified code | [optional] 
**CodeIntegrityCheckVersion** | Pointer to **NullableString** | The version of the Boot Manager | [optional] 
**CodeIntegrityPolicy** | Pointer to **NullableString** | The Code Integrity policy that is controlling the security of the boot environment | [optional] 
**ContentNamespaceUrl** | Pointer to **NullableString** | The DHA report version. (Namespace version) | [optional] 
**ContentVersion** | Pointer to **NullableString** | The HealthAttestation state schema version | [optional] 
**DataExcutionPolicy** | Pointer to **NullableString** | DEP Policy defines a set of hardware and software technologies that perform additional checks on memory | [optional] 
**DeviceHealthAttestationStatus** | Pointer to **NullableString** | The DHA report version. (Namespace version) | [optional] 
**EarlyLaunchAntiMalwareDriverProtection** | Pointer to **NullableString** | ELAM provides protection for the computers in your network when they start up | [optional] 
**HealthAttestationSupportedStatus** | Pointer to **NullableString** | This attribute indicates if DHA is supported for the device | [optional] 
**HealthStatusMismatchInfo** | Pointer to **NullableString** | This attribute appears if DHA-Service detects an integrity issue | [optional] 
**IssuedDateTime** | Pointer to **time.Time** | The DateTime when device was evaluated or issued to MDM | [optional] 
**LastUpdateDateTime** | Pointer to **NullableString** | The Timestamp of the last update. | [optional] 
**OperatingSystemKernelDebugging** | Pointer to **NullableString** | When operatingSystemKernelDebugging is enabled, the device is used in development and testing | [optional] 
**OperatingSystemRevListInfo** | Pointer to **NullableString** | The Operating System Revision List that was loaded during initial boot on the attested device | [optional] 
**Pcr0** | Pointer to **NullableString** | The measurement that is captured in PCR[0] | [optional] 
**PcrHashAlgorithm** | Pointer to **NullableString** | Informational attribute that identifies the HASH algorithm that was used by TPM | [optional] 
**ResetCount** | Pointer to **int64** | The number of times a PC device has hibernated or resumed | [optional] 
**RestartCount** | Pointer to **int64** | The number of times a PC device has rebooted | [optional] 
**SafeMode** | Pointer to **NullableString** | Safe mode is a troubleshooting option for Windows that starts your computer in a limited state | [optional] 
**SecureBoot** | Pointer to **NullableString** | When Secure Boot is enabled, the core components must have the correct cryptographic signatures | [optional] 
**SecureBootConfigurationPolicyFingerPrint** | Pointer to **NullableString** | Fingerprint of the Custom Secure Boot Configuration Policy | [optional] 
**TestSigning** | Pointer to **NullableString** | When test signing is allowed, the device does not enforce signature validation during boot | [optional] 
**TpmVersion** | Pointer to **NullableString** | The security version number of the Boot Application | [optional] 
**VirtualSecureMode** | Pointer to **NullableString** | VSM is a container that protects high value assets from a compromised kernel | [optional] 
**WindowsPE** | Pointer to **NullableString** | Operating system running with limited services that is used to prepare a computer for Windows | [optional] 

## Methods

### NewMicrosoftGraphDeviceHealthAttestationState

`func NewMicrosoftGraphDeviceHealthAttestationState() *MicrosoftGraphDeviceHealthAttestationState`

NewMicrosoftGraphDeviceHealthAttestationState instantiates a new MicrosoftGraphDeviceHealthAttestationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceHealthAttestationStateWithDefaults

`func NewMicrosoftGraphDeviceHealthAttestationStateWithDefaults() *MicrosoftGraphDeviceHealthAttestationState`

NewMicrosoftGraphDeviceHealthAttestationStateWithDefaults instantiates a new MicrosoftGraphDeviceHealthAttestationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestationIdentityKey

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetAttestationIdentityKey() string`

GetAttestationIdentityKey returns the AttestationIdentityKey field if non-nil, zero value otherwise.

### GetAttestationIdentityKeyOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetAttestationIdentityKeyOk() (*string, bool)`

GetAttestationIdentityKeyOk returns a tuple with the AttestationIdentityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationIdentityKey

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetAttestationIdentityKey(v string)`

SetAttestationIdentityKey sets AttestationIdentityKey field to given value.

### HasAttestationIdentityKey

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasAttestationIdentityKey() bool`

HasAttestationIdentityKey returns a boolean if a field has been set.

### SetAttestationIdentityKeyNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetAttestationIdentityKeyNil(b bool)`

 SetAttestationIdentityKeyNil sets the value for AttestationIdentityKey to be an explicit nil

### UnsetAttestationIdentityKey
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetAttestationIdentityKey()`

UnsetAttestationIdentityKey ensures that no value is present for AttestationIdentityKey, not even an explicit nil
### GetBitLockerStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBitLockerStatus() string`

GetBitLockerStatus returns the BitLockerStatus field if non-nil, zero value otherwise.

### GetBitLockerStatusOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBitLockerStatusOk() (*string, bool)`

GetBitLockerStatusOk returns a tuple with the BitLockerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitLockerStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBitLockerStatus(v string)`

SetBitLockerStatus sets BitLockerStatus field to given value.

### HasBitLockerStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBitLockerStatus() bool`

HasBitLockerStatus returns a boolean if a field has been set.

### SetBitLockerStatusNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBitLockerStatusNil(b bool)`

 SetBitLockerStatusNil sets the value for BitLockerStatus to be an explicit nil

### UnsetBitLockerStatus
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetBitLockerStatus()`

UnsetBitLockerStatus ensures that no value is present for BitLockerStatus, not even an explicit nil
### GetBootAppSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootAppSecurityVersion() string`

GetBootAppSecurityVersion returns the BootAppSecurityVersion field if non-nil, zero value otherwise.

### GetBootAppSecurityVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootAppSecurityVersionOk() (*string, bool)`

GetBootAppSecurityVersionOk returns a tuple with the BootAppSecurityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootAppSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootAppSecurityVersion(v string)`

SetBootAppSecurityVersion sets BootAppSecurityVersion field to given value.

### HasBootAppSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootAppSecurityVersion() bool`

HasBootAppSecurityVersion returns a boolean if a field has been set.

### SetBootAppSecurityVersionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootAppSecurityVersionNil(b bool)`

 SetBootAppSecurityVersionNil sets the value for BootAppSecurityVersion to be an explicit nil

### UnsetBootAppSecurityVersion
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetBootAppSecurityVersion()`

UnsetBootAppSecurityVersion ensures that no value is present for BootAppSecurityVersion, not even an explicit nil
### GetBootDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootDebugging() string`

GetBootDebugging returns the BootDebugging field if non-nil, zero value otherwise.

### GetBootDebuggingOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootDebuggingOk() (*string, bool)`

GetBootDebuggingOk returns a tuple with the BootDebugging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootDebugging(v string)`

SetBootDebugging sets BootDebugging field to given value.

### HasBootDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootDebugging() bool`

HasBootDebugging returns a boolean if a field has been set.

### SetBootDebuggingNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootDebuggingNil(b bool)`

 SetBootDebuggingNil sets the value for BootDebugging to be an explicit nil

### UnsetBootDebugging
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetBootDebugging()`

UnsetBootDebugging ensures that no value is present for BootDebugging, not even an explicit nil
### GetBootManagerSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerSecurityVersion() string`

GetBootManagerSecurityVersion returns the BootManagerSecurityVersion field if non-nil, zero value otherwise.

### GetBootManagerSecurityVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerSecurityVersionOk() (*string, bool)`

GetBootManagerSecurityVersionOk returns a tuple with the BootManagerSecurityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootManagerSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerSecurityVersion(v string)`

SetBootManagerSecurityVersion sets BootManagerSecurityVersion field to given value.

### HasBootManagerSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootManagerSecurityVersion() bool`

HasBootManagerSecurityVersion returns a boolean if a field has been set.

### SetBootManagerSecurityVersionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerSecurityVersionNil(b bool)`

 SetBootManagerSecurityVersionNil sets the value for BootManagerSecurityVersion to be an explicit nil

### UnsetBootManagerSecurityVersion
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetBootManagerSecurityVersion()`

UnsetBootManagerSecurityVersion ensures that no value is present for BootManagerSecurityVersion, not even an explicit nil
### GetBootManagerVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerVersion() string`

GetBootManagerVersion returns the BootManagerVersion field if non-nil, zero value otherwise.

### GetBootManagerVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerVersionOk() (*string, bool)`

GetBootManagerVersionOk returns a tuple with the BootManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootManagerVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerVersion(v string)`

SetBootManagerVersion sets BootManagerVersion field to given value.

### HasBootManagerVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootManagerVersion() bool`

HasBootManagerVersion returns a boolean if a field has been set.

### SetBootManagerVersionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerVersionNil(b bool)`

 SetBootManagerVersionNil sets the value for BootManagerVersion to be an explicit nil

### UnsetBootManagerVersion
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetBootManagerVersion()`

UnsetBootManagerVersion ensures that no value is present for BootManagerVersion, not even an explicit nil
### GetBootRevisionListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootRevisionListInfo() string`

GetBootRevisionListInfo returns the BootRevisionListInfo field if non-nil, zero value otherwise.

### GetBootRevisionListInfoOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootRevisionListInfoOk() (*string, bool)`

GetBootRevisionListInfoOk returns a tuple with the BootRevisionListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootRevisionListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootRevisionListInfo(v string)`

SetBootRevisionListInfo sets BootRevisionListInfo field to given value.

### HasBootRevisionListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootRevisionListInfo() bool`

HasBootRevisionListInfo returns a boolean if a field has been set.

### SetBootRevisionListInfoNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootRevisionListInfoNil(b bool)`

 SetBootRevisionListInfoNil sets the value for BootRevisionListInfo to be an explicit nil

### UnsetBootRevisionListInfo
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetBootRevisionListInfo()`

UnsetBootRevisionListInfo ensures that no value is present for BootRevisionListInfo, not even an explicit nil
### GetCodeIntegrity

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrity() string`

GetCodeIntegrity returns the CodeIntegrity field if non-nil, zero value otherwise.

### GetCodeIntegrityOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityOk() (*string, bool)`

GetCodeIntegrityOk returns a tuple with the CodeIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrity

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrity(v string)`

SetCodeIntegrity sets CodeIntegrity field to given value.

### HasCodeIntegrity

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasCodeIntegrity() bool`

HasCodeIntegrity returns a boolean if a field has been set.

### SetCodeIntegrityNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityNil(b bool)`

 SetCodeIntegrityNil sets the value for CodeIntegrity to be an explicit nil

### UnsetCodeIntegrity
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetCodeIntegrity()`

UnsetCodeIntegrity ensures that no value is present for CodeIntegrity, not even an explicit nil
### GetCodeIntegrityCheckVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityCheckVersion() string`

GetCodeIntegrityCheckVersion returns the CodeIntegrityCheckVersion field if non-nil, zero value otherwise.

### GetCodeIntegrityCheckVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityCheckVersionOk() (*string, bool)`

GetCodeIntegrityCheckVersionOk returns a tuple with the CodeIntegrityCheckVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrityCheckVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityCheckVersion(v string)`

SetCodeIntegrityCheckVersion sets CodeIntegrityCheckVersion field to given value.

### HasCodeIntegrityCheckVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasCodeIntegrityCheckVersion() bool`

HasCodeIntegrityCheckVersion returns a boolean if a field has been set.

### SetCodeIntegrityCheckVersionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityCheckVersionNil(b bool)`

 SetCodeIntegrityCheckVersionNil sets the value for CodeIntegrityCheckVersion to be an explicit nil

### UnsetCodeIntegrityCheckVersion
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetCodeIntegrityCheckVersion()`

UnsetCodeIntegrityCheckVersion ensures that no value is present for CodeIntegrityCheckVersion, not even an explicit nil
### GetCodeIntegrityPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityPolicy() string`

GetCodeIntegrityPolicy returns the CodeIntegrityPolicy field if non-nil, zero value otherwise.

### GetCodeIntegrityPolicyOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityPolicyOk() (*string, bool)`

GetCodeIntegrityPolicyOk returns a tuple with the CodeIntegrityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrityPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityPolicy(v string)`

SetCodeIntegrityPolicy sets CodeIntegrityPolicy field to given value.

### HasCodeIntegrityPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasCodeIntegrityPolicy() bool`

HasCodeIntegrityPolicy returns a boolean if a field has been set.

### SetCodeIntegrityPolicyNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityPolicyNil(b bool)`

 SetCodeIntegrityPolicyNil sets the value for CodeIntegrityPolicy to be an explicit nil

### UnsetCodeIntegrityPolicy
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetCodeIntegrityPolicy()`

UnsetCodeIntegrityPolicy ensures that no value is present for CodeIntegrityPolicy, not even an explicit nil
### GetContentNamespaceUrl

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentNamespaceUrl() string`

GetContentNamespaceUrl returns the ContentNamespaceUrl field if non-nil, zero value otherwise.

### GetContentNamespaceUrlOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentNamespaceUrlOk() (*string, bool)`

GetContentNamespaceUrlOk returns a tuple with the ContentNamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentNamespaceUrl

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentNamespaceUrl(v string)`

SetContentNamespaceUrl sets ContentNamespaceUrl field to given value.

### HasContentNamespaceUrl

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasContentNamespaceUrl() bool`

HasContentNamespaceUrl returns a boolean if a field has been set.

### SetContentNamespaceUrlNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentNamespaceUrlNil(b bool)`

 SetContentNamespaceUrlNil sets the value for ContentNamespaceUrl to be an explicit nil

### UnsetContentNamespaceUrl
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetContentNamespaceUrl()`

UnsetContentNamespaceUrl ensures that no value is present for ContentNamespaceUrl, not even an explicit nil
### GetContentVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentVersion() string`

GetContentVersion returns the ContentVersion field if non-nil, zero value otherwise.

### GetContentVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentVersionOk() (*string, bool)`

GetContentVersionOk returns a tuple with the ContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentVersion(v string)`

SetContentVersion sets ContentVersion field to given value.

### HasContentVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasContentVersion() bool`

HasContentVersion returns a boolean if a field has been set.

### SetContentVersionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentVersionNil(b bool)`

 SetContentVersionNil sets the value for ContentVersion to be an explicit nil

### UnsetContentVersion
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetContentVersion()`

UnsetContentVersion ensures that no value is present for ContentVersion, not even an explicit nil
### GetDataExcutionPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDataExcutionPolicy() string`

GetDataExcutionPolicy returns the DataExcutionPolicy field if non-nil, zero value otherwise.

### GetDataExcutionPolicyOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDataExcutionPolicyOk() (*string, bool)`

GetDataExcutionPolicyOk returns a tuple with the DataExcutionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataExcutionPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDataExcutionPolicy(v string)`

SetDataExcutionPolicy sets DataExcutionPolicy field to given value.

### HasDataExcutionPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasDataExcutionPolicy() bool`

HasDataExcutionPolicy returns a boolean if a field has been set.

### SetDataExcutionPolicyNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDataExcutionPolicyNil(b bool)`

 SetDataExcutionPolicyNil sets the value for DataExcutionPolicy to be an explicit nil

### UnsetDataExcutionPolicy
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetDataExcutionPolicy()`

UnsetDataExcutionPolicy ensures that no value is present for DataExcutionPolicy, not even an explicit nil
### GetDeviceHealthAttestationStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDeviceHealthAttestationStatus() string`

GetDeviceHealthAttestationStatus returns the DeviceHealthAttestationStatus field if non-nil, zero value otherwise.

### GetDeviceHealthAttestationStatusOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDeviceHealthAttestationStatusOk() (*string, bool)`

GetDeviceHealthAttestationStatusOk returns a tuple with the DeviceHealthAttestationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealthAttestationStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDeviceHealthAttestationStatus(v string)`

SetDeviceHealthAttestationStatus sets DeviceHealthAttestationStatus field to given value.

### HasDeviceHealthAttestationStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasDeviceHealthAttestationStatus() bool`

HasDeviceHealthAttestationStatus returns a boolean if a field has been set.

### SetDeviceHealthAttestationStatusNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDeviceHealthAttestationStatusNil(b bool)`

 SetDeviceHealthAttestationStatusNil sets the value for DeviceHealthAttestationStatus to be an explicit nil

### UnsetDeviceHealthAttestationStatus
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetDeviceHealthAttestationStatus()`

UnsetDeviceHealthAttestationStatus ensures that no value is present for DeviceHealthAttestationStatus, not even an explicit nil
### GetEarlyLaunchAntiMalwareDriverProtection

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtection() string`

GetEarlyLaunchAntiMalwareDriverProtection returns the EarlyLaunchAntiMalwareDriverProtection field if non-nil, zero value otherwise.

### GetEarlyLaunchAntiMalwareDriverProtectionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtectionOk() (*string, bool)`

GetEarlyLaunchAntiMalwareDriverProtectionOk returns a tuple with the EarlyLaunchAntiMalwareDriverProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyLaunchAntiMalwareDriverProtection

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtection(v string)`

SetEarlyLaunchAntiMalwareDriverProtection sets EarlyLaunchAntiMalwareDriverProtection field to given value.

### HasEarlyLaunchAntiMalwareDriverProtection

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasEarlyLaunchAntiMalwareDriverProtection() bool`

HasEarlyLaunchAntiMalwareDriverProtection returns a boolean if a field has been set.

### SetEarlyLaunchAntiMalwareDriverProtectionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtectionNil(b bool)`

 SetEarlyLaunchAntiMalwareDriverProtectionNil sets the value for EarlyLaunchAntiMalwareDriverProtection to be an explicit nil

### UnsetEarlyLaunchAntiMalwareDriverProtection
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetEarlyLaunchAntiMalwareDriverProtection()`

UnsetEarlyLaunchAntiMalwareDriverProtection ensures that no value is present for EarlyLaunchAntiMalwareDriverProtection, not even an explicit nil
### GetHealthAttestationSupportedStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthAttestationSupportedStatus() string`

GetHealthAttestationSupportedStatus returns the HealthAttestationSupportedStatus field if non-nil, zero value otherwise.

### GetHealthAttestationSupportedStatusOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthAttestationSupportedStatusOk() (*string, bool)`

GetHealthAttestationSupportedStatusOk returns a tuple with the HealthAttestationSupportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthAttestationSupportedStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthAttestationSupportedStatus(v string)`

SetHealthAttestationSupportedStatus sets HealthAttestationSupportedStatus field to given value.

### HasHealthAttestationSupportedStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasHealthAttestationSupportedStatus() bool`

HasHealthAttestationSupportedStatus returns a boolean if a field has been set.

### SetHealthAttestationSupportedStatusNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthAttestationSupportedStatusNil(b bool)`

 SetHealthAttestationSupportedStatusNil sets the value for HealthAttestationSupportedStatus to be an explicit nil

### UnsetHealthAttestationSupportedStatus
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetHealthAttestationSupportedStatus()`

UnsetHealthAttestationSupportedStatus ensures that no value is present for HealthAttestationSupportedStatus, not even an explicit nil
### GetHealthStatusMismatchInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthStatusMismatchInfo() string`

GetHealthStatusMismatchInfo returns the HealthStatusMismatchInfo field if non-nil, zero value otherwise.

### GetHealthStatusMismatchInfoOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthStatusMismatchInfoOk() (*string, bool)`

GetHealthStatusMismatchInfoOk returns a tuple with the HealthStatusMismatchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatusMismatchInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthStatusMismatchInfo(v string)`

SetHealthStatusMismatchInfo sets HealthStatusMismatchInfo field to given value.

### HasHealthStatusMismatchInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasHealthStatusMismatchInfo() bool`

HasHealthStatusMismatchInfo returns a boolean if a field has been set.

### SetHealthStatusMismatchInfoNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthStatusMismatchInfoNil(b bool)`

 SetHealthStatusMismatchInfoNil sets the value for HealthStatusMismatchInfo to be an explicit nil

### UnsetHealthStatusMismatchInfo
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetHealthStatusMismatchInfo()`

UnsetHealthStatusMismatchInfo ensures that no value is present for HealthStatusMismatchInfo, not even an explicit nil
### GetIssuedDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetIssuedDateTime() time.Time`

GetIssuedDateTime returns the IssuedDateTime field if non-nil, zero value otherwise.

### GetIssuedDateTimeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetIssuedDateTimeOk() (*time.Time, bool)`

GetIssuedDateTimeOk returns a tuple with the IssuedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetIssuedDateTime(v time.Time)`

SetIssuedDateTime sets IssuedDateTime field to given value.

### HasIssuedDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasIssuedDateTime() bool`

HasIssuedDateTime returns a boolean if a field has been set.

### GetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetLastUpdateDateTime() string`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetLastUpdateDateTimeOk() (*string, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetLastUpdateDateTime(v string)`

SetLastUpdateDateTime sets LastUpdateDateTime field to given value.

### HasLastUpdateDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasLastUpdateDateTime() bool`

HasLastUpdateDateTime returns a boolean if a field has been set.

### SetLastUpdateDateTimeNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetLastUpdateDateTimeNil(b bool)`

 SetLastUpdateDateTimeNil sets the value for LastUpdateDateTime to be an explicit nil

### UnsetLastUpdateDateTime
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetLastUpdateDateTime()`

UnsetLastUpdateDateTime ensures that no value is present for LastUpdateDateTime, not even an explicit nil
### GetOperatingSystemKernelDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemKernelDebugging() string`

GetOperatingSystemKernelDebugging returns the OperatingSystemKernelDebugging field if non-nil, zero value otherwise.

### GetOperatingSystemKernelDebuggingOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemKernelDebuggingOk() (*string, bool)`

GetOperatingSystemKernelDebuggingOk returns a tuple with the OperatingSystemKernelDebugging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemKernelDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemKernelDebugging(v string)`

SetOperatingSystemKernelDebugging sets OperatingSystemKernelDebugging field to given value.

### HasOperatingSystemKernelDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasOperatingSystemKernelDebugging() bool`

HasOperatingSystemKernelDebugging returns a boolean if a field has been set.

### SetOperatingSystemKernelDebuggingNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemKernelDebuggingNil(b bool)`

 SetOperatingSystemKernelDebuggingNil sets the value for OperatingSystemKernelDebugging to be an explicit nil

### UnsetOperatingSystemKernelDebugging
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetOperatingSystemKernelDebugging()`

UnsetOperatingSystemKernelDebugging ensures that no value is present for OperatingSystemKernelDebugging, not even an explicit nil
### GetOperatingSystemRevListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemRevListInfo() string`

GetOperatingSystemRevListInfo returns the OperatingSystemRevListInfo field if non-nil, zero value otherwise.

### GetOperatingSystemRevListInfoOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemRevListInfoOk() (*string, bool)`

GetOperatingSystemRevListInfoOk returns a tuple with the OperatingSystemRevListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemRevListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemRevListInfo(v string)`

SetOperatingSystemRevListInfo sets OperatingSystemRevListInfo field to given value.

### HasOperatingSystemRevListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasOperatingSystemRevListInfo() bool`

HasOperatingSystemRevListInfo returns a boolean if a field has been set.

### SetOperatingSystemRevListInfoNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemRevListInfoNil(b bool)`

 SetOperatingSystemRevListInfoNil sets the value for OperatingSystemRevListInfo to be an explicit nil

### UnsetOperatingSystemRevListInfo
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetOperatingSystemRevListInfo()`

UnsetOperatingSystemRevListInfo ensures that no value is present for OperatingSystemRevListInfo, not even an explicit nil
### GetPcr0

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcr0() string`

GetPcr0 returns the Pcr0 field if non-nil, zero value otherwise.

### GetPcr0Ok

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcr0Ok() (*string, bool)`

GetPcr0Ok returns a tuple with the Pcr0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcr0

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcr0(v string)`

SetPcr0 sets Pcr0 field to given value.

### HasPcr0

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasPcr0() bool`

HasPcr0 returns a boolean if a field has been set.

### SetPcr0Nil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcr0Nil(b bool)`

 SetPcr0Nil sets the value for Pcr0 to be an explicit nil

### UnsetPcr0
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetPcr0()`

UnsetPcr0 ensures that no value is present for Pcr0, not even an explicit nil
### GetPcrHashAlgorithm

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcrHashAlgorithm() string`

GetPcrHashAlgorithm returns the PcrHashAlgorithm field if non-nil, zero value otherwise.

### GetPcrHashAlgorithmOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcrHashAlgorithmOk() (*string, bool)`

GetPcrHashAlgorithmOk returns a tuple with the PcrHashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcrHashAlgorithm

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcrHashAlgorithm(v string)`

SetPcrHashAlgorithm sets PcrHashAlgorithm field to given value.

### HasPcrHashAlgorithm

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasPcrHashAlgorithm() bool`

HasPcrHashAlgorithm returns a boolean if a field has been set.

### SetPcrHashAlgorithmNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcrHashAlgorithmNil(b bool)`

 SetPcrHashAlgorithmNil sets the value for PcrHashAlgorithm to be an explicit nil

### UnsetPcrHashAlgorithm
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetPcrHashAlgorithm()`

UnsetPcrHashAlgorithm ensures that no value is present for PcrHashAlgorithm, not even an explicit nil
### GetResetCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetResetCount() int64`

GetResetCount returns the ResetCount field if non-nil, zero value otherwise.

### GetResetCountOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetResetCountOk() (*int64, bool)`

GetResetCountOk returns a tuple with the ResetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetResetCount(v int64)`

SetResetCount sets ResetCount field to given value.

### HasResetCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasResetCount() bool`

HasResetCount returns a boolean if a field has been set.

### GetRestartCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetRestartCount() int64`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetRestartCountOk() (*int64, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetRestartCount(v int64)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### GetSafeMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSafeMode() string`

GetSafeMode returns the SafeMode field if non-nil, zero value otherwise.

### GetSafeModeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSafeModeOk() (*string, bool)`

GetSafeModeOk returns a tuple with the SafeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSafeMode(v string)`

SetSafeMode sets SafeMode field to given value.

### HasSafeMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasSafeMode() bool`

HasSafeMode returns a boolean if a field has been set.

### SetSafeModeNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSafeModeNil(b bool)`

 SetSafeModeNil sets the value for SafeMode to be an explicit nil

### UnsetSafeMode
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetSafeMode()`

UnsetSafeMode ensures that no value is present for SafeMode, not even an explicit nil
### GetSecureBoot

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBoot() string`

GetSecureBoot returns the SecureBoot field if non-nil, zero value otherwise.

### GetSecureBootOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBootOk() (*string, bool)`

GetSecureBootOk returns a tuple with the SecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBoot

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBoot(v string)`

SetSecureBoot sets SecureBoot field to given value.

### HasSecureBoot

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasSecureBoot() bool`

HasSecureBoot returns a boolean if a field has been set.

### SetSecureBootNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBootNil(b bool)`

 SetSecureBootNil sets the value for SecureBoot to be an explicit nil

### UnsetSecureBoot
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetSecureBoot()`

UnsetSecureBoot ensures that no value is present for SecureBoot, not even an explicit nil
### GetSecureBootConfigurationPolicyFingerPrint

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrint() string`

GetSecureBootConfigurationPolicyFingerPrint returns the SecureBootConfigurationPolicyFingerPrint field if non-nil, zero value otherwise.

### GetSecureBootConfigurationPolicyFingerPrintOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrintOk() (*string, bool)`

GetSecureBootConfigurationPolicyFingerPrintOk returns a tuple with the SecureBootConfigurationPolicyFingerPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootConfigurationPolicyFingerPrint

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrint(v string)`

SetSecureBootConfigurationPolicyFingerPrint sets SecureBootConfigurationPolicyFingerPrint field to given value.

### HasSecureBootConfigurationPolicyFingerPrint

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasSecureBootConfigurationPolicyFingerPrint() bool`

HasSecureBootConfigurationPolicyFingerPrint returns a boolean if a field has been set.

### SetSecureBootConfigurationPolicyFingerPrintNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrintNil(b bool)`

 SetSecureBootConfigurationPolicyFingerPrintNil sets the value for SecureBootConfigurationPolicyFingerPrint to be an explicit nil

### UnsetSecureBootConfigurationPolicyFingerPrint
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetSecureBootConfigurationPolicyFingerPrint()`

UnsetSecureBootConfigurationPolicyFingerPrint ensures that no value is present for SecureBootConfigurationPolicyFingerPrint, not even an explicit nil
### GetTestSigning

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTestSigning() string`

GetTestSigning returns the TestSigning field if non-nil, zero value otherwise.

### GetTestSigningOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTestSigningOk() (*string, bool)`

GetTestSigningOk returns a tuple with the TestSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSigning

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTestSigning(v string)`

SetTestSigning sets TestSigning field to given value.

### HasTestSigning

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasTestSigning() bool`

HasTestSigning returns a boolean if a field has been set.

### SetTestSigningNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTestSigningNil(b bool)`

 SetTestSigningNil sets the value for TestSigning to be an explicit nil

### UnsetTestSigning
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetTestSigning()`

UnsetTestSigning ensures that no value is present for TestSigning, not even an explicit nil
### GetTpmVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTpmVersion() string`

GetTpmVersion returns the TpmVersion field if non-nil, zero value otherwise.

### GetTpmVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTpmVersionOk() (*string, bool)`

GetTpmVersionOk returns a tuple with the TpmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTpmVersion(v string)`

SetTpmVersion sets TpmVersion field to given value.

### HasTpmVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasTpmVersion() bool`

HasTpmVersion returns a boolean if a field has been set.

### SetTpmVersionNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTpmVersionNil(b bool)`

 SetTpmVersionNil sets the value for TpmVersion to be an explicit nil

### UnsetTpmVersion
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetTpmVersion()`

UnsetTpmVersion ensures that no value is present for TpmVersion, not even an explicit nil
### GetVirtualSecureMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetVirtualSecureMode() string`

GetVirtualSecureMode returns the VirtualSecureMode field if non-nil, zero value otherwise.

### GetVirtualSecureModeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetVirtualSecureModeOk() (*string, bool)`

GetVirtualSecureModeOk returns a tuple with the VirtualSecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSecureMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetVirtualSecureMode(v string)`

SetVirtualSecureMode sets VirtualSecureMode field to given value.

### HasVirtualSecureMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasVirtualSecureMode() bool`

HasVirtualSecureMode returns a boolean if a field has been set.

### SetVirtualSecureModeNil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetVirtualSecureModeNil(b bool)`

 SetVirtualSecureModeNil sets the value for VirtualSecureMode to be an explicit nil

### UnsetVirtualSecureMode
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetVirtualSecureMode()`

UnsetVirtualSecureMode ensures that no value is present for VirtualSecureMode, not even an explicit nil
### GetWindowsPE

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetWindowsPE() string`

GetWindowsPE returns the WindowsPE field if non-nil, zero value otherwise.

### GetWindowsPEOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetWindowsPEOk() (*string, bool)`

GetWindowsPEOk returns a tuple with the WindowsPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPE

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetWindowsPE(v string)`

SetWindowsPE sets WindowsPE field to given value.

### HasWindowsPE

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasWindowsPE() bool`

HasWindowsPE returns a boolean if a field has been set.

### SetWindowsPENil

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetWindowsPENil(b bool)`

 SetWindowsPENil sets the value for WindowsPE to be an explicit nil

### UnsetWindowsPE
`func (o *MicrosoftGraphDeviceHealthAttestationState) UnsetWindowsPE()`

UnsetWindowsPE ensures that no value is present for WindowsPE, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


