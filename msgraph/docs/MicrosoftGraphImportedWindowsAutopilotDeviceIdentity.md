# MicrosoftGraphImportedWindowsAutopilotDeviceIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AssignedUserPrincipalName** | Pointer to **NullableString** | UPN of the user the device will be assigned | [optional] 
**GroupTag** | Pointer to **NullableString** | Group Tag of the Windows autopilot device. | [optional] 
**HardwareIdentifier** | Pointer to **NullableString** | Hardware Blob of the Windows autopilot device. | [optional] 
**ImportId** | Pointer to **NullableString** | The Import Id of the Windows autopilot device. | [optional] 
**ProductKey** | Pointer to **NullableString** | Product Key of the Windows autopilot device. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Serial number of the Windows autopilot device. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState**](anyOf&lt;microsoft.graph.importedWindowsAutopilotDeviceIdentityState&gt;.md) | Current state of the imported device. | [optional] 

## Methods

### NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentity

`func NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentity() *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity`

NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentity instantiates a new MicrosoftGraphImportedWindowsAutopilotDeviceIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityWithDefaults

`func NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityWithDefaults() *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity`

NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityWithDefaults instantiates a new MicrosoftGraphImportedWindowsAutopilotDeviceIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignedUserPrincipalName

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalName() string`

GetAssignedUserPrincipalName returns the AssignedUserPrincipalName field if non-nil, zero value otherwise.

### GetAssignedUserPrincipalNameOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalNameOk() (*string, bool)`

GetAssignedUserPrincipalNameOk returns a tuple with the AssignedUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserPrincipalName

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalName(v string)`

SetAssignedUserPrincipalName sets AssignedUserPrincipalName field to given value.

### HasAssignedUserPrincipalName

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasAssignedUserPrincipalName() bool`

HasAssignedUserPrincipalName returns a boolean if a field has been set.

### SetAssignedUserPrincipalNameNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalNameNil(b bool)`

 SetAssignedUserPrincipalNameNil sets the value for AssignedUserPrincipalName to be an explicit nil

### UnsetAssignedUserPrincipalName
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetAssignedUserPrincipalName()`

UnsetAssignedUserPrincipalName ensures that no value is present for AssignedUserPrincipalName, not even an explicit nil
### GetGroupTag

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetGroupTag() string`

GetGroupTag returns the GroupTag field if non-nil, zero value otherwise.

### GetGroupTagOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetGroupTagOk() (*string, bool)`

GetGroupTagOk returns a tuple with the GroupTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTag

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetGroupTag(v string)`

SetGroupTag sets GroupTag field to given value.

### HasGroupTag

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasGroupTag() bool`

HasGroupTag returns a boolean if a field has been set.

### SetGroupTagNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetGroupTagNil(b bool)`

 SetGroupTagNil sets the value for GroupTag to be an explicit nil

### UnsetGroupTag
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetGroupTag()`

UnsetGroupTag ensures that no value is present for GroupTag, not even an explicit nil
### GetHardwareIdentifier

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifier() string`

GetHardwareIdentifier returns the HardwareIdentifier field if non-nil, zero value otherwise.

### GetHardwareIdentifierOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifierOk() (*string, bool)`

GetHardwareIdentifierOk returns a tuple with the HardwareIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareIdentifier

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifier(v string)`

SetHardwareIdentifier sets HardwareIdentifier field to given value.

### HasHardwareIdentifier

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasHardwareIdentifier() bool`

HasHardwareIdentifier returns a boolean if a field has been set.

### SetHardwareIdentifierNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifierNil(b bool)`

 SetHardwareIdentifierNil sets the value for HardwareIdentifier to be an explicit nil

### UnsetHardwareIdentifier
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetHardwareIdentifier()`

UnsetHardwareIdentifier ensures that no value is present for HardwareIdentifier, not even an explicit nil
### GetImportId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetImportId() string`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetImportIdOk() (*string, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetImportId(v string)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportIdNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetImportIdNil(b bool)`

 SetImportIdNil sets the value for ImportId to be an explicit nil

### UnsetImportId
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetImportId()`

UnsetImportId ensures that no value is present for ImportId, not even an explicit nil
### GetProductKey

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKeyNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetProductKeyNil(b bool)`

 SetProductKeyNil sets the value for ProductKey to be an explicit nil

### UnsetProductKey
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetProductKey()`

UnsetProductKey ensures that no value is present for ProductKey, not even an explicit nil
### GetSerialNumber

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetState

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetState() AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) GetStateOk() (*AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetState(v AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentity) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


