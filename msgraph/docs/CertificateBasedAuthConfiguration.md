# CertificateBasedAuthConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateAuthorities** | Pointer to [**[]MicrosoftGraphCertificateAuthority**](MicrosoftGraphCertificateAuthority.md) | Collection of certificate authorities which creates a trusted certificate chain. | [optional] 

## Methods

### NewCertificateBasedAuthConfiguration

`func NewCertificateBasedAuthConfiguration() *CertificateBasedAuthConfiguration`

NewCertificateBasedAuthConfiguration instantiates a new CertificateBasedAuthConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateBasedAuthConfigurationWithDefaults

`func NewCertificateBasedAuthConfigurationWithDefaults() *CertificateBasedAuthConfiguration`

NewCertificateBasedAuthConfigurationWithDefaults instantiates a new CertificateBasedAuthConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthorities

`func (o *CertificateBasedAuthConfiguration) GetCertificateAuthorities() []MicrosoftGraphCertificateAuthority`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *CertificateBasedAuthConfiguration) GetCertificateAuthoritiesOk() (*[]MicrosoftGraphCertificateAuthority, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *CertificateBasedAuthConfiguration) SetCertificateAuthorities(v []MicrosoftGraphCertificateAuthority)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *CertificateBasedAuthConfiguration) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


