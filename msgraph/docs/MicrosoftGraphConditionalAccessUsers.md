# MicrosoftGraphConditionalAccessUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeGroups** | Pointer to **[]string** | Group IDs excluded from scope of policy. | [optional] 
**ExcludeRoles** | Pointer to **[]string** | Role IDs excluded from scope of policy. | [optional] 
**ExcludeUsers** | Pointer to **[]string** | User IDs excluded from scope of policy and/or GuestsOrExternalUsers. | [optional] 
**IncludeGroups** | Pointer to **[]string** | Group IDs in scope of policy unless explicitly excluded, or All. | [optional] 
**IncludeRoles** | Pointer to **[]string** | Role IDs in scope of policy unless explicitly excluded, or All. | [optional] 
**IncludeUsers** | Pointer to **[]string** | User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessUsers

`func NewMicrosoftGraphConditionalAccessUsers() *MicrosoftGraphConditionalAccessUsers`

NewMicrosoftGraphConditionalAccessUsers instantiates a new MicrosoftGraphConditionalAccessUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessUsersWithDefaults

`func NewMicrosoftGraphConditionalAccessUsersWithDefaults() *MicrosoftGraphConditionalAccessUsers`

NewMicrosoftGraphConditionalAccessUsersWithDefaults instantiates a new MicrosoftGraphConditionalAccessUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeGroups

`func (o *MicrosoftGraphConditionalAccessUsers) GetExcludeGroups() []string`

GetExcludeGroups returns the ExcludeGroups field if non-nil, zero value otherwise.

### GetExcludeGroupsOk

`func (o *MicrosoftGraphConditionalAccessUsers) GetExcludeGroupsOk() (*[]string, bool)`

GetExcludeGroupsOk returns a tuple with the ExcludeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGroups

`func (o *MicrosoftGraphConditionalAccessUsers) SetExcludeGroups(v []string)`

SetExcludeGroups sets ExcludeGroups field to given value.

### HasExcludeGroups

`func (o *MicrosoftGraphConditionalAccessUsers) HasExcludeGroups() bool`

HasExcludeGroups returns a boolean if a field has been set.

### GetExcludeRoles

`func (o *MicrosoftGraphConditionalAccessUsers) GetExcludeRoles() []string`

GetExcludeRoles returns the ExcludeRoles field if non-nil, zero value otherwise.

### GetExcludeRolesOk

`func (o *MicrosoftGraphConditionalAccessUsers) GetExcludeRolesOk() (*[]string, bool)`

GetExcludeRolesOk returns a tuple with the ExcludeRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRoles

`func (o *MicrosoftGraphConditionalAccessUsers) SetExcludeRoles(v []string)`

SetExcludeRoles sets ExcludeRoles field to given value.

### HasExcludeRoles

`func (o *MicrosoftGraphConditionalAccessUsers) HasExcludeRoles() bool`

HasExcludeRoles returns a boolean if a field has been set.

### GetExcludeUsers

`func (o *MicrosoftGraphConditionalAccessUsers) GetExcludeUsers() []string`

GetExcludeUsers returns the ExcludeUsers field if non-nil, zero value otherwise.

### GetExcludeUsersOk

`func (o *MicrosoftGraphConditionalAccessUsers) GetExcludeUsersOk() (*[]string, bool)`

GetExcludeUsersOk returns a tuple with the ExcludeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsers

`func (o *MicrosoftGraphConditionalAccessUsers) SetExcludeUsers(v []string)`

SetExcludeUsers sets ExcludeUsers field to given value.

### HasExcludeUsers

`func (o *MicrosoftGraphConditionalAccessUsers) HasExcludeUsers() bool`

HasExcludeUsers returns a boolean if a field has been set.

### GetIncludeGroups

`func (o *MicrosoftGraphConditionalAccessUsers) GetIncludeGroups() []string`

GetIncludeGroups returns the IncludeGroups field if non-nil, zero value otherwise.

### GetIncludeGroupsOk

`func (o *MicrosoftGraphConditionalAccessUsers) GetIncludeGroupsOk() (*[]string, bool)`

GetIncludeGroupsOk returns a tuple with the IncludeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGroups

`func (o *MicrosoftGraphConditionalAccessUsers) SetIncludeGroups(v []string)`

SetIncludeGroups sets IncludeGroups field to given value.

### HasIncludeGroups

`func (o *MicrosoftGraphConditionalAccessUsers) HasIncludeGroups() bool`

HasIncludeGroups returns a boolean if a field has been set.

### GetIncludeRoles

`func (o *MicrosoftGraphConditionalAccessUsers) GetIncludeRoles() []string`

GetIncludeRoles returns the IncludeRoles field if non-nil, zero value otherwise.

### GetIncludeRolesOk

`func (o *MicrosoftGraphConditionalAccessUsers) GetIncludeRolesOk() (*[]string, bool)`

GetIncludeRolesOk returns a tuple with the IncludeRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRoles

`func (o *MicrosoftGraphConditionalAccessUsers) SetIncludeRoles(v []string)`

SetIncludeRoles sets IncludeRoles field to given value.

### HasIncludeRoles

`func (o *MicrosoftGraphConditionalAccessUsers) HasIncludeRoles() bool`

HasIncludeRoles returns a boolean if a field has been set.

### GetIncludeUsers

`func (o *MicrosoftGraphConditionalAccessUsers) GetIncludeUsers() []string`

GetIncludeUsers returns the IncludeUsers field if non-nil, zero value otherwise.

### GetIncludeUsersOk

`func (o *MicrosoftGraphConditionalAccessUsers) GetIncludeUsersOk() (*[]string, bool)`

GetIncludeUsersOk returns a tuple with the IncludeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUsers

`func (o *MicrosoftGraphConditionalAccessUsers) SetIncludeUsers(v []string)`

SetIncludeUsers sets IncludeUsers field to given value.

### HasIncludeUsers

`func (o *MicrosoftGraphConditionalAccessUsers) HasIncludeUsers() bool`

HasIncludeUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


