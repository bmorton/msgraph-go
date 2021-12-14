# MicrosoftGraphUpdateWindowsDeviceAccountActionParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarSyncEnabled** | Pointer to **NullableBool** | Not yet documented | [optional] 
**DeviceAccount** | Pointer to [**AnyOfmicrosoftGraphWindowsDeviceAccount**](anyOf&lt;microsoft.graph.windowsDeviceAccount&gt;.md) | Not yet documented | [optional] 
**DeviceAccountEmail** | Pointer to **NullableString** | Not yet documented | [optional] 
**ExchangeServer** | Pointer to **NullableString** | Not yet documented | [optional] 
**PasswordRotationEnabled** | Pointer to **NullableBool** | Not yet documented | [optional] 
**SessionInitiationProtocalAddress** | Pointer to **NullableString** | Not yet documented | [optional] 

## Methods

### NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameter

`func NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameter() *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter`

NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameter instantiates a new MicrosoftGraphUpdateWindowsDeviceAccountActionParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameterWithDefaults

`func NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameterWithDefaults() *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter`

NewMicrosoftGraphUpdateWindowsDeviceAccountActionParameterWithDefaults instantiates a new MicrosoftGraphUpdateWindowsDeviceAccountActionParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarSyncEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled() bool`

GetCalendarSyncEnabled returns the CalendarSyncEnabled field if non-nil, zero value otherwise.

### GetCalendarSyncEnabledOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabledOk() (*bool, bool)`

GetCalendarSyncEnabledOk returns a tuple with the CalendarSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarSyncEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(v bool)`

SetCalendarSyncEnabled sets CalendarSyncEnabled field to given value.

### HasCalendarSyncEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasCalendarSyncEnabled() bool`

HasCalendarSyncEnabled returns a boolean if a field has been set.

### SetCalendarSyncEnabledNil

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabledNil(b bool)`

 SetCalendarSyncEnabledNil sets the value for CalendarSyncEnabled to be an explicit nil

### UnsetCalendarSyncEnabled
`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetCalendarSyncEnabled()`

UnsetCalendarSyncEnabled ensures that no value is present for CalendarSyncEnabled, not even an explicit nil
### GetDeviceAccount

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccount() AnyOfmicrosoftGraphWindowsDeviceAccount`

GetDeviceAccount returns the DeviceAccount field if non-nil, zero value otherwise.

### GetDeviceAccountOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountOk() (*AnyOfmicrosoftGraphWindowsDeviceAccount, bool)`

GetDeviceAccountOk returns a tuple with the DeviceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccount

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(v AnyOfmicrosoftGraphWindowsDeviceAccount)`

SetDeviceAccount sets DeviceAccount field to given value.

### HasDeviceAccount

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasDeviceAccount() bool`

HasDeviceAccount returns a boolean if a field has been set.

### SetDeviceAccountNil

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountNil(b bool)`

 SetDeviceAccountNil sets the value for DeviceAccount to be an explicit nil

### UnsetDeviceAccount
`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetDeviceAccount()`

UnsetDeviceAccount ensures that no value is present for DeviceAccount, not even an explicit nil
### GetDeviceAccountEmail

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail() string`

GetDeviceAccountEmail returns the DeviceAccountEmail field if non-nil, zero value otherwise.

### GetDeviceAccountEmailOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmailOk() (*string, bool)`

GetDeviceAccountEmailOk returns a tuple with the DeviceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAccountEmail

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(v string)`

SetDeviceAccountEmail sets DeviceAccountEmail field to given value.

### HasDeviceAccountEmail

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasDeviceAccountEmail() bool`

HasDeviceAccountEmail returns a boolean if a field has been set.

### SetDeviceAccountEmailNil

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmailNil(b bool)`

 SetDeviceAccountEmailNil sets the value for DeviceAccountEmail to be an explicit nil

### UnsetDeviceAccountEmail
`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetDeviceAccountEmail()`

UnsetDeviceAccountEmail ensures that no value is present for DeviceAccountEmail, not even an explicit nil
### GetExchangeServer

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetExchangeServer() string`

GetExchangeServer returns the ExchangeServer field if non-nil, zero value otherwise.

### GetExchangeServerOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetExchangeServerOk() (*string, bool)`

GetExchangeServerOk returns a tuple with the ExchangeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeServer

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetExchangeServer(v string)`

SetExchangeServer sets ExchangeServer field to given value.

### HasExchangeServer

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasExchangeServer() bool`

HasExchangeServer returns a boolean if a field has been set.

### SetExchangeServerNil

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetExchangeServerNil(b bool)`

 SetExchangeServerNil sets the value for ExchangeServer to be an explicit nil

### UnsetExchangeServer
`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetExchangeServer()`

UnsetExchangeServer ensures that no value is present for ExchangeServer, not even an explicit nil
### GetPasswordRotationEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled() bool`

GetPasswordRotationEnabled returns the PasswordRotationEnabled field if non-nil, zero value otherwise.

### GetPasswordRotationEnabledOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabledOk() (*bool, bool)`

GetPasswordRotationEnabledOk returns a tuple with the PasswordRotationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRotationEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(v bool)`

SetPasswordRotationEnabled sets PasswordRotationEnabled field to given value.

### HasPasswordRotationEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasPasswordRotationEnabled() bool`

HasPasswordRotationEnabled returns a boolean if a field has been set.

### SetPasswordRotationEnabledNil

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabledNil(b bool)`

 SetPasswordRotationEnabledNil sets the value for PasswordRotationEnabled to be an explicit nil

### UnsetPasswordRotationEnabled
`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetPasswordRotationEnabled()`

UnsetPasswordRotationEnabled ensures that no value is present for PasswordRotationEnabled, not even an explicit nil
### GetSessionInitiationProtocalAddress

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress() string`

GetSessionInitiationProtocalAddress returns the SessionInitiationProtocalAddress field if non-nil, zero value otherwise.

### GetSessionInitiationProtocalAddressOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddressOk() (*string, bool)`

GetSessionInitiationProtocalAddressOk returns a tuple with the SessionInitiationProtocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInitiationProtocalAddress

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(v string)`

SetSessionInitiationProtocalAddress sets SessionInitiationProtocalAddress field to given value.

### HasSessionInitiationProtocalAddress

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasSessionInitiationProtocalAddress() bool`

HasSessionInitiationProtocalAddress returns a boolean if a field has been set.

### SetSessionInitiationProtocalAddressNil

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddressNil(b bool)`

 SetSessionInitiationProtocalAddressNil sets the value for SessionInitiationProtocalAddress to be an explicit nil

### UnsetSessionInitiationProtocalAddress
`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) UnsetSessionInitiationProtocalAddress()`

UnsetSessionInitiationProtocalAddress ensures that no value is present for SessionInitiationProtocalAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


