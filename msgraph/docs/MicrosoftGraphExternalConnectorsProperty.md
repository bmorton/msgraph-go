# MicrosoftGraphExternalConnectorsProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, &#x3D;, &amp;, ?, @, #, /, ~, &#39;, &#39;, &lt;, &gt;, &#x60;, ^. Optional. | [optional] 
**IsQueryable** | Pointer to **NullableBool** | Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional. | [optional] 
**IsRefinable** | Pointer to **NullableBool** | Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional. | [optional] 
**IsRetrievable** | Pointer to **NullableBool** | Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional. | [optional] 
**IsSearchable** | Pointer to **NullableBool** | Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional. | [optional] 
**Labels** | Pointer to [**[]AnyOfmicrosoftGraphExternalConnectorsLabel**](AnyOfmicrosoftGraphExternalConnectorsLabel.md) | Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional. | [optional] 
**Name** | Pointer to **string** | The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, &#x3D;, &amp;, ?, @, #, /, ~, &#39;, &#39;, &lt;, &gt;, &#x60;, ^.  Required. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsPropertyType**](anyOf&lt;microsoft.graph.externalConnectors.propertyType&gt;.md) | The data type of the property. Possible values are: string, int64, double, dateTime, boolean, stringCollection, int64Collection, doubleCollection, dateTimeCollection, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsProperty

`func NewMicrosoftGraphExternalConnectorsProperty() *MicrosoftGraphExternalConnectorsProperty`

NewMicrosoftGraphExternalConnectorsProperty instantiates a new MicrosoftGraphExternalConnectorsProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsPropertyWithDefaults

`func NewMicrosoftGraphExternalConnectorsPropertyWithDefaults() *MicrosoftGraphExternalConnectorsProperty`

NewMicrosoftGraphExternalConnectorsPropertyWithDefaults instantiates a new MicrosoftGraphExternalConnectorsProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *MicrosoftGraphExternalConnectorsProperty) GetAliases() []*string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetAliasesOk() (*[]*string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MicrosoftGraphExternalConnectorsProperty) SetAliases(v []*string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MicrosoftGraphExternalConnectorsProperty) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetIsQueryable

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsQueryable() bool`

GetIsQueryable returns the IsQueryable field if non-nil, zero value otherwise.

### GetIsQueryableOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsQueryableOk() (*bool, bool)`

GetIsQueryableOk returns a tuple with the IsQueryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueryable

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsQueryable(v bool)`

SetIsQueryable sets IsQueryable field to given value.

### HasIsQueryable

`func (o *MicrosoftGraphExternalConnectorsProperty) HasIsQueryable() bool`

HasIsQueryable returns a boolean if a field has been set.

### SetIsQueryableNil

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsQueryableNil(b bool)`

 SetIsQueryableNil sets the value for IsQueryable to be an explicit nil

### UnsetIsQueryable
`func (o *MicrosoftGraphExternalConnectorsProperty) UnsetIsQueryable()`

UnsetIsQueryable ensures that no value is present for IsQueryable, not even an explicit nil
### GetIsRefinable

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsRefinable() bool`

GetIsRefinable returns the IsRefinable field if non-nil, zero value otherwise.

### GetIsRefinableOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsRefinableOk() (*bool, bool)`

GetIsRefinableOk returns a tuple with the IsRefinable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefinable

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsRefinable(v bool)`

SetIsRefinable sets IsRefinable field to given value.

### HasIsRefinable

`func (o *MicrosoftGraphExternalConnectorsProperty) HasIsRefinable() bool`

HasIsRefinable returns a boolean if a field has been set.

### SetIsRefinableNil

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsRefinableNil(b bool)`

 SetIsRefinableNil sets the value for IsRefinable to be an explicit nil

### UnsetIsRefinable
`func (o *MicrosoftGraphExternalConnectorsProperty) UnsetIsRefinable()`

UnsetIsRefinable ensures that no value is present for IsRefinable, not even an explicit nil
### GetIsRetrievable

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsRetrievable() bool`

GetIsRetrievable returns the IsRetrievable field if non-nil, zero value otherwise.

### GetIsRetrievableOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsRetrievableOk() (*bool, bool)`

GetIsRetrievableOk returns a tuple with the IsRetrievable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetrievable

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsRetrievable(v bool)`

SetIsRetrievable sets IsRetrievable field to given value.

### HasIsRetrievable

`func (o *MicrosoftGraphExternalConnectorsProperty) HasIsRetrievable() bool`

HasIsRetrievable returns a boolean if a field has been set.

### SetIsRetrievableNil

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsRetrievableNil(b bool)`

 SetIsRetrievableNil sets the value for IsRetrievable to be an explicit nil

### UnsetIsRetrievable
`func (o *MicrosoftGraphExternalConnectorsProperty) UnsetIsRetrievable()`

UnsetIsRetrievable ensures that no value is present for IsRetrievable, not even an explicit nil
### GetIsSearchable

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsSearchable() bool`

GetIsSearchable returns the IsSearchable field if non-nil, zero value otherwise.

### GetIsSearchableOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetIsSearchableOk() (*bool, bool)`

GetIsSearchableOk returns a tuple with the IsSearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSearchable

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsSearchable(v bool)`

SetIsSearchable sets IsSearchable field to given value.

### HasIsSearchable

`func (o *MicrosoftGraphExternalConnectorsProperty) HasIsSearchable() bool`

HasIsSearchable returns a boolean if a field has been set.

### SetIsSearchableNil

`func (o *MicrosoftGraphExternalConnectorsProperty) SetIsSearchableNil(b bool)`

 SetIsSearchableNil sets the value for IsSearchable to be an explicit nil

### UnsetIsSearchable
`func (o *MicrosoftGraphExternalConnectorsProperty) UnsetIsSearchable()`

UnsetIsSearchable ensures that no value is present for IsSearchable, not even an explicit nil
### GetLabels

`func (o *MicrosoftGraphExternalConnectorsProperty) GetLabels() []*AnyOfmicrosoftGraphExternalConnectorsLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetLabelsOk() (*[]*AnyOfmicrosoftGraphExternalConnectorsLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MicrosoftGraphExternalConnectorsProperty) SetLabels(v []*AnyOfmicrosoftGraphExternalConnectorsLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MicrosoftGraphExternalConnectorsProperty) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphExternalConnectorsProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphExternalConnectorsProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphExternalConnectorsProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftGraphExternalConnectorsProperty) GetType() AnyOfmicrosoftGraphExternalConnectorsPropertyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphExternalConnectorsProperty) GetTypeOk() (*AnyOfmicrosoftGraphExternalConnectorsPropertyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphExternalConnectorsProperty) SetType(v AnyOfmicrosoftGraphExternalConnectorsPropertyType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphExternalConnectorsProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphExternalConnectorsProperty) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphExternalConnectorsProperty) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


