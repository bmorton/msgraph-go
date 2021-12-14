# MicrosoftGraphDeviceComplianceSettingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to **time.Time** | The DateTime when device compliance grace period expires | [optional] 
**DeviceId** | Pointer to **NullableString** | The Device Id that is being reported | [optional] 
**DeviceModel** | Pointer to **NullableString** | The device model that is being reported | [optional] 
**DeviceName** | Pointer to **NullableString** | The Device Name that is being reported | [optional] 
**Setting** | Pointer to **NullableString** | The setting class name and property name. | [optional] 
**SettingName** | Pointer to **NullableString** | The Setting Name that is being reported | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned. | [optional] 
**UserEmail** | Pointer to **NullableString** | The User email address that is being reported | [optional] 
**UserId** | Pointer to **NullableString** | The user Id that is being reported | [optional] 
**UserName** | Pointer to **NullableString** | The User Name that is being reported | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The User PrincipalName that is being reported | [optional] 

## Methods

### NewMicrosoftGraphDeviceComplianceSettingState

`func NewMicrosoftGraphDeviceComplianceSettingState() *MicrosoftGraphDeviceComplianceSettingState`

NewMicrosoftGraphDeviceComplianceSettingState instantiates a new MicrosoftGraphDeviceComplianceSettingState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceComplianceSettingStateWithDefaults

`func NewMicrosoftGraphDeviceComplianceSettingStateWithDefaults() *MicrosoftGraphDeviceComplianceSettingState`

NewMicrosoftGraphDeviceComplianceSettingStateWithDefaults instantiates a new MicrosoftGraphDeviceComplianceSettingState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime sets ComplianceGracePeriodExpirationDateTime field to given value.

### HasComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### GetDeviceId

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceModel

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetDeviceName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetSetting

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSetting(v string)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetSettingName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetState

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUserEmail

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphDeviceComplianceSettingState) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


