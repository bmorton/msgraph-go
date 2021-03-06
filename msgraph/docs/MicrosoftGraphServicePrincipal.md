# MicrosoftGraphServicePrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**AccountEnabled** | Pointer to **NullableBool** | true if the service principal account is enabled; otherwise, false. Supports $filter (eq, ne, not, in). | [optional] 
**AddIns** | Pointer to [**[]MicrosoftGraphAddIn**](MicrosoftGraphAddIn.md) | Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its &#39;FileHandler&#39; functionality. This will let services like Microsoft 365 call the application in the context of a document the user is working on. | [optional] 
**AlternativeNames** | Pointer to **[]string** | Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**AppDescription** | Pointer to **NullableString** | The description exposed by the associated application. | [optional] 
**AppDisplayName** | Pointer to **NullableString** | The display name exposed by the associated application. | [optional] 
**AppId** | Pointer to **NullableString** | The unique identifier for the associated application (its appId property). Supports $filter (eq, ne, not, in, startsWith). | [optional] 
**ApplicationTemplateId** | Pointer to **NullableString** | Unique identifier of the applicationTemplate that the servicePrincipal was created from. Read-only. Supports $filter (eq, ne, NOT, startsWith). | [optional] 
**AppOwnerOrganizationId** | Pointer to **NullableString** | Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications. Supports $filter (eq, ne, NOT, ge, le). | [optional] 
**AppRoleAssignmentRequired** | Pointer to **bool** | Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter (eq, ne, NOT). | [optional] 
**AppRoles** | Pointer to [**[]MicrosoftGraphAppRole**](MicrosoftGraphAppRole.md) | The roles exposed by the application which this service principal represents. For more information see the appRoles property definition on the application entity. Not nullable. | [optional] 
**Description** | Pointer to **NullableString** | Free text field to provide an internal end-user facing description of the service principal. End-user portals such MyApps will display the application description in this field. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search. | [optional] 
**DisabledByMicrosoftStatus** | Pointer to **NullableString** | Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not). | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**Homepage** | Pointer to **NullableString** | Home page or landing page of the application. | [optional] 
**Info** | Pointer to [**AnyOfmicrosoftGraphInformationalUrl**](anyOf&lt;microsoft.graph.informationalUrl&gt;.md) | Basic profile information of the acquired application such as app&#39;s marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values). | [optional] 
**KeyCredentials** | Pointer to [**[]MicrosoftGraphKeyCredential**](MicrosoftGraphKeyCredential.md) | The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge, le). | [optional] 
**LoginUrl** | Pointer to **NullableString** | Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The user launches the application from Microsoft 365, the Azure AD My Apps, or the Azure AD SSO URL. | [optional] 
**LogoutUrl** | Pointer to **NullableString** | Specifies the URL that will be used by Microsoft&#39;s authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols. | [optional] 
**Notes** | Pointer to **NullableString** | Free text field to capture information about the service principal, typically used for operational purposes. Maximum allowed size is 1024 characters. | [optional] 
**NotificationEmailAddresses** | Pointer to **[]string** | Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications. | [optional] 
**Oauth2PermissionScopes** | Pointer to [**[]MicrosoftGraphPermissionScope**](MicrosoftGraphPermissionScope.md) | The delegated permissions exposed by the application. For more information see the oauth2PermissionScopes property on the application entity&#39;s api property. Not nullable. | [optional] 
**PasswordCredentials** | Pointer to [**[]MicrosoftGraphPasswordCredential**](MicrosoftGraphPasswordCredential.md) | The collection of password credentials associated with the application. Not nullable. | [optional] 
**PreferredSingleSignOnMode** | Pointer to **NullableString** | Specifies the single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. The supported values are password, saml, notSupported, and oidc. | [optional] 
**PreferredTokenSigningKeyThumbprint** | Pointer to **NullableString** | Reserved for internal use only. Do not write or otherwise rely on this property. May be removed in future versions. | [optional] 
**ReplyUrls** | Pointer to **[]string** | The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable. | [optional] 
**SamlSingleSignOnSettings** | Pointer to [**AnyOfmicrosoftGraphSamlSingleSignOnSettings**](anyOf&lt;microsoft.graph.samlSingleSignOnSettings&gt;.md) | The collection for settings related to saml single sign-on. | [optional] 
**ServicePrincipalNames** | Pointer to **[]string** | Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD. For example,Client apps can specify a resource URI which is based on the values of this property to acquire an access token, which is the URI returned in the &#39;aud&#39; claim.The any operator is required for filter expressions on multi-valued properties. Not nullable.  Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**ServicePrincipalType** | Pointer to **NullableString** | Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally. The servicePrincipalType property can be set to three different values: __Application - A service principal that represents an application or service. The appId property identifies the associated app registration, and matches the appId of an application, possibly from a different tenant. If the associated app registration is missing, tokens are not issued for the service principal.__ManagedIdentity - A service principal that represents a managed identity. Service principals representing managed identities can be granted access and permissions, but cannot be updated or modified directly.__Legacy - A service principal that represents an app created before app registrations, or through legacy experiences. Legacy service principal can have credentials, service principal names, reply URLs, and other properties which are editable by an authorized user, but does not have an associated app registration. The appId value does not associate the service principal with an app registration. The service principal can only be used in the tenant where it was created.__SocialIdp - For internal use. | [optional] 
**SignInAudience** | Pointer to **NullableString** | Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values are:AzureADMyOrg: Users with a Microsoft work or school account in my organization???s Azure AD tenant (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization???s Azure AD tenant (multi-tenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or school account in any organization???s Azure AD tenant.PersonalMicrosoftAccount: Users with a personal Microsoft account only. | [optional] 
**Tags** | Pointer to **[]string** | Custom strings that can be used to categorize and identify the service principal. Not nullable. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**TokenEncryptionKeyId** | Pointer to **NullableString** | Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD issues tokens for this application encrypted using the key specified by this property. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user. | [optional] 
**AppRoleAssignedTo** | Pointer to [**[]MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | App role assignments for this app or service, granted to users, groups, and other service principals. Supports $expand. | [optional] 
**AppRoleAssignments** | Pointer to [**[]MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | App role assignment for another app or service, granted to this service principal. Supports $expand. | [optional] 
**ClaimsMappingPolicies** | Pointer to [**[]MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md) | The claimsMappingPolicies assigned to this service principal. Supports $expand. | [optional] 
**CreatedObjects** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Directory objects created by this service principal. Read-only. Nullable. | [optional] 
**DelegatedPermissionClassifications** | Pointer to [**[]MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md) | The permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand. | [optional] 
**Endpoints** | Pointer to [**[]MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md) | Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint endpoints that other applications can discover and use in their experiences. | [optional] 
**HomeRealmDiscoveryPolicies** | Pointer to [**[]MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md) | The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand. | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand. | [optional] 
**Oauth2PermissionGrants** | Pointer to [**[]MicrosoftGraphOAuth2PermissionGrant**](MicrosoftGraphOAuth2PermissionGrant.md) | Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user. Read-only. Nullable. | [optional] 
**OwnedObjects** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand. | [optional] 
**Owners** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand. | [optional] 
**TokenIssuancePolicies** | Pointer to [**[]MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md) | The tokenIssuancePolicies assigned to this service principal. | [optional] 
**TokenLifetimePolicies** | Pointer to [**[]MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md) | The tokenLifetimePolicies assigned to this service principal. | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 

## Methods

### NewMicrosoftGraphServicePrincipal

`func NewMicrosoftGraphServicePrincipal() *MicrosoftGraphServicePrincipal`

NewMicrosoftGraphServicePrincipal instantiates a new MicrosoftGraphServicePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServicePrincipalWithDefaults

`func NewMicrosoftGraphServicePrincipalWithDefaults() *MicrosoftGraphServicePrincipal`

NewMicrosoftGraphServicePrincipalWithDefaults instantiates a new MicrosoftGraphServicePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphServicePrincipal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphServicePrincipal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphServicePrincipal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphServicePrincipal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphServicePrincipal) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphServicePrincipal) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphServicePrincipal) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphServicePrincipal) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphServicePrincipal) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphServicePrincipal) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetAccountEnabled

`func (o *MicrosoftGraphServicePrincipal) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *MicrosoftGraphServicePrincipal) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *MicrosoftGraphServicePrincipal) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *MicrosoftGraphServicePrincipal) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabledNil

`func (o *MicrosoftGraphServicePrincipal) SetAccountEnabledNil(b bool)`

 SetAccountEnabledNil sets the value for AccountEnabled to be an explicit nil

### UnsetAccountEnabled
`func (o *MicrosoftGraphServicePrincipal) UnsetAccountEnabled()`

UnsetAccountEnabled ensures that no value is present for AccountEnabled, not even an explicit nil
### GetAddIns

`func (o *MicrosoftGraphServicePrincipal) GetAddIns() []MicrosoftGraphAddIn`

GetAddIns returns the AddIns field if non-nil, zero value otherwise.

### GetAddInsOk

`func (o *MicrosoftGraphServicePrincipal) GetAddInsOk() (*[]MicrosoftGraphAddIn, bool)`

GetAddInsOk returns a tuple with the AddIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddIns

`func (o *MicrosoftGraphServicePrincipal) SetAddIns(v []MicrosoftGraphAddIn)`

SetAddIns sets AddIns field to given value.

### HasAddIns

`func (o *MicrosoftGraphServicePrincipal) HasAddIns() bool`

HasAddIns returns a boolean if a field has been set.

### GetAlternativeNames

`func (o *MicrosoftGraphServicePrincipal) GetAlternativeNames() []string`

GetAlternativeNames returns the AlternativeNames field if non-nil, zero value otherwise.

### GetAlternativeNamesOk

`func (o *MicrosoftGraphServicePrincipal) GetAlternativeNamesOk() (*[]string, bool)`

GetAlternativeNamesOk returns a tuple with the AlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeNames

`func (o *MicrosoftGraphServicePrincipal) SetAlternativeNames(v []string)`

SetAlternativeNames sets AlternativeNames field to given value.

### HasAlternativeNames

`func (o *MicrosoftGraphServicePrincipal) HasAlternativeNames() bool`

HasAlternativeNames returns a boolean if a field has been set.

### GetAppDescription

`func (o *MicrosoftGraphServicePrincipal) GetAppDescription() string`

GetAppDescription returns the AppDescription field if non-nil, zero value otherwise.

### GetAppDescriptionOk

`func (o *MicrosoftGraphServicePrincipal) GetAppDescriptionOk() (*string, bool)`

GetAppDescriptionOk returns a tuple with the AppDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescription

`func (o *MicrosoftGraphServicePrincipal) SetAppDescription(v string)`

SetAppDescription sets AppDescription field to given value.

### HasAppDescription

`func (o *MicrosoftGraphServicePrincipal) HasAppDescription() bool`

HasAppDescription returns a boolean if a field has been set.

### SetAppDescriptionNil

`func (o *MicrosoftGraphServicePrincipal) SetAppDescriptionNil(b bool)`

 SetAppDescriptionNil sets the value for AppDescription to be an explicit nil

### UnsetAppDescription
`func (o *MicrosoftGraphServicePrincipal) UnsetAppDescription()`

UnsetAppDescription ensures that no value is present for AppDescription, not even an explicit nil
### GetAppDisplayName

`func (o *MicrosoftGraphServicePrincipal) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphServicePrincipal) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphServicePrincipal) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *MicrosoftGraphServicePrincipal) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *MicrosoftGraphServicePrincipal) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *MicrosoftGraphServicePrincipal) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetAppId

`func (o *MicrosoftGraphServicePrincipal) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphServicePrincipal) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphServicePrincipal) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphServicePrincipal) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *MicrosoftGraphServicePrincipal) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *MicrosoftGraphServicePrincipal) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetApplicationTemplateId

`func (o *MicrosoftGraphServicePrincipal) GetApplicationTemplateId() string`

GetApplicationTemplateId returns the ApplicationTemplateId field if non-nil, zero value otherwise.

### GetApplicationTemplateIdOk

`func (o *MicrosoftGraphServicePrincipal) GetApplicationTemplateIdOk() (*string, bool)`

GetApplicationTemplateIdOk returns a tuple with the ApplicationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTemplateId

`func (o *MicrosoftGraphServicePrincipal) SetApplicationTemplateId(v string)`

SetApplicationTemplateId sets ApplicationTemplateId field to given value.

### HasApplicationTemplateId

`func (o *MicrosoftGraphServicePrincipal) HasApplicationTemplateId() bool`

HasApplicationTemplateId returns a boolean if a field has been set.

### SetApplicationTemplateIdNil

`func (o *MicrosoftGraphServicePrincipal) SetApplicationTemplateIdNil(b bool)`

 SetApplicationTemplateIdNil sets the value for ApplicationTemplateId to be an explicit nil

### UnsetApplicationTemplateId
`func (o *MicrosoftGraphServicePrincipal) UnsetApplicationTemplateId()`

UnsetApplicationTemplateId ensures that no value is present for ApplicationTemplateId, not even an explicit nil
### GetAppOwnerOrganizationId

`func (o *MicrosoftGraphServicePrincipal) GetAppOwnerOrganizationId() string`

GetAppOwnerOrganizationId returns the AppOwnerOrganizationId field if non-nil, zero value otherwise.

### GetAppOwnerOrganizationIdOk

`func (o *MicrosoftGraphServicePrincipal) GetAppOwnerOrganizationIdOk() (*string, bool)`

GetAppOwnerOrganizationIdOk returns a tuple with the AppOwnerOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOwnerOrganizationId

`func (o *MicrosoftGraphServicePrincipal) SetAppOwnerOrganizationId(v string)`

SetAppOwnerOrganizationId sets AppOwnerOrganizationId field to given value.

### HasAppOwnerOrganizationId

`func (o *MicrosoftGraphServicePrincipal) HasAppOwnerOrganizationId() bool`

HasAppOwnerOrganizationId returns a boolean if a field has been set.

### SetAppOwnerOrganizationIdNil

`func (o *MicrosoftGraphServicePrincipal) SetAppOwnerOrganizationIdNil(b bool)`

 SetAppOwnerOrganizationIdNil sets the value for AppOwnerOrganizationId to be an explicit nil

### UnsetAppOwnerOrganizationId
`func (o *MicrosoftGraphServicePrincipal) UnsetAppOwnerOrganizationId()`

UnsetAppOwnerOrganizationId ensures that no value is present for AppOwnerOrganizationId, not even an explicit nil
### GetAppRoleAssignmentRequired

`func (o *MicrosoftGraphServicePrincipal) GetAppRoleAssignmentRequired() bool`

GetAppRoleAssignmentRequired returns the AppRoleAssignmentRequired field if non-nil, zero value otherwise.

### GetAppRoleAssignmentRequiredOk

`func (o *MicrosoftGraphServicePrincipal) GetAppRoleAssignmentRequiredOk() (*bool, bool)`

GetAppRoleAssignmentRequiredOk returns a tuple with the AppRoleAssignmentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignmentRequired

`func (o *MicrosoftGraphServicePrincipal) SetAppRoleAssignmentRequired(v bool)`

SetAppRoleAssignmentRequired sets AppRoleAssignmentRequired field to given value.

### HasAppRoleAssignmentRequired

`func (o *MicrosoftGraphServicePrincipal) HasAppRoleAssignmentRequired() bool`

HasAppRoleAssignmentRequired returns a boolean if a field has been set.

### GetAppRoles

`func (o *MicrosoftGraphServicePrincipal) GetAppRoles() []MicrosoftGraphAppRole`

GetAppRoles returns the AppRoles field if non-nil, zero value otherwise.

### GetAppRolesOk

`func (o *MicrosoftGraphServicePrincipal) GetAppRolesOk() (*[]MicrosoftGraphAppRole, bool)`

GetAppRolesOk returns a tuple with the AppRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoles

`func (o *MicrosoftGraphServicePrincipal) SetAppRoles(v []MicrosoftGraphAppRole)`

SetAppRoles sets AppRoles field to given value.

### HasAppRoles

`func (o *MicrosoftGraphServicePrincipal) HasAppRoles() bool`

HasAppRoles returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphServicePrincipal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphServicePrincipal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphServicePrincipal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphServicePrincipal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphServicePrincipal) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphServicePrincipal) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabledByMicrosoftStatus

`func (o *MicrosoftGraphServicePrincipal) GetDisabledByMicrosoftStatus() string`

GetDisabledByMicrosoftStatus returns the DisabledByMicrosoftStatus field if non-nil, zero value otherwise.

### GetDisabledByMicrosoftStatusOk

`func (o *MicrosoftGraphServicePrincipal) GetDisabledByMicrosoftStatusOk() (*string, bool)`

GetDisabledByMicrosoftStatusOk returns a tuple with the DisabledByMicrosoftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledByMicrosoftStatus

`func (o *MicrosoftGraphServicePrincipal) SetDisabledByMicrosoftStatus(v string)`

SetDisabledByMicrosoftStatus sets DisabledByMicrosoftStatus field to given value.

### HasDisabledByMicrosoftStatus

`func (o *MicrosoftGraphServicePrincipal) HasDisabledByMicrosoftStatus() bool`

HasDisabledByMicrosoftStatus returns a boolean if a field has been set.

### SetDisabledByMicrosoftStatusNil

`func (o *MicrosoftGraphServicePrincipal) SetDisabledByMicrosoftStatusNil(b bool)`

 SetDisabledByMicrosoftStatusNil sets the value for DisabledByMicrosoftStatus to be an explicit nil

### UnsetDisabledByMicrosoftStatus
`func (o *MicrosoftGraphServicePrincipal) UnsetDisabledByMicrosoftStatus()`

UnsetDisabledByMicrosoftStatus ensures that no value is present for DisabledByMicrosoftStatus, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphServicePrincipal) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphServicePrincipal) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphServicePrincipal) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphServicePrincipal) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphServicePrincipal) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphServicePrincipal) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetHomepage

`func (o *MicrosoftGraphServicePrincipal) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *MicrosoftGraphServicePrincipal) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *MicrosoftGraphServicePrincipal) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *MicrosoftGraphServicePrincipal) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### SetHomepageNil

`func (o *MicrosoftGraphServicePrincipal) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *MicrosoftGraphServicePrincipal) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil
### GetInfo

`func (o *MicrosoftGraphServicePrincipal) GetInfo() AnyOfmicrosoftGraphInformationalUrl`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MicrosoftGraphServicePrincipal) GetInfoOk() (*AnyOfmicrosoftGraphInformationalUrl, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MicrosoftGraphServicePrincipal) SetInfo(v AnyOfmicrosoftGraphInformationalUrl)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MicrosoftGraphServicePrincipal) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *MicrosoftGraphServicePrincipal) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *MicrosoftGraphServicePrincipal) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetKeyCredentials

`func (o *MicrosoftGraphServicePrincipal) GetKeyCredentials() []MicrosoftGraphKeyCredential`

GetKeyCredentials returns the KeyCredentials field if non-nil, zero value otherwise.

### GetKeyCredentialsOk

`func (o *MicrosoftGraphServicePrincipal) GetKeyCredentialsOk() (*[]MicrosoftGraphKeyCredential, bool)`

GetKeyCredentialsOk returns a tuple with the KeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCredentials

`func (o *MicrosoftGraphServicePrincipal) SetKeyCredentials(v []MicrosoftGraphKeyCredential)`

SetKeyCredentials sets KeyCredentials field to given value.

### HasKeyCredentials

`func (o *MicrosoftGraphServicePrincipal) HasKeyCredentials() bool`

HasKeyCredentials returns a boolean if a field has been set.

### GetLoginUrl

`func (o *MicrosoftGraphServicePrincipal) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *MicrosoftGraphServicePrincipal) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *MicrosoftGraphServicePrincipal) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *MicrosoftGraphServicePrincipal) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### SetLoginUrlNil

`func (o *MicrosoftGraphServicePrincipal) SetLoginUrlNil(b bool)`

 SetLoginUrlNil sets the value for LoginUrl to be an explicit nil

### UnsetLoginUrl
`func (o *MicrosoftGraphServicePrincipal) UnsetLoginUrl()`

UnsetLoginUrl ensures that no value is present for LoginUrl, not even an explicit nil
### GetLogoutUrl

`func (o *MicrosoftGraphServicePrincipal) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *MicrosoftGraphServicePrincipal) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *MicrosoftGraphServicePrincipal) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *MicrosoftGraphServicePrincipal) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### SetLogoutUrlNil

`func (o *MicrosoftGraphServicePrincipal) SetLogoutUrlNil(b bool)`

 SetLogoutUrlNil sets the value for LogoutUrl to be an explicit nil

### UnsetLogoutUrl
`func (o *MicrosoftGraphServicePrincipal) UnsetLogoutUrl()`

UnsetLogoutUrl ensures that no value is present for LogoutUrl, not even an explicit nil
### GetNotes

`func (o *MicrosoftGraphServicePrincipal) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphServicePrincipal) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MicrosoftGraphServicePrincipal) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MicrosoftGraphServicePrincipal) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *MicrosoftGraphServicePrincipal) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *MicrosoftGraphServicePrincipal) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetNotificationEmailAddresses

`func (o *MicrosoftGraphServicePrincipal) GetNotificationEmailAddresses() []string`

GetNotificationEmailAddresses returns the NotificationEmailAddresses field if non-nil, zero value otherwise.

### GetNotificationEmailAddressesOk

`func (o *MicrosoftGraphServicePrincipal) GetNotificationEmailAddressesOk() (*[]string, bool)`

GetNotificationEmailAddressesOk returns a tuple with the NotificationEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmailAddresses

`func (o *MicrosoftGraphServicePrincipal) SetNotificationEmailAddresses(v []string)`

SetNotificationEmailAddresses sets NotificationEmailAddresses field to given value.

### HasNotificationEmailAddresses

`func (o *MicrosoftGraphServicePrincipal) HasNotificationEmailAddresses() bool`

HasNotificationEmailAddresses returns a boolean if a field has been set.

### GetOauth2PermissionScopes

`func (o *MicrosoftGraphServicePrincipal) GetOauth2PermissionScopes() []MicrosoftGraphPermissionScope`

GetOauth2PermissionScopes returns the Oauth2PermissionScopes field if non-nil, zero value otherwise.

### GetOauth2PermissionScopesOk

`func (o *MicrosoftGraphServicePrincipal) GetOauth2PermissionScopesOk() (*[]MicrosoftGraphPermissionScope, bool)`

GetOauth2PermissionScopesOk returns a tuple with the Oauth2PermissionScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2PermissionScopes

`func (o *MicrosoftGraphServicePrincipal) SetOauth2PermissionScopes(v []MicrosoftGraphPermissionScope)`

SetOauth2PermissionScopes sets Oauth2PermissionScopes field to given value.

### HasOauth2PermissionScopes

`func (o *MicrosoftGraphServicePrincipal) HasOauth2PermissionScopes() bool`

HasOauth2PermissionScopes returns a boolean if a field has been set.

### GetPasswordCredentials

`func (o *MicrosoftGraphServicePrincipal) GetPasswordCredentials() []MicrosoftGraphPasswordCredential`

GetPasswordCredentials returns the PasswordCredentials field if non-nil, zero value otherwise.

### GetPasswordCredentialsOk

`func (o *MicrosoftGraphServicePrincipal) GetPasswordCredentialsOk() (*[]MicrosoftGraphPasswordCredential, bool)`

GetPasswordCredentialsOk returns a tuple with the PasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCredentials

`func (o *MicrosoftGraphServicePrincipal) SetPasswordCredentials(v []MicrosoftGraphPasswordCredential)`

SetPasswordCredentials sets PasswordCredentials field to given value.

### HasPasswordCredentials

`func (o *MicrosoftGraphServicePrincipal) HasPasswordCredentials() bool`

HasPasswordCredentials returns a boolean if a field has been set.

### GetPreferredSingleSignOnMode

`func (o *MicrosoftGraphServicePrincipal) GetPreferredSingleSignOnMode() string`

GetPreferredSingleSignOnMode returns the PreferredSingleSignOnMode field if non-nil, zero value otherwise.

### GetPreferredSingleSignOnModeOk

`func (o *MicrosoftGraphServicePrincipal) GetPreferredSingleSignOnModeOk() (*string, bool)`

GetPreferredSingleSignOnModeOk returns a tuple with the PreferredSingleSignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSingleSignOnMode

`func (o *MicrosoftGraphServicePrincipal) SetPreferredSingleSignOnMode(v string)`

SetPreferredSingleSignOnMode sets PreferredSingleSignOnMode field to given value.

### HasPreferredSingleSignOnMode

`func (o *MicrosoftGraphServicePrincipal) HasPreferredSingleSignOnMode() bool`

HasPreferredSingleSignOnMode returns a boolean if a field has been set.

### SetPreferredSingleSignOnModeNil

`func (o *MicrosoftGraphServicePrincipal) SetPreferredSingleSignOnModeNil(b bool)`

 SetPreferredSingleSignOnModeNil sets the value for PreferredSingleSignOnMode to be an explicit nil

### UnsetPreferredSingleSignOnMode
`func (o *MicrosoftGraphServicePrincipal) UnsetPreferredSingleSignOnMode()`

UnsetPreferredSingleSignOnMode ensures that no value is present for PreferredSingleSignOnMode, not even an explicit nil
### GetPreferredTokenSigningKeyThumbprint

`func (o *MicrosoftGraphServicePrincipal) GetPreferredTokenSigningKeyThumbprint() string`

GetPreferredTokenSigningKeyThumbprint returns the PreferredTokenSigningKeyThumbprint field if non-nil, zero value otherwise.

### GetPreferredTokenSigningKeyThumbprintOk

`func (o *MicrosoftGraphServicePrincipal) GetPreferredTokenSigningKeyThumbprintOk() (*string, bool)`

GetPreferredTokenSigningKeyThumbprintOk returns a tuple with the PreferredTokenSigningKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTokenSigningKeyThumbprint

`func (o *MicrosoftGraphServicePrincipal) SetPreferredTokenSigningKeyThumbprint(v string)`

SetPreferredTokenSigningKeyThumbprint sets PreferredTokenSigningKeyThumbprint field to given value.

### HasPreferredTokenSigningKeyThumbprint

`func (o *MicrosoftGraphServicePrincipal) HasPreferredTokenSigningKeyThumbprint() bool`

HasPreferredTokenSigningKeyThumbprint returns a boolean if a field has been set.

### SetPreferredTokenSigningKeyThumbprintNil

`func (o *MicrosoftGraphServicePrincipal) SetPreferredTokenSigningKeyThumbprintNil(b bool)`

 SetPreferredTokenSigningKeyThumbprintNil sets the value for PreferredTokenSigningKeyThumbprint to be an explicit nil

### UnsetPreferredTokenSigningKeyThumbprint
`func (o *MicrosoftGraphServicePrincipal) UnsetPreferredTokenSigningKeyThumbprint()`

UnsetPreferredTokenSigningKeyThumbprint ensures that no value is present for PreferredTokenSigningKeyThumbprint, not even an explicit nil
### GetReplyUrls

`func (o *MicrosoftGraphServicePrincipal) GetReplyUrls() []string`

GetReplyUrls returns the ReplyUrls field if non-nil, zero value otherwise.

### GetReplyUrlsOk

`func (o *MicrosoftGraphServicePrincipal) GetReplyUrlsOk() (*[]string, bool)`

GetReplyUrlsOk returns a tuple with the ReplyUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrls

`func (o *MicrosoftGraphServicePrincipal) SetReplyUrls(v []string)`

SetReplyUrls sets ReplyUrls field to given value.

### HasReplyUrls

`func (o *MicrosoftGraphServicePrincipal) HasReplyUrls() bool`

HasReplyUrls returns a boolean if a field has been set.

### GetSamlSingleSignOnSettings

`func (o *MicrosoftGraphServicePrincipal) GetSamlSingleSignOnSettings() AnyOfmicrosoftGraphSamlSingleSignOnSettings`

GetSamlSingleSignOnSettings returns the SamlSingleSignOnSettings field if non-nil, zero value otherwise.

### GetSamlSingleSignOnSettingsOk

`func (o *MicrosoftGraphServicePrincipal) GetSamlSingleSignOnSettingsOk() (*AnyOfmicrosoftGraphSamlSingleSignOnSettings, bool)`

GetSamlSingleSignOnSettingsOk returns a tuple with the SamlSingleSignOnSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSingleSignOnSettings

`func (o *MicrosoftGraphServicePrincipal) SetSamlSingleSignOnSettings(v AnyOfmicrosoftGraphSamlSingleSignOnSettings)`

SetSamlSingleSignOnSettings sets SamlSingleSignOnSettings field to given value.

### HasSamlSingleSignOnSettings

`func (o *MicrosoftGraphServicePrincipal) HasSamlSingleSignOnSettings() bool`

HasSamlSingleSignOnSettings returns a boolean if a field has been set.

### SetSamlSingleSignOnSettingsNil

`func (o *MicrosoftGraphServicePrincipal) SetSamlSingleSignOnSettingsNil(b bool)`

 SetSamlSingleSignOnSettingsNil sets the value for SamlSingleSignOnSettings to be an explicit nil

### UnsetSamlSingleSignOnSettings
`func (o *MicrosoftGraphServicePrincipal) UnsetSamlSingleSignOnSettings()`

UnsetSamlSingleSignOnSettings ensures that no value is present for SamlSingleSignOnSettings, not even an explicit nil
### GetServicePrincipalNames

`func (o *MicrosoftGraphServicePrincipal) GetServicePrincipalNames() []string`

GetServicePrincipalNames returns the ServicePrincipalNames field if non-nil, zero value otherwise.

### GetServicePrincipalNamesOk

`func (o *MicrosoftGraphServicePrincipal) GetServicePrincipalNamesOk() (*[]string, bool)`

GetServicePrincipalNamesOk returns a tuple with the ServicePrincipalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalNames

`func (o *MicrosoftGraphServicePrincipal) SetServicePrincipalNames(v []string)`

SetServicePrincipalNames sets ServicePrincipalNames field to given value.

### HasServicePrincipalNames

`func (o *MicrosoftGraphServicePrincipal) HasServicePrincipalNames() bool`

HasServicePrincipalNames returns a boolean if a field has been set.

### GetServicePrincipalType

`func (o *MicrosoftGraphServicePrincipal) GetServicePrincipalType() string`

GetServicePrincipalType returns the ServicePrincipalType field if non-nil, zero value otherwise.

### GetServicePrincipalTypeOk

`func (o *MicrosoftGraphServicePrincipal) GetServicePrincipalTypeOk() (*string, bool)`

GetServicePrincipalTypeOk returns a tuple with the ServicePrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalType

`func (o *MicrosoftGraphServicePrincipal) SetServicePrincipalType(v string)`

SetServicePrincipalType sets ServicePrincipalType field to given value.

### HasServicePrincipalType

`func (o *MicrosoftGraphServicePrincipal) HasServicePrincipalType() bool`

HasServicePrincipalType returns a boolean if a field has been set.

### SetServicePrincipalTypeNil

`func (o *MicrosoftGraphServicePrincipal) SetServicePrincipalTypeNil(b bool)`

 SetServicePrincipalTypeNil sets the value for ServicePrincipalType to be an explicit nil

### UnsetServicePrincipalType
`func (o *MicrosoftGraphServicePrincipal) UnsetServicePrincipalType()`

UnsetServicePrincipalType ensures that no value is present for ServicePrincipalType, not even an explicit nil
### GetSignInAudience

`func (o *MicrosoftGraphServicePrincipal) GetSignInAudience() string`

GetSignInAudience returns the SignInAudience field if non-nil, zero value otherwise.

### GetSignInAudienceOk

`func (o *MicrosoftGraphServicePrincipal) GetSignInAudienceOk() (*string, bool)`

GetSignInAudienceOk returns a tuple with the SignInAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInAudience

`func (o *MicrosoftGraphServicePrincipal) SetSignInAudience(v string)`

SetSignInAudience sets SignInAudience field to given value.

### HasSignInAudience

`func (o *MicrosoftGraphServicePrincipal) HasSignInAudience() bool`

HasSignInAudience returns a boolean if a field has been set.

### SetSignInAudienceNil

`func (o *MicrosoftGraphServicePrincipal) SetSignInAudienceNil(b bool)`

 SetSignInAudienceNil sets the value for SignInAudience to be an explicit nil

### UnsetSignInAudience
`func (o *MicrosoftGraphServicePrincipal) UnsetSignInAudience()`

UnsetSignInAudience ensures that no value is present for SignInAudience, not even an explicit nil
### GetTags

`func (o *MicrosoftGraphServicePrincipal) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MicrosoftGraphServicePrincipal) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MicrosoftGraphServicePrincipal) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MicrosoftGraphServicePrincipal) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTokenEncryptionKeyId

`func (o *MicrosoftGraphServicePrincipal) GetTokenEncryptionKeyId() string`

GetTokenEncryptionKeyId returns the TokenEncryptionKeyId field if non-nil, zero value otherwise.

### GetTokenEncryptionKeyIdOk

`func (o *MicrosoftGraphServicePrincipal) GetTokenEncryptionKeyIdOk() (*string, bool)`

GetTokenEncryptionKeyIdOk returns a tuple with the TokenEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEncryptionKeyId

`func (o *MicrosoftGraphServicePrincipal) SetTokenEncryptionKeyId(v string)`

SetTokenEncryptionKeyId sets TokenEncryptionKeyId field to given value.

### HasTokenEncryptionKeyId

`func (o *MicrosoftGraphServicePrincipal) HasTokenEncryptionKeyId() bool`

HasTokenEncryptionKeyId returns a boolean if a field has been set.

### SetTokenEncryptionKeyIdNil

`func (o *MicrosoftGraphServicePrincipal) SetTokenEncryptionKeyIdNil(b bool)`

 SetTokenEncryptionKeyIdNil sets the value for TokenEncryptionKeyId to be an explicit nil

### UnsetTokenEncryptionKeyId
`func (o *MicrosoftGraphServicePrincipal) UnsetTokenEncryptionKeyId()`

UnsetTokenEncryptionKeyId ensures that no value is present for TokenEncryptionKeyId, not even an explicit nil
### GetAppRoleAssignedTo

`func (o *MicrosoftGraphServicePrincipal) GetAppRoleAssignedTo() []MicrosoftGraphAppRoleAssignment`

GetAppRoleAssignedTo returns the AppRoleAssignedTo field if non-nil, zero value otherwise.

### GetAppRoleAssignedToOk

`func (o *MicrosoftGraphServicePrincipal) GetAppRoleAssignedToOk() (*[]MicrosoftGraphAppRoleAssignment, bool)`

GetAppRoleAssignedToOk returns a tuple with the AppRoleAssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignedTo

`func (o *MicrosoftGraphServicePrincipal) SetAppRoleAssignedTo(v []MicrosoftGraphAppRoleAssignment)`

SetAppRoleAssignedTo sets AppRoleAssignedTo field to given value.

### HasAppRoleAssignedTo

`func (o *MicrosoftGraphServicePrincipal) HasAppRoleAssignedTo() bool`

HasAppRoleAssignedTo returns a boolean if a field has been set.

### GetAppRoleAssignments

`func (o *MicrosoftGraphServicePrincipal) GetAppRoleAssignments() []MicrosoftGraphAppRoleAssignment`

GetAppRoleAssignments returns the AppRoleAssignments field if non-nil, zero value otherwise.

### GetAppRoleAssignmentsOk

`func (o *MicrosoftGraphServicePrincipal) GetAppRoleAssignmentsOk() (*[]MicrosoftGraphAppRoleAssignment, bool)`

GetAppRoleAssignmentsOk returns a tuple with the AppRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignments

`func (o *MicrosoftGraphServicePrincipal) SetAppRoleAssignments(v []MicrosoftGraphAppRoleAssignment)`

SetAppRoleAssignments sets AppRoleAssignments field to given value.

### HasAppRoleAssignments

`func (o *MicrosoftGraphServicePrincipal) HasAppRoleAssignments() bool`

HasAppRoleAssignments returns a boolean if a field has been set.

### GetClaimsMappingPolicies

`func (o *MicrosoftGraphServicePrincipal) GetClaimsMappingPolicies() []MicrosoftGraphClaimsMappingPolicy`

GetClaimsMappingPolicies returns the ClaimsMappingPolicies field if non-nil, zero value otherwise.

### GetClaimsMappingPoliciesOk

`func (o *MicrosoftGraphServicePrincipal) GetClaimsMappingPoliciesOk() (*[]MicrosoftGraphClaimsMappingPolicy, bool)`

GetClaimsMappingPoliciesOk returns a tuple with the ClaimsMappingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsMappingPolicies

`func (o *MicrosoftGraphServicePrincipal) SetClaimsMappingPolicies(v []MicrosoftGraphClaimsMappingPolicy)`

SetClaimsMappingPolicies sets ClaimsMappingPolicies field to given value.

### HasClaimsMappingPolicies

`func (o *MicrosoftGraphServicePrincipal) HasClaimsMappingPolicies() bool`

HasClaimsMappingPolicies returns a boolean if a field has been set.

### GetCreatedObjects

`func (o *MicrosoftGraphServicePrincipal) GetCreatedObjects() []MicrosoftGraphDirectoryObject`

GetCreatedObjects returns the CreatedObjects field if non-nil, zero value otherwise.

### GetCreatedObjectsOk

`func (o *MicrosoftGraphServicePrincipal) GetCreatedObjectsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetCreatedObjectsOk returns a tuple with the CreatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedObjects

`func (o *MicrosoftGraphServicePrincipal) SetCreatedObjects(v []MicrosoftGraphDirectoryObject)`

SetCreatedObjects sets CreatedObjects field to given value.

### HasCreatedObjects

`func (o *MicrosoftGraphServicePrincipal) HasCreatedObjects() bool`

HasCreatedObjects returns a boolean if a field has been set.

### GetDelegatedPermissionClassifications

`func (o *MicrosoftGraphServicePrincipal) GetDelegatedPermissionClassifications() []MicrosoftGraphDelegatedPermissionClassification`

GetDelegatedPermissionClassifications returns the DelegatedPermissionClassifications field if non-nil, zero value otherwise.

### GetDelegatedPermissionClassificationsOk

`func (o *MicrosoftGraphServicePrincipal) GetDelegatedPermissionClassificationsOk() (*[]MicrosoftGraphDelegatedPermissionClassification, bool)`

GetDelegatedPermissionClassificationsOk returns a tuple with the DelegatedPermissionClassifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedPermissionClassifications

`func (o *MicrosoftGraphServicePrincipal) SetDelegatedPermissionClassifications(v []MicrosoftGraphDelegatedPermissionClassification)`

SetDelegatedPermissionClassifications sets DelegatedPermissionClassifications field to given value.

### HasDelegatedPermissionClassifications

`func (o *MicrosoftGraphServicePrincipal) HasDelegatedPermissionClassifications() bool`

HasDelegatedPermissionClassifications returns a boolean if a field has been set.

### GetEndpoints

`func (o *MicrosoftGraphServicePrincipal) GetEndpoints() []MicrosoftGraphEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *MicrosoftGraphServicePrincipal) GetEndpointsOk() (*[]MicrosoftGraphEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *MicrosoftGraphServicePrincipal) SetEndpoints(v []MicrosoftGraphEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *MicrosoftGraphServicePrincipal) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetHomeRealmDiscoveryPolicies

`func (o *MicrosoftGraphServicePrincipal) GetHomeRealmDiscoveryPolicies() []MicrosoftGraphHomeRealmDiscoveryPolicy`

GetHomeRealmDiscoveryPolicies returns the HomeRealmDiscoveryPolicies field if non-nil, zero value otherwise.

### GetHomeRealmDiscoveryPoliciesOk

`func (o *MicrosoftGraphServicePrincipal) GetHomeRealmDiscoveryPoliciesOk() (*[]MicrosoftGraphHomeRealmDiscoveryPolicy, bool)`

GetHomeRealmDiscoveryPoliciesOk returns a tuple with the HomeRealmDiscoveryPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeRealmDiscoveryPolicies

`func (o *MicrosoftGraphServicePrincipal) SetHomeRealmDiscoveryPolicies(v []MicrosoftGraphHomeRealmDiscoveryPolicy)`

SetHomeRealmDiscoveryPolicies sets HomeRealmDiscoveryPolicies field to given value.

### HasHomeRealmDiscoveryPolicies

`func (o *MicrosoftGraphServicePrincipal) HasHomeRealmDiscoveryPolicies() bool`

HasHomeRealmDiscoveryPolicies returns a boolean if a field has been set.

### GetMemberOf

`func (o *MicrosoftGraphServicePrincipal) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MicrosoftGraphServicePrincipal) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *MicrosoftGraphServicePrincipal) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *MicrosoftGraphServicePrincipal) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetOauth2PermissionGrants

`func (o *MicrosoftGraphServicePrincipal) GetOauth2PermissionGrants() []MicrosoftGraphOAuth2PermissionGrant`

GetOauth2PermissionGrants returns the Oauth2PermissionGrants field if non-nil, zero value otherwise.

### GetOauth2PermissionGrantsOk

`func (o *MicrosoftGraphServicePrincipal) GetOauth2PermissionGrantsOk() (*[]MicrosoftGraphOAuth2PermissionGrant, bool)`

GetOauth2PermissionGrantsOk returns a tuple with the Oauth2PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2PermissionGrants

`func (o *MicrosoftGraphServicePrincipal) SetOauth2PermissionGrants(v []MicrosoftGraphOAuth2PermissionGrant)`

SetOauth2PermissionGrants sets Oauth2PermissionGrants field to given value.

### HasOauth2PermissionGrants

`func (o *MicrosoftGraphServicePrincipal) HasOauth2PermissionGrants() bool`

HasOauth2PermissionGrants returns a boolean if a field has been set.

### GetOwnedObjects

`func (o *MicrosoftGraphServicePrincipal) GetOwnedObjects() []MicrosoftGraphDirectoryObject`

GetOwnedObjects returns the OwnedObjects field if non-nil, zero value otherwise.

### GetOwnedObjectsOk

`func (o *MicrosoftGraphServicePrincipal) GetOwnedObjectsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnedObjectsOk returns a tuple with the OwnedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedObjects

`func (o *MicrosoftGraphServicePrincipal) SetOwnedObjects(v []MicrosoftGraphDirectoryObject)`

SetOwnedObjects sets OwnedObjects field to given value.

### HasOwnedObjects

`func (o *MicrosoftGraphServicePrincipal) HasOwnedObjects() bool`

HasOwnedObjects returns a boolean if a field has been set.

### GetOwners

`func (o *MicrosoftGraphServicePrincipal) GetOwners() []MicrosoftGraphDirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MicrosoftGraphServicePrincipal) GetOwnersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *MicrosoftGraphServicePrincipal) SetOwners(v []MicrosoftGraphDirectoryObject)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *MicrosoftGraphServicePrincipal) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetTokenIssuancePolicies

`func (o *MicrosoftGraphServicePrincipal) GetTokenIssuancePolicies() []MicrosoftGraphTokenIssuancePolicy`

GetTokenIssuancePolicies returns the TokenIssuancePolicies field if non-nil, zero value otherwise.

### GetTokenIssuancePoliciesOk

`func (o *MicrosoftGraphServicePrincipal) GetTokenIssuancePoliciesOk() (*[]MicrosoftGraphTokenIssuancePolicy, bool)`

GetTokenIssuancePoliciesOk returns a tuple with the TokenIssuancePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIssuancePolicies

`func (o *MicrosoftGraphServicePrincipal) SetTokenIssuancePolicies(v []MicrosoftGraphTokenIssuancePolicy)`

SetTokenIssuancePolicies sets TokenIssuancePolicies field to given value.

### HasTokenIssuancePolicies

`func (o *MicrosoftGraphServicePrincipal) HasTokenIssuancePolicies() bool`

HasTokenIssuancePolicies returns a boolean if a field has been set.

### GetTokenLifetimePolicies

`func (o *MicrosoftGraphServicePrincipal) GetTokenLifetimePolicies() []MicrosoftGraphTokenLifetimePolicy`

GetTokenLifetimePolicies returns the TokenLifetimePolicies field if non-nil, zero value otherwise.

### GetTokenLifetimePoliciesOk

`func (o *MicrosoftGraphServicePrincipal) GetTokenLifetimePoliciesOk() (*[]MicrosoftGraphTokenLifetimePolicy, bool)`

GetTokenLifetimePoliciesOk returns a tuple with the TokenLifetimePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifetimePolicies

`func (o *MicrosoftGraphServicePrincipal) SetTokenLifetimePolicies(v []MicrosoftGraphTokenLifetimePolicy)`

SetTokenLifetimePolicies sets TokenLifetimePolicies field to given value.

### HasTokenLifetimePolicies

`func (o *MicrosoftGraphServicePrincipal) HasTokenLifetimePolicies() bool`

HasTokenLifetimePolicies returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *MicrosoftGraphServicePrincipal) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *MicrosoftGraphServicePrincipal) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *MicrosoftGraphServicePrincipal) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *MicrosoftGraphServicePrincipal) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


