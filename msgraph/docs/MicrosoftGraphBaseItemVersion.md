# MicrosoftGraphBaseItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user which last modified the version. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Date and time the version was last modified. Read-only. | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) | Indicates the publication status of this particular version. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphBaseItemVersion

`func NewMicrosoftGraphBaseItemVersion() *MicrosoftGraphBaseItemVersion`

NewMicrosoftGraphBaseItemVersion instantiates a new MicrosoftGraphBaseItemVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphBaseItemVersionWithDefaults

`func NewMicrosoftGraphBaseItemVersionWithDefaults() *MicrosoftGraphBaseItemVersion`

NewMicrosoftGraphBaseItemVersionWithDefaults instantiates a new MicrosoftGraphBaseItemVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphBaseItemVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphBaseItemVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphBaseItemVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphBaseItemVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphBaseItemVersion) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphBaseItemVersion) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphBaseItemVersion) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphBaseItemVersion) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPublication

`func (o *MicrosoftGraphBaseItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphBaseItemVersion) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *MicrosoftGraphBaseItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *MicrosoftGraphBaseItemVersion) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *MicrosoftGraphBaseItemVersion) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *MicrosoftGraphBaseItemVersion) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


