# WindowsInformationProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureRightsManagementServicesAllowed** | Pointer to **bool** | Specifies whether to allow Azure RMS encryption for WIP | [optional] 
**DataRecoveryCertificate** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate**](anyOf&lt;microsoft.graph.windowsInformationProtectionDataRecoveryCertificate&gt;.md) | Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS) | [optional] 
**EnforcementLevel** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel**](anyOf&lt;microsoft.graph.windowsInformationProtectionEnforcementLevel&gt;.md) | WIP enforcement level.See the Enum definition for supported values. Possible values are: noProtection, encryptAndAuditOnly, encryptAuditAndPrompt, encryptAuditAndBlock. | [optional] 
**EnterpriseDomain** | Pointer to **NullableString** | Primary enterprise domain | [optional] 
**EnterpriseInternalProxyServers** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection.md) | This is the comma-separated list of internal proxy servers. For example, &#39;157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59&#39;. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies | [optional] 
**EnterpriseIPRanges** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection.md) | Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to | [optional] 
**EnterpriseIPRangesAreAuthoritative** | Pointer to **bool** | Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false | [optional] 
**EnterpriseNetworkDomainNames** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection.md) | This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to | [optional] 
**EnterpriseProtectedDomainNames** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection.md) | List of enterprise domains to be protected | [optional] 
**EnterpriseProxiedDomains** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection.md) | Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy | [optional] 
**EnterpriseProxyServers** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection.md) | This is a list of proxy servers. Any server not on this list is considered non-enterprise | [optional] 
**EnterpriseProxyServersAreAuthoritative** | Pointer to **bool** | Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false | [optional] 
**ExemptApps** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionApp**](AnyOfmicrosoftGraphWindowsInformationProtectionApp.md) | Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data. | [optional] 
**IconsVisible** | Pointer to **bool** | Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app | [optional] 
**IndexingEncryptedStoresOrItemsBlocked** | Pointer to **bool** | This switch is for the Windows Search Indexer, to allow or disallow indexing of items | [optional] 
**IsAssigned** | Pointer to **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**NeutralDomainResources** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection.md) | List of domain names that can used for work or personal resource | [optional] 
**ProtectedApps** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionApp**](AnyOfmicrosoftGraphWindowsInformationProtectionApp.md) | Protected applications can access enterprise data and the data handled by those applications are protected with encryption | [optional] 
**ProtectionUnderLockConfigRequired** | Pointer to **bool** | Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured | [optional] 
**RevokeOnUnenrollDisabled** | Pointer to **bool** | This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don&#39;t revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently. | [optional] 
**RightsManagementServicesTemplateId** | Pointer to **NullableString** | TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access | [optional] 
**SmbAutoEncryptedFileExtensions** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection.md) | Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](MicrosoftGraphTargetedManagedAppPolicyAssignment.md) | Navigation property to list of security groups targeted for policy. | [optional] 
**ExemptAppLockerFiles** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionAppLockerFile**](MicrosoftGraphWindowsInformationProtectionAppLockerFile.md) | Another way to input exempt apps through xml files | [optional] 
**ProtectedAppLockerFiles** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionAppLockerFile**](MicrosoftGraphWindowsInformationProtectionAppLockerFile.md) | Another way to input protected apps through xml files | [optional] 

## Methods

### NewWindowsInformationProtection

`func NewWindowsInformationProtection() *WindowsInformationProtection`

NewWindowsInformationProtection instantiates a new WindowsInformationProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsInformationProtectionWithDefaults

`func NewWindowsInformationProtectionWithDefaults() *WindowsInformationProtection`

NewWindowsInformationProtectionWithDefaults instantiates a new WindowsInformationProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureRightsManagementServicesAllowed

`func (o *WindowsInformationProtection) GetAzureRightsManagementServicesAllowed() bool`

GetAzureRightsManagementServicesAllowed returns the AzureRightsManagementServicesAllowed field if non-nil, zero value otherwise.

### GetAzureRightsManagementServicesAllowedOk

`func (o *WindowsInformationProtection) GetAzureRightsManagementServicesAllowedOk() (*bool, bool)`

GetAzureRightsManagementServicesAllowedOk returns a tuple with the AzureRightsManagementServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRightsManagementServicesAllowed

`func (o *WindowsInformationProtection) SetAzureRightsManagementServicesAllowed(v bool)`

SetAzureRightsManagementServicesAllowed sets AzureRightsManagementServicesAllowed field to given value.

### HasAzureRightsManagementServicesAllowed

`func (o *WindowsInformationProtection) HasAzureRightsManagementServicesAllowed() bool`

HasAzureRightsManagementServicesAllowed returns a boolean if a field has been set.

### GetDataRecoveryCertificate

`func (o *WindowsInformationProtection) GetDataRecoveryCertificate() AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate`

GetDataRecoveryCertificate returns the DataRecoveryCertificate field if non-nil, zero value otherwise.

### GetDataRecoveryCertificateOk

`func (o *WindowsInformationProtection) GetDataRecoveryCertificateOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate, bool)`

GetDataRecoveryCertificateOk returns a tuple with the DataRecoveryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRecoveryCertificate

`func (o *WindowsInformationProtection) SetDataRecoveryCertificate(v AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate)`

SetDataRecoveryCertificate sets DataRecoveryCertificate field to given value.

### HasDataRecoveryCertificate

`func (o *WindowsInformationProtection) HasDataRecoveryCertificate() bool`

HasDataRecoveryCertificate returns a boolean if a field has been set.

### SetDataRecoveryCertificateNil

`func (o *WindowsInformationProtection) SetDataRecoveryCertificateNil(b bool)`

 SetDataRecoveryCertificateNil sets the value for DataRecoveryCertificate to be an explicit nil

### UnsetDataRecoveryCertificate
`func (o *WindowsInformationProtection) UnsetDataRecoveryCertificate()`

UnsetDataRecoveryCertificate ensures that no value is present for DataRecoveryCertificate, not even an explicit nil
### GetEnforcementLevel

`func (o *WindowsInformationProtection) GetEnforcementLevel() AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *WindowsInformationProtection) GetEnforcementLevelOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementLevel

`func (o *WindowsInformationProtection) SetEnforcementLevel(v AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel)`

SetEnforcementLevel sets EnforcementLevel field to given value.

### HasEnforcementLevel

`func (o *WindowsInformationProtection) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### SetEnforcementLevelNil

`func (o *WindowsInformationProtection) SetEnforcementLevelNil(b bool)`

 SetEnforcementLevelNil sets the value for EnforcementLevel to be an explicit nil

### UnsetEnforcementLevel
`func (o *WindowsInformationProtection) UnsetEnforcementLevel()`

UnsetEnforcementLevel ensures that no value is present for EnforcementLevel, not even an explicit nil
### GetEnterpriseDomain

`func (o *WindowsInformationProtection) GetEnterpriseDomain() string`

GetEnterpriseDomain returns the EnterpriseDomain field if non-nil, zero value otherwise.

### GetEnterpriseDomainOk

`func (o *WindowsInformationProtection) GetEnterpriseDomainOk() (*string, bool)`

GetEnterpriseDomainOk returns a tuple with the EnterpriseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseDomain

`func (o *WindowsInformationProtection) SetEnterpriseDomain(v string)`

SetEnterpriseDomain sets EnterpriseDomain field to given value.

### HasEnterpriseDomain

`func (o *WindowsInformationProtection) HasEnterpriseDomain() bool`

HasEnterpriseDomain returns a boolean if a field has been set.

### SetEnterpriseDomainNil

`func (o *WindowsInformationProtection) SetEnterpriseDomainNil(b bool)`

 SetEnterpriseDomainNil sets the value for EnterpriseDomain to be an explicit nil

### UnsetEnterpriseDomain
`func (o *WindowsInformationProtection) UnsetEnterpriseDomain()`

UnsetEnterpriseDomain ensures that no value is present for EnterpriseDomain, not even an explicit nil
### GetEnterpriseInternalProxyServers

`func (o *WindowsInformationProtection) GetEnterpriseInternalProxyServers() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseInternalProxyServers returns the EnterpriseInternalProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseInternalProxyServersOk

`func (o *WindowsInformationProtection) GetEnterpriseInternalProxyServersOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseInternalProxyServersOk returns a tuple with the EnterpriseInternalProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseInternalProxyServers

`func (o *WindowsInformationProtection) SetEnterpriseInternalProxyServers(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseInternalProxyServers sets EnterpriseInternalProxyServers field to given value.

### HasEnterpriseInternalProxyServers

`func (o *WindowsInformationProtection) HasEnterpriseInternalProxyServers() bool`

HasEnterpriseInternalProxyServers returns a boolean if a field has been set.

### GetEnterpriseIPRanges

`func (o *WindowsInformationProtection) GetEnterpriseIPRanges() []*AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection`

GetEnterpriseIPRanges returns the EnterpriseIPRanges field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesOk

`func (o *WindowsInformationProtection) GetEnterpriseIPRangesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection, bool)`

GetEnterpriseIPRangesOk returns a tuple with the EnterpriseIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseIPRanges

`func (o *WindowsInformationProtection) SetEnterpriseIPRanges(v []*AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection)`

SetEnterpriseIPRanges sets EnterpriseIPRanges field to given value.

### HasEnterpriseIPRanges

`func (o *WindowsInformationProtection) HasEnterpriseIPRanges() bool`

HasEnterpriseIPRanges returns a boolean if a field has been set.

### GetEnterpriseIPRangesAreAuthoritative

`func (o *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritative() bool`

GetEnterpriseIPRangesAreAuthoritative returns the EnterpriseIPRangesAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesAreAuthoritativeOk

`func (o *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritativeOk() (*bool, bool)`

GetEnterpriseIPRangesAreAuthoritativeOk returns a tuple with the EnterpriseIPRangesAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseIPRangesAreAuthoritative

`func (o *WindowsInformationProtection) SetEnterpriseIPRangesAreAuthoritative(v bool)`

SetEnterpriseIPRangesAreAuthoritative sets EnterpriseIPRangesAreAuthoritative field to given value.

### HasEnterpriseIPRangesAreAuthoritative

`func (o *WindowsInformationProtection) HasEnterpriseIPRangesAreAuthoritative() bool`

HasEnterpriseIPRangesAreAuthoritative returns a boolean if a field has been set.

### GetEnterpriseNetworkDomainNames

`func (o *WindowsInformationProtection) GetEnterpriseNetworkDomainNames() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseNetworkDomainNames returns the EnterpriseNetworkDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseNetworkDomainNamesOk

`func (o *WindowsInformationProtection) GetEnterpriseNetworkDomainNamesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseNetworkDomainNamesOk returns a tuple with the EnterpriseNetworkDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseNetworkDomainNames

`func (o *WindowsInformationProtection) SetEnterpriseNetworkDomainNames(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseNetworkDomainNames sets EnterpriseNetworkDomainNames field to given value.

### HasEnterpriseNetworkDomainNames

`func (o *WindowsInformationProtection) HasEnterpriseNetworkDomainNames() bool`

HasEnterpriseNetworkDomainNames returns a boolean if a field has been set.

### GetEnterpriseProtectedDomainNames

`func (o *WindowsInformationProtection) GetEnterpriseProtectedDomainNames() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProtectedDomainNames returns the EnterpriseProtectedDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseProtectedDomainNamesOk

`func (o *WindowsInformationProtection) GetEnterpriseProtectedDomainNamesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProtectedDomainNamesOk returns a tuple with the EnterpriseProtectedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProtectedDomainNames

`func (o *WindowsInformationProtection) SetEnterpriseProtectedDomainNames(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProtectedDomainNames sets EnterpriseProtectedDomainNames field to given value.

### HasEnterpriseProtectedDomainNames

`func (o *WindowsInformationProtection) HasEnterpriseProtectedDomainNames() bool`

HasEnterpriseProtectedDomainNames returns a boolean if a field has been set.

### GetEnterpriseProxiedDomains

`func (o *WindowsInformationProtection) GetEnterpriseProxiedDomains() []*AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection`

GetEnterpriseProxiedDomains returns the EnterpriseProxiedDomains field if non-nil, zero value otherwise.

### GetEnterpriseProxiedDomainsOk

`func (o *WindowsInformationProtection) GetEnterpriseProxiedDomainsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection, bool)`

GetEnterpriseProxiedDomainsOk returns a tuple with the EnterpriseProxiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProxiedDomains

`func (o *WindowsInformationProtection) SetEnterpriseProxiedDomains(v []*AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection)`

SetEnterpriseProxiedDomains sets EnterpriseProxiedDomains field to given value.

### HasEnterpriseProxiedDomains

`func (o *WindowsInformationProtection) HasEnterpriseProxiedDomains() bool`

HasEnterpriseProxiedDomains returns a boolean if a field has been set.

### GetEnterpriseProxyServers

`func (o *WindowsInformationProtection) GetEnterpriseProxyServers() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProxyServers returns the EnterpriseProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersOk

`func (o *WindowsInformationProtection) GetEnterpriseProxyServersOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProxyServersOk returns a tuple with the EnterpriseProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProxyServers

`func (o *WindowsInformationProtection) SetEnterpriseProxyServers(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProxyServers sets EnterpriseProxyServers field to given value.

### HasEnterpriseProxyServers

`func (o *WindowsInformationProtection) HasEnterpriseProxyServers() bool`

HasEnterpriseProxyServers returns a boolean if a field has been set.

### GetEnterpriseProxyServersAreAuthoritative

`func (o *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritative() bool`

GetEnterpriseProxyServersAreAuthoritative returns the EnterpriseProxyServersAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersAreAuthoritativeOk

`func (o *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritativeOk() (*bool, bool)`

GetEnterpriseProxyServersAreAuthoritativeOk returns a tuple with the EnterpriseProxyServersAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProxyServersAreAuthoritative

`func (o *WindowsInformationProtection) SetEnterpriseProxyServersAreAuthoritative(v bool)`

SetEnterpriseProxyServersAreAuthoritative sets EnterpriseProxyServersAreAuthoritative field to given value.

### HasEnterpriseProxyServersAreAuthoritative

`func (o *WindowsInformationProtection) HasEnterpriseProxyServersAreAuthoritative() bool`

HasEnterpriseProxyServersAreAuthoritative returns a boolean if a field has been set.

### GetExemptApps

`func (o *WindowsInformationProtection) GetExemptApps() []*AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetExemptApps returns the ExemptApps field if non-nil, zero value otherwise.

### GetExemptAppsOk

`func (o *WindowsInformationProtection) GetExemptAppsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetExemptAppsOk returns a tuple with the ExemptApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptApps

`func (o *WindowsInformationProtection) SetExemptApps(v []*AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetExemptApps sets ExemptApps field to given value.

### HasExemptApps

`func (o *WindowsInformationProtection) HasExemptApps() bool`

HasExemptApps returns a boolean if a field has been set.

### GetIconsVisible

`func (o *WindowsInformationProtection) GetIconsVisible() bool`

GetIconsVisible returns the IconsVisible field if non-nil, zero value otherwise.

### GetIconsVisibleOk

`func (o *WindowsInformationProtection) GetIconsVisibleOk() (*bool, bool)`

GetIconsVisibleOk returns a tuple with the IconsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconsVisible

`func (o *WindowsInformationProtection) SetIconsVisible(v bool)`

SetIconsVisible sets IconsVisible field to given value.

### HasIconsVisible

`func (o *WindowsInformationProtection) HasIconsVisible() bool`

HasIconsVisible returns a boolean if a field has been set.

### GetIndexingEncryptedStoresOrItemsBlocked

`func (o *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlocked() bool`

GetIndexingEncryptedStoresOrItemsBlocked returns the IndexingEncryptedStoresOrItemsBlocked field if non-nil, zero value otherwise.

### GetIndexingEncryptedStoresOrItemsBlockedOk

`func (o *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlockedOk() (*bool, bool)`

GetIndexingEncryptedStoresOrItemsBlockedOk returns a tuple with the IndexingEncryptedStoresOrItemsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingEncryptedStoresOrItemsBlocked

`func (o *WindowsInformationProtection) SetIndexingEncryptedStoresOrItemsBlocked(v bool)`

SetIndexingEncryptedStoresOrItemsBlocked sets IndexingEncryptedStoresOrItemsBlocked field to given value.

### HasIndexingEncryptedStoresOrItemsBlocked

`func (o *WindowsInformationProtection) HasIndexingEncryptedStoresOrItemsBlocked() bool`

HasIndexingEncryptedStoresOrItemsBlocked returns a boolean if a field has been set.

### GetIsAssigned

`func (o *WindowsInformationProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *WindowsInformationProtection) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *WindowsInformationProtection) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *WindowsInformationProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetNeutralDomainResources

`func (o *WindowsInformationProtection) GetNeutralDomainResources() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetNeutralDomainResources returns the NeutralDomainResources field if non-nil, zero value otherwise.

### GetNeutralDomainResourcesOk

`func (o *WindowsInformationProtection) GetNeutralDomainResourcesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetNeutralDomainResourcesOk returns a tuple with the NeutralDomainResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralDomainResources

`func (o *WindowsInformationProtection) SetNeutralDomainResources(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetNeutralDomainResources sets NeutralDomainResources field to given value.

### HasNeutralDomainResources

`func (o *WindowsInformationProtection) HasNeutralDomainResources() bool`

HasNeutralDomainResources returns a boolean if a field has been set.

### GetProtectedApps

`func (o *WindowsInformationProtection) GetProtectedApps() []*AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetProtectedApps returns the ProtectedApps field if non-nil, zero value otherwise.

### GetProtectedAppsOk

`func (o *WindowsInformationProtection) GetProtectedAppsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetProtectedAppsOk returns a tuple with the ProtectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedApps

`func (o *WindowsInformationProtection) SetProtectedApps(v []*AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetProtectedApps sets ProtectedApps field to given value.

### HasProtectedApps

`func (o *WindowsInformationProtection) HasProtectedApps() bool`

HasProtectedApps returns a boolean if a field has been set.

### GetProtectionUnderLockConfigRequired

`func (o *WindowsInformationProtection) GetProtectionUnderLockConfigRequired() bool`

GetProtectionUnderLockConfigRequired returns the ProtectionUnderLockConfigRequired field if non-nil, zero value otherwise.

### GetProtectionUnderLockConfigRequiredOk

`func (o *WindowsInformationProtection) GetProtectionUnderLockConfigRequiredOk() (*bool, bool)`

GetProtectionUnderLockConfigRequiredOk returns a tuple with the ProtectionUnderLockConfigRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionUnderLockConfigRequired

`func (o *WindowsInformationProtection) SetProtectionUnderLockConfigRequired(v bool)`

SetProtectionUnderLockConfigRequired sets ProtectionUnderLockConfigRequired field to given value.

### HasProtectionUnderLockConfigRequired

`func (o *WindowsInformationProtection) HasProtectionUnderLockConfigRequired() bool`

HasProtectionUnderLockConfigRequired returns a boolean if a field has been set.

### GetRevokeOnUnenrollDisabled

`func (o *WindowsInformationProtection) GetRevokeOnUnenrollDisabled() bool`

GetRevokeOnUnenrollDisabled returns the RevokeOnUnenrollDisabled field if non-nil, zero value otherwise.

### GetRevokeOnUnenrollDisabledOk

`func (o *WindowsInformationProtection) GetRevokeOnUnenrollDisabledOk() (*bool, bool)`

GetRevokeOnUnenrollDisabledOk returns a tuple with the RevokeOnUnenrollDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOnUnenrollDisabled

`func (o *WindowsInformationProtection) SetRevokeOnUnenrollDisabled(v bool)`

SetRevokeOnUnenrollDisabled sets RevokeOnUnenrollDisabled field to given value.

### HasRevokeOnUnenrollDisabled

`func (o *WindowsInformationProtection) HasRevokeOnUnenrollDisabled() bool`

HasRevokeOnUnenrollDisabled returns a boolean if a field has been set.

### GetRightsManagementServicesTemplateId

`func (o *WindowsInformationProtection) GetRightsManagementServicesTemplateId() string`

GetRightsManagementServicesTemplateId returns the RightsManagementServicesTemplateId field if non-nil, zero value otherwise.

### GetRightsManagementServicesTemplateIdOk

`func (o *WindowsInformationProtection) GetRightsManagementServicesTemplateIdOk() (*string, bool)`

GetRightsManagementServicesTemplateIdOk returns a tuple with the RightsManagementServicesTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightsManagementServicesTemplateId

`func (o *WindowsInformationProtection) SetRightsManagementServicesTemplateId(v string)`

SetRightsManagementServicesTemplateId sets RightsManagementServicesTemplateId field to given value.

### HasRightsManagementServicesTemplateId

`func (o *WindowsInformationProtection) HasRightsManagementServicesTemplateId() bool`

HasRightsManagementServicesTemplateId returns a boolean if a field has been set.

### SetRightsManagementServicesTemplateIdNil

`func (o *WindowsInformationProtection) SetRightsManagementServicesTemplateIdNil(b bool)`

 SetRightsManagementServicesTemplateIdNil sets the value for RightsManagementServicesTemplateId to be an explicit nil

### UnsetRightsManagementServicesTemplateId
`func (o *WindowsInformationProtection) UnsetRightsManagementServicesTemplateId()`

UnsetRightsManagementServicesTemplateId ensures that no value is present for RightsManagementServicesTemplateId, not even an explicit nil
### GetSmbAutoEncryptedFileExtensions

`func (o *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensions() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetSmbAutoEncryptedFileExtensions returns the SmbAutoEncryptedFileExtensions field if non-nil, zero value otherwise.

### GetSmbAutoEncryptedFileExtensionsOk

`func (o *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensionsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetSmbAutoEncryptedFileExtensionsOk returns a tuple with the SmbAutoEncryptedFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAutoEncryptedFileExtensions

`func (o *WindowsInformationProtection) SetSmbAutoEncryptedFileExtensions(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetSmbAutoEncryptedFileExtensions sets SmbAutoEncryptedFileExtensions field to given value.

### HasSmbAutoEncryptedFileExtensions

`func (o *WindowsInformationProtection) HasSmbAutoEncryptedFileExtensions() bool`

HasSmbAutoEncryptedFileExtensions returns a boolean if a field has been set.

### GetAssignments

`func (o *WindowsInformationProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *WindowsInformationProtection) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *WindowsInformationProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *WindowsInformationProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetExemptAppLockerFiles

`func (o *WindowsInformationProtection) GetExemptAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetExemptAppLockerFiles returns the ExemptAppLockerFiles field if non-nil, zero value otherwise.

### GetExemptAppLockerFilesOk

`func (o *WindowsInformationProtection) GetExemptAppLockerFilesOk() (*[]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetExemptAppLockerFilesOk returns a tuple with the ExemptAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAppLockerFiles

`func (o *WindowsInformationProtection) SetExemptAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetExemptAppLockerFiles sets ExemptAppLockerFiles field to given value.

### HasExemptAppLockerFiles

`func (o *WindowsInformationProtection) HasExemptAppLockerFiles() bool`

HasExemptAppLockerFiles returns a boolean if a field has been set.

### GetProtectedAppLockerFiles

`func (o *WindowsInformationProtection) GetProtectedAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetProtectedAppLockerFiles returns the ProtectedAppLockerFiles field if non-nil, zero value otherwise.

### GetProtectedAppLockerFilesOk

`func (o *WindowsInformationProtection) GetProtectedAppLockerFilesOk() (*[]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetProtectedAppLockerFilesOk returns a tuple with the ProtectedAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppLockerFiles

`func (o *WindowsInformationProtection) SetProtectedAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetProtectedAppLockerFiles sets ProtectedAppLockerFiles field to given value.

### HasProtectedAppLockerFiles

`func (o *WindowsInformationProtection) HasProtectedAppLockerFiles() bool`

HasProtectedAppLockerFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


