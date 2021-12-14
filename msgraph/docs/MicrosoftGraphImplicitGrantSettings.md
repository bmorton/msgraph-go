# MicrosoftGraphImplicitGrantSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAccessTokenIssuance** | Pointer to **NullableBool** | Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow. | [optional] 
**EnableIdTokenIssuance** | Pointer to **NullableBool** | Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow. | [optional] 

## Methods

### NewMicrosoftGraphImplicitGrantSettings

`func NewMicrosoftGraphImplicitGrantSettings() *MicrosoftGraphImplicitGrantSettings`

NewMicrosoftGraphImplicitGrantSettings instantiates a new MicrosoftGraphImplicitGrantSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphImplicitGrantSettingsWithDefaults

`func NewMicrosoftGraphImplicitGrantSettingsWithDefaults() *MicrosoftGraphImplicitGrantSettings`

NewMicrosoftGraphImplicitGrantSettingsWithDefaults instantiates a new MicrosoftGraphImplicitGrantSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAccessTokenIssuance

`func (o *MicrosoftGraphImplicitGrantSettings) GetEnableAccessTokenIssuance() bool`

GetEnableAccessTokenIssuance returns the EnableAccessTokenIssuance field if non-nil, zero value otherwise.

### GetEnableAccessTokenIssuanceOk

`func (o *MicrosoftGraphImplicitGrantSettings) GetEnableAccessTokenIssuanceOk() (*bool, bool)`

GetEnableAccessTokenIssuanceOk returns a tuple with the EnableAccessTokenIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccessTokenIssuance

`func (o *MicrosoftGraphImplicitGrantSettings) SetEnableAccessTokenIssuance(v bool)`

SetEnableAccessTokenIssuance sets EnableAccessTokenIssuance field to given value.

### HasEnableAccessTokenIssuance

`func (o *MicrosoftGraphImplicitGrantSettings) HasEnableAccessTokenIssuance() bool`

HasEnableAccessTokenIssuance returns a boolean if a field has been set.

### SetEnableAccessTokenIssuanceNil

`func (o *MicrosoftGraphImplicitGrantSettings) SetEnableAccessTokenIssuanceNil(b bool)`

 SetEnableAccessTokenIssuanceNil sets the value for EnableAccessTokenIssuance to be an explicit nil

### UnsetEnableAccessTokenIssuance
`func (o *MicrosoftGraphImplicitGrantSettings) UnsetEnableAccessTokenIssuance()`

UnsetEnableAccessTokenIssuance ensures that no value is present for EnableAccessTokenIssuance, not even an explicit nil
### GetEnableIdTokenIssuance

`func (o *MicrosoftGraphImplicitGrantSettings) GetEnableIdTokenIssuance() bool`

GetEnableIdTokenIssuance returns the EnableIdTokenIssuance field if non-nil, zero value otherwise.

### GetEnableIdTokenIssuanceOk

`func (o *MicrosoftGraphImplicitGrantSettings) GetEnableIdTokenIssuanceOk() (*bool, bool)`

GetEnableIdTokenIssuanceOk returns a tuple with the EnableIdTokenIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIdTokenIssuance

`func (o *MicrosoftGraphImplicitGrantSettings) SetEnableIdTokenIssuance(v bool)`

SetEnableIdTokenIssuance sets EnableIdTokenIssuance field to given value.

### HasEnableIdTokenIssuance

`func (o *MicrosoftGraphImplicitGrantSettings) HasEnableIdTokenIssuance() bool`

HasEnableIdTokenIssuance returns a boolean if a field has been set.

### SetEnableIdTokenIssuanceNil

`func (o *MicrosoftGraphImplicitGrantSettings) SetEnableIdTokenIssuanceNil(b bool)`

 SetEnableIdTokenIssuanceNil sets the value for EnableIdTokenIssuance to be an explicit nil

### UnsetEnableIdTokenIssuance
`func (o *MicrosoftGraphImplicitGrantSettings) UnsetEnableIdTokenIssuance()`

UnsetEnableIdTokenIssuance ensures that no value is present for EnableIdTokenIssuance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


