# IosManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDataEncryptionType** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataEncryptionType**](anyOf&lt;microsoft.graph.managedAppDataEncryptionType&gt;.md) | Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked. | [optional] 
**CustomBrowserProtocol** | Pointer to **NullableString** | A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**FaceIdBlocked** | Pointer to **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. | [optional] 
**MinimumRequiredSdkVersion** | Pointer to **NullableString** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | List of apps to which the policy is deployed. | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) | Navigation property to deployment summary of the configuration. | [optional] 

## Methods

### NewIosManagedAppProtection

`func NewIosManagedAppProtection() *IosManagedAppProtection`

NewIosManagedAppProtection instantiates a new IosManagedAppProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIosManagedAppProtectionWithDefaults

`func NewIosManagedAppProtectionWithDefaults() *IosManagedAppProtection`

NewIosManagedAppProtectionWithDefaults instantiates a new IosManagedAppProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDataEncryptionType

`func (o *IosManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *IosManagedAppProtection) GetAppDataEncryptionTypeOk() (*AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDataEncryptionType

`func (o *IosManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType sets AppDataEncryptionType field to given value.

### HasAppDataEncryptionType

`func (o *IosManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionTypeNil

`func (o *IosManagedAppProtection) SetAppDataEncryptionTypeNil(b bool)`

 SetAppDataEncryptionTypeNil sets the value for AppDataEncryptionType to be an explicit nil

### UnsetAppDataEncryptionType
`func (o *IosManagedAppProtection) UnsetAppDataEncryptionType()`

UnsetAppDataEncryptionType ensures that no value is present for AppDataEncryptionType, not even an explicit nil
### GetCustomBrowserProtocol

`func (o *IosManagedAppProtection) GetCustomBrowserProtocol() string`

GetCustomBrowserProtocol returns the CustomBrowserProtocol field if non-nil, zero value otherwise.

### GetCustomBrowserProtocolOk

`func (o *IosManagedAppProtection) GetCustomBrowserProtocolOk() (*string, bool)`

GetCustomBrowserProtocolOk returns a tuple with the CustomBrowserProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBrowserProtocol

`func (o *IosManagedAppProtection) SetCustomBrowserProtocol(v string)`

SetCustomBrowserProtocol sets CustomBrowserProtocol field to given value.

### HasCustomBrowserProtocol

`func (o *IosManagedAppProtection) HasCustomBrowserProtocol() bool`

HasCustomBrowserProtocol returns a boolean if a field has been set.

### SetCustomBrowserProtocolNil

`func (o *IosManagedAppProtection) SetCustomBrowserProtocolNil(b bool)`

 SetCustomBrowserProtocolNil sets the value for CustomBrowserProtocol to be an explicit nil

### UnsetCustomBrowserProtocol
`func (o *IosManagedAppProtection) UnsetCustomBrowserProtocol()`

UnsetCustomBrowserProtocol ensures that no value is present for CustomBrowserProtocol, not even an explicit nil
### GetDeployedAppCount

`func (o *IosManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *IosManagedAppProtection) GetDeployedAppCountOk() (*int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAppCount

`func (o *IosManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount sets DeployedAppCount field to given value.

### HasDeployedAppCount

`func (o *IosManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### GetFaceIdBlocked

`func (o *IosManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *IosManagedAppProtection) GetFaceIdBlockedOk() (*bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceIdBlocked

`func (o *IosManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked sets FaceIdBlocked field to given value.

### HasFaceIdBlocked

`func (o *IosManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### GetMinimumRequiredSdkVersion

`func (o *IosManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *IosManagedAppProtection) GetMinimumRequiredSdkVersionOk() (*string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredSdkVersion

`func (o *IosManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion sets MinimumRequiredSdkVersion field to given value.

### HasMinimumRequiredSdkVersion

`func (o *IosManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersionNil

`func (o *IosManagedAppProtection) SetMinimumRequiredSdkVersionNil(b bool)`

 SetMinimumRequiredSdkVersionNil sets the value for MinimumRequiredSdkVersion to be an explicit nil

### UnsetMinimumRequiredSdkVersion
`func (o *IosManagedAppProtection) UnsetMinimumRequiredSdkVersion()`

UnsetMinimumRequiredSdkVersion ensures that no value is present for MinimumRequiredSdkVersion, not even an explicit nil
### GetApps

`func (o *IosManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *IosManagedAppProtection) GetAppsOk() (*[]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *IosManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *IosManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDeploymentSummary

`func (o *IosManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *IosManagedAppProtection) GetDeploymentSummaryOk() (*AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSummary

`func (o *IosManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary sets DeploymentSummary field to given value.

### HasDeploymentSummary

`func (o *IosManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummaryNil

`func (o *IosManagedAppProtection) SetDeploymentSummaryNil(b bool)`

 SetDeploymentSummaryNil sets the value for DeploymentSummary to be an explicit nil

### UnsetDeploymentSummary
`func (o *IosManagedAppProtection) UnsetDeploymentSummary()`

UnsetDeploymentSummary ensures that no value is present for DeploymentSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


