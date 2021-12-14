# ImportedWindowsAutopilotDeviceIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedUserPrincipalName** | Pointer to **NullableString** | UPN of the user the device will be assigned | [optional] 
**GroupTag** | Pointer to **NullableString** | Group Tag of the Windows autopilot device. | [optional] 
**HardwareIdentifier** | Pointer to **NullableString** | Hardware Blob of the Windows autopilot device. | [optional] 
**ImportId** | Pointer to **NullableString** | The Import Id of the Windows autopilot device. | [optional] 
**ProductKey** | Pointer to **NullableString** | Product Key of the Windows autopilot device. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Serial number of the Windows autopilot device. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState**](anyOf&lt;microsoft.graph.importedWindowsAutopilotDeviceIdentityState&gt;.md) | Current state of the imported device. | [optional] 

## Methods

### NewImportedWindowsAutopilotDeviceIdentity

`func NewImportedWindowsAutopilotDeviceIdentity() *ImportedWindowsAutopilotDeviceIdentity`

NewImportedWindowsAutopilotDeviceIdentity instantiates a new ImportedWindowsAutopilotDeviceIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedWindowsAutopilotDeviceIdentityWithDefaults

`func NewImportedWindowsAutopilotDeviceIdentityWithDefaults() *ImportedWindowsAutopilotDeviceIdentity`

NewImportedWindowsAutopilotDeviceIdentityWithDefaults instantiates a new ImportedWindowsAutopilotDeviceIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedUserPrincipalName

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalName() string`

GetAssignedUserPrincipalName returns the AssignedUserPrincipalName field if non-nil, zero value otherwise.

### GetAssignedUserPrincipalNameOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalNameOk() (*string, bool)`

GetAssignedUserPrincipalNameOk returns a tuple with the AssignedUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserPrincipalName

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalName(v string)`

SetAssignedUserPrincipalName sets AssignedUserPrincipalName field to given value.

### HasAssignedUserPrincipalName

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasAssignedUserPrincipalName() bool`

HasAssignedUserPrincipalName returns a boolean if a field has been set.

### SetAssignedUserPrincipalNameNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalNameNil(b bool)`

 SetAssignedUserPrincipalNameNil sets the value for AssignedUserPrincipalName to be an explicit nil

### UnsetAssignedUserPrincipalName
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetAssignedUserPrincipalName()`

UnsetAssignedUserPrincipalName ensures that no value is present for AssignedUserPrincipalName, not even an explicit nil
### GetGroupTag

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetGroupTag() string`

GetGroupTag returns the GroupTag field if non-nil, zero value otherwise.

### GetGroupTagOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetGroupTagOk() (*string, bool)`

GetGroupTagOk returns a tuple with the GroupTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTag

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetGroupTag(v string)`

SetGroupTag sets GroupTag field to given value.

### HasGroupTag

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasGroupTag() bool`

HasGroupTag returns a boolean if a field has been set.

### SetGroupTagNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetGroupTagNil(b bool)`

 SetGroupTagNil sets the value for GroupTag to be an explicit nil

### UnsetGroupTag
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetGroupTag()`

UnsetGroupTag ensures that no value is present for GroupTag, not even an explicit nil
### GetHardwareIdentifier

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifier() string`

GetHardwareIdentifier returns the HardwareIdentifier field if non-nil, zero value otherwise.

### GetHardwareIdentifierOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifierOk() (*string, bool)`

GetHardwareIdentifierOk returns a tuple with the HardwareIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareIdentifier

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifier(v string)`

SetHardwareIdentifier sets HardwareIdentifier field to given value.

### HasHardwareIdentifier

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasHardwareIdentifier() bool`

HasHardwareIdentifier returns a boolean if a field has been set.

### SetHardwareIdentifierNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifierNil(b bool)`

 SetHardwareIdentifierNil sets the value for HardwareIdentifier to be an explicit nil

### UnsetHardwareIdentifier
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetHardwareIdentifier()`

UnsetHardwareIdentifier ensures that no value is present for HardwareIdentifier, not even an explicit nil
### GetImportId

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetImportId() string`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetImportIdOk() (*string, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetImportId(v string)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportIdNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetImportIdNil(b bool)`

 SetImportIdNil sets the value for ImportId to be an explicit nil

### UnsetImportId
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetImportId()`

UnsetImportId ensures that no value is present for ImportId, not even an explicit nil
### GetProductKey

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKeyNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetProductKeyNil(b bool)`

 SetProductKeyNil sets the value for ProductKey to be an explicit nil

### UnsetProductKey
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetProductKey()`

UnsetProductKey ensures that no value is present for ProductKey, not even an explicit nil
### GetSerialNumber

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetState

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetState() AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImportedWindowsAutopilotDeviceIdentity) GetStateOk() (*AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetState(v AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityState)`

SetState sets State field to given value.

### HasState

`func (o *ImportedWindowsAutopilotDeviceIdentity) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ImportedWindowsAutopilotDeviceIdentity) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ImportedWindowsAutopilotDeviceIdentity) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


