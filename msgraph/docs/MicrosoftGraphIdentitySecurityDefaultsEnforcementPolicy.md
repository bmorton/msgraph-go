# MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** | Description for this policy. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name for this policy. Required. | [optional] 
**IsEnabled** | Pointer to **bool** | If set to true, Azure Active Directory security defaults is enabled for the tenant. | [optional] 

## Methods

### NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy

`func NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy() *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy`

NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy instantiates a new MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicyWithDefaults

`func NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicyWithDefaults() *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy`

NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicyWithDefaults instantiates a new MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsEnabled

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


