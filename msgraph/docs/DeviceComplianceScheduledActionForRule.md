# DeviceComplianceScheduledActionForRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | Pointer to **NullableString** | Name of the rule which this scheduled action applies to. Currently scheduled actions are created per policy instead of per rule, thus RuleName is always set to default value PasswordRequired. | [optional] 
**ScheduledActionConfigurations** | Pointer to [**[]MicrosoftGraphDeviceComplianceActionItem**](MicrosoftGraphDeviceComplianceActionItem.md) | The list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action. | [optional] 

## Methods

### NewDeviceComplianceScheduledActionForRule

`func NewDeviceComplianceScheduledActionForRule() *DeviceComplianceScheduledActionForRule`

NewDeviceComplianceScheduledActionForRule instantiates a new DeviceComplianceScheduledActionForRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceComplianceScheduledActionForRuleWithDefaults

`func NewDeviceComplianceScheduledActionForRuleWithDefaults() *DeviceComplianceScheduledActionForRule`

NewDeviceComplianceScheduledActionForRuleWithDefaults instantiates a new DeviceComplianceScheduledActionForRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleName

`func (o *DeviceComplianceScheduledActionForRule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *DeviceComplianceScheduledActionForRule) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *DeviceComplianceScheduledActionForRule) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *DeviceComplianceScheduledActionForRule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleNameNil

`func (o *DeviceComplianceScheduledActionForRule) SetRuleNameNil(b bool)`

 SetRuleNameNil sets the value for RuleName to be an explicit nil

### UnsetRuleName
`func (o *DeviceComplianceScheduledActionForRule) UnsetRuleName()`

UnsetRuleName ensures that no value is present for RuleName, not even an explicit nil
### GetScheduledActionConfigurations

`func (o *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurations() []MicrosoftGraphDeviceComplianceActionItem`

GetScheduledActionConfigurations returns the ScheduledActionConfigurations field if non-nil, zero value otherwise.

### GetScheduledActionConfigurationsOk

`func (o *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurationsOk() (*[]MicrosoftGraphDeviceComplianceActionItem, bool)`

GetScheduledActionConfigurationsOk returns a tuple with the ScheduledActionConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledActionConfigurations

`func (o *DeviceComplianceScheduledActionForRule) SetScheduledActionConfigurations(v []MicrosoftGraphDeviceComplianceActionItem)`

SetScheduledActionConfigurations sets ScheduledActionConfigurations field to given value.

### HasScheduledActionConfigurations

`func (o *DeviceComplianceScheduledActionForRule) HasScheduledActionConfigurations() bool`

HasScheduledActionConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


