# MicrosoftGraphFileSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHash** | Pointer to [**AnyOfmicrosoftGraphFileHash**](anyOf&lt;microsoft.graph.fileHash&gt;.md) | Complex type containing file hashes (cryptographic and location-sensitive). | [optional] 
**Name** | Pointer to **NullableString** | File name (without path). | [optional] 
**Path** | Pointer to **NullableString** | Full file path of the file/imageFile. | [optional] 
**RiskScore** | Pointer to **NullableString** | Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage. | [optional] 

## Methods

### NewMicrosoftGraphFileSecurityState

`func NewMicrosoftGraphFileSecurityState() *MicrosoftGraphFileSecurityState`

NewMicrosoftGraphFileSecurityState instantiates a new MicrosoftGraphFileSecurityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFileSecurityStateWithDefaults

`func NewMicrosoftGraphFileSecurityStateWithDefaults() *MicrosoftGraphFileSecurityState`

NewMicrosoftGraphFileSecurityStateWithDefaults instantiates a new MicrosoftGraphFileSecurityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHash

`func (o *MicrosoftGraphFileSecurityState) GetFileHash() AnyOfmicrosoftGraphFileHash`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *MicrosoftGraphFileSecurityState) GetFileHashOk() (*AnyOfmicrosoftGraphFileHash, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *MicrosoftGraphFileSecurityState) SetFileHash(v AnyOfmicrosoftGraphFileHash)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *MicrosoftGraphFileSecurityState) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHashNil

`func (o *MicrosoftGraphFileSecurityState) SetFileHashNil(b bool)`

 SetFileHashNil sets the value for FileHash to be an explicit nil

### UnsetFileHash
`func (o *MicrosoftGraphFileSecurityState) UnsetFileHash()`

UnsetFileHash ensures that no value is present for FileHash, not even an explicit nil
### GetName

`func (o *MicrosoftGraphFileSecurityState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphFileSecurityState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphFileSecurityState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphFileSecurityState) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphFileSecurityState) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphFileSecurityState) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *MicrosoftGraphFileSecurityState) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MicrosoftGraphFileSecurityState) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MicrosoftGraphFileSecurityState) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MicrosoftGraphFileSecurityState) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MicrosoftGraphFileSecurityState) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MicrosoftGraphFileSecurityState) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRiskScore

`func (o *MicrosoftGraphFileSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphFileSecurityState) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *MicrosoftGraphFileSecurityState) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *MicrosoftGraphFileSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScoreNil

`func (o *MicrosoftGraphFileSecurityState) SetRiskScoreNil(b bool)`

 SetRiskScoreNil sets the value for RiskScore to be an explicit nil

### UnsetRiskScore
`func (o *MicrosoftGraphFileSecurityState) UnsetRiskScore()`

UnsetRiskScore ensures that no value is present for RiskScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


