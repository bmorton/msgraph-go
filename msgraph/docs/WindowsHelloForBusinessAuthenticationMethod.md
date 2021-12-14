# WindowsHelloForBusinessAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The date and time that this Windows Hello for Business key was registered. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the device on which Windows Hello for Business is registered | [optional] 
**KeyStrength** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodKeyStrength**](anyOf&lt;microsoft.graph.authenticationMethodKeyStrength&gt;.md) | Key strength of this Windows Hello for Business key. Possible values are: normal, weak, unknown. | [optional] 
**Device** | Pointer to [**AnyOfmicrosoftGraphDevice**](anyOf&lt;microsoft.graph.device&gt;.md) | The registered device on which this Windows Hello for Business key resides. | [optional] 

## Methods

### NewWindowsHelloForBusinessAuthenticationMethod

`func NewWindowsHelloForBusinessAuthenticationMethod() *WindowsHelloForBusinessAuthenticationMethod`

NewWindowsHelloForBusinessAuthenticationMethod instantiates a new WindowsHelloForBusinessAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsHelloForBusinessAuthenticationMethodWithDefaults

`func NewWindowsHelloForBusinessAuthenticationMethodWithDefaults() *WindowsHelloForBusinessAuthenticationMethod`

NewWindowsHelloForBusinessAuthenticationMethodWithDefaults instantiates a new WindowsHelloForBusinessAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *WindowsHelloForBusinessAuthenticationMethod) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *WindowsHelloForBusinessAuthenticationMethod) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WindowsHelloForBusinessAuthenticationMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WindowsHelloForBusinessAuthenticationMethod) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetKeyStrength

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetKeyStrength() AnyOfmicrosoftGraphAuthenticationMethodKeyStrength`

GetKeyStrength returns the KeyStrength field if non-nil, zero value otherwise.

### GetKeyStrengthOk

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetKeyStrengthOk() (*AnyOfmicrosoftGraphAuthenticationMethodKeyStrength, bool)`

GetKeyStrengthOk returns a tuple with the KeyStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStrength

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetKeyStrength(v AnyOfmicrosoftGraphAuthenticationMethodKeyStrength)`

SetKeyStrength sets KeyStrength field to given value.

### HasKeyStrength

`func (o *WindowsHelloForBusinessAuthenticationMethod) HasKeyStrength() bool`

HasKeyStrength returns a boolean if a field has been set.

### SetKeyStrengthNil

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetKeyStrengthNil(b bool)`

 SetKeyStrengthNil sets the value for KeyStrength to be an explicit nil

### UnsetKeyStrength
`func (o *WindowsHelloForBusinessAuthenticationMethod) UnsetKeyStrength()`

UnsetKeyStrength ensures that no value is present for KeyStrength, not even an explicit nil
### GetDevice

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetDevice() AnyOfmicrosoftGraphDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WindowsHelloForBusinessAuthenticationMethod) GetDeviceOk() (*AnyOfmicrosoftGraphDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetDevice(v AnyOfmicrosoftGraphDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WindowsHelloForBusinessAuthenticationMethod) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *WindowsHelloForBusinessAuthenticationMethod) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *WindowsHelloForBusinessAuthenticationMethod) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


