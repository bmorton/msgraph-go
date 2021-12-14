# MicrosoftGraphTermStoreRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Relationship** | Pointer to [**AnyOfmicrosoftGraphTermStoreRelationType**](anyOf&lt;microsoft.graph.termStore.relationType&gt;.md) | The type of relation. Possible values are: pin, reuse. | [optional] 
**FromTerm** | Pointer to [**AnyOfmicrosoftGraphTermStoreTerm**](anyOf&lt;microsoft.graph.termStore.term&gt;.md) | The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set]. | [optional] 
**Set** | Pointer to [**AnyOfmicrosoftGraphTermStoreSet**](anyOf&lt;microsoft.graph.termStore.set&gt;.md) | The [set] in which the relation is relevant. | [optional] 
**ToTerm** | Pointer to [**AnyOfmicrosoftGraphTermStoreTerm**](anyOf&lt;microsoft.graph.termStore.term&gt;.md) | The to [term] of the relation. The term to which the relationship is defined. | [optional] 

## Methods

### NewMicrosoftGraphTermStoreRelation

`func NewMicrosoftGraphTermStoreRelation() *MicrosoftGraphTermStoreRelation`

NewMicrosoftGraphTermStoreRelation instantiates a new MicrosoftGraphTermStoreRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermStoreRelationWithDefaults

`func NewMicrosoftGraphTermStoreRelationWithDefaults() *MicrosoftGraphTermStoreRelation`

NewMicrosoftGraphTermStoreRelationWithDefaults instantiates a new MicrosoftGraphTermStoreRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTermStoreRelation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermStoreRelation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTermStoreRelation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTermStoreRelation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationship

`func (o *MicrosoftGraphTermStoreRelation) GetRelationship() AnyOfmicrosoftGraphTermStoreRelationType`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *MicrosoftGraphTermStoreRelation) GetRelationshipOk() (*AnyOfmicrosoftGraphTermStoreRelationType, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *MicrosoftGraphTermStoreRelation) SetRelationship(v AnyOfmicrosoftGraphTermStoreRelationType)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *MicrosoftGraphTermStoreRelation) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### SetRelationshipNil

`func (o *MicrosoftGraphTermStoreRelation) SetRelationshipNil(b bool)`

 SetRelationshipNil sets the value for Relationship to be an explicit nil

### UnsetRelationship
`func (o *MicrosoftGraphTermStoreRelation) UnsetRelationship()`

UnsetRelationship ensures that no value is present for Relationship, not even an explicit nil
### GetFromTerm

`func (o *MicrosoftGraphTermStoreRelation) GetFromTerm() AnyOfmicrosoftGraphTermStoreTerm`

GetFromTerm returns the FromTerm field if non-nil, zero value otherwise.

### GetFromTermOk

`func (o *MicrosoftGraphTermStoreRelation) GetFromTermOk() (*AnyOfmicrosoftGraphTermStoreTerm, bool)`

GetFromTermOk returns a tuple with the FromTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTerm

`func (o *MicrosoftGraphTermStoreRelation) SetFromTerm(v AnyOfmicrosoftGraphTermStoreTerm)`

SetFromTerm sets FromTerm field to given value.

### HasFromTerm

`func (o *MicrosoftGraphTermStoreRelation) HasFromTerm() bool`

HasFromTerm returns a boolean if a field has been set.

### SetFromTermNil

`func (o *MicrosoftGraphTermStoreRelation) SetFromTermNil(b bool)`

 SetFromTermNil sets the value for FromTerm to be an explicit nil

### UnsetFromTerm
`func (o *MicrosoftGraphTermStoreRelation) UnsetFromTerm()`

UnsetFromTerm ensures that no value is present for FromTerm, not even an explicit nil
### GetSet

`func (o *MicrosoftGraphTermStoreRelation) GetSet() AnyOfmicrosoftGraphTermStoreSet`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *MicrosoftGraphTermStoreRelation) GetSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *MicrosoftGraphTermStoreRelation) SetSet(v AnyOfmicrosoftGraphTermStoreSet)`

SetSet sets Set field to given value.

### HasSet

`func (o *MicrosoftGraphTermStoreRelation) HasSet() bool`

HasSet returns a boolean if a field has been set.

### SetSetNil

`func (o *MicrosoftGraphTermStoreRelation) SetSetNil(b bool)`

 SetSetNil sets the value for Set to be an explicit nil

### UnsetSet
`func (o *MicrosoftGraphTermStoreRelation) UnsetSet()`

UnsetSet ensures that no value is present for Set, not even an explicit nil
### GetToTerm

`func (o *MicrosoftGraphTermStoreRelation) GetToTerm() AnyOfmicrosoftGraphTermStoreTerm`

GetToTerm returns the ToTerm field if non-nil, zero value otherwise.

### GetToTermOk

`func (o *MicrosoftGraphTermStoreRelation) GetToTermOk() (*AnyOfmicrosoftGraphTermStoreTerm, bool)`

GetToTermOk returns a tuple with the ToTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTerm

`func (o *MicrosoftGraphTermStoreRelation) SetToTerm(v AnyOfmicrosoftGraphTermStoreTerm)`

SetToTerm sets ToTerm field to given value.

### HasToTerm

`func (o *MicrosoftGraphTermStoreRelation) HasToTerm() bool`

HasToTerm returns a boolean if a field has been set.

### SetToTermNil

`func (o *MicrosoftGraphTermStoreRelation) SetToTermNil(b bool)`

 SetToTermNil sets the value for ToTerm to be an explicit nil

### UnsetToTerm
`func (o *MicrosoftGraphTermStoreRelation) UnsetToTerm()`

UnsetToTerm ensures that no value is present for ToTerm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


