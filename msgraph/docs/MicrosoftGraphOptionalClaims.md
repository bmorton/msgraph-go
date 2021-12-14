# MicrosoftGraphOptionalClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to [**[]AnyOfmicrosoftGraphOptionalClaim**](AnyOfmicrosoftGraphOptionalClaim.md) | The optional claims returned in the JWT access token. | [optional] 
**IdToken** | Pointer to [**[]AnyOfmicrosoftGraphOptionalClaim**](AnyOfmicrosoftGraphOptionalClaim.md) | The optional claims returned in the JWT ID token. | [optional] 
**Saml2Token** | Pointer to [**[]AnyOfmicrosoftGraphOptionalClaim**](AnyOfmicrosoftGraphOptionalClaim.md) | The optional claims returned in the SAML token. | [optional] 

## Methods

### NewMicrosoftGraphOptionalClaims

`func NewMicrosoftGraphOptionalClaims() *MicrosoftGraphOptionalClaims`

NewMicrosoftGraphOptionalClaims instantiates a new MicrosoftGraphOptionalClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOptionalClaimsWithDefaults

`func NewMicrosoftGraphOptionalClaimsWithDefaults() *MicrosoftGraphOptionalClaims`

NewMicrosoftGraphOptionalClaimsWithDefaults instantiates a new MicrosoftGraphOptionalClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *MicrosoftGraphOptionalClaims) GetAccessToken() []*AnyOfmicrosoftGraphOptionalClaim`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *MicrosoftGraphOptionalClaims) GetAccessTokenOk() (*[]*AnyOfmicrosoftGraphOptionalClaim, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *MicrosoftGraphOptionalClaims) SetAccessToken(v []*AnyOfmicrosoftGraphOptionalClaim)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *MicrosoftGraphOptionalClaims) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetIdToken

`func (o *MicrosoftGraphOptionalClaims) GetIdToken() []*AnyOfmicrosoftGraphOptionalClaim`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *MicrosoftGraphOptionalClaims) GetIdTokenOk() (*[]*AnyOfmicrosoftGraphOptionalClaim, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *MicrosoftGraphOptionalClaims) SetIdToken(v []*AnyOfmicrosoftGraphOptionalClaim)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *MicrosoftGraphOptionalClaims) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetSaml2Token

`func (o *MicrosoftGraphOptionalClaims) GetSaml2Token() []*AnyOfmicrosoftGraphOptionalClaim`

GetSaml2Token returns the Saml2Token field if non-nil, zero value otherwise.

### GetSaml2TokenOk

`func (o *MicrosoftGraphOptionalClaims) GetSaml2TokenOk() (*[]*AnyOfmicrosoftGraphOptionalClaim, bool)`

GetSaml2TokenOk returns a tuple with the Saml2Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml2Token

`func (o *MicrosoftGraphOptionalClaims) SetSaml2Token(v []*AnyOfmicrosoftGraphOptionalClaim)`

SetSaml2Token sets Saml2Token field to given value.

### HasSaml2Token

`func (o *MicrosoftGraphOptionalClaims) HasSaml2Token() bool`

HasSaml2Token returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


