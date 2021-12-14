# MicrosoftGraphDeviceConfigurationSettingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **NullableString** | Current value of setting on device | [optional] 
**ErrorCode** | Pointer to **int64** | Error code for the setting | [optional] 
**ErrorDescription** | Pointer to **NullableString** | Error description | [optional] 
**InstanceDisplayName** | Pointer to **NullableString** | Name of setting instance that is being reported. | [optional] 
**Setting** | Pointer to **NullableString** | The setting that is being reported | [optional] 
**SettingName** | Pointer to **NullableString** | Localized/user friendly setting name that is being reported | [optional] 
**Sources** | Pointer to [**[]AnyOfmicrosoftGraphSettingSource**](AnyOfmicrosoftGraphSettingSource.md) | Contributing policies | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned. | [optional] 
**UserEmail** | Pointer to **NullableString** | UserEmail | [optional] 
**UserId** | Pointer to **NullableString** | UserId | [optional] 
**UserName** | Pointer to **NullableString** | UserName | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | UserPrincipalName. | [optional] 

## Methods

### NewMicrosoftGraphDeviceConfigurationSettingState

`func NewMicrosoftGraphDeviceConfigurationSettingState() *MicrosoftGraphDeviceConfigurationSettingState`

NewMicrosoftGraphDeviceConfigurationSettingState instantiates a new MicrosoftGraphDeviceConfigurationSettingState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceConfigurationSettingStateWithDefaults

`func NewMicrosoftGraphDeviceConfigurationSettingStateWithDefaults() *MicrosoftGraphDeviceConfigurationSettingState`

NewMicrosoftGraphDeviceConfigurationSettingStateWithDefaults instantiates a new MicrosoftGraphDeviceConfigurationSettingState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
### GetErrorCode

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil
### GetInstanceDisplayName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetInstanceDisplayName() string`

GetInstanceDisplayName returns the InstanceDisplayName field if non-nil, zero value otherwise.

### GetInstanceDisplayNameOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetInstanceDisplayNameOk() (*string, bool)`

GetInstanceDisplayNameOk returns a tuple with the InstanceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDisplayName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetInstanceDisplayName(v string)`

SetInstanceDisplayName sets InstanceDisplayName field to given value.

### HasInstanceDisplayName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasInstanceDisplayName() bool`

HasInstanceDisplayName returns a boolean if a field has been set.

### SetInstanceDisplayNameNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetInstanceDisplayNameNil(b bool)`

 SetInstanceDisplayNameNil sets the value for InstanceDisplayName to be an explicit nil

### UnsetInstanceDisplayName
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetInstanceDisplayName()`

UnsetInstanceDisplayName ensures that no value is present for InstanceDisplayName, not even an explicit nil
### GetSetting

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSetting(v string)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetSettingName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetSources

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSources() []*AnyOfmicrosoftGraphSettingSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSourcesOk() (*[]*AnyOfmicrosoftGraphSettingSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSources(v []*AnyOfmicrosoftGraphSettingSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUserEmail

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphDeviceConfigurationSettingState) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


