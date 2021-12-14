# MicrosoftGraphWorkbookFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Criteria** | Pointer to [**AnyOfmicrosoftGraphWorkbookFilterCriteria**](anyOf&lt;microsoft.graph.workbookFilterCriteria&gt;.md) | The currently applied filter on the given column. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookFilter

`func NewMicrosoftGraphWorkbookFilter() *MicrosoftGraphWorkbookFilter`

NewMicrosoftGraphWorkbookFilter instantiates a new MicrosoftGraphWorkbookFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookFilterWithDefaults

`func NewMicrosoftGraphWorkbookFilterWithDefaults() *MicrosoftGraphWorkbookFilter`

NewMicrosoftGraphWorkbookFilterWithDefaults instantiates a new MicrosoftGraphWorkbookFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCriteria

`func (o *MicrosoftGraphWorkbookFilter) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *MicrosoftGraphWorkbookFilter) GetCriteriaOk() (*AnyOfmicrosoftGraphWorkbookFilterCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *MicrosoftGraphWorkbookFilter) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *MicrosoftGraphWorkbookFilter) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteriaNil

`func (o *MicrosoftGraphWorkbookFilter) SetCriteriaNil(b bool)`

 SetCriteriaNil sets the value for Criteria to be an explicit nil

### UnsetCriteria
`func (o *MicrosoftGraphWorkbookFilter) UnsetCriteria()`

UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


