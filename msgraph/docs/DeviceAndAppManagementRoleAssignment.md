# DeviceAndAppManagementRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** | The list of ids of role member security groups. These are IDs from Azure Active Directory. | [optional] 

## Methods

### NewDeviceAndAppManagementRoleAssignment

`func NewDeviceAndAppManagementRoleAssignment() *DeviceAndAppManagementRoleAssignment`

NewDeviceAndAppManagementRoleAssignment instantiates a new DeviceAndAppManagementRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAndAppManagementRoleAssignmentWithDefaults

`func NewDeviceAndAppManagementRoleAssignmentWithDefaults() *DeviceAndAppManagementRoleAssignment`

NewDeviceAndAppManagementRoleAssignmentWithDefaults instantiates a new DeviceAndAppManagementRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *DeviceAndAppManagementRoleAssignment) GetMembers() []*string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DeviceAndAppManagementRoleAssignment) GetMembersOk() (*[]*string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DeviceAndAppManagementRoleAssignment) SetMembers(v []*string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DeviceAndAppManagementRoleAssignment) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


