# MicrosoftGraphListItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user which last modified the version. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Date and time the version was last modified. Read-only. | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) | Indicates the publication status of this particular version. Read-only. | [optional] 
**Fields** | Pointer to [**AnyOfmicrosoftGraphFieldValueSet**](anyOf&lt;microsoft.graph.fieldValueSet&gt;.md) | A collection of the fields and values for this version of the list item. | [optional] 

## Methods

### NewMicrosoftGraphListItemVersion

`func NewMicrosoftGraphListItemVersion() *MicrosoftGraphListItemVersion`

NewMicrosoftGraphListItemVersion instantiates a new MicrosoftGraphListItemVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphListItemVersionWithDefaults

`func NewMicrosoftGraphListItemVersionWithDefaults() *MicrosoftGraphListItemVersion`

NewMicrosoftGraphListItemVersionWithDefaults instantiates a new MicrosoftGraphListItemVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphListItemVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphListItemVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphListItemVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphListItemVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphListItemVersion) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphListItemVersion) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphListItemVersion) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphListItemVersion) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetPublication

`func (o *MicrosoftGraphListItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphListItemVersion) GetPublicationOk() (*AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublication

`func (o *MicrosoftGraphListItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication sets Publication field to given value.

### HasPublication

`func (o *MicrosoftGraphListItemVersion) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublicationNil

`func (o *MicrosoftGraphListItemVersion) SetPublicationNil(b bool)`

 SetPublicationNil sets the value for Publication to be an explicit nil

### UnsetPublication
`func (o *MicrosoftGraphListItemVersion) UnsetPublication()`

UnsetPublication ensures that no value is present for Publication, not even an explicit nil
### GetFields

`func (o *MicrosoftGraphListItemVersion) GetFields() AnyOfmicrosoftGraphFieldValueSet`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MicrosoftGraphListItemVersion) GetFieldsOk() (*AnyOfmicrosoftGraphFieldValueSet, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MicrosoftGraphListItemVersion) SetFields(v AnyOfmicrosoftGraphFieldValueSet)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MicrosoftGraphListItemVersion) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *MicrosoftGraphListItemVersion) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *MicrosoftGraphListItemVersion) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


