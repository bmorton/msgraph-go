# MicrosoftGraphWorkbookSortField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ascending** | Pointer to **bool** | Represents whether the sorting is done in an ascending fashion. | [optional] 
**Color** | Pointer to **NullableString** | Represents the color that is the target of the condition if the sorting is on font or cell color. | [optional] 
**DataOption** | Pointer to **string** | Represents additional sorting options for this field. The possible values are: Normal, TextAsNumber. | [optional] 
**Icon** | Pointer to [**AnyOfmicrosoftGraphWorkbookIcon**](anyOf&lt;microsoft.graph.workbookIcon&gt;.md) | Represents the icon that is the target of the condition if the sorting is on the cell&#39;s icon. | [optional] 
**Key** | Pointer to **int32** | Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset from the first column (or row). | [optional] 
**SortOn** | Pointer to **string** | Represents the type of sorting of this condition. The possible values are: Value, CellColor, FontColor, Icon. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookSortField

`func NewMicrosoftGraphWorkbookSortField() *MicrosoftGraphWorkbookSortField`

NewMicrosoftGraphWorkbookSortField instantiates a new MicrosoftGraphWorkbookSortField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookSortFieldWithDefaults

`func NewMicrosoftGraphWorkbookSortFieldWithDefaults() *MicrosoftGraphWorkbookSortField`

NewMicrosoftGraphWorkbookSortFieldWithDefaults instantiates a new MicrosoftGraphWorkbookSortField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAscending

`func (o *MicrosoftGraphWorkbookSortField) GetAscending() bool`

GetAscending returns the Ascending field if non-nil, zero value otherwise.

### GetAscendingOk

`func (o *MicrosoftGraphWorkbookSortField) GetAscendingOk() (*bool, bool)`

GetAscendingOk returns a tuple with the Ascending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAscending

`func (o *MicrosoftGraphWorkbookSortField) SetAscending(v bool)`

SetAscending sets Ascending field to given value.

### HasAscending

`func (o *MicrosoftGraphWorkbookSortField) HasAscending() bool`

HasAscending returns a boolean if a field has been set.

### GetColor

`func (o *MicrosoftGraphWorkbookSortField) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphWorkbookSortField) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MicrosoftGraphWorkbookSortField) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *MicrosoftGraphWorkbookSortField) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MicrosoftGraphWorkbookSortField) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MicrosoftGraphWorkbookSortField) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDataOption

`func (o *MicrosoftGraphWorkbookSortField) GetDataOption() string`

GetDataOption returns the DataOption field if non-nil, zero value otherwise.

### GetDataOptionOk

`func (o *MicrosoftGraphWorkbookSortField) GetDataOptionOk() (*string, bool)`

GetDataOptionOk returns a tuple with the DataOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOption

`func (o *MicrosoftGraphWorkbookSortField) SetDataOption(v string)`

SetDataOption sets DataOption field to given value.

### HasDataOption

`func (o *MicrosoftGraphWorkbookSortField) HasDataOption() bool`

HasDataOption returns a boolean if a field has been set.

### GetIcon

`func (o *MicrosoftGraphWorkbookSortField) GetIcon() AnyOfmicrosoftGraphWorkbookIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *MicrosoftGraphWorkbookSortField) GetIconOk() (*AnyOfmicrosoftGraphWorkbookIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *MicrosoftGraphWorkbookSortField) SetIcon(v AnyOfmicrosoftGraphWorkbookIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *MicrosoftGraphWorkbookSortField) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *MicrosoftGraphWorkbookSortField) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *MicrosoftGraphWorkbookSortField) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetKey

`func (o *MicrosoftGraphWorkbookSortField) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MicrosoftGraphWorkbookSortField) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MicrosoftGraphWorkbookSortField) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *MicrosoftGraphWorkbookSortField) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSortOn

`func (o *MicrosoftGraphWorkbookSortField) GetSortOn() string`

GetSortOn returns the SortOn field if non-nil, zero value otherwise.

### GetSortOnOk

`func (o *MicrosoftGraphWorkbookSortField) GetSortOnOk() (*string, bool)`

GetSortOnOk returns a tuple with the SortOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOn

`func (o *MicrosoftGraphWorkbookSortField) SetSortOn(v string)`

SetSortOn sets SortOn field to given value.

### HasSortOn

`func (o *MicrosoftGraphWorkbookSortField) HasSortOn() bool`

HasSortOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


