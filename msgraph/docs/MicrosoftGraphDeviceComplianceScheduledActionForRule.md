# MicrosoftGraphDeviceComplianceScheduledActionForRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**RuleName** | Pointer to **NullableString** | Name of the rule which this scheduled action applies to. Currently scheduled actions are created per policy instead of per rule, thus RuleName is always set to default value PasswordRequired. | [optional] 
**ScheduledActionConfigurations** | Pointer to [**[]MicrosoftGraphDeviceComplianceActionItem**](MicrosoftGraphDeviceComplianceActionItem.md) | The list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action. | [optional] 

## Methods

### NewMicrosoftGraphDeviceComplianceScheduledActionForRule

`func NewMicrosoftGraphDeviceComplianceScheduledActionForRule() *MicrosoftGraphDeviceComplianceScheduledActionForRule`

NewMicrosoftGraphDeviceComplianceScheduledActionForRule instantiates a new MicrosoftGraphDeviceComplianceScheduledActionForRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceComplianceScheduledActionForRuleWithDefaults

`func NewMicrosoftGraphDeviceComplianceScheduledActionForRuleWithDefaults() *MicrosoftGraphDeviceComplianceScheduledActionForRule`

NewMicrosoftGraphDeviceComplianceScheduledActionForRuleWithDefaults instantiates a new MicrosoftGraphDeviceComplianceScheduledActionForRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRuleName

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleNameNil

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetRuleNameNil(b bool)`

 SetRuleNameNil sets the value for RuleName to be an explicit nil

### UnsetRuleName
`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) UnsetRuleName()`

UnsetRuleName ensures that no value is present for RuleName, not even an explicit nil
### GetScheduledActionConfigurations

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetScheduledActionConfigurations() []MicrosoftGraphDeviceComplianceActionItem`

GetScheduledActionConfigurations returns the ScheduledActionConfigurations field if non-nil, zero value otherwise.

### GetScheduledActionConfigurationsOk

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetScheduledActionConfigurationsOk() (*[]MicrosoftGraphDeviceComplianceActionItem, bool)`

GetScheduledActionConfigurationsOk returns a tuple with the ScheduledActionConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledActionConfigurations

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetScheduledActionConfigurations(v []MicrosoftGraphDeviceComplianceActionItem)`

SetScheduledActionConfigurations sets ScheduledActionConfigurations field to given value.

### HasScheduledActionConfigurations

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) HasScheduledActionConfigurations() bool`

HasScheduledActionConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


