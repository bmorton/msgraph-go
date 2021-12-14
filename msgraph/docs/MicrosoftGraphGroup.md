# MicrosoftGraphGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
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

### NewMicrosoftGraphGroup

`func NewMicrosoftGraphGroup() *MicrosoftGraphGroup`

NewMicrosoftGraphGroup instantiates a new MicrosoftGraphGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphGroupWithDefaults

`func NewMicrosoftGraphGroupWithDefaults() *MicrosoftGraphGroup`

NewMicrosoftGraphGroupWithDefaults instantiates a new MicrosoftGraphGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphGroup) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphGroup) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphGroup) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphGroup) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphGroup) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphGroup) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetAssignedLabels

`func (o *MicrosoftGraphGroup) GetAssignedLabels() []*AnyOfmicrosoftGraphAssignedLabel`

GetAssignedLabels returns the AssignedLabels field if non-nil, zero value otherwise.

### GetAssignedLabelsOk

`func (o *MicrosoftGraphGroup) GetAssignedLabelsOk() (*[]*AnyOfmicrosoftGraphAssignedLabel, bool)`

GetAssignedLabelsOk returns a tuple with the AssignedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLabels

`func (o *MicrosoftGraphGroup) SetAssignedLabels(v []*AnyOfmicrosoftGraphAssignedLabel)`

SetAssignedLabels sets AssignedLabels field to given value.

### HasAssignedLabels

`func (o *MicrosoftGraphGroup) HasAssignedLabels() bool`

HasAssignedLabels returns a boolean if a field has been set.

### GetAssignedLicenses

`func (o *MicrosoftGraphGroup) GetAssignedLicenses() []*AnyOfmicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *MicrosoftGraphGroup) GetAssignedLicensesOk() (*[]*AnyOfmicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLicenses

`func (o *MicrosoftGraphGroup) SetAssignedLicenses(v []*AnyOfmicrosoftGraphAssignedLicense)`

SetAssignedLicenses sets AssignedLicenses field to given value.

### HasAssignedLicenses

`func (o *MicrosoftGraphGroup) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### GetClassification

`func (o *MicrosoftGraphGroup) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *MicrosoftGraphGroup) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *MicrosoftGraphGroup) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *MicrosoftGraphGroup) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *MicrosoftGraphGroup) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *MicrosoftGraphGroup) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphGroup) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphGroup) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphGroup) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphGroup) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphGroup) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphGroup) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExpirationDateTime

`func (o *MicrosoftGraphGroup) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphGroup) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphGroup) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphGroup) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *MicrosoftGraphGroup) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *MicrosoftGraphGroup) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetGroupTypes

`func (o *MicrosoftGraphGroup) GetGroupTypes() []string`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *MicrosoftGraphGroup) GetGroupTypesOk() (*[]string, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *MicrosoftGraphGroup) SetGroupTypes(v []string)`

SetGroupTypes sets GroupTypes field to given value.

### HasGroupTypes

`func (o *MicrosoftGraphGroup) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### GetHasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) GetHasMembersWithLicenseErrors() bool`

GetHasMembersWithLicenseErrors returns the HasMembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetHasMembersWithLicenseErrorsOk

`func (o *MicrosoftGraphGroup) GetHasMembersWithLicenseErrorsOk() (*bool, bool)`

GetHasMembersWithLicenseErrorsOk returns a tuple with the HasMembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) SetHasMembersWithLicenseErrors(v bool)`

SetHasMembersWithLicenseErrors sets HasMembersWithLicenseErrors field to given value.

### HasHasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) HasHasMembersWithLicenseErrors() bool`

HasHasMembersWithLicenseErrors returns a boolean if a field has been set.

### SetHasMembersWithLicenseErrorsNil

`func (o *MicrosoftGraphGroup) SetHasMembersWithLicenseErrorsNil(b bool)`

 SetHasMembersWithLicenseErrorsNil sets the value for HasMembersWithLicenseErrors to be an explicit nil

### UnsetHasMembersWithLicenseErrors
`func (o *MicrosoftGraphGroup) UnsetHasMembersWithLicenseErrors()`

UnsetHasMembersWithLicenseErrors ensures that no value is present for HasMembersWithLicenseErrors, not even an explicit nil
### GetIsAssignableToRole

`func (o *MicrosoftGraphGroup) GetIsAssignableToRole() bool`

GetIsAssignableToRole returns the IsAssignableToRole field if non-nil, zero value otherwise.

### GetIsAssignableToRoleOk

`func (o *MicrosoftGraphGroup) GetIsAssignableToRoleOk() (*bool, bool)`

GetIsAssignableToRoleOk returns a tuple with the IsAssignableToRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssignableToRole

`func (o *MicrosoftGraphGroup) SetIsAssignableToRole(v bool)`

SetIsAssignableToRole sets IsAssignableToRole field to given value.

### HasIsAssignableToRole

`func (o *MicrosoftGraphGroup) HasIsAssignableToRole() bool`

HasIsAssignableToRole returns a boolean if a field has been set.

### SetIsAssignableToRoleNil

`func (o *MicrosoftGraphGroup) SetIsAssignableToRoleNil(b bool)`

 SetIsAssignableToRoleNil sets the value for IsAssignableToRole to be an explicit nil

### UnsetIsAssignableToRole
`func (o *MicrosoftGraphGroup) UnsetIsAssignableToRole()`

UnsetIsAssignableToRole ensures that no value is present for IsAssignableToRole, not even an explicit nil
### GetLicenseProcessingState

`func (o *MicrosoftGraphGroup) GetLicenseProcessingState() AnyOfmicrosoftGraphLicenseProcessingState`

GetLicenseProcessingState returns the LicenseProcessingState field if non-nil, zero value otherwise.

### GetLicenseProcessingStateOk

`func (o *MicrosoftGraphGroup) GetLicenseProcessingStateOk() (*AnyOfmicrosoftGraphLicenseProcessingState, bool)`

GetLicenseProcessingStateOk returns a tuple with the LicenseProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseProcessingState

`func (o *MicrosoftGraphGroup) SetLicenseProcessingState(v AnyOfmicrosoftGraphLicenseProcessingState)`

SetLicenseProcessingState sets LicenseProcessingState field to given value.

### HasLicenseProcessingState

`func (o *MicrosoftGraphGroup) HasLicenseProcessingState() bool`

HasLicenseProcessingState returns a boolean if a field has been set.

### SetLicenseProcessingStateNil

`func (o *MicrosoftGraphGroup) SetLicenseProcessingStateNil(b bool)`

 SetLicenseProcessingStateNil sets the value for LicenseProcessingState to be an explicit nil

### UnsetLicenseProcessingState
`func (o *MicrosoftGraphGroup) UnsetLicenseProcessingState()`

UnsetLicenseProcessingState ensures that no value is present for LicenseProcessingState, not even an explicit nil
### GetMail

`func (o *MicrosoftGraphGroup) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *MicrosoftGraphGroup) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *MicrosoftGraphGroup) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *MicrosoftGraphGroup) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMailNil

`func (o *MicrosoftGraphGroup) SetMailNil(b bool)`

 SetMailNil sets the value for Mail to be an explicit nil

### UnsetMail
`func (o *MicrosoftGraphGroup) UnsetMail()`

UnsetMail ensures that no value is present for Mail, not even an explicit nil
### GetMailEnabled

`func (o *MicrosoftGraphGroup) GetMailEnabled() bool`

GetMailEnabled returns the MailEnabled field if non-nil, zero value otherwise.

### GetMailEnabledOk

`func (o *MicrosoftGraphGroup) GetMailEnabledOk() (*bool, bool)`

GetMailEnabledOk returns a tuple with the MailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailEnabled

`func (o *MicrosoftGraphGroup) SetMailEnabled(v bool)`

SetMailEnabled sets MailEnabled field to given value.

### HasMailEnabled

`func (o *MicrosoftGraphGroup) HasMailEnabled() bool`

HasMailEnabled returns a boolean if a field has been set.

### SetMailEnabledNil

`func (o *MicrosoftGraphGroup) SetMailEnabledNil(b bool)`

 SetMailEnabledNil sets the value for MailEnabled to be an explicit nil

### UnsetMailEnabled
`func (o *MicrosoftGraphGroup) UnsetMailEnabled()`

UnsetMailEnabled ensures that no value is present for MailEnabled, not even an explicit nil
### GetMailNickname

`func (o *MicrosoftGraphGroup) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *MicrosoftGraphGroup) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *MicrosoftGraphGroup) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *MicrosoftGraphGroup) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *MicrosoftGraphGroup) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *MicrosoftGraphGroup) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetMembershipRule

`func (o *MicrosoftGraphGroup) GetMembershipRule() string`

GetMembershipRule returns the MembershipRule field if non-nil, zero value otherwise.

### GetMembershipRuleOk

`func (o *MicrosoftGraphGroup) GetMembershipRuleOk() (*string, bool)`

GetMembershipRuleOk returns a tuple with the MembershipRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipRule

`func (o *MicrosoftGraphGroup) SetMembershipRule(v string)`

SetMembershipRule sets MembershipRule field to given value.

### HasMembershipRule

`func (o *MicrosoftGraphGroup) HasMembershipRule() bool`

HasMembershipRule returns a boolean if a field has been set.

### SetMembershipRuleNil

`func (o *MicrosoftGraphGroup) SetMembershipRuleNil(b bool)`

 SetMembershipRuleNil sets the value for MembershipRule to be an explicit nil

### UnsetMembershipRule
`func (o *MicrosoftGraphGroup) UnsetMembershipRule()`

UnsetMembershipRule ensures that no value is present for MembershipRule, not even an explicit nil
### GetMembershipRuleProcessingState

`func (o *MicrosoftGraphGroup) GetMembershipRuleProcessingState() string`

GetMembershipRuleProcessingState returns the MembershipRuleProcessingState field if non-nil, zero value otherwise.

### GetMembershipRuleProcessingStateOk

`func (o *MicrosoftGraphGroup) GetMembershipRuleProcessingStateOk() (*string, bool)`

GetMembershipRuleProcessingStateOk returns a tuple with the MembershipRuleProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipRuleProcessingState

`func (o *MicrosoftGraphGroup) SetMembershipRuleProcessingState(v string)`

SetMembershipRuleProcessingState sets MembershipRuleProcessingState field to given value.

### HasMembershipRuleProcessingState

`func (o *MicrosoftGraphGroup) HasMembershipRuleProcessingState() bool`

HasMembershipRuleProcessingState returns a boolean if a field has been set.

### SetMembershipRuleProcessingStateNil

`func (o *MicrosoftGraphGroup) SetMembershipRuleProcessingStateNil(b bool)`

 SetMembershipRuleProcessingStateNil sets the value for MembershipRuleProcessingState to be an explicit nil

### UnsetMembershipRuleProcessingState
`func (o *MicrosoftGraphGroup) UnsetMembershipRuleProcessingState()`

UnsetMembershipRuleProcessingState ensures that no value is present for MembershipRuleProcessingState, not even an explicit nil
### GetOnPremisesDomainName

`func (o *MicrosoftGraphGroup) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *MicrosoftGraphGroup) GetOnPremisesDomainNameOk() (*string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesDomainName

`func (o *MicrosoftGraphGroup) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName sets OnPremisesDomainName field to given value.

### HasOnPremisesDomainName

`func (o *MicrosoftGraphGroup) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainNameNil

`func (o *MicrosoftGraphGroup) SetOnPremisesDomainNameNil(b bool)`

 SetOnPremisesDomainNameNil sets the value for OnPremisesDomainName to be an explicit nil

### UnsetOnPremisesDomainName
`func (o *MicrosoftGraphGroup) UnsetOnPremisesDomainName()`

UnsetOnPremisesDomainName ensures that no value is present for OnPremisesDomainName, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphGroup) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphGroup) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphGroup) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphGroup) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *MicrosoftGraphGroup) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *MicrosoftGraphGroup) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesNetBiosName

`func (o *MicrosoftGraphGroup) GetOnPremisesNetBiosName() string`

GetOnPremisesNetBiosName returns the OnPremisesNetBiosName field if non-nil, zero value otherwise.

### GetOnPremisesNetBiosNameOk

`func (o *MicrosoftGraphGroup) GetOnPremisesNetBiosNameOk() (*string, bool)`

GetOnPremisesNetBiosNameOk returns a tuple with the OnPremisesNetBiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesNetBiosName

`func (o *MicrosoftGraphGroup) SetOnPremisesNetBiosName(v string)`

SetOnPremisesNetBiosName sets OnPremisesNetBiosName field to given value.

### HasOnPremisesNetBiosName

`func (o *MicrosoftGraphGroup) HasOnPremisesNetBiosName() bool`

HasOnPremisesNetBiosName returns a boolean if a field has been set.

### SetOnPremisesNetBiosNameNil

`func (o *MicrosoftGraphGroup) SetOnPremisesNetBiosNameNil(b bool)`

 SetOnPremisesNetBiosNameNil sets the value for OnPremisesNetBiosName to be an explicit nil

### UnsetOnPremisesNetBiosName
`func (o *MicrosoftGraphGroup) UnsetOnPremisesNetBiosName()`

UnsetOnPremisesNetBiosName ensures that no value is present for OnPremisesNetBiosName, not even an explicit nil
### GetOnPremisesProvisioningErrors

`func (o *MicrosoftGraphGroup) GetOnPremisesProvisioningErrors() []*AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *MicrosoftGraphGroup) GetOnPremisesProvisioningErrorsOk() (*[]*AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesProvisioningErrors

`func (o *MicrosoftGraphGroup) SetOnPremisesProvisioningErrors(v []*AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors sets OnPremisesProvisioningErrors field to given value.

### HasOnPremisesProvisioningErrors

`func (o *MicrosoftGraphGroup) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### GetOnPremisesSamAccountName

`func (o *MicrosoftGraphGroup) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *MicrosoftGraphGroup) GetOnPremisesSamAccountNameOk() (*string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSamAccountName

`func (o *MicrosoftGraphGroup) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName sets OnPremisesSamAccountName field to given value.

### HasOnPremisesSamAccountName

`func (o *MicrosoftGraphGroup) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountNameNil

`func (o *MicrosoftGraphGroup) SetOnPremisesSamAccountNameNil(b bool)`

 SetOnPremisesSamAccountNameNil sets the value for OnPremisesSamAccountName to be an explicit nil

### UnsetOnPremisesSamAccountName
`func (o *MicrosoftGraphGroup) UnsetOnPremisesSamAccountName()`

UnsetOnPremisesSamAccountName ensures that no value is present for OnPremisesSamAccountName, not even an explicit nil
### GetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphGroup) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *MicrosoftGraphGroup) GetOnPremisesSecurityIdentifierOk() (*string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphGroup) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier sets OnPremisesSecurityIdentifier field to given value.

### HasOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphGroup) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifierNil

`func (o *MicrosoftGraphGroup) SetOnPremisesSecurityIdentifierNil(b bool)`

 SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil

### UnsetOnPremisesSecurityIdentifier
`func (o *MicrosoftGraphGroup) UnsetOnPremisesSecurityIdentifier()`

UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphGroup) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphGroup) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphGroup) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphGroup) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *MicrosoftGraphGroup) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *MicrosoftGraphGroup) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetPreferredDataLocation

`func (o *MicrosoftGraphGroup) GetPreferredDataLocation() string`

GetPreferredDataLocation returns the PreferredDataLocation field if non-nil, zero value otherwise.

### GetPreferredDataLocationOk

`func (o *MicrosoftGraphGroup) GetPreferredDataLocationOk() (*string, bool)`

GetPreferredDataLocationOk returns a tuple with the PreferredDataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDataLocation

`func (o *MicrosoftGraphGroup) SetPreferredDataLocation(v string)`

SetPreferredDataLocation sets PreferredDataLocation field to given value.

### HasPreferredDataLocation

`func (o *MicrosoftGraphGroup) HasPreferredDataLocation() bool`

HasPreferredDataLocation returns a boolean if a field has been set.

### SetPreferredDataLocationNil

`func (o *MicrosoftGraphGroup) SetPreferredDataLocationNil(b bool)`

 SetPreferredDataLocationNil sets the value for PreferredDataLocation to be an explicit nil

### UnsetPreferredDataLocation
`func (o *MicrosoftGraphGroup) UnsetPreferredDataLocation()`

UnsetPreferredDataLocation ensures that no value is present for PreferredDataLocation, not even an explicit nil
### GetPreferredLanguage

`func (o *MicrosoftGraphGroup) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *MicrosoftGraphGroup) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *MicrosoftGraphGroup) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *MicrosoftGraphGroup) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *MicrosoftGraphGroup) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *MicrosoftGraphGroup) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetProxyAddresses

`func (o *MicrosoftGraphGroup) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *MicrosoftGraphGroup) GetProxyAddressesOk() (*[]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddresses

`func (o *MicrosoftGraphGroup) SetProxyAddresses(v []string)`

SetProxyAddresses sets ProxyAddresses field to given value.

### HasProxyAddresses

`func (o *MicrosoftGraphGroup) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### GetRenewedDateTime

`func (o *MicrosoftGraphGroup) GetRenewedDateTime() time.Time`

GetRenewedDateTime returns the RenewedDateTime field if non-nil, zero value otherwise.

### GetRenewedDateTimeOk

`func (o *MicrosoftGraphGroup) GetRenewedDateTimeOk() (*time.Time, bool)`

GetRenewedDateTimeOk returns a tuple with the RenewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewedDateTime

`func (o *MicrosoftGraphGroup) SetRenewedDateTime(v time.Time)`

SetRenewedDateTime sets RenewedDateTime field to given value.

### HasRenewedDateTime

`func (o *MicrosoftGraphGroup) HasRenewedDateTime() bool`

HasRenewedDateTime returns a boolean if a field has been set.

### SetRenewedDateTimeNil

`func (o *MicrosoftGraphGroup) SetRenewedDateTimeNil(b bool)`

 SetRenewedDateTimeNil sets the value for RenewedDateTime to be an explicit nil

### UnsetRenewedDateTime
`func (o *MicrosoftGraphGroup) UnsetRenewedDateTime()`

UnsetRenewedDateTime ensures that no value is present for RenewedDateTime, not even an explicit nil
### GetSecurityEnabled

`func (o *MicrosoftGraphGroup) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *MicrosoftGraphGroup) GetSecurityEnabledOk() (*bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityEnabled

`func (o *MicrosoftGraphGroup) SetSecurityEnabled(v bool)`

SetSecurityEnabled sets SecurityEnabled field to given value.

### HasSecurityEnabled

`func (o *MicrosoftGraphGroup) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### SetSecurityEnabledNil

`func (o *MicrosoftGraphGroup) SetSecurityEnabledNil(b bool)`

 SetSecurityEnabledNil sets the value for SecurityEnabled to be an explicit nil

### UnsetSecurityEnabled
`func (o *MicrosoftGraphGroup) UnsetSecurityEnabled()`

UnsetSecurityEnabled ensures that no value is present for SecurityEnabled, not even an explicit nil
### GetSecurityIdentifier

`func (o *MicrosoftGraphGroup) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *MicrosoftGraphGroup) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *MicrosoftGraphGroup) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *MicrosoftGraphGroup) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### SetSecurityIdentifierNil

`func (o *MicrosoftGraphGroup) SetSecurityIdentifierNil(b bool)`

 SetSecurityIdentifierNil sets the value for SecurityIdentifier to be an explicit nil

### UnsetSecurityIdentifier
`func (o *MicrosoftGraphGroup) UnsetSecurityIdentifier()`

UnsetSecurityIdentifier ensures that no value is present for SecurityIdentifier, not even an explicit nil
### GetTheme

`func (o *MicrosoftGraphGroup) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MicrosoftGraphGroup) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MicrosoftGraphGroup) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MicrosoftGraphGroup) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *MicrosoftGraphGroup) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *MicrosoftGraphGroup) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil
### GetVisibility

`func (o *MicrosoftGraphGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MicrosoftGraphGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MicrosoftGraphGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *MicrosoftGraphGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *MicrosoftGraphGroup) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *MicrosoftGraphGroup) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetAllowExternalSenders

`func (o *MicrosoftGraphGroup) GetAllowExternalSenders() bool`

GetAllowExternalSenders returns the AllowExternalSenders field if non-nil, zero value otherwise.

### GetAllowExternalSendersOk

`func (o *MicrosoftGraphGroup) GetAllowExternalSendersOk() (*bool, bool)`

GetAllowExternalSendersOk returns a tuple with the AllowExternalSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExternalSenders

`func (o *MicrosoftGraphGroup) SetAllowExternalSenders(v bool)`

SetAllowExternalSenders sets AllowExternalSenders field to given value.

### HasAllowExternalSenders

`func (o *MicrosoftGraphGroup) HasAllowExternalSenders() bool`

HasAllowExternalSenders returns a boolean if a field has been set.

### SetAllowExternalSendersNil

`func (o *MicrosoftGraphGroup) SetAllowExternalSendersNil(b bool)`

 SetAllowExternalSendersNil sets the value for AllowExternalSenders to be an explicit nil

### UnsetAllowExternalSenders
`func (o *MicrosoftGraphGroup) UnsetAllowExternalSenders()`

UnsetAllowExternalSenders ensures that no value is present for AllowExternalSenders, not even an explicit nil
### GetAutoSubscribeNewMembers

`func (o *MicrosoftGraphGroup) GetAutoSubscribeNewMembers() bool`

GetAutoSubscribeNewMembers returns the AutoSubscribeNewMembers field if non-nil, zero value otherwise.

### GetAutoSubscribeNewMembersOk

`func (o *MicrosoftGraphGroup) GetAutoSubscribeNewMembersOk() (*bool, bool)`

GetAutoSubscribeNewMembersOk returns a tuple with the AutoSubscribeNewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubscribeNewMembers

`func (o *MicrosoftGraphGroup) SetAutoSubscribeNewMembers(v bool)`

SetAutoSubscribeNewMembers sets AutoSubscribeNewMembers field to given value.

### HasAutoSubscribeNewMembers

`func (o *MicrosoftGraphGroup) HasAutoSubscribeNewMembers() bool`

HasAutoSubscribeNewMembers returns a boolean if a field has been set.

### SetAutoSubscribeNewMembersNil

`func (o *MicrosoftGraphGroup) SetAutoSubscribeNewMembersNil(b bool)`

 SetAutoSubscribeNewMembersNil sets the value for AutoSubscribeNewMembers to be an explicit nil

### UnsetAutoSubscribeNewMembers
`func (o *MicrosoftGraphGroup) UnsetAutoSubscribeNewMembers()`

UnsetAutoSubscribeNewMembers ensures that no value is present for AutoSubscribeNewMembers, not even an explicit nil
### GetHideFromAddressLists

`func (o *MicrosoftGraphGroup) GetHideFromAddressLists() bool`

GetHideFromAddressLists returns the HideFromAddressLists field if non-nil, zero value otherwise.

### GetHideFromAddressListsOk

`func (o *MicrosoftGraphGroup) GetHideFromAddressListsOk() (*bool, bool)`

GetHideFromAddressListsOk returns a tuple with the HideFromAddressLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromAddressLists

`func (o *MicrosoftGraphGroup) SetHideFromAddressLists(v bool)`

SetHideFromAddressLists sets HideFromAddressLists field to given value.

### HasHideFromAddressLists

`func (o *MicrosoftGraphGroup) HasHideFromAddressLists() bool`

HasHideFromAddressLists returns a boolean if a field has been set.

### SetHideFromAddressListsNil

`func (o *MicrosoftGraphGroup) SetHideFromAddressListsNil(b bool)`

 SetHideFromAddressListsNil sets the value for HideFromAddressLists to be an explicit nil

### UnsetHideFromAddressLists
`func (o *MicrosoftGraphGroup) UnsetHideFromAddressLists()`

UnsetHideFromAddressLists ensures that no value is present for HideFromAddressLists, not even an explicit nil
### GetHideFromOutlookClients

`func (o *MicrosoftGraphGroup) GetHideFromOutlookClients() bool`

GetHideFromOutlookClients returns the HideFromOutlookClients field if non-nil, zero value otherwise.

### GetHideFromOutlookClientsOk

`func (o *MicrosoftGraphGroup) GetHideFromOutlookClientsOk() (*bool, bool)`

GetHideFromOutlookClientsOk returns a tuple with the HideFromOutlookClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromOutlookClients

`func (o *MicrosoftGraphGroup) SetHideFromOutlookClients(v bool)`

SetHideFromOutlookClients sets HideFromOutlookClients field to given value.

### HasHideFromOutlookClients

`func (o *MicrosoftGraphGroup) HasHideFromOutlookClients() bool`

HasHideFromOutlookClients returns a boolean if a field has been set.

### SetHideFromOutlookClientsNil

`func (o *MicrosoftGraphGroup) SetHideFromOutlookClientsNil(b bool)`

 SetHideFromOutlookClientsNil sets the value for HideFromOutlookClients to be an explicit nil

### UnsetHideFromOutlookClients
`func (o *MicrosoftGraphGroup) UnsetHideFromOutlookClients()`

UnsetHideFromOutlookClients ensures that no value is present for HideFromOutlookClients, not even an explicit nil
### GetIsSubscribedByMail

`func (o *MicrosoftGraphGroup) GetIsSubscribedByMail() bool`

GetIsSubscribedByMail returns the IsSubscribedByMail field if non-nil, zero value otherwise.

### GetIsSubscribedByMailOk

`func (o *MicrosoftGraphGroup) GetIsSubscribedByMailOk() (*bool, bool)`

GetIsSubscribedByMailOk returns a tuple with the IsSubscribedByMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedByMail

`func (o *MicrosoftGraphGroup) SetIsSubscribedByMail(v bool)`

SetIsSubscribedByMail sets IsSubscribedByMail field to given value.

### HasIsSubscribedByMail

`func (o *MicrosoftGraphGroup) HasIsSubscribedByMail() bool`

HasIsSubscribedByMail returns a boolean if a field has been set.

### SetIsSubscribedByMailNil

`func (o *MicrosoftGraphGroup) SetIsSubscribedByMailNil(b bool)`

 SetIsSubscribedByMailNil sets the value for IsSubscribedByMail to be an explicit nil

### UnsetIsSubscribedByMail
`func (o *MicrosoftGraphGroup) UnsetIsSubscribedByMail()`

UnsetIsSubscribedByMail ensures that no value is present for IsSubscribedByMail, not even an explicit nil
### GetUnseenCount

`func (o *MicrosoftGraphGroup) GetUnseenCount() int32`

GetUnseenCount returns the UnseenCount field if non-nil, zero value otherwise.

### GetUnseenCountOk

`func (o *MicrosoftGraphGroup) GetUnseenCountOk() (*int32, bool)`

GetUnseenCountOk returns a tuple with the UnseenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnseenCount

`func (o *MicrosoftGraphGroup) SetUnseenCount(v int32)`

SetUnseenCount sets UnseenCount field to given value.

### HasUnseenCount

`func (o *MicrosoftGraphGroup) HasUnseenCount() bool`

HasUnseenCount returns a boolean if a field has been set.

### SetUnseenCountNil

`func (o *MicrosoftGraphGroup) SetUnseenCountNil(b bool)`

 SetUnseenCountNil sets the value for UnseenCount to be an explicit nil

### UnsetUnseenCount
`func (o *MicrosoftGraphGroup) UnsetUnseenCount()`

UnsetUnseenCount ensures that no value is present for UnseenCount, not even an explicit nil
### GetIsArchived

`func (o *MicrosoftGraphGroup) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *MicrosoftGraphGroup) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *MicrosoftGraphGroup) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *MicrosoftGraphGroup) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchivedNil

`func (o *MicrosoftGraphGroup) SetIsArchivedNil(b bool)`

 SetIsArchivedNil sets the value for IsArchived to be an explicit nil

### UnsetIsArchived
`func (o *MicrosoftGraphGroup) UnsetIsArchived()`

UnsetIsArchived ensures that no value is present for IsArchived, not even an explicit nil
### GetAppRoleAssignments

`func (o *MicrosoftGraphGroup) GetAppRoleAssignments() []MicrosoftGraphAppRoleAssignment`

GetAppRoleAssignments returns the AppRoleAssignments field if non-nil, zero value otherwise.

### GetAppRoleAssignmentsOk

`func (o *MicrosoftGraphGroup) GetAppRoleAssignmentsOk() (*[]MicrosoftGraphAppRoleAssignment, bool)`

GetAppRoleAssignmentsOk returns a tuple with the AppRoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRoleAssignments

`func (o *MicrosoftGraphGroup) SetAppRoleAssignments(v []MicrosoftGraphAppRoleAssignment)`

SetAppRoleAssignments sets AppRoleAssignments field to given value.

### HasAppRoleAssignments

`func (o *MicrosoftGraphGroup) HasAppRoleAssignments() bool`

HasAppRoleAssignments returns a boolean if a field has been set.

### GetCreatedOnBehalfOf

`func (o *MicrosoftGraphGroup) GetCreatedOnBehalfOf() AnyOfmicrosoftGraphDirectoryObject`

GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field if non-nil, zero value otherwise.

### GetCreatedOnBehalfOfOk

`func (o *MicrosoftGraphGroup) GetCreatedOnBehalfOfOk() (*AnyOfmicrosoftGraphDirectoryObject, bool)`

GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOnBehalfOf

`func (o *MicrosoftGraphGroup) SetCreatedOnBehalfOf(v AnyOfmicrosoftGraphDirectoryObject)`

SetCreatedOnBehalfOf sets CreatedOnBehalfOf field to given value.

### HasCreatedOnBehalfOf

`func (o *MicrosoftGraphGroup) HasCreatedOnBehalfOf() bool`

HasCreatedOnBehalfOf returns a boolean if a field has been set.

### SetCreatedOnBehalfOfNil

`func (o *MicrosoftGraphGroup) SetCreatedOnBehalfOfNil(b bool)`

 SetCreatedOnBehalfOfNil sets the value for CreatedOnBehalfOf to be an explicit nil

### UnsetCreatedOnBehalfOf
`func (o *MicrosoftGraphGroup) UnsetCreatedOnBehalfOf()`

UnsetCreatedOnBehalfOf ensures that no value is present for CreatedOnBehalfOf, not even an explicit nil
### GetMemberOf

`func (o *MicrosoftGraphGroup) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MicrosoftGraphGroup) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *MicrosoftGraphGroup) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *MicrosoftGraphGroup) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetMembers

`func (o *MicrosoftGraphGroup) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphGroup) GetMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MicrosoftGraphGroup) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *MicrosoftGraphGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) GetMembersWithLicenseErrors() []MicrosoftGraphDirectoryObject`

GetMembersWithLicenseErrors returns the MembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetMembersWithLicenseErrorsOk

`func (o *MicrosoftGraphGroup) GetMembersWithLicenseErrorsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMembersWithLicenseErrorsOk returns a tuple with the MembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) SetMembersWithLicenseErrors(v []MicrosoftGraphDirectoryObject)`

SetMembersWithLicenseErrors sets MembersWithLicenseErrors field to given value.

### HasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) HasMembersWithLicenseErrors() bool`

HasMembersWithLicenseErrors returns a boolean if a field has been set.

### GetOwners

`func (o *MicrosoftGraphGroup) GetOwners() []MicrosoftGraphDirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MicrosoftGraphGroup) GetOwnersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *MicrosoftGraphGroup) SetOwners(v []MicrosoftGraphDirectoryObject)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *MicrosoftGraphGroup) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetPermissionGrants

`func (o *MicrosoftGraphGroup) GetPermissionGrants() []MicrosoftGraphResourceSpecificPermissionGrant`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *MicrosoftGraphGroup) GetPermissionGrantsOk() (*[]MicrosoftGraphResourceSpecificPermissionGrant, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *MicrosoftGraphGroup) SetPermissionGrants(v []MicrosoftGraphResourceSpecificPermissionGrant)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *MicrosoftGraphGroup) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.

### GetSettings

`func (o *MicrosoftGraphGroup) GetSettings() []MicrosoftGraphGroupSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphGroup) GetSettingsOk() (*[]MicrosoftGraphGroupSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MicrosoftGraphGroup) SetSettings(v []MicrosoftGraphGroupSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MicrosoftGraphGroup) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *MicrosoftGraphGroup) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *MicrosoftGraphGroup) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *MicrosoftGraphGroup) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *MicrosoftGraphGroup) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### GetTransitiveMembers

`func (o *MicrosoftGraphGroup) GetTransitiveMembers() []MicrosoftGraphDirectoryObject`

GetTransitiveMembers returns the TransitiveMembers field if non-nil, zero value otherwise.

### GetTransitiveMembersOk

`func (o *MicrosoftGraphGroup) GetTransitiveMembersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMembersOk returns a tuple with the TransitiveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembers

`func (o *MicrosoftGraphGroup) SetTransitiveMembers(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMembers sets TransitiveMembers field to given value.

### HasTransitiveMembers

`func (o *MicrosoftGraphGroup) HasTransitiveMembers() bool`

HasTransitiveMembers returns a boolean if a field has been set.

### GetAcceptedSenders

`func (o *MicrosoftGraphGroup) GetAcceptedSenders() []MicrosoftGraphDirectoryObject`

GetAcceptedSenders returns the AcceptedSenders field if non-nil, zero value otherwise.

### GetAcceptedSendersOk

`func (o *MicrosoftGraphGroup) GetAcceptedSendersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetAcceptedSendersOk returns a tuple with the AcceptedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedSenders

`func (o *MicrosoftGraphGroup) SetAcceptedSenders(v []MicrosoftGraphDirectoryObject)`

SetAcceptedSenders sets AcceptedSenders field to given value.

### HasAcceptedSenders

`func (o *MicrosoftGraphGroup) HasAcceptedSenders() bool`

HasAcceptedSenders returns a boolean if a field has been set.

### GetCalendar

`func (o *MicrosoftGraphGroup) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MicrosoftGraphGroup) GetCalendarOk() (*AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *MicrosoftGraphGroup) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *MicrosoftGraphGroup) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendarNil

`func (o *MicrosoftGraphGroup) SetCalendarNil(b bool)`

 SetCalendarNil sets the value for Calendar to be an explicit nil

### UnsetCalendar
`func (o *MicrosoftGraphGroup) UnsetCalendar()`

UnsetCalendar ensures that no value is present for Calendar, not even an explicit nil
### GetCalendarView

`func (o *MicrosoftGraphGroup) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *MicrosoftGraphGroup) GetCalendarViewOk() (*[]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarView

`func (o *MicrosoftGraphGroup) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView sets CalendarView field to given value.

### HasCalendarView

`func (o *MicrosoftGraphGroup) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### GetConversations

`func (o *MicrosoftGraphGroup) GetConversations() []MicrosoftGraphConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *MicrosoftGraphGroup) GetConversationsOk() (*[]MicrosoftGraphConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *MicrosoftGraphGroup) SetConversations(v []MicrosoftGraphConversation)`

SetConversations sets Conversations field to given value.

### HasConversations

`func (o *MicrosoftGraphGroup) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### GetEvents

`func (o *MicrosoftGraphGroup) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MicrosoftGraphGroup) GetEventsOk() (*[]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *MicrosoftGraphGroup) SetEvents(v []MicrosoftGraphEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *MicrosoftGraphGroup) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetPhoto

`func (o *MicrosoftGraphGroup) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphGroup) GetPhotoOk() (*AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MicrosoftGraphGroup) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MicrosoftGraphGroup) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *MicrosoftGraphGroup) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *MicrosoftGraphGroup) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPhotos

`func (o *MicrosoftGraphGroup) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *MicrosoftGraphGroup) GetPhotosOk() (*[]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *MicrosoftGraphGroup) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *MicrosoftGraphGroup) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetRejectedSenders

`func (o *MicrosoftGraphGroup) GetRejectedSenders() []MicrosoftGraphDirectoryObject`

GetRejectedSenders returns the RejectedSenders field if non-nil, zero value otherwise.

### GetRejectedSendersOk

`func (o *MicrosoftGraphGroup) GetRejectedSendersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRejectedSendersOk returns a tuple with the RejectedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedSenders

`func (o *MicrosoftGraphGroup) SetRejectedSenders(v []MicrosoftGraphDirectoryObject)`

SetRejectedSenders sets RejectedSenders field to given value.

### HasRejectedSenders

`func (o *MicrosoftGraphGroup) HasRejectedSenders() bool`

HasRejectedSenders returns a boolean if a field has been set.

### GetThreads

`func (o *MicrosoftGraphGroup) GetThreads() []MicrosoftGraphConversationThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *MicrosoftGraphGroup) GetThreadsOk() (*[]MicrosoftGraphConversationThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *MicrosoftGraphGroup) SetThreads(v []MicrosoftGraphConversationThread)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *MicrosoftGraphGroup) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetDrive

`func (o *MicrosoftGraphGroup) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *MicrosoftGraphGroup) GetDriveOk() (*AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrive

`func (o *MicrosoftGraphGroup) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive sets Drive field to given value.

### HasDrive

`func (o *MicrosoftGraphGroup) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDriveNil

`func (o *MicrosoftGraphGroup) SetDriveNil(b bool)`

 SetDriveNil sets the value for Drive to be an explicit nil

### UnsetDrive
`func (o *MicrosoftGraphGroup) UnsetDrive()`

UnsetDrive ensures that no value is present for Drive, not even an explicit nil
### GetDrives

`func (o *MicrosoftGraphGroup) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *MicrosoftGraphGroup) GetDrivesOk() (*[]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *MicrosoftGraphGroup) SetDrives(v []MicrosoftGraphDrive)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *MicrosoftGraphGroup) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### GetSites

`func (o *MicrosoftGraphGroup) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *MicrosoftGraphGroup) GetSitesOk() (*[]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *MicrosoftGraphGroup) SetSites(v []MicrosoftGraphSite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *MicrosoftGraphGroup) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetExtensions

`func (o *MicrosoftGraphGroup) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphGroup) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphGroup) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphGroup) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGroupLifecyclePolicies

`func (o *MicrosoftGraphGroup) GetGroupLifecyclePolicies() []MicrosoftGraphGroupLifecyclePolicy`

GetGroupLifecyclePolicies returns the GroupLifecyclePolicies field if non-nil, zero value otherwise.

### GetGroupLifecyclePoliciesOk

`func (o *MicrosoftGraphGroup) GetGroupLifecyclePoliciesOk() (*[]MicrosoftGraphGroupLifecyclePolicy, bool)`

GetGroupLifecyclePoliciesOk returns a tuple with the GroupLifecyclePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLifecyclePolicies

`func (o *MicrosoftGraphGroup) SetGroupLifecyclePolicies(v []MicrosoftGraphGroupLifecyclePolicy)`

SetGroupLifecyclePolicies sets GroupLifecyclePolicies field to given value.

### HasGroupLifecyclePolicies

`func (o *MicrosoftGraphGroup) HasGroupLifecyclePolicies() bool`

HasGroupLifecyclePolicies returns a boolean if a field has been set.

### GetPlanner

`func (o *MicrosoftGraphGroup) GetPlanner() AnyOfmicrosoftGraphPlannerGroup`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *MicrosoftGraphGroup) GetPlannerOk() (*AnyOfmicrosoftGraphPlannerGroup, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanner

`func (o *MicrosoftGraphGroup) SetPlanner(v AnyOfmicrosoftGraphPlannerGroup)`

SetPlanner sets Planner field to given value.

### HasPlanner

`func (o *MicrosoftGraphGroup) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlannerNil

`func (o *MicrosoftGraphGroup) SetPlannerNil(b bool)`

 SetPlannerNil sets the value for Planner to be an explicit nil

### UnsetPlanner
`func (o *MicrosoftGraphGroup) UnsetPlanner()`

UnsetPlanner ensures that no value is present for Planner, not even an explicit nil
### GetOnenote

`func (o *MicrosoftGraphGroup) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *MicrosoftGraphGroup) GetOnenoteOk() (*AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnenote

`func (o *MicrosoftGraphGroup) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote sets Onenote field to given value.

### HasOnenote

`func (o *MicrosoftGraphGroup) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenoteNil

`func (o *MicrosoftGraphGroup) SetOnenoteNil(b bool)`

 SetOnenoteNil sets the value for Onenote to be an explicit nil

### UnsetOnenote
`func (o *MicrosoftGraphGroup) UnsetOnenote()`

UnsetOnenote ensures that no value is present for Onenote, not even an explicit nil
### GetTeam

`func (o *MicrosoftGraphGroup) GetTeam() AnyOfmicrosoftGraphTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MicrosoftGraphGroup) GetTeamOk() (*AnyOfmicrosoftGraphTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MicrosoftGraphGroup) SetTeam(v AnyOfmicrosoftGraphTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MicrosoftGraphGroup) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *MicrosoftGraphGroup) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *MicrosoftGraphGroup) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


