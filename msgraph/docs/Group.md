# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedLabels** | Pointer to [**[]AnyOfmicrosoftGraphAssignedLabel**](AnyOfmicrosoftGraphAssignedLabel.md) | The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only. | [optional] 
**AssignedLicenses** | Pointer to [**[]AnyOfmicrosoftGraphAssignedLicense**](AnyOfmicrosoftGraphAssignedLicense.md) | The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only. | [optional] 
**Classification** | Pointer to **NullableString** | Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith). | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only. | [optional] 
**Description** | Pointer to **NullableString** | An optional description for the group. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith) and $search. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**ExpirationDateTime** | Pointer to **NullableTime** | Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only. | [optional] 
**GroupTypes** | Pointer to **[]string** | Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it&#39;s either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, not). | [optional] 
**HasMembersWithLicenseErrors** | Pointer to **NullableBool** | Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq). | [optional] 
**IsAssignableToRole** | Pointer to **NullableBool** | Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, not). | [optional] 
**LicenseProcessingState** | Pointer to [**AnyOfmicrosoftGraphLicenseProcessingState**](anyOf&lt;microsoft.graph.licenseProcessingState&gt;.md) | Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only. | [optional] 
**Mail** | Pointer to **NullableString** | The SMTP address for the group, for example, &#39;serviceadmins@contoso.onmicrosoft.com&#39;. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**MailEnabled** | Pointer to **NullableBool** | Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, not). | [optional] 
**MailNickname** | Pointer to **NullableString** | The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] &#39; ; : . &lt;&gt; , SPACE. Required. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**MembershipRule** | Pointer to **NullableString** | The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, not, ge, le, startsWith). | [optional] 
**MembershipRuleProcessingState** | Pointer to **NullableString** | Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, not, in). | [optional] 
**OnPremisesDomainName** | Pointer to **NullableString** | Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only. | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, not, ge, le, in). | [optional] 
**OnPremisesNetBiosName** | Pointer to **NullableString** | Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only. | [optional] 
**OnPremisesProvisioningErrors** | Pointer to [**[]AnyOfmicrosoftGraphOnPremisesProvisioningError**](AnyOfmicrosoftGraphOnPremisesProvisioningError.md) | Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, not). | [optional] 
**OnPremisesSamAccountName** | Pointer to **NullableString** | Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith). Read-only. | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **NullableString** | Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only. | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**PreferredDataLocation** | Pointer to **NullableString** | The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator&#39;s preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default. | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values). | [optional] 
**ProxyAddresses** | Pointer to **[]string** | Email addresses for the group that direct to the same group mailbox. For example: [&#39;SMTP: bob@contoso.com&#39;, &#39;smtp: bob@sales.contoso.com&#39;]. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**RenewedDateTime** | Pointer to **NullableTime** | Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, not, ge, le, in). Read-only. | [optional] 
**SecurityEnabled** | Pointer to **NullableBool** | Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, not, in). | [optional] 
**SecurityIdentifier** | Pointer to **NullableString** | Security identifier of the group, used in Windows scenarios. Returned by default. | [optional] 
**Theme** | Pointer to **NullableString** | Specifies a Microsoft 365 group&#39;s color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default. | [optional] 
**Visibility** | Pointer to **NullableString** | Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can&#39;t be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable. | [optional] 
**AllowExternalSenders** | Pointer to **NullableBool** | Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). | [optional] 
**AutoSubscribeNewMembers** | Pointer to **NullableBool** | Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). | [optional] 
**HideFromAddressLists** | Pointer to **NullableBool** | True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). | [optional] 
**HideFromOutlookClients** | Pointer to **NullableBool** | True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). | [optional] 
**IsSubscribedByMail** | Pointer to **NullableBool** | Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). | [optional] 
**UnseenCount** | Pointer to **NullableInt32** | Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). | [optional] 
**IsArchived** | Pointer to **NullableBool** |  | [optional] 
**AppRoleAssignments** | Pointer to [**[]MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | Represents the app roles a group has been granted for an application. Supports $expand. | [optional] 
**CreatedOnBehalfOf** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) | The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only. | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand. | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand. | [optional] 
**MembersWithLicenseErrors** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | A list of group members with license errors from this group-based license assignment. Read-only. | [optional] 
**Owners** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand. | [optional] 
**PermissionGrants** | Pointer to [**[]MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md) | The permission that has been granted for a group to a specific application. Supports $expand. | [optional] 
**Settings** | Pointer to [**[]MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md) | Read-only. Nullable. | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 
**TransitiveMembers** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) |  | [optional] 
**AcceptedSenders** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The list of users or groups that are allowed to create post&#39;s or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post. | [optional] 
**Calendar** | Pointer to [**AnyOfmicrosoftGraphCalendar**](anyOf&lt;microsoft.graph.calendar&gt;.md) | The group&#39;s calendar. Read-only. | [optional] 
**CalendarView** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The calendar view for the calendar. Read-only. | [optional] 
**Conversations** | Pointer to [**[]MicrosoftGraphConversation**](MicrosoftGraphConversation.md) | The group&#39;s conversations. | [optional] 
**Events** | Pointer to [**[]MicrosoftGraphEvent**](MicrosoftGraphEvent.md) | The group&#39;s calendar events. | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphProfilePhoto**](anyOf&lt;microsoft.graph.profilePhoto&gt;.md) | The group&#39;s profile photo | [optional] 
**Photos** | Pointer to [**[]MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md) | The profile photos owned by the group. Read-only. Nullable. | [optional] 
**RejectedSenders** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable | [optional] 
**Threads** | Pointer to [**[]MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md) | The group&#39;s conversation threads. Nullable. | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) | The group&#39;s default drive. Read-only. | [optional] 
**Drives** | Pointer to [**[]MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | The group&#39;s drives. Read-only. | [optional] 
**Sites** | Pointer to [**[]MicrosoftGraphSite**](MicrosoftGraphSite.md) | The list of SharePoint sites in this group. Access the default site with /sites/root. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the group. Read-only. Nullable. | [optional] 
**GroupLifecyclePolicies** | Pointer to [**[]MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md) | The collection of lifecycle policies for this group. Read-only. Nullable. | [optional] 
**Planner** | Pointer to [**AnyOfmicrosoftGraphPlannerGroup**](anyOf&lt;microsoft.graph.plannerGroup&gt;.md) | Entry-point to Planner resource that might exist for a Unified Group. | [optional] 
**Onenote** | Pointer to [**AnyOfmicrosoftGraphOnenote**](anyOf&lt;microsoft.graph.onenote&gt;.md) | Read-only. | [optional] 
**Team** | Pointer to [**AnyOfmicrosoftGraphTeam**](anyOf&lt;microsoft.graph.team&gt;.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedLabels

`func (o *Group) GetAssignedLabels() []*AnyOfmicrosoftGraphAssignedLabel`

GetAssignedLabels returns the AssignedLabels field if non-nil, zero value otherwise.

### GetAssignedLabelsOk

`func (o *Group) GetAssignedLabelsOk() (*[]*AnyOfmicrosoftGraphAssignedLabel, bool)`

GetAssignedLabelsOk returns a tuple with the AssignedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLabels

`func (o *Group) SetAssignedLabels(v []*AnyOfmicrosoftGraphAssignedLabel)`

SetAssignedLabels sets AssignedLabels field to given value.

### HasAssignedLabels

`func (o *Group) HasAssignedLabels() bool`

HasAssignedLabels returns a boolean if a field has been set.

### GetAssignedLicenses

`func (o *Group) GetAssignedLicenses() []*AnyOfmicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *Group) GetAssignedLicensesOk() (*[]*AnyOfmicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLicenses

`func (o *Group) SetAssignedLicenses(v []*AnyOfmicrosoftGraphAssignedLicense)`

SetAssignedLicenses sets AssignedLicenses field to given value.

### HasAssignedLicenses

`func (o *Group) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### GetClassification

`func (o *Group) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Group) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Group) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Group) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *Group) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *Group) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetCreatedDateTime

`func (o *Group) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Group) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Group) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Group) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Group) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Group) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Group) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Group) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Group) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Group) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Group) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExpirationDateTime

`func (o *Group) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *Group) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *Group) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *Group) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *Group) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *Group) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetGroupTypes

`func (o *Group) GetGroupTypes() []string`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *Group) GetGroupTypesOk() (*[]string, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *Group) SetGroupTypes(v []string)`

SetGroupTypes sets GroupTypes field to given value.

### HasGroupTypes

`func (o *Group) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### GetHasMembersWithLicenseErrors

`func (o *Group) GetHasMembersWithLicenseErrors() bool`

GetHasMembersWithLicenseErrors returns the HasMembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetHasMembersWithLicenseErrorsOk

`func (o *Group) GetHasMembersWithLicenseErrorsOk() (*bool, bool)`

GetHasMembersWithLicenseErrorsOk returns a tuple with the HasMembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMembersWithLicenseErrors

`func (o *Group) SetHasMembersWithLicenseErrors(v bool)`

SetHasMembersWithLicenseErrors sets HasMembersWithLicenseErrors field to given value.

### HasHasMembersWithLicenseErrors

`func (o *Group) HasHasMembersWithLicenseErrors() bool`

HasHasMembersWithLicenseErrors returns a boolean if a field has been set.

### SetHasMembersWithLicenseErrorsNil

`func (o *Group) SetHasMembersWithLicenseErrorsNil(b bool)`

 SetHasMembersWithLicenseErrorsNil sets the value for HasMembersWithLicenseErrors to be an explicit nil

### UnsetHasMembersWithLicenseErrors
`func (o *Group) UnsetHasMembersWithLicenseErrors()`

UnsetHasMembersWithLicenseErrors ensures that no value is present for HasMembersWithLicenseErrors, not even an explicit nil
### GetIsAssignableToRole

`func (o *Group) GetIsAssignableToRole() bool`

GetIsAssignableToRole returns the IsAssignableToRole field if non-nil, zero value otherwise.

### GetIsAssignableToRoleOk

`func (o *Group) GetIsAssignableToRoleOk() (*bool, bool)`

GetIsAssignableToRoleOk returns a tuple with the IsAssignableToRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssignableToRole

`func (o *Group) SetIsAssignableToRole(v bool)`

SetIsAssignableToRole sets IsAssignableToRole field to given value.

### HasIsAssignableToRole

`func (o *Group) HasIsAssignableToRole() bool`

HasIsAssignableToRole returns a boolean if a field has been set.

### SetIsAssignableToRoleNil

`func (o *Group) SetIsAssignableToRoleNil(b bool)`

 SetIsAssignableToRoleNil sets the value for IsAssignableToRole to be an explicit nil

### UnsetIsAssignableToRole
`func (o *Group) UnsetIsAssignableToRole()`

UnsetIsAssignableToRole ensures that no value is present for IsAssignableToRole, not even an explicit nil
### GetLicenseProcessingState

`func (o *Group) GetLicenseProcessingState() AnyOfmicrosoftGraphLicenseProcessingState`

GetLicenseProcessingState returns the LicenseProcessingState field if non-nil, zero value otherwise.

### GetLicenseProcessingStateOk

`func (o *Group) GetLicenseProcessingStateOk() (*AnyOfmicrosoftGraphLicenseProcessingState, bool)`

GetLicenseProcessingStateOk returns a tuple with the LicenseProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProcessingState

`func (o *Group) SetLicenseProcessingState(v AnyOfmicrosoftGraphLicenseProcessingState)`

SetLicenseProcessingState sets LicenseProcessingState field to given value.

### HasLicenseProcessingState

`func (o *Group) HasLicenseProcessingState() bool`

HasLicenseProcessingState returns a boolean if a field has been set.

### SetLicenseProcessingStateNil

`func (o *Group) SetLicenseProcessingStateNil(b bool)`

 SetLicenseProcessingStateNil sets the value for LicenseProcessingState to be an explicit nil

### UnsetLicenseProcessingState
`func (o *Group) UnsetLicenseProcessingState()`

UnsetLicenseProcessingState ensures that no value is present for LicenseProcessingState, not even an explicit nil
### GetMail

`func (o *Group) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *Group) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *Group) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *Group) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *Group) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *Group) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetMailEnabled

`func (o *Group) GetMailEnabled() bool`

GetMailEnabled returns the MailEnabled field if non-nil, zero value otherwise.

### GetMailEnabledOk

`func (o *Group) GetMailEnabledOk() (*bool, bool)`

GetMailEnabledOk returns a tuple with the MailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailEnabled

`func (o *Group) SetMailEnabled(v bool)`

SetMailEnabled sets MailEnabled field to given value.

### HasMailEnabled

`func (o *Group) HasMailEnabled() bool`

HasMailEnabled returns a boolean if a field has been set.

### SetMailEnabledNil

`func (o *Group) SetMailEnabledNil(b bool)`

 SetMailEnabledNil sets the value for MailEnabled to be an explicit nil

### UnsetMailEnabled
`func (o *Group) UnsetMailEnabled()`

UnsetMailEnabled ensures that no value is present for MailEnabled, not even an explicit nil
### GetMailNickname

`func (o *Group) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *Group) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *Group) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *Group) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *Group) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *Group) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetMembershipRule

`func (o *Group) GetMembershipRule() string`

GetMembershipRule returns the MembershipRule field if non-nil, zero value otherwise.

### GetMembershipRuleOk

`func (o *Group) GetMembershipRuleOk() (*string, bool)`

GetMembershipRuleOk returns a tuple with the MembershipRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipRule

`func (o *Group) SetMembershipRule(v string)`

SetMembershipRule sets MembershipRule field to given value.

### HasMembershipRule

`func (o *Group) HasMembershipRule() bool`

HasMembershipRule returns a boolean if a field has been set.

### SetMembershipRuleNil

`func (o *Group) SetMembershipRuleNil(b bool)`

 SetMembershipRuleNil sets the value for MembershipRule to be an explicit nil

### UnsetMembershipRule
`func (o *Group) UnsetMembershipRule()`

UnsetMembershipRule ensures that no value is present for MembershipRule, not even an explicit nil
### GetMembershipRuleProcessingState

`func (o *Group) GetMembershipRuleProcessingState() string`

GetMembershipRuleProcessingState returns the MembershipRuleProcessingState field if non-nil, zero value otherwise.

### GetMembershipRuleProcessingStateOk

`func (o *Group) GetMembershipRuleProcessingStateOk() (*string, bool)`

GetMembershipRuleProcessingStateOk returns a tuple with the MembershipRuleProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipRuleProcessingState

`func (o *Group) SetMembershipRuleProcessingState(v string)`

SetMembershipRuleProcessingState sets MembershipRuleProcessingState field to given value.

### HasMembershipRuleProcessingState

`func (o *Group) HasMembershipRuleProcessingState() bool`

HasMembershipRuleProcessingState returns a boolean if a field has been set.

### SetMembershipRuleProcessingStateNil

`func (o *Group) SetMembershipRuleProcessingStateNil(b bool)`

 SetMembershipRuleProcessingStateNil sets the value for MembershipRuleProcessingState to be an explicit nil

### UnsetMembershipRuleProcessingState
`func (o *Group) UnsetMembershipRuleProcessingState()`

UnsetMembershipRuleProcessingState ensures that no value is present for MembershipRuleProcessingState, not even an explicit nil
### GetOnPremisesDomainName

`func (o *Group) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *Group) GetOnPremisesDomainNameOk() (*string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDomainName

`func (o *Group) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName sets OnPremisesDomainName field to given value.

### HasOnPremisesDomainName

`func (o *Group) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainNameNil

`func (o *Group) SetOnPremisesDomainNameNil(b bool)`

 SetOnPremisesDomainNameNil sets the value for OnPremisesDomainName to be an explicit nil

### UnsetOnPremisesDomainName
`func (o *Group) UnsetOnPremisesDomainName()`

UnsetOnPremisesDomainName ensures that no value is present for OnPremisesDomainName, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *Group) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Group) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Group) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *Group) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *Group) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *Group) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesNetBiosName

`func (o *Group) GetOnPremisesNetBiosName() string`

GetOnPremisesNetBiosName returns the OnPremisesNetBiosName field if non-nil, zero value otherwise.

### GetOnPremisesNetBiosNameOk

`func (o *Group) GetOnPremisesNetBiosNameOk() (*string, bool)`

GetOnPremisesNetBiosNameOk returns a tuple with the OnPremisesNetBiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesNetBiosName

`func (o *Group) SetOnPremisesNetBiosName(v string)`

SetOnPremisesNetBiosName sets OnPremisesNetBiosName field to given value.

### HasOnPremisesNetBiosName

`func (o *Group) HasOnPremisesNetBiosName() bool`

HasOnPremisesNetBiosName returns a boolean if a field has been set.

### SetOnPremisesNetBiosNameNil

`func (o *Group) SetOnPremisesNetBiosNameNil(b bool)`

 SetOnPremisesNetBiosNameNil sets the value for OnPremisesNetBiosName to be an explicit nil

### UnsetOnPremisesNetBiosName
`func (o *Group) UnsetOnPremisesNetBiosName()`

UnsetOnPremisesNetBiosName ensures that no value is present for OnPremisesNetBiosName, not even an explicit nil
### GetOnPremisesProvisioningErrors

`func (o *Group) GetOnPremisesProvisioningErrors() []*AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *Group) GetOnPremisesProvisioningErrorsOk() (*[]*AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesProvisioningErrors

`func (o *Group) SetOnPremisesProvisioningErrors(v []*AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors sets OnPremisesProvisioningErrors field to given value.

### HasOnPremisesProvisioningErrors

`func (o *Group) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### GetOnPremisesSamAccountName

`func (o *Group) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *Group) GetOnPremisesSamAccountNameOk() (*string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSamAccountName

`func (o *Group) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName sets OnPremisesSamAccountName field to given value.

### HasOnPremisesSamAccountName

`func (o *Group) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountNameNil

`func (o *Group) SetOnPremisesSamAccountNameNil(b bool)`

 SetOnPremisesSamAccountNameNil sets the value for OnPremisesSamAccountName to be an explicit nil

### UnsetOnPremisesSamAccountName
`func (o *Group) UnsetOnPremisesSamAccountName()`

UnsetOnPremisesSamAccountName ensures that no value is present for OnPremisesSamAccountName, not even an explicit nil
### GetOnPremisesSecurityIdentifier

`func (o *Group) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *Group) GetOnPremisesSecurityIdentifierOk() (*string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSecurityIdentifier

`func (o *Group) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier sets OnPremisesSecurityIdentifier field to given value.

### HasOnPremisesSecurityIdentifier

`func (o *Group) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifierNil

`func (o *Group) SetOnPremisesSecurityIdentifierNil(b bool)`

 SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil

### UnsetOnPremisesSecurityIdentifier
`func (o *Group) UnsetOnPremisesSecurityIdentifier()`

UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *Group) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Group) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *Group) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *Group) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *Group) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *Group) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetPreferredDataLocation

`func (o *Group) GetPreferredDataLocation() string`

GetPreferredDataLocation returns the PreferredDataLocation field if non-nil, zero value otherwise.

### GetPreferredDataLocationOk

`func (o *Group) GetPreferredDataLocationOk() (*string, bool)`

GetPreferredDataLocationOk returns a tuple with the PreferredDataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDataLocation

`func (o *Group) SetPreferredDataLocation(v string)`

SetPreferredDataLocation sets PreferredDataLocation field to given value.

### HasPreferredDataLocation

`func (o *Group) HasPreferredDataLocation() bool`

HasPreferredDataLocation returns a boolean if a field has been set.

### SetPreferredDataLocationNil

`func (o *Group) SetPreferredDataLocationNil(b bool)`

 SetPreferredDataLocationNil sets the value for PreferredDataLocation to be an explicit nil

### UnsetPreferredDataLocation
`func (o *Group) UnsetPreferredDataLocation()`

UnsetPreferredDataLocation ensures that no value is present for PreferredDataLocation, not even an explicit nil
### GetPreferredLanguage

`func (o *Group) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *Group) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *Group) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *Group) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *Group) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *Group) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetProxyAddresses

`func (o *Group) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *Group) GetProxyAddressesOk() (*[]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddresses

`func (o *Group) SetProxyAddresses(v []string)`

SetProxyAddresses sets ProxyAddresses field to given value.

### HasProxyAddresses

`func (o *Group) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### GetRenewedDateTime

`func (o *Group) GetRenewedDateTime() time.Time`

GetRenewedDateTime returns the RenewedDateTime field if non-nil, zero value otherwise.

### GetRenewedDateTimeOk

`func (o *Group) GetRenewedDateTimeOk() (*time.Time, bool)`

GetRenewedDateTimeOk returns a tuple with the RenewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewedDateTime

`func (o *Group) SetRenewedDateTime(v time.Time)`

SetRenewedDateTime sets RenewedDateTime field to given value.

### HasRenewedDateTime

`func (o *Group) HasRenewedDateTime() bool`

HasRenewedDateTime returns a boolean if a field has been set.

### SetRenewedDateTimeNil

`func (o *Group) SetRenewedDateTimeNil(b bool)`

 SetRenewedDateTimeNil sets the value for RenewedDateTime to be an explicit nil

### UnsetRenewedDateTime
`func (o *Group) UnsetRenewedDateTime()`

UnsetRenewedDateTime ensures that no value is present for RenewedDateTime, not even an explicit nil
### GetSecurityEnabled

`func (o *Group) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *Group) GetSecurityEnabledOk() (*bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEnabled

`func (o *Group) SetSecurityEnabled(v bool)`

SetSecurityEnabled sets SecurityEnabled field to given value.

### HasSecurityEnabled

`func (o *Group) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### SetSecurityEnabledNil

`func (o *Group) SetSecurityEnabledNil(b bool)`

 SetSecurityEnabledNil sets the value for SecurityEnabled to be an explicit nil

### UnsetSecurityEnabled
`func (o *Group) UnsetSecurityEnabled()`

UnsetSecurityEnabled ensures that no value is present for SecurityEnabled, not even an explicit nil
### GetSecurityIdentifier

`func (o *Group) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *Group) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *Group) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *Group) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### SetSecurityIdentifierNil

`func (o *Group) SetSecurityIdentifierNil(b bool)`

 SetSecurityIdentifierNil sets the value for SecurityIdentifier to be an explicit nil

### UnsetSecurityIdentifier
`func (o *Group) UnsetSecurityIdentifier()`

UnsetSecurityIdentifier ensures that no value is present for SecurityIdentifier, not even an explicit nil
### GetTheme

`func (o *Group) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Group) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Group) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Group) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *Group) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *Group) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil
### GetVisibility

`func (o *Group) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Group) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Group) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Group) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *Group) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *Group) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetAllowExternalSenders

`func (o *Group) GetAllowExternalSenders() bool`

GetAllowExternalSenders returns the AllowExternalSenders field if non-nil, zero value otherwise.

### GetAllowExternalSendersOk

`func (o *Group) GetAllowExternalSendersOk() (*bool, bool)`

GetAllowExternalSendersOk returns a tuple with the AllowExternalSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExternalSenders

`func (o *Group) SetAllowExternalSenders(v bool)`

SetAllowExternalSenders sets AllowExternalSenders field to given value.

### HasAllowExternalSenders

`func (o *Group) HasAllowExternalSenders() bool`

HasAllowExternalSenders returns a boolean if a field has been set.

### SetAllowExternalSendersNil

`func (o *Group) SetAllowExternalSendersNil(b bool)`

 SetAllowExternalSendersNil sets the value for AllowExternalSenders to be an explicit nil

### UnsetAllowExternalSenders
`func (o *Group) UnsetAllowExternalSenders()`

UnsetAllowExternalSenders ensures that no value is present for AllowExternalSenders, not even an explicit nil
### GetAutoSubscribeNewMembers

`func (o *Group) GetAutoSubscribeNewMembers() bool`

GetAutoSubscribeNewMembers returns the AutoSubscribeNewMembers field if non-nil, zero value otherwise.

### GetAutoSubscribeNewMembersOk

`func (o *Group) GetAutoSubscribeNewMembersOk() (*bool, bool)`

GetAutoSubscribeNewMembersOk returns a tuple with the AutoSubscribeNewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubscribeNewMembers

`func (o *Group) SetAutoSubscribeNewMembers(v bool)`

SetAutoSubscribeNewMembers sets AutoSubscribeNewMembers field to given value.

### HasAutoSubscribeNewMembers

`func (o *Group) HasAutoSubscribeNewMembers() bool`

HasAutoSubscribeNewMembers returns a boolean if a field has been set.

### SetAutoSubscribeNewMembersNil

`func (o *Group) SetAutoSubscribeNewMembersNil(b bool)`

 SetAutoSubscribeNewMembersNil sets the value for AutoSubscribeNewMembers to be an explicit nil

### UnsetAutoSubscribeNewMembers
`func (o *Group) UnsetAutoSubscribeNewMembers()`

UnsetAutoSubscribeNewMembers ensures that no value is present for AutoSubscribeNewMembers, not even an explicit nil
### GetHideFromAddressLists

`func (o *Group) GetHideFromAddressLists() bool`

GetHideFromAddressLists returns the HideFromAddressLists field if non-nil, zero value otherwise.

### GetHideFromAddressListsOk

`func (o *Group) GetHideFromAddressListsOk() (*bool, bool)`

GetHideFromAddressListsOk returns a tuple with the HideFromAddressLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromAddressLists

`func (o *Group) SetHideFromAddressLists(v bool)`

SetHideFromAddressLists sets HideFromAddressLists field to given value.

### HasHideFromAddressLists

`func (o *Group) HasHideFromAddressLists() bool`

HasHideFromAddressLists returns a boolean if a field has been set.

### SetHideFromAddressListsNil

`func (o *Group) SetHideFromAddressListsNil(b bool)`

 SetHideFromAddressListsNil sets the value for HideFromAddressLists to be an explicit nil

### UnsetHideFromAddressLists
`func (o *Group) UnsetHideFromAddressLists()`

UnsetHideFromAddressLists ensures that no value is present for HideFromAddressLists, not even an explicit nil
### GetHideFromOutlookClients

`func (o *Group) GetHideFromOutlookClients() bool`

GetHideFromOutlookClients returns the HideFromOutlookClients field if non-nil, zero value otherwise.

### GetHideFromOutlookClientsOk

`func (o *Group) GetHideFromOutlookClientsOk() (*bool, bool)`

GetHideFromOutlookClientsOk returns a tuple with the HideFromOutlookClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromOutlookClients

`func (o *Group) SetHideFromOutlookClients(v bool)`

SetHideFromOutlookClients sets HideFromOutlookClients field to given value.

### HasHideFromOutlookClients

`func (o *Group) HasHideFromOutlookClients() bool`

HasHideFromOutlookClients returns a boolean if a field has been set.

### SetHideFromOutlookClientsNil

`func (o *Group) SetHideFromOutlookClientsNil(b bool)`

 SetHideFromOutlookClientsNil sets the value for HideFromOutlookClients to be an explicit nil

### UnsetHideFromOutlookClients
`func (o *Group) UnsetHideFromOutlookClients()`

UnsetHideFromOutlookClients ensures that no value is present for HideFromOutlookClients, not even an explicit nil
### GetIsSubscribedByMail

`func (o *Group) GetIsSubscribedByMail() bool`

GetIsSubscribedByMail returns the IsSubscribedByMail field if non-nil, zero value otherwise.

### GetIsSubscribedByMailOk

`func (o *Group) GetIsSubscribedByMailOk() (*bool, bool)`

GetIsSubscribedByMailOk returns a tuple with the IsSubscribedByMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedByMail

`func (o *Group) SetIsSubscribedByMail(v bool)`

SetIsSubscribedByMail sets IsSubscribedByMail field to given value.

### HasIsSubscribedByMail

`func (o *Group) HasIsSubscribedByMail() bool`

HasIsSubscribedByMail returns a boolean if a field has been set.

### SetIsSubscribedByMailNil

`func (o *Group) SetIsSubscribedByMailNil(b bool)`

 SetIsSubscribedByMailNil sets the value for IsSubscribedByMail to be an explicit nil

### UnsetIsSubscribedByMail
`func (o *Group) UnsetIsSubscribedByMail()`

UnsetIsSubscribedByMail ensures that no value is present for IsSubscribedByMail, not even an explicit nil
### GetUnseenCount

`func (o *Group) GetUnseenCount() int32`

GetUnseenCount returns the UnseenCount field if non-nil, zero value otherwise.

### GetUnseenCountOk

`func (o *Group) GetUnseenCountOk() (*int32, bool)`

GetUnseenCountOk returns a tuple with the UnseenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseenCount

`func (o *Group) SetUnseenCount(v int32)`

SetUnseenCount sets UnseenCount field to given value.

### HasUnseenCount

`func (o *Group) HasUnseenCount() bool`

HasUnseenCount returns a boolean if a field has been set.

### SetUnseenCountNil

`func (o *Group) SetUnseenCountNil(b bool)`

 SetUnseenCountNil sets the value for UnseenCount to be an explicit nil

### UnsetUnseenCount
`func (o *Group) UnsetUnseenCount()`

UnsetUnseenCount ensures that no value is present for UnseenCount, not even an explicit nil
### GetIsArchived

`func (o *Group) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Group) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Group) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Group) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *Group) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *Group) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetAppRoleAssignments

`func (o *Group) GetAppRoleAssignments() []MicrosoftGraphAppRoleAssignment`

GetAppRoleAssignments returns the AppRoleAssignments field if non-nil, zero value otherwise.

### GetAppRoleAssignmentsOk

`func (o *Group) GetAppRoleAssignmentsOk() (*[]MicrosoftGraphAppRoleAssignment, bool)`

GetAppRoleAssignmentsOk returns a tuple with the AppRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignments

`func (o *Group) SetAppRoleAssignments(v []MicrosoftGraphAppRoleAssignment)`

SetAppRoleAssignments sets AppRoleAssignments field to given value.

### HasAppRoleAssignments

`func (o *Group) HasAppRoleAssignments() bool`

HasAppRoleAssignments returns a boolean if a field has been set.

### GetCreatedOnBehalfOf

`func (o *Group) GetCreatedOnBehalfOf() AnyOfmicrosoftGraphDirectoryObject`

GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field if non-nil, zero value otherwise.

### GetCreatedOnBehalfOfOk

`func (o *Group) GetCreatedOnBehalfOfOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOnBehalfOf

`func (o *Group) SetCreatedOnBehalfOf(v AnyOfmicrosoftGraphDirectoryObject)`

SetCreatedOnBehalfOf sets CreatedOnBehalfOf field to given value.

### HasCreatedOnBehalfOf

`func (o *Group) HasCreatedOnBehalfOf() bool`

HasCreatedOnBehalfOf returns a boolean if a field has been set.

### SetCreatedOnBehalfOfNil

`func (o *Group) SetCreatedOnBehalfOfNil(b bool)`

 SetCreatedOnBehalfOfNil sets the value for CreatedOnBehalfOf to be an explicit nil

### UnsetCreatedOnBehalfOf
`func (o *Group) UnsetCreatedOnBehalfOf()`

UnsetCreatedOnBehalfOf ensures that no value is present for CreatedOnBehalfOf, not even an explicit nil
### GetMemberOf

`func (o *Group) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *Group) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *Group) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *Group) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetMembers

`func (o *Group) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Group) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Group) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMembersWithLicenseErrors

`func (o *Group) GetMembersWithLicenseErrors() []MicrosoftGraphDirectoryObject`

GetMembersWithLicenseErrors returns the MembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetMembersWithLicenseErrorsOk

`func (o *Group) GetMembersWithLicenseErrorsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersWithLicenseErrorsOk returns a tuple with the MembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithLicenseErrors

`func (o *Group) SetMembersWithLicenseErrors(v []MicrosoftGraphDirectoryObject)`

SetMembersWithLicenseErrors sets MembersWithLicenseErrors field to given value.

### HasMembersWithLicenseErrors

`func (o *Group) HasMembersWithLicenseErrors() bool`

HasMembersWithLicenseErrors returns a boolean if a field has been set.

### GetOwners

`func (o *Group) GetOwners() []MicrosoftGraphDirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Group) GetOwnersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *Group) SetOwners(v []MicrosoftGraphDirectoryObject)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *Group) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetPermissionGrants

`func (o *Group) GetPermissionGrants() []MicrosoftGraphResourceSpecificPermissionGrant`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *Group) GetPermissionGrantsOk() (*[]MicrosoftGraphResourceSpecificPermissionGrant, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *Group) SetPermissionGrants(v []MicrosoftGraphResourceSpecificPermissionGrant)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *Group) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.

### GetSettings

`func (o *Group) GetSettings() []MicrosoftGraphGroupSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Group) GetSettingsOk() (*[]MicrosoftGraphGroupSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Group) SetSettings(v []MicrosoftGraphGroupSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Group) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *Group) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *Group) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *Group) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *Group) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### GetTransitiveMembers

`func (o *Group) GetTransitiveMembers() []MicrosoftGraphDirectoryObject`

GetTransitiveMembers returns the TransitiveMembers field if non-nil, zero value otherwise.

### GetTransitiveMembersOk

`func (o *Group) GetTransitiveMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMembersOk returns a tuple with the TransitiveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembers

`func (o *Group) SetTransitiveMembers(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMembers sets TransitiveMembers field to given value.

### HasTransitiveMembers

`func (o *Group) HasTransitiveMembers() bool`

HasTransitiveMembers returns a boolean if a field has been set.

### GetAcceptedSenders

`func (o *Group) GetAcceptedSenders() []MicrosoftGraphDirectoryObject`

GetAcceptedSenders returns the AcceptedSenders field if non-nil, zero value otherwise.

### GetAcceptedSendersOk

`func (o *Group) GetAcceptedSendersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAcceptedSendersOk returns a tuple with the AcceptedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedSenders

`func (o *Group) SetAcceptedSenders(v []MicrosoftGraphDirectoryObject)`

SetAcceptedSenders sets AcceptedSenders field to given value.

### HasAcceptedSenders

`func (o *Group) HasAcceptedSenders() bool`

HasAcceptedSenders returns a boolean if a field has been set.

### GetCalendar

`func (o *Group) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Group) GetCalendarOk() (*AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *Group) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *Group) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendarNil

`func (o *Group) SetCalendarNil(b bool)`

 SetCalendarNil sets the value for Calendar to be an explicit nil

### UnsetCalendar
`func (o *Group) UnsetCalendar()`

UnsetCalendar ensures that no value is present for Calendar, not even an explicit nil
### GetCalendarView

`func (o *Group) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *Group) GetCalendarViewOk() (*[]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarView

`func (o *Group) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView sets CalendarView field to given value.

### HasCalendarView

`func (o *Group) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### GetConversations

`func (o *Group) GetConversations() []MicrosoftGraphConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *Group) GetConversationsOk() (*[]MicrosoftGraphConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *Group) SetConversations(v []MicrosoftGraphConversation)`

SetConversations sets Conversations field to given value.

### HasConversations

`func (o *Group) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### GetEvents

`func (o *Group) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Group) GetEventsOk() (*[]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Group) SetEvents(v []MicrosoftGraphEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Group) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetPhoto

`func (o *Group) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Group) GetPhotoOk() (*AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Group) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *Group) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *Group) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *Group) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPhotos

`func (o *Group) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *Group) GetPhotosOk() (*[]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *Group) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *Group) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetRejectedSenders

`func (o *Group) GetRejectedSenders() []MicrosoftGraphDirectoryObject`

GetRejectedSenders returns the RejectedSenders field if non-nil, zero value otherwise.

### GetRejectedSendersOk

`func (o *Group) GetRejectedSendersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRejectedSendersOk returns a tuple with the RejectedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedSenders

`func (o *Group) SetRejectedSenders(v []MicrosoftGraphDirectoryObject)`

SetRejectedSenders sets RejectedSenders field to given value.

### HasRejectedSenders

`func (o *Group) HasRejectedSenders() bool`

HasRejectedSenders returns a boolean if a field has been set.

### GetThreads

`func (o *Group) GetThreads() []MicrosoftGraphConversationThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *Group) GetThreadsOk() (*[]MicrosoftGraphConversationThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *Group) SetThreads(v []MicrosoftGraphConversationThread)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *Group) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetDrive

`func (o *Group) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *Group) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *Group) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *Group) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *Group) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *Group) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetDrives

`func (o *Group) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *Group) GetDrivesOk() (*[]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *Group) SetDrives(v []MicrosoftGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *Group) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetSites

`func (o *Group) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Group) GetSitesOk() (*[]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *Group) SetSites(v []MicrosoftGraphSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *Group) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetExtensions

`func (o *Group) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Group) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Group) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Group) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGroupLifecyclePolicies

`func (o *Group) GetGroupLifecyclePolicies() []MicrosoftGraphGroupLifecyclePolicy`

GetGroupLifecyclePolicies returns the GroupLifecyclePolicies field if non-nil, zero value otherwise.

### GetGroupLifecyclePoliciesOk

`func (o *Group) GetGroupLifecyclePoliciesOk() (*[]MicrosoftGraphGroupLifecyclePolicy, bool)`

GetGroupLifecyclePoliciesOk returns a tuple with the GroupLifecyclePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLifecyclePolicies

`func (o *Group) SetGroupLifecyclePolicies(v []MicrosoftGraphGroupLifecyclePolicy)`

SetGroupLifecyclePolicies sets GroupLifecyclePolicies field to given value.

### HasGroupLifecyclePolicies

`func (o *Group) HasGroupLifecyclePolicies() bool`

HasGroupLifecyclePolicies returns a boolean if a field has been set.

### GetPlanner

`func (o *Group) GetPlanner() AnyOfmicrosoftGraphPlannerGroup`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *Group) GetPlannerOk() (*AnyOfmicrosoftGraphPlannerGroup, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanner

`func (o *Group) SetPlanner(v AnyOfmicrosoftGraphPlannerGroup)`

SetPlanner sets Planner field to given value.

### HasPlanner

`func (o *Group) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlannerNil

`func (o *Group) SetPlannerNil(b bool)`

 SetPlannerNil sets the value for Planner to be an explicit nil

### UnsetPlanner
`func (o *Group) UnsetPlanner()`

UnsetPlanner ensures that no value is present for Planner, not even an explicit nil
### GetOnenote

`func (o *Group) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *Group) GetOnenoteOk() (*AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnenote

`func (o *Group) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote sets Onenote field to given value.

### HasOnenote

`func (o *Group) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenoteNil

`func (o *Group) SetOnenoteNil(b bool)`

 SetOnenoteNil sets the value for Onenote to be an explicit nil

### UnsetOnenote
`func (o *Group) UnsetOnenote()`

UnsetOnenote ensures that no value is present for Onenote, not even an explicit nil
### GetTeam

`func (o *Group) GetTeam() AnyOfmicrosoftGraphTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Group) GetTeamOk() (*AnyOfmicrosoftGraphTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *Group) SetTeam(v AnyOfmicrosoftGraphTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *Group) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *Group) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *Group) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


