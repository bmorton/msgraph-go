# MicrosoftGraphDeviceCompliancePolicyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the policy for this policyBase | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Platform type that the policy applies to | [optional] 
**SettingCount** | Pointer to **int32** | Count of how many setting a policy holds | [optional] 
**SettingStates** | Pointer to [**[]AnyOfmicrosoftGraphDeviceCompliancePolicySettingState**](AnyOfmicrosoftGraphDeviceCompliancePolicySettingState.md) |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the policy | [optional] 
**Version** | Pointer to **int32** | The version of the policy | [optional] 

## Methods

### NewMicrosoftGraphDeviceCompliancePolicyState

`func NewMicrosoftGraphDeviceCompliancePolicyState() *MicrosoftGraphDeviceCompliancePolicyState`

NewMicrosoftGraphDeviceCompliancePolicyState instantiates a new MicrosoftGraphDeviceCompliancePolicyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceCompliancePolicyStateWithDefaults

`func NewMicrosoftGraphDeviceCompliancePolicyStateWithDefaults() *MicrosoftGraphDeviceCompliancePolicyState`

NewMicrosoftGraphDeviceCompliancePolicyStateWithDefaults instantiates a new MicrosoftGraphDeviceCompliancePolicyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDeviceCompliancePolicyState) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetPlatformTypeOk() (*AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformTypeNil

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetPlatformTypeNil(b bool)`

 SetPlatformTypeNil sets the value for PlatformType to be an explicit nil

### UnsetPlatformType
`func (o *MicrosoftGraphDeviceCompliancePolicyState) UnsetPlatformType()`

UnsetPlatformType ensures that no value is present for PlatformType, not even an explicit nil
### GetSettingCount

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingCount() int32`

GetSettingCount returns the SettingCount field if non-nil, zero value otherwise.

### GetSettingCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingCountOk() (*int32, bool)`

GetSettingCountOk returns a tuple with the SettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCount

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetSettingCount(v int32)`

SetSettingCount sets SettingCount field to given value.

### HasSettingCount

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasSettingCount() bool`

HasSettingCount returns a boolean if a field has been set.

### GetSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingStates() []*AnyOfmicrosoftGraphDeviceCompliancePolicySettingState`

GetSettingStates returns the SettingStates field if non-nil, zero value otherwise.

### GetSettingStatesOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingStatesOk() (*[]*AnyOfmicrosoftGraphDeviceCompliancePolicySettingState, bool)`

GetSettingStatesOk returns a tuple with the SettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetSettingStates(v []*AnyOfmicrosoftGraphDeviceCompliancePolicySettingState)`

SetSettingStates sets SettingStates field to given value.

### HasSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasSettingStates() bool`

HasSettingStates returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphDeviceCompliancePolicyState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


