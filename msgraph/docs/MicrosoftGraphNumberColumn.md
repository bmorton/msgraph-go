# MicrosoftGraphNumberColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecimalPlaces** | Pointer to **NullableString** | How many decimal places to display. See below for information about the possible values. | [optional] 
**DisplayAs** | Pointer to **NullableString** | How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number. | [optional] 
**Maximum** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The maximum permitted value. | [optional] 
**Minimum** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The minimum permitted value. | [optional] 

## Methods

### NewMicrosoftGraphNumberColumn

`func NewMicrosoftGraphNumberColumn() *MicrosoftGraphNumberColumn`

NewMicrosoftGraphNumberColumn instantiates a new MicrosoftGraphNumberColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphNumberColumnWithDefaults

`func NewMicrosoftGraphNumberColumnWithDefaults() *MicrosoftGraphNumberColumn`

NewMicrosoftGraphNumberColumnWithDefaults instantiates a new MicrosoftGraphNumberColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimalPlaces

`func (o *MicrosoftGraphNumberColumn) GetDecimalPlaces() string`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *MicrosoftGraphNumberColumn) GetDecimalPlacesOk() (*string, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *MicrosoftGraphNumberColumn) SetDecimalPlaces(v string)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *MicrosoftGraphNumberColumn) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *MicrosoftGraphNumberColumn) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *MicrosoftGraphNumberColumn) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetDisplayAs

`func (o *MicrosoftGraphNumberColumn) GetDisplayAs() string`

GetDisplayAs returns the DisplayAs field if non-nil, zero value otherwise.

### GetDisplayAsOk

`func (o *MicrosoftGraphNumberColumn) GetDisplayAsOk() (*string, bool)`

GetDisplayAsOk returns a tuple with the DisplayAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAs

`func (o *MicrosoftGraphNumberColumn) SetDisplayAs(v string)`

SetDisplayAs sets DisplayAs field to given value.

### HasDisplayAs

`func (o *MicrosoftGraphNumberColumn) HasDisplayAs() bool`

HasDisplayAs returns a boolean if a field has been set.

### SetDisplayAsNil

`func (o *MicrosoftGraphNumberColumn) SetDisplayAsNil(b bool)`

 SetDisplayAsNil sets the value for DisplayAs to be an explicit nil

### UnsetDisplayAs
`func (o *MicrosoftGraphNumberColumn) UnsetDisplayAs()`

UnsetDisplayAs ensures that no value is present for DisplayAs, not even an explicit nil
### GetMaximum

`func (o *MicrosoftGraphNumberColumn) GetMaximum() AnyOfnumberstringstring`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *MicrosoftGraphNumberColumn) GetMaximumOk() (*AnyOfnumberstringstring, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *MicrosoftGraphNumberColumn) SetMaximum(v AnyOfnumberstringstring)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *MicrosoftGraphNumberColumn) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *MicrosoftGraphNumberColumn) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *MicrosoftGraphNumberColumn) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetMinimum

`func (o *MicrosoftGraphNumberColumn) GetMinimum() AnyOfnumberstringstring`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *MicrosoftGraphNumberColumn) GetMinimumOk() (*AnyOfnumberstringstring, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *MicrosoftGraphNumberColumn) SetMinimum(v AnyOfnumberstringstring)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *MicrosoftGraphNumberColumn) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *MicrosoftGraphNumberColumn) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *MicrosoftGraphNumberColumn) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


