# DeviceEnrollmentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | Created date time in UTC of the device enrollment configuration | [optional] 
**Description** | Pointer to **NullableString** | The description of the device enrollment configuration | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the device enrollment configuration | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Last modified date time in UTC of the device enrollment configuration | [optional] 
**Priority** | Pointer to **int32** | Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject only to the configuration with the lowest priority value. | [optional] 
**Version** | Pointer to **int32** | The version of the device enrollment configuration | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md) | The list of group assignments for the device configuration profile | [optional] 

## Methods

### NewDeviceEnrollmentConfiguration

`func NewDeviceEnrollmentConfiguration() *DeviceEnrollmentConfiguration`

NewDeviceEnrollmentConfiguration instantiates a new DeviceEnrollmentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentConfigurationWithDefaults

`func NewDeviceEnrollmentConfigurationWithDefaults() *DeviceEnrollmentConfiguration`

NewDeviceEnrollmentConfigurationWithDefaults instantiates a new DeviceEnrollmentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *DeviceEnrollmentConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *DeviceEnrollmentConfiguration) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *DeviceEnrollmentConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *DeviceEnrollmentConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceEnrollmentConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceEnrollmentConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceEnrollmentConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceEnrollmentConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeviceEnrollmentConfiguration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeviceEnrollmentConfiguration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *DeviceEnrollmentConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceEnrollmentConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceEnrollmentConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceEnrollmentConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DeviceEnrollmentConfiguration) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DeviceEnrollmentConfiguration) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedDateTime

`func (o *DeviceEnrollmentConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *DeviceEnrollmentConfiguration) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *DeviceEnrollmentConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *DeviceEnrollmentConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetPriority

`func (o *DeviceEnrollmentConfiguration) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeviceEnrollmentConfiguration) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeviceEnrollmentConfiguration) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DeviceEnrollmentConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceEnrollmentConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceEnrollmentConfiguration) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceEnrollmentConfiguration) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceEnrollmentConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAssignments

`func (o *DeviceEnrollmentConfiguration) GetAssignments() []MicrosoftGraphEnrollmentConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *DeviceEnrollmentConfiguration) GetAssignmentsOk() (*[]MicrosoftGraphEnrollmentConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *DeviceEnrollmentConfiguration) SetAssignments(v []MicrosoftGraphEnrollmentConfigurationAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *DeviceEnrollmentConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


