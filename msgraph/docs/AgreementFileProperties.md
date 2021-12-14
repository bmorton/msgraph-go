# AgreementFileProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**FileData** | Pointer to [**AnyOfmicrosoftGraphAgreementFileData**](anyOf&lt;microsoft.graph.agreementFileData&gt;.md) |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**IsMajorVersion** | Pointer to **NullableBool** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgreementFileProperties

`func NewAgreementFileProperties() *AgreementFileProperties`

NewAgreementFileProperties instantiates a new AgreementFileProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementFilePropertiesWithDefaults

`func NewAgreementFilePropertiesWithDefaults() *AgreementFileProperties`

NewAgreementFilePropertiesWithDefaults instantiates a new AgreementFileProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *AgreementFileProperties) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AgreementFileProperties) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AgreementFileProperties) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AgreementFileProperties) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *AgreementFileProperties) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *AgreementFileProperties) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDisplayName

`func (o *AgreementFileProperties) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AgreementFileProperties) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AgreementFileProperties) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AgreementFileProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AgreementFileProperties) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AgreementFileProperties) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFileData

`func (o *AgreementFileProperties) GetFileData() AnyOfmicrosoftGraphAgreementFileData`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *AgreementFileProperties) GetFileDataOk() (*AnyOfmicrosoftGraphAgreementFileData, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *AgreementFileProperties) SetFileData(v AnyOfmicrosoftGraphAgreementFileData)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *AgreementFileProperties) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *AgreementFileProperties) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *AgreementFileProperties) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetFileName

`func (o *AgreementFileProperties) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AgreementFileProperties) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AgreementFileProperties) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AgreementFileProperties) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AgreementFileProperties) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AgreementFileProperties) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetIsDefault

`func (o *AgreementFileProperties) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AgreementFileProperties) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AgreementFileProperties) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AgreementFileProperties) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *AgreementFileProperties) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *AgreementFileProperties) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsMajorVersion

`func (o *AgreementFileProperties) GetIsMajorVersion() bool`

GetIsMajorVersion returns the IsMajorVersion field if non-nil, zero value otherwise.

### GetIsMajorVersionOk

`func (o *AgreementFileProperties) GetIsMajorVersionOk() (*bool, bool)`

GetIsMajorVersionOk returns a tuple with the IsMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMajorVersion

`func (o *AgreementFileProperties) SetIsMajorVersion(v bool)`

SetIsMajorVersion sets IsMajorVersion field to given value.

### HasIsMajorVersion

`func (o *AgreementFileProperties) HasIsMajorVersion() bool`

HasIsMajorVersion returns a boolean if a field has been set.

### SetIsMajorVersionNil

`func (o *AgreementFileProperties) SetIsMajorVersionNil(b bool)`

 SetIsMajorVersionNil sets the value for IsMajorVersion to be an explicit nil

### UnsetIsMajorVersion
`func (o *AgreementFileProperties) UnsetIsMajorVersion()`

UnsetIsMajorVersion ensures that no value is present for IsMajorVersion, not even an explicit nil
### GetLanguage

`func (o *AgreementFileProperties) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AgreementFileProperties) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AgreementFileProperties) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AgreementFileProperties) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AgreementFileProperties) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AgreementFileProperties) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


