# MicrosoftGraphTermStoreSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of set creation. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description that gives details on the term usage. | [optional] 
**LocalizedNames** | Pointer to [**[]AnyOfmicrosoftGraphTermStoreLocalizedName**](AnyOfmicrosoftGraphTermStoreLocalizedName.md) | Name of the set for each languageTag. | [optional] 
**Properties** | Pointer to [**[]AnyOfmicrosoftGraphKeyValue**](AnyOfmicrosoftGraphKeyValue.md) | Custom properties for the set. | [optional] 
**Children** | Pointer to [**[]MicrosoftGraphTermStoreTerm**](MicrosoftGraphTermStoreTerm.md) | Children terms of set in term [store]. | [optional] 
**ParentGroup** | Pointer to [**MicrosoftGraphTermStoreGroup**](MicrosoftGraphTermStoreGroup.md) |  | [optional] 
**Relations** | Pointer to [**[]MicrosoftGraphTermStoreRelation**](MicrosoftGraphTermStoreRelation.md) | Indicates which terms have been pinned or reused directly under the set. | [optional] 
**Terms** | Pointer to [**[]MicrosoftGraphTermStoreTerm**](MicrosoftGraphTermStoreTerm.md) | All the terms under the set. | [optional] 

## Methods

### NewMicrosoftGraphTermStoreSet

`func NewMicrosoftGraphTermStoreSet() *MicrosoftGraphTermStoreSet`

NewMicrosoftGraphTermStoreSet instantiates a new MicrosoftGraphTermStoreSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermStoreSetWithDefaults

`func NewMicrosoftGraphTermStoreSetWithDefaults() *MicrosoftGraphTermStoreSet`

NewMicrosoftGraphTermStoreSetWithDefaults instantiates a new MicrosoftGraphTermStoreSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTermStoreSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermStoreSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTermStoreSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTermStoreSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphTermStoreSet) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTermStoreSet) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTermStoreSet) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTermStoreSet) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphTermStoreSet) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphTermStoreSet) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphTermStoreSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTermStoreSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTermStoreSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTermStoreSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTermStoreSet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTermStoreSet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalizedNames

`func (o *MicrosoftGraphTermStoreSet) GetLocalizedNames() []*AnyOfmicrosoftGraphTermStoreLocalizedName`

GetLocalizedNames returns the LocalizedNames field if non-nil, zero value otherwise.

### GetLocalizedNamesOk

`func (o *MicrosoftGraphTermStoreSet) GetLocalizedNamesOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedName, bool)`

GetLocalizedNamesOk returns a tuple with the LocalizedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedNames

`func (o *MicrosoftGraphTermStoreSet) SetLocalizedNames(v []*AnyOfmicrosoftGraphTermStoreLocalizedName)`

SetLocalizedNames sets LocalizedNames field to given value.

### HasLocalizedNames

`func (o *MicrosoftGraphTermStoreSet) HasLocalizedNames() bool`

HasLocalizedNames returns a boolean if a field has been set.

### GetProperties

`func (o *MicrosoftGraphTermStoreSet) GetProperties() []*AnyOfmicrosoftGraphKeyValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MicrosoftGraphTermStoreSet) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MicrosoftGraphTermStoreSet) SetProperties(v []*AnyOfmicrosoftGraphKeyValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MicrosoftGraphTermStoreSet) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetChildren

`func (o *MicrosoftGraphTermStoreSet) GetChildren() []MicrosoftGraphTermStoreTerm`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MicrosoftGraphTermStoreSet) GetChildrenOk() (*[]MicrosoftGraphTermStoreTerm, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MicrosoftGraphTermStoreSet) SetChildren(v []MicrosoftGraphTermStoreTerm)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MicrosoftGraphTermStoreSet) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParentGroup

`func (o *MicrosoftGraphTermStoreSet) GetParentGroup() MicrosoftGraphTermStoreGroup`

GetParentGroup returns the ParentGroup field if non-nil, zero value otherwise.

### GetParentGroupOk

`func (o *MicrosoftGraphTermStoreSet) GetParentGroupOk() (*MicrosoftGraphTermStoreGroup, bool)`

GetParentGroupOk returns a tuple with the ParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroup

`func (o *MicrosoftGraphTermStoreSet) SetParentGroup(v MicrosoftGraphTermStoreGroup)`

SetParentGroup sets ParentGroup field to given value.

### HasParentGroup

`func (o *MicrosoftGraphTermStoreSet) HasParentGroup() bool`

HasParentGroup returns a boolean if a field has been set.

### GetRelations

`func (o *MicrosoftGraphTermStoreSet) GetRelations() []MicrosoftGraphTermStoreRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *MicrosoftGraphTermStoreSet) GetRelationsOk() (*[]MicrosoftGraphTermStoreRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *MicrosoftGraphTermStoreSet) SetRelations(v []MicrosoftGraphTermStoreRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *MicrosoftGraphTermStoreSet) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetTerms

`func (o *MicrosoftGraphTermStoreSet) GetTerms() []MicrosoftGraphTermStoreTerm`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *MicrosoftGraphTermStoreSet) GetTermsOk() (*[]MicrosoftGraphTermStoreTerm, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *MicrosoftGraphTermStoreSet) SetTerms(v []MicrosoftGraphTermStoreTerm)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *MicrosoftGraphTermStoreSet) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


