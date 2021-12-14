# MicrosoftGraphWorkbookChartAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**MajorUnit** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number. | [optional] 
**Maximum** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number. | [optional] 
**Minimum** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number. | [optional] 
**MinorUnit** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the interval between two minor tick marks. &#39;Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisFormat**](anyOf&lt;microsoft.graph.workbookChartAxisFormat&gt;.md) | Represents the formatting of a chart object, which includes line and font formatting. Read-only. | [optional] 
**MajorGridlines** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlines**](anyOf&lt;microsoft.graph.workbookChartGridlines&gt;.md) | Returns a gridlines object that represents the major gridlines for the specified axis. Read-only. | [optional] 
**MinorGridlines** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlines**](anyOf&lt;microsoft.graph.workbookChartGridlines&gt;.md) | Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only. | [optional] 
**Title** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisTitle**](anyOf&lt;microsoft.graph.workbookChartAxisTitle&gt;.md) | Represents the axis title. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartAxis

`func NewMicrosoftGraphWorkbookChartAxis() *MicrosoftGraphWorkbookChartAxis`

NewMicrosoftGraphWorkbookChartAxis instantiates a new MicrosoftGraphWorkbookChartAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartAxisWithDefaults

`func NewMicrosoftGraphWorkbookChartAxisWithDefaults() *MicrosoftGraphWorkbookChartAxis`

NewMicrosoftGraphWorkbookChartAxisWithDefaults instantiates a new MicrosoftGraphWorkbookChartAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartAxis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartAxis) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartAxis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMajorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorUnit() AnyOfobject`

GetMajorUnit returns the MajorUnit field if non-nil, zero value otherwise.

### GetMajorUnitOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorUnitOk() (*AnyOfobject, bool)`

GetMajorUnitOk returns a tuple with the MajorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorUnit(v AnyOfobject)`

SetMajorUnit sets MajorUnit field to given value.

### HasMajorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) HasMajorUnit() bool`

HasMajorUnit returns a boolean if a field has been set.

### SetMajorUnitNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorUnitNil(b bool)`

 SetMajorUnitNil sets the value for MajorUnit to be an explicit nil

### UnsetMajorUnit
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetMajorUnit()`

UnsetMajorUnit ensures that no value is present for MajorUnit, not even an explicit nil
### GetMaximum

`func (o *MicrosoftGraphWorkbookChartAxis) GetMaximum() AnyOfobject`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMaximumOk() (*AnyOfobject, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *MicrosoftGraphWorkbookChartAxis) SetMaximum(v AnyOfobject)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *MicrosoftGraphWorkbookChartAxis) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetMinimum

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinimum() AnyOfobject`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinimumOk() (*AnyOfobject, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinimum(v AnyOfobject)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *MicrosoftGraphWorkbookChartAxis) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMinorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorUnit() AnyOfobject`

GetMinorUnit returns the MinorUnit field if non-nil, zero value otherwise.

### GetMinorUnitOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorUnitOk() (*AnyOfobject, bool)`

GetMinorUnitOk returns a tuple with the MinorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorUnit(v AnyOfobject)`

SetMinorUnit sets MinorUnit field to given value.

### HasMinorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) HasMinorUnit() bool`

HasMinorUnit returns a boolean if a field has been set.

### SetMinorUnitNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorUnitNil(b bool)`

 SetMinorUnitNil sets the value for MinorUnit to be an explicit nil

### UnsetMinorUnit
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetMinorUnit()`

UnsetMinorUnit ensures that no value is present for MinorUnit, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphWorkbookChartAxis) GetFormat() AnyOfmicrosoftGraphWorkbookChartAxisFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAxisFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartAxis) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAxisFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartAxis) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetMajorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines`

GetMajorGridlines returns the MajorGridlines field if non-nil, zero value otherwise.

### GetMajorGridlinesOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorGridlinesOk() (*AnyOfmicrosoftGraphWorkbookChartGridlines, bool)`

GetMajorGridlinesOk returns a tuple with the MajorGridlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines)`

SetMajorGridlines sets MajorGridlines field to given value.

### HasMajorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) HasMajorGridlines() bool`

HasMajorGridlines returns a boolean if a field has been set.

### SetMajorGridlinesNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorGridlinesNil(b bool)`

 SetMajorGridlinesNil sets the value for MajorGridlines to be an explicit nil

### UnsetMajorGridlines
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetMajorGridlines()`

UnsetMajorGridlines ensures that no value is present for MajorGridlines, not even an explicit nil
### GetMinorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines`

GetMinorGridlines returns the MinorGridlines field if non-nil, zero value otherwise.

### GetMinorGridlinesOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorGridlinesOk() (*AnyOfmicrosoftGraphWorkbookChartGridlines, bool)`

GetMinorGridlinesOk returns a tuple with the MinorGridlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines)`

SetMinorGridlines sets MinorGridlines field to given value.

### HasMinorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) HasMinorGridlines() bool`

HasMinorGridlines returns a boolean if a field has been set.

### SetMinorGridlinesNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorGridlinesNil(b bool)`

 SetMinorGridlinesNil sets the value for MinorGridlines to be an explicit nil

### UnsetMinorGridlines
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetMinorGridlines()`

UnsetMinorGridlines ensures that no value is present for MinorGridlines, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphWorkbookChartAxis) GetTitle() AnyOfmicrosoftGraphWorkbookChartAxisTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetTitleOk() (*AnyOfmicrosoftGraphWorkbookChartAxisTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphWorkbookChartAxis) SetTitle(v AnyOfmicrosoftGraphWorkbookChartAxisTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphWorkbookChartAxis) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphWorkbookChartAxis) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphWorkbookChartAxis) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


