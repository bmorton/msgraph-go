# MicrosoftGraphWorkforceIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the person who last modified the entity. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**ApiVersion** | Pointer to **NullableInt32** | API version for the call back URL. Start with 1. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the workforce integration. | [optional] 
**Encryption** | Pointer to [**AnyOfmicrosoftGraphWorkforceIntegrationEncryption**](anyOf&lt;microsoft.graph.workforceIntegrationEncryption&gt;.md) | The workforce integration encryption resource. | [optional] 
**IsActive** | Pointer to **NullableBool** | Indicates whether this workforce integration is currently active and available. | [optional] 
**SupportedEntities** | Pointer to [**AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities**](anyOf&lt;microsoft.graph.workforceIntegrationSupportedEntities&gt;.md) | The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue. | [optional] 
**Url** | Pointer to **NullableString** | Workforce Integration URL for callbacks from the Shifts service. | [optional] 

## Methods

### NewMicrosoftGraphWorkforceIntegration

`func NewMicrosoftGraphWorkforceIntegration() *MicrosoftGraphWorkforceIntegration`

NewMicrosoftGraphWorkforceIntegration instantiates a new MicrosoftGraphWorkforceIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkforceIntegrationWithDefaults

`func NewMicrosoftGraphWorkforceIntegrationWithDefaults() *MicrosoftGraphWorkforceIntegration`

NewMicrosoftGraphWorkforceIntegrationWithDefaults instantiates a new MicrosoftGraphWorkforceIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkforceIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkforceIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkforceIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkforceIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphWorkforceIntegration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWorkforceIntegration) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWorkforceIntegration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphWorkforceIntegration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphWorkforceIntegration) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphWorkforceIntegration) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphWorkforceIntegration) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphWorkforceIntegration) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWorkforceIntegration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphWorkforceIntegration) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetApiVersion

`func (o *MicrosoftGraphWorkforceIntegration) GetApiVersion() int32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MicrosoftGraphWorkforceIntegration) GetApiVersionOk() (*int32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MicrosoftGraphWorkforceIntegration) SetApiVersion(v int32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MicrosoftGraphWorkforceIntegration) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *MicrosoftGraphWorkforceIntegration) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *MicrosoftGraphWorkforceIntegration) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphWorkforceIntegration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWorkforceIntegration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphWorkforceIntegration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphWorkforceIntegration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphWorkforceIntegration) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphWorkforceIntegration) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEncryption

`func (o *MicrosoftGraphWorkforceIntegration) GetEncryption() AnyOfmicrosoftGraphWorkforceIntegrationEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *MicrosoftGraphWorkforceIntegration) GetEncryptionOk() (*AnyOfmicrosoftGraphWorkforceIntegrationEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *MicrosoftGraphWorkforceIntegration) SetEncryption(v AnyOfmicrosoftGraphWorkforceIntegrationEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *MicrosoftGraphWorkforceIntegration) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### SetEncryptionNil

`func (o *MicrosoftGraphWorkforceIntegration) SetEncryptionNil(b bool)`

 SetEncryptionNil sets the value for Encryption to be an explicit nil

### UnsetEncryption
`func (o *MicrosoftGraphWorkforceIntegration) UnsetEncryption()`

UnsetEncryption ensures that no value is present for Encryption, not even an explicit nil
### GetIsActive

`func (o *MicrosoftGraphWorkforceIntegration) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *MicrosoftGraphWorkforceIntegration) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *MicrosoftGraphWorkforceIntegration) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *MicrosoftGraphWorkforceIntegration) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *MicrosoftGraphWorkforceIntegration) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *MicrosoftGraphWorkforceIntegration) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetSupportedEntities

`func (o *MicrosoftGraphWorkforceIntegration) GetSupportedEntities() AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities`

GetSupportedEntities returns the SupportedEntities field if non-nil, zero value otherwise.

### GetSupportedEntitiesOk

`func (o *MicrosoftGraphWorkforceIntegration) GetSupportedEntitiesOk() (*AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities, bool)`

GetSupportedEntitiesOk returns a tuple with the SupportedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEntities

`func (o *MicrosoftGraphWorkforceIntegration) SetSupportedEntities(v AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities)`

SetSupportedEntities sets SupportedEntities field to given value.

### HasSupportedEntities

`func (o *MicrosoftGraphWorkforceIntegration) HasSupportedEntities() bool`

HasSupportedEntities returns a boolean if a field has been set.

### SetSupportedEntitiesNil

`func (o *MicrosoftGraphWorkforceIntegration) SetSupportedEntitiesNil(b bool)`

 SetSupportedEntitiesNil sets the value for SupportedEntities to be an explicit nil

### UnsetSupportedEntities
`func (o *MicrosoftGraphWorkforceIntegration) UnsetSupportedEntities()`

UnsetSupportedEntities ensures that no value is present for SupportedEntities, not even an explicit nil
### GetUrl

`func (o *MicrosoftGraphWorkforceIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MicrosoftGraphWorkforceIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MicrosoftGraphWorkforceIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MicrosoftGraphWorkforceIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MicrosoftGraphWorkforceIntegration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MicrosoftGraphWorkforceIntegration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


