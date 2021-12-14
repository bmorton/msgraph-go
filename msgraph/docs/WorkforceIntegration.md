# WorkforceIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **NullableInt32** | API version for the call back URL. Start with 1. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the workforce integration. | [optional] 
**Encryption** | Pointer to [**AnyOfmicrosoftGraphWorkforceIntegrationEncryption**](anyOf&lt;microsoft.graph.workforceIntegrationEncryption&gt;.md) | The workforce integration encryption resource. | [optional] 
**IsActive** | Pointer to **NullableBool** | Indicates whether this workforce integration is currently active and available. | [optional] 
**SupportedEntities** | Pointer to [**AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities**](anyOf&lt;microsoft.graph.workforceIntegrationSupportedEntities&gt;.md) | The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue. | [optional] 
**Url** | Pointer to **NullableString** | Workforce Integration URL for callbacks from the Shifts service. | [optional] 

## Methods

### NewWorkforceIntegration

`func NewWorkforceIntegration() *WorkforceIntegration`

NewWorkforceIntegration instantiates a new WorkforceIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkforceIntegrationWithDefaults

`func NewWorkforceIntegrationWithDefaults() *WorkforceIntegration`

NewWorkforceIntegrationWithDefaults instantiates a new WorkforceIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkforceIntegration) GetApiVersion() int32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkforceIntegration) GetApiVersionOk() (*int32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkforceIntegration) SetApiVersion(v int32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkforceIntegration) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *WorkforceIntegration) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *WorkforceIntegration) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDisplayName

`func (o *WorkforceIntegration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkforceIntegration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkforceIntegration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkforceIntegration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkforceIntegration) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkforceIntegration) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEncryption

`func (o *WorkforceIntegration) GetEncryption() AnyOfmicrosoftGraphWorkforceIntegrationEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *WorkforceIntegration) GetEncryptionOk() (*AnyOfmicrosoftGraphWorkforceIntegrationEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *WorkforceIntegration) SetEncryption(v AnyOfmicrosoftGraphWorkforceIntegrationEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *WorkforceIntegration) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### SetEncryptionNil

`func (o *WorkforceIntegration) SetEncryptionNil(b bool)`

 SetEncryptionNil sets the value for Encryption to be an explicit nil

### UnsetEncryption
`func (o *WorkforceIntegration) UnsetEncryption()`

UnsetEncryption ensures that no value is present for Encryption, not even an explicit nil
### GetIsActive

`func (o *WorkforceIntegration) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WorkforceIntegration) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WorkforceIntegration) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WorkforceIntegration) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *WorkforceIntegration) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *WorkforceIntegration) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetSupportedEntities

`func (o *WorkforceIntegration) GetSupportedEntities() AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities`

GetSupportedEntities returns the SupportedEntities field if non-nil, zero value otherwise.

### GetSupportedEntitiesOk

`func (o *WorkforceIntegration) GetSupportedEntitiesOk() (*AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities, bool)`

GetSupportedEntitiesOk returns a tuple with the SupportedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEntities

`func (o *WorkforceIntegration) SetSupportedEntities(v AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities)`

SetSupportedEntities sets SupportedEntities field to given value.

### HasSupportedEntities

`func (o *WorkforceIntegration) HasSupportedEntities() bool`

HasSupportedEntities returns a boolean if a field has been set.

### SetSupportedEntitiesNil

`func (o *WorkforceIntegration) SetSupportedEntitiesNil(b bool)`

 SetSupportedEntitiesNil sets the value for SupportedEntities to be an explicit nil

### UnsetSupportedEntities
`func (o *WorkforceIntegration) UnsetSupportedEntities()`

UnsetSupportedEntities ensures that no value is present for SupportedEntities, not even an explicit nil
### GetUrl

`func (o *WorkforceIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkforceIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkforceIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WorkforceIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WorkforceIntegration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WorkforceIntegration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


