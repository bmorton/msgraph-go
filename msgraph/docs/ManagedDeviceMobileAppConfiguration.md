# ManagedDeviceMobileAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | DateTime the object was created. | [optional] 
**Description** | Pointer to **NullableString** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**TargetedMobileApps** | Pointer to **[]string** | the associated app. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment**](MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment.md) | The list of group assignemenets for app configration. | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus**](MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus.md) | List of ManagedDeviceMobileAppConfigurationDeviceStatus. | [optional] 
**DeviceStatusSummary** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary**](anyOf&lt;microsoft.graph.managedDeviceMobileAppConfigurationDeviceSummary&gt;.md) | App configuration device status summary. | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus**](MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus.md) | List of ManagedDeviceMobileAppConfigurationUserStatus. | [optional] 
**UserStatusSummary** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary**](anyOf&lt;microsoft.graph.managedDeviceMobileAppConfigurationUserSummary&gt;.md) | App configuration user status summary. | [optional] 

## Methods

### NewManagedDeviceMobileAppConfiguration

`func NewManagedDeviceMobileAppConfiguration() *ManagedDeviceMobileAppConfiguration`

NewManagedDeviceMobileAppConfiguration instantiates a new ManagedDeviceMobileAppConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDeviceMobileAppConfigurationWithDefaults

`func NewManagedDeviceMobileAppConfigurationWithDefaults() *ManagedDeviceMobileAppConfiguration`

NewManagedDeviceMobileAppConfigurationWithDefaults instantiates a new ManagedDeviceMobileAppConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *ManagedDeviceMobileAppConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ManagedDeviceMobileAppConfiguration) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ManagedDeviceMobileAppConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ManagedDeviceMobileAppConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *ManagedDeviceMobileAppConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedDeviceMobileAppConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedDeviceMobileAppConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagedDeviceMobileAppConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ManagedDeviceMobileAppConfiguration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ManagedDeviceMobileAppConfiguration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ManagedDeviceMobileAppConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManagedDeviceMobileAppConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManagedDeviceMobileAppConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManagedDeviceMobileAppConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *ManagedDeviceMobileAppConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *ManagedDeviceMobileAppConfiguration) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *ManagedDeviceMobileAppConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *ManagedDeviceMobileAppConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetTargetedMobileApps

`func (o *ManagedDeviceMobileAppConfiguration) GetTargetedMobileApps() []*string`

GetTargetedMobileApps returns the TargetedMobileApps field if non-nil, zero value otherwise.

### GetTargetedMobileAppsOk

`func (o *ManagedDeviceMobileAppConfiguration) GetTargetedMobileAppsOk() (*[]*string, bool)`

GetTargetedMobileAppsOk returns a tuple with the TargetedMobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedMobileApps

`func (o *ManagedDeviceMobileAppConfiguration) SetTargetedMobileApps(v []*string)`

SetTargetedMobileApps sets TargetedMobileApps field to given value.

### HasTargetedMobileApps

`func (o *ManagedDeviceMobileAppConfiguration) HasTargetedMobileApps() bool`

HasTargetedMobileApps returns a boolean if a field has been set.

### GetVersion

`func (o *ManagedDeviceMobileAppConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedDeviceMobileAppConfiguration) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedDeviceMobileAppConfiguration) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagedDeviceMobileAppConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAssignments

`func (o *ManagedDeviceMobileAppConfiguration) GetAssignments() []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ManagedDeviceMobileAppConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *ManagedDeviceMobileAppConfiguration) SetAssignments(v []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *ManagedDeviceMobileAppConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDeviceStatuses

`func (o *ManagedDeviceMobileAppConfiguration) GetDeviceStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *ManagedDeviceMobileAppConfiguration) GetDeviceStatusesOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatuses

`func (o *ManagedDeviceMobileAppConfiguration) SetDeviceStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus)`

SetDeviceStatuses sets DeviceStatuses field to given value.

### HasDeviceStatuses

`func (o *ManagedDeviceMobileAppConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### GetDeviceStatusSummary

`func (o *ManagedDeviceMobileAppConfiguration) GetDeviceStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary`

GetDeviceStatusSummary returns the DeviceStatusSummary field if non-nil, zero value otherwise.

### GetDeviceStatusSummaryOk

`func (o *ManagedDeviceMobileAppConfiguration) GetDeviceStatusSummaryOk() (*AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary, bool)`

GetDeviceStatusSummaryOk returns a tuple with the DeviceStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatusSummary

`func (o *ManagedDeviceMobileAppConfiguration) SetDeviceStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary)`

SetDeviceStatusSummary sets DeviceStatusSummary field to given value.

### HasDeviceStatusSummary

`func (o *ManagedDeviceMobileAppConfiguration) HasDeviceStatusSummary() bool`

HasDeviceStatusSummary returns a boolean if a field has been set.

### SetDeviceStatusSummaryNil

`func (o *ManagedDeviceMobileAppConfiguration) SetDeviceStatusSummaryNil(b bool)`

 SetDeviceStatusSummaryNil sets the value for DeviceStatusSummary to be an explicit nil

### UnsetDeviceStatusSummary
`func (o *ManagedDeviceMobileAppConfiguration) UnsetDeviceStatusSummary()`

UnsetDeviceStatusSummary ensures that no value is present for DeviceStatusSummary, not even an explicit nil
### GetUserStatuses

`func (o *ManagedDeviceMobileAppConfiguration) GetUserStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *ManagedDeviceMobileAppConfiguration) GetUserStatusesOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatuses

`func (o *ManagedDeviceMobileAppConfiguration) SetUserStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus)`

SetUserStatuses sets UserStatuses field to given value.

### HasUserStatuses

`func (o *ManagedDeviceMobileAppConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### GetUserStatusSummary

`func (o *ManagedDeviceMobileAppConfiguration) GetUserStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary`

GetUserStatusSummary returns the UserStatusSummary field if non-nil, zero value otherwise.

### GetUserStatusSummaryOk

`func (o *ManagedDeviceMobileAppConfiguration) GetUserStatusSummaryOk() (*AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary, bool)`

GetUserStatusSummaryOk returns a tuple with the UserStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatusSummary

`func (o *ManagedDeviceMobileAppConfiguration) SetUserStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary)`

SetUserStatusSummary sets UserStatusSummary field to given value.

### HasUserStatusSummary

`func (o *ManagedDeviceMobileAppConfiguration) HasUserStatusSummary() bool`

HasUserStatusSummary returns a boolean if a field has been set.

### SetUserStatusSummaryNil

`func (o *ManagedDeviceMobileAppConfiguration) SetUserStatusSummaryNil(b bool)`

 SetUserStatusSummaryNil sets the value for UserStatusSummary to be an explicit nil

### UnsetUserStatusSummary
`func (o *ManagedDeviceMobileAppConfiguration) UnsetUserStatusSummary()`

UnsetUserStatusSummary ensures that no value is present for UserStatusSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


