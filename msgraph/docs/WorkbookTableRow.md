# WorkbookTableRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Returns the index number of the row within the rows collection of the table. Zero-indexed. Read-only. | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string. | [optional] 

## Methods

### NewWorkbookTableRow

`func NewWorkbookTableRow() *WorkbookTableRow`

NewWorkbookTableRow instantiates a new WorkbookTableRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookTableRowWithDefaults

`func NewWorkbookTableRowWithDefaults() *WorkbookTableRow`

NewWorkbookTableRowWithDefaults instantiates a new WorkbookTableRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *WorkbookTableRow) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WorkbookTableRow) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WorkbookTableRow) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *WorkbookTableRow) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetValues

`func (o *WorkbookTableRow) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkbookTableRow) GetValuesOk() (*AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WorkbookTableRow) SetValues(v AnyOfobject)`

SetValues sets Values field to given value.

### HasValues

`func (o *WorkbookTableRow) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *WorkbookTableRow) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *WorkbookTableRow) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


