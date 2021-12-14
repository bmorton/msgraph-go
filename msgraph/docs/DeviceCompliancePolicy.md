# DeviceCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | DateTime the object was created. | [optional] 
**Description** | Pointer to **NullableString** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicyAssignment**](MicrosoftGraphDeviceCompliancePolicyAssignment.md) | The collection of assignments for this compliance policy. | [optional] 
**DeviceSettingStateSummaries** | Pointer to [**[]MicrosoftGraphSettingStateDeviceSummary**](MicrosoftGraphSettingStateDeviceSummary.md) | Compliance Setting State Device Summary | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphDeviceComplianceDeviceStatus**](MicrosoftGraphDeviceComplianceDeviceStatus.md) | List of DeviceComplianceDeviceStatus. | [optional] 
**DeviceStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceDeviceOverview**](anyOf&lt;microsoft.graph.deviceComplianceDeviceOverview&gt;.md) | Device compliance devices status overview | [optional] 
**ScheduledActionsForRule** | Pointer to [**[]MicrosoftGraphDeviceComplianceScheduledActionForRule**](MicrosoftGraphDeviceComplianceScheduledActionForRule.md) | The list of scheduled action per rule for this compliance policy. This is a required property when creating any individual per-platform compliance policies. | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphDeviceComplianceUserStatus**](MicrosoftGraphDeviceComplianceUserStatus.md) | List of DeviceComplianceUserStatus. | [optional] 
**UserStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceUserOverview**](anyOf&lt;microsoft.graph.deviceComplianceUserOverview&gt;.md) | Device compliance users status overview | [optional] 

## Methods

### NewDeviceCompliancePolicy

`func NewDeviceCompliancePolicy() *DeviceCompliancePolicy`

NewDeviceCompliancePolicy instantiates a new DeviceCompliancePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCompliancePolicyWithDefaults

`func NewDeviceCompliancePolicyWithDefaults() *DeviceCompliancePolicy`

NewDeviceCompliancePolicyWithDefaults instantiates a new DeviceCompliancePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *DeviceCompliancePolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *DeviceCompliancePolicy) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *DeviceCompliancePolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *DeviceCompliancePolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceCompliancePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceCompliancePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceCompliancePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceCompliancePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeviceCompliancePolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeviceCompliancePolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *DeviceCompliancePolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceCompliancePolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceCompliancePolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceCompliancePolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *DeviceCompliancePolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *DeviceCompliancePolicy) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *DeviceCompliancePolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *DeviceCompliancePolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceCompliancePolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceCompliancePolicy) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceCompliancePolicy) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceCompliancePolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAssignments

`func (o *DeviceCompliancePolicy) GetAssignments() []MicrosoftGraphDeviceCompliancePolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *DeviceCompliancePolicy) GetAssignmentsOk() (*[]MicrosoftGraphDeviceCompliancePolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *DeviceCompliancePolicy) SetAssignments(v []MicrosoftGraphDeviceCompliancePolicyAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *DeviceCompliancePolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDeviceSettingStateSummaries

`func (o *DeviceCompliancePolicy) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *DeviceCompliancePolicy) GetDeviceSettingStateSummariesOk() (*[]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSettingStateSummaries

`func (o *DeviceCompliancePolicy) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries sets DeviceSettingStateSummaries field to given value.

### HasDeviceSettingStateSummaries

`func (o *DeviceCompliancePolicy) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### GetDeviceStatuses

`func (o *DeviceCompliancePolicy) GetDeviceStatuses() []MicrosoftGraphDeviceComplianceDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *DeviceCompliancePolicy) GetDeviceStatusesOk() (*[]MicrosoftGraphDeviceComplianceDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatuses

`func (o *DeviceCompliancePolicy) SetDeviceStatuses(v []MicrosoftGraphDeviceComplianceDeviceStatus)`

SetDeviceStatuses sets DeviceStatuses field to given value.

### HasDeviceStatuses

`func (o *DeviceCompliancePolicy) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### GetDeviceStatusOverview

`func (o *DeviceCompliancePolicy) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceComplianceDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *DeviceCompliancePolicy) GetDeviceStatusOverviewOk() (*AnyOfmicrosoftGraphDeviceComplianceDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatusOverview

`func (o *DeviceCompliancePolicy) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceDeviceOverview)`

SetDeviceStatusOverview sets DeviceStatusOverview field to given value.

### HasDeviceStatusOverview

`func (o *DeviceCompliancePolicy) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverviewNil

`func (o *DeviceCompliancePolicy) SetDeviceStatusOverviewNil(b bool)`

 SetDeviceStatusOverviewNil sets the value for DeviceStatusOverview to be an explicit nil

### UnsetDeviceStatusOverview
`func (o *DeviceCompliancePolicy) UnsetDeviceStatusOverview()`

UnsetDeviceStatusOverview ensures that no value is present for DeviceStatusOverview, not even an explicit nil
### GetScheduledActionsForRule

`func (o *DeviceCompliancePolicy) GetScheduledActionsForRule() []MicrosoftGraphDeviceComplianceScheduledActionForRule`

GetScheduledActionsForRule returns the ScheduledActionsForRule field if non-nil, zero value otherwise.

### GetScheduledActionsForRuleOk

`func (o *DeviceCompliancePolicy) GetScheduledActionsForRuleOk() (*[]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool)`

GetScheduledActionsForRuleOk returns a tuple with the ScheduledActionsForRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledActionsForRule

`func (o *DeviceCompliancePolicy) SetScheduledActionsForRule(v []MicrosoftGraphDeviceComplianceScheduledActionForRule)`

SetScheduledActionsForRule sets ScheduledActionsForRule field to given value.

### HasScheduledActionsForRule

`func (o *DeviceCompliancePolicy) HasScheduledActionsForRule() bool`

HasScheduledActionsForRule returns a boolean if a field has been set.

### GetUserStatuses

`func (o *DeviceCompliancePolicy) GetUserStatuses() []MicrosoftGraphDeviceComplianceUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *DeviceCompliancePolicy) GetUserStatusesOk() (*[]MicrosoftGraphDeviceComplianceUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatuses

`func (o *DeviceCompliancePolicy) SetUserStatuses(v []MicrosoftGraphDeviceComplianceUserStatus)`

SetUserStatuses sets UserStatuses field to given value.

### HasUserStatuses

`func (o *DeviceCompliancePolicy) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### GetUserStatusOverview

`func (o *DeviceCompliancePolicy) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceComplianceUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *DeviceCompliancePolicy) GetUserStatusOverviewOk() (*AnyOfmicrosoftGraphDeviceComplianceUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatusOverview

`func (o *DeviceCompliancePolicy) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceComplianceUserOverview)`

SetUserStatusOverview sets UserStatusOverview field to given value.

### HasUserStatusOverview

`func (o *DeviceCompliancePolicy) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverviewNil

`func (o *DeviceCompliancePolicy) SetUserStatusOverviewNil(b bool)`

 SetUserStatusOverviewNil sets the value for UserStatusOverview to be an explicit nil

### UnsetUserStatusOverview
`func (o *DeviceCompliancePolicy) UnsetUserStatusOverview()`

UnsetUserStatusOverview ensures that no value is present for UserStatusOverview, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


