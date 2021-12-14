# MicrosoftGraphWorkbookChartLegend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Overlay** | Pointer to **NullableBool** | Boolean value for whether the chart legend should overlap with the main body of the chart. | [optional] 
**Position** | Pointer to **NullableString** | Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom. | [optional] 
**Visible** | Pointer to **bool** | A boolean value the represents the visibility of a ChartLegend object. | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLegendFormat**](anyOf&lt;microsoft.graph.workbookChartLegendFormat&gt;.md) | Represents the formatting of a chart legend, which includes fill and font formatting. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookChartLegend

`func NewMicrosoftGraphWorkbookChartLegend() *MicrosoftGraphWorkbookChartLegend`

NewMicrosoftGraphWorkbookChartLegend instantiates a new MicrosoftGraphWorkbookChartLegend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookChartLegendWithDefaults

`func NewMicrosoftGraphWorkbookChartLegendWithDefaults() *MicrosoftGraphWorkbookChartLegend`

NewMicrosoftGraphWorkbookChartLegendWithDefaults instantiates a new MicrosoftGraphWorkbookChartLegend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookChartLegend) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartLegend) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartLegend) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookChartLegend) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOverlay

`func (o *MicrosoftGraphWorkbookChartLegend) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *MicrosoftGraphWorkbookChartLegend) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *MicrosoftGraphWorkbookChartLegend) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *MicrosoftGraphWorkbookChartLegend) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### SetOverlayNil

`func (o *MicrosoftGraphWorkbookChartLegend) SetOverlayNil(b bool)`

 SetOverlayNil sets the value for Overlay to be an explicit nil

### UnsetOverlay
`func (o *MicrosoftGraphWorkbookChartLegend) UnsetOverlay()`

UnsetOverlay ensures that no value is present for Overlay, not even an explicit nil
### GetPosition

`func (o *MicrosoftGraphWorkbookChartLegend) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MicrosoftGraphWorkbookChartLegend) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MicrosoftGraphWorkbookChartLegend) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MicrosoftGraphWorkbookChartLegend) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *MicrosoftGraphWorkbookChartLegend) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *MicrosoftGraphWorkbookChartLegend) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetVisible

`func (o *MicrosoftGraphWorkbookChartLegend) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *MicrosoftGraphWorkbookChartLegend) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *MicrosoftGraphWorkbookChartLegend) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *MicrosoftGraphWorkbookChartLegend) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetFormat

`func (o *MicrosoftGraphWorkbookChartLegend) GetFormat() AnyOfmicrosoftGraphWorkbookChartLegendFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartLegend) GetFormatOk() (*AnyOfmicrosoftGraphWorkbookChartLegendFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartLegend) SetFormat(v AnyOfmicrosoftGraphWorkbookChartLegendFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartLegend) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphWorkbookChartLegend) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphWorkbookChartLegend) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


