# MicrosoftGraphDriveItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user which last modified the version. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Date and time the version was last modified. Read-only. | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) | Indicates the publication status of this particular version. Read-only. | [optional] 
**Content** | Pointer to **NullableString** | The content stream for this version of the item. | [optional] 
**Size** | Pointer to **NullableInt64** | Indicates the size of the content stream for this version of the item. | [optional] 

## Methods

### NewMicrosoftGraphDriveItemVersion

`func NewMicrosoftGraphDriveItemVersion() *MicrosoftGraphDriveItemVersion`

NewMicrosoftGraphDriveItemVersion instantiates a new MicrosoftGraphDriveItemVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDriveItemVersionWithDefaults

`func NewMicrosoftGraphDriveItemVersionWithDefaults() *MicrosoftGraphDriveItemVersion`

NewMicrosoftGraphDriveItemVersionWithDefaults instantiates a new MicrosoftGraphDriveItemVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDriveItemVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDriveItemVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDriveItemVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDriveItemVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphDriveItemVersion) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphDriveItemVersion) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDriveItemVersion) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphDriveItemVersion) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPublication

`func (o *MicrosoftGraphDriveItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphDriveItemVersion) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *MicrosoftGraphDriveItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *MicrosoftGraphDriveItemVersion) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *MicrosoftGraphDriveItemVersion) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *MicrosoftGraphDriveItemVersion) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetContent

`func (o *MicrosoftGraphDriveItemVersion) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphDriveItemVersion) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphDriveItemVersion) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphDriveItemVersion) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphDriveItemVersion) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphDriveItemVersion) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphDriveItemVersion) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphDriveItemVersion) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphDriveItemVersion) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphDriveItemVersion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphDriveItemVersion) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphDriveItemVersion) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


