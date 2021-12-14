# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLanguageTag** | Pointer to **string** | Default language of the term store. | [optional] 
**LanguageTags** | Pointer to **[]string** | List of languages for the term store. | [optional] 
**Groups** | Pointer to [**[]MicrosoftGraphTermStoreGroup**](MicrosoftGraphTermStoreGroup.md) | Collection of all groups available in the term store. | [optional] 
**Sets** | Pointer to [**[]MicrosoftGraphTermStoreSet**](MicrosoftGraphTermStoreSet.md) | Collection of all sets available in the term store. | [optional] 

## Methods

### NewStore

`func NewStore() *Store`

NewStore instantiates a new Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWithDefaults

`func NewStoreWithDefaults() *Store`

NewStoreWithDefaults instantiates a new Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLanguageTag

`func (o *Store) GetDefaultLanguageTag() string`

GetDefaultLanguageTag returns the DefaultLanguageTag field if non-nil, zero value otherwise.

### GetDefaultLanguageTagOk

`func (o *Store) GetDefaultLanguageTagOk() (*string, bool)`

GetDefaultLanguageTagOk returns a tuple with the DefaultLanguageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageTag

`func (o *Store) SetDefaultLanguageTag(v string)`

SetDefaultLanguageTag sets DefaultLanguageTag field to given value.

### HasDefaultLanguageTag

`func (o *Store) HasDefaultLanguageTag() bool`

HasDefaultLanguageTag returns a boolean if a field has been set.

### GetLanguageTags

`func (o *Store) GetLanguageTags() []string`

GetLanguageTags returns the LanguageTags field if non-nil, zero value otherwise.

### GetLanguageTagsOk

`func (o *Store) GetLanguageTagsOk() (*[]string, bool)`

GetLanguageTagsOk returns a tuple with the LanguageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageTags

`func (o *Store) SetLanguageTags(v []string)`

SetLanguageTags sets LanguageTags field to given value.

### HasLanguageTags

`func (o *Store) HasLanguageTags() bool`

HasLanguageTags returns a boolean if a field has been set.

### GetGroups

`func (o *Store) GetGroups() []MicrosoftGraphTermStoreGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Store) GetGroupsOk() (*[]MicrosoftGraphTermStoreGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Store) SetGroups(v []MicrosoftGraphTermStoreGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Store) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSets

`func (o *Store) GetSets() []MicrosoftGraphTermStoreSet`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *Store) GetSetsOk() (*[]MicrosoftGraphTermStoreSet, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *Store) SetSets(v []MicrosoftGraphTermStoreSet)`

SetSets sets Sets field to given value.

### HasSets

`func (o *Store) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


