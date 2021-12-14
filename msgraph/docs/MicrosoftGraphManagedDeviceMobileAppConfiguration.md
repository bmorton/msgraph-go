# MicrosoftGraphManagedDeviceMobileAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphManagedDeviceMobileAppConfiguration

`func NewMicrosoftGraphManagedDeviceMobileAppConfiguration() *MicrosoftGraphManagedDeviceMobileAppConfiguration`

NewMicrosoftGraphManagedDeviceMobileAppConfiguration instantiates a new MicrosoftGraphManagedDeviceMobileAppConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedDeviceMobileAppConfigurationWithDefaults

`func NewMicrosoftGraphManagedDeviceMobileAppConfigurationWithDefaults() *MicrosoftGraphManagedDeviceMobileAppConfiguration`

NewMicrosoftGraphManagedDeviceMobileAppConfigurationWithDefaults instantiates a new MicrosoftGraphManagedDeviceMobileAppConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetTargetedMobileApps

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetTargetedMobileApps() []*string`

GetTargetedMobileApps returns the TargetedMobileApps field if non-nil, zero value otherwise.

### GetTargetedMobileAppsOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetTargetedMobileAppsOk() (*[]*string, bool)`

GetTargetedMobileAppsOk returns a tuple with the TargetedMobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedMobileApps

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetTargetedMobileApps(v []*string)`

SetTargetedMobileApps sets TargetedMobileApps field to given value.

### HasTargetedMobileApps

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasTargetedMobileApps() bool`

HasTargetedMobileApps returns a boolean if a field has been set.

### GetVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetAssignments() []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetAssignments(v []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDeviceStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatusesOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDeviceStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus)`

SetDeviceStatuses sets DeviceStatuses field to given value.

### HasDeviceStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### GetDeviceStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary`

GetDeviceStatusSummary returns the DeviceStatusSummary field if non-nil, zero value otherwise.

### GetDeviceStatusSummaryOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatusSummaryOk() (*AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary, bool)`

GetDeviceStatusSummaryOk returns a tuple with the DeviceStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDeviceStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary)`

SetDeviceStatusSummary sets DeviceStatusSummary field to given value.

### HasDeviceStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDeviceStatusSummary() bool`

HasDeviceStatusSummary returns a boolean if a field has been set.

### SetDeviceStatusSummaryNil

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDeviceStatusSummaryNil(b bool)`

 SetDeviceStatusSummaryNil sets the value for DeviceStatusSummary to be an explicit nil

### UnsetDeviceStatusSummary
`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) UnsetDeviceStatusSummary()`

UnsetDeviceStatusSummary ensures that no value is present for DeviceStatusSummary, not even an explicit nil
### GetUserStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatusesOk() (*[]MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetUserStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus)`

SetUserStatuses sets UserStatuses field to given value.

### HasUserStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### GetUserStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary`

GetUserStatusSummary returns the UserStatusSummary field if non-nil, zero value otherwise.

### GetUserStatusSummaryOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatusSummaryOk() (*AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary, bool)`

GetUserStatusSummaryOk returns a tuple with the UserStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetUserStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary)`

SetUserStatusSummary sets UserStatusSummary field to given value.

### HasUserStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasUserStatusSummary() bool`

HasUserStatusSummary returns a boolean if a field has been set.

### SetUserStatusSummaryNil

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetUserStatusSummaryNil(b bool)`

 SetUserStatusSummaryNil sets the value for UserStatusSummary to be an explicit nil

### UnsetUserStatusSummary
`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) UnsetUserStatusSummary()`

UnsetUserStatusSummary ensures that no value is present for UserStatusSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


