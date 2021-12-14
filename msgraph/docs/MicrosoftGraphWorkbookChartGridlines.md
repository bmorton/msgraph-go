# MicrosoftGraphWorkbookChartGridlines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Visible** | Pointer to **bool** | Boolean value representing if the axis gridlines are visible or not. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlinesFormat**](anyOf&lt;microsoft.graph.workbookChartGridlinesFormat&gt;.md) | Represents the formatting of chart gridlines. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartGridlines

`func NewMicrosoftGraphWorkbookChartGridlines() *MicrosoftGraphWorkbookChartGridlines`

NewMicrosoftGraphWorkbookChartGridlines instantiates a new MicrosoftGraphWorkbookChartGridlines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartGridlinesWithDefaults

`func NewMicrosoftGraphWorkbookChartGridlinesWithDefaults() *MicrosoftGraphWorkbookChartGridlines`

NewMicrosoftGraphWorkbookChartGridlinesWithDefaults instantiates a new MicrosoftGraphWorkbookChartGridlines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartGridlines) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartGridlines) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartGridlines) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartGridlines) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVisible

`func (o *MicrosoftGraphWorkbookChartGridlines) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *MicrosoftGraphWorkbookChartGridlines) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *MicrosoftGraphWorkbookChartGridlines) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *MicrosoftGraphWorkbookChartGridlines) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetFormat

`func (o *MicrosoftGraphWorkbookChartGridlines) GetFormat() AnyOfmicrosoftGraphWorkbookChartGridlinesFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartGridlines) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartGridlinesFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartGridlines) SetFormat(v AnyOfmicrosoftGraphWorkbookChartGridlinesFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartGridlines) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartGridlines) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartGridlines) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


