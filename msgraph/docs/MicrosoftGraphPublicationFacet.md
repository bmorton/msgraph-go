# MicrosoftGraphPublicationFacet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **NullableString** | The state of publication for this document. Either published or checkout. Read-only. | [optional] 
**VersionId** | Pointer to **NullableString** | The unique identifier for the version that is visible to the current caller. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPublicationFacet

`func NewMicrosoftGraphPublicationFacet() *MicrosoftGraphPublicationFacet`

NewMicrosoftGraphPublicationFacet instantiates a new MicrosoftGraphPublicationFacet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPublicationFacetWithDefaults

`func NewMicrosoftGraphPublicationFacetWithDefaults() *MicrosoftGraphPublicationFacet`

NewMicrosoftGraphPublicationFacetWithDefaults instantiates a new MicrosoftGraphPublicationFacet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *MicrosoftGraphPublicationFacet) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MicrosoftGraphPublicationFacet) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MicrosoftGraphPublicationFacet) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MicrosoftGraphPublicationFacet) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *MicrosoftGraphPublicationFacet) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *MicrosoftGraphPublicationFacet) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetVersionId

`func (o *MicrosoftGraphPublicationFacet) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *MicrosoftGraphPublicationFacet) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *MicrosoftGraphPublicationFacet) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *MicrosoftGraphPublicationFacet) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### SetVersionIdNil

`func (o *MicrosoftGraphPublicationFacet) SetVersionIdNil(b bool)`

 SetVersionIdNil sets the value for VersionId to be an explicit nil

### UnsetVersionId
`func (o *MicrosoftGraphPublicationFacet) UnsetVersionId()`

UnsetVersionId ensures that no value is present for VersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


