# MicrosoftGraphTextColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleLines** | Pointer to **NullableBool** | Whether to allow multiple lines of text. | [optional] 
**AppendChangesToExistingText** | Pointer to **NullableBool** | Whether updates to this column should replace existing text, or append to it. | [optional] 
**LinesForEditing** | Pointer to **NullableInt32** | The size of the text box. | [optional] 
**MaxLength** | Pointer to **NullableInt32** | The maximum number of characters for the value. | [optional] 
**TextType** | Pointer to **NullableString** | The type of text being stored. Must be one of plain or richText | [optional] 

## Methods

### NewMicrosoftGraphTextColumn

`func NewMicrosoftGraphTextColumn() *MicrosoftGraphTextColumn`

NewMicrosoftGraphTextColumn instantiates a new MicrosoftGraphTextColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTextColumnWithDefaults

`func NewMicrosoftGraphTextColumnWithDefaults() *MicrosoftGraphTextColumn`

NewMicrosoftGraphTextColumnWithDefaults instantiates a new MicrosoftGraphTextColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleLines

`func (o *MicrosoftGraphTextColumn) GetAllowMultipleLines() bool`

GetAllowMultipleLines returns the AllowMultipleLines field if non-nil, zero value otherwise.

### GetAllowMultipleLinesOk

`func (o *MicrosoftGraphTextColumn) GetAllowMultipleLinesOk() (*bool, bool)`

GetAllowMultipleLinesOk returns a tuple with the AllowMultipleLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleLines

`func (o *MicrosoftGraphTextColumn) SetAllowMultipleLines(v bool)`

SetAllowMultipleLines sets AllowMultipleLines field to given value.

### HasAllowMultipleLines

`func (o *MicrosoftGraphTextColumn) HasAllowMultipleLines() bool`

HasAllowMultipleLines returns a boolean if a field has been set.

### SetAllowMultipleLinesNil

`func (o *MicrosoftGraphTextColumn) SetAllowMultipleLinesNil(b bool)`

 SetAllowMultipleLinesNil sets the value for AllowMultipleLines to be an explicit nil

### UnsetAllowMultipleLines
`func (o *MicrosoftGraphTextColumn) UnsetAllowMultipleLines()`

UnsetAllowMultipleLines ensures that no value is present for AllowMultipleLines, not even an explicit nil
### GetAppendChangesToExistingText

`func (o *MicrosoftGraphTextColumn) GetAppendChangesToExistingText() bool`

GetAppendChangesToExistingText returns the AppendChangesToExistingText field if non-nil, zero value otherwise.

### GetAppendChangesToExistingTextOk

`func (o *MicrosoftGraphTextColumn) GetAppendChangesToExistingTextOk() (*bool, bool)`

GetAppendChangesToExistingTextOk returns a tuple with the AppendChangesToExistingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendChangesToExistingText

`func (o *MicrosoftGraphTextColumn) SetAppendChangesToExistingText(v bool)`

SetAppendChangesToExistingText sets AppendChangesToExistingText field to given value.

### HasAppendChangesToExistingText

`func (o *MicrosoftGraphTextColumn) HasAppendChangesToExistingText() bool`

HasAppendChangesToExistingText returns a boolean if a field has been set.

### SetAppendChangesToExistingTextNil

`func (o *MicrosoftGraphTextColumn) SetAppendChangesToExistingTextNil(b bool)`

 SetAppendChangesToExistingTextNil sets the value for AppendChangesToExistingText to be an explicit nil

### UnsetAppendChangesToExistingText
`func (o *MicrosoftGraphTextColumn) UnsetAppendChangesToExistingText()`

UnsetAppendChangesToExistingText ensures that no value is present for AppendChangesToExistingText, not even an explicit nil
### GetLinesForEditing

`func (o *MicrosoftGraphTextColumn) GetLinesForEditing() int32`

GetLinesForEditing returns the LinesForEditing field if non-nil, zero value otherwise.

### GetLinesForEditingOk

`func (o *MicrosoftGraphTextColumn) GetLinesForEditingOk() (*int32, bool)`

GetLinesForEditingOk returns a tuple with the LinesForEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesForEditing

`func (o *MicrosoftGraphTextColumn) SetLinesForEditing(v int32)`

SetLinesForEditing sets LinesForEditing field to given value.

### HasLinesForEditing

`func (o *MicrosoftGraphTextColumn) HasLinesForEditing() bool`

HasLinesForEditing returns a boolean if a field has been set.

### SetLinesForEditingNil

`func (o *MicrosoftGraphTextColumn) SetLinesForEditingNil(b bool)`

 SetLinesForEditingNil sets the value for LinesForEditing to be an explicit nil

### UnsetLinesForEditing
`func (o *MicrosoftGraphTextColumn) UnsetLinesForEditing()`

UnsetLinesForEditing ensures that no value is present for LinesForEditing, not even an explicit nil
### GetMaxLength

`func (o *MicrosoftGraphTextColumn) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *MicrosoftGraphTextColumn) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *MicrosoftGraphTextColumn) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *MicrosoftGraphTextColumn) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *MicrosoftGraphTextColumn) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *MicrosoftGraphTextColumn) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetTextType

`func (o *MicrosoftGraphTextColumn) GetTextType() string`

GetTextType returns the TextType field if non-nil, zero value otherwise.

### GetTextTypeOk

`func (o *MicrosoftGraphTextColumn) GetTextTypeOk() (*string, bool)`

GetTextTypeOk returns a tuple with the TextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextType

`func (o *MicrosoftGraphTextColumn) SetTextType(v string)`

SetTextType sets TextType field to given value.

### HasTextType

`func (o *MicrosoftGraphTextColumn) HasTextType() bool`

HasTextType returns a boolean if a field has been set.

### SetTextTypeNil

`func (o *MicrosoftGraphTextColumn) SetTextTypeNil(b bool)`

 SetTextTypeNil sets the value for TextType to be an explicit nil

### UnsetTextType
`func (o *MicrosoftGraphTextColumn) UnsetTextType()`

UnsetTextType ensures that no value is present for TextType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


