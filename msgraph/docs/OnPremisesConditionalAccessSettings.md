# OnPremisesConditionalAccessSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if on premises conditional access is enabled for this organization | [optional] 
**ExcludedGroups** | Pointer to **[]string** | User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy. | [optional] 
**IncludedGroups** | Pointer to **[]string** | User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access. | [optional] 
**OverrideDefaultRule** | Pointer to **bool** | Override the default access rule when allowing a device to ensure access is granted. | [optional] 

## Methods

### NewOnPremisesConditionalAccessSettings

`func NewOnPremisesConditionalAccessSettings() *OnPremisesConditionalAccessSettings`

NewOnPremisesConditionalAccessSettings instantiates a new OnPremisesConditionalAccessSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremisesConditionalAccessSettingsWithDefaults

`func NewOnPremisesConditionalAccessSettingsWithDefaults() *OnPremisesConditionalAccessSettings`

NewOnPremisesConditionalAccessSettingsWithDefaults instantiates a new OnPremisesConditionalAccessSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OnPremisesConditionalAccessSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OnPremisesConditionalAccessSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OnPremisesConditionalAccessSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OnPremisesConditionalAccessSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExcludedGroups

`func (o *OnPremisesConditionalAccessSettings) GetExcludedGroups() []string`

GetExcludedGroups returns the ExcludedGroups field if non-nil, zero value otherwise.

### GetExcludedGroupsOk

`func (o *OnPremisesConditionalAccessSettings) GetExcludedGroupsOk() (*[]string, bool)`

GetExcludedGroupsOk returns a tuple with the ExcludedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedGroups

`func (o *OnPremisesConditionalAccessSettings) SetExcludedGroups(v []string)`

SetExcludedGroups sets ExcludedGroups field to given value.

### HasExcludedGroups

`func (o *OnPremisesConditionalAccessSettings) HasExcludedGroups() bool`

HasExcludedGroups returns a boolean if a field has been set.

### GetIncludedGroups

`func (o *OnPremisesConditionalAccessSettings) GetIncludedGroups() []string`

GetIncludedGroups returns the IncludedGroups field if non-nil, zero value otherwise.

### GetIncludedGroupsOk

`func (o *OnPremisesConditionalAccessSettings) GetIncludedGroupsOk() (*[]string, bool)`

GetIncludedGroupsOk returns a tuple with the IncludedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedGroups

`func (o *OnPremisesConditionalAccessSettings) SetIncludedGroups(v []string)`

SetIncludedGroups sets IncludedGroups field to given value.

### HasIncludedGroups

`func (o *OnPremisesConditionalAccessSettings) HasIncludedGroups() bool`

HasIncludedGroups returns a boolean if a field has been set.

### GetOverrideDefaultRule

`func (o *OnPremisesConditionalAccessSettings) GetOverrideDefaultRule() bool`

GetOverrideDefaultRule returns the OverrideDefaultRule field if non-nil, zero value otherwise.

### GetOverrideDefaultRuleOk

`func (o *OnPremisesConditionalAccessSettings) GetOverrideDefaultRuleOk() (*bool, bool)`

GetOverrideDefaultRuleOk returns a tuple with the OverrideDefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDefaultRule

`func (o *OnPremisesConditionalAccessSettings) SetOverrideDefaultRule(v bool)`

SetOverrideDefaultRule sets OverrideDefaultRule field to given value.

### HasOverrideDefaultRule

`func (o *OnPremisesConditionalAccessSettings) HasOverrideDefaultRule() bool`

HasOverrideDefaultRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


