# MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceErrorCode** | Pointer to **int32** | Device error code reported by Device Directory Service(DDS). | [optional] 
**DeviceErrorName** | Pointer to **NullableString** | Device error name reported by Device Directory Service(DDS). | [optional] 
**DeviceImportStatus** | Pointer to [**AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityImportStatus**](anyOf&lt;microsoft.graph.importedWindowsAutopilotDeviceIdentityImportStatus&gt;.md) | Device status reported by Device Directory Service(DDS). Possible values are: unknown, pending, partial, complete, error. | [optional] 
**DeviceRegistrationId** | Pointer to **NullableString** | Device Registration ID for successfully added device reported by Device Directory Service(DDS). | [optional] 

## Methods

### NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityState

`func NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityState() *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState`

NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityState instantiates a new MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityStateWithDefaults

`func NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityStateWithDefaults() *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState`

NewMicrosoftGraphImportedWindowsAutopilotDeviceIdentityStateWithDefaults instantiates a new MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceErrorCode

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorCode() int32`

GetDeviceErrorCode returns the DeviceErrorCode field if non-nil, zero value otherwise.

### GetDeviceErrorCodeOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorCodeOk() (*int32, bool)`

GetDeviceErrorCodeOk returns a tuple with the DeviceErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrorCode

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorCode(v int32)`

SetDeviceErrorCode sets DeviceErrorCode field to given value.

### HasDeviceErrorCode

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) HasDeviceErrorCode() bool`

HasDeviceErrorCode returns a boolean if a field has been set.

### GetDeviceErrorName

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorName() string`

GetDeviceErrorName returns the DeviceErrorName field if non-nil, zero value otherwise.

### GetDeviceErrorNameOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceErrorNameOk() (*string, bool)`

GetDeviceErrorNameOk returns a tuple with the DeviceErrorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrorName

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorName(v string)`

SetDeviceErrorName sets DeviceErrorName field to given value.

### HasDeviceErrorName

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) HasDeviceErrorName() bool`

HasDeviceErrorName returns a boolean if a field has been set.

### SetDeviceErrorNameNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceErrorNameNil(b bool)`

 SetDeviceErrorNameNil sets the value for DeviceErrorName to be an explicit nil

### UnsetDeviceErrorName
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) UnsetDeviceErrorName()`

UnsetDeviceErrorName ensures that no value is present for DeviceErrorName, not even an explicit nil
### GetDeviceImportStatus

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceImportStatus() AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityImportStatus`

GetDeviceImportStatus returns the DeviceImportStatus field if non-nil, zero value otherwise.

### GetDeviceImportStatusOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceImportStatusOk() (*AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityImportStatus, bool)`

GetDeviceImportStatusOk returns a tuple with the DeviceImportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceImportStatus

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceImportStatus(v AnyOfmicrosoftGraphImportedWindowsAutopilotDeviceIdentityImportStatus)`

SetDeviceImportStatus sets DeviceImportStatus field to given value.

### HasDeviceImportStatus

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) HasDeviceImportStatus() bool`

HasDeviceImportStatus returns a boolean if a field has been set.

### SetDeviceImportStatusNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceImportStatusNil(b bool)`

 SetDeviceImportStatusNil sets the value for DeviceImportStatus to be an explicit nil

### UnsetDeviceImportStatus
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) UnsetDeviceImportStatus()`

UnsetDeviceImportStatus ensures that no value is present for DeviceImportStatus, not even an explicit nil
### GetDeviceRegistrationId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceRegistrationId() string`

GetDeviceRegistrationId returns the DeviceRegistrationId field if non-nil, zero value otherwise.

### GetDeviceRegistrationIdOk

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) GetDeviceRegistrationIdOk() (*string, bool)`

GetDeviceRegistrationIdOk returns a tuple with the DeviceRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceRegistrationId(v string)`

SetDeviceRegistrationId sets DeviceRegistrationId field to given value.

### HasDeviceRegistrationId

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) HasDeviceRegistrationId() bool`

HasDeviceRegistrationId returns a boolean if a field has been set.

### SetDeviceRegistrationIdNil

`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) SetDeviceRegistrationIdNil(b bool)`

 SetDeviceRegistrationIdNil sets the value for DeviceRegistrationId to be an explicit nil

### UnsetDeviceRegistrationId
`func (o *MicrosoftGraphImportedWindowsAutopilotDeviceIdentityState) UnsetDeviceRegistrationId()`

UnsetDeviceRegistrationId ensures that no value is present for DeviceRegistrationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


