# MicrosoftGraphMdmWindowsInformationProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time the policy was created. | [optional] 
**Description** | Pointer to **NullableString** | The policy&#39;s description. | [optional] 
**DisplayName** | Pointer to **string** | Policy display name. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Last time the policy was modified. | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 
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

### NewMicrosoftGraphMdmWindowsInformationProtectionPolicy

`func NewMicrosoftGraphMdmWindowsInformationProtectionPolicy() *MicrosoftGraphMdmWindowsInformationProtectionPolicy`

NewMicrosoftGraphMdmWindowsInformationProtectionPolicy instantiates a new MicrosoftGraphMdmWindowsInformationProtectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMdmWindowsInformationProtectionPolicyWithDefaults

`func NewMicrosoftGraphMdmWindowsInformationProtectionPolicyWithDefaults() *MicrosoftGraphMdmWindowsInformationProtectionPolicy`

NewMicrosoftGraphMdmWindowsInformationProtectionPolicyWithDefaults instantiates a new MicrosoftGraphMdmWindowsInformationProtectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAzureRightsManagementServicesAllowed() bool`

GetAzureRightsManagementServicesAllowed returns the AzureRightsManagementServicesAllowed field if non-nil, zero value otherwise.

### GetAzureRightsManagementServicesAllowedOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAzureRightsManagementServicesAllowedOk() (*bool, bool)`

GetAzureRightsManagementServicesAllowedOk returns a tuple with the AzureRightsManagementServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetAzureRightsManagementServicesAllowed(v bool)`

SetAzureRightsManagementServicesAllowed sets AzureRightsManagementServicesAllowed field to given value.

### HasAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasAzureRightsManagementServicesAllowed() bool`

HasAzureRightsManagementServicesAllowed returns a boolean if a field has been set.

### GetDataRecoveryCertificate

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDataRecoveryCertificate() AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate`

GetDataRecoveryCertificate returns the DataRecoveryCertificate field if non-nil, zero value otherwise.

### GetDataRecoveryCertificateOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDataRecoveryCertificateOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate, bool)`

GetDataRecoveryCertificateOk returns a tuple with the DataRecoveryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRecoveryCertificate

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDataRecoveryCertificate(v AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate)`

SetDataRecoveryCertificate sets DataRecoveryCertificate field to given value.

### HasDataRecoveryCertificate

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasDataRecoveryCertificate() bool`

HasDataRecoveryCertificate returns a boolean if a field has been set.

### SetDataRecoveryCertificateNil

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDataRecoveryCertificateNil(b bool)`

 SetDataRecoveryCertificateNil sets the value for DataRecoveryCertificate to be an explicit nil

### UnsetDataRecoveryCertificate
`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) UnsetDataRecoveryCertificate()`

UnsetDataRecoveryCertificate ensures that no value is present for DataRecoveryCertificate, not even an explicit nil
### GetEnforcementLevel

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnforcementLevel() AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnforcementLevelOk() (*AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementLevel

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnforcementLevel(v AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel)`

SetEnforcementLevel sets EnforcementLevel field to given value.

### HasEnforcementLevel

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### SetEnforcementLevelNil

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnforcementLevelNil(b bool)`

 SetEnforcementLevelNil sets the value for EnforcementLevel to be an explicit nil

### UnsetEnforcementLevel
`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) UnsetEnforcementLevel()`

UnsetEnforcementLevel ensures that no value is present for EnforcementLevel, not even an explicit nil
### GetEnterpriseDomain

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseDomain() string`

GetEnterpriseDomain returns the EnterpriseDomain field if non-nil, zero value otherwise.

### GetEnterpriseDomainOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseDomainOk() (*string, bool)`

GetEnterpriseDomainOk returns a tuple with the EnterpriseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseDomain

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseDomain(v string)`

SetEnterpriseDomain sets EnterpriseDomain field to given value.

### HasEnterpriseDomain

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseDomain() bool`

HasEnterpriseDomain returns a boolean if a field has been set.

### SetEnterpriseDomainNil

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseDomainNil(b bool)`

 SetEnterpriseDomainNil sets the value for EnterpriseDomain to be an explicit nil

### UnsetEnterpriseDomain
`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) UnsetEnterpriseDomain()`

UnsetEnterpriseDomain ensures that no value is present for EnterpriseDomain, not even an explicit nil
### GetEnterpriseInternalProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseInternalProxyServers() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseInternalProxyServers returns the EnterpriseInternalProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseInternalProxyServersOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseInternalProxyServersOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseInternalProxyServersOk returns a tuple with the EnterpriseInternalProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseInternalProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseInternalProxyServers(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseInternalProxyServers sets EnterpriseInternalProxyServers field to given value.

### HasEnterpriseInternalProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseInternalProxyServers() bool`

HasEnterpriseInternalProxyServers returns a boolean if a field has been set.

### GetEnterpriseIPRanges

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRanges() []*AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection`

GetEnterpriseIPRanges returns the EnterpriseIPRanges field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRangesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection, bool)`

GetEnterpriseIPRangesOk returns a tuple with the EnterpriseIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseIPRanges

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseIPRanges(v []*AnyOfmicrosoftGraphWindowsInformationProtectionIPRangeCollection)`

SetEnterpriseIPRanges sets EnterpriseIPRanges field to given value.

### HasEnterpriseIPRanges

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseIPRanges() bool`

HasEnterpriseIPRanges returns a boolean if a field has been set.

### GetEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRangesAreAuthoritative() bool`

GetEnterpriseIPRangesAreAuthoritative returns the EnterpriseIPRangesAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesAreAuthoritativeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRangesAreAuthoritativeOk() (*bool, bool)`

GetEnterpriseIPRangesAreAuthoritativeOk returns a tuple with the EnterpriseIPRangesAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseIPRangesAreAuthoritative(v bool)`

SetEnterpriseIPRangesAreAuthoritative sets EnterpriseIPRangesAreAuthoritative field to given value.

### HasEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseIPRangesAreAuthoritative() bool`

HasEnterpriseIPRangesAreAuthoritative returns a boolean if a field has been set.

### GetEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseNetworkDomainNames() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseNetworkDomainNames returns the EnterpriseNetworkDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseNetworkDomainNamesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseNetworkDomainNamesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseNetworkDomainNamesOk returns a tuple with the EnterpriseNetworkDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseNetworkDomainNames(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseNetworkDomainNames sets EnterpriseNetworkDomainNames field to given value.

### HasEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseNetworkDomainNames() bool`

HasEnterpriseNetworkDomainNames returns a boolean if a field has been set.

### GetEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProtectedDomainNames() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProtectedDomainNames returns the EnterpriseProtectedDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseProtectedDomainNamesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProtectedDomainNamesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProtectedDomainNamesOk returns a tuple with the EnterpriseProtectedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProtectedDomainNames(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProtectedDomainNames sets EnterpriseProtectedDomainNames field to given value.

### HasEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProtectedDomainNames() bool`

HasEnterpriseProtectedDomainNames returns a boolean if a field has been set.

### GetEnterpriseProxiedDomains

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxiedDomains() []*AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection`

GetEnterpriseProxiedDomains returns the EnterpriseProxiedDomains field if non-nil, zero value otherwise.

### GetEnterpriseProxiedDomainsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxiedDomainsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection, bool)`

GetEnterpriseProxiedDomainsOk returns a tuple with the EnterpriseProxiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProxiedDomains

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProxiedDomains(v []*AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection)`

SetEnterpriseProxiedDomains sets EnterpriseProxiedDomains field to given value.

### HasEnterpriseProxiedDomains

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProxiedDomains() bool`

HasEnterpriseProxiedDomains returns a boolean if a field has been set.

### GetEnterpriseProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServers() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProxyServers returns the EnterpriseProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServersOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProxyServersOk returns a tuple with the EnterpriseProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProxyServers(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProxyServers sets EnterpriseProxyServers field to given value.

### HasEnterpriseProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProxyServers() bool`

HasEnterpriseProxyServers returns a boolean if a field has been set.

### GetEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServersAreAuthoritative() bool`

GetEnterpriseProxyServersAreAuthoritative returns the EnterpriseProxyServersAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersAreAuthoritativeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServersAreAuthoritativeOk() (*bool, bool)`

GetEnterpriseProxyServersAreAuthoritativeOk returns a tuple with the EnterpriseProxyServersAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProxyServersAreAuthoritative(v bool)`

SetEnterpriseProxyServersAreAuthoritative sets EnterpriseProxyServersAreAuthoritative field to given value.

### HasEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProxyServersAreAuthoritative() bool`

HasEnterpriseProxyServersAreAuthoritative returns a boolean if a field has been set.

### GetExemptApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptApps() []*AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetExemptApps returns the ExemptApps field if non-nil, zero value otherwise.

### GetExemptAppsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptAppsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetExemptAppsOk returns a tuple with the ExemptApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetExemptApps(v []*AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetExemptApps sets ExemptApps field to given value.

### HasExemptApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasExemptApps() bool`

HasExemptApps returns a boolean if a field has been set.

### GetIconsVisible

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIconsVisible() bool`

GetIconsVisible returns the IconsVisible field if non-nil, zero value otherwise.

### GetIconsVisibleOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIconsVisibleOk() (*bool, bool)`

GetIconsVisibleOk returns a tuple with the IconsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconsVisible

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetIconsVisible(v bool)`

SetIconsVisible sets IconsVisible field to given value.

### HasIconsVisible

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasIconsVisible() bool`

HasIconsVisible returns a boolean if a field has been set.

### GetIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIndexingEncryptedStoresOrItemsBlocked() bool`

GetIndexingEncryptedStoresOrItemsBlocked returns the IndexingEncryptedStoresOrItemsBlocked field if non-nil, zero value otherwise.

### GetIndexingEncryptedStoresOrItemsBlockedOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIndexingEncryptedStoresOrItemsBlockedOk() (*bool, bool)`

GetIndexingEncryptedStoresOrItemsBlockedOk returns a tuple with the IndexingEncryptedStoresOrItemsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetIndexingEncryptedStoresOrItemsBlocked(v bool)`

SetIndexingEncryptedStoresOrItemsBlocked sets IndexingEncryptedStoresOrItemsBlocked field to given value.

### HasIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasIndexingEncryptedStoresOrItemsBlocked() bool`

HasIndexingEncryptedStoresOrItemsBlocked returns a boolean if a field has been set.

### GetIsAssigned

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIsAssignedOk() (*bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigned

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetIsAssigned(v bool)`

SetIsAssigned sets IsAssigned field to given value.

### HasIsAssigned

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### GetNeutralDomainResources

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetNeutralDomainResources() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetNeutralDomainResources returns the NeutralDomainResources field if non-nil, zero value otherwise.

### GetNeutralDomainResourcesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetNeutralDomainResourcesOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetNeutralDomainResourcesOk returns a tuple with the NeutralDomainResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralDomainResources

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetNeutralDomainResources(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetNeutralDomainResources sets NeutralDomainResources field to given value.

### HasNeutralDomainResources

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasNeutralDomainResources() bool`

HasNeutralDomainResources returns a boolean if a field has been set.

### GetProtectedApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedApps() []*AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetProtectedApps returns the ProtectedApps field if non-nil, zero value otherwise.

### GetProtectedAppsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedAppsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetProtectedAppsOk returns a tuple with the ProtectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetProtectedApps(v []*AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetProtectedApps sets ProtectedApps field to given value.

### HasProtectedApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasProtectedApps() bool`

HasProtectedApps returns a boolean if a field has been set.

### GetProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectionUnderLockConfigRequired() bool`

GetProtectionUnderLockConfigRequired returns the ProtectionUnderLockConfigRequired field if non-nil, zero value otherwise.

### GetProtectionUnderLockConfigRequiredOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectionUnderLockConfigRequiredOk() (*bool, bool)`

GetProtectionUnderLockConfigRequiredOk returns a tuple with the ProtectionUnderLockConfigRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetProtectionUnderLockConfigRequired(v bool)`

SetProtectionUnderLockConfigRequired sets ProtectionUnderLockConfigRequired field to given value.

### HasProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasProtectionUnderLockConfigRequired() bool`

HasProtectionUnderLockConfigRequired returns a boolean if a field has been set.

### GetRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRevokeOnUnenrollDisabled() bool`

GetRevokeOnUnenrollDisabled returns the RevokeOnUnenrollDisabled field if non-nil, zero value otherwise.

### GetRevokeOnUnenrollDisabledOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRevokeOnUnenrollDisabledOk() (*bool, bool)`

GetRevokeOnUnenrollDisabledOk returns a tuple with the RevokeOnUnenrollDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetRevokeOnUnenrollDisabled(v bool)`

SetRevokeOnUnenrollDisabled sets RevokeOnUnenrollDisabled field to given value.

### HasRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasRevokeOnUnenrollDisabled() bool`

HasRevokeOnUnenrollDisabled returns a boolean if a field has been set.

### GetRightsManagementServicesTemplateId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRightsManagementServicesTemplateId() string`

GetRightsManagementServicesTemplateId returns the RightsManagementServicesTemplateId field if non-nil, zero value otherwise.

### GetRightsManagementServicesTemplateIdOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRightsManagementServicesTemplateIdOk() (*string, bool)`

GetRightsManagementServicesTemplateIdOk returns a tuple with the RightsManagementServicesTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightsManagementServicesTemplateId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetRightsManagementServicesTemplateId(v string)`

SetRightsManagementServicesTemplateId sets RightsManagementServicesTemplateId field to given value.

### HasRightsManagementServicesTemplateId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasRightsManagementServicesTemplateId() bool`

HasRightsManagementServicesTemplateId returns a boolean if a field has been set.

### SetRightsManagementServicesTemplateIdNil

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetRightsManagementServicesTemplateIdNil(b bool)`

 SetRightsManagementServicesTemplateIdNil sets the value for RightsManagementServicesTemplateId to be an explicit nil

### UnsetRightsManagementServicesTemplateId
`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) UnsetRightsManagementServicesTemplateId()`

UnsetRightsManagementServicesTemplateId ensures that no value is present for RightsManagementServicesTemplateId, not even an explicit nil
### GetSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetSmbAutoEncryptedFileExtensions() []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetSmbAutoEncryptedFileExtensions returns the SmbAutoEncryptedFileExtensions field if non-nil, zero value otherwise.

### GetSmbAutoEncryptedFileExtensionsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetSmbAutoEncryptedFileExtensionsOk() (*[]*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetSmbAutoEncryptedFileExtensionsOk returns a tuple with the SmbAutoEncryptedFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetSmbAutoEncryptedFileExtensions(v []*AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetSmbAutoEncryptedFileExtensions sets SmbAutoEncryptedFileExtensions field to given value.

### HasSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasSmbAutoEncryptedFileExtensions() bool`

HasSmbAutoEncryptedFileExtensions returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAssignmentsOk() (*[]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetExemptAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetExemptAppLockerFiles returns the ExemptAppLockerFiles field if non-nil, zero value otherwise.

### GetExemptAppLockerFilesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptAppLockerFilesOk() (*[]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetExemptAppLockerFilesOk returns a tuple with the ExemptAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetExemptAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetExemptAppLockerFiles sets ExemptAppLockerFiles field to given value.

### HasExemptAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasExemptAppLockerFiles() bool`

HasExemptAppLockerFiles returns a boolean if a field has been set.

### GetProtectedAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetProtectedAppLockerFiles returns the ProtectedAppLockerFiles field if non-nil, zero value otherwise.

### GetProtectedAppLockerFilesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedAppLockerFilesOk() (*[]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetProtectedAppLockerFilesOk returns a tuple with the ProtectedAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetProtectedAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetProtectedAppLockerFiles sets ProtectedAppLockerFiles field to given value.

### HasProtectedAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasProtectedAppLockerFiles() bool`

HasProtectedAppLockerFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


