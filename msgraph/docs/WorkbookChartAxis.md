# WorkbookChartAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MajorUnit** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the interval between two major tick marks. Can be set to a numeric value or an empty string.  The returned value is always a number. | [optional] 
**Maximum** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the maximum value on the value axis.  Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number. | [optional] 
**Minimum** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the minimum value on the value axis. Can be set to a numeric value or an empty string (for automatic axis values).  The returned value is always a number. | [optional] 
**MinorUnit** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the interval between two minor tick marks. &#39;Can be set to a numeric value or an empty string (for automatic axis values). The returned value is always a number. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisFormat**](anyOf&lt;microsoft.graph.workbookChartAxisFormat&gt;.md) | Represents the formatting of a chart object, which includes line and font formatting. Read-only. | [optional] 
**MajorGridlines** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlines**](anyOf&lt;microsoft.graph.workbookChartGridlines&gt;.md) | Returns a gridlines object that represents the major gridlines for the specified axis. Read-only. | [optional] 
**MinorGridlines** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlines**](anyOf&lt;microsoft.graph.workbookChartGridlines&gt;.md) | Returns a Gridlines object that represents the minor gridlines for the specified axis. Read-only. | [optional] 
**Title** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisTitle**](anyOf&lt;microsoft.graph.workbookChartAxisTitle&gt;.md) | Represents the axis title. Read-only. | [optional] 

## Methods

### NewWorkbookChartAxis

`func NewWorkbookChartAxis() *WorkbookChartAxis`

NewWorkbookChartAxis instantiates a new WorkbookChartAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookChartAxisWithDefaults

`func NewWorkbookChartAxisWithDefaults() *WorkbookChartAxis`

NewWorkbookChartAxisWithDefaults instantiates a new WorkbookChartAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajorUnit

`func (o *WorkbookChartAxis) GetMajorUnit() AnyOfobject`

GetMajorUnit returns the MajorUnit field if non-nil, zero value otherwise.

### GetMajorUnitOk

`func (o *WorkbookChartAxis) GetMajorUnitOk() (*AnyOfobject, bool)`

GetMajorUnitOk returns a tuple with the MajorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorUnit

`func (o *WorkbookChartAxis) SetMajorUnit(v AnyOfobject)`

SetMajorUnit sets MajorUnit field to given value.

### HasMajorUnit

`func (o *WorkbookChartAxis) HasMajorUnit() bool`

HasMajorUnit returns a boolean if a field has been set.

### SetMajorUnitNil

`func (o *WorkbookChartAxis) SetMajorUnitNil(b bool)`

 SetMajorUnitNil sets the value for MajorUnit to be an explicit nil

### UnsetMajorUnit
`func (o *WorkbookChartAxis) UnsetMajorUnit()`

UnsetMajorUnit ensures that no value is present for MajorUnit, not even an explicit nil
### GetMaximum

`func (o *WorkbookChartAxis) GetMaximum() AnyOfobject`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *WorkbookChartAxis) GetMaximumOk() (*AnyOfobject, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *WorkbookChartAxis) SetMaximum(v AnyOfobject)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *WorkbookChartAxis) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *WorkbookChartAxis) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *WorkbookChartAxis) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetMinimum

`func (o *WorkbookChartAxis) GetMinimum() AnyOfobject`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *WorkbookChartAxis) GetMinimumOk() (*AnyOfobject, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *WorkbookChartAxis) SetMinimum(v AnyOfobject)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *WorkbookChartAxis) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *WorkbookChartAxis) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *WorkbookChartAxis) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMinorUnit

`func (o *WorkbookChartAxis) GetMinorUnit() AnyOfobject`

GetMinorUnit returns the MinorUnit field if non-nil, zero value otherwise.

### GetMinorUnitOk

`func (o *WorkbookChartAxis) GetMinorUnitOk() (*AnyOfobject, bool)`

GetMinorUnitOk returns a tuple with the MinorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorUnit

`func (o *WorkbookChartAxis) SetMinorUnit(v AnyOfobject)`

SetMinorUnit sets MinorUnit field to given value.

### HasMinorUnit

`func (o *WorkbookChartAxis) HasMinorUnit() bool`

HasMinorUnit returns a boolean if a field has been set.

### SetMinorUnitNil

`func (o *WorkbookChartAxis) SetMinorUnitNil(b bool)`

 SetMinorUnitNil sets the value for MinorUnit to be an explicit nil

### UnsetMinorUnit
`func (o *WorkbookChartAxis) UnsetMinorUnit()`

UnsetMinorUnit ensures that no value is present for MinorUnit, not even an explicit nil
### GetFormat

`func (o *WorkbookChartAxis) GetFormat() AnyOfmicrosoftGraphWorkbookChartAxisFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChartAxis) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartAxisFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkbookChartAxis) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAxisFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WorkbookChartAxis) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *WorkbookChartAxis) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *WorkbookChartAxis) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetMajorGridlines

`func (o *WorkbookChartAxis) GetMajorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines`

GetMajorGridlines returns the MajorGridlines field if non-nil, zero value otherwise.

### GetMajorGridlinesOk

`func (o *WorkbookChartAxis) GetMajorGridlinesOk() (*AnyOfmicrosoftGraphWorkbookChartGridlines, bool)`

GetMajorGridlinesOk returns a tuple with the MajorGridlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorGridlines

`func (o *WorkbookChartAxis) SetMajorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines)`

SetMajorGridlines sets MajorGridlines field to given value.

### HasMajorGridlines

`func (o *WorkbookChartAxis) HasMajorGridlines() bool`

HasMajorGridlines returns a boolean if a field has been set.

### SetMajorGridlinesNil

`func (o *WorkbookChartAxis) SetMajorGridlinesNil(b bool)`

 SetMajorGridlinesNil sets the value for MajorGridlines to be an explicit nil

### UnsetMajorGridlines
`func (o *WorkbookChartAxis) UnsetMajorGridlines()`

UnsetMajorGridlines ensures that no value is present for MajorGridlines, not even an explicit nil
### GetMinorGridlines

`func (o *WorkbookChartAxis) GetMinorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines`

GetMinorGridlines returns the MinorGridlines field if non-nil, zero value otherwise.

### GetMinorGridlinesOk

`func (o *WorkbookChartAxis) GetMinorGridlinesOk() (*AnyOfmicrosoftGraphWorkbookChartGridlines, bool)`

GetMinorGridlinesOk returns a tuple with the MinorGridlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorGridlines

`func (o *WorkbookChartAxis) SetMinorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines)`

SetMinorGridlines sets MinorGridlines field to given value.

### HasMinorGridlines

`func (o *WorkbookChartAxis) HasMinorGridlines() bool`

HasMinorGridlines returns a boolean if a field has been set.

### SetMinorGridlinesNil

`func (o *WorkbookChartAxis) SetMinorGridlinesNil(b bool)`

 SetMinorGridlinesNil sets the value for MinorGridlines to be an explicit nil

### UnsetMinorGridlines
`func (o *WorkbookChartAxis) UnsetMinorGridlines()`

UnsetMinorGridlines ensures that no value is present for MinorGridlines, not even an explicit nil
### GetTitle

`func (o *WorkbookChartAxis) GetTitle() AnyOfmicrosoftGraphWorkbookChartAxisTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkbookChartAxis) GetTitleOk() (*AnyOfmicrosoftGraphWorkbookChartAxisTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkbookChartAxis) SetTitle(v AnyOfmicrosoftGraphWorkbookChartAxisTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkbookChartAxis) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkbookChartAxis) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkbookChartAxis) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


