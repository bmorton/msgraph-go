# MicrosoftGraphPasswordProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceChangePasswordNextSignIn** | Pointer to **NullableBool** | true if the user must change her password on the next login; otherwise false. If not set, default is false. NOTE:  For Azure B2C tenants, set to false and instead use custom policies and user flows to force password reset at first sign in. See Force password reset at first logon. | [optional] 
**ForceChangePasswordNextSignInWithMfa** | Pointer to **NullableBool** | If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false. | [optional] 
**Password** | Pointer to **NullableString** | The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required. | [optional] 

## Methods

### NewMicrosoftGraphPasswordProfile

`func NewMicrosoftGraphPasswordProfile() *MicrosoftGraphPasswordProfile`

NewMicrosoftGraphPasswordProfile instantiates a new MicrosoftGraphPasswordProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPasswordProfileWithDefaults

`func NewMicrosoftGraphPasswordProfileWithDefaults() *MicrosoftGraphPasswordProfile`

NewMicrosoftGraphPasswordProfileWithDefaults instantiates a new MicrosoftGraphPasswordProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceChangePasswordNextSignIn

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignIn() bool`

GetForceChangePasswordNextSignIn returns the ForceChangePasswordNextSignIn field if non-nil, zero value otherwise.

### GetForceChangePasswordNextSignInOk

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInOk() (*bool, bool)`

GetForceChangePasswordNextSignInOk returns a tuple with the ForceChangePasswordNextSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangePasswordNextSignIn

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignIn(v bool)`

SetForceChangePasswordNextSignIn sets ForceChangePasswordNextSignIn field to given value.

### HasForceChangePasswordNextSignIn

`func (o *MicrosoftGraphPasswordProfile) HasForceChangePasswordNextSignIn() bool`

HasForceChangePasswordNextSignIn returns a boolean if a field has been set.

### SetForceChangePasswordNextSignInNil

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInNil(b bool)`

 SetForceChangePasswordNextSignInNil sets the value for ForceChangePasswordNextSignIn to be an explicit nil

### UnsetForceChangePasswordNextSignIn
`func (o *MicrosoftGraphPasswordProfile) UnsetForceChangePasswordNextSignIn()`

UnsetForceChangePasswordNextSignIn ensures that no value is present for ForceChangePasswordNextSignIn, not even an explicit nil
### GetForceChangePasswordNextSignInWithMfa

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInWithMfa() bool`

GetForceChangePasswordNextSignInWithMfa returns the ForceChangePasswordNextSignInWithMfa field if non-nil, zero value otherwise.

### GetForceChangePasswordNextSignInWithMfaOk

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInWithMfaOk() (*bool, bool)`

GetForceChangePasswordNextSignInWithMfaOk returns a tuple with the ForceChangePasswordNextSignInWithMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangePasswordNextSignInWithMfa

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInWithMfa(v bool)`

SetForceChangePasswordNextSignInWithMfa sets ForceChangePasswordNextSignInWithMfa field to given value.

### HasForceChangePasswordNextSignInWithMfa

`func (o *MicrosoftGraphPasswordProfile) HasForceChangePasswordNextSignInWithMfa() bool`

HasForceChangePasswordNextSignInWithMfa returns a boolean if a field has been set.

### SetForceChangePasswordNextSignInWithMfaNil

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInWithMfaNil(b bool)`

 SetForceChangePasswordNextSignInWithMfaNil sets the value for ForceChangePasswordNextSignInWithMfa to be an explicit nil

### UnsetForceChangePasswordNextSignInWithMfa
`func (o *MicrosoftGraphPasswordProfile) UnsetForceChangePasswordNextSignInWithMfa()`

UnsetForceChangePasswordNextSignInWithMfa ensures that no value is present for ForceChangePasswordNextSignInWithMfa, not even an explicit nil
### GetPassword

`func (o *MicrosoftGraphPasswordProfile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MicrosoftGraphPasswordProfile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MicrosoftGraphPasswordProfile) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MicrosoftGraphPasswordProfile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MicrosoftGraphPasswordProfile) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MicrosoftGraphPasswordProfile) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


