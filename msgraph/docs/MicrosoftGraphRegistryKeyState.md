# MicrosoftGraphRegistryKeyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hive** | Pointer to [**AnyOfmicrosoftGraphRegistryHive**](anyOf&lt;microsoft.graph.registryHive&gt;.md) | A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault. | [optional] 
**Key** | Pointer to **NullableString** | Current (i.e. changed) registry key (excludes HIVE). | [optional] 
**OldKey** | Pointer to **NullableString** | Previous (i.e. before changed) registry key (excludes HIVE). | [optional] 
**OldValueData** | Pointer to **NullableString** | Previous (i.e. before changed) registry key value data (contents). | [optional] 
**OldValueName** | Pointer to **NullableString** | Previous (i.e. before changed) registry key value name. | [optional] 
**Operation** | Pointer to [**AnyOfmicrosoftGraphRegistryOperation**](anyOf&lt;microsoft.graph.registryOperation&gt;.md) | Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete. | [optional] 
**ProcessId** | Pointer to **NullableInt32** | Process ID (PID) of the process that modified the registry key (process details will appear in the alert &#39;processes&#39; collection). | [optional] 
**ValueData** | Pointer to **NullableString** | Current (i.e. changed) registry key value data (contents). | [optional] 
**ValueName** | Pointer to **NullableString** | Current (i.e. changed) registry key value name | [optional] 
**ValueType** | Pointer to [**AnyOfmicrosoftGraphRegistryValueType**](anyOf&lt;microsoft.graph.registryValueType&gt;.md) | Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz. | [optional] 

## Methods

### NewMicrosoftGraphRegistryKeyState

`func NewMicrosoftGraphRegistryKeyState() *MicrosoftGraphRegistryKeyState`

NewMicrosoftGraphRegistryKeyState instantiates a new MicrosoftGraphRegistryKeyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRegistryKeyStateWithDefaults

`func NewMicrosoftGraphRegistryKeyStateWithDefaults() *MicrosoftGraphRegistryKeyState`

NewMicrosoftGraphRegistryKeyStateWithDefaults instantiates a new MicrosoftGraphRegistryKeyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHive

`func (o *MicrosoftGraphRegistryKeyState) GetHive() AnyOfmicrosoftGraphRegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *MicrosoftGraphRegistryKeyState) GetHiveOk() (*AnyOfmicrosoftGraphRegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *MicrosoftGraphRegistryKeyState) SetHive(v AnyOfmicrosoftGraphRegistryHive)`

SetHive sets Hive field to given value.

### HasHive

`func (o *MicrosoftGraphRegistryKeyState) HasHive() bool`

HasHive returns a boolean if a field has been set.

### SetHiveNil

`func (o *MicrosoftGraphRegistryKeyState) SetHiveNil(b bool)`

 SetHiveNil sets the value for Hive to be an explicit nil

### UnsetHive
`func (o *MicrosoftGraphRegistryKeyState) UnsetHive()`

UnsetHive ensures that no value is present for Hive, not even an explicit nil
### GetKey

`func (o *MicrosoftGraphRegistryKeyState) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MicrosoftGraphRegistryKeyState) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MicrosoftGraphRegistryKeyState) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MicrosoftGraphRegistryKeyState) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MicrosoftGraphRegistryKeyState) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MicrosoftGraphRegistryKeyState) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetOldKey

`func (o *MicrosoftGraphRegistryKeyState) GetOldKey() string`

GetOldKey returns the OldKey field if non-nil, zero value otherwise.

### GetOldKeyOk

`func (o *MicrosoftGraphRegistryKeyState) GetOldKeyOk() (*string, bool)`

GetOldKeyOk returns a tuple with the OldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldKey

`func (o *MicrosoftGraphRegistryKeyState) SetOldKey(v string)`

SetOldKey sets OldKey field to given value.

### HasOldKey

`func (o *MicrosoftGraphRegistryKeyState) HasOldKey() bool`

HasOldKey returns a boolean if a field has been set.

### SetOldKeyNil

`func (o *MicrosoftGraphRegistryKeyState) SetOldKeyNil(b bool)`

 SetOldKeyNil sets the value for OldKey to be an explicit nil

### UnsetOldKey
`func (o *MicrosoftGraphRegistryKeyState) UnsetOldKey()`

UnsetOldKey ensures that no value is present for OldKey, not even an explicit nil
### GetOldValueData

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueData() string`

GetOldValueData returns the OldValueData field if non-nil, zero value otherwise.

### GetOldValueDataOk

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueDataOk() (*string, bool)`

GetOldValueDataOk returns a tuple with the OldValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValueData

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueData(v string)`

SetOldValueData sets OldValueData field to given value.

### HasOldValueData

`func (o *MicrosoftGraphRegistryKeyState) HasOldValueData() bool`

HasOldValueData returns a boolean if a field has been set.

### SetOldValueDataNil

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueDataNil(b bool)`

 SetOldValueDataNil sets the value for OldValueData to be an explicit nil

### UnsetOldValueData
`func (o *MicrosoftGraphRegistryKeyState) UnsetOldValueData()`

UnsetOldValueData ensures that no value is present for OldValueData, not even an explicit nil
### GetOldValueName

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueName() string`

GetOldValueName returns the OldValueName field if non-nil, zero value otherwise.

### GetOldValueNameOk

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueNameOk() (*string, bool)`

GetOldValueNameOk returns a tuple with the OldValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValueName

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueName(v string)`

SetOldValueName sets OldValueName field to given value.

### HasOldValueName

`func (o *MicrosoftGraphRegistryKeyState) HasOldValueName() bool`

HasOldValueName returns a boolean if a field has been set.

### SetOldValueNameNil

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueNameNil(b bool)`

 SetOldValueNameNil sets the value for OldValueName to be an explicit nil

### UnsetOldValueName
`func (o *MicrosoftGraphRegistryKeyState) UnsetOldValueName()`

UnsetOldValueName ensures that no value is present for OldValueName, not even an explicit nil
### GetOperation

`func (o *MicrosoftGraphRegistryKeyState) GetOperation() AnyOfmicrosoftGraphRegistryOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MicrosoftGraphRegistryKeyState) GetOperationOk() (*AnyOfmicrosoftGraphRegistryOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MicrosoftGraphRegistryKeyState) SetOperation(v AnyOfmicrosoftGraphRegistryOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MicrosoftGraphRegistryKeyState) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *MicrosoftGraphRegistryKeyState) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *MicrosoftGraphRegistryKeyState) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetProcessId

`func (o *MicrosoftGraphRegistryKeyState) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *MicrosoftGraphRegistryKeyState) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *MicrosoftGraphRegistryKeyState) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *MicrosoftGraphRegistryKeyState) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### SetProcessIdNil

`func (o *MicrosoftGraphRegistryKeyState) SetProcessIdNil(b bool)`

 SetProcessIdNil sets the value for ProcessId to be an explicit nil

### UnsetProcessId
`func (o *MicrosoftGraphRegistryKeyState) UnsetProcessId()`

UnsetProcessId ensures that no value is present for ProcessId, not even an explicit nil
### GetValueData

`func (o *MicrosoftGraphRegistryKeyState) GetValueData() string`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *MicrosoftGraphRegistryKeyState) GetValueDataOk() (*string, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *MicrosoftGraphRegistryKeyState) SetValueData(v string)`

SetValueData sets ValueData field to given value.

### HasValueData

`func (o *MicrosoftGraphRegistryKeyState) HasValueData() bool`

HasValueData returns a boolean if a field has been set.

### SetValueDataNil

`func (o *MicrosoftGraphRegistryKeyState) SetValueDataNil(b bool)`

 SetValueDataNil sets the value for ValueData to be an explicit nil

### UnsetValueData
`func (o *MicrosoftGraphRegistryKeyState) UnsetValueData()`

UnsetValueData ensures that no value is present for ValueData, not even an explicit nil
### GetValueName

`func (o *MicrosoftGraphRegistryKeyState) GetValueName() string`

GetValueName returns the ValueName field if non-nil, zero value otherwise.

### GetValueNameOk

`func (o *MicrosoftGraphRegistryKeyState) GetValueNameOk() (*string, bool)`

GetValueNameOk returns a tuple with the ValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueName

`func (o *MicrosoftGraphRegistryKeyState) SetValueName(v string)`

SetValueName sets ValueName field to given value.

### HasValueName

`func (o *MicrosoftGraphRegistryKeyState) HasValueName() bool`

HasValueName returns a boolean if a field has been set.

### SetValueNameNil

`func (o *MicrosoftGraphRegistryKeyState) SetValueNameNil(b bool)`

 SetValueNameNil sets the value for ValueName to be an explicit nil

### UnsetValueName
`func (o *MicrosoftGraphRegistryKeyState) UnsetValueName()`

UnsetValueName ensures that no value is present for ValueName, not even an explicit nil
### GetValueType

`func (o *MicrosoftGraphRegistryKeyState) GetValueType() AnyOfmicrosoftGraphRegistryValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *MicrosoftGraphRegistryKeyState) GetValueTypeOk() (*AnyOfmicrosoftGraphRegistryValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *MicrosoftGraphRegistryKeyState) SetValueType(v AnyOfmicrosoftGraphRegistryValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *MicrosoftGraphRegistryKeyState) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *MicrosoftGraphRegistryKeyState) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *MicrosoftGraphRegistryKeyState) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


