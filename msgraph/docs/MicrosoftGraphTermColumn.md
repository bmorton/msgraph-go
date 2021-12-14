# MicrosoftGraphTermColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleValues** | Pointer to **NullableBool** | Specifies whether the column will allow more than one value. | [optional] 
**ShowFullyQualifiedName** | Pointer to **NullableBool** | Specifies whether to display the entire term path or only the term label. | [optional] 
**ParentTerm** | Pointer to [**AnyOfmicrosoftGraphTermStoreTerm**](anyOf&lt;microsoft.graph.termStore.term&gt;.md) |  | [optional] 
**TermSet** | Pointer to [**AnyOfmicrosoftGraphTermStoreSet**](anyOf&lt;microsoft.graph.termStore.set&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphTermColumn

`func NewMicrosoftGraphTermColumn() *MicrosoftGraphTermColumn`

NewMicrosoftGraphTermColumn instantiates a new MicrosoftGraphTermColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermColumnWithDefaults

`func NewMicrosoftGraphTermColumnWithDefaults() *MicrosoftGraphTermColumn`

NewMicrosoftGraphTermColumnWithDefaults instantiates a new MicrosoftGraphTermColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleValues

`func (o *MicrosoftGraphTermColumn) GetAllowMultipleValues() bool`

GetAllowMultipleValues returns the AllowMultipleValues field if non-nil, zero value otherwise.

### GetAllowMultipleValuesOk

`func (o *MicrosoftGraphTermColumn) GetAllowMultipleValuesOk() (*bool, bool)`

GetAllowMultipleValuesOk returns a tuple with the AllowMultipleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleValues

`func (o *MicrosoftGraphTermColumn) SetAllowMultipleValues(v bool)`

SetAllowMultipleValues sets AllowMultipleValues field to given value.

### HasAllowMultipleValues

`func (o *MicrosoftGraphTermColumn) HasAllowMultipleValues() bool`

HasAllowMultipleValues returns a boolean if a field has been set.

### SetAllowMultipleValuesNil

`func (o *MicrosoftGraphTermColumn) SetAllowMultipleValuesNil(b bool)`

 SetAllowMultipleValuesNil sets the value for AllowMultipleValues to be an explicit nil

### UnsetAllowMultipleValues
`func (o *MicrosoftGraphTermColumn) UnsetAllowMultipleValues()`

UnsetAllowMultipleValues ensures that no value is present for AllowMultipleValues, not even an explicit nil
### GetShowFullyQualifiedName

`func (o *MicrosoftGraphTermColumn) GetShowFullyQualifiedName() bool`

GetShowFullyQualifiedName returns the ShowFullyQualifiedName field if non-nil, zero value otherwise.

### GetShowFullyQualifiedNameOk

`func (o *MicrosoftGraphTermColumn) GetShowFullyQualifiedNameOk() (*bool, bool)`

GetShowFullyQualifiedNameOk returns a tuple with the ShowFullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFullyQualifiedName

`func (o *MicrosoftGraphTermColumn) SetShowFullyQualifiedName(v bool)`

SetShowFullyQualifiedName sets ShowFullyQualifiedName field to given value.

### HasShowFullyQualifiedName

`func (o *MicrosoftGraphTermColumn) HasShowFullyQualifiedName() bool`

HasShowFullyQualifiedName returns a boolean if a field has been set.

### SetShowFullyQualifiedNameNil

`func (o *MicrosoftGraphTermColumn) SetShowFullyQualifiedNameNil(b bool)`

 SetShowFullyQualifiedNameNil sets the value for ShowFullyQualifiedName to be an explicit nil

### UnsetShowFullyQualifiedName
`func (o *MicrosoftGraphTermColumn) UnsetShowFullyQualifiedName()`

UnsetShowFullyQualifiedName ensures that no value is present for ShowFullyQualifiedName, not even an explicit nil
### GetParentTerm

`func (o *MicrosoftGraphTermColumn) GetParentTerm() AnyOfmicrosoftGraphTermStoreTerm`

GetParentTerm returns the ParentTerm field if non-nil, zero value otherwise.

### GetParentTermOk

`func (o *MicrosoftGraphTermColumn) GetParentTermOk() (*AnyOfmicrosoftGraphTermStoreTerm, bool)`

GetParentTermOk returns a tuple with the ParentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTerm

`func (o *MicrosoftGraphTermColumn) SetParentTerm(v AnyOfmicrosoftGraphTermStoreTerm)`

SetParentTerm sets ParentTerm field to given value.

### HasParentTerm

`func (o *MicrosoftGraphTermColumn) HasParentTerm() bool`

HasParentTerm returns a boolean if a field has been set.

### SetParentTermNil

`func (o *MicrosoftGraphTermColumn) SetParentTermNil(b bool)`

 SetParentTermNil sets the value for ParentTerm to be an explicit nil

### UnsetParentTerm
`func (o *MicrosoftGraphTermColumn) UnsetParentTerm()`

UnsetParentTerm ensures that no value is present for ParentTerm, not even an explicit nil
### GetTermSet

`func (o *MicrosoftGraphTermColumn) GetTermSet() AnyOfmicrosoftGraphTermStoreSet`

GetTermSet returns the TermSet field if non-nil, zero value otherwise.

### GetTermSetOk

`func (o *MicrosoftGraphTermColumn) GetTermSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool)`

GetTermSetOk returns a tuple with the TermSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSet

`func (o *MicrosoftGraphTermColumn) SetTermSet(v AnyOfmicrosoftGraphTermStoreSet)`

SetTermSet sets TermSet field to given value.

### HasTermSet

`func (o *MicrosoftGraphTermColumn) HasTermSet() bool`

HasTermSet returns a boolean if a field has been set.

### SetTermSetNil

`func (o *MicrosoftGraphTermColumn) SetTermSetNil(b bool)`

 SetTermSetNil sets the value for TermSet to be an explicit nil

### UnsetTermSet
`func (o *MicrosoftGraphTermColumn) UnsetTermSet()`

UnsetTermSet ensures that no value is present for TermSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


