# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddIns** | Pointer to [**[]MicrosoftGraphAddIn**](MicrosoftGraphAddIn.md) | Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its &#39;FileHandler&#39; functionality. This will let services like Office 365 call the application in the context of a document the user is working on. | [optional] 
**Api** | Pointer to [**AnyOfmicrosoftGraphApiApplication**](anyOf&lt;microsoft.graph.apiApplication&gt;.md) | Specifies settings for an application that implements a web API. | [optional] 
**AppId** | Pointer to **NullableString** | The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only. | [optional] 
**ApplicationTemplateId** | Pointer to **NullableString** | Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne). | [optional] 
**AppRoles** | Pointer to [**[]MicrosoftGraphAppRole**](MicrosoftGraphAppRole.md) | The collection of roles assigned to the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy. | [optional] 
**Description** | Pointer to **NullableString** | Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search. | [optional] 
**DisabledByMicrosoftStatus** | Pointer to **NullableString** | Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not). | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**GroupMembershipClaims** | Pointer to **NullableString** | Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of). | [optional] 
**IdentifierUris** | Pointer to **[]string** | Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you&#39;ll reference in your API&#39;s code, and it must be globally unique. You can use the default value provided, which is in the form api://&lt;application-client-id&gt;, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith). | [optional] 
**Info** | Pointer to [**AnyOfmicrosoftGraphInformationalUrl**](anyOf&lt;microsoft.graph.informationalUrl&gt;.md) | Basic profile information of the application such as  app&#39;s marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values). | [optional] 
**IsDeviceOnlyAuthSupported** | Pointer to **NullableBool** | Specifies whether this application supports device authentication without a user. The default is false. | [optional] 
**IsFallbackPublicClient** | Pointer to **NullableBool** | Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property. | [optional] 
**KeyCredentials** | Pointer to [**[]MicrosoftGraphKeyCredential**](MicrosoftGraphKeyCredential.md) | The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le). | [optional] 
**Logo** | Pointer to **string** | The main logo for the application. Not nullable. | [optional] 
**Notes** | Pointer to **NullableString** | Notes relevant for the management of the application. | [optional] 
**Oauth2RequirePostResponse** | Pointer to **bool** |  | [optional] 
**OptionalClaims** | Pointer to [**AnyOfmicrosoftGraphOptionalClaims**](anyOf&lt;microsoft.graph.optionalClaims&gt;.md) | Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app. | [optional] 
**ParentalControlSettings** | Pointer to [**AnyOfmicrosoftGraphParentalControlSettings**](anyOf&lt;microsoft.graph.parentalControlSettings&gt;.md) | Specifies parental control settings for an application. | [optional] 
**PasswordCredentials** | Pointer to [**[]MicrosoftGraphPasswordCredential**](MicrosoftGraphPasswordCredential.md) | The collection of password credentials associated with the application. Not nullable. | [optional] 
**PublicClient** | Pointer to [**AnyOfmicrosoftGraphPublicClientApplication**](anyOf&lt;microsoft.graph.publicClientApplication&gt;.md) | Specifies settings for installed clients such as desktop or mobile devices. | [optional] 
**PublisherDomain** | Pointer to **NullableString** | The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application&#39;s publisher domain. Supports $filter (eq, ne, ge, le, startsWith). | [optional] 
**RequiredResourceAccess** | Pointer to [**[]MicrosoftGraphRequiredResourceAccess**](MicrosoftGraphRequiredResourceAccess.md) | Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. Not nullable. Supports $filter (eq, not, ge, le). | [optional] 
**SignInAudience** | Pointer to **NullableString** | Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table below. Supports $filter (eq, ne, not). | [optional] 
**Spa** | Pointer to [**AnyOfmicrosoftGraphSpaApplication**](anyOf&lt;microsoft.graph.spaApplication&gt;.md) | Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens. | [optional] 
**Tags** | Pointer to **[]string** | Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**TokenEncryptionKeyId** | Pointer to **NullableString** | Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user. | [optional] 
**VerifiedPublisher** | Pointer to [**AnyOfmicrosoftGraphVerifiedPublisher**](anyOf&lt;microsoft.graph.verifiedPublisher&gt;.md) | Specifies the verified publisher of the application. | [optional] 
**Web** | Pointer to [**AnyOfmicrosoftGraphWebApplication**](anyOf&lt;microsoft.graph.webApplication&gt;.md) | Specifies settings for a web application. | [optional] 
**CreatedOnBehalfOf** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) | Read-only. | [optional] 
**ExtensionProperties** | Pointer to [**[]MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md) | Read-only. Nullable. | [optional] 
**HomeRealmDiscoveryPolicies** | Pointer to [**[]MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md) |  | [optional] 
**Owners** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Directory objects that are owners of the application. Read-only. Nullable. Supports $expand. | [optional] 
**TokenIssuancePolicies** | Pointer to [**[]MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md) |  | [optional] 
**TokenLifetimePolicies** | Pointer to [**[]MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md) | The tokenLifetimePolicies assigned to this application. Supports $expand. | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddIns

`func (o *Application) GetAddIns() []MicrosoftGraphAddIn`

GetAddIns returns the AddIns field if non-nil, zero value otherwise.

### GetAddInsOk

`func (o *Application) GetAddInsOk() (*[]MicrosoftGraphAddIn, bool)`

GetAddInsOk returns a tuple with the AddIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddIns

`func (o *Application) SetAddIns(v []MicrosoftGraphAddIn)`

SetAddIns sets AddIns field to given value.

### HasAddIns

`func (o *Application) HasAddIns() bool`

HasAddIns returns a boolean if a field has been set.

### GetApi

`func (o *Application) GetApi() AnyOfmicrosoftGraphApiApplication`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *Application) GetApiOk() (*AnyOfmicrosoftGraphApiApplication, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *Application) SetApi(v AnyOfmicrosoftGraphApiApplication)`

SetApi sets Api field to given value.

### HasApi

`func (o *Application) HasApi() bool`

HasApi returns a boolean if a field has been set.

### SetApiNil

`func (o *Application) SetApiNil(b bool)`

 SetApiNil sets the value for Api to be an explicit nil

### UnsetApi
`func (o *Application) UnsetApi()`

UnsetApi ensures that no value is present for Api, not even an explicit nil
### GetAppId

`func (o *Application) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Application) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Application) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Application) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *Application) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *Application) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetApplicationTemplateId

`func (o *Application) GetApplicationTemplateId() string`

GetApplicationTemplateId returns the ApplicationTemplateId field if non-nil, zero value otherwise.

### GetApplicationTemplateIdOk

`func (o *Application) GetApplicationTemplateIdOk() (*string, bool)`

GetApplicationTemplateIdOk returns a tuple with the ApplicationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTemplateId

`func (o *Application) SetApplicationTemplateId(v string)`

SetApplicationTemplateId sets ApplicationTemplateId field to given value.

### HasApplicationTemplateId

`func (o *Application) HasApplicationTemplateId() bool`

HasApplicationTemplateId returns a boolean if a field has been set.

### SetApplicationTemplateIdNil

`func (o *Application) SetApplicationTemplateIdNil(b bool)`

 SetApplicationTemplateIdNil sets the value for ApplicationTemplateId to be an explicit nil

### UnsetApplicationTemplateId
`func (o *Application) UnsetApplicationTemplateId()`

UnsetApplicationTemplateId ensures that no value is present for ApplicationTemplateId, not even an explicit nil
### GetAppRoles

`func (o *Application) GetAppRoles() []MicrosoftGraphAppRole`

GetAppRoles returns the AppRoles field if non-nil, zero value otherwise.

### GetAppRolesOk

`func (o *Application) GetAppRolesOk() (*[]MicrosoftGraphAppRole, bool)`

GetAppRolesOk returns a tuple with the AppRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoles

`func (o *Application) SetAppRoles(v []MicrosoftGraphAppRole)`

SetAppRoles sets AppRoles field to given value.

### HasAppRoles

`func (o *Application) HasAppRoles() bool`

HasAppRoles returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Application) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Application) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Application) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Application) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Application) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Application) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Application) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Application) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabledByMicrosoftStatus

`func (o *Application) GetDisabledByMicrosoftStatus() string`

GetDisabledByMicrosoftStatus returns the DisabledByMicrosoftStatus field if non-nil, zero value otherwise.

### GetDisabledByMicrosoftStatusOk

`func (o *Application) GetDisabledByMicrosoftStatusOk() (*string, bool)`

GetDisabledByMicrosoftStatusOk returns a tuple with the DisabledByMicrosoftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledByMicrosoftStatus

`func (o *Application) SetDisabledByMicrosoftStatus(v string)`

SetDisabledByMicrosoftStatus sets DisabledByMicrosoftStatus field to given value.

### HasDisabledByMicrosoftStatus

`func (o *Application) HasDisabledByMicrosoftStatus() bool`

HasDisabledByMicrosoftStatus returns a boolean if a field has been set.

### SetDisabledByMicrosoftStatusNil

`func (o *Application) SetDisabledByMicrosoftStatusNil(b bool)`

 SetDisabledByMicrosoftStatusNil sets the value for DisabledByMicrosoftStatus to be an explicit nil

### UnsetDisabledByMicrosoftStatus
`func (o *Application) UnsetDisabledByMicrosoftStatus()`

UnsetDisabledByMicrosoftStatus ensures that no value is present for DisabledByMicrosoftStatus, not even an explicit nil
### GetDisplayName

`func (o *Application) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Application) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Application) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Application) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Application) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Application) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetGroupMembershipClaims

`func (o *Application) GetGroupMembershipClaims() string`

GetGroupMembershipClaims returns the GroupMembershipClaims field if non-nil, zero value otherwise.

### GetGroupMembershipClaimsOk

`func (o *Application) GetGroupMembershipClaimsOk() (*string, bool)`

GetGroupMembershipClaimsOk returns a tuple with the GroupMembershipClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipClaims

`func (o *Application) SetGroupMembershipClaims(v string)`

SetGroupMembershipClaims sets GroupMembershipClaims field to given value.

### HasGroupMembershipClaims

`func (o *Application) HasGroupMembershipClaims() bool`

HasGroupMembershipClaims returns a boolean if a field has been set.

### SetGroupMembershipClaimsNil

`func (o *Application) SetGroupMembershipClaimsNil(b bool)`

 SetGroupMembershipClaimsNil sets the value for GroupMembershipClaims to be an explicit nil

### UnsetGroupMembershipClaims
`func (o *Application) UnsetGroupMembershipClaims()`

UnsetGroupMembershipClaims ensures that no value is present for GroupMembershipClaims, not even an explicit nil
### GetIdentifierUris

`func (o *Application) GetIdentifierUris() []string`

GetIdentifierUris returns the IdentifierUris field if non-nil, zero value otherwise.

### GetIdentifierUrisOk

`func (o *Application) GetIdentifierUrisOk() (*[]string, bool)`

GetIdentifierUrisOk returns a tuple with the IdentifierUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierUris

`func (o *Application) SetIdentifierUris(v []string)`

SetIdentifierUris sets IdentifierUris field to given value.

### HasIdentifierUris

`func (o *Application) HasIdentifierUris() bool`

HasIdentifierUris returns a boolean if a field has been set.

### GetInfo

`func (o *Application) GetInfo() AnyOfmicrosoftGraphInformationalUrl`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Application) GetInfoOk() (*AnyOfmicrosoftGraphInformationalUrl, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Application) SetInfo(v AnyOfmicrosoftGraphInformationalUrl)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Application) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *Application) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *Application) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetIsDeviceOnlyAuthSupported

`func (o *Application) GetIsDeviceOnlyAuthSupported() bool`

GetIsDeviceOnlyAuthSupported returns the IsDeviceOnlyAuthSupported field if non-nil, zero value otherwise.

### GetIsDeviceOnlyAuthSupportedOk

`func (o *Application) GetIsDeviceOnlyAuthSupportedOk() (*bool, bool)`

GetIsDeviceOnlyAuthSupportedOk returns a tuple with the IsDeviceOnlyAuthSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeviceOnlyAuthSupported

`func (o *Application) SetIsDeviceOnlyAuthSupported(v bool)`

SetIsDeviceOnlyAuthSupported sets IsDeviceOnlyAuthSupported field to given value.

### HasIsDeviceOnlyAuthSupported

`func (o *Application) HasIsDeviceOnlyAuthSupported() bool`

HasIsDeviceOnlyAuthSupported returns a boolean if a field has been set.

### SetIsDeviceOnlyAuthSupportedNil

`func (o *Application) SetIsDeviceOnlyAuthSupportedNil(b bool)`

 SetIsDeviceOnlyAuthSupportedNil sets the value for IsDeviceOnlyAuthSupported to be an explicit nil

### UnsetIsDeviceOnlyAuthSupported
`func (o *Application) UnsetIsDeviceOnlyAuthSupported()`

UnsetIsDeviceOnlyAuthSupported ensures that no value is present for IsDeviceOnlyAuthSupported, not even an explicit nil
### GetIsFallbackPublicClient

`func (o *Application) GetIsFallbackPublicClient() bool`

GetIsFallbackPublicClient returns the IsFallbackPublicClient field if non-nil, zero value otherwise.

### GetIsFallbackPublicClientOk

`func (o *Application) GetIsFallbackPublicClientOk() (*bool, bool)`

GetIsFallbackPublicClientOk returns a tuple with the IsFallbackPublicClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFallbackPublicClient

`func (o *Application) SetIsFallbackPublicClient(v bool)`

SetIsFallbackPublicClient sets IsFallbackPublicClient field to given value.

### HasIsFallbackPublicClient

`func (o *Application) HasIsFallbackPublicClient() bool`

HasIsFallbackPublicClient returns a boolean if a field has been set.

### SetIsFallbackPublicClientNil

`func (o *Application) SetIsFallbackPublicClientNil(b bool)`

 SetIsFallbackPublicClientNil sets the value for IsFallbackPublicClient to be an explicit nil

### UnsetIsFallbackPublicClient
`func (o *Application) UnsetIsFallbackPublicClient()`

UnsetIsFallbackPublicClient ensures that no value is present for IsFallbackPublicClient, not even an explicit nil
### GetKeyCredentials

`func (o *Application) GetKeyCredentials() []MicrosoftGraphKeyCredential`

GetKeyCredentials returns the KeyCredentials field if non-nil, zero value otherwise.

### GetKeyCredentialsOk

`func (o *Application) GetKeyCredentialsOk() (*[]MicrosoftGraphKeyCredential, bool)`

GetKeyCredentialsOk returns a tuple with the KeyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCredentials

`func (o *Application) SetKeyCredentials(v []MicrosoftGraphKeyCredential)`

SetKeyCredentials sets KeyCredentials field to given value.

### HasKeyCredentials

`func (o *Application) HasKeyCredentials() bool`

HasKeyCredentials returns a boolean if a field has been set.

### GetLogo

`func (o *Application) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Application) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Application) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Application) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetNotes

`func (o *Application) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Application) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Application) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Application) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Application) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Application) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOauth2RequirePostResponse

`func (o *Application) GetOauth2RequirePostResponse() bool`

GetOauth2RequirePostResponse returns the Oauth2RequirePostResponse field if non-nil, zero value otherwise.

### GetOauth2RequirePostResponseOk

`func (o *Application) GetOauth2RequirePostResponseOk() (*bool, bool)`

GetOauth2RequirePostResponseOk returns a tuple with the Oauth2RequirePostResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2RequirePostResponse

`func (o *Application) SetOauth2RequirePostResponse(v bool)`

SetOauth2RequirePostResponse sets Oauth2RequirePostResponse field to given value.

### HasOauth2RequirePostResponse

`func (o *Application) HasOauth2RequirePostResponse() bool`

HasOauth2RequirePostResponse returns a boolean if a field has been set.

### GetOptionalClaims

`func (o *Application) GetOptionalClaims() AnyOfmicrosoftGraphOptionalClaims`

GetOptionalClaims returns the OptionalClaims field if non-nil, zero value otherwise.

### GetOptionalClaimsOk

`func (o *Application) GetOptionalClaimsOk() (*AnyOfmicrosoftGraphOptionalClaims, bool)`

GetOptionalClaimsOk returns a tuple with the OptionalClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalClaims

`func (o *Application) SetOptionalClaims(v AnyOfmicrosoftGraphOptionalClaims)`

SetOptionalClaims sets OptionalClaims field to given value.

### HasOptionalClaims

`func (o *Application) HasOptionalClaims() bool`

HasOptionalClaims returns a boolean if a field has been set.

### SetOptionalClaimsNil

`func (o *Application) SetOptionalClaimsNil(b bool)`

 SetOptionalClaimsNil sets the value for OptionalClaims to be an explicit nil

### UnsetOptionalClaims
`func (o *Application) UnsetOptionalClaims()`

UnsetOptionalClaims ensures that no value is present for OptionalClaims, not even an explicit nil
### GetParentalControlSettings

`func (o *Application) GetParentalControlSettings() AnyOfmicrosoftGraphParentalControlSettings`

GetParentalControlSettings returns the ParentalControlSettings field if non-nil, zero value otherwise.

### GetParentalControlSettingsOk

`func (o *Application) GetParentalControlSettingsOk() (*AnyOfmicrosoftGraphParentalControlSettings, bool)`

GetParentalControlSettingsOk returns a tuple with the ParentalControlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentalControlSettings

`func (o *Application) SetParentalControlSettings(v AnyOfmicrosoftGraphParentalControlSettings)`

SetParentalControlSettings sets ParentalControlSettings field to given value.

### HasParentalControlSettings

`func (o *Application) HasParentalControlSettings() bool`

HasParentalControlSettings returns a boolean if a field has been set.

### SetParentalControlSettingsNil

`func (o *Application) SetParentalControlSettingsNil(b bool)`

 SetParentalControlSettingsNil sets the value for ParentalControlSettings to be an explicit nil

### UnsetParentalControlSettings
`func (o *Application) UnsetParentalControlSettings()`

UnsetParentalControlSettings ensures that no value is present for ParentalControlSettings, not even an explicit nil
### GetPasswordCredentials

`func (o *Application) GetPasswordCredentials() []MicrosoftGraphPasswordCredential`

GetPasswordCredentials returns the PasswordCredentials field if non-nil, zero value otherwise.

### GetPasswordCredentialsOk

`func (o *Application) GetPasswordCredentialsOk() (*[]MicrosoftGraphPasswordCredential, bool)`

GetPasswordCredentialsOk returns a tuple with the PasswordCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCredentials

`func (o *Application) SetPasswordCredentials(v []MicrosoftGraphPasswordCredential)`

SetPasswordCredentials sets PasswordCredentials field to given value.

### HasPasswordCredentials

`func (o *Application) HasPasswordCredentials() bool`

HasPasswordCredentials returns a boolean if a field has been set.

### GetPublicClient

`func (o *Application) GetPublicClient() AnyOfmicrosoftGraphPublicClientApplication`

GetPublicClient returns the PublicClient field if non-nil, zero value otherwise.

### GetPublicClientOk

`func (o *Application) GetPublicClientOk() (*AnyOfmicrosoftGraphPublicClientApplication, bool)`

GetPublicClientOk returns a tuple with the PublicClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicClient

`func (o *Application) SetPublicClient(v AnyOfmicrosoftGraphPublicClientApplication)`

SetPublicClient sets PublicClient field to given value.

### HasPublicClient

`func (o *Application) HasPublicClient() bool`

HasPublicClient returns a boolean if a field has been set.

### SetPublicClientNil

`func (o *Application) SetPublicClientNil(b bool)`

 SetPublicClientNil sets the value for PublicClient to be an explicit nil

### UnsetPublicClient
`func (o *Application) UnsetPublicClient()`

UnsetPublicClient ensures that no value is present for PublicClient, not even an explicit nil
### GetPublisherDomain

`func (o *Application) GetPublisherDomain() string`

GetPublisherDomain returns the PublisherDomain field if non-nil, zero value otherwise.

### GetPublisherDomainOk

`func (o *Application) GetPublisherDomainOk() (*string, bool)`

GetPublisherDomainOk returns a tuple with the PublisherDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherDomain

`func (o *Application) SetPublisherDomain(v string)`

SetPublisherDomain sets PublisherDomain field to given value.

### HasPublisherDomain

`func (o *Application) HasPublisherDomain() bool`

HasPublisherDomain returns a boolean if a field has been set.

### SetPublisherDomainNil

`func (o *Application) SetPublisherDomainNil(b bool)`

 SetPublisherDomainNil sets the value for PublisherDomain to be an explicit nil

### UnsetPublisherDomain
`func (o *Application) UnsetPublisherDomain()`

UnsetPublisherDomain ensures that no value is present for PublisherDomain, not even an explicit nil
### GetRequiredResourceAccess

`func (o *Application) GetRequiredResourceAccess() []MicrosoftGraphRequiredResourceAccess`

GetRequiredResourceAccess returns the RequiredResourceAccess field if non-nil, zero value otherwise.

### GetRequiredResourceAccessOk

`func (o *Application) GetRequiredResourceAccessOk() (*[]MicrosoftGraphRequiredResourceAccess, bool)`

GetRequiredResourceAccessOk returns a tuple with the RequiredResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredResourceAccess

`func (o *Application) SetRequiredResourceAccess(v []MicrosoftGraphRequiredResourceAccess)`

SetRequiredResourceAccess sets RequiredResourceAccess field to given value.

### HasRequiredResourceAccess

`func (o *Application) HasRequiredResourceAccess() bool`

HasRequiredResourceAccess returns a boolean if a field has been set.

### GetSignInAudience

`func (o *Application) GetSignInAudience() string`

GetSignInAudience returns the SignInAudience field if non-nil, zero value otherwise.

### GetSignInAudienceOk

`func (o *Application) GetSignInAudienceOk() (*string, bool)`

GetSignInAudienceOk returns a tuple with the SignInAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInAudience

`func (o *Application) SetSignInAudience(v string)`

SetSignInAudience sets SignInAudience field to given value.

### HasSignInAudience

`func (o *Application) HasSignInAudience() bool`

HasSignInAudience returns a boolean if a field has been set.

### SetSignInAudienceNil

`func (o *Application) SetSignInAudienceNil(b bool)`

 SetSignInAudienceNil sets the value for SignInAudience to be an explicit nil

### UnsetSignInAudience
`func (o *Application) UnsetSignInAudience()`

UnsetSignInAudience ensures that no value is present for SignInAudience, not even an explicit nil
### GetSpa

`func (o *Application) GetSpa() AnyOfmicrosoftGraphSpaApplication`

GetSpa returns the Spa field if non-nil, zero value otherwise.

### GetSpaOk

`func (o *Application) GetSpaOk() (*AnyOfmicrosoftGraphSpaApplication, bool)`

GetSpaOk returns a tuple with the Spa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpa

`func (o *Application) SetSpa(v AnyOfmicrosoftGraphSpaApplication)`

SetSpa sets Spa field to given value.

### HasSpa

`func (o *Application) HasSpa() bool`

HasSpa returns a boolean if a field has been set.

### SetSpaNil

`func (o *Application) SetSpaNil(b bool)`

 SetSpaNil sets the value for Spa to be an explicit nil

### UnsetSpa
`func (o *Application) UnsetSpa()`

UnsetSpa ensures that no value is present for Spa, not even an explicit nil
### GetTags

`func (o *Application) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Application) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Application) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Application) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTokenEncryptionKeyId

`func (o *Application) GetTokenEncryptionKeyId() string`

GetTokenEncryptionKeyId returns the TokenEncryptionKeyId field if non-nil, zero value otherwise.

### GetTokenEncryptionKeyIdOk

`func (o *Application) GetTokenEncryptionKeyIdOk() (*string, bool)`

GetTokenEncryptionKeyIdOk returns a tuple with the TokenEncryptionKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEncryptionKeyId

`func (o *Application) SetTokenEncryptionKeyId(v string)`

SetTokenEncryptionKeyId sets TokenEncryptionKeyId field to given value.

### HasTokenEncryptionKeyId

`func (o *Application) HasTokenEncryptionKeyId() bool`

HasTokenEncryptionKeyId returns a boolean if a field has been set.

### SetTokenEncryptionKeyIdNil

`func (o *Application) SetTokenEncryptionKeyIdNil(b bool)`

 SetTokenEncryptionKeyIdNil sets the value for TokenEncryptionKeyId to be an explicit nil

### UnsetTokenEncryptionKeyId
`func (o *Application) UnsetTokenEncryptionKeyId()`

UnsetTokenEncryptionKeyId ensures that no value is present for TokenEncryptionKeyId, not even an explicit nil
### GetVerifiedPublisher

`func (o *Application) GetVerifiedPublisher() AnyOfmicrosoftGraphVerifiedPublisher`

GetVerifiedPublisher returns the VerifiedPublisher field if non-nil, zero value otherwise.

### GetVerifiedPublisherOk

`func (o *Application) GetVerifiedPublisherOk() (*AnyOfmicrosoftGraphVerifiedPublisher, bool)`

GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisher

`func (o *Application) SetVerifiedPublisher(v AnyOfmicrosoftGraphVerifiedPublisher)`

SetVerifiedPublisher sets VerifiedPublisher field to given value.

### HasVerifiedPublisher

`func (o *Application) HasVerifiedPublisher() bool`

HasVerifiedPublisher returns a boolean if a field has been set.

### SetVerifiedPublisherNil

`func (o *Application) SetVerifiedPublisherNil(b bool)`

 SetVerifiedPublisherNil sets the value for VerifiedPublisher to be an explicit nil

### UnsetVerifiedPublisher
`func (o *Application) UnsetVerifiedPublisher()`

UnsetVerifiedPublisher ensures that no value is present for VerifiedPublisher, not even an explicit nil
### GetWeb

`func (o *Application) GetWeb() AnyOfmicrosoftGraphWebApplication`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *Application) GetWebOk() (*AnyOfmicrosoftGraphWebApplication, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *Application) SetWeb(v AnyOfmicrosoftGraphWebApplication)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *Application) HasWeb() bool`

HasWeb returns a boolean if a field has been set.

### SetWebNil

`func (o *Application) SetWebNil(b bool)`

 SetWebNil sets the value for Web to be an explicit nil

### UnsetWeb
`func (o *Application) UnsetWeb()`

UnsetWeb ensures that no value is present for Web, not even an explicit nil
### GetCreatedOnBehalfOf

`func (o *Application) GetCreatedOnBehalfOf() AnyOfmicrosoftGraphDirectoryObject`

GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field if non-nil, zero value otherwise.

### GetCreatedOnBehalfOfOk

`func (o *Application) GetCreatedOnBehalfOfOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOnBehalfOf

`func (o *Application) SetCreatedOnBehalfOf(v AnyOfmicrosoftGraphDirectoryObject)`

SetCreatedOnBehalfOf sets CreatedOnBehalfOf field to given value.

### HasCreatedOnBehalfOf

`func (o *Application) HasCreatedOnBehalfOf() bool`

HasCreatedOnBehalfOf returns a boolean if a field has been set.

### SetCreatedOnBehalfOfNil

`func (o *Application) SetCreatedOnBehalfOfNil(b bool)`

 SetCreatedOnBehalfOfNil sets the value for CreatedOnBehalfOf to be an explicit nil

### UnsetCreatedOnBehalfOf
`func (o *Application) UnsetCreatedOnBehalfOf()`

UnsetCreatedOnBehalfOf ensures that no value is present for CreatedOnBehalfOf, not even an explicit nil
### GetExtensionProperties

`func (o *Application) GetExtensionProperties() []MicrosoftGraphExtensionProperty`

GetExtensionProperties returns the ExtensionProperties field if non-nil, zero value otherwise.

### GetExtensionPropertiesOk

`func (o *Application) GetExtensionPropertiesOk() (*[]MicrosoftGraphExtensionProperty, bool)`

GetExtensionPropertiesOk returns a tuple with the ExtensionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionProperties

`func (o *Application) SetExtensionProperties(v []MicrosoftGraphExtensionProperty)`

SetExtensionProperties sets ExtensionProperties field to given value.

### HasExtensionProperties

`func (o *Application) HasExtensionProperties() bool`

HasExtensionProperties returns a boolean if a field has been set.

### GetHomeRealmDiscoveryPolicies

`func (o *Application) GetHomeRealmDiscoveryPolicies() []MicrosoftGraphHomeRealmDiscoveryPolicy`

GetHomeRealmDiscoveryPolicies returns the HomeRealmDiscoveryPolicies field if non-nil, zero value otherwise.

### GetHomeRealmDiscoveryPoliciesOk

`func (o *Application) GetHomeRealmDiscoveryPoliciesOk() (*[]MicrosoftGraphHomeRealmDiscoveryPolicy, bool)`

GetHomeRealmDiscoveryPoliciesOk returns a tuple with the HomeRealmDiscoveryPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeRealmDiscoveryPolicies

`func (o *Application) SetHomeRealmDiscoveryPolicies(v []MicrosoftGraphHomeRealmDiscoveryPolicy)`

SetHomeRealmDiscoveryPolicies sets HomeRealmDiscoveryPolicies field to given value.

### HasHomeRealmDiscoveryPolicies

`func (o *Application) HasHomeRealmDiscoveryPolicies() bool`

HasHomeRealmDiscoveryPolicies returns a boolean if a field has been set.

### GetOwners

`func (o *Application) GetOwners() []MicrosoftGraphDirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Application) GetOwnersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *Application) SetOwners(v []MicrosoftGraphDirectoryObject)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *Application) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetTokenIssuancePolicies

`func (o *Application) GetTokenIssuancePolicies() []MicrosoftGraphTokenIssuancePolicy`

GetTokenIssuancePolicies returns the TokenIssuancePolicies field if non-nil, zero value otherwise.

### GetTokenIssuancePoliciesOk

`func (o *Application) GetTokenIssuancePoliciesOk() (*[]MicrosoftGraphTokenIssuancePolicy, bool)`

GetTokenIssuancePoliciesOk returns a tuple with the TokenIssuancePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIssuancePolicies

`func (o *Application) SetTokenIssuancePolicies(v []MicrosoftGraphTokenIssuancePolicy)`

SetTokenIssuancePolicies sets TokenIssuancePolicies field to given value.

### HasTokenIssuancePolicies

`func (o *Application) HasTokenIssuancePolicies() bool`

HasTokenIssuancePolicies returns a boolean if a field has been set.

### GetTokenLifetimePolicies

`func (o *Application) GetTokenLifetimePolicies() []MicrosoftGraphTokenLifetimePolicy`

GetTokenLifetimePolicies returns the TokenLifetimePolicies field if non-nil, zero value otherwise.

### GetTokenLifetimePoliciesOk

`func (o *Application) GetTokenLifetimePoliciesOk() (*[]MicrosoftGraphTokenLifetimePolicy, bool)`

GetTokenLifetimePoliciesOk returns a tuple with the TokenLifetimePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifetimePolicies

`func (o *Application) SetTokenLifetimePolicies(v []MicrosoftGraphTokenLifetimePolicy)`

SetTokenLifetimePolicies sets TokenLifetimePolicies field to given value.

### HasTokenLifetimePolicies

`func (o *Application) HasTokenLifetimePolicies() bool`

HasTokenLifetimePolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


