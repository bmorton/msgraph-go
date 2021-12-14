# MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In. | [optional] 
**DeviceTag** | Pointer to **NullableString** | Tags containing app metadata. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the device on which this app is registered. | [optional] 
**PhoneAppVersion** | Pointer to **NullableString** | Numerical version of this instance of the Authenticator app. | [optional] 
**Device** | Pointer to [**AnyOfmicrosoftGraphDevice**](anyOf&lt;microsoft.graph.device&gt;.md) | The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In. | [optional] 

## Methods

### NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod

`func NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod() *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod`

NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod instantiates a new MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethodWithDefaults

`func NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethodWithDefaults() *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod`

NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethodWithDefaults instantiates a new MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDeviceTag

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetDeviceTag() string`

GetDeviceTag returns the DeviceTag field if non-nil, zero value otherwise.

### GetDeviceTagOk

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetDeviceTagOk() (*string, bool)`

GetDeviceTagOk returns a tuple with the DeviceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTag

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetDeviceTag(v string)`

SetDeviceTag sets DeviceTag field to given value.

### HasDeviceTag

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) HasDeviceTag() bool`

HasDeviceTag returns a boolean if a field has been set.

### SetDeviceTagNil

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetDeviceTagNil(b bool)`

 SetDeviceTagNil sets the value for DeviceTag to be an explicit nil

### UnsetDeviceTag
`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) UnsetDeviceTag()`

UnsetDeviceTag ensures that no value is present for DeviceTag, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPhoneAppVersion

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersion() string`

GetPhoneAppVersion returns the PhoneAppVersion field if non-nil, zero value otherwise.

### GetPhoneAppVersionOk

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersionOk() (*string, bool)`

GetPhoneAppVersionOk returns a tuple with the PhoneAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneAppVersion

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersion(v string)`

SetPhoneAppVersion sets PhoneAppVersion field to given value.

### HasPhoneAppVersion

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) HasPhoneAppVersion() bool`

HasPhoneAppVersion returns a boolean if a field has been set.

### SetPhoneAppVersionNil

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersionNil(b bool)`

 SetPhoneAppVersionNil sets the value for PhoneAppVersion to be an explicit nil

### UnsetPhoneAppVersion
`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) UnsetPhoneAppVersion()`

UnsetPhoneAppVersion ensures that no value is present for PhoneAppVersion, not even an explicit nil
### GetDevice

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetDevice() AnyOfmicrosoftGraphDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) GetDeviceOk() (*AnyOfmicrosoftGraphDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetDevice(v AnyOfmicrosoftGraphDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


