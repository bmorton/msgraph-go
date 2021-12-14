# MicrosoftGraphCalculatedColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** | For dateTime output types, the format of the value. Must be one of dateOnly or dateTime. | [optional] 
**Formula** | Pointer to **NullableString** | The formula used to compute the value for this column. | [optional] 
**OutputType** | Pointer to **NullableString** | The output type used to format values in this column. Must be one of boolean, currency, dateTime, number, or text. | [optional] 

## Methods

### NewMicrosoftGraphCalculatedColumn

`func NewMicrosoftGraphCalculatedColumn() *MicrosoftGraphCalculatedColumn`

NewMicrosoftGraphCalculatedColumn instantiates a new MicrosoftGraphCalculatedColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCalculatedColumnWithDefaults

`func NewMicrosoftGraphCalculatedColumnWithDefaults() *MicrosoftGraphCalculatedColumn`

NewMicrosoftGraphCalculatedColumnWithDefaults instantiates a new MicrosoftGraphCalculatedColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *MicrosoftGraphCalculatedColumn) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphCalculatedColumn) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphCalculatedColumn) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphCalculatedColumn) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphCalculatedColumn) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphCalculatedColumn) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetFormula

`func (o *MicrosoftGraphCalculatedColumn) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *MicrosoftGraphCalculatedColumn) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *MicrosoftGraphCalculatedColumn) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *MicrosoftGraphCalculatedColumn) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### SetFormulaNil

`func (o *MicrosoftGraphCalculatedColumn) SetFormulaNil(b bool)`

 SetFormulaNil sets the value for Formula to be an explicit nil

### UnsetFormula
`func (o *MicrosoftGraphCalculatedColumn) UnsetFormula()`

UnsetFormula ensures that no value is present for Formula, not even an explicit nil
### GetOutputType

`func (o *MicrosoftGraphCalculatedColumn) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *MicrosoftGraphCalculatedColumn) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *MicrosoftGraphCalculatedColumn) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *MicrosoftGraphCalculatedColumn) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### SetOutputTypeNil

`func (o *MicrosoftGraphCalculatedColumn) SetOutputTypeNil(b bool)`

 SetOutputTypeNil sets the value for OutputType to be an explicit nil

### UnsetOutputType
`func (o *MicrosoftGraphCalculatedColumn) UnsetOutputType()`

UnsetOutputType ensures that no value is present for OutputType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


