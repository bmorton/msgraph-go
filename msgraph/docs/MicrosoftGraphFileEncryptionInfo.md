# MicrosoftGraphFileEncryptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | Pointer to **NullableString** | The key used to encrypt the file content. | [optional] 
**FileDigest** | Pointer to **NullableString** | The file digest prior to encryption. | [optional] 
**FileDigestAlgorithm** | Pointer to **NullableString** | The file digest algorithm. | [optional] 
**InitializationVector** | Pointer to **NullableString** | The initialization vector used for the encryption algorithm. | [optional] 
**Mac** | Pointer to **NullableString** | The hash of the encrypted file content + IV (content hash). | [optional] 
**MacKey** | Pointer to **NullableString** | The key used to get mac. | [optional] 
**ProfileIdentifier** | Pointer to **NullableString** | The the profile identifier. | [optional] 

## Methods

### NewMicrosoftGraphFileEncryptionInfo

`func NewMicrosoftGraphFileEncryptionInfo() *MicrosoftGraphFileEncryptionInfo`

NewMicrosoftGraphFileEncryptionInfo instantiates a new MicrosoftGraphFileEncryptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFileEncryptionInfoWithDefaults

`func NewMicrosoftGraphFileEncryptionInfoWithDefaults() *MicrosoftGraphFileEncryptionInfo`

NewMicrosoftGraphFileEncryptionInfoWithDefaults instantiates a new MicrosoftGraphFileEncryptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *MicrosoftGraphFileEncryptionInfo) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *MicrosoftGraphFileEncryptionInfo) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *MicrosoftGraphFileEncryptionInfo) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetFileDigest

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigest() string`

GetFileDigest returns the FileDigest field if non-nil, zero value otherwise.

### GetFileDigestOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestOk() (*string, bool)`

GetFileDigestOk returns a tuple with the FileDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDigest

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigest(v string)`

SetFileDigest sets FileDigest field to given value.

### HasFileDigest

`func (o *MicrosoftGraphFileEncryptionInfo) HasFileDigest() bool`

HasFileDigest returns a boolean if a field has been set.

### SetFileDigestNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestNil(b bool)`

 SetFileDigestNil sets the value for FileDigest to be an explicit nil

### UnsetFileDigest
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetFileDigest()`

UnsetFileDigest ensures that no value is present for FileDigest, not even an explicit nil
### GetFileDigestAlgorithm

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestAlgorithm() string`

GetFileDigestAlgorithm returns the FileDigestAlgorithm field if non-nil, zero value otherwise.

### GetFileDigestAlgorithmOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestAlgorithmOk() (*string, bool)`

GetFileDigestAlgorithmOk returns a tuple with the FileDigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDigestAlgorithm

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestAlgorithm(v string)`

SetFileDigestAlgorithm sets FileDigestAlgorithm field to given value.

### HasFileDigestAlgorithm

`func (o *MicrosoftGraphFileEncryptionInfo) HasFileDigestAlgorithm() bool`

HasFileDigestAlgorithm returns a boolean if a field has been set.

### SetFileDigestAlgorithmNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestAlgorithmNil(b bool)`

 SetFileDigestAlgorithmNil sets the value for FileDigestAlgorithm to be an explicit nil

### UnsetFileDigestAlgorithm
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetFileDigestAlgorithm()`

UnsetFileDigestAlgorithm ensures that no value is present for FileDigestAlgorithm, not even an explicit nil
### GetInitializationVector

`func (o *MicrosoftGraphFileEncryptionInfo) GetInitializationVector() string`

GetInitializationVector returns the InitializationVector field if non-nil, zero value otherwise.

### GetInitializationVectorOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetInitializationVectorOk() (*string, bool)`

GetInitializationVectorOk returns a tuple with the InitializationVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationVector

`func (o *MicrosoftGraphFileEncryptionInfo) SetInitializationVector(v string)`

SetInitializationVector sets InitializationVector field to given value.

### HasInitializationVector

`func (o *MicrosoftGraphFileEncryptionInfo) HasInitializationVector() bool`

HasInitializationVector returns a boolean if a field has been set.

### SetInitializationVectorNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetInitializationVectorNil(b bool)`

 SetInitializationVectorNil sets the value for InitializationVector to be an explicit nil

### UnsetInitializationVector
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetInitializationVector()`

UnsetInitializationVector ensures that no value is present for InitializationVector, not even an explicit nil
### GetMac

`func (o *MicrosoftGraphFileEncryptionInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MicrosoftGraphFileEncryptionInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MicrosoftGraphFileEncryptionInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### SetMacNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetMacNil(b bool)`

 SetMacNil sets the value for Mac to be an explicit nil

### UnsetMac
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetMac()`

UnsetMac ensures that no value is present for Mac, not even an explicit nil
### GetMacKey

`func (o *MicrosoftGraphFileEncryptionInfo) GetMacKey() string`

GetMacKey returns the MacKey field if non-nil, zero value otherwise.

### GetMacKeyOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetMacKeyOk() (*string, bool)`

GetMacKeyOk returns a tuple with the MacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKey

`func (o *MicrosoftGraphFileEncryptionInfo) SetMacKey(v string)`

SetMacKey sets MacKey field to given value.

### HasMacKey

`func (o *MicrosoftGraphFileEncryptionInfo) HasMacKey() bool`

HasMacKey returns a boolean if a field has been set.

### SetMacKeyNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetMacKeyNil(b bool)`

 SetMacKeyNil sets the value for MacKey to be an explicit nil

### UnsetMacKey
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetMacKey()`

UnsetMacKey ensures that no value is present for MacKey, not even an explicit nil
### GetProfileIdentifier

`func (o *MicrosoftGraphFileEncryptionInfo) GetProfileIdentifier() string`

GetProfileIdentifier returns the ProfileIdentifier field if non-nil, zero value otherwise.

### GetProfileIdentifierOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetProfileIdentifierOk() (*string, bool)`

GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIdentifier

`func (o *MicrosoftGraphFileEncryptionInfo) SetProfileIdentifier(v string)`

SetProfileIdentifier sets ProfileIdentifier field to given value.

### HasProfileIdentifier

`func (o *MicrosoftGraphFileEncryptionInfo) HasProfileIdentifier() bool`

HasProfileIdentifier returns a boolean if a field has been set.

### SetProfileIdentifierNil

`func (o *MicrosoftGraphFileEncryptionInfo) SetProfileIdentifierNil(b bool)`

 SetProfileIdentifierNil sets the value for ProfileIdentifier to be an explicit nil

### UnsetProfileIdentifier
`func (o *MicrosoftGraphFileEncryptionInfo) UnsetProfileIdentifier()`

UnsetProfileIdentifier ensures that no value is present for ProfileIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


