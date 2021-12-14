# MicrosoftGraphDeviceManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**IntuneAccountId** | Pointer to **string** | Intune Account Id for given tenant | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementSettings**](anyOf&lt;microsoft.graph.deviceManagementSettings&gt;.md) | Account level settings. | [optional] 
**IntuneBrand** | Pointer to [**AnyOfmicrosoftGraphIntuneBrand**](anyOf&lt;microsoft.graph.intuneBrand&gt;.md) | intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal. | [optional] 
**SubscriptionState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementSubscriptionState**](anyOf&lt;microsoft.graph.deviceManagementSubscriptionState&gt;.md) | Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut. | [optional] 
**TermsAndConditions** | Pointer to [**[]MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md) | The terms and conditions associated with device management of the company. | [optional] 
**DeviceCompliancePolicies** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicy**](MicrosoftGraphDeviceCompliancePolicy.md) | The device compliance policies. | [optional] 
**DeviceCompliancePolicyDeviceStateSummary** | Pointer to [**AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary**](anyOf&lt;microsoft.graph.deviceCompliancePolicyDeviceStateSummary&gt;.md) | The device compliance state summary for this account. | [optional] 
**DeviceCompliancePolicySettingStateSummaries** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](MicrosoftGraphDeviceCompliancePolicySettingStateSummary.md) | The summary states of compliance policy settings for this account. | [optional] 
**DeviceConfigurationDeviceStateSummaries** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary**](anyOf&lt;microsoft.graph.deviceConfigurationDeviceStateSummary&gt;.md) | The device configuration device state summary for this account. | [optional] 
**DeviceConfigurations** | Pointer to [**[]MicrosoftGraphDeviceConfiguration**](MicrosoftGraphDeviceConfiguration.md) | The device configurations. | [optional] 
**IosUpdateStatuses** | Pointer to [**[]MicrosoftGraphIosUpdateDeviceStatus**](MicrosoftGraphIosUpdateDeviceStatus.md) | The IOS software update installation statuses for this account. | [optional] 
**SoftwareUpdateStatusSummary** | Pointer to [**AnyOfmicrosoftGraphSoftwareUpdateStatusSummary**](anyOf&lt;microsoft.graph.softwareUpdateStatusSummary&gt;.md) | The software update status summary. | [optional] 
**ComplianceManagementPartners** | Pointer to [**[]MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md) | The list of Compliance Management Partners configured by the tenant. | [optional] 
**ConditionalAccessSettings** | Pointer to [**AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings**](anyOf&lt;microsoft.graph.onPremisesConditionalAccessSettings&gt;.md) | The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access | [optional] 
**DeviceCategories** | Pointer to [**[]MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md) | The list of device categories with the tenant. | [optional] 
**DeviceEnrollmentConfigurations** | Pointer to [**[]MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md) | The list of device enrollment configurations | [optional] 
**DeviceManagementPartners** | Pointer to [**[]MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md) | The list of Device Management Partners configured by the tenant. | [optional] 
**ExchangeConnectors** | Pointer to [**[]MicrosoftGraphDeviceManagementExchangeConnector**](MicrosoftGraphDeviceManagementExchangeConnector.md) | The list of Exchange Connectors configured by the tenant. | [optional] 
**MobileThreatDefenseConnectors** | Pointer to [**[]MicrosoftGraphMobileThreatDefenseConnector**](MicrosoftGraphMobileThreatDefenseConnector.md) | The list of Mobile threat Defense connectors configured by the tenant. | [optional] 
**ApplePushNotificationCertificate** | Pointer to [**AnyOfmicrosoftGraphApplePushNotificationCertificate**](anyOf&lt;microsoft.graph.applePushNotificationCertificate&gt;.md) | Apple push notification certificate. | [optional] 
**DetectedApps** | Pointer to [**[]MicrosoftGraphDetectedApp**](MicrosoftGraphDetectedApp.md) | The list of detected apps associated with a device. | [optional] 
**ManagedDeviceOverview** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceOverview**](anyOf&lt;microsoft.graph.managedDeviceOverview&gt;.md) | Device overview | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md) | The list of managed devices. | [optional] 
**ImportedWindowsAutopilotDeviceIdentities** | Pointer to [**[]MicrosoftGraphImportedWindowsAutopilotDeviceIdentity**](MicrosoftGraphImportedWindowsAutopilotDeviceIdentity.md) | Collection of imported Windows autopilot devices. | [optional] 
**WindowsAutopilotDeviceIdentities** | Pointer to [**[]MicrosoftGraphWindowsAutopilotDeviceIdentity**](MicrosoftGraphWindowsAutopilotDeviceIdentity.md) | The Windows autopilot device identities contained collection. | [optional] 
**NotificationMessageTemplates** | Pointer to [**[]MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md) | The Notification Message Templates. | [optional] 
**ResourceOperations** | Pointer to [**[]MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md) | The Resource Operations. | [optional] 
**RoleAssignments** | Pointer to [**[]MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md) | The Role Assignments. | [optional] 
**RoleDefinitions** | Pointer to [**[]MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md) | The Role Definitions. | [optional] 
**RemoteAssistancePartners** | Pointer to [**[]MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md) | The remote assist partners. | [optional] 
**Reports** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementReports**](anyOf&lt;microsoft.graph.deviceManagementReports&gt;.md) | Reports singleton | [optional] 
**TelecomExpenseManagementPartners** | Pointer to [**[]MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md) | The telecom expense management partners. | [optional] 
**TroubleshootingEvents** | Pointer to [**[]MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md) | The list of troubleshooting events for the tenant. | [optional] 
**WindowsInformationProtectionAppLearningSummaries** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](MicrosoftGraphWindowsInformationProtectionAppLearningSummary.md) | The windows information protection app learning summaries. | [optional] 
**WindowsInformationProtectionNetworkLearningSummaries** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary.md) | The windows information protection network learning summaries. | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagement

`func NewMicrosoftGraphDeviceManagement() *MicrosoftGraphDeviceManagement`

NewMicrosoftGraphDeviceManagement instantiates a new MicrosoftGraphDeviceManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementWithDefaults

`func NewMicrosoftGraphDeviceManagementWithDefaults() *MicrosoftGraphDeviceManagement`

NewMicrosoftGraphDeviceManagementWithDefaults instantiates a new MicrosoftGraphDeviceManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceManagement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceManagement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntuneAccountId

`func (o *MicrosoftGraphDeviceManagement) GetIntuneAccountId() string`

GetIntuneAccountId returns the IntuneAccountId field if non-nil, zero value otherwise.

### GetIntuneAccountIdOk

`func (o *MicrosoftGraphDeviceManagement) GetIntuneAccountIdOk() (*string, bool)`

GetIntuneAccountIdOk returns a tuple with the IntuneAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneAccountId

`func (o *MicrosoftGraphDeviceManagement) SetIntuneAccountId(v string)`

SetIntuneAccountId sets IntuneAccountId field to given value.

### HasIntuneAccountId

`func (o *MicrosoftGraphDeviceManagement) HasIntuneAccountId() bool`

HasIntuneAccountId returns a boolean if a field has been set.

### GetSettings

`func (o *MicrosoftGraphDeviceManagement) GetSettings() AnyOfmicrosoftGraphDeviceManagementSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphDeviceManagement) GetSettingsOk() (*AnyOfmicrosoftGraphDeviceManagementSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MicrosoftGraphDeviceManagement) SetSettings(v AnyOfmicrosoftGraphDeviceManagementSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MicrosoftGraphDeviceManagement) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *MicrosoftGraphDeviceManagement) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *MicrosoftGraphDeviceManagement) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetIntuneBrand

`func (o *MicrosoftGraphDeviceManagement) GetIntuneBrand() AnyOfmicrosoftGraphIntuneBrand`

GetIntuneBrand returns the IntuneBrand field if non-nil, zero value otherwise.

### GetIntuneBrandOk

`func (o *MicrosoftGraphDeviceManagement) GetIntuneBrandOk() (*AnyOfmicrosoftGraphIntuneBrand, bool)`

GetIntuneBrandOk returns a tuple with the IntuneBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntuneBrand

`func (o *MicrosoftGraphDeviceManagement) SetIntuneBrand(v AnyOfmicrosoftGraphIntuneBrand)`

SetIntuneBrand sets IntuneBrand field to given value.

### HasIntuneBrand

`func (o *MicrosoftGraphDeviceManagement) HasIntuneBrand() bool`

HasIntuneBrand returns a boolean if a field has been set.

### SetIntuneBrandNil

`func (o *MicrosoftGraphDeviceManagement) SetIntuneBrandNil(b bool)`

 SetIntuneBrandNil sets the value for IntuneBrand to be an explicit nil

### UnsetIntuneBrand
`func (o *MicrosoftGraphDeviceManagement) UnsetIntuneBrand()`

UnsetIntuneBrand ensures that no value is present for IntuneBrand, not even an explicit nil
### GetSubscriptionState

`func (o *MicrosoftGraphDeviceManagement) GetSubscriptionState() AnyOfmicrosoftGraphDeviceManagementSubscriptionState`

GetSubscriptionState returns the SubscriptionState field if non-nil, zero value otherwise.

### GetSubscriptionStateOk

`func (o *MicrosoftGraphDeviceManagement) GetSubscriptionStateOk() (*AnyOfmicrosoftGraphDeviceManagementSubscriptionState, bool)`

GetSubscriptionStateOk returns a tuple with the SubscriptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionState

`func (o *MicrosoftGraphDeviceManagement) SetSubscriptionState(v AnyOfmicrosoftGraphDeviceManagementSubscriptionState)`

SetSubscriptionState sets SubscriptionState field to given value.

### HasSubscriptionState

`func (o *MicrosoftGraphDeviceManagement) HasSubscriptionState() bool`

HasSubscriptionState returns a boolean if a field has been set.

### SetSubscriptionStateNil

`func (o *MicrosoftGraphDeviceManagement) SetSubscriptionStateNil(b bool)`

 SetSubscriptionStateNil sets the value for SubscriptionState to be an explicit nil

### UnsetSubscriptionState
`func (o *MicrosoftGraphDeviceManagement) UnsetSubscriptionState()`

UnsetSubscriptionState ensures that no value is present for SubscriptionState, not even an explicit nil
### GetTermsAndConditions

`func (o *MicrosoftGraphDeviceManagement) GetTermsAndConditions() []MicrosoftGraphTermsAndConditions`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *MicrosoftGraphDeviceManagement) GetTermsAndConditionsOk() (*[]MicrosoftGraphTermsAndConditions, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *MicrosoftGraphDeviceManagement) SetTermsAndConditions(v []MicrosoftGraphTermsAndConditions)`

SetTermsAndConditions sets TermsAndConditions field to given value.

### HasTermsAndConditions

`func (o *MicrosoftGraphDeviceManagement) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### GetDeviceCompliancePolicies

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicies() []MicrosoftGraphDeviceCompliancePolicy`

GetDeviceCompliancePolicies returns the DeviceCompliancePolicies field if non-nil, zero value otherwise.

### GetDeviceCompliancePoliciesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePoliciesOk() (*[]MicrosoftGraphDeviceCompliancePolicy, bool)`

GetDeviceCompliancePoliciesOk returns a tuple with the DeviceCompliancePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCompliancePolicies

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicies(v []MicrosoftGraphDeviceCompliancePolicy)`

SetDeviceCompliancePolicies sets DeviceCompliancePolicies field to given value.

### HasDeviceCompliancePolicies

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCompliancePolicies() bool`

HasDeviceCompliancePolicies returns a boolean if a field has been set.

### GetDeviceCompliancePolicyDeviceStateSummary

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary() AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary`

GetDeviceCompliancePolicyDeviceStateSummary returns the DeviceCompliancePolicyDeviceStateSummary field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicyDeviceStateSummaryOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicyDeviceStateSummaryOk() (*AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary, bool)`

GetDeviceCompliancePolicyDeviceStateSummaryOk returns a tuple with the DeviceCompliancePolicyDeviceStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCompliancePolicyDeviceStateSummary

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(v AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary)`

SetDeviceCompliancePolicyDeviceStateSummary sets DeviceCompliancePolicyDeviceStateSummary field to given value.

### HasDeviceCompliancePolicyDeviceStateSummary

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCompliancePolicyDeviceStateSummary() bool`

HasDeviceCompliancePolicyDeviceStateSummary returns a boolean if a field has been set.

### SetDeviceCompliancePolicyDeviceStateSummaryNil

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicyDeviceStateSummaryNil(b bool)`

 SetDeviceCompliancePolicyDeviceStateSummaryNil sets the value for DeviceCompliancePolicyDeviceStateSummary to be an explicit nil

### UnsetDeviceCompliancePolicyDeviceStateSummary
`func (o *MicrosoftGraphDeviceManagement) UnsetDeviceCompliancePolicyDeviceStateSummary()`

UnsetDeviceCompliancePolicyDeviceStateSummary ensures that no value is present for DeviceCompliancePolicyDeviceStateSummary, not even an explicit nil
### GetDeviceCompliancePolicySettingStateSummaries

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicySettingStateSummaries() []MicrosoftGraphDeviceCompliancePolicySettingStateSummary`

GetDeviceCompliancePolicySettingStateSummaries returns the DeviceCompliancePolicySettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicySettingStateSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicySettingStateSummariesOk() (*[]MicrosoftGraphDeviceCompliancePolicySettingStateSummary, bool)`

GetDeviceCompliancePolicySettingStateSummariesOk returns a tuple with the DeviceCompliancePolicySettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCompliancePolicySettingStateSummaries

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(v []MicrosoftGraphDeviceCompliancePolicySettingStateSummary)`

SetDeviceCompliancePolicySettingStateSummaries sets DeviceCompliancePolicySettingStateSummaries field to given value.

### HasDeviceCompliancePolicySettingStateSummaries

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCompliancePolicySettingStateSummaries() bool`

HasDeviceCompliancePolicySettingStateSummaries returns a boolean if a field has been set.

### GetDeviceConfigurationDeviceStateSummaries

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurationDeviceStateSummaries() AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary`

GetDeviceConfigurationDeviceStateSummaries returns the DeviceConfigurationDeviceStateSummaries field if non-nil, zero value otherwise.

### GetDeviceConfigurationDeviceStateSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurationDeviceStateSummariesOk() (*AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary, bool)`

GetDeviceConfigurationDeviceStateSummariesOk returns a tuple with the DeviceConfigurationDeviceStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurationDeviceStateSummaries

`func (o *MicrosoftGraphDeviceManagement) SetDeviceConfigurationDeviceStateSummaries(v AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary)`

SetDeviceConfigurationDeviceStateSummaries sets DeviceConfigurationDeviceStateSummaries field to given value.

### HasDeviceConfigurationDeviceStateSummaries

`func (o *MicrosoftGraphDeviceManagement) HasDeviceConfigurationDeviceStateSummaries() bool`

HasDeviceConfigurationDeviceStateSummaries returns a boolean if a field has been set.

### SetDeviceConfigurationDeviceStateSummariesNil

`func (o *MicrosoftGraphDeviceManagement) SetDeviceConfigurationDeviceStateSummariesNil(b bool)`

 SetDeviceConfigurationDeviceStateSummariesNil sets the value for DeviceConfigurationDeviceStateSummaries to be an explicit nil

### UnsetDeviceConfigurationDeviceStateSummaries
`func (o *MicrosoftGraphDeviceManagement) UnsetDeviceConfigurationDeviceStateSummaries()`

UnsetDeviceConfigurationDeviceStateSummaries ensures that no value is present for DeviceConfigurationDeviceStateSummaries, not even an explicit nil
### GetDeviceConfigurations

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurations() []MicrosoftGraphDeviceConfiguration`

GetDeviceConfigurations returns the DeviceConfigurations field if non-nil, zero value otherwise.

### GetDeviceConfigurationsOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurationsOk() (*[]MicrosoftGraphDeviceConfiguration, bool)`

GetDeviceConfigurationsOk returns a tuple with the DeviceConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurations

`func (o *MicrosoftGraphDeviceManagement) SetDeviceConfigurations(v []MicrosoftGraphDeviceConfiguration)`

SetDeviceConfigurations sets DeviceConfigurations field to given value.

### HasDeviceConfigurations

`func (o *MicrosoftGraphDeviceManagement) HasDeviceConfigurations() bool`

HasDeviceConfigurations returns a boolean if a field has been set.

### GetIosUpdateStatuses

`func (o *MicrosoftGraphDeviceManagement) GetIosUpdateStatuses() []MicrosoftGraphIosUpdateDeviceStatus`

GetIosUpdateStatuses returns the IosUpdateStatuses field if non-nil, zero value otherwise.

### GetIosUpdateStatusesOk

`func (o *MicrosoftGraphDeviceManagement) GetIosUpdateStatusesOk() (*[]MicrosoftGraphIosUpdateDeviceStatus, bool)`

GetIosUpdateStatusesOk returns a tuple with the IosUpdateStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosUpdateStatuses

`func (o *MicrosoftGraphDeviceManagement) SetIosUpdateStatuses(v []MicrosoftGraphIosUpdateDeviceStatus)`

SetIosUpdateStatuses sets IosUpdateStatuses field to given value.

### HasIosUpdateStatuses

`func (o *MicrosoftGraphDeviceManagement) HasIosUpdateStatuses() bool`

HasIosUpdateStatuses returns a boolean if a field has been set.

### GetSoftwareUpdateStatusSummary

`func (o *MicrosoftGraphDeviceManagement) GetSoftwareUpdateStatusSummary() AnyOfmicrosoftGraphSoftwareUpdateStatusSummary`

GetSoftwareUpdateStatusSummary returns the SoftwareUpdateStatusSummary field if non-nil, zero value otherwise.

### GetSoftwareUpdateStatusSummaryOk

`func (o *MicrosoftGraphDeviceManagement) GetSoftwareUpdateStatusSummaryOk() (*AnyOfmicrosoftGraphSoftwareUpdateStatusSummary, bool)`

GetSoftwareUpdateStatusSummaryOk returns a tuple with the SoftwareUpdateStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateStatusSummary

`func (o *MicrosoftGraphDeviceManagement) SetSoftwareUpdateStatusSummary(v AnyOfmicrosoftGraphSoftwareUpdateStatusSummary)`

SetSoftwareUpdateStatusSummary sets SoftwareUpdateStatusSummary field to given value.

### HasSoftwareUpdateStatusSummary

`func (o *MicrosoftGraphDeviceManagement) HasSoftwareUpdateStatusSummary() bool`

HasSoftwareUpdateStatusSummary returns a boolean if a field has been set.

### SetSoftwareUpdateStatusSummaryNil

`func (o *MicrosoftGraphDeviceManagement) SetSoftwareUpdateStatusSummaryNil(b bool)`

 SetSoftwareUpdateStatusSummaryNil sets the value for SoftwareUpdateStatusSummary to be an explicit nil

### UnsetSoftwareUpdateStatusSummary
`func (o *MicrosoftGraphDeviceManagement) UnsetSoftwareUpdateStatusSummary()`

UnsetSoftwareUpdateStatusSummary ensures that no value is present for SoftwareUpdateStatusSummary, not even an explicit nil
### GetComplianceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) GetComplianceManagementPartners() []MicrosoftGraphComplianceManagementPartner`

GetComplianceManagementPartners returns the ComplianceManagementPartners field if non-nil, zero value otherwise.

### GetComplianceManagementPartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetComplianceManagementPartnersOk() (*[]MicrosoftGraphComplianceManagementPartner, bool)`

GetComplianceManagementPartnersOk returns a tuple with the ComplianceManagementPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) SetComplianceManagementPartners(v []MicrosoftGraphComplianceManagementPartner)`

SetComplianceManagementPartners sets ComplianceManagementPartners field to given value.

### HasComplianceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) HasComplianceManagementPartners() bool`

HasComplianceManagementPartners returns a boolean if a field has been set.

### GetConditionalAccessSettings

`func (o *MicrosoftGraphDeviceManagement) GetConditionalAccessSettings() AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings`

GetConditionalAccessSettings returns the ConditionalAccessSettings field if non-nil, zero value otherwise.

### GetConditionalAccessSettingsOk

`func (o *MicrosoftGraphDeviceManagement) GetConditionalAccessSettingsOk() (*AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings, bool)`

GetConditionalAccessSettingsOk returns a tuple with the ConditionalAccessSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccessSettings

`func (o *MicrosoftGraphDeviceManagement) SetConditionalAccessSettings(v AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings)`

SetConditionalAccessSettings sets ConditionalAccessSettings field to given value.

### HasConditionalAccessSettings

`func (o *MicrosoftGraphDeviceManagement) HasConditionalAccessSettings() bool`

HasConditionalAccessSettings returns a boolean if a field has been set.

### SetConditionalAccessSettingsNil

`func (o *MicrosoftGraphDeviceManagement) SetConditionalAccessSettingsNil(b bool)`

 SetConditionalAccessSettingsNil sets the value for ConditionalAccessSettings to be an explicit nil

### UnsetConditionalAccessSettings
`func (o *MicrosoftGraphDeviceManagement) UnsetConditionalAccessSettings()`

UnsetConditionalAccessSettings ensures that no value is present for ConditionalAccessSettings, not even an explicit nil
### GetDeviceCategories

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCategories() []MicrosoftGraphDeviceCategory`

GetDeviceCategories returns the DeviceCategories field if non-nil, zero value otherwise.

### GetDeviceCategoriesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCategoriesOk() (*[]MicrosoftGraphDeviceCategory, bool)`

GetDeviceCategoriesOk returns a tuple with the DeviceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCategories

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCategories(v []MicrosoftGraphDeviceCategory)`

SetDeviceCategories sets DeviceCategories field to given value.

### HasDeviceCategories

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCategories() bool`

HasDeviceCategories returns a boolean if a field has been set.

### GetDeviceEnrollmentConfigurations

`func (o *MicrosoftGraphDeviceManagement) GetDeviceEnrollmentConfigurations() []MicrosoftGraphDeviceEnrollmentConfiguration`

GetDeviceEnrollmentConfigurations returns the DeviceEnrollmentConfigurations field if non-nil, zero value otherwise.

### GetDeviceEnrollmentConfigurationsOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceEnrollmentConfigurationsOk() (*[]MicrosoftGraphDeviceEnrollmentConfiguration, bool)`

GetDeviceEnrollmentConfigurationsOk returns a tuple with the DeviceEnrollmentConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentConfigurations

`func (o *MicrosoftGraphDeviceManagement) SetDeviceEnrollmentConfigurations(v []MicrosoftGraphDeviceEnrollmentConfiguration)`

SetDeviceEnrollmentConfigurations sets DeviceEnrollmentConfigurations field to given value.

### HasDeviceEnrollmentConfigurations

`func (o *MicrosoftGraphDeviceManagement) HasDeviceEnrollmentConfigurations() bool`

HasDeviceEnrollmentConfigurations returns a boolean if a field has been set.

### GetDeviceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) GetDeviceManagementPartners() []MicrosoftGraphDeviceManagementPartner`

GetDeviceManagementPartners returns the DeviceManagementPartners field if non-nil, zero value otherwise.

### GetDeviceManagementPartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceManagementPartnersOk() (*[]MicrosoftGraphDeviceManagementPartner, bool)`

GetDeviceManagementPartnersOk returns a tuple with the DeviceManagementPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) SetDeviceManagementPartners(v []MicrosoftGraphDeviceManagementPartner)`

SetDeviceManagementPartners sets DeviceManagementPartners field to given value.

### HasDeviceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) HasDeviceManagementPartners() bool`

HasDeviceManagementPartners returns a boolean if a field has been set.

### GetExchangeConnectors

`func (o *MicrosoftGraphDeviceManagement) GetExchangeConnectors() []MicrosoftGraphDeviceManagementExchangeConnector`

GetExchangeConnectors returns the ExchangeConnectors field if non-nil, zero value otherwise.

### GetExchangeConnectorsOk

`func (o *MicrosoftGraphDeviceManagement) GetExchangeConnectorsOk() (*[]MicrosoftGraphDeviceManagementExchangeConnector, bool)`

GetExchangeConnectorsOk returns a tuple with the ExchangeConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeConnectors

`func (o *MicrosoftGraphDeviceManagement) SetExchangeConnectors(v []MicrosoftGraphDeviceManagementExchangeConnector)`

SetExchangeConnectors sets ExchangeConnectors field to given value.

### HasExchangeConnectors

`func (o *MicrosoftGraphDeviceManagement) HasExchangeConnectors() bool`

HasExchangeConnectors returns a boolean if a field has been set.

### GetMobileThreatDefenseConnectors

`func (o *MicrosoftGraphDeviceManagement) GetMobileThreatDefenseConnectors() []MicrosoftGraphMobileThreatDefenseConnector`

GetMobileThreatDefenseConnectors returns the MobileThreatDefenseConnectors field if non-nil, zero value otherwise.

### GetMobileThreatDefenseConnectorsOk

`func (o *MicrosoftGraphDeviceManagement) GetMobileThreatDefenseConnectorsOk() (*[]MicrosoftGraphMobileThreatDefenseConnector, bool)`

GetMobileThreatDefenseConnectorsOk returns a tuple with the MobileThreatDefenseConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileThreatDefenseConnectors

`func (o *MicrosoftGraphDeviceManagement) SetMobileThreatDefenseConnectors(v []MicrosoftGraphMobileThreatDefenseConnector)`

SetMobileThreatDefenseConnectors sets MobileThreatDefenseConnectors field to given value.

### HasMobileThreatDefenseConnectors

`func (o *MicrosoftGraphDeviceManagement) HasMobileThreatDefenseConnectors() bool`

HasMobileThreatDefenseConnectors returns a boolean if a field has been set.

### GetApplePushNotificationCertificate

`func (o *MicrosoftGraphDeviceManagement) GetApplePushNotificationCertificate() AnyOfmicrosoftGraphApplePushNotificationCertificate`

GetApplePushNotificationCertificate returns the ApplePushNotificationCertificate field if non-nil, zero value otherwise.

### GetApplePushNotificationCertificateOk

`func (o *MicrosoftGraphDeviceManagement) GetApplePushNotificationCertificateOk() (*AnyOfmicrosoftGraphApplePushNotificationCertificate, bool)`

GetApplePushNotificationCertificateOk returns a tuple with the ApplePushNotificationCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePushNotificationCertificate

`func (o *MicrosoftGraphDeviceManagement) SetApplePushNotificationCertificate(v AnyOfmicrosoftGraphApplePushNotificationCertificate)`

SetApplePushNotificationCertificate sets ApplePushNotificationCertificate field to given value.

### HasApplePushNotificationCertificate

`func (o *MicrosoftGraphDeviceManagement) HasApplePushNotificationCertificate() bool`

HasApplePushNotificationCertificate returns a boolean if a field has been set.

### SetApplePushNotificationCertificateNil

`func (o *MicrosoftGraphDeviceManagement) SetApplePushNotificationCertificateNil(b bool)`

 SetApplePushNotificationCertificateNil sets the value for ApplePushNotificationCertificate to be an explicit nil

### UnsetApplePushNotificationCertificate
`func (o *MicrosoftGraphDeviceManagement) UnsetApplePushNotificationCertificate()`

UnsetApplePushNotificationCertificate ensures that no value is present for ApplePushNotificationCertificate, not even an explicit nil
### GetDetectedApps

`func (o *MicrosoftGraphDeviceManagement) GetDetectedApps() []MicrosoftGraphDetectedApp`

GetDetectedApps returns the DetectedApps field if non-nil, zero value otherwise.

### GetDetectedAppsOk

`func (o *MicrosoftGraphDeviceManagement) GetDetectedAppsOk() (*[]MicrosoftGraphDetectedApp, bool)`

GetDetectedAppsOk returns a tuple with the DetectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedApps

`func (o *MicrosoftGraphDeviceManagement) SetDetectedApps(v []MicrosoftGraphDetectedApp)`

SetDetectedApps sets DetectedApps field to given value.

### HasDetectedApps

`func (o *MicrosoftGraphDeviceManagement) HasDetectedApps() bool`

HasDetectedApps returns a boolean if a field has been set.

### GetManagedDeviceOverview

`func (o *MicrosoftGraphDeviceManagement) GetManagedDeviceOverview() AnyOfmicrosoftGraphManagedDeviceOverview`

GetManagedDeviceOverview returns the ManagedDeviceOverview field if non-nil, zero value otherwise.

### GetManagedDeviceOverviewOk

`func (o *MicrosoftGraphDeviceManagement) GetManagedDeviceOverviewOk() (*AnyOfmicrosoftGraphManagedDeviceOverview, bool)`

GetManagedDeviceOverviewOk returns a tuple with the ManagedDeviceOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDeviceOverview

`func (o *MicrosoftGraphDeviceManagement) SetManagedDeviceOverview(v AnyOfmicrosoftGraphManagedDeviceOverview)`

SetManagedDeviceOverview sets ManagedDeviceOverview field to given value.

### HasManagedDeviceOverview

`func (o *MicrosoftGraphDeviceManagement) HasManagedDeviceOverview() bool`

HasManagedDeviceOverview returns a boolean if a field has been set.

### SetManagedDeviceOverviewNil

`func (o *MicrosoftGraphDeviceManagement) SetManagedDeviceOverviewNil(b bool)`

 SetManagedDeviceOverviewNil sets the value for ManagedDeviceOverview to be an explicit nil

### UnsetManagedDeviceOverview
`func (o *MicrosoftGraphDeviceManagement) UnsetManagedDeviceOverview()`

UnsetManagedDeviceOverview ensures that no value is present for ManagedDeviceOverview, not even an explicit nil
### GetManagedDevices

`func (o *MicrosoftGraphDeviceManagement) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *MicrosoftGraphDeviceManagement) GetManagedDevicesOk() (*[]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevices

`func (o *MicrosoftGraphDeviceManagement) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices sets ManagedDevices field to given value.

### HasManagedDevices

`func (o *MicrosoftGraphDeviceManagement) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### GetImportedWindowsAutopilotDeviceIdentities

`func (o *MicrosoftGraphDeviceManagement) GetImportedWindowsAutopilotDeviceIdentities() []MicrosoftGraphImportedWindowsAutopilotDeviceIdentity`

GetImportedWindowsAutopilotDeviceIdentities returns the ImportedWindowsAutopilotDeviceIdentities field if non-nil, zero value otherwise.

### GetImportedWindowsAutopilotDeviceIdentitiesOk

`func (o *MicrosoftGraphDeviceManagement) GetImportedWindowsAutopilotDeviceIdentitiesOk() (*[]MicrosoftGraphImportedWindowsAutopilotDeviceIdentity, bool)`

GetImportedWindowsAutopilotDeviceIdentitiesOk returns a tuple with the ImportedWindowsAutopilotDeviceIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedWindowsAutopilotDeviceIdentities

`func (o *MicrosoftGraphDeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(v []MicrosoftGraphImportedWindowsAutopilotDeviceIdentity)`

SetImportedWindowsAutopilotDeviceIdentities sets ImportedWindowsAutopilotDeviceIdentities field to given value.

### HasImportedWindowsAutopilotDeviceIdentities

`func (o *MicrosoftGraphDeviceManagement) HasImportedWindowsAutopilotDeviceIdentities() bool`

HasImportedWindowsAutopilotDeviceIdentities returns a boolean if a field has been set.

### GetWindowsAutopilotDeviceIdentities

`func (o *MicrosoftGraphDeviceManagement) GetWindowsAutopilotDeviceIdentities() []MicrosoftGraphWindowsAutopilotDeviceIdentity`

GetWindowsAutopilotDeviceIdentities returns the WindowsAutopilotDeviceIdentities field if non-nil, zero value otherwise.

### GetWindowsAutopilotDeviceIdentitiesOk

`func (o *MicrosoftGraphDeviceManagement) GetWindowsAutopilotDeviceIdentitiesOk() (*[]MicrosoftGraphWindowsAutopilotDeviceIdentity, bool)`

GetWindowsAutopilotDeviceIdentitiesOk returns a tuple with the WindowsAutopilotDeviceIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsAutopilotDeviceIdentities

`func (o *MicrosoftGraphDeviceManagement) SetWindowsAutopilotDeviceIdentities(v []MicrosoftGraphWindowsAutopilotDeviceIdentity)`

SetWindowsAutopilotDeviceIdentities sets WindowsAutopilotDeviceIdentities field to given value.

### HasWindowsAutopilotDeviceIdentities

`func (o *MicrosoftGraphDeviceManagement) HasWindowsAutopilotDeviceIdentities() bool`

HasWindowsAutopilotDeviceIdentities returns a boolean if a field has been set.

### GetNotificationMessageTemplates

`func (o *MicrosoftGraphDeviceManagement) GetNotificationMessageTemplates() []MicrosoftGraphNotificationMessageTemplate`

GetNotificationMessageTemplates returns the NotificationMessageTemplates field if non-nil, zero value otherwise.

### GetNotificationMessageTemplatesOk

`func (o *MicrosoftGraphDeviceManagement) GetNotificationMessageTemplatesOk() (*[]MicrosoftGraphNotificationMessageTemplate, bool)`

GetNotificationMessageTemplatesOk returns a tuple with the NotificationMessageTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessageTemplates

`func (o *MicrosoftGraphDeviceManagement) SetNotificationMessageTemplates(v []MicrosoftGraphNotificationMessageTemplate)`

SetNotificationMessageTemplates sets NotificationMessageTemplates field to given value.

### HasNotificationMessageTemplates

`func (o *MicrosoftGraphDeviceManagement) HasNotificationMessageTemplates() bool`

HasNotificationMessageTemplates returns a boolean if a field has been set.

### GetResourceOperations

`func (o *MicrosoftGraphDeviceManagement) GetResourceOperations() []MicrosoftGraphResourceOperation`

GetResourceOperations returns the ResourceOperations field if non-nil, zero value otherwise.

### GetResourceOperationsOk

`func (o *MicrosoftGraphDeviceManagement) GetResourceOperationsOk() (*[]MicrosoftGraphResourceOperation, bool)`

GetResourceOperationsOk returns a tuple with the ResourceOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOperations

`func (o *MicrosoftGraphDeviceManagement) SetResourceOperations(v []MicrosoftGraphResourceOperation)`

SetResourceOperations sets ResourceOperations field to given value.

### HasResourceOperations

`func (o *MicrosoftGraphDeviceManagement) HasResourceOperations() bool`

HasResourceOperations returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *MicrosoftGraphDeviceManagement) GetRoleAssignments() []MicrosoftGraphDeviceAndAppManagementRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *MicrosoftGraphDeviceManagement) GetRoleAssignmentsOk() (*[]MicrosoftGraphDeviceAndAppManagementRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *MicrosoftGraphDeviceManagement) SetRoleAssignments(v []MicrosoftGraphDeviceAndAppManagementRoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *MicrosoftGraphDeviceManagement) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### GetRoleDefinitions

`func (o *MicrosoftGraphDeviceManagement) GetRoleDefinitions() []MicrosoftGraphRoleDefinition`

GetRoleDefinitions returns the RoleDefinitions field if non-nil, zero value otherwise.

### GetRoleDefinitionsOk

`func (o *MicrosoftGraphDeviceManagement) GetRoleDefinitionsOk() (*[]MicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionsOk returns a tuple with the RoleDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleDefinitions

`func (o *MicrosoftGraphDeviceManagement) SetRoleDefinitions(v []MicrosoftGraphRoleDefinition)`

SetRoleDefinitions sets RoleDefinitions field to given value.

### HasRoleDefinitions

`func (o *MicrosoftGraphDeviceManagement) HasRoleDefinitions() bool`

HasRoleDefinitions returns a boolean if a field has been set.

### GetRemoteAssistancePartners

`func (o *MicrosoftGraphDeviceManagement) GetRemoteAssistancePartners() []MicrosoftGraphRemoteAssistancePartner`

GetRemoteAssistancePartners returns the RemoteAssistancePartners field if non-nil, zero value otherwise.

### GetRemoteAssistancePartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetRemoteAssistancePartnersOk() (*[]MicrosoftGraphRemoteAssistancePartner, bool)`

GetRemoteAssistancePartnersOk returns a tuple with the RemoteAssistancePartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAssistancePartners

`func (o *MicrosoftGraphDeviceManagement) SetRemoteAssistancePartners(v []MicrosoftGraphRemoteAssistancePartner)`

SetRemoteAssistancePartners sets RemoteAssistancePartners field to given value.

### HasRemoteAssistancePartners

`func (o *MicrosoftGraphDeviceManagement) HasRemoteAssistancePartners() bool`

HasRemoteAssistancePartners returns a boolean if a field has been set.

### GetReports

`func (o *MicrosoftGraphDeviceManagement) GetReports() AnyOfmicrosoftGraphDeviceManagementReports`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *MicrosoftGraphDeviceManagement) GetReportsOk() (*AnyOfmicrosoftGraphDeviceManagementReports, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *MicrosoftGraphDeviceManagement) SetReports(v AnyOfmicrosoftGraphDeviceManagementReports)`

SetReports sets Reports field to given value.

### HasReports

`func (o *MicrosoftGraphDeviceManagement) HasReports() bool`

HasReports returns a boolean if a field has been set.

### SetReportsNil

`func (o *MicrosoftGraphDeviceManagement) SetReportsNil(b bool)`

 SetReportsNil sets the value for Reports to be an explicit nil

### UnsetReports
`func (o *MicrosoftGraphDeviceManagement) UnsetReports()`

UnsetReports ensures that no value is present for Reports, not even an explicit nil
### GetTelecomExpenseManagementPartners

`func (o *MicrosoftGraphDeviceManagement) GetTelecomExpenseManagementPartners() []MicrosoftGraphTelecomExpenseManagementPartner`

GetTelecomExpenseManagementPartners returns the TelecomExpenseManagementPartners field if non-nil, zero value otherwise.

### GetTelecomExpenseManagementPartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetTelecomExpenseManagementPartnersOk() (*[]MicrosoftGraphTelecomExpenseManagementPartner, bool)`

GetTelecomExpenseManagementPartnersOk returns a tuple with the TelecomExpenseManagementPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelecomExpenseManagementPartners

`func (o *MicrosoftGraphDeviceManagement) SetTelecomExpenseManagementPartners(v []MicrosoftGraphTelecomExpenseManagementPartner)`

SetTelecomExpenseManagementPartners sets TelecomExpenseManagementPartners field to given value.

### HasTelecomExpenseManagementPartners

`func (o *MicrosoftGraphDeviceManagement) HasTelecomExpenseManagementPartners() bool`

HasTelecomExpenseManagementPartners returns a boolean if a field has been set.

### GetTroubleshootingEvents

`func (o *MicrosoftGraphDeviceManagement) GetTroubleshootingEvents() []MicrosoftGraphDeviceManagementTroubleshootingEvent`

GetTroubleshootingEvents returns the TroubleshootingEvents field if non-nil, zero value otherwise.

### GetTroubleshootingEventsOk

`func (o *MicrosoftGraphDeviceManagement) GetTroubleshootingEventsOk() (*[]MicrosoftGraphDeviceManagementTroubleshootingEvent, bool)`

GetTroubleshootingEventsOk returns a tuple with the TroubleshootingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootingEvents

`func (o *MicrosoftGraphDeviceManagement) SetTroubleshootingEvents(v []MicrosoftGraphDeviceManagementTroubleshootingEvent)`

SetTroubleshootingEvents sets TroubleshootingEvents field to given value.

### HasTroubleshootingEvents

`func (o *MicrosoftGraphDeviceManagement) HasTroubleshootingEvents() bool`

HasTroubleshootingEvents returns a boolean if a field has been set.

### GetWindowsInformationProtectionAppLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionAppLearningSummaries() []MicrosoftGraphWindowsInformationProtectionAppLearningSummary`

GetWindowsInformationProtectionAppLearningSummaries returns the WindowsInformationProtectionAppLearningSummaries field if non-nil, zero value otherwise.

### GetWindowsInformationProtectionAppLearningSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionAppLearningSummariesOk() (*[]MicrosoftGraphWindowsInformationProtectionAppLearningSummary, bool)`

GetWindowsInformationProtectionAppLearningSummariesOk returns a tuple with the WindowsInformationProtectionAppLearningSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsInformationProtectionAppLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(v []MicrosoftGraphWindowsInformationProtectionAppLearningSummary)`

SetWindowsInformationProtectionAppLearningSummaries sets WindowsInformationProtectionAppLearningSummaries field to given value.

### HasWindowsInformationProtectionAppLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) HasWindowsInformationProtectionAppLearningSummaries() bool`

HasWindowsInformationProtectionAppLearningSummaries returns a boolean if a field has been set.

### GetWindowsInformationProtectionNetworkLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries() []MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary`

GetWindowsInformationProtectionNetworkLearningSummaries returns the WindowsInformationProtectionNetworkLearningSummaries field if non-nil, zero value otherwise.

### GetWindowsInformationProtectionNetworkLearningSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionNetworkLearningSummariesOk() (*[]MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary, bool)`

GetWindowsInformationProtectionNetworkLearningSummariesOk returns a tuple with the WindowsInformationProtectionNetworkLearningSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsInformationProtectionNetworkLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(v []MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary)`

SetWindowsInformationProtectionNetworkLearningSummaries sets WindowsInformationProtectionNetworkLearningSummaries field to given value.

### HasWindowsInformationProtectionNetworkLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) HasWindowsInformationProtectionNetworkLearningSummaries() bool`

HasWindowsInformationProtectionNetworkLearningSummaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


