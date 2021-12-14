# MicrosoftGraphDeviceCompliancePolicySettingState

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

### NewMicrosoftGraphDeviceCompliancePolicySettingState

`func NewMicrosoftGraphDeviceCompliancePolicySettingState() *MicrosoftGraphDeviceCompliancePolicySettingState`

NewMicrosoftGraphDeviceCompliancePolicySettingState instantiates a new MicrosoftGraphDeviceCompliancePolicySettingState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceCompliancePolicySettingStateWithDefaults

`func NewMicrosoftGraphDeviceCompliancePolicySettingStateWithDefaults() *MicrosoftGraphDeviceCompliancePolicySettingState`

NewMicrosoftGraphDeviceCompliancePolicySettingStateWithDefaults instantiates a new MicrosoftGraphDeviceCompliancePolicySettingState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
### GetErrorCode

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil
### GetInstanceDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetInstanceDisplayName() string`

GetInstanceDisplayName returns the InstanceDisplayName field if non-nil, zero value otherwise.

### GetInstanceDisplayNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetInstanceDisplayNameOk() (*string, bool)`

GetInstanceDisplayNameOk returns a tuple with the InstanceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetInstanceDisplayName(v string)`

SetInstanceDisplayName sets InstanceDisplayName field to given value.

### HasInstanceDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasInstanceDisplayName() bool`

HasInstanceDisplayName returns a boolean if a field has been set.

### SetInstanceDisplayNameNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetInstanceDisplayNameNil(b bool)`

 SetInstanceDisplayNameNil sets the value for InstanceDisplayName to be an explicit nil

### UnsetInstanceDisplayName
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetInstanceDisplayName()`

UnsetInstanceDisplayName ensures that no value is present for InstanceDisplayName, not even an explicit nil
### GetSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSetting(v string)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetSources

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSources() []*AnyOfmicrosoftGraphSettingSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSourcesOk() (*[]*AnyOfmicrosoftGraphSettingSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSources(v []*AnyOfmicrosoftGraphSettingSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetStateOk() (*AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUserEmail

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


