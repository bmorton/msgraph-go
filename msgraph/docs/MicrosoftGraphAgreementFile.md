# MicrosoftGraphAgreementFile

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
**Localizations** | Pointer to [**[]MicrosoftGraphAgreementFileLocalization**](MicrosoftGraphAgreementFileLocalization.md) |  | [optional] 

## Methods

### NewMicrosoftGraphAgreementFile

`func NewMicrosoftGraphAgreementFile() *MicrosoftGraphAgreementFile`

NewMicrosoftGraphAgreementFile instantiates a new MicrosoftGraphAgreementFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAgreementFileWithDefaults

`func NewMicrosoftGraphAgreementFileWithDefaults() *MicrosoftGraphAgreementFile`

NewMicrosoftGraphAgreementFileWithDefaults instantiates a new MicrosoftGraphAgreementFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAgreementFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAgreementFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAgreementFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAgreementFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphAgreementFile) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAgreementFile) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAgreementFile) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAgreementFile) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAgreementFile) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAgreementFile) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAgreementFile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAgreementFile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAgreementFile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAgreementFile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAgreementFile) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAgreementFile) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFileData

`func (o *MicrosoftGraphAgreementFile) GetFileData() AnyOfmicrosoftGraphAgreementFileData`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *MicrosoftGraphAgreementFile) GetFileDataOk() (*AnyOfmicrosoftGraphAgreementFileData, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *MicrosoftGraphAgreementFile) SetFileData(v AnyOfmicrosoftGraphAgreementFileData)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *MicrosoftGraphAgreementFile) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *MicrosoftGraphAgreementFile) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *MicrosoftGraphAgreementFile) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetFileName

`func (o *MicrosoftGraphAgreementFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphAgreementFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MicrosoftGraphAgreementFile) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MicrosoftGraphAgreementFile) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MicrosoftGraphAgreementFile) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MicrosoftGraphAgreementFile) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetIsDefault

`func (o *MicrosoftGraphAgreementFile) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphAgreementFile) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphAgreementFile) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphAgreementFile) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *MicrosoftGraphAgreementFile) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *MicrosoftGraphAgreementFile) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsMajorVersion

`func (o *MicrosoftGraphAgreementFile) GetIsMajorVersion() bool`

GetIsMajorVersion returns the IsMajorVersion field if non-nil, zero value otherwise.

### GetIsMajorVersionOk

`func (o *MicrosoftGraphAgreementFile) GetIsMajorVersionOk() (*bool, bool)`

GetIsMajorVersionOk returns a tuple with the IsMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorVersion

`func (o *MicrosoftGraphAgreementFile) SetIsMajorVersion(v bool)`

SetIsMajorVersion sets IsMajorVersion field to given value.

### HasIsMajorVersion

`func (o *MicrosoftGraphAgreementFile) HasIsMajorVersion() bool`

HasIsMajorVersion returns a boolean if a field has been set.

### SetIsMajorVersionNil

`func (o *MicrosoftGraphAgreementFile) SetIsMajorVersionNil(b bool)`

 SetIsMajorVersionNil sets the value for IsMajorVersion to be an explicit nil

### UnsetIsMajorVersion
`func (o *MicrosoftGraphAgreementFile) UnsetIsMajorVersion()`

UnsetIsMajorVersion ensures that no value is present for IsMajorVersion, not even an explicit nil
### GetLanguage

`func (o *MicrosoftGraphAgreementFile) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphAgreementFile) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MicrosoftGraphAgreementFile) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MicrosoftGraphAgreementFile) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *MicrosoftGraphAgreementFile) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *MicrosoftGraphAgreementFile) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetLocalizations

`func (o *MicrosoftGraphAgreementFile) GetLocalizations() []MicrosoftGraphAgreementFileLocalization`

GetLocalizations returns the Localizations field if non-nil, zero value otherwise.

### GetLocalizationsOk

`func (o *MicrosoftGraphAgreementFile) GetLocalizationsOk() (*[]MicrosoftGraphAgreementFileLocalization, bool)`

GetLocalizationsOk returns a tuple with the Localizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizations

`func (o *MicrosoftGraphAgreementFile) SetLocalizations(v []MicrosoftGraphAgreementFileLocalization)`

SetLocalizations sets Localizations field to given value.

### HasLocalizations

`func (o *MicrosoftGraphAgreementFile) HasLocalizations() bool`

HasLocalizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


