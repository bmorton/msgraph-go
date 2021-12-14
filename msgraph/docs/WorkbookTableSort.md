# WorkbookTableSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]AnyOfmicrosoftGraphWorkbookSortField**](AnyOfmicrosoftGraphWorkbookSortField.md) | Represents the current conditions used to last sort the table. Read-only. | [optional] 
**MatchCase** | Pointer to **bool** | Represents whether the casing impacted the last sort of the table. Read-only. | [optional] 
**Method** | Pointer to **string** | Represents Chinese character ordering method last used to sort the table. The possible values are: PinYin, StrokeCount. Read-only. | [optional] 

## Methods

### NewWorkbookTableSort

`func NewWorkbookTableSort() *WorkbookTableSort`

NewWorkbookTableSort instantiates a new WorkbookTableSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookTableSortWithDefaults

`func NewWorkbookTableSortWithDefaults() *WorkbookTableSort`

NewWorkbookTableSortWithDefaults instantiates a new WorkbookTableSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *WorkbookTableSort) GetFields() []*AnyOfmicrosoftGraphWorkbookSortField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkbookTableSort) GetFieldsOk() (*[]*AnyOfmicrosoftGraphWorkbookSortField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkbookTableSort) SetFields(v []*AnyOfmicrosoftGraphWorkbookSortField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkbookTableSort) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMatchCase

`func (o *WorkbookTableSort) GetMatchCase() bool`

GetMatchCase returns the MatchCase field if non-nil, zero value otherwise.

### GetMatchCaseOk

`func (o *WorkbookTableSort) GetMatchCaseOk() (*bool, bool)`

GetMatchCaseOk returns a tuple with the MatchCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCase

`func (o *WorkbookTableSort) SetMatchCase(v bool)`

SetMatchCase sets MatchCase field to given value.

### HasMatchCase

`func (o *WorkbookTableSort) HasMatchCase() bool`

HasMatchCase returns a boolean if a field has been set.

### GetMethod

`func (o *WorkbookTableSort) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WorkbookTableSort) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WorkbookTableSort) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WorkbookTableSort) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


