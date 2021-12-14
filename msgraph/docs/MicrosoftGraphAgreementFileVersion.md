# MicrosoftGraphAgreementFileVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**FileData** | Pointer to [**AnyOfmicrosoftGraphAgreementFileData**](anyOf&lt;microsoft.graph.agreementFileData&gt;.md) |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**IsMajorVersion** | Pointer to **NullableBool** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMicrosoftGraphAgreementFileVersion

`func NewMicrosoftGraphAgreementFileVersion() *MicrosoftGraphAgreementFileVersion`

NewMicrosoftGraphAgreementFileVersion instantiates a new MicrosoftGraphAgreementFileVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAgreementFileVersionWithDefaults

`func NewMicrosoftGraphAgreementFileVersionWithDefaults() *MicrosoftGraphAgreementFileVersion`

NewMicrosoftGraphAgreementFileVersionWithDefaults instantiates a new MicrosoftGraphAgreementFileVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAgreementFileVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAgreementFileVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAgreementFileVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAgreementFileVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphAgreementFileVersion) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAgreementFileVersion) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAgreementFileVersion) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAgreementFileVersion) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAgreementFileVersion) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAgreementFileVersion) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAgreementFileVersion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAgreementFileVersion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAgreementFileVersion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAgreementFileVersion) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAgreementFileVersion) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAgreementFileVersion) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFileData

`func (o *MicrosoftGraphAgreementFileVersion) GetFileData() AnyOfmicrosoftGraphAgreementFileData`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *MicrosoftGraphAgreementFileVersion) GetFileDataOk() (*AnyOfmicrosoftGraphAgreementFileData, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *MicrosoftGraphAgreementFileVersion) SetFileData(v AnyOfmicrosoftGraphAgreementFileData)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *MicrosoftGraphAgreementFileVersion) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *MicrosoftGraphAgreementFileVersion) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *MicrosoftGraphAgreementFileVersion) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetFileName

`func (o *MicrosoftGraphAgreementFileVersion) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphAgreementFileVersion) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MicrosoftGraphAgreementFileVersion) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MicrosoftGraphAgreementFileVersion) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MicrosoftGraphAgreementFileVersion) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MicrosoftGraphAgreementFileVersion) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetIsDefault

`func (o *MicrosoftGraphAgreementFileVersion) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphAgreementFileVersion) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphAgreementFileVersion) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphAgreementFileVersion) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *MicrosoftGraphAgreementFileVersion) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *MicrosoftGraphAgreementFileVersion) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsMajorVersion

`func (o *MicrosoftGraphAgreementFileVersion) GetIsMajorVersion() bool`

GetIsMajorVersion returns the IsMajorVersion field if non-nil, zero value otherwise.

### GetIsMajorVersionOk

`func (o *MicrosoftGraphAgreementFileVersion) GetIsMajorVersionOk() (*bool, bool)`

GetIsMajorVersionOk returns a tuple with the IsMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorVersion

`func (o *MicrosoftGraphAgreementFileVersion) SetIsMajorVersion(v bool)`

SetIsMajorVersion sets IsMajorVersion field to given value.

### HasIsMajorVersion

`func (o *MicrosoftGraphAgreementFileVersion) HasIsMajorVersion() bool`

HasIsMajorVersion returns a boolean if a field has been set.

### SetIsMajorVersionNil

`func (o *MicrosoftGraphAgreementFileVersion) SetIsMajorVersionNil(b bool)`

 SetIsMajorVersionNil sets the value for IsMajorVersion to be an explicit nil

### UnsetIsMajorVersion
`func (o *MicrosoftGraphAgreementFileVersion) UnsetIsMajorVersion()`

UnsetIsMajorVersion ensures that no value is present for IsMajorVersion, not even an explicit nil
### GetLanguage

`func (o *MicrosoftGraphAgreementFileVersion) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphAgreementFileVersion) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MicrosoftGraphAgreementFileVersion) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MicrosoftGraphAgreementFileVersion) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *MicrosoftGraphAgreementFileVersion) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *MicrosoftGraphAgreementFileVersion) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


