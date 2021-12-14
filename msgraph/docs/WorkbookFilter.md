# WorkbookFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**AnyOfmicrosoftGraphWorkbookFilterCriteria**](anyOf&lt;microsoft.graph.workbookFilterCriteria&gt;.md) | The currently applied filter on the given column. Read-only. | [optional] 

## Methods

### NewWorkbookFilter

`func NewWorkbookFilter() *WorkbookFilter`

NewWorkbookFilter instantiates a new WorkbookFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookFilterWithDefaults

`func NewWorkbookFilterWithDefaults() *WorkbookFilter`

NewWorkbookFilterWithDefaults instantiates a new WorkbookFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *WorkbookFilter) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *WorkbookFilter) GetCriteriaOk() (*AnyOfmicrosoftGraphWorkbookFilterCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *WorkbookFilter) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *WorkbookFilter) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteriaNil

`func (o *WorkbookFilter) SetCriteriaNil(b bool)`

 SetCriteriaNil sets the value for Criteria to be an explicit nil

### UnsetCriteria
`func (o *WorkbookFilter) UnsetCriteria()`

UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


