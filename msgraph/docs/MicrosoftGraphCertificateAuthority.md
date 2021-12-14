# MicrosoftGraphCertificateAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Required. The base64 encoded string representing the public certificate. | [optional] 
**CertificateRevocationListUrl** | Pointer to **NullableString** | The URL of the certificate revocation list. | [optional] 
**DeltaCertificateRevocationListUrl** | Pointer to **NullableString** | The URL contains the list of all revoked certificates since the last time a full certificate revocaton list was created. | [optional] 
**IsRootAuthority** | Pointer to **bool** | Required. true if the trusted certificate is a root authority, false if the trusted certificate is an intermediate authority. | [optional] 
**Issuer** | Pointer to **string** | The issuer of the certificate, calculated from the certificate value. Read-only. | [optional] 
**IssuerSki** | Pointer to **string** | The subject key identifier of the certificate, calculated from the certificate value. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphCertificateAuthority

`func NewMicrosoftGraphCertificateAuthority() *MicrosoftGraphCertificateAuthority`

NewMicrosoftGraphCertificateAuthority instantiates a new MicrosoftGraphCertificateAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCertificateAuthorityWithDefaults

`func NewMicrosoftGraphCertificateAuthorityWithDefaults() *MicrosoftGraphCertificateAuthority`

NewMicrosoftGraphCertificateAuthorityWithDefaults instantiates a new MicrosoftGraphCertificateAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *MicrosoftGraphCertificateAuthority) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *MicrosoftGraphCertificateAuthority) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *MicrosoftGraphCertificateAuthority) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *MicrosoftGraphCertificateAuthority) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateRevocationListUrl

`func (o *MicrosoftGraphCertificateAuthority) GetCertificateRevocationListUrl() string`

GetCertificateRevocationListUrl returns the CertificateRevocationListUrl field if non-nil, zero value otherwise.

### GetCertificateRevocationListUrlOk

`func (o *MicrosoftGraphCertificateAuthority) GetCertificateRevocationListUrlOk() (*string, bool)`

GetCertificateRevocationListUrlOk returns a tuple with the CertificateRevocationListUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRevocationListUrl

`func (o *MicrosoftGraphCertificateAuthority) SetCertificateRevocationListUrl(v string)`

SetCertificateRevocationListUrl sets CertificateRevocationListUrl field to given value.

### HasCertificateRevocationListUrl

`func (o *MicrosoftGraphCertificateAuthority) HasCertificateRevocationListUrl() bool`

HasCertificateRevocationListUrl returns a boolean if a field has been set.

### SetCertificateRevocationListUrlNil

`func (o *MicrosoftGraphCertificateAuthority) SetCertificateRevocationListUrlNil(b bool)`

 SetCertificateRevocationListUrlNil sets the value for CertificateRevocationListUrl to be an explicit nil

### UnsetCertificateRevocationListUrl
`func (o *MicrosoftGraphCertificateAuthority) UnsetCertificateRevocationListUrl()`

UnsetCertificateRevocationListUrl ensures that no value is present for CertificateRevocationListUrl, not even an explicit nil
### GetDeltaCertificateRevocationListUrl

`func (o *MicrosoftGraphCertificateAuthority) GetDeltaCertificateRevocationListUrl() string`

GetDeltaCertificateRevocationListUrl returns the DeltaCertificateRevocationListUrl field if non-nil, zero value otherwise.

### GetDeltaCertificateRevocationListUrlOk

`func (o *MicrosoftGraphCertificateAuthority) GetDeltaCertificateRevocationListUrlOk() (*string, bool)`

GetDeltaCertificateRevocationListUrlOk returns a tuple with the DeltaCertificateRevocationListUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCertificateRevocationListUrl

`func (o *MicrosoftGraphCertificateAuthority) SetDeltaCertificateRevocationListUrl(v string)`

SetDeltaCertificateRevocationListUrl sets DeltaCertificateRevocationListUrl field to given value.

### HasDeltaCertificateRevocationListUrl

`func (o *MicrosoftGraphCertificateAuthority) HasDeltaCertificateRevocationListUrl() bool`

HasDeltaCertificateRevocationListUrl returns a boolean if a field has been set.

### SetDeltaCertificateRevocationListUrlNil

`func (o *MicrosoftGraphCertificateAuthority) SetDeltaCertificateRevocationListUrlNil(b bool)`

 SetDeltaCertificateRevocationListUrlNil sets the value for DeltaCertificateRevocationListUrl to be an explicit nil

### UnsetDeltaCertificateRevocationListUrl
`func (o *MicrosoftGraphCertificateAuthority) UnsetDeltaCertificateRevocationListUrl()`

UnsetDeltaCertificateRevocationListUrl ensures that no value is present for DeltaCertificateRevocationListUrl, not even an explicit nil
### GetIsRootAuthority

`func (o *MicrosoftGraphCertificateAuthority) GetIsRootAuthority() bool`

GetIsRootAuthority returns the IsRootAuthority field if non-nil, zero value otherwise.

### GetIsRootAuthorityOk

`func (o *MicrosoftGraphCertificateAuthority) GetIsRootAuthorityOk() (*bool, bool)`

GetIsRootAuthorityOk returns a tuple with the IsRootAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootAuthority

`func (o *MicrosoftGraphCertificateAuthority) SetIsRootAuthority(v bool)`

SetIsRootAuthority sets IsRootAuthority field to given value.

### HasIsRootAuthority

`func (o *MicrosoftGraphCertificateAuthority) HasIsRootAuthority() bool`

HasIsRootAuthority returns a boolean if a field has been set.

### GetIssuer

`func (o *MicrosoftGraphCertificateAuthority) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *MicrosoftGraphCertificateAuthority) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *MicrosoftGraphCertificateAuthority) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *MicrosoftGraphCertificateAuthority) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetIssuerSki

`func (o *MicrosoftGraphCertificateAuthority) GetIssuerSki() string`

GetIssuerSki returns the IssuerSki field if non-nil, zero value otherwise.

### GetIssuerSkiOk

`func (o *MicrosoftGraphCertificateAuthority) GetIssuerSkiOk() (*string, bool)`

GetIssuerSkiOk returns a tuple with the IssuerSki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerSki

`func (o *MicrosoftGraphCertificateAuthority) SetIssuerSki(v string)`

SetIssuerSki sets IssuerSki field to given value.

### HasIssuerSki

`func (o *MicrosoftGraphCertificateAuthority) HasIssuerSki() bool`

HasIssuerSki returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


