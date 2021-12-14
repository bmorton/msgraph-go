# MicrosoftGraphWorkforceIntegrationEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol**](anyOf&lt;microsoft.graph.workforceIntegrationEncryptionProtocol&gt;.md) | Possible values are: sharedSecret, unknownFutureValue. | [optional] 
**Secret** | Pointer to **NullableString** | Encryption shared secret. | [optional] 

## Methods

### NewMicrosoftGraphWorkforceIntegrationEncryption

`func NewMicrosoftGraphWorkforceIntegrationEncryption() *MicrosoftGraphWorkforceIntegrationEncryption`

NewMicrosoftGraphWorkforceIntegrationEncryption instantiates a new MicrosoftGraphWorkforceIntegrationEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkforceIntegrationEncryptionWithDefaults

`func NewMicrosoftGraphWorkforceIntegrationEncryptionWithDefaults() *MicrosoftGraphWorkforceIntegrationEncryption`

NewMicrosoftGraphWorkforceIntegrationEncryptionWithDefaults instantiates a new MicrosoftGraphWorkforceIntegrationEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetProtocol() AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetProtocolOk() (*AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetProtocol(v AnyOfmicrosoftGraphWorkforceIntegrationEncryptionProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *MicrosoftGraphWorkforceIntegrationEncryption) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSecret

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *MicrosoftGraphWorkforceIntegrationEncryption) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *MicrosoftGraphWorkforceIntegrationEncryption) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


