# MicrosoftGraphFileHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashType** | Pointer to [**AnyOfmicrosoftGraphFileHashType**](anyOf&lt;microsoft.graph.fileHashType&gt;.md) | File hash type. Possible values are: unknown, sha1, sha256, md5, authenticodeHash256, lsHash, ctph, peSha1, peSha256. | [optional] 
**HashValue** | Pointer to **NullableString** | Value of the file hash. | [optional] 

## Methods

### NewMicrosoftGraphFileHash

`func NewMicrosoftGraphFileHash() *MicrosoftGraphFileHash`

NewMicrosoftGraphFileHash instantiates a new MicrosoftGraphFileHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFileHashWithDefaults

`func NewMicrosoftGraphFileHashWithDefaults() *MicrosoftGraphFileHash`

NewMicrosoftGraphFileHashWithDefaults instantiates a new MicrosoftGraphFileHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashType

`func (o *MicrosoftGraphFileHash) GetHashType() AnyOfmicrosoftGraphFileHashType`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *MicrosoftGraphFileHash) GetHashTypeOk() (*AnyOfmicrosoftGraphFileHashType, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *MicrosoftGraphFileHash) SetHashType(v AnyOfmicrosoftGraphFileHashType)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *MicrosoftGraphFileHash) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### SetHashTypeNil

`func (o *MicrosoftGraphFileHash) SetHashTypeNil(b bool)`

 SetHashTypeNil sets the value for HashType to be an explicit nil

### UnsetHashType
`func (o *MicrosoftGraphFileHash) UnsetHashType()`

UnsetHashType ensures that no value is present for HashType, not even an explicit nil
### GetHashValue

`func (o *MicrosoftGraphFileHash) GetHashValue() string`

GetHashValue returns the HashValue field if non-nil, zero value otherwise.

### GetHashValueOk

`func (o *MicrosoftGraphFileHash) GetHashValueOk() (*string, bool)`

GetHashValueOk returns a tuple with the HashValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashValue

`func (o *MicrosoftGraphFileHash) SetHashValue(v string)`

SetHashValue sets HashValue field to given value.

### HasHashValue

`func (o *MicrosoftGraphFileHash) HasHashValue() bool`

HasHashValue returns a boolean if a field has been set.

### SetHashValueNil

`func (o *MicrosoftGraphFileHash) SetHashValueNil(b bool)`

 SetHashValueNil sets the value for HashValue to be an explicit nil

### UnsetHashValue
`func (o *MicrosoftGraphFileHash) UnsetHashValue()`

UnsetHashValue ensures that no value is present for HashValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


