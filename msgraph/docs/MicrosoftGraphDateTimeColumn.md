# MicrosoftGraphDateTimeColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayAs** | Pointer to **NullableString** | How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default. | [optional] 
**Format** | Pointer to **NullableString** | Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime | [optional] 

## Methods

### NewMicrosoftGraphDateTimeColumn

`func NewMicrosoftGraphDateTimeColumn() *MicrosoftGraphDateTimeColumn`

NewMicrosoftGraphDateTimeColumn instantiates a new MicrosoftGraphDateTimeColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDateTimeColumnWithDefaults

`func NewMicrosoftGraphDateTimeColumnWithDefaults() *MicrosoftGraphDateTimeColumn`

NewMicrosoftGraphDateTimeColumnWithDefaults instantiates a new MicrosoftGraphDateTimeColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayAs

`func (o *MicrosoftGraphDateTimeColumn) GetDisplayAs() string`

GetDisplayAs returns the DisplayAs field if non-nil, zero value otherwise.

### GetDisplayAsOk

`func (o *MicrosoftGraphDateTimeColumn) GetDisplayAsOk() (*string, bool)`

GetDisplayAsOk returns a tuple with the DisplayAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAs

`func (o *MicrosoftGraphDateTimeColumn) SetDisplayAs(v string)`

SetDisplayAs sets DisplayAs field to given value.

### HasDisplayAs

`func (o *MicrosoftGraphDateTimeColumn) HasDisplayAs() bool`

HasDisplayAs returns a boolean if a field has been set.

### SetDisplayAsNil

`func (o *MicrosoftGraphDateTimeColumn) SetDisplayAsNil(b bool)`

 SetDisplayAsNil sets the value for DisplayAs to be an explicit nil

### UnsetDisplayAs
`func (o *MicrosoftGraphDateTimeColumn) UnsetDisplayAs()`

UnsetDisplayAs ensures that no value is present for DisplayAs, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphDateTimeColumn) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphDateTimeColumn) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphDateTimeColumn) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphDateTimeColumn) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphDateTimeColumn) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphDateTimeColumn) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


