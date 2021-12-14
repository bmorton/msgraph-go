# Term

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of term creation. Read-only. | [optional] 
**Descriptions** | Pointer to [**[]AnyOfmicrosoftGraphTermStoreLocalizedDescription**](AnyOfmicrosoftGraphTermStoreLocalizedDescription.md) | Description about term that is dependent on the languageTag. | [optional] 
**Labels** | Pointer to [**[]AnyOfmicrosoftGraphTermStoreLocalizedLabel**](AnyOfmicrosoftGraphTermStoreLocalizedLabel.md) | Label metadata for a term. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Last date and time of term modification. Read-only. | [optional] 
**Properties** | Pointer to [**[]AnyOfmicrosoftGraphKeyValue**](AnyOfmicrosoftGraphKeyValue.md) | Collection of properties on the term. | [optional] 
**Children** | Pointer to [**[]MicrosoftGraphTermStoreTerm**](MicrosoftGraphTermStoreTerm.md) | Children of current term. | [optional] 
**Relations** | Pointer to [**[]MicrosoftGraphTermStoreRelation**](MicrosoftGraphTermStoreRelation.md) | To indicate which terms are related to the current term as either pinned or reused. | [optional] 
**Set** | Pointer to [**AnyOfmicrosoftGraphTermStoreSet**](anyOf&lt;microsoft.graph.termStore.set&gt;.md) | The [set] in which the term is created. | [optional] 

## Methods

### NewTerm

`func NewTerm() *Term`

NewTerm instantiates a new Term object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermWithDefaults

`func NewTermWithDefaults() *Term`

NewTermWithDefaults instantiates a new Term object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *Term) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Term) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Term) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Term) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Term) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Term) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescriptions

`func (o *Term) GetDescriptions() []*AnyOfmicrosoftGraphTermStoreLocalizedDescription`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Term) GetDescriptionsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedDescription, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Term) SetDescriptions(v []*AnyOfmicrosoftGraphTermStoreLocalizedDescription)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *Term) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetLabels

`func (o *Term) GetLabels() []*AnyOfmicrosoftGraphTermStoreLocalizedLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Term) GetLabelsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Term) SetLabels(v []*AnyOfmicrosoftGraphTermStoreLocalizedLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Term) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *Term) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *Term) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *Term) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *Term) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *Term) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *Term) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetProperties

`func (o *Term) GetProperties() []*AnyOfmicrosoftGraphKeyValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Term) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Term) SetProperties(v []*AnyOfmicrosoftGraphKeyValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Term) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetChildren

`func (o *Term) GetChildren() []MicrosoftGraphTermStoreTerm`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Term) GetChildrenOk() (*[]MicrosoftGraphTermStoreTerm, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Term) SetChildren(v []MicrosoftGraphTermStoreTerm)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Term) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRelations

`func (o *Term) GetRelations() []MicrosoftGraphTermStoreRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *Term) GetRelationsOk() (*[]MicrosoftGraphTermStoreRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *Term) SetRelations(v []MicrosoftGraphTermStoreRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *Term) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetSet

`func (o *Term) GetSet() AnyOfmicrosoftGraphTermStoreSet`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *Term) GetSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *Term) SetSet(v AnyOfmicrosoftGraphTermStoreSet)`

SetSet sets Set field to given value.

### HasSet

`func (o *Term) HasSet() bool`

HasSet returns a boolean if a field has been set.

### SetSetNil

`func (o *Term) SetSetNil(b bool)`

 SetSetNil sets the value for Set to be an explicit nil

### UnsetSet
`func (o *Term) UnsetSet()`

UnsetSet ensures that no value is present for Set, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


