# MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time that this Windows Hello for Business key was registered. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the device on which Windows Hello for Business is registered | [optional] 
**KeyStrength** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodKeyStrength**](anyOf&lt;microsoft.graph.authenticationMethodKeyStrength&gt;.md) | Key strength of this Windows Hello for Business key. Possible values are: normal, weak, unknown. | [optional] 
**Device** | Pointer to [**AnyOfmicrosoftGraphDevice**](anyOf&lt;microsoft.graph.device&gt;.md) | The registered device on which this Windows Hello for Business key resides. | [optional] 

## Methods

### NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod

`func NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod() *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod`

NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod instantiates a new MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethodWithDefaults

`func NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethodWithDefaults() *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod`

NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethodWithDefaults instantiates a new MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetKeyStrength

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetKeyStrength() AnyOfmicrosoftGraphAuthenticationMethodKeyStrength`

GetKeyStrength returns the KeyStrength field if non-nil, zero value otherwise.

### GetKeyStrengthOk

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetKeyStrengthOk() (*AnyOfmicrosoftGraphAuthenticationMethodKeyStrength, bool)`

GetKeyStrengthOk returns a tuple with the KeyStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStrength

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetKeyStrength(v AnyOfmicrosoftGraphAuthenticationMethodKeyStrength)`

SetKeyStrength sets KeyStrength field to given value.

### HasKeyStrength

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) HasKeyStrength() bool`

HasKeyStrength returns a boolean if a field has been set.

### SetKeyStrengthNil

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetKeyStrengthNil(b bool)`

 SetKeyStrengthNil sets the value for KeyStrength to be an explicit nil

### UnsetKeyStrength
`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) UnsetKeyStrength()`

UnsetKeyStrength ensures that no value is present for KeyStrength, not even an explicit nil
### GetDevice

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetDevice() AnyOfmicrosoftGraphDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) GetDeviceOk() (*AnyOfmicrosoftGraphDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetDevice(v AnyOfmicrosoftGraphDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


