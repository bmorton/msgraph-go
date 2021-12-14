# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountEnabled** | Pointer to **NullableBool** | true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter (eq, ne, not, and in). | [optional] 
**AgeGroup** | Pointer to **NullableString** | Sets the age group of the user. Allowed values: null, minor, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in). | [optional] 
**AssignedLicenses** | Pointer to [**[]MicrosoftGraphAssignedLicense**](MicrosoftGraphAssignedLicense.md) | The licenses that are assigned to the user, including inherited (group-based) licenses.  Not nullable. Returned only on $select. Supports $filter (eq and not). | [optional] 
**AssignedPlans** | Pointer to [**[]MicrosoftGraphAssignedPlan**](MicrosoftGraphAssignedPlan.md) | The plans that are assigned to the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq and not). | [optional] 
**BusinessPhones** | Pointer to **[]string** | The telephone numbers for the user. NOTE: Although this is a string collection, only one number can be set for this property. Read-only for users synced from on-premises directory. Returned by default. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**City** | Pointer to **NullableString** | The city in which the user is located. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**CompanyName** | Pointer to **NullableString** | The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**ConsentProvidedForMinor** | Pointer to **NullableString** | Sets whether consent has been obtained for minors. Allowed values: null, granted, denied and notRequired. Refer to the legal age group property definitions for further information. Returned only on $select. Supports $filter (eq, ne, not, and in). | [optional] 
**Country** | Pointer to **NullableString** | The country/region in which the user is located; for example, US or UK. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The created date of the user object. Read-only. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in). | [optional] 
**CreationType** | Pointer to **NullableString** | Indicates whether the user account was created through one of the following methods:  As a regular school or work account (null). As an external account (Invitation). As a local account for an Azure Active Directory B2C tenant (LocalAccount). Through self-service sign-up by an internal user using email verification (EmailVerified). Through self-service sign-up by an external user signing up through a link that is part of a user flow (SelfServiceSignUp). Read-only.Returned only on $select. Supports $filter (eq, ne, not, in). | [optional] 
**Department** | Pointer to **NullableString** | The name for the department in which the user works. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, and eq on null values). | [optional] 
**DisplayName** | Pointer to **NullableString** | The name displayed in the address book for the user. This is usually the combination of the user&#39;s first name, middle initial and last name. This property is required when a user is created and it cannot be cleared during updates. Maximum length is 256 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values), $orderBy, and $search. | [optional] 
**EmployeeHireDate** | Pointer to **NullableTime** | The date and time when the user was hired or will start work in case of a future hire. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in). | [optional] 
**EmployeeId** | Pointer to **NullableString** | The employee identifier assigned to the user by the organization. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values). | [optional] 
**EmployeeOrgData** | Pointer to [**AnyOfmicrosoftGraphEmployeeOrgData**](anyOf&lt;microsoft.graph.employeeOrgData&gt;.md) | Represents organization data (e.g. division and costCenter) associated with a user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in). | [optional] 
**EmployeeType** | Pointer to **NullableString** | Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith). | [optional] 
**ExternalUserState** | Pointer to **NullableString** | For an external user invited to the tenant using the invitation API, this property represents the invited user&#39;s invitation status. For invited users, the state can be PendingAcceptance or Accepted, or null for all other users. Returned only on $select. Supports $filter (eq, ne, not , in). | [optional] 
**ExternalUserStateChangeDateTime** | Pointer to **NullableTime** | Shows the timestamp for the latest change to the externalUserState property. Returned only on $select. Supports $filter (eq, ne, not , in). | [optional] 
**FaxNumber** | Pointer to **NullableString** | The fax number of the user. Returned only on $select. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values). | [optional] 
**GivenName** | Pointer to **NullableString** | The given name (first name) of the user. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values). | [optional] 
**Identities** | Pointer to [**[]AnyOfmicrosoftGraphObjectIdentity**](AnyOfmicrosoftGraphObjectIdentity.md) | Represents the identities that can be used to sign in to this user account. An identity can be provided by Microsoft (also known as a local account), by organizations, or by social identity providers such as Facebook, Google, and Microsoft, and tied to a user account. May contain multiple items with the same signInType value. Returned only on $select. Supports $filter (eq) including on null values, only where the signInType is not userPrincipalName. | [optional] 
**ImAddresses** | Pointer to **[]string** | The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user. Read-only. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**IsResourceAccount** | Pointer to **NullableBool** | Do not use – reserved for future use. | [optional] 
**JobTitle** | Pointer to **NullableString** | The user&#39;s job title. Maximum length is 128 characters. Returned by default. Supports $filter (eq, ne, not , ge, le, in, startsWith, and eq on null values). | [optional] 
**LastPasswordChangeDateTime** | Pointer to **NullableTime** | The time when this Azure AD user last changed their password or when their password was created, whichever date the latest action was performed. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select. | [optional] 
**LegalAgeGroupClassification** | Pointer to **NullableString** | Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select. | [optional] 
**LicenseAssignmentStates** | Pointer to [**[]AnyOfmicrosoftGraphLicenseAssignmentState**](AnyOfmicrosoftGraphLicenseAssignmentState.md) | State of license assignments for this user. Read-only. Returned only on $select. | [optional] 
**Mail** | Pointer to **NullableString** | The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.Changes to this property will also update the user&#39;s proxyAddresses collection to include the value as an SMTP address. For Azure AD B2C accounts, this property can be updated up to only ten times with unique SMTP addresses. This property cannot contain accent characters.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith, and eq on null values). | [optional] 
**MailNickname** | Pointer to **NullableString** | The mail alias for the user. This property must be specified when a user is created. Maximum length is 64 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**MobilePhone** | Pointer to **NullableString** | The primary cellular telephone number for the user. Read-only for users synced from on-premises directory. Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**OfficeLocation** | Pointer to **NullableString** | The office location in the user&#39;s place of business. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**OnPremisesDistinguishedName** | Pointer to **NullableString** | Contains the on-premises Active Directory distinguished name or DN. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. | [optional] 
**OnPremisesDomainName** | Pointer to **NullableString** | Contains the on-premises domainFQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. | [optional] 
**OnPremisesExtensionAttributes** | Pointer to [**AnyOfmicrosoftGraphOnPremisesExtensionAttributes**](anyOf&lt;microsoft.graph.onPremisesExtensionAttributes&gt;.md) | Contains extensionAttributes 1-15 for the user. Note that the individual extension attributes are neither selectable nor filterable. For an onPremisesSyncEnabled user, the source of authority for this set of properties is the on-premises and is read-only. For a cloud-only user (where onPremisesSyncEnabled is false), these properties may be set during creation or update. These extension attributes are also known as Exchange custom attributes 1-15. Returned only on $select. Supports $filter (eq, not, ge, le, in, and eq on null values). | [optional] 
**OnPremisesImmutableId** | Pointer to **NullableString** | This property is used to associate an on-premises Active Directory user account to their Azure AD user object. This property must be specified when creating a new user account in the Graph if you are using a federated domain for the user&#39;s userPrincipalName (UPN) property. NOTE: The $ and _ characters cannot be used when specifying this property. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in).. | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | Indicates the last time at which the object was synced with the on-premises directory; for example: 2013-02-16T03:04:54Z. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in). | [optional] 
**OnPremisesProvisioningErrors** | Pointer to [**[]AnyOfmicrosoftGraphOnPremisesProvisioningError**](AnyOfmicrosoftGraphOnPremisesProvisioningError.md) | Errors when using Microsoft synchronization product during provisioning. Returned only on $select. Supports $filter (eq, not, ge, le). | [optional] 
**OnPremisesSamAccountName** | Pointer to **NullableString** | Contains the on-premises samAccountName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith). | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **NullableString** | Contains the on-premises security identifier (SID) for the user that was synchronized from on-premises to the cloud. Read-only. Returned only on $select. Supports $filter (eq) on null values only. | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**OnPremisesUserPrincipalName** | Pointer to **NullableString** | Contains the on-premises userPrincipalName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect. Read-only. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith). | [optional] 
**OtherMails** | Pointer to **[]string** | A list of additional email addresses for the user; for example: [&#39;bob@contoso.com&#39;, &#39;Robert@fabrikam.com&#39;]. NOTE: This property cannot contain accent characters. Returned only on $select. Supports $filter (eq, not, ge, le, in, startsWith). | [optional] 
**PasswordPolicies** | Pointer to **NullableString** | Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword. Returned only on $select. For more information on the default password policies, see Azure AD pasword policies. Supports $filter (ne, not, and eq on null values). | [optional] 
**PasswordProfile** | Pointer to [**AnyOfmicrosoftGraphPasswordProfile**](anyOf&lt;microsoft.graph.passwordProfile&gt;.md) | Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required. NOTE: For Azure B2C tenants, the forceChangePasswordNextSignIn property should be set to false and instead use custom policies and user flows to force password reset at first logon. See Force password reset at first logon.Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code. Maximum length is 40 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The preferred language for the user. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values) | [optional] 
**ProvisionedPlans** | Pointer to [**[]MicrosoftGraphProvisionedPlan**](MicrosoftGraphProvisionedPlan.md) | The plans that are provisioned for the user. Read-only. Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le). | [optional] 
**ProxyAddresses** | Pointer to **[]string** | For example: [&#39;SMTP: bob@contoso.com&#39;, &#39;smtp: bob@sales.contoso.com&#39;]. For Azure AD B2C accounts, this property has a limit of ten unique addresses. Read-only, Not nullable. Returned only on $select. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**ShowInAddressList** | Pointer to **NullableBool** | true if the Outlook global address list should contain this user, otherwise false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false. Returned only on $select. Supports $filter (eq, ne, not, in). | [optional] 
**SignInSessionsValidFromDateTime** | Pointer to **NullableTime** | Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).  If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint. Read-only. Use revokeSignInSessions to reset. Returned only on $select. | [optional] 
**State** | Pointer to **NullableString** | The state or province in the user&#39;s address. Maximum length is 128 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**StreetAddress** | Pointer to **NullableString** | The street address of the user&#39;s place of business. Maximum length is 1024 characters. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**Surname** | Pointer to **NullableString** | The user&#39;s surname (family name or last name). Maximum length is 64 characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**UsageLocation** | Pointer to **NullableString** | A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: US, JP, and GB. Not nullable. Returned only on $select. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user&#39;s email name. The general format is alias@domain, where domain must be present in the tenant&#39;s collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization.NOTE: This property cannot contain accent characters. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, endsWith) and $orderBy. | [optional] 
**UserType** | Pointer to **NullableString** | A string value that can be used to classify user types in your directory, such as Member and Guest. Returned only on $select. Supports $filter (eq, ne, not, in, and eq on null values). NOTE: For more information about the permissions for member and guest users, see What are the default user permissions in Azure Active Directory? | [optional] 
**MailboxSettings** | Pointer to [**AnyOfmicrosoftGraphMailboxSettings**](anyOf&lt;microsoft.graph.mailboxSettings&gt;.md) | Settings for the primary mailbox of the signed-in user. You can get or update settings for sending automatic replies to incoming messages, locale and time zone.Returned only on $select. | [optional] 
**DeviceEnrollmentLimit** | Pointer to **int32** | The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000. | [optional] 
**AboutMe** | Pointer to **NullableString** | A freeform text entry field for the user to describe themselves. Returned only on $select. | [optional] 
**Birthday** | Pointer to **time.Time** | The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select. | [optional] 
**HireDate** | Pointer to **time.Time** | The hire date of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned only on $select.  Note: This property is specific to SharePoint Online. We recommend using the native employeeHireDate property to set and update hire date values using Microsoft Graph APIs. | [optional] 
**Interests** | Pointer to **[]string** | A list for the user to describe their interests. Returned only on $select. | [optional] 
**MySite** | Pointer to **NullableString** | The URL for the user&#39;s personal site. Returned only on $select. | [optional] 
**PastProjects** | Pointer to **[]string** | A list for the user to enumerate their past projects. Returned only on $select. | [optional] 
**PreferredName** | Pointer to **NullableString** | The preferred name for the user. Returned only on $select. | [optional] 
**Responsibilities** | Pointer to **[]string** | A list for the user to enumerate their responsibilities. Returned only on $select. | [optional] 
**Schools** | Pointer to **[]string** | A list for the user to enumerate the schools they have attended. Returned only on $select. | [optional] 
**Skills** | Pointer to **[]string** | A list for the user to enumerate their skills. Returned only on $select. | [optional] 
**AppRoleAssignments** | Pointer to [**[]MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | Represents the app roles a user has been granted for an application. Supports $expand. | [optional] 
**CreatedObjects** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Directory objects that were created by the user. Read-only. Nullable. | [optional] 
**DirectReports** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The users and contacts that report to the user. (The users and contacts that have their manager property set to this user.) Read-only. Nullable. Supports $expand. | [optional] 
**LicenseDetails** | Pointer to [**[]MicrosoftGraphLicenseDetails**](MicrosoftGraphLicenseDetails.md) | A collection of this user&#39;s license details. Read-only. | [optional] 
**Manager** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) | The user or contact that is this user&#39;s manager. Read-only. (HTTP Methods: GET, PUT, DELETE.). Supports $expand. | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand. | [optional] 
**Oauth2PermissionGrants** | Pointer to [**[]MicrosoftGraphOAuth2PermissionGrant**](MicrosoftGraphOAuth2PermissionGrant.md) |  | [optional] 
**OwnedDevices** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Devices that are owned by the user. Read-only. Nullable. Supports $expand. | [optional] 
**OwnedObjects** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Directory objects that are owned by the user. Read-only. Nullable. Supports $expand. | [optional] 
**RegisteredDevices** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Devices that are registered for the user. Read-only. Nullable. Supports $expand. | [optional] 
**ScopedRoleMemberOf** | Pointer to [**[]MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | The scoped-role administrative unit memberships for this user. Read-only. Nullable. | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 
**Calendar** | Pointer to [**AnyOfmicrosoftGraphCalendar**](anyOf&lt;microsoft.graph.calendar&gt;.md) | The user&#39;s primary calendar. Read-only. | [optional] 
**CalendarGroups** | Pointer to [**[]MicrosoftGraphCalendarGroup**](MicrosoftGraphCalendarGroup.md) | The user&#39;s calendar groups. Read-only. Nullable. | [optional] 
**Calendars** | Pointer to [**[]MicrosoftGraphCalendar**](MicrosoftGraphCalendar.md) | The user&#39;s calendars. Read-only. Nullable. | [optional] 
**CalendarView** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The calendar view for the calendar. Read-only. Nullable. | [optional] 
**ContactFolders** | Pointer to [**[]MicrosoftGraphContactFolder**](MicrosoftGraphContactFolder.md) | The user&#39;s contacts folders. Read-only. Nullable. | [optional] 
**Contacts** | Pointer to [**[]MicrosoftGraphContact**](MicrosoftGraphContact.md) | The user&#39;s contacts. Read-only. Nullable. | [optional] 
**Events** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The user&#39;s events. Default is to show Events under the Default Calendar. Read-only. Nullable. | [optional] 
**InferenceClassification** | Pointer to [**AnyOfmicrosoftGraphInferenceClassification**](anyOf&lt;microsoft.graph.inferenceClassification&gt;.md) | Relevance classification of the user&#39;s messages based on explicit designations which override inferred relevance or importance. | [optional] 
**MailFolders** | Pointer to [**[]MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md) | The user&#39;s mail folders. Read-only. Nullable. | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphMessage**](MicrosoftGraphMessage.md) | The messages in a mailbox or folder. Read-only. Nullable. | [optional] 
**Outlook** | Pointer to [**AnyOfmicrosoftGraphOutlookUser**](anyOf&lt;microsoft.graph.outlookUser&gt;.md) | Read-only. | [optional] 
**People** | Pointer to [**[]MicrosoftGraphPerson**](MicrosoftGraphPerson.md) | People that are relevant to the user. Read-only. Nullable. | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) | The user&#39;s OneDrive. Read-only. | [optional] 
**Drives** | Pointer to [**[]MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | A collection of drives available for this user. Read-only. | [optional] 
**FollowedSites** | Pointer to [**[]MicrosoftGraphSite**](MicrosoftGraphSite.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the user. Read-only. Nullable. | [optional] 
**AgreementAcceptances** | Pointer to [**[]MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md) | The user&#39;s terms of use acceptance statuses. Read-only. Nullable. | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md) | The managed devices associated with the user. | [optional] 
**ManagedAppRegistrations** | Pointer to [**[]MicrosoftGraphManagedAppRegistration**](MicrosoftGraphManagedAppRegistration.md) | Zero or more managed app registrations that belong to the user. | [optional] 
**DeviceManagementTroubleshootingEvents** | Pointer to [**[]MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md) | The list of troubleshooting events for this user. | [optional] 
**Planner** | Pointer to [**AnyOfmicrosoftGraphPlannerUser**](anyOf&lt;microsoft.graph.plannerUser&gt;.md) | Entry-point to the Planner resource that might exist for a user. Read-only. | [optional] 
**Insights** | Pointer to [**AnyOfmicrosoftGraphOfficeGraphInsights**](anyOf&lt;microsoft.graph.officeGraphInsights&gt;.md) | Read-only. Nullable. | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphUserSettings**](anyOf&lt;microsoft.graph.userSettings&gt;.md) | Read-only. Nullable. | [optional] 
**Onenote** | Pointer to [**AnyOfmicrosoftGraphOnenote**](anyOf&lt;microsoft.graph.onenote&gt;.md) | Read-only. | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphProfilePhoto**](anyOf&lt;microsoft.graph.profilePhoto&gt;.md) | The user&#39;s profile photo. Read-only. | [optional] 
**Photos** | Pointer to [**[]MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md) | Read-only. Nullable. | [optional] 
**Activities** | Pointer to [**[]MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md) | The user&#39;s activities across devices. Read-only. Nullable. | [optional] 
**OnlineMeetings** | Pointer to [**[]MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md) |  | [optional] 
**Presence** | Pointer to [**AnyOfmicrosoftGraphPresence**](anyOf&lt;microsoft.graph.presence&gt;.md) |  | [optional] 
**Authentication** | Pointer to [**AnyOfmicrosoftGraphAuthentication**](anyOf&lt;microsoft.graph.authentication&gt;.md) |  | [optional] 
**Chats** | Pointer to [**[]MicrosoftGraphChat**](MicrosoftGraphChat.md) |  | [optional] 
**JoinedTeams** | Pointer to [**[]MicrosoftGraphTeam**](MicrosoftGraphTeam.md) | The Microsoft Teams teams that the user is a member of. Read-only. Nullable. | [optional] 
**Teamwork** | Pointer to [**AnyOfmicrosoftGraphUserTeamwork**](anyOf&lt;microsoft.graph.userTeamwork&gt;.md) | A container for Microsoft Teams features available for the user. Read-only. Nullable. | [optional] 
**Todo** | Pointer to [**AnyOfmicrosoftGraphTodo**](anyOf&lt;microsoft.graph.todo&gt;.md) | Represents the To Do services available to a user. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEnabled

`func (o *User) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *User) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *User) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *User) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabledNil

`func (o *User) SetAccountEnabledNil(b bool)`

 SetAccountEnabledNil sets the value for AccountEnabled to be an explicit nil

### UnsetAccountEnabled
`func (o *User) UnsetAccountEnabled()`

UnsetAccountEnabled ensures that no value is present for AccountEnabled, not even an explicit nil
### GetAgeGroup

`func (o *User) GetAgeGroup() string`

GetAgeGroup returns the AgeGroup field if non-nil, zero value otherwise.

### GetAgeGroupOk

`func (o *User) GetAgeGroupOk() (*string, bool)`

GetAgeGroupOk returns a tuple with the AgeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGroup

`func (o *User) SetAgeGroup(v string)`

SetAgeGroup sets AgeGroup field to given value.

### HasAgeGroup

`func (o *User) HasAgeGroup() bool`

HasAgeGroup returns a boolean if a field has been set.

### SetAgeGroupNil

`func (o *User) SetAgeGroupNil(b bool)`

 SetAgeGroupNil sets the value for AgeGroup to be an explicit nil

### UnsetAgeGroup
`func (o *User) UnsetAgeGroup()`

UnsetAgeGroup ensures that no value is present for AgeGroup, not even an explicit nil
### GetAssignedLicenses

`func (o *User) GetAssignedLicenses() []MicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *User) GetAssignedLicensesOk() (*[]MicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLicenses

`func (o *User) SetAssignedLicenses(v []MicrosoftGraphAssignedLicense)`

SetAssignedLicenses sets AssignedLicenses field to given value.

### HasAssignedLicenses

`func (o *User) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### GetAssignedPlans

`func (o *User) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *User) GetAssignedPlansOk() (*[]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedPlans

`func (o *User) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans sets AssignedPlans field to given value.

### HasAssignedPlans

`func (o *User) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### GetBusinessPhones

`func (o *User) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *User) GetBusinessPhonesOk() (*[]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhones

`func (o *User) SetBusinessPhones(v []string)`

SetBusinessPhones sets BusinessPhones field to given value.

### HasBusinessPhones

`func (o *User) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### GetCity

`func (o *User) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *User) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *User) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *User) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *User) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *User) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCompanyName

`func (o *User) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *User) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *User) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *User) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *User) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *User) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetConsentProvidedForMinor

`func (o *User) GetConsentProvidedForMinor() string`

GetConsentProvidedForMinor returns the ConsentProvidedForMinor field if non-nil, zero value otherwise.

### GetConsentProvidedForMinorOk

`func (o *User) GetConsentProvidedForMinorOk() (*string, bool)`

GetConsentProvidedForMinorOk returns a tuple with the ConsentProvidedForMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentProvidedForMinor

`func (o *User) SetConsentProvidedForMinor(v string)`

SetConsentProvidedForMinor sets ConsentProvidedForMinor field to given value.

### HasConsentProvidedForMinor

`func (o *User) HasConsentProvidedForMinor() bool`

HasConsentProvidedForMinor returns a boolean if a field has been set.

### SetConsentProvidedForMinorNil

`func (o *User) SetConsentProvidedForMinorNil(b bool)`

 SetConsentProvidedForMinorNil sets the value for ConsentProvidedForMinor to be an explicit nil

### UnsetConsentProvidedForMinor
`func (o *User) UnsetConsentProvidedForMinor()`

UnsetConsentProvidedForMinor ensures that no value is present for ConsentProvidedForMinor, not even an explicit nil
### GetCountry

`func (o *User) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *User) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *User) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *User) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *User) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCreatedDateTime

`func (o *User) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *User) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *User) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *User) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *User) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *User) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetCreationType

`func (o *User) GetCreationType() string`

GetCreationType returns the CreationType field if non-nil, zero value otherwise.

### GetCreationTypeOk

`func (o *User) GetCreationTypeOk() (*string, bool)`

GetCreationTypeOk returns a tuple with the CreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationType

`func (o *User) SetCreationType(v string)`

SetCreationType sets CreationType field to given value.

### HasCreationType

`func (o *User) HasCreationType() bool`

HasCreationType returns a boolean if a field has been set.

### SetCreationTypeNil

`func (o *User) SetCreationTypeNil(b bool)`

 SetCreationTypeNil sets the value for CreationType to be an explicit nil

### UnsetCreationType
`func (o *User) UnsetCreationType()`

UnsetCreationType ensures that no value is present for CreationType, not even an explicit nil
### GetDepartment

`func (o *User) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *User) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *User) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *User) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *User) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *User) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *User) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *User) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmployeeHireDate

`func (o *User) GetEmployeeHireDate() time.Time`

GetEmployeeHireDate returns the EmployeeHireDate field if non-nil, zero value otherwise.

### GetEmployeeHireDateOk

`func (o *User) GetEmployeeHireDateOk() (*time.Time, bool)`

GetEmployeeHireDateOk returns a tuple with the EmployeeHireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeHireDate

`func (o *User) SetEmployeeHireDate(v time.Time)`

SetEmployeeHireDate sets EmployeeHireDate field to given value.

### HasEmployeeHireDate

`func (o *User) HasEmployeeHireDate() bool`

HasEmployeeHireDate returns a boolean if a field has been set.

### SetEmployeeHireDateNil

`func (o *User) SetEmployeeHireDateNil(b bool)`

 SetEmployeeHireDateNil sets the value for EmployeeHireDate to be an explicit nil

### UnsetEmployeeHireDate
`func (o *User) UnsetEmployeeHireDate()`

UnsetEmployeeHireDate ensures that no value is present for EmployeeHireDate, not even an explicit nil
### GetEmployeeId

`func (o *User) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *User) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *User) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.

### HasEmployeeId

`func (o *User) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### SetEmployeeIdNil

`func (o *User) SetEmployeeIdNil(b bool)`

 SetEmployeeIdNil sets the value for EmployeeId to be an explicit nil

### UnsetEmployeeId
`func (o *User) UnsetEmployeeId()`

UnsetEmployeeId ensures that no value is present for EmployeeId, not even an explicit nil
### GetEmployeeOrgData

`func (o *User) GetEmployeeOrgData() AnyOfmicrosoftGraphEmployeeOrgData`

GetEmployeeOrgData returns the EmployeeOrgData field if non-nil, zero value otherwise.

### GetEmployeeOrgDataOk

`func (o *User) GetEmployeeOrgDataOk() (*AnyOfmicrosoftGraphEmployeeOrgData, bool)`

GetEmployeeOrgDataOk returns a tuple with the EmployeeOrgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeOrgData

`func (o *User) SetEmployeeOrgData(v AnyOfmicrosoftGraphEmployeeOrgData)`

SetEmployeeOrgData sets EmployeeOrgData field to given value.

### HasEmployeeOrgData

`func (o *User) HasEmployeeOrgData() bool`

HasEmployeeOrgData returns a boolean if a field has been set.

### SetEmployeeOrgDataNil

`func (o *User) SetEmployeeOrgDataNil(b bool)`

 SetEmployeeOrgDataNil sets the value for EmployeeOrgData to be an explicit nil

### UnsetEmployeeOrgData
`func (o *User) UnsetEmployeeOrgData()`

UnsetEmployeeOrgData ensures that no value is present for EmployeeOrgData, not even an explicit nil
### GetEmployeeType

`func (o *User) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *User) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *User) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *User) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### SetEmployeeTypeNil

`func (o *User) SetEmployeeTypeNil(b bool)`

 SetEmployeeTypeNil sets the value for EmployeeType to be an explicit nil

### UnsetEmployeeType
`func (o *User) UnsetEmployeeType()`

UnsetEmployeeType ensures that no value is present for EmployeeType, not even an explicit nil
### GetExternalUserState

`func (o *User) GetExternalUserState() string`

GetExternalUserState returns the ExternalUserState field if non-nil, zero value otherwise.

### GetExternalUserStateOk

`func (o *User) GetExternalUserStateOk() (*string, bool)`

GetExternalUserStateOk returns a tuple with the ExternalUserState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserState

`func (o *User) SetExternalUserState(v string)`

SetExternalUserState sets ExternalUserState field to given value.

### HasExternalUserState

`func (o *User) HasExternalUserState() bool`

HasExternalUserState returns a boolean if a field has been set.

### SetExternalUserStateNil

`func (o *User) SetExternalUserStateNil(b bool)`

 SetExternalUserStateNil sets the value for ExternalUserState to be an explicit nil

### UnsetExternalUserState
`func (o *User) UnsetExternalUserState()`

UnsetExternalUserState ensures that no value is present for ExternalUserState, not even an explicit nil
### GetExternalUserStateChangeDateTime

`func (o *User) GetExternalUserStateChangeDateTime() time.Time`

GetExternalUserStateChangeDateTime returns the ExternalUserStateChangeDateTime field if non-nil, zero value otherwise.

### GetExternalUserStateChangeDateTimeOk

`func (o *User) GetExternalUserStateChangeDateTimeOk() (*time.Time, bool)`

GetExternalUserStateChangeDateTimeOk returns a tuple with the ExternalUserStateChangeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserStateChangeDateTime

`func (o *User) SetExternalUserStateChangeDateTime(v time.Time)`

SetExternalUserStateChangeDateTime sets ExternalUserStateChangeDateTime field to given value.

### HasExternalUserStateChangeDateTime

`func (o *User) HasExternalUserStateChangeDateTime() bool`

HasExternalUserStateChangeDateTime returns a boolean if a field has been set.

### SetExternalUserStateChangeDateTimeNil

`func (o *User) SetExternalUserStateChangeDateTimeNil(b bool)`

 SetExternalUserStateChangeDateTimeNil sets the value for ExternalUserStateChangeDateTime to be an explicit nil

### UnsetExternalUserStateChangeDateTime
`func (o *User) UnsetExternalUserStateChangeDateTime()`

UnsetExternalUserStateChangeDateTime ensures that no value is present for ExternalUserStateChangeDateTime, not even an explicit nil
### GetFaxNumber

`func (o *User) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *User) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *User) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *User) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### SetFaxNumberNil

`func (o *User) SetFaxNumberNil(b bool)`

 SetFaxNumberNil sets the value for FaxNumber to be an explicit nil

### UnsetFaxNumber
`func (o *User) UnsetFaxNumber()`

UnsetFaxNumber ensures that no value is present for FaxNumber, not even an explicit nil
### GetGivenName

`func (o *User) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *User) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *User) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *User) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenNameNil

`func (o *User) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *User) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetIdentities

`func (o *User) GetIdentities() []*AnyOfmicrosoftGraphObjectIdentity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *User) GetIdentitiesOk() (*[]*AnyOfmicrosoftGraphObjectIdentity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *User) SetIdentities(v []*AnyOfmicrosoftGraphObjectIdentity)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *User) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetImAddresses

`func (o *User) GetImAddresses() []*string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *User) GetImAddressesOk() (*[]*string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImAddresses

`func (o *User) SetImAddresses(v []*string)`

SetImAddresses sets ImAddresses field to given value.

### HasImAddresses

`func (o *User) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### GetIsResourceAccount

`func (o *User) GetIsResourceAccount() bool`

GetIsResourceAccount returns the IsResourceAccount field if non-nil, zero value otherwise.

### GetIsResourceAccountOk

`func (o *User) GetIsResourceAccountOk() (*bool, bool)`

GetIsResourceAccountOk returns a tuple with the IsResourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResourceAccount

`func (o *User) SetIsResourceAccount(v bool)`

SetIsResourceAccount sets IsResourceAccount field to given value.

### HasIsResourceAccount

`func (o *User) HasIsResourceAccount() bool`

HasIsResourceAccount returns a boolean if a field has been set.

### SetIsResourceAccountNil

`func (o *User) SetIsResourceAccountNil(b bool)`

 SetIsResourceAccountNil sets the value for IsResourceAccount to be an explicit nil

### UnsetIsResourceAccount
`func (o *User) UnsetIsResourceAccount()`

UnsetIsResourceAccount ensures that no value is present for IsResourceAccount, not even an explicit nil
### GetJobTitle

`func (o *User) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *User) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *User) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *User) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *User) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *User) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetLastPasswordChangeDateTime

`func (o *User) GetLastPasswordChangeDateTime() time.Time`

GetLastPasswordChangeDateTime returns the LastPasswordChangeDateTime field if non-nil, zero value otherwise.

### GetLastPasswordChangeDateTimeOk

`func (o *User) GetLastPasswordChangeDateTimeOk() (*time.Time, bool)`

GetLastPasswordChangeDateTimeOk returns a tuple with the LastPasswordChangeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChangeDateTime

`func (o *User) SetLastPasswordChangeDateTime(v time.Time)`

SetLastPasswordChangeDateTime sets LastPasswordChangeDateTime field to given value.

### HasLastPasswordChangeDateTime

`func (o *User) HasLastPasswordChangeDateTime() bool`

HasLastPasswordChangeDateTime returns a boolean if a field has been set.

### SetLastPasswordChangeDateTimeNil

`func (o *User) SetLastPasswordChangeDateTimeNil(b bool)`

 SetLastPasswordChangeDateTimeNil sets the value for LastPasswordChangeDateTime to be an explicit nil

### UnsetLastPasswordChangeDateTime
`func (o *User) UnsetLastPasswordChangeDateTime()`

UnsetLastPasswordChangeDateTime ensures that no value is present for LastPasswordChangeDateTime, not even an explicit nil
### GetLegalAgeGroupClassification

`func (o *User) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *User) GetLegalAgeGroupClassificationOk() (*string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAgeGroupClassification

`func (o *User) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification sets LegalAgeGroupClassification field to given value.

### HasLegalAgeGroupClassification

`func (o *User) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

### SetLegalAgeGroupClassificationNil

`func (o *User) SetLegalAgeGroupClassificationNil(b bool)`

 SetLegalAgeGroupClassificationNil sets the value for LegalAgeGroupClassification to be an explicit nil

### UnsetLegalAgeGroupClassification
`func (o *User) UnsetLegalAgeGroupClassification()`

UnsetLegalAgeGroupClassification ensures that no value is present for LegalAgeGroupClassification, not even an explicit nil
### GetLicenseAssignmentStates

`func (o *User) GetLicenseAssignmentStates() []*AnyOfmicrosoftGraphLicenseAssignmentState`

GetLicenseAssignmentStates returns the LicenseAssignmentStates field if non-nil, zero value otherwise.

### GetLicenseAssignmentStatesOk

`func (o *User) GetLicenseAssignmentStatesOk() (*[]*AnyOfmicrosoftGraphLicenseAssignmentState, bool)`

GetLicenseAssignmentStatesOk returns a tuple with the LicenseAssignmentStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAssignmentStates

`func (o *User) SetLicenseAssignmentStates(v []*AnyOfmicrosoftGraphLicenseAssignmentState)`

SetLicenseAssignmentStates sets LicenseAssignmentStates field to given value.

### HasLicenseAssignmentStates

`func (o *User) HasLicenseAssignmentStates() bool`

HasLicenseAssignmentStates returns a boolean if a field has been set.

### GetMail

`func (o *User) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *User) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *User) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *User) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *User) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *User) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetMailNickname

`func (o *User) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *User) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *User) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *User) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *User) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *User) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetMobilePhone

`func (o *User) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *User) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *User) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *User) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *User) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *User) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetOfficeLocation

`func (o *User) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *User) GetOfficeLocationOk() (*string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeLocation

`func (o *User) SetOfficeLocation(v string)`

SetOfficeLocation sets OfficeLocation field to given value.

### HasOfficeLocation

`func (o *User) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocationNil

`func (o *User) SetOfficeLocationNil(b bool)`

 SetOfficeLocationNil sets the value for OfficeLocation to be an explicit nil

### UnsetOfficeLocation
`func (o *User) UnsetOfficeLocation()`

UnsetOfficeLocation ensures that no value is present for OfficeLocation, not even an explicit nil
### GetOnPremisesDistinguishedName

`func (o *User) GetOnPremisesDistinguishedName() string`

GetOnPremisesDistinguishedName returns the OnPremisesDistinguishedName field if non-nil, zero value otherwise.

### GetOnPremisesDistinguishedNameOk

`func (o *User) GetOnPremisesDistinguishedNameOk() (*string, bool)`

GetOnPremisesDistinguishedNameOk returns a tuple with the OnPremisesDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDistinguishedName

`func (o *User) SetOnPremisesDistinguishedName(v string)`

SetOnPremisesDistinguishedName sets OnPremisesDistinguishedName field to given value.

### HasOnPremisesDistinguishedName

`func (o *User) HasOnPremisesDistinguishedName() bool`

HasOnPremisesDistinguishedName returns a boolean if a field has been set.

### SetOnPremisesDistinguishedNameNil

`func (o *User) SetOnPremisesDistinguishedNameNil(b bool)`

 SetOnPremisesDistinguishedNameNil sets the value for OnPremisesDistinguishedName to be an explicit nil

### UnsetOnPremisesDistinguishedName
`func (o *User) UnsetOnPremisesDistinguishedName()`

UnsetOnPremisesDistinguishedName ensures that no value is present for OnPremisesDistinguishedName, not even an explicit nil
### GetOnPremisesDomainName

`func (o *User) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *User) GetOnPremisesDomainNameOk() (*string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDomainName

`func (o *User) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName sets OnPremisesDomainName field to given value.

### HasOnPremisesDomainName

`func (o *User) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainNameNil

`func (o *User) SetOnPremisesDomainNameNil(b bool)`

 SetOnPremisesDomainNameNil sets the value for OnPremisesDomainName to be an explicit nil

### UnsetOnPremisesDomainName
`func (o *User) UnsetOnPremisesDomainName()`

UnsetOnPremisesDomainName ensures that no value is present for OnPremisesDomainName, not even an explicit nil
### GetOnPremisesExtensionAttributes

`func (o *User) GetOnPremisesExtensionAttributes() AnyOfmicrosoftGraphOnPremisesExtensionAttributes`

GetOnPremisesExtensionAttributes returns the OnPremisesExtensionAttributes field if non-nil, zero value otherwise.

### GetOnPremisesExtensionAttributesOk

`func (o *User) GetOnPremisesExtensionAttributesOk() (*AnyOfmicrosoftGraphOnPremisesExtensionAttributes, bool)`

GetOnPremisesExtensionAttributesOk returns a tuple with the OnPremisesExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesExtensionAttributes

`func (o *User) SetOnPremisesExtensionAttributes(v AnyOfmicrosoftGraphOnPremisesExtensionAttributes)`

SetOnPremisesExtensionAttributes sets OnPremisesExtensionAttributes field to given value.

### HasOnPremisesExtensionAttributes

`func (o *User) HasOnPremisesExtensionAttributes() bool`

HasOnPremisesExtensionAttributes returns a boolean if a field has been set.

### SetOnPremisesExtensionAttributesNil

`func (o *User) SetOnPremisesExtensionAttributesNil(b bool)`

 SetOnPremisesExtensionAttributesNil sets the value for OnPremisesExtensionAttributes to be an explicit nil

### UnsetOnPremisesExtensionAttributes
`func (o *User) UnsetOnPremisesExtensionAttributes()`

UnsetOnPremisesExtensionAttributes ensures that no value is present for OnPremisesExtensionAttributes, not even an explicit nil
### GetOnPremisesImmutableId

`func (o *User) GetOnPremisesImmutableId() string`

GetOnPremisesImmutableId returns the OnPremisesImmutableId field if non-nil, zero value otherwise.

### GetOnPremisesImmutableIdOk

`func (o *User) GetOnPremisesImmutableIdOk() (*string, bool)`

GetOnPremisesImmutableIdOk returns a tuple with the OnPremisesImmutableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesImmutableId

`func (o *User) SetOnPremisesImmutableId(v string)`

SetOnPremisesImmutableId sets OnPremisesImmutableId field to given value.

### HasOnPremisesImmutableId

`func (o *User) HasOnPremisesImmutableId() bool`

HasOnPremisesImmutableId returns a boolean if a field has been set.

### SetOnPremisesImmutableIdNil

`func (o *User) SetOnPremisesImmutableIdNil(b bool)`

 SetOnPremisesImmutableIdNil sets the value for OnPremisesImmutableId to be an explicit nil

### UnsetOnPremisesImmutableId
`func (o *User) UnsetOnPremisesImmutableId()`

UnsetOnPremisesImmutableId ensures that no value is present for OnPremisesImmutableId, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *User) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *User) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *User) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *User) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *User) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *User) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesProvisioningErrors

`func (o *User) GetOnPremisesProvisioningErrors() []*AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *User) GetOnPremisesProvisioningErrorsOk() (*[]*AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesProvisioningErrors

`func (o *User) SetOnPremisesProvisioningErrors(v []*AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors sets OnPremisesProvisioningErrors field to given value.

### HasOnPremisesProvisioningErrors

`func (o *User) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### GetOnPremisesSamAccountName

`func (o *User) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *User) GetOnPremisesSamAccountNameOk() (*string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSamAccountName

`func (o *User) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName sets OnPremisesSamAccountName field to given value.

### HasOnPremisesSamAccountName

`func (o *User) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountNameNil

`func (o *User) SetOnPremisesSamAccountNameNil(b bool)`

 SetOnPremisesSamAccountNameNil sets the value for OnPremisesSamAccountName to be an explicit nil

### UnsetOnPremisesSamAccountName
`func (o *User) UnsetOnPremisesSamAccountName()`

UnsetOnPremisesSamAccountName ensures that no value is present for OnPremisesSamAccountName, not even an explicit nil
### GetOnPremisesSecurityIdentifier

`func (o *User) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *User) GetOnPremisesSecurityIdentifierOk() (*string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSecurityIdentifier

`func (o *User) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier sets OnPremisesSecurityIdentifier field to given value.

### HasOnPremisesSecurityIdentifier

`func (o *User) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifierNil

`func (o *User) SetOnPremisesSecurityIdentifierNil(b bool)`

 SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil

### UnsetOnPremisesSecurityIdentifier
`func (o *User) UnsetOnPremisesSecurityIdentifier()`

UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *User) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *User) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *User) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *User) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *User) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *User) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetOnPremisesUserPrincipalName

`func (o *User) GetOnPremisesUserPrincipalName() string`

GetOnPremisesUserPrincipalName returns the OnPremisesUserPrincipalName field if non-nil, zero value otherwise.

### GetOnPremisesUserPrincipalNameOk

`func (o *User) GetOnPremisesUserPrincipalNameOk() (*string, bool)`

GetOnPremisesUserPrincipalNameOk returns a tuple with the OnPremisesUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesUserPrincipalName

`func (o *User) SetOnPremisesUserPrincipalName(v string)`

SetOnPremisesUserPrincipalName sets OnPremisesUserPrincipalName field to given value.

### HasOnPremisesUserPrincipalName

`func (o *User) HasOnPremisesUserPrincipalName() bool`

HasOnPremisesUserPrincipalName returns a boolean if a field has been set.

### SetOnPremisesUserPrincipalNameNil

`func (o *User) SetOnPremisesUserPrincipalNameNil(b bool)`

 SetOnPremisesUserPrincipalNameNil sets the value for OnPremisesUserPrincipalName to be an explicit nil

### UnsetOnPremisesUserPrincipalName
`func (o *User) UnsetOnPremisesUserPrincipalName()`

UnsetOnPremisesUserPrincipalName ensures that no value is present for OnPremisesUserPrincipalName, not even an explicit nil
### GetOtherMails

`func (o *User) GetOtherMails() []string`

GetOtherMails returns the OtherMails field if non-nil, zero value otherwise.

### GetOtherMailsOk

`func (o *User) GetOtherMailsOk() (*[]string, bool)`

GetOtherMailsOk returns a tuple with the OtherMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMails

`func (o *User) SetOtherMails(v []string)`

SetOtherMails sets OtherMails field to given value.

### HasOtherMails

`func (o *User) HasOtherMails() bool`

HasOtherMails returns a boolean if a field has been set.

### GetPasswordPolicies

`func (o *User) GetPasswordPolicies() string`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *User) GetPasswordPoliciesOk() (*string, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicies

`func (o *User) SetPasswordPolicies(v string)`

SetPasswordPolicies sets PasswordPolicies field to given value.

### HasPasswordPolicies

`func (o *User) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### SetPasswordPoliciesNil

`func (o *User) SetPasswordPoliciesNil(b bool)`

 SetPasswordPoliciesNil sets the value for PasswordPolicies to be an explicit nil

### UnsetPasswordPolicies
`func (o *User) UnsetPasswordPolicies()`

UnsetPasswordPolicies ensures that no value is present for PasswordPolicies, not even an explicit nil
### GetPasswordProfile

`func (o *User) GetPasswordProfile() AnyOfmicrosoftGraphPasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *User) GetPasswordProfileOk() (*AnyOfmicrosoftGraphPasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProfile

`func (o *User) SetPasswordProfile(v AnyOfmicrosoftGraphPasswordProfile)`

SetPasswordProfile sets PasswordProfile field to given value.

### HasPasswordProfile

`func (o *User) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### SetPasswordProfileNil

`func (o *User) SetPasswordProfileNil(b bool)`

 SetPasswordProfileNil sets the value for PasswordProfile to be an explicit nil

### UnsetPasswordProfile
`func (o *User) UnsetPasswordProfile()`

UnsetPasswordProfile ensures that no value is present for PasswordProfile, not even an explicit nil
### GetPostalCode

`func (o *User) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *User) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *User) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *User) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *User) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *User) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPreferredLanguage

`func (o *User) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *User) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *User) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *User) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *User) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *User) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetProvisionedPlans

`func (o *User) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *User) GetProvisionedPlansOk() (*[]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedPlans

`func (o *User) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans sets ProvisionedPlans field to given value.

### HasProvisionedPlans

`func (o *User) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### GetProxyAddresses

`func (o *User) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *User) GetProxyAddressesOk() (*[]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddresses

`func (o *User) SetProxyAddresses(v []string)`

SetProxyAddresses sets ProxyAddresses field to given value.

### HasProxyAddresses

`func (o *User) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### GetShowInAddressList

`func (o *User) GetShowInAddressList() bool`

GetShowInAddressList returns the ShowInAddressList field if non-nil, zero value otherwise.

### GetShowInAddressListOk

`func (o *User) GetShowInAddressListOk() (*bool, bool)`

GetShowInAddressListOk returns a tuple with the ShowInAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowInAddressList

`func (o *User) SetShowInAddressList(v bool)`

SetShowInAddressList sets ShowInAddressList field to given value.

### HasShowInAddressList

`func (o *User) HasShowInAddressList() bool`

HasShowInAddressList returns a boolean if a field has been set.

### SetShowInAddressListNil

`func (o *User) SetShowInAddressListNil(b bool)`

 SetShowInAddressListNil sets the value for ShowInAddressList to be an explicit nil

### UnsetShowInAddressList
`func (o *User) UnsetShowInAddressList()`

UnsetShowInAddressList ensures that no value is present for ShowInAddressList, not even an explicit nil
### GetSignInSessionsValidFromDateTime

`func (o *User) GetSignInSessionsValidFromDateTime() time.Time`

GetSignInSessionsValidFromDateTime returns the SignInSessionsValidFromDateTime field if non-nil, zero value otherwise.

### GetSignInSessionsValidFromDateTimeOk

`func (o *User) GetSignInSessionsValidFromDateTimeOk() (*time.Time, bool)`

GetSignInSessionsValidFromDateTimeOk returns a tuple with the SignInSessionsValidFromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInSessionsValidFromDateTime

`func (o *User) SetSignInSessionsValidFromDateTime(v time.Time)`

SetSignInSessionsValidFromDateTime sets SignInSessionsValidFromDateTime field to given value.

### HasSignInSessionsValidFromDateTime

`func (o *User) HasSignInSessionsValidFromDateTime() bool`

HasSignInSessionsValidFromDateTime returns a boolean if a field has been set.

### SetSignInSessionsValidFromDateTimeNil

`func (o *User) SetSignInSessionsValidFromDateTimeNil(b bool)`

 SetSignInSessionsValidFromDateTimeNil sets the value for SignInSessionsValidFromDateTime to be an explicit nil

### UnsetSignInSessionsValidFromDateTime
`func (o *User) UnsetSignInSessionsValidFromDateTime()`

UnsetSignInSessionsValidFromDateTime ensures that no value is present for SignInSessionsValidFromDateTime, not even an explicit nil
### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *User) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *User) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreetAddress

`func (o *User) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *User) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *User) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *User) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *User) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *User) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetSurname

`func (o *User) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *User) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *User) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *User) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *User) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *User) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetUsageLocation

`func (o *User) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *User) GetUsageLocationOk() (*string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLocation

`func (o *User) SetUsageLocation(v string)`

SetUsageLocation sets UsageLocation field to given value.

### HasUsageLocation

`func (o *User) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocationNil

`func (o *User) SetUsageLocationNil(b bool)`

 SetUsageLocationNil sets the value for UsageLocation to be an explicit nil

### UnsetUsageLocation
`func (o *User) UnsetUsageLocation()`

UnsetUsageLocation ensures that no value is present for UsageLocation, not even an explicit nil
### GetUserPrincipalName

`func (o *User) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *User) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *User) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *User) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *User) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *User) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetUserType

`func (o *User) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *User) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *User) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *User) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetMailboxSettings

`func (o *User) GetMailboxSettings() AnyOfmicrosoftGraphMailboxSettings`

GetMailboxSettings returns the MailboxSettings field if non-nil, zero value otherwise.

### GetMailboxSettingsOk

`func (o *User) GetMailboxSettingsOk() (*AnyOfmicrosoftGraphMailboxSettings, bool)`

GetMailboxSettingsOk returns a tuple with the MailboxSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxSettings

`func (o *User) SetMailboxSettings(v AnyOfmicrosoftGraphMailboxSettings)`

SetMailboxSettings sets MailboxSettings field to given value.

### HasMailboxSettings

`func (o *User) HasMailboxSettings() bool`

HasMailboxSettings returns a boolean if a field has been set.

### SetMailboxSettingsNil

`func (o *User) SetMailboxSettingsNil(b bool)`

 SetMailboxSettingsNil sets the value for MailboxSettings to be an explicit nil

### UnsetMailboxSettings
`func (o *User) UnsetMailboxSettings()`

UnsetMailboxSettings ensures that no value is present for MailboxSettings, not even an explicit nil
### GetDeviceEnrollmentLimit

`func (o *User) GetDeviceEnrollmentLimit() int32`

GetDeviceEnrollmentLimit returns the DeviceEnrollmentLimit field if non-nil, zero value otherwise.

### GetDeviceEnrollmentLimitOk

`func (o *User) GetDeviceEnrollmentLimitOk() (*int32, bool)`

GetDeviceEnrollmentLimitOk returns a tuple with the DeviceEnrollmentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentLimit

`func (o *User) SetDeviceEnrollmentLimit(v int32)`

SetDeviceEnrollmentLimit sets DeviceEnrollmentLimit field to given value.

### HasDeviceEnrollmentLimit

`func (o *User) HasDeviceEnrollmentLimit() bool`

HasDeviceEnrollmentLimit returns a boolean if a field has been set.

### GetAboutMe

`func (o *User) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *User) GetAboutMeOk() (*string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboutMe

`func (o *User) SetAboutMe(v string)`

SetAboutMe sets AboutMe field to given value.

### HasAboutMe

`func (o *User) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### SetAboutMeNil

`func (o *User) SetAboutMeNil(b bool)`

 SetAboutMeNil sets the value for AboutMe to be an explicit nil

### UnsetAboutMe
`func (o *User) UnsetAboutMe()`

UnsetAboutMe ensures that no value is present for AboutMe, not even an explicit nil
### GetBirthday

`func (o *User) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *User) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *User) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *User) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetHireDate

`func (o *User) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *User) GetHireDateOk() (*time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *User) SetHireDate(v time.Time)`

SetHireDate sets HireDate field to given value.

### HasHireDate

`func (o *User) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### GetInterests

`func (o *User) GetInterests() []*string`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *User) GetInterestsOk() (*[]*string, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *User) SetInterests(v []*string)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *User) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetMySite

`func (o *User) GetMySite() string`

GetMySite returns the MySite field if non-nil, zero value otherwise.

### GetMySiteOk

`func (o *User) GetMySiteOk() (*string, bool)`

GetMySiteOk returns a tuple with the MySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMySite

`func (o *User) SetMySite(v string)`

SetMySite sets MySite field to given value.

### HasMySite

`func (o *User) HasMySite() bool`

HasMySite returns a boolean if a field has been set.

### SetMySiteNil

`func (o *User) SetMySiteNil(b bool)`

 SetMySiteNil sets the value for MySite to be an explicit nil

### UnsetMySite
`func (o *User) UnsetMySite()`

UnsetMySite ensures that no value is present for MySite, not even an explicit nil
### GetPastProjects

`func (o *User) GetPastProjects() []*string`

GetPastProjects returns the PastProjects field if non-nil, zero value otherwise.

### GetPastProjectsOk

`func (o *User) GetPastProjectsOk() (*[]*string, bool)`

GetPastProjectsOk returns a tuple with the PastProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastProjects

`func (o *User) SetPastProjects(v []*string)`

SetPastProjects sets PastProjects field to given value.

### HasPastProjects

`func (o *User) HasPastProjects() bool`

HasPastProjects returns a boolean if a field has been set.

### GetPreferredName

`func (o *User) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *User) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *User) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *User) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredNameNil

`func (o *User) SetPreferredNameNil(b bool)`

 SetPreferredNameNil sets the value for PreferredName to be an explicit nil

### UnsetPreferredName
`func (o *User) UnsetPreferredName()`

UnsetPreferredName ensures that no value is present for PreferredName, not even an explicit nil
### GetResponsibilities

`func (o *User) GetResponsibilities() []*string`

GetResponsibilities returns the Responsibilities field if non-nil, zero value otherwise.

### GetResponsibilitiesOk

`func (o *User) GetResponsibilitiesOk() (*[]*string, bool)`

GetResponsibilitiesOk returns a tuple with the Responsibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibilities

`func (o *User) SetResponsibilities(v []*string)`

SetResponsibilities sets Responsibilities field to given value.

### HasResponsibilities

`func (o *User) HasResponsibilities() bool`

HasResponsibilities returns a boolean if a field has been set.

### GetSchools

`func (o *User) GetSchools() []*string`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *User) GetSchoolsOk() (*[]*string, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchools

`func (o *User) SetSchools(v []*string)`

SetSchools sets Schools field to given value.

### HasSchools

`func (o *User) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### GetSkills

`func (o *User) GetSkills() []*string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *User) GetSkillsOk() (*[]*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *User) SetSkills(v []*string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *User) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetAppRoleAssignments

`func (o *User) GetAppRoleAssignments() []MicrosoftGraphAppRoleAssignment`

GetAppRoleAssignments returns the AppRoleAssignments field if non-nil, zero value otherwise.

### GetAppRoleAssignmentsOk

`func (o *User) GetAppRoleAssignmentsOk() (*[]MicrosoftGraphAppRoleAssignment, bool)`

GetAppRoleAssignmentsOk returns a tuple with the AppRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignments

`func (o *User) SetAppRoleAssignments(v []MicrosoftGraphAppRoleAssignment)`

SetAppRoleAssignments sets AppRoleAssignments field to given value.

### HasAppRoleAssignments

`func (o *User) HasAppRoleAssignments() bool`

HasAppRoleAssignments returns a boolean if a field has been set.

### GetCreatedObjects

`func (o *User) GetCreatedObjects() []MicrosoftGraphDirectoryObject`

GetCreatedObjects returns the CreatedObjects field if non-nil, zero value otherwise.

### GetCreatedObjectsOk

`func (o *User) GetCreatedObjectsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetCreatedObjectsOk returns a tuple with the CreatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedObjects

`func (o *User) SetCreatedObjects(v []MicrosoftGraphDirectoryObject)`

SetCreatedObjects sets CreatedObjects field to given value.

### HasCreatedObjects

`func (o *User) HasCreatedObjects() bool`

HasCreatedObjects returns a boolean if a field has been set.

### GetDirectReports

`func (o *User) GetDirectReports() []MicrosoftGraphDirectoryObject`

GetDirectReports returns the DirectReports field if non-nil, zero value otherwise.

### GetDirectReportsOk

`func (o *User) GetDirectReportsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetDirectReportsOk returns a tuple with the DirectReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectReports

`func (o *User) SetDirectReports(v []MicrosoftGraphDirectoryObject)`

SetDirectReports sets DirectReports field to given value.

### HasDirectReports

`func (o *User) HasDirectReports() bool`

HasDirectReports returns a boolean if a field has been set.

### GetLicenseDetails

`func (o *User) GetLicenseDetails() []MicrosoftGraphLicenseDetails`

GetLicenseDetails returns the LicenseDetails field if non-nil, zero value otherwise.

### GetLicenseDetailsOk

`func (o *User) GetLicenseDetailsOk() (*[]MicrosoftGraphLicenseDetails, bool)`

GetLicenseDetailsOk returns a tuple with the LicenseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseDetails

`func (o *User) SetLicenseDetails(v []MicrosoftGraphLicenseDetails)`

SetLicenseDetails sets LicenseDetails field to given value.

### HasLicenseDetails

`func (o *User) HasLicenseDetails() bool`

HasLicenseDetails returns a boolean if a field has been set.

### GetManager

`func (o *User) GetManager() AnyOfmicrosoftGraphDirectoryObject`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *User) GetManagerOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *User) SetManager(v AnyOfmicrosoftGraphDirectoryObject)`

SetManager sets Manager field to given value.

### HasManager

`func (o *User) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *User) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *User) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetMemberOf

`func (o *User) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *User) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *User) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *User) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetOauth2PermissionGrants

`func (o *User) GetOauth2PermissionGrants() []MicrosoftGraphOAuth2PermissionGrant`

GetOauth2PermissionGrants returns the Oauth2PermissionGrants field if non-nil, zero value otherwise.

### GetOauth2PermissionGrantsOk

`func (o *User) GetOauth2PermissionGrantsOk() (*[]MicrosoftGraphOAuth2PermissionGrant, bool)`

GetOauth2PermissionGrantsOk returns a tuple with the Oauth2PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2PermissionGrants

`func (o *User) SetOauth2PermissionGrants(v []MicrosoftGraphOAuth2PermissionGrant)`

SetOauth2PermissionGrants sets Oauth2PermissionGrants field to given value.

### HasOauth2PermissionGrants

`func (o *User) HasOauth2PermissionGrants() bool`

HasOauth2PermissionGrants returns a boolean if a field has been set.

### GetOwnedDevices

`func (o *User) GetOwnedDevices() []MicrosoftGraphDirectoryObject`

GetOwnedDevices returns the OwnedDevices field if non-nil, zero value otherwise.

### GetOwnedDevicesOk

`func (o *User) GetOwnedDevicesOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnedDevicesOk returns a tuple with the OwnedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedDevices

`func (o *User) SetOwnedDevices(v []MicrosoftGraphDirectoryObject)`

SetOwnedDevices sets OwnedDevices field to given value.

### HasOwnedDevices

`func (o *User) HasOwnedDevices() bool`

HasOwnedDevices returns a boolean if a field has been set.

### GetOwnedObjects

`func (o *User) GetOwnedObjects() []MicrosoftGraphDirectoryObject`

GetOwnedObjects returns the OwnedObjects field if non-nil, zero value otherwise.

### GetOwnedObjectsOk

`func (o *User) GetOwnedObjectsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnedObjectsOk returns a tuple with the OwnedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedObjects

`func (o *User) SetOwnedObjects(v []MicrosoftGraphDirectoryObject)`

SetOwnedObjects sets OwnedObjects field to given value.

### HasOwnedObjects

`func (o *User) HasOwnedObjects() bool`

HasOwnedObjects returns a boolean if a field has been set.

### GetRegisteredDevices

`func (o *User) GetRegisteredDevices() []MicrosoftGraphDirectoryObject`

GetRegisteredDevices returns the RegisteredDevices field if non-nil, zero value otherwise.

### GetRegisteredDevicesOk

`func (o *User) GetRegisteredDevicesOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevices

`func (o *User) SetRegisteredDevices(v []MicrosoftGraphDirectoryObject)`

SetRegisteredDevices sets RegisteredDevices field to given value.

### HasRegisteredDevices

`func (o *User) HasRegisteredDevices() bool`

HasRegisteredDevices returns a boolean if a field has been set.

### GetScopedRoleMemberOf

`func (o *User) GetScopedRoleMemberOf() []MicrosoftGraphScopedRoleMembership`

GetScopedRoleMemberOf returns the ScopedRoleMemberOf field if non-nil, zero value otherwise.

### GetScopedRoleMemberOfOk

`func (o *User) GetScopedRoleMemberOfOk() (*[]MicrosoftGraphScopedRoleMembership, bool)`

GetScopedRoleMemberOfOk returns a tuple with the ScopedRoleMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedRoleMemberOf

`func (o *User) SetScopedRoleMemberOf(v []MicrosoftGraphScopedRoleMembership)`

SetScopedRoleMemberOf sets ScopedRoleMemberOf field to given value.

### HasScopedRoleMemberOf

`func (o *User) HasScopedRoleMemberOf() bool`

HasScopedRoleMemberOf returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *User) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *User) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *User) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *User) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### GetCalendar

`func (o *User) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *User) GetCalendarOk() (*AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *User) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *User) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendarNil

`func (o *User) SetCalendarNil(b bool)`

 SetCalendarNil sets the value for Calendar to be an explicit nil

### UnsetCalendar
`func (o *User) UnsetCalendar()`

UnsetCalendar ensures that no value is present for Calendar, not even an explicit nil
### GetCalendarGroups

`func (o *User) GetCalendarGroups() []MicrosoftGraphCalendarGroup`

GetCalendarGroups returns the CalendarGroups field if non-nil, zero value otherwise.

### GetCalendarGroupsOk

`func (o *User) GetCalendarGroupsOk() (*[]MicrosoftGraphCalendarGroup, bool)`

GetCalendarGroupsOk returns a tuple with the CalendarGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarGroups

`func (o *User) SetCalendarGroups(v []MicrosoftGraphCalendarGroup)`

SetCalendarGroups sets CalendarGroups field to given value.

### HasCalendarGroups

`func (o *User) HasCalendarGroups() bool`

HasCalendarGroups returns a boolean if a field has been set.

### GetCalendars

`func (o *User) GetCalendars() []MicrosoftGraphCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *User) GetCalendarsOk() (*[]MicrosoftGraphCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendars

`func (o *User) SetCalendars(v []MicrosoftGraphCalendar)`

SetCalendars sets Calendars field to given value.

### HasCalendars

`func (o *User) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### GetCalendarView

`func (o *User) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *User) GetCalendarViewOk() (*[]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarView

`func (o *User) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView sets CalendarView field to given value.

### HasCalendarView

`func (o *User) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### GetContactFolders

`func (o *User) GetContactFolders() []MicrosoftGraphContactFolder`

GetContactFolders returns the ContactFolders field if non-nil, zero value otherwise.

### GetContactFoldersOk

`func (o *User) GetContactFoldersOk() (*[]MicrosoftGraphContactFolder, bool)`

GetContactFoldersOk returns a tuple with the ContactFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFolders

`func (o *User) SetContactFolders(v []MicrosoftGraphContactFolder)`

SetContactFolders sets ContactFolders field to given value.

### HasContactFolders

`func (o *User) HasContactFolders() bool`

HasContactFolders returns a boolean if a field has been set.

### GetContacts

`func (o *User) GetContacts() []MicrosoftGraphContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *User) GetContactsOk() (*[]MicrosoftGraphContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *User) SetContacts(v []MicrosoftGraphContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *User) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetEvents

`func (o *User) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *User) GetEventsOk() (*[]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *User) SetEvents(v []MicrosoftGraphEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *User) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetInferenceClassification

`func (o *User) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassification`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *User) GetInferenceClassificationOk() (*AnyOfmicrosoftGraphInferenceClassification, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceClassification

`func (o *User) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassification)`

SetInferenceClassification sets InferenceClassification field to given value.

### HasInferenceClassification

`func (o *User) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassificationNil

`func (o *User) SetInferenceClassificationNil(b bool)`

 SetInferenceClassificationNil sets the value for InferenceClassification to be an explicit nil

### UnsetInferenceClassification
`func (o *User) UnsetInferenceClassification()`

UnsetInferenceClassification ensures that no value is present for InferenceClassification, not even an explicit nil
### GetMailFolders

`func (o *User) GetMailFolders() []MicrosoftGraphMailFolder`

GetMailFolders returns the MailFolders field if non-nil, zero value otherwise.

### GetMailFoldersOk

`func (o *User) GetMailFoldersOk() (*[]MicrosoftGraphMailFolder, bool)`

GetMailFoldersOk returns a tuple with the MailFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailFolders

`func (o *User) SetMailFolders(v []MicrosoftGraphMailFolder)`

SetMailFolders sets MailFolders field to given value.

### HasMailFolders

`func (o *User) HasMailFolders() bool`

HasMailFolders returns a boolean if a field has been set.

### GetMessages

`func (o *User) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *User) GetMessagesOk() (*[]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *User) SetMessages(v []MicrosoftGraphMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *User) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetOutlook

`func (o *User) GetOutlook() AnyOfmicrosoftGraphOutlookUser`

GetOutlook returns the Outlook field if non-nil, zero value otherwise.

### GetOutlookOk

`func (o *User) GetOutlookOk() (*AnyOfmicrosoftGraphOutlookUser, bool)`

GetOutlookOk returns a tuple with the Outlook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlook

`func (o *User) SetOutlook(v AnyOfmicrosoftGraphOutlookUser)`

SetOutlook sets Outlook field to given value.

### HasOutlook

`func (o *User) HasOutlook() bool`

HasOutlook returns a boolean if a field has been set.

### SetOutlookNil

`func (o *User) SetOutlookNil(b bool)`

 SetOutlookNil sets the value for Outlook to be an explicit nil

### UnsetOutlook
`func (o *User) UnsetOutlook()`

UnsetOutlook ensures that no value is present for Outlook, not even an explicit nil
### GetPeople

`func (o *User) GetPeople() []MicrosoftGraphPerson`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *User) GetPeopleOk() (*[]MicrosoftGraphPerson, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *User) SetPeople(v []MicrosoftGraphPerson)`

SetPeople sets People field to given value.

### HasPeople

`func (o *User) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetDrive

`func (o *User) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *User) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *User) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *User) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *User) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *User) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetDrives

`func (o *User) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *User) GetDrivesOk() (*[]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *User) SetDrives(v []MicrosoftGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *User) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetFollowedSites

`func (o *User) GetFollowedSites() []MicrosoftGraphSite`

GetFollowedSites returns the FollowedSites field if non-nil, zero value otherwise.

### GetFollowedSitesOk

`func (o *User) GetFollowedSitesOk() (*[]MicrosoftGraphSite, bool)`

GetFollowedSitesOk returns a tuple with the FollowedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedSites

`func (o *User) SetFollowedSites(v []MicrosoftGraphSite)`

SetFollowedSites sets FollowedSites field to given value.

### HasFollowedSites

`func (o *User) HasFollowedSites() bool`

HasFollowedSites returns a boolean if a field has been set.

### GetExtensions

`func (o *User) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *User) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *User) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *User) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetAgreementAcceptances

`func (o *User) GetAgreementAcceptances() []MicrosoftGraphAgreementAcceptance`

GetAgreementAcceptances returns the AgreementAcceptances field if non-nil, zero value otherwise.

### GetAgreementAcceptancesOk

`func (o *User) GetAgreementAcceptancesOk() (*[]MicrosoftGraphAgreementAcceptance, bool)`

GetAgreementAcceptancesOk returns a tuple with the AgreementAcceptances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAcceptances

`func (o *User) SetAgreementAcceptances(v []MicrosoftGraphAgreementAcceptance)`

SetAgreementAcceptances sets AgreementAcceptances field to given value.

### HasAgreementAcceptances

`func (o *User) HasAgreementAcceptances() bool`

HasAgreementAcceptances returns a boolean if a field has been set.

### GetManagedDevices

`func (o *User) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *User) GetManagedDevicesOk() (*[]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevices

`func (o *User) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices sets ManagedDevices field to given value.

### HasManagedDevices

`func (o *User) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### GetManagedAppRegistrations

`func (o *User) GetManagedAppRegistrations() []MicrosoftGraphManagedAppRegistration`

GetManagedAppRegistrations returns the ManagedAppRegistrations field if non-nil, zero value otherwise.

### GetManagedAppRegistrationsOk

`func (o *User) GetManagedAppRegistrationsOk() (*[]MicrosoftGraphManagedAppRegistration, bool)`

GetManagedAppRegistrationsOk returns a tuple with the ManagedAppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppRegistrations

`func (o *User) SetManagedAppRegistrations(v []MicrosoftGraphManagedAppRegistration)`

SetManagedAppRegistrations sets ManagedAppRegistrations field to given value.

### HasManagedAppRegistrations

`func (o *User) HasManagedAppRegistrations() bool`

HasManagedAppRegistrations returns a boolean if a field has been set.

### GetDeviceManagementTroubleshootingEvents

`func (o *User) GetDeviceManagementTroubleshootingEvents() []MicrosoftGraphDeviceManagementTroubleshootingEvent`

GetDeviceManagementTroubleshootingEvents returns the DeviceManagementTroubleshootingEvents field if non-nil, zero value otherwise.

### GetDeviceManagementTroubleshootingEventsOk

`func (o *User) GetDeviceManagementTroubleshootingEventsOk() (*[]MicrosoftGraphDeviceManagementTroubleshootingEvent, bool)`

GetDeviceManagementTroubleshootingEventsOk returns a tuple with the DeviceManagementTroubleshootingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementTroubleshootingEvents

`func (o *User) SetDeviceManagementTroubleshootingEvents(v []MicrosoftGraphDeviceManagementTroubleshootingEvent)`

SetDeviceManagementTroubleshootingEvents sets DeviceManagementTroubleshootingEvents field to given value.

### HasDeviceManagementTroubleshootingEvents

`func (o *User) HasDeviceManagementTroubleshootingEvents() bool`

HasDeviceManagementTroubleshootingEvents returns a boolean if a field has been set.

### GetPlanner

`func (o *User) GetPlanner() AnyOfmicrosoftGraphPlannerUser`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *User) GetPlannerOk() (*AnyOfmicrosoftGraphPlannerUser, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanner

`func (o *User) SetPlanner(v AnyOfmicrosoftGraphPlannerUser)`

SetPlanner sets Planner field to given value.

### HasPlanner

`func (o *User) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlannerNil

`func (o *User) SetPlannerNil(b bool)`

 SetPlannerNil sets the value for Planner to be an explicit nil

### UnsetPlanner
`func (o *User) UnsetPlanner()`

UnsetPlanner ensures that no value is present for Planner, not even an explicit nil
### GetInsights

`func (o *User) GetInsights() AnyOfmicrosoftGraphOfficeGraphInsights`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *User) GetInsightsOk() (*AnyOfmicrosoftGraphOfficeGraphInsights, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *User) SetInsights(v AnyOfmicrosoftGraphOfficeGraphInsights)`

SetInsights sets Insights field to given value.

### HasInsights

`func (o *User) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### SetInsightsNil

`func (o *User) SetInsightsNil(b bool)`

 SetInsightsNil sets the value for Insights to be an explicit nil

### UnsetInsights
`func (o *User) UnsetInsights()`

UnsetInsights ensures that no value is present for Insights, not even an explicit nil
### GetSettings

`func (o *User) GetSettings() AnyOfmicrosoftGraphUserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *User) GetSettingsOk() (*AnyOfmicrosoftGraphUserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *User) SetSettings(v AnyOfmicrosoftGraphUserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *User) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *User) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *User) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetOnenote

`func (o *User) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *User) GetOnenoteOk() (*AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnenote

`func (o *User) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote sets Onenote field to given value.

### HasOnenote

`func (o *User) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenoteNil

`func (o *User) SetOnenoteNil(b bool)`

 SetOnenoteNil sets the value for Onenote to be an explicit nil

### UnsetOnenote
`func (o *User) UnsetOnenote()`

UnsetOnenote ensures that no value is present for Onenote, not even an explicit nil
### GetPhoto

`func (o *User) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *User) GetPhotoOk() (*AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *User) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *User) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *User) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *User) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPhotos

`func (o *User) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *User) GetPhotosOk() (*[]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *User) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *User) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetActivities

`func (o *User) GetActivities() []MicrosoftGraphUserActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *User) GetActivitiesOk() (*[]MicrosoftGraphUserActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *User) SetActivities(v []MicrosoftGraphUserActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *User) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetOnlineMeetings

`func (o *User) GetOnlineMeetings() []MicrosoftGraphOnlineMeeting`

GetOnlineMeetings returns the OnlineMeetings field if non-nil, zero value otherwise.

### GetOnlineMeetingsOk

`func (o *User) GetOnlineMeetingsOk() (*[]MicrosoftGraphOnlineMeeting, bool)`

GetOnlineMeetingsOk returns a tuple with the OnlineMeetings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMeetings

`func (o *User) SetOnlineMeetings(v []MicrosoftGraphOnlineMeeting)`

SetOnlineMeetings sets OnlineMeetings field to given value.

### HasOnlineMeetings

`func (o *User) HasOnlineMeetings() bool`

HasOnlineMeetings returns a boolean if a field has been set.

### GetPresence

`func (o *User) GetPresence() AnyOfmicrosoftGraphPresence`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *User) GetPresenceOk() (*AnyOfmicrosoftGraphPresence, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *User) SetPresence(v AnyOfmicrosoftGraphPresence)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *User) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### SetPresenceNil

`func (o *User) SetPresenceNil(b bool)`

 SetPresenceNil sets the value for Presence to be an explicit nil

### UnsetPresence
`func (o *User) UnsetPresence()`

UnsetPresence ensures that no value is present for Presence, not even an explicit nil
### GetAuthentication

`func (o *User) GetAuthentication() AnyOfmicrosoftGraphAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *User) GetAuthenticationOk() (*AnyOfmicrosoftGraphAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *User) SetAuthentication(v AnyOfmicrosoftGraphAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *User) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *User) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *User) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil
### GetChats

`func (o *User) GetChats() []MicrosoftGraphChat`

GetChats returns the Chats field if non-nil, zero value otherwise.

### GetChatsOk

`func (o *User) GetChatsOk() (*[]MicrosoftGraphChat, bool)`

GetChatsOk returns a tuple with the Chats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChats

`func (o *User) SetChats(v []MicrosoftGraphChat)`

SetChats sets Chats field to given value.

### HasChats

`func (o *User) HasChats() bool`

HasChats returns a boolean if a field has been set.

### GetJoinedTeams

`func (o *User) GetJoinedTeams() []MicrosoftGraphTeam`

GetJoinedTeams returns the JoinedTeams field if non-nil, zero value otherwise.

### GetJoinedTeamsOk

`func (o *User) GetJoinedTeamsOk() (*[]MicrosoftGraphTeam, bool)`

GetJoinedTeamsOk returns a tuple with the JoinedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedTeams

`func (o *User) SetJoinedTeams(v []MicrosoftGraphTeam)`

SetJoinedTeams sets JoinedTeams field to given value.

### HasJoinedTeams

`func (o *User) HasJoinedTeams() bool`

HasJoinedTeams returns a boolean if a field has been set.

### GetTeamwork

`func (o *User) GetTeamwork() AnyOfmicrosoftGraphUserTeamwork`

GetTeamwork returns the Teamwork field if non-nil, zero value otherwise.

### GetTeamworkOk

`func (o *User) GetTeamworkOk() (*AnyOfmicrosoftGraphUserTeamwork, bool)`

GetTeamworkOk returns a tuple with the Teamwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamwork

`func (o *User) SetTeamwork(v AnyOfmicrosoftGraphUserTeamwork)`

SetTeamwork sets Teamwork field to given value.

### HasTeamwork

`func (o *User) HasTeamwork() bool`

HasTeamwork returns a boolean if a field has been set.

### SetTeamworkNil

`func (o *User) SetTeamworkNil(b bool)`

 SetTeamworkNil sets the value for Teamwork to be an explicit nil

### UnsetTeamwork
`func (o *User) UnsetTeamwork()`

UnsetTeamwork ensures that no value is present for Teamwork, not even an explicit nil
### GetTodo

`func (o *User) GetTodo() AnyOfmicrosoftGraphTodo`

GetTodo returns the Todo field if non-nil, zero value otherwise.

### GetTodoOk

`func (o *User) GetTodoOk() (*AnyOfmicrosoftGraphTodo, bool)`

GetTodoOk returns a tuple with the Todo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodo

`func (o *User) SetTodo(v AnyOfmicrosoftGraphTodo)`

SetTodo sets Todo field to given value.

### HasTodo

`func (o *User) HasTodo() bool`

HasTodo returns a boolean if a field has been set.

### SetTodoNil

`func (o *User) SetTodoNil(b bool)`

 SetTodoNil sets the value for Todo to be an explicit nil

### UnsetTodo
`func (o *User) UnsetTodo()`

UnsetTodo ensures that no value is present for Todo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


