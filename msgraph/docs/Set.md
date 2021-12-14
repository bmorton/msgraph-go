# Set

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of set creation. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description that gives details on the term usage. | [optional] 
**LocalizedNames** | Pointer to [**[]AnyOfmicrosoftGraphTermStoreLocalizedName**](AnyOfmicrosoftGraphTermStoreLocalizedName.md) | Name of the set for each languageTag. | [optional] 
**Properties** | Pointer to [**[]AnyOfmicrosoftGraphKeyValue**](AnyOfmicrosoftGraphKeyValue.md) | Custom properties for the set. | [optional] 
**Children** | Pointer to [**[]MicrosoftGraphTermStoreTerm**](MicrosoftGraphTermStoreTerm.md) | Children terms of set in term [store]. | [optional] 
**ParentGroup** | Pointer to [**MicrosoftGraphTermStoreGroup**](MicrosoftGraphTermStoreGroup.md) |  | [optional] 
**Relations** | Pointer to [**[]MicrosoftGraphTermStoreRelation**](MicrosoftGraphTermStoreRelation.md) | Indicates which terms have been pinned or reused directly under the set. | [optional] 
**Terms** | Pointer to [**[]MicrosoftGraphTermStoreTerm**](MicrosoftGraphTermStoreTerm.md) | All the terms under the set. | [optional] 

## Methods

### NewSet

`func NewSet() *Set`

NewSet instantiates a new Set object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWithDefaults

`func NewSetWithDefaults() *Set`

NewSetWithDefaults instantiates a new Set object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *Set) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Set) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Set) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Set) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Set) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Set) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *Set) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Set) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Set) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Set) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Set) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Set) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalizedNames

`func (o *Set) GetLocalizedNames() []*AnyOfmicrosoftGraphTermStoreLocalizedName`

GetLocalizedNames returns the LocalizedNames field if non-nil, zero value otherwise.

### GetLocalizedNamesOk

`func (o *Set) GetLocalizedNamesOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedName, bool)`

GetLocalizedNamesOk returns a tuple with the LocalizedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedNames

`func (o *Set) SetLocalizedNames(v []*AnyOfmicrosoftGraphTermStoreLocalizedName)`

SetLocalizedNames sets LocalizedNames field to given value.

### HasLocalizedNames

`func (o *Set) HasLocalizedNames() bool`

HasLocalizedNames returns a boolean if a field has been set.

### GetProperties

`func (o *Set) GetProperties() []*AnyOfmicrosoftGraphKeyValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Set) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Set) SetProperties(v []*AnyOfmicrosoftGraphKeyValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Set) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetChildren

`func (o *Set) GetChildren() []MicrosoftGraphTermStoreTerm`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Set) GetChildrenOk() (*[]MicrosoftGraphTermStoreTerm, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Set) SetChildren(v []MicrosoftGraphTermStoreTerm)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Set) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParentGroup

`func (o *Set) GetParentGroup() MicrosoftGraphTermStoreGroup`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *Set) GetParentGroupOk() (*MicrosoftGraphTermStoreGroup, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *Set) SetParentGroup(v MicrosoftGraphTermStoreGroup)`

SetParentGroup sets ParentGroup field to given value.

### HasParentGroup

`func (o *Set) HasParentGroup() bool`

HasParentGroup returns a boolean if a field has been set.

### GetRelations

`func (o *Set) GetRelations() []MicrosoftGraphTermStoreRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *Set) GetRelationsOk() (*[]MicrosoftGraphTermStoreRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *Set) SetRelations(v []MicrosoftGraphTermStoreRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *Set) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetTerms

`func (o *Set) GetTerms() []MicrosoftGraphTermStoreTerm`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Set) GetTermsOk() (*[]MicrosoftGraphTermStoreTerm, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Set) SetTerms(v []MicrosoftGraphTermStoreTerm)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Set) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


