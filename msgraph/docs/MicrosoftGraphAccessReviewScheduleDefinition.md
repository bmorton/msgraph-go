# MicrosoftGraphAccessReviewScheduleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AdditionalNotificationRecipients** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewNotificationRecipientItem**](AnyOfmicrosoftGraphAccessReviewNotificationRecipientItem.md) | Defines the list of additional users or group members to be notified of the access review progress. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) | User who created this review. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Timestamp when the access review series was created. Supports $select and $orderBy. Read-only. | [optional] 
**DescriptionForAdmins** | Pointer to **NullableString** | Description provided by review creators to provide more context of the review to admins. Supports $select. | [optional] 
**DescriptionForReviewers** | Pointer to **NullableString** | Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the access review series. Supports $select and $orderBy. Required on create. | [optional] 
**FallbackReviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user&#39;s manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select. | [optional] 
**InstanceEnumerationScope** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | This property is required when scoping a review to guest users&#39; access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Timestamp when the access review series was last modified. Supports $select. Read-only. | [optional] 
**Reviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. | [optional] 
**Scope** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Defines the entities whose access is reviewed.  For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API. | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphAccessReviewScheduleSettings**](anyOf&lt;microsoft.graph.accessReviewScheduleSettings&gt;.md) | The settings for an access review series, see type definition below. Supports $select. Required on create. | [optional] 
**Status** | Pointer to **NullableString** | This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only. | [optional] 
**Instances** | Pointer to [**[]MicrosoftGraphAccessReviewInstance**](MicrosoftGraphAccessReviewInstance.md) | If the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence. | [optional] 

## Methods

### NewMicrosoftGraphAccessReviewScheduleDefinition

`func NewMicrosoftGraphAccessReviewScheduleDefinition() *MicrosoftGraphAccessReviewScheduleDefinition`

NewMicrosoftGraphAccessReviewScheduleDefinition instantiates a new MicrosoftGraphAccessReviewScheduleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessReviewScheduleDefinitionWithDefaults

`func NewMicrosoftGraphAccessReviewScheduleDefinitionWithDefaults() *MicrosoftGraphAccessReviewScheduleDefinition`

NewMicrosoftGraphAccessReviewScheduleDefinitionWithDefaults instantiates a new MicrosoftGraphAccessReviewScheduleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdditionalNotificationRecipients

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetAdditionalNotificationRecipients() []*AnyOfmicrosoftGraphAccessReviewNotificationRecipientItem`

GetAdditionalNotificationRecipients returns the AdditionalNotificationRecipients field if non-nil, zero value otherwise.

### GetAdditionalNotificationRecipientsOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetAdditionalNotificationRecipientsOk() (*[]*AnyOfmicrosoftGraphAccessReviewNotificationRecipientItem, bool)`

GetAdditionalNotificationRecipientsOk returns a tuple with the AdditionalNotificationRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalNotificationRecipients

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetAdditionalNotificationRecipients(v []*AnyOfmicrosoftGraphAccessReviewNotificationRecipientItem)`

SetAdditionalNotificationRecipients sets AdditionalNotificationRecipients field to given value.

### HasAdditionalNotificationRecipients

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasAdditionalNotificationRecipients() bool`

HasAdditionalNotificationRecipients returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetCreatedBy() AnyOfmicrosoftGraphUserIdentity`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetCreatedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetCreatedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescriptionForAdmins

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetDescriptionForAdmins() string`

GetDescriptionForAdmins returns the DescriptionForAdmins field if non-nil, zero value otherwise.

### GetDescriptionForAdminsOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetDescriptionForAdminsOk() (*string, bool)`

GetDescriptionForAdminsOk returns a tuple with the DescriptionForAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionForAdmins

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetDescriptionForAdmins(v string)`

SetDescriptionForAdmins sets DescriptionForAdmins field to given value.

### HasDescriptionForAdmins

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasDescriptionForAdmins() bool`

HasDescriptionForAdmins returns a boolean if a field has been set.

### SetDescriptionForAdminsNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetDescriptionForAdminsNil(b bool)`

 SetDescriptionForAdminsNil sets the value for DescriptionForAdmins to be an explicit nil

### UnsetDescriptionForAdmins
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetDescriptionForAdmins()`

UnsetDescriptionForAdmins ensures that no value is present for DescriptionForAdmins, not even an explicit nil
### GetDescriptionForReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetDescriptionForReviewers() string`

GetDescriptionForReviewers returns the DescriptionForReviewers field if non-nil, zero value otherwise.

### GetDescriptionForReviewersOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetDescriptionForReviewersOk() (*string, bool)`

GetDescriptionForReviewersOk returns a tuple with the DescriptionForReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionForReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetDescriptionForReviewers(v string)`

SetDescriptionForReviewers sets DescriptionForReviewers field to given value.

### HasDescriptionForReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasDescriptionForReviewers() bool`

HasDescriptionForReviewers returns a boolean if a field has been set.

### SetDescriptionForReviewersNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetDescriptionForReviewersNil(b bool)`

 SetDescriptionForReviewersNil sets the value for DescriptionForReviewers to be an explicit nil

### UnsetDescriptionForReviewers
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetDescriptionForReviewers()`

UnsetDescriptionForReviewers ensures that no value is present for DescriptionForReviewers, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFallbackReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetFallbackReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetFallbackReviewers returns the FallbackReviewers field if non-nil, zero value otherwise.

### GetFallbackReviewersOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetFallbackReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetFallbackReviewersOk returns a tuple with the FallbackReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetFallbackReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetFallbackReviewers sets FallbackReviewers field to given value.

### HasFallbackReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasFallbackReviewers() bool`

HasFallbackReviewers returns a boolean if a field has been set.

### GetInstanceEnumerationScope

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetInstanceEnumerationScope() AnyOfobject`

GetInstanceEnumerationScope returns the InstanceEnumerationScope field if non-nil, zero value otherwise.

### GetInstanceEnumerationScopeOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetInstanceEnumerationScopeOk() (*AnyOfobject, bool)`

GetInstanceEnumerationScopeOk returns a tuple with the InstanceEnumerationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceEnumerationScope

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetInstanceEnumerationScope(v AnyOfobject)`

SetInstanceEnumerationScope sets InstanceEnumerationScope field to given value.

### HasInstanceEnumerationScope

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasInstanceEnumerationScope() bool`

HasInstanceEnumerationScope returns a boolean if a field has been set.

### SetInstanceEnumerationScopeNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetInstanceEnumerationScopeNil(b bool)`

 SetInstanceEnumerationScopeNil sets the value for InstanceEnumerationScope to be an explicit nil

### UnsetInstanceEnumerationScope
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetInstanceEnumerationScope()`

UnsetInstanceEnumerationScope ensures that no value is present for InstanceEnumerationScope, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetScope

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetScope() AnyOfobject`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetScopeOk() (*AnyOfobject, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetScope(v AnyOfobject)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSettings

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetSettings() AnyOfmicrosoftGraphAccessReviewScheduleSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetSettingsOk() (*AnyOfmicrosoftGraphAccessReviewScheduleSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetSettings(v AnyOfmicrosoftGraphAccessReviewScheduleSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphAccessReviewScheduleDefinition) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInstances

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetInstances() []MicrosoftGraphAccessReviewInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) GetInstancesOk() (*[]MicrosoftGraphAccessReviewInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) SetInstances(v []MicrosoftGraphAccessReviewInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *MicrosoftGraphAccessReviewScheduleDefinition) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


