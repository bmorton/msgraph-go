# MicrosoftGraphTermStoreGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of group creation. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description giving details on the term usage. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of group. | [optional] 
**ParentSiteId** | Pointer to **NullableString** | Id of the parent site of this group. | [optional] 
**Scope** | Pointer to [**AnyOfmicrosoftGraphTermStoreTermGroupScope**](anyOf&lt;microsoft.graph.termStore.termGroupScope&gt;.md) | Returns type of group. Possible values are &#39;global&#39;, &#39;system&#39; and &#39;siteCollection&#39;. | [optional] 
**Sets** | Pointer to [**[]MicrosoftGraphTermStoreSet**](MicrosoftGraphTermStoreSet.md) | All sets under the group in a term [store]. | [optional] 

## Methods

### NewMicrosoftGraphTermStoreGroup

`func NewMicrosoftGraphTermStoreGroup() *MicrosoftGraphTermStoreGroup`

NewMicrosoftGraphTermStoreGroup instantiates a new MicrosoftGraphTermStoreGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermStoreGroupWithDefaults

`func NewMicrosoftGraphTermStoreGroupWithDefaults() *MicrosoftGraphTermStoreGroup`

NewMicrosoftGraphTermStoreGroupWithDefaults instantiates a new MicrosoftGraphTermStoreGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTermStoreGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermStoreGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTermStoreGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTermStoreGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphTermStoreGroup) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTermStoreGroup) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTermStoreGroup) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTermStoreGroup) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphTermStoreGroup) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphTermStoreGroup) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphTermStoreGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTermStoreGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTermStoreGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTermStoreGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTermStoreGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTermStoreGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTermStoreGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTermStoreGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTermStoreGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTermStoreGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphTermStoreGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphTermStoreGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetParentSiteId

`func (o *MicrosoftGraphTermStoreGroup) GetParentSiteId() string`

GetParentSiteId returns the ParentSiteId field if non-nil, zero value otherwise.

### GetParentSiteIdOk

`func (o *MicrosoftGraphTermStoreGroup) GetParentSiteIdOk() (*string, bool)`

GetParentSiteIdOk returns a tuple with the ParentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSiteId

`func (o *MicrosoftGraphTermStoreGroup) SetParentSiteId(v string)`

SetParentSiteId sets ParentSiteId field to given value.

### HasParentSiteId

`func (o *MicrosoftGraphTermStoreGroup) HasParentSiteId() bool`

HasParentSiteId returns a boolean if a field has been set.

### SetParentSiteIdNil

`func (o *MicrosoftGraphTermStoreGroup) SetParentSiteIdNil(b bool)`

 SetParentSiteIdNil sets the value for ParentSiteId to be an explicit nil

### UnsetParentSiteId
`func (o *MicrosoftGraphTermStoreGroup) UnsetParentSiteId()`

UnsetParentSiteId ensures that no value is present for ParentSiteId, not even an explicit nil
### GetScope

`func (o *MicrosoftGraphTermStoreGroup) GetScope() AnyOfmicrosoftGraphTermStoreTermGroupScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphTermStoreGroup) GetScopeOk() (*AnyOfmicrosoftGraphTermStoreTermGroupScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphTermStoreGroup) SetScope(v AnyOfmicrosoftGraphTermStoreTermGroupScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphTermStoreGroup) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphTermStoreGroup) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphTermStoreGroup) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSets

`func (o *MicrosoftGraphTermStoreGroup) GetSets() []MicrosoftGraphTermStoreSet`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *MicrosoftGraphTermStoreGroup) GetSetsOk() (*[]MicrosoftGraphTermStoreSet, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *MicrosoftGraphTermStoreGroup) SetSets(v []MicrosoftGraphTermStoreSet)`

SetSets sets Sets field to given value.

### HasSets

`func (o *MicrosoftGraphTermStoreGroup) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


