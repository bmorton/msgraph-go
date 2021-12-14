# MicrosoftGraphTermStoreStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DefaultLanguageTag** | Pointer to **string** | Default language of the term store. | [optional] 
**LanguageTags** | Pointer to **[]string** | List of languages for the term store. | [optional] 
**Groups** | Pointer to [**[]MicrosoftGraphTermStoreGroup**](MicrosoftGraphTermStoreGroup.md) | Collection of all groups available in the term store. | [optional] 
**Sets** | Pointer to [**[]MicrosoftGraphTermStoreSet**](MicrosoftGraphTermStoreSet.md) | Collection of all sets available in the term store. | [optional] 

## Methods

### NewMicrosoftGraphTermStoreStore

`func NewMicrosoftGraphTermStoreStore() *MicrosoftGraphTermStoreStore`

NewMicrosoftGraphTermStoreStore instantiates a new MicrosoftGraphTermStoreStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermStoreStoreWithDefaults

`func NewMicrosoftGraphTermStoreStoreWithDefaults() *MicrosoftGraphTermStoreStore`

NewMicrosoftGraphTermStoreStoreWithDefaults instantiates a new MicrosoftGraphTermStoreStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTermStoreStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermStoreStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTermStoreStore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTermStoreStore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultLanguageTag

`func (o *MicrosoftGraphTermStoreStore) GetDefaultLanguageTag() string`

GetDefaultLanguageTag returns the DefaultLanguageTag field if non-nil, zero value otherwise.

### GetDefaultLanguageTagOk

`func (o *MicrosoftGraphTermStoreStore) GetDefaultLanguageTagOk() (*string, bool)`

GetDefaultLanguageTagOk returns a tuple with the DefaultLanguageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageTag

`func (o *MicrosoftGraphTermStoreStore) SetDefaultLanguageTag(v string)`

SetDefaultLanguageTag sets DefaultLanguageTag field to given value.

### HasDefaultLanguageTag

`func (o *MicrosoftGraphTermStoreStore) HasDefaultLanguageTag() bool`

HasDefaultLanguageTag returns a boolean if a field has been set.

### GetLanguageTags

`func (o *MicrosoftGraphTermStoreStore) GetLanguageTags() []string`

GetLanguageTags returns the LanguageTags field if non-nil, zero value otherwise.

### GetLanguageTagsOk

`func (o *MicrosoftGraphTermStoreStore) GetLanguageTagsOk() (*[]string, bool)`

GetLanguageTagsOk returns a tuple with the LanguageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageTags

`func (o *MicrosoftGraphTermStoreStore) SetLanguageTags(v []string)`

SetLanguageTags sets LanguageTags field to given value.

### HasLanguageTags

`func (o *MicrosoftGraphTermStoreStore) HasLanguageTags() bool`

HasLanguageTags returns a boolean if a field has been set.

### GetGroups

`func (o *MicrosoftGraphTermStoreStore) GetGroups() []MicrosoftGraphTermStoreGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MicrosoftGraphTermStoreStore) GetGroupsOk() (*[]MicrosoftGraphTermStoreGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MicrosoftGraphTermStoreStore) SetGroups(v []MicrosoftGraphTermStoreGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MicrosoftGraphTermStoreStore) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSets

`func (o *MicrosoftGraphTermStoreStore) GetSets() []MicrosoftGraphTermStoreSet`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *MicrosoftGraphTermStoreStore) GetSetsOk() (*[]MicrosoftGraphTermStoreSet, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *MicrosoftGraphTermStoreStore) SetSets(v []MicrosoftGraphTermStoreSet)`

SetSets sets Sets field to given value.

### HasSets

`func (o *MicrosoftGraphTermStoreStore) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


