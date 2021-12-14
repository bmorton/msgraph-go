# MicrosoftGraphTermStoreTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of term creation. Read-only. | [optional] 
**Descriptions** | Pointer to [**[]AnyOfmicrosoftGraphTermStoreLocalizedDescription**](AnyOfmicrosoftGraphTermStoreLocalizedDescription.md) | Description about term that is dependent on the languageTag. | [optional] 
**Labels** | Pointer to [**[]AnyOfmicrosoftGraphTermStoreLocalizedLabel**](AnyOfmicrosoftGraphTermStoreLocalizedLabel.md) | Label metadata for a term. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Last date and time of term modification. Read-only. | [optional] 
**Properties** | Pointer to [**[]AnyOfmicrosoftGraphKeyValue**](AnyOfmicrosoftGraphKeyValue.md) | Collection of properties on the term. | [optional] 
**Children** | Pointer to [**[]MicrosoftGraphTermStoreTerm**](MicrosoftGraphTermStoreTerm.md) | Children of current term. | [optional] 
**Relations** | Pointer to [**[]MicrosoftGraphTermStoreRelation**](MicrosoftGraphTermStoreRelation.md) | To indicate which terms are related to the current term as either pinned or reused. | [optional] 
**Set** | Pointer to [**AnyOfmicrosoftGraphTermStoreSet**](anyOf&lt;microsoft.graph.termStore.set&gt;.md) | The [set] in which the term is created. | [optional] 

## Methods

### NewMicrosoftGraphTermStoreTerm

`func NewMicrosoftGraphTermStoreTerm() *MicrosoftGraphTermStoreTerm`

NewMicrosoftGraphTermStoreTerm instantiates a new MicrosoftGraphTermStoreTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermStoreTermWithDefaults

`func NewMicrosoftGraphTermStoreTermWithDefaults() *MicrosoftGraphTermStoreTerm`

NewMicrosoftGraphTermStoreTermWithDefaults instantiates a new MicrosoftGraphTermStoreTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTermStoreTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermStoreTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTermStoreTerm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTermStoreTerm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphTermStoreTerm) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTermStoreTerm) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTermStoreTerm) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTermStoreTerm) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphTermStoreTerm) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphTermStoreTerm) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescriptions

`func (o *MicrosoftGraphTermStoreTerm) GetDescriptions() []*AnyOfmicrosoftGraphTermStoreLocalizedDescription`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *MicrosoftGraphTermStoreTerm) GetDescriptionsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedDescription, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *MicrosoftGraphTermStoreTerm) SetDescriptions(v []*AnyOfmicrosoftGraphTermStoreLocalizedDescription)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *MicrosoftGraphTermStoreTerm) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetLabels

`func (o *MicrosoftGraphTermStoreTerm) GetLabels() []*AnyOfmicrosoftGraphTermStoreLocalizedLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MicrosoftGraphTermStoreTerm) GetLabelsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MicrosoftGraphTermStoreTerm) SetLabels(v []*AnyOfmicrosoftGraphTermStoreLocalizedLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MicrosoftGraphTermStoreTerm) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTermStoreTerm) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTermStoreTerm) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTermStoreTerm) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTermStoreTerm) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphTermStoreTerm) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphTermStoreTerm) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetProperties

`func (o *MicrosoftGraphTermStoreTerm) GetProperties() []*AnyOfmicrosoftGraphKeyValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MicrosoftGraphTermStoreTerm) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MicrosoftGraphTermStoreTerm) SetProperties(v []*AnyOfmicrosoftGraphKeyValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MicrosoftGraphTermStoreTerm) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetChildren

`func (o *MicrosoftGraphTermStoreTerm) GetChildren() []MicrosoftGraphTermStoreTerm`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MicrosoftGraphTermStoreTerm) GetChildrenOk() (*[]MicrosoftGraphTermStoreTerm, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MicrosoftGraphTermStoreTerm) SetChildren(v []MicrosoftGraphTermStoreTerm)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MicrosoftGraphTermStoreTerm) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRelations

`func (o *MicrosoftGraphTermStoreTerm) GetRelations() []MicrosoftGraphTermStoreRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *MicrosoftGraphTermStoreTerm) GetRelationsOk() (*[]MicrosoftGraphTermStoreRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *MicrosoftGraphTermStoreTerm) SetRelations(v []MicrosoftGraphTermStoreRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *MicrosoftGraphTermStoreTerm) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetSet

`func (o *MicrosoftGraphTermStoreTerm) GetSet() AnyOfmicrosoftGraphTermStoreSet`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *MicrosoftGraphTermStoreTerm) GetSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *MicrosoftGraphTermStoreTerm) SetSet(v AnyOfmicrosoftGraphTermStoreSet)`

SetSet sets Set field to given value.

### HasSet

`func (o *MicrosoftGraphTermStoreTerm) HasSet() bool`

HasSet returns a boolean if a field has been set.

### SetSetNil

`func (o *MicrosoftGraphTermStoreTerm) SetSetNil(b bool)`

 SetSetNil sets the value for Set to be an explicit nil

### UnsetSet
`func (o *MicrosoftGraphTermStoreTerm) UnsetSet()`

UnsetSet ensures that no value is present for Set, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


