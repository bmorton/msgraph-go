# MicrosoftGraphAgreementFileLocalization

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
**Versions** | Pointer to [**[]MicrosoftGraphAgreementFileVersion**](MicrosoftGraphAgreementFileVersion.md) |  | [optional] 

## Methods

### NewMicrosoftGraphAgreementFileLocalization

`func NewMicrosoftGraphAgreementFileLocalization() *MicrosoftGraphAgreementFileLocalization`

NewMicrosoftGraphAgreementFileLocalization instantiates a new MicrosoftGraphAgreementFileLocalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAgreementFileLocalizationWithDefaults

`func NewMicrosoftGraphAgreementFileLocalizationWithDefaults() *MicrosoftGraphAgreementFileLocalization`

NewMicrosoftGraphAgreementFileLocalizationWithDefaults instantiates a new MicrosoftGraphAgreementFileLocalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAgreementFileLocalization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAgreementFileLocalization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAgreementFileLocalization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphAgreementFileLocalization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAgreementFileLocalization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAgreementFileLocalization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAgreementFileLocalization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAgreementFileLocalization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAgreementFileLocalization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFileData

`func (o *MicrosoftGraphAgreementFileLocalization) GetFileData() AnyOfmicrosoftGraphAgreementFileData`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetFileDataOk() (*AnyOfmicrosoftGraphAgreementFileData, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *MicrosoftGraphAgreementFileLocalization) SetFileData(v AnyOfmicrosoftGraphAgreementFileData)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *MicrosoftGraphAgreementFileLocalization) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetFileName

`func (o *MicrosoftGraphAgreementFileLocalization) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MicrosoftGraphAgreementFileLocalization) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MicrosoftGraphAgreementFileLocalization) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetIsDefault

`func (o *MicrosoftGraphAgreementFileLocalization) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphAgreementFileLocalization) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphAgreementFileLocalization) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsMajorVersion

`func (o *MicrosoftGraphAgreementFileLocalization) GetIsMajorVersion() bool`

GetIsMajorVersion returns the IsMajorVersion field if non-nil, zero value otherwise.

### GetIsMajorVersionOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetIsMajorVersionOk() (*bool, bool)`

GetIsMajorVersionOk returns a tuple with the IsMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorVersion

`func (o *MicrosoftGraphAgreementFileLocalization) SetIsMajorVersion(v bool)`

SetIsMajorVersion sets IsMajorVersion field to given value.

### HasIsMajorVersion

`func (o *MicrosoftGraphAgreementFileLocalization) HasIsMajorVersion() bool`

HasIsMajorVersion returns a boolean if a field has been set.

### SetIsMajorVersionNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetIsMajorVersionNil(b bool)`

 SetIsMajorVersionNil sets the value for IsMajorVersion to be an explicit nil

### UnsetIsMajorVersion
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetIsMajorVersion()`

UnsetIsMajorVersion ensures that no value is present for IsMajorVersion, not even an explicit nil
### GetLanguage

`func (o *MicrosoftGraphAgreementFileLocalization) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MicrosoftGraphAgreementFileLocalization) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MicrosoftGraphAgreementFileLocalization) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *MicrosoftGraphAgreementFileLocalization) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *MicrosoftGraphAgreementFileLocalization) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetVersions

`func (o *MicrosoftGraphAgreementFileLocalization) GetVersions() []MicrosoftGraphAgreementFileVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MicrosoftGraphAgreementFileLocalization) GetVersionsOk() (*[]MicrosoftGraphAgreementFileVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MicrosoftGraphAgreementFileLocalization) SetVersions(v []MicrosoftGraphAgreementFileVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MicrosoftGraphAgreementFileLocalization) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


