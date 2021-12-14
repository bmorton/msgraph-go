# MicrosoftGraphWorkbookTableColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Index** | Pointer to **int32** | Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | Returns the name of the table column. | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string. | [optional] 
**Filter** | Pointer to [**AnyOfmicrosoftGraphWorkbookFilter**](anyOf&lt;microsoft.graph.workbookFilter&gt;.md) | Retrieve the filter applied to the column. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookTableColumn

`func NewMicrosoftGraphWorkbookTableColumn() *MicrosoftGraphWorkbookTableColumn`

NewMicrosoftGraphWorkbookTableColumn instantiates a new MicrosoftGraphWorkbookTableColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookTableColumnWithDefaults

`func NewMicrosoftGraphWorkbookTableColumnWithDefaults() *MicrosoftGraphWorkbookTableColumn`

NewMicrosoftGraphWorkbookTableColumnWithDefaults instantiates a new MicrosoftGraphWorkbookTableColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookTableColumn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookTableColumn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookTableColumn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *MicrosoftGraphWorkbookTableColumn) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MicrosoftGraphWorkbookTableColumn) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MicrosoftGraphWorkbookTableColumn) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphWorkbookTableColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookTableColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookTableColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookTableColumn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookTableColumn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValues

`func (o *MicrosoftGraphWorkbookTableColumn) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetValuesOk() (*AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookTableColumn) SetValues(v AnyOfobject)`

SetValues sets Values field to given value.

### HasValues

`func (o *MicrosoftGraphWorkbookTableColumn) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *MicrosoftGraphWorkbookTableColumn) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *MicrosoftGraphWorkbookTableColumn) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetFilter

`func (o *MicrosoftGraphWorkbookTableColumn) GetFilter() AnyOfmicrosoftGraphWorkbookFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetFilterOk() (*AnyOfmicrosoftGraphWorkbookFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MicrosoftGraphWorkbookTableColumn) SetFilter(v AnyOfmicrosoftGraphWorkbookFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MicrosoftGraphWorkbookTableColumn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *MicrosoftGraphWorkbookTableColumn) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *MicrosoftGraphWorkbookTableColumn) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


