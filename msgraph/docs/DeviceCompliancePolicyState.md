# DeviceCompliancePolicyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The name of the policy for this policyBase | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Platform type that the policy applies to | [optional] 
**SettingCount** | Pointer to **int32** | Count of how many setting a policy holds | [optional] 
**SettingStates** | Pointer to [**[]AnyOfmicrosoftGraphDeviceCompliancePolicySettingState**](AnyOfmicrosoftGraphDeviceCompliancePolicySettingState.md) |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the policy | [optional] 
**Version** | Pointer to **int32** | The version of the policy | [optional] 

## Methods

### NewDeviceCompliancePolicyState

`func NewDeviceCompliancePolicyState() *DeviceCompliancePolicyState`

NewDeviceCompliancePolicyState instantiates a new DeviceCompliancePolicyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCompliancePolicyStateWithDefaults

`func NewDeviceCompliancePolicyStateWithDefaults() *DeviceCompliancePolicyState`

NewDeviceCompliancePolicyStateWithDefaults instantiates a new DeviceCompliancePolicyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DeviceCompliancePolicyState) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceCompliancePolicyState) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceCompliancePolicyState) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceCompliancePolicyState) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DeviceCompliancePolicyState) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DeviceCompliancePolicyState) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPlatformType

`func (o *DeviceCompliancePolicyState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *DeviceCompliancePolicyState) GetPlatformTypeOk() (*AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *DeviceCompliancePolicyState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *DeviceCompliancePolicyState) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformTypeNil

`func (o *DeviceCompliancePolicyState) SetPlatformTypeNil(b bool)`

 SetPlatformTypeNil sets the value for PlatformType to be an explicit nil

### UnsetPlatformType
`func (o *DeviceCompliancePolicyState) UnsetPlatformType()`

UnsetPlatformType ensures that no value is present for PlatformType, not even an explicit nil
### GetSettingCount

`func (o *DeviceCompliancePolicyState) GetSettingCount() int32`

GetSettingCount returns the SettingCount field if non-nil, zero value otherwise.

### GetSettingCountOk

`func (o *DeviceCompliancePolicyState) GetSettingCountOk() (*int32, bool)`

GetSettingCountOk returns a tuple with the SettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCount

`func (o *DeviceCompliancePolicyState) SetSettingCount(v int32)`

SetSettingCount sets SettingCount field to given value.

### HasSettingCount

`func (o *DeviceCompliancePolicyState) HasSettingCount() bool`

HasSettingCount returns a boolean if a field has been set.

### GetSettingStates

`func (o *DeviceCompliancePolicyState) GetSettingStates() []*AnyOfmicrosoftGraphDeviceCompliancePolicySettingState`

GetSettingStates returns the SettingStates field if non-nil, zero value otherwise.

### GetSettingStatesOk

`func (o *DeviceCompliancePolicyState) GetSettingStatesOk() (*[]*AnyOfmicrosoftGraphDeviceCompliancePolicySettingState, bool)`

GetSettingStatesOk returns a tuple with the SettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingStates

`func (o *DeviceCompliancePolicyState) SetSettingStates(v []*AnyOfmicrosoftGraphDeviceCompliancePolicySettingState)`

SetSettingStates sets SettingStates field to given value.

### HasSettingStates

`func (o *DeviceCompliancePolicyState) HasSettingStates() bool`

HasSettingStates returns a boolean if a field has been set.

### GetState

`func (o *DeviceCompliancePolicyState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeviceCompliancePolicyState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeviceCompliancePolicyState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState sets State field to given value.

### HasState

`func (o *DeviceCompliancePolicyState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *DeviceCompliancePolicyState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *DeviceCompliancePolicyState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetVersion

`func (o *DeviceCompliancePolicyState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceCompliancePolicyState) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceCompliancePolicyState) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceCompliancePolicyState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


