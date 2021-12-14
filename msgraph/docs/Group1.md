# Group1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | Date and time of group creation. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Description giving details on the term usage. | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of group. | [optional] 
**ParentSiteId** | Pointer to **NullableString** | Id of the parent site of this group. | [optional] 
**Scope** | Pointer to [**AnyOfmicrosoftGraphTermStoreTermGroupScope**](anyOf&lt;microsoft.graph.termStore.termGroupScope&gt;.md) | Returns type of group. Possible values are &#39;global&#39;, &#39;system&#39; and &#39;siteCollection&#39;. | [optional] 
**Sets** | Pointer to [**[]MicrosoftGraphTermStoreSet**](MicrosoftGraphTermStoreSet.md) | All sets under the group in a term [store]. | [optional] 

## Methods

### NewGroup1

`func NewGroup1() *Group1`

NewGroup1 instantiates a new Group1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroup1WithDefaults

`func NewGroup1WithDefaults() *Group1`

NewGroup1WithDefaults instantiates a new Group1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *Group1) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Group1) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Group1) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Group1) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *Group1) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *Group1) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *Group1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Group1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Group1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Group1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Group1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Group1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Group1) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Group1) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetParentSiteId

`func (o *Group1) GetParentSiteId() string`

GetParentSiteId returns the ParentSiteId field if non-nil, zero value otherwise.

### GetParentSiteIdOk

`func (o *Group1) GetParentSiteIdOk() (*string, bool)`

GetParentSiteIdOk returns a tuple with the ParentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSiteId

`func (o *Group1) SetParentSiteId(v string)`

SetParentSiteId sets ParentSiteId field to given value.

### HasParentSiteId

`func (o *Group1) HasParentSiteId() bool`

HasParentSiteId returns a boolean if a field has been set.

### SetParentSiteIdNil

`func (o *Group1) SetParentSiteIdNil(b bool)`

 SetParentSiteIdNil sets the value for ParentSiteId to be an explicit nil

### UnsetParentSiteId
`func (o *Group1) UnsetParentSiteId()`

UnsetParentSiteId ensures that no value is present for ParentSiteId, not even an explicit nil
### GetScope

`func (o *Group1) GetScope() AnyOfmicrosoftGraphTermStoreTermGroupScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Group1) GetScopeOk() (*AnyOfmicrosoftGraphTermStoreTermGroupScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Group1) SetScope(v AnyOfmicrosoftGraphTermStoreTermGroupScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Group1) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *Group1) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *Group1) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSets

`func (o *Group1) GetSets() []MicrosoftGraphTermStoreSet`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *Group1) GetSetsOk() (*[]MicrosoftGraphTermStoreSet, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *Group1) SetSets(v []MicrosoftGraphTermStoreSet)`

SetSets sets Sets field to given value.

### HasSets

`func (o *Group1) HasSets() bool`

HasSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


