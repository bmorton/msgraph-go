# MicrosoftGraphObjectIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **NullableString** | Specifies the issuer of the identity, for example facebook.com.For local accounts (where signInType is not federated), this property is the local B2C tenant default domain name, for example contoso.onmicrosoft.com.For external users from other Azure AD organization, this will be the domain of the federated organization, for example contoso.com.Supports $filter. 512 character limit. | [optional] 
**IssuerAssignedId** | Pointer to **NullableString** | Specifies the unique identifier assigned to the user by the issuer. The combination of issuer and issuerAssignedId must be unique within the organization. Represents the sign-in name for the user, when signInType is set to emailAddress or userName (also known as local accounts).When signInType is set to: emailAddress, (or a custom string that starts with emailAddress like emailAddress1) issuerAssignedId must be a valid email addressuserName, issuerAssignedId must be a valid local part of an email addressSupports $filter. 100 character limit. | [optional] 
**SignInType** | Pointer to **NullableString** | Specifies the user sign-in types in your directory, such as emailAddress, userName, federated, or userPrincipalName. federated represents a unique identifier for a user from an issuer, that can be in any format chosen by the issuer. Setting or updating a userPrincipalName identity will update the value of the userPrincipalName property on the user object. The validations performed on the userPrincipalName property on the user object, for example, verified domains and acceptable characters, will be performed when setting or updating a userPrincipalName identity. Additional validation is enforced on issuerAssignedId when the sign-in type is set to emailAddress or userName. This property can also be set to any custom string. | [optional] 

## Methods

### NewMicrosoftGraphObjectIdentity

`func NewMicrosoftGraphObjectIdentity() *MicrosoftGraphObjectIdentity`

NewMicrosoftGraphObjectIdentity instantiates a new MicrosoftGraphObjectIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphObjectIdentityWithDefaults

`func NewMicrosoftGraphObjectIdentityWithDefaults() *MicrosoftGraphObjectIdentity`

NewMicrosoftGraphObjectIdentityWithDefaults instantiates a new MicrosoftGraphObjectIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *MicrosoftGraphObjectIdentity) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MicrosoftGraphObjectIdentity) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MicrosoftGraphObjectIdentity) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *MicrosoftGraphObjectIdentity) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *MicrosoftGraphObjectIdentity) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *MicrosoftGraphObjectIdentity) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetIssuerAssignedId

`func (o *MicrosoftGraphObjectIdentity) GetIssuerAssignedId() string`

GetIssuerAssignedId returns the IssuerAssignedId field if non-nil, zero value otherwise.

### GetIssuerAssignedIdOk

`func (o *MicrosoftGraphObjectIdentity) GetIssuerAssignedIdOk() (*string, bool)`

GetIssuerAssignedIdOk returns a tuple with the IssuerAssignedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAssignedId

`func (o *MicrosoftGraphObjectIdentity) SetIssuerAssignedId(v string)`

SetIssuerAssignedId sets IssuerAssignedId field to given value.

### HasIssuerAssignedId

`func (o *MicrosoftGraphObjectIdentity) HasIssuerAssignedId() bool`

HasIssuerAssignedId returns a boolean if a field has been set.

### SetIssuerAssignedIdNil

`func (o *MicrosoftGraphObjectIdentity) SetIssuerAssignedIdNil(b bool)`

 SetIssuerAssignedIdNil sets the value for IssuerAssignedId to be an explicit nil

### UnsetIssuerAssignedId
`func (o *MicrosoftGraphObjectIdentity) UnsetIssuerAssignedId()`

UnsetIssuerAssignedId ensures that no value is present for IssuerAssignedId, not even an explicit nil
### GetSignInType

`func (o *MicrosoftGraphObjectIdentity) GetSignInType() string`

GetSignInType returns the SignInType field if non-nil, zero value otherwise.

### GetSignInTypeOk

`func (o *MicrosoftGraphObjectIdentity) GetSignInTypeOk() (*string, bool)`

GetSignInTypeOk returns a tuple with the SignInType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInType

`func (o *MicrosoftGraphObjectIdentity) SetSignInType(v string)`

SetSignInType sets SignInType field to given value.

### HasSignInType

`func (o *MicrosoftGraphObjectIdentity) HasSignInType() bool`

HasSignInType returns a boolean if a field has been set.

### SetSignInTypeNil

`func (o *MicrosoftGraphObjectIdentity) SetSignInTypeNil(b bool)`

 SetSignInTypeNil sets the value for SignInType to be an explicit nil

### UnsetSignInType
`func (o *MicrosoftGraphObjectIdentity) UnsetSignInType()`

UnsetSignInType ensures that no value is present for SignInType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


