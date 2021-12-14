# MicrosoftAuthenticatorAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In. | [optional] 
**DeviceTag** | Pointer to **NullableString** | Tags containing app metadata. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the device on which this app is registered. | [optional] 
**PhoneAppVersion** | Pointer to **NullableString** | Numerical version of this instance of the Authenticator app. | [optional] 
**Device** | Pointer to [**AnyOfmicrosoftGraphDevice**](anyOf&lt;microsoft.graph.device&gt;.md) | The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In. | [optional] 

## Methods

### NewMicrosoftAuthenticatorAuthenticationMethod

`func NewMicrosoftAuthenticatorAuthenticationMethod() *MicrosoftAuthenticatorAuthenticationMethod`

NewMicrosoftAuthenticatorAuthenticationMethod instantiates a new MicrosoftAuthenticatorAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftAuthenticatorAuthenticationMethodWithDefaults

`func NewMicrosoftAuthenticatorAuthenticationMethodWithDefaults() *MicrosoftAuthenticatorAuthenticationMethod`

NewMicrosoftAuthenticatorAuthenticationMethodWithDefaults instantiates a new MicrosoftAuthenticatorAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftAuthenticatorAuthenticationMethod) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftAuthenticatorAuthenticationMethod) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDeviceTag

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetDeviceTag() string`

GetDeviceTag returns the DeviceTag field if non-nil, zero value otherwise.

### GetDeviceTagOk

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetDeviceTagOk() (*string, bool)`

GetDeviceTagOk returns a tuple with the DeviceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTag

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetDeviceTag(v string)`

SetDeviceTag sets DeviceTag field to given value.

### HasDeviceTag

`func (o *MicrosoftAuthenticatorAuthenticationMethod) HasDeviceTag() bool`

HasDeviceTag returns a boolean if a field has been set.

### SetDeviceTagNil

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetDeviceTagNil(b bool)`

 SetDeviceTagNil sets the value for DeviceTag to be an explicit nil

### UnsetDeviceTag
`func (o *MicrosoftAuthenticatorAuthenticationMethod) UnsetDeviceTag()`

UnsetDeviceTag ensures that no value is present for DeviceTag, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftAuthenticatorAuthenticationMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftAuthenticatorAuthenticationMethod) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPhoneAppVersion

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersion() string`

GetPhoneAppVersion returns the PhoneAppVersion field if non-nil, zero value otherwise.

### GetPhoneAppVersionOk

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersionOk() (*string, bool)`

GetPhoneAppVersionOk returns a tuple with the PhoneAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneAppVersion

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersion(v string)`

SetPhoneAppVersion sets PhoneAppVersion field to given value.

### HasPhoneAppVersion

`func (o *MicrosoftAuthenticatorAuthenticationMethod) HasPhoneAppVersion() bool`

HasPhoneAppVersion returns a boolean if a field has been set.

### SetPhoneAppVersionNil

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersionNil(b bool)`

 SetPhoneAppVersionNil sets the value for PhoneAppVersion to be an explicit nil

### UnsetPhoneAppVersion
`func (o *MicrosoftAuthenticatorAuthenticationMethod) UnsetPhoneAppVersion()`

UnsetPhoneAppVersion ensures that no value is present for PhoneAppVersion, not even an explicit nil
### GetDevice

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetDevice() AnyOfmicrosoftGraphDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MicrosoftAuthenticatorAuthenticationMethod) GetDeviceOk() (*AnyOfmicrosoftGraphDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetDevice(v AnyOfmicrosoftGraphDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MicrosoftAuthenticatorAuthenticationMethod) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *MicrosoftAuthenticatorAuthenticationMethod) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *MicrosoftAuthenticatorAuthenticationMethod) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


